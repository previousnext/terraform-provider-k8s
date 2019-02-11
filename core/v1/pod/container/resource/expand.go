package resource

import (
	corev1 "k8s.io/api/core/v1"
	k8sresource "k8s.io/apimachinery/pkg/api/resource"
)

// Expand will return a structured object.
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
