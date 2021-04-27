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

export function logInUser(data) {
	return axios.post('/api/v1/auth/login', data, postHeaders);
}

export function refreshToken() {
	return axios.post('/api/v1/auth/refresh');
}

export function getUserInfo() {
	return axios.get('/api/v1/users/me');
}

export function changeUserPassword(data) {
	return axios.post('/api/v1/users/changepassword', data, postHeaders);
}

export function getAllUsers() {
	return axios.get('/api/v1/users/all');
}

export function createUser(data) {
	return axios.post('/api/v1/users/', data, postHeaders);
}

export function activateUser(userId) {
	return axios.patch(`/api/v1/users/activate/${userId}`);
}

export function deactivateUser(userId) {
	return axios.patch(`/api/v1/users/deactivate/${userId}`);
}

export function setUserPassword(data) {
	return axios.patch(`/api/v1/users/setpassword`, data, postHeaders);
}

export function forceUserPasswordChange(userId) {
	return axios.patch(`/api/v1/users/forcepasswordchange/${userId}`);
}

export function setUserPermissions(data) {
	return axios.patch('/api/v1/users/setpermissions', data, postHeaders);
}
