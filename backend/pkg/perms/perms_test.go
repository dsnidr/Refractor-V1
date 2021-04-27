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

package perms

import (
	"github.com/sniddunc/bitperms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasHigherAccess(t *testing.T) {
	type args struct {
		user1Perms bitperms.PermissionValue
		user2Perms bitperms.PermissionValue
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "perms.hashigheraccess.1",
			args: args{
				user1Perms: bitperms.PermissionValue(0),
				user2Perms: bitperms.PermissionValue(0),
			},
			want: false,
		},
		{
			name: "perms.hashigheraccess.2",
			args: args{
				user1Perms: bitperms.PermissionValue(SUPER_ADMIN),
				user2Perms: bitperms.PermissionValue(SUPER_ADMIN),
			},
			want: false,
		},
		{
			name: "perms.hashigheraccess.3",
			args: args{
				user1Perms: bitperms.PermissionValue(SUPER_ADMIN),
				user2Perms: bitperms.PermissionValue(FULL_ACCESS),
			},
			want: true,
		},
		{
			name: "perms.hashigheraccess.4",
			args: args{
				user1Perms: bitperms.PermissionValue(SUPER_ADMIN),
				user2Perms: bitperms.PermissionValue(LOG_WARNING | LOG_KICK | LOG_BAN),
			},
			want: true,
		},
		{
			name: "perms.hashigheraccess.5",
			args: args{
				user1Perms: bitperms.PermissionValue(FULL_ACCESS),
				user2Perms: bitperms.PermissionValue(DEFAULT_PERMS),
			},
			want: true,
		},
		{
			name: "perms.hashigheraccess.6",
			args: args{
				user1Perms: bitperms.PermissionValue(FULL_ACCESS),
				user2Perms: bitperms.PermissionValue(DEFAULT_PERMS),
			},
			want: true,
		},
		{
			name: "perms.hashigheraccess.7",
			args: args{
				user1Perms: bitperms.PermissionValue(DEFAULT_PERMS),
				user2Perms: bitperms.PermissionValue(LOG_WARNING | LOG_KICK | LOG_BAN),
			},
			want: false,
		},
		{
			name: "perms.hashigheraccess.8",
			args: args{
				user1Perms: bitperms.PermissionValue(FULL_ACCESS),
				user2Perms: bitperms.PermissionValue(SUPER_ADMIN),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasHigherAccess(tt.args.user1Perms, tt.args.user2Perms)

			assert.Equal(t, tt.want, got)
		})
	}
}
