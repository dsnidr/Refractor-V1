export const GET_SERVERS = 'GET_SERVERS';
export const getServers = () => ({
	type: GET_SERVERS,
});

export const SET_SERVERS = 'SET_SERVERS';
export const setServers = (servers) => ({
	type: SET_SERVERS,
	payload: servers,
});

export const ADD_PLAYER_TO_SERVER = 'ADD_PLAYER_TO_SERVER';
export const addPlayerToServer = (serverId, player) => ({
	type: ADD_PLAYER_TO_SERVER,
	serverId: serverId,
	payload: player,
});

export const REMOVE_PLAYER_FROM_SERVER = 'REMOVE_PLAYER_FROM_SERVER';
export const removePlayerFromServer = (serverId, player) => ({
	type: REMOVE_PLAYER_FROM_SERVER,
	serverId: serverId,
	payload: player,
});
