package volume

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "List of Volumes used to mount into containers.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Name of this Volume.",
				},
				"pvc": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Name of the PersistentVolumeClaim to mount.",
				},
				"configmap": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Name of the ConfigMap to mount.",
				},
			},
		},
	}
}
