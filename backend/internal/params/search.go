package params

import (
	"fmt"
	"github.com/sniddunc/refractor/pkg/config"
	"net/url"
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

var validPlayerSearchTypes = []string{"name", "playfabid", "mcuuid"}

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
