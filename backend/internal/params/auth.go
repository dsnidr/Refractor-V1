package params

import (
	"backend/pkg/config"
	"fmt"
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
