/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package chat

import (
	"fmt"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
)

type chatService struct {
	repo             refractor.ChatRepository
	playerRepo       refractor.PlayerRepository
	websocketService refractor.WebsocketService
	rconService      refractor.RCONService
	log              log.Logger
}

func NewChatService(chatRepo refractor.ChatRepository, playerRepo refractor.PlayerRepository,
	websocketService refractor.WebsocketService, rconService refractor.RCONService, log log.Logger) refractor.ChatService {
	return &chatService{
		repo:             chatRepo,
		playerRepo:       playerRepo,
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

	player, err := s.playerRepo.FindOne(refractor.FindArgs{
		gameConfig.PlayerGameIDField: message.PlayerGameID,
	})
	if err != nil {
		if err == refractor.ErrNotFound {
			s.log.Warn("OnChatReceive player not find by %s = %s", gameConfig.PlayerGameIDField, message.PlayerGameID)
			return
		}

		s.log.Error("Could not find player by %s. Error: %v", gameConfig.PlayerGameIDField, err)
		return
	}

	// Log chat message
	_, _ = s.LogMessage(&refractor.ChatMessage{
		PlayerID: player.PlayerID,
		ServerID: serverID,
		Message:  message.Message,
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

func (s *chatService) LogMessage(message *refractor.ChatMessage) (*refractor.ChatMessage, *refractor.ServiceResponse) {
	newMessage, err := s.repo.Create(message)
	if err != nil {
		s.log.Error("Could not create chat message. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	// Get player name
	currentName, _, err := s.playerRepo.GetPlayerNames(newMessage.PlayerID)
	if err != nil {
		s.log.Error("Could not get player's current name. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	newMessage.PlayerName = currentName

	return newMessage, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Chat message logged",
	}
}
