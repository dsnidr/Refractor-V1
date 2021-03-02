import { all, call, put, takeLatest } from 'redux-saga/effects';
import {
	CREATE_SERVER,
	GET_SERVERS,
	setServers,
	UPDATE_SERVER,
} from './serverActions';
import {
	createServer,
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

function* watchGetServers() {
	yield takeLatest(GET_SERVERS, getServersAsync);
}

function* watchUpdateServer() {
	yield takeLatest(UPDATE_SERVER, updateServerAsync);
}

function* watchCreateServer() {
	yield takeLatest(CREATE_SERVER, createServerAsync);
}

export default function* gameSagas() {
	yield all([
		call(watchGetServers),
		call(watchUpdateServer),
		call(watchCreateServer),
	]);
}
