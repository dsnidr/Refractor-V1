export function getModalStateFromProps(nextProps, prevState) {
	if (nextProps.success) {
		prevState = {
			...prevState,
			errors: {},
			success: nextProps.success,
		};
	}

	if (nextProps.errors) {
		prevState = {
			...prevState,
			errors: nextProps.errors,
			success: {},
		};
	}

	if (nextProps.player) {
		prevState.player = nextProps.player;
	}

	if (nextProps.serverId) {
		prevState.serverId = nextProps.serverId;
	}

	return prevState;
}
