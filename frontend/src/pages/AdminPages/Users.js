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
import Heading from '../../components/Heading';
import { setLoading } from '../../redux/loading/loadingActions';
import { connect } from 'react-redux';
import {
	activateUser,
	deactivateUser,
	getAllUsers,
} from '../../redux/user/userActions';
import styled, { css } from 'styled-components';
import Button from '../../components/Button';
import BasicModal from '../../components/modals/BasicModal';

const UserInfo = styled.div`
	${(props) => css`
		padding: 2rem;
		margin-bottom: 2rem;
		border-radius: ${props.theme.borderRadiusNormal};
		background-color: ${props.theme.colorAccent};
		font-size: 2rem;
		display: flex;
		justify-content: space-between;
		align-items: center;
		text-decoration: none;
		color: ${props.theme.colorTextSecondary};

		:hover {
			cursor: pointer;
		}

		h1 {
			font-size: 2rem;
			font-weight: 400;
		}
		div {
			display: flex;
			// Override button styling
			> * {
				width: auto;
				margin-right: 1rem;
				font-size: 1.3rem;
			}
			> :last-child {
				margin-right: 0;
			}
		}
		:last-of-type {
			margin-bottom: 0;
		}

		:first-of-type {
			margin-top: 1rem;
		}
	`}
`;

const UserDisplay = styled.div`
	${(props) => css`
		> button {
			margin-top: 2rem;
		}
	`}
`;

class Users extends Component {
	constructor(props) {
		super(props);

		this.state = {
			modals: {
				activateUser: {
					show: false,
					ctx: {},
					success: null,
					error: null,
				},
				deactivateUser: {
					show: false,
					ctx: {},
					success: null,
					error: null,
				},
			},
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		if (!nextProps.users) {
			nextProps.setLoading(true);
			nextProps.getAllUsers();
		} else if (nextProps.users && nextProps.isLoading) {
			nextProps.setLoading(false);
		}

		return prevState;
	}

	onUserClick = (userId) => () => {
		this.props.history.push(`/user/${userId}`);
	};

	hideModal = (name) => () => {
		this.setState((prevState) => ({
			...prevState,
			modals: {
				...prevState.modals,
				[name]: {
					...prevState.modals[name],
					show: false,
					ctx: {},
				},
			},
		}));
	};

	onActivateUserClick = (user) => (e) => {
		e.stopPropagation();

		this.setState((prevState) => ({
			...prevState,
			modals: {
				...prevState.modals,
				activateUser: {
					...prevState.modals.activateUser,
					show: true,
					ctx: user,
				},
			},
		}));
	};

	onDeactivateUserClick = (user) => (e) => {
		e.stopPropagation();

		this.setState((prevState) => ({
			...prevState,
			modals: {
				...prevState.modals,
				deactivateUser: {
					...prevState.modals.deactivateUser,
					show: true,
					ctx: user,
				},
			},
		}));
	};

	activateUser = () => {
		const { ctx: user } = this.state.modals.activateUser;

		this.props.activateUser(user.id);

		this.hideModal('activateUser')();
	};

	deactivateUser = () => {
		const { ctx: user } = this.state.modals.deactivateUser;

		this.props.deactivateUser(user.id);

		this.hideModal('deactivateUser')();
	};

	onAddUserClick = () => {
		this.props.history.push('/users/add');
	};

	render() {
		const { users: usersObj } = this.props;

		if (!usersObj) {
			return null;
		}

		const users = Object.values(usersObj);

		const { activateUser, deactivateUser } = this.state.modals;

		return (
			<>
				<BasicModal
					show={activateUser.show}
					heading={`Activate ${activateUser.ctx.username}`}
					message={`Are you sure you wish to activate ${activateUser.ctx.username}?`}
					submitLabel={'Activate User'}
					success={activateUser.success}
					error={activateUser.error}
					onClose={this.hideModal('activateUser')}
					onSubmit={this.activateUser}
				/>

				<BasicModal
					show={deactivateUser.show}
					heading={`Deactivate ${deactivateUser.ctx.username}`}
					message={`Are you sure you wish to deactivate ${deactivateUser.ctx.username}?`}
					submitLabel={'Deactivate User'}
					success={deactivateUser.success}
					error={deactivateUser.error}
					onClose={this.hideModal('deactivateUser')}
					onSubmit={this.deactivateUser}
				/>

				<div>
					<Heading headingStyle={'title'}>Users</Heading>
				</div>

				<UserDisplay>
					<Heading headingStyle={'subtitle'}>Activated users</Heading>
					{users.map((user) => {
						if (!user.activated) {
							return null;
						}

						return (
							<UserInfo
								key={user.id}
								to={`/user/${user.id}`}
								onClick={this.onUserClick(user.id)}
							>
								<h1>{user.username}</h1>

								<div>
									<Button
										size={'normal'}
										color={'danger'}
										onClick={this.onDeactivateUserClick(
											user
										)}
									>
										Deactivate
									</Button>
								</div>
							</UserInfo>
						);
					})}

					<Button size={'normal'} onClick={this.onAddUserClick}>
						Add User
					</Button>
				</UserDisplay>

				<UserDisplay>
					<Heading headingStyle={'subtitle'}>
						Deactivated users
					</Heading>
					{users.map((user) => {
						if (user.activated) {
							return null;
						}

						return (
							<UserInfo
								key={user.id}
								to={`/user/${user.id}`}
								onClick={this.onUserClick(user.id)}
							>
								<h1>{user.username}</h1>

								<div>
									<Button
										size={'normal'}
										color={'alert'}
										onClick={this.onActivateUserClick(user)}
									>
										Reactivate
									</Button>
								</div>
							</UserInfo>
						);
					})}
				</UserDisplay>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	self: state.user.self,
	users: state.user.others,
	isLoading: state.loading.users,
});

const mapDispatchToProps = (dispatch) => ({
	setLoading: (isLoading) => dispatch(setLoading('users', isLoading)),
	getAllUsers: () => dispatch(getAllUsers()),
	activateUser: (userId) => dispatch(activateUser(userId)),
	deactivateUser: (userId) => dispatch(deactivateUser(userId)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Users);
