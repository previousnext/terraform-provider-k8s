package volume

import (
	"github.com/hashicorp/terraform/helper/schema"
	corev1 "k8s.io/api/core/v1"
)

func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "List of Volumes used to mount into containers.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Name of this Volume.",
				},
				"hostpath": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Path on the host to mount into the container.",
				},
				"pvc": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Name of the PersistentVolumeClaim to mount.",
				},
				"configmap": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Name of the ConfigMap to mount.",
				},
			},
		},
	}
}

func Expand(in []interface{}) ([]corev1.Volume, error) {
	if len(in) == 0 {
		return []corev1.Volume{}, nil
	}

	volumes := make([]corev1.Volume, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value["name"]; ok {
			volumes[key].Name = name.(string)
		}

		if hostpath, ok := value["hostpath"]; ok && hostpath != "" {
			volumes[key].HostPath = &corev1.HostPathVolumeSource{
				Path: hostpath.(string),
			}
		}

		if pvc, ok := value["pvc"]; ok && pvc != "" {
			volumes[key].PersistentVolumeClaim = &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: pvc.(string),
			}
		}

		if cfg, ok := value["configmap"]; ok && cfg != "" {
			volumes[key].ConfigMap = &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: cfg.(string),
				},
			}
		}
	}

	return volumes, nil
}

func Flatten(in []corev1.Volume) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row["name"] = value.Name
		}

		if value.HostPath != nil && value.HostPath.Path != "" {
			row["hostpath"] = value.HostPath.Path
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
