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

		${respondTo.medium`
        	width: 15%;
		`}
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

		${respondTo.medium`
          	width: 30%;
          	margin-left: 35%;
          	margin-right: 35%;
	  	`}
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
		const { server, messages, error, message, lastScrollPos } = this.state;
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
						<ChatContentWrapper
							onScroll={this.onScroll}
							ref={this.scrollRef}
							id={'chat-content'}
						>
							<ChatContent>
								{messages.map((msg, i) => (
									<ChatMessage>
										<ChatMessageName>
											{msg.name}
										</ChatMessageName>
										{msg.message}
									</ChatMessage>
								))}
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
