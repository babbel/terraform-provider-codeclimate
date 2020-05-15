package main

import (
	"github.com/babbel/terraform-provider-codeclimate/codeclimate"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return codeclimate.Provider()
		}})
}
