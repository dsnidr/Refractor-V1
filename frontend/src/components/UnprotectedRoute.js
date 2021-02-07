import React from "react";
import { Route, Redirect } from "react-router";
import { connect } from "react-redux";

// UnprotectedRoute redirects to Component only if the user is not authenticated
const UnprotectedRoute = ({ component: Component, user, ...rest }) => (
    <Route
        {...rest}
        render={(props) =>
            user === null || !user.isAuthenticated ? (
                <Component {...props} />
            ) : (
                <Redirect to="/" />
            )
        }
    />
);

const mapStateToProps = (state) => ({
    user: state.user.self,
});

export default connect(mapStateToProps)(UnprotectedRoute);