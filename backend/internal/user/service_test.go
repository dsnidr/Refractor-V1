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

package user

import (
	"github.com/sniddunc/bitperms"
	"github.com/sniddunc/refractor/internal/mock"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/pkg/perms"
	"github.com/sniddunc/refractor/refractor"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/url"
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
				Permissions:         perms.DEFAULT_PERMS,
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
			assert.Equal(t, tt.wantUser.Permissions, newUser.Permissions, "Permission values are not equal")
			assert.True(t, tt.wantRes.Equals(res), "tt.wantRes = %v and res = %v should be equal", tt.wantRes, res)

			assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(tt.args.body.Password)),
				"Password hash comparison failed. Password = %v Hash = %v", tt.args.body.Password, newUser.Password)
		})
	}
}

func Test_userService_SetUserPermissions(t *testing.T) {
	type fields struct {
		mockUsers map[int64]*mock.MockUser
	}
	type args struct {
		body params.SetUserPermissionsParams
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
							Permissions:         perms.DEFAULT_PERMS,
							Activated:           true,
							NeedsPasswordChange: true,
						},
					},
				},
			},
			args: args{
				body: params.SetUserPermissionsParams{
					UserID:           1,
					PermissionString: bitperms.PermissionValue(perms.FULL_ACCESS).Serialize(),
					Permissions:      perms.FULL_ACCESS,
					UserMeta: &params.UserMeta{
						UserID:      2,
						Permissions: perms.SUPER_ADMIN,
					},
				},
			},
			wantUser: &refractor.User{
				UserID:              1,
				Email:               "test@test.com",
				Username:            "testuser.1",
				Password:            "",
				Permissions:         perms.FULL_ACCESS,
				Activated:           true,
				NeedsPasswordChange: true,
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Permissions set. Any new access rights will come into effect next time the reloads.",
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
							Permissions:         perms.DEFAULT_PERMS,
							Activated:           true,
							NeedsPasswordChange: true,
						},
					},
				},
			},
			args: args{
				body: params.SetUserPermissionsParams{
					UserID:           1,
					PermissionString: bitperms.PermissionValue(perms.FULL_ACCESS).Serialize(),
					Permissions:      perms.FULL_ACCESS,
					UserMeta: &params.UserMeta{
						UserID:      2,
						Permissions: perms.DEFAULT_PERMS,
					},
				},
			},
			wantUser: nil,
			wantRes: &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageNoPermission,
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
							Permissions:         perms.DEFAULT_PERMS,
							Activated:           true,
							NeedsPasswordChange: true,
						},
					},
				},
			},
			args: args{
				body: params.SetUserPermissionsParams{
					UserID:           1,
					PermissionString: bitperms.PermissionValue(perms.SUPER_ADMIN).Serialize(),
					Permissions:      perms.SUPER_ADMIN,
					UserMeta: &params.UserMeta{
						UserID:      2,
						Permissions: perms.DEFAULT_PERMS,
					},
				},
			},
			wantUser: nil,
			wantRes: &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageNoPermission,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := mock.NewMockUserRepository(tt.fields.mockUsers)
			mockLogger, _ := log.NewLogger(true, false)
			userService := NewUserService(mockUserRepo, mockLogger)

			newUser, res := userService.SetUserPermissions(tt.args.body)

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
							Permissions:         perms.FULL_ACCESS,
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
				Permissions:         perms.FULL_ACCESS,
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

func Test_userService_ChangeUserPassword(t *testing.T) {
	type fields struct {
		mockUsers map[int64]*mock.MockUser
	}
	type args struct {
		id   int64
		body params.ChangeUserPassword
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *refractor.ServiceResponse
	}{
		{
			name: "userservice.changepassword.1",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "password",
						User: &refractor.User{
							UserID:              1,
							Email:               "test@test.com",
							Username:            "testuser.1",
							Password:            mock.HashPassword("password"),
							Permissions:         perms.FULL_ACCESS,
							Activated:           true,
							NeedsPasswordChange: true,
						},
					},
				},
			},
			args: args{
				id: 1,
				body: params.ChangeUserPassword{
					CurrentPassword:    "password",
					NewPassword:        "newpassword2",
					NewPasswordConfirm: "newpassword",
				},
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Password changed",
			},
		},
		{
			name: "userservice.changepassword.2",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "password",
						User: &refractor.User{
							UserID:              1,
							Email:               "test@test.com",
							Username:            "testuser.1",
							Password:            mock.HashPassword("password"),
							Permissions:         perms.FULL_ACCESS,
							Activated:           true,
							NeedsPasswordChange: true,
						},
					},
				},
			},
			args: args{
				id: 1,
				body: params.ChangeUserPassword{
					CurrentPassword:    "password",
					NewPassword:        "password",
					NewPasswordConfirm: "password",
				},
			},
			wantRes: &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				ValidationErrors: url.Values{
					"newPassword": []string{"You can't re-use your current password"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := mock.NewMockUserRepository(tt.fields.mockUsers)
			mockLogger, _ := log.NewLogger(true, false)
			userService := NewUserService(mockUserRepo, mockLogger)

			updatedUser, res := userService.ChangeUserPassword(tt.args.id, tt.args.body)
			assert.True(t, tt.wantRes.Equals(res), "tt.wantRes = %v and res = %v should be equal", tt.wantRes, res)

			if tt.wantRes.Success {
				// Make sure password was changed
				assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte(tt.args.body.NewPassword)))
			}
		})
	}
}

