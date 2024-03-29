package apiservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
)

// Generate the APIService.
func Generate(d *schema.ResourceData) (apiregistrationv1.APIService, error) {
	var (
		name             = d.Get(FieldName).(string)
		rawLabels        = d.Get(FieldLabels).(map[string]interface{})
		group            = d.Get(FieldGroup).(string)
		version          = d.Get(FieldVersion).(string)
		serviceName      = d.Get(FieldServiceName).(string)
		serviceNamespace = d.Get(FieldServiceNamespace).(string)
		insecure         = d.Get(FieldInsecureSkipTLSVerify).(bool)
		groupPriority    = d.Get(FieldGroupPriorityMinimum).(int)
		versionPriority  = d.Get(FieldVersionPriority).(int)
	)

	crd := apiregistrationv1.APIService{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: interfaceutils.ExpandMap(rawLabels),
		},
		Spec: apiregistrationv1.APIServiceSpec{
			Service: &apiregistrationv1.ServiceReference{
				Name:      serviceName,
				Namespace: serviceNamespace,
			},
			Group:                 group,
			Version:               version,
			InsecureSkipTLSVerify: insecure,
			GroupPriorityMinimum:  int32(groupPriority),
			VersionPriority:       int32(versionPriority),
		},
	}

	return crd, nil
}
