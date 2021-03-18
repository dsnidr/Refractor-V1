export function reasonIsValid(reason) {
	return !(
		!reason ||
		typeof reason !== 'string' ||
		reason.trim().length === 0
	);
}
