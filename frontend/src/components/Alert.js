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

const ErrorMessage = styled.div`
	${(props) => css`
		width: 100%;
		background-color: ${props.theme.colorDanger};
		color: ${props.theme.colorTextDanger};
		padding: 0.75rem;
		border-radius: ${props.theme.borderRadiusNormal};
		margin-bottom: 2rem;
		font-size: 1.4rem;
	`}
`;

const SuccessMessage = styled.div`
	${(props) => css`
		width: 100%;
		background-color: ${props.theme.colorSuccess};
		color: ${props.theme.colorTextSuccess};
		padding: 0.75rem;
		border-radius: ${props.theme.borderRadiusNormal};
		margin-bottom: 2rem;
		font-size: 1.4rem;
	`}
`;

const GeneralError = (props) => {
	let { type, message } = props;

	if (!message) {
		return null;
	}

	if (typeof message === 'object') {
		return null;
	}

	if (Array.isArray(message)) {
		message = message[0];
	}

	switch (type) {
		case 'error':
			return <ErrorMessage>{message}</ErrorMessage>;
		case 'success':
			return <SuccessMessage>{message}</SuccessMessage>;
		default:
			return null;
	}
};

export default GeneralError;
