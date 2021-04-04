import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';
import {
	forceUserPasswordChange,
	getAllUsers,
	setUserPassword,
	setUserPermissions,
} from '../../redux/user/userActions';
import TextInput from '../../components/TextInput';
import Button from '../../components/Button';
import respondTo from '../../mixins/respondTo';
import Checkbox from '../../components/Checkbox';
import {
	FULL_ACCESS,
	LOG_WARNING,
	LOG_MUTE,
	LOG_KICK,
	LOG_BAN,
	EDIT_OWN_INFRACTIONS,
	EDIT_ANY_INFRACTION,
	DELETE_OWN_INFRACTIONS,
	DELETE_ANY_INFRACTION,
	flags,
	isRestricted,
	getGrantedPerms,
} from '../../permissions/permissions';
import Alert from '../../components/Alert';

const ManagementSection = styled.div`
	${(props) => css`
		> :first-child {
			margin: 1rem 0;
		}
	`}
`;

const ManagementRow = styled.div`
	${(props) => css`
		display: flex;

		> :last-child {
			margin-right: 0;
		}

		button {
			width: auto;
			height: 5rem;
		}

		${respondTo.medium`
          .text-input {
            width: 50rem;
          }

          > * {
            margin-right: 2rem;
          }
		`}
	`}
`;

const PermissionsManager = styled.div`
	${(props) => css`
		button {
			margin-top: 2rem;
		}
	`}
`;

const PermissionCheckboxes = styled.div`
	display: grid;
	grid-template-columns: 1fr;

	${respondTo.small`
    	grid-template-columns: 1fr 1fr;
	`}

	${respondTo.medium`
	  	grid-template-columns: 1fr 1fr 1fr;
	`}
	
	${respondTo.large`
		grid-template-columns: 1fr 1fr 1fr 1fr;
	`}
`;

const PermissionButtons = styled.div`
	display: flex;

	> :last-child {
		margin-left: 1rem;
	}
`;

