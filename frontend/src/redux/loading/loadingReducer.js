import { SET_LOADING } from './loadingActions';

const initialState = false;

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_LOADING:
			return action.payload;
		default:
			return state;
	}
};

export default reducer;
