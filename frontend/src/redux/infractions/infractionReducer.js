import {
	SET_INFRACTION_SEARCH_RESULTS,
	SET_RECENT_INFRACTIONS,
} from './infractionActions';

const initialState = {
	searchResults: [],
	recentInfractions: [],
};

const infractionReducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_INFRACTION_SEARCH_RESULTS:
			return {
				...state,
				searchResults: action.payload,
			};
		case SET_RECENT_INFRACTIONS:
			return {
				...state,
				recentInfractions: action.payload,
			};
		default:
			return state;
	}
};

export default infractionReducer;
