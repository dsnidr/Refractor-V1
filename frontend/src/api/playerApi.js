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

export function getRecentPlayers() {
	return axios.get('/api/v1/players/recent');
}

export function watchPlayer(playerID) {
	return axios.post(`/api/v1/players/${playerID}/watch`);
}

export function unwatchPlayer(playerID) {
	return axios.post(`/api/v1/players/${playerID}/unwatch`);
}
