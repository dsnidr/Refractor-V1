package params

import (
	"fmt"
	"github.com/sniddunc/refractor/pkg/config"
	"net"
	"net/url"
	"strconv"
)

// CreateServerParams holds the data we expect when creating a server
type CreateServerParams struct {
	Name         string `form:"name"`
	Game         string `form:"game"`
	Address      string `form:"address"`
	RCONPort     string `form:"rconPort"`
	RCONPassword string `form:"rconPassword"`
}

// Validate validates the data inside the attached struct
func (body *CreateServerParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	if len(body.Game) < config.ServerGameMinLen || len(body.Game) > config.ServerGameMaxLen {
		errors.Set("game", fmt.Sprintf("Server game name must be between %d and %d characters in length",
			config.ServerGameMinLen, config.ServerGameMaxLen))
	}

	if len(body.Name) < config.ServerNameMinLen || len(body.Name) > config.ServerNameMaxLen {
		errors.Set("name", fmt.Sprintf("Server name must be between %d and %d characters in length",
			config.ServerNameMinLen, config.ServerNameMaxLen))
	}

	if net.ParseIP(body.Address) == nil {
		errors.Set("address", "The provided server address was not a valid IP address")
	}

	// Since port numbers are 16 bit integers, we can check if the provided port is valid by
	// trying to parse it to an int16.
	if _, err := strconv.ParseUint(body.RCONPort, 10, 16); err != nil {
		errors.Set("rconPort", "The provided RCON port was not a valid port number")
	}

	if len(body.RCONPassword) < config.ServerPasswordMinLen || len(body.RCONPassword) > config.ServerPasswordMaxLen {
		errors.Set("rconPassword", fmt.Sprintf("RCON passwords must be between %d and %d characters in length",
			config.ServerPasswordMinLen, config.ServerPasswordMaxLen))
	}

	return len(errors) == 0, errors
}

type UpdateServerParams struct {
	Name         string `form:"name"`
	Address      string `form:"address"`
	RCONPort     string `form:"rconPort"`
	RCONPassword string `form:"rconPassword"`
}

func (body *UpdateServerParams) Validate() (bool, url.Values) {
	errors := url.Values{}

	if body.Name != "" {
		if len(body.Name) < config.ServerNameMinLen || len(body.Name) > config.ServerNameMaxLen {
			errors.Set("name", fmt.Sprintf("Server name must be between %d and %d characters in length",
				config.ServerNameMinLen, config.ServerNameMaxLen))
		}
	}

	if body.Address != "" {
		if net.ParseIP(body.Address) == nil {
			errors.Set("address", "The provided server address was not a valid IP address")
		}
	}

	if body.RCONPort != "" {
		// Since port numbers are 16 bit ints, we can check if the provided port is valid by
		// trying to parse it to an int16.
		if _, err := strconv.ParseUint(body.RCONPort, 10, 16); err != nil {
			errors.Set("rconPort", "The provided RCON port was not a valid port number")
		}
	}

	if body.RCONPassword != "" {
		if len(body.RCONPassword) < config.ServerPasswordMinLen || len(body.RCONPassword) > config.ServerPasswordMaxLen {
			errors.Set("rconPassword", fmt.Sprintf("RCON passwords must be between %d and %d characters in length",
				config.ServerPasswordMinLen, config.ServerPasswordMaxLen))
		}
	}

	return len(errors) == 0, errors
}
