package cronjob

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/client-go/kubernetes"
)

func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	cronJob, err := generateCronJob(d)
	if err != nil {
		return err
	}

	_, err = conn.BatchV1beta1().CronJobs(cronJob.ObjectMeta.Namespace).Update(&cronJob)
	if err != nil {
		return fmt.Errorf("failed to update batch/v1beta1/cronjob: %s", err)
	}

	return nil
}
