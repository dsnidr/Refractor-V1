import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';
import { getAllUsers } from '../../redux/user/userActions';
import { Link } from 'react-router-dom';
import TextInput from '../../components/TextInput';
import Button from '../../components/Button';

const ManagementRow = styled(Link)`
	${(props) => css``}
`;

class User extends Component {
	constructor(props) {
		super(props);

		this.state = {};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
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

		return prevState;
	}

	render() {
		const { user } = this.state;

		if (!user) {
			return (
				<div>
					<Heading headingStyle={'title'}>User not found</Heading>
				</div>
			);
		}

		return (
			<>
				<div>
					<Heading headingStyle={'title'}>
						Viewing user: {user.username}
					</Heading>
				</div>

				<div>
					<Heading headingStyle={'subtitle'}>
						Password management
					</Heading>

					<ManagementRow>
						<TextInput
							name={'newPassword'}
							placeholder={'New password'}
						/>
						<Button size={'normal'} color={'primary'}>
							Set password
						</Button>
					</ManagementRow>
				</div>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	users: state.user.others,
});

const mapDispatchToProps = (dispatch) => ({
	getAllUsers: () => dispatch(getAllUsers()),
});

export default connect(mapStateToProps, mapDispatchToProps)(User);
