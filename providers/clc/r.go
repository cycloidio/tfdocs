package clc

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "clc_group",
			Category:         "Resources",
			ShortDescription: `Manages a CLC server group.`,
			Description:      ``,
			Keywords: []string{
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name (or GUID) of this server group. Will resolve to existing if present.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Required, string) The name or ID of the parent group. Will error if absent or unable to resolve.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `(Required, string) The datacenter location of both parent group and this group. Examples: "WA1", "VA1"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, string) Description for server group (visible in control portal only)`,
				},
				resource.Attribute{
					Name:        "custom_fields",
					Description: `(Optional) See [CustomFields](#custom_fields) below for details. <a id="custom_fields"></a> ## CustomFields ` + "`" + `custom_fields` + "`" + ` is a block within the configuration that may be repeated to bind custom fields for a server. CustomFields need be set up in advance. Each ` + "`" + `custom_fields` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required, string) The ID of the custom field to set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, string) The value for the specified field.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clc_load_balancer",
			Category:         "Resources",
			ShortDescription: `Manages a CLC load balacner.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "data_center",
					Description: `(Required, string) The datacenter location of both parent group and this group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required, string) Either "enabled" or "disabled"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, string) Description for server group (visible in control portal only)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clc_load_balancer_pool",
			Category:         "Resources",
			ShortDescription: `Manages a CLC load balancer pool.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Required, string) The id of the load balancer.`,
				},
				resource.Attribute{
					Name:        "data_center",
					Description: `(Required, string) The datacenter location for this pool.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required, int) Either 80 or 443`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional, string) The configured balancing method. Either "roundRobin" (default) or "leastConnection".`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `(Optional, string) The configured persistence method. Either "standard" (default) or "sticky".`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clc_public_ip",
			Category:         "Resources",
			ShortDescription: `Manages a CLC public ip.`,
			Description:      ``,
			Keywords: []string{
				"public",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required, string) The name or ID of the server to bind IP to.`,
				},
				resource.Attribute{
					Name:        "internal_ip_address",
					Description: `(Required, string) The internal IP of the NIC to attach to. If not provided, a new internal NIC will be provisioned and used.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) See [Ports](#ports) below for details.`,
				},
				resource.Attribute{
					Name:        "source_restrictions",
					Description: `(Optional) See [SourceRestrictions](#source_restrictions) below for details. <a id="ports"></a> ## Ports ` + "`" + `ports` + "`" + ` is a block within the configuration that may be repeated to specify open ports on the target IP. Each ` + "`" + `ports` + "`" + ` block supports the following:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clc_server",
			Category:         "Resources",
			ShortDescription: `Manages the lifecycle of a CLC server.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_template",
					Description: `(Required, string) The basename of the server. A unique name will be generated by the platform.`,
				},
				resource.Attribute{
					Name:        "source_server_id",
					Description: `(Required, string) The name or ID of the base OS image. Examples: "ubuntu-14-64-template", "rhel-7-64-template", "win2012r2dtc-64"`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, string) The name or ID of the server group to spawn server into.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Required, int) The number of virtual cores`,
				},
				resource.Attribute{
					Name:        "memory_mb",
					Description: `(Required, int) Provisioned RAM`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, string) The virtualization type One of "standard", "hyperscale", "bareMetal"`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, string) The root/administrator password. Will be generated by platform if not provided.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, string) Description for server (visible in control portal only)`,
				},
				resource.Attribute{
					Name:        "power_state",
					Description: `(Optional, string) See [PowerStates](#power_states) below for details. If absent, defaults to ` + "`" + `started` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `(Optional, string) Set internal IP address. If absent, allocated and assigned from pool.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional, string) GUID of network to use. (Must be set up in advance from control portal.) When absent, the default network will be used.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional, string) Backup and replication strategy for disks. One of "standard", "premium"`,
				},
				resource.Attribute{
					Name:        "aa_policy_id",
					Description: `(Optional, string | hyperscale) Anti-Affinity policy ID`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `(Optional, string | bareMetal) Hardware configuration ID`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional, string | bareMetal) Operating system to install.`,
				},
				resource.Attribute{
					Name:        "additional_disks",
					Description: `(Optional) See [Disks](#disks) below for details.`,
				},
				resource.Attribute{
					Name:        "custom_fields",
					Description: `(Optional) See [CustomFields](#custom_fields) below for details.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Misc state storage for non-CLC metadata. <a id="server-types"></a> ## Server Types #### standard Cloud servers ` + "`" + `standard` + "`" + ` offer basic, commodity level performance with mixed spindle/SSD storage profiles. Additional features storage backups, snapshot/clone/archive, and widespread availability. #### hyperscale Hyperscale ` + "`" + `hyperscale` + "`" + ` servers offer significantly higher IOPS than standard servers for CPU and IO intensive servers. See the [FAQ](https://www.ctl.io/knowledge-base/servers/hyperscale-server-faq/) for more details. Physical host redundancy can be managed via [Anti-Affinity policies](https://www.ctl.io/knowledge-base/servers/centurylink-cloud-anti-affinity-policies/). #### bareMetal Bare metal ` + "`" + `bareMetal` + "`" + ` offers optimal compute performance and is available in select datacenters in CLC for approved customers. For more info see the [FAQ](https://www.ctl.io/knowledge-base/servers/bare-metal-faq/). For ` + "`" + `bareMetal` + "`" + `, the required fields ` + "`" + `source_server_id` + "`" + `, ` + "`" + `cpu` + "`" + `, and ` + "`" + `memory_mb` + "`" + ` are ignored and instead the following fields are required: - configuration_id - os_type Values for ` + "`" + `configuration_id` + "`" + ` and ` + "`" + `os_type` + "`" + ` are specific to each datacenter and are available via the API endpoints [here](https://www.ctl.io/api-docs/v2/#data-centers-get-data-center-bare-metal-capabilities). <a id="power_states"></a> ## PowerStates ` + "`" + `power_state` + "`" + ` may be used to set initial power state or modify existing instances.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `freeze machine: memory, processes, billing, monitoring.`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `shutdown gracefully`,
				},
				resource.Attribute{
					Name:        "reboot",
					Description: `restart gracefully`,
				},
				resource.Attribute{
					Name:        "reset",
					Description: `restart forcefully <a id="disks"></a> ## Disks ` + "`" + `additional_disks` + "`" + ` is a block within the configuration that may be repeated to specify the attached disks on a server. Each ` + "`" + `additional_disks` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, string) Either "raw" or "partitioned".`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `(Required, int) Size of allocated disk.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required, string, type:` + "`" + `partitioned` + "`" + `) The mountpoint for the disk. <a id="custom_fields"></a> ## CustomFields ` + "`" + `custom_fields` + "`" + ` is a block within the configuration that may be repeated to bind custom fields for a server. CustomFields need be set up in advance. Each ` + "`" + `custom_fields` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required, string) The ID of the custom field to set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, string) The value for the specified field. <a id="packages"></a> ## Packages ` + "`" + `packages` + "`" + ` is a block within the configuration that may be repeated to specify packages and their associated parameters to be run at instantiation. Packages facilitate various tasks like ssh key installation, kernel upgrades, etc. Package ID as well as parameters are configured via this block. Example: ` + "`" + `` + "`" + `` + "`" + `hcl # Configure the CLC Provider provider "clc_server" "ubuntu" { # ... packages { id = "77abb844-579d-478d-3955-c69ab4a7ba1a" SshKey = "ssh-rsa AAAAB3NzaC1yc2EAAAABIwAA..." } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"clc_group":              0,
		"clc_load_balancer":      1,
		"clc_load_balancer_pool": 2,
		"clc_public_ip":          3,
		"clc_server":             4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
