package resource

import (
	"github.com/hashicorp/terraform/helper/schema"
	corev1 "k8s.io/api/core/v1"
	k8sresource "k8s.io/apimachinery/pkg/api/resource"
)

func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "CPU and Memory constraints to apply.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"cpu": {
					Type:        schema.TypeString,
					Description: "CPU resource constraint.",
					Optional:    true,
				},
				"memory": {
					Type:        schema.TypeString,
					Description: "Memory resource constraint.",
					Optional:    true,
				},
			},
		},
	}
}

func Expand(in []interface{}) (corev1.ResourceList, error) {
	if len(in) == 0 {
		return corev1.ResourceList{}, nil
	}

	list := corev1.ResourceList{}

	for _, v := range in {
		value := v.(map[string]interface{})

		if cpu, ok := value["cpu"]; ok {
			quantity, err := k8sresource.ParseQuantity(cpu.(string))
			if err != nil {
				return list, err
			}

			list["cpu"] = quantity
		}

		if memory, ok := value["memory"]; ok {
			quantity, err := k8sresource.ParseQuantity(memory.(string))
			if err != nil {
				return list, err
			}

			list["memory"] = quantity
		}
	}

	return list, nil
}

func Flatten(in corev1.ResourceList) map[string]interface{} {
	flattened := map[string]interface{}{}

	if cpu, ok := in["cpu"]; ok {
		flattened["cpu"] = cpu.String()
	}

	if memory, ok := in["memory"]; ok {
		flattened["memory"] = memory.String()
	}

	return flattened
}
