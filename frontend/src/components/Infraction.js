import React from 'react';
import styled, { css, keyframes } from 'styled-components';
import { respondTo } from '../mixins/respondTo';
import PropTypes from 'prop-types';
import { buildTimeRemainingString } from '../utils/timeUtils';

const borderFlashKeyframes = (color) => keyframes`
	0% {
	  border: 0;
	}
  
  	50% {
	  border: 5px solid ${color};
	}
  
  	100% {
	  border: 0;
	}
`;

const InfractionBox = styled.div`
	${(props) => css`
		padding: 1rem;
		background-color: ${props.theme.colorAccent};
		border-radius: ${props.theme.borderRadiusNormal};

		display: grid;
		grid-template-columns: 1fr 1fr 1fr 1fr 0.2fr;
		grid-template-rows: 2rem auto;
		grid-row-gap: 0.5rem;

		animation: ${props.highlight
			? css`
					${borderFlashKeyframes(
						props.theme.colorBorderPrimary
					)} 2s ease-in-out
			  `
			: ''};
	`}
`;

const MetaDisplay = styled.div`
	${(props) => css`
		font-size: 1.4rem;

		span {
			color: ${props.theme.colorPrimary};
		}
	`}
`;

const UtilBox = styled.div`
	${(props) => css`
      display: flex;
      justify-content: right;
	  
	  grid-column: 5;

      > * {
        font-size: 1rem;
        color: ${props.theme.colorTextPrimary};
        margin-left: 1rem;
        user-select: none;

        :hover {
          cursor: pointer;
        }
	`}
`;

const InfractionInfo = styled.div`
	${(props) => css`
		grid-row: 2;
		grid-column: span 5;

		font-size: 1.4rem;
	`}
`;

const Infraction = (props) => {
	let duration = props.duration;
	if (!isNaN(duration)) {
		if (duration === 0) {
			duration = 'permanent';
		} else {
			duration = `${duration} minutes`;
		}
	}

	return (
		<InfractionBox highlight={!!props.highlight} ref={props.highlightRef}>
			<MetaDisplay>
				<span>Date:</span> {props.date}
			</MetaDisplay>
			<MetaDisplay>
				<span>Issued by:</span> {props.issuer}
			</MetaDisplay>
			{duration && (
				<MetaDisplay>
					<span>Duration:</span> {duration}
				</MetaDisplay>
			)}
			{props.remaining && (
				<MetaDisplay>
					<span>Time left:</span>{' '}
					{buildTimeRemainingString(props.remaining)}
				</MetaDisplay>
			)}
			<UtilBox>
				<div onClick={props.onEditClick}>Edit</div>
				<div onClick={props.onDeleteClick}>Delete</div>
			</UtilBox>
			<InfractionInfo>{props.reason}</InfractionInfo>
		</InfractionBox>
	);
};

Infraction.propTypes = {
	date: PropTypes.any.isRequired,
	issuer: PropTypes.string.isRequired,
	duration: PropTypes.number,
	remaining: PropTypes.any,
	reason: PropTypes.string.isRequired,
	onEditClick: PropTypes.func,
	onDeleteClick: PropTypes.func,
};

export default Infraction;
