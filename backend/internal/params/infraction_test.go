/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package params

import (
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
	"testing"
)

func TestCreateWarningParams_Validate(t *testing.T) {
	type fields struct {
		PlayerID int64
		ServerID int64
		Reason   string
	}
	tests := []struct {
		name      string
		fields    fields
		wantValid bool
	}{
		{
			name: "params.infractions.warning.1",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Reason:   strings.Repeat("a", config.InfractionReasonMinLen),
			},
			wantValid: true,
		},
		{
			name: "params.infractions.warning.2",
			fields: fields{
				PlayerID: 0,
				ServerID: 0,
				Reason:   "",
			},
			wantValid: false,
		},
		{
			name: "params.infractions.warning.3",
			fields: fields{
				PlayerID: math.MinInt32,
				ServerID: math.MaxInt32 + 1,
				Reason:   "",
			},
			wantValid: false,
		},
		{
			name: "params.infractions.warning.4",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Reason:   strings.Repeat("a", config.InfractionReasonMinLen-1),
			},
			wantValid: false,
		},
		{
			name: "params.infractions.warning.5",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Reason:   strings.Repeat("a", config.InfractionReasonMaxLen+1),
			},
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &CreateWarningParams{
				PlayerID: tt.fields.PlayerID,
				ServerID: tt.fields.ServerID,
				Reason:   tt.fields.Reason,
			}

			valid, errors := body.Validate()
			assert.Equal(t, tt.wantValid, valid, "Validate returned the wrong values. Errors: %v", errors)
		})
	}
}

func TestCreateMuteParams_Validate(t *testing.T) {
	type fields struct {
		PlayerID int64
		ServerID int64
		Reason   string
		Duration int
	}
	tests := []struct {
		name      string
		fields    fields
		wantValid bool
	}{
		{
			name: "params.infractions.mute.1",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Duration: 1440,
				Reason:   strings.Repeat("a", config.InfractionReasonMinLen),
			},
			wantValid: true,
		},
		{
			name: "params.infractions.mute.2",
			fields: fields{
				PlayerID: 0,
				ServerID: 0,
				Duration: -1,
				Reason:   "",
			},
			wantValid: false,
		},
		{
			name: "params.infractions.mute.3",
			fields: fields{
				PlayerID: math.MinInt32,
				ServerID: math.MaxInt32 + 1,
				Duration: math.MaxInt32 + 1,
				Reason:   "",
			},
			wantValid: false,
		},
		{
			name: "params.infractions.mute.4",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Duration: -1,
				Reason:   strings.Repeat("a", config.InfractionReasonMinLen-1),
			},
			wantValid: false,
		},
		{
			name: "params.infractions.mute.5",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Duration: 0,
				Reason:   strings.Repeat("a", config.InfractionReasonMaxLen+1),
			},
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &CreateMuteParams{
				PlayerID: tt.fields.PlayerID,
				ServerID: tt.fields.ServerID,
				Reason:   tt.fields.Reason,
				Duration: tt.fields.Duration,
			}

			valid, errors := body.Validate()
			assert.Equal(t, tt.wantValid, valid, "Validate returned the wrong values. Errors: %v", errors)
		})
	}
}

func TestCreateKickParams_Validate(t *testing.T) {
	type fields struct {
		PlayerID int64
		ServerID int64
		Reason   string
	}
	tests := []struct {
		name      string
		fields    fields
		wantValid bool
	}{
		{
			name: "params.infractions.kick.1",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Reason:   strings.Repeat("a", config.InfractionReasonMinLen),
			},
			wantValid: true,
		},
		{
			name: "params.infractions.kick.2",
			fields: fields{
				PlayerID: 0,
				ServerID: 0,
				Reason:   "",
			},
			wantValid: false,
		},
		{
			name: "params.infractions.kick.3",
			fields: fields{
				PlayerID: math.MinInt32,
				ServerID: math.MaxInt32 + 1,
				Reason:   "",
			},
			wantValid: false,
		},
		{
			name: "params.infractions.kick.4",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Reason:   strings.Repeat("a", config.InfractionReasonMinLen-1),
			},
			wantValid: false,
		},
		{
			name: "params.infractions.kick.5",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Reason:   strings.Repeat("a", config.InfractionReasonMaxLen+1),
			},
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &CreateKickParams{
				PlayerID: tt.fields.PlayerID,
				ServerID: tt.fields.ServerID,
				Reason:   tt.fields.Reason,
			}

			valid, errors := body.Validate()
			assert.Equal(t, tt.wantValid, valid, "Validate returned the wrong values. Errors: %v", errors)
		})
	}
}

func TestCreateBanParams_Validate(t *testing.T) {
	type fields struct {
		PlayerID int64
		ServerID int64
		Reason   string
		Duration int
	}
	tests := []struct {
		name      string
		fields    fields
		wantValid bool
	}{
		{
			name: "params.infractions.ban.1",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Duration: 1440,
				Reason:   strings.Repeat("a", config.InfractionReasonMinLen),
			},
			wantValid: true,
		},
		{
			name: "params.infractions.ban.2",
			fields: fields{
				PlayerID: 0,
				ServerID: 0,
				Duration: -1,
				Reason:   "",
			},
			wantValid: false,
		},
		{
			name: "params.infractions.ban.3",
			fields: fields{
				PlayerID: math.MinInt32,
				ServerID: math.MaxInt32 + 1,
				Duration: math.MaxInt32 + 1,
				Reason:   "",
			},
			wantValid: false,
		},
		{
			name: "params.infractions.ban.4",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Duration: -1,
				Reason:   strings.Repeat("a", config.InfractionReasonMinLen-1),
			},
			wantValid: false,
		},
		{
			name: "params.infractions.ban.5",
			fields: fields{
				PlayerID: 1,
				ServerID: 1,
				Duration: 0,
				Reason:   strings.Repeat("a", config.InfractionReasonMaxLen+1),
			},
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &CreateMuteParams{
				PlayerID: tt.fields.PlayerID,
				ServerID: tt.fields.ServerID,
				Reason:   tt.fields.Reason,
				Duration: tt.fields.Duration,
			}

			valid, errors := body.Validate()
			assert.Equal(t, tt.wantValid, valid, "Validate returned the wrong values. Errors: %v", errors)
		})
	}
}
