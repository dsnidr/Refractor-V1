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

import axios from 'axios';

const postHeaders = {
	'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
};

export function createWarning(data) {
	return axios.post('/api/v1/infractions/warning', data, postHeaders);
}

export function createMute(data) {
	return axios.post('/api/v1/infractions/mute', data, postHeaders);
}

export function createKick(data) {
	return axios.post('/api/v1/infractions/kick', data, postHeaders);
}

export function createBan(data) {
	return axios.post('/api/v1/infractions/ban', data, postHeaders);
}

export function updateInfraction(infractionId, data) {
	return axios.patch(
		`/api/v1/infractions/${infractionId}`,
		data,
		postHeaders
	);
}

export function deleteInfraction(infractionId, data) {
	return axios.delete(`/api/v1/infractions/${infractionId}`);
}

export function searchInfractions(data) {
	return axios.post('/api/v1/search/infractions', data, postHeaders);
}

export function getRecentInfractions() {
	return axios.get('/api/v1/infractions/recent');
}
