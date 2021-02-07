export const LOG_IN = 'LOG_IN';
export const logIn = (credentials) => ({
	type: LOG_IN,
	payload: credentials,
});

export const SET_USER = 'SET_USER';
export const setUser = (data) => ({
	type: SET_USER,
	payload: data,
});
