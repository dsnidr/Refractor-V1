export const SET_LOADING = 'SET_LOADING';
export const setLoading = (scope, isLoading) => ({
	type: SET_LOADING,
	scope: scope,
	payload: isLoading,
});
