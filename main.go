package main

import (
	_ "github.com/hiddify/hiddify-app-demo-extension/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
