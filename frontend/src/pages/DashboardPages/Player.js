import React, { Component } from 'react';
import Heading from '../../components/Heading';
import { connect } from 'react-redux';
import {
	getPlayerSummary,
	unwatchPlayer,
	watchPlayer,
} from '../../redux/players/playerActions';
import styled, { css } from 'styled-components';
import respondTo from '../../mixins/respondTo';
import { timestampToDateTime, getTimeRemaining } from '../../utils/timeUtils';
import Button from '../../components/Button';
import Infraction from '../../components/Infraction';
import WarnModal from '../../components/modals/WarnModal';
import KickModal from '../../components/modals/KickModal';
import BanModal from '../../components/modals/BanModal';
import MuteModal from '../../components/modals/MuteModal';
import { setLoading } from '../../redux/loading/loadingActions';
import EditInfractionModal from '../../components/modals/EditInfractionModal';
import BasicModal from '../../components/modals/BasicModal';
import DeleteInfractionModal from '../../components/modals/DeleteInfractionModal';
import queryString from 'querystring';

const PlayerInfo = styled.div`
	${(props) => css`
		display: flex;
		font-size: 1.6rem;
		margin-top: 1rem;
		justify-content: space-between;
		flex-direction: column;

		${respondTo.small`
        	flex-direction: row;
      	`}

		${respondTo.large`
        	justify-content: normal;
        
        	> * {
          	margin-right: 4rem;
        	}
      	`}
	`}
`;

const LogButtons = styled.div`
	${(props) => css`
		display: flex;

		margin-top: 2rem;

		> * {
			margin-right: 1rem;
		}
	`}
`;

const InfoDisplay = styled.div`
	${(props) => css`
		font-size: 1.6rem;
		display: flex;
		justify-content: space-between;

		span {
			color: ${props.theme.colorTextPrimary};
			margin-right: 1rem;
		}
	`}
`;

const InfractionSection = styled.div`
	${(props) => css`
		display: flex;
		flex-direction: column;

		> * {
			margin-bottom: 1rem;
		}

		> :last-child {
			margin-bottom: 0;
		}
	`}
`;

const PreviousNames = styled.div`
	${(props) => css`
		font-size: 1.6rem;
	`}
`;

const HeaderBox = styled.div`
	> button:first-of-type {
		margin-left: auto;
	}
`;

