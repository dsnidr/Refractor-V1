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
	"github.com/sniddunc/bitperms"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/perms"
	"github.com/sniddunc/refractor/pkg/validation"
	"net/url"
	"strconv"
)

// UserMeta is a struct intended to be attached to other param structs. Its purpose is to provide user metadata
// on the user who is sending a request.
type UserMeta struct {
	UserID      int64
	Permissions int64
}

// CreateUserParams holds the data we expect when creating a new user.
type CreateUserParams struct {
	Email           string `json:"email" form:"email"`
	Username        string `json:"username" form:"username"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"passwordConfirm" form:"passwordConfirm"`
}

// Validate validates the data inside the attached struct
func (body *CreateUserParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	if !validation.IsEmailValid(body.Email) {
		errors.Set("email", "Invalid email address")
	}

	if len(body.Username) < config.UsernameMinLen || len(body.Username) > config.UsernameMaxLen {
		errors.Set("username", fmt.Sprintf("Username must be between %d and %d characters in length",
			config.UsernameMinLen, config.UsernameMaxLen))
	}

	if len(body.Password) < config.PasswordMinLen || len(body.Password) > config.PasswordMaxLen {
		errors.Set("password", fmt.Sprintf("Password must be between %d and %d characters in length",
			config.PasswordMinLen, config.PasswordMaxLen))
	}

	if body.Password != body.PasswordConfirm {
		errors.Set("passwordConfirm", "Passwords don't match")
	}

	return len(errors) == 0, errors
}

// SetUserPermissionsParams holds the data we expect when setting a user's permissions.
type SetUserPermissionsParams struct {
	UserID           int64 `json:"id" form:"id"`
	Permissions      int64
	PermissionString string `json:"permissions" form:"permissions"`
	*UserMeta
}

// Validate validates the data inside the attached struct
func (body *SetUserPermissionsParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	permSigned, err := strconv.ParseUint(body.PermissionString, 10, 64)
	if err != nil {
		errors.Set("permissions", "Invalid permissions value. Must be a string representing a uint64")
	} else {
		body.Permissions = int64(permSigned)

		newPerms := bitperms.PermissionValue(body.Permissions)
		if newPerms.HasFlag(perms.SUPER_ADMIN) {
			errors.Set("permissions", "For security reasons, you cannot make a user a super admin")
		}
	}

	if body.UserID < 1 {
		errors.Set("id", "Invalid user ID provided")
	}

	return len(errors) == 0, errors
}

// ChangeUserPassword holds the data we expect when changing a user's password
type ChangeUserPassword struct {
	CurrentPassword    string `json:"currentPassword" form:"currentPassword"`
	NewPassword        string `json:"newPassword" form:"newPassword"`
	NewPasswordConfirm string `json:"newPasswordConfirm" form:"newPasswordConfirm"`
}

// Validate validates the data inside the attached struct
func (body *ChangeUserPassword) Validate() (bool, url.Values) {
	errors := url.Values{}

	if len(body.NewPassword) < config.PasswordMinLen || len(body.NewPassword) > config.PasswordMaxLen {
		errors.Set("newPassword", fmt.Sprintf("Password must be between %d and %d characters in length",
			config.PasswordMinLen, config.PasswordMaxLen))
	}

	if body.NewPassword != body.NewPasswordConfirm {
		errors.Set("newPasswordConfirm", "Passwords don't match")
	}

	return len(errors) == 0, errors
}

// SetUserPasswordParams holds the data we expect when setting a user's password to a new value
type SetUserPasswordParams struct {
	UserID      int64  `json:"id" form:"id"`
	NewPassword string `json:"newPassword" form:"newPassword"`
	*UserMeta
}

// Validate validates the data inside the attached struct
func (body *SetUserPasswordParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	if body.UserID < 1 {
		errors.Set("id", "Invalid user ID provided")
	}

	if len(body.NewPassword) < config.PasswordMinLen || len(body.NewPassword) > config.PasswordMaxLen {
		errors.Set("newPassword", fmt.Sprintf("Password must be between %d and %d characters in length",
			config.PasswordMinLen, config.PasswordMaxLen))
	}

	return len(errors) == 0, errors
}
