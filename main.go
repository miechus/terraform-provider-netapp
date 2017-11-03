package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-netapp/netapp"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: netapp.Provider})
}
