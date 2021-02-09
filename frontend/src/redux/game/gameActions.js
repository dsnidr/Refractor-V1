export const SET_GAMES = 'SET_GAMES';
export const setGames = (games) => ({
	type: SET_GAMES,
	payload: games,
});

export const GET_GAMES = 'GET_GAMES';
export const getGames = () => ({
	type: GET_GAMES,
});
