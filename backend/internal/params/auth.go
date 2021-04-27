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

type LoginParams struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (body *LoginParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	if len(body.Username) < config.UsernameMinLen || len(body.Username) > config.UsernameMaxLen {
		errors.Set("username", fmt.Sprintf("Username must be between %d and %d characters in length",
			config.UsernameMinLen, config.UsernameMaxLen))
	}

	if len(body.Password) < config.PasswordMinLen || len(body.Password) > config.PasswordMaxLen {
		errors.Set("password", fmt.Sprintf("Password must be between %d and %d characters in length",
			config.PasswordMinLen, config.PasswordMaxLen))
	}

	return len(errors) == 0, errors
}
