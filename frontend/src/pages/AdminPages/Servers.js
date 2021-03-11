import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import respondTo from '../../mixins/respondTo';
import Heading from '../../components/Heading';
import Button from '../../components/Button';
import { connect } from 'react-redux';
import { push } from 'connected-react-router';
import { Link } from 'react-router-dom';
import { useHistory } from 'react-router';
import StatusTag from '../../components/StatusTag';
import Modal from '../../components/Modal';
import BasicModal from '../../components/modals/BasicModal';
import { deleteServer } from '../../redux/servers/serverActions';

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

		a {
			text-decoration: none !important;
			margin-right: 0.5rem;
		}
	`}
`;

class Servers extends Component {
	constructor(props) {
		super(props);

		this.state = {
			redirectTo: null,
			modals: {
				deleteServer: {
					show: false,
					ctx: {},
				},
			},
		};
	}

	onAddServerClick = () => {
		this.props.history.push('/servers/add');
	};

	showModal = (modalName, context) => () => {
		this.setState((prevState) => ({
			...prevState,
			modals: {
				...prevState.modals,
				[modalName]: {
					...prevState.modals[modalName],
					show: true,
					ctx: context,
				},
			},
		}));
	};

	hideModal = (modalName) => () => {
		this.setState((prevState) => ({
			...prevState,
			modals: {
				...prevState.modals,
				[modalName]: {
					...prevState.modals[modalName],
					show: false,
					ctx: {},
				},
			},
		}));
	};

	onDeleteServerSubmit = () => {
		const { ctx: server } = this.state.modals.deleteServer;

		this.props.deleteServer(server.id);

		this.hideModal('deleteServer')();
	};

	render() {
		const { deleteServer: deleteServerModal } = this.state.modals;
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
				<BasicModal
					show={deleteServerModal.show}
					heading={`Delete ${deleteServerModal.ctx.name}`}
					message={`Are you sure you wish to delete ${deleteServerModal.ctx.name}? This action cannot be undone.`}
					success={null}
					error={null}
					onClose={this.hideModal('deleteServer')}
					onSubmit={this.onDeleteServerSubmit}
				/>

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
									<StatusTag status={server.online} />
								</div>
								<ServerButtonBox>
									<Link to={`/servers/edit/${server.id}`}>
										<Button
											size={'small'}
											color={'primary'}
										>
											Edit
										</Button>
									</Link>
									<Link>
										<Button
											onClick={this.showModal(
												'deleteServer',
												server
											)}
											size={'small'}
											color={'danger'}
										>
											Delete
										</Button>
									</Link>
								</ServerButtonBox>
							</>
						))}
					</ServerTable>

					<Button
						size="normal"
						color="primary"
						onClick={this.onAddServerClick}
					>
						Add Server
					</Button>
				</div>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	servers: state.servers,
});

const mapDispatchToProps = (dispatch) => ({
	deleteServer: (serverId) => dispatch(deleteServer(serverId)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Servers);
