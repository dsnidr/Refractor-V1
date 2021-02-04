package api

import (
	"github.com/labstack/echo/v4"
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
	panic("implement me")
}
