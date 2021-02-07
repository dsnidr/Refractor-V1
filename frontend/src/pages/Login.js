import React, { Component } from 'react';
import { connect } from 'react-redux';
import respondTo from '../mixins/respondTo';
import styled, { css } from 'styled-components';
import Wrapper from '../components/Wrapper';

const FormWrapper = styled.form`
	${(props) => css`
		border: 1px solid ${props.theme.colorBorderSecondary};
		border-radius: ${props.theme.borderRadiusNormal};
		background: ${props.theme.colorBackground};
		width: 100%;
		min-height: 100vh;
		display: flex;
		flex-direction: column;
		align-items: center;

		${props.theme.boxShadowPrimary}

		${respondTo.medium`
      		width: 50rem;
      		min-height: 40rem;
    	`}
		
		> * {
			width: 80%;
		}
	`}
`;

const HeadingBox = styled.div`
	${(props) => css`
		height: 10rem;
		font-size: 3.2rem;
		display: flex;
		justify-content: center;
		align-items: center;
		font-weight: 500;
		color: ${props.theme.colorTextPrimary};
	`}
`;

class Login extends Component {
	render() {
		return (
			<Wrapper>
				<FormWrapper>
					<HeadingBox>LOG IN</HeadingBox>
				</FormWrapper>
			</Wrapper>
		);
	}
}

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(Login);
