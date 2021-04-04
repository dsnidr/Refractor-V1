package config

import "math"

var (
	// Auth
	UsernameMinLen = 1
	UsernameMaxLen = 20
	PasswordMinLen = 8
	PasswordMaxLen = 80

	// Server
	ServerNameMinLen     = 1
	ServerNameMaxLen     = 32
	ServerGameMinLen     = 1
	ServerGameMaxLen     = 32
	ServerPasswordMinLen = 1
	ServerPasswordMaxLen = 64

	// Infractions
	InfractionReasonMinLen       = 1
	InfractionReasonMaxLen       = 4096
	InfractionDurationMax        = math.MaxInt32
	RecentInfractionsReturnCount = 20

	// Search
	SearchTermMinLen = 1
	SearchTermMaxLen = 64
	SearchOffsetMin  = 0
	SearchOffsetMax  = 2147483647 // max int32 value
	SearchLimitMin   = 1
	SearchLimitMax   = 100

	// Players
	RecentPlayersMaxSize = 22
)
