/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
		if (data.payload) {
			data.payload.forEach((server) => {
				servers[server.id] = server;
			});
		}

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

export default function* serverSagas() {
	yield all([
		call(watchGetServers),
		call(watchUpdateServer),
		call(watchCreateServer),
		call(watchDeleteServer),
	]);
}
