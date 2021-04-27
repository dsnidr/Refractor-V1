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
