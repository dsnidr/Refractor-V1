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
import GameSelector from '../../components/GameSelector';
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
			filters: {
				startdate: new Date(new Date().getTime() - 1000 * 60 * 60), // 1 hour ago
				enddate: new Date(), // now
			},
			errors: {},
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {}

	onDateChange = (key) => (date) => {
		this.setState((prevState) => ({
			...prevState,
			filters: {
				...prevState.filters,
				[key]: date,
			},
		}));
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
						/>
						<DateTimeSelector
							title={'start date'}
							onChange={this.onDateChange('startdate')}
							value={filters.startdate}
						/>
						<DateTimeSelector
							title={'end date'}
							onChange={this.onDateChange('enddate')}
							value={filters.enddate}
						/>
						<ServerSelector />
						<PlayerSelector
							title={'player'}
							value={
								filters.player
									? filters.player.currentName
									: 'Select...'
							}
						/>
						<Button>View Records</Button>
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
