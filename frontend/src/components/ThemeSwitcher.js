import React, { Component } from 'react';
import themes from '../themes';
import { connect } from 'react-redux';
import styled, { css } from 'styled-components';
import { setTheme } from '../redux/theme/themeActions';

const Switcher = styled.select`
	${(props) => css`
		margin: 0 2rem;
		background: ${props.theme.inputs.fillInBackground
			? props.theme.colorBorderPrimary
			: 'none'};
		color: ${props.theme.colorTextPrimary};
		border: 1px solid ${props.theme.colorBorderPrimary};
		border-radius: ${props.theme.borderRadiusNormal};
		width: 6rem;
		text-align: center;
		-webkit-appearance: none;
		-moz-appearance: none;
		text-indent: 1px;
		text-overflow: '';
	`}
`;

class ThemeSwitcher extends Component {
	onSelectionChange = (e) => {
		this.props.setTheme(e.target.value);
	};

	render() {
		const options = [];

		Object.keys(themes).forEach((name) =>
			options.push(
				<option key={name} value={name}>
					{name}
				</option>
			)
		);

		return (
			<Switcher
				onChange={this.onSelectionChange}
				value={this.props.theme}
			>
				{options}
			</Switcher>
		);
	}
}

const mapStateToProps = (state) => ({
	theme: state.theme,
});

const mapDispatchToProps = (dispatch) => ({
	setTheme: (theme) => dispatch(setTheme(theme)),
});

export default connect(mapStateToProps, mapDispatchToProps)(ThemeSwitcher);
