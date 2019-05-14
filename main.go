package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/tokend/terraform-provider-tokend/tokend"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: tokend.Provider})
}
