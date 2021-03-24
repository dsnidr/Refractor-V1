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
		Payload: infractionResultPayload{
			Results: infractions,
			Count:   count,
		},
	})
}
