import {
	SET_CURRENT_USER,
	SET_USER_ACTIVATED,
	SET_USER_DEACTIVATED,
} from './constants';
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
		case SET_USER_ACTIVATED:
			return {
				...state,
				others: {
					...state.others,
					[action.userId]: {
						...state.others[action.userId],
						activated: true,
					},
				},
			};
		case SET_USER_DEACTIVATED:
			return {
				...state,
				others: {
					...state.others,
					[action.userId]: {
						...state.others[action.userId],
						activated: false,
					},
				},
			};
		default:
			return state;
	}
};

export default reducer;
