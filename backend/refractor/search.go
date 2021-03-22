package refractor

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
)

type SearchService interface {
	SearchPlayers(body params.SearchPlayersParams) (int, []*Player, *ServiceResponse)
}

type SearchHandler interface {
	SearchPlayers(c echo.Context) error
}
