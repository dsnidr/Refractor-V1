package refractor

import (
	"github.com/sniddunc/refractor/pkg/broadcast"
	"net"
)

type WebsocketMessage struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

type ChatSendSubscriber func(msgBody *ChatSendBody)

type ChatSendBody struct {
	ServerID   int64  `json:"serverId"`
	Message    string `json:"message"`
	Sender     string `json:"sender"`

	// SentByUser is true if this message was sent by another Refractor user
	SentByUser bool   `json:"sendByUser"`
}

type WebsocketService interface {
	Broadcast(message *WebsocketMessage)
	CreateClient(userID int64, conn net.Conn)
	StartPool()
	OnPlayerJoin(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
	OnPlayerQuit(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
	OnServerOnline(serverID int64)
	OnServerOffline(serverID int64)
	SubscribeChatSend(subscriber ChatSendSubscriber)
}
