import React, { Component } from 'react';
import styled, { css } from 'styled-components';
import PropTypes from 'prop-types';

const Error = styled.div`
	${(props) => css`
		color: ${props.theme.colorDanger};
		font-size: 1.2rem;
		height: 2rem;
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

const arrowSvgEncoded = (color) =>
	`url(\"data:image/svg+xml;utf8,<svg fill='${color}' height='25' viewBox='0 0 25 25' width='25' xmlns='http://www.w3.org/2000/svg'><path d='M7 10l5 5 5-5z'/><path d='M0 0h25v25H0z' fill='none'/></svg>\")`;

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
		position: relative;

		// Style arrow
		-webkit-appearance: none;
		-moz-appearance: none;
		background: transparent;
		background-image: ${arrowSvgEncoded(
			props.theme.inputs.selectArrowColor
		)};
		background-repeat: no-repeat;
		background-position-x: 100%;
		background-position-y: 50%;
		color: ${props.theme.colorTextPrimary};
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
