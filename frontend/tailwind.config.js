/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        satpol: {
          orange: '#F08300',
          'orange-hover': '#d97400',
          dark: '#2b2f3f',
          black: '#101116',
          gray: '#6f7775',
          cream: '#f9f4e8',
          border: '#eee9db',
          navy: '#1e293b'
        }
      },
      fontFamily: {
        sans: ['Manrope', 'Nunito', 'Arial', 'sans-serif']
      }
    },
  },
  plugins: [],
}
