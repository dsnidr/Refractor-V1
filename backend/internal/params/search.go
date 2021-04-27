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

package params

import (
	"fmt"
	"github.com/sniddunc/refractor/pkg/config"
	"net/url"
	"strconv"
	"strings"
)

type SearchParams struct {
	Offset int `json:"offset" form:"offset"`
	Limit  int `json:"limit" form:"limit"`
}

func (body *SearchParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	if body.Offset < config.SearchOffsetMin {
		errors.Set("offset", fmt.Sprintf("Offset value too small. Minimum: %d", config.SearchOffsetMin))
	} else if body.Offset > config.SearchOffsetMax {
		errors.Set("offset", fmt.Sprintf("Offset value too large. Maximum: %d", config.SearchOffsetMax))
	}

	if body.Limit < config.SearchLimitMin {
		errors.Set("limit", fmt.Sprintf("Limit value too small. Minimum: %d", config.SearchLimitMin))
	} else if body.Limit > config.SearchLimitMax {
		errors.Set("limit", fmt.Sprintf("Limit value too large. Maximum: %d", config.SearchLimitMax))
	}

	return len(errors) == 0, errors
}

type SearchPlayersParams struct {
	SearchTerm string `json:"term" form:"term"`
	SearchType string `json:"type" form:"type"`
	SearchParams
}

var validPlayerSearchTypes = []string{"name", "id", "playfabid", "mcuuid"}

func (body *SearchPlayersParams) Validate() (bool, url.Values) {
	if ok, errors := body.SearchParams.Validate(); !ok {
		return ok, errors
	}

	errors := url.Values{}

	body.SearchTerm = strings.TrimSpace(body.SearchTerm)
	body.SearchType = strings.TrimSpace(body.SearchType)

	if body.SearchTerm == "" {
		errors.Set("term", "You must provide a search term")
	}

	if len(body.SearchTerm) < config.SearchTermMinLen || len(body.SearchTerm) > config.SearchTermMaxLen {
		errors.Set("term", fmt.Sprintf("Search term must be between %d and %d characters in length", config.SearchTermMinLen, config.SearchTermMaxLen))
	}

	if body.SearchType == "" {
		errors.Set("type", "Please select a search type")
	} else {
		typeIsValid := false
		for _, validType := range validPlayerSearchTypes {
			if body.SearchType == validType {
				typeIsValid = true
				break
			}
		}

		if !typeIsValid {
			errors.Set("type", "Invalid search type. Valid types are: "+strings.Join(validPlayerSearchTypes, ", "))
		}
	}

	return len(errors) == 0, errors
}

type SearchInfractionsParams struct {
	Type     string `json:"type" form:"type"`
	PlayerID string `json:"playerId" form:"playerId"`
	UserID   string `json:"userId" form:"userId"`
	Game     string `json:"game" form:"game"`
	ServerID string `json:"serverId" form:"serverId"`
	*ParsedIDs
	SearchParams
}

type ParsedIDs struct {
	PlayerID int64
	UserID   int64
	ServerID int64
}

var validInfractionTypes = []string{"WARNING", "MUTE", "KICK", "BAN"}

func (body *SearchInfractionsParams) Validate() (bool, url.Values) {
	if ok, errors := body.SearchParams.Validate(); !ok {
		return ok, errors
	}
	body.ParsedIDs = &ParsedIDs{}

	errors := url.Values{}

	// Validate infraction type filter
	if body.Type != "" {
		valid := false
		for _, validType := range validInfractionTypes {
			if body.Type == validType {
				valid = true
				break
			}
		}

		if !valid {
			errors.Set("type", "Invalid type")
		}
	}

	// Validate and parse PlayerID
	if body.PlayerID != "" {
		playerID, err := strconv.ParseInt(body.PlayerID, 10, 64)
		if err != nil {
			errors.Set("playerId", config.MessageInvalidIDProvided)
		} else {
			body.ParsedIDs.PlayerID = playerID
		}
	}

	// Validate and parse UserID
	if body.UserID != "" {
		playerID, err := strconv.ParseInt(body.UserID, 10, 64)
		if err != nil {
			errors.Set("userId", config.MessageInvalidIDProvided)
		} else {
			body.ParsedIDs.UserID = playerID
		}
	}

	// Validate and parse ServerID
	if body.ServerID != "" {
		playerID, err := strconv.ParseInt(body.ServerID, 10, 64)
		if err != nil {
			errors.Set("serverId", config.MessageInvalidIDProvided)
		} else {
			body.ParsedIDs.ServerID = playerID
		}
	}

	// Validate game length (we don't check if the game exists since this is done at the service layer)
	if body.Game != "" {
		if len(body.Game) < config.ServerGameMinLen || len(body.Game) > config.ServerGameMaxLen {
			errors.Set("game", fmt.Sprintf("Game name must be between %d and %d characters in length",
				config.ServerGameMinLen, config.ServerGameMaxLen))
		}
	}

	return len(errors) == 0, errors
}
