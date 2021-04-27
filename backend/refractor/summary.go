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

import "github.com/labstack/echo/v4"

type PlayerSummary struct {
	Warnings []*Infraction `json:"warnings"`
	Mutes    []*Infraction `json:"mutes"`
	Kicks    []*Infraction `json:"kicks"`
	Bans     []*Infraction `json:"bans"`
	*Player
}

type SummaryService interface {
	GetPlayerSummary(id int64) (*PlayerSummary, *ServiceResponse)
}

type SummaryHandler interface {
	GetPlayerSummary(c echo.Context) error
}
