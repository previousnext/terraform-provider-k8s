package objectmeta

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta/annotations"
	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta/labels"
)

const (
	// FieldObjectMeta identifier for the objectmeta field.
	FieldObjectMeta = "objectmeta"
	// FieldName identifier for the objectmeta Name field.
	FieldName = "name"
	// FieldNamespace identifier for the objectmeta Namespace field.
	FieldNamespace = "namespace"
	// FieldAnnotations identifier for the objectmeta Annotations field.
	FieldAnnotations = "annotations"
	// FieldLabels identifier for the objectmeta Labels field.
	FieldLabels = "labels"
)

// Fields which define an Object.
func Fields() *schema.Schema {
	return &schema.Schema{
		Description: "ObjectMeta relating to the object.",
		Type:        schema.TypeList,
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				FieldName: {
					Type:        schema.TypeString,
					Description: "Name of the object.",
					Required:    true,
				},
				FieldNamespace: {
					Type:        schema.TypeString,
					Description: "Namespace which the object will reside.",
					Optional:    true,
				},
				FieldAnnotations: annotations.Fields(),
				FieldLabels:      labels.Fields(),
			},
		},
	}
}
