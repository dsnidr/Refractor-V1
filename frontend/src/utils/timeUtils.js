export function timestampToDateTime(timestamp) {
	const date = new Date(timestamp * 1000);

	return `${date.toLocaleString('en-GB', { hour12: true })}`;
}
