import { getUserInfo, logInUser } from '../../api/authApi';
import { setToken } from '../../utils/tokenUtils';
import { all, call, put, takeLatest } from 'redux-saga/effects';
import { LOG_IN, setUser } from './userActions';
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

function* watchLogIn() {
	yield takeLatest(LOG_IN, logInAsync);
}

function* watchGetUserInfo() {
	yield takeLatest(GET_USER_INFO, getUserInfoAsync);
}

export default function* userSagas() {
	yield all([call(watchLogIn), call(watchGetUserInfo)]);
}
