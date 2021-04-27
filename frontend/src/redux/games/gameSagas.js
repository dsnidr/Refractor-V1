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
