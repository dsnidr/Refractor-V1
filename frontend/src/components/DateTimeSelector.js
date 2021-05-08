import React from 'react';
import styled, { css } from 'styled-components';
import PropTypes from 'prop-types';
import DateTime from 'react-datetime';
import 'react-datetime/css/react-datetime.css';

const Title = styled.span`
	${(props) => css`
		font-size: 1rem;
		color: ${props.theme.colorTextPrimary};
		position: absolute;
		top: -1.6rem;
		left: 0.5rem;
	`}
`;

const DateTimeTheme = styled.div`
	${(props) => css`
		position: relative;

		.rdt {
			.form-control {
				width: 100%;
				height: 4rem;
				color: ${props.theme.colorTextPrimary};
				font-size: 1.4rem;
				padding: 1rem;

				background: none;
				background-color: ${props.theme.inputs.fillInBackground
					? props.theme.colorBorderPrimary
					: null};

				border: 1px solid ${props.theme.colorBorderPrimary};
				border-radius: ${props.theme.borderRadiusNormal};
			}

			.rdtPicker {
				background-color: ${props.theme.colorBackground};
				border: 1px solid ${props.theme.colorBorderPrimary};
				border-radius: ${props.theme.borderRadiusNormal};
				color: ${props.theme.colorTextSecondary};

				.rdtActive {
					background-color: ${props.theme.colorPrimaryDark};
				}

				.rdtSwitch:hover {
					cursor: pointer;
				}
			}

			.rdtTimeToggle:hover {
				background-color: ${props.theme.colorBackground};
			}
		}
	`}
`;

const DateTimeSelector = (props) => {
	console.log('DateTimeSelector value', props.value);

	return (
		<DateTimeTheme>
			<Title>{props.title}</Title>
			<DateTime onChange={props.onChange} value={props.value} />
		</DateTimeTheme>
	);
};

DateTimeSelector.propTypes = {
	title: PropTypes.string,
	onChange: PropTypes.func.isRequired,
	value: PropTypes.any.isRequired,
};

export default DateTimeSelector;
