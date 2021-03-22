import {
	SET_CURRENT_PLAYER,
	SET_PLAYER_WATCHED,
	SET_RECENT_PLAYERS,
	SET_SEARCH_RESULTS,
} from './playerActions';

const initialState = {
	currentPlayer: null,
	searchResults: [],
	recentPlayers: [],
};

const playerReducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_CURRENT_PLAYER:
			return {
				...state,
				currentPlayer: action.payload,
			};
		case SET_SEARCH_RESULTS:
			return {
				...state,
				searchResults: action.payload,
			};
		case SET_RECENT_PLAYERS:
			return {
				...state,
				recentPlayers: action.payload,
			};
		case SET_PLAYER_WATCHED:
			return setPlayerWatched(state, action.playerId, action.payload);
		default:
			return state;
	}
};

function setPlayerWatched(state, playerId, watched) {
	console.log(state, playerId, watched);

	console.log(state.currentPlayer.id === playerId);

	if (state.currentPlayer && state.currentPlayer.id === playerId) {
		return {
			...state,
			currentPlayer: {
				...state.currentPlayer,
				watched: watched,
			},
		};
	}

	return state;
}

export default playerReducer;
