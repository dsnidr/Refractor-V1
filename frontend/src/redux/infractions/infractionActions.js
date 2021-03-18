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
