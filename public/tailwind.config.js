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
        50: "white",
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
        50: "#CC3336",
        100: "#CC3336",
        200: "#CC3336",
        300: "#CC3336",
        400: "#CC3336",
        500: "#CC3336",
        600: "#A3292B",
        700: "#7A1F20",
        800: "#521415",
        900: "#290A0B",
        950: "#140505"
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
