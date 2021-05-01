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

import React from 'react';
import { connect } from 'react-redux';
import Select from './Select';

const GameSelector = (props) => {
	if (!props.games) {
		return null;
	}

	const games = Object.values(props.games);

	return (
		<Select
			name={props.name || 'selectedGame'}
			onChange={props.onChange}
			error={props.error}
			title={'game'}
		>
			<option value={''}>{props.default || 'Select game'}</option>
			{games.map((game) => (
				<option key={game.name} value={game.name}>
					{game.name}
				</option>
			))}
		</Select>
	);
};

const mapStateToProps = (state) => ({
	games: state.games,
});

export default connect(mapStateToProps)(GameSelector);
