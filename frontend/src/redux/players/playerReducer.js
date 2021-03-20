import { SET_CURRENT_PLAYER } from './playerActions';

const initialState = {
	currentPlayer: null,
};

const playerReducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_CURRENT_PLAYER:
			return {
				...state,
				currentPlayer: action.payload,
			};
		default:
			return state;
	}
};

export default playerReducer;
