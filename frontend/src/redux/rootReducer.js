import { combineReducers } from 'redux';
import { connectRouter } from 'connected-react-router';

import themeReducer from './theme/themeReducer';
import userReducer from './user/userReducer';
import errorReducer from './error/errorReducer';
import successReducer from './success/successReducer';
import gameReducer from './games/gameReducer';
import serverReducer from './servers/serverReducer';
import loadingReducer from './loading/loadingReducer';

const createRootReducer = (history) =>
	combineReducers({
		router: connectRouter(history),
		loading: loadingReducer,
		theme: themeReducer,
		error: errorReducer,
		success: successReducer,
		user: userReducer,
		games: gameReducer,
		servers: serverReducer,
	});

export default createRootReducer;
