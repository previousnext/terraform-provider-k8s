package objectmeta

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta/annotations"
	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta/labels"
)

// Expand will return a structured object.
func Expand(in []interface{}) metav1.ObjectMeta {
	var meta metav1.ObjectMeta

	raw := in[0].(map[string]interface{})

	if val, ok := raw[FieldName]; ok {
		meta.Name = val.(string)
	}

	if val, ok := raw[FieldNamespace]; ok {
		meta.Namespace = val.(string)
	}

	if val, ok := raw[FieldAnnotations]; ok {
		meta.Annotations = annotations.Expand(val.(map[string]interface{}))
	}

	if val, ok := raw[FieldLabels]; ok {
		meta.Labels = labels.Expand(val.(map[string]interface{}))
	}

	return meta
}
