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
				RconPort:     tt.fields.RconPort,
				RconPassword: tt.fields.RconPassword,
			}

			got, errors := body.Validate()
			assert.Equal(t, tt.want, got, "Validate returned the wrong values. Errors: %v", errors)
		})
	}
}
