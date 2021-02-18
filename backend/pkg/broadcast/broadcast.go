package broadcast

import (
	"github.com/sniddunc/refractor/pkg/regexutils"
	"regexp"
)

type Broadcast struct {
	Type   string
	Fields map[string]string
}

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
