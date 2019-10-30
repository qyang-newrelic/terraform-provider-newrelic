package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/qyang-newrelic/terraform-provider-newrelic"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: newrelic.Provider})
}
