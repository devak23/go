package main

import "github.com/gorilla/websocket"

type client struct {
	socket *websocket.Conn // create socket for this client
	send   chan []byte     // the channel on which messages are sent
	room   *room
}

type room struct {
	forward chan []byte // holds the incoming message that should be
	// forwarded to other clients
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

// define a read operation that reads the data from the socket
// and forwards it on the channel in the room
func (client *client) read() {
	defer client.socket.Close()
	for {
		_, msg, err := client.socket.ReadMessage()
		if err != nil {
			return
		}

		client.room.forward <- msg
	}
}

// define a write operation for the client which continuously accepts messages
// from the sent channel and writes it to the socket
func (client *client) write() {
	defer client.socket.Close()
	for msg := range client.send {
		err := client.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
