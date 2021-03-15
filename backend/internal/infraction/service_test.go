package infraction

import (
	"github.com/sniddunc/refractor/internal/mock"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/internal/player"
	"github.com/sniddunc/refractor/internal/server"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_infractionService_CreateWarning(t *testing.T) {
	testLogger, _ := log.NewLogger(true, false)

	type fields struct {
		mockPlayers map[int64]*refractor.Player
		mockServers map[int64]*refractor.Server
	}
	type args struct {
		userID int64
		body   params.CreateWarningParams
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantInfraction *refractor.Infraction
		wantRes        *refractor.ServiceResponse
	}{
		{
			name: "infraction.createwarning.1",
			fields: fields{
				mockPlayers: map[int64]*refractor.Player{
					1: {
						PlayerID: 1,
					},
				},
				mockServers: map[int64]*refractor.Server{
					1: {
						ServerID: 1,
					},
				},
			},
			args: args{
				userID: 1,
				body: params.CreateWarningParams{
					PlayerID: 1,
					ServerID: 1,
					Reason:   "Test warning reason",
				},
			},
			wantInfraction: &refractor.Infraction{
				InfractionID: 1,
				PlayerID:     1,
				UserID:       1,
				ServerID:     1,
				Type:         refractor.INFRACTION_TYPE_WARNING,
				Reason:       "Test warning reason",
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Warning created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPlayerRepo := mock.NewMockPlayerRepository(tt.fields.mockPlayers)
			playerService := player.NewPlayerService(mockPlayerRepo, testLogger)
			mockServerRepo := mock.NewMockServerRepository(tt.fields.mockServers)
			serverService := server.NewServerService(mockServerRepo, nil, testLogger)
			mockInfractionRepo := mock.NewMockInfractionRepository(map[int64]*refractor.Infraction{})
			infractionService := NewInfractionService(mockInfractionRepo, playerService, serverService, testLogger)

			warning, res := infractionService.CreateWarning(tt.args.userID, tt.args.body)

			assert.True(t, infractionsAreEqual(tt.wantInfraction, warning), "Infractions were not equal\nWant = %v\nGot  = %v", tt.wantInfraction, warning)
			assert.True(t, tt.wantRes.Equals(res), "tt.wantRes = %v and res = %v should be equal", tt.wantRes, res)
		})
	}
}

// infractionsAreEqual compares the following fields to determine is two infractions are equal:
// InfractionID, PlayerID, ServerID, UserID, Type, Reason, SystemAction
func infractionsAreEqual(infraction1 *refractor.Infraction, infraction2 *refractor.Infraction) bool {
	if infraction1.InfractionID != infraction2.InfractionID {
		return false
	}

	if infraction1.PlayerID != infraction2.PlayerID {
		return false
	}

	if infraction1.ServerID != infraction2.ServerID {
		return false
	}

	if infraction1.UserID != infraction2.UserID {
		return false
	}

	if infraction1.Type != infraction2.Type {
		return false
	}

	if infraction1.Reason != infraction2.Reason {
		return false
	}

	if infraction1.Duration != infraction2.Duration {
		return false
	}

	if infraction1.SystemAction != infraction2.SystemAction {
		return false
	}

	return true
}
