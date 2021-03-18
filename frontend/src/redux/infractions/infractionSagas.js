import { all, call, put, takeLatest } from 'redux-saga/effects';
import {
	CREATE_BAN,
	CREATE_KICK,
	CREATE_MUTE,
	CREATE_WARNING,
} from './infractionActions';
import { setErrors } from '../error/errorActions';
import { setSuccess } from '../success/successActions';
import { createWarning } from '../../api/infractionApi';

function* createWarningAsync(action) {
	try {
		const { data } = yield call(createWarning, {
			playerId: action.playerId,
			serverId: action.serverId,
			...action.payload,
		});

		console.log('Warning created', data);

		yield put(setSuccess('createwarning', 'Warning logged'));
		yield put(setErrors('createwarning', undefined));
	} catch (err) {
		console.log('Could not create warning', err);
		const { data } = err.response;

		yield put(setSuccess('createwarning', undefined));
		yield put(
			setErrors(
				'createwarning',
				!data.errors ? data.message : data.errors
			)
		);
	}
}

function* createMuteAsync(action) {}

function* createKickAsync(action) {}

function* createBanAsync(action) {}

function* watchCreateWarning() {
	yield takeLatest(CREATE_WARNING, createWarningAsync);
}

function* watchCreateMute() {
	yield takeLatest(CREATE_MUTE, createMuteAsync);
}

function* watchCreateKick() {
	yield takeLatest(CREATE_KICK, createKickAsync);
}

function* watchCreateBan() {
	yield takeLatest(CREATE_BAN, createBanAsync);
}

export default function* infractionSagas() {
	yield all([
		call(watchCreateWarning),
		call(watchCreateMute),
		call(watchCreateKick),
		call(watchCreateBan),
	]);
}
