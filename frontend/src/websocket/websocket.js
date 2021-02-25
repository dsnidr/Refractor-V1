import { w3cwebsocket as W3CWebSocket } from 'websocket';

let pingInterval = null;

// Actions is an object of expected redux dispatch attached actions
export function newWebsocket(websocketURI, actions, handleOpen, handleClose) {
	const wsClient = new W3CWebSocket(websocketURI);
	wsClient.onopen = onOpen(wsClient, handleOpen);
	wsClient.onclose = onClose(wsClient, handleClose);
	wsClient.onmessage = onMessage(wsClient, actions);

	return wsClient;
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

const onMessage = (client, actions) => (msg) => {
	const wsMsg = JSON.parse(msg.data);
	const { type, body } = wsMsg;

	console.log('Received message: ', wsMsg);

	// TODO: Handle message types and take proper actions
	switch (type) {
		case 'player-join':
			actions.addPlayer(body.serverId, {
				id: body.id,
				playerGameId: body.playerGameId,
				currentName: body.name,
			});
			break;
		case 'player-quit':
			actions.removePlayer(body.serverId, {
				id: body.id,
				playerGameId: body.playerGameId,
				currentName: body.name,
			});
			break;
	}
};
