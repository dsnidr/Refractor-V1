/*
This file is part of Refractor.

Refractor is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

export default {
	// Main colors
	colorPrimary: '#197fc2',
	colorPrimaryLight: '#3bb1ff',
	colorPrimaryDark: '#197fc2',
	colorDanger: '#FF0035',
	colorAlert: '#FBB13C',
	colorWarning: '#197fc2',
	colorSuccess: '#30db64',
	colorAccent: '#e6e6e6',
	colorDisabled: '#a2b2b8',

	// Background colors
	colorBackground: '#ffffff',
	colorBackgroundDark: '#197fc2',
	colorBackgroundAlt: '#202a2e',

	// Text colors
	colorTextPrimary: '#171717',
	colorTextSecondary: '#171717',
	colorTextLight: '#F2E2D2',
	colorTextDanger: '#f7f7f7',
	colorTextAlert: '#3B3002',
	colorTextWarning: '#F7DEEE',
	colorTextSuccess: '#00171F',
	colorTextDisabled: '#667073',

	// Border colors
	colorBorderPrimary: '#e6e6e6',
	colorBorderSecondary: '#ffffff',

	// Border radius
	borderRadiusNormal: '3px',

	// Shadow definitions
	boxShadowPrimary: {
		'-webkit-box-shadow': '0px 0px 40px -20px rgba(0, 0, 0, 1)',
		'-moz-box-shadow': '0px 0px 40px -20px rgba(0, 0, 0, 1)',
		'box-shadow': '0px 0px 40px -20px rgba(0, 0, 0, 1)',
	},

	///////////////////////////
	// Component Options
	inputs: {
		fillInBackground: true,
		selectArrowColor: 'rgba(0, 168, 232, 1)',
	},

	///////////////////////////
	// Dashboard Options
	dashboard: {
		wrapperPadding: '0',
		gridGap: '0',
		contentPadding: '3rem',
		borderRadius: false,
		topBarShadow: true,
	},
};
