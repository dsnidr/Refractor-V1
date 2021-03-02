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
