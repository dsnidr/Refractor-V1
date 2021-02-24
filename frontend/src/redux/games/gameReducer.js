import { SET_SERVERS } from '../servers/serverActions';

const initialState = null;

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_SERVERS:
			return action.payload;
		default:
			return state;
	}
};

export default reducer;
