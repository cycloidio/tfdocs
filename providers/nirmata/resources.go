package nirmata

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nirmata_aws_cloud_credentials",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the AWS cloud credentials.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Use "default credentials" as the value for this field.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Required) Select the access type for the AWS credentials (it is either access_key or assume_role).`,
				},
				resource.Attribute{
					Name:        "aws_access_key_id",
					Description: `(Optional) The AWS access key ID. This value is required if the access_type is access_key.`,
				},
				resource.Attribute{
					Name:        "aws_secret_key",
					Description: `(Optional) Enter the AWS secret access key. This value is required if the access_type is access_key.`,
				},
				resource.Attribute{
					Name:        "aws_role_arn",
					Description: `(Optional) The Amazon Resource Name (ARN) is the unique identifier of AWS resources. It is the role that is assumed for access type. This value is required if the access_type is assume_role.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_catalog",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the catalog.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) This field indicates the description for the catalog.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) This field indicates the labels to be set on the catalog.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_catalog_git_app",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the application in the catalog.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) This field indicates the namespace for the git application.`,
				},
				resource.Attribute{
					Name:        "git_credentials",
					Description: `(Required) This field indicates the git credentials name.`,
				},
				resource.Attribute{
					Name:        "git_repository",
					Description: `(Required) This is the repository URL as used in the git clone command.`,
				},
				resource.Attribute{
					Name:        "git_branch",
					Description: `(Required) Enter the git branch name to track.`,
				},
				resource.Attribute{
					Name:        "git_directory_list",
					Description: `(Optional) This field indicates the git directories to track.`,
				},
				resource.Attribute{
					Name:        "git_include_list",
					Description: `(Optional) the file extensions to track.`,
				},
				resource.Attribute{
					Name:        "fixed_kustomization",
					Description: `(Optional) enable fixed kustomize to select kustomizations for your application.`,
				},
				resource.Attribute{
					Name:        "target_based_kustomization",
					Description: `(Optional) enable target based kustomize to select kustomizations for your application.`,
				},
				resource.Attribute{
					Name:        "kustomization_file_path",
					Description: `(Required if fixed_kustomization or target_based_kustomization is set) the kustomization file path. kustomization file path required if fixed_kustomization or target_based_kustomization selected.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name to identify your application.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) Enter the name of the catalog.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required) Enter the application name.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) Enter the channel from which the application should be deployed.`,
				},
				resource.Attribute{
					Name:        "environments",
					Description: `(Required) Enter the list of environments to deploy an application.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Enter the version for the application.`,
				},
				resource.Attribute{
					Name:        "rollout_name",
					Description: `(Required) Enter a unique name for rollout.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) Enter the name of the catalog.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required) Enter the application name.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) Enter the channel from which the application should be deployed.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Enter the version of the application.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_catalog_helm_app",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the application in the catalog.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required) Enter the application name.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Enter the repository URL.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required). Enter the location of the chart (for example, https://charts.bitnami.com/bitnami/airflow-0.0.1.tgz).`,
				},
				resource.Attribute{
					Name:        "app_version",
					Description: `(Required). Enter the version of the application (for example, "0.0.1").`,
				},
				resource.Attribute{
					Name:        "chart_version",
					Description: `(Required).Enter the version of the chart (for example, "1.2.3").`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_catalog_promote_version",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rollout_name",
					Description: `(Required) Enter a unique name for rollout.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) Enter the name of the catalog.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required) Enter the application name.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) Enter the channel from which the application should be deployed.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Enter the version of the application.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_catalog_run_app",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name to identify your application.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) Enter the name of the catalog.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required) Enter the application name.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) Enter the channel from which the application should be deployed.`,
				},
				resource.Attribute{
					Name:        "environments",
					Description: `(Required) Enter the list of environments to deploy an application.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) This field indicates the version of the application.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_catalog_yaml_app",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the application in the catalog.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) Enter the name of the catalog.`,
				},
				resource.Attribute{
					Name:        "yamls",
					Description: `(Required) Enter the path for the yaml file.`,
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
					Description: `(Required) Enter a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) Enter the type of cluster to create.`,
				},
				resource.Attribute{
					Name:        "nodepools",
					Description: `Indicates a list of [nodepool](#nodepool) types.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) This field indicates the labels to be set on the cluster.`,
				},
				resource.Attribute{
					Name:        "delete_action",
					Description: `(Optional) This field indicates that if delete_action is set to ` + "`" + `remove` + "`" + `, then the cluster gets removed from the Nirmata platform and not from the original provider. If delete_action is set to ` + "`" + `delete` + "`" + `, then the cluster gets deleted from the Nirmata platform as well as from the original provider.`,
				},
				resource.Attribute{
					Name:        "creation_timeout_minutes",
					Description: `(Optional) This field is set to maximum time to create a cluster. ## Nested Blocks ### nodepool`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Required) Enter the number of worker nodes for the cluster`,
				},
				resource.Attribute{
					Name:        "enable_auto_scaling",
					Description: `(Optional) This field indicates to enable autoscaling for the cluster. The default value is "disable".`,
				},
				resource.Attribute{
					Name:        "min_count",
					Description: `(Optional) This field indicates to set the minimum node count value for the cluster. To set this value, you must set ` + "`" + `enable_auto_scaling` + "`" + ` to true.`,
				},
				resource.Attribute{
					Name:        "max_count",
					Description: `(Optional) This field indicates the set max node count value for the cluster. To set this value, you must set ` + "`" + `enable_auto_scaling` + "`" + ` to true.`,
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
					Description: `(Required) Enter a unique name for the cluster add-on service.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) Enter the kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) Enter the catalog name.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required) Enter the application name.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) This field indicates the default value to the application name.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) This field indicates the defaults to the application name and the cluster name.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) Enter the channel from which the application should be deployed.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) This field indicates the labels set on cluster add-on.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_aks_registered",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) Enter the cluster type to apply to the cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) This field indicates the labels to be set on cluster.`,
				},
				resource.Attribute{
					Name:        "delete_action",
					Description: `(Optional) This field indicates whether to delete or remove the cluster on destroy. The default value is ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) This field indicates the url of the kubernetes cluster API server.`,
				},
				resource.Attribute{
					Name:        "owner_info",
					Description: `(Optional) The [owner_info](#owner_info) for this cluster, if it has to be overridden.`,
				},
				resource.Attribute{
					Name:        "access_control_list",
					Description: `(Optional) List of additional [ACLs](#access_control_list) for this cluster. ### owner_info`,
				},
				resource.Attribute{
					Name:        "owner_type",
					Description: `(Required) The type of the owner. Valid values are user or team.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) The name of the user/team. ### access_control_list`,
				},
				resource.Attribute{
					Name:        "entity_type",
					Description: `(Required) The type of entity. Valid values are user or team.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Required) The permission. Valid values are admin, edit, view.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user/team.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_eks_imported",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) Enter the cloud credentials that is used to locate and import the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) Enter the cluster type to apply for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) Enter the type of cluster (example, gke/eks).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Enter the region of the cluster that is located in.`,
				},
				resource.Attribute{
					Name:        "delete_action",
					Description: `(Optional) This field indicates whether to delete or remove the cluster on destroy. The default value is ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_metadata",
					Description: `(Optional) This field indicates the key-value pairs that will be provisioned as a ConfigMap called system-metadata-map in the cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) This field indicates the labels to be set on the cluster.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) This field indicates the url of the kubernetes cluster API server.`,
				},
				resource.Attribute{
					Name:        "owner_info",
					Description: `(Optional) The [owner_info](#owner_info) for this cluster, if it has to be overridden.`,
				},
				resource.Attribute{
					Name:        "access_control_list",
					Description: `(Optional) List of additional [ACLs](#access_control_list) for this cluster. ### owner_info`,
				},
				resource.Attribute{
					Name:        "owner_type",
					Description: `(Required) The type of the owner. Valid values are user or team.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) The email of the user or the name of the team. ### access_control_list`,
				},
				resource.Attribute{
					Name:        "entity_type",
					Description: `(Required) The type of entity. Valid values are user or team.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Required) The permission. Valid values are admin, edit, view.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The email of the user or the name of the team.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_eks_registered",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) Enter the cluster type to be applied for the cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) This field indicates the labels to be set on the cluster.`,
				},
				resource.Attribute{
					Name:        "delete_action",
					Description: `(Optional) This field indicates whether to delete or remove the cluster on destroy. The default value is ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) This field indicates the url of the kubernetes cluster API server.`,
				},
				resource.Attribute{
					Name:        "owner_info",
					Description: `(Optional) The [owner_info](#owner_info) for this cluster, if it has to be overridden.`,
				},
				resource.Attribute{
					Name:        "access_control_list",
					Description: `(Optional) List of additional [ACLs](#access_control_list) for this cluster. ### owner_info`,
				},
				resource.Attribute{
					Name:        "owner_type",
					Description: `(Required) The type of the owner. Valid values are user or team.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) The email of the user or the name of the team. ### access_control_list`,
				},
				resource.Attribute{
					Name:        "entity_type",
					Description: `(Required) The type of entity. Valid values are user or team.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Required) The permission. Valid values are admin, edit, view.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The email of the user or the name of the team.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_gke_imported",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) Enter the cloud credentials to locate and import the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) Enter the cluster type to apply.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Enter the region where the cluster is located.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) Enter the type of cluster (gke/eks).`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Required) Enter the project where the cluster is located.`,
				},
				resource.Attribute{
					Name:        "delete_action",
					Description: `(Optional) This field indicates whether to delete or remove the cluster on destroy. The default value is set to ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_metadata",
					Description: `(Optional) This field indicates the key-value pairs that will be provisioned as a ConfigMap called system-metadata-map in the cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) This field indicates the labels set on cluster.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) This field indicates the url of the kubernetes cluster API server.`,
				},
				resource.Attribute{
					Name:        "owner_info",
					Description: `(Optional) The [owner_info](#owner_info) for this cluster, if it has to be overridden.`,
				},
				resource.Attribute{
					Name:        "access_control_list",
					Description: `(Optional) List of additional [ACLs](#access_control_list) for this cluster. ### owner_info`,
				},
				resource.Attribute{
					Name:        "owner_type",
					Description: `(Required) The type of the owner. Valid values are user or team.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) The email of the user or the name of the team. ### access_control_list`,
				},
				resource.Attribute{
					Name:        "entity_type",
					Description: `(Required) The type of entity. Valid values are user or team.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Required) The permission. Valid values are admin, edit, view.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The email of the user or the name of the team. ### vault_auth`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for vault authentication.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Enter the vault authentication path. The variable $(cluster.name) is allowed in the path for uniquenes.`,
				},
				resource.Attribute{
					Name:        "addon_name",
					Description: `(Required) Enter the associated Vault Agent Injector add-on.`,
				},
				resource.Attribute{
					Name:        "credentials_name",
					Description: `(Required) Enter the Vault credentials to be used.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) Enter a list of application roles to configure for the add-on services.`,
				},
				resource.Attribute{
					Name:        "delete_auth_path",
					Description: `(Optional) This field indicates the delete authentication path on cluster delete. #### roles`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for roles.`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Required) Enter the allowed service account name.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Enter the allowed namespace.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) Enter the applied policies.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_gke_registered",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) Enter the cluster type to apply.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) This field indicates the labels to be set on cluster.`,
				},
				resource.Attribute{
					Name:        "delete_action",
					Description: `(Optional) This field indicates whether to delete or remove the cluster on destroy. The default value is to ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) This field indicates the url of the kubernetes cluster API server.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_cluster_kind_registered",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) Enter the cluster type to be applied to the cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) This field indicates the labels to be set on the cluster.`,
				},
				resource.Attribute{
					Name:        "delete_action",
					Description: `(Optional) This field indicates whether to delete or remove the cluster on destroy. The default value is ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) This field indicates the url of the kubernetes cluster API server.`,
				},
				resource.Attribute{
					Name:        "owner_info",
					Description: `(Optional) The [owner_info](#owner_info) for this cluster, if it has to be overridden.`,
				},
				resource.Attribute{
					Name:        "access_control_list",
					Description: `(Optional) List of additional [ACLs](#access_control_list) for this cluster.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `clusters.cluster.server.`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `users.user.client-certificate.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `users.user.client-key.`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `clusters.cluster.certificate-authority-data. ### owner_info`,
				},
				resource.Attribute{
					Name:        "owner_type",
					Description: `(Required) The type of the owner. Valid values are user or team.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) The email of the user or the name of the team. ### access_control_list`,
				},
				resource.Attribute{
					Name:        "entity_type",
					Description: `(Required) The type of entity. Valid values are user or team.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Required) The permission. Valid values are admin, edit, view.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The email of the user or the name of the team.`,
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
					Description: `(Required) Enter a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Enter the EKS version (example, 1.18).`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) Enter the cloud credentials that is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Enter the AWS region for the cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Enter the AWS VPC ID for the cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Enter a list of AWS VPC subnets to be used for the cluster.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) Enter a list of AWS VPC security groups to be used for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_role_arn",
					Description: `(Required) Enter the cluster role ARN.`,
				},
				resource.Attribute{
					Name:        "enable_private_endpoint",
					Description: `(Optional) This field indicates if the cluster API server endpoint should be in a private network.`,
				},
				resource.Attribute{
					Name:        "enable_identity_provider",
					Description: `(Optional) This value enables IAM roles for service accounts.`,
				},
				resource.Attribute{
					Name:        "auto_sync_namespaces",
					Description: `(Optional) This value enables automatic synchronization of cluster namespaces for Nirmata.`,
				},
				resource.Attribute{
					Name:        "enable_secrets_encryption",
					Description: `(Optional) This value enables encryption at REST for secrets.`,
				},
				resource.Attribute{
					Name:        "kms_key_arn",
					Description: `(Optional) This field indicates the KMS key ARN to be used for the secrets encryption.`,
				},
				resource.Attribute{
					Name:        "log_types",
					Description: `(Optional) This field indicates the log types to be collected.`,
				},
				resource.Attribute{
					Name:        "enable_fargate",
					Description: `(Optional) This value enables Fargate to provision nodes based on the workload resource requests.`,
				},
				resource.Attribute{
					Name:        "system_metadata",
					Description: `(Optional) This is the key-value pairs that will be provisioned as a ConfigMap called system-metadata-map in the cluster.`,
				},
				resource.Attribute{
					Name:        "allow_override_credentials",
					Description: `(Optional) This allows the passing of cloud credentials when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "cluster_field_override",
					Description: `(Optional) This allows the override of cluster settings ('network' and 'subnetwork') when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "nodepool_field_override",
					Description: `(Optional) This allows the override of node fields ('disk_size' and 'machine_type') when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "nodepools",
					Description: `(Optional) This field indicates a list of [nodepool](#nodepool) types.`,
				},
				resource.Attribute{
					Name:        "addons",
					Description: `(Optional) This field indicates a list of add-on services.`,
				},
				resource.Attribute{
					Name:        "vault_auth",
					Description: `(Optional) This field indicates the vault authentication configuration. ## Nested Blocks ### nodepool`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the node pool.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Enter the EC2 instance type (example, "t3.medium").`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) Enter the worker node disk size in GB (example, 60).`,
				},
				resource.Attribute{
					Name:        "ssh_key_name",
					Description: `(Required) Enter the SSK key pair to access nodes.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) Enter the node security groups.`,
				},
				resource.Attribute{
					Name:        "iam_role",
					Description: `(Required) Enter the IAM role to be used for nodes.`,
				},
				resource.Attribute{
					Name:        "ami_type",
					Description: `(Required) Enter the EKS-optimized Amazon Machine Image (AMI) type to be used for node images.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) Enter the Amazon Machine Image (AMI) IS to be used for node images.`,
				},
				resource.Attribute{
					Name:        "node_annotations",
					Description: `(Optional) This value indicates the annotations to be set on each node in this pool. This setting is permanent.`,
				},
				resource.Attribute{
					Name:        "node_labels",
					Description: `(Optional) This value indicates the labels to be set on each Kubernetes node in this node pool. This setting is permanent. ### addons`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the add-on service.`,
				},
				resource.Attribute{
					Name:        "addon_selector",
					Description: `(Required) Enter the catalog application name.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) Enter the catalog name.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) Enter the channel from which the application should be deployed.`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `(Optional) This feild indicates a sequence number that control installation order. ### vault_auth`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the vault authentication.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Enter the vault authentication path. The variable $(cluster.name) is allowed for uniquenes.`,
				},
				resource.Attribute{
					Name:        "addon_name",
					Description: `(Required) Enter the associated Vault Agent Injector add-on.`,
				},
				resource.Attribute{
					Name:        "credentials_name",
					Description: `(Required) Enter the Vault credentials to be used for the authentication.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) Enter a list of application roles to be configured for the add-on services.`,
				},
				resource.Attribute{
					Name:        "delete_auth_path",
					Description: `(Optional) This field indicates the delete authentication path on cluster delete. #### roles`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for roles.`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Required) Enter the allowed service account name.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Enter the allowed namespace.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) Enter the applied policies.`,
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
					Description: `(Required) Enter a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Enter the GKE version (example, 1.16.12-gke.3)`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) Enter the cloud credentials to be used for the cluster.`,
				},
				resource.Attribute{
					Name:        "location_type",
					Description: `(Required) Enter the location type as Regional or Zonal. A regional cluster has multiple replicas of the control plane running in multiple zones within a given region. A zonal cluster has a single replica of the control plane running in a single zone.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) This field indicates the GCP region into which the cluster should be deployed (e.g. "us-central1"). This value is required when location_type is ` + "`" + `Regional` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) This field indicates the GCP zone into which the cluster should be deployed (example, "us-central1-a"). This value is required if location_type is ` + "`" + `Zonal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_locations",
					Description: `(Optional) This field indicates the nodes that should be deployed. Selecting more than one zone increases availability. (example, ["asia-east1-a"]). This value is required if location_type is ` + "`" + `Regional` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_secrets_encryption",
					Description: `(Optional) This value enables envelope encryption for Kubernetes Secrets.`,
				},
				resource.Attribute{
					Name:        "enable_workload_identity",
					Description: `(Optional) This field indicates that the Workload Identity is the recommended way to access Google Cloud services from an applications running within GKE, due to its improved security properties and manageability.`,
				},
				resource.Attribute{
					Name:        "workload_pool",
					Description: `(Optional) This field indicates that the Workload Identity relies on a Workload Pool to aggregate identity across multiple clusters. This value is required if enable_secrets_encryption is set to true.`,
				},
				resource.Attribute{
					Name:        "secrets_encryption_key",
					Description: `(Optional) This field indicates that the Resource ID of the key you want to use (example, projects/project-name/locations/global/keyRings/my-keyring/cryptoKeys/my-key). This value is required if enable_workload_identity is set to true.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Enter the cluster network (example, "default")`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Required) Enter the node subnetwork (example, "default")`,
				},
				resource.Attribute{
					Name:        "cluster_ipv4_cidr",
					Description: `(Optional) This field indicates that all pods in the cluster are assigned an IP address from this range. Enter a range (in CIDR notation) within a network range, a mask, or leave this field blank to use a default range. This setting is permanent.`,
				},
				resource.Attribute{
					Name:        "services_ipv4_cidr",
					Description: `(Optional) This field indicates that the cluster services will be assigned an IP address from this IP address range. Enter a range (in CIDR notation) within a network range, a mask, or leave this field blank to use a default range. This setting is permanent.`,
				},
				resource.Attribute{
					Name:        "cloud_run",
					Description: `(Optional) The Cloud Run for Anthos enables you to easily deploy stateless apps and functions to the cluster using the Cloud Run experience. Cloud Run for Anthos automatically manages underlying resources and scales your app, based on requests.`,
				},
				resource.Attribute{
					Name:        "enable_network_policy",
					Description: `(Optional) The Kubernetes Network Policy API allows the cluster administrator to specify what pods are allowed to communicate with each other. Google Kubernetes Engine has partnered with Tigera to provide Project Calico to enforce network policies within your cluster.`,
				},
				resource.Attribute{
					Name:        "enable_http_load_balancing",
					Description: `(Optional) The HTTP Load Balancing add-on is required to use the Google Cloud Load Balancer with Kubernetes Ingress. It enables, a controller install to coordinate applying load balancing configuration changes to your GCP project.`,
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
					Description: `(Optional) This field specifies regular time for maintenance and enables maintenance windows. Normally, routine Kubernetes Engine maintenance may run at any time on your cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_start_time",
					Description: `(Optional) This field indicates the start time for the maintenance window.`,
				},
				resource.Attribute{
					Name:        "maintenance_duration",
					Description: `(Optional) This field indicates the duration for the maintenance window in hours.`,
				},
				resource.Attribute{
					Name:        "maintenance_recurrence",
					Description: `(Optional) This field indicates the Recurrence Rule specification (RRULE) for the maintenance window. Example RRule to run maintenance during weekends: 'FREQ=WEEKLY;BYDAY=SA,SU'.`,
				},
				resource.Attribute{
					Name:        "maintenance_exclusion_timewindow",
					Description: `(Optional) This field specifies the time when routine, non-emergency maintenance will not happen. It is set up to 3 maintenance exclusions. Normally, routine Kubernetes Engine maintenance may run at any time on your cluster.`,
				},
				resource.Attribute{
					Name:        "system_metadata",
					Description: `(Optional) This field indicates that the key-value pairs that will be provisioned as a ConfigMap called system-metadata-map in the cluster.`,
				},
				resource.Attribute{
					Name:        "allow_override_credentials",
					Description: `(Optional) This value allows passing in cloud credentials when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "cluster_field_override",
					Description: `(Optional) This value allows override of cluster settings ('network' and 'subnetwork') when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "nodepool_field_override",
					Description: `(Optional) This value allows override of node fields ('disk_size' and 'machine_type') when a cluster is created using this cluster type.`,
				},
				resource.Attribute{
					Name:        "nodepools",
					Description: `(Optional) This field indicates a list of [nodepool](#nodepool) types.`,
				},
				resource.Attribute{
					Name:        "addons",
					Description: `(Optional) This field indicates a list of add-on services.`,
				},
				resource.Attribute{
					Name:        "vault_auth",
					Description: `(Optional) This field indicates the vault authentication configuration. ## Nested Blocks ### nodepool`,
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
					Description: `(Required) Enter a unique name for the add-on service.`,
				},
				resource.Attribute{
					Name:        "addon_selector",
					Description: `(Required) Enter the catalog application name.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) Enter the catalog name.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) Enter the channel from which the application should be deployed.`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `(Optional) This field indicates a sequence number to control the installation order. ### vault_auth`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the vault authentication.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Enter the vault authentication path. The variable $(cluster.name) is allowed in the path for uniquenes.`,
				},
				resource.Attribute{
					Name:        "addon_name",
					Description: `(Required) Enter the associated Vault Agent Injector add-on.`,
				},
				resource.Attribute{
					Name:        "credentials_name",
					Description: `(Required) Enter the Vault credentials to use.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) Enter a list of application roles to configure for the add-on services.`,
				},
				resource.Attribute{
					Name:        "delete_auth_path",
					Description: `(Optional) This field indicates the delete authentication path on cluster delete. #### roles`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for roles.`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Required) Enter the allowed service account name.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Enter the allowed namespace.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) Enter the applied policies.`,
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
					Description: `(Required) Enter a unique name for the cluster.`,
				},
				resource.Attribute{
					Name:        "cloud",
					Description: `(Optional) Enter the cloud provider. Defaults to ` + "`" + `Other` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "addons",
					Description: `(Optional) Enter a list of add-on services.`,
				},
				resource.Attribute{
					Name:        "vault_auth",
					Description: `(Optional) Enter the vault authentication configuration.`,
				},
				resource.Attribute{
					Name:        "system_metadata",
					Description: `(Optional) The key-value pairs that will be provisioned as a ConfigMap called system-metadata-map in the cluster type. ## Nested Blocks ### addons`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the add-on service.`,
				},
				resource.Attribute{
					Name:        "addon_selector",
					Description: `(Required) Enter the catalog application name.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) Enter the catalog name.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) Enter the channel from which the application should be deployed.`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `(Optional) This field indicates a sequence number to control the installation order. ### vault_auth`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Enter the vault authentication path. The variable $(cluster.name) is allowed in the path for uniquenes.`,
				},
				resource.Attribute{
					Name:        "addon_name",
					Description: `(Required) Enter the associated Vault Agent Injector add-on.`,
				},
				resource.Attribute{
					Name:        "credentials_name",
					Description: `(Required) Enter the Vault credentials to be used.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) Enter a list of application roles to configure for add-on services.`,
				},
				resource.Attribute{
					Name:        "delete_auth_path",
					Description: `(Optional) This field indicates the delete authentication path on cluster delete. #### roles`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Required) Enter the allowed service account name.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Ener the allowed namespace.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) Enter the applied policies.`,
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
					Description: `(Required) Enter a unique name for the environment.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Enter the environnment type.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) Enter the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) This field indicates the labels to be set on the add-on application's environment.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) This field indicates the cluster namespace bound to this environment. It defaults to the environment name.`,
				},
				resource.Attribute{
					Name:        "environment_update_action",
					Description: `(Optional) By default, this value is set to notify and to update if changes need to apply automatically.`,
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
					Description: `(Required) Enter a unique name for the environment type.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) This field uses the default environment type.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) This field indicates that the labels to set on the add-on application's environment type.`,
				},
				resource.Attribute{
					Name:        "resource_limits",
					Description: `(Required) Enter a map of resource limits for the environment. The commonly used resources include ` + "`" + `cpu` + "`" + `, ` + "`" + `memory` + "`" + `, and ` + "`" + `storage` + "`" + `. Refer the [Kubernetes docs](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) for a complete reference.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nirmata_policy_set",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter a unique name for the policy set.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) This field indicates that the policy is set as default. The default policy set will be automatically deployed to new clusters.`,
				},
				resource.Attribute{
					Name:        "git_credentials",
					Description: `(Required) Enter the git credential name.`,
				},
				resource.Attribute{
					Name:        "git_repository",
					Description: `(Required) Enter the repository URL as used in the git clone command.`,
				},
				resource.Attribute{
					Name:        "git_branch",
					Description: `(Required) Enter the Git branch to track. It indicates the name.`,
				},
				resource.Attribute{
					Name:        "git_directory_list",
					Description: `(Optional) This field indicates the directories to track.`,
				},
				resource.Attribute{
					Name:        "fixed_kustomization",
					Description: `(Optional) This field enables fixed kustomize to select kustomizations for your application.`,
				},
				resource.Attribute{
					Name:        "target_based_kustomization",
					Description: `(Optional) This field enables target based kustomize to select kustomizations for your application.`,
				},
				resource.Attribute{
					Name:        "kustomization_file_path",
					Description: `(Required if fixed_kustomization or target_based_kustomization is set) Enter the kustomization file path. kustomization file path is required if you select fixed_kustomization or target_based_kustomization.`,
				},
				resource.Attribute{
					Name:        "delete_from_cluster",
					Description: `(Optional) This field indicates the delete from cluster. # nirmata_deploy_policy_set Resource ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl resource "nirmata_deploy_policy_set" "tf-policy-set-deploy" { policy_set_name = "policy-set-name" cluster = "cluster-name" delete_from_cluster = true depends_on = [nirmata_policy_set.create-policy-set] } ` + "`" + `` + "`" + `` + "`" + ` ## Argument Reference`,
				},
				resource.Attribute{
					Name:        "policy_set_name",
					Description: `(Required) Enter deploy policy set name.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) Enter the cluster name in which the policy is set to deploy.`,
				},
				resource.Attribute{
					Name:        "delete_from_cluster",
					Description: `(Optional) This field indicates the delete from cluster.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"nirmata_aws_cloud_credentials":   0,
		"nirmata_catalog":                 1,
		"nirmata_catalog_git_app":         2,
		"nirmata_catalog_helm_app":        3,
		"nirmata_catalog_promote_version": 4,
		"nirmata_catalog_run_app":         5,
		"nirmata_catalog_yaml_app":        6,
		"nirmata_cluster":                 7,
		"nirmata_cluster_addon":           8,
		"nirmata_cluster_aks_registered":  9,
		"nirmata_cluster_eks_imported":    10,
		"nirmata_cluster_eks_registered":  11,
		"nirmata_cluster_gke_imported":    12,
		"nirmata_cluster_gke_registered":  13,
		"nirmata_cluster_kind_registered": 14,
		"nirmata_cluster_type_eks":        15,
		"nirmata_cluster_type_gke":        16,
		"nirmata_cluster_type_registered": 17,
		"nirmata_environment":             18,
		"nirmata_environment_type":        19,
		"nirmata_policy_set":              20,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
