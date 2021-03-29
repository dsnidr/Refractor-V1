package refractor

import "github.com/sniddunc/refractor/pkg/broadcast"

type ChatService interface {
	OnChatReceive(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
}
