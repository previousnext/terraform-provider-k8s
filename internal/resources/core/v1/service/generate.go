package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/core/v1/service/port"
)

// Generate the Service.
func Generate(d *schema.ResourceData) (corev1.Service, error) {
	var (
		name           = d.Get(FieldName).(string)
		namespace      = d.Get(FieldNamespace).(string)
		rawLabels      = d.Get(FieldLabels).(map[string]interface{})
		rawAnnotations = d.Get(FieldAnnotations).(map[string]interface{})
		rawType        = d.Get(FieldType).(string)
		rawPorts       = d.Get(FieldPort).([]interface{})
		rawSelector    = d.Get(FieldSelector).(map[string]interface{})
	)

	service := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Namespace:   namespace,
			Labels:      interfaceutils.ExpandMap(rawLabels),
			Annotations: interfaceutils.ExpandMap(rawAnnotations),
		},
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceType(rawType),
			Ports:    port.Expand(rawPorts),
			Selector: interfaceutils.ExpandMap(rawSelector),
		},
	}

	return service, nil
}
