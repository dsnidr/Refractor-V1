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

export const SET_ERRORS = 'SET_ERRORS';
export const setErrors = (field, errors) => ({
	type: SET_ERRORS,
	field: field,
	payload: errors,
});

export const CLEAR_ERRORS = 'CLEAR_ERRORS';
export const clearErrors = (field) => ({
	type: CLEAR_ERRORS,
	field: field,
});

export const CLEAR_ALL_ERRORS = 'CLEAR_ALL_ERRORS';
export const clearAllErrors = () => ({
	type: CLEAR_ALL_ERRORS,
});
