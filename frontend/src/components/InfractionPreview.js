import React from 'react';
import styled, { css, keyframes } from 'styled-components';
import { respondTo } from '../mixins/respondTo';
import PropTypes from 'prop-types';
import { buildTimeRemainingString } from '../utils/timeUtils';
import { Link } from 'react-router-dom';
import { lighten } from 'polished';
import { typeHasDuration } from '../utils/infractionUtils';

const InfractionBox = styled(Link)`
	${(props) => css`
		padding: 1rem;
		background-color: ${props.theme.colorAccent};
		border-radius: ${props.theme.borderRadiusNormal};

		display: grid;
		grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
		grid-template-rows: 2rem auto;
		grid-row-gap: 0.5rem;

		// Override annoying react router link styling. Why do they do this?
		text-decoration: none;
		text-underline: none;
		color: ${props.theme.colorTextSecondary};
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

const InfractionReason = styled.div`
	${(props) => css`
		grid-row: 2;
		grid-column: span 5;

		font-size: 1.4rem;
	`}
`;

const InfractionPreview = (props) => {
	let duration = props.duration;
	if (typeof duration === 'number') {
		if (duration === 0) {
			duration = 'permanent';
		} else {
			duration = `${duration} minutes`;
		}
	} else {
		duration = null;
	}

	if (!typeHasDuration(props.type)) {
		duration = null;
	}

	return (
		<InfractionBox
			to={props.to}
			highlight={!!props.highlight}
			ref={props.highlightRef}
		>
			<MetaDisplay>
				<span>Type: </span>
				{props.type}
			</MetaDisplay>
			<MetaDisplay>
				<span>Player: </span>
				{props.player}
			</MetaDisplay>
			<MetaDisplay>
				<span>Issuer: </span>
				{props.issuer}
			</MetaDisplay>
			<MetaDisplay>
				<span>Date: </span>
				{props.date}
			</MetaDisplay>
			{duration && (
				<MetaDisplay>
					<span>Duration: </span>
					{duration}
				</MetaDisplay>
			)}
			<InfractionReason>{props.reason}</InfractionReason>
		</InfractionBox>
	);
};

InfractionPreview.propTypes = {
	type: PropTypes.string.isRequired,
	date: PropTypes.any.isRequired,
	issuer: PropTypes.string.isRequired,
	duration: PropTypes.any,
	reason: PropTypes.string.isRequired,
	to: PropTypes.string,
};

export default InfractionPreview;
