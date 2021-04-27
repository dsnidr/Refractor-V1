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

package validation

import (
	"strings"
	"testing"
)

func TestIsEmailValid(t *testing.T) {
	type args struct {
		email string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "validation.email.1",
			args: args{
				email: strings.Repeat("a", emailMinLen-2) + "@" + strings.Repeat("d", emailMinLen-2),
			},
			want: true,
		},
		{
			name: "validation.email.2",
			args: args{
				email: strings.Repeat("a", emailMaxLen-5) + "@a.co",
			},
			want: true,
		},
		{
			name: "validation.email.3",
			args: args{
				email: strings.Repeat("a", emailMaxLen-5+1) + "@a.co",
			},
			want: false,
		},
		{
			name: "validation.email.4",
			args: args{
				email: strings.Repeat("a", emailMinLen) + "@",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmailValid(tt.args.email); got != tt.want {
				t.Errorf("IsEmailValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
