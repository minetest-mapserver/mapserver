package public

import "embed"

//go:embed colors/* css/* pics/* sql/* *.html *.txt js/*
//go:embed node_modules/bootstrap/dist/css/bootstrap.min.css
//go:embed node_modules/vue/dist/vue.min.js
//go:embed node_modules/vue-router/dist/vue-router.min.js
//go:embed node_modules/vue-i18n/dist/vue-i18n.min.js
//go:embed node_modules/@fortawesome/fontawesome-free/css/all.min.css
//go:embed node_modules/@fortawesome/fontawesome-free/webfonts/*
//go:embed node_modules/leaflet/dist/images/*
//go:embed node_modules/leaflet/dist/leaflet.js
//go:embed node_modules/leaflet/dist/leaflet.css
var Files embed.FS
