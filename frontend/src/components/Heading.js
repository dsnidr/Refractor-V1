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

import React from 'react';
import styled, { css, ThemeProvider, withTheme } from 'styled-components';

const StyledHeading = styled.h1`
	${(props) => css`
		font-size: ${props.theme.fontSize};
		font-weight: ${props.theme.fontWeight};
		color: ${props.theme.textColor};
	`}
`;

const themes = {
	title: (theme) => ({
		fontSize: '3.2rem',
		textColor: theme.colorTextLighter,
		fontWeight: 400,
	}),
	subtitle: (theme) => ({
		fontSize: '2rem',
		textColor: theme.colorTextPrimary,
		fontWeight: 400,
	}),
	secondary: (theme) => ({
		fontSize: '1.6rem',
		textColor: theme.colorTextSecondary,
		fontWeight: 400,
	}),
};

const Heading = (props) => {
	const theme = props.headingStyle
		? themes[props.headingStyle]
		: themes.title;

	return (
		<ThemeProvider theme={theme(props.theme)}>
			<StyledHeading>{props.children}</StyledHeading>
		</ThemeProvider>
	);
};

export default withTheme(Heading);