class Player extends Component {
	constructor(props) {
		super(props);

		this.state = {
			player: null,
			modals: {
				warn: {
					show: false,
					ctx: {},
				},
				mute: {
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
				edit: {
					show: false,
					ctx: {},
				},
				del: {
					show: false,
					ctx: {},
				},
			},
		};

		this.warnModalRef = React.createRef();
		this.muteModalRef = React.createRef();
		this.kickModalRef = React.createRef();
		this.banModalRef = React.createRef();

		this.infractionRef = React.createRef();
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		const id = parseInt(nextProps.match.params.id);
		if (!id) {
			prevState.error = 'Player not found';
		}

		const highlightString = queryString.parse(nextProps.location.search);
		if (highlightString['?highlight']) {
			const highlightId = parseInt(highlightString['?highlight']);
			if (highlightId) {
				prevState.highlightId = highlightId;
			}
		}

		if (!prevState.player || prevState.player.id !== id) {
			nextProps.getPlayerSummary(id);
		}

		if (nextProps.player && nextProps.player.id === id) {
			prevState.player = nextProps.player;
		}

		return prevState;
	}

	componentDidMount() {
		const { highlightId } = this.state;

		// This timeout is a bit hacky, but we need it since the ref gets passed down to a child
		// component and this function is called before the ref is typically assigned.
		// The delay also seems to make the UX a bit better rather than instantly scrolling.
		setTimeout(() => {
			if (
				highlightId &&
				this.infractionRef &&
				this.infractionRef.current
			) {
				this.infractionRef.current.scrollIntoView({
					behavior: 'smooth',
					block: 'center',
				});
			}
		}, 500);
	}

	showModal = (modal, ctx) => () => {
		this.setState(
			(prevState) => ({
				...prevState,
				modals: {
					...prevState.modals,
					[modal]: {
						...prevState.modals[modal],
						show: true,
						ctx,
					},
				},
			}),
			() => {
				const ref = this[`${modal}ModalRef`];
				if (ref && ref.current) {
					ref.current.focus();
				}
			}
		);
	};

	closeModal = (modal) => () => {
		this.setState((prevState) => ({
			...prevState,
			modals: {
				...prevState.modals,
				[modal]: {
					...prevState.modals[modal],
					show: false,
					ctx: {},
				},
			},
		}));
	};

	toggleWatchClicked = () => {
		const { player } = this.state;

		if (!!player.watched) {
			this.props.unwatchPlayer(player.id);
		} else {
			this.props.watchPlayer(player.id);
		}
	};

	render() {
		const { player, error, modals, highlightId } = this.state;
		const { warn, mute, kick, ban, edit, del } = modals;
		const { self } = this.props;

		if (error) {
			return (
				<div>
					<Heading headingStyle={'title'}>{error}</Heading>
				</div>
			);
		}

		if (!player) {
			return (
				<div>
					<Heading headingStyle={'title'}>Loading...</Heading>
				</div>
			);
		}

		// Convert flat player infraction structures to arrays
		const warnings = Object.values(player.warnings).filter(
			(el) => el !== undefined
		);
		const mutes = Object.values(player.mutes).filter(
			(el) => el !== undefined
		);
		const kicks = Object.values(player.kicks).filter(
			(el) => el !== undefined
		);
		const bans = Object.values(player.bans).filter(
			(el) => el !== undefined
		);

		const infractionCount =
			warnings.length + mutes.length + kicks.length + bans.length;

		let previousNames = [];
		if (Array.isArray(player.previousNames)) {
			previousNames = player.previousNames.filter(
				(prevName) => prevName !== player.currentName
			);
		}

		return (
			<>
				<WarnModal
					player={warn.ctx}
					show={warn.show}
					onClose={this.closeModal('warn')}
					inputRef={this.warnModalRef}
					reload={true}
				/>

				<MuteModal
					player={mute.ctx}
					show={mute.show}
					onClose={this.closeModal('mute')}
					inputRef={this.muteModalRef}
					reload={true}
				/>

				<KickModal
					player={kick.ctx}
					show={kick.show}
					onClose={this.closeModal('kick')}
					inputRef={this.kickModalRef}
					reload={true}
				/>

				<BanModal
					player={ban.ctx}
					show={ban.show}
					onClose={this.closeModal('ban')}
					inputRef={this.banModalRef}
					reload={true}
				/>

				<EditInfractionModal
					infraction={edit.ctx}
					show={edit.show}
					onClose={this.closeModal('edit')}
				/>

				<DeleteInfractionModal
					infraction={del.ctx}
					show={del.show}
					onClose={this.closeModal('del')}
				/>

				<HeaderBox>
					<Button
						size={'small'}
						color={!!player.watched ? 'danger' : 'alert'}
						onClick={this.toggleWatchClicked}
					>
						{!!player.watched ? 'Unwatch' : 'Watch'}
					</Button>

					<Heading headingStyle={'title'}>
						Viewing: {player.currentName}
					</Heading>
					<LogButtons>
						<Button
							size={'normal'}
							color={'primary'}
							onClick={this.showModal('warn', player)}
						>
							Log Warning
						</Button>
						<Button
							size={'normal'}
							color={'primary'}
							onClick={this.showModal('mute', player)}
						>
							Log Mute
						</Button>
						<Button
							size={'normal'}
							color={'alert'}
							onClick={this.showModal('kick', player)}
						>
							Log Kick
						</Button>
						<Button
							size={'normal'}
							color={'danger'}
							onClick={this.showModal('ban', player)}
						>
							Log Ban
						</Button>
					</LogButtons>
				</HeaderBox>

				<div>
					<Heading headingStyle={'subtitle'}>Player Info</Heading>
					<PlayerInfo>
						{player.playFabId && (
							<InfoDisplay>
								<span>PlayFabID:</span>
								<p>{player.playFabId}</p>
							</InfoDisplay>
						)}

						{player.mcuuid && (
							<InfoDisplay>
								<span>MC-UUID:</span>
								<p>{player.mcuuid}</p>
							</InfoDisplay>
						)}

						<InfoDisplay>
							<span>Infractions:</span>
							<p>{infractionCount}</p>
						</InfoDisplay>

						<InfoDisplay>
							<span>Last seen:</span>
							<p>{timestampToDateTime(player.lastSeen)}</p>
						</InfoDisplay>
					</PlayerInfo>
				</div>

				<InfractionSection>
					{warnings.length > 0 ? (
						<>
							<Heading headingStyle={'subtitle'}>
								Warnings
							</Heading>
							{warnings.map(
								(warning) =>
									warning !== undefined && (
										<Infraction
											highlight={
												warning.id === highlightId
											}
											highlightRef={
												warning.id === highlightId
													? this.infractionRef
													: null
											}
											date={timestampToDateTime(
												warning.timestamp
											)}
											perms={self.permissions}
											issuer={warning.staffName}
											isOwnInfraction={
												self.id === warning.userId
											}
											reason={warning.reason}
											onEditClick={this.showModal(
												'edit',
												warning
											)}
											onDeleteClick={this.showModal(
												'del',
												warning
											)}
										/>
									)
							)}
						</>
					) : (
						<Heading headingStyle="subtitle">
							No warnings on record
						</Heading>
					)}
				</InfractionSection>

				<InfractionSection>
					{mutes.length > 0 ? (
						<>
							<Heading headingStyle={'subtitle'}>Mutes</Heading>
							{mutes.map(
								(mute) =>
									mute !== undefined && (
										<Infraction
											highlight={mute.id === highlightId}
											highlightRef={
												mute.id === highlightId
													? this.infractionRef
													: null
											}
											date={timestampToDateTime(
												mute.timestamp
											)}
											perms={self.permissions}
											isOwnInfraction={
												self.id === mute.userId
											}
											issuer={mute.staffName}
											reason={mute.reason}
											duration={mute.duration}
											remaining={getTimeRemaining(
												mute.timestamp,
												mute.duration
											)}
											onEditClick={this.showModal(
												'edit',
												mute
											)}
											onDeleteClick={this.showModal(
												'del',
												mute
											)}
										/>
									)
							)}
						</>
					) : (
						<Heading headingStyle="subtitle">
							No mutes on record
						</Heading>
					)}
				</InfractionSection>

				<InfractionSection>
					{kicks.length > 0 ? (
						<>
							<Heading headingStyle={'subtitle'}>Kicks</Heading>
							{kicks.map(
								(kick) =>
									kick !== undefined && (
										<Infraction
											highlight={kick.id === highlightId}
											highlightRef={
												kick.id === highlightId
													? this.infractionRef
													: null
											}
											date={timestampToDateTime(
												kick.timestamp
											)}
											perms={self.permissions}
											isOwnInfraction={
												self.id === kick.userId
											}
											issuer={kick.staffName}
											reason={kick.reason}
											onEditClick={this.showModal(
												'edit',
												kick
											)}
											onDeleteClick={this.showModal(
												'del',
												kick
											)}
										/>
									)
							)}
						</>
					) : (
						<Heading headingStyle="subtitle">
							No kicks on record
						</Heading>
					)}
				</InfractionSection>

				<InfractionSection>
					{bans.length > 0 ? (
						<>
							<Heading headingStyle={'subtitle'}>Bans</Heading>
							{bans.map(
								(ban) =>
									ban !== undefined && (
										<Infraction
											highlight={ban.id === highlightId}
											highlightRef={
												ban.id === highlightId
													? this.infractionRef
													: null
											}
											date={timestampToDateTime(
												ban.timestamp
											)}
											perms={self.permissions}
											issuer={ban.staffName}
											isOwnInfraction={
												self.id === ban.userId
											}
											reason={ban.reason}
											duration={ban.duration}
											remaining={getTimeRemaining(
												ban.timestamp,
												ban.duration
											)}
											onEditClick={this.showModal(
												'edit',
												ban
											)}
											onDeleteClick={this.showModal(
												'del',
												ban
											)}
										/>
									)
							)}
						</>
					) : (
						<Heading headingStyle="subtitle">
							No bans on record
						</Heading>
					)}
				</InfractionSection>

				{player.previousNames ? (
					<div>
						<Heading headingStyle="subtitle">
							Previous Names
						</Heading>
						<PreviousNames>
							{previousNames.join(', ')}
						</PreviousNames>
					</div>
				) : null}
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	self: state.user.self,
	player: state.players.currentPlayer,
	loading: state.loading.playersummary,
});

const mapDispatchToProps = (dispatch) => ({
	getPlayerSummary: (playerId) => dispatch(getPlayerSummary(playerId)),
	setLoading: (isLoading) => dispatch(setLoading('main', isLoading)),
	unwatchPlayer: (playerId) => dispatch(unwatchPlayer(playerId)),
	watchPlayer: (playerId) => dispatch(watchPlayer(playerId)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Player);
