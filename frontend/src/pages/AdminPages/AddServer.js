import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';
import Alert from '../../components/Alert';
import TextInput from '../../components/TextInput';
import Button from '../../components/Button';
import { createServer } from '../../redux/servers/serverActions';

const Form = styled.div`
	display: grid;
	grid-template-columns: 1fr 1fr;
	grid-template-rows: 6rem 6rem;
	grid-gap: 1rem;
`;

const InputWrapper = styled.div`
	flex: 0 0 50%;
`;

class AddServer extends Component {
	constructor(props) {
		super(props);

		this.state = {
			errors: {},
			buttonDisabled: false,
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		// Check for success
		if (nextProps.success) {
			return {
				...prevState,
				success: nextProps.success,
				errors: {},
				buttonDisabled: true,
			};
		}

		// Check for errors
		if (nextProps.errors) {
			return {
				...prevState,
				errors: {
					...nextProps.errors,
				},
				success: {},
			};
		}
	}

	onChange = (e) => {
		this.setState({
			[e.target.name]: e.target.value,
		});
	};

	onAddServerClick = () => {
		const { name, address, rconPort, rconPassword } = this.state;

		this.props.addServer({
			name,
			address,
			rconPort,
			rconPassword,
		});
	};

	render() {
		const { errors, success } = this.state;

		return (
			<>
				<div>
					<Heading headingStyle={'title'}>Add a new server</Heading>
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
								type={'text'}
								name={'name'}
								placeholder={'Server Name'}
								size={'small'}
								onChange={this.onChange}
								value={this.state.name}
								error={this.state.errors.name}
							/>
						</InputWrapper>

						<InputWrapper>
							<TextInput
								type={'text'}
								name={'address'}
								placeholder={'Server Address'}
								size={'small'}
								onChange={this.onChange}
								value={this.state.address}
								error={this.state.errors.address}
							/>
						</InputWrapper>

						<InputWrapper>
							<TextInput
								type={'password'}
								name={'rconPassword'}
								placeholder={'RCON Password'}
								size={'small'}
								onChange={this.onChange}
								value={this.state.rconPassword}
								error={this.state.errors.rconPassword}
							/>
						</InputWrapper>

						<InputWrapper>
							<TextInput
								type={'text'}
								name={'rconPort'}
								placeholder={'RCON Port'}
								size={'small'}
								onChange={this.onChange}
								value={this.state.rconPort}
								error={this.state.errors.rconPort}
							/>
						</InputWrapper>

						<Button
							size={'normal'}
							color={'primary'}
							onClick={this.onAddServerClick}
							disabled={this.state.buttonDisabled}
						>
							Save
						</Button>
					</Form>
				</div>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	errors: state.error.createserver,
	success: state.success.createserver,
});

const mapDispatchToProps = (dispatch) => ({
	createServer: (data) => dispatch(createServer(data)),
});

export default connect(mapStateToProps, mapDispatchToProps)(AddServer);
