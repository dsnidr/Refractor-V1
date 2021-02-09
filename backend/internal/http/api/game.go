package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/refractor"
)

type gameHandler struct {
	service refractor.GameService
}

func NewGameHandler(service refractor.GameService) refractor.GameHandler {
	return &gameHandler{
		service: service,
	}
}

func (h *gameHandler) GetAllGames(c echo.Context) error {
	allGames, res := h.service.GetAllGameInfo()
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: allGames,
	})
}