class User extends Component {
	constructor(props) {
		super(props);

		this.state = {
			perms: {},
			errors: {},
			initialPerms: null,
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		if (nextProps.passwordErrors) {
			prevState.errors = {
				...prevState.errors,
				newPassword: nextProps.passwordErrors.newPassword,
			};
		}

		if (nextProps.passwordSuccess) {
			prevState.errors = {
				...prevState.errors,
				newPassword: false,
			};
		}

		// Load permissions if user was found
		if (!prevState.initialPerms && prevState.user) {
			const grantedPerms = getGrantedPerms(prevState.user.permissions);

			const perms = {};

			grantedPerms.forEach((perm) => {
				perms[perm] = true;
			});

			prevState.initialPerms = perms;
			prevState.perms = perms;
		}

		if (!nextProps.users) {
			nextProps.getAllUsers();
			return prevState;
		}

		// Get user data
		const id = parseInt(nextProps.match.params.id);
		if (!id) {
			return prevState;
		}

		if (!prevState.user && !nextProps.users[id]) {
			return prevState;
		}

		prevState.user = nextProps.users[id];

		// TODO: Load initial perms into state

		return prevState;
	}

	handlePermChange = (cbName) => (e) => {
		this.setState((prevState) => ({
			...prevState,
			perms: {
				...prevState.perms,
				[cbName]: e.target.checked,
			},
		}));
	};

	handleAdminPermChange = () => {
		if (!!this.state.perms[FULL_ACCESS]) {
			return this.setState((prevState) => ({
				...prevState,
				perms: {
					...prevState.perms,
					[FULL_ACCESS]: false,
				},
			}));
		}

		const newPerms = {};

		Object.keys(flags).forEach((flag) => {
			if (isRestricted(flag)) {
				return;
			}

			newPerms[flag] = true;
		});

		this.setState((prevState) => ({
			...prevState,
			perms: newPerms,
		}));
	};

	onPermSaveClick = () => {
		const { user, perms } = this.state;

		/* global BigInt */
		let flag = BigInt(
			0b0000000000000000000000000000000000000000000000000000000000000000
		);

		Object.keys(perms).forEach((key) => {
			if (!perms[key]) {
				return;
			}

			flag = flag | flags[key];
		});

		this.props.setPermissions(user.id, flag);
	};

	revertPermChanges = () => {
		this.setState((prevState) => ({
			...prevState,
			perms: {
				...prevState.initialPerms,
			},
		}));
	};

	onTextInputChange = (e) => {
		this.setState((prevState) => ({
			...prevState,
			[e.target.name]: e.target.value,
		}));
	};

	onSetPasswordClick = () => {
		const { newPassword, user } = this.state;

		this.props.setUserPassword(user.id, { newPassword });
	};

	onForcePasswordChangeClick = () => {
		const { user } = this.state;

		this.props.forcePasswordChange(user.id);
	};

	render() {
		const { user, perms, errors } = this.state;
		const {
			passwordSuccess,
			passwordErrors,
			permErrors,
			permSuccess,
		} = this.props;

		const adminBoxChecked = !!perms[FULL_ACCESS];

		if (!user) {
			return (
				<div>
					<Heading headingStyle={'title'}>User not found</Heading>
				</div>
			);
		}

		console.log(perms);

		return (
			<>
				<div>
					<Heading headingStyle={'title'}>
						Viewing user: {user.username}
					</Heading>
				</div>

				<ManagementSection>
					<Alert
						type="error"
						message={
							typeof passwordErrors === 'string'
								? passwordErrors
								: null
						}
					/>
					<Alert type="success" message={passwordSuccess} />

					<Heading headingStyle={'subtitle'}>
						Password Management
					</Heading>

					<ManagementRow>
						<TextInput
							name={'newPassword'}
							type={'password'}
							placeholder={'New password'}
							onChange={this.onTextInputChange}
							error={errors.newPassword}
						/>
						<Button
							size={'normal'}
							color={'primary'}
							onClick={this.onSetPasswordClick}
						>
							Set password
						</Button>
					</ManagementRow>

					<ManagementRow>
						<Button
							size={'normal'}
							color={'alert'}
							onClick={this.onForcePasswordChangeClick}
						>
							Force password change
						</Button>
					</ManagementRow>
				</ManagementSection>

				<ManagementSection>
					<Alert
						type="error"
						message={
							typeof permErrors === 'string' ? permErrors : null
						}
					/>
					<Alert type="success" message={permSuccess} />

					<Heading headingStyle={'subtitle'}>
						Permission Management
					</Heading>

					<PermissionsManager>
						<PermissionCheckboxes>
							<Checkbox
								label={<strong>ADMIN (FULL ACCESS)</strong>}
								checked={!!perms[FULL_ACCESS]}
								onChange={this.handleAdminPermChange}
							/>
							<Checkbox
								label={'Log warning'}
								checked={!!perms[LOG_WARNING]}
								disabled={adminBoxChecked}
								onChange={this.handlePermChange(LOG_WARNING)}
							/>
							<Checkbox
								label={'Log mute'}
								checked={!!perms[LOG_MUTE]}
								disabled={adminBoxChecked}
								onChange={this.handlePermChange(LOG_MUTE)}
							/>
							<Checkbox
								label={'Log kick'}
								checked={!!perms[LOG_KICK]}
								disabled={adminBoxChecked}
								onChange={this.handlePermChange(LOG_KICK)}
							/>
							<Checkbox
								label={'Log ban'}
								checked={!!perms[LOG_BAN]}
								disabled={adminBoxChecked}
								onChange={this.handlePermChange(LOG_BAN)}
							/>
							<Checkbox
								label={'Edit own infractions'}
								checked={!!perms[EDIT_OWN_INFRACTIONS]}
								disabled={adminBoxChecked}
								onChange={this.handlePermChange(
									EDIT_OWN_INFRACTIONS
								)}
							/>
							<Checkbox
								label={'Edit any infractions'}
								checked={!!perms[EDIT_ANY_INFRACTION]}
								disabled={adminBoxChecked}
								onChange={this.handlePermChange(
									EDIT_ANY_INFRACTION
								)}
							/>
							<Checkbox
								label={'Delete own infractions'}
								checked={!!perms[DELETE_OWN_INFRACTIONS]}
								disabled={adminBoxChecked}
								onChange={this.handlePermChange(
									DELETE_OWN_INFRACTIONS
								)}
							/>
							<Checkbox
								label={'Delete any infraction'}
								checked={!!perms[DELETE_ANY_INFRACTION]}
								disabled={adminBoxChecked}
								onChange={this.handlePermChange(
									DELETE_ANY_INFRACTION
								)}
							/>
						</PermissionCheckboxes>

						<PermissionButtons>
							<Button
								size={'normal'}
								onClick={this.revertPermChanges}
							>
								Reset
							</Button>

							<Button
								size={'normal'}
								color={'alert'}
								onClick={this.onPermSaveClick}
							>
								Save
							</Button>
						</PermissionButtons>
					</PermissionsManager>
				</ManagementSection>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	users: state.user.others,
	passwordErrors: state.error.passwordmgmt,
	passwordSuccess: state.success.passwordmgmt,
	permErrors: state.error.setpermissions,
	permSuccess: state.success.setpermissions,
});

const mapDispatchToProps = (dispatch) => ({
	getAllUsers: () => dispatch(getAllUsers()),
	setUserPassword: (userId, data) => dispatch(setUserPassword(userId, data)),
	forcePasswordChange: (userId) => dispatch(forceUserPasswordChange(userId)),
	setPermissions: (userId, permissions) =>
		dispatch(setUserPermissions(userId, permissions)),
});

export default connect(mapStateToProps, mapDispatchToProps)(User);
