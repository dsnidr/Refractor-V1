import { SET_THEME } from './constants';

const initialState = 'dark';

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_THEME:
			localStorage.setItem('theme', action.payload);
			return action.payload;
		default:
			return state;
	}
};

export default reducer;
