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

package refractor

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
)

type SearchService interface {
	SearchPlayers(body params.SearchPlayersParams) (int, []*Player, *ServiceResponse)
	SearchInfractions(body params.SearchInfractionsParams) (int, []*Infraction, *ServiceResponse)
	SearchChatMessages(body params.SearchChatMessagesParams) (int, []*ChatMessage, *ServiceResponse)
}

type SearchHandler interface {
	SearchPlayers(c echo.Context) error
	SearchInfractions(c echo.Context) error
	SearchChatMessages(c echo.Context) error
}
