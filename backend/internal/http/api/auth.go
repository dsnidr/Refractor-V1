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
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"time"
)

type authHandler struct {
	service refractor.AuthService
	secure  bool
}

func NewAuthHandler(authService refractor.AuthService, secure bool) refractor.AuthHandler {
	return &authHandler{
		service: authService,
		secure:  secure,
	}
}

func (h *authHandler) LogInUser(c echo.Context) error {
	body := params.LoginParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	tokenPair, res := h.service.LogInUser(body)
	if !res.Success {
		return c.JSON(res.StatusCode, Response{
			Success: res.Success,
			Message: res.Message,
		})
	}

	// Create an HTTPOnly cookie for the refresh token
	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    tokenPair.RefreshToken,
		Expires:  time.Now().Add(14 * (24 * time.Hour)),
		HttpOnly: true,
	}

	// If deployed to a secure environment (https) set SameSite and Secure cookie properties
	if h.secure {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Secure = true
	}

	c.SetCookie(&cookie)

	// Send back auth token to the user
	return c.JSON(http.StatusOK, Response{
		Success: true,
		Payload: tokenPair.AuthToken,
	})
}

func (h *authHandler) RefreshUser(c echo.Context) error {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: config.MessageUnableRefreshCreds,
		})
	}

	tokenPair, res := h.service.RefreshUser(cookie.Value)
	if !res.Success {
		return c.JSON(res.StatusCode, Response{
			Success: res.Success,
			Message: res.Message,
		})
	}

	// Create an HTTPOnly cookie for the refresh token
	newCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    tokenPair.RefreshToken,
		Expires:  time.Now().Add(14 * (24 * time.Hour)),
		HttpOnly: true,
	}

	// If deployed to a secure environment (https) set SameSite and Secure cookie properties
	if h.secure {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Secure = true
	}

	c.SetCookie(&newCookie)

	// Send back auth token to the user
	return c.JSON(http.StatusOK, Response{
		Success: true,
		Payload: tokenPair.AuthToken,
	})
}

func (h *authHandler) CheckAuth(c echo.Context) error {
	// Auth is checked in middleware, so if it reaches this point auth is valid
	return c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Authenticated",
	})
}
