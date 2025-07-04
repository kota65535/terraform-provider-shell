package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/kota65535/terraform-provider-shell/shell"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: shell.Provider})
}
