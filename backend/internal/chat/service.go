package chat

import (
	"fmt"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
)

type chatService struct {
	log              log.Logger
	websocketService refractor.WebsocketService
	rconService      refractor.RCONService
}

func NewChatService(websocketService refractor.WebsocketService, rconService refractor.RCONService, log log.Logger) refractor.ChatService {
	return &chatService{
		websocketService: websocketService,
		rconService:      rconService,
		log:              log,
	}
}

func (s *chatService) OnChatReceive(message *refractor.ChatReceiveBody, serverID int64, gameConfig *refractor.GameConfig) {
	s.websocketService.Broadcast(&refractor.WebsocketMessage{
		Type: "chat",
		Body: message,
	})
}

func (s *chatService) OnUserSendChat(msgBody *refractor.ChatSendBody) {
	if !msgBody.SentByUser {
		return
	}

	fmt.Println("A chat message was sent", msgBody.SentByUser)

	s.websocketService.Broadcast(&refractor.WebsocketMessage{
		Type: "chat",
		Body: &refractor.ChatReceiveBody{
			ServerID:     msgBody.ServerID,
			PlayerGameID: "",
			Name:         msgBody.Sender,
			Message:      msgBody.Message,
			SentByUser:   msgBody.SentByUser,
		},
	})
}
