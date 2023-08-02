/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/layouts/*.html",
    "./views/pages/*.html",
    "./views/components/**/*.html"
  ],
  theme: {
    fontFamily: {
      sans: ['"Inter"', 'sans-serif'],
    },
    extend: {
      colors: {
        derisk: {
          primary: '#DB2777',
          100: '#FFFFFF',
          200: '#9DA2A8',
          300: '#5F646C',
          500: '#33363D',
          600: '#23252C',
          800: '#18191E',
          900: '#0E0E11',
          950: '#000000',
        },
      },
    },
  },
  plugins: [],
}
