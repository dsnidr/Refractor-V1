import { all, call } from 'redux-saga/effects';

import userSagas from './user/userSagas';
import gameSagas from './games/gameSagas';
import serverSagas from './servers/serverSagas';

export default function* rootSaga() {
	yield all([call(userSagas), call(gameSagas), call(serverSagas)]);
}
