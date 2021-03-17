import React from 'react';
import { connect } from 'react-redux';
import { flags, hasPermission } from '../permissions/permissions';
import PropTypes from 'prop-types';

const RequirePerms = (props) => {
	let { perms, mode, children, user } = props;

	if (!user) {
		return null;
	}

	if (!Array.isArray(perms)) {
		return children;
	}

	const userPerms = BigInt(user.permissions);

	// If user is super admin or has full access, just return the children since
	// they for sure have permission.
	if (
		hasPermission(userPerms, flags.SUPER_ADMIN) ||
		hasPermission(userPerms, flags.FULL_ACCESS)
	) {
		return children;
	}

	if (!mode) {
		mode = 'all';
	}

	/* global BigInt */
	switch (mode) {
		case 'all':
			let returnChildren = true;
			for (let i = 0; i < perms.length && returnChildren; i++) {
				let perm = perms[i];

				if (!hasPermission(userPerms, perm)) {
					returnChildren = false;
				}
			}

			if (returnChildren) {
				return children;
			}

			break;
		case 'any':
			for (let i = 0; i < perms.length; i++) {
				let perm = perms[i];

				if (hasPermission(userPerms, perm)) {
					return children;
				}
			}
			break;
		default:
			console.log('RequirePerms mode not defined', mode);
			break;
	}

	return null;
};

RequirePerms.propTypes = {
	perms: PropTypes.arrayOf.BigIntLiteral,
	mode: PropTypes.string,
	children: PropTypes.any,
};

const mapStateToProps = (state) => ({
	user: state.user.self,
});

export default connect(mapStateToProps, null)(RequirePerms);
