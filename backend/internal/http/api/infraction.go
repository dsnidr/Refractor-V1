package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/jwt"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"strconv"
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
		Success: res.Success,
		Message: res.Message,
		Errors:  res.ValidationErrors,
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
		Success: res.Success,
		Message: res.Message,
		Errors:  res.ValidationErrors,
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
		Success: res.Success,
		Message: res.Message,
		Errors:  res.ValidationErrors,
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
		Success: res.Success,
		Message: res.Message,
		Errors:  res.ValidationErrors,
		Payload: ban,
	})
}

func (h *infractionHandler) DeleteInfraction(c echo.Context) error {
	idString := c.Param("id")

	infractionID, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: config.MessageInvalidIDProvided,
		})
	}

	claims := c.Get("claims").(*jwt.Claims)

	res := h.service.DeleteInfraction(infractionID, params.UserMeta{
		UserID:      claims.UserID,
		Permissions: claims.Permissions,
	})

	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
	})
}

func (h *infractionHandler) UpdateInfraction(c echo.Context) error {
	idString := c.Param("id")

	infractionID, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: config.MessageInvalidIDProvided,
		})
	}

	// Validate request body
	body := params.UpdateInfractionParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	claims := c.Get("claims").(*jwt.Claims)

	body.UserMeta = &params.UserMeta{
		UserID:      claims.UserID,
		Permissions: claims.Permissions,
	}

	updatedInfraction, res := h.service.UpdateInfraction(infractionID, body)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Errors:  res.ValidationErrors,
		Payload: updatedInfraction,
	})
}

func (h *infractionHandler) GetPlayerInfractions(infractionType string) echo.HandlerFunc {
	return func(c echo.Context) error {
		idString := c.Param("id")

		playerID, err := strconv.ParseInt(idString, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: config.MessageInvalidIDProvided,
			})
		}

		infractions, res := h.service.GetPlayerInfractionsType(infractionType, playerID)
		return c.JSON(res.StatusCode, Response{
			Success: res.Success,
			Message: res.Message,
			Payload: infractions,
		})
	}
}

func (h *infractionHandler) GetRecentInfractions(c echo.Context) error {
	infractions, res := h.service.GetRecentInfractions(config.RecentInfractionsReturnCount)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: infractions,
	})
}
