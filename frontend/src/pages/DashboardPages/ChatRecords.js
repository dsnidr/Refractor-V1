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

		> :nth-child(1) {
			min-width: 3rem;
			max-width: 10rem;
			width: auto;

			:hover {
				cursor: pointer;
				background-color: ${props.theme.colorPrimary};
			}
		}

		> :nth-child(2) {
			width: 20rem;

			:hover {
				cursor: pointer;
				background-color: ${props.theme.colorPrimary};
			}
		}

		> :nth-child(3) {
			flex: 1;
		}
	`}
`;

class ChatRecords extends Component {
	constructor(props) {
		super(props);

		this.state = {
			page: 0,
			currentFilters: {},
			filters: {
				startDate: new Date(new Date().getTime() - 1000 * 60 * 60), // 1 hour ago
				endDate: new Date(), // now
			},
			searchWasRun: false,
			errors: {},
		};
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
		this.setState((prevState) => ({
			...prevState,
			filters: {
				...prevState.filters,
				[key]: date,
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

	onViewRecordsClicked = () => {
		const {
			message,
			startDate,
			endDate,
			serverId,
			player,
		} = this.state.filters;

		console.log(
			message,
			startDate,
			endDate,
			serverId,
			player ? player.id : undefined
		);
	};

	render() {
		const { errors, filters } = this.state;

		return (
			<>
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
						/>
						<DateTimeSelector
							title={'start date'}
							onChange={this.onDateChange('startDate')}
							value={filters.startDate}
						/>
						<DateTimeSelector
							title={'end date'}
							onChange={this.onDateChange('endDate')}
							value={filters.endDate}
						/>
						<ServerSelector
							onChange={this.onChange}
							name={'serverId'}
						/>
						<PlayerSelector
							title={'player'}
							name={'playerId'}
							value={
								filters.player
									? filters.player.currentName
									: 'Select...'
							}
							onSelect={this.onPlayerChange}
						/>
						<Button onClick={this.onViewRecordsClicked}>
							View Records
						</Button>
					</FilterBox>

					<ResultsBox>
						<Heading headingStyle={'subtitle'}>Results</Heading>

						<Results>
							<Result>
								<div>
									<MobileLabel>ID: </MobileLabel>1
								</div>
								<div>
									<MobileLabel>Date: </MobileLabel>
									{timestampToDateTime(18000)}
								</div>
								<div>Lorem ipsum dolor si amet</div>
							</Result>
							<Result>
								<div>
									<MobileLabel>ID: </MobileLabel>2
								</div>
								<div>
									<MobileLabel>Date: </MobileLabel>
									{timestampToDateTime(19000)}
								</div>
								<div>Lorem ipsum dolor si amet</div>
							</Result>
							<Result>
								<div>
									<MobileLabel>ID: </MobileLabel>3
								</div>
								<div>
									<MobileLabel>Date: </MobileLabel>
									{timestampToDateTime(20000)}
								</div>
								<div>Lorem ipsum dolor si amet</div>
							</Result>
							<Result>
								<div>
									<MobileLabel>ID: </MobileLabel>4
								</div>
								<div>
									<MobileLabel>Date: </MobileLabel>
									{timestampToDateTime(21000)}
								</div>
								<div>Lorem ipsum dolor si amet</div>
							</Result>
						</Results>
					</ResultsBox>
				</ChatRecordsBox>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	success: state.success.chatrecords,
	error: state.error.chatrecords,
	servers: state.servers,
	games: state.games,
});

const mapDispatchToProps = (dispatch) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(ChatRecords);
