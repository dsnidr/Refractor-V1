import axios from 'axios';

export function getPlayerSummary(playerId) {
	return axios.get(`/api/v1/players/summary/${playerId}`);
}
