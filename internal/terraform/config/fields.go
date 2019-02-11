package config

import "github.com/hashicorp/terraform/helper/schema"

const (
	// FieldHost identifier for host field.
	FieldHost = "host"
	// FieldInsecure identifier for insecure field.
	FieldInsecure = "insecure"
	// FieldClientCertificate identifier for client_certificate field.
	FieldClientCertificate = "client_certificate"
	// FieldClientKey identifier for client_key field.
	FieldClientKey = "client_key"
	// FieldClusterCACertificate identifier for cluster_ca_certificate field.
	FieldClusterCACertificate = "cluster_ca_certificate"
	// FieldConfigPath identifier for config_path field.
	FieldConfigPath = "config_path"
)

// Fields which are used to configure the Kubernetes client.
func Fields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		FieldHost: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_HOST", ""),
			Description: "The hostname (in form of URI) of Kubernetes master.",
		},
		FieldInsecure: {
			Type:        schema.TypeBool,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_INSECURE", false),
			Description: "Whether server should be accessed without verifying the TLS certificate.",
		},
		FieldClientCertificate: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_CLIENT_CERT_DATA", ""),
			Description: "PEM-encoded client certificate for TLS authentication.",
		},
		FieldClientKey: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_CLIENT_KEY_DATA", ""),
			Description: "PEM-encoded client certificate key for TLS authentication.",
		},
		FieldClusterCACertificate: {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("KUBE_CLUSTER_CA_CERT_DATA", ""),
			Description: "PEM-encoded root certificates bundle for TLS authentication.",
		},
	}
}
