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

type User struct {
	UserID              int64  `json:"id"`
	Email               string `json:"email"`
	Username            string `json:"username"`
	Password            string `json:"-"`
	Permissions         int64  `json:"permissions"`
	Activated           bool   `json:"activated"`
	NeedsPasswordChange bool   `json:"needsPasswordChange"`
}

type UserInfo struct {
	ID                  int64  `json:"id"`
	Email               string `json:"email"`
	Username            string `json:"username"`
	Activated           bool   `json:"activated"`
	Permissions         int64  `json:"permissions"`
	NeedsPasswordChange bool   `json:"needsPasswordChange"`
}

type UserRepository interface {
	Create(user *User) error
	FindByID(id int64) (*User, error)
	Exists(args FindArgs) (bool, error)
	FindOne(args FindArgs) (*User, error)
	FindMany(args FindArgs) ([]*User, error)
	Update(id int64, args UpdateArgs) (*User, error)
	FindAll() ([]*User, error)
	GetCount() int
}

type UserService interface {
	CreateUser(body params.CreateUserParams) (*User, *ServiceResponse)
	GetUserByID(id int64) (*User, *ServiceResponse)
	GetUserInfo(id int64) (*UserInfo, *ServiceResponse)
	SetUserPermissions(body params.SetUserPermissionsParams) (*User, *ServiceResponse)
	ChangeUserPassword(id int64, body params.ChangeUserPassword) (*User, *ServiceResponse)
	ForceUserPasswordChange(id int64, userMeta *params.UserMeta) *ServiceResponse
	GetAllUsers() ([]*User, *ServiceResponse)
	UpdateUser(id int64, args UpdateArgs) (*User, *ServiceResponse)
	SetUserPassword(body params.SetUserPasswordParams) (*User, *ServiceResponse)
}

type UserHandler interface {
	GetOwnUserInfo(c echo.Context) error
	ChangeUserPassword(c echo.Context) error
	GetAllUsers(c echo.Context) error
	CreateUser(c echo.Context) error
	ActivateUser(c echo.Context) error
	DeactivateUser(c echo.Context) error
	ForcePasswordChange(c echo.Context) error
	SetUserPassword(c echo.Context) error
	SetUserPermissions(c echo.Context) error
}
