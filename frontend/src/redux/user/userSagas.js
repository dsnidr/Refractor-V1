/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import {
	activateUser,
	changeUserPassword,
	createUser,
	deactivateUser,
	forceUserPasswordChange,
	getAllUsers,
	getUserInfo,
	logInUser,
	refreshToken,
	setUserPassword,
	setUserPermissions,
} from '../../api/userApi';
import { decodeToken, getToken, setToken } from '../../utils/tokenUtils';
import { all, call, put, takeLatest } from 'redux-saga/effects';
import {
	ACTIVATE_USER,
	ADD_USER,
	CHANGE_USER_PASSWORD,
	changePassword,
	DEACTIVATE_USER,
	FORCE_USER_PASSWORD_CHANGE,
	GET_ALL_USERS,
	LOG_IN,
	REFRESH_TOKENS,
	SET_USER_PASSWORD,
	SET_USER_PERMISSIONS,
	setAllUsers,
	setUser,
} from './userActions';
import {
	GET_USER_INFO,
	SET_USER_ACTIVATED,
	SET_USER_DEACTIVATED,
} from './constants';
import { setErrors } from '../error/errorActions';
import { setSuccess } from '../success/successActions';
import { setLoading } from '../loading/loadingActions';

function* logInAsync(action) {
	try {
		const { data } = yield call(logInUser, action.payload);

		yield setToken(data.payload);
		yield getUserInfoAsync();
		yield put(setSuccess('auth', 'Successfully logged in'));
		yield put(setErrors('auth', undefined));
		yield put(setLoading('login', false));
	} catch (err) {
		yield put(setErrors('auth', err.response.data.message));
		yield put(setSuccess('auth', undefined));
	}
}

function* refreshTokensAsync(action) {
	try {
		const { data } = yield call(refreshToken);

		setToken(data.payload);
	} catch (err) {
		console.log('Could not refresh tokens', err);
	}
}

function* getUserInfoAsync() {
	try {
		const { data } = yield call(getUserInfo);

		yield put(
			setUser({
				isAuthenticated: true,
				...data.payload,
			})
		);
	} catch (err) {
		console.log('Could not get user info', err);
		yield put(
			setUser({
				isAuthenticated: false,
			})
		);
	}
}

function* changeUserPasswordAsync(action) {
	try {
		const { data } = yield call(changeUserPassword, action.payload);

		yield put(setSuccess('changepassword', data.message));
		yield put(setErrors('changepassword', undefined));
		yield getUserInfoAsync();
	} catch (err) {
		const { data } = err.response;

		yield put(setSuccess('changepassword', undefined));
		yield put(
			setErrors(
				'changepassword',
				data.errors ? data.errors : data.message
			)
		);
	}
}

function* getAllUsersAsync() {
	try {
		const { data } = yield call(getAllUsers);

		const allUsers = {};

		data.payload.forEach((user) => {
			allUsers[user.id] = user;
		});

		yield put(setAllUsers(allUsers));
	} catch (err) {
		console.log('Could not get all users', err);
	}
}

function* addUserAsync(action) {
	try {
		const { data } = yield call(createUser, action.payload);

		yield put(setSuccess('adduser', data.message));
		yield put(setErrors('adduser', undefined));
	} catch (err) {
		console.log('Could not create new user', err);
		const { data } = err.response;

		yield put(setSuccess('adduser', undefined));
		yield put(
			setErrors(
				'adduser',
				!data.errors
					? `Could not create new user: ${err.response.data.message}`
					: data.errors
			)
		);
	}
}

function* activateUserAsync(action) {
	try {
		const { data } = yield call(activateUser, action.userId);

		yield put({
			type: SET_USER_ACTIVATED,
			userId: action.userId,
		});

		yield put(setSuccess('usermgmt', data.message));
		yield put(setErrors('usermgmt', undefined));
	} catch (err) {
		console.log('Could not activate user', err);
		const { data } = err.response;

		yield put(setErrors('usermgmt', data.message));
		yield put(setSuccess('usermgmt', undefined));
	}
}

