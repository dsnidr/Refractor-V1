import React from 'react';
import styled, { css } from 'styled-components';
import PropTypes from 'prop-types';

const Shortcuts = styled.div`
	${(props) => css`
		font-size: 1.2rem;
		color: ${props.theme.colorPrimary};

		span {
			margin-right: 1rem;
			user-select: none;

			:hover {
				cursor: pointer;
				color: ${props.theme.colorTextSecondary};
			}
		}
	`}
`;

const DurationShortcuts = (props) => {
	const { durations, onClick } = props;

	return (
		<Shortcuts>
			{durations.map((duration, i) => (
				<span key={i} minutes={duration.minutes} onClick={onClick}>
					{duration.display}
				</span>
			))}
		</Shortcuts>
	);
};

DurationShortcuts.propTypes = {
	durations: PropTypes.array.isRequired,
	onClick: PropTypes.func,
};

export default DurationShortcuts;
