const colors = require('tailwindcss/colors');

module.exports = {
  content: ['./www/*.html'],
  theme: {
    extend: {
      colors: {
        blue: colors.blue,
        green: colors.emerald,
        short: "rgb(69, 159, 246)",
        standard: "rgb(139, 93, 246)",
        extended: "rgb(246, 168, 35)",
        tutorial: "rgb(115, 128, 140)",
        weewoosiren: "rgb(234 88 12)"
      },
      height: {
        stretch: 'stretch',
      },
      spacing: {
        '05rem': '0.5rem',
        '1rem': '1rem',
        '2rem': '2rem',
        '3rem': '3rem',
      },
      backgroundColor: {
        'white': '#ffffff',
        'dark': '#393b40',
        'darker': '#242629',
        'darkerer': "#1c1d1f",
        'darkest': '#151617',

        'dark_tab': "#323633",
        'darker_tab': "#262927",
        'dark_tab_hover': "#3a3d3a",
        'darker_tab_hover': "#2e312e",
      },
      borderColor: {
        "dark_tab": "#111211"
      }
    },
  },
  plugins: [require('@tailwindcss/forms')],
};
