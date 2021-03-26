import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';
import InfractionPreview from '../../components/InfractionPreview';
import Select from '../../components/Select';
import ServerSelector from '../../components/ServerSelector';
import Button from '../../components/Button';
import PlayerSelector from '../../components/PlayerSelector';

const RecentInfractionsBox = styled.div`
	> :first-child {
		margin-bottom: 1rem;
	}
`;

const RecentInfractions = styled.div`
	${(props) => css`
		display: flex;
		flex-direction: column;
		height: clamp(10rem, 20rem, 30vh);
		overflow-y: scroll;

		> :nth-child(even) {
			background-color: ${props.theme.colorBackground};
		}
	`}
`;

const InfractionSearchBox = styled.div`
	> :nth-child(1) {
		margin-bottom: 0.5rem;
	}

	> :nth-child(2) {
		margin-bottom: 3rem;
		font-size: 1.2rem;
	}
`;

const SearchBox = styled.div`
	display: grid;
	grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
	grid-column-gap: 1rem;

	button {
		height: 4rem;
	}
`;

const ResultsBox = styled.div`
	${(props) => css`
		> :first-child {
			margin-bottom: 1rem;
		}

		> :nth-child(odd) {
			background-color: ${props.theme.colorBackground};
		}
	`}
`;

class Infractions extends Component {
	constructor(props) {
		super(props);

		this.state = {};
	}

	onPlayerSelectionChanged = (player) => {
		console.log('PLAYER SELECTED:', player);
	};

	render() {
		return (
			<>
				<div>
					<Heading headingStyle={'title'}>Infractions</Heading>
				</div>

				<RecentInfractionsBox>
					<Heading headingStyle={'subtitle'}>
						Recent Infractions
					</Heading>

					<RecentInfractions>
						<InfractionPreview
							to={`/player/1?highlight=6`}
							type={'Ban'}
							player={'LMG'}
							date={'2021-03-23'}
							issuer={'Void'}
							duration={'permanent'}
							reason={'test infraction reason'}
						/>
						<InfractionPreview
							type={'Warning'}
							player={'LMG'}
							date={'2021-03-23'}
							issuer={'machetemike'}
							reason={'test infraction reason'}
						/>
						<InfractionPreview
							type={'Kick'}
							player={'LMG'}
							date={'2021-03-23'}
							issuer={'Gladman'}
							reason={'test infraction reason'}
						/>
						<InfractionPreview
							type={'Ban'}
							player={'LMG'}
							date={'2021-03-23'}
							issuer={'Void'}
							duration={'permanent'}
							reason={'test infraction reason'}
						/>
						<InfractionPreview
							type={'Ban'}
							player={'LMG'}
							date={'2021-03-23'}
							issuer={'Void'}
							duration={'permanent'}
							reason={'test infraction reason'}
						/>
						<InfractionPreview
							type={'Ban'}
							player={'LMG'}
							date={'2021-03-23'}
							issuer={'Void'}
							duration={'permanent'}
							reason={'test infraction reason'}
						/>
						<InfractionPreview
							type={'Ban'}
							player={'LMG'}
							date={'2021-03-23'}
							issuer={'Void'}
							duration={'permanent'}
							reason={'test infraction reason'}
						/>
						<InfractionPreview
							type={'Ban'}
							player={'LMG'}
							date={'2021-03-23'}
							issuer={'Void'}
							duration={'permanent'}
							reason={'test infraction reason'}
						/>
					</RecentInfractions>
				</RecentInfractionsBox>

				<InfractionSearchBox>
					<Heading headingStyle={'subtitle'}>
						Search Infractions
					</Heading>

					<p>
						Apply the filters you want and click Search. To leave a
						filter out, set it to it's default value of Select...
					</p>

					<SearchBox>
						<Select title={'infraction type'}>
							<option value={null}>Select...</option>
							<option value={'WARNING'}>Warning</option>
							<option value={'MUTE'}>Mute</option>
							<option value={'KICK'}>Kick</option>
							<option value={'BAN'}>Ban</option>
						</Select>
						<PlayerSelector
							title={'player'}
							onSelect={this.onPlayerSelectionChanged}
						/>
						<Select title={'user'}>
							<option value={null}>Select...</option>
							<option value={1}>User1</option>
							<option value={2}>User2</option>
							<option value={3}>User3</option>
							<option value={4}>User4</option>
						</Select>
						<Select title={'game'}>
							<option value={null}>Select...</option>
							<option value={'Mordhau'}>Mordhau</option>
							<option value={'Minecraft'}>Minecraft</option>
						</Select>
						<ServerSelector default={'Select...'} />
					</SearchBox>

					<Button size={'normal'} color={'primary'}>
						Search
					</Button>
				</InfractionSearchBox>

				<ResultsBox>
					<Heading headingStyle={'subtitle'}>Results</Heading>

					<InfractionPreview
						type={'Ban'}
						player={'LMG'}
						date={'2021-03-23'}
						issuer={'Void'}
						duration={'permanent'}
						reason={'test infraction reason'}
					/>

					<InfractionPreview
						type={'Ban'}
						player={'LMG'}
						date={'2021-03-23'}
						issuer={'Void'}
						duration={'permanent'}
						reason={'test infraction reason'}
					/>

					<InfractionPreview
						type={'Ban'}
						player={'LMG'}
						date={'2021-03-23'}
						issuer={'Void'}
						duration={'permanent'}
						reason={'test infraction reason'}
					/>

					<InfractionPreview
						type={'Ban'}
						player={'LMG'}
						date={'2021-03-23'}
						issuer={'Void'}
						duration={'permanent'}
						reason={'test infraction reason'}
					/>
				</ResultsBox>
			</>
		);
	}
}

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(Infractions);
