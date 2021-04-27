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

import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import respondTo from '../mixins/respondTo';
import PropTypes from 'prop-types';

const ModalContainer = styled.div`
	${(props) => css`
		position: fixed;
		overflow-y: scroll;
		z-index: 100000;
		top: 0;
		left: 0;
		background: rgba(0, 0, 0, 0.65) !important;
		width: 100%;
		height: 100%;

		> div {
			background-color: ${props.theme.colorBackground};
		}
	`}
`;

const ModalTopContent = styled.div`
	${(props) => css`
		position: fixed;
		z-index: 100001;
		width: 100%;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		padding: 2rem;
		border-radius: ${props.theme.borderRadiusNormal};
		display: flex;
		flex-direction: column;
		min-height: ${props.tall ? '60rem' : '20rem'};

		${respondTo.medium`
      		width: ${props.wide ? '100rem;' : '50rem;'};
    	`}

		h1 {
			flex: 0 0 4rem;
			border-bottom: 1px solid ${props.theme.colorPrimary};
			margin-bottom: 2rem;
			font-weight: 400;
		}

		${ModalButtonBox} {
			flex: 0 0 4rem;
			margin-top: 1rem;
		}
	`}
`;

export const ModalButtonBox = styled.div`
	${(props) => css`
		width: 100%;
		bottom: 0;
		display: flex;

		> * {
			flex: 1;
		}

		> :first-child {
			margin-right: 0.5rem;
		}

		> :last-child {
			margin-left: 0.5rem;
		}
	`}
`;

export const ModalContent = styled.div`
	${(props) => css`
		flex: 1;
		font-size: 1.6rem;
		margin-bottom: 1rem;
		position: relative;

		> * {
			margin-top: 1rem;
		}
	`}
`;

class Modal extends Component {
	onModalClick = (e) => {
		e.stopPropagation();
	};

	render() {
		const { props } = this;

		if (!props.show) {
			return null;
		}

		return (
			<ModalContainer onClick={props.onContainerClick}>
				<ModalTopContent
					onClick={props.onModalClick || this.onModalClick}
					wide={!!props.wide}
					tall={!!props.tall}
				>
					{props.children}
				</ModalTopContent>
			</ModalContainer>
		);
	}
}

Modal.propTypes = {
	show: PropTypes.bool.isRequired,
	onContainerClick: PropTypes.func,
};

export default Modal;
