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
