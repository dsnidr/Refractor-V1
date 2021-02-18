package regexutils

import "regexp"

func MapNamedMatches(pattern *regexp.Regexp, data string) map[string]string {
	matches := pattern.FindStringSubmatch(data)

	if len(matches) < 1 {
		return nil
	}

	matches = matches[1:] // skip first match since it's the entire match, not just the submatches

	namedMatches := map[string]string{}

	for i, name := range pattern.SubexpNames() {
		// skip the first global match
		if i == 0 {
			continue
		}

		namedMatches[name] = matches[i-1]
	}

	return namedMatches
}
