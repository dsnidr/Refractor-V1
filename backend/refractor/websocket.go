package refractor

import "net"

type WebsocketMessage struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

type WebsocketService interface {
	Broadcast(message *WebsocketMessage)
	CreateClient(userID int64, conn net.Conn)
	StartPool()
}
