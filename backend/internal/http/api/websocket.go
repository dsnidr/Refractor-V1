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
