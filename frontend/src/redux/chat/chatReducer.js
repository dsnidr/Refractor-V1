import { ADD_CHAT_MESSAGE } from './chatActions';

const initialState = {};

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case ADD_CHAT_MESSAGE:
			return addChatMessage(state, action.payload);
		default:
			return state;
	}
};

function addChatMessage(state, messageBody) {
	let serverMessages = state[messageBody.serverId];

	if (!serverMessages) {
		serverMessages = [];
	}

	serverMessages.push(messageBody);

	return {
		...state,
		[messageBody.serverId]: serverMessages,
	};
}

export default reducer;
