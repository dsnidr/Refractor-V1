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
