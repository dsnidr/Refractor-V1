import { GET_USER_INFO, SET_CURRENT_USER } from './constants';

export const LOG_IN = 'LOG_IN';
export const logIn = (credentials) => ({
	type: LOG_IN,
	payload: credentials,
});

export const setUser = (data) => ({
	type: SET_CURRENT_USER,
	payload: data,
});

export const getUserInfo = () => ({
	type: GET_USER_INFO,
	payload: null,
});

export const CHANGE_USER_PASSWORD = 'CHANGE_USER_PASSWORD';
export const changePassword = (data) => ({
	type: CHANGE_USER_PASSWORD,
	payload: data,
});