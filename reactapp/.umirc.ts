import { defineConfig } from 'umi';

export default defineConfig({
  // layout: {},
  nodeModulesTransform: {
    type: 'none',
  },
  routes: [
    { path: '/', component: '@/pages/index' },
    { path: '/Login', component: '@/pages/Login' },
  ],
  fastRefresh: {},
  history: {
    type: 'hash'
  }
});
