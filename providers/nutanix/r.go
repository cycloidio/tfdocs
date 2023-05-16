package nutanix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nutanix_access_control_policy",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create an access control policy based on the input parameters.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"control",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "entity_filter_expression_list",
					Description: `(Required) A list of Entity filter expressions. ### Scope Filter Expression List The scope_filter_expression_list attribute supports the following.`,
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
			Type:             "nutanix_address_group",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create a address group based on the input parameters.`,
			Description:      ``,
			Keywords: []string{
				"address",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_category_key",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Category key resource to Create a category key name.`,
			Description:      ``,
			Keywords: []string{
				"category",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) The version of the API. See detailed information in [Nutanix Image](https://nutanix.github.io/Automation/experimental/swagger-redoc-sandbox/#tag/category/paths/~1categories~1{name}/get).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) The version of the API. See detailed information in [Nutanix Image](https://nutanix.github.io/Automation/experimental/swagger-redoc-sandbox/#tag/category/paths/~1categories~1{name}/get).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_category_value",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Category value resource to Create a category value.`,
			Description:      ``,
			Keywords: []string{
				"category",
				"value",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value for the category value.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) The version of the API. See detailed information in [Nutanix Image](http://developer.nutanix.com/reference/prism_central/v3/#category).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) The version of the API. See detailed information in [Nutanix Image](http://developer.nutanix.com/reference/prism_central/v3/#category).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_floating_ip",
			Category:         "Resources",
			ShortDescription: `Create Floating IPs .`,
			Description:      ``,
			Keywords: []string{
				"floating",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "external_subnet_reference_uuid",
					Description: `(Optional) The reference to a subnet. Should not be used with {external_subnet_reference_name} .`,
				},
				resource.Attribute{
					Name:        "external_subnet_reference_name",
					Description: `(Optional) The reference to a subnet. Should not be used with {external_subnet_reference_uuid} .`,
				},
				resource.Attribute{
					Name:        "vm_nic_reference_uuid",
					Description: `(Optional) The reference to a vm_nic .`,
				},
				resource.Attribute{
					Name:        "vpc_reference_uuid",
					Description: `(Optional) The reference to a vpc. Should not be used with {vpc_reference_name}.`,
				},
				resource.Attribute{
					Name:        "vpc_reference_name",
					Description: `(Optional) The reference to a vpc. Should not be used with {vpc_reference_uuid}.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Private IP with which floating IP is associated. Should be used with vpc_reference . ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The floating_ips kind metadata.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `The floating_ips kind metadata.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_foundation_central_api_key",
			Category:         "Resources",
			ShortDescription: `Create a new api key which will be used by remote nodes to authenticate with Foundation Central .`,
			Description:      ``,
			Keywords: []string{
				"foundation",
				"central",
				"api",
				"key",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_foundation_central_image_cluster",
			Category:         "Resources",
			ShortDescription: `Image Nodes and Create a cluster out of nodes registered with Foundation Central.`,
			Description:      ``,
			Keywords: []string{
				"foundation",
				"central",
				"image",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_foundation_image",
			Category:         "Resources",
			ShortDescription: `Uploads hypervisor or AOS image to foundation.`,
			Description:      ``,
			Keywords: []string{
				"foundation",
				"image",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_foundation_image_nodes",
			Category:         "Resources",
			ShortDescription: `Images node(s) and optionally creates clusters.`,
			Description:      ``,
			Keywords: []string{
				"foundation",
				"image",
				"nodes",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_foundation_ipmi_config",
			Category:         "Resources",
			ShortDescription: `Configures IPMI IP address on BMC of nodes.`,
			Description:      ``,
			Keywords: []string{
				"foundation",
				"ipmi",
				"config",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_image",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Image resource to Create a Image.`,
			Description:      ``,
			Keywords: []string{
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "data_source_reference",
					Description: `(Optional) Reference to a data source. ### Version The version attribute supports the following:`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_karbon_cluster",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Karbon Cluster resource to Create a k8s cluster.`,
			Description:      ``,
			Keywords: []string{
				"karbon",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "reclaim_policy",
					Description: `(Optional) Reclaim policy for persistent volumes provisioned using the specified storage class.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.file_system",
					Description: `(Optional) Karbon uses either the ext4 or xfs file-system on the volume disk.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.flash_mode",
					Description: `(Optional) Pins the persistent volumes to the flash tier in case of a ` + "`" + `true` + "`" + ` value.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.password",
					Description: `(Required) The password of the Prism Element user that the API calls use to provision volumes.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.prism_element_cluster_uuid",
					Description: `(Required) The universally unique identifier (UUID) of the Prism Element cluster.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.storage_container",
					Description: `(Required) Name of the storage container the storage container uses to provision volumes.`,
				},
				resource.Attribute{
					Name:        "volumes_config.#.username",
					Description: `(Required) Username of the Prism Element user that the API calls use to provision volumes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_karbon_private_registry",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Karbon Registry resource to Create a private registry entry in Karbon.`,
			Description:      ``,
			Keywords: []string{
				"karbon",
				"private",
				"registry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_authorize_dbserver",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to authorize db server VMs for cloning of the database instance in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"authorize",
				"dbserver",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_clone",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to perform clone of the database instance in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"clone",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_clone_refresh",
			Category:         "Resources",
			ShortDescription: `NDB allows you to create and refresh clones to a point in time either by using transactional logs or by using snapshots. This operation submits a request to perform refresh clone of the database in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"clone",
				"refresh",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_cluster",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to add a Nutanix cluster to Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_database",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create, update and delete database instance in Nutanix database service (NDB). Note: For 1.8.0 release, only postgress database type is qualified and officially supported.`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"database",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_database_restore",
			Category:         "Resources",
			ShortDescription: `Restoring allows you to restore a source instance registered with NDB to a snapshot or point in time supported by the source instance time machine. You can restore an instance by using a snapshot ID, the point-in-time recovery (PITR) timestamp, or the latest snapshot. This operation submits a request to restore the database instance in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"database",
				"restore",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_database_scale",
			Category:         "Resources",
			ShortDescription: `Scaling the database extends the storage size proportionally across the attached virtual disks or volume groups. Scaling is supported for both single and HA instances. This operation submits a request to scale out the database instance in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"database",
				"scale",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_database_snapshot",
			Category:         "Resources",
			ShortDescription: `NDB time machine allows you to capture and replicate snapshots of the source database across multiple clusters (as defined in the DAM policy) at the time and frequency specified in the schedule. This operation submits a request to perform snapshot of the database instance in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"database",
				"snapshot",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_dbserver_vm",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create, update and delete database server VMs in Nutanix database service (NDB). Note: For 1.8.0 release, only postgress database type is qualified and officially supported.`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"dbserver",
				"vm",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_register_dbserver",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to register the database server VM in Nutanix database service (NDB). Note: For 1.8.0 release, only postgress database type is qualified and officially supported.`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"register",
				"dbserver",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_log_catchups",
			Category:         "Resources",
			ShortDescription: `A log catch-up operation copies transaction logs from the source database based on a specified schedule. The schedule can be provided during database registration or provisioning or can be modified later. This operation submits a request to perform log catchups of the database instance in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"log",
				"catchups",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_maintenance_task",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create, update and delete maintenance task association with database servers vms in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"maintenance",
				"task",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_maintenance_window",
			Category:         "Resources",
			ShortDescription: `A maintenance window allows you to set a schedule that is used to automate repeated maintenance tasks such as OS patching and database patching. NDB allows you to create a maintenance window and then associate the maintenance window with a list of database server VMs or an instance. This operation submits a request to create, update and delete maintenance window in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"maintenance",
				"window",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_network",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create, update and delete networks in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"network",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_profile",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create, update and delete profiles in Nutanix database service (NDB). Note: For 1.8.0 release, only postgress database type is qualified and officially supported.`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_register_database",
			Category:         "Resources",
			ShortDescription: `It helps to register a source (production) database running on a Nutanix cluster with NDB. When you register a database with NDB, the database server VM (VM that hosts the source database) is also registered with NDB. After you have registered a database with NDB, a time machine is created for that database. This operation submits a request to register the database in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"register",
				"database",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_sla",
			Category:         "Resources",
			ShortDescription: `SLAs are data retention policies that allow you to specify how long the daily, weekly, monthly, and quarterly snapshots are retained in NDB. This operation submits a request to create, update and delete slas in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"sla",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_software_version_profile",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create, update and delete software profile versions in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"software",
				"version",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_stretched_vlan",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create, update and delete stretched vlans in Nutanix database service (NDB). We can add a stretched VLAN to NDB by selecting the existing VLANs from each Nutanix cluster.`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"stretched",
				"vlan",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_tag",
			Category:         "Resources",
			ShortDescription: `NDB allows you to assign metadata to entities (clones, time machines, databases, and database servers) by using tags. When you are cloning a database, you can associate tags with the database that you are creating. This operation submits a request to create, update and delete tags in Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"tag",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_ndb_tms_cluster",
			Category:         "Resources",
			ShortDescription: `NDB multi-cluster allows you to manage time machine data availability across all the registered Nutanix clusters in NDB. This operation submits a request to add, update and delete clusters in time machine data availability for Nutanix database service (NDB).`,
			Description:      ``,
			Keywords: []string{
				"ndb",
				"tms",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_network_security_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Network Security Rule resource to Create a Network Security Rule .`,
			Description:      ``,
			Keywords: []string{
				"network",
				"security",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_pbr",
			Category:         "Resources",
			ShortDescription: `Create Policy Based Routing within VPCs .`,
			Description:      ``,
			Keywords: []string{
				"pbr",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of policy`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) priority of policy`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Required) Protocol type of policy based routing. Must be one of {TCP, UDP, ICMP, PROTOCOL_NUMBER, ALL} .`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Routing policy action. Must be one of {DENY, PERMIT, REROUTE} .`,
				},
				resource.Attribute{
					Name:        "service_ip_list",
					Description: `(Optional) IP addresses of network services. This field is valid only when action is REROUTE.`,
				},
				resource.Attribute{
					Name:        "vpc_reference_uuid",
					Description: `(Required) The reference to a vpc . Should not be used with {vpc_name} .`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `(Required) The reference to a vpc. Should not be used with {vpc_reference_uuid}`,
				},
				resource.Attribute{
					Name:        "is_bidirectional",
					Description: `(Optional) Additionally create Policy in reverse direction. Should be used with {TCP, UDP with start and end port ranges and ICMP with icmp code and type}. Supported with 2022.x. ## source source address of an IP packet. This could be either an ip prefix or an address_type .`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) address type of source. Should be one of {INTERNET, ALL}.`,
				},
				resource.Attribute{
					Name:        "subnet_ip",
					Description: `(Optional) IP subnet provided as an address.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Optional) prefix length of provided subnet. ## destination destination address of an IP packet. This could be either an ip prefix or an address_type .`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) address type of source. Should be one of {INTERNET, ALL}.`,
				},
				resource.Attribute{
					Name:        "subnet_ip",
					Description: `(Optional) IP subnet provided as an address.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Optional) prefix length of provided subnet. ## protocol_parameters Routing policy IP protocol parameters`,
				},
				resource.Attribute{
					Name:        "tcp",
					Description: `(Optional) TCP parameters in routing policy`,
				},
				resource.Attribute{
					Name:        "udp",
					Description: `(Optional) UDP parameters in routing policy`,
				},
				resource.Attribute{
					Name:        "icmp",
					Description: `(Optional) ICMP parameters in routing policy.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `(Optional) Protocol number in routing policy ## tcp, udp`,
				},
				resource.Attribute{
					Name:        "source_port_range",
					Description: `(Required) Range of TCP/UDP ports.`,
				},
				resource.Attribute{
					Name:        "destination_port_range",
					Description: `(Required) Range of TCP/UDP ports. ## source_port_range, destination_port_range`,
				},
				resource.Attribute{
					Name:        "start_port",
					Description: `(Required) start port number`,
				},
				resource.Attribute{
					Name:        "end_port",
					Description: `(Required) end port number ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The routing policies kind metadata.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `The routing policies kind metadata.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_project",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Category key resource to Create a Project.`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the project.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for project.`,
				},
				resource.Attribute{
					Name:        "use_project_internal",
					Description: `(Optional) flag to use project internal for user role mapping`,
				},
				resource.Attribute{
					Name:        "cluster_uuid",
					Description: `(Optional) The UUID of cluster. (Required when using project_internal flag).`,
				},
				resource.Attribute{
					Name:        "enable_collab",
					Description: `(Optional) flag to allow collaboration of projects. (Use with project_internal flag) ### Resource Domain`,
				},
				resource.Attribute{
					Name:        "resource_domain",
					Description: `(Required) The status for a resource domain (limits and values)`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources",
					Description: `(Required) Array of the utilization/limit for resource types`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.limit",
					Description: `(Required) The resource consumption limit (unspecified is unlimited)`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.resource_type",
					Description: `(Required) The type of resource (for example storage, CPUs) ### Account Reference List`,
				},
				resource.Attribute{
					Name:        "account_reference_list",
					Description: `(Optional/Computed) List of accounts associated with the project.`,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.kind",
					Description: `(Optional) The kind name. Default value is ` + "`" + `account` + "`" + ``,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.uuid",
					Description: `(Required) The UUID of an account.`,
				},
				resource.Attribute{
					Name:        "account_reference_list.#.name",
					Description: `(Optional/Computed) The name of an account. ### Environment Reference List`,
				},
				resource.Attribute{
					Name:        "environment_reference_list",
					Description: `(Optional/Computed) List of environments associated with the project.`,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.kind",
					Description: `(Optional) The kind name. Default value is ` + "`" + `environment` + "`" + ``,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.uuid",
					Description: `(Required) The UUID of an environment.`,
				},
				resource.Attribute{
					Name:        "environment_reference_list.#.name",
					Description: `(Optional/Computed) The name of an environment. ### Default Subnet Reference Map`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference",
					Description: `(Required) Reference to a subnet.`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.kind",
					Description: `(Optional) The kind name. Default value is ` + "`" + `subnet` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.uuid",
					Description: `(Required) The UUID of a subnet.`,
				},
				resource.Attribute{
					Name:        "default_subnet_reference.name",
					Description: `(Optional/Computed) The name of a subnet. ### user_reference_list`,
				},
				resource.Attribute{
					Name:        "user_reference_list",
					Description: `(Optional/Computed) List of users in the project.`,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.kind",
					Description: `(Optional) The kind name. Default value is ` + "`" + `user` + "`" + ``,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.uuid",
					Description: `(Required) The UUID of a user`,
				},
				resource.Attribute{
					Name:        "user_reference_list.#.name",
					Description: `(Optional/Computed) The name of a user. ### External User Group Reference List`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list",
					Description: `(Optional/Computed) List of directory service user groups. These groups are not managed by Nutanix.`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.kind",
					Description: `(Optional) The kind name. Default value is ` + "`" + `user_group` + "`" + ``,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.uuid",
					Description: `(Required) The UUID of a user_group`,
				},
				resource.Attribute{
					Name:        "external_user_group_reference_list.#.name",
					Description: `(Optional/Computed) The name of a user_group ### Subnet Reference List`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list",
					Description: `(Optional/Computed) List of subnets for the project.`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.kind",
					Description: `(Optional) The kind name. Default value is ` + "`" + `subnet` + "`" + ``,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.uuid",
					Description: `(Required) The UUID of a subnet`,
				},
				resource.Attribute{
					Name:        "subnet_reference_list.#.name",
					Description: `(Optional/Computed) The name of a subnet. ### External Network List`,
				},
				resource.Attribute{
					Name:        "external_network_list",
					Description: `(Optional/Computed) List of external networks associated with the project.`,
				},
				resource.Attribute{
					Name:        "external_network_list.#.uuid",
					Description: `(Required) The UUID of a network.`,
				},
				resource.Attribute{
					Name:        "external_network_list.#.name",
					Description: `(Optional/Computed) The name of a network. ### Tunnel Reference List`,
				},
				resource.Attribute{
					Name:        "tunnel_reference_list",
					Description: `(Optional/Computed) List of tunnels associated with the project.`,
				},
				resource.Attribute{
					Name:        "tunnel_reference_list.#.kind",
					Description: `(Optional) The kind name. Default value is ` + "`" + `tunnel` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tunnel_reference_list.#.uuid",
					Description: `(Required) The UUID of a tunnel`,
				},
				resource.Attribute{
					Name:        "tunnel_reference_list.#.name",
					Description: `(Optional/Computed) The name of a tunnel. ### Cluster Reference List`,
				},
				resource.Attribute{
					Name:        "cluster_reference_list",
					Description: `(Optional/Computed) List of clusters associated with the project..`,
				},
				resource.Attribute{
					Name:        "cluster_reference_list.#.kind",
					Description: `(Optional) The kind name. Default value is ` + "`" + `cluster` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cluster_reference_list.#.uuid",
					Description: `(Required) The UUID of a cluster`,
				},
				resource.Attribute{
					Name:        "cluster_reference_list.#.name",
					Description: `(Optional/Computed) The name of a cluster. ### VPC Reference List`,
				},
				resource.Attribute{
					Name:        "vpc_reference_list",
					Description: `(Optional/Computed) List of VPCs associated with the project..`,
				},
				resource.Attribute{
					Name:        "vpc_reference_list.#.kind",
					Description: `(Optional) The kind name. Default value is ` + "`" + `vpc` + "`" + ``,
				},
				resource.Attribute{
					Name:        "vpc_reference_list.#.uuid",
					Description: `(Required) The UUID of a vpc`,
				},
				resource.Attribute{
					Name:        "vpc_reference_list.#.name",
					Description: `(Optional/Computed) The name of a vpc. ### Default Environment Reference Map`,
				},
				resource.Attribute{
					Name:        "default_environment_reference",
					Description: `(Optional/Computed) Reference to a environment.`,
				},
				resource.Attribute{
					Name:        "default_environment_reference.kind",
					Description: `(Optional) The kind name. Default value is ` + "`" + `environment` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_environment_reference.uuid",
					Description: `(Required) The UUID of a environment`,
				},
				resource.Attribute{
					Name:        "default_environment_reference.name",
					Description: `(Optional/Computed) The name of a environment. ### ACP`,
				},
				resource.Attribute{
					Name:        "acp",
					Description: `(Optional) The list of ACPs to be attached to the users belonging to a project. It is mandate to provide cluster_uuid while using ACP. It helps to get the context list based on user role.`,
				},
				resource.Attribute{
					Name:        "acp.#.name",
					Description: `(Required) Name of the Access Control Policy.`,
				},
				resource.Attribute{
					Name:        "acp.#.description",
					Description: `The description of the association of a role to a user in a given context.`,
				},
				resource.Attribute{
					Name:        "acp.#.user_reference_list",
					Description: `The User(s) being assigned a given role.`,
				},
				resource.Attribute{
					Name:        "acp.#.user_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `user` + "`" + ``,
				},
				resource.Attribute{
					Name:        "acp.#.user_reference_list.#.uuid",
					Description: `(Required) The UUID of a user`,
				},
				resource.Attribute{
					Name:        "acp.#.user_reference_list.#.name",
					Description: `(Optional/Computed) The name of a user.`,
				},
				resource.Attribute{
					Name:        "acp.#.user_group_reference_list",
					Description: `The User group(s) being assigned a given role`,
				},
				resource.Attribute{
					Name:        "acp.#.user_group_reference_list.#.kind",
					Description: `The kind name. Default value is ` + "`" + `user_group` + "`" + ``,
				},
				resource.Attribute{
					Name:        "acp.#.user_group_reference_list.#.uuid",
					Description: `(Required) The UUID of a user group`,
				},
				resource.Attribute{
					Name:        "acp.#.user_group_reference_list.#.name",
					Description: `(Optional/Computed) The name of a user group.`,
				},
				resource.Attribute{
					Name:        "acp.#.role_reference",
					Description: `Reference to a role.`,
				},
				resource.Attribute{
					Name:        "acp.#.role_reference.kind",
					Description: `The kind name. Default value is ` + "`" + `role` + "`" + ``,
				},
				resource.Attribute{
					Name:        "acp.#.role_reference.uuid",
					Description: `(Required) The UUID of a role`,
				},
				resource.Attribute{
					Name:        "acp.#.role_reference.name",
					Description: `(Optional/Computed) The name of a role. ### User List`,
				},
				resource.Attribute{
					Name:        "user_list",
					Description: `(Optional) The list of user specification to be associated with the project. It is only required when user is not added in the PC.`,
				},
				resource.Attribute{
					Name:        "user_list.#.directory_service_user",
					Description: `(Optional) A Directory Service user.`,
				},
				resource.Attribute{
					Name:        "user_list.#.directory_service_user.user_principal_name",
					Description: `(Required) The UserPrincipalName of the user from the directory service.`,
				},
				resource.Attribute{
					Name:        "user_list.#.directory_service_user.directory_service_reference",
					Description: `(Required) Reference to a directory_service .`,
				},
				resource.Attribute{
					Name:        "user_list.#.directory_service_user.directory_service_reference.uuid",
					Description: `(Required) The uuid to a directory_service.`,
				},
				resource.Attribute{
					Name:        "user_list.#.directory_service_user.directory_service_reference.kind",
					Description: `(Optional) The kind to a directory_service.`,
				},
				resource.Attribute{
					Name:        "user_list.#.identity_provider_user",
					Description: `(Optional) An Identity Provider user.`,
				},
				resource.Attribute{
					Name:        "user_list.#.identity_provider_user.username",
					Description: `(Required) The username from the identity provider. Name Id for SAML Identity Provider.`,
				},
				resource.Attribute{
					Name:        "user_list.#.identity_provider_user.identity_provider_reference",
					Description: `(Required) The reference to a identity_provider.`,
				},
				resource.Attribute{
					Name:        "user_list.#.identity_provider_user.identity_provider_reference.uuid",
					Description: `(Required) The uuid to a identity_provider.`,
				},
				resource.Attribute{
					Name:        "user_list.#.identity_provider_user.identity_provider_reference.kind",
					Description: `(Optional) The kind to a identity_provider.`,
				},
				resource.Attribute{
					Name:        "user_list.#.metadata",
					Description: `(Required) Metadata Reference for user`,
				},
				resource.Attribute{
					Name:        "user_list.#.metadata.uuid",
					Description: `(Required) UUID of the USER`,
				},
				resource.Attribute{
					Name:        "user_list.#.metadata.Kind",
					Description: `Kind of the USER. ### User Group`,
				},
				resource.Attribute{
					Name:        "user_group",
					Description: `(Optional) The list of user group specification to be associated with the project. It is only Required when user group is not added in the PC.`,
				},
				resource.Attribute{
					Name:        "user_group.#.directory_service_user_group",
					Description: `(Optional) A Directory Service user group.`,
				},
				resource.Attribute{
					Name:        "user_group.#.directory_service_user_group.distinguished_name",
					Description: `(Required) The Distinguished name for the user group.`,
				},
				resource.Attribute{
					Name:        "user_group.#.saml_user_group",
					Description: `(Optional) A SAML Service user group.`,
				},
				resource.Attribute{
					Name:        "user_group.#.saml_user_group.idp_uuid",
					Description: `(Required) The UUID of the Identity Provider that the group belongs to.`,
				},
				resource.Attribute{
					Name:        "user_group.#.saml_user_group.name",
					Description: `(Required) The name of the SAML group which the IDP provides as attribute in SAML response.`,
				},
				resource.Attribute{
					Name:        "user_group.#.directory_service_ou",
					Description: `(Optional) A Directory Service user group.`,
				},
				resource.Attribute{
					Name:        "user_group.#.directory_service_ou.distinguished_name",
					Description: `(Required) The Distinguished name for the user group.`,
				},
				resource.Attribute{
					Name:        "user_group.#.metadata",
					Description: `(Required) Metadata Reference for user group`,
				},
				resource.Attribute{
					Name:        "user_group.#.metadata.uuid",
					Description: `(Required) UUID of the USER Group`,
				},
				resource.Attribute{
					Name:        "user_group.#.metadata.Kind",
					Description: `Kind of the USER Group. ## Attributes Reference The following attributes are exported: ### Resource Domain`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.units",
					Description: `The units of the resource type`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.value",
					Description: `The amount of resource consumed ### ACP ACPs will be exported if use_project_internal flag is set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of ACP`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of ACP`,
				},
				resource.Attribute{
					Name:        "user_reference_list",
					Description: `List of Reference of users.`,
				},
				resource.Attribute{
					Name:        "user_group_reference_list",
					Description: `List of Reference of users groups.`,
				},
				resource.Attribute{
					Name:        "role_reference",
					Description: `Reference to role.`,
				},
				resource.Attribute{
					Name:        "context_filter_list",
					Description: `The list of context filters. These are OR filters. The scope-expression-list defines the context, and the filter works in conjunction with the entity-expression-list. The context_list attribute supports the following:`,
				},
				resource.Attribute{
					Name:        "entity_filter_expression_list",
					Description: `(Required) A list of Entity filter expressions. ### Scope Filter Expression List The scope_filter_expression_list attribute supports the following.`,
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
					Description: `(Required) the UUID. Note: Few attributes which are added to support ACPs for Project are dependent on PC version. Features such as VPC, Cluster Reference requires pc2022.4 while Tunnel Reference requires pc2022.6 . See detailed information in [Nutanix Project](https://www.nutanix.dev/reference/prism_central/v3/api/projects/postprojects/).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_domain.resources.#.units",
					Description: `The units of the resource type`,
				},
				resource.Attribute{
					Name:        "resource_domain.resources.#.value",
					Description: `The amount of resource consumed ### ACP ACPs will be exported if use_project_internal flag is set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of ACP`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of ACP`,
				},
				resource.Attribute{
					Name:        "user_reference_list",
					Description: `List of Reference of users.`,
				},
				resource.Attribute{
					Name:        "user_group_reference_list",
					Description: `List of Reference of users groups.`,
				},
				resource.Attribute{
					Name:        "role_reference",
					Description: `Reference to role.`,
				},
				resource.Attribute{
					Name:        "context_filter_list",
					Description: `The list of context filters. These are OR filters. The scope-expression-list defines the context, and the filter works in conjunction with the entity-expression-list. The context_list attribute supports the following:`,
				},
				resource.Attribute{
					Name:        "entity_filter_expression_list",
					Description: `(Required) A list of Entity filter expressions. ### Scope Filter Expression List The scope_filter_expression_list attribute supports the following.`,
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
					Description: `(Required) the UUID. Note: Few attributes which are added to support ACPs for Project are dependent on PC version. Features such as VPC, Cluster Reference requires pc2022.4 while Tunnel Reference requires pc2022.6 . See detailed information in [Nutanix Project](https://www.nutanix.dev/reference/prism_central/v3/api/projects/postprojects/).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_protection_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Category key resource to Create a Protection Rule.`,
			Description:      ``,
			Keywords: []string{
				"protection",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the protection rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for protection rule. ### Availability Zone Connectivity List`,
				},
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
					Description: `(Optional/Computed) A list of category key and list of values. ## Attributes Reference The following attributes are exported: ### Metadata The metadata attribute exports the following:`,
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
					Description: `(Required) the UUID. See detailed information in [Nutanix Protection Rule](https://www.nutanix.dev/reference/prism_central/v3/api/protection-rules/postprotectionrules/).`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(Required) the UUID. See detailed information in [Nutanix Protection Rule](https://www.nutanix.dev/reference/prism_central/v3/api/protection-rules/postprotectionrules/).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_recovery_plan",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Category key resource to Create a Recovery Plan.`,
			Description:      ``,
			Keywords: []string{
				"recovery",
				"plan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the Recovery Plan.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for Recovery Plan. ### Stage List`,
				},
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
					Description: `(Optional/Computed) The name. ## Attributes Reference The following attributes are exported: ### Metadata The metadata attribute exports the following:`,
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
					Description: `(Required) the UUID. See detailed information in [Nutanix Recovery Plan](https://www.nutanix.dev/reference/prism_central/v3/api/recovery-plans/postrecoveryplans/).`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(Required) the UUID. See detailed information in [Nutanix Recovery Plan](https://www.nutanix.dev/reference/prism_central/v3/api/recovery-plans/postrecoveryplans/).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_role",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create a role based on the input parameters.`,
			Description:      ``,
			Keywords: []string{
				"role",
			},
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
			Type:             "nutanix_service_group",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create a service group based on the input parameters.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "end_port",
					Description: `(Optional) End Port (Int) #### UDP Port Range The udp_port_range_list attribute supports the following:`,
				},
				resource.Attribute{
					Name:        "end_port",
					Description: `(Optional) End Port (Int) See detailed information in [Nutanix Service Groups](https://www.nutanix.dev/reference/prism_central/v3/api/service-groups/postservicegroups).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_static_routes",
			Category:         "Resources",
			ShortDescription: `Create Static Routes within VPCs .`,
			Description:      ``,
			Keywords: []string{
				"static",
				"routes",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `(Required) Reference to a VPC UUID. Should not be used with vpc_name.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `(Required) vpc Name. Should not be used with vpc_uuid.`,
				},
				resource.Attribute{
					Name:        "static_routes_list",
					Description: `(Optional) Static Routes.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination ip with prefix.`,
				},
				resource.Attribute{
					Name:        "external_subnet_reference_uuid",
					Description: `(Optional) Reference to a subnet. Supported with 2022.x .`,
				},
				resource.Attribute{
					Name:        "vpn_connection_reference_uuid",
					Description: `(Optional) Reference to a vpn connection. ### default_route_nexthop`,
				},
				resource.Attribute{
					Name:        "external_subnet_reference_uuid",
					Description: `(Required) Reference to a subnet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The vpc_route_table kind metadata.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `The vpc_route_table kind metadata.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_subnet",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create a subnet based on the input parameters. A subnet is a block of IP addresses.`,
			Description:      ``,
			Keywords: []string{
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_user",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create a user based on the input parameters.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "entity_filter_expression_list",
					Description: `(Required) A list of Entity filter expressions. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "Resources",
			ShortDescription: `This operation add a User group to the system.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `The user_group kind metadata.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `The user_group kind metadata.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_virtual_machine",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Virtual Machine resource to Create a virtual machine.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"machine",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ngt_enabled_capability_list",
					Description: `(Optional) Application names that are enabled.`,
				},
				resource.Attribute{
					Name:        "guest_customization_cloud_init_meta_data",
					Description: `(Optional) The contents of the meta_data configuration for cloud-init. This can be formatted as YAML or JSON. The value must be base64 encoded.`,
				},
				resource.Attribute{
					Name:        "disk_size_bytes",
					Description: `(Optional) Size of the disk in Bytes.`,
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
					Description: `Reference to a volume group. The disk_size (the disk size_mib and the disk_size_bytes attributes) is only honored by creating an empty disk. When you are creating from an image, the size is ignored and the disk becomes the size of the image from which it was cloned. In VM creation, you can't set either disk size_mib or disk_size_bytes when you set data_source_reference but, you can update the disk_size after creation (second apply). ### Device Properties The device_properties attribute supports the following.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `(ReadOnly) GPU {segment:bus:device:function} (sbdf) address if assigned.`,
				},
				resource.Attribute{
					Name:        "fraction",
					Description: `(ReadOnly) Fraction of the physical GPU assigned.`,
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
			Type:             "nutanix_vpc",
			Category:         "Resources",
			ShortDescription: `Create Virtual Private Cloud .`,
			Description:      ``,
			Keywords: []string{
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the VPC.`,
				},
				resource.Attribute{
					Name:        "external_subnet_reference_uuid",
					Description: `(Required) List of external subnets uuid attached to this VPC. Should not be used with external_subnet_reference_name.`,
				},
				resource.Attribute{
					Name:        "external_subnet_reference_name",
					Description: `(Required) List of external subnets name attached to this VPC. Should not be used with external_subnet_reference_uuid.`,
				},
				resource.Attribute{
					Name:        "externally_routable_prefix_list",
					Description: `(Optional) List Externally Routable IP Addresses. Required when external subnet with NoNAT is used.`,
				},
				resource.Attribute{
					Name:        "common_domain_name_server_ip_list",
					Description: `(Optional) List of domain name server IPs. ## externally_routable_prefix_list Externally Routable IP Addresses`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The name for the VPC.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Required) prefix length. ## common_domain_name_server_ip_list List of domain name server IPs`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) ip address. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The vpc kind metadata.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
				resource.Attribute{
					Name:        "external_subnet_list_status",
					Description: `Status of List of external subnets attached to this VPC ### Metadata The metadata attribute exports the following:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `The vpc kind metadata.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
				resource.Attribute{
					Name:        "external_subnet_list_status",
					Description: `Status of List of external subnets attached to this VPC ### Metadata The metadata attribute exports the following:`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"nutanix_access_control_policy":            0,
		"nutanix_address_group":                    1,
		"nutanix_category_key":                     2,
		"nutanix_category_value":                   3,
		"nutanix_floating_ip":                      4,
		"nutanix_foundation_central_api_key":       5,
		"nutanix_foundation_central_image_cluster": 6,
		"nutanix_foundation_image":                 7,
		"nutanix_foundation_image_nodes":           8,
		"nutanix_foundation_ipmi_config":           9,
		"nutanix_image":                            10,
		"nutanix_karbon_cluster":                   11,
		"nutanix_karbon_private_registry":          12,
		"nutanix_ndb_authorize_dbserver":           13,
		"nutanix_ndb_clone":                        14,
		"nutanix_ndb_clone_refresh":                15,
		"nutanix_ndb_cluster":                      16,
		"nutanix_ndb_database":                     17,
		"nutanix_ndb_database_restore":             18,
		"nutanix_ndb_database_scale":               19,
		"nutanix_ndb_database_snapshot":            20,
		"nutanix_ndb_dbserver_vm":                  21,
		"nutanix_ndb_register_dbserver":            22,
		"nutanix_ndb_log_catchups":                 23,
		"nutanix_ndb_maintenance_task":             24,
		"nutanix_ndb_maintenance_window":           25,
		"nutanix_ndb_network":                      26,
		"nutanix_ndb_profile":                      27,
		"nutanix_ndb_register_database":            28,
		"nutanix_ndb_sla":                          29,
		"nutanix_ndb_software_version_profile":     30,
		"nutanix_ndb_stretched_vlan":               31,
		"nutanix_ndb_tag":                          32,
		"nutanix_ndb_tms_cluster":                  33,
		"nutanix_network_security_rule":            34,
		"nutanix_pbr":                              35,
		"nutanix_project":                          36,
		"nutanix_protection_rule":                  37,
		"nutanix_recovery_plan":                    38,
		"nutanix_role":                             39,
		"nutanix_service_group":                    40,
		"nutanix_static_routes":                    41,
		"nutanix_subnet":                           42,
		"nutanix_user":                             43,
		"nutanix_user_groups":                      44,
		"nutanix_virtual_machine":                  45,
		"nutanix_vpc":                              46,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
