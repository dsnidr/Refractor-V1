import React, { Component } from 'react';
import Modal, { ModalButtonBox, ModalContent } from '../Modal';
import PropTypes from 'prop-types';
import Alert from '../Alert';
import Button from '../Button';

class BasicModal extends Component {
	render() {
		const {
			show,
			onClose,
			onSubmit,
			heading,
			message,
			submitLabel,
			cancelLabel,
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
						{cancelLabel || 'Cancel'}
					</Button>
					<Button size="normal" color="danger" onClick={onSubmit}>
						{submitLabel}
					</Button>
				</ModalButtonBox>
			</Modal>
		);
	}
}

BasicModal.propTypes = {
	show: PropTypes.bool.isRequired,
	onClose: PropTypes.func,
	onSubmit: PropTypes.func,
	heading: PropTypes.string.isRequired,
	message: PropTypes.string.isRequired,
	submitLabel: PropTypes.string.isRequired,
	cancelLabel: PropTypes.string,
	success: PropTypes.any.isRequired,
	error: PropTypes.any.isRequired,
};

export default BasicModal;
