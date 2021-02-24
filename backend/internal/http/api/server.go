package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
)

type serverHandler struct {
	service refractor.ServerService
	log     log.Logger
}

func NewServerHandler(service refractor.ServerService, log log.Logger) refractor.ServerHandler {
	return &serverHandler{
		service: service,
		log:     log,
	}
}

func (h *serverHandler) CreateServer(c echo.Context) error {
	body := params.CreateServerParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	server, res := h.service.CreateServer(body)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: server,
		Errors:  res.ValidationErrors,
	})
}

func (h *serverHandler) GetAllServers(c echo.Context) error {
	allServers, res := h.service.GetAllServers()
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: allServers,
	})
}

func (h *serverHandler) GetAllServerData(c echo.Context) error {
	allServerData, res := h.service.GetAllServerData()
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: allServerData,
	})
}