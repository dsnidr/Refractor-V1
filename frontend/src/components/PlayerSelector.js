import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { connect } from 'react-redux';
import {
	searchPlayers,
	setSearchResults,
} from '../redux/players/playerActions';
import Modal, { ModalButtonBox, ModalContent } from './Modal';
import Button from './Button';
import Heading from './Heading';
import TextInput from './TextInput';
import Select from './Select';
import respondTo from '../mixins/respondTo';
import {
	DisabledPageSwitcherButton,
	MobileLabel,
	PageSwitcher,
	PageSwitcherButton,
	PageSwitcherLabel,
	ResultHeading,
	ResultID,
	ResultLastSeen,
	ResultName,
	ResultPlatform,
	SearchResults,
} from '../pages/DashboardPages/Players';
import { Link } from 'react-router-dom';
import { timestampToDateTime } from '../utils/timeUtils';
import { setSuccess } from '../redux/success/successActions';
import { setErrors } from '../redux/error/errorActions';
import PropTypes from 'prop-types';

const Wrapper = styled.div``;

const limitInterval = 5;

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
		height: 18rem;

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
          	height: auto;
		  
		  	h1 {
            	grid-column: span 4;
		  	}
		`}

		select {
			height: 5rem;
		}

		button {
			height: 5rem !important;
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

const Result = styled.div`
	${(props) => css`
		display: flex;
		justify-content: space-between;

		background-color: ${props.theme.colorAccent};
		padding: 1rem 1rem;
		font-size: 1.4rem;
		border-radius: ${props.theme.borderRadiusNormal};

		:nth-child(even) {
			background-color: ${props.theme.colorBackground};
		}

		:hover {
			background-color: ${props.theme.colorAccent};
		}

		${respondTo.medium`
		  	display: grid;
			grid-template-columns: 6rem 4fr 2fr 2fr;
			grid-template-rows: 1fr;
		`};

		${props.highlight
			? `background: ${props.theme.colorPrimaryDark} !important;`
			: ''}

		${MobileLabel} {
			display: none;
		}
	`}
`;

const CustomResultHeading = styled(ResultHeading)`
	font-size: 1.6rem;
	padding: 1rem;
`;

const CustomPageSwitcher = styled(PageSwitcher)`
	position: absolute;
	bottom: 0;
	width: 100%;
`;

const CustomResultPlatform = styled(ResultPlatform)`
	display: none;

	${respondTo.medium`
	  display: block;
	`}
`;

const CustomResultLastSeen = styled(ResultLastSeen)`
	display: none;

	${respondTo.medium`
	  display: block;
	`}
