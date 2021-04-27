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
import { deleteInfraction } from '../../redux/infractions/infractionActions';
import Button from '../Button';
import Alert from '../Alert';
import { setErrors } from '../../redux/error/errorActions';
import { setSuccess } from '../../redux/success/successActions';

class DeleteInfractionModal extends Component {
	constructor(props) {
		super(props);

		this.state = {
			infraction: null,
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
			prevState.infraction = nextProps.infraction;
		}

		return prevState;
	}

	onClose = () => {
		this.setState((prevState) => ({
			...prevState,
			infraction: null,
		}));

		// Clear errors and success
		this.props.clearSuccess();
		this.props.clearErrors();

		if (this.props.onClose) {
			this.props.onClose();
		}
	};

	onSubmit = () => {
		const { infraction } = this.state;

		this.props.deleteInfraction(infraction.id);
	};

	render() {
		const { infraction, success } = this.state;
		const { show } = this.props;

		if (!infraction) {
			return null;
		}

		if (success) {
			setTimeout(() => window.location.reload(), 1500);
		}

		return (
			<Modal show={show} onContainerClick={this.onClose}>
				<h1>Delete infraction ID {infraction.id}</h1>
				<ModalContent>
					<Alert type={'success'} message={success} />
					<p>
						Are you sure you want to delete this infraction? This
						action cannot be undone.
					</p>
				</ModalContent>
				<ModalButtonBox>
					<Button
						size="normal"
						color="primary"
						onClick={this.onClose}
					>
						Cancel
					</Button>
					<Button
						size="normal"
						color="danger"
						onClick={this.onSubmit}
					>
						Delete
					</Button>
				</ModalButtonBox>
			</Modal>
		);
	}
}

const mapStateToProps = (state) => ({
	success: state.success.deleteinfraction,
	errors: state.error.deleteinfraction,
});

const mapDispatchToProps = (dispatch) => ({
	deleteInfraction: (id) => dispatch(deleteInfraction(id)),
	clearErrors: () => dispatch(setErrors('deleteinfraction', undefined)),
	clearSuccess: () => dispatch(setSuccess('deleteinfraction', undefined)),
});

DeleteInfractionModal.propTypes = {
	show: PropTypes.bool.isRequired,
	infraction: PropTypes.any.isRequired,
	onClose: PropTypes.func,
};

export default connect(
	mapStateToProps,
	mapDispatchToProps
)(DeleteInfractionModal);
