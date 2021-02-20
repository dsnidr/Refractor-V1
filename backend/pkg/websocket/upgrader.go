package websocket

import (
	"github.com/gobwas/ws"
	"net"
	"net/http"
)

func Upgrade(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
