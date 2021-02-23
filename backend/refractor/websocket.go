package refractor

import (
	"github.com/sniddunc/refractor/pkg/broadcast"
	"net"
)

type WebsocketMessage struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

type WebsocketService interface {
	Broadcast(message *WebsocketMessage)
	CreateClient(userID int64, conn net.Conn)
	StartPool()
	OnPlayerJoin(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
	OnPlayerQuit(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
	OnServerOnline(serverID int64)
	OnServerOffline(serverID int64)
}
