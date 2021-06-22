package ingress

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldIP is a field identifier.
	FieldIP = "ip"
	// FieldHostname is a field identifier.
	FieldHostname = "hostname"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldIP: {
					Type:     schema.TypeString,
					Computed: true,
				},
				FieldHostname: {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}
