package refractor

import (
	rcon "github.com/sniddunc/mordhau-rcon"
	"github.com/sniddunc/refractor/pkg/broadcast"
)

// RCONClient wraps around a mordhau-rcon Client and has an extra field containing the server
type RCONClient struct {
	Server *Server
	*rcon.Client
}

type BroadcastSubscriber func(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
type StatusSubscriber func(serverID int64)

type RCONService interface {
	CreateClient(*Server) error
	GetClients() map[int64]*RCONClient
	DeleteClient(serverID int64)
	SubscribeJoin(subscriber BroadcastSubscriber)
	SubscribeQuit(subscriber BroadcastSubscriber)
	SubscribeOnline(subscriber StatusSubscriber)
	SubscribeOffline(subscriber StatusSubscriber)
	SubscribeChat(subscriber BroadcastSubscriber)
	SendChatMessage(msgBody *ChatSendBody)
}
