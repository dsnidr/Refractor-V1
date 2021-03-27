import { SET_INFRACTION_SEARCH_RESULTS } from './infractionActions';

const initialState = {
	searchResults: [],
};

const infractionReducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_INFRACTION_SEARCH_RESULTS:
			return {
				...state,
				searchResults: action.payload,
			};
		default:
			return state;
	}
};

export default infractionReducer;
