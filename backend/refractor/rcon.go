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
	rcon "github.com/sniddunc/mordhau-rcon"
	"github.com/sniddunc/refractor/pkg/broadcast"
)

// RCONClient wraps around a mordhau-rcon Client and has an extra field containing the server
type RCONClient struct {
	Server *Server
	*rcon.Client
}

type BroadcastSubscriber func(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
type ChatReceiveSubscriber func(msgBody *ChatReceiveBody, serverID int64, gameConfig *GameConfig)
type PlayerListPollSubscriber func(serverID int64, gameConfig *GameConfig, players []*Player)
type StatusSubscriber func(serverID int64)

type RCONService interface {
	CreateClient(*Server) error
	GetClients() map[int64]*RCONClient
	DeleteClient(serverID int64)
	SendChatMessage(msgBody *ChatSendBody)
	SubscribeJoin(subscriber BroadcastSubscriber)
	SubscribeQuit(subscriber BroadcastSubscriber)
	SubscribeOnline(subscriber StatusSubscriber)
	SubscribeOffline(subscriber StatusSubscriber)
	SubscribeChat(subscriber ChatReceiveSubscriber)
	SubscribePlayerListPoll(subscriber PlayerListPollSubscriber)
}
