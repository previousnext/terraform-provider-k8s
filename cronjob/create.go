package cronjob

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"

	"github.com/previousnext/terraform-provider-k8s/utils/id"
)

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	cronJob, err := generateCronJob(d)
	if err != nil {
		return err
	}

	out, err := conn.BatchV1beta1().CronJobs(cronJob.ObjectMeta.Namespace).Create(&cronJob)
	if err != nil {
		return fmt.Errorf("failed to create batch/v1beta1/cronjob: %s", err)
	}

	d.SetId(id.Join(out.ObjectMeta))

	return nil
}
