import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { reasonIsValid } from '../../utils/infractionUtils';
import Modal, { ModalButtonBox, ModalContent } from '../Modal';
import Alert from '../Alert';
import TextArea from '../TextArea';
import Button from '../Button';

class KickModal extends Component {
	constructor(props) {
		super(props);

		this.state = {
			player: null,
			reason: '',
			error: null,
			success: null,
		};
	}

	static getDerivedStateFromProps(nextProps, prevState) {
		if (nextProps.player) {
			prevState.player = nextProps.player;
		}

		return prevState;
	}

	onClose = () => {
		this.setState((prevState) => ({
			...prevState,
			player: null,
			reason: '',
			error: null,
			success: null,
		}));

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
		const { player } = this.state;

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

		// TODO: Create the kick
		console.log('Creating warning:', player.id, reason);
	};

	render() {
		const { player, success, error } = this.state;
		const { show, inputRef } = this.props;

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
	show: PropTypes.bool.isRequired,
	onClose: PropTypes.func.isRequired,
	onSuccess: PropTypes.func,
	inputRef: PropTypes.object,
};

export default KickModal;
