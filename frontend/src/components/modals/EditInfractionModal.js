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
import { connect } from 'react-redux';
import Modal, { ModalButtonBox, ModalContent } from '../Modal';
import PropTypes from 'prop-types';
import Alert from '../Alert';
import Button from '../Button';
import {
	updateInfraction,
} from '../../redux/infractions/infractionActions';
import { setErrors } from '../../redux/error/errorActions';
import { setSuccess } from '../../redux/success/successActions';
import TextArea from '../TextArea';
import TextInput from '../TextInput';
import DurationShortcuts from '../DurationShortcuts';
import { isNullOrUndefined } from '../../utils/typeUtils';
import { reasonIsValid } from '../../utils/infractionUtils';

class EditInfractionModal extends Component {
	constructor(props) {
		super(props);

		this.state = {
			editableFields: [],
			errors: {},
			success: null,
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		if (nextProps.success) {
			prevState = {
				...prevState,
				errors: {},
				success: nextProps.success,
			};
		}

		if (nextProps.errors) {
			prevState = {
				...prevState,
				errors: nextProps.errors,
				success: null,
			};
		}

		if (nextProps.infraction) {
			const infraction = nextProps.infraction;

			switch (infraction.type) {
				case 'WARNING':
					prevState.editableFields = ['reason'];
					break;
				case 'MUTE':
					prevState.editableFields = ['reason', 'duration'];
					break;
				case 'KICK':
					prevState.editableFields = ['reason'];
					break;
				case 'BAN':
					prevState.editableFields = ['reason', 'duration'];
					break;
				default:
					prevState.editableFields = [];
					break;
			}

			prevState.infraction = infraction;

			// Set current infraction values
			if (infraction.reason && isNullOrUndefined(prevState.reason)) {
				prevState.reason = infraction.reason;
			}

			if (infraction.duration && isNullOrUndefined(prevState.duration)) {
				prevState.duration = infraction.duration;
			}
		}

		return prevState;
	}

	onClose = () => {
		const resetFields = {};
		this.state.editableFields.forEach(
			(field) => (resetFields[field] = null)
		);

		this.setState({
			infraction: null,
			editableFields: [],
			errors: {},
			success: null,
			...resetFields,
		});

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

	onReasonChange = (e) => {
		e.persist();

		this.setState((prevState) => ({
			...prevState,
			reason: e.target.value,
		}));
	};

	onSubmit = () => {
		let { reason } = this.state;
		const { infraction, duration, editableFields } = this.state;

		// Basic validation
		let errorsExist = false;
		const errors = {};

		if (editableFields.includes('reason') && !reasonIsValid(reason)) {
			errorsExist = true;
			errors.reason = 'Please enter a reason for the ban';
		}

		if (
			editableFields.includes('duration') &&
			(duration === null ||
				duration === '' ||
				duration < 0 ||
				isNaN(duration))
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

		// Build submit data
		const submitData = {};
		editableFields.forEach(
			(field) => (submitData[field] = this.state[field])
		);

		this.props.updateInfraction(infraction.id, submitData);
	};

	render() {
		const {
			infraction,
			editableFields,
			success,
			errors,
			reason,
			duration,
		} = this.state;
		const { show } = this.props;

		if (!infraction) {
			return null;
		}

		if (success) {
			setTimeout(() => window.location.reload(), 1500);
		}

		return (
			<Modal show={show} onContainerClick={this.onClose}>
				<h1>
					Editing {String(infraction.type).toLowerCase()} ID{' '}
					{infraction.id}
				</h1>
				<ModalContent>
					<Alert type={'success'} message={success} />
					{editableFields.includes('reason') && (
						<TextArea
							placeholder={'Reason'}
							onChange={this.onReasonChange}
							error={errors.reason}
							defaultValue={reason}
						/>
					)}
					{editableFields.includes('duration') && (
						<>
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
								placeholder={'Duration (minutes)'}
								onKeyPress={this.onDurationKeyPress}
								onChange={this.onDurationChange}
								value={duration}
							/>
						</>
					)}
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
						Save Changes
					</Button>
				</ModalButtonBox>
			</Modal>
		);
	}
}

EditInfractionModal.propTypes = {
	infraction: PropTypes.object,
	show: PropTypes.bool.isRequired,
	onClose: PropTypes.func.isRequired,
	onSuccess: PropTypes.func,
};

const mapStateToProps = (state) => ({
	success: state.success.updateinfraction,
	errors: state.error.updateinfraction,
});

const mapDispatchToProps = (dispatch) => ({
	updateInfraction: (infractionId, data) =>
		dispatch(updateInfraction(infractionId, data)),
	clearErrors: () => dispatch(setErrors('updateinfraction', undefined)),
	clearSuccess: () => dispatch(setSuccess('updateinfraction', undefined)),
});

export default connect(
	mapStateToProps,
	mapDispatchToProps
)(EditInfractionModal);
