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
)

const (
	SUPER_ADMIN            = int64(0b0100000000000000000000000000000000000000000000000000000000000000)
	FULL_ACCESS            = int64(0b0010000000000000000000000000000000000000000000000000000000000000)
	LOG_WARNING            = int64(0b0001000000000000000000000000000000000000000000000000000000000000)
	LOG_MUTE               = int64(0b0000100000000000000000000000000000000000000000000000000000000000)
	LOG_KICK               = int64(0b0000010000000000000000000000000000000000000000000000000000000000)
	LOG_BAN                = int64(0b0000001000000000000000000000000000000000000000000000000000000000)
	EDIT_OWN_INFRACTIONS   = int64(0b0000000100000000000000000000000000000000000000000000000000000000)
	EDIT_ANY_INFRACTION    = int64(0b0000000010000000000000000000000000000000000000000000000000000000)
	DELETE_OWN_INFRACTIONS = int64(0b0000000001000000000000000000000000000000000000000000000000000000)
	DELETE_ANY_INFRACTION  = int64(0b0000000000100000000000000000000000000000000000000000000000000000)
	VIEW_CHAT_RECORDS      = int64(0b0000000000010000000000000000000000000000000000000000000000000000)

	DEFAULT_PERMS = LOG_WARNING | LOG_MUTE | LOG_KICK | LOG_BAN | EDIT_OWN_INFRACTIONS | VIEW_CHAT_RECORDS // 2238289014803136512
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

	if UserIsAdmin(user1Perms) && (!UserIsAdmin(user2Perms) && !UserIsSuperAdmin(user2Perms)) {
		return true
	}

	return false
}
