package websocket

import "net"

type Client struct {
	ID     int64
	UserID int64
	Conn   net.Conn
	Pool   *Pool
}

var nextClientID int64 = 1

func NewClient(userID int64, conn net.Conn, pool *Pool) *Client {
	client := &Client{
		ID:     nextClientID,
		UserID: userID,
		Conn:   conn,
		Pool:   pool,
	}

	nextClientID++

	return client
}
