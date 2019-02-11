package rule

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/extensions/v1beta1/ingress/rule/path"
)

const (
	// FieldRule identifier for rule.
	FieldRule = "rule"
	// FieldHost identifier for host.
	FieldHost = "host"
	// FieldPath identifier for path.
	FieldPath = "path"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Rules to apply to an Ingress",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldHost: {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Hostname to respond to",
				},
				FieldPath: path.Fields(),
			},
		},
	}
}
