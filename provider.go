package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/datasources"
	"github.com/previousnext/terraform-provider-k8s/internal/resources"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
)

// Provider returns this providers resources.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:         config.Fields(),
		DataSourcesMap: datasources.Map(),
		ResourcesMap:   resources.Map(),
		ConfigureFunc:  config.Func,
	}
}
