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

import { all, call, put, takeLatest } from 'redux-saga/effects';
import { SEARCH_CHAT_RECORDS, setChatSearchResults } from './chatActions';
import { setSuccess } from '../success/successActions';
import { setErrors } from '../error/errorActions';
import { searchChatRecords } from '../../api/chatApi';

function* searchChatRecordsAsync(action) {
	try {
		const { data } = yield call(searchChatRecords, action.payload);

		yield put(setChatSearchResults(data.payload));

		yield put(setSuccess('chatrecords', 'Results fetched'));
		yield put(setErrors('chatrecords', undefined));
	} catch (err) {
		console.log('Could not search chat records', err);
		const { data } = err.response;

		yield put(setChatSearchResults([]));

		yield put(setSuccess('chatrecords', undefined));
		yield put(
			setErrors(
				'chatrecords',
				!data.errors
					? `Could not search chat records: ${err.response.data.message}`
					: data.errors
			)
		);
	}
}

function* watchSearchChatRecords() {
	yield takeLatest(SEARCH_CHAT_RECORDS, searchChatRecordsAsync);
}

export default function* infractionSagas() {
	yield all([call(watchSearchChatRecords)]);
}
