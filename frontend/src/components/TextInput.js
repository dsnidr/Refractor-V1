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
import styled, { css, ThemeProvider } from 'styled-components';

const Wrapper = styled.div`
	${(props) => css`
		width: 100%;
		display: flex;
		flex-direction: column;
		height: ${props.theme.height};
	`}
`;

const Error = styled.div`
	${(props) => css`
		color: ${props.theme.colorDanger};
		font-size: 1.2rem;
		height: ${props.theme.errorHeight};
		padding: 0.5rem;
	`}
`;

const InputMain = styled.div`
	${(props) => css`
		background-color: ${props.theme.inputs.fillInBackground
			? props.theme.colorBorderPrimary
			: null};
		height: ${props.theme.inputHeight};
		display: flex;
		border: 1px solid ${props.theme.colorBorderPrimary};
		border-radius: ${props.theme.borderRadiusNormal};
	`}
`;

const IconBox = styled.div`
	${(props) => css`
		width: 5rem;
		display: flex;
		align-items: center;
		padding: 0.8rem;
		svg,
		img {
			width: 2rem;
			height: 2rem;
			fill: ${props.theme.colorTextPrimary};
		}
	`}
`;

const InputSpacer = styled.div`
	padding-left: 1.5rem;
`;

const Input = styled.input`
	${(props) => css`
		width: 100%;
		background: none;
		border: none;
		outline: none;
		font-size: 1.6rem;
		color: ${props.theme.colorTextPrimary};
	`}
`;

const themes = {
	small: (theme) => ({
		height: '7rem',
		inputHeight: '4rem',
		errorHeight: '3rem',
	}),
	normal: (theme) => ({
		height: '8rem',
		inputHeight: '5rem',
		errorHeight: '3rem',
	}),
};

class TextInput extends Component {
	render() {
		const props = this.props;

		const theme = props.size ? themes[props.size] : themes.normal;
		let error = props.error;
		let hasIcon = !!props.icon;

		return (
			<ThemeProvider theme={theme(props.theme)}>
				<Wrapper className={'text-input'}>
					<InputMain>
						{hasIcon ? (
							<IconBox>{props.icon}</IconBox>
						) : (
							<InputSpacer />
						)}
						<Input
							name={props.name}
							value={props.value}
							onChange={props.onChange}
							onBlur={props.onBlur}
							type={props.type}
							placeholder={props.placeholder}
							min={props.min}
							max={props.max}
							onKeyPress={props.onKeyPress}
						/>
					</InputMain>
					<Error>{error ? error : null}</Error>
				</Wrapper>
			</ThemeProvider>
		);
	}
}

export default TextInput;
