import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  resolve:{
    alias: {
      crypto: 'crypto-browserify',
      buffer: "buffer",
      stream: "stream-browserify"

    }
  },
  define: {
    global: "window",
    process: { env: {}, version: "v1.0.0" },
  },
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:5000",
        secure: false,
      },
    },
  },
});
