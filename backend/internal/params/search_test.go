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

func TestSearchInfractionsParams_Validate(t *testing.T) {
	type fields struct {
		Type         string
		PlayerID     string
		UserID       string
		Game         string
		ServerID     string
		SearchParams SearchParams
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "params.search.infractions.validate.1",
			fields: fields{
				Type:     "WARNING",
				PlayerID: "2362793428",
				UserID:   "1",
				Game:     strings.Repeat("a", config.ServerGameMinLen),
				ServerID: "9",
				SearchParams: SearchParams{
					Offset: config.SearchOffsetMin,
					Limit:  config.SearchLimitMin,
				},
			},
			want: true,
		},
		{
			name: "params.search.infractions.validate.2",
			fields: fields{
				Type:     "INVALID",
				PlayerID: "2362793428",
				UserID:   "1",
				Game:     strings.Repeat("a", config.ServerGameMinLen),
				ServerID: "9",
				SearchParams: SearchParams{
					Offset: config.SearchOffsetMin,
					Limit:  config.SearchLimitMin,
				},
			},
			want: false,
		},
		{
			name: "params.search.infractions.validate.3",
			fields: fields{
				Type:     "WARNING",
				PlayerID: "invalid",
				UserID:   "1",
				Game:     strings.Repeat("a", config.ServerGameMinLen),
				ServerID: "9",
				SearchParams: SearchParams{
					Offset: config.SearchOffsetMin,
					Limit:  config.SearchLimitMin,
				},
			},
			want: false,
		},
		{
			name: "params.search.infractions.validate.4",
			fields: fields{
				Type:     "WARNING",
				PlayerID: "2362793428",
				UserID:   "invalid",
				Game:     strings.Repeat("a", config.ServerGameMinLen),
				ServerID: "9",
				SearchParams: SearchParams{
					Offset: config.SearchOffsetMin,
					Limit:  config.SearchLimitMin,
				},
			},
			want: false,
		},
		{
			name: "params.search.infractions.validate.5",
			fields: fields{
				Type:     "WARNING",
				PlayerID: "2362793428",
				UserID:   "1",
				Game:     strings.Repeat("a", config.ServerGameMaxLen+1),
				ServerID: "9",
				SearchParams: SearchParams{
					Offset: config.SearchOffsetMin,
					Limit:  config.SearchLimitMin,
				},
			},
			want: false,
		},
		{
			name: "params.search.infractions.validate.5",
			fields: fields{
				Type:     "WARNING",
				PlayerID: "2362793428",
				UserID:   "1",
				Game:     strings.Repeat("a", config.ServerGameMinLen),
				ServerID: "invalid",
				SearchParams: SearchParams{
					Offset: config.SearchOffsetMin,
					Limit:  config.SearchLimitMin,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &SearchInfractionsParams{
				Type:         tt.fields.Type,
				PlayerID:     tt.fields.PlayerID,
				UserID:       tt.fields.UserID,
				Game:         tt.fields.Game,
				ServerID:     tt.fields.ServerID,
				SearchParams: tt.fields.SearchParams,
			}

			got, errors := body.Validate()
			if got != tt.want {
				t.Errorf("Validate() got = %v, want %v\nErrors: %v", got, tt.want, errors)
			}
		})
	}
}
