import {
	ADD_PLAYER_TO_SERVER,
	REMOVE_PLAYER_FROM_SERVER,
	SET_SERVER_STATUS,
	SET_SERVERS,
} from './serverActions';

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
		default:
			return state;
	}
};

function addPlayerToServer(state, serverId, player) {
	if (!state[serverId]) {
		return;
	}

	let players = [];

	if (state[serverId].players) {
		players = state[serverId].players;
	}

	players.push(player);

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
		return;
	}

	let players = [];

	if (state[serverId].players) {
		players = state[serverId].players;
	}

	players = players.filter((arrPlayer) => arrPlayer.id !== player.id);

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
		return;
	}

	return {
		...state,
		[serverId]: {
			...state[serverId],
			online: isOnline,
		},
	};
}

export default reducer;
