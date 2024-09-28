package main

import (
	_ "github.com/hiddify/hiddify-app-example-extension/hiddify_extension"

	"github.com/hiddify/hiddify-core/cmd"
)

func main() {
	cmd.StartExtension()
}
