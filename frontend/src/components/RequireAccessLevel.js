import React from 'react';
import { connect } from 'react-redux';

const RequireAccessLevel = (props) => {
	const { minAccessLevel, or, children } = props;

	if (props.user.accessLevel < minAccessLevel && or !== true) {
		return null;
	}

	return children;
};

const mapStateToProps = (state) => ({
	user: state.user.self,
});

export default connect(mapStateToProps, null)(RequireAccessLevel);
