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

package config

import "math"

var (
	// Auth
	UsernameMinLen = 1
	UsernameMaxLen = 20
	PasswordMinLen = 8
	PasswordMaxLen = 80

	// Server
	ServerNameMinLen     = 1
	ServerNameMaxLen     = 32
	ServerGameMinLen     = 1
	ServerGameMaxLen     = 32
	ServerPasswordMinLen = 1
	ServerPasswordMaxLen = 64

	// Infractions
	InfractionReasonMinLen       = 1
	InfractionReasonMaxLen       = 4096
	InfractionDurationMax        = math.MaxInt32
	RecentInfractionsReturnCount = 20

	// Search
	SearchTermMinLen = 1
	SearchTermMaxLen = 64
	SearchOffsetMin  = 0
	SearchOffsetMax  = 2147483647 // max int32 value
	SearchLimitMin   = 1
	SearchLimitMax   = 100

	// Players
	RecentPlayersMaxSize = 22
)
