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

import React, { Component } from 'react';
import { connect } from 'react-redux';

import { store, history } from './redux/store';
import { Router, Switch } from 'react-router';
import { setTheme } from './redux/theme/themeActions';
import { css, ThemeProvider } from 'styled-components';
import { getUserInfo, refreshTokens, setUser } from './redux/user/userActions';
import styled from 'styled-components';
import themes from './themes';
import UnprotectedRoute from './components/UnprotectedRoute';
import ProtectedRoute from './components/ProtectedRoute';
import Spinner from './components/Spinner';
import { decodeToken, destroyToken, getToken } from './utils/tokenUtils';

import Login from './pages/Login';
import Dashboard from './pages/Dashboard';
import ChangePassword from './pages/ChangePassword';

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
			// Check if permissions fetched from the server are the same as the permissions in the token.
			// This is necessary to detect user permission changes and prevent desync.
			// If they are not equal, a refresh is forced.
			const decoded = decodeToken(getToken());

			if (nextProps.user.permissions !== decoded.permissions) {
				nextProps.refreshTokens();
			}

			// Mark token as checked
			nextState.tokenChecked = true;
		}

		return nextState;
	}

	render() {
		return this.state.tokenChecked ? (
			<AppContainer>
				<ThemeProvider theme={themes[this.props.theme]}>
					{this.props.isLoading ? <Spinner /> : null}

					<Switch>
						<UnprotectedRoute
							exact
							path={'/login'}
							component={Login}
						/>
						<ProtectedRoute
							exact
							path={'/changepassword'}
							bypassPasswordChange={true}
							component={ChangePassword}
						/>
						<ProtectedRoute path={'/'} component={Dashboard} />
					</Switch>
				</ThemeProvider>
			</AppContainer>
		) : null;
	}
}

const mapStateToProps = (state) => ({
	isLoading: state.loading.main,
	user: state.user.self,
	theme: state.theme,
});

const mapDispatchToProps = (dispatch) => ({
	getUserInfo: () => dispatch(getUserInfo()),
	setUser: (user) => dispatch(setUser(user)),
	refreshTokens: () => dispatch(refreshTokens()),
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
