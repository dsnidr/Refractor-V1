package api

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/bitperms"
	"github.com/sniddunc/refractor/pkg/jwt"
	"github.com/sniddunc/refractor/pkg/perms"
	"net/http"
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

// RequireAccessLevel is used to ensure a user has an access level of at least the required access level.
// RequireAccessLevel must be called after AttachClaims.
// Normally middleware would not be attached to the API struct, but in this case logging is important so
// binding it to the API struct greatly simplifies that.
//func (api *API) RequireAccessLevel(minAccessLevel int) echo.MiddlewareFunc {
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			claims := c.Get("claims").(*jwt.Claims)
//
//			if claims.AccessLevel < minAccessLevel {
//				api.log.Warn(`User of ID %d tried to access a route handler which requires a higher access level then
//								they possess. Have = %d Needs = %d`, claims.UserID, claims.AccessLevel, minAccessLevel)
//				return c.String(http.StatusUnauthorized, "Unauthorized")
//			}
//
//			return next(c)
//		}
//	}
//}

func (api *API) RequirePerms(flag uint64) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := c.Get("claims").(*jwt.Claims)

			userPerms := bitperms.PermissionValue(claims.Permissions)

			if !userPerms.HasFlag(flag) && !perms.UserHasFullAccess(userPerms) {
				api.log.Warn("User ID %d tried to access an endpoint which they did not have permission to use: %s %s",
					claims.UserID, c.Request().Method, c.Request().URL.String())

				return c.String(http.StatusUnauthorized, "Unauthorized")
			}

			return next(c)
		}
	}
}
