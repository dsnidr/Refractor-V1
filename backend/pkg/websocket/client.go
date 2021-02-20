package websocket

import (
	"encoding/json"
	"github.com/gobwas/ws/wsutil"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"io"
	"net"
)

type Client struct {
	ID     int64
	UserID int64
	Conn   net.Conn
	Pool   *Pool
	log    log.Logger
}

var nextClientID int64 = 1

func NewClient(userID int64, conn net.Conn, pool *Pool, log log.Logger) *Client {
	client := &Client{
		ID:     nextClientID,
		UserID: userID,
		Conn:   conn,
		Pool:   pool,
		log:    log,
	}

	nextClientID++

	return client
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

			c.log.Warn("Could not read message from client %d. Error: %v\n", c.UserID, err)
			continue
		}

		var msg *refractor.WebsocketMessage
		if err = json.Unmarshal(msgBytes, &msg); err != nil {
			c.log.Error("Could not unmarshal message: %v Error: %v\n", string(msgBytes), err)
			continue
		}

		c.log.Info("Message received from client ID %d: %v\n", c.ID, msg)
	}
}
