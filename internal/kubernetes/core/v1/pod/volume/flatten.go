package volume

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/volume/nfs"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/volume/hostpath"
)

// Flatten structured object into unstructured.
func Flatten(in []corev1.Volume) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row[FieldName] = value.Name
		}

		if value.PersistentVolumeClaim != nil && value.PersistentVolumeClaim.ClaimName != "" {
			row[FieldPVC] = value.PersistentVolumeClaim.ClaimName
		}

		if value.ConfigMap != nil && value.ConfigMap.LocalObjectReference.Name != "" {
			row[FieldConfigMap] = value.ConfigMap.LocalObjectReference.Name
		}

		if value.EmptyDir != nil {
			row[FieldEmptyDir] = value.EmptyDir.Medium
		}

		if value.NFS != nil {
			row[FieldNFS] = nfs.Flatten(value.NFS)
		}

		if value.HostPath != nil {
			row[FieldHostPath] = hostpath.Flatten(value.HostPath)
		}

		flattened[key] = row
	}

	return flattened
}
