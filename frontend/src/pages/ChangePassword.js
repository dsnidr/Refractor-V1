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
import Wrapper from '../components/Wrapper';
import styled, { css } from 'styled-components';
import respondTo from '../mixins/respondTo';
import { Field, Form, Formik } from 'formik';
import * as Yup from 'yup';
import { connect } from 'react-redux';
import { Redirect } from 'react-router';
import { setLoading } from '../redux/loading/loadingActions';
import Alert from '../components/Alert';
import TextInput from '../components/TextInput';
import Button from '../components/Button';
import { changePassword } from '../redux/user/userActions';

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
		padding-bottom: 3rem;

		${props.theme.boxShadowPrimary}
		${respondTo.medium`
      		width: 50rem;
      		min-height: 40rem;
      		height: auto;
    	`}
    	> * {
			width: 80%;
		}
	`}
`;

const HeadingBox = styled.div`
	${(props) => css`
		height: 10rem;
		font-size: 2.8rem;
		display: flex;
		justify-content: center;
		align-items: center;
		font-weight: 500;
		color: ${props.theme.colorTextPrimary};
	`}
`;

const initialValues = {
	currentPassword: '',
	newPassword: '',
	newPasswordConfirm: '',
};

class ChangePassword extends Component {
	constructor(props) {
		super(props);

		this.state = {
			errors: {},
			success: {},
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		if (nextProps.success) {
			const {
				history: { push },
			} = nextProps;

			setTimeout(() => push('/'), 1500);
		}

		if (nextProps.errors) {
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

		// Try to change password if no input level errors occurred
		this.props.changePassword(values);

		// Make the user wait a second before retrying. This gives them the chance to read error
		// messages rather than simply spamming the reset button, as we know they will...
		setTimeout(() => {
			actions.setSubmitting(false);
		}, 1500);
	};

	render() {
		return (
			<Wrapper>
				<Formik
					initialValues={initialValues}
					validationSchema={Yup.object({
						currentPassword: Yup.string().required(
							'Please enter your current password'
						),
						newPassword: Yup.string()
							.required('Please enter a new password')
							.min(8, 'That password is too short')
							.max(80, 'That password is too long')
							.notOneOf(
								[Yup.ref('currentPassword'), null],
								"You can't re-use your current password"
							),
						newPasswordConfirm: Yup.string()
							.required('Please re-enter your new password')
							.oneOf(
								[Yup.ref('newPassword'), null],
								'Passwords must match'
							),
					})}
					onSubmit={(values, actions) =>
						this.submitForm(values, actions)
					}
				>
					{({ errors, isSubmitting, handleReset, handleSubmit }) => (
						<FormWrapper
							onReset={handleReset}
							onSubmit={handleSubmit}
						>
							<HeadingBox>CHANGE PASSWORD</HeadingBox>
							<Alert type={'error'} message={this.props.errors} />
							<Alert
								type={'success'}
								message={this.props.success}
							/>

							<Field name={'currentPassword'}>
								{({ field }) => (
									<TextInput
										{...field}
										type={'password'}
										placeholder={'Current password'}
										error={
											errors.currentPassword
												? errors.currentPassword
												: this.state.errors
														.currentPassword
										}
									/>
								)}
							</Field>

							<Field name={'newPassword'}>
								{({ field }) => (
									<TextInput
										{...field}
										type={'password'}
										placeholder={'New password'}
										error={errors.newPassword}
									/>
								)}
							</Field>

							<Field name={'newPasswordConfirm'}>
								{({ field }) => (
									<TextInput
										{...field}
										type={'password'}
										placeholder={'Re-enter new password'}
										error={errors.newPasswordConfirm}
									/>
								)}
							</Field>

							<Button type={'submit'} disabled={isSubmitting}>
								CHANGE PASSWORD
							</Button>
						</FormWrapper>
					)}
				</Formik>
			</Wrapper>
		);
	}
}

const mapStateToProps = (state) => ({
	success: state.success.changepassword,
	errors: state.error.changepassword,
});

const mapDispatchToProps = (dispatch) => ({
	setLoading: (isLoading) => dispatch(setLoading(isLoading)),
	changePassword: (data) => dispatch(changePassword(data)),
});

export default connect(mapStateToProps, mapDispatchToProps)(ChangePassword);
