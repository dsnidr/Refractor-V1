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
	CREATE_BAN,
	CREATE_KICK,
	CREATE_MUTE,
	CREATE_WARNING,
	DELETE_INFRACTION,
	GET_RECENT_INFRACTIONS,
	SEARCH_INFRACTIONS,
	setRecentInfractions,
	setSearchResults,
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
	getRecentInfractions,
	searchInfractions,
	updateInfraction,
} from '../../api/infractionApi';
import { updateOnlinePlayer } from '../servers/serverActions';

function* createWarningAsync(action) {
	try {
		const { data } = yield call(createWarning, {
			playerId: action.playerId,
			serverId: action.serverId,
			...action.payload,
		});

		yield put(setSuccess('createwarning', 'Warning logged'));
		yield put(setErrors('createwarning', undefined));

		yield put(updateOnlinePlayer(action.playerId, (player) => {
			return {
				...player,
				infractionCount: player.infractionCount + 1
			}
		}))
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

		yield put(updateOnlinePlayer(action.playerId, (player) => {
			return {
				...player,
				infractionCount: player.infractionCount + 1
			}
		}))
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

		yield put(updateOnlinePlayer(action.playerId, (player) => {
			return {
				...player,
				infractionCount: player.infractionCount + 1
			}
		}))
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

		yield put(updateOnlinePlayer(action.playerId, (player) => {
			return {
				...player,
				infractionCount: player.infractionCount + 1
			}
		}))
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

function* searchInfractionsAsync(action) {
	try {
		const { data } = yield call(searchInfractions, action.payload);

		yield put(setSearchResults(data.payload));

		yield put(setSuccess('searchinfractions', 'Results fetched'));
		yield put(setErrors('searchinfractions', undefined));
	} catch (err) {
		console.log('Could not search infractions', err);
		const { data } = err.response;

		yield put(setSearchResults([]));

		yield put(setSuccess('searchinfractions', undefined));
		yield put(
			setErrors(
				'searchinfractions',
				!data.errors
					? `Could not search infractions: ${err.response.data.message}`
					: data.errors
			)
		);
	}
}

function* getRecentInfractionsAsync(action) {
	try {
		const { data } = yield call(getRecentInfractions);

		yield put(setRecentInfractions(data.payload));
	} catch (err) {
		console.log('Could not get recent infractions. Error: %v', err);
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

function* watchSearchInfractions() {
	yield takeLatest(SEARCH_INFRACTIONS, searchInfractionsAsync);
}

function* watchGetRecentInfractions() {
	yield takeLatest(GET_RECENT_INFRACTIONS, getRecentInfractionsAsync);
}

export default function* infractionSagas() {
	yield all([
		call(watchCreateWarning),
		call(watchCreateMute),
		call(watchCreateKick),
		call(watchCreateBan),
		call(watchUpdateInfraction),
		call(watchDeleteInfraction),
		call(watchSearchInfractions),
		call(watchGetRecentInfractions),
	]);
}
