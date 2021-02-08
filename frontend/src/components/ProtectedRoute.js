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
				console.log('User is authenticated');
				if (user.needsPasswordChange && !bypassPasswordChange) {
					console.log('User needs password change');
					setError(
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

				console.log(
					'User is logged in and does not need a password change'
				);

				return <Component {...props} />;
			}

			console.log('User is not logged in');

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
	setError: (error) => dispatch(setErrors('auth', error)),
});

export default connect(mapStateToProps, mapDispatchToProps)(ProtectedRoute);
