const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.html",
  ],
  theme: {
    colors: {
      transparent: 'transparent',
      current: 'currentColor',
      white: colors.white,
      black: colors.black,
      yellow: colors.yellow,
      gray: {
        50: "#EEEFF1",
        100: "#DEDFE3",
        200: "#BABCC5",
        300: "#999CA8",
        400: "#75798A",
        500: "#585B68",
        600: "#464853",
        700: "#363840",
        800: "#232429",
        900: "#222426"
      },
      rose: colors.rose,
      primary: {
        50: '#cc3336',
        100: '#cc3336',
        200: '#cc3336',
        300: '#cc3336',
        400: '#cc3336',
        500: '#cc3336',
        600: '#cc3336',
        700: '#cc3336',
        800: '#cc3336',
        900: '#cc3336',
        950: '#cc3336',
      },
      slate: colors.slate
    },
    extend: {
      colors: {
        slate: {
          600: '#73777d',
          700: '#222426',
        }
      },
      borderWidth: {
        '1': '1px',
      },
      boxShadow: {
        'DEFAULT': 'none',
      },
      borderRadius: {
        'lg': '0',
      }
    },
  },
  plugins: [require("@tailwindcss/typography"),require('@tailwindcss/forms')],
  darkMode: 'class',
}
