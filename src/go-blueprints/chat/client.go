package main

import "github.com/gorilla/websocket"

type client struct {
	socket *websocket.Conn	// create socket for this client
	send chan []byte 		// the channels on which messages are sent
	room *room
}

type room struct {
	forward chan []byte 	// holds the incoming message that should be
							// forwarded to other clients
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}

		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}