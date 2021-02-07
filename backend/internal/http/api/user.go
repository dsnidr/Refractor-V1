package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/pkg/jwt"
	"github.com/sniddunc/refractor/refractor"
)

type userHandler struct {
	service refractor.UserService
}

func NewUserHandler(userService refractor.UserService) refractor.UserHandler {
	return &userHandler{
		service: userService,
	}
}

func (h *userHandler) GetOwnUserInfo(c echo.Context) error {
	claims := c.Get("claims").(*jwt.Claims)

	userInfo, res := h.service.GetUserInfo(claims.UserID)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: userInfo,
	})
}
