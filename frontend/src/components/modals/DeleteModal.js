import React, { Component } from 'react';
import Modal, { ModalButtonBox, ModalContent } from '../Modal';
import PropTypes from 'prop-types';
import Alert from '../Alert';
import Button from '../Button';

class DeleteModal extends Component {
	render() {
		const {
			show,
			onClose,
			onSubmit,
			heading,
			message,
			success,
			error,
		} = this.props;

		return (
			<Modal show={show} onContainerClick={onClose}>
				<h1>{heading}</h1>
				<ModalContent>
					<Alert type="success" message={success} />
					<Alert type="error" message={error} />
					{message}
				</ModalContent>
				<ModalButtonBox>
					<Button size="normal" color="primary" onClick={onClose}>
						Cancel
					</Button>
					<Button size="normal" color="danger" onClick={onSubmit}>
						Delete
					</Button>
				</ModalButtonBox>
			</Modal>
		);
	}
}

DeleteModal.propTypes = {
	show: PropTypes.bool.isRequired,
	onClose: PropTypes.func,
	onSubmit: PropTypes.func,
	heading: PropTypes.string.isRequired,
	message: PropTypes.string.isRequired,
	success: PropTypes.any.isRequired,
	error: PropTypes.any.isRequired,
};

export default DeleteModal;
