import React from 'react';
import styled, { css } from 'styled-components';

const Label = styled.label`
	display: inline-block;
	vertical-align: middle;
`;

const CheckboxContainer = styled.div`
	display: inline-block;
	vertical-align: middle;
`;

const Icon = styled.svg`
	fill: none;
	stroke: white;
	stroke-width: 2px;
`;
// Hide checkbox visually but remain accessible to screen readers.
// Source: https://polished.js.org/docs/#hidevisually
const HiddenCheckbox = styled.input.attrs({ type: 'checkbox' })`
	border: 0;
	clip: rect(0 0 0 0);
	clippath: inset(50%);
	height: 1px;
	margin: -1px;
	overflow: hidden;
	padding: 0;
	position: absolute;
	white-space: nowrap;
	width: 1px;
`;

const StyledCheckbox = styled.div`
	${(props) => css`
		display: inline-block;
		width: 1.6rem;
		height: 1.6rem;
		background: ${props.checked
			? props.theme.colorPrimaryDark
			: props.theme.colorAccent};
		border-radius: ${props.theme.borderRadiusNormal};
		transition: all 150ms;

		${Icon} {
			visibility: ${props.checked ? 'visible' : 'hidden'};
		}
	`}
`;

const LabelSpan = styled.span`
	${(props) => css`
		margin-left: 1rem;
		user-select: none;
		font-size: 1.4rem;
		color: ${props.disabled ? props.theme.colorTextDisabled : 'auto'};
	`}
`;

const Checkbox = ({ className, checked, disabled, ...props }) => (
	<Label>
		<CheckboxContainer className={className}>
			<HiddenCheckbox checked={checked} disabled={disabled} {...props} />
			<StyledCheckbox checked={checked} disabled={disabled}>
				<Icon viewBox="0 0 24 24">
					<polyline points="20 6 9 17 4 12" />
				</Icon>
			</StyledCheckbox>
		</CheckboxContainer>
		<LabelSpan disabled={disabled}>{props.label}</LabelSpan>
	</Label>
);

export default Checkbox;
