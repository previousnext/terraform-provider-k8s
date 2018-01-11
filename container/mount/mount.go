package mount

import (
	"github.com/hashicorp/terraform/helper/schema"
	corev1 "k8s.io/api/core/v1"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Mount a volume into a path inside the container.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Name of the volume to mount.",
				},
				"path": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Path to mount the volume into.",
				},
			},
		},
	}
}

// Expand will return a structured object.
func Expand(in []interface{}) ([]corev1.VolumeMount, error) {
	if len(in) == 0 {
		return []corev1.VolumeMount{}, nil
	}

	mounts := make([]corev1.VolumeMount, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value["name"]; ok {
			mounts[key].Name = name.(string)
		}

		if path, ok := value["path"]; ok {
			mounts[key].MountPath = path.(string)
		}
	}

	return mounts, nil
}

// Flatten structured object into unstructured.
func Flatten(in []corev1.VolumeMount) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row["name"] = value.Name
		}

		if value.MountPath != "" {
			row["value"] = value.MountPath
		}

		flattened[key] = row
	}

	return flattened
}
