package hostalias

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/pod/container/hostalias/hostname"
)

const (
	// FieldIP is used to identify the IP field.
	FieldIP = "ip"
	// FieldHostname is used to identify the hostname field.
	FieldHostname = "hostname"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "Environment variables which can be set for a container",
		Type:        schema.TypeList,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldIP: {
					Type:        schema.TypeString,
					Required:    true,
					Description: "IP Address of the host.",
				},
				FieldHostname: hostname.Fields(),
			},
		},
	}
}
