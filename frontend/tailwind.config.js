/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        satpol: {
          // Warna Praja Wibawa & Tamalate Sovereign Fusion
          navy: '#0b172a',
          'navy-dark': '#060d19',
          'navy-surface': '#0f223d',
          'navy-light': '#1e3a5f',
          // Aksen Oranye / Emas Praja Wibawa
          orange: '#F08300',
          'orange-hover': '#d97400',
          gold: '#F08300',
          'gold-subtle': '#fff7ed',
          // Tamalate Sovereign Blue & Porcelain Elements
          azure: '#1D4ED8',
          'azure-vivid': '#246BFD',
          'azure-tint': '#EBF2FF',
          canvas: '#F4F8FF',
          surface: '#FFFFFF',
          porcelain: '#F8FAFC',
          // Latar Netral & Abu Bersih (Menggantikan Soft Pink & Krem Monoton)
          slate: '#f1f5f9',
          'slate-border': '#e2e8f0',
          'slate-ink': '#0f172a',
          'slate-muted': '#64748b',
          // Kompatibilitas Legacy Praja
          brown: '#3d261a',
          'brown-dark': '#1a120c',
          'brown-light': '#5a3927',
          'brown-subtle': '#f4eee9',
          cream: '#f8fafc',
          border: '#e2e8f0',
          dark: '#0b172a',
          black: '#070d18'
        }
      },
      borderRadius: {
        'card': '1rem',
        'squircle': '1.125rem',
        'squircle-lg': '1.5rem',
        'button': '0.625rem'
      },
      boxShadow: {
        'glass': '0 8px 30px -4px rgba(15, 23, 42, 0.08), 0 2px 8px rgba(15, 23, 42, 0.04)',
        'glass-hover': '0 16px 40px -8px rgba(15, 23, 42, 0.12), 0 4px 12px rgba(15, 23, 42, 0.06)',
        'glass-gold': '0 10px 30px -4px rgba(240, 131, 0, 0.3)',
        'glass-azure': '0 10px 30px -4px rgba(36, 107, 253, 0.28)'
      },
      fontFamily: {
        sans: ['Plus Jakarta Sans', 'Inter', 'system-ui', 'sans-serif'],
        mono: ['JetBrains Mono', 'monospace']
      }
    },
  },
  plugins: [],
}
