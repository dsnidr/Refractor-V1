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
import Button from '../Button';
import { reasonIsValid } from '../../utils/infractionUtils';
import { connect } from 'react-redux';
import TextArea from '../TextArea';
import { createWarning } from '../../redux/infractions/infractionActions';
import { setErrors } from '../../redux/error/errorActions';
import { setSuccess } from '../../redux/success/successActions';
import { getModalStateFromProps } from './modalHelpers';
import ServerSelector from '../ServerSelector';

class WarnModal extends Component {
	constructor(props) {
		super(props);

		this.state = {
			player: null,
			serverId: null,
			reason: '',
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
			serverId: null,
			reason: '',
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

	onReasonChange = (e) => {
		if (e.persist) {
			e.persist();
		}

		this.setState((prevState) => ({
			...prevState,
			reason: e.target.value,
		}));
	};

	onSubmit = () => {
		let { reason } = this.state;
		const { player, serverId } = this.state;

		// Basic validation
		let errorsExist = false;
		const errors = {};

		// Basic validation
		if (!reasonIsValid(reason)) {
			errorsExist = true;
			errors.reason = 'Please enter a reason for the ban';
		}

		if (!serverId) {
			errorsExist = true;
			errors.server = 'Please select a server';
		}

		this.setState((prevState) => ({
			...prevState,
			errors: errors,
		}));

		if (errorsExist) {
			return;
		}

		reason = reason.trim();

		// Create the warning
		this.props.createWarning(serverId, player.id, { reason });
	};

	focus = () => {
		this.inputRef.focus();
	};

	onServerSelectionChange = (e) => {
		this.setState((prevState) => ({
			...prevState,
			serverId: parseInt(e.target.value),
		}));
	};

	render() {
		const { player, success, errors } = this.state;
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
				<h1>Log a warning for {player.currentName}</h1>
				<ModalContent>
					<Alert type="success" message={success} />
					{!this.props.serverId && (
						<ServerSelector
							onChange={this.onServerSelectionChange}
							error={errors.server}
						/>
					)}
					<TextArea
						placeholder={'Reason for warning'}
						onChange={this.onReasonChange}
						error={errors.reason}
						ref={inputRef}
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
						Submit Warning
					</Button>
				</ModalButtonBox>
			</Modal>
		);
	}
}

WarnModal.propTypes = {
	player: PropTypes.object.isRequired,
	serverId: PropTypes.number,
	show: PropTypes.bool.isRequired,
	onClose: PropTypes.func.isRequired,
	onSuccess: PropTypes.func,
	inputRef: PropTypes.object,
	reload: PropTypes.bool,
};

const mapStateToProps = (state) => ({
	success: state.success.createwarning,
	errors: state.error.createwarning,
});

const mapDispatchToProps = (dispatch) => ({
	createWarning: (serverId, playerId, data) =>
		dispatch(createWarning(serverId, playerId, data)),
	clearErrors: () => dispatch(setErrors('createwarning', undefined)),
	clearSuccess: () => dispatch(setSuccess('createwarning', undefined)),
});

export default connect(mapStateToProps, mapDispatchToProps)(WarnModal);
