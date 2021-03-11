package perms

import "github.com/sniddunc/bitperms"

const (
	SUPER_ADMIN            = uint64(0b1000000000000000000000000000000000000000000000000000000000000000)
	FULL_ACCESS            = uint64(0b0100000000000000000000000000000000000000000000000000000000000000)
	LOG_WARNING            = uint64(0b0010000000000000000000000000000000000000000000000000000000000000)
	LOG_MUTE               = uint64(0b0001000000000000000000000000000000000000000000000000000000000000)
	LOG_KICK               = uint64(0b0000100000000000000000000000000000000000000000000000000000000000)
	LOG_BAN                = uint64(0b0000010000000000000000000000000000000000000000000000000000000000)
	EDIT_OWN_INFRACTIONS   = uint64(0b0000001000000000000000000000000000000000000000000000000000000000)
	EDIT_ANY_INFRACTION    = uint64(0b0000000100000000000000000000000000000000000000000000000000000000)
	DELETE_OWN_INFRACTIONS = uint64(0b0000000010000000000000000000000000000000000000000000000000000000)
	DELETE_ANY_INFRACTION  = uint64(0b0000000001000000000000000000000000000000000000000000000000000000)

	DEFAULT_PERMS = LOG_WARNING | LOG_MUTE | LOG_KICK | LOG_BAN | EDIT_OWN_INFRACTIONS
)

// UserIsSuperAdmin returns true if the user has the super admin flag set.
func UserIsSuperAdmin(userPerms bitperms.PermissionValue) bool {
	return userPerms.HasFlag(SUPER_ADMIN)
}

// UserIsAdmin returns true if the user has the full access flag set.
func UserIsAdmin(userPerms bitperms.PermissionValue) bool {
	return userPerms.HasFlag(FULL_ACCESS)
}

// UserHasFullAccess returns true if the user has the super admin or the full access flag set.
func UserHasFullAccess(userPerms bitperms.PermissionValue) bool {
	return userPerms.HasFlag(SUPER_ADMIN) || userPerms.HasFlag(FULL_ACCESS)
}

// HasHigherAccess returns true if user1 has a higher access level than user2.
// user1 is determined to have a higher access level if:
//
//  a) user1 is an admin and user2 is not or
//
//  b) user1 is a super admin and user2 is not
func HasHigherAccess(user1Perms bitperms.PermissionValue, user2Perms bitperms.PermissionValue) bool {
	if UserIsSuperAdmin(user1Perms) && !UserIsSuperAdmin(user2Perms) {
		return true
	}

	if UserIsAdmin(user1Perms) && !UserIsAdmin(user2Perms) {
		return true
	}

	return false
}
