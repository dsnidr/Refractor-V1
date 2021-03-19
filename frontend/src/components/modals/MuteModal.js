import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Modal, { ModalButtonBox, ModalContent } from '../Modal';
import Alert from '../Alert';
import TextArea from '../TextArea';
import Button from '../Button';
import { reasonIsValid } from '../../utils/infractionUtils';
import styled, { css } from 'styled-components';
import TextInput from '../TextInput';
import { createMute } from '../../redux/infractions/infractionActions';
import { setErrors } from '../../redux/error/errorActions';
import { setSuccess } from '../../redux/success/successActions';
import { connect } from 'react-redux';
import { getModalStateFromProps } from './modalHelpers';

const Shortcuts = styled.div`
	${(props) => css`
		font-size: 1.2rem;
		color: ${props.theme.colorPrimary};

		span {
			margin-right: 1rem;
			user-select: none;

			:hover {
				cursor: pointer;
				color: ${props.theme.colorTextSecondary};
			}
		}
	`}
`;

class MuteModal extends Component {
	constructor(props) {
		super(props);

		this.state = {
			player: null,
			serverId: null,
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
			serverId: null,
			reason: '',
			duration: '',
			errors: {},
			success: null,
		}));

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
		let { reason, errors: prevErrors } = this.state;
		const { player, serverId, duration } = this.state;

		let errorsExist = false;
		const errors = {
			...prevErrors,
		};

		// Basic validation
		if (!reasonIsValid(reason)) {
			errorsExist = true;
			errors.reason = 'Please enter a reason for the ban';
		}

		if (duration === null || duration < 0 || isNaN(duration)) {
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

		this.props.createMute(serverId, player.id, { reason, duration });
	};

	render() {
		const { player, success, errors, duration } = this.state;
		const { show, inputRef } = this.props;

		if (success) {
			setTimeout(() => this.onClose(), 1500);
		}

		return (
			<Modal show={show} onContainerClick={this.onClose}>
				<h1>Log a mute for {player.currentName}</h1>
				<ModalContent>
					<Alert type="success" message={success} />
					<TextArea
						placeholder={'Reason for mute'}
						onChange={this.onReasonChange}
						error={errors.reason}
						ref={inputRef}
					/>
					<Shortcuts>
						<span
							minutes={30}
							onClick={this.onDurationShortcutClick}
						>
							30 minutes
						</span>
						<span
							minutes={60}
							onClick={this.onDurationShortcutClick}
						>
							1 hour
						</span>
						<span
							minutes={1440}
							onClick={this.onDurationShortcutClick}
						>
							1 day
						</span>
						<span
							minutes={10080}
							onClick={this.onDurationShortcutClick}
						>
							1 week
						</span>
						<span
							minutes={0}
							onClick={this.onDurationShortcutClick}
						>
							Permanent
						</span>
					</Shortcuts>
					<TextInput
						type={'text'}
						placeholder={'Mute duration (minutes)'}
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
						Submit Mute
					</Button>
				</ModalButtonBox>
			</Modal>
		);
	}
}

MuteModal.propTypes = {
	player: PropTypes.object,
	serverId: PropTypes.number.isRequired,
	show: PropTypes.bool.isRequired,
	onClose: PropTypes.func.isRequired,
	onSuccess: PropTypes.func,
	inputRef: PropTypes.object,
};

const mapStateToProps = (state) => ({
	success: state.success.createmute,
	errors: state.error.createmute,
});

const mapDispatchToProps = (dispatch) => ({
	createMute: (serverId, playerId, data) =>
		dispatch(createMute(serverId, playerId, data)),
	clearErrors: () => dispatch(setErrors('createmute', undefined)),
	clearSuccess: () => dispatch(setSuccess('createmute', undefined)),
});

export default connect(mapStateToProps, mapDispatchToProps)(MuteModal);
