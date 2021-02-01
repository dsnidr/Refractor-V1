package validation

import "regexp"

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const (
	emailMinLen = 3
	emailMaxLen = 254
)

// IsEmailValid takes in an email string and returns true if it's valid, false if it isn't.
func IsEmailValid(email string) bool {
	if len(email) < emailMinLen || len(email) > emailMaxLen {
		return false
	}

	return emailRegex.MatchString(email)
}
