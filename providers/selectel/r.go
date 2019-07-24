package selectel

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "selectel_vpc_crossregion_subnet_v2",
			Category:         "VPC Resources",
			ShortDescription: `Manages a V2 Cross-region subnet resource within VPC Selectel VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"crossregion",
				"subnet",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) An associated Selectel VPC project. Changing this creates a new Cross-region subnet.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Required) An array of regions where the Cross-region subnet resides. Changing this creates a new Cross-region subnet. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) A cross-region subnet CIDR representation. Changing this creates a new Cross-region subnet. The ` + "`" + `regions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) A region of where the Cross-region subnet resides. Changing this creates a new Cross-region subnet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `Shows information about servers that use this Cross-region subnet. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Shows if the Cross-region subnet is used or not.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `Shows information about OpenStack Networking subnets that use this Cross-region subnet. Contains ` + "`" + `cidr` + "`" + `, ` + "`" + `network_id` + "`" + `, ` + "`" + `project_id` + "`" + `, ` + "`" + `region` + "`" + `, ` + "`" + `subnet_id` + "`" + `, ` + "`" + `vlan_id` + "`" + ` and ` + "`" + `vtep_ip_address` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `Shows id of the associated VLAN in the OpenStack Networking service for this Cross-region subnet. ## Import Cross-region subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_crossregion_subnet_v2.crossregion_subnet_1 2060 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "servers",
					Description: `Shows information about servers that use this Cross-region subnet. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Shows if the Cross-region subnet is used or not.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `Shows information about OpenStack Networking subnets that use this Cross-region subnet. Contains ` + "`" + `cidr` + "`" + `, ` + "`" + `network_id` + "`" + `, ` + "`" + `project_id` + "`" + `, ` + "`" + `region` + "`" + `, ` + "`" + `subnet_id` + "`" + `, ` + "`" + `vlan_id` + "`" + ` and ` + "`" + `vtep_ip_address` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `Shows id of the associated VLAN in the OpenStack Networking service for this Cross-region subnet. ## Import Cross-region subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_crossregion_subnet_v2.crossregion_subnet_1 2060 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Shows information about servers that use this license. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields. ## Import Licenses can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_license_v2.license_1 4123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Shows if the license is used or not.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `Shows information about servers that use this license. Contains ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `status` + "`" + ` fields. ## Import Licenses can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ env SEL_TOKEN=SELECTEL_API_TOKEN terraform import selectel_vpc_license_v2.license_1 4123 ` + "`" + `` + "`" + `` + "`" + ``,
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

		"selectel_vpc_crossregion_subnet_v2": 0,
		"selectel_vpc_floatingip_v2":         1,
		"selectel_vpc_keypair_v2":            2,
		"selectel_vpc_license_v2":            3,
		"selectel_vpc_project_v2":            4,
		"selectel_vpc_role_v2":               5,
		"selectel_vpc_subnet_v2":             6,
		"selectel_vpc_token_v2":              7,
		"selectel_vpc_user_v2":               8,
		"selectel_vpc_vrrp_subnet_v2":        9,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
