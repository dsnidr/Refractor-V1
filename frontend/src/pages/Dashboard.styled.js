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

import styled, { css } from 'styled-components';
import respondTo from '../mixins/respondTo';
import { lighten } from 'polished';

export const Wrapper = styled.div`
	${(props) => css`
		width: 100%;
		min-height: 100vh;
		background-color: ${props.theme.colorBackgroundDark};
		display: grid;
		grid-template-columns: auto 5fr;
		grid-template-rows: 6rem auto;
		grid-row-gap: ${props.theme.dashboard.gridGap};
		padding: ${props.theme.dashboard.wrapperPadding};

		${respondTo.medium`
      		grid-gap: ${props.theme.dashboard.gridGap};
    	`}
	`}
`;

export const Sidebar = styled.div`
	${(props) => css`
		width: 24rem;
		grid-row: 2;
		grid-column: 1;
		background-color: ${props.theme.colorBackgroundAlt};
		border-radius: ${props.theme.dashboard.borderRadius
			? props.theme.borderRadiusNormal
			: 0};
		flex-direction: column;
		user-select: none;
		display: none;

		${respondTo.medium`
      		display: flex;
    	`}
	`}
`;

export const SidebarSection = styled.div`
	${(props) => css`
		padding-bottom: 1rem;

		:first-of-type {
			margin-top: 2rem;
		}

		h1 {
			padding: 0.5rem;
			margin-left: 2rem;
			margin-bottom: 0.5rem;
			font-size: 1.2rem;
			font-weight: 400;
			color: ${props.theme.colorTextLight};
			border-top: 1px solid ${props.theme.colorBackgroundDark};
		}
	`}
`;

export const Topbar = styled.div`
	${(props) => css`
		grid-row: 1;
		grid-column: span 2;
		background-color: ${props.theme.colorBackgroundAlt};
		border-radius: ${props.theme.dashboard.borderRadius
			? props.theme.borderRadiusNormal
			: 0};
		${props.theme.dashboard.topBarShadow
			? `box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.25)`
			: null};
		z-index: 10;
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 0 3rem;

		h1 {
			font-weight: 400;
			font-size: 2.4rem;
			color: ${props.theme.colorPrimary};
		}
	`}
`;

export const TopbarItems = styled.div`
	display: flex;
`;

export const TopbarItem = styled.div`
	${(props) => css`
		font-size: 1.2rem;
		color: ${props.theme.colorTextLight};
		padding: 0.5rem;
		border-radius: ${props.theme.borderRadiusNormal};

		:hover {
			cursor: pointer;
			background-color: ${lighten(0.02, props.theme.colorBackgroundAlt)};
		}
	`}
`;

export const UsernameTopbarItem = styled(TopbarItem)`
	${(props) => css`
		display: none;

		${respondTo.medium`
          margin-right: 2rem;
          color: ${props.theme.colorDisabled};

          :hover {
            cursor: revert;
            background: none;
          }
		`}
	`}
`;

export const Content = styled.div`
	${(props) => css`
		grid-row: 2;
		grid-column: 2;
		background-color: ${props.theme.colorBackgroundDark};
		border-radius: ${props.theme.dashboard.borderRadius
			? props.theme.borderRadiusNormal
			: 0};

		> * {
			padding: 3rem;
			margin-bottom: 2rem;
			background-color: ${props.theme.colorBackground};
			border-radius: ${props.theme.borderRadiusNormal};
			color: ${props.theme.colorTextSecondary};
		}

		${respondTo.medium`
      		padding: ${props.theme.dashboard.contentPadding};
    	`}
	`}
`;

export const DrawerToggle = styled.div`
	${(props) => css`
		display: flex;
		height: 1.6rem;
		flex-direction: column;
		justify-content: space-between;

		div {
			align-self: flex-start;
			width: 2rem;
			border-bottom: 2px solid ${props.theme.colorPrimary};
		}

		&:hover {
			cursor: pointer;
		}

		${respondTo.medium`
      		display: none;
    	`}
	`}
`;

export const Drawer = styled.div`
	${(props) => css`
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100vh;
		z-index: 99999;
		display: flex;

		${respondTo.medium`
		  display: none;
		`}
	`}
`;

export const DrawerOverlay = styled.div`
	${(props) => css`
		background-color: rgba(0, 0, 0, 0.5);
		width: 100%;
		height: 100vh;
		z-index: 5000;
		overflow: visible;
	`}
`;

export const DrawerItems = styled.div`
	${(props) => css`
		position: fixed;
		top: 0;
		left: 0;
		background-color: ${props.theme.colorBackgroundAlt};
		border-right: 1px solid ${props.theme.colorPrimary};
		width: 65%;
		height: 100%;
		z-index: 10000;
		display: flex;
		flex-direction: column;
		color: ${props.theme.colorPrimary};
	`}
`;

export const WebsocketError = styled.div`
	${(props) => css`
		position: fixed;
		overflow-y: scroll;
		z-index: 100000;
		top: 0;
		left: 0;
		background: rgba(0, 0, 0, 0.65) !important;
		width: 100%;
		height: 100%;

		display: flex;
		justify-content: center;
		align-items: center;

		h1 {
			font-size: 3rem;
			text-transform: uppercase;
			color: ${props.theme.colorDanger};
			font-weight: 300;
			text-align: center;
		}

		p {
			font-size: 2rem;
			color: ${props.theme.colorTextSecondary};
			margin-top: 2rem;
			text-align: center;
		}
	`}
`;
