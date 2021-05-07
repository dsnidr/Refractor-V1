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

import { ADD_CHAT_MESSAGE, SET_CHAT_SEARCH_RESULTS } from './chatActions';

const initialState = {
	searchResults: [],
};

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case ADD_CHAT_MESSAGE:
			return addChatMessage(state, action.payload);
		case SET_CHAT_SEARCH_RESULTS:
			return {
				...state,
				searchResults: action.payload,
			};
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
