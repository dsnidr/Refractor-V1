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

export function timestampToDateTime(unixTimestamp) {
	const date = new Date(unixTimestamp * 1000);

	return `${date.toLocaleString('en-GB', { hour12: true })}`;
}

const SECONDS_IN_DAY = 86400;
const SECONDS_IN_HOUR = 3600;
const SECONDS_IN_MINUTE = 60;
const HOURS_IN_DAY = 24;

export function getTimeRemaining(unixTimestamp, duration) {
	const endDate = new Date(unixTimestamp * 1000 + duration * 60 * 1000);
	const currentDate = new Date();

	return (endDate.getTime() - currentDate.getTime()) / 1000;
}

export function buildTimeRemainingString(timeRemaining) {
	if (timeRemaining < 0) {
		return 'EXPIRED';
	}

	let delta = timeRemaining;

	// determine days left
	let days = Math.floor(delta / SECONDS_IN_DAY);
	delta -= days * SECONDS_IN_DAY;

	// determine hours left
	let hours = Math.floor(delta / SECONDS_IN_HOUR) % HOURS_IN_DAY;
	delta -= hours * SECONDS_IN_HOUR;

	// determine minutes left
	let minutes = Math.floor(delta / SECONDS_IN_MINUTE) % SECONDS_IN_MINUTE;
	delta -= minutes * SECONDS_IN_MINUTE;

	// the remainder is seconds
	delta -= minutes * 60;

	let seconds = Math.floor(delta % 60);

	// Build output string
	let output = '';

	if (days > 0) {
		output += `${days} days, `;
	}

	if (hours > 0) {
		output += `${hours} hours, `;
	}

	if (minutes > 0) {
		output += `${minutes} minutes, `;
	}

	if (seconds > 0) {
		output += `${seconds} seconds, `;
	}

	// Trim trailing space and comma
	output = output.substr(0, output.length - 2);

	return output;
}
