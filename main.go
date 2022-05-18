package main

import (
	"github.com/afresh-technologies/terraform-provider-onepassword/onepassword"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: onepassword.Provider,
	})
}
