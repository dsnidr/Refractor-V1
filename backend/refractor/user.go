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
	AccessLevel         int    `json:"accessLevel"`
	Activated           bool   `json:"activated"`
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
	GetUserInfo(id int64) (*UserInfo, *ServiceResponse)
	SetUserAccessLevel(body params.SetUserAccessLevelParams) (*User, *ServiceResponse)
}

type UserHandler interface {
	GetOwnUserInfo(c echo.Context) error
}

type UserInfo struct {
	ID                  int64  `json:"id"`
	Email               string `json:"email"`
	Username            string `json:"username"`
	Activated           bool   `json:"activated"`
	AccessLevel         int    `json:"accessLevel"`
	NeedsPasswordChange bool   `json:"needsPasswordChange"`
}
