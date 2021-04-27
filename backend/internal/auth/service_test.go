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
							Permissions:         0,
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
							Permissions:         0,
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
							Permissions:         0,
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
