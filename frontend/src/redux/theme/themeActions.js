import { SET_THEME } from './constants';

export const setTheme = (theme) => ({
	type: SET_THEME,
	payload: theme,
});
