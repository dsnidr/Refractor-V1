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
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"strconv"
)

type summaryHandler struct {
	service refractor.SummaryService
}

func NewSummaryHandler(service refractor.SummaryService) refractor.SummaryHandler {
	return &summaryHandler{
		service: service,
	}
}

func (h *summaryHandler) GetPlayerSummary(c echo.Context) error {
	idString := c.Param("id")

	playerID, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: config.MessageInvalidIDProvided,
		})
	}

	summary, res := h.service.GetPlayerSummary(playerID)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: summary,
	})
}
