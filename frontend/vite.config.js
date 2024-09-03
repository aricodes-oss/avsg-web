import { defineConfig } from 'vite';
import path from 'path';

import handlebars from 'vite-plugin-handlebars';

export default defineConfig({
  root: path.resolve(__dirname, 'src'),
  build: {
    outDir: '../../embeds/dist',
    emptyOutDir: true,
  },
  server: {
    host: '0.0.0.0',
    port: 8082,
  },
  plugins: [
    handlebars({
      partialDirectory: path.resolve(__dirname, 'src', 'partials'),
    }),
  ],
});
