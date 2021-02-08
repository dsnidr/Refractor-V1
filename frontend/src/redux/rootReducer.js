import { combineReducers } from 'redux';
import { connectRouter } from 'connected-react-router';

import themeReducer from './theme/themeReducer';
import userReducer from './user/userReducer';
import errorReducer from './error/errorReducer';
import successReducer from './success/successReducer';

const createRootReducer = (history) =>
	combineReducers({
		router: connectRouter(history),
		theme: themeReducer,
		error: errorReducer,
		success: successReducer,
		user: userReducer,
	});

export default createRootReducer;
