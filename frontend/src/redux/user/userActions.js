import { GET_USER_INFO, SET_CURRENT_USER } from './constants';

export const LOG_IN = 'LOG_IN';
export const logIn = (credentials) => ({
	type: LOG_IN,
	payload: credentials,
});

export const REFRESH_TOKENS = 'REFRESH_TOKENS';
export const refreshTokens = () => ({
	type: REFRESH_TOKENS,
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

export const GET_ALL_USERS = 'GET_ALL_USERS';
export const getAllUsers = () => ({
	type: GET_ALL_USERS,
});

export const SET_ALL_USERS = 'SET_ALL_USERS';
export const setAllUsers = (users) => ({
	type: SET_ALL_USERS,
	payload: users,
});

export const ADD_USER = 'ADD_USER';
export const addUser = (data) => ({
	type: ADD_USER,
	payload: data,
});

export const ACTIVATE_USER = 'ACTIVATE_USER';
export const activateUser = (userId) => ({
	type: ACTIVATE_USER,
	userId: userId,
});

export const DEACTIVATE_USER = 'DEACTIVATE_USER';
export const deactivateUser = (userId) => ({
	type: DEACTIVATE_USER,
	userId: userId,
});

export const SET_USER_PASSWORD = 'SET_USER_PASSWORD';
export const setUserPassword = (userId, data) => ({
	type: SET_USER_PASSWORD,
	userId: userId,
	payload: data,
});

export const FORCE_USER_PASSWORD_CHANGE = 'FORCE_USER_PASSWORD_CHANGE';
export const forceUserPasswordChange = (userId) => ({
	type: FORCE_USER_PASSWORD_CHANGE,
	userId: userId,
});

export const SET_USER_PERMISSIONS = 'SET_USER_PERMISSIONS';
export const setUserPermissions = (userId, permissions) => ({
	type: SET_USER_PERMISSIONS,
	userId: userId,
	payload: permissions,
});
