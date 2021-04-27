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
