package crd

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/previousnext/terraform-provider-k8s/internal/interfaceutils"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/apiextensions/v1beta1/crd/names"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/apiextensions/v1beta1/crd/property"
	"github.com/previousnext/terraform-provider-k8s/internal/resources/apiextensions/v1beta1/crd/required"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Generate the ServiceAccount.
func Generate(d *schema.ResourceData) (apiextensionsv1.CustomResourceDefinition, error) {
	var (
		name          = d.Get(FieldName).(string)
		rawLabels     = d.Get(FieldLabels).(map[string]interface{})
		group         = d.Get(FieldGroup).(string)
		version       = d.Get(FieldVersion).(string)
		scope         = d.Get(FieldScope).(string)
		rawNames      = d.Get(FieldNames).([]interface{})
		rawProperties = d.Get(FieldProperty).([]interface{})
		rawRequired   = d.Get(FieldRequired).([]interface{})
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
					Schema: &apiextensionsv1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
							Properties: property.Expand(rawProperties),
							Required:   required.Expand(rawRequired),
							Type:       "object",
						},
					},
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