func Test_userService_UpdateUser(t *testing.T) {
	type fields struct {
		mockUsers map[int64]*mock.MockUser
	}
	type args struct {
		id   int64
		args refractor.UpdateArgs
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser *refractor.User
		wantRes  *refractor.ServiceResponse
	}{
		{
			name: "userservice.updateuser.1",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "password",
						User: &refractor.User{
							UserID:              1,
							Email:               "test@test.com",
							Username:            "testuser.1",
							Password:            "doesntmatter",
							Permissions:         perms.FULL_ACCESS,
							Activated:           true,
							NeedsPasswordChange: false,
						},
					},
				},
			},
			args: args{
				id: 1,
				args: refractor.UpdateArgs{
					"Activated":           false,
					"Permissions":         perms.FULL_ACCESS,
					"NeedsPasswordChange": true,
				},
			},
			wantUser: &refractor.User{
				UserID:              1,
				Email:               "test@test.com",
				Username:            "testuser.1",
				Password:            "doesntmatter",
				Permissions:         perms.FULL_ACCESS,
				Activated:           false,
				NeedsPasswordChange: true,
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "User updated",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := mock.NewMockUserRepository(tt.fields.mockUsers)
			mockLogger, _ := log.NewLogger(true, false)
			userService := NewUserService(mockUserRepo, mockLogger)

			updatedUser, res := userService.UpdateUser(tt.args.id, tt.args.args)

			assert.Equal(t, tt.wantUser, updatedUser, "updatedUser does not match the expected user")
			assert.True(t, tt.wantRes.Equals(res), "tt.wantRes = %v and res = %v should be equal", tt.wantRes, res)
		})
	}
}

func Test_userService_SetUserPassword(t *testing.T) {
	type fields struct {
		mockUsers map[int64]*mock.MockUser
	}
	type args struct {
		body params.SetUserPasswordParams
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser *refractor.User
		wantRes  *refractor.ServiceResponse
	}{
		{
			name: "userservice.setuserpassword.1",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "password",
						User: &refractor.User{
							UserID:              1,
							Email:               "test@test.com",
							Username:            "testuser.1",
							Password:            mock.HashPassword("password"),
							Permissions:         perms.FULL_ACCESS,
							Activated:           true,
							NeedsPasswordChange: false,
						},
					},
					2: {
						UnhashedPassword: "changeme",
						User: &refractor.User{
							UserID:              2,
							Email:               "test2@test.com",
							Username:            "testuser.2",
							Password:            mock.HashPassword("changeme"),
							Permissions:         perms.DEFAULT_PERMS,
							Activated:           true,
							NeedsPasswordChange: false,
						},
					},
				},
			},
			args: args{
				body: params.SetUserPasswordParams{
					UserID:      2,
					NewPassword: "changedpassword",
					UserMeta: &params.UserMeta{
						UserID:      1,
						Permissions: perms.FULL_ACCESS,
					},
				},
			},
			wantUser: &refractor.User{
				UserID:              1,
				Email:               "test@test.com",
				Username:            "testuser.1",
				Password:            "doesntmatter",
				Permissions:         perms.FULL_ACCESS,
				Activated:           false,
				NeedsPasswordChange: true,
			},
			wantRes: &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "New password set",
			},
		},
		{
			name: "userservice.setuserpassword.2",
			fields: fields{
				mockUsers: map[int64]*mock.MockUser{
					1: {
						UnhashedPassword: "password",
						User: &refractor.User{
							UserID:              1,
							Email:               "test@test.com",
							Username:            "testuser.1",
							Password:            mock.HashPassword("password"),
							Permissions:         perms.FULL_ACCESS,
							Activated:           true,
							NeedsPasswordChange: false,
						},
					},
					2: {
						UnhashedPassword: "changeme",
						User: &refractor.User{
							UserID:              1,
							Email:               "test2@test.com",
							Username:            "testuser.2",
							Password:            mock.HashPassword("changeme"),
							Permissions:         perms.FULL_ACCESS,
							Activated:           true,
							NeedsPasswordChange: false,
						},
					},
				},
			},
			args: args{
				body: params.SetUserPasswordParams{
					UserID:      2,
					NewPassword: "changedpassword",
					UserMeta: &params.UserMeta{
						UserID:      1,
						Permissions: perms.DEFAULT_PERMS,
					},
				},
			},
			wantUser: nil,
			wantRes: &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    "You do not have permission to set a new password for this user. This incident was recorded.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := mock.NewMockUserRepository(tt.fields.mockUsers)
			mockLogger, _ := log.NewLogger(true, false)
			userService := NewUserService(mockUserRepo, mockLogger)

			updatedUser, res := userService.SetUserPassword(tt.args.body)

			assert.True(t, tt.wantRes.Equals(res), "tt.wantRes = %v and res = %v should be equal", tt.wantRes, res)

			if tt.wantRes.Success {
				// Make sure password was changed
				assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte(tt.args.body.NewPassword)))
			}
		})
	}
}
