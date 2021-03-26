import React, { Component } from 'react';
import Heading from '../../components/Heading';
import TextInput from '../../components/TextInput';
import styled, { css } from 'styled-components';
import Select from '../../components/Select';
import Button from '../../components/Button';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import respondTo from '../../mixins/respondTo';
import { timestampToDateTime } from '../../utils/timeUtils';
import { setLoading } from '../../redux/loading/loadingActions';
import {
	getRecentPlayers,
	searchPlayers,
} from '../../redux/players/playerActions';

const SearchBox = styled.form`
	${(props) => css`
		display: grid;
		grid-template-rows: 4rem 5rem 5rem 5rem 5rem;
		grid-template-columns: auto;
		grid-row-gap: 1rem;

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
	`}
`;

export const SearchResults = styled.div`
	${(props) => css`
		display: flex;
		flex-direction: column;

		// Override react-router link styling
		a {
			text-decoration: none !important;
			color: ${props.theme.colorTextSecondary} !important;
		}
	`}
`;

export const ResultID = styled.div`
	grid-column: 1;
	grid-row: 1;

	${respondTo.medium`
		grid-column: 1;
		grid-row: 1;
	`};
`;

export const ResultName = styled.div`
	grid-column: 1;
	grid-row: 2;

	${respondTo.medium`
		grid-column: 2;
		grid-row: 1;
	`};
`;

export const ResultPlatform = styled.div`
	grid-column: 1;
	grid-row: 3;

	${respondTo.medium`
		grid-column: 3;
		grid-row: 1;
	`};
`;

export const ResultLastSeen = styled.div`
	grid-column: 1;
	grid-row: 4;

	${respondTo.medium`
		grid-column: 4;
		grid-row: 1;
	`};
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
		display: grid;
		grid-template-columns: 1fr;
		grid-template-rows: auto auto auto;

		background-color: ${props.theme.colorAccent};
		padding: 2rem 2rem;
		font-size: 1.8rem;
		border-radius: ${props.theme.borderRadiusNormal};

		:nth-child(odd) {
			background-color: ${props.theme.colorBackground};
		}

		:hover {
			background-color: ${props.theme.colorAccent};
		}

		${respondTo.medium`
			grid-template-columns: 6rem 4fr 1fr 2fr;
			grid-template-rows: 1fr;
		`};
	`}
`;

export const ResultHeading = styled(Result)`
	${(props) => css`
		display: none;

		${respondTo.medium`
			display: grid;
			font-weight: 500;
			color: ${props.theme.colorPrimary};
			background-color: ${props.theme.colorAccent} !important;
		`};
	`}
`;

export const PageSwitcher = styled.div`
	${(props) => css`
		font-size: 1.6rem;
		color: ${props.theme.colorTextPrimary};

		> div {
			display: flex;
			justify-content: center;
		}
	`}
`;

export const PageSwitcherButton = styled.div`
	${(props) => css`
		height: 3rem;
		display: flex;
		align-items: center;
		justify-content: center;
		color: ${props.theme.colorTextPrimary};
		user-select: none;

		:hover {
			cursor: pointer;
		}
	`}
`;

export const DisabledPageSwitcherButton = styled.div`
	${(props) => css`
		height: 3rem;
		display: flex;
		align-items: center;
		justify-content: center;
		color: ${props.theme.colorDisabled};
		user-select: none;
	`}
`;

export const PageSwitcherLabel = styled.div`
	${(props) => css`
		color: ${props.theme.colorTextSecondary};

		width: 8rem;
		height: 3rem;
		display: flex;
		align-items: center;
		justify-content: center;
		user-select: none;
	`}
`;

const RecentPlayers = styled.div`
	${(props) => css`
		display: grid;
		grid-template-columns: 1fr;
		grid-row-gap: 0.5rem;
		grid-column-gap: 0.5rem;
		margin-top: 1rem;

		${respondTo.medium`
		  	grid-template-columns: 1fr 1fr;
	  	`}

		${respondTo.large`
		  	grid-template-columns: 1fr 1fr 1fr;
	  	`}

		// Override react-router link styling
        a {
			text-decoration: none !important;
			color: ${props.theme.colorTextSecondary} !important;
		}
	`}
