package pod

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/container"
	"github.com/previousnext/terraform-provider-k8s/core/v1/pod/volume"
)

// Expand will return a structured object.
func Expand(in []interface{}) (corev1.PodSpec, error) {
	var spec corev1.PodSpec

	raw := in[0].(map[string]interface{})

	if val, ok := raw[FieldInitContainer]; ok {
		initContainers, err := container.Expand(val.([]interface{}))
		if err != nil {
			return spec, err
		}

		spec.InitContainers = initContainers
	}

	if val, ok := raw[FieldContainer]; ok {
		containers, err := container.Expand(val.([]interface{}))
		if err != nil {
			return spec, err
		}

		spec.Containers = containers
	}

	if val, ok := raw[FieldVolume]; ok {
		volumes, err := volume.Expand(val.([]interface{}))
		if err != nil {
			return spec, err
		}

		spec.Volumes = volumes
	}

	if val, ok := raw[FieldServiceAccount]; ok {
		spec.ServiceAccountName = val.(string)
	}

	return spec, nil
}
