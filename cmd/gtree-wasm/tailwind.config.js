/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./css/*.css", "./js/*.js"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'), // require: npm install -D @tailwindcss/forms
    require('@tailwindcss/aspect-ratio'),
  ],
}
