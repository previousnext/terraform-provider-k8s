package hostaliases

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/previousnext/terraform-provider-k8s/hostaliases/hostnames"
	corev1 "k8s.io/api/core/v1"
)

// Fields returns the fields for this package.
func Fields() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Environment variables which can be set for a container",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "IP Address of the host.",
				},
				"hostnames": hostnames.Fields(),
			},
		},
	}
}

// Expand will return a structured object.
func Expand(in []interface{}) []corev1.HostAlias {
	if len(in) == 0 {
		return []corev1.HostAlias{}
	}

	aliases := make([]corev1.HostAlias, len(in))

	for key, v := range in {
		value := v.(map[string]interface{})

		if ip, ok := value["ip"]; ok && ip != "" {
			aliases[key].IP = ip.(string)
		}

		if hosts, ok := value["hostnames"]; ok {
			aliases[key].Hostnames = hostnames.Expand(hosts.([]interface{}))
		}
	}

	return aliases
}

// Flatten structured object into unstructured.
func Flatten(in []corev1.HostAlias) []interface{} {
	flattened := make([]interface{}, len(in))

	for key, value := range in {
		row := map[string]interface{}{}

		if value.IP != "" {
			row["ip"] = value.IP
		}

		if len(value.Hostnames) > 0 {
			row["hostnames"] = value.Hostnames
		}

		flattened[key] = row
	}

	return flattened
}
