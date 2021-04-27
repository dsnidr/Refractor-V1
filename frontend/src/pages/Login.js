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
import { connect } from 'react-redux';
import respondTo from '../mixins/respondTo';
import styled, { css } from 'styled-components';
import Wrapper from '../components/Wrapper';
import { Field, Form, Formik } from 'formik';
import * as Yup from 'yup';
import TextInput from '../components/TextInput';
import Alert from '../components/Alert';
import { ReactComponent as Avatar } from '../assets/avatar.svg';
import { ReactComponent as Padlock } from '../assets/padlock.svg';
import Button from '../components/Button';
import { logIn } from '../redux/user/userActions';
import { setLoading } from '../redux/loading/loadingActions';
import { Redirect } from 'react-router';
import Spinner from '../components/Spinner';

const FormWrapper = styled.form`
	${(props) => css`
		border: 1px solid ${props.theme.colorBorderSecondary};
		border-radius: ${props.theme.borderRadiusNormal};
		background: ${props.theme.colorBackground};
		width: 100%;
		min-height: 100vh;
		display: flex;
		flex-direction: column;
		align-items: center;

		${props.theme.boxShadowPrimary}

		${respondTo.medium`
      		width: 50rem;
      		min-height: 40rem;
    	`}
		
		> * {
			width: 80%;
		}
	`}
`;

const HeadingBox = styled.div`
	${(props) => css`
		height: 10rem;
		font-size: 3.2rem;
		display: flex;
		justify-content: center;
		align-items: center;
		font-weight: 500;
		color: ${props.theme.colorTextPrimary};
	`}
`;

const initialValues = {
	username: '',
	password: '',
};

class Login extends Component {
	constructor(props) {
		super(props);

		this.state = {
			errors: {},
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		if (nextProps.success) {
			nextProps.setLoading(false);

			return {
				...prevState,
				success: nextProps.success,
				errors: {},
			};
		}

		if (nextProps.errors) {
			nextProps.setLoading(false);

			return {
				...prevState,
				errors:
					typeof nextProps.errors === 'string'
						? {
								...prevState.errors,
								general: nextProps.errors,
						  }
						: {
								...prevState.errors,
								...nextProps.errors,
						  },
				success: {},
			};
		}

		return prevState;
	}

	submitForm = (values, actions) => {
		this.props.setLoading(true);

		// Try to log in if no input level errors occurred
		this.props.logIn({
			...values,
		});

		// Make the user wait a second before retrying. This gives them the chance to read error
		// messages rather than simply spamming the log in button, as we know they will...
		setTimeout(() => {
			actions.setSubmitting(false);
		}, 750);
	};

	render() {
		console.log(this.state);
		return (
			<Wrapper>
				{this.props.isLoading && <Spinner />}
				{this.props.success && <Redirect to={'/'} />}

				<Formik
					initialValues={initialValues}
					validationSchema={Yup.object({
						username: Yup.string().required(
							'Please enter your username'
						),
						password: Yup.string()
							.required('Please enter your password')
							.min(8, 'This password is too short')
							.max(80, 'This password is too long'),
					})}
					onSubmit={(values, actions) =>
						this.submitForm(values, actions)
					}
				>
					{({
						values,
						errors,
						isSubmitting,
						handleReset,
						handleSubmit,
					}) => {
						return (
							<FormWrapper
								onReset={handleReset}
								onSubmit={handleSubmit}
							>
								<HeadingBox>LOG IN</HeadingBox>
								<Alert
									type="error"
									message={this.state.errors.general}
								/>

								<Field name="username">
									{({ field }) => (
										<TextInput
											{...field}
											type="text"
											icon={<Avatar />}
											placeholder="Username"
											error={errors.username}
										/>
									)}
								</Field>

								<Field name="password">
									{({ field }) => (
										<TextInput
											{...field}
											type="password"
											icon={<Padlock />}
											placeholder="Password"
											error={errors.password}
										/>
									)}
								</Field>

								<Button
									color="primary"
									disabled={isSubmitting}
									type={'submit'}
								>
									LOG IN
								</Button>
							</FormWrapper>
						);
					}}
				</Formik>
			</Wrapper>
		);
	}
}

const mapStateToProps = (state) => ({
	errors: state.error.auth,
	isLoading: state.loading.login,
});

const mapDispatchToProps = (dispatch) => ({
	logIn: (credentials) => dispatch(logIn(credentials)),
	setLoading: (isLoading) => dispatch(setLoading('login', isLoading)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Login);
