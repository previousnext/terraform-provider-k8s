package path

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldPath identifier for path.
	FieldPath = "path"
	// FieldService identifier for service.
	FieldService = "service"
	// FieldPort identifier for port.
	FieldPort = "port"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Paths to apply to an Ingress Rule",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldPath: {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Path to push Ingress traffic to",
				},
				FieldService: {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Service to route the Ingress to",
				},
				FieldPort: {
					Type:        schema.TypeInt,
					Required:    true,
					Description: "Port for Ingress traffic",
				},
			},
		},
	}
}
