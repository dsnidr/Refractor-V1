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
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/refractor"
)

type searchHandler struct {
	service refractor.SearchService
}

func NewSearchHandler(service refractor.SearchService) refractor.SearchHandler {
	return &searchHandler{
		service: service,
	}
}

type playerResultPayload struct {
	Results []*refractor.Player `json:"results"`
	Count   int                 `json:"count"`
}

func (h *searchHandler) SearchPlayers(c echo.Context) error {
	body := params.SearchPlayersParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	count, players, res := h.service.SearchPlayers(body)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Errors:  res.ValidationErrors,
		Payload: playerResultPayload{
			Results: players,
			Count:   count,
		},
	})
}

type infractionResultPayload struct {
	Results []*refractor.Infraction `json:"results"`
	Count   int                     `json:"count"`
}

func (h *searchHandler) SearchInfractions(c echo.Context) error {
	body := params.SearchInfractionsParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	count, infractions, res := h.service.SearchInfractions(body)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Errors:  res.ValidationErrors,
		Payload: infractionResultPayload{
			Results: infractions,
			Count:   count,
		},
	})
}
