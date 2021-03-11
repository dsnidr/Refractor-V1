import {
	activateUser,
	changeUserPassword,
	createUser,
	deactivateUser,
	getAllUsers,
	getUserInfo,
	logInUser,
} from '../../api/userApi';
import { setToken } from '../../utils/tokenUtils';
import { all, call, put, takeLatest } from 'redux-saga/effects';
import {
	ACTIVATE_USER,
	ADD_USER,
	CHANGE_USER_PASSWORD,
	changePassword,
	DEACTIVATE_USER,
	GET_ALL_USERS,
	LOG_IN,
	SET_USER_PASSWORD,
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

function* setUserPasswordAsync(action) {}

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

function* watchActivateUser() {
	yield takeLatest(ACTIVATE_USER, activateUserAsync);
}

function* watchDeactivateUser() {
	yield takeLatest(DEACTIVATE_USER, deactivateUserAsync);
}

function* watchSetUserPassword() {
	yield takeLatest(SET_USER_PASSWORD, setUserPasswordAsync);
}

export default function* userSagas() {
	yield all([
		call(watchLogIn),
		call(watchGetUserInfo),
		call(watchChangeUserPassword),
		call(watchGetAllUsers),
		call(watchAddUser),
		call(watchActivateUser),
		call(watchDeactivateUser),
		call(watchSetUserPassword),
	]);
}
