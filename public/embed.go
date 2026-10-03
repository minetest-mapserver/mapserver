package public

import (
	"embed"
	"io/fs"
)

// frontend build output, created with `npm run build`
//
//go:embed all:dist
var dist embed.FS

// Files contains the built frontend (index.html, assets, pics)
var Files, _ = fs.Sub(dist, "dist")
