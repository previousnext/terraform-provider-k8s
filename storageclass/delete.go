package storageclass

import (
	"github.com/hashicorp/terraform/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/utils/id"
)

func resourceDelete(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	_, name, err := id.Split(d.Id())
	if err != nil {
		return err
	}

	return conn.StorageV1beta1().StorageClasses().Delete(name, &metav1.DeleteOptions{})
}
