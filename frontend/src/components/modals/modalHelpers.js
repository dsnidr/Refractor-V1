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

export function getModalStateFromProps(nextProps, prevState) {
	if (nextProps.success) {
		prevState = {
			...prevState,
			errors: {},
			success: nextProps.success,
		};
	}

	if (nextProps.errors) {
		prevState = {
			...prevState,
			errors: nextProps.errors,
			success: null,
		};
	}

	if (nextProps.player) {
		prevState.player = nextProps.player;
	}

	if (nextProps.serverId) {
		prevState.serverId = nextProps.serverId;
	}

	return prevState;
}
