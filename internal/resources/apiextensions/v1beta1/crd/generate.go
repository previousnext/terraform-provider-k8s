package crd

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/apiextensions/v1beta1/crd/names"
)

// Generate the ServiceAccount.
func Generate(d *schema.ResourceData) (apiextensionsv1.CustomResourceDefinition, error) {
	var (
		name      = d.Get(FieldName).(string)
		rawLabels = d.Get(FieldLabels).(map[string]interface{})
		group     = d.Get(FieldGroup).(string)
		version   = d.Get(FieldVersion).(string)
		scope     = d.Get(FieldScope).(string)
		rawNames  = d.Get(FieldNames).([]interface{})
	)

	crd := apiextensionsv1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: interfaceutils.ExpandMap(rawLabels),
		},
		Spec: apiextensionsv1.CustomResourceDefinitionSpec{
			Group: group,
			Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
				{
					Name:    version,
					Served:  true,
					Storage: true,
					// @todo, Determine better approach for schema.
                                        Schema: &apiextensionsv1.CustomResourceValidation{},
					Subresources: &apiextensionsv1.CustomResourceSubresources{
						Status: &apiextensionsv1.CustomResourceSubresourceStatus{},
					},
				},
			},
			Scope: apiextensionsv1.ResourceScope(scope),
			Names: names.Expand(rawNames),
		},
	}

	return crd, nil
}
