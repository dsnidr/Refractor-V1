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

import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Modal, { ModalButtonBox, ModalContent } from '../Modal';
import Alert from '../Alert';
import TextArea from '../TextArea';
import Button from '../Button';
import { reasonIsValid } from '../../utils/infractionUtils';
import styled, { css } from 'styled-components';
import TextInput from '../TextInput';
import { createBan } from '../../redux/infractions/infractionActions';
import { setErrors } from '../../redux/error/errorActions';
import { setSuccess } from '../../redux/success/successActions';
import { connect } from 'react-redux';
import { getModalStateFromProps } from './modalHelpers';
import ServerSelector from '../ServerSelector';
import DurationShortcuts from '../DurationShortcuts';

class BanModal extends Component {
	constructor(props) {
		super(props);

		this.state = {
			player: null,
			reason: '',
			duration: '',
			errors: {},
			success: null,
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		return getModalStateFromProps(nextProps, prevState);
	}

	onClose = () => {
		this.setState((prevState) => ({
			...prevState,
			player: null,
			reason: '',
			duration: '',
			errors: {},
			success: null,
		}));

		// Clear errors and success messages
		this.props.clearErrors();
		this.props.clearSuccess();

		if (this.props.onClose) {
			this.props.onClose();
		}
	};

	onDurationShortcutClick = (event) => {
		const minutes = parseInt(event.target.getAttribute('minutes'));

		this.setState((prevState) => ({
			...prevState,
			duration: minutes,
		}));
	};

	onReasonChange = (e) => {
		e.persist();

		this.setState((prevState) => ({
			...prevState,
			reason: e.target.value,
		}));
	};

	onDurationChange = (e) => {
		let value = parseInt(e.target.value);

		if (isNaN(value)) {
			value = '';
		}

		this.setState((prevState) => ({
			...prevState,
			duration: value,
		}));
	};

	onDurationKeyPress = (e) => {
		const keyCode = e.keyCode || e.which;
		const keyValue = String.fromCharCode(keyCode);

		// Make sure that the key pressed was a number
		if (!/[0-9]/.test(keyValue)) {
			e.preventDefault();
		}
	};

	onSubmit = () => {
		let { reason } = this.state;
		const { player, serverId, duration } = this.state;

		// Basic validation
		let errorsExist = false;
		const errors = {};

		if (!reasonIsValid(reason)) {
			errorsExist = true;
			errors.reason = 'Please enter a reason for the ban';
		}

		if (!serverId) {
			errorsExist = true;
			errors.server = 'Please select a server';
		}

		if (
			duration === null ||
			duration === '' ||
			duration < 0 ||
			isNaN(duration)
		) {
			errorsExist = true;
			errors.duration = "Please enter the ban's duration";
		}

		this.setState((prevState) => ({
			...prevState,
			errors: errors,
		}));

		if (errorsExist) {
			return;
		}

		reason = reason.trim();

		this.props.createBan(serverId, player.id, { reason, duration });
	};

	onServerSelectionChange = (e) => {
		this.setState((prevState) => ({
			...prevState,
			serverId: parseInt(e.target.value),
		}));
	};

	render() {
		const { player, success, errors, duration } = this.state;
		const { show, inputRef, reload } = this.props;

		if (success) {
			setTimeout(() => {
				if (reload) {
					window.location.reload();
				} else {
					this.onClose();
				}
			}, 1500);
		}

		return (
			<Modal show={show} onContainerClick={this.onClose}>
				<h1>Log a ban for {player.currentName}</h1>
				<ModalContent>
					<Alert type="success" message={success} />
					{!this.props.serverId && (
						<ServerSelector
							onChange={this.onServerSelectionChange}
							error={errors.server}
						/>
					)}
					<TextArea
						placeholder={'Reason for ban'}
						onChange={this.onReasonChange}
						error={errors.reason}
						ref={inputRef}
					/>
					<DurationShortcuts
						durations={[
							{ minutes: 60, display: '1 hour' },
							{ minutes: 1440, display: '1 day' },
							{ minutes: 10080, display: '1 week' },
							{ minutes: 40320, display: '1 month' },
							{ minutes: 0, display: 'permanent' },
						]}
						onClick={this.onDurationShortcutClick}
					/>
					<TextInput
						type={'text'}
						placeholder={'Ban duration (minutes)'}
						onKeyPress={this.onDurationKeyPress}
						onChange={this.onDurationChange}
						value={duration}
						error={errors.duration}
					/>
				</ModalContent>
				<ModalButtonBox>
					<Button size="normal" color="danger" onClick={this.onClose}>
						Cancel
					</Button>
					<Button
						size="normal"
						color="primary"
						onClick={this.onSubmit}
					>
						Submit Ban
					</Button>
				</ModalButtonBox>
			</Modal>
		);
	}
}

BanModal.propTypes = {
	player: PropTypes.object,
	serverId: PropTypes.number.isRequired,
	show: PropTypes.bool.isRequired,
	onClose: PropTypes.func.isRequired,
	onSuccess: PropTypes.func,
	inputRef: PropTypes.object,
	reload: PropTypes.bool,
};

const mapStateToProps = (state) => ({
	success: state.success.createban,
	errors: state.error.createban,
});

const mapDispatchToProps = (dispatch) => ({
	createBan: (serverId, playerId, data) =>
		dispatch(createBan(serverId, playerId, data)),
	clearErrors: () => dispatch(setErrors('createban', undefined)),
	clearSuccess: () => dispatch(setSuccess('createban', undefined)),
});

export default connect(mapStateToProps, mapDispatchToProps)(BanModal);
