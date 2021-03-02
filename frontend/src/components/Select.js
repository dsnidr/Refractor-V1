import React, { Component } from 'react';
import styled, { css } from 'styled-components';

const Error = styled.div`
	${(props) => css`
		color: ${props.theme.colorDanger};
		font-size: 1.2rem;
		height: ${props.theme.errorHeight};
		padding: 0.5rem;
	`}
`;

const StyledSelect = styled.select`
	${(props) => css`
		width: 100%;
		height: 100%;
		background: ${props.theme.inputs.fillInBackground
			? props.theme.colorBorderPrimary
			: 'none'};
		border: 1px solid ${props.theme.colorBorderPrimary};
		border-radius: ${props.theme.borderRadiusNormal};
		font-size: 1.6rem;
		color: ${props.theme.colorTextPrimary};
		padding-left: 1rem;

		-webkit-appearance: none;
		-moz-appearance: none;
	`}
`;

class Select extends Component {
	render() {
		const { props } = this;

		return (
			<div>
				<StyledSelect name={props.name} onChange={props.onChange}>
					{this.props.children}
				</StyledSelect>
				<Error>{props.error ? props.error : null}</Error>
			</div>
		);
	}
}

export default Select;
