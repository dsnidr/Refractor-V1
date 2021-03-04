import React, { Component } from 'react';
import Heading from '../../components/Heading';
import { setLoading } from '../../redux/loading/loadingActions';
import { connect } from 'react-redux';
import { getAllUsers } from '../../redux/user/userActions';
import Spinner from '../../components/Spinner';
import styled, { css } from 'styled-components';
import Button from '../../components/Button';

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
	`}
`;

const UserList = styled.div`
	${(props) => css`
		margin: 2rem 0;
	`}
`;

class Users extends Component {
	static getDerivedStateFromProps(nextProps, prevState) {
		if (!nextProps.users) {
			nextProps.setLoading(true);
			nextProps.getAllUsers();
		}

		if (nextProps.users) {
			nextProps.setLoading(false);
		}

		return prevState;
	}

	render() {
		const { users } = this.props;

		if (!users) {
			return null;
		}

		return (
			<>
				{this.props.isLoading && <Spinner />}

				<div>
					<Heading headingStyle={'title'}>Users</Heading>
				</div>

				<div>
					{users.map((user) => (
						<UserInfo key={user.id}>
							<h1>{user.username}</h1>

							<div>
								{this.props.self.accessLevel >
									user.accessLevel && (
									<>
										<Button color="primary" size="normal">
											Set Access Level
										</Button>

										<Button color="primary" size="normal">
											Change Password
										</Button>

										<Button
											color="danger"
											size="normal"
											disabled={!user.activated}
										>
											Deactivate
										</Button>
									</>
								)}
							</div>
						</UserInfo>
					))}
				</div>
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
});

export default connect(mapStateToProps, mapDispatchToProps)(Users);
