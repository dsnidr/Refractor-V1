import axios from 'axios';

export function getAllServerData() {
	return axios.get(`/api/v1/servers/data`);
}
