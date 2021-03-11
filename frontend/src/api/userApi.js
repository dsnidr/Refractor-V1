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
