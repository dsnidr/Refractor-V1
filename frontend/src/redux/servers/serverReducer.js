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
	ADD_PLAYER_TO_SERVER,
	REMOVE_PLAYER_FROM_SERVER,
	SET_SERVER_STATUS,
	SET_SERVERS, UPDATE_ONLINE_PLAYER
} from './serverActions';
import { REMOVE_SERVER } from './constants';

const initialState = null;

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_SERVERS:
			return action.payload;
		case ADD_PLAYER_TO_SERVER:
			return addPlayerToServer(state, action.serverId, action.payload);
		case REMOVE_PLAYER_FROM_SERVER:
			return removePlayerFromServer(
				state,
				action.serverId,
				action.payload
			);
		case SET_SERVER_STATUS:
			return setServerStatus(state, action.serverId, action.payload);
		case REMOVE_SERVER:
			return removeServer(state, action.serverId);
		case UPDATE_ONLINE_PLAYER:
			return updateOnlinePlayer(state, action.playerId, action.payload);
		default:
			return state;
	}
};

function addPlayerToServer(state, serverId, player) {
	if (!state[serverId]) {
		return state;
	}

	let players = {};

	if (state[serverId].players) {
		players = state[serverId].players;
	}

	players[player.id] = player

	return {
		...state,
		[serverId]: {
			...state[serverId],
			players: players,
		},
	};
}

function removePlayerFromServer(state, serverId, player) {
	if (!state[serverId]) {
		return state;
	}

	let players = {};

	if (state[serverId].players) {
		players = state[serverId].players;
	}

	delete players[player.id]

	return {
		...state,
		[serverId]: {
			...state[serverId],
			players: players,
		},
	};
}

function setServerStatus(state, serverId, isOnline) {
	if (!state[serverId]) {
		return state;
	}

	return {
		...state,
		[serverId]: {
			...state[serverId],
			online: isOnline,
		},
	};
}

function removeServer(state, serverId) {
	if (!state[serverId]) {
		return state;
	}

	delete state[serverId];

	return state;
}

function updateOnlinePlayer(state, targetPlayerId, updateFields) {
	if (!targetPlayerId || (typeof updateFields !== 'function' && !updateFields)) {
		return state
	}

	Object.keys(state).forEach(serverId => {
		const server = state[serverId]

		if (!server.players) {
			return
		}

		console.log(updateFields)

		Object.keys(server.players).forEach(playerIdStr => {
			const playerId = parseInt(playerIdStr)

			if (playerId !== targetPlayerId) {
				return
			}

			if (typeof updateFields === 'function') {
				updateFields = updateFields(state[serverId].players[targetPlayerId])
			}

			state[serverId].players = {
				...state[serverId].players,
				[targetPlayerId]: {
					...state[serverId].players[targetPlayerId],
					...updateFields
				}
			}
		})
	})

	return state
}

export default reducer;
