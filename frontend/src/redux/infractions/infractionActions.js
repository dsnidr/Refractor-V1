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
