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
import styled, { css } from 'styled-components';

const Wrapper = styled.div`
	${(props) => css`
		width: 100%;
		display: flex;
		flex-direction: column;
		min-height: 12rem;
	`}
`;

const CustomTextArea = styled.textarea`
	${(props) => css`
		width: 100%;
		min-height: 10.25rem; // .25 to allow for 4 lines to fit before scrolling starts
		max-height: 15rem;
		background: ${props.theme.inputs.fillInBackground
			? props.theme.colorBorderPrimary
			: 'none'};
		border: 1px solid ${props.theme.colorBorderPrimary};
		border-radius: ${props.theme.borderRadiusNormal};
		outline: none;
		font-size: 1.6rem;
		padding: 1rem;
		color: ${props.theme.colorTextPrimary};
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

class TextArea extends Component {
	constructor(props) {
		super(props);

		this.inputRef = React.createRef();
	}

	focus = () => {
		this.inputRef.current.focus();
	};

	render() {
		const { placeholder, onChange, error, defaultValue } = this.props;

		return (
			<Wrapper>
				<CustomTextArea
					placeholder={placeholder}
					onChange={onChange}
					defaultValue={defaultValue}
					ref={this.inputRef}
				/>
				<Error>{error}</Error>
			</Wrapper>
		);
	}
}

export default TextArea;
