import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { connect } from 'react-redux';
import { searchPlayers } from '../redux/players/playerActions';
import Modal, { ModalButtonBox, ModalContent } from './Modal';
import Button from './Button';
import Heading from './Heading';
import TextInput from './TextInput';
import Select from './Select';
import respondTo from '../mixins/respondTo';

const Wrapper = styled.div``;

const StyledPlayerSelector = styled.div`
	${(props) => css`
		height: 4rem;
		width: 100%;
		font-size: 1.6rem;
		position: relative;

		display: flex;
		align-items: center;
		padding-left: 1rem;

		color: ${props.theme.colorTextPrimary};
		border: 1px solid ${props.theme.colorBorderPrimary};
		border-radius: ${props.theme.borderRadiusNormal};
		user-select: none;

		:hover {
			cursor: pointer;
		}
	`}
`;

const PlayerName = styled.span`
	overflow-x: hidden;
`;

const SearchBox = styled.form`
	${(props) => css`
		display: grid;
		grid-template-rows: 4rem 5rem 5rem 5rem 5rem;
		grid-row-gap: 1rem;
		grid-template-columns: auto;

		:nth-child(1) {
			grid-column: 1;
		}

		:nth-child(2) {
			grid-column: 2;
		}

		h1 {
			grid-row: 1;
		}

		${respondTo.medium`
			grid-template-rows: 4rem 5rem;
			grid-template-columns: 7fr 2fr 1fr;
			grid-column-gap: 1rem;
			grid-row-gap: 0;
		  
		  	h1 {
            	grid-column: span 4;
		  	}
		`}

		select {
			height: 5rem;
		}

		button {
			height: 5rem;
		}
	`}
`;

const Title = styled.span`
	${(props) => css`
		font-size: 1rem;
		color: ${props.theme.colorTextPrimary};
		position: absolute;
		top: -1.6rem;
		left: 0.5rem;
	`}
`;

class PlayerSelector extends Component {
	constructor(props) {
		super(props);

		this.state = {
			showModal: false,
			selectedPlayer: null,
			errors: {},
		};
	}

	onClose = () => {
		this.setState((prevState) => ({
			...prevState,
			showModal: false,
			errors: {},
		}));
	};

	onSelectClicked = () => {
		this.setState((prevState) => ({
			...prevState,
			showModal: true,
		}));
	};

	render() {
		const { showModal, errors } = this.state;
		const { title } = this.props;

		return (
			<Wrapper>
				<Modal
					show={showModal}
					onContainerClick={this.onClose}
					wide={true}
					tall={true}
				>
					<h1>Select a player</h1>
					<ModalContent>
						<SearchBox onSubmit={this.onSearchClick}>
							<TextInput
								type="text"
								name="term"
								placeholder="Search term"
								onChange={this.onChange}
								error={errors.term}
							/>

							<Select
								name="type"
								onChange={this.onChange}
								error={errors.type}
							>
								<option value="name">Name</option>
								<option value="playfabid">PlayFabID</option>
								<option value="mcuuid">Minecraft UUID</option>
							</Select>

							<Button size={'small'} onClick={this.onSearchClick}>
								Search
							</Button>
						</SearchBox>
					</ModalContent>
					<ModalButtonBox>
						<Button
							size="normal"
							color="danger"
							onClick={this.onClose}
						>
							Cancel
						</Button>
						<Button size="normal" color="primary">
							Select
						</Button>
					</ModalButtonBox>
				</Modal>

				<StyledPlayerSelector onClick={this.onSelectClicked}>
					<Title>{title}</Title>
					<PlayerName>Select...</PlayerName>
				</StyledPlayerSelector>
			</Wrapper>
		);
	}
}

const mapStateToProps = (state) => ({
	results: state.players.searchResults,
});

const mapDispatchToProps = (dispatch) => ({
	searchPlayers: (searchData) => dispatch(searchPlayers(searchData)),
});

export default connect(mapStateToProps, mapDispatchToProps)(PlayerSelector);
