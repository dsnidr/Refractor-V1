import { all, call } from 'redux-saga/effects';

import userSagas from './user/userSagas';
import gameSagas from './games/gameSagas';

export default function* rootSaga() {
	yield all([call(userSagas), call(gameSagas)]);
}