`;

class PlayerSelector extends Component {
	constructor(props) {
		super(props);

		this.state = {
			showModal: false,
			selectedPlayer: null,
			errors: {},
			page: 0,
			type: 'name',
			term: '',
			currentSearchType: '',
			currentSearchTerm: '',
			searchWasRun: false,
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		if (nextProps.success) {
			prevState.searchWasRun = true;
			prevState.errors = {};
		} else if (nextProps.errors) {
			prevState.errors = {
				...prevState.errors,
				...nextProps.errors,
			};
		}

		return prevState;
	}

	onSearchClick = (e) => {
		e.preventDefault();

		const data = {
			limit: limitInterval,
			offset: 0,
			type: this.state.type,
			term: this.state.term,
		};

		// Basic validation
		let errors = {};

		if (!data.type) {
			errors.type = 'Please select a search type';
		}

		if (!data.term) {
			errors.term = 'Please enter a search term';
		}

		this.setState((prevState) => ({
			...prevState,
			errors: errors,
		}));

		if (Object.keys(errors).length > 0) {
			return;
		}

		data.term = data.term.toUpperCase();

		this.props.searchPlayers(data);

		// Set current search fields
		this.setState((prevState) => ({
			...prevState,
			page: 0,
			currentSearchType: data.type,
			currentSearchTerm: data.term,
		}));
	};

	onChange = (e) => {
		e.persist();
		e.stopPropagation();

		this.setState((prevState) => ({
			...prevState,
			[e.target.name]: e.target.value,
		}));
	};

	onNextPage = () => {
		const nextPage = this.state.page + 1;

		const data = {
			limit: limitInterval,
			offset: nextPage * limitInterval,
			type: this.state.currentSearchType,
			term: this.state.currentSearchTerm,
		};

		// Update page in state
		this.setState((prevState) => ({
			...prevState,
			page: nextPage,
		}));

		this.props.searchPlayers(data);
	};

	onPrevPage = () => {
		const prevPage = this.state.page - 1;

		const data = {
			limit: limitInterval,
			offset: prevPage * limitInterval,
			type: this.state.currentSearchType,
			term: this.state.currentSearchTerm,
		};

		// Update page in state
		this.setState((prevState) => ({
			...prevState,
			page: prevPage,
		}));

		this.props.searchPlayers(data);
	};

	onClose = () => {
		// Clear search results
		this.props.clearResults();

		this.setState({
			showModal: false,
			selectedPlayer: null,
			errors: {},
			page: 0,
			type: 'name',
			term: '',
			currentSearchType: '',
			currentSearchTerm: '',
			searchWasRun: false,
		});

		// Clear success and errors
		this.props.clearSuccess();
		this.props.clearErrors();
	};

	onSelectPlayer = (player) => () => {
		this.setState((prevState) => ({
			...prevState,
			selectedPlayer: player,
		}));
	};

	onSelectClicked = () => {
		this.setState((prevState) => ({
			...prevState,
			showModal: true,
		}));
	};

	onSubmit = () => {
		if (this.props.onSelect) {
			this.props.onSelect(this.state.selectedPlayer);
		}

		this.onClose();
	};

	getPlatform = (player) => {
		if (player.playFabId.length > 0) return <span>PlayFab</span>;
		else if (player.mcuuid.length > 0) return <span>Minecraft</span>;

		return <span>Unknown</span>;
	};

	render() {
		const {
			showModal,
			errors,
			page,
			searchWasRun,
			selectedPlayer,
		} = this.state;
		const { title, value, results: stateRes } = this.props;
		const { results, count } = stateRes;

		const amountOfPages = Math.ceil(count / limitInterval);

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
								<option value="id">ID</option>
								<option value="playfabid">PlayFabID</option>
								<option value="mcuuid">Minecraft UUID</option>
							</Select>

							<Button
								size={'inline'}
								onClick={this.onSearchClick}
							>
								Search
							</Button>
						</SearchBox>

						{results && results.length > 0 ? (
							<SearchResults>
								<>
									<CustomResultHeading>
										<ResultID>ID</ResultID>
										<ResultName>Current Name</ResultName>
										<ResultPlatform>
											Platform
										</ResultPlatform>
										<ResultLastSeen>
											Last Seen
										</ResultLastSeen>
									</CustomResultHeading>
									{results.map((result, index) => (
										<Result
											id={index}
											onClick={this.onSelectPlayer(
												result
											)}
											highlight={
												selectedPlayer &&
												selectedPlayer.id === result.id
											}
										>
											<ResultID>
												<MobileLabel>ID: </MobileLabel>
												{result.id}
											</ResultID>
											<ResultName>
												<MobileLabel>
													Name:{' '}
												</MobileLabel>
												{result.currentName}
											</ResultName>
											<ResultPlatform>
												<MobileLabel>
													Platform:{' '}
												</MobileLabel>
												{this.getPlatform(result)}
											</ResultPlatform>
											<ResultLastSeen>
												<MobileLabel>
													Last Seen:{' '}
												</MobileLabel>
												{timestampToDateTime(
													result.lastSeen
												)}
											</ResultLastSeen>
										</Result>
									))}
								</>
							</SearchResults>
						) : (
							!!searchWasRun && (
								<Heading headingStyle={'subtitle'}>
									No results found
								</Heading>
							)
						)}

						{this.state.searchWasRun ? (
							<CustomPageSwitcher>
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
									{results &&
									results.length > 0 &&
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
							</CustomPageSwitcher>
						) : null}
					</ModalContent>
					<ModalButtonBox>
						<Button
							size="normal"
							color="danger"
							onClick={this.onClose}
						>
							Cancel
						</Button>
						<Button
							size="normal"
							color="primary"
							disabled={!!!selectedPlayer}
							onClick={this.onSubmit}
						>
							Select
						</Button>
					</ModalButtonBox>
				</Modal>

				<StyledPlayerSelector onClick={this.onSelectClicked}>
					<Title>{title}</Title>
					<PlayerName>{value}</PlayerName>
				</StyledPlayerSelector>
			</Wrapper>
		);
	}
}

PlayerSelector.propTypes = {
	onSelect: PropTypes.func.isRequired,
	title: PropTypes.string,
	value: PropTypes.string.isRequired,
};

const mapStateToProps = (state) => ({
	results: state.players.searchResults,
	errors: state.error.searchplayers,
	success: state.success.searchplayers,
});

const mapDispatchToProps = (dispatch) => ({
	searchPlayers: (searchData) => dispatch(searchPlayers(searchData)),
	clearResults: () => dispatch(setSearchResults([])),
	clearSuccess: () => dispatch(setSuccess('searchplayers', undefined)),
	clearErrors: () => dispatch(setErrors('searchplayers', undefined)),
});

export default connect(mapStateToProps, mapDispatchToProps)(PlayerSelector);
