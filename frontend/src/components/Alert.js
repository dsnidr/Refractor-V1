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
	const { type, message } = props;

	if (!message) {
		return null;
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
