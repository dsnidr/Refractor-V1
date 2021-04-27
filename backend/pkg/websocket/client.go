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
	"encoding/json"
	"github.com/gobwas/ws/wsutil"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"io"
	"net"
)

type ChatSendHandler func(msgBody *SendChatBody)

type Client struct {
	ID              int64
	UserID          int64
	Conn            net.Conn
	Pool            *Pool
	ChatSendHandler ChatSendHandler
	log             log.Logger
}

var nextClientID int64 = 1

func NewClient(userID int64, conn net.Conn, pool *Pool, log log.Logger, chatSendHandler ChatSendHandler) *Client {
	client := &Client{
		ID:              nextClientID,
		UserID:          userID,
		Conn:            conn,
		Pool:            pool,
		ChatSendHandler: chatSendHandler,
		log:             log,
	}

	nextClientID++

	return client
}

type SendChatBody struct {
	ServerID int64 `json:"serverId"`
	UserID   int64
	Message  string `json:"message"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		_ = c.Conn.Close()
	}()

	for {
		msgBytes, _, err := wsutil.ReadClientData(c.Conn)
		if err != nil {
			if err == io.EOF {
				return
			}

			c.log.Warn("Could not read message from client %d. Error: %v", c.UserID, err)
			continue
		}

		var msg *refractor.WebsocketMessage
		if err = json.Unmarshal(msgBytes, &msg); err != nil {
			c.log.Error("Could not unmarshal message: %v Error: %v", string(msgBytes), err)
			continue
		}

		if msg.Type == "ping" {
			// Send back a pong message
			reply := &refractor.WebsocketMessage{
				Type: "pong",
				Body: "",
			}

			msgBytes, err := json.Marshal(reply)
			if err != nil {
				c.log.Error("Could not marshal ping reply message to client ID %d. Error: %v", c.ID, err)
				continue
			}

			// Send pong message
			if err := wsutil.WriteServerText(c.Conn, msgBytes); err != nil {
				c.log.Error("Could not send ping reply message to client ID %d. Error: %v", c.ID, err)
				continue
			}

			// Continue as there's no need to log a ping message
			continue
		}

		if msg.Type == "chat" {
			data, err := json.Marshal(msg.Body)
			if err != nil {
				c.log.Error("Could not marshal chat message body (intermediary). Error: %v", err)
				continue
			}

			msgBody := &SendChatBody{}

			if err := json.Unmarshal(data, msgBody); err != nil {
				c.log.Error("Could not unmarshal chat message body (intermediary). Error: %v", err)
				continue
			}

			msgBody.UserID = c.UserID

			c.ChatSendHandler(msgBody)
		}

		c.log.Info("Message received from client ID %d: %v", c.ID, msg)
	}
}
