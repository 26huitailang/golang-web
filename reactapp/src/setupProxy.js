// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
// eslint-disable-next-line @typescript-eslint/no-var-requires
const { createProxyMiddleware } = require('http-proxy-middleware');

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
module.exports = function (app) {
  app.use(createProxyMiddleware(
    '/apiV1',
    {
      target: 'http://localhost:8001',
      changeOrigin: true,
    },
  ));
  app.use(createProxyMiddleware(
    '/login',
    {
      target: 'http://localhost:8001',
      changeOrigin: true,
      // pathRewrite: {
      //   '/login': '/login',
      // },
    },
  ));
  app.use(createProxyMiddleware('/logout',
    {
      target: 'http://localhost:8001',
      changeOrigin: true,
    }));
};
