/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./cmd/server/templates/**/*.{html,js}",
    "./cmd/server/handlers/**/*.{html,js}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};
