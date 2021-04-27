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
	SET_CURRENT_USER,
	SET_USER_ACTIVATED,
	SET_USER_DEACTIVATED,
} from './constants';
import { SET_ALL_USERS } from './userActions';

const initialState = {
	self: null,
	others: null,
};

const reducer = (state = initialState, action) => {
	switch (action.type) {
		case SET_CURRENT_USER:
			return {
				...state,
				self: action.payload,
			};
		case SET_ALL_USERS:
			return {
				...state,
				others: action.payload,
			};
		case SET_USER_ACTIVATED:
			return {
				...state,
				others: {
					...state.others,
					[action.userId]: {
						...state.others[action.userId],
						activated: true,
					},
				},
			};
		case SET_USER_DEACTIVATED:
			return {
				...state,
				others: {
					...state.others,
					[action.userId]: {
						...state.others[action.userId],
						activated: false,
					},
				},
			};
		default:
			return state;
	}
};

export default reducer;
