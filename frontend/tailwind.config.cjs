/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#FFD700', // Yellow
          dark: '#E6C200',
          light: '#FFDF33',
        },
        secondary: {
          DEFAULT: '#000000', // Black
          light: '#333333',
        },
        background: {
          DEFAULT: '#FFFFFF', // White
          dark: '#F5F5F5',
        },
      },
    },
  },
  plugins: [],
}
