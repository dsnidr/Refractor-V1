import axios from 'axios';
import decodeJWT from 'jwt-decode';
import { refreshToken } from '../api/authApi';

export function setToken(token) {
	// Store token in local storage
	localStorage.setItem('token', token);

	// Instruct axios to send this token with all requests
	axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
}

export function destroyToken() {
	localStorage.removeItem('token');

	axios.defaults.headers.common['Authorization'] = undefined;
}

export function getToken() {
	const token = localStorage.getItem('token');

	if (!token) {
		return null;
	}

	// Instruct axios to send this token with all requests
	axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;

	return token;
}

export function decodeToken(token) {
	if (!token) {
		return false;
	}

	try {
		return decodeJWT(token);
	} catch (err) {
		console.log('decodeToken err', err);
		return null;
	}
}

export function tokenIsCurrent(decodedToken) {
	if (!decodedToken) {
		return false;
	}

	const now = Math.round(new Date().getTime() / 1000);

	return decodedToken.exp > now;
}

function tokenIsValid(token) {
	const decoded = decodeToken(token);

	if (!decoded) {
		return false;
	}

	if (!tokenIsCurrent(decoded)) {
		return false;
	}

	return true;
}

// This massively ugly function is an axios interceptor. Basically what it does is when we receive a 401 unauthorized
// status code from a request to the API, we assume the user's auth token is expired so we try to refresh it.
export function setInterceptors(store) {
	axios.interceptors.response.use(
		(response) => {
			return response;
		},
		(error) => {
			// Configure the problematic request not to retry to prevent looping
			const originalRequest = error.config;

			// If the request was unauthorized because the JWT used was invalid, assume it's expired and try to refresh it.
			if (
				error.response.status === 401 &&
				!originalRequest._retry &&
				!tokenIsValid(localStorage.getItem('token'))
			) {
				originalRequest._retry = true;

				return refreshToken()
					.then((res) => {
						// If refresh was successful, store the new token
						if (res.status === 200) {
							setToken(res.data.payload);

							// Update bearer token on originalRequest and then retry it
							originalRequest.headers.Authorization = `Bearer ${res.data.payload}`;

							return axios(originalRequest);
						}
					})
					.catch(() => {
						// If token could not be refreshed, log out.
						destroyToken();

						// Set user to null to trigger logout
						store.dispatch({
							type: 'SET_USER',
							payload: {
								isAuthenticated: false,
							},
						});
					});
			}

			return Promise.reject(error);
		}
	);
}
