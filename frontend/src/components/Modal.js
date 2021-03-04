import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import respondTo from '../mixins/respondTo';
import PropTypes from 'prop-types';

const ModalContainer = styled.div`
	${(props) => css`
		position: fixed;
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
		min-height: 20rem;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		padding: 2rem;
		border-radius: ${props.theme.borderRadiusNormal};
		display: flex;
		flex-direction: column;

		${respondTo.medium`
      		width: 50rem;
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
