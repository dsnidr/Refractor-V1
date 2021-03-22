import axios from 'axios';

const postHeaders = {
	'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
};

export function getPlayerSummary(playerId) {
	return axios.get(`/api/v1/players/summary/${playerId}`);
}

export function searchPlayers(data) {
	return axios.post('/api/v1/search/players', data, postHeaders);
}
