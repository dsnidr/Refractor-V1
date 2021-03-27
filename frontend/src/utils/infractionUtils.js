export function reasonIsValid(reason) {
	return !(
		!reason ||
		typeof reason !== 'string' ||
		reason.trim().length === 0
	);
}

export function typeHasDuration(type) {
	return type === 'MUTE' || type === 'BAN';
}
