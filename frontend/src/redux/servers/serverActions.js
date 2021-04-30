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

export const SET_SERVER_STATUS = 'SET_SERVER_STATUS';
export const setServerStatus = (serverId, isOnline) => ({
	type: SET_SERVER_STATUS,
	serverId: serverId,
	payload: isOnline,
});

export const UPDATE_SERVER = 'UPDATE_SERVER';
export const updateServer = (serverId, editedData) => ({
	type: UPDATE_SERVER,
	serverId: serverId,
	payload: editedData,
});

export const CREATE_SERVER = 'CREATE_SERVER';
export const createServer = (data) => ({
	type: CREATE_SERVER,
	payload: data,
});

export const DELETE_SERVER = 'DELETE_SERVER';
export const deleteServer = (serverId) => ({
	type: DELETE_SERVER,
	serverId: serverId,
});

export const REMOVE_SERVER = 'REMOVE_SERVER';
export const removeServer = (serverId) => ({
	type: REMOVE_SERVER,
	serverId: serverId,
});

export const UPDATE_ONLINE_PLAYER = 'UPDATE_ONLINE_PLAYER'

// updateOnlinePlayer updates a player in state if they are online in any server.
// fields can either be an object with updated fields which will be spread out in
// order to do the update, or it can be a function.
//
// If fields is a function, the player will be provided as an argument. Any passed
// in function must return the updated player.
//
// Example:
//
// updateOnlinePlayer(1, (player) =>
//   { ...player, infractionCount: player.infractionCount + 1 });
export const updateOnlinePlayer = (playerId, fields) => ({
	type: UPDATE_ONLINE_PLAYER,
	playerId: playerId,
	payload: fields
})