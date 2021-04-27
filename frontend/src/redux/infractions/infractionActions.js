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

export const CREATE_WARNING = 'CREATE_WARNING';
export const createWarning = (serverId, playerId, data) => ({
	type: CREATE_WARNING,
	serverId: serverId,
	playerId: playerId,
	payload: data,
});

export const CREATE_MUTE = 'CREATE_MUTE';
export const createMute = (serverId, playerId, data) => ({
	type: CREATE_MUTE,
	serverId: serverId,
	playerId: playerId,
	payload: data,
});

export const CREATE_KICK = 'CREATE_KICK';
export const createKick = (serverId, playerId, data) => ({
	type: CREATE_KICK,
	serverId: serverId,
	playerId: playerId,
	payload: data,
});

export const CREATE_BAN = 'CREATE_BAN';
export const createBan = (serverId, playerId, data) => ({
	type: CREATE_BAN,
	serverId: serverId,
	playerId: playerId,
	payload: data,
});

export const UPDATE_INFRACTION = 'UPDATE_INFRACTION';
export const updateInfraction = (infractionId, data) => ({
	type: UPDATE_INFRACTION,
	infractionId: infractionId,
	payload: data,
});

export const DELETE_INFRACTION = 'DELETE_INFRACTION';
export const deleteInfraction = (infractionId, data) => ({
	type: DELETE_INFRACTION,
	infractionId: infractionId,
});

export const SEARCH_INFRACTIONS = 'SEARCH_INFRACTIONS';
export const searchInfractions = (searchData) => ({
	type: SEARCH_INFRACTIONS,
	payload: searchData,
});

export const SET_INFRACTION_SEARCH_RESULTS = 'SET_INFRACTION_SEARCH_RESULTS';
export const setSearchResults = (results) => ({
	type: SET_INFRACTION_SEARCH_RESULTS,
	payload: results,
});

export const GET_RECENT_INFRACTIONS = 'GET_RECENT_INFRACTIONS';
export const getRecentInfractions = () => ({
	type: GET_RECENT_INFRACTIONS,
});

export const SET_RECENT_INFRACTIONS = 'SET_RECENT_INFRACTIONS';
export const setRecentInfractions = (recentInfractions) => ({
	type: SET_RECENT_INFRACTIONS,
	payload: recentInfractions,
});
