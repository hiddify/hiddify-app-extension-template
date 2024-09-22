package hiddify_extension

import (
	"embed"
)

// Embed all translations in the resources directory
//
//go:embed translations/*.json
var Resources embed.FS
