import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';
import Alert from '../../components/Alert';
import TextInput from '../../components/TextInput';
import Button from '../../components/Button';
import { addUser } from '../../redux/user/userActions';

const Form = styled.div`
	display: grid;
	grid-template-columns: 1fr 1fr;
	grid-template-rows: 6rem 6rem;
	grid-gap: 1rem;
`;

const InputWrapper = styled.div`
	flex: 0 0 50%;
`;

class AddUser extends Component {
	constructor(props) {
		super(props);

		this.state = {
			errors: {},
			buttonDisabled: false,
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		if (nextProps.success) {
			setTimeout(() => {
				document.location.href = '/users';
			}, 3000);

			return {
				...prevState,
				errors: {},
				success: nextProps.success,
				buttonDisabled: true,
			};
		}

		if (nextProps.errors) {
			return {
				...prevState,
				errors: {
					...nextProps.errors,
				},
			};
		}

		return prevState;
	}

	onChange = (e) => {
		this.setState((prevState) => ({
			[e.target.name]: e.target.value,
		}));
	};

	onAddUserClicked = () => {
		const { username, email, password, passwordConfirm } = this.state;

		this.props.addUser({
			username,
			email,
			password,
			passwordConfirm,
		});
	};

	render() {
		const { errors, success, buttonDisabled } = this.state;

		return (
			<>
				<div>
					<Heading headingStyle={'title'}>Add a new user</Heading>
				</div>

				<div>
					<Alert
						type="error"
						message={typeof errors === 'string' ? errors : null}
					/>
					<Alert type="success" message={success} />

					<Form>
						<InputWrapper>
							<TextInput
								type="text"
								name="username"
								placeholder="Username"
								size="small"
								onChange={this.onChange}
								error={errors.username}
							/>
						</InputWrapper>
						<InputWrapper>
							<TextInput
								type="text"
								name="email"
								placeholder="Email Address"
								size="small"
								onChange={this.onChange}
								error={errors.email}
							/>
						</InputWrapper>
						<InputWrapper>
							<TextInput
								type="password"
								name="password"
								placeholder="Password"
								size="small"
								onChange={this.onChange}
								error={errors.password}
							/>
						</InputWrapper>
						<InputWrapper>
							<TextInput
								type="password"
								name="passwordConfirm"
								placeholder="Confirm Password"
								size="small"
								onChange={this.onChange}
								error={errors.passwordConfirm}
							/>
						</InputWrapper>

						<Button
							size="normal"
							color="primary"
							onClick={this.onAddUserClicked}
							disabled={buttonDisabled}
						>
							Add User
						</Button>
					</Form>
				</div>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	success: state.success.adduser,
	errors: state.error.adduser,
});

const mapDispatchToProps = (dispatch) => ({
	addUser: (data) => dispatch(addUser(data)),
});

export default connect(mapStateToProps, mapDispatchToProps)(AddUser);
