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

import { all, call } from 'redux-saga/effects';

import userSagas from './user/userSagas';
import gameSagas from './games/gameSagas';
import serverSagas from './servers/serverSagas';
import infractionSagas from './infractions/infractionSagas';
import playerSagas from './players/playerSagas';
import chatSagas from './chat/chatSagas';

export default function* rootSaga() {
	yield all([
		call(userSagas),
		call(gameSagas),
		call(serverSagas),
		call(infractionSagas),
		call(playerSagas),
		call(chatSagas),
	]);
}
