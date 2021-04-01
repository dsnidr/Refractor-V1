export const ADD_CHAT_MESSAGE = 'ADD_CHAT_MESSAGE';
export const addChatMessage = (message) => ({
	type: ADD_CHAT_MESSAGE,
	payload: message,
});

export const SEND_CHAT_MESSAGE = 'SEND_CHAT_MESSAGE';
export const sendChatMessage = (message, serverId) => ({
	type: SEND_CHAT_MESSAGE,
	serverId: serverId,
	payload: message,
});
