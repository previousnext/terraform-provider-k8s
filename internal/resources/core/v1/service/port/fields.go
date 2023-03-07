package port

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// FieldName is a field identifier.
	FieldName = "name"
	// FieldPort is a field identifier.
	FieldPort = "port"
	// FieldTargetPort is a field identifier.
	FieldTargetPort = "target_port"
	// FieldNodePort is a field identifier.
	FieldNodePort = "node_port"
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
				FieldPort: {
					Description: "Port to receive requests",
					Type:        schema.TypeInt,
					Required:    true,
				},
				FieldTargetPort: {
					Description: "Port to receive requests",
					Type:        schema.TypeString,
					Required:    true,
				},
				FieldNodePort: {
					Description: "Port to receive requests",
					Type:        schema.TypeInt,
					Optional:    true,
					Computed:    true,
				},
			},
		},
	}
}
