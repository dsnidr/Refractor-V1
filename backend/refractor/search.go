package refractor

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
)

type SearchService interface {
	SearchPlayers(body params.SearchPlayersParams) (int, []*Player, *ServiceResponse)
	SearchInfractions(body params.SearchInfractionsParams) (int, []*Infraction, *ServiceResponse)
}

type SearchHandler interface {
	SearchPlayers(c echo.Context) error
}
