import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

function pathResolve(dir) {
  return path.join(__dirname, dir)
}

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': pathResolve('src')
    }
  },
  server: {
    cors: true,
    open: true,
    port: 3333,
    proxy: {
      '/api': {
        target: 'http://localhost:8199',
        changeOrigin: true,
        secure: false,
      },
      '/image': {
        target: 'http://localhost:8199',
        changeOrigin: true,
        secure: false,
      }
    }
  }
})
