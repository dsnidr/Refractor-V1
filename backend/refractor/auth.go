package refractor

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
)

type AuthHandler interface {
	LogInUser(c echo.Context) error
	RefreshUser(c echo.Context) error
	CheckAuth(c echo.Context) error
}

type AuthService interface {
	LogInUser(body params.LoginParams) (*TokenPair, *ServiceResponse)
	RefreshUser(refreshToken string) (*TokenPair, *ServiceResponse)
}
