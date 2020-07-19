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
					Description: `(Required) User component of the new rule. - ` + "`" + `#<id>` + "`" + ` matches a single user id - ` + "`" + `@<id>` + "`" + ` matches a group id - ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) Resource component of the new rule. Any combination of valid resources, separated by a ` + "`" + `+` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `(Optional) Rights component of the new rule. Any combination of valid Rights, separated by a ` + "`" + `+` + "`" + `. The following rights are valid: - USE - MANAGE - ADMIN - CREATE ## Import To import an existing ACL #134 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_acl" "importacl" { user = "@1" resource = "HOST+CLUSTER+DATASTORE/`,
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
					Name:        "template",
					Description: `(Required) Group template content in OpenNebula XML or String format. Used to provide SUSNTONE arguments.`,
				},
				resource.Attribute{
					Name:        "delete_on_destruction",
					Description: `(Optional) Flag to delete the group on destruction. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "admins",
					Description: `(Optional) List of Administrator user IDs part of the group.`,
				},
				resource.Attribute{
					Name:        "quotas",
					Description: `(Optional) See [Quotas parameters](#quotas-parameters) below for details ### Quotas parameters ` + "`" + `quotas` + "`" + ` supports the following arguments:`,
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
					Description: `(Optional) Maximum disk global size (in MB) allowed on a ` + "`" + `SYSTEM` + "`" + ` datastore. Defaults to ` + "`" + `default quota` + "`" + `. ## Attribute Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the group. ## Import To import an existing group #134 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_group" "importgroup" { name = "importedgroup" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_group.importgroup 134 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the group. ## Import To import an existing group #134 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_group" "importgroup" { name = "importedgroup" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_group.importgroup 134 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Timeout (in Minutes) for Image availability. Defaults to 10 minutes. ## Attribute Reference The following attributes are exported:`,
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
					Description: `Group Name which owns the image. ## Import To import an existing image #14 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_image" "importimage" { name = "importedimage" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_image.importimage 14 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Group Name which owns the image. ## Import To import an existing image #14 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_image" "importimage" { name = "importedimage" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_image.importimage 14 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Security group tags (Key = Value). ### Rule parameters ` + "`" + `rule` + "`" + ` supports the following arguments:`,
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
					Description: `(Optional) Type of ICMP traffic to apply to when 'protocol' is ` + "`" + `ICMP` + "`" + `. See https://docs.opennebula.org/5.12/operation/network_management/security_groups.html for more details on allowed values. ## Attribute Reference The following attribute are exported:`,
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
					Description: `Group Name which owns the security group. ## Import To import an existing security group #142 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_security_group" "importsg" { name = "importedsg" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_security_group.importsg 142 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Group Name which owns the security group. ## Import To import an existing security group #142 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_security_group" "importsg" { name = "importedsg" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_security_group.importsg 142 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "vpcu",
					Description: `(Optional) Number of CPU cores presented to the VM.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) Amount of RAM assigned to the VM in MB.`,
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
					Name:        "vmgroup",
					Description: `(Optional) See [VM group parameters](#vm-group-parameters) below for details. Changing this argument triggers a new resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Template tags (Key = Value).`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Deprecated) Text describing the OpenNebula template object, in Opennebula's XML string format. ### Graphics parameters ` + "`" + `graphics` + "`" + ` supports the following arguments:`,
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
					Description: `(Required) ID of the image to attach to the virtual machine.`,
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
					Description: `(Optional) OpenNebula image driver. Minimum 1 item. Maximum 8 items. ### NIC parameters ` + "`" + `nic` + "`" + ` supports the following arguments`,
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
					Description: `(Optional) List of security group IDs to use on the virtual network. Minimum 1 item. Maximum 8 items. ### VM group parameters ` + "`" + `vmgroup` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "vmgroup_id",
					Description: `(Required) ID of the VM group to use.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) role of the VM group to use. ## Attribute Reference The following attribute are exported:`,
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
					Description: `Registration time of the template. ## Import To import an existing virtual machine template #54 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_template" "importtpl" { name = "importedtpl" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_template.importtppl 54 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Registration time of the template. ## Import To import an existing virtual machine template #54 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_template" "importtpl" { name = "importedtpl" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_template.importtppl 54 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `ID of the virtual data center. ## Import To import an existing virtual data center #13 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_virtual_data_center" "importvdc" { name = "importedvdc" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_virtual_data_center.importvdc 13 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual data center. ## Import To import an existing virtual data center #13 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_virtual_data_center" "importvdc" { name = "importedvdc" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_virtual_data_center.importvdc 13 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "vmgroup",
					Description: `(Optional) See [VM group parameters](#vm-group-parameters) below for details. Changing this argument triggers a new resource.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Name of the group which owns the virtual machine. Defaults to the caller primary group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Virtual Machine tags (Key = Value).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout (in Minutes) for VM availability. Defaults to 3 minutes. ### Graphics parameters ` + "`" + `graphics` + "`" + ` supports the following arguments:`,
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
					Description: `(Required) ID of the image to attach to the virtual machine.`,
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
					Description: `(Optional) OpenNebula image driver. Minimum 1 item. Maximum 8 items. ### NIC parameters ` + "`" + `nic` + "`" + ` supports the following arguments`,
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
					Description: `(Optional) List of security group IDs to use on the virtual network. Minimum 1 item. Maximum 8 items. ### VM group parameters ` + "`" + `vmgroup` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "vmgroup_id",
					Description: `(Required) ID of the VM group to use.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) role of the VM group to use. ## Attribute Reference The following attribute are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `(Deprecated) Name of the virtual machine instance created on the cluster.`,
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
					Description: `LCM State of the virtual machine. ## Instantiate from a template When the attribute ` + "`" + `template_id` + "`" + ` is set, here is the behavior: For all parameters excepted context: parameters present in VM overrides parameters defined in template. For context: it merges them. ## Import To import an existing virtual machine #42 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_virtual_machine" "importvm" { name = "importedvm" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_virtual_machine.importvm 42 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `(Deprecated) Name of the virtual machine instance created on the cluster.`,
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
					Description: `LCM State of the virtual machine. ## Instantiate from a template When the attribute ` + "`" + `template_id` + "`" + ` is set, here is the behavior: For all parameters excepted context: parameters present in VM overrides parameters defined in template. For context: it merges them. ## Import To import an existing virtual machine #42 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_virtual_machine" "importvm" { name = "importedvm" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_virtual_machine.importvm 42 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Virtual Machine group tags. ### Role parameters ` + "`" + `role` + "`" + ` supports the following arguments:`,
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
					Description: `(Optional) Policy to apply between Virtual Machines using this role. Allowed Values: ` + "`" + `NONE` + "`" + `, ` + "`" + `AFFINED` + "`" + `, ` + "`" + `ANTI_AFFINED` + "`" + `. ## Attribute Reference The following attribute are exported:`,
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
					Description: `See [Role Attribute Reference](#role-attribute-reference) below for details ## Role Attribute Reference The Following attributes are exported under ` + "`" + `role` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the role.`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `List of Virtual Machine IDs using this role. ## Import To import an existing virtual machine group #42 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_virtual_machine_group" "importvmg" { name = "importedvmg" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_virtual_machine_group.importvmg 42 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See [Role Attribute Reference](#role-attribute-reference) below for details ## Role Attribute Reference The Following attributes are exported under ` + "`" + `role` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the role.`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `List of Virtual Machine IDs using this role. ## Import To import an existing virtual machine group #42 into Terraform, add this declaration to your .tf file: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_virtual_machine_group" "importvmg" { name = "importedvmg" } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_virtual_machine_group.importvmg 42 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) ID of the parent virtual network to reserve from. Conflicts with all parameters excepted ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `permissions` + "`" + `, ` + "`" + `security_groups` + "`" + ` and ` + "`" + `group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reservation_size",
					Description: `(Optional) Size (in address) reserved. Conflicts with all parameters excepted ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `permissions` + "`" + `, ` + "`" + `security_groups` + "`" + ` and ` + "`" + `group` + "`" + `.`,
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
					Description: `(Optional) List of cluster IDs where the virtual network can be use. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
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
					Description: `(Optional) Text String containing a comma separated list of DNS IPs. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ar",
					Description: `(Optional) List of address ranges. See [Address Range Parameters](#address-range-parameters) below for more details. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hold_size",
					Description: `(Optional) Carve a network reservation of this size from the reservation starting from ` + "`" + `ip_hold` + "`" + `. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_hold",
					Description: `(Optional) Start IP of the range to be held. Conflicts with ` + "`" + `reservation_vnet` + "`" + ` and ` + "`" + `reservation_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Name of the group which owns the virtual network. Defaults to the caller primary group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Virtual Network tags (Key = Value). ### Address Range parameters ` + "`" + `ar` + "`" + ` supports the following arguments:`,
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
					Description: `(Optional) Prefix length. Only needed for ` + "`" + `IP6_STATIC` + "`" + ` or ` + "`" + `IP4_6_STATIC` + "`" + ` ## Attribute Reference The following attribute are exported:`,
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
					Description: `Group Name which owns the virtual network. ## Import To import an existing virtual network #1234 into Terraform, add this declaration to your .tf file (don't specify the reservation_size): ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_virtual_network" "importtest" { name = "importedvnet" reservation_vnet = 394 security_groups = ["0"] } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_virtual_network.importtest 1234 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Group Name which owns the virtual network. ## Import To import an existing virtual network #1234 into Terraform, add this declaration to your .tf file (don't specify the reservation_size): ` + "`" + `` + "`" + `` + "`" + `hcl resource "opennebula_virtual_network" "importtest" { name = "importedvnet" reservation_vnet = 394 security_groups = ["0"] } ` + "`" + `` + "`" + `` + "`" + ` And then run: ` + "`" + `` + "`" + `` + "`" + ` terraform import opennebula_virtual_network.importtest 1234 ` + "`" + `` + "`" + `` + "`" + ` Verify that Terraform does not perform any change: ` + "`" + `` + "`" + `` + "`" + ` terraform plan ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"opennebula_acl":                   0,
		"opennebula_group":                 1,
		"opennebula_image":                 2,
		"opennebula_security group":        3,
		"opennebula_template":              4,
		"opennebula_virtual data center":   5,
		"opennebula_virtual machine":       6,
		"opennebula_virtual machine group": 7,
		"opennebula_virtual network":       8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
