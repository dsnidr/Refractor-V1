import { SET_LOADING } from './loadingActions';

const initialState = {
	main: false,
	login: false,
	users: false,
};

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_LOADING:
			state[action.scope] = action.payload;
			return state;
		default:
			return state;
	}
};

export default reducer;
