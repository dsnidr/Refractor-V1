import React, { Component } from 'react';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';
import styled, { css } from 'styled-components';
import respondTo from '../../mixins/respondTo';
import { Link } from 'react-router-dom';

const ServerSummary = styled.div`
	${(props) => css`
		display: flex;
		flex-direction: column;
		font-size: 1.4rem;

		> * {
			margin-right: 2rem;
		}

		${respondTo.medium`
      flex-direction: row;
    `}
	`}
`;

const InfoSpan = styled.span`
	${(props) => css`
		color: ${props.theme.colorTextPrimary};
	`}
`;

const PlayerList = styled.div`
	${(props) => css`
		margin-top: 1rem;
		display: grid;
		grid-template-columns: auto;
		grid-auto-rows: auto;
		grid-row-gap: 1rem;

		${respondTo.medium`
      		grid-template-columns: 1fr 1fr 1fr;
      		grid-gap: 2rem;
		`}

		${respondTo.large`
            grid-template-columns: 1fr 1fr 1fr 1fr;
            grid-gap: 2rem;
        `}
		
    	${respondTo.extralarge`
            grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
            grid-gap: 2rem;
        `}
	`}
`;

const Player = styled.div`
	${(props) => css`
		background-color: ${props.theme.colorAccent};
		border-radius: ${props.theme.borderRadiusNormal};
		white-space: nowrap;
		text-overflow: ellipsis;
		overflow: hidden;
		grid-row: auto;

		a {
			color: ${props.theme.colorTextSecondary} !important;
			text-decoration: none !important;
		}

		h1 {
			padding: 1rem;
			font-weight: 400;
			font-size: 1.7rem;

			${respondTo.medium`
        		padding: 1.5rem;
      		`}
		}
	`}
`;

const PlayerButtons = styled.div`
	${(props) => css`
		display: flex;
		height: 2rem;

		${respondTo.medium`
      		height: 3rem;
    	`}

		> * {
			flex: 1;
			display: flex;
			align-items: center;
			justify-content: center;
			user-select: none;
			border: 1px solid ${props.theme.colorBackground};
			color: ${props.theme.colorPrimary};
			font-size: 1.4rem;

			:hover {
				cursor: pointer;
			}
		}

		> *:nth-child(1) {
			border-bottom-left-radius: ${props.theme.borderRadiusNormal};
			border-right: none;

			:hover {
				background-color: ${props.theme.colorWarning};
				color: ${props.theme.colorTextWarning};
			}
		}

		> *:nth-child(2) {
			:hover {
				background-color: ${props.theme.colorAlert};
				color: ${props.theme.colorTextAlert};
			}
		}

		> *:nth-child(3) {
			border-left: none;
			border-bottom-right-radius: ${props.theme.borderRadiusNormal};

			:hover {
				background-color: ${props.theme.colorDanger};
				color: ${props.theme.colorTextDanger};
			}
		}
	`}
`;

class Server extends Component {
	constructor(props) {
		super(props);

		this.state = {};
	}

	render() {
		return (
			<>
				<div>
					<Heading headingStyle={'title'}>SERVER TITLE</Heading>
					<ServerSummary>
						<p>
							<InfoSpan>{`Players: `}</InfoSpan>
							<InfoSpan>{`Status: `}</InfoSpan>
							<InfoSpan>{`Address: `}</InfoSpan>
						</p>
					</ServerSummary>
				</div>

				<div>
					<Heading headingStyle="subtitle">PLAYER COUNT</Heading>

					<PlayerList>
						<Player>
							<Link to={`/player/idhere`}>
								<h1>Player name</h1>
							</Link>
							<PlayerButtons>
								<div>Warn</div>
								<div>Kick</div>
								<div>Ban</div>
							</PlayerButtons>
						</Player>
					</PlayerList>
				</div>
			</>
		);
	}
}

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(Server);
