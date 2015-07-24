package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/panther-it/terraform-provider-vsphere"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vsphere.Provider,
	})
}
