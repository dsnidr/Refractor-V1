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
