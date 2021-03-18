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
				success: {},
			};
		}

		if (nextProps.player) {
			prevState.player = nextProps.player;
		}

		if (nextProps.serverId) {
			prevState.serverId = nextProps.serverId;
		}

		return prevState;
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
		if (!reasonIsValid(reason)) {
			return this.setState((prevState) => ({
				...prevState,
				errors: {
					reason: 'Please enter a reason for the warning',
				},
			}));
		}

		// Clear error
		this.setState((prevState) => ({
			...prevState,
			errors: {},
		}));

		reason = reason.trim();

		// Create the warning
		this.props.createWarning(serverId, player.id, { reason });
	};

	focus = () => {
		this.inputRef.focus();
	};

	render() {
		const { player, success, errors } = this.state;
		const { show, inputRef } = this.props;

		console.log(success, errors);

		return (
			<Modal show={show} onContainerClick={this.onClose}>
				<h1>Log a warning for {player.currentName}</h1>
				<ModalContent>
					<Alert type="success" message={success} />
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
	serverId: PropTypes.number.isRequired,
	show: PropTypes.bool.isRequired,
	onClose: PropTypes.func.isRequired,
	onSuccess: PropTypes.func,
	inputRef: PropTypes.object,
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
