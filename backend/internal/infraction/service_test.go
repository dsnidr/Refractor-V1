package infraction

import (
	"database/sql"
	"github.com/sniddunc/refractor/internal/mock"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/internal/player"
	"github.com/sniddunc/refractor/internal/server"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/pkg/perms"
	"github.com/sniddunc/refractor/refractor"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func Test_infractionService_CreateWarning(t *testing.T) {
	testLogger, _ := log.NewLogger(true, false)

	type fields struct {
		mockPlayers map[int64]*refractor.DBPlayer
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
				mockPlayers: map[int64]*refractor.DBPlayer{
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
				Message:    "Infraction created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPlayerRepo := mock.NewMockPlayerRepository(tt.fields.mockPlayers)
			playerService := player.NewPlayerService(mockPlayerRepo, testLogger)
			mockServerRepo := mock.NewMockServerRepository(tt.fields.mockServers)
			serverService := server.NewServerService(mockServerRepo, nil, testLogger)
			mockInfractionRepo := mock.NewMockInfractionRepository(map[int64]*refractor.DBInfraction{})
			infractionService := NewInfractionService(mockInfractionRepo, playerService, serverService, testLogger)

			warning, res := infractionService.CreateWarning(tt.args.userID, tt.args.body)

			assert.True(t, infractionsAreEqual(tt.wantInfraction, warning), "Infractions were not equal\nWant = %v\nGot  = %v", tt.wantInfraction, warning)
			assert.True(t, tt.wantRes.Equals(res), "tt.wantRes = %v and res = %v should be equal", tt.wantRes, res)
		})
	}
}

func Test_infractionService_DeleteInfraction(t *testing.T) {
	testLogger, _ := log.NewLogger(true, false)

	type fields struct {
		mockInfractions map[int64]*refractor.DBInfraction
	}
	type args struct {
		id   int64
		user params.UserMeta
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *refractor.ServiceResponse
	}{
		{
			name: "infraction.deleteinfraction.1",
			fields: fields{
				mockInfractions: map[int64]*refractor.DBInfraction{
					1: {
						InfractionID: 1,
						PlayerID:     1,
						UserID:       1,
						ServerID:     1,
						Type:         refractor.INFRACTION_TYPE_WARNING,
						Reason:       sql.NullString{String: strings.Repeat("a", config.InfractionReasonMinLen), Valid: true},
						Duration:     sql.NullInt32{Int32: 0, Valid: true},
						Timestamp:    0,
						SystemAction: false,
					},
				},
			},
			args: args{
				id: 1,
				user: params.UserMeta{
					UserID:      1,
					Permissions: perms.DELETE_OWN_INFRACTIONS,
				},
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Infraction deleted",
			},
		},
		{
			name: "infraction.deleteinfraction.2",
			fields: fields{
				mockInfractions: map[int64]*refractor.DBInfraction{
					1: {
						InfractionID: 1,
						PlayerID:     1,
						UserID:       1,
						ServerID:     1,
						Type:         refractor.INFRACTION_TYPE_WARNING,
						Reason:       sql.NullString{String: strings.Repeat("a", config.InfractionReasonMinLen), Valid: true},
						Duration:     sql.NullInt32{Int32: 0, Valid: true},
						Timestamp:    0,
						SystemAction: false,
					},
				},
			},
			args: args{
				id: 1,
				user: params.UserMeta{
					UserID:      1,
					Permissions: perms.FULL_ACCESS,
				},
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Infraction deleted",
			},
		},
		{
			name: "infraction.deleteinfraction.3",
			fields: fields{
				mockInfractions: map[int64]*refractor.DBInfraction{
					1: {
						InfractionID: 1,
						PlayerID:     1,
						UserID:       1,
						ServerID:     1,
						Type:         refractor.INFRACTION_TYPE_WARNING,
						Reason:       sql.NullString{String: strings.Repeat("a", config.InfractionReasonMinLen), Valid: true},
						Duration:     sql.NullInt32{Int32: 0, Valid: true},
						Timestamp:    0,
						SystemAction: false,
					},
				},
			},
			args: args{
				id: 1,
				user: params.UserMeta{
					UserID:      2,
					Permissions: perms.DELETE_ANY_INFRACTION,
				},
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Infraction deleted",
			},
		},
		{
			name: "infraction.deleteinfraction.4",
			fields: fields{
				mockInfractions: map[int64]*refractor.DBInfraction{
					1: {
						InfractionID: 1,
						PlayerID:     1,
						UserID:       1,
						ServerID:     1,
						Type:         refractor.INFRACTION_TYPE_WARNING,
						Reason:       sql.NullString{String: strings.Repeat("a", config.InfractionReasonMinLen), Valid: true},
						Duration:     sql.NullInt32{Int32: 0, Valid: true},
						Timestamp:    0,
						SystemAction: false,
					},
				},
			},
			args: args{
				id: 1,
				user: params.UserMeta{
					UserID:      2,
					Permissions: perms.DELETE_OWN_INFRACTIONS,
				},
			},
			wantRes: &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageNoPermission,
			},
		},
		{
			name: "infraction.deleteinfraction.5",
			fields: fields{
				mockInfractions: map[int64]*refractor.DBInfraction{
					1: {
						InfractionID: 1,
						PlayerID:     1,
						UserID:       1,
						ServerID:     1,
						Type:         refractor.INFRACTION_TYPE_WARNING,
						Reason:       sql.NullString{String: strings.Repeat("a", config.InfractionReasonMinLen), Valid: true},
						Duration:     sql.NullInt32{Int32: 0, Valid: true},
						Timestamp:    0,
						SystemAction: false,
					},
				},
			},
			args: args{
				id: 1,
				user: params.UserMeta{
					UserID:      2,
					Permissions: 0,
				},
			},
			wantRes: &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageNoPermission,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInfractionRepo := mock.NewMockInfractionRepository(tt.fields.mockInfractions)
			infractionService := NewInfractionService(mockInfractionRepo, nil, nil, testLogger)

			res := infractionService.DeleteInfraction(tt.args.id, tt.args.user)

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
