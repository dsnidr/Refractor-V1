package refractor

import (
	"backend/internal/params"
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	LogInUser(c echo.Context) error
	RefreshUser(c echo.Context) error
	CheckAuth(c echo.Context) error
}

type AuthService interface {
	LogInUser(body params.LoginParams) *ServiceResponse
	RefreshUser(refreshToken string) *ServiceResponse
}
