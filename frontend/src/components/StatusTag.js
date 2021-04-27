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

const StyledStatusTag = styled.div`
	${(props) => css`
		display: inline-block;
		padding: 0.2rem;
		font-size: 1.2rem;
		border-radius: ${props.theme.borderRadiusNormal};
		user-select: none;
	`}
`;

const OnlineStatusTag = styled(StyledStatusTag)`
	${(props) => css`
		background-color: ${props.theme.colorSuccess};
		color: ${props.theme.colorTextSuccess};
	`}
`;

const OfflineStatusTag = styled(StyledStatusTag)`
	${(props) => css`
		background-color: ${props.theme.colorDanger};
		color: ${props.theme.colorTextDanger};
	`}
`;

const StatusTag = (props) => {
	if (props.status === true) {
		return <OnlineStatusTag>Online</OnlineStatusTag>;
	} else {
		return <OfflineStatusTag>Offline</OfflineStatusTag>;
	}
};

export default StatusTag;
