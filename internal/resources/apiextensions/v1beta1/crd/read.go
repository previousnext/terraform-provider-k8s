package crd

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/previousnext/terraform-provider-k8s/internal/resources/apiextensions/v1beta1/crd/names"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/config"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the StorageClass.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*config.Client)

	_, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	crd, err := conn.APIExtensions().ApiextensionsV1beta1().CustomResourceDefinitions().Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	d.Set(FieldName, crd.ObjectMeta.Name)
	d.Set(FieldLabels, crd.ObjectMeta.Labels)
	d.Set(FieldAnnotations, crd.ObjectMeta.Annotations)

	d.Set(FieldGroup, crd.Spec.Group)
	d.Set(FieldVersion, crd.Spec.Version)
	d.Set(FieldScope, crd.Spec.Scope)
	d.Set(FieldNames, names.Flatten(crd.Spec.Names))

	return nil
}
