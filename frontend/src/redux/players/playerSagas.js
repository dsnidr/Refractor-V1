import { all, call, put, takeLatest } from 'redux-saga/effects';
import { GET_PLAYER_SUMMARY, setCurrentPlayer } from './playerActions';
import { getPlayerSummary } from '../../api/playerApi';
import { setErrors } from '../error/errorActions';

function* getPlayerSummaryAsync(action) {
	try {
		const { data } = yield call(getPlayerSummary, action.playerId);

		// Flatten data
		const summary = {
			...data.payload,
			warnings: {},
			mutes: {},
			kicks: {},
			bans: {},
		};

		if (data.payload.warnings) {
			data.payload.warnings.forEach((w) => (summary.warnings[w.id] = w));
		}

		if (data.payload.mutes) {
			data.payload.mutes.forEach((m) => (summary.mutes[m.id] = m));
		}

		if (data.payload.kicks) {
			data.payload.kicks.forEach((k) => (summary.kicks[k.id] = k));
		}

		if (data.payload.bans) {
			data.payload.bans.forEach((b) => (summary.bans[b.id] = b));
		}

		yield put(setCurrentPlayer(summary));
	} catch (err) {
		console.log('Could not get player summary', err);
		yield setErrors('playersummary', 'Could not get player summary');
	}
}

function* watchGetPlayerSummary() {
	yield takeLatest(GET_PLAYER_SUMMARY, getPlayerSummaryAsync);
}

export default function* playerSagas() {
	yield all([call(watchGetPlayerSummary)]);
}
