import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    svelte(),
    {
      name: 'custom-server-logger',
      configureServer() {
        // console.log('\n🚀 Visit your host screen at: http://localhost:5173/#/host\n');
        // console.log('\n🚀 Visit Player at: http://localhost:5173/#/host\n');
      }
    }
  ],
  server: {
    host: true,        // Binds to your IP (0.0.0.0)
    port: 5173,        // You can change this if needed
    strictPort: true,  // Errors if port is already used
  }
})