function* deactivateUserAsync(action) {
	try {
		const { data } = yield call(deactivateUser, action.userId);

		yield put({
			type: SET_USER_DEACTIVATED,
			userId: action.userId,
		});

		yield put(setSuccess('usermgmt', data.message));
		yield put(setErrors('usermgmt', undefined));
	} catch (err) {
		console.log('Could not deactivate user', err);
		const { data } = err.response;

		yield put(setErrors('usermgmt', data.message));
		yield put(setSuccess('usermgmt', undefined));
	}
}

function* setUserPasswordAsync(action) {
	try {
		const { data } = yield call(setUserPassword, {
			id: action.userId,
			...action.payload,
		});

		yield put(setSuccess('passwordmgmt', data.message));
		yield put(setErrors('passwordmgmt', undefined));
	} catch (err) {
		console.log('Could not set user password', err);
		const { data } = err.response;

		yield put(setErrors('passwordmgmt', data.message));
		yield put(
			setErrors('passwordmgmt', !data.errors ? data.message : data.errors)
		);
	}
}

function* forceUserPasswordChangeAsync(action) {
	try {
		const { data } = yield call(forceUserPasswordChange, action.userId);

		yield put(
			setSuccess(
				'passwordmgmt',
				'User will be forced to change their password the next time they log in.'
			)
		);
		yield put(setErrors('passwordmgmt', undefined));
	} catch (err) {
		const { data } = err.response;
		console.log('Could not force user password change', err);

		yield put(setErrors('passwordmgmt', data.message));
		yield put(setSuccess('passwordmgmt', undefined));
	}
}

function* setUserPermissionsAsync(action) {
	try {
		const { data } = yield call(setUserPermissions, {
			id: action.userId,
			permissions: action.payload.toString(),
		});

		yield put(setSuccess('setpermissions', data.message));
		yield put(setErrors('setpermissions', undefined));
	} catch (err) {
		console.log('Could not set user permissions', err);
		const { data } = err.response;

		yield put(setSuccess('setpermissions', undefined));
		yield put(
			setErrors(
				'setpermissions',
				!data.errors
					? `Could not set permissions for user: ${err.response.data.message}`
					: data.errors
			)
		);
	}
}

function* watchLogIn() {
	yield takeLatest(LOG_IN, logInAsync);
}

function* watchRefreshTokens() {
	yield takeLatest(REFRESH_TOKENS, refreshTokensAsync);
}

function* watchGetUserInfo() {
	yield takeLatest(GET_USER_INFO, getUserInfoAsync);
}

function* watchChangeUserPassword() {
	yield takeLatest(CHANGE_USER_PASSWORD, changeUserPasswordAsync);
}

function* watchGetAllUsers() {
	yield takeLatest(GET_ALL_USERS, getAllUsersAsync);
}

function* watchAddUser() {
	yield takeLatest(ADD_USER, addUserAsync);
}

function* watchActivateUser() {
	yield takeLatest(ACTIVATE_USER, activateUserAsync);
}

function* watchDeactivateUser() {
	yield takeLatest(DEACTIVATE_USER, deactivateUserAsync);
}

function* watchSetUserPassword() {
	yield takeLatest(SET_USER_PASSWORD, setUserPasswordAsync);
}

function* watchForceUserPasswordChange() {
	yield takeLatest(FORCE_USER_PASSWORD_CHANGE, forceUserPasswordChangeAsync);
}

function* watchSetUserPermissions() {
	yield takeLatest(SET_USER_PERMISSIONS, setUserPermissionsAsync);
}

export default function* userSagas() {
	yield all([
		call(watchLogIn),
		call(watchRefreshTokens),
		call(watchGetUserInfo),
		call(watchChangeUserPassword),
		call(watchGetAllUsers),
		call(watchAddUser),
		call(watchActivateUser),
		call(watchDeactivateUser),
		call(watchSetUserPassword),
		call(watchForceUserPasswordChange),
		call(watchSetUserPermissions),
	]);
}
