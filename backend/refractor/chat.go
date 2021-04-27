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

type ChatReceiveBody struct {
	ServerID     int64  `json:"serverId"`
	PlayerGameID string `json:"playerGameID"`
	Name         string `json:"name"`
	Message      string `json:"message"`
	SentByUser   bool   `json:"sentByUser"`
}

type ChatService interface {
	OnChatReceive(msgBody *ChatReceiveBody, serverID int64, gameConfig *GameConfig)
	OnUserSendChat(msgBody *ChatSendBody)
}
