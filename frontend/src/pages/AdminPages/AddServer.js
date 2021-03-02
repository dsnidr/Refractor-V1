import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';
import Alert from '../../components/Alert';
import TextInput from '../../components/TextInput';
import Button from '../../components/Button';
import { createServer } from '../../redux/servers/serverActions';
import Select from '../../components/Select';

const InputWrapper = styled.div`
	flex: 0 0 50%;
`;

const TopInputWrapper = styled.div`
	flex: 0 0 100%;
`;

const Form = styled.div`
	display: grid;
	grid-template-columns: 1fr 1fr;
	grid-template-rows: 6rem 6rem 6rem;
	grid-gap: 1rem;

	${TopInputWrapper} {
		grid-row: 1;
		grid-column: span 2;

		select {
			height: 4rem;
		}
	}
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
		// Set the default game value
		if (nextProps.games) {
			prevState.game = Object.values(nextProps.games)[0].name;
		}

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
		const { game, name, address, rconPort, rconPassword } = this.state;

		this.props.createServer({
			game,
			name,
			address,
			rconPort,
			rconPassword,
		});
	};

	render() {
		const { errors, success } = this.state;

		const games = Object.values(this.props.games);

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
						<TopInputWrapper>
							<Select
								name={'game'}
								onChange={this.onChange}
								error={errors.game}
							>
								{games.map((game) => (
									<option key={game.id} value={game.name}>
										{game.name}
									</option>
								))}
							</Select>
						</TopInputWrapper>

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
							Add Server
						</Button>
					</Form>
				</div>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	games: state.games,
	errors: state.error.createserver,
	success: state.success.createserver,
});

const mapDispatchToProps = (dispatch) => ({
	createServer: (data) => dispatch(createServer(data)),
});

export default connect(mapStateToProps, mapDispatchToProps)(AddServer);
