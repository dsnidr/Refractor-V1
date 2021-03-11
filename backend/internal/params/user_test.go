package params

import (
	"github.com/sniddunc/refractor/pkg/perms"
	"github.com/stretchr/testify/assert"
	"net/url"
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

func TestSetUserPermissionsParams_Validate(t *testing.T) {
	type fields struct {
		UserID       int64
		Permissions  uint64
		SetterUserID int64
		UserMeta     *UserMeta
	}
	tests := []struct {
		name      string
		fields    fields
		wantValid bool
	}{
		{
			name: "params.user.setpermissions.1",
			fields: fields{
				UserID:      1,
				Permissions: 0,
				UserMeta: &UserMeta{
					UserID:      2,
					Permissions: perms.FULL_ACCESS,
				},
			},
			wantValid: true,
		},
		{
			name: "params.user.setpermissions.2",
			fields: fields{
				UserID:      -1,
				Permissions: 0,
				UserMeta: &UserMeta{
					UserID:      2,
					Permissions: 0,
				},
			},
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &SetUserPermissionsParams{
				UserID:      tt.fields.UserID,
				Permissions: tt.fields.Permissions,
				UserMeta:    tt.fields.UserMeta,
			}
			valid, errors := body.Validate()

			assert.Equal(t, valid, tt.wantValid, "valid = %v and tt.wantValid = %v should be equal.\nErrors: %v", valid, tt.wantValid, errors)
		})
	}
}

func TestChangeUserPassword_Validate(t *testing.T) {
	type fields struct {
		CurrentPassword    string
		NewPassword        string
		NewPasswordConfirm string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "params.user.changepwd.1",
			fields: fields{
				CurrentPassword:    "currentPass (not tested)",
				NewPassword:        "password",
				NewPasswordConfirm: "password",
			},
			want: true,
		},
		{
			name: "params.user.changepwd.2",
			fields: fields{
				CurrentPassword:    "currentPass (not tested)",
				NewPassword:        strings.Repeat("p", config.PasswordMaxLen),
				NewPasswordConfirm: strings.Repeat("p", config.PasswordMaxLen),
			},
			want: true,
		},
		{
			name: "params.user.changepwd.3",
			fields: fields{
				CurrentPassword:    "currentPass (not tested)",
				NewPassword:        strings.Repeat("p", config.PasswordMinLen),
				NewPasswordConfirm: strings.Repeat("p", config.PasswordMinLen),
			},
			want: true,
		},
		{
			name: "params.user.changepwd.4",
			fields: fields{
				CurrentPassword:    "currentPass (not tested)",
				NewPassword:        strings.Repeat("p", config.PasswordMaxLen+1),
				NewPasswordConfirm: strings.Repeat("p", config.PasswordMaxLen+1),
			},
			want: false,
		},
		{
			name: "params.user.changepwd.5",
			fields: fields{
				CurrentPassword:    "currentPass (not tested)",
				NewPassword:        strings.Repeat("p", config.PasswordMinLen-1),
				NewPasswordConfirm: strings.Repeat("p", config.PasswordMinLen-1),
			},
			want: false,
		},
		{
			name: "params.user.changepwd.6",
			fields: fields{
				CurrentPassword:    "currentPass (not tested)",
				NewPassword:        "passwords",
				NewPasswordConfirm: "do not match",
			},
			want: false,
		},
		{
			name: "params.user.changepwd.7",
			fields: fields{
				CurrentPassword:    "currentPass (not tested)",
				NewPassword:        "2short",
				NewPasswordConfirm: "2short",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &ChangeUserPassword{
				CurrentPassword:    tt.fields.CurrentPassword,
				NewPassword:        tt.fields.NewPassword,
				NewPasswordConfirm: tt.fields.NewPasswordConfirm,
			}

			got, _ := body.Validate()
			assert.Equal(t, tt.want, got, "Validate returned the wrong value")
		})
	}
}

func TestSetUserPasswordParams_Validate(t *testing.T) {
	type fields struct {
		UserID      int64
		NewPassword string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
		want1  url.Values
	}{
		{
			name: "params.user.setuserpassword.1",
			fields: fields{
				UserID:      1,
				NewPassword: "validpassword",
			},
			want: true,
		},
		{
			name: "params.user.setuserpassword.2",
			fields: fields{
				UserID:      999999999,
				NewPassword: strings.Repeat("a", config.PasswordMinLen),
			},
			want: true,
		},
		{
			name: "params.user.setuserpassword.3",
			fields: fields{
				UserID:      0,
				NewPassword: strings.Repeat("a", config.PasswordMaxLen),
			},
			want: false,
		},
		{
			name: "params.user.setuserpassword.4",
			fields: fields{
				UserID:      -1,
				NewPassword: "validpassword",
			},
			want: false,
		},
		{
			name: "params.user.setuserpassword.5",
			fields: fields{
				UserID:      2,
				NewPassword: strings.Repeat("a", config.PasswordMinLen-1),
			},
			want: false,
		},
		{
			name: "params.user.setuserpassword.6",
			fields: fields{
				UserID:      894268726,
				NewPassword: strings.Repeat("a", config.PasswordMaxLen+1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &SetUserPasswordParams{
				UserID:      tt.fields.UserID,
				NewPassword: tt.fields.NewPassword,
			}

			got, _ := body.Validate()
			assert.Equal(t, tt.want, got, "Validate returned the wrong value")
		})
	}
}
