import { w3cwebsocket as W3CWebSocket } from 'websocket';
import { store } from '../redux/store';
import {
	addPlayerToServer,
	removePlayerFromServer,
	setServerStatus,
} from '../redux/servers/serverActions';
import { addChatMessage } from '../redux/chat/chatActions';

let currentWebsocket = null;
let pingInterval = null;

// Actions is an object of expected redux dispatch attached actions
export function newWebsocket(websocketURI, actions, handleOpen, handleClose) {
	const wsClient = new W3CWebSocket(websocketURI);
	wsClient.onopen = onOpen(wsClient, handleOpen);
	wsClient.onclose = onClose(wsClient, handleClose);
	wsClient.onmessage = onMessage(wsClient, actions);

	currentWebsocket = wsClient;
	return currentWebsocket;
}

export function getCurrentWebsocket() {
	return currentWebsocket;
}

const onOpen = (client, handleOpen) => () => {
	console.log('Websocket connection opened');

	pingInterval = setInterval(() => {
		console.log('Pinging server');
		client.send(
			JSON.stringify({
				type: 'ping',
				body: '',
			})
		);
	}, 40000);

	handleOpen();
};

const onClose = (client, handleClose) => (data) => {
	console.log(
		'WebSocket Connection Closed',
		new Date().toLocaleString('en-GB'),
		'Reason:',
		data.reason
	);

	clearInterval(pingInterval);

	handleClose(data);
};

const onMessage = () => (msg) => {
	const wsMsg = JSON.parse(msg.data);
	const { type, body } = wsMsg;

	console.log('Received message: ', wsMsg);

	switch (type) {
		case 'player-join':
			store.dispatch(
				addPlayerToServer(body.serverId, {
					id: body.id,
					playerGameId: body.playerGameId,
					currentName: body.name,
				})
			);
			break;
		case 'player-quit':
			store.dispatch(
				removePlayerFromServer(body.serverId, {
					id: body.id,
					playerGameId: body.playerGameId,
					currentName: body.name,
				})
			);
			break;
		case 'server-online':
			store.dispatch(setServerStatus(body, true));
			break;
		case 'server-offline':
			store.dispatch(setServerStatus(body, false));
			break;
		case 'chat':
			store.dispatch(addChatMessage(body));
			break;
		default:
			console.log('Unknown message type received:', type);
	}
};
