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

import React from 'react';
import styled, { css, keyframes } from 'styled-components';
import PropTypes from 'prop-types';
import { buildTimeRemainingString } from '../utils/timeUtils';
import {
	hasFullAccess,
	hasPermission,
	flags,
} from '../permissions/permissions';
import respondTo from '../mixins/respondTo';

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
		width: 100%;
		padding: 1rem;
		background-color: ${props.theme.colorAccent};
		border-radius: ${props.theme.borderRadiusNormal};
		overflow: hidden;

		display: grid;
		grid-template-rows: auto auto auto auto;
		grid-template-columns: auto;

		grid-row-gap: 0.5rem;

		${respondTo.medium`
          	grid-template-columns: 1fr 1fr 1fr 1fr 0.2fr;
          	grid-template-rows: 2rem auto;
		`}

		animation: ${props.highlight
			? css`
					${borderFlashKeyframes(
						props.theme.colorPrimaryLight
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
		grid-row: 1;

		> * {
			font-size: 1rem;
			color: ${props.theme.colorTextPrimary};
			user-select: none;
			margin-right: 1rem;

			:hover {
				cursor: pointer;
			}
		}

		${respondTo.medium`
			grid-column: 5;
			justify-content: right;
			
			> * {
			  margin-left: 1rem;
			  margin-right: 0;
			}
		`}
	`}
`;

const InfractionInfo = styled.div`
	${(props) => css`
		${respondTo.medium`
          grid-row: 2;
          grid-column: span 5;
	  	`}

		width: 100%;
		overflow-wrap: break-word;
		word-wrap: break-word;
		word-break: break-word;
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

	const { perms, isOwnInfraction } = props;

	/* global BigInt */
	const userPerms = BigInt(perms);

	const showEdit =
		hasFullAccess(userPerms) ||
		(isOwnInfraction && flags.EDIT_OWN_INFRACTIONS) ||
		hasPermission(userPerms, flags.EDIT_ANY_INFRACTION);

	const showDelete =
		hasFullAccess(userPerms) ||
		(isOwnInfraction && flags.DELETE_OWN_INFRACTIONS) ||
		hasPermission(userPerms, flags.DELETE_ANY_INFRACTION);

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
					{props.duration !== 0
						? buildTimeRemainingString(props.remaining)
						: 'permanent'}
				</MetaDisplay>
			)}
			<UtilBox>
				{showEdit && <div onClick={props.onEditClick}>Edit</div>}
				{showDelete && <div onClick={props.onDeleteClick}>Delete</div>}
			</UtilBox>
			<InfractionInfo>{props.reason}</InfractionInfo>
		</InfractionBox>
	);
};

Infraction.propTypes = {
	date: PropTypes.any.isRequired,
	issuer: PropTypes.string.isRequired,
	duration: PropTypes.any,
	remaining: PropTypes.any,
	reason: PropTypes.string.isRequired,
	onEditClick: PropTypes.func,
	onDeleteClick: PropTypes.func,
	perms: PropTypes.any,
	isOwnInfraction: PropTypes.bool,
};

export default Infraction;
