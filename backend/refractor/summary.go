package refractor

import "github.com/labstack/echo/v4"

type PlayerSummary struct {
	Warnings []*Infraction `json:"warnings"`
	Mutes    []*Infraction `json:"mutes"`
	Kicks    []*Infraction `json:"kicks"`
	Bans     []*Infraction `json:"bans"`
	*Player
}

type SummaryService interface {
	GetPlayerSummary(id int64) (*PlayerSummary, *ServiceResponse)
}

type SummaryHandler interface {
	GetPlayerSummary(c echo.Context) error
}
