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

// UnprotectedRoute redirects to Component only if the user is not authenticated
const UnprotectedRoute = ({ component: Component, user, ...rest }) => (
	<Route
		{...rest}
		render={(props) =>
			user === null || !user.isAuthenticated ? (
				<Component {...props} />
			) : (
				<Redirect to="/" />
			)
		}
	/>
);

const mapStateToProps = (state) => ({
	user: state.user.self,
});

export default connect(mapStateToProps)(UnprotectedRoute);
