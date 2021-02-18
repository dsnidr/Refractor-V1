package websocket

import "github.com/sniddunc/refractor/refractor"

type Pool struct {
	Clients    map[int64]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *refractor.WebsocketMessage
}

func NewPool() *Pool {
	return &Pool{
		Clients:    map[int64]*Client{},
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *refractor.WebsocketMessage),
	}
}
