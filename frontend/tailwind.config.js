/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        satpol: {
          // Warna Cokelat Praja (Menggantikan Biru/Navy)
          brown: '#3d261a',
          'brown-dark': '#1a120c',
          'brown-light': '#5a3927',
          'brown-subtle': '#f4eee9',
          // Aksen Oranye/Emas Praja Wibawa
          orange: '#F08300',
          'orange-hover': '#d97400',
          gold: '#F08300',
          // Latar Krem & Border Klasik Sebelumnya
          cream: '#f9f4e8',
          border: '#eee9db',
          dark: '#1a120c',
          black: '#120e0a'
        }
      },
      borderRadius: {
        'card': '0.75rem',
        'button': '0.5rem'
      },
      fontFamily: {
        sans: ['Plus Jakarta Sans', 'Inter', 'system-ui', 'sans-serif']
      }
    },
  },
  plugins: [],
}
