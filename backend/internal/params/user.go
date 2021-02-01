package params

import (
	"fmt"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/validation"
	"net/url"
)

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
