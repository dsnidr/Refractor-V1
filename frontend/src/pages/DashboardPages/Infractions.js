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
import styled, { css } from 'styled-components';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';
import InfractionPreview from '../../components/InfractionPreview';
import Select from '../../components/Select';
import ServerSelector from '../../components/ServerSelector';
import Button from '../../components/Button';
import PlayerSelector from '../../components/PlayerSelector';
import { getAllUsers } from '../../redux/user/userActions';
import {
	flags,
	hasFullAccess,
	hasPermission,
} from '../../permissions/permissions';
import {
	getRecentInfractions,
	searchInfractions,
	setSearchResults,
} from '../../redux/infractions/infractionActions';
import Alert from '../../components/Alert';
import { setSuccess } from '../../redux/success/successActions';
import { setErrors } from '../../redux/error/errorActions';
import { timestampToDateTime } from '../../utils/timeUtils';
import {
	DisabledPageSwitcherButton,
	PageSwitcher,
	PageSwitcherButton,
	PageSwitcherLabel,
} from './Players';

const RecentInfractionsBox = styled.div`
	> :first-child {
		margin-bottom: 1rem;
	}
`;

const RecentInfractions = styled.div`
	${(props) => css`
		display: flex;
		flex-direction: column;
		height: clamp(10rem, 20rem, 30vh);
		overflow-y: scroll;

		> :nth-child(even) {
			background-color: ${props.theme.colorBackground};
		}
	`}
`;

const InfractionSearchBox = styled.div`
	> :nth-child(1) {
		margin-bottom: 0.5rem;
	}

	> :nth-child(2) {
		margin-bottom: 3rem;
		font-size: 1.4rem;
	}
`;

const SearchBox = styled.div`
	display: grid;
	grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
	grid-column-gap: 1rem;

	button {
		height: 4rem;
	}
`;

const ResultsBox = styled.div`
	${(props) => css`
		> :nth-child(2) {
			margin-bottom: 1rem;
			font-size: 1.2rem;
		}

		> :nth-child(odd) {
			background-color: ${props.theme.colorBackground};
		}
	`}
`;

const limitInterval = 10;

