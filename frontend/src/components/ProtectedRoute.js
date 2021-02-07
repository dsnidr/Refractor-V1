import React from 'react';
import { Redirect, Route } from 'react-router';
import { connect } from 'react-redux';

const ProtectedRoute = ({
	component: Component,
	user,
	bypassPasswordChange,
	setErrorAlert,
	...rest
}) => (
	<Route
		{...rest}
		render={(props) => {
			if (user !== null && user.isAuthenticated === true) {
				if (user.needsPasswordChange && !bypassPasswordChange) {
					setErrorAlert(
						'You need to change your password before accessing the dashboard.'
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

const mapStateToProps = (state) => ({
	user: state.user.self,
});

export default connect(mapStateToProps)(ProtectedRoute);
