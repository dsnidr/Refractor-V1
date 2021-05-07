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
