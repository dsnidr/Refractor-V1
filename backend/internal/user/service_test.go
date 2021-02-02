package user

import (
	"github.com/sniddunc/refractor/internal/mock"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"testing"
)

func Test_userService_CreateUser(t *testing.T) {
	type fields struct {
		mockUsers map[int64]*mock.MockUser
	}
	type args struct {
		body params.CreateUserParams
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser *refractor.User
		wantRes  *refractor.ServiceResponse
	}{
		{
			name: "userservice.create.1",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{},
			},
			args: args{
				body: params.CreateUserParams{
					Email:           "test@test.com",
					Username:        "testuser.1",
					Password:        "password",
					PasswordConfirm: "password",
				},
			},
			wantUser: &refractor.User{
				UserID:              1,
				Email:               "test@test.com",
				Username:            "testuser.1",
				Password:            mock.HashPassword("password"),
				AccessLevel:         0,
				Activated:           true,
				NeedsPasswordChange: true,
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "User created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := mock.NewMockUserRepository(tt.fields.mockUsers)
			mockLogger, _ := log.NewLogger(true, false)
			userService := NewUserService(mockUserRepo, mockLogger)

			newUser, res := userService.CreateUser(tt.args.body)

			assert.Equal(t, tt.wantUser.UserID, newUser.UserID, "UserIDs are not equal")
			assert.Equal(t, tt.wantUser.Email, newUser.Email, "Emails are not equal")
			assert.Equal(t, tt.wantUser.Username, newUser.Username, "Usernames are not equal")
			assert.Equal(t, tt.wantUser.AccessLevel, newUser.AccessLevel, "AccessLevels are not equal")
			assert.True(t, tt.wantRes.Equals(res), "tt.wantRes = %v and res = %v should be equal", tt.wantRes, res)

			assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(tt.args.body.Password)),
				"Password hash comparison failed. Password = %v Hash = %v", tt.args.body.Password, newUser.Password)
		})
	}
}

func Test_userService_SetUserAccessLevel(t *testing.T) {
	type fields struct {
		mockUsers map[int64]*mock.MockUser
	}
	type args struct {
		body params.SetUserAccessLevelParams
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser *refractor.User
		wantRes  *refractor.ServiceResponse
	}{
		{
			name: "userservice.setaccesslevel.1",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "password",
						User: &refractor.User{
							UserID:              1,
							Email:               "test@test.com",
							Username:            "testuser.1",
							Password:            "",
							AccessLevel:         0,
							Activated:           true,
							NeedsPasswordChange: true,
						},
					},
				},
			},
			args: args{
				body: params.SetUserAccessLevelParams{
					UserID:      1,
					AccessLevel: config.AL_ADMIN,
					UserMeta: &params.UserMeta{
						UserID:      2,
						AccessLevel: config.AL_SUPERADMIN,
					},
				},
			},
			wantUser: &refractor.User{
				UserID:              1,
				Email:               "test@test.com",
				Username:            "testuser.1",
				Password:            "",
				AccessLevel:         config.AL_ADMIN,
				Activated:           true,
				NeedsPasswordChange: true,
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Access level set. Any new access rights will come into effect next time the user logs in",
			},
		},
		{
			name: "userservice.setaccesslevel.2",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "password",
						User: &refractor.User{
							UserID:              1,
							Email:               "test@test.com",
							Username:            "testuser.1",
							Password:            "",
							AccessLevel:         0,
							Activated:           true,
							NeedsPasswordChange: true,
						},
					},
				},
			},
			args: args{
				body: params.SetUserAccessLevelParams{
					UserID:      1,
					AccessLevel: config.AL_ADMIN,
					UserMeta: &params.UserMeta{
						UserID:      2,
						AccessLevel: config.AL_USER,
					},
				},
			},
			wantUser: nil,
			wantRes: &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    "You do not have permission to set the access level of this user",
			},
		},
		{
			name: "userservice.setaccesslevel.3",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "password",
						User: &refractor.User{
							UserID:              1,
							Email:               "test@test.com",
							Username:            "testuser.1",
							Password:            "",
							AccessLevel:         0,
							Activated:           true,
							NeedsPasswordChange: true,
						},
					},
				},
			},
			args: args{
				body: params.SetUserAccessLevelParams{
					UserID:      1,
					AccessLevel: config.AL_SUPERADMIN,
					UserMeta: &params.UserMeta{
						UserID:      2,
						AccessLevel: config.AL_ADMIN,
					},
				},
			},
			wantUser: nil,
			wantRes: &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    "You do not have permission to set the access level of this user",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := mock.NewMockUserRepository(tt.fields.mockUsers)
			mockLogger, _ := log.NewLogger(true, false)
			userService := NewUserService(mockUserRepo, mockLogger)

			newUser, res := userService.SetUserAccessLevel(tt.args.body)

			assert.Equal(t, tt.wantUser, newUser, "Structs are not equal")
			assert.True(t, tt.wantRes.Equals(res), "tt.wantRes = %v and res = %v should be equal", tt.wantRes, res)
		})
	}
}

func Test_userService_GetUserInfo(t *testing.T) {
	type fields struct {
		mockUsers map[int64]*mock.MockUser
	}
	tests := []struct {
		name     string
		fields   fields
		userID   int64
		wantInfo *refractor.UserInfo
		wantRes  *refractor.ServiceResponse
	}{
		{
			name: "userservice.create.1",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "",
						User: &refractor.User{
							UserID:              1,
							Email:               "test@test.com",
							Username:            "testuser.1",
							Password:            "password",
							AccessLevel:         config.AL_ADMIN,
							Activated:           true,
							NeedsPasswordChange: true,
						},
					},
				},
			},
			userID: 1,
			wantInfo: &refractor.UserInfo{
				ID:                  1,
				Email:               "test@test.com",
				Username:            "testuser.1",
				Activated:           true,
				AccessLevel:         config.AL_ADMIN,
				NeedsPasswordChange: true,
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "User info retrieved",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := mock.NewMockUserRepository(tt.fields.mockUsers)
			mockLogger, _ := log.NewLogger(true, false)
			userService := NewUserService(mockUserRepo, mockLogger)

			userInfo, res := userService.GetUserInfo(tt.userID)

			assert.Equal(t, tt.wantInfo, userInfo, "Structs are not equal")
			assert.True(t, tt.wantRes.Equals(res), "tt.wantRes = %v and res = %v should be equal", tt.wantRes, res)
		})
	}
}
