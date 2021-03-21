import { all, call, put, takeLatest } from 'redux-saga/effects';
import {
	CREATE_BAN,
	CREATE_KICK,
	CREATE_MUTE,
	CREATE_WARNING,
	DELETE_INFRACTION,
	UPDATE_INFRACTION,
} from './infractionActions';
import { setErrors } from '../error/errorActions';
import { setSuccess } from '../success/successActions';
import {
	createBan,
	createKick,
	createMute,
	createWarning,
	deleteInfraction,
	updateInfraction,
} from '../../api/infractionApi';

function* createWarningAsync(action) {
	try {
		const { data } = yield call(createWarning, {
			playerId: action.playerId,
			serverId: action.serverId,
			...action.payload,
		});

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

function* createMuteAsync(action) {
	try {
		const { data } = yield call(createMute, {
			playerId: action.playerId,
			serverId: action.serverId,
			...action.payload,
		});

		yield put(setSuccess('createmute', 'Mute logged'));
		yield put(setErrors('createmute', undefined));
	} catch (err) {
		console.log('Could not create mute', err);
		const { data } = err.response;

		yield put(setSuccess('createmute', undefined));
		yield put(
			setErrors('createmute', !data.errors ? data.message : data.errors)
		);
	}
}

function* createKickAsync(action) {
	try {
		const { data } = yield call(createKick, {
			playerId: action.playerId,
			serverId: action.serverId,
			...action.payload,
		});

		yield put(setSuccess('createkick', 'Kick logged'));
		yield put(setErrors('createkick', undefined));
	} catch (err) {
		console.log('Could not create kick', err);
		const { data } = err.response;

		yield put(setSuccess('createkick', undefined));
		yield put(
			setErrors('createkick', !data.errors ? data.message : data.errors)
		);
	}
}

function* createBanAsync(action) {
	try {
		const { data } = yield call(createBan, {
			playerId: action.playerId,
			serverId: action.serverId,
			...action.payload,
		});

		yield put(setSuccess('createban', 'Ban logged'));
		yield put(setErrors('createban', undefined));
	} catch (err) {
		console.log('Could not create ban', err);
		const { data } = err.response;

		yield put(setSuccess('createban', undefined));
		yield put(
			setErrors('createban', !data.errors ? data.message : data.errors)
		);
	}
}

function* updateInfractionAsync(action) {
	try {
		yield call(updateInfraction, action.infractionId, action.payload);

		yield put(
			setSuccess('updateinfraction', 'Infraction has been updated')
		);
		yield put(setErrors('updateinfraction', undefined));
	} catch (err) {
		console.log('Could not update infraction', err);
		const { data } = err.response;

		yield put(setSuccess('updateinfraction', undefined));
		yield put(
			setErrors(
				'updateinfraction',
				!data.errors ? data.message : data.errors
			)
		);
	}
}

function* deleteInfractionAsync(action) {
	try {
		yield call(deleteInfraction, action.infractionId);

		yield put(setSuccess('deleteinfraction', 'Infraction deleted'));
		yield put(setErrors('deleteinfraction', undefined));
	} catch (err) {
		console.log('Could not delete infraction', err);
		const { data } = err.response;

		yield put(setSuccess('deleteinfraction', undefined));
		yield put(setErrors('deleteinfraction', data.message));
	}
}

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

function* watchUpdateInfraction() {
	yield takeLatest(UPDATE_INFRACTION, updateInfractionAsync);
}

function* watchDeleteInfraction() {
	yield takeLatest(DELETE_INFRACTION, deleteInfractionAsync);
}

export default function* infractionSagas() {
	yield all([
		call(watchCreateWarning),
		call(watchCreateMute),
		call(watchCreateKick),
		call(watchCreateBan),
		call(watchUpdateInfraction),
		call(watchDeleteInfraction),
	]);
}
