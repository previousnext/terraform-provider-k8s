package pod

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/container"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/pod/volume"
)

// Flatten structured object into unstructured.
func Flatten(template corev1.PodTemplateSpec) []interface{} {
	out := make([]interface{}, 1)

	row := map[string]interface{}{}

	if len(template.ObjectMeta.Labels) > 0 {
		row[FieldLabels] = template.ObjectMeta.Labels
	}

	if len(template.ObjectMeta.Annotations) > 0 {
		row[FieldAnnotations] = template.ObjectMeta.Annotations
	}

	if len(template.Spec.NodeSelector) > 0 {
		row[FieldNodeSelector] = template.Spec.NodeSelector
	}

	if len(template.Spec.InitContainers) > 0 {
		row[FieldInitContainer] = container.Flatten(template.Spec.InitContainers)
	}

	if len(template.Spec.Containers) > 0 {
		row[FieldContainer] = container.Flatten(template.Spec.Containers)
	}

	if len(template.Spec.Volumes) > 0 {
		row[FieldVolume] = volume.Flatten(template.Spec.Volumes)
	}

	if template.Spec.ServiceAccountName != "" {
		row[FieldServiceAccount] = template.Spec.ServiceAccountName
	}

	if len(template.Spec.ImagePullSecrets) > 0 {
		row[FieldPullSecret] = template.Spec.ImagePullSecrets[0].Name
	}

	if template.Spec.HostPID {
		row[FieldHostPID] = template.Spec.HostPID
	}

	if template.Spec.PriorityClassName != "" {
		row[FieldPriorityClassName] = template.Spec.PriorityClassName
	}

	out[0] = row

	return out
}
