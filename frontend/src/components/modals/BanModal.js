import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Modal, { ModalButtonBox, ModalContent } from '../Modal';
import Alert from '../Alert';
import TextArea from '../TextArea';
import Button from '../Button';
import { reasonIsValid } from '../../utils/infractionUtils';
import styled, { css } from 'styled-components';
import TextInput from '../TextInput';

const Shortcuts = styled.div`
	${(props) => css`
		font-size: 1.2rem;
		color: ${props.theme.colorPrimary};

		span {
			margin-right: 1rem;

			:hover {
				cursor: pointer;
				color: ${props.theme.colorTextSecondary};
			}
		}
	`}
`;

class BanModal extends Component {
	constructor(props) {
		super(props);

		this.state = {
			player: null,
			reason: '',
			duration: '',
			errors: {
				reason: null,
				duration: null,
			},
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
		let { reason } = this.state;
		const { player, duration } = this.state;

		let errorsExist = false;
		const errors = {};

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

		// TODO: Create the ban
		console.log('Create ban', player.id, reason, duration);
	};

	render() {
		const { player, success, errors, duration } = this.state;
		const { show, inputRef } = this.props;

		return (
			<Modal show={show} onContainerClick={this.onClose}>
				<h1>Log a ban for {player.currentName}</h1>
				<ModalContent>
					<Alert type="success" message={success} />
					<TextArea
						placeholder={'Reason for ban'}
						onChange={this.onReasonChange}
						error={errors.reason}
						ref={inputRef}
					/>
					<Shortcuts>
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
							minutes={40320}
							onClick={this.onDurationShortcutClick}
						>
							1 month
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
						placeholder={'Ban duration (minutes)'}
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
						Submit Ban
					</Button>
				</ModalButtonBox>
			</Modal>
		);
	}
}

BanModal.propTypes = {
	player: PropTypes.object,
	show: PropTypes.bool.isRequired,
	onClose: PropTypes.func.isRequired,
	onSuccess: PropTypes.func,
	inputRef: PropTypes.object,
};

export default BanModal;
