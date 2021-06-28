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

package websocket

import (
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/pkg/websocket"
	"github.com/sniddunc/refractor/refractor"
	"net"
)

type websocketService struct {
	pool                    *websocket.Pool
	userRepo                refractor.UserRepository
	playerRepo              refractor.PlayerRepository
	playerInfractionService refractor.PlayerInfractionService
	log                     log.Logger
	chatSendSubscribers     []refractor.ChatSendSubscriber
}

func NewWebsocketService(playerRepo refractor.PlayerRepository, userRepo refractor.UserRepository,
	playerInfractionService refractor.PlayerInfractionService, log log.Logger) refractor.WebsocketService {
	return &websocketService{
		pool:                    websocket.NewPool(log),
		playerRepo:              playerRepo,
		playerInfractionService: playerInfractionService,
		userRepo:                userRepo,
		log:                     log,
		chatSendSubscribers:     []refractor.ChatSendSubscriber{},
	}
}

func (s *websocketService) Broadcast(message *refractor.WebsocketMessage) {
	s.pool.Broadcast <- message
}

func (s *websocketService) CreateClient(userID int64, conn net.Conn) {
	client := websocket.NewClient(userID, conn, s.pool, s.log, s.sendChatHandler)

	s.pool.Register <- client
	client.Read()
}

func (s *websocketService) sendChatHandler(msgBody *websocket.SendChatBody) {
	// Get user's name
	user, err := s.userRepo.FindByID(msgBody.UserID)
	if err != nil {
		s.log.Error("Could not get user by ID %d. Error: %v", msgBody.UserID, err)
		return
	}

	transformed := &refractor.ChatSendBody{
		ServerID:   msgBody.ServerID,
		Message:    msgBody.Message,
		Sender:     user.Username,
		SentByUser: true,
	}

	for _, sub := range s.chatSendSubscribers {
		sub(transformed)
	}
}

func (s *websocketService) StartPool() {
	s.pool.Start()
}

type playerJoinQuitData struct {
	ServerID        int64  `json:"serverId"`
	PlayerID        int64  `json:"id"`
	PlayerGameID    string `json:"playerGameId"`
	Name            string `json:"name"`
	InfractionCount int    `json:"infractionCount,omitempty"`
	Watched         bool   `json:"watched"`
}

func (s *websocketService) OnPlayerJoin(fields broadcast.Fields, serverID int64, gameConfig *refractor.GameConfig) {
	idField := gameConfig.PlayerGameIDField

	player, err := s.playerRepo.FindOne(refractor.FindArgs{
		idField: fields[idField],
	})

	if err != nil {
		s.log.Warn("Could not GetPlayer. PlayerGameIDField = %s, field value = %v", idField, fields[idField])
		return
	}

	count, _ := s.playerInfractionService.GetPlayerInfractionCount(player.PlayerID)

	s.Broadcast(&refractor.WebsocketMessage{
		Type: "player-join",
		Body: playerJoinQuitData{
			ServerID:        serverID,
			PlayerID:        player.PlayerID,
			PlayerGameID:    fields[idField],
			Name:            player.CurrentName,
			InfractionCount: count,
			Watched:         player.Watched,
		},
	})
}

func (s *websocketService) OnPlayerQuit(fields broadcast.Fields, serverID int64, gameConfig *refractor.GameConfig) {
	idField := gameConfig.PlayerGameIDField

	player, err := s.playerRepo.FindOne(refractor.FindArgs{
		idField: fields[idField],
	})

	if err != nil {
		s.log.Warn("Could not GetPlayer. PlayerGameIDField = %s, field value = %v", idField, fields[idField])
		return
	}

	if player == nil {
		s.log.Warn("GetPlayer player returned was nil!")
		return
	}

	s.Broadcast(&refractor.WebsocketMessage{
		Type: "player-quit",
		Body: playerJoinQuitData{
			ServerID:     serverID,
			PlayerID:     player.PlayerID,
			PlayerGameID: fields[idField],
			Name:         player.CurrentName,
		},
	})
}

func (s *websocketService) OnServerOnline(serverID int64) {
	s.Broadcast(&refractor.WebsocketMessage{
		Type: "server-online",
		Body: serverID,
	})
}

func (s *websocketService) OnServerOffline(serverID int64) {
	s.Broadcast(&refractor.WebsocketMessage{
		Type: "server-offline",
		Body: serverID,
	})
}

type infractionCreateBody struct {
	InfractionID int64  `json:"id"`
	PlayerID     int64  `json:"playerId"`
	ServerID     int64  `json:"serverId"`
	Type         string `json:"type"`
	Reason       string `json:"reason"`
	Duration     int    `json:"duration"`
	SystemAction bool   `json:"systemAction"`
	StaffName    string `json:"staffName"`
	PlayerName   string `json:"playerName"`
}

func (s *websocketService) OnInfractionCreate(infraction *refractor.Infraction) {
	sendParams := &refractor.WebsocketDirectMessage{
		ClientID: 0,
		Message: &refractor.WebsocketMessage{
			Type: "infraction-create",
			Body: &infractionCreateBody{
				InfractionID: infraction.InfractionID,
				PlayerID:     infraction.PlayerID,
				ServerID:     infraction.ServerID,
				Type:         infraction.Type,
				Reason:       infraction.Reason,
				Duration:     infraction.Duration,
				SystemAction: infraction.SystemAction,
				StaffName:    infraction.StaffName,
				PlayerName:   infraction.PlayerName,
			},
		},
	}

	for clientID, client := range s.pool.Clients {
		if client.UserID == infraction.UserID {
			// Do not send to the user who created the infraction
			continue
		}

		sendParams.ClientID = clientID

		// Send direct
		s.pool.SendDirect <- sendParams
	}
}

func (s *websocketService) SubscribeChatSend(subscriber refractor.ChatSendSubscriber) {
	s.chatSendSubscribers = append(s.chatSendSubscribers, subscriber)
}
