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
import Button from '../../components/Button';
import ReactTooltip from 'react-tooltip';

const Header = styled.div`
	button {
		margin-top: 1rem;
	}
`;

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

		${props.watched ? `color: ${props.theme.colorDanger}` : ``}
	`}
`;

const PlayerHeader = styled.div`
	${(props) => css`
		position: relative;

		h1 {
			padding: 1rem;
			font-weight: 400;
			font-size: 1.7rem;
			flex: 0 0 90%;

			${respondTo.medium`
          padding: 1.5rem;
		`}

			${props.watched ? `color: ${props.theme.colorDanger}` : ''}
		}

		span {
			position: absolute;
			top: 1rem;
			right: 1rem;
			font-size: 1.5rem;
			background-color: ${props.theme.colorBackgroundAlt};
			width: 2rem;
			height: 2rem;
			border-radius: ${props.theme.borderRadiusNormal};
			color: ${props.theme.colorTextLight};

			display: flex;
			align-items: center;
			justify-content: center;

			${respondTo.medium`
          top: calc(50% - 1rem);
		`}
		}

		:hover {
			cursor: pointer;
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

const ServerButtonWrapper = styled.div`
	display: flex;

	> * {
		margin-right: 1rem;
	}
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

		const { servers, games } = nextProps;

		if (!servers || !servers[id] || !games) {
			return prevState;
		}

		prevState.server = servers[id];
		prevState.game = games[servers[id].game];
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

	onOpenChatClick = () => {
		const { server } = this.state;

		this.props.history.push(`/server/${server.id}/chat`);
	};

	onPlayerClick = (playerId) => () => {
		this.props.history.push(`/player/${playerId}`);
	};

	onViewChatRecordsClick = () => {};

	render() {
		const { server, modals, game } = this.state;
		const { warn, kick, ban } = modals;

		if (!server) {
			return null;
		}

		let players = [];
		if (server.players) {
			players = Object.values(server.players);
		}

		return (
			<>
				<WarnModal
					player={warn.ctx}
					serverId={server.id}
					show={warn.show}
					onClose={this.closeModal('warn')}
					inputRef={this.warnModalRef}
					reload={false}
				/>

				<KickModal
					player={kick.ctx}
					serverId={server.id}
					show={kick.show}
					onClose={this.closeModal('kick')}
					inputRef={this.kickModalRef}
					reload={false}
				/>

				<BanModal
					player={ban.ctx}
					serverId={server.id}
					show={ban.show}
					onClose={this.closeModal('ban')}
					inputRef={this.banModalRef}
					reload={false}
				/>

				<ReactTooltip delayShow={100} />

				<Header>
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
					{game.config.enableChat && (
						<ServerButtonWrapper>
							<Button
								size={'small'}
								onClick={this.onOpenChatClick}
							>
								Open Chat
							</Button>
						</ServerButtonWrapper>
					)}
				</Header>

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
						{players.map((player) => {
							return (
								<Player watched={player.watched}>
									<PlayerHeader
										onClick={this.onPlayerClick(player.id)}
									>
										<h1>{player.currentName}</h1>
										{player.infractionCount > 0 && (
											<span
												data-tip={
													'The number of infractions this player has'
												}
											>
												{player.infractionCount}
											</span>
										)}
									</PlayerHeader>
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
							);
						})}
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
