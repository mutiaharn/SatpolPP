/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        satpol: {
          // 4-Color Master Palette (Apple HIG Sovereign Praja)
          // 1. Praja Gold / Amber Accent (#F08300)
          gold: '#F08300',
          orange: '#F08300',
          'orange-hover': '#d97400',
          'gold-subtle': 'rgba(240, 131, 0, 0.08)',
          // 2. Obsidian Navy Dark Anchor (#0B132A)
          navy: '#0b132a',
          'navy-dark': '#070e1c',
          'navy-surface': '#0f172a',
          'navy-light': '#1e293b',
          dark: '#0b132a',
          black: '#070e1c',
          // 3. Pristine White / Canvas Surface (#FFFFFF & #F8FAFC)
          surface: '#FFFFFF',
          canvas: '#F8FAFC',
          porcelain: '#F8FAFC',
          cream: '#F8FAFC',
          // 4. Neutral Steel Slate (#64748B & #E2E8F0)
          slate: '#f1f5f9',
          'slate-border': '#e2e8f0',
          'slate-ink': '#0f172a',
          'slate-muted': '#64748b',
          border: '#e2e8f0',
          // Legacy aliases safely mapped to 4-color palette
          azure: '#F08300',
          'azure-vivid': '#F08300',
          'azure-tint': 'rgba(240, 131, 0, 0.08)',
          brown: '#0b132a',
          'brown-dark': '#070e1c',
          'brown-light': '#1e293b',
          'brown-subtle': '#f8fafc'
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
        'glass-azure': '0 10px 30px -4px rgba(240, 131, 0, 0.3)'
      },
      fontFamily: {
        sans: ['Plus Jakarta Sans', 'Inter', 'system-ui', 'sans-serif'],
        mono: ['JetBrains Mono', 'monospace']
      }
    },
  },
  plugins: [],
}
