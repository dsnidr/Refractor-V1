import {
	CLEAR_ALL_SUCCESS,
	CLEAR_SUCCESS,
	SET_SUCCESS,
} from './successActions';

const initialState = {};

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_SUCCESS:
			return {
				...state,
				[action.field]: action.payload,
			};
		case CLEAR_SUCCESS:
			return {
				...state,
				[action.field]: null,
			};
		case CLEAR_ALL_SUCCESS:
			return {};
		default:
			return state;
	}
};

export default reducer;
