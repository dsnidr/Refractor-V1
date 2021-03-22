package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/pkg/broadcast"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"strconv"
)

type playerHandler struct {
	service refractor.PlayerService
}

func NewPlayerHandler(service refractor.PlayerService) refractor.PlayerHandler {
	return &playerHandler{
		service: service,
	}
}

func (h *playerHandler) GetRecentPlayers(c echo.Context) error {
	recentPlayers, res := h.service.GetRecentPlayers()
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: recentPlayers,
	})
}

func (h *playerHandler) SwitchPlayerWatch(watch bool) echo.HandlerFunc {
	return func(c echo.Context) error {
		idString := c.Param("id")

		playerID, err := strconv.ParseInt(idString, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: config.MessageInvalidIDProvided,
			})
		}

		res := h.service.SetPlayerWatch(playerID, watch)
		return c.JSON(res.StatusCode, Response{
			Success: res.Success,
			Message: res.Message,
		})
	}
}

func (h *playerHandler) OnPlayerJoin(fields broadcast.Fields, serverID int64, gameConfig *refractor.GameConfig) {
	h.service.OnPlayerJoin(serverID, fields[gameConfig.PlayerGameIDField], fields["Name"], gameConfig)
}

func (h *playerHandler) OnPlayerQuit(fields broadcast.Fields, serverID int64, gameConfig *refractor.GameConfig) {
	h.service.OnPlayerQuit(serverID, fields[gameConfig.PlayerGameIDField], gameConfig)
}
