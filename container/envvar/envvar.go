package envvar

import (
	"github.com/hashicorp/terraform/helper/schema"
	corev1 "k8s.io/api/core/v1"
)

func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Environment variables which can be set for a container",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Name of environment variable.",
				},
				"value": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Environment variable value.",
				},
				"field_path": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Field path to use for environment variable.",
				},
			},
		},
	}
}

func Expand(in []interface{}) []corev1.EnvVar {
	if len(in) == 0 {
		return []corev1.EnvVar{}
	}

	mounts := make([]corev1.EnvVar, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value["name"]; ok && name != "" {
			mounts[key].Name = name.(string)
		}

		if val, ok := value["value"]; ok && val != "" {
			mounts[key].Value = val.(string)
		}

		if path, ok := value["field_path"]; ok && path != "" {
			mounts[key].ValueFrom = &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: path.(string),
				},
			}
		}
	}

	return mounts
}

func Flatten(in []corev1.EnvVar) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Name != "" {
			row["name"] = value.Name
		}

		if value.Value != "" {
			row["value"] = value.Value
		}

		if value.ValueFrom != nil && value.ValueFrom.FieldRef != nil && value.ValueFrom.FieldRef.FieldPath != "" {
			row["field_path"] = value.ValueFrom.FieldRef.FieldPath
		}

		flattened[key] = row
	}

	return flattened
}
