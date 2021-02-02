package user

import (
	"github.com/sniddunc/refractor/internal/mock"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		userRepo refractor.UserRepository
		logger   log.Logger
	}
	tests := []struct {
		name string
		args args
		want refractor.UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.userRepo, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

			assert.ObjectsAreEqual(newUser, tt.wantUser)
			assert.True(t, tt.wantRes.Equals(res), "tt.wantRes = %v and res = %v should be equal", tt.wantRes, res)
		})
	}
}

//func Test_userService_SetUserAccessLevel(t *testing.T) {
//	type fields struct {
//		repo refractor.UserRepository
//		log  log.Logger
//	}
//	type args struct {
//		body params.SetUserAccessLevelParams
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		wantUser   *refractor.User
//		wantRes  *refractor.ServiceResponse
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &userService{
//				repo: tt.fields.repo,
//				log:  tt.fields.log,
//			}
//			got, got1 := s.SetUserAccessLevel(tt.args.body)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("SetUserAccessLevel() got = %v, want %v", got, tt.want)
//			}
//			if !reflect.DeepEqual(got1, tt.want1) {
//				t.Errorf("SetUserAccessLevel() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
