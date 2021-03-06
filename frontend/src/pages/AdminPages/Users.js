import React, { Component } from 'react';
import Heading from '../../components/Heading';
import { setLoading } from '../../redux/loading/loadingActions';
import { connect } from 'react-redux';
import { getAllUsers } from '../../redux/user/userActions';
import Spinner from '../../components/Spinner';
import styled, { css } from 'styled-components';
import Button from '../../components/Button';
import { Link } from 'react-router-dom';

const UserInfo = styled(Link)`
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

class Users extends Component {
	static getDerivedStateFromProps(nextProps, prevState) {
		if (!nextProps.users) {
			nextProps.setLoading(true);
			nextProps.getAllUsers();
		} else if (nextProps.users && nextProps.isLoading) {
			nextProps.setLoading(false);
		}

		return prevState;
	}

	render() {
		const { users: usersObj } = this.props;

		if (!usersObj) {
			return null;
		}

		const users = Object.values(usersObj);

		return (
			<>
				<div>
					<Heading headingStyle={'title'}>Users</Heading>
				</div>

				<div>
					{users.map((user) => (
						<UserInfo key={user.id} to={`/user/${user.id}`}>
							<h1>{user.username}</h1>
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
