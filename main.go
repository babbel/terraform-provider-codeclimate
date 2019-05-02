package main

import (
	"github.com/babbel/terraform-provider-codeclimate/codeclimate"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: codeclimate.Provider})
}
