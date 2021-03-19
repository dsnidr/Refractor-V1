import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { reasonIsValid } from '../../utils/infractionUtils';
import Modal, { ModalButtonBox, ModalContent } from '../Modal';
import Alert from '../Alert';
import TextArea from '../TextArea';
import Button from '../Button';
import { createKick } from '../../redux/infractions/infractionActions';
import { setErrors } from '../../redux/error/errorActions';
import { setSuccess } from '../../redux/success/successActions';
import { connect } from 'react-redux';
import { getModalStateFromProps } from './modalHelpers';

class KickModal extends Component {
	constructor(props) {
		super(props);

		this.state = {
			player: null,
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
		if (!reasonIsValid(reason)) {
			return this.setState((prevState) => ({
				...prevState,
				error: 'Please enter a reason for the kick',
			}));
		}

		// Clear error
		this.setState((prevState) => ({
			...prevState,
			error: null,
		}));

		reason = reason.trim();

		this.props.createKick(serverId, player.id, { reason });
	};

	render() {
		const { player, success, error } = this.state;
		const { show, inputRef } = this.props;

		if (success) {
			setTimeout(() => this.onClose(), 1500);
		}

		return (
			<Modal show={show} onContainerClick={this.onClose}>
				<h1>Log a kick for {player.currentName}</h1>
				<ModalContent>
					<Alert type="success" message={success} />
					<TextArea
						placeholder={'Reason for kick'}
						onChange={this.onReasonChange}
						error={error}
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
						Submit Kick
					</Button>
				</ModalButtonBox>
			</Modal>
		);
	}
}

KickModal.propTypes = {
	player: PropTypes.object,
	serverId: PropTypes.number.isRequired,
	show: PropTypes.bool.isRequired,
	onClose: PropTypes.func.isRequired,
	onSuccess: PropTypes.func,
	inputRef: PropTypes.object,
};

const mapStateToProps = (state) => ({
	success: state.success.createkick,
	errors: state.error.createkick,
});

const mapDispatchToProps = (dispatch) => ({
	createKick: (serverId, playerId, data) =>
		dispatch(createKick(serverId, playerId, data)),
	clearErrors: () => dispatch(setErrors('createkick', undefined)),
	clearSuccess: () => dispatch(setSuccess('createkick', undefined)),
});

export default connect(mapStateToProps, mapDispatchToProps)(KickModal);
