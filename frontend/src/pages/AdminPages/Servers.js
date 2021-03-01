import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import respondTo from '../../mixins/respondTo';
import Heading from '../../components/Heading';
import { Link } from 'react-router-dom';
import Button from '../../components/Button';
import { connect } from 'react-redux';
import { editServer } from '../../redux/servers/serverActions';

const ServerTable = styled.table`
	${(props) => css`
		width: 100%;
		text-align: left;
		font-size: 1.6rem;
		font-weight: 400;
		margin-bottom: 2rem;

		display: grid;
		grid-template-rows: auto;
		grid-template-columns: 1fr 1fr 1fr 0.4fr;
		grid-row-gap: 1rem;

		${respondTo.medium`
		  flex-direction: column;
		`}
	`}
`;

const ServerTableHeading = styled.div`
	${(props) => css`
		color: ${props.theme.colorTextPrimary};
		grid-row: 1;
	`}
`;

const ServerButtonBox = styled.div`
	${(props) => css`
		display: flex;

		> * {
			margin-right: 0.5rem;
		}
	`}
`;

class Servers extends Component {
	render() {
		const { servers: serversObj } = this.props;

		if (!serversObj) {
			return (
				<div>
					<Heading headingStyle={'title'}>No servers found</Heading>
				</div>
			);
		}

		const servers = Object.values(serversObj);

		return (
			<>
				<div>
					<Heading headingStyle={'title'}>Servers</Heading>
				</div>

				<div>
					<ServerTable>
						<ServerTableHeading>Server Name</ServerTableHeading>
						<ServerTableHeading>Players</ServerTableHeading>
						<ServerTableHeading>Status</ServerTableHeading>
						<ServerTableHeading />

						{servers.map((server) => (
							<>
								<div>{server.name}</div>
								<div>
									{server.players ? server.players.length : 0}
								</div>
								<div>
									{server.online ? 'Online' : 'Offline'}
								</div>
								<ServerButtonBox>
									<Button size={'small'} color={'primary'}>
										Edit
									</Button>
									<Button size={'small'} color={'danger'}>
										Delete
									</Button>
								</ServerButtonBox>
							</>
						))}
					</ServerTable>
				</div>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	servers: state.servers,
});

const mapDispatchToProps = (dispatch) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(Servers);
