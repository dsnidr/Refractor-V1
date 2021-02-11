package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/refractor"
)

type gameServerHandler struct {
	service refractor.GameServerService
}

func NewGameServerHandler(service refractor.GameServerService) refractor.GameServerHandler {
	return &gameServerHandler{
		service: service,
	}
}

func (h *gameServerHandler) GetAllGameServers(c echo.Context) error {
	gameServers, res := h.service.GetGameServers()
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: gameServers,
	})
}
