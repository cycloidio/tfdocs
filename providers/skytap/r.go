package skytap

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "skytap_environment",
			Category:         "Resources",
			ShortDescription: `Provides a Skytap Environment resource.`,
			Description:      ``,
			Keywords: []string{
				"environment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_id",
					Description: `(Required, Force New) ID of the template you want to create an environment from. If updating with a new one then the environment will be recreated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User-defined name of the environment. Limited to 255 characters. UTF-8 character type. Will default to source template’s name if null is provided.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) User-defined description of the environment. Limited to 1000 characters. Null allowed. UTF-8 character type.`,
				},
				resource.Attribute{
					Name:        "outbound_traffic",
					Description: `(Optional) Indicates whether networks in the environment can send outbound traffic.`,
				},
				resource.Attribute{
					Name:        "routable",
					Description: `(Optional) Indicates whether networks within the environment can route traffic to one another.`,
				},
				resource.Attribute{
					Name:        "suspend_on_idle",
					Description: `(Optional) The number of seconds an environment can be idle before it is automatically suspended. Valid range: 300 to 86400 seconds (5 minutes to 1 day).`,
				},
				resource.Attribute{
					Name:        "suspend_at_time",
					Description: `(Optional) The date and time that the environment will be automatically suspended. Format: yyyy/mm/dd hh:mm:ss. By default, the suspend time uses the UTC offset for the time zone defined in your user account settings. Optionally, a different UTC offset can be supplied (for example: 2018/07/20 14:20:00 -0000). The value in the API response is converted to your time zone.`,
				},
				resource.Attribute{
					Name:        "shutdown_on_idle",
					Description: `(Optional) The number of seconds an environment can be idle before it is automatically shut down. Valid range: 300 to 86400 seconds (5 minutes to 1 day).`,
				},
				resource.Attribute{
					Name:        "shutdown_at_time",
					Description: `(Optional) The date and time that the environment will be automatically shut down. Format: yyyy/mm/dd hh:mm:ss. By default, the suspend time uses the UTC offset for the time zone defined in your user account settings. Optionally, a different UTC offset can be supplied (for example: 2018/07/20 14:20:00 -0000). The value in the API response is converted to your time zone.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of environment tags.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) Set of labels for the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Environment user data, available from the metadata server and the skytap api The ` + "`" + `label` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Label category that provide contextual meaning.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Label value to be use for reporting ~>`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the environment`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 10 mins) Used when updating the environment`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when destroying the environment ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skytap_icnr_tunnel",
			Category:         "Resources",
			ShortDescription: `Provides ICNR Tunnel resource.`,
			Description:      ``,
			Keywords: []string{
				"icnr",
				"tunnel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source",
					Description: `(Required, Force New) Source network from where the connection was initiated. This network does not need to be “tunnelable” (visible to other networks).`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required, Force New) Target network to which the connection was made. The network needs to be “tunnelable” (visible to other networks) ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) for certain operations:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the tunnel`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 10 mins) Used when updating the tunnel`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when destroying the tunnel ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skytap_label_category",
			Category:         "Resources",
			ShortDescription: `Provides a Skytap Label Category resource.`,
			Description:      ``,
			Keywords: []string{
				"label",
				"category",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User-defined label category name.`,
				},
				resource.Attribute{
					Name:        "single_value",
					Description: `(Required) With single value labels can have only one value for a category, with false labels can have multiple values. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) for certain operations:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the category`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 10 mins) Used when updating the category`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when destroying the category ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skytap_network",
			Category:         "Resources",
			ShortDescription: `Provides a Skytap Network resource.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required, Force New) ID of the environment you want to attach the network to. If updating with a new one then the network will be recreated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User-defined name of the network. Limited to 255 characters. UTF-8 character type.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain name for the Skytap network. Limited to 64 characters. ~>`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) Defines the subnet address and subnet mask size in CIDR format (for example, 10.0.0.0/24). IP addresses for the VMs are assigned from this subnet and standard network services (DNS resolution, CIFS share, routes to Internet) are defined appropriately for it. ~>`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Optional, Computed) Gateway IP address.`,
				},
				resource.Attribute{
					Name:        "tunnelable",
					Description: `(Optional) If true, this network can be connected to other networks. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) for certain operations:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the network`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 10 mins) Used when updating the network`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when destroying the network ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skytap_project",
			Category:         "Resources",
			ShortDescription: `Provides a Skytap Project resource.`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User-defined project name.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `(Optional) User-defined description of project.`,
				},
				resource.Attribute{
					Name:        "auto_add_role_name",
					Description: `(Optional) If this field is set to ` + "`" + `viewer` + "`" + `, ` + "`" + `participant` + "`" + `, ` + "`" + `editor` + "`" + `, or ` + "`" + `manager` + "`" + `, new users added to your Skytap account are automatically added to this project with the specified project role. Existing users aren’t affected by this setting. If the field is set to ` + "`" + `null` + "`" + `, new users aren’t automatically added to the project. For additional details, see [Automatically adding new users to a project](https://help.skytap.com/csh-project-automatic-role.html).`,
				},
				resource.Attribute{
					Name:        "show_project_members",
					Description: `(Optional) Determines whether projects members can view a list of other project members. False by default. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) for certain operations:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the project`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 10 mins) Used when updating the project`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) destroying the project ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skytap_vm",
			Category:         "Resources",
			ShortDescription: `Provides a Skytap Virtual Machine (VM) resource.`,
			Description:      ``,
			Keywords: []string{
				"vm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required, Force New) ID of the environment you want to add the VM to. If updating with a new one then the VM will be recreated.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Required, Force New) ID of the template you want to create the vm from. If updating with a new one then the VM will be recreated.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Required, Force New) ID of the VM you want to create the VM from. If updating with a new one then the VM will be recreated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, Computed) User-defined name. Limited to 100 characters.`,
				},
				resource.Attribute{
					Name:        "cpus",
					Description: `(Optional, Computed) Number of CPUs allocated to this virtual machine. Valid range is 1 to 12. Maximum limit depends on the ` + "`" + `max_cpus` + "`" + ` setting.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `(Optional, Computed) Amount of RAM allocated to this VM. Valid range is 256 and 131,072 (MB). Maximum limit depends on ` + "`" + `max_ram` + "`" + ` setting.`,
				},
				resource.Attribute{
					Name:        "os_disk_size",
					Description: `(Optional, Computed) The size of the OS disk. The disk size is in MiB; it will be converted to GiB in the Skytap UI. The maximum disk size is 2,096,128 MiB (1.999 TiB).`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional) Array of virtual disks within the VM. The disk size is in MiB; it will be converted to GiB in the Skytap UI. The maximum disk size is 2,096,128 MiB (1.999 TiB).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Specify the size of the disk. The new disk’s size should be provided in MiB. The minimum disk size is 2048 MiB; the maximum is 2096128 MiB (1.999 TiB). ~>`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Optional, Computed, ForceNew) A Skytap network adapter is a virtualized network interface card (also known as a network adapter). It is logically contained in a VM and attached to a network.`,
				},
				resource.Attribute{
					Name:        "interface_type",
					Description: `(Required, Force New) Type of network that this network adapter is attached to.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required, Force New) ID of the network that this network adapter is attached to.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required, Force New) The IP address (for example, 10.1.0.37). Skytap will not assign the same IP address to multiple interfaces on the same network.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required, Force New) Limited to 32 characters. Valid characters are lowercase letters, numbers, and hyphens. Cannot begin or end with hyphens. Cannot be ` + "`" + `gw` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "published_service",
					Description: `(Optional, Force New) Generally, a published service represents a binding of a port on a network interface to an IP and port that is routable and accessible from the public Internet. This mechanism is used to selectively expose ports on the guest to the public Internet. ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Force New) A unique name for the published service.`,
				},
				resource.Attribute{
					Name:        "internal_port",
					Description: `(Required, Force New) The port that is exposed on the interface. Typically this will be dictated by standard usage (e.g., port 80 for http traffic, port 22 for SSH).`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) VM user data, available from the metadata server and the skytap api`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) Set of labels for the instance. Structure is documented below. The ` + "`" + `label` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Label category that provide contextual meaning.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Label value to be use for reporting ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) for certain operations:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when launching the VM`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 10 mins) Used when stopping and starting the VM when necessary during update`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the instance ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"skytap_environment":    0,
		"skytap_icnr_tunnel":    1,
		"skytap_label_category": 2,
		"skytap_network":        3,
		"skytap_project":        4,
		"skytap_vm":             5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
