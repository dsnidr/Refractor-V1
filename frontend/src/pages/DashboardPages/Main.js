import React, { Component } from 'react';
import { connect } from 'react-redux';
import Heading from '../../components/Heading';

class Main extends Component {
	render() {
		return (
			<>
				<div>
					<Heading headingStyle="title">
						Hello, {this.props.user.username}.
					</Heading>
				</div>
			</>
		);
	}
}

const mapStateToProps = (state) => ({
	user: state.user.self,
});

export default connect(mapStateToProps)(Main);
