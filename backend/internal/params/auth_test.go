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

package params

import (
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLoginParams_Validate(t *testing.T) {
	type fields struct {
		Username string
		Password string
	}
	tests := []struct {
		name      string
		fields    fields
		wantValid bool
	}{
		{
			name: "params.auth.1",
			fields: fields{
				Username: "username",
				Password: "password",
			},
			wantValid: true,
		},
		{
			name: "params.auth.2",
			fields: fields{
				Username: strings.Repeat("a", config.UsernameMinLen),
				Password: strings.Repeat("p", config.PasswordMinLen),
			},
			wantValid: true,
		},
		{
			name: "params.auth.3",
			fields: fields{
				Username: strings.Repeat("a", config.UsernameMaxLen),
				Password: strings.Repeat("p", config.PasswordMaxLen),
			},
			wantValid: true,
		},
		{
			name: "params.auth.4",
			fields: fields{
				Username: strings.Repeat("a", config.UsernameMinLen-1),
				Password: strings.Repeat("p", config.PasswordMinLen-1),
			},
			wantValid: false,
		},
		{
			name: "params.auth.5",
			fields: fields{
				Username: strings.Repeat("a", config.UsernameMaxLen+1),
				Password: strings.Repeat("p", config.PasswordMaxLen+1),
			},
			wantValid: false,
		},
		{
			name: "params.auth.6",
			fields: fields{
				Username: strings.Repeat("a", config.UsernameMaxLen+1),
				Password: strings.Repeat("p", config.PasswordMinLen-1),
			},
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := &LoginParams{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}

			gotValid, _ := body.Validate()
			assert.Equal(t, gotValid, tt.wantValid, "gotValid and wantValid should be equal")
		})
	}
}
