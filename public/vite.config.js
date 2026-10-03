import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { writeFileSync } from 'node:fs';

// mapserver backend, used by the dev-server proxy
const backend = process.env.MAPSERVER_URL || 'http://localhost:8080';

// keep the (empty) dist/.gitkeep, go:embed fails on a missing/empty directory
const keepDistDir = {
  name: 'keep-dist-dir',
  closeBundle(){
    writeFileSync(new URL('./dist/.gitkeep', import.meta.url), '');
  }
};

export default defineConfig({
  plugins: [vue(), keepDistDir],
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
