import React, { Component } from 'react';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';
import styled, { css } from 'styled-components';
import respondTo from '../../mixins/respondTo';
import { Link } from 'react-router-dom';
import StatusTag from '../../components/StatusTag';
import RequirePerms from '../../components/RequirePerms';
import { flags } from '../../permissions/permissions';
import WarnModal from '../../components/modals/WarnModal';
import KickModal from '../../components/modals/KickModal';
import BanModal from '../../components/modals/BanModal';

const ServerSummary = styled.div`
	${(props) => css`
		display: flex;
		flex-direction: column;
		font-size: 1.4rem;

		> * {
			margin-right: 2rem;
		}

		${respondTo.medium`
      flex-direction: row;
    `}
	`}
`;

const InfoSpan = styled.span`
	${(props) => css`
		color: ${props.theme.colorTextPrimary};
	`}
`;

const PlayerList = styled.div`
	${(props) => css`
		margin-top: 1rem;
		display: grid;
		grid-template-columns: auto;
		grid-auto-rows: auto;
		grid-row-gap: 1rem;

		${respondTo.medium`
      		grid-template-columns: 1fr 1fr 1fr;
      		grid-gap: 2rem;
		`}

		${respondTo.large`
            grid-template-columns: 1fr 1fr 1fr 1fr;
            grid-gap: 2rem;
        `}
		
    	${respondTo.extralarge`
            grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
            grid-gap: 2rem;
        `}
	`}
`;

const Player = styled.div`
	${(props) => css`
		background-color: ${props.theme.colorAccent};
		border-radius: ${props.theme.borderRadiusNormal};
		white-space: nowrap;
		text-overflow: ellipsis;
		overflow: hidden;
		grid-row: auto;

		a {
			color: ${props.theme.colorTextSecondary} !important;
			text-decoration: none !important;
		}

		h1 {
			padding: 1rem;
			font-weight: 400;
			font-size: 1.7rem;

			${respondTo.medium`
        		padding: 1.5rem;
      		`}
		}
	`}
`;

const PlayerButtons = styled.div`
	${(props) => css`
		display: flex;
		height: 2rem;

		${respondTo.medium`
      		height: 3rem;
    	`}

		> * {
			flex: 1;
			display: flex;
			align-items: center;
			justify-content: center;
			user-select: none;
			border: 1px solid ${props.theme.colorBackground};
			color: ${props.theme.colorPrimary};
			font-size: 1.4rem;

			:hover {
				cursor: pointer;
			}
		}

		> *:nth-child(1) {
			border-bottom-left-radius: ${props.theme.borderRadiusNormal};
			border-right: none;

			:hover {
				background-color: ${props.theme.colorWarning};
				color: ${props.theme.colorTextWarning};
			}
		}

		> *:nth-child(2) {
			:hover {
				background-color: ${props.theme.colorAlert};
				color: ${props.theme.colorTextAlert};
			}
		}

		> *:nth-child(3) {
			border-left: none;
			border-bottom-right-radius: ${props.theme.borderRadiusNormal};

			:hover {
				background-color: ${props.theme.colorDanger};
				color: ${props.theme.colorTextDanger};
			}
		}
	`}
`;

class Server extends Component {
	constructor(props) {
		super(props);

		this.state = {
			modals: {
				warn: {
					show: false,
					ctx: {},
				},
				kick: {
					show: false,
					ctx: {},
				},
				ban: {
					show: false,
					ctx: {},
				},
			},
		};

		this.warnModalRef = React.createRef();
		this.kickModalRef = React.createRef();
		this.banModalRef = React.createRef();
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
		return prevState;
	}

	showModal = (type, ctx) => () => {
		this.setState(
			(prevState) => ({
				...prevState,
				modals: {
					...prevState.modals,
					[type]: {
						...prevState.modals[type],
						show: true,
						ctx: ctx,
					},
				},
			}),
			() => {
				const ref = this[`${type}ModalRef`];
				if (ref.current) {
					ref.current.focus();
				}
			}
		);
	};

	closeModal = (type) => () => {
		this.setState((prevState) => ({
			...prevState,
			modals: {
				...prevState.modals,
				[type]: {
					...prevState.modals[type],
					show: false,
					ctx: {},
				},
			},
		}));
	};

	render() {
		const { server, modals } = this.state;
		const { warn, kick, ban } = modals;

		if (!server) {
			return null;
		}

		const players = server.players || [];

		return (
			<>
				<WarnModal
					player={warn.ctx}
					serverId={server.id}
					show={warn.show}
					onClose={this.closeModal('warn')}
					inputRef={this.warnModalRef}
				/>

				<KickModal
					player={kick.ctx}
					show={kick.show}
					onClose={this.closeModal('kick')}
					inputRef={this.kickModalRef}
				/>

				<BanModal
					player={ban.ctx}
					show={ban.show}
					onClose={this.closeModal('ban')}
					inputRef={this.banModalRef}
				/>

				<div>
					<Heading headingStyle={'title'}>{server.name}</Heading>
					<ServerSummary>
						<p>
							<InfoSpan>{`Players: `}</InfoSpan> {players.length}
						</p>
						<p>
							<InfoSpan>{`Status: `}</InfoSpan>{' '}
							<StatusTag status={server.online} />
						</p>
						<p>
							<InfoSpan>{`Address: `}</InfoSpan> {server.address}
						</p>
					</ServerSummary>
				</div>

				<div>
					{players.length > 0 ? (
						<Heading headingStyle="subtitle">
							Online players:
						</Heading>
					) : (
						<Heading headingStyle="subtitle">
							No players are online
						</Heading>
					)}

					<PlayerList>
						{players.map((player) => (
							<Player>
								<Link to={`/player/${player.id}`}>
									<h1>{player.currentName}</h1>
								</Link>
								<PlayerButtons>
									<RequirePerms
										mode={'all'}
										perms={[flags.LOG_WARNING]}
									>
										<div
											onClick={this.showModal(
												'warn',
												player
											)}
										>
											Warn
										</div>
									</RequirePerms>
									<RequirePerms
										mode={'all'}
										perms={[flags.LOG_KICK]}
									>
										<div
											onClick={this.showModal(
												'kick',
												player
											)}
										>
											Kick
										</div>
									</RequirePerms>
									<RequirePerms
										mode={'all'}
										perms={[flags.LOG_BAN]}
									>
										<div
											onClick={this.showModal(
												'ban',
												player
											)}
										>
											Ban
										</div>
									</RequirePerms>
								</PlayerButtons>
							</Player>
						))}
					</PlayerList>
				</div>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	servers: state.servers,
	games: state.games,
});

const mapDispatchToProps = (dispatch) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(Server);
