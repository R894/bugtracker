import { configDefaults, defineConfig } from 'vitest/config'
import react from '@vitejs/plugin-react'
import tsconfigPaths from 'vite-tsconfig-paths'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [tsconfigPaths(), react()],
  test: {
    environment: 'jsdom',
    exclude: [...configDefaults.exclude, '**/*.spec.{js,jsx,ts,tsx}'],
    setupFiles: ['vitestSetup.ts'],
    globals: true
  },
})
