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

export function getAllServerData() {
	return axios.get(`/api/v1/servers/data`);
}

export function updateServer(serverId, data) {
	return axios.patch(`/api/v1/servers/${serverId}`, data, postHeaders);
}

export function createServer(data) {
	return axios.post(`/api/v1/servers/`, data, postHeaders);
}

export function deleteServer(serverId) {
	return axios.delete(`/api/v1/servers/${serverId}`);
}
