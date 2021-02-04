package api

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/pkg/jwt"
)

// AttachClaims extracts the claims from a JWT and attaches them to the context to be used in handlers.
// AttachClaims must be called after the JWT middleware as the token is required to already be in the context.
func AttachClaims() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwtgo.Token)

			if claims, ok := token.Claims.(*jwt.Claims); ok && token.Valid {
				c.Set("claims", claims)
			}

			return next(c)
		}
	}
}
