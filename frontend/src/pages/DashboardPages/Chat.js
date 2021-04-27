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
import styled, { css } from 'styled-components';
import Heading from '../../components/Heading';
import respondTo from '../../mixins/respondTo';
import TextInput from '../../components/TextInput';
import Button from '../../components/Button';
import { getCurrentWebsocket } from '../../websocket/websocket';
import Alert from '../../components/Alert';
import { position } from 'polished';
import { animateScroll } from 'react-scroll';
import { addChatMessage } from '../../redux/chat/chatActions';

const ChatWindow = styled.form`
	${(props) => css`
		height: clamp(20vh, 50vh, 60vh);
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		border-radius: ${props.theme.borderRadiusNormal};
	`}
`;

const ChatContentWrapper = styled.div`
	overflow-y: scroll;
	position: relative;

	// Chrome, Safari, Opera
	::-webkit-scrollbar {
		display: none;
	}

	-ms-overflow-style: none; // Edge
	scrollbar-width: none; // Firefox
`;

const ChatContent = styled.div`
	${(props) => css`
		display: flex;
		flex-direction: column;
		overflow-y: scroll;
		font-size: 1.4rem;
		position: relative;

		// Chrome, Safari, Opera
		::-webkit-scrollbar {
			display: none;
		}

		-ms-overflow-style: none; // Edge
		scrollbar-width: none; // Firefox

		> :nth-child(even) {
			background-color: ${props.theme.colorBackground};
		}

		> :nth-child(odd) {
			background-color: ${props.theme.colorAccent};
		}
	`}
`;

const ChatMessage = styled.div`
	${(props) => css`
		width: 100%;
		display: flex;
		flex-direction: column;
		margin-bottom: 1rem;
		padding: 0 1rem;

		${respondTo.medium`
	  		flex-direction: row;
		  	margin-bottom: 0;
		`}
	`}
`;

const ChatMessageName = styled.div`
	${(props) => css`
		color: ${props.theme.colorTextPrimary};
		width: 100%;
		font-weight: 500;
		margin-right: 1rem;

		${respondTo.medium`
        	width: 15%;
		`}

		${props.ownMessage
			? `border-right: 4px solid ${props.theme.colorPrimary};`
			: null}

        ${props.sentByUser
			? `border-right: 4px solid ${props.theme.colorAlert};`
			: null}
	`}
`;

const ChatBox = styled.input`
	${(props) => css`
		height: 4rem;
		border: 0;
		font-size: 1.6rem;
		color: ${props.theme.colorTextPrimary};
		padding: 2rem 1rem;
		background-color: ${props.theme.inputs.fillInBackground
			? props.theme.colorBorderPrimary
			: props.theme.colorBackgroundDark};

		outline: none;
	`}
`;

const CustomAlert = styled(Alert)`
	position: absolute;
	top: 0;
`;

const ScrollToBottomButton = styled.div`
	${(props) => css`
		position: sticky;
		bottom: 0;
		width: 80%;
		margin-left: 10%;
		margin-right: 10%;
		height: 2rem;
		text-align: center;
		border-top-left-radius: 1rem;
		border-top-right-radius: 1rem;
		background-color: ${props.theme.colorPrimaryDark};
		font-size: 1.2rem;
		user-select: none;

		> :hover {
			cursor: pointer;
		}

		${respondTo.medium`
          	width: 30%;
          	margin-left: 35%;
          	margin-right: 35%;
	  	`}
	`}
`;

