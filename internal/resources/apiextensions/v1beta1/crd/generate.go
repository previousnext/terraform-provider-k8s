package crd

import (
	"github.com/hashicorp/terraform/helper/schema"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/apiextensions/v1beta1/crd/names"
)

// Generate the ServiceAccount.
func Generate(d *schema.ResourceData) (apiextensionsv1beta1.CustomResourceDefinition, error) {
	var (
		name      = d.Get(FieldName).(string)
		rawLabels = d.Get(FieldLabels).(map[string]interface{})
		group     = d.Get(FieldGroup).(string)
		version   = d.Get(FieldVersion).(string)
		scope     = d.Get(FieldScope).(string)
		rawNames  = d.Get(FieldNames).([]interface{})
	)

	crd := apiextensionsv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: interfaceutils.ExpandMap(rawLabels),
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   group,
			Version: version,
			Scope:   apiextensionsv1beta1.ResourceScope(scope),
			Names:   names.Expand(rawNames),
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
		},
	}

	return crd, nil
}
