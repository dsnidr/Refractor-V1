export const GET_SERVERS = 'GET_SERVERS';
export const getServers = () => ({
	type: GET_SERVERS,
});

export const SET_SERVERS = 'SET_SERVERS';
export const setServers = (servers) => ({
	type: SET_SERVERS,
	payload: servers,
});
