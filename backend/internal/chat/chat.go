package chat

import (
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
)

type chatService struct {
	log              log.Logger
	websocketService refractor.WebsocketService
}

func NewChatService(websocketService refractor.WebsocketService, log log.Logger) refractor.ChatService {
	return &chatService{
		websocketService: websocketService,
		log:              log,
	}
}

type chatMessage struct {
	ServerID     int64  `json:"serverId"`
	PlayerGameID string `json:"playerGameID"`
	Name         string `json:"name"`
	Message      string `json:"message"`
}

func (s *chatService) OnChatReceive(fields broadcast.Fields, serverID int64, gameConfig *refractor.GameConfig) {
	s.log.Info("Fields: %v", fields)

	name := fields["Name"]
	message := fields["Message"]
	playerGameID := fields[gameConfig.PlayerGameIDField]

	s.websocketService.Broadcast(&refractor.WebsocketMessage{
		Type: "chat",
		Body: chatMessage{
			ServerID:     serverID,
			PlayerGameID: playerGameID,
			Name:         name,
			Message:      message,
		},
	})
}
