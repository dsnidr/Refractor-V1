const initialState = null;

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_GAMES:
			return action.payload;
		default:
			return state;
	}
};

export default reducer;
