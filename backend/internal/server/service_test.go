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

package server

import (
	"github.com/sniddunc/refractor/internal/game"
	"github.com/sniddunc/refractor/internal/mock"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_serverService_CreateServer(t *testing.T) {
	testLogger, _ := log.NewLogger(true, false)

	type fields struct {
		mockServers map[int64]*refractor.Server
	}
	type args struct {
		body params.CreateServerParams
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantServer *refractor.Server
		wantRes    *refractor.ServiceResponse
	}{
		{
			name: "server.createserver.1",
			fields: fields{
				mockServers: map[int64]*refractor.Server{},
			},
			args: args{
				body: params.CreateServerParams{
					Name:         "testserver.1",
					Game:         mock.NewMockGame().GetName(),
					Address:      "127.0.0.1",
					RCONPort:     "7777",
					RCONPassword: "RconPassword",
				},
			},
			wantServer: &refractor.Server{
				ServerID:     1,
				Name:         "testserver.1",
				Game:         mock.NewMockGame().GetName(),
				Address:      "127.0.0.1",
				RCONPort:     "7777",
				RCONPassword: "RconPassword",
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Server created",
			},
		},
		{
			name: "server.createserver.2",
			fields: fields{
				mockServers: map[int64]*refractor.Server{},
			},
			args: args{
				body: params.CreateServerParams{
					Name:         "testserver.1",
					Game:         "invalid game",
					Address:      "127.0.0.1",
					RCONPort:     "7777",
					RCONPassword: "RconPassword",
				},
			},
			wantServer: nil,
			wantRes: &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				ValidationErrors: url.Values{
					"game": []string{"Invalid game"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServerRepo := mock.NewMockServerRepository(tt.fields.mockServers)
			gameService := game.NewGameService()
			gameService.AddGame(mock.NewMockGame())
			serverService := NewServerService(mockServerRepo, gameService, testLogger)

			server, res := serverService.CreateServer(tt.args.body)

			assert.Equal(t, tt.wantServer, server, "Structs are not equal")
			assert.True(t, tt.wantRes.Equals(res), "tt.wantRes = %v and res = %v should be equal", tt.wantRes, res)
		})
	}
}

func Test_serverService_EditServer(t *testing.T) {
	testLogger, _ := log.NewLogger(true, false)

	type fields struct {
		mockServers map[int64]*refractor.Server
	}
	type args struct {
		id   int64
		body params.UpdateServerParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *refractor.Server
		wantRes *refractor.ServiceResponse
	}{
		{
			name: "server.editserver.1",
			fields: fields{
				mockServers: map[int64]*refractor.Server{
					1: {
						ServerID:     1,
						Name:         "Test Server",
						Game:         "Test Game",
						Address:      "127.0.0,1",
						RCONPort:     "1000",
						RCONPassword: "Password",
					},
				},
			},
			args: args{
				id: 1,
				body: params.UpdateServerParams{
					Name:         "Updated server name",
					Address:      "192.168.0.1",
					RCONPort:     "2383",
					RCONPassword: "UpdatedPassword",
				},
			},
			want: &refractor.Server{
				ServerID:     1,
				Name:         "Updated server name",
				Game:         "Test Game",
				Address:      "192.168.0.1",
				RCONPort:     "2383",
				RCONPassword: "UpdatedPassword",
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Server updated. RCON changes will come into effect the next time Refractor is restarted.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServerRepo := mock.NewMockServerRepository(tt.fields.mockServers)
			serverService := NewServerService(mockServerRepo, nil, testLogger)

			gotServer, gotRes := serverService.UpdateServer(tt.args.id, tt.args.body)

			assert.Equal(t, tt.want, gotServer, "Servers did not match")
			assert.Equal(t, tt.wantRes, gotRes, "Responses did not match")
		})
	}
}
