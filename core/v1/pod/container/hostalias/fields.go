package hostalias

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/pod/container/hostalias/hostname"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Environment variables which can be set for a container",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "IP Address of the host.",
				},
				"hostname": hostname.Fields(),
			},
		},
	}
}
