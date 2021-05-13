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
import Heading from '../../components/Heading';
import { connect } from 'react-redux';
import styled, { css } from 'styled-components';
import Alert from '../../components/Alert';
import ServerSelector from '../../components/ServerSelector';
import respondTo from '../../mixins/respondTo';
import PlayerSelector from '../../components/PlayerSelector';
import TextInput from '../../components/TextInput';
import Button from '../../components/Button';
import DateTimeSelector from '../../components/DateTimeSelector';
import { timestampToDateTime } from '../../utils/timeUtils';
import {
	searchChatRecords,
	setChatSearchResults,
} from '../../redux/chat/chatActions';
import { Link } from 'react-router-dom';
import {
	DisabledPageSwitcherButton,
	PageSwitcher,
	PageSwitcherButton,
	PageSwitcherLabel,
} from './Players';
import ReactTooltip from 'react-tooltip';
import { setSearchResults } from '../../redux/infractions/infractionActions';
import { setSuccess } from '../../redux/success/successActions';
import { setErrors } from '../../redux/error/errorActions';

const ChatRecordsBox = styled.div``;

const FilterBox = styled.div`
	display: grid;
	grid-template-rows: 1fr 1fr 1fr 1fr 1fr 1fr;
	grid-template-columns: 1fr;

	${respondTo.medium`
		grid-template-rows: 1fr 1fr;
		grid-template-columns: 1fr 1fr 1fr 1fr;
	  	grid-column-gap: 1rem;
	  
	    > :nth-child(1) {
		  grid-column: span 2;
		}

	  	> :nth-child(4) {
		  grid-column: span 2;
		}
	  
	  	button {
		  height: 4rem;
		}
  	`}
`;

const ResultsBox = styled.div``;

const Results = styled.div`
	${(props) => css`
		display: flex;
		flex-direction: column;
		background-color: ${props.theme.colorAccent};
		margin-top: 1rem;
		border-radius: ${props.theme.borderRadiusNormal};
		//font-family: 'Roboto Mono', 'Poppins', sans-serif;

		> :nth-child(even) {
			background-color: ${props.theme.colorBackground};
		}
	`}
`;

export const MobileLabel = styled.div`
	${(props) => css`
		display: inline;
		color: ${props.theme.colorPrimary};

		${respondTo.medium`
			display: none;
		`};
	`}
`;

const Result = styled.div`
	${(props) => css`
		font-size: 1.2rem;
		display: flex;
		flex-direction: column;

		${respondTo.medium`
		  flex-direction: row;
		`}

		> * {
			padding: 0.5rem;
		}

		a {
			text-decoration: none !important;
			color: ${props.theme.colorTextSecondary};
		}

		> :nth-child(1) {
			width: 7rem;
		}

		> :nth-child(2) {
			width: 20rem;

			:hover {
				cursor: pointer;
				background-color: ${props.theme.colorPrimary};
			}
		}

		> :nth-child(3) {
			width: 20rem;

			:hover {
				cursor: pointer;
				background-color: ${props.theme.colorPrimary};
			}
		}

		> :nth-child(4) {
			width: 100%;
			overflow-wrap: break-word;
			word-wrap: break-word;
			word-break: break-word;
			flex: 1;
		}
	`}
`;

const limitInterval = 20;

class ChatRecords extends Component {
	constructor(props) {
		super(props);

		this.state = {
			page: 0,
			currentFilters: {},
			filters: {
				startDate: new Date(new Date().getTime() - 1000 * 60 * 60), // 1 hour ago
				endDate: new Date(), // now
				message: '',
				player: null,
				serverId: '',
			},
			searchWasRun: false,
			errors: {},
		};
	}

	componentDidMount() {
		this.props.clearResults();
	}

