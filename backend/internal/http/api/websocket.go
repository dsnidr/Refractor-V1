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
	"github.com/sniddunc/refractor/pkg/jwt"
	"github.com/sniddunc/refractor/pkg/websocket"
	"net/http"
	"os"
)

func (api *API) websocketHandler(c echo.Context) error {
	jwtString := c.QueryParam("auth")

	// Authenticate the user attempting this connection
	claims, err := jwt.ExtractAuthClaims(jwtString, os.Getenv("JWT_SECRET"))
	if err != nil {
		api.log.Warn("Failed to extract token claims for websocket auth. Error: %v", err)
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	conn, err := websocket.Upgrade(c.Response(), c.Request())
	if err != nil {
		api.log.Error("Could not upgrade WS request. Error: %v", err)
		return nil
	}

	// Leverage the websocket service to create a new client for us
	api.websocketService.CreateClient(claims.UserID, conn)

	return nil
}
