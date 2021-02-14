package auth

import (
	"github.com/sniddunc/refractor/internal/mock"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_authService_LogInUser(t *testing.T) {
	const testJWTSecret = "ns@&y72786b!Y&ysht@WBGstaygs"
	testLogger, _ := log.NewLogger(true, false)

	type fields struct {
		mockUsers map[int64]*mock.MockUser
	}
	type args struct {
		body params.LoginParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *refractor.ServiceResponse
	}{
		{
			name: "auth.loginuser.1",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "password",
						User: &refractor.User{
							UserID:              1,
							Email:               "test@test.com",
							Username:            "testuser.1",
							Password:            mock.HashPassword("password"),
							AccessLevel:         0,
							Activated:           true,
							NeedsPasswordChange: false,
						},
					},
				},
			},
			args: args{
				body: params.LoginParams{
					Username: "testuser.1",
					Password: "password",
				},
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Successfully logged in",
			},
		},
		{
			name: "auth.loginuser.2",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "password",
						User: &refractor.User{
							UserID:              1,
							Email:               "test2@test.com",
							Username:            "testuser.1",
							Password:            mock.HashPassword("password"),
							AccessLevel:         0,
							Activated:           true,
							NeedsPasswordChange: false,
						},
					},
				},
			},
			args: args{
				body: params.LoginParams{
					Username: "testuser.1",
					Password: "wrongpassword",
				},
			},
			wantRes: &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid credentials were provided",
			},
		},
		{
			name: "auth.loginuser.3",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "password",
						User: &refractor.User{
							UserID:              1,
							Email:               "test2@test.com",
							Username:            "testuser.1",
							Password:            mock.HashPassword("password"),
							AccessLevel:         0,
							Activated:           true,
							NeedsPasswordChange: false,
						},
					},
				},
			},
			args: args{
				body: params.LoginParams{
					Username: "wrongusername",
					Password: "password",
				},
			},
			wantRes: &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid credentials were provided",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := mock.NewMockUserRepository(tt.fields.mockUsers)
			authService := NewAuthService(userRepo, testLogger, testJWTSecret)

			_, res := authService.LogInUser(tt.args.body)

			assert.Equal(t, tt.wantRes, res, "Responses were not equal")
		})
	}
}
