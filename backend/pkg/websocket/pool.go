package websocket

import (
	"encoding/json"
	"github.com/gobwas/ws/wsutil"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
)

type Pool struct {
	Clients    map[int64]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *refractor.WebsocketMessage
	log        log.Logger
}

func NewPool() *Pool {
	return &Pool{
		Clients:    map[int64]*Client{},
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *refractor.WebsocketMessage),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client.ID] = client

			pool.log.Info("Client connected: %v\n", client.Conn.RemoteAddr())
			break
		case client := <-pool.Unregister:
			pool.Clients[client.ID] = nil
			delete(pool.Clients, client.ID)

			pool.log.Info("Client disconnected: %v\n", client.Conn.RemoteAddr())
			break
		case msg := <-pool.Broadcast:
			msgBytes, err := json.Marshal(msg)
			if err != nil {
				pool.log.Error("Could not marshal broadcast message. Error: %v\n", err)
				continue
			}

			for _, client := range pool.Clients {
				if err := wsutil.WriteServerText(client.Conn, msgBytes); err != nil {
					pool.log.Warn("Could not send broadcast message to client ID %d. Error: %v\n", client.ID, err)
					continue
				}
			}

			break
		}
	}
}
