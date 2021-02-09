import axios from 'axios';

export function getAllGames() {
	return axios.get('/api/v1/games/');
}
