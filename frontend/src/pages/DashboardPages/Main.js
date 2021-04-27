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
