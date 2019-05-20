package rule

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role/rule/apigroups"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role/rule/resourcenames"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role/rule/resources"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/rbac/v1/role/rule/verbs"
)

const (
	// FieldAPIGroups is used to identify the api groups field.
	FieldAPIGroups = "api_groups"
	// FieldResources is used to identify the resources field.
	FieldResources = "resources"
	// FieldResourceNames is used to identify the resourceNames field.
	FieldResourceNames = "resource_names"
	// FieldVerbs is used to identify the verbs field.
	FieldVerbs = "verbs"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Rules to apply to a Role",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldAPIGroups:     apigroups.Fields(),
				FieldResources:     resources.Fields(),
				FieldResourceNames: resourcenames.Fields(),
				FieldVerbs:         verbs.Fields(),
			},
		},
	}
}
