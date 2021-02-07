import React, {Component} from "react";
import { Provider } from "react-redux";

import { store, history } from "./redux/store";
import {Router, Switch} from "react-router";

class App extends Component {
    constructor(props) {
        super(props);

        this.state = {
            tokenChecked: false,
        }
    }

    render() {
        return (
            <Provider store={store}>
                <Router history={history}>
                    <Switch>
                        
                    </Switch>
                </Router>
            </Provider>
        )
    }
}

export default App;
