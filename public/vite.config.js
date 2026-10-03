import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

// mapserver backend, used by the dev-server proxy
const backend = process.env.MAPSERVER_URL || 'http://localhost:8080';

export default defineConfig({
  plugins: [vue()],
  // relative paths: the mapserver may be served below a sub-path (reverse proxy)
  base: './',
  publicDir: 'static',
  build: {
    outDir: 'dist',
    emptyOutDir: true,
    sourcemap: true
  },
  server: {
    proxy: {
      '/api': {
        target: backend,
        ws: true
      }
    }
  }
});
