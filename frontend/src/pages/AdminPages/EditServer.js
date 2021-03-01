import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';
import Alert from '../../components/Alert';
import TextInput from '../../components/TextInput';
import Button from '../../components/Button';
import { updateServer } from '../../redux/servers/serverActions';

const Form = styled.div`
	display: grid;
	grid-template-columns: 1fr 1fr;
	grid-template-rows: 6rem 6rem;
	grid-gap: 1rem;
`;

const InputWrapper = styled.div`
	flex: 0 0 50%;
`;

class EditServer extends Component {
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
			setTimeout(() => {
				document.location.href = '/servers';
			}, 4000);

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

		// Get server data if no success or errors were found
		if (!prevState.server) {
			const id = parseInt(nextProps.match.params.id);
			if (!id) {
				return prevState;
			}

			const { servers } = nextProps;

			if (!servers || !servers[id]) {
				return prevState;
			}

			let server = servers[id];

			return {
				...prevState,
				server,
				name: server.name,
				address: server.address,
				rconPort: server.rconPort,
				rconPassword: '',
			};
		}
	}

	onChange = (e) => {
		this.setState({
			[e.target.name]: e.target.value,
		});
	};

	onSaveClick = () => {
		const { server, name, address, rconPassword, rconPort } = this.state;

		const errors = {};

		if (!name) {
			errors.name = "Please enter the server's name";
		}

		if (!address) {
			errors.address = "Please enter the server's address";
		}

		if (!rconPassword) {
			errors.rconPassword = "Please enter the server's RCON password";
		}

		if (!rconPort) {
			errors.rconPort = "Please enter the server's RCON port";
		}

		if (Object.keys(errors).length > 0) {
			return this.setState((prevState) => ({
				...prevState,
				errors: errors,
			}));
		}

		this.props.updateServer(server.id, {
			name,
			address,
			rconPassword,
			rconPort,
		});
	};

	render() {
		const { server, errors, success } = this.state;

		if (!server) {
			return <Heading headingStyle={'title'}>Server not found</Heading>;
		}

		console.log(errors, success);

		return (
			<>
				<div>
					<Heading headingStyle={'title'}>
						Editing server: {server.name}
					</Heading>
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
							onClick={this.onSaveClick}
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
	servers: state.servers,
	errors: state.error.editserver,
	success: state.success.editserver,
});

const mapDispatchToProps = (dispatch) => ({
	updateServer: (serverId, data) => dispatch(updateServer(serverId, data)),
});

export default connect(mapStateToProps, mapDispatchToProps)(EditServer);
