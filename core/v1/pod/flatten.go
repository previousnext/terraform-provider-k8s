package pod

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/volume"
)

// Flatten structured object into unstructured.
func Flatten(spec corev1.PodSpec) []interface{} {
	out := make([]interface{}, 1)

	row := map[string]interface{}{}

	if len(spec.InitContainers) > 0 {
		row[FieldInitContainer] = container.Flatten(spec.InitContainers)
	}

	if len(spec.Containers) > 0 {
		row[FieldContainer] = container.Flatten(spec.Containers)
	}

	if len(spec.Volumes) > 0 {
		row[FieldVolume] = volume.Flatten(spec.Volumes)
	}

	if spec.ServiceAccountName != "" {
		row[FieldServiceAccount] = spec.ServiceAccountName
	}

	out[0] = row

	return out
}
