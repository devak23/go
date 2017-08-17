package main

import (
  "log"
  "github.com/gorilla/websocket"
  "net/http"
)

const (
  socketBufferSize = 1024
  messageBufferSize = 256
)

var upgrader = &websocket.Upgrader { ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize }

func (r *room) ServeHTTP (res http.ResponseWriter, req *http.Request) {
  socket, err := upgrader.Upgrade(res, req, nil)
  if err != nil {
    log.Fatal("ServeHTTP error:",err)
    return
  }

  client := &client {
    socket: socket,
    send: make(chan []byte, messageBufferSize),
    room: r,
  }

  r.join <- client
  defer func() { r.leave <- client }()
  go client.write()
  client.read()
}
