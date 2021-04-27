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
