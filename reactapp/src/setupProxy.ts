// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
// eslint-disable-next-line @typescript-eslint/no-var-requires
const { proxy } = require('http-proxy-middleware');

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
module.exports = function (app) {
  app.use(proxy(
    '/apiV1',
    {
      target: 'http://localhost:8080',
      changeOrigin: true,
    },
  ));
  app.use(proxy(
    '/login',
    {
      target: 'http://localhost:8080',
      changeOrigin: true,
      pathRewrite: {
        '/login': '/login',
      },
    },
  ));
  app.use(proxy('/logout',
    {
      target: 'http://localhost:8080',
      changeOrigin: true,
    }));
};
