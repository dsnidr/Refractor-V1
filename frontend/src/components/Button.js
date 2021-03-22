import React from 'react';
import styled, { css, ThemeProvider, withTheme } from 'styled-components';
import { lighten, darken } from 'polished';

const StyledButton = styled.button`
	${(props) => css`
		width: 100%;
		height: 5rem;
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: ${props.theme.backgroundColor};
		color: ${props.theme.textColor};
		border: 1px solid ${props.theme.borderColor};
		border-radius: ${props.theme.borderRadiusNormal};
		font-size: 1.6rem;
		text-decoration: none !important;

		:hover {
			background-color: ${props.theme.hover.backgroundColor};
			cursor: pointer;
		}

		:disabled {
			background-color: ${props.theme.colorDisabled};
			color: ${props.theme.colorTextDisabled};
			border: 1px solid ${darken(0.1, props.theme.colorDisabled)};
			opacity: 0.8;
		}

		transition: all 0.2s;
	`}
`;

const SmallStyledButton = styled(StyledButton)`
	width: auto;
	height: auto;
	padding: 0 1rem;
	font-size: 1.4rem;
`;

const NormalStyledButton = styled(StyledButton)`
	width: 12rem;
	height: 4rem;
	padding: 0 1rem;
	font-size: 1.4rem;
`;

const themes = {
	primary: (theme) => ({
		backgroundColor: theme.colorPrimary,
		textColor: theme.colorBackground,
		borderColor: theme.colorPrimary,

		hover: {
			backgroundColor: lighten(0.05, theme.colorPrimary),
		},
	}),
	danger: (theme) => ({
		backgroundColor: theme.colorDanger,
		textColor: theme.colorTextDanger,
		borderColor: theme.colorDanger,

		hover: {
			backgroundColor: lighten(0.05, theme.colorDanger),
		},
	}),
	alert: (theme) => ({
		backgroundColor: theme.colorAlert,
		textColor: theme.colorTextAlert,
		borderColor: theme.colorAlert,

		hover: {
			backgroundColor: lighten(0.05, theme.colorAlert),
		},
	}),
	success: (theme) => ({
		backgroundColor: theme.colorSuccess,
		textColor: theme.colorTextSuccess,
		borderColor: theme.colorSuccess,

		hover: {
			backgroundColor: lighten(0.05, theme.colorSuccess),
		},
	}),
};

const Button = (props) => {
	const theme = props.color ? themes[props.color] : themes.primary;

	let Component;

	switch (props.size) {
		case 'small':
			Component = SmallStyledButton;
			break;
		case 'normal':
			Component = NormalStyledButton;
			break;
		default:
			Component = StyledButton;
			break;
	}

	return (
		<ThemeProvider theme={theme(props.theme)}>
			<Component
				onClick={props.onClick}
				disabled={props.disabled}
				type={props.type ? props.type : 'button'}
			>
				{props.children}
			</Component>
		</ThemeProvider>
	);
};

export default withTheme(Button);
