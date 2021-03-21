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
