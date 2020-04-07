package main

import (
	c "github.com/StuartLox/terraform-provider-kafka-connect/connect"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: c.Provider})
}
