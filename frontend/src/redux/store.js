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

import { createStore, applyMiddleware, compose } from 'redux';
import createSagaMiddleware from 'redux-saga';
import { createBrowserHistory } from 'history';

import createRootReducer from './rootReducer';
import rootSaga from './rootSaga';

const history = createBrowserHistory();
const rootReducer = createRootReducer(history);

const sagaMiddleware = createSagaMiddleware();
const middleware = [sagaMiddleware];

const useMiddleware = () => {
	if (process.env.NODE_ENV === 'development') {
		return compose(
			applyMiddleware(...middleware),
			window.__REDUX_DEVTOOLS_EXTENSION__ &&
				window.__REDUX_DEVTOOLS_EXTENSION__()
		);
	}

	return applyMiddleware(...middleware);
};

// eslint-disable-next-line react-hooks/rules-of-hooks
const store = createStore(rootReducer, useMiddleware());
sagaMiddleware.run(rootSaga);

export { store, history };
