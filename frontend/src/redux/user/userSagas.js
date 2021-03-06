import {
	changeUserPassword,
	createUser,
	getAllUsers,
	getUserInfo,
	logInUser,
} from '../../api/userApi';
import { setToken } from '../../utils/tokenUtils';
import { all, call, put, takeLatest } from 'redux-saga/effects';
import {
	ADD_USER,
	CHANGE_USER_PASSWORD,
	changePassword,
	GET_ALL_USERS,
	LOG_IN,
	setAllUsers,
	setUser,
} from './userActions';
import { GET_USER_INFO } from './constants';
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

function* watchLogIn() {
	yield takeLatest(LOG_IN, logInAsync);
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

export default function* userSagas() {
	yield all([
		call(watchLogIn),
		call(watchGetUserInfo),
		call(watchChangeUserPassword),
		call(watchGetAllUsers),
		call(watchAddUser),
	]);
}
