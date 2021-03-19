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
