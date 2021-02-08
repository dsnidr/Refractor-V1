export const SET_SUCCESS = 'SET_SUCCESS';
export const setSuccess = (field, message) => ({
	type: SET_SUCCESS,
	field: field,
	payload: message,
});

export const CLEAR_SUCCESS = 'CLEAR_SUCCESS';
export const clearSuccess = (field) => ({
	type: CLEAR_SUCCESS,
	field: field,
});

export const CLEAR_ALL_SUCCESS = 'CLEAR_ALL_SUCCESS';
export const clearAllSuccess = () => ({
	type: CLEAR_ALL_SUCCESS,
});
