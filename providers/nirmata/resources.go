package nirmata

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nirmata_aws_role_credentials",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the credentials.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) use as the default credentials.`,
				},
				resource.Attribute{
					Name:        "aws_access_key_id",
					Description: `(Required) The AWS access key ID.`,
				},
				resource.Attribute{
					Name:        "aws_secret_key",
					Description: `(Required) The AWS secret access key.`,
				},
				resource.Attribute{
					Name:        "aws_role_arn",
					Description: `(Optional) The Amazon Resource Name (ARN) of the role to assume.`,
				},
				resource.Attribute{
					Name:        "aws_external_id",
					Description: `(Optional) A unique identifier that might be required when you assume a role in another account.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) the type of cluster to create.`,
				},
				resource.Attribute{
					Name:        "nodepools",
					Description: `A list of [nodepool](#nodepool) types.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) labels to set on cluster. ## Nested Blocks ### nodepool`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Required) the number of worker nodes for the cluster`,
				},
				resource.Attribute{
					Name:        "enable_auto_scaling",
					Description: `(Optional) Enable autoscaling for cluster. default valie is disable.`,
				},
				resource.Attribute{
					Name:        "min_count",
					Description: `(Optional) Set minimun node count value for cluster. ` + "`" + `enable_auto_scaling` + "`" + ` must set true to set min_count.`,
				},
				resource.Attribute{
					Name:        "max_count",
					Description: `(Optional) Set max node count value for cluster. ` + "`" + `enable_auto_scaling` + "`" + ` must set true to set max_count.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_addon",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) unique name for the cluster add-on service.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) the kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) the catalog.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required) the application.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) defaults to the application name.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) defaults to the application name and cluster name.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) the release channel`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) labels to set on cluster add-on.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_imported",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) the cloud credentials to use to locate and import the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) the cluster type to apply.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) the region the cluster is located in.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Required) the project the cluster is located in.`,
				},
				resource.Attribute{
					Name:        "delete_action",
					Description: `(Optional) whether to delete or remove the cluster on destroy. Defaults to ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_metadata",
					Description: `(Optional) key-value pairs that will be provisioned as a ConfigMap called system-metadata-map in the cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) labels to set on cluster. ### vault_auth`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) the vault authentication path. The variable $(cluster.name) is allowed in the path for uniquenes.`,
				},
				resource.Attribute{
					Name:        "addon_name",
					Description: `(Required) the associated Vault Agent Injector add-on`,
				},
				resource.Attribute{
					Name:        "credentials_name",
					Description: `(Required) the Vault credentials to use`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) a list of application roles to configure for add-on services`,
				},
				resource.Attribute{
					Name:        "delete_auth_path",
					Description: `(Optional) delete auth path on cluster delete #### roles`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Required) the allowed service account name`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) the allowed namespace`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) the applied policies`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_registered",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) the cluster type to apply.`,
				},
				resource.Attribute{
					Name:        "controller_yamls",
					Description: `(Computed) the controller YAML`,
				},
				resource.Attribute{
					Name:        "controller_yamls_folder",
					Description: `(Computed) a local temporary folder with the controller YAML files`,
				},
				resource.Attribute{
					Name:        "controller_yamls_count",
					Description: `(Computed) the controller YAML file count`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) labels to set on cluster.`,
				},
				resource.Attribute{
					Name:        "delete_action",
					Description: `(Optional) whether to delete or remove the cluster on destroy. Defaults to ` + "`" + `remove` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_type_eks",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) the EKS version (e.g. 1.18)`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) the cloud credentials to use.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) the AWS region for the cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) the AWS VPC ID for the cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) a list of AWS VPC subnets to use for the cluster.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) a list of AWS VPC security groups to use for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_role_arn",
					Description: `(Required) the cluster role ARN.`,
				},
				resource.Attribute{
					Name:        "enable_private_endpoint",
					Description: `(Optional) specify if the cluster API seerver endpoint should be in a private network.`,
				},
				resource.Attribute{
					Name:        "enable_identity_provider",
					Description: `(Optional) enable IAM roles for service accounts.`,
				},
				resource.Attribute{
					Name:        "auto_sync_namespaces",
					Description: `(Optional) enable automatic synchronization of cluster namespaces to Nirmata.`,
				},
				resource.Attribute{
					Name:        "enable_secrets_encryption",
					Description: `(Optional) enable encryption at rest for secrets.`,
				},
				resource.Attribute{
					Name:        "kms_key_arn",
					Description: `(Optional) the KMS key ARN to use for secrets encryption.`,
				},
				resource.Attribute{
					Name:        "log_types",
					Description: `(Optional) the log types to collect.`,
				},
				resource.Attribute{
					Name:        "enable_fargate",
					Description: `(Optional) enable Fargate to provision nodes based on workload resource requests.`,
				},
				resource.Attribute{
					Name:        "system_metadata",
					Description: `(Optional) key-value pairs that will be provisioned as a ConfigMap called system-metadata-map in the cluster.`,
				},
				resource.Attribute{
					Name:        "allow_override_credentials",
					Description: `(Optional) Allow passing in cloud credentials when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "cluster_field_override",
					Description: `(Optional) Allow override of cluster settings ('network' and 'subnetwork') when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "nodepool_field_override",
					Description: `(Optional) Allow override of node fields ('disk_size' and 'machine_type') when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "nodepools",
					Description: `(Optional) A list of [nodepool](#nodepool) types.`,
				},
				resource.Attribute{
					Name:        "addons",
					Description: `(Optional) a list of add-on services.`,
				},
				resource.Attribute{
					Name:        "vault_auth",
					Description: `(Optional) vault authentication configuration. ## Nested Blocks ### nodepool`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the node pool.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) the EC2 instance type (e.g. "t3.medium").`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) the worker node disk size in GB (e.g. 60).`,
				},
				resource.Attribute{
					Name:        "ssh_key_name",
					Description: `(Required) the SSK key pair to access nodes.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) the Node security groups.`,
				},
				resource.Attribute{
					Name:        "iam_role",
					Description: `(Required) the IAM role to use for nodes.`,
				},
				resource.Attribute{
					Name:        "ami_type",
					Description: `(Required) the EKS-optimized Amazon Machine Image (AMI) type to use for node images.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) an Amazon Machine Image (AMI) IS to use for node images.`,
				},
				resource.Attribute{
					Name:        "node_annotations",
					Description: `(Optional) Annotations to set on each node in this pool. This setting is permanent.`,
				},
				resource.Attribute{
					Name:        "node_labels",
					Description: `(Optional) Labels to set on each Kubernetes node in this node pool. This setting is permanent. ### addons`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the add-on service`,
				},
				resource.Attribute{
					Name:        "addon_selector",
					Description: `(Required) the catalog application name`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) the catalog name`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) the release channel`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `(Optional) a sequence number to control installation order ### vault_auth`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) the vault authentication path. The variable $(cluster.name) is allowed for uniquenes.`,
				},
				resource.Attribute{
					Name:        "addon_name",
					Description: `(Required) the associated Vault Agent Injector add-on`,
				},
				resource.Attribute{
					Name:        "credentials_name",
					Description: `(Required) the Vault credentials to use`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) a list of application roles to configure for add-on services`,
				},
				resource.Attribute{
					Name:        "delete_auth_path",
					Description: `(Optional) delete auth path on cluster delete #### roles`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Required) the allowed service account name`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) the allowed namespace`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) the applied policies`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_type_gke",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) the GKE version (e.g. 1.16.12-gke.3)`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) the cloud credentials to use.`,
				},
				resource.Attribute{
					Name:        "location_type",
					Description: `(Required) Regional or Zonal. A regional cluster has multiple replicas of the control plane, running in multiple zones within a given region. A zonal cluster has a single replica of the control plane running in a single zone.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) the GCP region into which the cluster should be deployed (e.g. "us-central1"). Required when location_type is ` + "`" + `Regional` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) the GCP zone into which the cluster should be deployed (e.g. "us-central1-a"). Required if location_type is ` + "`" + `Zonal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_locations",
					Description: `(Optional) the nodes should be deployed. Selecting more than one zone increases availability. (e.g. ["asia-east1-a"]). Required if location_type is ` + "`" + `Regional` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_secrets_encryption",
					Description: `(Optional) enables envelope encryption for Kubernetes Secrets.`,
				},
				resource.Attribute{
					Name:        "enable_workload_identity",
					Description: `(Optional) Workload Identity is the recommended way to access Google Cloud services from applications running within GKE due to its improved security properties and manageability.`,
				},
				resource.Attribute{
					Name:        "workload_pool",
					Description: `(Optional) Workload Identity relies on a Workload Pool to aggregate identity across multiple clusters. Required if enable_secrets_encryption is true`,
				},
				resource.Attribute{
					Name:        "secrets_encryption_key",
					Description: `(Optional) the Resource ID of the key you want to use (e.g. projects/project-name/locations/global/keyRings/my-keyring/cryptoKeys/my-key). Required if enable_workload_identity is true.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) the cluster network (e.g. "default")`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Required) the node subnetwork (e.g. "default")`,
				},
				resource.Attribute{
					Name:        "cluster_ipv4_cidr",
					Description: `(Optional) All pods in the cluster are assigned an IP address from this range. Enter a range (in CIDR notation) within a network range, a mask, or leave this field blank to use a default range. This setting is permanent.`,
				},
				resource.Attribute{
					Name:        "services_ipv4_cidr",
					Description: `(Optional) Cluster services will be assigned an IP address from this IP address range. Enter a range (in CIDR notation) within a network range, a mask, or leave this field blank to use a default range. This setting is permanent.`,
				},
				resource.Attribute{
					Name:        "cloud_run",
					Description: `(Optional) Cloud Run for Anthos enables you to easily deploy stateless apps and functions to this cluster using the Cloud Run experience. Cloud Run for Anthos automatically manages underlying resources and scales your app based on requests.`,
				},
				resource.Attribute{
					Name:        "enable_network_policy",
					Description: `(Optional) The Kubernetes Network Policy API allows the cluster administrator to specify what pods are allowed to communicate with each other. Google Kubernetes Engine has partnered with Tigera to provide Project Calico to enforce network policies within your cluster.`,
				},
				resource.Attribute{
					Name:        "enable_http_load_balancing",
					Description: `(Optional) The HTTP Load Balancing add-on is required to use the Google Cloud Load Balancer with Kubernetes Ingress. If enabled, a controller will be installed to coordinate applying load balancing configuration changes to your GCP project`,
				},
				resource.Attribute{
					Name:        "enable_vertical_pod_autoscaling",
					Description: `(Optional) Vertical Pod Autoscaling automatically analyzes and adjusts your containers' CPU requests and memory requests.`,
				},
				resource.Attribute{
					Name:        "enable_horizontal_pod_autoscaling",
					Description: ``,
				},
				resource.Attribute{
					Name:        "enable_maintenance_policy",
					Description: `(Optional) To specify regular times for maintenance, enable maintenance windows. Normally, routine Kubernetes Engine maintenance may run at any time on your cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_start_time",
					Description: `(Optional) Start time for the maintenance window.`,
				},
				resource.Attribute{
					Name:        "maintenance_duration",
					Description: `(Optional) Duration for the maintenance window in hours.`,
				},
				resource.Attribute{
					Name:        "maintenance_recurrence",
					Description: `(Optional) Recurrence rule specification (RRULE) for the maintenance window. Example RRule to run maintenance during weekends: 'FREQ=WEEKLY;BYDAY=SA,SU'.`,
				},
				resource.Attribute{
					Name:        "maintenance_exclusion_timewindow",
					Description: `(Optional) To specify times when routine, non-emergency maintenance won't happen, set up to 3 maintenance exclusions. Normally, routine Kubernetes Engine maintenance may run at any time on your cluster.`,
				},
				resource.Attribute{
					Name:        "system_metadata",
					Description: `(Optional) key-value pairs that will be provisioned as a ConfigMap called system-metadata-map in the cluster.`,
				},
				resource.Attribute{
					Name:        "allow_override_credentials",
					Description: `(Optional) Allow passing in cloud credentials when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "cluster_field_override",
					Description: `(Optional) Allow override of cluster settings ('network' and 'subnetwork') when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "nodepool_field_override",
					Description: `(Optional) Allow override of node fields ('disk_size' and 'machine_type') when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "nodepools",
					Description: `(Optional) A list of [nodepool](#nodepool) types.`,
				},
				resource.Attribute{
					Name:        "addons",
					Description: `(Optional) a list of add-on services.`,
				},
				resource.Attribute{
					Name:        "vault_auth",
					Description: `(Optional) vault authentication configuration. ## Nested Blocks ### nodepool`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required) the GCP machine type (e.g. "e2-standard-2")`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) the worker node disk size in GB (e.g. 60)`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Optional) Applications running on the VM use the service account to call Google Cloud APIs. Use Permissions on the console menu to create a service account or use the default service account if available. Service account is permanent`,
				},
				resource.Attribute{
					Name:        "enable_preemptible_nodes",
					Description: `(Optional) Preemptible nodes are Compute Engine instances that last up to 24 hours and provide no availability guarantees, but are priced lower. This setting is permanent.`,
				},
				resource.Attribute{
					Name:        "node_annotations",
					Description: `(Optional) Annotations to set on each node in this pool. This setting is permanent.`,
				},
				resource.Attribute{
					Name:        "node_labels",
					Description: `(Optional) Labels to set on each Kubernetes node in this node pool. This setting is permanent. ### addons`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the add-on service`,
				},
				resource.Attribute{
					Name:        "addon_selector",
					Description: `(Required) the catalog application name`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) the catalog name`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) the release channel`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `(Optional) a sequence number to control installation order ### vault_auth`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) the vault authentication path. The variable $(cluster.name) is allowed in the path for uniquenes.`,
				},
				resource.Attribute{
					Name:        "addon_name",
					Description: `(Required) the associated Vault Agent Injector add-on`,
				},
				resource.Attribute{
					Name:        "credentials_name",
					Description: `(Required) the Vault credentials to use`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) a list of application roles to configure for add-on services`,
				},
				resource.Attribute{
					Name:        "delete_auth_path",
					Description: `(Optional) delete auth path on cluster delete #### roles`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Required) the allowed service account name`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) the allowed namespace`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) the applied policies`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_type_registered",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "cloud",
					Description: `(Optional) the cloud provider. Defaults to ` + "`" + `Other` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "addons",
					Description: `(Optional) a list of add-on services.`,
				},
				resource.Attribute{
					Name:        "vault_auth",
					Description: `(Optional) vault authentication configuration.`,
				},
				resource.Attribute{
					Name:        "system_metadata",
					Description: `(Optional) key-value pairs that will be provisioned as a ConfigMap called system-metadata-map in the cluster type. ## Nested Blocks ### addons`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the add-on service`,
				},
				resource.Attribute{
					Name:        "addon_selector",
					Description: `(Required) the catalog application name`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) the catalog name`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) the release channel`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `(Optional) a sequence number to control installation order ### vault_auth`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) the vault authentication path. The variable $(cluster.name) is allowed in the path for uniquenes.`,
				},
				resource.Attribute{
					Name:        "addon_name",
					Description: `(Required) the associated Vault Agent Injector add-on`,
				},
				resource.Attribute{
					Name:        "credentials_name",
					Description: `(Required) the Vault credentials to use`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) a list of application roles to configure for add-on services`,
				},
				resource.Attribute{
					Name:        "delete_auth_path",
					Description: `(Optional) delete auth path on cluster delete #### roles`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Required) the allowed service account name`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) the allowed namespace`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) the applied policies`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_environment",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the environment.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) the environnment type.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) the kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) the cluster namespace bound to this environment. Defaults to the environment name.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_environment_type",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) a unique name for the environment type.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) use as the default environment type.`,
				},
				resource.Attribute{
					Name:        "resource_limits",
					Description: `(Required) a map of resource limits for the environment. Commonly used resources include ` + "`" + `cpu` + "`" + `, ` + "`" + `memory` + "`" + `, and ` + "`" + `storage` + "`" + `. Check the [Kubernetes docs](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) for a complete reference.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"nirmata_aws_role_credentials":    0,
		"nirmata_cluster":                 1,
		"nirmata_cluster_addon":           2,
		"nirmata_cluster_imported":        3,
		"nirmata_cluster_registered":      4,
		"nirmata_cluster_type_eks":        5,
		"nirmata_cluster_type_gke":        6,
		"nirmata_cluster_type_registered": 7,
		"nirmata_environment":             8,
		"nirmata_environment_type":        9,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
