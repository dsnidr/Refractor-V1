import { CLEAR_ALL_ERRORS, CLEAR_ERRORS, SET_ERRORS } from './errorActions';

const initialState = {};

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_ERRORS:
			return {
				...state,
				[action.field]: action.payload,
			};
		case CLEAR_ERRORS:
			return {
				...state,
				[action.field]: null,
			};
		case CLEAR_ALL_ERRORS:
			return {};
		default:
			return state;
	}
};

export default reducer;
