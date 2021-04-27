/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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

func (api *API) RequirePerms(flag int64) echo.MiddlewareFunc {
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

func (api *API) RequireOneOfPerms(flags ...int64) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := c.Get("claims").(*jwt.Claims)

			userPerms := bitperms.PermissionValue(claims.Permissions)

			if perms.UserHasFullAccess(userPerms) {
				return next(c)
			}

			for _, flag := range flags {
				if userPerms.HasFlag(flag) {
					return next(c)
				}
			}

			api.log.Warn("User ID %d tried to access an endpoint which they did not have permission to use: %s %s",
				claims.UserID, c.Request().Method, c.Request().URL.String())

			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
	}
}
