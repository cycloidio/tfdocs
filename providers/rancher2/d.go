package rancher2

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rancher2_app",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 app.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The app name (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project where the app is deployed (string)`,
				},
				resource.Attribute{
					Name:        "target_namespace",
					Description: `(Optional/Computed) The namespace name where the app is deployed (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Computed) Catalog name of the app (string)`,
				},
				resource.Attribute{
					Name:        "answers",
					Description: `(Computed) Answers for the app (map)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) Description for the app (string)`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Computed) The URL of the helm catalog app (string)`,
				},
				resource.Attribute{
					Name:        "revision_id",
					Description: `(Computed) Current revision id for the app (string)`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Computed) Template name of the app (string)`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `(Computed) Template version of the app (string)`,
				},
				resource.Attribute{
					Name:        "values_yaml",
					Description: `(Computed) values.yaml file content for the app (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for the catalog (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for the catalog (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Computed) Catalog name of the app (string)`,
				},
				resource.Attribute{
					Name:        "answers",
					Description: `(Computed) Answers for the app (map)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) Description for the app (string)`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Computed) The URL of the helm catalog app (string)`,
				},
				resource.Attribute{
					Name:        "revision_id",
					Description: `(Computed) Current revision id for the app (string)`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Computed) Template name of the app (string)`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `(Computed) Template version of the app (string)`,
				},
				resource.Attribute{
					Name:        "values_yaml",
					Description: `(Computed) values.yaml file content for the app (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for the catalog (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for the catalog (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_catalog",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 catalog.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The catalog name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the catalog. ` + "`" + `cluster` + "`" + `, ` + "`" + `global` + "`" + `, and ` + "`" + `project` + "`" + ` are supported. Default ` + "`" + `global` + "`" + ` (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Computed) The branch of the catalog repo to use (string)`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Computed) The cluster id of the catalog (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) A catalog description (string)`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Computed) The kind of the catalog. Just helm by the moment (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Computed/Sensitive) The password to access the catalog if needed (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Computed) The project id of the catalog (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Computed/Sensitive) The username to access the catalog if needed (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) The url of the catalog repo (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for the catalog (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for the catalog (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Computed) The branch of the catalog repo to use (string)`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Computed) The cluster id of the catalog (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) A catalog description (string)`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Computed) The kind of the catalog. Just helm by the moment (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Computed/Sensitive) The password to access the catalog if needed (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Computed) The project id of the catalog (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Computed/Sensitive) The username to access the catalog if needed (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) The url of the catalog repo (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for the catalog (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for the catalog (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the certificate (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project id where to assign the certificate (string)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Optional) The namespace id where to assign the namespaced certificate (string) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "certs",
					Description: `(Computed) Base64 encoded certs (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) A certificate description (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for certificate object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for certificate object (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "certs",
					Description: `(Computed) Base64 encoded certs (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) A certificate description (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for certificate object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for certificate object (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cloud_credential",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 Cloud Credentials.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Cloud Credential name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for the Cloud Credential (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for the Cloud Credential (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for the Cloud Credential (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for the Cloud Credential (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Cluster (string) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "cluster_registration_token",
					Description: `(Computed) Cluster Registration Token generated for the cluster (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "default_project_id",
					Description: `(Computed) Default project ID for the cluster (string)`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Computed) The driver used for the Cluster. ` + "`" + `imported` + "`" + `, ` + "`" + `azurekubernetesservice` + "`" + `, ` + "`" + `amazonelasticcontainerservice` + "`" + `, ` + "`" + `googlekubernetesengine` + "`" + ` and ` + "`" + `rancherKubernetesEngine` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `(Computed) Kube Config generated for the cluster (string)`,
				},
				resource.Attribute{
					Name:        "system_project_id",
					Description: `(Computed) System project ID for the cluster (string)`,
				},
				resource.Attribute{
					Name:        "rke_config",
					Description: `(Computed) The rke configuration for ` + "`" + `rke` + "`" + ` Clusters. Conflicts with ` + "`" + `aks_config` + "`" + `, ` + "`" + `eks_config` + "`" + ` and ` + "`" + `gke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "aks_config",
					Description: `(Computed) The Azure aks configuration for ` + "`" + `aks` + "`" + ` Clusters. Conflicts with ` + "`" + `eks_config` + "`" + `, ` + "`" + `gke_config` + "`" + ` and ` + "`" + `rke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "eks_config",
					Description: `(Computed) The Amazon eks configuration for ` + "`" + `eks` + "`" + ` Clusters. Conflicts with ` + "`" + `aks_config` + "`" + `, ` + "`" + `gke_config` + "`" + ` and ` + "`" + `rke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "gke_config",
					Description: `(Computed) The Google gke configuration for ` + "`" + `gke` + "`" + ` Clusters. Conflicts with ` + "`" + `aks_config` + "`" + `, ` + "`" + `eks_config` + "`" + ` and ` + "`" + `rke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) The description for Cluster (string)`,
				},
				resource.Attribute{
					Name:        "cluster_auth_endpoint",
					Description: `(Computed) Enabling the [local cluster authorized endpoint](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#local-cluster-auth-endpoint) allows direct communication with the cluster, bypassing the Rancher API proxy. (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "default_pod_security_policy_template_id",
					Description: `(Computed) [Default pod security policy template id](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#pod-security-policy-support). ` + "`" + `restricted` + "`" + ` and ` + "`" + `unrestricted` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "enable_cluster_monitoring",
					Description: `(Computed) Enable built-in cluster monitoring. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_network_policy",
					Description: `(Computed) Enable project network isolation. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Node Pool object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Node Pool object (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "cluster_registration_token",
					Description: `(Computed) Cluster Registration Token generated for the cluster (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "default_project_id",
					Description: `(Computed) Default project ID for the cluster (string)`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Computed) The driver used for the Cluster. ` + "`" + `imported` + "`" + `, ` + "`" + `azurekubernetesservice` + "`" + `, ` + "`" + `amazonelasticcontainerservice` + "`" + `, ` + "`" + `googlekubernetesengine` + "`" + ` and ` + "`" + `rancherKubernetesEngine` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `(Computed) Kube Config generated for the cluster (string)`,
				},
				resource.Attribute{
					Name:        "system_project_id",
					Description: `(Computed) System project ID for the cluster (string)`,
				},
				resource.Attribute{
					Name:        "rke_config",
					Description: `(Computed) The rke configuration for ` + "`" + `rke` + "`" + ` Clusters. Conflicts with ` + "`" + `aks_config` + "`" + `, ` + "`" + `eks_config` + "`" + ` and ` + "`" + `gke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "aks_config",
					Description: `(Computed) The Azure aks configuration for ` + "`" + `aks` + "`" + ` Clusters. Conflicts with ` + "`" + `eks_config` + "`" + `, ` + "`" + `gke_config` + "`" + ` and ` + "`" + `rke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "eks_config",
					Description: `(Computed) The Amazon eks configuration for ` + "`" + `eks` + "`" + ` Clusters. Conflicts with ` + "`" + `aks_config` + "`" + `, ` + "`" + `gke_config` + "`" + ` and ` + "`" + `rke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "gke_config",
					Description: `(Computed) The Google gke configuration for ` + "`" + `gke` + "`" + ` Clusters. Conflicts with ` + "`" + `aks_config` + "`" + `, ` + "`" + `eks_config` + "`" + ` and ` + "`" + `rke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) The description for Cluster (string)`,
				},
				resource.Attribute{
					Name:        "cluster_auth_endpoint",
					Description: `(Computed) Enabling the [local cluster authorized endpoint](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#local-cluster-auth-endpoint) allows direct communication with the cluster, bypassing the Rancher API proxy. (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "default_pod_security_policy_template_id",
					Description: `(Computed) [Default pod security policy template id](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#pod-security-policy-support). ` + "`" + `restricted` + "`" + ` and ` + "`" + `unrestricted` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "enable_cluster_monitoring",
					Description: `(Computed) Enable built-in cluster monitoring. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_network_policy",
					Description: `(Computed) Enable project network isolation. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Node Pool object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Node Pool object (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cluster_driver",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 Cluster Driver resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the cluster driver (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional/Computed) The URL to download the machine driver binary for 64-bit Linux (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Computed) Specify if the cluster driver state (bool)`,
				},
				resource.Attribute{
					Name:        "builtin",
					Description: `(Computed) Specify wheter the cluster driver is an internal cluster driver or not (bool)`,
				},
				resource.Attribute{
					Name:        "actual_url",
					Description: `(Computed) Actual url of the cluster driver (string)`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `(Computed) Verify that the downloaded driver matches the expected checksum (string)`,
				},
				resource.Attribute{
					Name:        "ui_url",
					Description: `(Computed) The URL to load for customized Add Clusters screen for this driver (string)`,
				},
				resource.Attribute{
					Name:        "whitelist_domains",
					Description: `(Computed) Domains to whitelist for the ui (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Computed) Specify if the cluster driver state (bool)`,
				},
				resource.Attribute{
					Name:        "builtin",
					Description: `(Computed) Specify wheter the cluster driver is an internal cluster driver or not (bool)`,
				},
				resource.Attribute{
					Name:        "actual_url",
					Description: `(Computed) Actual url of the cluster driver (string)`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `(Computed) Verify that the downloaded driver matches the expected checksum (string)`,
				},
				resource.Attribute{
					Name:        "ui_url",
					Description: `(Computed) The URL to load for customized Add Clusters screen for this driver (string)`,
				},
				resource.Attribute{
					Name:        "whitelist_domains",
					Description: `(Computed) Domains to whitelist for the ui (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cluster_logging",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 Cluster Logging resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The cluster id to configure logging (string) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Computed) The kind of the Cluster Logging. ` + "`" + `elasticsearch` + "`" + `, ` + "`" + `fluentd` + "`" + `, ` + "`" + `kafka` + "`" + `, ` + "`" + `splunk` + "`" + ` and ` + "`" + `syslog` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "elasticsearch_config",
					Description: `(Computed) The elasticsearch config for Cluster Logging. For ` + "`" + `kind = elasticsearch` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "fluentd_config",
					Description: `(Computed) The fluentd config for Cluster Logging. For ` + "`" + `kind = fluentd` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kafka_config",
					Description: `(Computed) The kafka config for Cluster Logging. For ` + "`" + `kind = kafka` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the cluster logging config (string)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Computed) The namespace id from cluster logging (string)`,
				},
				resource.Attribute{
					Name:        "output_flush_interval",
					Description: `(Computed) How often buffered logs would be flushed. Default: ` + "`" + `3` + "`" + ` seconds (int)`,
				},
				resource.Attribute{
					Name:        "output_tags",
					Description: `(computed) The output tags for Cluster Logging (map)`,
				},
				resource.Attribute{
					Name:        "splunk_config",
					Description: `(Computed) The splunk config for Cluster Logging. For ` + "`" + `kind = splunk` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "syslog_config",
					Description: `(Computed) The syslog config for Cluster Logging. For ` + "`" + `kind = syslog` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Cluster Logging object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Cluster Logging object (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Computed) The kind of the Cluster Logging. ` + "`" + `elasticsearch` + "`" + `, ` + "`" + `fluentd` + "`" + `, ` + "`" + `kafka` + "`" + `, ` + "`" + `splunk` + "`" + ` and ` + "`" + `syslog` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "elasticsearch_config",
					Description: `(Computed) The elasticsearch config for Cluster Logging. For ` + "`" + `kind = elasticsearch` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "fluentd_config",
					Description: `(Computed) The fluentd config for Cluster Logging. For ` + "`" + `kind = fluentd` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kafka_config",
					Description: `(Computed) The kafka config for Cluster Logging. For ` + "`" + `kind = kafka` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the cluster logging config (string)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Computed) The namespace id from cluster logging (string)`,
				},
				resource.Attribute{
					Name:        "output_flush_interval",
					Description: `(Computed) How often buffered logs would be flushed. Default: ` + "`" + `3` + "`" + ` seconds (int)`,
				},
				resource.Attribute{
					Name:        "output_tags",
					Description: `(computed) The output tags for Cluster Logging (map)`,
				},
				resource.Attribute{
					Name:        "splunk_config",
					Description: `(Computed) The splunk config for Cluster Logging. For ` + "`" + `kind = splunk` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "syslog_config",
					Description: `(Computed) The syslog config for Cluster Logging. For ` + "`" + `kind = syslog` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Cluster Logging object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Cluster Logging object (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cluster_role_template_binding",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 cluster role template binding.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The cluster id where bind cluster role template (string)`,
				},
				resource.Attribute{
					Name:        "role_template_id",
					Description: `(Optional/Computed) The role template id from create cluster role template binding (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Computed) The group ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "group_principal_id",
					Description: `(Computed) The group_principal ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Computed) The user ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_principal_id",
					Description: `(Computed) The user_principal ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Computed) The group ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "group_principal_id",
					Description: `(Computed) The group_principal ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Computed) The user ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_principal_id",
					Description: `(Computed) The user_principal ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_etcd_backup",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 etcd backup.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID to config Etcd Backup (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Etcd Backup (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "backup_config",
					Description: `(Computed) Backup config for etcd backup (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Computed) Filename of the Etcd Backup (string)`,
				},
				resource.Attribute{
					Name:        "manual",
					Description: `(Computed) Manual execution of the Etcd Backup. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Computed) Description for the Etcd Backup (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Etcd Backup object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Etcd Backup object (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "backup_config",
					Description: `(Computed) Backup config for etcd backup (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Computed) Filename of the Etcd Backup (string)`,
				},
				resource.Attribute{
					Name:        "manual",
					Description: `(Computed) Manual execution of the Etcd Backup. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Computed) Description for the Etcd Backup (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Etcd Backup object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Etcd Backup object (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_global_role_binding",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 global role binding.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the global role binding (string)`,
				},
				resource.Attribute{
					Name:        "global_role_id",
					Description: `(Optional/Computed) The global role id (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Computed) The user ID to assign global role binding (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Computed) The user ID to assign global role binding (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_multi_cluster_app",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 multi cluster app.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The multi cluster app name (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Computed) The multi cluster app catalog name (string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Computed) The multi cluster app roles (list)`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Computed) The multi cluster app target projects (list)`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Computed) The multi cluster app template name (string)`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `(Computed) The multi cluster app template version (string)`,
				},
				resource.Attribute{
					Name:        "template_version_id",
					Description: `(Computed) The multi cluster app template version ID (string)`,
				},
				resource.Attribute{
					Name:        "answers",
					Description: `(Computed) The multi cluster app answers (list)`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Computed) The multi cluster app members (list)`,
				},
				resource.Attribute{
					Name:        "revision_history_limit",
					Description: `(Computed) The multi cluster app revision history limit (int)`,
				},
				resource.Attribute{
					Name:        "revision_id",
					Description: `(Computed) Current revision id for the multi cluster app (string)`,
				},
				resource.Attribute{
					Name:        "upgrade_strategy",
					Description: `(Computed) The multi cluster app upgrade strategy (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for multi cluster app object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for multi cluster app object (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Computed) The multi cluster app catalog name (string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Computed) The multi cluster app roles (list)`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Computed) The multi cluster app target projects (list)`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Computed) The multi cluster app template name (string)`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `(Computed) The multi cluster app template version (string)`,
				},
				resource.Attribute{
					Name:        "template_version_id",
					Description: `(Computed) The multi cluster app template version ID (string)`,
				},
				resource.Attribute{
					Name:        "answers",
					Description: `(Computed) The multi cluster app answers (list)`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Computed) The multi cluster app members (list)`,
				},
				resource.Attribute{
					Name:        "revision_history_limit",
					Description: `(Computed) The multi cluster app revision history limit (int)`,
				},
				resource.Attribute{
					Name:        "revision_id",
					Description: `(Computed) Current revision id for the multi cluster app (string)`,
				},
				resource.Attribute{
					Name:        "upgrade_strategy",
					Description: `(Computed) The multi cluster app upgrade strategy (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for multi cluster app object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for multi cluster app object (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_namespace",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 namespace.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the namespace (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project id where namespace is assigned (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "container_resource_limit",
					Description: `(Computed) Default containers resource limits on namespace (List maxitem:1)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) A namespace description (string)`,
				},
				resource.Attribute{
					Name:        "resource_quota",
					Description: `(Computed) Resource quota for namespace. Rancher v2.1.x or higher (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Node Pool object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Node Pool object (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "container_resource_limit",
					Description: `(Computed) Default containers resource limits on namespace (List maxitem:1)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) A namespace description (string)`,
				},
				resource.Attribute{
					Name:        "resource_quota",
					Description: `(Computed) Resource quota for namespace. Rancher v2.1.x or higher (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Node Pool object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Node Pool object (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_node_driver",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 Node Driver resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the node driver (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional/Computed) The URL to download the machine driver binary for 64-bit Linux (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Computed) Specify if the node driver state (bool)`,
				},
				resource.Attribute{
					Name:        "builtin",
					Description: `(Computed) Specify wheter the node driver is an internal cluster driver or not (bool)`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `(Computed) Verify that the downloaded driver matches the expected checksum (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) Description of the node driver (string)`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Computed) External ID (string)`,
				},
				resource.Attribute{
					Name:        "ui_url",
					Description: `(Computed) The URL to load for customized Add Node screen for this driver (string)`,
				},
				resource.Attribute{
					Name:        "whitelist_domains",
					Description: `(Computed) Domains to whitelist for the ui (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Computed) Specify if the node driver state (bool)`,
				},
				resource.Attribute{
					Name:        "builtin",
					Description: `(Computed) Specify wheter the node driver is an internal cluster driver or not (bool)`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `(Computed) Verify that the downloaded driver matches the expected checksum (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) Description of the node driver (string)`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Computed) External ID (string)`,
				},
				resource.Attribute{
					Name:        "ui_url",
					Description: `(Computed) The URL to load for customized Add Node screen for this driver (string)`,
				},
				resource.Attribute{
					Name:        "whitelist_domains",
					Description: `(Computed) Domains to whitelist for the ui (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_node_pool",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 Node Pool resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The rke cluster id to use Node Pool (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Node Pool (string)`,
				},
				resource.Attribute{
					Name:        "node_template_id",
					Description: `(Optional/Computed) The Node Template ID to use for node creation (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "hostname_prefix",
					Description: `(Computed) The prefix for created nodes of the Node Pool (string)`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Computed) The number of nodes to create on Node Pool (int)`,
				},
				resource.Attribute{
					Name:        "control_plane",
					Description: `(Computed) RKE control plane role for created nodes (bool)`,
				},
				resource.Attribute{
					Name:        "etcd",
					Description: `(Computed) RKE etcd role for created nodes (bool)`,
				},
				resource.Attribute{
					Name:        "worker",
					Description: `(Computed) RKE role role for created nodes (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Node Pool object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Node Pool object (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "hostname_prefix",
					Description: `(Computed) The prefix for created nodes of the Node Pool (string)`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Computed) The number of nodes to create on Node Pool (int)`,
				},
				resource.Attribute{
					Name:        "control_plane",
					Description: `(Computed) RKE control plane role for created nodes (bool)`,
				},
				resource.Attribute{
					Name:        "etcd",
					Description: `(Computed) RKE etcd role for created nodes (bool)`,
				},
				resource.Attribute{
					Name:        "worker",
					Description: `(Computed) RKE role role for created nodes (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Node Pool object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Node Pool object (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_node_template",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 Node Template resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Node Template (string) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "cloud_credential_id",
					Description: `(Computed) Cloud credential ID for the Node Template. Required from rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) Description for the Node Template (string)`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Computed) The driver of the node template (string)`,
				},
				resource.Attribute{
					Name:        "engine_env",
					Description: `(Computed) Engine environment for the node template (string)`,
				},
				resource.Attribute{
					Name:        "engine_insecure_registry",
					Description: `(Computed) Insecure registry for the node template (list)`,
				},
				resource.Attribute{
					Name:        "engine_install_url",
					Description: `(Computed) Docker engine install URL for the node template (string)`,
				},
				resource.Attribute{
					Name:        "engine_label",
					Description: `(Computed) Engine label for the node template (string)`,
				},
				resource.Attribute{
					Name:        "engine_opt",
					Description: `(Computed) Engine options for the node template (map)`,
				},
				resource.Attribute{
					Name:        "engine_registry_mirror",
					Description: `(Computed) Engine registry mirror for the node template (list)`,
				},
				resource.Attribute{
					Name:        "engine_storage_driver",
					Description: `(Computed) Engine storage driver for the node template (string)`,
				},
				resource.Attribute{
					Name:        "use_internal_ip_address",
					Description: `(Computed) Engine storage driver for the node template (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Node Template object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Node Template object (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "cloud_credential_id",
					Description: `(Computed) Cloud credential ID for the Node Template. Required from rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) Description for the Node Template (string)`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Computed) The driver of the node template (string)`,
				},
				resource.Attribute{
					Name:        "engine_env",
					Description: `(Computed) Engine environment for the node template (string)`,
				},
				resource.Attribute{
					Name:        "engine_insecure_registry",
					Description: `(Computed) Insecure registry for the node template (list)`,
				},
				resource.Attribute{
					Name:        "engine_install_url",
					Description: `(Computed) Docker engine install URL for the node template (string)`,
				},
				resource.Attribute{
					Name:        "engine_label",
					Description: `(Computed) Engine label for the node template (string)`,
				},
				resource.Attribute{
					Name:        "engine_opt",
					Description: `(Computed) Engine options for the node template (map)`,
				},
				resource.Attribute{
					Name:        "engine_registry_mirror",
					Description: `(Computed) Engine registry mirror for the node template (list)`,
				},
				resource.Attribute{
					Name:        "engine_storage_driver",
					Description: `(Computed) Engine storage driver for the node template (string)`,
				},
				resource.Attribute{
					Name:        "use_internal_ip_address",
					Description: `(Computed) Engine storage driver for the node template (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Node Template object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Node Template object (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_project",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) ID of the Rancher 2 cluster (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The project name (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) Cluster-wide unique ID of the Rancher 2 project (string)`,
				},
				resource.Attribute{
					Name:        "container_resource_limit",
					Description: `(Computed) Default containers resource limits on project (List maxitem:1)`,
				},
				resource.Attribute{
					Name:        "enable_project_monitoring",
					Description: `(Computed) Enable built-in project monitoring. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "pod_security_policy_template_id",
					Description: `(Computed) Default Pod Security Policy ID for the project (string)`,
				},
				resource.Attribute{
					Name:        "resource_quota",
					Description: `(Computed) Resource quota for project. Rancher v2.1.x or higher (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Computed) UUID of the project as stored by Rancher 2 (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) The project's description (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the rancher2 project (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the rancher2 project (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) Cluster-wide unique ID of the Rancher 2 project (string)`,
				},
				resource.Attribute{
					Name:        "container_resource_limit",
					Description: `(Computed) Default containers resource limits on project (List maxitem:1)`,
				},
				resource.Attribute{
					Name:        "enable_project_monitoring",
					Description: `(Computed) Enable built-in project monitoring. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "pod_security_policy_template_id",
					Description: `(Computed) Default Pod Security Policy ID for the project (string)`,
				},
				resource.Attribute{
					Name:        "resource_quota",
					Description: `(Computed) Resource quota for project. Rancher v2.1.x or higher (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Computed) UUID of the project as stored by Rancher 2 (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) The project's description (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the rancher2 project (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the rancher2 project (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_project_logging",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 Project Logging resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project id to configure logging (string) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Computed) The kind of the Cluster Logging. ` + "`" + `elasticsearch` + "`" + `, ` + "`" + `fluentd` + "`" + `, ` + "`" + `kafka` + "`" + `, ` + "`" + `splunk` + "`" + ` and ` + "`" + `syslog` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "elasticsearch_config",
					Description: `(Computed) The elasticsearch config for Cluster Logging. For ` + "`" + `kind = elasticsearch` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "fluentd_config",
					Description: `(Computed) The fluentd config for Cluster Logging. For ` + "`" + `kind = fluentd` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kafka_config",
					Description: `(Computed) The kafka config for Cluster Logging. For ` + "`" + `kind = kafka` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the cluster logging config (string)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Computed) The namespace id from cluster logging (string)`,
				},
				resource.Attribute{
					Name:        "output_flush_interval",
					Description: `(Computed) How often buffered logs would be flushed. Default: ` + "`" + `3` + "`" + ` seconds (int)`,
				},
				resource.Attribute{
					Name:        "output_tags",
					Description: `(computed) The output tags for Cluster Logging (map)`,
				},
				resource.Attribute{
					Name:        "splunk_config",
					Description: `(Computed) The splunk config for Cluster Logging. For ` + "`" + `kind = splunk` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "syslog_config",
					Description: `(Computed) The syslog config for Cluster Logging. For ` + "`" + `kind = syslog` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Cluster Logging object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Cluster Logging object (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Computed) The kind of the Cluster Logging. ` + "`" + `elasticsearch` + "`" + `, ` + "`" + `fluentd` + "`" + `, ` + "`" + `kafka` + "`" + `, ` + "`" + `splunk` + "`" + ` and ` + "`" + `syslog` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "elasticsearch_config",
					Description: `(Computed) The elasticsearch config for Cluster Logging. For ` + "`" + `kind = elasticsearch` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "fluentd_config",
					Description: `(Computed) The fluentd config for Cluster Logging. For ` + "`" + `kind = fluentd` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kafka_config",
					Description: `(Computed) The kafka config for Cluster Logging. For ` + "`" + `kind = kafka` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the cluster logging config (string)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Computed) The namespace id from cluster logging (string)`,
				},
				resource.Attribute{
					Name:        "output_flush_interval",
					Description: `(Computed) How often buffered logs would be flushed. Default: ` + "`" + `3` + "`" + ` seconds (int)`,
				},
				resource.Attribute{
					Name:        "output_tags",
					Description: `(computed) The output tags for Cluster Logging (map)`,
				},
				resource.Attribute{
					Name:        "splunk_config",
					Description: `(Computed) The splunk config for Cluster Logging. For ` + "`" + `kind = splunk` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "syslog_config",
					Description: `(Computed) The syslog config for Cluster Logging. For ` + "`" + `kind = syslog` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Cluster Logging object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Cluster Logging object (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_project_role_template_binding",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 project role template binding.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project id where bind project role template (string)`,
				},
				resource.Attribute{
					Name:        "role_template_id",
					Description: `(Optional/Computed) The role template id from create project role template binding (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Computed) The group ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "group_principal_id",
					Description: `(Computed) The group_principal ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Computed) The user ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_principal_id",
					Description: `(Computed) The user_principal ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Computed) The group ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "group_principal_id",
					Description: `(Computed) The group_principal ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Computed) The user ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_principal_id",
					Description: `(Computed) The user_principal ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_registry",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 docker registry.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the registry (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project id where to assign the registry (string)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Optional) The namespace id where to assign the namespaced registry (string) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "registries",
					Description: `(Computed) Registries data for registry (list)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) A registry description (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Registry object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Registry object (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "registries",
					Description: `(Computed) Registries data for registry (list)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) A registry description (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for Registry object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for Registry object (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_secret",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 secret.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the secret (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project id where to assign the secret (string)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Optional) The namespace id where to assign the namespaced secret (string) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Computed) Secret key/value data. Base64 encoding required for values (map)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) A secret description (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for secret object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for secret object (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Computed) Secret key/value data. Base64 encoding required for values (map)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) A secret description (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for secret object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for secret object (map)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_setting",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The setting name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `the settting's value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `the settting's value.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_user",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the global role binding (string)`,
				},
				resource.Attribute{
					Name:        "global_role_id",
					Description: `(Optional/Computed) The global role id (string) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Computed) The user ID to assign global role binding (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Computed) The user ID to assign global role binding (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels of the resource (map)`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"rancher2_app":                           0,
		"rancher2_catalog":                       1,
		"rancher2_certificate":                   2,
		"rancher2_cloud_credential":              3,
		"rancher2_cluster":                       4,
		"rancher2_cluster_driver":                5,
		"rancher2_cluster_logging":               6,
		"rancher2_cluster_role_template_binding": 7,
		"rancher2_etcd_backup":                   8,
		"rancher2_global_role_binding":           9,
		"rancher2_multi_cluster_app":             10,
		"rancher2_namespace":                     11,
		"rancher2_node_driver":                   12,
		"rancher2_node_pool":                     13,
		"rancher2_node_template":                 14,
		"rancher2_project":                       15,
		"rancher2_project_logging":               16,
		"rancher2_project_role_template_binding": 17,
		"rancher2_registry":                      18,
		"rancher2_secret":                        19,
		"rancher2_setting":                       20,
		"rancher2_user":                          21,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