`;

const RecentPlayer = styled.div`
	${(props) => css`
		display: flex;
		align-items: center;
		padding: 0.5rem;
		font-size: 1.4rem;
		background-color: ${props.theme.colorBackgroundDark};
		border-radius: ${props.theme.borderRadiusNormal};
	`}
`;

const limitInterval = 10;

class Players extends Component {
	constructor(props) {
		super(props);

		this.state = {
			page: 0,
			type: 'name', // name is the default
			term: '',
			currentSearchType: '',
			currentSearchTerm: '',
			errors: {
				term: false,
				type: false,
			},
			searchWasRun: false,
		};
	}

	componentDidMount() {
		this.props.getRecentPlayers();
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		if (nextProps.success) {
			prevState.searchWasRun = true;
			prevState.errors = {};

			nextProps.setLoading(false);
		} else if (nextProps.errors) {
			prevState.errors = {
				...prevState.errors,
				...nextProps.errors,
			};

			nextProps.setLoading(false);
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

		this.props.setLoading(true);
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

		this.props.setLoading(true);
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

		this.props.setLoading(true);
		this.props.searchPlayers(data);
	};

	getPlatform = (player) => {
		if (player.playFabId.length > 0) return <span>PlayFab</span>;
		else if (player.mcuuid.length > 0) return <span>Minecraft</span>;

		return <span>Unknown</span>;
	};

	render() {
		const { errors, page } = this.state;
		const { results, count } = this.props.results;
		const { recentPlayers } = this.props;

		const amountOfPages = Math.ceil(count / limitInterval);

		return (
			<>
				<div>
					<Heading headingStyle="title">Players</Heading>
				</div>

				<div>
					<Heading headingStyle={'subtitle'}>Recent Players</Heading>
					{recentPlayers.length === 0 && (
						<Heading headingStyle={'secondary'}>
							There were no recent players
						</Heading>
					)}
					<RecentPlayers>
						{recentPlayers.map((rp) => (
							<Link to={`/player/${rp.id}`}>
								<RecentPlayer>{rp.currentName}</RecentPlayer>
							</Link>
						))}
					</RecentPlayers>
				</div>

				<SearchBox onSubmit={this.onSearchClick}>
					<Heading headingStyle={'subtitle'}>Search Players</Heading>

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
				{results && results.length > 0 ? (
					<SearchResults>
						<>
							<ResultHeading>
								<ResultID>ID</ResultID>
								<ResultName>Current Name</ResultName>
								<ResultPlatform>Platform</ResultPlatform>
								<ResultLastSeen>Last Seen</ResultLastSeen>
							</ResultHeading>
							{results.map((result, index) => (
								<Link to={`/player/${result.id}`}>
									<Result id={index}>
										<ResultID>
											<MobileLabel>ID: </MobileLabel>
											{result.id}
										</ResultID>
										<ResultName>
											<MobileLabel>Name: </MobileLabel>
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
								</Link>
							))}
						</>
					</SearchResults>
				) : (
					this.state.searchWasRun && (
						<Heading headingStyle={'subtitle'}>
							No results found
						</Heading>
					)
				)}

				{this.state.searchWasRun ? (
					<PageSwitcher>
						<div>
							{page > 0 ? (
								<PageSwitcherButton onClick={this.onPrevPage}>
									Prev
								</PageSwitcherButton>
							) : (
								<DisabledPageSwitcherButton>
									Prev
								</DisabledPageSwitcherButton>
							)}
							<PageSwitcherLabel>{page + 1}</PageSwitcherLabel>
							{results &&
							results.length > 0 &&
							page !== amountOfPages - 1 ? (
								<PageSwitcherButton onClick={this.onNextPage}>
									Next
								</PageSwitcherButton>
							) : (
								<DisabledPageSwitcherButton>
									Next
								</DisabledPageSwitcherButton>
							)}
						</div>
					</PageSwitcher>
				) : null}
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	results: state.players.searchResults,
	recentPlayers: state.players.recentPlayers,
	errors: state.error.searchplayers,
	success: state.success.searchplayers,
});

const mapDispatchToProps = (dispatch) => ({
	setLoading: (isLoading) => dispatch(setLoading('main', isLoading)),
	searchPlayers: (searchData) => dispatch(searchPlayers(searchData)),
	getRecentPlayers: () => dispatch(getRecentPlayers()),
});

export default connect(mapStateToProps, mapDispatchToProps)(Players);
