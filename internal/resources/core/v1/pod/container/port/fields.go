package port

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// FieldName used to identify the type of port.
	FieldName = "name"
	// FieldContainerPort used to identify the port which is being exposed from the container.
	FieldContainerPort = "container"
	// FieldHostPort used to identify the port which is being exposed from the host.
	FieldHostPort = "host"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "Mount a volume into a path inside the container.",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldName: {
					Description: "Name of the port",
					Type:        schema.TypeString,
					Required:    true,
				},
				FieldContainerPort: {
					Description: "Port to receive requests",
					Type:        schema.TypeInt,
					Optional:    true,
				},
				FieldHostPort: {
					Description: "Port to receive requests from the host",
					Type:        schema.TypeInt,
					Optional:    true,
				},
			},
		},
	}
}
