package pod

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/container"
	"github.com/previousnext/terraform-provider-k8s/internal/kubernetes/core/v1/pod/volume"
)

// Expand will return a structured object.
func Expand(in []interface{}) (corev1.PodTemplateSpec, error) {
	var template corev1.PodTemplateSpec

	if len(in) == 0 {
		return template, nil
	}

	raw := in[0].(map[string]interface{})

	if val, ok := raw[FieldLabels]; ok {
		template.ObjectMeta.Labels = interfaceutils.ExpandMap(val.(map[string]interface{}))

		// Also apply an anti affinity rule to this pod based on the labels.
		template.Spec.Affinity = &corev1.Affinity{
			PodAffinity: &corev1.PodAffinity{
				PreferredDuringSchedulingIgnoredDuringExecution: []corev1.WeightedPodAffinityTerm{
					{
						Weight: 100,
						PodAffinityTerm: corev1.PodAffinityTerm{
							LabelSelector: &metav1.LabelSelector{
								MatchLabels: template.ObjectMeta.Labels,
							},
							TopologyKey: "kubernetes.io/hostname",
						},
					},
				},
			},
		}
	}

	if val, ok := raw[FieldAnnotations]; ok {
		template.ObjectMeta.Annotations = interfaceutils.ExpandMap(val.(map[string]interface{}))
	}

	if val, ok := raw[FieldInitContainer]; ok {
		initContainers, err := container.Expand(val.([]interface{}))
		if err != nil {
			return template, err
		}

		template.Spec.InitContainers = initContainers
	}

	if val, ok := raw[FieldContainer]; ok {
		containers, err := container.Expand(val.([]interface{}))
		if err != nil {
			return template, err
		}

		template.Spec.Containers = containers
	}

	if val, ok := raw[FieldVolume]; ok {
		volumes, err := volume.Expand(val.([]interface{}))
		if err != nil {
			return template, err
		}

		template.Spec.Volumes = volumes
	}

	if val, ok := raw[FieldServiceAccount]; ok {
		template.Spec.ServiceAccountName = val.(string)
	}

	if val, ok := raw[FieldPullSecret]; ok {
		template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{
			{
				Name: val.(string),
			},
		}
	}

	return template, nil
}
