package public

import "embed"

//go:embed js/* css/* pics/* index.html
//go:embed node_modules/vue/dist/vue.global.prod.js
//go:embed node_modules/vue-router/dist/vue-router.global.prod.js
//go:embed node_modules/@fortawesome/fontawesome-free/css/all.min.css
//go:embed node_modules/@fortawesome/fontawesome-free/webfonts/*
//go:embed node_modules/leaflet/dist/leaflet.js
//go:embed node_modules/leaflet/dist/leaflet.css
//go:embed node_modules/leaflet/dist/images
//go:embed node_modules/leaflet.awesome-markers/dist/leaflet.awesome-markers.css
//go:embed colors/*
//go:embed sql/*
//go:embed *.txt
var Files embed.FS
