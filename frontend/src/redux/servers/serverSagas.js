import { all, call, put, takeLatest } from 'redux-saga/effects';
import {
	CREATE_SERVER,
	DELETE_SERVER,
	GET_SERVERS,
	removeServer,
	setServers,
	UPDATE_SERVER,
} from './serverActions';
import {
	createServer,
	deleteServer,
	getAllServerData,
	updateServer,
} from '../../api/serverApi';
import { setErrors } from '../error/errorActions';
import { setSuccess } from '../success/successActions';

function* getServersAsync(action) {
	try {
		const { data } = yield call(getAllServerData);

		const servers = {};

		// Flatten the server structure
		data.payload.forEach((server) => {
			servers[server.id] = server;
		});

		yield put(setServers(servers));
	} catch (err) {
		console.log('Could not get server data', err);
		yield put(
			setErrors(
				'servers',
				`Could not get server data: ${err.response.data.message}`
			)
		);
	}
}

function* updateServerAsync(action) {
	try {
		yield call(updateServer, action.serverId, action.payload);

		yield put(setErrors('editserver', undefined));
		yield put(setSuccess('editserver', 'Server updated successfully'));
	} catch (err) {
		console.log('Could not update server', err);
		const { data } = err.response;

		yield put(setSuccess('editserver', undefined));
		yield put(
			setErrors(
				'editserver',
				!data.errors
					? `Could not update server: ${err.response.data.message}`
					: data.errors
			)
		);
	}
}

function* createServerAsync(action) {
	try {
		yield call(createServer, action.payload);

		yield put(setErrors('createserver', undefined));
		yield put(setSuccess('createserver', 'Server added successfully'));
	} catch (err) {
		console.log('Could not create server', err);
		const { data } = err.response;

		yield put(setSuccess('createserver', undefined));
		yield put(
			setErrors(
				'createserver',
				!data.errors
					? `Could not create server: ${err.response.data.message}`
					: data.errors
			)
		);
	}
}

function* deleteServerAsync(action) {
	try {
		yield put(removeServer(action.serverId));
		yield call(deleteServer, action.serverId);

		yield put(setSuccess('deleteserver', 'Server deleted'));
		yield put(setErrors('deleteserver', undefined));
	} catch (err) {
		yield put(setErrors('deleteserver', err.response.data.message));
		yield put(setSuccess('deleteserver', undefined));
	}
}

function* watchGetServers() {
	yield takeLatest(GET_SERVERS, getServersAsync);
}

function* watchUpdateServer() {
	yield takeLatest(UPDATE_SERVER, updateServerAsync);
}

function* watchCreateServer() {
	yield takeLatest(CREATE_SERVER, createServerAsync);
}

function* watchDeleteServer() {
	yield takeLatest(DELETE_SERVER, deleteServerAsync);
}

export default function* gameSagas() {
	yield all([
		call(watchGetServers),
		call(watchUpdateServer),
		call(watchCreateServer),
		call(watchDeleteServer),
	]);
}
