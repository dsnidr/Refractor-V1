import React, { Component } from 'react';
import { connect } from 'react-redux';
import {
	Content,
	Drawer,
	DrawerItems,
	DrawerOverlay,
	DrawerToggle,
	Sidebar,
	SidebarSection,
	Topbar,
	TopbarItem,
	Wrapper,
} from './Dashboard.styled';
import SidebarItem from '../components/SidebarItem';
import { ReactComponent as House } from '../assets/house.svg';
import { ReactComponent as ServerIcon } from '../assets/server.svg';
import { ReactComponent as Avatar } from '../assets/avatar.svg';
import { ReactComponent as SingleServer } from '../assets/server-1.svg';
import { ReactComponent as List } from '../assets/list.svg';
import { destroyToken } from '../utils/tokenUtils';
import RequireAccessLevel from '../components/RequireAccessLevel';
import ThemeSwitcher from '../components/ThemeSwitcher';
import Main from './DashboardPages/Main';
import { Route, Switch } from 'react-router';
import { getGames } from '../redux/game/gameActions';
import Server from './DashboardPages/Server';

class Dashboard extends Component {
	constructor(props) {
		super(props);

		this.state = {
			drawerOpen: false,
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		if (nextProps.games === null) {
			nextProps.getGames();
		}

		return prevState;
	}

	toggleDrawer = () => {
		this.setState({
			drawerOpen: !this.state.drawerOpen,
		});
	};

	closeDrawer = () => {
		this.setState({
			drawerOpen: false,
		});
	};

	onLogOutClick = () => {
		destroyToken();
		window.location.href = '/';
	};

	render() {
		let { games } = this.props;

		console.log(games);

		if (games === null) {
			return null;
		}

		games = Object.values(games);

		console.log(games);

		return (
			<>
				{this.state.drawerOpen ? (
					<Drawer>
						<DrawerItems>
							<SidebarItem to="/" icon={<House />}>
								Dashboard
							</SidebarItem>
							<SidebarItem to="/players" icon={<Avatar />}>
								Players
							</SidebarItem>
							<SidebarItem to="/infractions" icon={<List />}>
								Infractions
							</SidebarItem>

							<div>
								{this.props.games.map((game) => (
									<SidebarSection key={game.name}>
										<SidebarItem icon={<SingleServer />}>
											Test Server
										</SidebarItem>
									</SidebarSection>
								))}
								<SidebarSection>
									<h1>&#62; admin</h1>
									<SidebarItem
										to="/servers"
										icon={<ServerIcon />}
									>
										Servers
									</SidebarItem>
									<SidebarItem to="/users" icon={<Avatar />}>
										Users
									</SidebarItem>
								</SidebarSection>
							</div>
						</DrawerItems>
						<DrawerOverlay onClick={this.closeDrawer} />
					</Drawer>
				) : null}

				<Wrapper>
					<Topbar>
						<DrawerToggle onClick={this.toggleDrawer}>
							<div />
							<div />
							<div />
						</DrawerToggle>
						<h1>REFRACTOR</h1>
						<div className={'items'}>
							<TopbarItem onClick={this.onLogOutClick}>
								LOG OUT
							</TopbarItem>
						</div>
					</Topbar>

					<Sidebar>
						<div>
							<SidebarItem to="/" icon={<House />}>
								Dashboard
							</SidebarItem>
							<SidebarItem to="/players" icon={<Avatar />}>
								Players
							</SidebarItem>
							<SidebarItem to="/infractions" icon={<List />}>
								Infractions
							</SidebarItem>
						</div>
						<div>
							{games.map((game) => (
								<SidebarSection key={game.name}>
									<h1>&#62; {game.name.toLowerCase()}</h1>
									{game.servers.map((server) => (
										<SidebarItem
											key={server.id}
											icon={<SingleServer />}
											to={`/server/${server.id}`}
										>
											{server.name}
										</SidebarItem>
									))}
								</SidebarSection>
							))}
							<RequireAccessLevel minAccessLevel={10}>
								<SidebarSection>
									<h1>&#62; admin</h1>
									<SidebarItem
										to="/servers"
										icon={<ServerIcon />}
									>
										Servers
									</SidebarItem>
									<SidebarItem to="/users" icon={<Avatar />}>
										Users
									</SidebarItem>
								</SidebarSection>
							</RequireAccessLevel>

							<SidebarSection>
								<h1>&#62; theme</h1>
								<ThemeSwitcher />
							</SidebarSection>
						</div>
					</Sidebar>
					<Content>
						<Switch>
							<Route exact path="/" component={Main} />
							<Route
								exact
								path="/server/:id"
								component={Server}
							/>
						</Switch>
					</Content>
				</Wrapper>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	games: state.game,
});

const mapDispatchToProps = (dispatch) => ({
	getGames: () => dispatch(getGames()),
});

export default connect(mapStateToProps, mapDispatchToProps)(Dashboard);
