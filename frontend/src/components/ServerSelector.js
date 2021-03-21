import React from 'react';
import { connect } from 'react-redux';
import Select from './Select';

const ServerSelector = (props) => {
	if (!props.games) {
		return null;
	}

	const games = Object.values(props.games);

	return (
		<Select
			name={'selectedServer'}
			onChange={props.onChange}
			error={props.error}
		>
			{games.map((game) =>
				game.servers.map((server) => (
					<option key={server.id} value={server.id}>
						{game.name} - {server.name}
					</option>
				))
			)}
		</Select>
	);
};

const mapStateToProps = (state) => ({
	games: state.games,
});

export default connect(mapStateToProps)(ServerSelector);
