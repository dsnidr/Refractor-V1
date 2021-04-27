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
import { rgba, darken } from 'polished';

const SpinnerBox = styled.div`
	${(props) => css`
		width: 100%;
		height: 100vh;
		position: absolute;
		z-index: 1000;
		top: 0;
		left: 0;
		display: flex;
		justify-content: center;
		align-items: center;
		background: ${rgba(props.theme.colorBackground, 0.9)};
	`}
`;

const LDSRing = styled.div`
	${(props) => css`
		display: inline-block;
		position: relative;
		width: 80px;
		height: 80px;
		div {
			box-sizing: border-box;
			display: block;
			position: absolute;
			width: 64px;
			height: 64px;
			margin: 8px;
			border: 8px solid ${darken(50, props.theme.colorPrimary)};
			border-radius: 50%;
			animation: lds-ring 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite;
			border-color: ${props.theme.colorPrimary} transparent transparent
				transparent;
		}
		div:nth-child(1) {
			animation-delay: -0.45s;
		}
		div:nth-child(2) {
			animation-delay: -0.3s;
		}
		div:nth-child(3) {
			animation-delay: -0.15s;
		}
		@keyframes lds-ring {
			0% {
				transform: rotate(0deg);
			}
			100% {
				transform: rotate(360deg);
			}
		}
	`}
`;

const Spinner = () => {
	return (
		<SpinnerBox>
			<LDSRing>
				<div></div>
				<div></div>
				<div></div>
				<div></div>
			</LDSRing>
		</SpinnerBox>
	);
};

export default Spinner;
