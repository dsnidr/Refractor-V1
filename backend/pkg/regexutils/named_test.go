package regexutils

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestMapNamedMatches(t *testing.T) {
	type args struct {
		pattern *regexp.Regexp
		data    string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "regexutils.namedmatches.1",
			args: args{
				pattern: regexp.MustCompile("^(?P<ID>[0-9]+)$"),
				data:    "1",
			},
			want: map[string]string{
				"ID": "1",
			},
		},
		{
			name: "regexutils.namedmatches.2",
			args: args{
				pattern: regexp.MustCompile("^(?P<ID>[0-9]+),(?P<Username>[0-9a-zA-Z]+)$"),
				data:    "1,test",
			},
			want: map[string]string{
				"ID":       "1",
				"Username": "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := MapNamedMatches(tt.args.pattern, tt.args.data)

			assert.Equal(t, tt.want, fields)
		})
	}
}
