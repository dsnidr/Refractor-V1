package broadcast

import (
	"github.com/sniddunc/refractor/pkg/regexutils"
	"regexp"
)

type Fields map[string]string

type Broadcast struct {
	Type   string
	Fields Fields
}

const (
	TYPE_JOIN = "JOIN"
	TYPE_QUIT = "QUIT"
)

func GetBroadcastType(broadcast string, patterns map[string]*regexp.Regexp) *Broadcast {
	for bcastType, pattern := range patterns {
		if pattern.MatchString(broadcast) {
			namedMatches := regexutils.MapNamedMatches(pattern, broadcast)

			return &Broadcast{
				Type:   bcastType,
				Fields: namedMatches,
			}
		}
	}

	return nil
}
