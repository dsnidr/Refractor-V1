import { all, call } from 'redux-saga/effects';

import userSagas from './user/userSagas';
import gameSagas from './game/gameSagas';

export default function* rootSaga() {
	yield all([call(userSagas), call(gameSagas)]);
}
