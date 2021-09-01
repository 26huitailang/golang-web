import { defineConfig } from 'umi';

export default defineConfig({
  // layout: {},
  nodeModulesTransform: {
    type: 'none',
  },
  routes: [
    { path: '/', component: '@/pages/index' },
    { path: '/login', component: '@/pages/Login' },
  ],
  fastRefresh: {},
  history: {
    type: 'hash'
  },
  proxy: {
    '/login': {
      target: 'http://localhost:8001',
      pathRewrite: {'': ''},
      changeOrigin: true,
    }
  }
});
