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
import {
	decodeToken,
	destroyToken,
	getToken,
	tokenIsCurrent,
} from '../utils/tokenUtils';
import RequireAccessLevel from '../components/RequireAccessLevel';
import ThemeSwitcher from '../components/ThemeSwitcher';
import Main from './DashboardPages/Main';
import { Route, Switch } from 'react-router';
import { getGames } from '../redux/games/gameActions';
import Server from './DashboardPages/Server';
import { refreshToken } from '../api/authApi';
import { newWebsocket } from '../websocket/websocket';
import {
	addPlayerToServer,
	getServers,
	removePlayerFromServer,
} from '../redux/servers/serverActions';

let reconnectInterval;
let reconnectTaskStarted = false;

class Dashboard extends Component {
	constructor(props) {
		super(props);

		this.state = {
			drawerOpen: false,
			wsClient: null,
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		if (nextProps.games === null) {
			nextProps.getGames();
		}

		if (nextProps.servers === null) {
			nextProps.getServers();
		}

		// We don't want to create a websocket until we know the user was fetched.
		// This is because we rely on the getUserInfo call made in App.js during the auth check process
		// to refresh our JWTs if needed. If we didn't wait for a user to be defined in props then we
		// could send the servers an expired which would fail the websocket auth check.
		if (!nextProps.user || prevState.wsClient !== null) {
			return prevState;
		}

		// Create and configure a websocket connection if the above statement didn't return
		let webSocketUri = `${
			window.location.protocol === 'https:' ? 'wss' : 'ws'
		}://${process.env.REACT_APP_DOMAIN}/ws?auth=${getToken()}`;

		// handleClose function which is responsible for re-creating the websocket client after it was terminated.
		// This is NOT where the initial connection is created.
		const handleClose = () => {
			if (reconnectTaskStarted) {
				return;
			}

			reconnectTaskStarted = true;
			reconnectInterval = setTimeout(async () => {
				console.log('Attempting websocket reconnection...');

				const decodedToken = decodeToken(getToken());
				if (!tokenIsCurrent(decodedToken)) {
					try {
						await refreshToken();
					} catch (err) {
						console.log('Could not refresh expired token', err);
					}
				}

				webSocketUri = `${
					window.location.protocol === 'https:' ? 'wss' : 'ws'
				}://${process.env.REACT_APP_DOMAIN}/ws?auth=${getToken()}`;

				prevState.wsClient = newWebsocket(
					webSocketUri,
					{
						addPlayer: nextProps.addPlayer,
						removePlayer: nextProps.removePlayer,
					},
					() => {
						reconnectTaskStarted = false;
						clearInterval(reconnectInterval);
					},
					() => {
						handleClose();
					}
				);
			}, 15000);
		};

		// Initial websocket client connection
		prevState.wsClient = newWebsocket(
			webSocketUri,
			{
				addPlayer: nextProps.addPlayer,
				removePlayer: nextProps.removePlayer,
			},
			() => {
				clearInterval(reconnectInterval);
			},
			handleClose
		);

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

		if (games === null) {
			return null;
		}

		games = Object.values(games);

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
	user: state.user.self,
	games: state.games,
	servers: state.servers,
});

const mapDispatchToProps = (dispatch) => ({
	getGames: () => dispatch(getGames()),
	getServers: () => dispatch(getServers()),
	addPlayer: (serverId, player) =>
		dispatch(addPlayerToServer(serverId, player)),
	removePlayer: (serverId, player) =>
		dispatch(removePlayerFromServer(serverId, player)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Dashboard);
