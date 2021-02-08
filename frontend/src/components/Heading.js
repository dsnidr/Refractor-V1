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
