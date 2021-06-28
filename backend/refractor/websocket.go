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

package refractor

import (
	"github.com/sniddunc/refractor/pkg/broadcast"
	"net"
)

type WebsocketMessage struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

type WebsocketDirectMessage struct {
	ClientID int64
	Message  *WebsocketMessage
}

type ChatSendSubscriber func(msgBody *ChatSendBody)

type ChatSendBody struct {
	ServerID int64  `json:"serverId"`
	Message  string `json:"message"`
	Sender   string `json:"sender"`

	// SentByUser is true if this message was sent by another Refractor user
	SentByUser bool `json:"sendByUser"`
}

type WebsocketService interface {
	Broadcast(message *WebsocketMessage)
	CreateClient(userID int64, conn net.Conn)
	StartPool()
	OnPlayerJoin(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
	OnPlayerQuit(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
	OnServerOnline(serverID int64)
	OnServerOffline(serverID int64)
	OnInfractionCreate(infraction *Infraction)
	SubscribeChatSend(subscriber ChatSendSubscriber)
}
