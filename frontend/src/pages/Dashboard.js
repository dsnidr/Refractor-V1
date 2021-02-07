import React, { Component } from 'react';
import { connect } from 'react-redux';

class Dashboard extends Component {
	constructor(props) {
		super(props);

		this.state = {};
	}

	render() {
		return <h1>Dashboard</h1>;
	}
}

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(Dashboard);
