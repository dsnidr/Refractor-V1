import { combineReducers } from 'redux';

import themeReducer from './theme/themeReducer';
import userReducer from './user/userReducer';
import errorReducer from './error/errorReducer';
import successReducer from './success/successReducer';

export default combineReducers({
	theme: themeReducer,
	error: errorReducer,
	success: successReducer,
	user: userReducer,
});
