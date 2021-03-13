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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasHigherAccess(tt.args.user1Perms, tt.args.user2Perms)

			assert.Equal(t, tt.want, got)
		})
	}
}
