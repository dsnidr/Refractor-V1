import { all, call, put, takeLatest } from 'redux-saga/effects';
import { GET_GAMES, setGames } from './gameActions';
import { setErrors } from '../error/errorActions';
import { getAllGames } from '../../api/gameApi';

function* getGamesAsync(action) {
	try {
		const { data } = yield call(getAllGames);

		const newGameState = {};

		data.payload.forEach((game) => {
			newGameState[game.name] = game;
		});

		yield put(setGames(newGameState));
	} catch (err) {
		console.log('Could not get all games', err);
		yield put(
			setErrors(
				'games',
				err.response.data.errors
					? err.response.data.errors
					: err.response.data.message
			)
		);
	}
}

function* watchGetGames() {
	yield takeLatest(GET_GAMES, getGamesAsync);
}

export default function* gameSagas() {
	yield all([call(watchGetGames)]);
}
