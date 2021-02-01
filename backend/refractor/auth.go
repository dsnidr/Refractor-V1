package refractor

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/jwt"
)

type AuthHandler interface {
	LogInUser(c echo.Context) error
	RefreshUser(c echo.Context) error
	CheckAuth(c echo.Context) error
}

type AuthService interface {
	LogInUser(body params.LoginParams) (*jwt.TokenPair, *ServiceResponse)
	RefreshUser(refreshToken string) *ServiceResponse
}