class Infractions extends Component {
	constructor(props) {
		super(props);

		this.state = {
			page: 0,
			currentSearchData: null,
			errors: {},
			searchWasRun: false,
			fields: {},
			currentSearchFields: {},
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		/* global BigInt */
		if (
			nextProps.self &&
			!nextProps.otherUsers &&
			hasFullAccess(BigInt(nextProps.self.permissions))
		) {
			nextProps.getAllUsers();
		}

		if (nextProps.success) {
			prevState.searchWasRun = true;
			prevState.errors = {};
		}

		if (nextProps.errors) {
			prevState.errors = {
				...prevState.errors,
				...nextProps.errors,
			};
		}
	}

	componentDidMount() {
		this.props.clearResults();
		this.props.getRecentInfractions();
	}

	onPlayerSelectionChanged = (player) => {
		this.setState((prevState) => ({
			...prevState,
			fields: {
				...prevState.fields,
				player: player,
			},
		}));
	};

	onSelectChange = (e) => {
		this.setState((prevState) => ({
			...prevState,
			fields: {
				...prevState.fields,
				[e.target.name]: e.target.value,
			},
		}));
	};

	onSearchClick = (e) => {
		e.preventDefault();

		// Clear previous results, successes and errors
		this.props.clearResults();
		this.props.clearSuccess();
		this.props.clearErrors();

		const { type, player, user, game, server } = this.state.fields;
		const searchData = {};
		const errors = {};

		// A bit messy, but I'm running out of time and this works.
		// Could refactor later to be less if spam.
		if (type) {
			searchData.type = type.toString();
		}

		if (player) {
			searchData.playerId = player.id.toString();
		}

		if (user) {
			searchData.userId = user.toString();
		}

		if (game) {
			searchData.game = game.toString();
		}

		if (server) {
			searchData.serverId = server.id.toString();
		}

		if (Object.keys(searchData).length === 0) {
			errors.general = 'You must set at least one search filter';
		}

		this.setState((prevState) => ({
			...prevState,
			errors: errors,
		}));

		if (Object.keys(errors).length > 0) {
			return;
		}

		// Set limit and offset
		searchData.limit = limitInterval;
		searchData.offset = 0;

		// Record current search data
		this.setState((prevState) => ({
			...prevState,
			currentSearchFields: searchData,
		}));

		this.props.searchInfractions(searchData);
	};

	onNextPage = () => {
		const { page, currentSearchFields } = this.state;

		const nextPage = page + 1;

		const searchData = {
			...currentSearchFields,
			limit: limitInterval,
			offset: nextPage * limitInterval,
		};

		// Update page in state
		this.setState((prevState) => ({
			...prevState,
			page: nextPage,
		}));

		this.props.searchInfractions(searchData);
	};

	onPrevPage = () => {
		const { page, currentSearchFields } = this.state;

		const prevPage = page - 1;

		const searchData = {
			...currentSearchFields,
			limit: limitInterval,
			offset: prevPage * limitInterval,
		};

		// Update page in state
		this.setState((prevState) => ({
			...prevState,
			page: prevPage,
		}));

		this.props.searchInfractions(searchData);
	};

	render() {
		const {
			results: searchResults,
			self,
			otherUsers,
			games: allGames,
			recentInfractions,
		} = this.props;
		const { results, count } = searchResults;
		const { searchWasRun, fields, errors, page } = this.state;
		const { player } = fields;

		const users = [];
		if (!!otherUsers) {
			users.push(...Object.values(otherUsers));
		} else {
			users.push(self);
		}

		const games = Object.keys(allGames);

		const amountOfPages = Math.ceil(count / limitInterval);

		return (
			<>
				<div>
					<Heading headingStyle={'title'}>Infractions</Heading>
				</div>

				{recentInfractions && (
					<RecentInfractionsBox>
						<Heading headingStyle={'subtitle'}>
							Recent Infractions
						</Heading>

						<RecentInfractions>
							{recentInfractions.map((infraction, index) => (
								<InfractionPreview
									key={`ri${index}`}
									to={`/player/${infraction.playerId}?highlight=${infraction.id}`}
									type={infraction.type}
									player={infraction.playerName}
									date={timestampToDateTime(
										infraction.timestamp
									)}
									issuer={infraction.staffName}
									duration={infraction.duration}
									reason={infraction.reason}
								/>
							))}
						</RecentInfractions>
					</RecentInfractionsBox>
				)}

				<InfractionSearchBox>
					<Heading headingStyle={'subtitle'}>
						Search Infractions
					</Heading>

					<p>
						Apply the filters you want and click Search. To leave a
						filter out, set it to "Select...".
					</p>

					<Alert type="error" message={errors.general} />

					<SearchBox>
						<Select
							title={'infraction type'}
							name={'type'}
							onChange={this.onSelectChange}
						>
							<option value={''}>Select...</option>
							<option value={'WARNING'}>Warning</option>
							<option value={'MUTE'}>Mute</option>
							<option value={'KICK'}>Kick</option>
							<option value={'BAN'}>Ban</option>
						</Select>
						<PlayerSelector
							title={'player'}
							onSelect={this.onPlayerSelectionChanged}
							value={player ? player.currentName : 'Select...'}
						/>
						<Select
							title={'user'}
							onChange={this.onSelectChange}
							name={'user'}
						>
							<option value={''}>Select...</option>
							{users.map((user, index) => (
								<option key={`user${index}`} value={user.id}>
									{user.username}
								</option>
							))}
						</Select>
						<Select
							title={'game'}
							onChange={this.onSelectChange}
							name={'game'}
						>
							<option value={null}>Select...</option>
							{games.map((game, index) => (
								<option key={`game${index}`} value={game}>
									{game}
								</option>
							))}
						</Select>
						<ServerSelector
							default={'Select...'}
							onChange={this.onSelectChange}
							name={'server'}
						/>
					</SearchBox>

					<Button
						size={'normal'}
						color={'primary'}
						onClick={this.onSearchClick}
					>
						Search
					</Button>
				</InfractionSearchBox>

				{results && results.length > 0 ? (
					<ResultsBox>
						<Heading headingStyle={'subtitle'}>Results</Heading>
						<p>found {count} results</p>

						{results.map((result, index) => (
							<InfractionPreview
								to={`/player/${result.playerId}?highlight=${result.id}`}
								type={result.type}
								player={result.playerName}
								date={timestampToDateTime(result.timestamp)}
								issuer={result.staffName}
								duration={result.duration}
								reason={result.reason}
							/>
						))}
					</ResultsBox>
				) : (
					searchWasRun && (
						<Heading headingStyle={'subtitle'}>
							No results found
						</Heading>
					)
				)}

				{this.state.searchWasRun && results && results.length > 0 ? (
					<PageSwitcher>
						<div>
							{page > 0 ? (
								<PageSwitcherButton onClick={this.onPrevPage}>
									Prev
								</PageSwitcherButton>
							) : (
								<DisabledPageSwitcherButton>
									Prev
								</DisabledPageSwitcherButton>
							)}
							<PageSwitcherLabel>{page + 1}</PageSwitcherLabel>
							{results &&
							results.length > 0 &&
							page !== amountOfPages - 1 ? (
								<PageSwitcherButton onClick={this.onNextPage}>
									Next
								</PageSwitcherButton>
							) : (
								<DisabledPageSwitcherButton>
									Next
								</DisabledPageSwitcherButton>
							)}
						</div>
					</PageSwitcher>
				) : null}
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	self: state.user.self,
	otherUsers: state.user.others,
	games: state.games,
	success: state.success.searchinfractions,
	errors: state.error.searchinfractions,
	results: state.infractions.searchResults,
	recentInfractions: state.infractions.recentInfractions,
});

const mapDispatchToProps = (dispatch) => ({
	getAllUsers: () => dispatch(getAllUsers()),
	searchInfractions: (searchData) => dispatch(searchInfractions(searchData)),
	clearResults: () => dispatch(setSearchResults([])),
	clearSuccess: () => dispatch(setSuccess('searchinfractions', undefined)),
	clearErrors: () => dispatch(setErrors('searchinfractions', undefined)),
	getRecentInfractions: () => dispatch(getRecentInfractions()),
});

export default connect(mapStateToProps, mapDispatchToProps)(Infractions);
