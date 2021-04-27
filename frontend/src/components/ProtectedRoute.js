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

import React from 'react';
import { Redirect, Route } from 'react-router';
import { connect } from 'react-redux';
import { setErrors } from '../redux/error/errorActions';
import PropTypes from 'prop-types';

const ProtectedRoute = ({
	component: Component,
	user,
	bypassPasswordChange,
	setError,
	...rest
}) => (
	<Route
		{...rest}
		render={(props) => {
			if (user !== null && user.isAuthenticated === true) {
				if (user.needsPasswordChange && !bypassPasswordChange) {
					setError(
						'You need to change your password before accessing Refractor.'
					);

					return (
						<Redirect
							to={{
								pathname: '/changepassword',
								state: { from: props.location },
							}}
						/>
					);
				}

				return <Component {...props} />;
			}

			return (
				<Redirect
					to={{ pathname: '/login', state: { from: props.location } }}
				/>
			);
		}}
	/>
);

ProtectedRoute.propTypes = {
	component: PropTypes.any.isRequired,
	bypassPasswordChange: PropTypes.bool,
};

const mapStateToProps = (state) => ({
	user: state.user.self,
});

const mapDispatchToProps = (dispatch) => ({
	setError: (error) => dispatch(setErrors('changepassword', error)),
});

export default connect(mapStateToProps, mapDispatchToProps)(ProtectedRoute);
