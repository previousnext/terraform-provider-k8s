package storageclass

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/apimachinery/objectmeta"
	"github.com/previousnext/terraform-provider-k8s/internal/terraform/id"
)

// Read the StorageClass.
func Read(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	_, name, err := id.Split(d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to get ID")
	}

	storageclass, err := conn.StorageV1().StorageClasses().Get(name, metav1.GetOptions{})
	if kerrors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	d.Set(objectmeta.FieldObjectMeta, objectmeta.Flatten(storageclass.ObjectMeta))
	d.Set(FieldProvisioner, storageclass.Provisioner)

	return nil
}
