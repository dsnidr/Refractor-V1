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

export function reasonIsValid(reason) {
	return !(
		!reason ||
		typeof reason !== 'string' ||
		reason.trim().length === 0
	);
}

export function typeHasDuration(type) {
	return type === 'MUTE' || type === 'BAN';
}

export function buildInfractionTitle(infraction) {
	return `${infraction.type}`;
}

export function buildInfractionText(infraction) {
	const nameSection = `${infraction.staffName} has ${getInfractionVerb(
		infraction.type
	)} player ${infraction.playerName}`;
	const reasonSection = `for: ${infraction.reason}`;

	return `${nameSection} ${reasonSection}`;
}

function getInfractionVerb(infractionType) {
	switch (infractionType) {
		case 'WARNING':
			return 'warned';
		case 'KICK':
			return 'kicked';
		case 'MUTE':
			return 'muted';
		case 'BAN':
			return 'banned';
	}
}
