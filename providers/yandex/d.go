package yandex

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_disk",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Compute disk.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "disk_id",
					Description: `(Optional) The ID of a specific disk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the disk. ~>`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Optional description of this disk.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the disk belongs to.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of the zone where the disk resides.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the disk, specified in Gb.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the source image that was used to create this disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Source snapshot that was used to create this disk.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the disk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the disk.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Map of labels applied to this disk.`,
				},
				resource.Attribute{
					Name:        "product_ids",
					Description: `License IDs that indicate which licenses are attached to this disk.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `IDs of instances to which this disk is attached.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Disk creation timestamp.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Optional description of this disk.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the disk belongs to.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of the zone where the disk resides.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the disk, specified in Gb.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the source image that was used to create this disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Source snapshot that was used to create this disk.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the disk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the disk.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Map of labels applied to this disk.`,
				},
				resource.Attribute{
					Name:        "product_ids",
					Description: `License IDs that indicate which licenses are attached to this disk.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `IDs of instances to which this disk is attached.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Disk creation timestamp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_image",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Compute image.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The ID of a specific image.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `(Optional) The family name of an image. Used to search the latest image in a family.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the image. ~>`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If a value is not provided, the default provider folder is used. ~>`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An optional description of this image.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The OS family name of the image.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `Minimum size of the disk which is created from this image.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image, specified in Gb.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image.`,
				},
				resource.Attribute{
					Name:        "product_ids",
					Description: `License IDs that indicate which licenses are attached to this image.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `Operating system type that the image contains.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Image creation timestamp.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `An optional description of this image.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The OS family name of the image.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `Minimum size of the disk which is created from this image.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image, specified in Gb.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image.`,
				},
				resource.Attribute{
					Name:        "product_ids",
					Description: `License IDs that indicate which licenses are attached to this image.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `Operating system type that the image contains.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Image creation timestamp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_instance",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Compute Instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) The ID of a specific instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the instance. ~>`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the instance.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `FQDN of the instance.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Availability zone where the instance resides.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata key/value pairs to make available from within the instance.`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `Type of virtual machine to create. Default is 'standard-v1'.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.memory",
					Description: `Memory size allocated for the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.cores",
					Description: `Number of CPU cores allocated for the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.core_fraction",
					Description: `Baseline performance for a core, set as a percent.`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `The boot disk for the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `The networks attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.ip_address",
					Description: `An internal IP address of the instance, either manually or dynamically assigned.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.nat_ip_address",
					Description: `An assigned external IP address if the instance has NAT enabled.`,
				},
				resource.Attribute{
					Name:        "secondary_disk",
					Description: `List of secondary disks attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduling_policy",
					Description: `Scheduling policy configuration. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `ID of the service account authorized for this instance.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Instance creation timestamp. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `Whether the disk is auto-deleted when the instance is deleted. The default value is false.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Name that can be used to access an attached disk under ` + "`" + `/dev/disk/by-id/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Access to the disk resource. By default a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `ID of the attached disk.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `Parameters used for creating a disk alongside the instance. The structure is documented below. The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the boot disk.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `A disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `A snapshot to initialize this disk from. The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `The index of the network interface, generated by the server.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `MAC address that is assigned to the network interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The private IP address to assign to the instance. If empty, the address is automatically assigned from the specified subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet to attach this interface to. The subnet must reside in the same zone where this instance was created.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Assigned for the instance's public address that is used to access the internet over NAT.`,
				},
				resource.Attribute{
					Name:        "nat_ip_address",
					Description: `Public IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "nat_ip_version",
					Description: `IP version for the public address. The ` + "`" + `secondary_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `Specifies whether the disk is auto-deleted when the instance is deleted.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `This value can be used to reference the device from within the instance for mounting, resizing, and so on.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Access to the Disk resource. By default, a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `ID of the disk that is attached to the instance. The ` + "`" + `scheduling_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) Specifies if the instance is preemptible. Defaults to false.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the instance.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `FQDN of the instance.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Availability zone where the instance resides.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata key/value pairs to make available from within the instance.`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `Type of virtual machine to create. Default is 'standard-v1'.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.memory",
					Description: `Memory size allocated for the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.cores",
					Description: `Number of CPU cores allocated for the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.core_fraction",
					Description: `Baseline performance for a core, set as a percent.`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `The boot disk for the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `The networks attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.ip_address",
					Description: `An internal IP address of the instance, either manually or dynamically assigned.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.nat_ip_address",
					Description: `An assigned external IP address if the instance has NAT enabled.`,
				},
				resource.Attribute{
					Name:        "secondary_disk",
					Description: `List of secondary disks attached to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduling_policy",
					Description: `Scheduling policy configuration. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `ID of the service account authorized for this instance.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Instance creation timestamp. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `Whether the disk is auto-deleted when the instance is deleted. The default value is false.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Name that can be used to access an attached disk under ` + "`" + `/dev/disk/by-id/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Access to the disk resource. By default a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `ID of the attached disk.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `Parameters used for creating a disk alongside the instance. The structure is documented below. The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the boot disk.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `A disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `A snapshot to initialize this disk from. The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `The index of the network interface, generated by the server.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `MAC address that is assigned to the network interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The private IP address to assign to the instance. If empty, the address is automatically assigned from the specified subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet to attach this interface to. The subnet must reside in the same zone where this instance was created.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Assigned for the instance's public address that is used to access the internet over NAT.`,
				},
				resource.Attribute{
					Name:        "nat_ip_address",
					Description: `Public IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "nat_ip_version",
					Description: `IP version for the public address. The ` + "`" + `secondary_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `Specifies whether the disk is auto-deleted when the instance is deleted.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `This value can be used to reference the device from within the instance for mounting, resizing, and so on.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Access to the Disk resource. By default, a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `ID of the disk that is attached to the instance. The ` + "`" + `scheduling_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) Specifies if the instance is preemptible. Defaults to false.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_instance_group",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Compute Instance Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_group_id",
					Description: `(Required) The ID of a specific instance group. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the instance group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the instance group.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `The ID of the folder that the instance group belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the instance group.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health check specification. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `Load balancing specification. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "deploy_policy",
					Description: `The deployment policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "allocation_policy",
					Description: `The allocation policy of the instance group by zone and region. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances in the specified instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `The instance template that the instance group belongs to. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `The ID of the service account authorized for this instance group.`,
				},
				resource.Attribute{
					Name:        "scale_policy",
					Description: `The scaling policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "load_balancer_state",
					Description: `Information about which entities can be attached to this load balancer. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The instance group creation timestamp. --- The ` + "`" + `load_balancer_state` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `The ID of the target group used for load balancing.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The status message of the target group. --- The ` + "`" + `scale_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fixed_scale",
					Description: `The fixed scaling policy of the instance group. The structure is documented below. --- The ` + "`" + `fixed_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the instance group. --- The ` + "`" + `instance_template` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the instance template.`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `The ID of the hardware platform configuration for the instance.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `The service account ID for the instance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The set of metadata ` + "`" + `key:value` + "`" + ` pairs assigned to this instance template. This includes custom metadata and predefined keys.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.memory",
					Description: `The memory size allocated to the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.cores",
					Description: `Number of CPU cores allocated to the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.core_fraction",
					Description: `Baseline core performance as a percent.`,
				},
				resource.Attribute{
					Name:        "scheduling_policy",
					Description: `The scheduling policy for the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `An array with the network interfaces that will be attached to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "secondary_disk",
					Description: `An array with the secondary disks that will be attached to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `The specifications for boot disks that will be attached to the instance. The structure is documented below. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `This value can be used to reference the device under ` + "`" + `/dev/disk/by-id/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The access mode to the disk resource. By default a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `The parameters used for creating a disk alongside the instance. The structure is documented below. --- The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot to initialize this disk from. --- The ` + "`" + `secondary_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `This value can be used to reference the device under ` + "`" + `/dev/disk/by-id/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The access mode to the disk resource. By default a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `The parameters used for creating a disk alongside the instance. The structure is documented below. --- The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot to initialize this disk from. --- The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The ID of the network.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `The IDs of the subnets.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `A public address that can be used to access the internet over NAT. --- The ` + "`" + `scheduling_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `Specifies if the instance is preemptible. Defaults to false. --- The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the managed instance.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The Fully Qualified Domain Name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The status message of the instance.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The ID of the availability zone where the instance resides.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `An array with the network interfaces attached to the managed instance. The structure is documented below. --- The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `The index of the network interface as generated by the server.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address assigned to the network interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The private IP address to assign to the instance. If empty, the address is automatically assigned from the specified subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet to attach this interface to. The subnet must reside in the same zone where this instance was created.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `The instance's public address for accessing the internet over NAT.`,
				},
				resource.Attribute{
					Name:        "nat_ip_address",
					Description: `The public IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "nat_ip_version",
					Description: `The IP version for the public address. --- The ` + "`" + `allocation_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. --- The ` + "`" + `deploy_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "max_unavailable",
					Description: `The maximum number of running instances that can be taken offline (stopped or deleted) at the same time during the update process.`,
				},
				resource.Attribute{
					Name:        "max_expansion",
					Description: `The maximum number of instances that can be temporarily allocated above the group's target size during the update process.`,
				},
				resource.Attribute{
					Name:        "max_deleting",
					Description: `The maximum number of instances that can be deleted at the same time.`,
				},
				resource.Attribute{
					Name:        "max_creating",
					Description: `The maximum number of instances that can be created at the same time.`,
				},
				resource.Attribute{
					Name:        "startup_duration",
					Description: `The amount of time in seconds to allow for an instance to start. Instance will be considered up and running (and start receiving traffic) only after the startup_duration has elapsed and all health checks are passed. --- The ` + "`" + `load_balancer` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_group_name",
					Description: `The name of the target group.`,
				},
				resource.Attribute{
					Name:        "target_group_description",
					Description: `A description of the target group.`,
				},
				resource.Attribute{
					Name:        "target_group_labels",
					Description: `A set of key/value label pairs. --- The ` + "`" + `health_check` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The interval between health checks in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout for the managed instance to return a response for the health check in seconds.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of successful health checks before the managed instance is declared healthy.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of failed health checks before the managed instance is declared unhealthy.`,
				},
				resource.Attribute{
					Name:        "tcp_options",
					Description: `TCP check options. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "http_options",
					Description: `HTTP check options. The structure is documented below. --- The ` + "`" + `http_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used for HTTP health checks.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The URL path used for health check requests. --- The ` + "`" + `tcp_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port to use for TCP health checks.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the instance group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the instance group.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `The ID of the folder that the instance group belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the instance group.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health check specification. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `Load balancing specification. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "deploy_policy",
					Description: `The deployment policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "allocation_policy",
					Description: `The allocation policy of the instance group by zone and region. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances in the specified instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `The instance template that the instance group belongs to. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `The ID of the service account authorized for this instance group.`,
				},
				resource.Attribute{
					Name:        "scale_policy",
					Description: `The scaling policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "load_balancer_state",
					Description: `Information about which entities can be attached to this load balancer. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The instance group creation timestamp. --- The ` + "`" + `load_balancer_state` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `The ID of the target group used for load balancing.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The status message of the target group. --- The ` + "`" + `scale_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fixed_scale",
					Description: `The fixed scaling policy of the instance group. The structure is documented below. --- The ` + "`" + `fixed_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the instance group. --- The ` + "`" + `instance_template` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the instance template.`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `The ID of the hardware platform configuration for the instance.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `The service account ID for the instance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The set of metadata ` + "`" + `key:value` + "`" + ` pairs assigned to this instance template. This includes custom metadata and predefined keys.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.memory",
					Description: `The memory size allocated to the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.cores",
					Description: `Number of CPU cores allocated to the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.core_fraction",
					Description: `Baseline core performance as a percent.`,
				},
				resource.Attribute{
					Name:        "scheduling_policy",
					Description: `The scheduling policy for the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `An array with the network interfaces that will be attached to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "secondary_disk",
					Description: `An array with the secondary disks that will be attached to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `The specifications for boot disks that will be attached to the instance. The structure is documented below. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `This value can be used to reference the device under ` + "`" + `/dev/disk/by-id/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The access mode to the disk resource. By default a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `The parameters used for creating a disk alongside the instance. The structure is documented below. --- The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot to initialize this disk from. --- The ` + "`" + `secondary_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `This value can be used to reference the device under ` + "`" + `/dev/disk/by-id/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The access mode to the disk resource. By default a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `The parameters used for creating a disk alongside the instance. The structure is documented below. --- The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot to initialize this disk from. --- The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The ID of the network.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `The IDs of the subnets.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `A public address that can be used to access the internet over NAT. --- The ` + "`" + `scheduling_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `Specifies if the instance is preemptible. Defaults to false. --- The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the managed instance.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The Fully Qualified Domain Name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The status message of the instance.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The ID of the availability zone where the instance resides.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `An array with the network interfaces attached to the managed instance. The structure is documented below. --- The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `The index of the network interface as generated by the server.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address assigned to the network interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The private IP address to assign to the instance. If empty, the address is automatically assigned from the specified subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet to attach this interface to. The subnet must reside in the same zone where this instance was created.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `The instance's public address for accessing the internet over NAT.`,
				},
				resource.Attribute{
					Name:        "nat_ip_address",
					Description: `The public IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "nat_ip_version",
					Description: `The IP version for the public address. --- The ` + "`" + `allocation_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. --- The ` + "`" + `deploy_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "max_unavailable",
					Description: `The maximum number of running instances that can be taken offline (stopped or deleted) at the same time during the update process.`,
				},
				resource.Attribute{
					Name:        "max_expansion",
					Description: `The maximum number of instances that can be temporarily allocated above the group's target size during the update process.`,
				},
				resource.Attribute{
					Name:        "max_deleting",
					Description: `The maximum number of instances that can be deleted at the same time.`,
				},
				resource.Attribute{
					Name:        "max_creating",
					Description: `The maximum number of instances that can be created at the same time.`,
				},
				resource.Attribute{
					Name:        "startup_duration",
					Description: `The amount of time in seconds to allow for an instance to start. Instance will be considered up and running (and start receiving traffic) only after the startup_duration has elapsed and all health checks are passed. --- The ` + "`" + `load_balancer` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_group_name",
					Description: `The name of the target group.`,
				},
				resource.Attribute{
					Name:        "target_group_description",
					Description: `A description of the target group.`,
				},
				resource.Attribute{
					Name:        "target_group_labels",
					Description: `A set of key/value label pairs. --- The ` + "`" + `health_check` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The interval between health checks in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout for the managed instance to return a response for the health check in seconds.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of successful health checks before the managed instance is declared healthy.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of failed health checks before the managed instance is declared unhealthy.`,
				},
				resource.Attribute{
					Name:        "tcp_options",
					Description: `TCP check options. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "http_options",
					Description: `HTTP check options. The structure is documented below. --- The ` + "`" + `http_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used for HTTP health checks.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The URL path used for health check requests. --- The ` + "`" + `tcp_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port to use for TCP health checks.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Compute Snapshot.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of a specific snapshot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the snapshot. ~>`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An optional description of this snapshot.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the snapshot belongs to.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `The size of the snapshot, specified in Gb.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Minimum required size of the disk which is created from this snapshot.`,
				},
				resource.Attribute{
					Name:        "source_disk_id",
					Description: `ID of the source disk.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this snapshot.`,
				},
				resource.Attribute{
					Name:        "product_ids",
					Description: `License IDs that indicate which licenses are attached to this snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Snapshot creation timestamp.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `An optional description of this snapshot.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the snapshot belongs to.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `The size of the snapshot, specified in Gb.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Minimum required size of the disk which is created from this snapshot.`,
				},
				resource.Attribute{
					Name:        "source_disk_id",
					Description: `ID of the source disk.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this snapshot.`,
				},
				resource.Attribute{
					Name:        "product_ids",
					Description: `License IDs that indicate which licenses are attached to this snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Snapshot creation timestamp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_policy",
			Category:         "Data Sources",
			ShortDescription: `Generates an IAM policy that can be referenced by other resources and applied to them.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_data",
					Description: `The above bindings serialized in a format suitable for referencing from a resource that supports IAM. [IAM]: https://cloud.yandex.com/docs/iam/ [IAM Roles]: https://cloud.yandex.com/docs/iam/concepts/access-control/roles`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_data",
					Description: `The above bindings serialized in a format suitable for referencing from a resource that supports IAM. [IAM]: https://cloud.yandex.com/docs/iam/ [IAM Roles]: https://cloud.yandex.com/docs/iam/concepts/access-control/roles`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_role",
			Category:         "Data Sources",
			ShortDescription: `Generates an IAM role that can be referenced by other resources, applying the role to them.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_data",
					Description: `The above bindings serialized in a format suitable for referencing from a resource that supports IAM. [IAM]: https://cloud.yandex.com/docs/iam/ [IAM Roles]: https://cloud.yandex.com/docs/iam/concepts/access-control/roles`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role_data",
					Description: `The above bindings serialized in a format suitable for referencing from a resource that supports IAM. [IAM]: https://cloud.yandex.com/docs/iam/ [IAM Roles]: https://cloud.yandex.com/docs/iam/concepts/access-control/roles`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex IAM service account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_user",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex IAM user account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `ID of IAM user account.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `Login name of IAM user account.`,
				},
				resource.Attribute{
					Name:        "default_email",
					Description: `Email address of user account. [IAM Users]: https://cloud.yandex.com/docs/iam/concepts/users/users#passport`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `ID of IAM user account.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `Login name of IAM user account.`,
				},
				resource.Attribute{
					Name:        "default_email",
					Description: `Email address of user account. [IAM Users]: https://cloud.yandex.com/docs/iam/concepts/users/users#passport`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_cloud",
			Category:         "Data Sources",
			ShortDescription: `Retrieve Yandex RM Cloud details.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_id",
					Description: `(Optional) ID of the cloud.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the cloud. ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cloud.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the cloud.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Cloud creation timestamp.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cloud.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the cloud.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Cloud creation timestamp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_folder",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex RM Folder.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the folder.`,
				},
				resource.Attribute{
					Name:        "cloud_id",
					Description: `ID of the cloud that contains the folder.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the folder.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this folder.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Folder creation timestamp. [Folder]: https://cloud.yandex.com/docs/resource-manager/concepts/resources-hierarchy#folder`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the folder.`,
				},
				resource.Attribute{
					Name:        "cloud_id",
					Description: `ID of the cloud that contains the folder.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the folder.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels applied to this folder.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Folder creation timestamp. [Folder]: https://cloud.yandex.com/docs/resource-manager/concepts/resources-hierarchy#folder`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_vpc_network",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex VPC network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the network.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels assigned to this network.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this network. [VPC Networks]: https://cloud.yandex.com/docs/vpc/concepts/network`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the network.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels assigned to this network.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this network. [VPC Networks]: https://cloud.yandex.com/docs/vpc/concepts/network`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_vpc_route_table",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex VPC route table.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) - Name of the route table. ~>`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the route table.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network this route table belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this route table.`,
				},
				resource.Attribute{
					Name:        "static_route",
					Description: `List of static route records of the route table. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this route table. The ` + "`" + `static_route` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "destination_prefix",
					Description: `Route prefix in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "next_hop_address",
					Description: `Address of the next hop. [VPC Route Table]: https://cloud.yandex.com/docs/vpc/concepts/`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the route table.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network this route table belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this route table.`,
				},
				resource.Attribute{
					Name:        "static_route",
					Description: `List of static route records of the route table. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this route table. The ` + "`" + `static_route` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "destination_prefix",
					Description: `Route prefix in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "next_hop_address",
					Description: `Address of the next hop. [VPC Route Table]: https://cloud.yandex.com/docs/vpc/concepts/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_vpc_subnet",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex VPC subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) - Name of the subnet. ~>`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the subnet.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network this subnet belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this subnet.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Name of the availability zone for this subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the route table to assign to this subnet.`,
				},
				resource.Attribute{
					Name:        "v4_cidr_blocks",
					Description: `The blocks of internal IPv4 addresses owned by this subnet.`,
				},
				resource.Attribute{
					Name:        "v6_cidr_blocks",
					Description: `The blocks of internal IPv6 addresses owned by this subnet.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this subnet. ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the subnet.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder that the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network this subnet belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this subnet.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Name of the availability zone for this subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the route table to assign to this subnet.`,
				},
				resource.Attribute{
					Name:        "v4_cidr_blocks",
					Description: `The blocks of internal IPv4 addresses owned by this subnet.`,
				},
				resource.Attribute{
					Name:        "v6_cidr_blocks",
					Description: `The blocks of internal IPv6 addresses owned by this subnet.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this subnet. ~>`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"yandex_compute_disk":           0,
		"yandex_compute_image":          1,
		"yandex_compute_instance":       2,
		"yandex_compute_instance_group": 3,
		"yandex_compute_snapshot":       4,
		"yandex_iam_policy":             5,
		"yandex_iam_role":               6,
		"yandex_iam_service_account":    7,
		"yandex_iam_user":               8,
		"yandex_resourcemanager_cloud":  9,
		"yandex_resourcemanager_folder": 10,
		"yandex_vpc_network":            11,
		"yandex_vpc_route_table":        12,
		"yandex_vpc_subnet":             13,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
