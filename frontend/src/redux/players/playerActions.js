/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

export const GET_PLAYER_SUMMARY = 'GET_PLAYER_SUMMARY';
export const getPlayerSummary = (playerId) => ({
	type: GET_PLAYER_SUMMARY,
	playerId: playerId,
});

export const SET_CURRENT_PLAYER = 'SET_CURRENT_PLAYER';
export const setCurrentPlayer = (playerData) => ({
	type: SET_CURRENT_PLAYER,
	payload: playerData,
});

export const SEARCH_PLAYERS = 'SEARCH_PLAYERS';
export const searchPlayers = (data) => ({
	type: SEARCH_PLAYERS,
	payload: data,
});

export const SET_PLAYER_SEARCH_RESULTS = 'SET_PLAYER_SEARCH_RESULTS';
export const setSearchResults = (results) => ({
	type: SET_PLAYER_SEARCH_RESULTS,
	payload: results,
});

export const GET_RECENT_PLAYERS = 'GET_RECENT_PLAYERS';
export const getRecentPlayers = () => ({
	type: GET_RECENT_PLAYERS,
});

export const SET_RECENT_PLAYERS = 'SET_RECENT_PLAYERS';
export const setRecentPlayers = (recentPlayers) => ({
	type: SET_RECENT_PLAYERS,
	payload: recentPlayers,
});

export const WATCH_PLAYER = 'WATCH_PLAYER';
export const watchPlayer = (playerId) => ({
	type: WATCH_PLAYER,
	playerId: playerId,
});

export const UNWATCH_PLAYER = 'UNWATCH_PLAYER';
export const unwatchPlayer = (playerId) => ({
	type: UNWATCH_PLAYER,
	playerId: playerId,
});

export const SET_PLAYER_WATCHED = 'SET_PLAYER_WATCHED';
export const setPlayerWatched = (playerId, watched) => ({
	type: SET_PLAYER_WATCHED,
	playerId: playerId,
	payload: watched,
});