import React, { Component } from 'react';
import { connect } from 'react-redux';
import styled, { css } from 'styled-components';
import Heading from '../../components/Heading';
import respondTo from '../../mixins/respondTo';
import TextInput from '../../components/TextInput';
import Button from '../../components/Button';
import { getCurrentWebsocket } from '../../websocket/websocket';
import Alert from '../../components/Alert';
import { position } from 'polished';

const ChatWindow = styled.form`
	${(props) => css`
		background-color: ${props.theme.colorBackgroundDark};
		height: clamp(20vh, 50vh, 60vh);
		display: flex;
		flex-direction: column;
		border-radius: ${props.theme.borderRadiusNormal};
	`}
`;

const ChatContent = styled.div`
	${(props) => css`
		flex: 1;
		padding: 1rem;
		overflow-y: scroll;

		// Chrome, Safari, Opera
		::-webkit-scrollbar {
			display: none;
		}

		-ms-overflow-style: none; // Edge
		scrollbar-width: none; // Firefox
	`}
`;

const ChatBox = styled.input`
	${(props) => css`
		height: 4rem;
		border: 0;
		font-size: 1.6rem;
		border-top: 2px solid ${props.theme.colorBackground};
		background-color: ${props.theme.colorBackgroundDark};
		color: ${props.theme.colorTextSecondary};
		padding-left: 1rem;

		outline: none;
	`}
`;

const CustomAlert = styled(Alert)`
	position: absolute;
	top: 0;
`;

class Chat extends Component {
	constructor(props) {
		super(props);

		this.state = {
			server: null,
			messages: [],
			lastSend: 0,
			error: false,
			alertShown: false,
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		// Get server data
		const id = parseInt(nextProps.match.params.id);
		if (!id) {
			return prevState;
		}

		const { servers } = nextProps;

		if (!servers || !servers[id]) {
			return prevState;
		}

		prevState.server = servers[id];

		if (nextProps.chat && nextProps.chat[id]) {
			prevState.messages = nextProps.chat[id];
		}

		return prevState;
	}

	onMessageChange = (e) => {
		this.setState((prevState) => ({
			...prevState,
			message: e.target.value,
		}));
	};

	sendMessage = (e) => {
		e.preventDefault();

		const SEND_INTERVAL_MS = 500;
		const { server, message, lastSend, error, alertShown } = this.state;

		if (!message || message.trim().length === 0) {
			return;
		}

		if (
			error !== null &&
			Date.now() - lastSend <= SEND_INTERVAL_MS &&
			!alertShown
		) {
			this.setState((prevState) => ({
				...prevState,
				error: 'You are sending messages too fast',
				alertShown: true,
			}));

			setTimeout(() => {
				this.setState((prevState) => ({
					...prevState,
					error: false,
					alertShown: false,
				}));
			}, 1000);
		}

		const wsClient = getCurrentWebsocket();

		wsClient.send(
			JSON.stringify({
				type: 'chat',
				body: {
					serverId: server.id,
					message: message,
				},
			})
		);

		this.setState((prevState) => ({
			...prevState,
			lastSend: Date.now(),
		}));
	};

	render() {
		const { server, messages, error } = this.state;

		if (!server) {
			return <Heading headingStyle={'title'}>Server not found</Heading>;
		}

		return (
			<>
				<div>
					<Heading headingStyle={'title'}>
						{server.name}: Chat
					</Heading>
				</div>

				<div>
					<CustomAlert type={'error'} message={error} />
					<ChatWindow onSubmit={this.sendMessage}>
						<ChatContent>
							{messages.map((msg, i) => (
								<p key={`msg${i}`}>{JSON.stringify(msg)}</p>
							))}
						</ChatContent>
						<ChatBox
							placeholder={'Type a message and hit enter'}
							name={'message'}
							onChange={this.onMessageChange}
						/>
					</ChatWindow>
				</div>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	servers: state.servers,
	chat: state.chat,
});

const mapDispatchToProps = (dispatch) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(Chat);
