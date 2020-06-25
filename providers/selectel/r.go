package selectel

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "selectel_domains_domain_v1",
			Category:         "Domains Resources",
			ShortDescription: `Manages a V1 domain resource within Selectel Domains API Service.`,
			Description:      ``,
			Keywords: []string{
				"domains",
				"domain",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the domain. Changing this creates a new domain name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the domain.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Identifier of the Selectel API user. ## Import Domain can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_domains_domain_v1.domain_1 45623 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the domain.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Identifier of the Selectel API user. ## Import Domain can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_domains_domain_v1.domain_1 45623 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_domains_record_v1",
			Category:         "Domains Resources",
			ShortDescription: `Manages a V1 record resource within Selectel Domains API Service.`,
			Description:      ``,
			Keywords: []string{
				"domains",
				"record",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Represents an identifier of the associated domain. Changing this creates a new domain record.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Represents a name of the domain record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Represents a type of the record. Possible values: A, AAAA, TXT, CNAME, NS, SOA, MX, SRV.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) Represents a time-to-live for the record. Must be the value between 60 and 604800.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Represents a content of the record. Absent for SRV records.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Represents an email of the domain's admin. For SOA records only.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Represents priority of the records preferences. Lower value means more preferred. MX/SRV records only.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Represents a relative weight for records with the same priority, higher value means higher chance of getting picked. For SRV records only.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Represents TCP or UDP port on which the service is to be found. For SRV records only.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) Represents a canonical hostname of the machine providing the service. For SRV records only. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Represents a content of the record.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Represents an email of the domain's admin.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Represents priority of the records preferences.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Represents a relative weight for records with the same priority, higher value means higher chance of getting picked.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Represents TCP or UDP port on which the service is to be found.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Represents a canonical hostname of the machine providing the service. ## Import Domain records can be imported using a combined ID in the following format: ` + "`" + `` + "`" + `<domain_id>/<record_id>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_domains_record_v1.record_1 45623/123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `Represents a content of the record.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Represents an email of the domain's admin.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Represents priority of the records preferences.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Represents a relative weight for records with the same priority, higher value means higher chance of getting picked.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Represents TCP or UDP port on which the service is to be found.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Represents a canonical hostname of the machine providing the service. ## Import Domain records can be imported using a combined ID in the following format: ` + "`" + `` + "`" + `<domain_id>/<record_id>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_domains_record_v1.record_1 45623/123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_mks_cluster_v1",
			Category:         "MKS Resources",
			ShortDescription: `Manages a V1 cluster resource within Selectel Managed Kubernetes Service.`,
			Description:      ``,
			Keywords: []string{
				"mks",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the cluster. Changing this creates a new cluster.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) An associated Selectel VPC project. Changing this creates a new cluster.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) A Selectel VPC region of where the cluster is located. Changing this creates a new cluster.`,
				},
				resource.Attribute{
					Name:        "kube_version",
					Description: `(Required) The current Kubernetes version of the cluster. Changing this upgrades the current version of the cluster. To upgrade a patch version, the desired version should match the latest available patch version for the current minor release. To upgrade a minor version, the desired version should match the next available minor release with the latest patch version.`,
				},
				resource.Attribute{
					Name:        "enable_autorepair",
					Description: `(Optional) Reflects if worker nodes are allowed to be reinstalled automatically. Accepts true or false. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "enable_patch_version_auto_upgrade",
					Description: `(Optional) Specifies if Kubernetes patch version of the cluster is allowed to be upgraded automatically. Accepts true or false. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) An associated OpenStack Networking service network ID. Changing this creates a new cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) associated OpenStack Networking service subnet ID. Changing this creates a new cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_window_start",
					Description: `(Optional) Represents UTC time in "hh:mm:ss" format of when the cluster will start its maintenance tasks. Changing this updates maintenance window start time. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "maintenance_window_end",
					Description: `Shows UTC time in "hh:mm:ss" format of when the cluster will end its maintenance tasks.`,
				},
				resource.Attribute{
					Name:        "kube_api_ip",
					Description: `Shows the IP of the Kubernetes API.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Shows the current status of the cluster. ## Import Cluster can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN SEL_PROJECT_ID=SELECTEL_VPC_PROJECT_ID SEL_REGION=SELECTEL_VPC_REGION terraform import selectel_mks_cluster_v1.cluster_1 b311ce58-2658-46b5-b733-7a0f418703f2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "maintenance_window_end",
					Description: `Shows UTC time in "hh:mm:ss" format of when the cluster will end its maintenance tasks.`,
				},
				resource.Attribute{
					Name:        "kube_api_ip",
					Description: `Shows the IP of the Kubernetes API.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Shows the current status of the cluster. ## Import Cluster can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN SEL_PROJECT_ID=SELECTEL_VPC_PROJECT_ID SEL_REGION=SELECTEL_VPC_REGION terraform import selectel_mks_cluster_v1.cluster_1 b311ce58-2658-46b5-b733-7a0f418703f2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_mks_nodegroup_v1",
			Category:         "MKS Resources",
			ShortDescription: `Manages a V1 nodegroup resource within Selectel Managed Kubernetes Service.`,
			Description:      ``,
			Keywords: []string{
				"mks",
				"nodegroup",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) An associated MKS cluster. Changing this creates a new nodegroup.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) An associated Selectel VPC project. Changing this creates a new nodegroup.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) A Selectel VPC region of where the nodegroup is located. Changing this creates a new nodegroup.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Contains a list of all nodes in the nodegroup. ## Import Nodegroup can be imported using a combined ID using the following format: ` + "`" + `` + "`" + `<cluster_id>/<nodegroup_id>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN SEL_PROJECT_ID=SELECTEL_VPC_PROJECT_ID SEL_REGION=SELECTEL_VPC_REGION terraform import selectel_mks_nodegroup_v1.nodegroup_1 b311ce58-2658-46b5-b733-7a0f418703f2/63ed5342-b22c-4c7a-9d41-c1fe4a142c13 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nodes",
					Description: `Contains a list of all nodes in the nodegroup. ## Import Nodegroup can be imported using a combined ID using the following format: ` + "`" + `` + "`" + `<cluster_id>/<nodegroup_id>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN SEL_PROJECT_ID=SELECTEL_VPC_PROJECT_ID SEL_REGION=SELECTEL_VPC_REGION terraform import selectel_mks_nodegroup_v1.nodegroup_1 b311ce58-2658-46b5-b733-7a0f418703f2/63ed5342-b22c-4c7a-9d41-c1fe4a142c13 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_vpc_floatingip_v2",
			Category:         "VPC Resources",
			ShortDescription: `Manages a V2 floating IP resource within Selectel VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"floatingip",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) An associated Selectel VPC project. Changing this creates a new floating IP.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) A region of where the floating IP resides. Changing this creates a new floating IP. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Contains id of the Networking service port.`,
				},
				resource.Attribute{
					Name:        "floating_ip_address",
					Description: `Contains floating IP address.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_address",
					Description: `Contains internal IP address of the Networking service port.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Shows if the license is used or not.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `Shows information about servers that use this floating IP. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_floatingip_v2.floatingip_1 aa402146-d83e-4c8c-8b74-1f415d4b8253 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "port_id",
					Description: `Contains id of the Networking service port.`,
				},
				resource.Attribute{
					Name:        "floating_ip_address",
					Description: `Contains floating IP address.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_address",
					Description: `Contains internal IP address of the Networking service port.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Shows if the license is used or not.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `Shows information about servers that use this floating IP. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_floatingip_v2.floatingip_1 aa402146-d83e-4c8c-8b74-1f415d4b8253 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_vpc_keypair_v2",
			Category:         "VPC Resources",
			ShortDescription: `Manages a V2 keypair resource within Selectel VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"keypair",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the keypair. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) A pregenerated OpenSSH-formatted public key. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) List of region names where keypair is need to be created. Keypair will be created in all available regions if omitted. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) An associated Selectel VPC user. Changing this creates a new keypair. ## Attributes Reference There are no additional attributes for this resource. ## Import Keypairs can be imported by specifying ` + "`" + `user_id` + "`" + ` and ` + "`" + `name` + "`" + ` arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_keypair_v2.keypair_1 <user_id>/<name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_vpc_license_v2",
			Category:         "VPC Resources",
			ShortDescription: `Manages a V2 license resource within Selectel VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"license",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) An associated Selectel VPC project. Changing this creates a new license.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) A region of where the license resides. Changing this creates a new license.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of license. Changing this creates a new license. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Shows if the license is used or not.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `Shows information about servers that use this license. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `Represents id of the associated network in the Networking service.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Represents id of the associated network in the Networking service.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Represents id of the associated port in the Networking service. ## Import Licenses can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_license_v2.license_1 4123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Shows if the license is used or not.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `Shows information about servers that use this license. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `Represents id of the associated network in the Networking service.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Represents id of the associated network in the Networking service.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Represents id of the associated port in the Networking service. ## Import Licenses can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_license_v2.license_1 4123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_vpc_project_v2",
			Category:         "VPC Resources",
			ShortDescription: `Manages a V2 project resource within Selectel VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"project",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project.`,
				},
				resource.Attribute{
					Name:        "custom_url",
					Description: `(Optional) The custom url for the project. Needs to be the 3rd-level domain for the ` + "`" + `selvpc.ru` + "`" + `. Example: ` + "`" + `terraform-project-001.selvpc.ru` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "theme",
					Description: `(Optional) An additional theme settings for this project. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "auto_quotas",
					Description: `(Optional) A boolean parameter that specifies if project should get automatically calculated quotas.`,
				},
				resource.Attribute{
					Name:        "quotas",
					Description: `(Optional) An array of desired quotas for this project. The structure is described below. The ` + "`" + `theme` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) A background color in hex format.`,
				},
				resource.Attribute{
					Name:        "logo",
					Description: `(Optional) An url of the project header logo. The ` + "`" + `quotas` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Required) A name of the billing resource to set quotas for.`,
				},
				resource.Attribute{
					Name:        "resource_quotas",
					Description: `(Required) An array of desired billing quotas for this particular resource. The structure is described below. The ` + "`" + `resource_quotas` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) A Selectel VPC region for the resource quota.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) A Selectel VPC zone for the resource quota.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) A value of the resource quota. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `An url of the Selectel VP project. It is set by the Selectel and can't be changed by the user.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Shows if project is active or it was disabled by the Selectel.`,
				},
				resource.Attribute{
					Name:        "all_quotas",
					Description: `Contains all quotas. They can differ from the configurable ` + "`" + `quota` + "`" + ` argument since the project will have all available resource quotas automatically applied. ## Import Projects can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_project_v2.project_1 0a343062504b4d06a0fac375e466db25 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `An url of the Selectel VP project. It is set by the Selectel and can't be changed by the user.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Shows if project is active or it was disabled by the Selectel.`,
				},
				resource.Attribute{
					Name:        "all_quotas",
					Description: `Contains all quotas. They can differ from the configurable ` + "`" + `quota` + "`" + ` argument since the project will have all available resource quotas automatically applied. ## Import Projects can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_project_v2.project_1 0a343062504b4d06a0fac375e466db25 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_vpc_role_v2",
			Category:         "VPC Resources",
			ShortDescription: `Manages a V2 role resource within Selectel VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"role",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) An associated Selectel VPC project. Changing this creates a new role.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) An associated Selectel VPC user. Changing this creates a new role. ## Attributes Reference There are no additional attributes for this resource. ## Import Roles can be imported by specifying ` + "`" + `project_id` + "`" + ` and ` + "`" + `user_id` + "`" + ` arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_role_v2.role_1 <project_id>/<user_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_vpc_subnet_v2",
			Category:         "VPC Resources",
			ShortDescription: `Manages a V2 subnet resource within Selectel VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"subnet",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) An associated Selectel VPC project. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) A region of where the subnet resides. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Optional) A prefix length of the subnet. Defaults to 29. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) A version of the IP protocol of the subnet. Defaults to "ipv4". Changing this creates a new subnet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Shows subnet CIDR representation.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `Shows associated OpenStack Networking service network ID.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Shows associated OpenStack Networking service subnet ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Shows if the subnet is used or not.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `Shows information about servers that use this subnet. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_subnet_v2.subnet_1 2060 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `Shows subnet CIDR representation.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `Shows associated OpenStack Networking service network ID.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Shows associated OpenStack Networking service subnet ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Shows if the subnet is used or not.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `Shows information about servers that use this subnet. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_subnet_v2.subnet_1 2060 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_vpc_token_v2",
			Category:         "VPC Resources",
			ShortDescription: `Manages a V2 token resource within Selectel VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"token",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) An associated Selectel VPC project. Changing this creates a new token.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Optional) An associated Selectel VPC account. Changing this creates a new token. ## Attributes Reference There are no additional attributes for this resource. ## Import Tokens can't be imported at this time.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_vpc_user_v2",
			Category:         "VPC Resources",
			ShortDescription: `Manages a V2 user resource within Selectel VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"user",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the user. Changing this updates the name of the existing user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of the user. Changing this updates the password of the existing user.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enabled state of the user. Changing this updates the enabled state of the existing user. ## Attributes Reference There are no additional attributes for this resource. ## Import Users can't be imported at this time.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "selectel_vpc_vrrp_subnet_v2",
			Category:         "VPC Resources",
			ShortDescription: `Manages a V2 VRRP subnet resource within Selectel VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"vrrp",
				"subnet",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) An associated Selectel VPC project. Changing this creates a new VRRP subnet.`,
				},
				resource.Attribute{
					Name:        "master_region",
					Description: `(Required) A master region of where the VRRP subnet resides. Changing this creates a new VRRP subnet.`,
				},
				resource.Attribute{
					Name:        "slave_region",
					Description: `(Required) A slave region of where the VRRP subnet resides. Changing this creates a new VRRP subnet.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Optional) A prefix length of the VRRP subnet. Defaults to 29. Changing this creates a new VRRP subnet.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) A version of the IP protocol of the VRRP subnet. Defaults to "ipv4". Changing this creates a new VRRP subnet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Shows VRRP subnet CIDR representation.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `Shows information about OpenStack Networking subnets that use this VRRP subnet. Contains ` + "`" + `network_id` + "`" + `, ` + "`" + `subnet_id` + "`" + ` and ` + "`" + `region` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Shows if the VRRP subnet is used or not.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `Shows information about servers that use this VRRP subnet. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields. ## Import VRRP subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_vrrp_subnet_v2.vrrp_subnet_1 2060 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `Shows VRRP subnet CIDR representation.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `Shows information about OpenStack Networking subnets that use this VRRP subnet. Contains ` + "`" + `network_id` + "`" + `, ` + "`" + `subnet_id` + "`" + ` and ` + "`" + `region` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Shows if the VRRP subnet is used or not.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `Shows information about servers that use this VRRP subnet. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields. ## Import VRRP subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_vrrp_subnet_v2.vrrp_subnet_1 2060 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"selectel_domains_domain_v1":  0,
		"selectel_domains_record_v1":  1,
		"selectel_mks_cluster_v1":     2,
		"selectel_mks_nodegroup_v1":   3,
		"selectel_vpc_floatingip_v2":  4,
		"selectel_vpc_keypair_v2":     5,
		"selectel_vpc_license_v2":     6,
		"selectel_vpc_project_v2":     7,
		"selectel_vpc_role_v2":        8,
		"selectel_vpc_subnet_v2":      9,
		"selectel_vpc_token_v2":       10,
		"selectel_vpc_user_v2":        11,
		"selectel_vpc_vrrp_subnet_v2": 12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
