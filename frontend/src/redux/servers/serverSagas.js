import { all, call, put, takeLatest } from 'redux-saga/effects';
import { EDIT_SERVER, GET_SERVERS, setServers } from './serverActions';
import { editServer, getAllServerData } from '../../api/serverApi';
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
		console.log('Could not get servers', err);
		yield put(
			setErrors(
				'servers',
				`Could not get servers: ${err.response.data.message}`
			)
		);
	}
}

function* editServerAsync(action) {
	try {
		const { data } = yield call(
			editServer,
			action.serverId,
			action.payload
		);

		yield put(setSuccess('servers', 'Server edited successfully'));

		console.log(data);
	} catch (err) {
		console.log('Could not edit server', err);
		const { data } = err.response.data;

		yield put(
			setErrors(
				'servers',
				!data.errors
					? `Could not edit server: ${err.response.data.message}`
					: data.errors
			)
		);
	}
}

function* watchGetServers() {
	yield takeLatest(GET_SERVERS, getServersAsync);
}

function* watchEditServer() {
	yield takeLatest(EDIT_SERVER, editServerAsync);
}

export default function* gameSagas() {
	yield all([call(watchGetServers), call(watchEditServer)]);
}
