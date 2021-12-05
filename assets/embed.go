package assets

import "embed"

//go:embed "templates/*"
var Templates embed.FS

//go:embed "public/*"
var Public embed.FS
