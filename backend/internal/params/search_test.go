package params

import (
	"github.com/sniddunc/refractor/pkg/config"
	"strings"
	"testing"
)

func TestSearchPlayerParams_Validate(t *testing.T) {
	type fields struct {
		SearchTerm   string
		SearchType   string
		SearchParams SearchParams
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "TestSearchPlayerParams_Validate-01",
			fields: fields{
				SearchTerm: strings.Repeat("a", config.SearchTermMinLen),
				SearchType: "name",
				SearchParams: SearchParams{
					Offset: config.SearchOffsetMin,
					Limit:  config.SearchLimitMin,
				},
			},
			want: true,
		},
		{
			name: "TestSearchPlayerParams_Validate-02",
			fields: fields{
				SearchTerm: strings.Repeat("a", config.SearchTermMaxLen),
				SearchType: "playfabid",
				SearchParams: SearchParams{
					Offset: config.SearchOffsetMax,
					Limit:  config.SearchLimitMax,
				},
			},
			want: true,
		},
		{
			name: "TestSearchPlayerParams_Validate-03",
			fields: fields{
				SearchTerm: strings.Repeat("a", config.SearchTermMinLen-1),
				SearchType: "invalid",
				SearchParams: SearchParams{
					Offset: config.SearchOffsetMin - 1,
					Limit:  config.SearchLimitMin - 1,
				},
			},
			want: false,
		},
		{
			name: "TestSearchPlayerParams_Validate-04",
			fields: fields{
				SearchTerm: strings.Repeat("a", config.SearchTermMaxLen+1),
				SearchType: "invalid",
				SearchParams: SearchParams{
					Offset: config.SearchOffsetMax + 1,
					Limit:  config.SearchLimitMax + 1,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &SearchPlayersParams{
				SearchTerm:   tt.fields.SearchTerm,
				SearchType:   tt.fields.SearchType,
				SearchParams: tt.fields.SearchParams,
			}

			got, errors := body.Validate()
			if got != tt.want {
				t.Errorf("Validate() got = %v, want %v\nErrors: %v", got, tt.want, errors)
			}
		})
	}
}
