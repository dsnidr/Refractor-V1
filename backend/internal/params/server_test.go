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
	"strings"
	"testing"
)

func TestCreateServerParams_Validate(t *testing.T) {
	type fields struct {
		Game         string
		Name         string
		Address      string
		RconPort     string
		RconPassword string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "params.createserver.1",
			fields: fields{
				Game:         "testgame",
				Name:         "Test Server",
				Address:      "127.0.0.1",
				RconPort:     "17262",
				RconPassword: "password",
			},
			want: true,
		},
		{
			name: "params.createserver.2",
			fields: fields{
				Game:         "testgame",
				Name:         "Local Test Server",
				Address:      "192.168.1.1",
				RconPort:     "1778",
				RconPassword: "rc*a72etyhast2@[]q@{@&Yhsa",
			},
			want: true,
		},
		{
			name: "params.createserver.3",
			fields: fields{
				Game:         "testgame",
				Name:         strings.Repeat("a", config.ServerNameMinLen),
				Address:      "0.0.0.0",
				RconPort:     "23",
				RconPassword: strings.Repeat("p", config.ServerPasswordMinLen),
			},
			want: true,
		},
		{
			name: "params.createserver.4",
			fields: fields{
				Game:         "testgame",
				Name:         strings.Repeat("a", config.ServerNameMaxLen),
				Address:      "111.111.111.111",
				RconPort:     "65535",
				RconPassword: strings.Repeat("p", config.ServerPasswordMaxLen),
			},
			want: true,
		},
		{
			name: "params.createserver.5",
			fields: fields{
				Game:         "testgame",
				Name:         "hey! this isn't a valid server name!",
				Address:      "haha invalid IP address funny",
				RconPort:     ":) totally a legit port",
				RconPassword: "L",
			},
			want: false,
		},
		{
			name: "params.createserver.6",
			fields: fields{
				Game:         "testgame",
				Name:         "hey! this isn't a valid server name!",
				Address:      "192.168.1.2",
				RconPort:     "4322",
				RconPassword: "password",
			},
			want: false,
		},
		{
			name: "params.createserver.7",
			fields: fields{
				Game:         "testgame",
				Name:         "valid name",
				Address:      "192.168.1.2",
				RconPort:     "4322",
				RconPassword: strings.Repeat("p", config.ServerPasswordMinLen-1),
			},
			want: false,
		},
		{
			name: "params.createserver.8",
			fields: fields{
				Game:         "testgame",
				Name:         "valid name",
				Address:      "192.168.1.2",
				RconPort:     "4322",
				RconPassword: strings.Repeat("p", config.ServerPasswordMaxLen+1),
			},
			want: false,
		},
		{
			name: "params.createserver.9",
			fields: fields{
				Game:         strings.Repeat("a", config.ServerGameMinLen),
				Name:         "valid name",
				Address:      "192.168.1.2",
				RconPort:     "4322",
				RconPassword: strings.Repeat("p", config.ServerPasswordMinLen),
			},
			want: true,
		},
		{
			name: "params.createserver.10",
			fields: fields{
				Game:         strings.Repeat("a", config.ServerGameMaxLen),
				Name:         "valid name",
				Address:      "192.168.1.2",
				RconPort:     "4322",
				RconPassword: strings.Repeat("p", config.ServerPasswordMinLen),
			},
			want: true,
		},
		{
			name: "params.createserver.11",
			fields: fields{
				Game:         strings.Repeat("a", config.ServerGameMinLen-1),
				Name:         "valid name",
				Address:      "192.168.1.2",
				RconPort:     "4322",
				RconPassword: strings.Repeat("p", config.ServerPasswordMinLen),
			},
			want: false,
		},
		{
			name: "params.createserver.12",
			fields: fields{
				Game:         strings.Repeat("a", config.ServerGameMaxLen+1),
				Name:         "valid name",
				Address:      "192.168.1.2",
				RconPort:     "4322",
				RconPassword: strings.Repeat("p", config.ServerPasswordMinLen),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &CreateServerParams{
				Game:         tt.fields.Game,
				Name:         tt.fields.Name,
				Address:      tt.fields.Address,
				RCONPort:     tt.fields.RconPort,
				RCONPassword: tt.fields.RconPassword,
			}

			got, errors := body.Validate()
			assert.Equal(t, tt.want, got, "Validate returned the wrong values. Errors: %v", errors)
		})
	}
}

func TestUpdateServerParams_Validate(t *testing.T) {
	type fields struct {
		Name         string
		Address      string
		RCONPort     string
		RCONPassword string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "params.server.1",
			fields: fields{
				Name:         "Test Server",
				Address:      "127.0.0.1",
				RCONPort:     "17262",
				RCONPassword: "password",
			},
			want: true,
		},
		{
			name: "params.server.2",
			fields: fields{
				Name:         "Local Test Server",
				Address:      "192.168.1.1",
				RCONPort:     "1778",
				RCONPassword: "rc*a72etyhast2@[]q@{@&Yhsa",
			},
			want: true,
		},
		{
			name: "params.server.3",
			fields: fields{
				Name:         strings.Repeat("a", config.ServerNameMinLen),
				Address:      "0.0.0.0",
				RCONPort:     "23",
				RCONPassword: strings.Repeat("p", config.ServerPasswordMinLen),
			},
			want: true,
		},
		{
			name: "params.server.4",
			fields: fields{
				Name:         strings.Repeat("a", config.ServerNameMaxLen),
				Address:      "111.111.111.111",
				RCONPort:     "65535",
				RCONPassword: strings.Repeat("p", config.ServerPasswordMaxLen),
			},
			want: true,
		},
		{
			name: "params.server.5",
			fields: fields{
				Name:         "hey! this isn't a valid server name!",
				Address:      "invalid IP address goes here",
				RCONPort:     ":) totally a legit port",
				RCONPassword: "L",
			},
			want: false,
		},
		{
			name: "params.server.6",
			fields: fields{
				Name:         "hey! this isn't a valid server name!",
				Address:      "192.168.1.2",
				RCONPort:     "4322",
				RCONPassword: "password",
			},
			want: false,
		},
		{
			name: "params.server.7",
			fields: fields{
				Name:         "valid name",
				Address:      "192.168.1.2",
				RCONPort:     "4322",
				RCONPassword: strings.Repeat("p", config.ServerPasswordMinLen-1),
			},
			want: true,
		},
		{
			name: "params.server.8",
			fields: fields{
				Name:         "valid name",
				Address:      "192.168.1.2",
				RCONPort:     "4322",
				RCONPassword: strings.Repeat("p", config.ServerPasswordMaxLen+1),
			},
			want: false,
		},
		{
			name:   "params.server.9",
			fields: fields{},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &UpdateServerParams{
				Name:         tt.fields.Name,
				Address:      tt.fields.Address,
				RCONPort:     tt.fields.RCONPort,
				RCONPassword: tt.fields.RCONPassword,
			}

			got, errors := body.Validate()
			assert.Equal(t, tt.want, got, "Validate returned the wrong values. Errors: %v", errors)
		})
	}
}
