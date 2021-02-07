import React, { Component } from 'react';
import { Provider, connect } from 'react-redux';

import { store, history } from './redux/store';
import { Router, Switch } from 'react-router';
import { setTheme } from './redux/theme/themeActions';
import { css, ThemeProvider } from 'styled-components';
import { getUserInfo, setUser } from './redux/user/userActions';
import styled from 'styled-components';
import themes from './themes';
import UnprotectedRoute from './components/UnprotectedRoute';
import ProtectedRoute from './components/ProtectedRoute';

import Login from './pages/Login';
import Dashboard from './pages/Dashboard';
import Spinner from './components/Spinner';
import { decodeToken, destroyToken, getToken } from './utils/tokenUtils';

// Load previously selected theme
let theme = localStorage.getItem('theme');

if (!theme) {
	localStorage.setItem('theme', 'dark');
	theme = 'dark';
}

// Set the theme
store.dispatch(setTheme(theme));

const AppContainer = styled.div`
	${(props) => css`
		background: ${props.theme.colorBackground};
	`}
`;

class App extends Component {
	constructor(props) {
		super(props);

		this.state = {
			tokenChecked: false,
		};
	}

	componentDidMount() {
		const token = getToken();

		if (token) {
			// Try to decode token
			const decoded = decodeToken(token);

			// If token is invalid, destroy it and mark the user as not authenticated and return
			if (!decoded) {
				destroyToken();
				this.props.setUser({ isAuthenticated: false });
				return;
			}

			// If token could be decoded, try to get user info.
			//
			// If the token is expired, this will kick off our token refresh logic
			// and if it's still current, it will fetch the user info.
			//
			// If the refresh logic fails, the user will be marked as not authenticated by the redux saga.
			this.props.getUserInfo();
		} else {
			// If no token is present, mark the user as not authenticated
			this.props.setUser({ isAuthenticated: false });
		}
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		const nextState = prevState;

		if (nextProps.user !== null) {
			nextState.tokenChecked = true;
		}

		return nextState;
	}

	render() {
		return (
			<AppContainer>
				<ThemeProvider theme={themes[this.props.theme]}>
					{this.props.isLoading ? <Spinner /> : null}
					<Router history={history}>
						<Switch>
							<UnprotectedRoute
								exact
								path={'/login'}
								component={Login}
							/>
							<ProtectedRoute path={'/'} component={Dashboard} />
						</Switch>
					</Router>
				</ThemeProvider>
			</AppContainer>
		);
	}
}

const mapStateToProps = (state) => ({
	isLoading: state.loading,
	user: state.user.self,
	theme: state.theme,
});

const mapDispatchToProps = (dispatch) => ({
	getUserInfo: () => dispatch(getUserInfo()),
	setUser: (user) => dispatch(setUser(user)),
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
