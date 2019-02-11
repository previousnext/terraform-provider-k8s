package objectmeta

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Flatten structured object into unstructured.
func Flatten(metadata metav1.ObjectMeta) []interface{} {
	out := make([]interface{}, 1)

	row := map[string]interface{}{}

	if metadata.Name != "" {
		row[FieldName] = metadata.Name
	}

	if metadata.Namespace != "" {
		row[FieldNamespace] = metadata.Namespace
	}

	if len(metadata.Annotations) > 0 {
		row[FieldAnnotations] = metadata.Annotations
	}

	if len(metadata.Labels) > 0 {
		row[FieldLabels] = metadata.Labels
	}

	out[0] = row

	return out
}
