const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = function (app) {
    app.use(
        "/ws",
        createProxyMiddleware({
            target: "http://localhost:5000",
            ws: true,
        })
    );

    app.use(
        "/api/",
        createProxyMiddleware({
            target: "http://localhost:5000",
        })
    );
};