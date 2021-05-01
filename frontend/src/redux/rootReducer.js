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

import { combineReducers } from 'redux';
import { connectRouter } from 'connected-react-router';

import themeReducer from './theme/themeReducer';
import userReducer from './user/userReducer';
import errorReducer from './error/errorReducer';
import successReducer from './success/successReducer';
import gameReducer from './games/gameReducer';
import serverReducer from './servers/serverReducer';
import loadingReducer from './loading/loadingReducer';
import playerReducer from './players/playerReducer';
import infractionReducer from './infractions/infractionReducer';
import chatReducer from './chat/chatReducer';
import { reducer as toastrReducer } from 'react-redux-toastr';

const createRootReducer = (history) =>
	combineReducers({
		router: connectRouter(history),
		loading: loadingReducer,
		theme: themeReducer,
		error: errorReducer,
		success: successReducer,
		user: userReducer,
		games: gameReducer,
		servers: serverReducer,
		players: playerReducer,
		infractions: infractionReducer,
		chat: chatReducer,
		toastr: toastrReducer,
	});

export default createRootReducer;
