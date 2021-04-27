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
import respondTo from '../mixins/respondTo';

const StyledWrapper = styled.div`
	${(props) => css`
		width: 100%;
		min-height: 100vh;
		background-color: ${props.theme.colorBackgroundDark};
		${respondTo.medium`
		  display: flex;
		  align-items: center;
		  justify-content: center;
		`}
	`}
`;

const Wrapper = (props) => <StyledWrapper>{props.children}</StyledWrapper>;

export default Wrapper;
