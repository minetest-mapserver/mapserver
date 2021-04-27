package public

import "embed"

//go:embed colors/*
//go:embed css/*
//go:embed pics/*
//go:embed sql/*
//go:embed webfonts/*
//go:embed *.html
//go:embed *.txt
//go:embed js/*
var Files embed.FS
