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
