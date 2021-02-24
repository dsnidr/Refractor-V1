import { all, call, put, takeLatest } from 'redux-saga/effects';
import { GET_SERVERS, setServers } from './serverActions';
import { getAllServerData } from '../../api/serverApi';
import { setErrors } from '../error/errorActions';

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

function* watchGetServers() {
	yield takeLatest(GET_SERVERS, getServersAsync);
}

export default function* gameSagas() {
	yield all([call(watchGetServers)]);
}
