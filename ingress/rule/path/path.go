package path

import (
	"github.com/hashicorp/terraform/helper/schema"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Paths to apply to an Ingress Rule",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"path": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Path to push Ingress traffic to.",
				},
				"service_name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Service to route the Ingress to.",
				},

				"port": {
					Type:        schema.TypeInt,
					Required:    true,
					Description: "Port for Ingress traffic.",
				},
			},
		},
	}
}

func Expand(in []interface{}) []extensionsv1beta1.HTTPIngressPath {
	if len(in) == 0 {
		return []extensionsv1beta1.HTTPIngressPath{}
	}

	paths := make([]extensionsv1beta1.HTTPIngressPath, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if name, ok := value["path"]; ok && name != "" {
			paths[key].Path = name.(string)
		}

		if serviceName, ok := value["service_name"]; ok {
			paths[key].Backend.ServiceName = serviceName.(string)
		}

		if port, ok := value["port"]; ok {
			paths[key].Backend.ServicePort = intstr.FromInt(port.(int))
		}
	}

	return paths
}

func Flatten(in []extensionsv1beta1.HTTPIngressPath) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.Path != "" {
			row["path"] = value.Path
		}

		if value.Backend.ServiceName != "" {
			row["service_name"] = value.Backend.ServiceName
		}

		if value.Backend.ServicePort.String() != "" {
			row["port"] = value.Backend.ServicePort
		}
		flattened[key] = row
	}

	return flattened
}
