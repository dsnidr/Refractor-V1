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

export const SET_SUCCESS = 'SET_SUCCESS';
export const setSuccess = (field, message) => ({
	type: SET_SUCCESS,
	field: field,
	payload: message,
});

export const CLEAR_SUCCESS = 'CLEAR_SUCCESS';
export const clearSuccess = (field) => ({
	type: CLEAR_SUCCESS,
	field: field,
});

export const CLEAR_ALL_SUCCESS = 'CLEAR_ALL_SUCCESS';
export const clearAllSuccess = () => ({
	type: CLEAR_ALL_SUCCESS,
});
