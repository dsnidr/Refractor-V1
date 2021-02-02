package params

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"

	"github.com/sniddunc/refractor/pkg/config"
)

func TestCreateUserParams_Validate(t *testing.T) {
	type fields struct {
		Email           string
		Username        string
		Password        string
		PasswordConfirm string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "params.user.create.1",
			fields: fields{
				Email:           "test@test.com",
				Username:        "username",
				Password:        "password",
				PasswordConfirm: "password",
			},
			want: true,
		},
		{
			name: "params.user.create.2",
			fields: fields{
				Email:           "test@test.com",
				Username:        strings.Repeat("u", config.UsernameMinLen),
				Password:        strings.Repeat("p", config.PasswordMinLen),
				PasswordConfirm: strings.Repeat("p", config.PasswordMinLen),
			},
			want: true,
		},
		{
			name: "params.user.create.3",
			fields: fields{
				Email:           "test@test.com",
				Username:        strings.Repeat("u", config.UsernameMaxLen),
				Password:        strings.Repeat("p", config.PasswordMaxLen),
				PasswordConfirm: strings.Repeat("p", config.PasswordMaxLen),
			},
			want: true,
		},
		{
			name: "params.user.create.4",
			fields: fields{
				Email:           "test@test.com",
				Username:        strings.Repeat("u", config.UsernameMaxLen),
				Password:        strings.Repeat("p", config.PasswordMaxLen),
				PasswordConfirm: "doesntmatch",
			},
			want: false,
		},
		{
			name: "params.user.create.5",
			fields: fields{
				Email:           "test@test.com",
				Username:        strings.Repeat("u", config.UsernameMaxLen+1),
				Password:        strings.Repeat("p", config.PasswordMaxLen-1),
				PasswordConfirm: "doesntmatch",
			},
			want: false,
		},
		{
			name: "params.user.create.5",
			fields: fields{
				Email:           "test@test.com",
				Username:        strings.Repeat("u", config.UsernameMaxLen-1),
				Password:        strings.Repeat("p", config.PasswordMaxLen+1),
				PasswordConfirm: "doesntmatch",
			},
			want: false,
		},
		{
			name: "params.user.create.6",
			fields: fields{
				Email:           "invalidemail",
				Username:        "testusername",
				Password:        "testpass",
				PasswordConfirm: "testpass",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &CreateUserParams{
				Email:           tt.fields.Email,
				Username:        tt.fields.Username,
				Password:        tt.fields.Password,
				PasswordConfirm: tt.fields.PasswordConfirm,
			}
			got, errors := body.Validate()
			if got != tt.want {
				t.Errorf("CreateUserParams.Validate() got = %v, want %v, errors %v", got, tt.want, errors)
			}
		})
	}
}

func TestSetUserAccessLevelParams_Validate(t *testing.T) {
	type fields struct {
		UserID       int64
		AccessLevel  int
		SetterUserID int64
		UserMeta     *UserMeta
	}
	tests := []struct {
		name      string
		fields    fields
		wantValid bool
	}{
		{
			name: "params.user.setaccesslevel.1",
			fields: fields{
				UserID:      1,
				AccessLevel: config.AL_USER,
				UserMeta: &UserMeta{
					UserID:      2,
					AccessLevel: config.AL_ADMIN,
				},
			},
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &SetUserAccessLevelParams{
				UserID:      tt.fields.UserID,
				AccessLevel: tt.fields.AccessLevel,
				UserMeta:    tt.fields.UserMeta,
			}
			valid, _ := body.Validate()

			assert.Equal(t, valid, tt.wantValid, "valid = %v and tt.wantValid = %v should be equal", valid, tt.wantValid)
		})
	}
}
