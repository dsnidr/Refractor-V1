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
)

// CreateWarningParams holds the data we expect when creating a new warning
type CreateWarningParams struct {
	PlayerID int64  `json:"playerId" form:"playerId"`
	ServerID int64  `json:"serverId" form:"serverId"`
	Reason   string `json:"reason" form:"reason"`
}

func (body *CreateWarningParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	if body.PlayerID < 1 {
		errors.Set("playerId", "Invalid player ID")
	}

	if body.ServerID < 1 {
		errors.Set("serverId", "Invalid server ID")
	}

	if body.Reason == "" {
		errors.Set("reason", "Reason is a required field")
	} else if len(body.Reason) < config.InfractionReasonMinLen || len(body.Reason) > config.InfractionReasonMaxLen {
		errors.Set("reason", fmt.Sprintf("Reason must be between %d and %d characters in length", config.InfractionReasonMinLen, config.InfractionReasonMaxLen))
	}

	return len(errors) == 0, errors
}

// CreateMuteParams holds the data we expect when creating a new kick
type CreateMuteParams struct {
	PlayerID int64  `json:"playerId" form:"playerId"`
	ServerID int64  `json:"serverId" form:"serverId"`
	Reason   string `json:"reason" form:"reason"`
	Duration int    `json:"duration" form:"duration"`
}

func (body *CreateMuteParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	if body.PlayerID < 1 {
		errors.Set("playerId", "Invalid player ID")
	}

	if body.ServerID < 1 {
		errors.Set("serverId", "Invalid server ID")
	}

	if body.Reason == "" {
		errors.Set("reason", "Reason is a required field")
	} else if len(body.Reason) < config.InfractionReasonMinLen || len(body.Reason) > config.InfractionReasonMaxLen {
		errors.Set("reason", fmt.Sprintf("Reason must be between %d and %d characters in length", config.InfractionReasonMinLen, config.InfractionReasonMaxLen))
	}

	if body.Duration > config.InfractionDurationMax {
		errors.Set("duration", fmt.Sprintf("The maximum duration a mute can have is %d minutes", config.InfractionDurationMax))
	}

	if body.Duration < 0 {
		errors.Set("duration", "Invalid duration")
	}

	return len(errors) == 0, errors
}

// CreateKickParams holds the data we expect when creating a new kick
type CreateKickParams struct {
	PlayerID int64  `json:"playerId" form:"playerId"`
	ServerID int64  `json:"serverId" form:"serverId"`
	Reason   string `json:"reason" form:"reason"`
}

func (body *CreateKickParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	if body.PlayerID < 1 {
		errors.Set("playerId", "Invalid player ID")
	}

	if body.ServerID < 1 {
		errors.Set("serverId", "Invalid server ID")
	}

	if body.Reason == "" {
		errors.Set("reason", "Reason is a required field")
	} else if len(body.Reason) < config.InfractionReasonMinLen || len(body.Reason) > config.InfractionReasonMaxLen {
		errors.Set("reason", fmt.Sprintf("Reason must be between %d and %d characters in length", config.InfractionReasonMinLen, config.InfractionReasonMaxLen))
	}

	return len(errors) == 0, errors
}

// CreateBanParams holds the data we expect when creating a new ban
type CreateBanParams struct {
	PlayerID int64  `json:"playerId" form:"playerId"`
	ServerID int64  `json:"serverId" form:"serverId"`
	Reason   string `json:"reason" form:"reason"`
	Duration int    `json:"duration" form:"duration"`
}

func (body *CreateBanParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	if body.PlayerID < 1 {
		errors.Set("playerId", "Invalid player ID")
	}

	if body.ServerID < 1 {
		errors.Set("serverId", "Invalid server ID")
	}

	if body.Reason == "" {
		errors.Set("reason", "Reason is a required field")
	} else if len(body.Reason) < config.InfractionReasonMinLen || len(body.Reason) > config.InfractionReasonMaxLen {
		errors.Set("reason", fmt.Sprintf("Reason must be between %d and %d characters in length", config.InfractionReasonMinLen, config.InfractionReasonMaxLen))
	}

	if body.Duration > config.InfractionDurationMax {
		errors.Set("duration", fmt.Sprintf("The maximum duration a ban can have is %d minutes", config.InfractionDurationMax))
	}

	if body.Duration < 0 {
		errors.Set("duration", "Invalid duration")
	}

	return len(errors) == 0, errors
}

// UpdateInfractionParams holds the data we expect when updating an infraction
type UpdateInfractionParams struct {
	Reason   *string `json:"reason" form:"reason"`
	Duration *int    `json:"duration" form:"duration"`
	*UserMeta
}

func (body *UpdateInfractionParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	if body.Reason != nil {
		reason := *body.Reason

		if reason != "" && (len(reason) < config.InfractionReasonMinLen || len(reason) > config.InfractionReasonMaxLen) {
			errors.Set("reason", fmt.Sprintf("Reason must be between %d and %d characters in length", config.InfractionReasonMinLen, config.InfractionReasonMaxLen))
		}
	}

	if body.Duration != nil {
		duration := *body.Duration

		if duration > config.InfractionDurationMax {
			errors.Set("duration", fmt.Sprintf("The maximum duration a ban can have is %d minutes", config.InfractionDurationMax))
		}

		if duration < 0 || duration > config.InfractionDurationMax {
			errors.Set("duration", "Invalid duration")
		}
	}

	return len(errors) == 0, errors
}
