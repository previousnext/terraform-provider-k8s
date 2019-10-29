package apiservice

import (
	"github.com/hashicorp/terraform/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiregistrationv1beta1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
)

// Generate the APIService.
func Generate(d *schema.ResourceData) (apiregistrationv1beta1.APIService, error) {
	var (
		name             = d.Get(FieldName).(string)
		rawLabels        = d.Get(FieldLabels).(map[string]interface{})
		group            = d.Get(FieldGroup).(string)
		version          = d.Get(FieldVersion).(string)
		serviceName      = d.Get(FieldServiceName).(string)
		serviceNamespace = d.Get(FieldServiceNamespace).(string)
		insecure         = d.Get(FieldInsecureSkipTLSVerify).(bool)
		groupPriority    = d.Get(FieldGroupPriorityMinimum).(int32)
		versionPriority  = d.Get(FieldVersionPriority).(int32)
	)

	crd := apiregistrationv1beta1.APIService{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: interfaceutils.ExpandMap(rawLabels),
		},
		Spec: apiregistrationv1beta1.APIServiceSpec{
			Service: &apiregistrationv1beta1.ServiceReference{
				Name:      serviceName,
				Namespace: serviceNamespace,
			},
			Group:                 group,
			Version:               version,
			InsecureSkipTLSVerify: insecure,
			GroupPriorityMinimum:  groupPriority,
			VersionPriority:       versionPriority,
		},
	}

	return crd, nil
}
