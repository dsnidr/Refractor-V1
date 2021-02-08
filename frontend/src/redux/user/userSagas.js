import { changeUserPassword, getUserInfo, logInUser } from '../../api/authApi';
import { setToken } from '../../utils/tokenUtils';
import { all, call, put, takeLatest } from 'redux-saga/effects';
import {
	CHANGE_USER_PASSWORD,
	changePassword,
	LOG_IN,
	setUser,
} from './userActions';
import { GET_USER_INFO } from './constants';
import { setErrors } from '../error/errorActions';
import { setSuccess } from '../success/successActions';

function* logInAsync(action) {
	try {
		const { data } = yield call(logInUser, action.payload);

		yield setToken(data.payload);
		yield getUserInfoAsync();
		yield put(setSuccess('auth', 'Successfully logged in'));
	} catch (err) {
		yield put(setErrors('auth', err.response.data.message));
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
		yield getUserInfoAsync();
	} catch (err) {
		const { data } = err.response;

		let errors;
		if (data.message) {
			errors = data.message;
		} else if (data.errors) {
			errors = data.errors;
		}

		yield put(setErrors('changepassword', errors));
	}
}

function* watchLogIn() {
	yield takeLatest(LOG_IN, logInAsync);
}

function* watchGetUserInfo() {
	yield takeLatest(GET_USER_INFO, getUserInfoAsync);
}

function* watchChangeUserPassword() {
	yield takeLatest(CHANGE_USER_PASSWORD, changeUserPasswordAsync);
}

export default function* userSagas() {
	yield all([
		call(watchLogIn),
		call(watchGetUserInfo),
		call(watchChangeUserPassword),
	]);
}
