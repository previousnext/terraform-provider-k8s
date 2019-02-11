package volume

import corev1 "k8s.io/api/core/v1"

// Flatten structured object into unstructured.
func Flatten(in []corev1.Volume) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row["name"] = value.Name
		}

		if value.PersistentVolumeClaim != nil && value.PersistentVolumeClaim.ClaimName != "" {
			row["pvc"] = value.PersistentVolumeClaim.ClaimName
		}

		if value.ConfigMap != nil && value.ConfigMap.LocalObjectReference.Name != "" {
			row["configmap"] = value.ConfigMap.LocalObjectReference.Name
		}

		flattened[key] = row
	}

	return flattened
}