const Legend = styled.div`
	${(props) => css`
		display: flex;
		font-size: 1.2rem;
		line-height: 1.2rem;
		margin-top: 1rem;

		> * {
			margin-right: 2rem;
		}

		> :nth-child(1) {
			border-left: 4px solid ${props.theme.colorPrimary};
		}

		> :nth-child(2) {
			border-left: 4px solid ${props.theme.colorAlert};
		}
	`}
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
			lastScrollPos: 0,
		};

		this.scrollRef = React.createRef();
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		// Get server data
		const id = parseInt(nextProps.match.params.id);
		if (!id) {
			return prevState;
		}

		const { servers, games } = nextProps;

		if (!servers || !servers[id] || !games) {
			return prevState;
		}

		prevState.server = servers[id];

		const game = games[servers[id].game];

		if (game && game.config && !game.config.enableChat) {
			prevState.chatDisabled = true;
		}

		if (nextProps.chat && nextProps.chat[id]) {
			prevState.messages = nextProps.chat[id];
		}

		return prevState;
	}

	componentDidMount() {
		this.scrollToBottom(750)();
	}

	componentDidUpdate(prevProps, prevState, snapshot) {
		if (!this.scrollRef.current) {
			return;
		}

		if (this.state.lastScrollPos !== prevState.lastScrollPos) {
			return;
		}

		if (this.isAnchoredToBottom()) {
			this.scrollToBottom(0)();
		}
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

		this.props.addMessage({
			serverId: server.id,
			name: 'You',
			ownMessage: true,
			message: message,
		});

		this.setState((prevState) => ({
			...prevState,
			lastSend: Date.now(),
			message: '',
		}));
	};

	onScroll = (e) => {
		const { lastScrollPos, scrollDirection } = this.state;
		const currentTarget = e.currentTarget;

		if (!currentTarget) {
			return;
		}

		let direction = scrollDirection;
		if (lastScrollPos > currentTarget.scrollTop) {
			// Scrolling up
			direction = 'up';
		} else if (lastScrollPos < currentTarget.scrollTop) {
			// Scrolling down
			direction = 'down';
		}

		this.setState((prevState) => ({
			...prevState,
			scrollDirection: direction,
			lastScrollPos: currentTarget.scrollTop,
		}));
	};

	scrollToBottom = (ms) => () => {
		animateScroll.scrollToBottom({
			containerId: 'chat-content',
			duration: ms,
		});
	};

	isAnchoredToBottom = () => {
		const { lastScrollPos } = this.state;

		let isAnchored = true;
		if (this.scrollRef.current) {
			const scrollheight = this.scrollRef.current.scrollHeight;
			const offsetHeight = this.scrollRef.current.offsetHeight;
			const tolerance = 50;
			isAnchored = !(
				lastScrollPos <
				scrollheight - offsetHeight - tolerance
			);
		}

		return isAnchored;
	};

	render() {
		const { server, messages, error, message, chatDisabled } = this.state;
		if (!server) {
			return <Heading headingStyle={'title'}>Server not found</Heading>;
		}

		if (chatDisabled) {
			return (
				<Heading headingStyle={'title'}>
					Live chat is not enabled for this game
				</Heading>
			);
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
						<ChatContentWrapper
							onScroll={this.onScroll}
							ref={this.scrollRef}
							id={'chat-content'}
						>
							<ChatContent>
								{messages.map((msg, i) => {
									if (
										msg.sentByUser &&
										msg.name === this.props.self.username
									) {
										return null;
									}

									return (
										<ChatMessage key={`chatmsg${i}`}>
											<ChatMessageName
												sentByUser={!!msg.sentByUser}
												ownMessage={!!msg.ownMessage}
											>
												{msg.name}
											</ChatMessageName>
											{msg.message}
										</ChatMessage>
									);
								})}
							</ChatContent>
							{!this.isAnchoredToBottom() && (
								<ScrollToBottomButton
									onClick={this.scrollToBottom(750)}
								>
									Jump to bottom
								</ScrollToBottomButton>
							)}
						</ChatContentWrapper>
						<ChatBox
							placeholder={'Type a message and hit enter'}
							name={'message'}
							onChange={this.onMessageChange}
							value={message}
						/>
					</ChatWindow>
					<Legend>
						<div>= you</div>
						<div>= other user</div>
					</Legend>
				</div>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	self: state.user.self,
	servers: state.servers,
	games: state.games,
	chat: state.chat,
});

const mapDispatchToProps = (dispatch) => ({
	addMessage: (messageBody) => dispatch(addChatMessage(messageBody)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Chat);
