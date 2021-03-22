import {
	SET_CURRENT_PLAYER,
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
		default:
			return state;
	}
};

export default playerReducer;
