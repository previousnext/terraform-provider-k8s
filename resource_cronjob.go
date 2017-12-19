package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	utilsenv "github.com/previousnext/terraform-provider-k8s/utils/env"
	utilsid "github.com/previousnext/terraform-provider-k8s/utils/id"
	utilsslice "github.com/previousnext/terraform-provider-k8s/utils/slicestring"
	volhostpath "github.com/previousnext/terraform-provider-k8s/volumes/hostpath"
)

func resourceCronJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceCronJobCreate,
		Read:   resourceCronJobRead,
		Update: resourceCronJobUpdate,
		Delete: resourceCronJobDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the CronJob.",
				Required:    true,
			},
			"namespace": {
				Type:        schema.TypeString,
				Description: "Namespace which the CronJob will be run in.",
				Required:    true,
			},
			"schedule": {
				Type:        schema.TypeString,
				Description: "How often to run this CronJob.",
				Required:    true,
			},
			"image": {
				Type:        schema.TypeString,
				Description: "The image to execute for this CronJob.",
				Required:    true,
			},
			"command": {
				Type:        schema.TypeList,
				Description: "Command is the command line to execute inside the container as a part of the job.",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"env": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of environment variables to set in the container. Cannot be updated.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the environment variable.",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Value of the environment variable.",
						},
					},
				},
			},
			"hostpath": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of host path volumes which we want to mount.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the mount.",
						},
						"source": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Source path from the host filesystem.",
						},
						"target": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Target mount for within the container.",
						},
					},
				},
			},
		},
	}
}

func resourceCronJobCreate(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	cronJob, err := generateCronJob(d)
	if err != nil {
		return err
	}

	out, err := conn.BatchV1beta1().CronJobs(cronJob.ObjectMeta.Namespace).Create(&cronJob)
	if err != nil {
		return fmt.Errorf("failed to create batch/v1beta1/cronjob: %s", err)
	}

	d.SetId(utilsid.Join(out.ObjectMeta))

	return nil
}

func resourceCronJobRead(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	namespace, name, err := utilsid.Split(d.Id())
	if err != nil {
		return err
	}

	cronJob, err := conn.BatchV1beta1().CronJobs(namespace).Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		// This is how we tell Terraform that the resource does not exist.
		d.SetId("")
		return nil
	} else if err != nil {
		return err
	}

	if len(cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers) == 0 {
		return fmt.Errorf("cannot find a container associated with cronjob")
	}

	d.Set("name", cronJob.ObjectMeta.Name)
	d.Set("namespace", cronJob.ObjectMeta.Namespace)
	d.Set("schedule", cronJob.Spec.Schedule)
	d.Set("image", cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image)
	d.Set("command", cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Command)
	d.Set("env", utilsenv.Flatten(cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Env))
	d.Set("hostpath", volhostpath.Flatten(cronJob.Spec.JobTemplate.Spec.Template.Spec.Volumes, cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].VolumeMounts))

	return nil
}

func resourceCronJobUpdate(d *schema.ResourceData, m interface{}) error {
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

func resourceCronJobDelete(d *schema.ResourceData, m interface{}) error {
	conn := m.(*kubernetes.Clientset)

	namespace, name, err := utilsid.Split(d.Id())
	if err != nil {
		return err
	}

	return conn.BatchV1beta1().CronJobs(namespace).Delete(name, &metav1.DeleteOptions{})
}

// Helper function for generating the CronJob object.
func generateCronJob(d *schema.ResourceData) (batchv1beta1.CronJob, error) {
	var (
		name      = d.Get("name").(string)
		namespace = d.Get("namespace").(string)
		schedule  = d.Get("schedule").(string)
		image     = d.Get("image").(string)
		command   = d.Get("command").([]interface{})
		env       = d.Get("env").([]interface{})
		hostpath  = d.Get("hostpath").([]interface{})
	)

	cronJob := batchv1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: batchv1beta1.CronJobSpec{
			Schedule: schedule,
			JobTemplate: batchv1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							RestartPolicy: corev1.RestartPolicyNever,
						},
					},
				},
			},
		},
	}

	envs, err := utilsenv.Expand(env)
	if err != nil {
		return cronJob, err
	}

	volumes, mounts, err := volhostpath.Expand(hostpath)
	if err != nil {
		return cronJob, err
	}
	cronJob.Spec.JobTemplate.Spec.Template.Spec.Volumes = append(cronJob.Spec.JobTemplate.Spec.Template.Spec.Volumes, volumes...)

	container := corev1.Container{
		Name:         name,
		Image:        image,
		Command:      utilsslice.Expand(command),
		Env:          envs,
		VolumeMounts: mounts,
	}

	cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers = append(cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers, container)

	return cronJob, nil
}
