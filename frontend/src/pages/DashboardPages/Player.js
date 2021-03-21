import React, { Component } from 'react';
import Heading from '../../components/Heading';
import { connect } from 'react-redux';
import { getPlayerSummary } from '../../redux/players/playerActions';
import styled, { css } from 'styled-components';
import respondTo from '../../mixins/respondTo';
import { timestampToDateTime, getTimeRemaining } from '../../utils/timeUtils';
import Button from '../../components/Button';
import Infraction from '../../components/Infraction';

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

		> :first-child {
			margin-bottom: 1rem;
		}
	`}
`;

class Player extends Component {
	constructor(props) {
		super(props);

		this.state = {
			player: null,
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		const id = parseInt(nextProps.match.params.id);
		if (!id) {
			return {
				...prevState,
				error: 'Player not found',
			};
		}

		if (!prevState.player) {
			nextProps.getPlayerSummary(id);

			if (nextProps.player) {
				prevState.player = nextProps.player;
			}
		}

		return prevState;
	}

	render() {
		const { player, error } = this.state;

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
				<div>
					<Heading headingStyle={'title'}>
						Viewing: {player.currentName}
					</Heading>
					<LogButtons>
						<Button size={'normal'} color={'primary'}>
							Log Warning
						</Button>
						<Button size={'normal'} color={'primary'}>
							Log Mute
						</Button>
						<Button size={'normal'} color={'alert'}>
							Log Kick
						</Button>
						<Button size={'normal'} color={'danger'}>
							Log Ban
						</Button>
					</LogButtons>
				</div>

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
											date={timestampToDateTime(
												warning.timestamp
											)}
											issuer={warning.staffName}
											reason={warning.reason}
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
											date={timestampToDateTime(
												mute.timestamp
											)}
											issuer={mute.staffName}
											reason={mute.reason}
											duration={mute.duration}
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
											date={timestampToDateTime(
												kick.timestamp
											)}
											issuer={kick.staffName}
											reason={kick.reason}
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
											date={timestampToDateTime(
												ban.timestamp
											)}
											issuer={ban.staffName}
											reason={ban.reason}
											duration={ban.duration}
											remaining={getTimeRemaining(
												ban.timestamp,
												ban.duration
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
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	player: state.players.currentPlayer,
});

const mapDispatchToProps = (dispatch) => ({
	getPlayerSummary: (playerId) => dispatch(getPlayerSummary(playerId)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Player);
