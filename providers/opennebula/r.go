package opennebula

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "opennebula_acl",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula acl resource.`,
			Description:      ``,
			Keywords: []string{
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user",
					Description: `(Required) User component of the new rule.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) Resource component of the new rule. Any combination of valid resources, separated by a ` + "`" + `+` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `(Optional) Rights component of the new rule. Any combination of valid Rights, separated by a ` + "`" + `+` + "`" + `. The following rights are valid:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone component of the new rule.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_group",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula group resource.`,
			Description:      ``,
			Keywords: []string{
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the group.`,
				},
				resource.Attribute{
					Name:        "admins",
					Description: `(Optional) List of Administrator user IDs part of the group.`,
				},
				resource.Attribute{
					Name:        "quotas",
					Description: `(Optional) See [Quotas parameters](#quotas-parameters) below for details`,
				},
				resource.Attribute{
					Name:        "sunstone",
					Description: `(Optional) Allow users and group admins to access specific views. See [Sunstone parameters](#sunstone-parameters) below for details`,
				},
				resource.Attribute{
					Name:        "opennebula",
					Description: `(Optional) OpenNebula core configuration. See [Opennebula parameters](#opennebula-parameters) below for details`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Group tags (Key = value)`,
				},
				resource.Attribute{
					Name:        "template_section",
					Description: `(Optional) Allow to add a custom vector. See [Template section parameters](#template-section-parameters) ### Quotas parameters ` + "`" + `quotas` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "datastore_quotas",
					Description: `(Optional) List of datastore quotas. See [Datastore quotas parameters](#datastore-quotas-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "network_quotas",
					Description: `(Optional) List of network quotas. See [Network quotas parameters](#network-quotas-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "image_quotas",
					Description: `(Optional) List of image quotas. See [Image quotas parameters](#image-quotas-parameters) below for details`,
				},
				resource.Attribute{
					Name:        "vm_quotas",
					Description: `(Optional) See [Virtual Machine quotas parameters](#virtual-machine-quotas-parameters) below for details #### Datastore quotas parameters ` + "`" + `datastore` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Datastore ID.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `(Optional) Maximum number of images allowed on the datastore. Defaults to ` + "`" + `default quota` + "`" + ``,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Total size in MB allowed on the datastore. Defaults to ` + "`" + `default quota` + "`" + ` #### Network quotas parameters ` + "`" + `network` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Network ID.`,
				},
				resource.Attribute{
					Name:        "leases",
					Description: `(Optional) Maximum number of ip leases allowed on the network. Defaults to ` + "`" + `default quota` + "`" + ` #### Image quotas parameters ` + "`" + `image` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Image ID.`,
				},
				resource.Attribute{
					Name:        "running_vms",
					Description: `(Optional) Maximum number of Virtual Machines in ` + "`" + `RUNNING` + "`" + ` state with this image ID attached. Defaults to ` + "`" + `default quota` + "`" + ` #### Virtual Machine quotas parameters ` + "`" + `vm` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional) Total of CPUs allowed. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) Total of memory (in MB) allowed. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `(Optional) Maximum number of Virtual Machines allowed. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "running_cpu",
					Description: `(Optional) Virtual Machine CPUs allowed in ` + "`" + `RUNNING` + "`" + ` state. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "running_memory",
					Description: `(Optional) Virtual Machine Memory (in MB) allowed in ` + "`" + `RUNNING` + "`" + ` state. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "running_vms",
					Description: `(Optional) Number of Virtual Machines allowed in ` + "`" + `RUNNING` + "`" + ` state. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Maximum disk global size (in MB) allowed on a ` + "`" + `SYSTEM` + "`" + ` datastore. Defaults to ` + "`" + `default quota` + "`" + `. ### Sunstone parameters`,
				},
				resource.Attribute{
					Name:        "default_view",
					Description: `(Optional) Default Sunstone view for regular users.`,
				},
				resource.Attribute{
					Name:        "views",
					Description: `(Optional) List of available views for regular users.`,
				},
				resource.Attribute{
					Name:        "group_admin_default_view",
					Description: `(Optional) Default Sunstone view for group admin users.`,
				},
				resource.Attribute{
					Name:        "group_admin_views",
					Description: `(Optional) List of available views for the group admins. #### Opennebula parameters`,
				},
				resource.Attribute{
					Name:        "default_image_persistent",
					Description: `(Optional) Control the default value for the ` + "`" + `persistent` + "`" + ` attribute on image creation ( clone and disk save-as).`,
				},
				resource.Attribute{
					Name:        "default_image_persistent_new",
					Description: `(Optional) Control the default value for the ` + "`" + `persistent` + "`" + ` attribute on image creation ( only new images).`,
				},
				resource.Attribute{
					Name:        "api_list_order",
					Description: `(Optional) Sets order of elements by ID in list API calls: asc or desc respectively for ascending or descending order. ### Template section parameters ` + "`" + `template_section` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The vector name.`,
				},
				resource.Attribute{
					Name:        "elements",
					Description: `(Optional) Collection of custom tags. ## Attribute Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the group.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Import ` + "`" + `opennebula_group` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_group.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the group.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Import ` + "`" + `opennebula_group` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_group.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_image",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula image resource.`,
			Description:      ``,
			Keywords: []string{
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the image.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the image.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Permissions applied to the image. Defaults to the UMASK in OpenNebula (in UNIX Format: owner-group-other => Use-Manage-Admin).`,
				},
				resource.Attribute{
					Name:        "clone_from_image",
					Description: `(Optional) ID or name of the image to clone from. Conflicts with ` + "`" + `path` + "`" + `, ` + "`" + `size` + "`" + ` and ` + "`" + `type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datastore_id",
					Description: `(Required) ID of the datastore used to store the image. The ` + "`" + `datastore_id` + "`" + ` must be an active ` + "`" + `IMAGE` + "`" + ` datastore.`,
				},
				resource.Attribute{
					Name:        "persistent",
					Description: `(Optional) Flag which indicates if the Image has to be persistent. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lock",
					Description: `(Optional) Lock the image with a specific lock level. Supported values: ` + "`" + `USE` + "`" + `, ` + "`" + `MANAGE` + "`" + `, ` + "`" + `ADMIN` + "`" + `, ` + "`" + `ALL` + "`" + ` or ` + "`" + `UNLOCK` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path or URL of the original image to use. Conflicts with ` + "`" + `clone_from_image` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the image. Supported values: ` + "`" + `OS` + "`" + `, ` + "`" + `CDROM` + "`" + `, ` + "`" + `DATABLOCK` + "`" + `, ` + "`" + `KERNEL` + "`" + `, ` + "`" + `RAMDISK` + "`" + ` or ` + "`" + `CONTEXT` + "`" + `. Conflicts with ` + "`" + `clone_from_image` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size of the image in MB. Conflicts with ` + "`" + `clone_from_image` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dev_prefix",
					Description: `(Optional) Device prefix on Virtual Machine. Usually one of these: ` + "`" + `hd` + "`" + `, ` + "`" + `sd` + "`" + ` or ` + "`" + `vd` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) Device target on Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Optional) OpenNebula Driver to use.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Image format. Example: ` + "`" + `raw` + "`" + `, ` + "`" + `qcow2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Name of the group which owns the image. Defaults to the caller primary group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Image tags (Key = value)`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Deprecated) Timeout (in Minutes) for Image availability. Defaults to 10 minutes.`,
				},
				resource.Attribute{
					Name:        "template_section",
					Description: `(Optional) Allow to add a custom vector. See [Template section parameters](#template-section-parameters) ### Template section parameters ` + "`" + `template_section` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The vector name.`,
				},
				resource.Attribute{
					Name:        "elements",
					Description: `(Optional) Collection of custom tags. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the image.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the image.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the image.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the image.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Import ` + "`" + `opennebula_image` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_image.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the image.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the image.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the image.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the image.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Import ` + "`" + `opennebula_image` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_image.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_security group",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula security group resource.`,
			Description:      ``,
			Keywords: []string{
				"security group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the security group.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Permissions applied on security group. Defaults to the UMASK in OpenNebula (in UNIX Format: owner-group-other => Use-Manage-Admin).`,
				},
				resource.Attribute{
					Name:        "commit",
					Description: `(Optional) Flag to commit changes on Virtual Machine on security group update. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) List of rules. See [Rule parameters](#rule-parameters) below for details`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Name of the group which owns the security group. Defaults to the caller primary group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Security group tags (Key = Value).`,
				},
				resource.Attribute{
					Name:        "template_section",
					Description: `(Optional) Allow to add a custom vector. See [Template section parameters](#template-section-parameters) ### Rule parameters ` + "`" + `rule` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol for the rule. Supported values: ` + "`" + `ALL` + "`" + `, ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` or ` + "`" + `IPSEC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `(Required) Direction of the traffic flow to allow, must be ` + "`" + `INBOUND` + "`" + ` or ` + "`" + `OUTBOUND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) VNET ID to be used as the source/destination IP addresses.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP (or starting IP if used with 'size') to apply the rule to.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Number of IPs to apply the rule from, starting with ` + "`" + `ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `(Optional) Comma separated list of ports and port ranges.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) Type of ICMP traffic to apply to when 'protocol' is ` + "`" + `ICMP` + "`" + `. See <https://docs.opennebula.org/5.12/operation/network_management/security_groups.html> for more details on allowed values. ### Template section parameters ` + "`" + `template_section` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The vector name.`,
				},
				resource.Attribute{
					Name:        "elements",
					Description: `(Optional) Collection of custom tags. ## Attribute Reference The following attribute are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the security group.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the security group.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the security group.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the security group.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the security group.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Import ` + "`" + `opennebula_security_group` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_security_group.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the security group.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the security group.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the security group.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the security group.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the security group.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Import ` + "`" + `opennebula_security_group` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_security_group.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_service",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula service resource.`,
			Description:      ``,
			Keywords: []string{
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the service.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Permissions applied on service. Defaults to the UMASK in OpenNebula (in UNIX Format: owner-group-other => Use-Manage-Admin).`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Required) Service will be instantiated from the template ID.`,
				},
				resource.Attribute{
					Name:        "extra_template",
					Description: `(Optional) Service information to be merged with the template during instantiate.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Set the id of the user owner of the newly created service. The corresponding ` + "`" + `uname` + "`" + ` will be computed.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `(Optional) Set the name of the user owner of the newly created service. The corresponding ` + "`" + `uid` + "`" + ` will be computed.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `(Optional) Set the id of the group owner of the newly created service. The corresponding ` + "`" + `gname` + "`" + ` will be computed.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `(Optional) Set the name of the group owner of the newly created service. The corresponding ` + "`" + `gid` + "`" + ` will be computed. ## Attribute Reference The following attribute are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the service.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the service.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the service.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the service.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the service.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `Map with the service name of each networks along with the id of the network.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Array with roles information containing: ` + "`" + `cardinality` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `nodes` + "`" + ` and ` + "`" + `state` + "`" + `. ## Import ` + "`" + `opennebula_service` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_service.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the service.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the service.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the service.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the service.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the service.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `Map with the service name of each networks along with the id of the network.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Array with roles information containing: ` + "`" + `cardinality` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `nodes` + "`" + ` and ` + "`" + `state` + "`" + `. ## Import ` + "`" + `opennebula_service` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_service.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_service_template",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula service template resource.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the service template.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Permissions applied on service template. Defaults to the UMASK in OpenNebula (in UNIX Format: owner-group-other => Use-Manage-Admin).`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) Service template definition in JSON format.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Set the id of the user owner of the newly created service template. The corresponding ` + "`" + `uname` + "`" + ` will be computed.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `(Optional) Set the name of the user owner of the newly created service template. The corresponding ` + "`" + `uid` + "`" + ` will be computed.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `(Optional) Set the id of the group owner of the newly created service template. The corresponding ` + "`" + `gname` + "`" + ` will be computed.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `(Optional) Set the name of the group owner of the newly created service template. The corresponding ` + "`" + `gid` + "`" + ` will be computed. ## Attribute Reference The following attribute are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the service.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the service.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the service.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the service.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the service. ## Import ` + "`" + `opennebula_security_group` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_service_template.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the service.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the service.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the service.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the service.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the service. ## Import ` + "`" + `opennebula_security_group` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_service_template.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_template",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula template resource.`,
			Description:      ``,
			Keywords: []string{
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual machine template.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Permissions applied on template. Defaults to the UMASK in OpenNebula (in UNIX Format: owner-group-other => Use-Manage-Admin).`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Name of the group which owns the template. Defaults to the caller primary group.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional) Amount of CPU shares assigned to the VM.`,
				},
				resource.Attribute{
					Name:        "vcpu",
					Description: `(Optional) Number of CPU cores presented to the VM.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) Amount of RAM assigned to the VM in MB.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `(Optional) See [Features parameters](#features-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Optional) Array of free form key=value pairs, rendered and added to the CONTEXT variables for the VM. Recommended to include: ` + "`" + `NETWORK = "YES"` + "`" + ` and ` + "`" + `SET_HOSTNAME = "$NAME"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "graphics",
					Description: `(Optional) See [Graphics parameters](#graphics-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `(Optional) See [OS parameters](#os-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional) Can be specified multiple times to attach several disks. See [Disks parameters](#disks-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "nic",
					Description: `(Optional) Can be specified multiple times to attach several NICs. See [Nic parameters](#nic-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "raw",
					Description: `(Optional) Allow to pass hypervisor level tuning content. See [Raw parameters](#raw-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "vmgroup",
					Description: `(Optional) See [VM group parameters](#vm-group-parameters) below for details. Changing this argument triggers a new resource.`,
				},
				resource.Attribute{
					Name:        "user_inputs",
					Description: `(Optional) Ask the user instantiating the template to define the values described.`,
				},
				resource.Attribute{
					Name:        "sched_requirements",
					Description: `(Optional) Scheduling requirements to deploy the resource following specific rule`,
				},
				resource.Attribute{
					Name:        "sched_ds_requirements",
					Description: `(Optional) Storage placement requirements to deploy the resource following specific rule.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Template tags (Key = Value).`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Deprecated) Text describing the OpenNebula template object, in Opennebula's XML string format.`,
				},
				resource.Attribute{
					Name:        "lock",
					Description: `(Optional) Lock the template with a specific lock level. Supported values: ` + "`" + `USE` + "`" + `, ` + "`" + `MANAGE` + "`" + `, ` + "`" + `ADMIN` + "`" + `, ` + "`" + `ALL` + "`" + ` or ` + "`" + `UNLOCK` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_section",
					Description: `(Optional) Allow to add a custom vector. See [Template section parameters](#template-section-parameters) ### Graphics parameters ` + "`" + `graphics` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Generally set to VNC.`,
				},
				resource.Attribute{
					Name:        "listen",
					Description: `(Required) Binding address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Binding Port.`,
				},
				resource.Attribute{
					Name:        "keymap",
					Description: `(Optional) Keyboard mapping. ### OS parameters ` + "`" + `os` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "arch",
					Description: `(Required) Hardware architecture of the Virtual machine. Must fit the host architecture.`,
				},
				resource.Attribute{
					Name:        "boot",
					Description: `(Optional) ` + "`" + `OS` + "`" + ` disk to use to boot on. ### Features parameters ` + "`" + `features` + "`" + ` supports the following arguments`,
				},
				resource.Attribute{
					Name:        "pea",
					Description: `(Optional) Physical address extension mode allows 32-bit guests to address more than 4 GB of memory. (Can be ` + "`" + `YES` + "`" + ` or ` + "`" + `NO` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "acpi",
					Description: `(Optional) Useful for power management, for example, with KVM guests it is required for graceful shutdown to work. (Can be ` + "`" + `YES` + "`" + ` or ` + "`" + `NO` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "apic",
					Description: `(Optional) Enables the advanced programmable IRQ management. Useful for SMP machines. (Can be ` + "`" + `YES` + "`" + ` or ` + "`" + `NO` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "localtime",
					Description: `(Optional) The guest clock will be synchronized to the host’s configured timezone when booted. Useful for Windows VMs. (Can be ` + "`" + `YES` + "`" + ` or ` + "`" + `NO` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "hyperv",
					Description: `(Optional) Add hyperv extensions to the VM. The options can be configured in the driver configuration, HYPERV_OPTIONS.`,
				},
				resource.Attribute{
					Name:        "guest_agent",
					Description: `(Optional) Enables the QEMU Guest Agent communication. This only creates the socket inside the VM, the Guest Agent itself must be installed and started in the VM. (Can be ` + "`" + `YES` + "`" + ` or ` + "`" + `NO` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "virtio_scsi_queues",
					Description: `(Optional) Numer of vCPU queues for the virtio-scsi controller.`,
				},
				resource.Attribute{
					Name:        "iothreads",
					Description: `(Optional) umber of iothreads for virtio disks. By default threads will be assign to disk by round robin algorithm. Disk thread id can be forced by disk IOTHREAD attribute. ### Disk parameters ` + "`" + `disk` + "`" + ` supports the following arguments`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) ID of the image to attach to the virtual machine. Conflicts with ` + "`" + `volatile_type` + "`" + ` and ` + "`" + `volatile_format` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size (in MB) of the image attached to the virtual machine. Not possible to change a cloned image size.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) Target name device on the virtual machine. Depends of the image ` + "`" + `dev_prefix` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Optional) OpenNebula image driver.`,
				},
				resource.Attribute{
					Name:        "dev_prefix",
					Description: `(Optional) Prefix for the emulated device this image will be mounted at. For instance, attribute of the Image will be used.`,
				},
				resource.Attribute{
					Name:        "cache",
					Description: `(Optional) Selects the cache mechanism for the disk. Values are default, none, writethrough, writeback, directsync and unsafe.`,
				},
				resource.Attribute{
					Name:        "discard",
					Description: `(Optional) Controls what’s done with with trim commands to the disk, the values can be ignore or discard.`,
				},
				resource.Attribute{
					Name:        "io",
					Description: `(Optional) Set IO policy. Values are threads, native.`,
				},
				resource.Attribute{
					Name:        "volatile_type",
					Description: `(Optional) Type of the volatile disk: ` + "`" + `swap` + "`" + ` or ` + "`" + `fs` + "`" + `. Type ` + "`" + `swap` + "`" + ` is not supported in vcenter. Conflicts with ` + "`" + `image_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "volatile_format",
					Description: `(Optional) Format of the volatile disk: ` + "`" + `raw` + "`" + ` or ` + "`" + `qcow2` + "`" + `. Conflicts with ` + "`" + `image_id` + "`" + `. Minimum 1 item. Maximum 8 items. ### NIC parameters ` + "`" + `nic` + "`" + ` supports the following arguments`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the virtual network to attach to the virtual machine.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) MAC of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `(Optional) Nic model driver. Example: ` + "`" + `virtio` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "physical_device",
					Description: `(Optional) Physical device hosting the virtual network.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) List of security group IDs to use on the virtual network. Minimum 1 item. Maximum 8 items. ### Raw parameters ` + "`" + `raw` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) - Hypervisor. Supported values: ` + "`" + `kvm` + "`" + `, ` + "`" + `lxd` + "`" + `, ` + "`" + `vmware` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Required) - Raw data to pass to the hypervisor. ### VM group parameters ` + "`" + `vmgroup` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "vmgroup_id",
					Description: `(Required) ID of the VM group to use.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) role of the VM group to use. ### Template section parameters ` + "`" + `template_section` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The vector name.`,
				},
				resource.Attribute{
					Name:        "elements",
					Description: `(Optional) Collection of custom tags. ## Attribute Reference The following attribute are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the template.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the template.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the template.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the template.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the template.`,
				},
				resource.Attribute{
					Name:        "reg_time",
					Description: `Registration time of the template.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Import ` + "`" + `opennebula_template` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_template.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the template.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the template.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the template.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the template.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the template.`,
				},
				resource.Attribute{
					Name:        "reg_time",
					Description: `Registration time of the template.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Import ` + "`" + `opennebula_template` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_template.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_user",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula user resource.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password of the user. It is required for all ` + "`" + `auth_driver` + "`" + ` excepted 'ldap'`,
				},
				resource.Attribute{
					Name:        "auth_driver",
					Description: `(Optional) Authentication Driver for User management. DEfaults to 'core'.`,
				},
				resource.Attribute{
					Name:        "primary_group",
					Description: `(Optional) Primary group ID of the User. Defaults to 0 (oneadmin).`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of secondary groups ID of the user.`,
				},
				resource.Attribute{
					Name:        "quotas",
					Description: `(Optional) See [Quotas parameters](#quotas-parameters) below for details`,
				},
				resource.Attribute{
					Name:        "ssh_public_key",
					Description: `(Optional) SSH public key.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Group tags (Key = value)`,
				},
				resource.Attribute{
					Name:        "template_section",
					Description: `(Optional) Allow to add a custom vector. See [Template section parameters](#template-section-parameters) ### Quotas parameters ` + "`" + `quotas` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "datastore_quotas",
					Description: `(Optional) List of datastore quotas. See [Datastore quotas parameters](#datastore-quotas-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "network_quotas",
					Description: `(Optional) List of network quotas. See [Network quotas parameters](#network-quotas-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "image_quotas",
					Description: `(Optional) List of image quotas. See [Image quotas parameters](#image-quotas-parameters) below for details`,
				},
				resource.Attribute{
					Name:        "vm_quotas",
					Description: `(Optional) See [Virtual Machine quotas parameters](#virtual-machine-quotas-parameters) below for details #### Datastore quotas parameters ` + "`" + `datastore` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Datastore ID.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `(Optional) Maximum number of images allowed on the datastore. Defaults to ` + "`" + `default quota` + "`" + ``,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Total size in MB allowed on the datastore. Defaults to ` + "`" + `default quota` + "`" + ` #### Network quotas parameters ` + "`" + `network` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Network ID.`,
				},
				resource.Attribute{
					Name:        "leases",
					Description: `(Optional) Maximum number of ip leases allowed on the network. Defaults to ` + "`" + `default quota` + "`" + ` #### Image quotas parameters ` + "`" + `image` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Image ID.`,
				},
				resource.Attribute{
					Name:        "running_vms",
					Description: `(Optional) Maximum number of Virtual Machines in ` + "`" + `RUNNING` + "`" + ` state with this image ID attached. Defaults to ` + "`" + `default quota` + "`" + ` #### Virtual Machine quotas parameters ` + "`" + `vm` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional) Total of CPUs allowed. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) Total of memory (in MB) allowed. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `(Optional) Maximum number of Virtual Machines allowed. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "running_cpu",
					Description: `(Optional) Virtual Machine CPUs allowed in ` + "`" + `RUNNING` + "`" + ` state. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "running_memory",
					Description: `(Optional) Virtual Machine Memory (in MB) allowed in ` + "`" + `RUNNING` + "`" + ` state. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "running_vms",
					Description: `(Optional) Number of Virtual Machines allowed in ` + "`" + `RUNNING` + "`" + ` state. Defaults to ` + "`" + `default quota` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Maximum disk global size (in MB) allowed on a ` + "`" + `SYSTEM` + "`" + ` datastore. Defaults to ` + "`" + `default quota` + "`" + `. ### Template section parameters ` + "`" + `template_section` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The vector name.`,
				},
				resource.Attribute{
					Name:        "elements",
					Description: `(Optional) Collection of custom tags. ## Attribute Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Import ` + "`" + `opennebula_user` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_user.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Import ` + "`" + `opennebula_user` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_user.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_virtual data center",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula virtual data center resource.`,
			Description:      ``,
			Keywords: []string{
				"virtual data center",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual data center.`,
				},
				resource.Attribute{
					Name:        "groups_ids",
					Description: `(Optional) List of group IDs part of the virtual data center.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `(Optional) List of zones. See [Zones parameters](#zones-parameters) below for details ### Zones parameters ` + "`" + `zones` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Zone ID from where resource to add in virtual data center are located. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "host_ids",
					Description: `(Optional) List of hosts from Zone ID to add in virtual data center.`,
				},
				resource.Attribute{
					Name:        "datastore_ids",
					Description: `(Optional) List of datastore from Zone ID to add in virtual data center.`,
				},
				resource.Attribute{
					Name:        "vnet_ids",
					Description: `(Optional) List of virtual networks from Zone ID to add in virtual data center.`,
				},
				resource.Attribute{
					Name:        "cluster_ids",
					Description: `(Optional) List of clusters from Zone ID to add in virtual data center. ## Attribute Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual data center. ## Import ` + "`" + `opennebula_virtual_data_center` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_virtual_data_center.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual data center. ## Import ` + "`" + `opennebula_virtual_data_center` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_virtual_data_center.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_virtual machine",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula virtual machine resource.`,
			Description:      ``,
			Keywords: []string{
				"virtual machine",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Permissions applied on virtual machine. Defaults to the UMASK in OpenNebula (in UNIX Format: owner-group-other => Use-Manage-Admin).`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Optional) If set, VM are instantiated from the template ID. See [Instantiate from a template](#instantiate-from-a-template) for details. Changing this argument triggers a new resource.`,
				},
				resource.Attribute{
					Name:        "pending",
					Description: `(Optional) Pending state during VM creation. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional) Amount of CPU shares assigned to the VM.`,
				},
				resource.Attribute{
					Name:        "vpcu",
					Description: `(Optional) Number of CPU cores presented to the VM.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) Amount of RAM assigned to the VM in MB.`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Optional) Array of free form key=value pairs, rendered and added to the CONTEXT variables for the VM. Recommended to include: ` + "`" + `NETWORK = "YES"` + "`" + ` and ` + "`" + `SET_HOSTNAME = "$NAME"` + "`" + `. If a ` + "`" + `template_id` + "`" + ` is set, see [Instantiate from a template](#instantiate-from-a-template) for details.`,
				},
				resource.Attribute{
					Name:        "graphics",
					Description: `(Optional) See [Graphics parameters](#graphics-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `(Optional) See [OS parameters](#os-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional) Can be specified multiple times to attach several disks. See [Disk parameters](#disk-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "nic",
					Description: `(Optional) Can be specified multiple times to attach several NICs. See [Nic parameters](#nic-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "keep_nic_order",
					Description: `(Optional) Indicates if the provider should keep NIC list ordering at update.`,
				},
				resource.Attribute{
					Name:        "vmgroup",
					Description: `(Optional) See [VM group parameters](#vm-group-parameters) below for details. Changing this argument triggers a new resource.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Name of the group which owns the virtual machine. Defaults to the caller primary group.`,
				},
				resource.Attribute{
					Name:        "sched_requirements",
					Description: `(Optional) Scheduling requirements to deploy the resource following specific rule.`,
				},
				resource.Attribute{
					Name:        "sched_ds_requirements",
					Description: `(Optional) Storage placement requirements to deploy the resource following specific rule.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Virtual Machine tags (Key = Value).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Deprecated) Timeout (in Minutes) for VM availability. Defaults to 3 minutes.`,
				},
				resource.Attribute{
					Name:        "lock",
					Description: `(Optional) Lock the VM with a specific lock level. Supported values: ` + "`" + `USE` + "`" + `, ` + "`" + `MANAGE` + "`" + `, ` + "`" + `ADMIN` + "`" + `, ` + "`" + `ALL` + "`" + ` or ` + "`" + `UNLOCK` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "on_disk_change",
					Description: `(Optional) Select the behavior for changing disk images. Supported values: ` + "`" + `RECREATE` + "`" + ` or ` + "`" + `SWAP` + "`" + ` (default). ` + "`" + `RECREATE` + "`" + ` forces recreation of the vm and ` + "`" + `SWAP` + "`" + ` adopts the standard behavior of hot-swapping the disks. NOTE: This property does not affect the behavior of adding new disks.`,
				},
				resource.Attribute{
					Name:        "hard_shutdown",
					Description: `(Optional) If the VM doesn't have ACPI support, it immediately poweroff/terminate/reboot/undeploy the VM. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "template_section",
					Description: `(Optional) Allow to add a custom vector. See [Template section parameters](#template-section-parameters) ### Graphics parameters ` + "`" + `graphics` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Generally set to VNC.`,
				},
				resource.Attribute{
					Name:        "listen",
					Description: `(Required) Binding address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Binding Port.`,
				},
				resource.Attribute{
					Name:        "keymap",
					Description: `(Optional) Keyboard mapping. ### OS parameters ` + "`" + `os` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "arch",
					Description: `(Required) Hardware architecture of the Virtual machine. Must fit the host architecture.`,
				},
				resource.Attribute{
					Name:        "boot",
					Description: `(Optional) ` + "`" + `OS` + "`" + ` disk to use to boot on. ### Disk parameters ` + "`" + `disk` + "`" + ` supports the following arguments`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) ID of the image to attach to the virtual machine. Defaults to -1 if not set: this skip Image attchment to the VM. Conflicts with ` + "`" + `volatile_type` + "`" + ` and ` + "`" + `volatile_format` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size (in MB) of the image. If set, it will resize the image disk to the targeted size. The size must be greater than the current one.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) Target name device on the virtual machine. Depends of the image ` + "`" + `dev_prefix` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Optional) OpenNebula image driver.`,
				},
				resource.Attribute{
					Name:        "volatile_type",
					Description: `(Optional) Type of the disk: ` + "`" + `swap` + "`" + ` or ` + "`" + `fs` + "`" + `. Type ` + "`" + `swap` + "`" + ` is not supported in vcenter. Conflicts with ` + "`" + `image_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "volatile_format",
					Description: `(Optional) Format of the Image: ` + "`" + `raw` + "`" + ` or ` + "`" + `qcow2` + "`" + `. Conflicts with ` + "`" + `image_id` + "`" + `. Minimum 1 item. Maximum 8 items. A disk update will be triggered in adding or removing a ` + "`" + `disk` + "`" + ` section, or by a modification of any of these parameters: ` + "`" + `image_id` + "`" + `, ` + "`" + `target` + "`" + `, ` + "`" + `driver` + "`" + ` ### NIC parameters ` + "`" + `nic` + "`" + ` supports the following arguments`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the virtual network to attach to the virtual machine.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) MAC of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `(Optional) Nic model driver. Example: ` + "`" + `virtio` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "virtio_queues",
					Description: `(Optional) Virtio multi-queue size. Only if ` + "`" + `model` + "`" + ` is ` + "`" + `virtio` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "physical_device",
					Description: `(Optional) Physical device hosting the virtual network.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) List of security group IDs to use on the virtual network. Minimum 1 item. Maximum 8 items. A NIC update will be triggered in adding or removing a ` + "`" + `nic` + "`" + ` section, or by a modification of any of these parameters: ` + "`" + `network_id` + "`" + `, ` + "`" + `ip` + "`" + `, ` + "`" + `mac` + "`" + `, ` + "`" + `security_groups` + "`" + `, ` + "`" + `physical_device` + "`" + ` ### VM group parameters ` + "`" + `vmgroup` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "vmgroup_id",
					Description: `(Required) ID of the VM group to use.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) role of the VM group to use. ### Template section parameters ` + "`" + `template_section` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The vector name.`,
				},
				resource.Attribute{
					Name:        "elements",
					Description: `(Optional) Collection of custom tags. ## Attribute Reference The following attribute are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "lcmstate",
					Description: `LCM State of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "template_disk",
					Description: `when ` + "`" + `template_id` + "`" + ` is used and the template define some disks, this contains the template disks description.`,
				},
				resource.Attribute{
					Name:        "template_nic",
					Description: `when ` + "`" + `template_id` + "`" + ` is used and the template define some NICs, this contains the template NICs description.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration.`,
				},
				resource.Attribute{
					Name:        "template_tags",
					Description: `When ` + "`" + `template_id` + "`" + ` was set this keeps the template tags.`,
				},
				resource.Attribute{
					Name:        "template_section_names",
					Description: `When ` + "`" + `template_id` + "`" + ` was set this keeps the template section names only. ### Template NIC`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the image attached to the virtual machine.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `nic attachment identifier`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `network name`,
				},
				resource.Attribute{
					Name:        "computed_ip",
					Description: `IP of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "computed_mac",
					Description: `MAC of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "computed_model",
					Description: `Nic model driver.`,
				},
				resource.Attribute{
					Name:        "computed_virtio_queues",
					Description: `Virtio multi-queue size.`,
				},
				resource.Attribute{
					Name:        "computed_physical_device",
					Description: `Physical device hosting the virtual network.`,
				},
				resource.Attribute{
					Name:        "computed_security_groups",
					Description: `List of security group IDs to use on the virtual. ### Template disk`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image attached to the virtual machine.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `disk attachment identifier`,
				},
				resource.Attribute{
					Name:        "computed_size",
					Description: `Size (in MB) of the image attached to the virtual machine. Not possible to change a cloned image size.`,
				},
				resource.Attribute{
					Name:        "computed_target",
					Description: `Target name device on the virtual machine. Depends of the image ` + "`" + `dev_prefix` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "computed_driver",
					Description: `OpenNebula image driver. ### NIC`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `nic attachment identifier`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `network name`,
				},
				resource.Attribute{
					Name:        "computed_ip",
					Description: `IP of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "computed_mac",
					Description: `MAC of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "computed_model",
					Description: `Nic model driver.`,
				},
				resource.Attribute{
					Name:        "computed_virtio_queues",
					Description: `Virtio multi-queue size.`,
				},
				resource.Attribute{
					Name:        "computed_physical_device",
					Description: `Physical device hosting the virtual network.`,
				},
				resource.Attribute{
					Name:        "computed_security_groups",
					Description: `List of security group IDs to use on the virtual. ### Disk`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `disk attachment identifier`,
				},
				resource.Attribute{
					Name:        "computed_size",
					Description: `Size (in MB) of the image attached to the virtual machine. Not possible to change a cloned image size.`,
				},
				resource.Attribute{
					Name:        "computed_target",
					Description: `Target name device on the virtual machine. Depends of the image ` + "`" + `dev_prefix` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "computed_driver",
					Description: `OpenNebula image driver.`,
				},
				resource.Attribute{
					Name:        "computed_volatile_format",
					Description: `Format of the Image: ` + "`" + `raw` + "`" + ` or ` + "`" + `qcow2` + "`" + `. ## Instantiate from a template When the attribute ` + "`" + `template_id` + "`" + ` is set, here is the behavior: For all parameters excepted context: parameters present in VM overrides parameters defined in template. For context: it merges them. For disks and NICs defined in the template, if they are not overriden, are described in ` + "`" + `template_disk` + "`" + ` and ` + "`" + `template_nic` + "`" + ` attributes of the instantiated VM and are not modifiable anymore. ## Import ` + "`" + `opennebula_virtual_machine` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_virtual_machine.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "lcmstate",
					Description: `LCM State of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "template_disk",
					Description: `when ` + "`" + `template_id` + "`" + ` is used and the template define some disks, this contains the template disks description.`,
				},
				resource.Attribute{
					Name:        "template_nic",
					Description: `when ` + "`" + `template_id` + "`" + ` is used and the template define some NICs, this contains the template NICs description.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration.`,
				},
				resource.Attribute{
					Name:        "template_tags",
					Description: `When ` + "`" + `template_id` + "`" + ` was set this keeps the template tags.`,
				},
				resource.Attribute{
					Name:        "template_section_names",
					Description: `When ` + "`" + `template_id` + "`" + ` was set this keeps the template section names only. ### Template NIC`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the image attached to the virtual machine.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `nic attachment identifier`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `network name`,
				},
				resource.Attribute{
					Name:        "computed_ip",
					Description: `IP of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "computed_mac",
					Description: `MAC of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "computed_model",
					Description: `Nic model driver.`,
				},
				resource.Attribute{
					Name:        "computed_virtio_queues",
					Description: `Virtio multi-queue size.`,
				},
				resource.Attribute{
					Name:        "computed_physical_device",
					Description: `Physical device hosting the virtual network.`,
				},
				resource.Attribute{
					Name:        "computed_security_groups",
					Description: `List of security group IDs to use on the virtual. ### Template disk`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image attached to the virtual machine.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `disk attachment identifier`,
				},
				resource.Attribute{
					Name:        "computed_size",
					Description: `Size (in MB) of the image attached to the virtual machine. Not possible to change a cloned image size.`,
				},
				resource.Attribute{
					Name:        "computed_target",
					Description: `Target name device on the virtual machine. Depends of the image ` + "`" + `dev_prefix` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "computed_driver",
					Description: `OpenNebula image driver. ### NIC`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `nic attachment identifier`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `network name`,
				},
				resource.Attribute{
					Name:        "computed_ip",
					Description: `IP of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "computed_mac",
					Description: `MAC of the virtual machine on this network.`,
				},
				resource.Attribute{
					Name:        "computed_model",
					Description: `Nic model driver.`,
				},
				resource.Attribute{
					Name:        "computed_virtio_queues",
					Description: `Virtio multi-queue size.`,
				},
				resource.Attribute{
					Name:        "computed_physical_device",
					Description: `Physical device hosting the virtual network.`,
				},
				resource.Attribute{
					Name:        "computed_security_groups",
					Description: `List of security group IDs to use on the virtual. ### Disk`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `disk attachment identifier`,
				},
				resource.Attribute{
					Name:        "computed_size",
					Description: `Size (in MB) of the image attached to the virtual machine. Not possible to change a cloned image size.`,
				},
				resource.Attribute{
					Name:        "computed_target",
					Description: `Target name device on the virtual machine. Depends of the image ` + "`" + `dev_prefix` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "computed_driver",
					Description: `OpenNebula image driver.`,
				},
				resource.Attribute{
					Name:        "computed_volatile_format",
					Description: `Format of the Image: ` + "`" + `raw` + "`" + ` or ` + "`" + `qcow2` + "`" + `. ## Instantiate from a template When the attribute ` + "`" + `template_id` + "`" + ` is set, here is the behavior: For all parameters excepted context: parameters present in VM overrides parameters defined in template. For context: it merges them. For disks and NICs defined in the template, if they are not overriden, are described in ` + "`" + `template_disk` + "`" + ` and ` + "`" + `template_nic` + "`" + ` attributes of the instantiated VM and are not modifiable anymore. ## Import ` + "`" + `opennebula_virtual_machine` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_virtual_machine.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_virtual machine group",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula virtual machine group resource.`,
			Description:      ``,
			Keywords: []string{
				"virtual machine group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual machine group.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Permissions applied on virtual machine group. Defaults to the UMASK in OpenNebula (in UNIX Format: owner-group-other => Use-Manage-Admin.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) List of roles. See [Role parameters](#role-parameters) below for details.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Name of the group which owns the virtual machine group. Defaults to the caller primary group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Virtual Machine group tags.`,
				},
				resource.Attribute{
					Name:        "template_section",
					Description: `(Optional) Allow to add a custom vector. See [Template section parameters](#template-section-parameters) ### Role parameters ` + "`" + `role` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the role.`,
				},
				resource.Attribute{
					Name:        "host_affined",
					Description: `(Optional) List of Hosts affined to Virtual Machines using this role.`,
				},
				resource.Attribute{
					Name:        "host_anti_affined",
					Description: `(Optional) List of Hosts not-affined to Virtual Machines using this role.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) Policy to apply between Virtual Machines using this role. Allowed Values: ` + "`" + `NONE` + "`" + `, ` + "`" + `AFFINED` + "`" + `, ` + "`" + `ANTI_AFFINED` + "`" + `. ### Template section parameters ` + "`" + `template_section` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The vector name.`,
				},
				resource.Attribute{
					Name:        "elements",
					Description: `(Optional) Collection of custom tags. ## Attribute Reference The following attribute are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `See [Role Attribute Reference](#role-attribute-reference) below for details`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Role Attribute Reference The Following attributes are exported under ` + "`" + `role` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the role. ## Import ` + "`" + `opennebula_virtual_machine_group` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_virtual_machine_group.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the virtual machine.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `See [Role Attribute Reference](#role-attribute-reference) below for details`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ## Role Attribute Reference The Following attributes are exported under ` + "`" + `role` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the role. ## Import ` + "`" + `opennebula_virtual_machine_group` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_virtual_machine_group.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_virtual network",
			Category:         "Resources",
			ShortDescription: `Provides an OpenNebula virtual network resource.`,
			Description:      ``,
			Keywords: []string{
				"virtual network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the virtual network.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Permissions applied on virtual network. Defaults to the UMASK in OpenNebula (in UNIX Format: owner-group-other => Use-Manage-Admin).`,
				},
				resource.Attribute{
					Name:        "reservation_vnet",
					Description: `(Optional) ID of the parent virtual network to reserve from. Conflicts with all parameters except ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `permissions` + "`" + `, ` + "`" + `security_groups` + "`" + `, ` + "`" + `group` + "`" + `, ` + "`" + `reservation_ar_id` + "`" + `, ` + "`" + `reservation_first_ip` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reservation_size",
					Description: `(Optional) Size (in address) reserved. Conflicts with all parameters except ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `permissions` + "`" + `, ` + "`" + `security_groups` + "`" + `, ` + "`" + `group` + "`" + `, ` + "`" + `reservation_ar_id` + "`" + `, ` + "`" + `reservation_first_ip` + "`" + ` and ` + "`" + `reservation_vnet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reservation_ar_id",
					Description: `(Optional) ID of the address range from which to reserve the addresses. Conflicts with all parameters except ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `permissions` + "`" + `, ` + "`" + `security_groups` + "`" + `, ` + "`" + `group` + "`" + `, ` + "`" + `reservation_size` + "`" + `, ` + "`" + `reservation_first_ip` + "`" + ` and ` + "`" + `reservation_vnet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reservation_first_ip",
					Description: `(Optional) The first IPv4 address to start the reservation range. Conflicts with all parameters except ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `permissions` + "`" + `, ` + "`" + `security_groups` + "`" + `, ` + "`" + `group` + "`" + `, ` + "`" + `reservation_ar_id` + "`" + `, ` + "`" + `reservation_size` + "`" + ` and ` + "`" + `reservation_vnet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) List of security group IDs to apply on the virtual network.`,
				},
				resource.Attribute{
					Name:        "bridge",
					Description: `(Optional) Name of the bridge interface to which the virtual network should be associated. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "physical_device",
					Description: `(Optional) Name of the physical device interface to which the virtual network should be associated. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Virtual network type. One of these: ` + "`" + `dummy` + "`" + `, ` + "`" + `bridge` + "`" + `'` + "`" + `fw` + "`" + `, ` + "`" + `ebtables` + "`" + `, ` + "`" + `802.1Q` + "`" + `, ` + "`" + `vxlan` + "`" + ` or ` + "`" + `ovswitch` + "`" + `. Defaults to ` + "`" + `bridge` + "`" + `. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `(Deprecated) List of cluster IDs where the virtual network can be use. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_ids",
					Description: `(Optional) List of cluster IDs where the virtual network can be use. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `. Minimum 1 item.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Optional) ID of VLAN. Only if ` + "`" + `type` + "`" + ` is ` + "`" + `802.1Q` + "`" + `, ` + "`" + `vxlan` + "`" + ` or ` + "`" + `ovswitch` + "`" + `. Conflicts with ` + "`" + `reservation_vnet` + "`" + `, ` + "`" + `reservation_size` + "`" + ` and ` + "`" + `automatic_vlan_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "automatic_vlan_id",
					Description: `(Optional) Flag to let OpenNebula scheduler to attribute the VLAN ID. Conflicts with ` + "`" + `reservation_vnet` + "`" + `, ` + "`" + `reservation_size` + "`" + ` and ` + "`" + `vlan_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) Virtual network MTU. Defaults to ` + "`" + `1500` + "`" + `. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "guest_mtu",
					Description: `(Optional) MTU of the network caord on the virtual machine.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Optional) IP of the gateway. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_mask",
					Description: `(Optional) Network mask. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional) Text String containing a space separated list of DNS IPs. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_address",
					Description: `(Optional) Base network address. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "search_domain",
					Description: `(Optional) Default search domains for DNS resolution. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ar",
					Description: `(Deprecated) List of address ranges. See [Address Range Parameters](#address-range-parameters) below for more details. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hold_ips",
					Description: `(Deprecated) Hold Ips from any Address Range of the Virtual Network. The IP must be available to be held` + "`" + `. Conflicts with` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Name of the group which owns the virtual network. Defaults to the caller primary group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Virtual Network tags (Key = Value).`,
				},
				resource.Attribute{
					Name:        "lock",
					Description: `(Optional) Lock the virtual network with a specific lock level. Supported values: ` + "`" + `USE` + "`" + `, ` + "`" + `MANAGE` + "`" + `, ` + "`" + `ADMIN` + "`" + `, ` + "`" + `ALL` + "`" + ` or ` + "`" + `UNLOCK` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_section",
					Description: `(Optional) Allow to add a custom vector. See [Template section parameters](#template-section-parameters) ### Address Range parameters ` + "`" + `ar` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "ar_type",
					Description: `(Optional) Address range type. Supported values: ` + "`" + `IP4` + "`" + `, ` + "`" + `IP6` + "`" + `, ` + "`" + `IP6_STATIC` + "`" + `, ` + "`" + `IP4_6` + "`" + ` or ` + "`" + `IP4_6_STATIC` + "`" + ` or ` + "`" + `ETHER` + "`" + `. Defaults to ` + "`" + `IP4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip4",
					Description: `(Optional) Starting IPv4 address of the range. Required if ` + "`" + `ar_type` + "`" + ` is ` + "`" + `IP4` + "`" + ` or ` + "`" + `IP4_6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip6",
					Description: `(Optional) Starting IPv6 address of the range. Required if ` + "`" + `ar_type` + "`" + ` is ` + "`" + `IP6_STATIC` + "`" + ` or ` + "`" + `IP4_6_STATIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Address range size.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) Starting MAC Address of the range.`,
				},
				resource.Attribute{
					Name:        "global_prefix",
					Description: `(Optional) Global prefix for ` + "`" + `IP6` + "`" + ` or ` + "`" + `IP_4_6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ula_prefix",
					Description: `(Optional) ULA prefix for ` + "`" + `IP6` + "`" + ` or ` + "`" + `IP_4_6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Optional) Prefix length. Only needed for ` + "`" + `IP6_STATIC` + "`" + ` or ` + "`" + `IP4_6_STATIC` + "`" + ` ### Template section parameters ` + "`" + `template_section` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The vector name.`,
				},
				resource.Attribute{
					Name:        "elements",
					Description: `(Optional) Collection of custom tags. ## Attribute Reference The following attribute are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual network.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the virtual network.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the virtual network.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the virtual network.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the virtual network.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ### Address range computed attributes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the address range`,
				},
				resource.Attribute{
					Name:        "computed_ip6",
					Description: `Starting IPv6 address of the range.`,
				},
				resource.Attribute{
					Name:        "computed_mac",
					Description: `Starting MAC Address of the range.`,
				},
				resource.Attribute{
					Name:        "computed_global_prefix",
					Description: `Global prefix for type ` + "`" + `IP6` + "`" + ` or ` + "`" + `IP_4_6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "computed_ula_prefix",
					Description: `ULA prefix for type ` + "`" + `IP6` + "`" + ` or ` + "`" + `IP_4_6` + "`" + `. ## Import ` + "`" + `opennebula_virtual_network` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_virtual_network.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual network.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `User ID whom owns the virtual network.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `Group ID which owns the virtual network.`,
				},
				resource.Attribute{
					Name:        "uname",
					Description: `User Name whom owns the virtual network.`,
				},
				resource.Attribute{
					Name:        "gname",
					Description: `Group Name which owns the virtual network.`,
				},
				resource.Attribute{
					Name:        "tags_all",
					Description: `Result of the applied ` + "`" + `default_tags` + "`" + ` and then resource ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `Default tags defined in the provider configuration. ### Address range computed attributes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the address range`,
				},
				resource.Attribute{
					Name:        "computed_ip6",
					Description: `Starting IPv6 address of the range.`,
				},
				resource.Attribute{
					Name:        "computed_mac",
					Description: `Starting MAC Address of the range.`,
				},
				resource.Attribute{
					Name:        "computed_global_prefix",
					Description: `Global prefix for type ` + "`" + `IP6` + "`" + ` or ` + "`" + `IP_4_6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "computed_ula_prefix",
					Description: `ULA prefix for type ` + "`" + `IP6` + "`" + ` or ` + "`" + `IP_4_6` + "`" + `. ## Import ` + "`" + `opennebula_virtual_network` + "`" + ` can be imported using its ID: ` + "`" + `` + "`" + `` + "`" + `shell terraform import opennebula_virtual_network.example 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"opennebula_acl":                   0,
		"opennebula_group":                 1,
		"opennebula_image":                 2,
		"opennebula_security group":        3,
		"opennebula_service":               4,
		"opennebula_service_template":      5,
		"opennebula_template":              6,
		"opennebula_user":                  7,
		"opennebula_virtual data center":   8,
		"opennebula_virtual machine":       9,
		"opennebula_virtual machine group": 10,
		"opennebula_virtual network":       11,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
