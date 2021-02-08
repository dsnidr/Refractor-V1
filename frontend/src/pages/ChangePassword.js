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

		this.state = {};
	}

	submitForm = (values, actions) => {
		this.props.setLoading(true);

		// Try to change password if no input level errors occurred
		// TODO: Route
		console.log('Changing password');

		// Make the user wait a second before retrying. This gives them the chance to read error
		// messages rather than simply spamming the log in button, as we know they will...
		setTimeout(() => {
			actions.setSubmitting(false);
		}, 750);
	};

	render() {
		return (
			<Wrapper>
				{this.props.success && <Redirect to={'/'} />}
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
							<Alert type={'error'} message={this.state.errors} />

							<Field name={'currentPassword'}>
								{({ field }) => (
									<TextInput
										{...field}
										type={'password'}
										placeholder={'Current password'}
										error={errors.currentPassword}
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

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => ({
	setLoading: (isLoading) => dispatch(setLoading(isLoading)),
});

export default connect(mapStateToProps, mapDispatchToProps)(ChangePassword);
