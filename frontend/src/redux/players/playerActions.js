export const GET_PLAYER_SUMMARY = 'GET_PLAYER_SUMMARY';
export const getPlayerSummary = (playerId) => ({
	type: GET_PLAYER_SUMMARY,
	playerId: playerId,
});

export const SET_CURRENT_PLAYER = 'SET_CURRENT_PLAYER';
export const setCurrentPlayer = (playerData) => ({
	type: SET_CURRENT_PLAYER,
	payload: playerData,
});
