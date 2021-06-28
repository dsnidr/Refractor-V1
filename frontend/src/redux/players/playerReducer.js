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

import {
	SET_CURRENT_PLAYER,
	SET_PLAYER_WATCHED,
	SET_RECENT_PLAYERS,
	SET_PLAYER_SEARCH_RESULTS,
} from './playerActions';

const initialState = {
	currentPlayer: null,
	searchResults: [],
	recentPlayers: [],
};

const playerReducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_CURRENT_PLAYER:
			return {
				...state,
				currentPlayer: action.payload,
			};
		case SET_PLAYER_SEARCH_RESULTS:
			return {
				...state,
				searchResults: action.payload,
			};
		case SET_RECENT_PLAYERS:
			return {
				...state,
				recentPlayers: action.payload,
			};
		case SET_PLAYER_WATCHED:
			return setPlayerWatched(state, action.playerId, action.payload);
		default:
			return state;
	}
};

function setPlayerWatched(state, playerId, watched) {
	if (state.currentPlayer && state.currentPlayer.id === playerId) {
		return {
			...state,
			currentPlayer: {
				...state.currentPlayer,
				watched: watched,
			},
		};
	}

	return state;
}

export default playerReducer;
