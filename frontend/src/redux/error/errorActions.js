export const SET_ERRORS = 'SET_ERRORS';
export const setErrors = (field, errors) => ({
	type: SET_ERRORS,
	field: field,
	payload: errors,
});

export const CLEAR_ERRORS = 'CLEAR_ERRORS';
export const clearErrors = (field) => ({
	type: CLEAR_ERRORS,
	field: field,
});

export const CLEAR_ALL_ERRORS = 'CLEAR_ALL_ERRORS';
export const clearAllErrors = () => ({
	type: CLEAR_ALL_ERRORS,
});
