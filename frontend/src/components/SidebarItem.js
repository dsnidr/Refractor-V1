import React from 'react';
import styled, { css } from 'styled-components';
import { Link } from 'react-router-dom';
import { lighten } from 'polished';

const StyledSidebarItem = styled(Link)`
	${(props) => css`
		height: 4rem;
		font-size: 1.4rem;
		margin: 0 1rem;
		padding: 0 2rem;
		text-decoration: none !important;
		border-radius: ${props.theme.dashboard.borderRadius
			? props.theme.borderRadiusNormal
			: 0};
		display: flex;
		align-items: center;

		:first-child {
			margin-top: 1rem;
		}

		:hover {
			background-color: ${lighten(0.02, props.theme.colorBackgroundAlt)};
		}
	`}
`;

const SidebarItemIcon = styled.div`
	${(props) => css`
		fill: ${props.theme.colorTextLight};
		width: 1.6rem;
		height: 1.6rem;
		margin-right: 1rem;
	`}
`;

const SidebarItemLabel = styled.span`
	${(props) => css`
		align-self: stretch;
		color: ${props.theme.colorTextLight};
		font-size: 1.4rem;
		display: flex;
		align-items: center;
		padding: 1rem;
	`}
`;

const SidebarItem = (props) => (
	<StyledSidebarItem to={props.to}>
		<SidebarItemIcon>{props.icon}</SidebarItemIcon>
		<SidebarItemLabel>{props.children}</SidebarItemLabel>
	</StyledSidebarItem>
);

export default SidebarItem;
