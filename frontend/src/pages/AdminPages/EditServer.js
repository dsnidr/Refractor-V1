import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';
import Alert from '../../components/Alert';
import TextInput from '../../components/TextInput';
import Button from '../../components/Button';

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
		console.log('Save clicked');
	};

	render() {
		const { server } = this.state;

		if (!server) {
			return <Heading headingStyle={'title'}>Server not found</Heading>;
		}

		return (
			<>
				<div>
					<Heading headingStyle={'title'}>
						Editing server: {server.name}
					</Heading>
				</div>

				<div>
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
});

const mapDispatchToProps = (dispatch) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(EditServer);
