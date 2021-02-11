package refractor

import "github.com/labstack/echo/v4"

type GameServer struct {
	Name    string        `json:"name"`
	Servers []*ServerInfo `json:"servers"`
}

type GameServerService interface {
	GetGameServers() ([]*GameServer, *ServiceResponse)
}

type GameServerHandler interface {
	GetAllGameServers(c echo.Context) error
}
