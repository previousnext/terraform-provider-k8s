package subject

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// FieldSubject identifier for the subject field.
	FieldSubject = "subject"
	// FieldKind identifier for the kind field.
	FieldKind = "kind"
	// FieldName identifier for the name field.
	FieldName = "name"
	// FieldAPIGroup identifier for the API group field.
	FieldAPIGroup = "api_group"
	// FieldNamespace identifier for the namespace field.
	FieldNamespace = "namespace"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Subjects which will receive this RoleBinding.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldKind: {
					Type:        schema.TypeString,
					Description: "Kind eg. User, Group, ServiceAccount etc",
					Required:    true,
				},
				FieldName: {
					Type:        schema.TypeString,
					Description: "Name of the allowed",
					Required:    true,
				},
				FieldAPIGroup: {
					Type:        schema.TypeString,
					Description: "API group of the allowed",
					Optional:    true,
				},
				FieldNamespace: {
					Type:        schema.TypeString,
					Description: "Namespace of the allowed",
					Optional:    true,
				},
			},
		},
	}
}
