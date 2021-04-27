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

package rcon

import (
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/refractor"
)

func (s *rconService) HandleJoinBroadcast(bcast *broadcast.Broadcast, serverID int64, gameConfig *refractor.GameConfig) {
	for _, sub := range s.joinSubscribers {
		sub(bcast.Fields, serverID, gameConfig)
	}
}

func (s *rconService) HandleQuitBroadcast(bcast *broadcast.Broadcast, serverID int64, gameConfig *refractor.GameConfig) {
	for _, sub := range s.quitSubscribers {
		sub(bcast.Fields, serverID, gameConfig)
	}
}

func (s *rconService) HandleChatBroadcast(bcast *broadcast.Broadcast, serverID int64, gameConfig *refractor.GameConfig) {
	fields := bcast.Fields

	for _, sub := range s.chatSubscribers {
		msgBody := &refractor.ChatReceiveBody{
			ServerID:     serverID,
			PlayerGameID: fields[gameConfig.PlayerGameIDField],
			Name:         fields["Name"],
			Message:      fields["Message"],
			SentByUser:   false,
		}

		sub(msgBody, serverID, gameConfig)
	}
}
