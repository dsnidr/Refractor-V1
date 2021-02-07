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
