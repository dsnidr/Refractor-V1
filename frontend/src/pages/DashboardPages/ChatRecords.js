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

const ChatRecordsBox = styled.div``;

const FilterBox = styled.div`
	display: grid;
	grid-template-rows: 1fr 1fr 1fr 1fr 1fr 1fr;
	grid-template-columns: 1fr;

	${respondTo.medium`
		grid-template-rows: 1fr 1fr;
		grid-template-columns: 1fr 1fr 1fr 1fr;
	  	grid-column-gap: 1rem;
	  
	    > :first-child {
		  grid-column: span 4;
		}
	  
	  	> :nth-child(4) {
		  grid-column: span 2;
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
						<GameSelector />
						<PlayerSelector
							title={'player'}
							value={
								filters.player
									? filters.player.currentName
									: 'Select...'
							}
						/>
					</FilterBox>
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
