package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/jwt"
	"github.com/sniddunc/refractor/refractor"
)

type infractionHandler struct {
	service refractor.InfractionService
}

func NewInfractionHandler(service refractor.InfractionService) refractor.InfractionHandler {
	return &infractionHandler{
		service: service,
	}
}

func (h *infractionHandler) CreateWarning(c echo.Context) error {
	body := params.CreateWarningParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	claims := c.Get("claims").(*jwt.Claims)

	warning, res := h.service.CreateWarning(claims.UserID, body)
	return c.JSON(res.StatusCode, Response{
		Success: true,
		Message: res.Message,
		Payload: warning,
	})
}

func (h *infractionHandler) CreateMute(c echo.Context) error {
	body := params.CreateMuteParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	claims := c.Get("claims").(*jwt.Claims)

	mute, res := h.service.CreateMute(claims.UserID, body)
	return c.JSON(res.StatusCode, Response{
		Success: true,
		Message: res.Message,
		Payload: mute,
	})
}

func (h *infractionHandler) CreateKick(c echo.Context) error {
	body := params.CreateKickParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	claims := c.Get("claims").(*jwt.Claims)

	kick, res := h.service.CreateKick(claims.UserID, body)
	return c.JSON(res.StatusCode, Response{
		Success: true,
		Message: res.Message,
		Payload: kick,
	})
}

func (h *infractionHandler) CreateBan(c echo.Context) error {
	body := params.CreateBanParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	claims := c.Get("claims").(*jwt.Claims)

	ban, res := h.service.CreateBan(claims.UserID, body)
	return c.JSON(res.StatusCode, Response{
		Success: true,
		Message: res.Message,
		Payload: ban,
	})
}

func (h *infractionHandler) DeleteInfraction(c echo.Context) error {
	return nil
}
