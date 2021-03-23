import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import PropTypes from 'prop-types';

const Error = styled.div`
	${(props) => css`
		color: ${props.theme.colorDanger};
		font-size: 1.2rem;
		height: ${props.theme.errorHeight};
		padding: 0.5rem;
	`}
`;

const Title = styled.span`
	${(props) => css`
		font-size: 1rem;
		color: ${props.theme.colorTextPrimary};
		position: absolute;
		top: -1.6rem;
		left: 0.5rem;
	`}
`;

const StyledSelect = styled.select`
	${(props) => css`
		width: 100%;
		height: 4rem;
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

const SelectBox = styled.div`
	position: relative;
`;

class Select extends Component {
	render() {
		const { props } = this;

		return (
			<SelectBox>
				<Title>{props.title}</Title>
				<StyledSelect name={props.name} onChange={props.onChange}>
					{this.props.children}
				</StyledSelect>
				<Error>{props.error ? props.error : null}</Error>
			</SelectBox>
		);
	}
}

Select.propTypes = {
	name: PropTypes.string,
	onChange: PropTypes.func,
	error: PropTypes.any,
	title: PropTypes.string,
};

export default Select;
