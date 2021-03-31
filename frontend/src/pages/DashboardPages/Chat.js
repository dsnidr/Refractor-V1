import React, { Component } from 'react';
import { connect } from 'react-redux';
import styled, { css } from 'styled-components';
import Heading from '../../components/Heading';
import respondTo from '../../mixins/respondTo';
import TextInput from '../../components/TextInput';

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

class Chat extends Component {
	constructor(props) {
		super(props);

		this.state = {
			server: null,
			messages: [],
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

	onMessageSend = (e) => {
		e.preventDefault();

		console.log('Sending message');
	};

	render() {
		const { server, messages } = this.state;

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
					<ChatWindow onSubmit={this.onMessageSend}>
						<ChatContent>
							{messages.map((msg, i) => (
								<p key={`msg${i}`}>{JSON.stringify(msg)}</p>
							))}
						</ChatContent>
						<ChatBox placeholder={'Type a message and hit enter'} />
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
