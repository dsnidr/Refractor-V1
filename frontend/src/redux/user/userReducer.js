import { SET_CURRENT_USER } from './constants';
import { SET_ALL_USERS } from './userActions';

const initialState = {
	self: null,
	others: null,
};

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_CURRENT_USER:
			return {
				...state,
				self: action.payload,
			};
		case SET_ALL_USERS:
			return {
				...state,
				others: action.payload,
			};
		default:
			return state;
	}
};

export default reducer;