	static getDerivedStateFromProps(nextProps, prevState) {
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

	onDateChange = (key) => (date) => {
		if (!date.toDate) {
			return;
		}

		this.setState((prevState) => ({
			...prevState,
			filters: {
				...prevState.filters,
				[key]: date.toDate(),
			},
		}));
	};

	onChange = (e) => {
		this.setState((prevState) => ({
			...prevState,
			filters: {
				...prevState.filters,
				[e.target.name]: e.target.value,
			},
		}));
	};

	onPlayerChange = (player) => {
		this.setState((prevState) => ({
			...prevState,
			filters: {
				...prevState.filters,
				player: player,
			},
		}));
	};

	onNextPage = () => {
		const { page, currentFilters } = this.state;

		const nextPage = page + 1;

		const searchData = {
			...currentFilters,
			limit: limitInterval,
			offset: nextPage * limitInterval,
		};

		// Update page in state
		this.setState((prevState) => ({
			...prevState,
			page: nextPage,
		}));

		this.props.searchChatRecords(searchData);
	};

	onPrevPage = () => {
		const { page, currentFilters } = this.state;

		const prevPage = page - 1;

		const searchData = {
			...currentFilters,
			limit: limitInterval,
			offset: prevPage * limitInterval,
		};

		// Update page in state
		this.setState((prevState) => ({
			...prevState,
			page: prevPage,
		}));

		this.props.searchChatRecords(searchData);
	};

	onResultsKeyDown = (e) => {
		const { page, searchWasRun } = this.state;
		const { results: searchResults } = this.props;

		if (!searchWasRun || !searchResults) {
			return;
		}

		const { results, count } = searchResults;

		if (results.length <= 0) {
			return;
		}

		const amountOfPages = Math.ceil(count / limitInterval);

		switch (e.keyCode) {
			case 39:
				if (page !== amountOfPages - 1) {
					this.onNextPage();
				}
				break;
			case 37:
				if (page > 0) {
					this.onPrevPage();
				}
				break;
			default:
				return;
		}
	};

	onDateClick = (timestamp) => () => {
		// When a date is clicked we want to show the messages within 2 minutes
		// of either side of the date clicked.
		const nanoTimestamp = Math.floor(timestamp * 1000);

		const newStartTime = nanoTimestamp - 1000 * 60 * 2;
		const newEndTime = nanoTimestamp + 1000 * 60 * 2;

		this.setState(
			(prevState) => ({
				...prevState,
				filters: {
					message: '',
					player: null,
					startDate: new Date(newStartTime),
					endDate: new Date(newEndTime),
				},
			}),
			() => {
				this.onViewRecordsClicked();
			}
		);
	};

	onViewRecordsClicked = () => {
		const {
			message,
			startDate,
			endDate,
			serverId,
			player,
		} = this.state.filters;

		const searchData = {};

		if (message) {
			searchData.message = message.toString();
		}

		if (startDate) {
			searchData.startDate = Math.floor(startDate.getTime() / 1000);
		}

		if (endDate) {
			searchData.endDate = Math.floor(endDate.getTime() / 1000);
		}

		if (serverId) {
			searchData.serverId = serverId.toString();
		}

		if (player) {
			searchData.playerId = player.id.toString();
		}

		// Set limit and offset
		searchData.limit = limitInterval;
		searchData.offset = 0;

		// Record current search data
		this.setState(
			(prevState) => ({
				...prevState,
				currentFilters: searchData,
				searchWasRun: true,
			}),
			() => {
				this.props.searchChatRecords(searchData);
			}
		);
	};

	render() {
		const { errors, filters, searchWasRun, page } = this.state;
		const { message, startDate, endDate, serverId, player } = filters;
		const { results } = this.props;
		const { results: searchResults, count } = results;

		const amountOfPages = Math.ceil(count / limitInterval);

		return (
			<>
				<ReactTooltip delayShow={500} multiline={true} />

				<div>
					<Heading headingStyle={'title'}>Chat Records</Heading>
				</div>

				<ChatRecordsBox>
					<Alert type="error" message={errors.general} />

					<FilterBox>
						<TextInput
							name={'message'}
							title={'message'}
							size={'small'}
							placeholder={'Message'}
							onChange={this.onChange}
							value={message}
						/>
						<DateTimeSelector
							title={'start date'}
							onChange={this.onDateChange('startDate')}
							value={startDate}
						/>
						<DateTimeSelector
							title={'end date'}
							onChange={this.onDateChange('endDate')}
							value={endDate}
						/>
						<ServerSelector
							onChange={this.onChange}
							name={'serverId'}
							value={serverId}
						/>
						<PlayerSelector
							title={'player'}
							name={'playerId'}
							value={player ? player.currentName : 'Select...'}
							onSelect={this.onPlayerChange}
						/>
						<Button onClick={this.onViewRecordsClicked}>
							View Records
						</Button>
					</FilterBox>

					{!searchResults && searchWasRun && (
						<div>
							<Heading headingStyle={'subtitle'}>
								No results found
							</Heading>
						</div>
					)}

					{searchResults && (
						<ResultsBox
							onKeyDown={this.onResultsKeyDown}
							tabIndex={'0'}
						>
							<Heading headingStyle={'subtitle'}>Results</Heading>
							<p>found {count} results</p>

							<Results>
								{searchResults.map((result) => (
									<Result>
										<div>
											<MobileLabel>ID: </MobileLabel>
											{result.id}
										</div>
										<div
											data-tip={`Click on a timestamp to display all messages
												within two minutes of the clicked time.`}
											onClick={this.onDateClick(
												result.timestamp
											)}
										>
											<MobileLabel>Date: </MobileLabel>
											{timestampToDateTime(
												result.timestamp
											)}
										</div>
										<Link to={`/player/${result.playerId}`}>
											<MobileLabel>Player: </MobileLabel>
											{result.playerName}
										</Link>
										<div>{result.message}</div>
									</Result>
								))}
							</Results>
						</ResultsBox>
					)}

					{searchWasRun &&
					searchResults &&
					searchResults.length > 0 ? (
						<PageSwitcher>
							<div>
								{page > 0 ? (
									<PageSwitcherButton
										onClick={this.onPrevPage}
									>
										Prev
									</PageSwitcherButton>
								) : (
									<DisabledPageSwitcherButton>
										Prev
									</DisabledPageSwitcherButton>
								)}
								<PageSwitcherLabel>
									{page + 1}
								</PageSwitcherLabel>
								{searchResults &&
								searchResults.length > 0 &&
								page !== amountOfPages - 1 ? (
									<PageSwitcherButton
										onClick={this.onNextPage}
									>
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
				</ChatRecordsBox>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	success: state.success.chatrecords,
	error: state.error.chatrecords,
	results: state.chat.searchResults,
});

const mapDispatchToProps = (dispatch) => ({
	searchChatRecords: (searchData) => dispatch(searchChatRecords(searchData)),
	clearResults: () => dispatch(setChatSearchResults([])),
	clearSuccess: () => dispatch(setSuccess('chatrecords', undefined)),
	clearErrors: () => dispatch(setErrors('chatrecords', undefined)),
});

export default connect(mapStateToProps, mapDispatchToProps)(ChatRecords);
