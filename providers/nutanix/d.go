package nutanix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nutanix_access_control_policies",
			Category:         "Data Sources",
			ShortDescription: `Describes a list of access control policies`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "entity_filter_expression_list",
					Description: `A list of Entity filter expressions. ### Scope Filter Expression List The scope_filter_expression_list attribute supports the following.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_access_control_policy",
			Category:         "Data Sources",
			ShortDescription: `Describes an Access Control Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "entity_filter_expression_list",
					Description: `A list of Entity filter expressions. ### Scope Filter Expression List The scope_filter_expression_list attribute supports the following.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "entity_filter_expression_list",
					Description: `A list of Entity filter expressions. ### Scope Filter Expression List The scope_filter_expression_list attribute supports the following.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_category_key",
			Category:         "Data Sources",
			ShortDescription: `Describe a Nutanix Category Key and its values (if it has them).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_cluster",
			Category:         "Data Sources",
			ShortDescription: `Describes a Cluster`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The API version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The API version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_clusters",
			Category:         "Data Sources",
			ShortDescription: `Describes a Clusters`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The API version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The API version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_host",
			Category:         "Data Sources",
			ShortDescription: `Describes a host`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The API version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The API version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_hosts",
			Category:         "Data Sources",
			ShortDescription: `Describes a list of hosts`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The API version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_image",
			Category:         "Data Sources",
			ShortDescription: `Describes a Image`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_karbon_cluster",
			Category:         "Data Sources",
			ShortDescription: `Describes a Karbon Cluster`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "reclaim_policy",
					Description: `Reclaim policy for persistent volumes provisioned using the specified storage class.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.file_system",
					Description: `Karbon uses either the ext4 or xfs file-system on the volume disk.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.flash_mode",
					Description: `Pins the persistent volumes to the flash tier in case of a ` + "`" + `true` + "`" + ` value.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.password",
					Description: `The password of the Prism Element user that the API calls use to provision volumes.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.prism_element_cluster_uuid",
					Description: `The universally unique identifier (UUID) of the Prism Element cluster.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.storage_container",
					Description: `Name of the storage container the storage container uses to provision volumes.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.username",
					Description: `Username of the Prism Element user that the API calls use to provision volumes. ### External LB Config The external load balancer configuration in the case of a multi-master-external-load-balancer type master deployment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "reclaim_policy",
					Description: `Reclaim policy for persistent volumes provisioned using the specified storage class.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.file_system",
					Description: `Karbon uses either the ext4 or xfs file-system on the volume disk.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.flash_mode",
					Description: `Pins the persistent volumes to the flash tier in case of a ` + "`" + `true` + "`" + ` value.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.password",
					Description: `The password of the Prism Element user that the API calls use to provision volumes.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.prism_element_cluster_uuid",
					Description: `The universally unique identifier (UUID) of the Prism Element cluster.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.storage_container",
					Description: `Name of the storage container the storage container uses to provision volumes.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.username",
					Description: `Username of the Prism Element user that the API calls use to provision volumes. ### External LB Config The external load balancer configuration in the case of a multi-master-external-load-balancer type master deployment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_karbon_cluster_kubeconfig",
			Category:         "Data Sources",
			ShortDescription: `Describes the SSH config from a Karbon Cluster`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_karbon_cluster_ssh",
			Category:         "Data Sources",
			ShortDescription: `Describes the SSH config from a Karbon Cluster`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_karbon_clusters",
			Category:         "Data Sources",
			ShortDescription: `Describes a Karbon Clusters`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "reclaim_policy",
					Description: `Reclaim policy for persistent volumes provisioned using the specified storage class.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.file_system",
					Description: `Karbon uses either the ext4 or xfs file-system on the volume disk.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.password",
					Description: `TPassword of the Prism Element user that the API calls use to provision volumes.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.prism_element_cluster_uuid",
					Description: `The universally unique identifier (UUID) of the Prism Element cluster.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.storage_container",
					Description: `TName of the storage container the storage container uses to provision volumes.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.username",
					Description: `TThe username of the Prism Element user that the API calls use to provision volumes. ### External LB Config The external load balancer configuration in the case of a multi-master-external-load-balancer type master deployment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "reclaim_policy",
					Description: `Reclaim policy for persistent volumes provisioned using the specified storage class.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.file_system",
					Description: `Karbon uses either the ext4 or xfs file-system on the volume disk.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.password",
					Description: `TPassword of the Prism Element user that the API calls use to provision volumes.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.prism_element_cluster_uuid",
					Description: `The universally unique identifier (UUID) of the Prism Element cluster.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.storage_container",
					Description: `TName of the storage container the storage container uses to provision volumes.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.username",
					Description: `TThe username of the Prism Element user that the API calls use to provision volumes. ### External LB Config The external load balancer configuration in the case of a multi-master-external-load-balancer type master deployment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_karbon_private_registries",
			Category:         "Data Sources",
			ShortDescription: `Describes a List of Karbon private registry entry`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_karbon_private_registry",
			Category:         "Data Sources",
			ShortDescription: `Describes a Karbon private registry entry`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_network_security_rule",
			Category:         "Data Sources",
			ShortDescription: `Describes a Network security rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_security_rule_id",
					Description: `(Required) The ID for the rule you want to retrieve.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_security_rule_id",
					Description: `(Required) The ID for the rule you want to retrieve.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_permission",
			Category:         "Data Sources",
			ShortDescription: `Describe a Nutanix Permission and its values (if it has them).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "last_update_time",
					Description: `UTC date and time in RFC-3339 format when the permission was last updated.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Permission UUID.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `UTC date and time in RFC-3339 format when the permission was created.`,
				},
				resource.Attribute{
					Name:        "spec_version",
					Description: `Version number of the latest spec.`,
				},
				resource.Attribute{
					Name:        "spec_hash",
					Description: `Hash of the spec. This will be returned from server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Permission name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the key name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `value of the key. ### Reference The ` + "`" + `project_reference` + "`" + `, ` + "`" + `owner_reference` + "`" + ` attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind name (Default value: ` + "`" + `project` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) the UUID. See detailed information in [Nutanix Permission](https://www.nutanix.dev/reference/prism_central/v3/api/permissions/getpermissionsuuid/).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "last_update_time",
					Description: `UTC date and time in RFC-3339 format when the permission was last updated.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Permission UUID.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `UTC date and time in RFC-3339 format when the permission was created.`,
				},
				resource.Attribute{
					Name:        "spec_version",
					Description: `Version number of the latest spec.`,
				},
				resource.Attribute{
					Name:        "spec_hash",
					Description: `Hash of the spec. This will be returned from server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Permission name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the key name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `value of the key. ### Reference The ` + "`" + `project_reference` + "`" + `, ` + "`" + `owner_reference` + "`" + ` attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind name (Default value: ` + "`" + `project` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) the UUID. See detailed information in [Nutanix Permission](https://www.nutanix.dev/reference/prism_central/v3/api/permissions/getpermissionsuuid/).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_permissions",
			Category:         "Data Sources",
			ShortDescription: `Provides a datasource to retrieve all the permissions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "last_update_time",
					Description: `UTC date and time in RFC-3339 format when the permission was last updated.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `permission UUID.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `UTC date and time in RFC-3339 format when the permission was created.`,
				},
				resource.Attribute{
					Name:        "spec_version",
					Description: `Version number of the latest spec.`,
				},
				resource.Attribute{
					Name:        "spec_hash",
					Description: `Hash of the spec. This will be returned from server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `permission name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the key name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `value of the key. ### Reference The ` + "`" + `project_reference` + "`" + `, ` + "`" + `owner_reference` + "`" + ` attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind name (Default value: ` + "`" + `project` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) the UUID. See detailed information in [Nutanix Permissions](https://www.nutanix.dev/reference/prism_central/v3/api/permissions/postpermissionslist/).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_project",
			Category:         "Data Sources",
			ShortDescription: `Describe a Nutanix Project and its values (if it has them).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_domain.resources.#.units",
					Description: `The units of the resource type`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.value",
					Description: `The amount of resource consumed ### Account Reference List`,
				},
				resource.Attribute{
					Name:        "account_reference_list",
					Description: `List of accounts associated with the project.`,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `account` + "`" + ``,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.uuid",
					Description: `The UUID of an account.`,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.name",
					Description: `The name of an account. ### Environment Reference List`,
				},
				resource.Attribute{
					Name:        "environment_reference_list",
					Description: `List of environments associated with the project.`,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `environment` + "`" + ``,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.uuid",
					Description: `The UUID of an environment.`,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.name",
					Description: `The name of an environment. ### Default Subnet Reference Map`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference",
					Description: `Reference to a subnet.`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.kind",
					Description: `The kind name. Default value is ` + "`" + `subnet` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.uuid",
					Description: `The UUID of a subnet.`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.name",
					Description: `The name of a subnet. ### user_reference_list`,
				},
				resource.Attribute{
					Name:        "user_reference_list",
					Description: `List of users in the project.`,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `user` + "`" + ``,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.uuid",
					Description: `The UUID of a user`,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.name",
					Description: `The name of a user. ### External User Group Reference List`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list",
					Description: `List of directory service user groups. These groups are not managed by Nutanix.`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `user_group` + "`" + ``,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.uuid",
					Description: `The UUID of a user_group`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.name",
					Description: `The name of a user_group ### Subnet Reference List`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list",
					Description: `List of subnets for the project.`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `subnet` + "`" + ``,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.uuid",
					Description: `The UUID of a subnet`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.name",
					Description: `The name of a subnet. ### External Network List`,
				},
				resource.Attribute{
					Name:        "external_network_list",
					Description: `List of external networks associated with the project.`,
				},
				resource.Attribute{
					Name:        "external_network_list.#.uuid",
					Description: `The UUID of a network.`,
				},
				resource.Attribute{
					Name:        "external_network_list.#.name",
					Description: `The name of a network. ### Resource Domain`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.units",
					Description: `The units of the resource type`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.value",
					Description: `The amount of resource consumed ### Metadata The metadata attribute exports the following:`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `UTC date and time in RFC-3339 format when the project was last updated.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Project UUID.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `UTC date and time in RFC-3339 format when the project was created.`,
				},
				resource.Attribute{
					Name:        "spec_version",
					Description: `Version number of the latest spec.`,
				},
				resource.Attribute{
					Name:        "spec_hash",
					Description: `Hash of the spec. This will be returned from server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Project name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the key name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `value of the key. ### Reference The ` + "`" + `project_reference` + "`" + `, ` + "`" + `owner_reference` + "`" + ` attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind name (Default value: ` + "`" + `project` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) the UUID. See detailed information in [Nutanix Project](https://www.nutanix.dev/reference/prism_central/v3/api/projects/getprojectsuuid/).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_domain.resources.#.units",
					Description: `The units of the resource type`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.value",
					Description: `The amount of resource consumed ### Account Reference List`,
				},
				resource.Attribute{
					Name:        "account_reference_list",
					Description: `List of accounts associated with the project.`,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `account` + "`" + ``,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.uuid",
					Description: `The UUID of an account.`,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.name",
					Description: `The name of an account. ### Environment Reference List`,
				},
				resource.Attribute{
					Name:        "environment_reference_list",
					Description: `List of environments associated with the project.`,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `environment` + "`" + ``,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.uuid",
					Description: `The UUID of an environment.`,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.name",
					Description: `The name of an environment. ### Default Subnet Reference Map`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference",
					Description: `Reference to a subnet.`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.kind",
					Description: `The kind name. Default value is ` + "`" + `subnet` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.uuid",
					Description: `The UUID of a subnet.`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.name",
					Description: `The name of a subnet. ### user_reference_list`,
				},
				resource.Attribute{
					Name:        "user_reference_list",
					Description: `List of users in the project.`,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `user` + "`" + ``,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.uuid",
					Description: `The UUID of a user`,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.name",
					Description: `The name of a user. ### External User Group Reference List`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list",
					Description: `List of directory service user groups. These groups are not managed by Nutanix.`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `user_group` + "`" + ``,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.uuid",
					Description: `The UUID of a user_group`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.name",
					Description: `The name of a user_group ### Subnet Reference List`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list",
					Description: `List of subnets for the project.`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `subnet` + "`" + ``,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.uuid",
					Description: `The UUID of a subnet`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.name",
					Description: `The name of a subnet. ### External Network List`,
				},
				resource.Attribute{
					Name:        "external_network_list",
					Description: `List of external networks associated with the project.`,
				},
				resource.Attribute{
					Name:        "external_network_list.#.uuid",
					Description: `The UUID of a network.`,
				},
				resource.Attribute{
					Name:        "external_network_list.#.name",
					Description: `The name of a network. ### Resource Domain`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.units",
					Description: `The units of the resource type`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.value",
					Description: `The amount of resource consumed ### Metadata The metadata attribute exports the following:`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `UTC date and time in RFC-3339 format when the project was last updated.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Project UUID.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `UTC date and time in RFC-3339 format when the project was created.`,
				},
				resource.Attribute{
					Name:        "spec_version",
					Description: `Version number of the latest spec.`,
				},
				resource.Attribute{
					Name:        "spec_hash",
					Description: `Hash of the spec. This will be returned from server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Project name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the key name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `value of the key. ### Reference The ` + "`" + `project_reference` + "`" + `, ` + "`" + `owner_reference` + "`" + ` attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind name (Default value: ` + "`" + `project` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) the UUID. See detailed information in [Nutanix Project](https://www.nutanix.dev/reference/prism_central/v3/api/projects/getprojectsuuid/).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_projects",
			Category:         "Data Sources",
			ShortDescription: `Describes a Projects`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_domain.resources.#.units",
					Description: `The units of the resource type`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.value",
					Description: `The amount of resource consumed ### Account Reference List`,
				},
				resource.Attribute{
					Name:        "account_reference_list",
					Description: `List of accounts associated with the project.`,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `account` + "`" + ``,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.uuid",
					Description: `The UUID of an account.`,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.name",
					Description: `The name of an account. ### Environment Reference List`,
				},
				resource.Attribute{
					Name:        "environment_reference_list",
					Description: `List of environments associated with the project.`,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `environment` + "`" + ``,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.uuid",
					Description: `The UUID of an environment.`,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.name",
					Description: `The name of an environment. ### Default Subnet Reference Map`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference",
					Description: `Reference to a subnet.`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.kind",
					Description: `The kind name. Default value is ` + "`" + `subnet` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.uuid",
					Description: `The UUID of a subnet.`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.name",
					Description: `The name of a subnet. ### user_reference_list`,
				},
				resource.Attribute{
					Name:        "user_reference_list",
					Description: `List of users in the project.`,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `user` + "`" + ``,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.uuid",
					Description: `The UUID of a user`,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.name",
					Description: `The name of a user. ### External User Group Reference List`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list",
					Description: `List of directory service user groups. These groups are not managed by Nutanix.`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `user_group` + "`" + ``,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.uuid",
					Description: `The UUID of a user_group`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.name",
					Description: `The name of a user_group ### Subnet Reference List`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list",
					Description: `List of subnets for the project.`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `subnet` + "`" + ``,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.uuid",
					Description: `The UUID of a subnet`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.name",
					Description: `The name of a subnet. ### External Network List`,
				},
				resource.Attribute{
					Name:        "external_network_list",
					Description: `List of external networks associated with the project.`,
				},
				resource.Attribute{
					Name:        "external_network_list.#.uuid",
					Description: `The UUID of a network.`,
				},
				resource.Attribute{
					Name:        "external_network_list.#.name",
					Description: `The name of a network. ### Resource Domain`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.units",
					Description: `The units of the resource type`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.value",
					Description: `The amount of resource consumed ### Metadata The metadata attribute exports the following:`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `UTC date and time in RFC-3339 format when the project was last updated.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Project UUID.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `UTC date and time in RFC-3339 format when the project was created.`,
				},
				resource.Attribute{
					Name:        "spec_version",
					Description: `Version number of the latest spec.`,
				},
				resource.Attribute{
					Name:        "spec_hash",
					Description: `Hash of the spec. This will be returned from server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Project name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the key name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `value of the key. ### Reference The ` + "`" + `project_reference` + "`" + `, ` + "`" + `owner_reference` + "`" + ` attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind name (Default value: ` + "`" + `project` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) the name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) the UUID. See detailed information in [Nutanix Projects](https://www.nutanix.dev/reference/prism_central/v3/api/projects/postprojectslist).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_protection_rule",
			Category:         "Data Sources",
			ShortDescription: `Describe a Nutanix Protection Rule and its values (if it has them).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone_connectivity_list",
					Description: `(Required) This encodes the datapipes between various availability zones and\nthe backup policy of the pipes.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.destination_availability_zone_index",
					Description: `(Optional/Computed) Index of the availability zone.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.source_availability_zone_index",
					Description: `(Optional/Computed) Index of the availability zone.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list",
					Description: `(Optional/Computed) Snapshot schedules for the pair of the availability zones.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.recovery_point_objective_secs",
					Description: `(Required) "A recovery point objective (RPO) is the maximum acceptable amount of data loss.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy",
					Description: `(Optional/Computed) This describes the snapshot retention policy for this availability zone.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy.0.num_snapshots",
					Description: `(Optional/Computed) Number of snapshots need to be retained.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy.0.rollup_retention_policy_multiple",
					Description: `(Optional/Computed) Multiplier to 'snapshot_interval_type'.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy.0.rollup_retention_policy_snapshot_interval_type",
					Description: `(Optional/Computed)`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.auto_suspend_timeout_secs",
					Description: `(Optional/Computed) Auto suspend timeout in case of connection failure between the sites.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.snapshot_type",
					Description: `(Optional/Computed) Crash consistent or Application Consistent snapshot.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.remote_snapshot_retention_policy",
					Description: `(Optional/Computed) This describes the snapshot retention policy for this availability zone. ### Ordered Availability Zone List`,
				},
				resource.Attribute{
					Name:        "ordered_availability_zone_list",
					Description: `(Required) A list of availability zones, each of which, receives a replica\nof the data for the entities protected by this protection rule.`,
				},
				resource.Attribute{
					Name:        "ordered_availability_zone_list.#.cluster_uuid",
					Description: `(Optional/Computed) UUID of specific cluster to which we will be replicating.`,
				},
				resource.Attribute{
					Name:        "ordered_availability_zone_list.#.availability_zone_url",
					Description: `(Optional/Computed) The FQDN or IP address of the availability zone. ### Category Filter`,
				},
				resource.Attribute{
					Name:        "category_filter",
					Description: `(Optional/Computed)`,
				},
				resource.Attribute{
					Name:        "category_filter.0.type",
					Description: `(Optional/Computed) The type of the filter being used.`,
				},
				resource.Attribute{
					Name:        "category_filter.0.kind_list",
					Description: `(Optional/Computed) List of kinds associated with this filter.`,
				},
				resource.Attribute{
					Name:        "category_filter.0.params",
					Description: `(Optional/Computed) A list of category key and list of values. ### Metadata The metadata attribute exports the following:`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `UTC date and time in RFC-3339 format when vm was last updated.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `vm UUID.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `UTC date and time in RFC-3339 format when vm was created.`,
				},
				resource.Attribute{
					Name:        "spec_version",
					Description: `Version number of the latest spec.`,
				},
				resource.Attribute{
					Name:        "spec_hash",
					Description: `Hash of the spec. This will be returned from server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `vm name. ### Categories The categories attribute supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the key name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `value of the key. ### Reference The ` + "`" + `project_reference` + "`" + `, ` + "`" + `owner_reference` + "`" + ` attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind name (Default value: ` + "`" + `project` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) the name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) the UUID. See detailed information in [Nutanix Protection Rule](https://www.nutanix.dev/reference/prism_central/v3/api/protection-rules/getprotectionrulesuuid/).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone_connectivity_list",
					Description: `(Required) This encodes the datapipes between various availability zones and\nthe backup policy of the pipes.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.destination_availability_zone_index",
					Description: `(Optional/Computed) Index of the availability zone.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.source_availability_zone_index",
					Description: `(Optional/Computed) Index of the availability zone.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list",
					Description: `(Optional/Computed) Snapshot schedules for the pair of the availability zones.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.recovery_point_objective_secs",
					Description: `(Required) "A recovery point objective (RPO) is the maximum acceptable amount of data loss.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy",
					Description: `(Optional/Computed) This describes the snapshot retention policy for this availability zone.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy.0.num_snapshots",
					Description: `(Optional/Computed) Number of snapshots need to be retained.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy.0.rollup_retention_policy_multiple",
					Description: `(Optional/Computed) Multiplier to 'snapshot_interval_type'.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy.0.rollup_retention_policy_snapshot_interval_type",
					Description: `(Optional/Computed)`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.auto_suspend_timeout_secs",
					Description: `(Optional/Computed) Auto suspend timeout in case of connection failure between the sites.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.snapshot_type",
					Description: `(Optional/Computed) Crash consistent or Application Consistent snapshot.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.remote_snapshot_retention_policy",
					Description: `(Optional/Computed) This describes the snapshot retention policy for this availability zone. ### Ordered Availability Zone List`,
				},
				resource.Attribute{
					Name:        "ordered_availability_zone_list",
					Description: `(Required) A list of availability zones, each of which, receives a replica\nof the data for the entities protected by this protection rule.`,
				},
				resource.Attribute{
					Name:        "ordered_availability_zone_list.#.cluster_uuid",
					Description: `(Optional/Computed) UUID of specific cluster to which we will be replicating.`,
				},
				resource.Attribute{
					Name:        "ordered_availability_zone_list.#.availability_zone_url",
					Description: `(Optional/Computed) The FQDN or IP address of the availability zone. ### Category Filter`,
				},
				resource.Attribute{
					Name:        "category_filter",
					Description: `(Optional/Computed)`,
				},
				resource.Attribute{
					Name:        "category_filter.0.type",
					Description: `(Optional/Computed) The type of the filter being used.`,
				},
				resource.Attribute{
					Name:        "category_filter.0.kind_list",
					Description: `(Optional/Computed) List of kinds associated with this filter.`,
				},
				resource.Attribute{
					Name:        "category_filter.0.params",
					Description: `(Optional/Computed) A list of category key and list of values. ### Metadata The metadata attribute exports the following:`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `UTC date and time in RFC-3339 format when vm was last updated.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `vm UUID.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `UTC date and time in RFC-3339 format when vm was created.`,
				},
				resource.Attribute{
					Name:        "spec_version",
					Description: `Version number of the latest spec.`,
				},
				resource.Attribute{
					Name:        "spec_hash",
					Description: `Hash of the spec. This will be returned from server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `vm name. ### Categories The categories attribute supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the key name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `value of the key. ### Reference The ` + "`" + `project_reference` + "`" + `, ` + "`" + `owner_reference` + "`" + ` attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind name (Default value: ` + "`" + `project` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) the name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) the UUID. See detailed information in [Nutanix Protection Rule](https://www.nutanix.dev/reference/prism_central/v3/api/protection-rules/getprotectionrulesuuid/).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_protection_rules",
			Category:         "Data Sources",
			ShortDescription: `Describes a protection rules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone_connectivity_list",
					Description: `(Required) This encodes the datapipes between various availability zones and\nthe backup policy of the pipes.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.destination_availability_zone_index",
					Description: `(Optional/Computed) Index of the availability zone.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.source_availability_zone_index",
					Description: `(Optional/Computed) Index of the availability zone.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list",
					Description: `(Optional/Computed) Snapshot schedules for the pair of the availability zones.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.recovery_point_objective_secs",
					Description: `(Required) "A recovery point objective (RPO) is the maximum acceptable amount of data loss.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy",
					Description: `(Optional/Computed) This describes the snapshot retention policy for this availability zone.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy.0.num_snapshots",
					Description: `(Optional/Computed) Number of snapshots need to be retained.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy.0.rollup_retention_policy_multiple",
					Description: `(Optional/Computed) Multiplier to 'snapshot_interval_type'.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.local_snapshot_retention_policy.0.rollup_retention_policy_snapshot_interval_type",
					Description: `(Optional/Computed)`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.auto_suspend_timeout_secs",
					Description: `(Optional/Computed) Auto suspend timeout in case of connection failure between the sites.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.snapshot_type",
					Description: `(Optional/Computed) Crash consistent or Application Consistent snapshot.`,
				},
				resource.Attribute{
					Name:        "availability_zone_connectivity_list.snapshot_schedule_list.#.remote_snapshot_retention_policy",
					Description: `(Optional/Computed) This describes the snapshot retention policy for this availability zone. ### Ordered Availability Zone List`,
				},
				resource.Attribute{
					Name:        "ordered_availability_zone_list",
					Description: `(Required) A list of availability zones, each of which, receives a replica\nof the data for the entities protected by this protection rule.`,
				},
				resource.Attribute{
					Name:        "ordered_availability_zone_list.#.cluster_uuid",
					Description: `(Optional/Computed) UUID of specific cluster to which we will be replicating.`,
				},
				resource.Attribute{
					Name:        "ordered_availability_zone_list.#.availability_zone_url",
					Description: `(Optional/Computed) The FQDN or IP address of the availability zone. ### Category Filter`,
				},
				resource.Attribute{
					Name:        "category_filter",
					Description: `(Optional/Computed)`,
				},
				resource.Attribute{
					Name:        "category_filter.0.type",
					Description: `(Optional/Computed) The type of the filter being used.`,
				},
				resource.Attribute{
					Name:        "category_filter.0.kind_list",
					Description: `(Optional/Computed) List of kinds associated with this filter.`,
				},
				resource.Attribute{
					Name:        "category_filter.0.params",
					Description: `(Optional/Computed) A list of category key and list of values. ### Metadata The metadata attribute exports the following:`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `UTC date and time in RFC-3339 format when vm was last updated.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `vm UUID.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `UTC date and time in RFC-3339 format when vm was created.`,
				},
				resource.Attribute{
					Name:        "spec_version",
					Description: `Version number of the latest spec.`,
				},
				resource.Attribute{
					Name:        "spec_hash",
					Description: `Hash of the spec. This will be returned from server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `vm name. ### Categories The categories attribute supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the key name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `value of the key. ### Reference The ` + "`" + `project_reference` + "`" + `, ` + "`" + `owner_reference` + "`" + ` attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind name (Default value: ` + "`" + `project` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) the name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) the UUID. See detailed information in [Nutanix Protection Rules](https://www.nutanix.dev/reference/prism_central/v3/api/projects/postprojectslist).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_recovery_plan",
			Category:         "Data Sources",
			ShortDescription: `Describe a Nutanix Recovery Plan and its values (if it has them).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "stage_list",
					Description: `(Required) Input for the stages of the Recovery Plan. Each stage will perform a predefined type of task.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_uuid",
					Description: `(Optional/Computed) UUID of stage.`,
				},
				resource.Attribute{
					Name:        "stage_list.delay_time_secs",
					Description: `(Optional/Computed) Amount of time in seconds to delay the execution of next stage after execution of current stage.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work",
					Description: `(Required) A stage specifies the work to be performed when the Recovery Plan is executed.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities",
					Description: `(Optional/Computed) Information about entities to be recovered.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities.0.entity_info_list",
					Description: `(Optional/Computed) Information about entities to be recovered as part of this stage. For VM, entity information will include set of scripts to be executed after recovery of VM. Only one of categories or any_entity_reference has to be provided.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities.0.entity_info_list.#.any_entity_reference_kind",
					Description: `(Optional/Computed) Reference to a kind.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities.0.entity_info_list.#.any_entity_reference_uuid",
					Description: `(Optional/Computed) Reference to a uuid.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities.0.entity_info_list.#.any_entity_reference_name",
					Description: `(Optional/Computed) Reference to a name.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities.0.entity_info_list.#.categories",
					Description: `(Optional/Computed) Categories for filtering entities. ### Parameters`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Required) Parameters for the Recovery Plan.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list",
					Description: `(Optional/Computed) Floating IP assignment for VMs upon recovery in an Availability Zone. This is applicable only for the public cloud Availability Zones.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.availability_zone_url",
					Description: `(Required) URL of the Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list",
					Description: `(Required) IP assignment for VMs upon recovery in the specified Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.test_floating_ip_config",
					Description: `(Optional/Computed) Configuration for assigning floating IP to a VM on the execution of the Recovery Plan.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.test_floating_ip_config.ip",
					Description: `(Optional/Computed) IP to be assigned to VM, in case of failover.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.test_floating_ip_config.should_allocate_dynamically",
					Description: `(Optional/Computed) Whether to allocate the floating IPs for the VMs dynamically.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.recovery_floating_ip_config",
					Description: `(Optional/Computed) Configuration for assigning floating IP to a VM on the execution of the Recovery Plan.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.recovery_floating_ip_config.ip",
					Description: `(Optional/Computed) IP to be assigned to VM, in case of failover.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.recovery_floating_ip_config.should_allocate_dynamically",
					Description: `(Optional/Computed) Whether to allocate the floating IPs for the VMs dynamically.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_reference",
					Description: `(Required) Reference to a vm.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_reference.kind",
					Description: `(Required) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_reference.uuid",
					Description: `(Required) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_nic_information",
					Description: `(Required) Information about vnic to which floating IP has to be assigned.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_nic_information.ip",
					Description: `(Optional/Computed) IP address associated with vnic for which floating IP has to be assigned on failover.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_nic_information.uuid",
					Description: `(Required) Uuid of the vnic of the VM to which floating IP has to be assigned.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list",
					Description: `(Required) Network mappings to be used for the Recovery Plan. This will be represented by array of network mappings across the Availability Zones.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list",
					Description: `(Required) Mapping of networks across the Availability Zones.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.availability_zone_url",
					Description: `(Optional/Computed) URL of the Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network",
					Description: `(Optional/Computed) Network configuration to be used for performing network mapping and IP preservation/mapping on Recovery Plan execution.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.virtual_network_reference",
					Description: `(Optional/Computed) The reference to a virtual_network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.virtual_network_reference.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.virtual_network_reference.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.virtual_network_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.use_vpc_reference",
					Description: `(Optional/Computed) The reference to a VPC.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.vpc_reference",
					Description: `(Optional/Computed) The reference to a VPC.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.vpc_reference.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.vpc_reference.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.vpc_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.subnet_list",
					Description: `(Optional/Computed) List of subnets for the network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.subnet_list.#.gateway_ip",
					Description: `(Required) Gateway IP address for the subnet.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.subnet_list.#.external_connectivity_state",
					Description: `(Optional/Computed) External connectivity state of the subnet. This is applicable only for the subnet to be created in public cloud Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.subnet_list.#.prefix_length",
					Description: `(Required) Prefix length for the subnet.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.name",
					Description: `(Required) Name of the network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network",
					Description: `(Optional/Computed) Network configuration to be used for performing network mapping and IP preservation/mapping on Recovery Plan execution.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.virtual_network_reference",
					Description: `(Optional/Computed) The reference to a virtual_network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.virtual_network_reference.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.virtual_network_reference.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.virtual_network_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.subnet_list",
					Description: `(Optional/Computed) List of subnets for the network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.subnet_list.#.gateway_ip",
					Description: `(Required) Gateway IP address for the subnet.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.subnet_list.#.external_connectivity_state",
					Description: `(Optional/Computed) External connectivity state of the subnet. This is applicable only for the subnet to be created in public cloud Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.subnet_list.#.prefix_length",
					Description: `(Required) Prefix length for the subnet.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.name",
					Description: `(Required) Name of the network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list",
					Description: `(Optional/Computed) Static IP configuration for the VMs to be applied post recovery in the recovery network for migrate/ failover action on the Recovery Plan.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.vm_reference",
					Description: `(Optional/Computed) The reference to a vm.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.vm_reference.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.vm_reference.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.vm_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.ip_config_list",
					Description: `(Optional/Computed) List of IP configurations for a VM.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.ip_config_list.#.ip_address",
					Description: `(Required) IP address.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list",
					Description: `(Optional/Computed) Static IP configuration for the VMs to be applied post recovery in the test network for test failover action on the Recovery Plan.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.vm_reference",
					Description: `(Optional/Computed) The reference to a vm.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.vm_reference.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.vm_reference.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.vm_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.ip_config_list",
					Description: `(Optional/Computed) List of IP configurations for a VM.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.ip_config_list.#.ip_address",
					Description: `(Required) IP address.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.cluster_reference_list",
					Description: `(Optional/Computed) The clusters where the recovery and test networks reside. This is required to specify network mapping across clusters for a Recovery Plan created to handle failover within the same Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.cluster_reference_list.0.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.cluster_reference_list.0.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.name",
					Description: `(Optional/Computed) The name. ### Metadata The metadata attribute exports the following:`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `UTC date and time in RFC-3339 format when vm was last updated.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `vm UUID.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `UTC date and time in RFC-3339 format when vm was created.`,
				},
				resource.Attribute{
					Name:        "spec_version",
					Description: `Version number of the latest spec.`,
				},
				resource.Attribute{
					Name:        "spec_hash",
					Description: `Hash of the spec. This will be returned from server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `vm name. ### Categories The categories attribute supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the key name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `value of the key. ### Reference The ` + "`" + `project_reference` + "`" + `, ` + "`" + `owner_reference` + "`" + ` attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind name (Default value: ` + "`" + `project` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) the name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) the UUID. See detailed information in [Nutanix Recovery Plan](https://www.nutanix.dev/reference/prism_central/v3/api/recovery-plans/getrecoveryplansuuid/).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "stage_list",
					Description: `(Required) Input for the stages of the Recovery Plan. Each stage will perform a predefined type of task.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_uuid",
					Description: `(Optional/Computed) UUID of stage.`,
				},
				resource.Attribute{
					Name:        "stage_list.delay_time_secs",
					Description: `(Optional/Computed) Amount of time in seconds to delay the execution of next stage after execution of current stage.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work",
					Description: `(Required) A stage specifies the work to be performed when the Recovery Plan is executed.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities",
					Description: `(Optional/Computed) Information about entities to be recovered.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities.0.entity_info_list",
					Description: `(Optional/Computed) Information about entities to be recovered as part of this stage. For VM, entity information will include set of scripts to be executed after recovery of VM. Only one of categories or any_entity_reference has to be provided.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities.0.entity_info_list.#.any_entity_reference_kind",
					Description: `(Optional/Computed) Reference to a kind.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities.0.entity_info_list.#.any_entity_reference_uuid",
					Description: `(Optional/Computed) Reference to a uuid.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities.0.entity_info_list.#.any_entity_reference_name",
					Description: `(Optional/Computed) Reference to a name.`,
				},
				resource.Attribute{
					Name:        "stage_list.stage_work.0.recover_entities.0.entity_info_list.#.categories",
					Description: `(Optional/Computed) Categories for filtering entities. ### Parameters`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Required) Parameters for the Recovery Plan.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list",
					Description: `(Optional/Computed) Floating IP assignment for VMs upon recovery in an Availability Zone. This is applicable only for the public cloud Availability Zones.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.availability_zone_url",
					Description: `(Required) URL of the Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list",
					Description: `(Required) IP assignment for VMs upon recovery in the specified Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.test_floating_ip_config",
					Description: `(Optional/Computed) Configuration for assigning floating IP to a VM on the execution of the Recovery Plan.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.test_floating_ip_config.ip",
					Description: `(Optional/Computed) IP to be assigned to VM, in case of failover.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.test_floating_ip_config.should_allocate_dynamically",
					Description: `(Optional/Computed) Whether to allocate the floating IPs for the VMs dynamically.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.recovery_floating_ip_config",
					Description: `(Optional/Computed) Configuration for assigning floating IP to a VM on the execution of the Recovery Plan.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.recovery_floating_ip_config.ip",
					Description: `(Optional/Computed) IP to be assigned to VM, in case of failover.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.recovery_floating_ip_config.should_allocate_dynamically",
					Description: `(Optional/Computed) Whether to allocate the floating IPs for the VMs dynamically.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_reference",
					Description: `(Required) Reference to a vm.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_reference.kind",
					Description: `(Required) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_reference.uuid",
					Description: `(Required) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_nic_information",
					Description: `(Required) Information about vnic to which floating IP has to be assigned.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_nic_information.ip",
					Description: `(Optional/Computed) IP address associated with vnic for which floating IP has to be assigned on failover.`,
				},
				resource.Attribute{
					Name:        "parameters.0.floating_ip_assignment_list.#.vm_ip_assignment_list.#.vm_nic_information.uuid",
					Description: `(Required) Uuid of the vnic of the VM to which floating IP has to be assigned.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list",
					Description: `(Required) Network mappings to be used for the Recovery Plan. This will be represented by array of network mappings across the Availability Zones.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list",
					Description: `(Required) Mapping of networks across the Availability Zones.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.availability_zone_url",
					Description: `(Optional/Computed) URL of the Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network",
					Description: `(Optional/Computed) Network configuration to be used for performing network mapping and IP preservation/mapping on Recovery Plan execution.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.virtual_network_reference",
					Description: `(Optional/Computed) The reference to a virtual_network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.virtual_network_reference.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.virtual_network_reference.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.virtual_network_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.use_vpc_reference",
					Description: `(Optional/Computed) The reference to a VPC.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.vpc_reference",
					Description: `(Optional/Computed) The reference to a VPC.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.vpc_reference.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.vpc_reference.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.vpc_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.subnet_list",
					Description: `(Optional/Computed) List of subnets for the network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.subnet_list.#.gateway_ip",
					Description: `(Required) Gateway IP address for the subnet.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.subnet_list.#.external_connectivity_state",
					Description: `(Optional/Computed) External connectivity state of the subnet. This is applicable only for the subnet to be created in public cloud Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.subnet_list.#.prefix_length",
					Description: `(Required) Prefix length for the subnet.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_network.0.name",
					Description: `(Required) Name of the network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network",
					Description: `(Optional/Computed) Network configuration to be used for performing network mapping and IP preservation/mapping on Recovery Plan execution.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.virtual_network_reference",
					Description: `(Optional/Computed) The reference to a virtual_network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.virtual_network_reference.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.virtual_network_reference.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.virtual_network_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.subnet_list",
					Description: `(Optional/Computed) List of subnets for the network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.subnet_list.#.gateway_ip",
					Description: `(Required) Gateway IP address for the subnet.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.subnet_list.#.external_connectivity_state",
					Description: `(Optional/Computed) External connectivity state of the subnet. This is applicable only for the subnet to be created in public cloud Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.subnet_list.#.prefix_length",
					Description: `(Required) Prefix length for the subnet.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_network.0.name",
					Description: `(Required) Name of the network.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list",
					Description: `(Optional/Computed) Static IP configuration for the VMs to be applied post recovery in the recovery network for migrate/ failover action on the Recovery Plan.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.vm_reference",
					Description: `(Optional/Computed) The reference to a vm.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.vm_reference.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.vm_reference.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.vm_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.ip_config_list",
					Description: `(Optional/Computed) List of IP configurations for a VM.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.recovery_ip_assignment_list.0.ip_config_list.#.ip_address",
					Description: `(Required) IP address.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list",
					Description: `(Optional/Computed) Static IP configuration for the VMs to be applied post recovery in the test network for test failover action on the Recovery Plan.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.vm_reference",
					Description: `(Optional/Computed) The reference to a vm.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.vm_reference.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.vm_reference.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.vm_reference.name",
					Description: `(Optional/Computed) The name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.ip_config_list",
					Description: `(Optional/Computed) List of IP configurations for a VM.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.ip_config_list.#.ip_address",
					Description: `(Required) IP address.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.cluster_reference_list",
					Description: `(Optional/Computed) The clusters where the recovery and test networks reside. This is required to specify network mapping across clusters for a Recovery Plan created to handle failover within the same Availability Zone.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.cluster_reference_list.0.kind",
					Description: `(Optional/Computed) The kind name.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.cluster_reference_list.0.uuid",
					Description: `(Optional/Computed) The uuid.`,
				},
				resource.Attribute{
					Name:        "parameters.0.network_mapping_list.#.availability_zone_network_mapping_list.#.test_ip_assignment_list.0.name",
					Description: `(Optional/Computed) The name. ### Metadata The metadata attribute exports the following:`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `UTC date and time in RFC-3339 format when vm was last updated.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `vm UUID.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `UTC date and time in RFC-3339 format when vm was created.`,
				},
				resource.Attribute{
					Name:        "spec_version",
					Description: `Version number of the latest spec.`,
				},
				resource.Attribute{
					Name:        "spec_hash",
					Description: `Hash of the spec. This will be returned from server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `vm name. ### Categories The categories attribute supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the key name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `value of the key. ### Reference The ` + "`" + `project_reference` + "`" + `, ` + "`" + `owner_reference` + "`" + ` attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind name (Default value: ` + "`" + `project` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) the name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) the UUID. See detailed information in [Nutanix Recovery Plan](https://www.nutanix.dev/reference/prism_central/v3/api/recovery-plans/getrecoveryplansuuid/).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_role",
			Category:         "Data Sources",
			ShortDescription: `This operation retrieves a role on the input parameters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_roles",
			Category:         "Data Sources",
			ShortDescription: `Describes a list of roles`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_subnet",
			Category:         "Data Sources",
			ShortDescription: `This operation retrieves a subnet based on the input parameters. A subnet is a block of IP addresses.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_subnets",
			Category:         "Data Sources",
			ShortDescription: `Describes a list of subnets`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_user",
			Category:         "Data Sources",
			ShortDescription: `This operation retrieves a user based on the input parameters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_user_group",
			Category:         "Data Sources",
			ShortDescription: `This operation retrieves a user based on the input parameters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_group_distinguished_name",
					Description: `(Optional) The distinguished name for the user group ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_user_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a datasource to retrieve all the user groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_users",
			Category:         "Data Sources",
			ShortDescription: `This operation retrieves a list of all the users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_virtual_machine",
			Category:         "Data Sources",
			ShortDescription: `Describes a Virtual Machine`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
				resource.Attribute{
					Name:        "ngt_enabled_capability_list",
					Description: `Application names that are enabled.`,
				},
				resource.Attribute{
					Name:        "guest_customization_cloud_init_meta_data",
					Description: `The contents of the meta_data configuration for cloud-init. This can be formatted as YAML or JSON. The value must be base64 encoded.`,
				},
				resource.Attribute{
					Name:        "disk_size_bytes",
					Description: `Size of the disk in Bytes.`,
				},
				resource.Attribute{
					Name:        "disk_size_mib",
					Description: `Size of the disk in MiB. Must match the size specified in 'disk_size_bytes' - rounded up to the nearest MiB - when that field is present.`,
				},
				resource.Attribute{
					Name:        "device_properties",
					Description: `Properties to a device.`,
				},
				resource.Attribute{
					Name:        "data_source_reference",
					Description: `Reference to a data source.`,
				},
				resource.Attribute{
					Name:        "volume_group_reference",
					Description: `Reference to a volume group. ### Device Properties The device_properties attribute supports the following.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `GPU {segment:bus:device:function} (sbdf) address if assigned.`,
				},
				resource.Attribute{
					Name:        "fraction",
					Description: `Fraction of the physical GPU assigned.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
				resource.Attribute{
					Name:        "ngt_enabled_capability_list",
					Description: `Application names that are enabled.`,
				},
				resource.Attribute{
					Name:        "guest_customization_cloud_init_meta_data",
					Description: `The contents of the meta_data configuration for cloud-init. This can be formatted as YAML or JSON. The value must be base64 encoded.`,
				},
				resource.Attribute{
					Name:        "disk_size_bytes",
					Description: `Size of the disk in Bytes.`,
				},
				resource.Attribute{
					Name:        "disk_size_mib",
					Description: `Size of the disk in MiB. Must match the size specified in 'disk_size_bytes' - rounded up to the nearest MiB - when that field is present.`,
				},
				resource.Attribute{
					Name:        "device_properties",
					Description: `Properties to a device.`,
				},
				resource.Attribute{
					Name:        "data_source_reference",
					Description: `Reference to a data source.`,
				},
				resource.Attribute{
					Name:        "volume_group_reference",
					Description: `Reference to a volume group. ### Device Properties The device_properties attribute supports the following.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `GPU {segment:bus:device:function} (sbdf) address if assigned.`,
				},
				resource.Attribute{
					Name:        "fraction",
					Description: `Fraction of the physical GPU assigned.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"nutanix_access_control_policies":   0,
		"nutanix_access_control_policy":     1,
		"nutanix_category_key":              2,
		"nutanix_cluster":                   3,
		"nutanix_clusters":                  4,
		"nutanix_host":                      5,
		"nutanix_hosts":                     6,
		"nutanix_image":                     7,
		"nutanix_karbon_cluster":            8,
		"nutanix_karbon_cluster_kubeconfig": 9,
		"nutanix_karbon_cluster_ssh":        10,
		"nutanix_karbon_clusters":           11,
		"nutanix_karbon_private_registries": 12,
		"nutanix_karbon_private_registry":   13,
		"nutanix_network_security_rule":     14,
		"nutanix_permission":                15,
		"nutanix_permissions":               16,
		"nutanix_project":                   17,
		"nutanix_projects":                  18,
		"nutanix_protection_rule":           19,
		"nutanix_protection_rules":          20,
		"nutanix_recovery_plan":             21,
		"nutanix_role":                      22,
		"nutanix_roles":                     23,
		"nutanix_subnet":                    24,
		"nutanix_subnets":                   25,
		"nutanix_user":                      26,
		"nutanix_user_group":                27,
		"nutanix_user_groups":               28,
		"nutanix_users":                     29,
		"nutanix_virtual_machine":           30,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
