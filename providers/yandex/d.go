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
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ~>`,
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
					Description: `(Optional) Name of the instance.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ~>`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the instance.`,
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
					Description: `A set of key/value label pairs assigned to the instance.`,
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
					Name:        "resources.0.gpus",
					Description: `Number of GPU cores allocated for the instance.`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `The boot disk for the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_acceleration_type",
					Description: `Type of network acceleration. The default is ` + "`" + `standard` + "`" + `. Values: ` + "`" + `standard` + "`" + `, ` + "`" + `software_accelerated` + "`" + ``,
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
					Name:        "ipv4",
					Description: `Show if IPv4 address is assigned to the network interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The assignd private IP address to the network interface.`,
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
					Name:        "fqdn",
					Description: `FQDN of the instance.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Availability zone where the instance resides.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs assigned to the instance.`,
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
					Name:        "resources.0.gpus",
					Description: `Number of GPU cores allocated for the instance.`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `The boot disk for the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_acceleration_type",
					Description: `Type of network acceleration. The default is ` + "`" + `standard` + "`" + `. Values: ` + "`" + `standard` + "`" + `, ` + "`" + `software_accelerated` + "`" + ``,
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
					Name:        "ipv4",
					Description: `Show if IPv4 address is assigned to the network interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The assignd private IP address to the network interface.`,
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
					Description: `The ID of a specific instance group. ## Attributes Reference`,
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
					Description: `The instance group creation timestamp.`,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `A set of key/value variables pairs to assign to the instance group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance group. --- The ` + "`" + `load_balancer_state` + "`" + ` block supports:`,
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
					Description: `The fixed scaling policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "auto_scale",
					Description: `The auto scaling policy of the instance group. The structure is documented below. --- The ` + "`" + `fixed_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the instance group. --- The ` + "`" + `auto_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "initial_size",
					Description: `The initial number of instances in the instance group.`,
				},
				resource.Attribute{
					Name:        "measurement_duration",
					Description: `The amount of time, in seconds, that metrics are averaged for. If the average value at the end of the interval is higher than the ` + "`" + `cpu_utilization_target` + "`" + `, the instance group will increase the number of virtual machines in the group.`,
				},
				resource.Attribute{
					Name:        "min_zone_size",
					Description: `The minimum number of virtual machines in a single availability zone.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum number of virtual machines in the group.`,
				},
				resource.Attribute{
					Name:        "warmup_duration",
					Description: `The warm-up time of the virtual machine, in seconds. During this time, traffic is fed to the virtual machine, but load metrics are not taken into account.`,
				},
				resource.Attribute{
					Name:        "stabilization_duration",
					Description: `The minimum time interval, in seconds, to monitor the load before an instance group can reduce the number of virtual machines in the group. During this time, the group will not decrease even if the average load falls below the value of ` + "`" + `cpu_utilization_target` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_utilization_target",
					Description: `Target CPU load level.`,
				},
				resource.Attribute{
					Name:        "custom_rule",
					Description: `A list of custom rules. The structure is documented below. --- The ` + "`" + `custom_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `Rule type: ` + "`" + `UTILIZATION` + "`" + ` - This type means that the metric applies to one instance. First, Instance Groups calculates the average metric value for each instance, then averages the values for instances in one availability zone. This type of metric must have the ` + "`" + `instance_id` + "`" + ` label. ` + "`" + `WORKLOAD` + "`" + ` - This type means that the metric applies to instances in one availability zone. This type of metric must have the ` + "`" + `zone_id` + "`" + ` label.`,
				},
				resource.Attribute{
					Name:        "metric_type",
					Description: `Metric type, ` + "`" + `GAUGE` + "`" + ` or ` + "`" + `COUNTER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `The name of metric.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Target metric value level.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels of metric. --- The ` + "`" + `instance_template` + "`" + ` block supports:`,
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
					Name:        "resources.0.gpus",
					Description: `Number of GPU cores allocated to the instance.`,
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
					Description: `The specifications for boot disk that will be attached to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_settings",
					Description: `Network acceleration settings. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name template of the instance.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname temaplate for the instance. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
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
					Name:        "ipv4",
					Description: `Is IPv4 address assigned.`,
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
					Description: `An array with the network interfaces attached to the managed instance. The structure is documented below.`,
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
					Description: `A set of key/value label pairs.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The status message of the target group. --- The ` + "`" + `health_check` + "`" + ` block supports:`,
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
					Description: `The port to use for TCP health checks. --- The ` + "`" + `network_settings` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Network acceleration type. By default a network is in ` + "`" + `STANDARD` + "`" + ` mode.`,
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
					Description: `The instance group creation timestamp.`,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `A set of key/value variables pairs to assign to the instance group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance group. --- The ` + "`" + `load_balancer_state` + "`" + ` block supports:`,
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
					Description: `The fixed scaling policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "auto_scale",
					Description: `The auto scaling policy of the instance group. The structure is documented below. --- The ` + "`" + `fixed_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the instance group. --- The ` + "`" + `auto_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "initial_size",
					Description: `The initial number of instances in the instance group.`,
				},
				resource.Attribute{
					Name:        "measurement_duration",
					Description: `The amount of time, in seconds, that metrics are averaged for. If the average value at the end of the interval is higher than the ` + "`" + `cpu_utilization_target` + "`" + `, the instance group will increase the number of virtual machines in the group.`,
				},
				resource.Attribute{
					Name:        "min_zone_size",
					Description: `The minimum number of virtual machines in a single availability zone.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum number of virtual machines in the group.`,
				},
				resource.Attribute{
					Name:        "warmup_duration",
					Description: `The warm-up time of the virtual machine, in seconds. During this time, traffic is fed to the virtual machine, but load metrics are not taken into account.`,
				},
				resource.Attribute{
					Name:        "stabilization_duration",
					Description: `The minimum time interval, in seconds, to monitor the load before an instance group can reduce the number of virtual machines in the group. During this time, the group will not decrease even if the average load falls below the value of ` + "`" + `cpu_utilization_target` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_utilization_target",
					Description: `Target CPU load level.`,
				},
				resource.Attribute{
					Name:        "custom_rule",
					Description: `A list of custom rules. The structure is documented below. --- The ` + "`" + `custom_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `Rule type: ` + "`" + `UTILIZATION` + "`" + ` - This type means that the metric applies to one instance. First, Instance Groups calculates the average metric value for each instance, then averages the values for instances in one availability zone. This type of metric must have the ` + "`" + `instance_id` + "`" + ` label. ` + "`" + `WORKLOAD` + "`" + ` - This type means that the metric applies to instances in one availability zone. This type of metric must have the ` + "`" + `zone_id` + "`" + ` label.`,
				},
				resource.Attribute{
					Name:        "metric_type",
					Description: `Metric type, ` + "`" + `GAUGE` + "`" + ` or ` + "`" + `COUNTER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `The name of metric.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Target metric value level.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels of metric. --- The ` + "`" + `instance_template` + "`" + ` block supports:`,
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
					Name:        "resources.0.gpus",
					Description: `Number of GPU cores allocated to the instance.`,
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
					Description: `The specifications for boot disk that will be attached to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_settings",
					Description: `Network acceleration settings. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name template of the instance.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname temaplate for the instance. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
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
					Name:        "ipv4",
					Description: `Is IPv4 address assigned.`,
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
					Description: `An array with the network interfaces attached to the managed instance. The structure is documented below.`,
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
					Description: `A set of key/value label pairs.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The status message of the target group. --- The ` + "`" + `health_check` + "`" + ` block supports:`,
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
					Description: `The port to use for TCP health checks. --- The ` + "`" + `network_settings` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Network acceleration type. By default a network is in ` + "`" + `STANDARD` + "`" + ` mode.`,
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
			Type:             "yandex_container_registry",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Container Registry.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "registry_id",
					Description: `(Optional) The ID of a specific registry.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the registry.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the registry.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this registry.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this registry.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Status of the registry.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this registry.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this registry.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_dataproc_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Data Proc cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) The ID of the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Data Proc cluster. ~>`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Name of the Object Storage bucket used for Data Proc jobs.`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `Configuration and resources of the cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The Data Proc cluster creation timestamp.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs assigned to the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `Service account used by the Data Proc agent to access resources of Yandex.Cloud.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the availability zone where the cluster resides. --- The ` + "`" + `cluster_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `Version of Data Proc image.`,
				},
				resource.Attribute{
					Name:        "hadoop",
					Description: `Data Proc specific options. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "subcluster_spec",
					Description: `Configuration of the Data Proc subcluster. The structure is documented below. --- The ` + "`" + `hadoop` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `List of services launched on Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `A set of key/value pairs used to configure cluster services.`,
				},
				resource.Attribute{
					Name:        "ssh_public_keys",
					Description: `List of SSH public keys distributed to the hosts of the cluster. --- The ` + "`" + `subcluster_spec` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Data Proc subcluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Data Proc subcluster.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role of the subcluster in the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to each host of the Data Proc subcluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which hosts of the subcluster belong.`,
				},
				resource.Attribute{
					Name:        "hosts_count",
					Description: `Number of hosts within Data Proc subcluster. --- The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resource_preset_id",
					Description: `The ID of the preset for computational resources available to a host. All available presets are listed in the [documentation](https://cloud.yandex.com/docs/data-proc/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `Type of the storage of a host.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `Name of the Object Storage bucket used for Data Proc jobs.`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `Configuration and resources of the cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The Data Proc cluster creation timestamp.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs assigned to the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `Service account used by the Data Proc agent to access resources of Yandex.Cloud.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the availability zone where the cluster resides. --- The ` + "`" + `cluster_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `Version of Data Proc image.`,
				},
				resource.Attribute{
					Name:        "hadoop",
					Description: `Data Proc specific options. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "subcluster_spec",
					Description: `Configuration of the Data Proc subcluster. The structure is documented below. --- The ` + "`" + `hadoop` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `List of services launched on Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `A set of key/value pairs used to configure cluster services.`,
				},
				resource.Attribute{
					Name:        "ssh_public_keys",
					Description: `List of SSH public keys distributed to the hosts of the cluster. --- The ` + "`" + `subcluster_spec` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Data Proc subcluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Data Proc subcluster.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role of the subcluster in the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to each host of the Data Proc subcluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which hosts of the subcluster belong.`,
				},
				resource.Attribute{
					Name:        "hosts_count",
					Description: `Number of hosts within Data Proc subcluster. --- The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resource_preset_id",
					Description: `The ID of the preset for computational resources available to a host. All available presets are listed in the [documentation](https://cloud.yandex.com/docs/data-proc/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `Type of the storage of a host.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_function",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Cloud Function.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `Runtime for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "entrypoint",
					Description: `Entrypoint for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory in megabytes (`,
				},
				resource.Attribute{
					Name:        "execution_timeout",
					Description: `Execution timeout in seconds for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `Service account ID for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `A set of key/value environment variables for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for Yandex Cloud Function. Tag "$latest" isn't returned.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version for Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Image size for Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "loggroup_id",
					Description: `Log group ID size for Yandex Cloud Function.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `Runtime for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "entrypoint",
					Description: `Entrypoint for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory in megabytes (`,
				},
				resource.Attribute{
					Name:        "execution_timeout",
					Description: `Execution timeout in seconds for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `Service account ID for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `A set of key/value environment variables for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for Yandex Cloud Function. Tag "$latest" isn't returned.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version for Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Image size for Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "loggroup_id",
					Description: `Log group ID size for Yandex Cloud Function.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_function_trigger",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Cloud Functions Trigger.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder ID for the Yandex Cloud Functions Trigger ~>`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `Folder ID for the Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function",
					Description: `[Yandex.Cloud Function](https://cloud.yandex.com/docs/functions/concepts/function) settings definition for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.id",
					Description: `Yandex.Cloud Function ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.service_account_id",
					Description: `Service account ID for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.tag",
					Description: `Tag for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.retry_attempts",
					Description: `Retry attempts for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.retry_interval",
					Description: `Retry interval in seconds for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "dlq",
					Description: `Dead Letter Queue settings definition for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "dlq.0.queue_id",
					Description: `Queue ID for Dead Letter Queue for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "dlq.0.service_account_id",
					Description: `Service Account ID for Dead Letter Queue for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "iot",
					Description: `[IoT](https://cloud.yandex.com/docs/functions/concepts/trigger/iot-core-trigger) settings definition for Yandex Cloud Functions Trigger, if present`,
				},
				resource.Attribute{
					Name:        "iot.0.registry_id",
					Description: `IoT Registry ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "iot.0.device_id",
					Description: `IoT Device ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "iot.0.topic",
					Description: `IoT Topic for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue",
					Description: `[Message Queue](https://cloud.yandex.com/docs/functions/concepts/trigger/ymq-trigger) settings definition for Yandex Cloud Functions Trigger, if present`,
				},
				resource.Attribute{
					Name:        "message_queue.0.queue_id",
					Description: `Message Queue ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.service_account_id",
					Description: `Message Queue Service Account ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.batch_cutoff",
					Description: `Batch Duration in seconds for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.batch_size",
					Description: `Batch Size for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.visibility_timeout",
					Description: `Visibility timeout for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage",
					Description: `[Object Storage](https://cloud.yandex.com/docs/functions/concepts/trigger/os-trigger) settings definition for Yandex Cloud Functions Trigger, if present`,
				},
				resource.Attribute{
					Name:        "object_storage.0.bucket_id",
					Description: `Object Storage Bucket ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.prefix",
					Description: `Prefix for Object Storage for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.suffix",
					Description: `Suffix for Object Storage for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.create",
					Description: `Boolean flag for setting create event for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.update",
					Description: `Boolean flag for setting update event for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.delete",
					Description: `Boolean flag for setting delete event for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `[Timer](https://cloud.yandex.com/docs/functions/concepts/trigger/timer) settings definition for Yandex Cloud Functions Trigger, if present`,
				},
				resource.Attribute{
					Name:        "timer.0.cron_expression",
					Description: `Cron expression for timer for Yandex Cloud Functions Trigger`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `Folder ID for the Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function",
					Description: `[Yandex.Cloud Function](https://cloud.yandex.com/docs/functions/concepts/function) settings definition for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.id",
					Description: `Yandex.Cloud Function ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.service_account_id",
					Description: `Service account ID for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.tag",
					Description: `Tag for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.retry_attempts",
					Description: `Retry attempts for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.retry_interval",
					Description: `Retry interval in seconds for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "dlq",
					Description: `Dead Letter Queue settings definition for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "dlq.0.queue_id",
					Description: `Queue ID for Dead Letter Queue for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "dlq.0.service_account_id",
					Description: `Service Account ID for Dead Letter Queue for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "iot",
					Description: `[IoT](https://cloud.yandex.com/docs/functions/concepts/trigger/iot-core-trigger) settings definition for Yandex Cloud Functions Trigger, if present`,
				},
				resource.Attribute{
					Name:        "iot.0.registry_id",
					Description: `IoT Registry ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "iot.0.device_id",
					Description: `IoT Device ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "iot.0.topic",
					Description: `IoT Topic for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue",
					Description: `[Message Queue](https://cloud.yandex.com/docs/functions/concepts/trigger/ymq-trigger) settings definition for Yandex Cloud Functions Trigger, if present`,
				},
				resource.Attribute{
					Name:        "message_queue.0.queue_id",
					Description: `Message Queue ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.service_account_id",
					Description: `Message Queue Service Account ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.batch_cutoff",
					Description: `Batch Duration in seconds for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.batch_size",
					Description: `Batch Size for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.visibility_timeout",
					Description: `Visibility timeout for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage",
					Description: `[Object Storage](https://cloud.yandex.com/docs/functions/concepts/trigger/os-trigger) settings definition for Yandex Cloud Functions Trigger, if present`,
				},
				resource.Attribute{
					Name:        "object_storage.0.bucket_id",
					Description: `Object Storage Bucket ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.prefix",
					Description: `Prefix for Object Storage for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.suffix",
					Description: `Suffix for Object Storage for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.create",
					Description: `Boolean flag for setting create event for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.update",
					Description: `Boolean flag for setting update event for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.delete",
					Description: `Boolean flag for setting delete event for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `[Timer](https://cloud.yandex.com/docs/functions/concepts/trigger/timer) settings definition for Yandex Cloud Functions Trigger, if present`,
				},
				resource.Attribute{
					Name:        "timer.0.cron_expression",
					Description: `Cron expression for timer for Yandex Cloud Functions Trigger`,
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
					Description: `Email address of user account. [IAM User]: https://cloud.yandex.com/docs/iam/concepts/#passport`,
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
					Description: `Email address of user account. [IAM User]: https://cloud.yandex.com/docs/iam/concepts/#passport`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iot_core",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex.Cloud IoT Core Device.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "registry_id",
					Description: `IoT Core Registry ID for the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `A set of key/value aliases pairs to assign to the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A set of certificate's fingerprints for the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "passwords",
					Description: `A set of passwords's id for the IoT Core Device`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "registry_id",
					Description: `IoT Core Registry ID for the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `A set of key/value aliases pairs to assign to the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A set of certificate's fingerprints for the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "passwords",
					Description: `A set of passwords's id for the IoT Core Device`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iot_registry",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex.Cloud IoT Core Registry.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the IoT Core Registry`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the IoT Core Registry.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the IoT Core Registry`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A set of certificate's fingerprints for the IoT Core Registry`,
				},
				resource.Attribute{
					Name:        "passwords",
					Description: `A set of passwords's id for the IoT Core Registry`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the IoT Core Registry`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the IoT Core Registry.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the IoT Core Registry`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A set of certificate's fingerprints for the IoT Core Registry`,
				},
				resource.Attribute{
					Name:        "passwords",
					Description: `A set of passwords's id for the IoT Core Registry`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_kubernetes_cluster",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) ID of a specific Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of a specific Kubernetes cluster. ~>`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The ID of the cluster network.`,
				},
				resource.Attribute{
					Name:        "cluster_ipv4_range",
					Description: `IP range for allocating pod addresses.`,
				},
				resource.Attribute{
					Name:        "node_ipv4_cidr_mask_size",
					Description: `Size of the masks that are assigned to each node in the cluster.`,
				},
				resource.Attribute{
					Name:        "service_ipv4_range",
					Description: `IP range Kubernetes services Kubernetes cluster IP addresses will be allocated from`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `Service account to be used for provisioning Compute Cloud and VPC resources for Kubernetes cluster. Selected service account should have ` + "`" + `edit` + "`" + ` role on the folder where the Kubernetes cluster will be located and on the folder where selected network resides.`,
				},
				resource.Attribute{
					Name:        "node_service_account_id",
					Description: `Service account to be used by the worker nodes of the Kubernetes cluster to access Container Registry or to push node logs and metrics.`,
				},
				resource.Attribute{
					Name:        "release_channel",
					Description: `Cluster release channel.`,
				},
				resource.Attribute{
					Name:        "master",
					Description: `Kubernetes master configuration options. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The Kubernetes cluster creation timestamp.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Health of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "network_policy_provider",
					Description: `Network policy provider for the cluster, if present. Possible values: ` + "`" + `CALICO` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_provider",
					Description: `cluster KMS provider parameters. --- The ` + "`" + `master` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of Kubernetes master.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Boolean flag. When ` + "`" + `true` + "`" + `, Kubernetes master have visible ipv4 address.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `Maintenance policy for Kubernetes master. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "zonal",
					Description: `Information about cluster zonal master. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "regional",
					Description: `Information about cluster regional master. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "internal_v4_address",
					Description: `An IPv4 internal network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "external_v4_address",
					Description: `An IPv4 external network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "internal_v4_endpoint",
					Description: `Internal endpoint that can be used to connect to the master from cloud networks.`,
				},
				resource.Attribute{
					Name:        "external_v4_endpoint",
					Description: `External endpoint that can be used to access Kubernetes cluster API from the internet (outside of the cloud).`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `PEM-encoded public certificate that is the root of trust for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "version_info",
					Description: `Information about cluster version. The structure is documented below. --- The ` + "`" + `maintenance_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_upgrade",
					Description: `Boolean flag that specifies if master can be upgraded automatically.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `Set of day intervals, when maintenance is allowed, when update for master is allowed. When omitted, it defaults to any time. Weekly maintenance policy expands to one element, with only two fields set: ` + "`" + `start_time` + "`" + `, ` + "`" + `duration` + "`" + `, and ` + "`" + `day` + "`" + ` field omitted. Daily maintenance policy expands to list of elements, with all fields set, that specify time interval for selected days. Only one interval is possible for any week day. Some days can be omitted, when there is no allowed interval for maintenance specified. --- The ` + "`" + `zonal` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of the availability zone where the master compute instance resides. --- The ` + "`" + `regional` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `ID of the availability region where the master compute instances resides. --- The ` + "`" + `version_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "current_version",
					Description: `Current Kubernetes version, major.minor (e.g. 1.15).`,
				},
				resource.Attribute{
					Name:        "new_revision_available",
					Description: `True/false flag. Newer revisions may include Kubernetes patches (e.g 1.15.1 -> 1.15.2) as well as some internal component updates - new features or bug fixes in yandex-specific components either on the master or nodes.`,
				},
				resource.Attribute{
					Name:        "new_revision_summary",
					Description: `Human readable description of the changes to be applied when updating to the latest revision. Empty if new_revision_available is false.`,
				},
				resource.Attribute{
					Name:        "version_deprecated",
					Description: `True/false flag. The current version is on the deprecation schedule, component (master or node group) should be upgraded. --- The ` + "`" + `kms_provider` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `KMS key ID. ---`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `A description of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The ID of the cluster network.`,
				},
				resource.Attribute{
					Name:        "cluster_ipv4_range",
					Description: `IP range for allocating pod addresses.`,
				},
				resource.Attribute{
					Name:        "node_ipv4_cidr_mask_size",
					Description: `Size of the masks that are assigned to each node in the cluster.`,
				},
				resource.Attribute{
					Name:        "service_ipv4_range",
					Description: `IP range Kubernetes services Kubernetes cluster IP addresses will be allocated from`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `Service account to be used for provisioning Compute Cloud and VPC resources for Kubernetes cluster. Selected service account should have ` + "`" + `edit` + "`" + ` role on the folder where the Kubernetes cluster will be located and on the folder where selected network resides.`,
				},
				resource.Attribute{
					Name:        "node_service_account_id",
					Description: `Service account to be used by the worker nodes of the Kubernetes cluster to access Container Registry or to push node logs and metrics.`,
				},
				resource.Attribute{
					Name:        "release_channel",
					Description: `Cluster release channel.`,
				},
				resource.Attribute{
					Name:        "master",
					Description: `Kubernetes master configuration options. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The Kubernetes cluster creation timestamp.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Health of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "network_policy_provider",
					Description: `Network policy provider for the cluster, if present. Possible values: ` + "`" + `CALICO` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_provider",
					Description: `cluster KMS provider parameters. --- The ` + "`" + `master` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of Kubernetes master.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Boolean flag. When ` + "`" + `true` + "`" + `, Kubernetes master have visible ipv4 address.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `Maintenance policy for Kubernetes master. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "zonal",
					Description: `Information about cluster zonal master. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "regional",
					Description: `Information about cluster regional master. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "internal_v4_address",
					Description: `An IPv4 internal network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "external_v4_address",
					Description: `An IPv4 external network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "internal_v4_endpoint",
					Description: `Internal endpoint that can be used to connect to the master from cloud networks.`,
				},
				resource.Attribute{
					Name:        "external_v4_endpoint",
					Description: `External endpoint that can be used to access Kubernetes cluster API from the internet (outside of the cloud).`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `PEM-encoded public certificate that is the root of trust for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "version_info",
					Description: `Information about cluster version. The structure is documented below. --- The ` + "`" + `maintenance_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_upgrade",
					Description: `Boolean flag that specifies if master can be upgraded automatically.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `Set of day intervals, when maintenance is allowed, when update for master is allowed. When omitted, it defaults to any time. Weekly maintenance policy expands to one element, with only two fields set: ` + "`" + `start_time` + "`" + `, ` + "`" + `duration` + "`" + `, and ` + "`" + `day` + "`" + ` field omitted. Daily maintenance policy expands to list of elements, with all fields set, that specify time interval for selected days. Only one interval is possible for any week day. Some days can be omitted, when there is no allowed interval for maintenance specified. --- The ` + "`" + `zonal` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of the availability zone where the master compute instance resides. --- The ` + "`" + `regional` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `ID of the availability region where the master compute instances resides. --- The ` + "`" + `version_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "current_version",
					Description: `Current Kubernetes version, major.minor (e.g. 1.15).`,
				},
				resource.Attribute{
					Name:        "new_revision_available",
					Description: `True/false flag. Newer revisions may include Kubernetes patches (e.g 1.15.1 -> 1.15.2) as well as some internal component updates - new features or bug fixes in yandex-specific components either on the master or nodes.`,
				},
				resource.Attribute{
					Name:        "new_revision_summary",
					Description: `Human readable description of the changes to be applied when updating to the latest revision. Empty if new_revision_available is false.`,
				},
				resource.Attribute{
					Name:        "version_deprecated",
					Description: `True/false flag. The current version is on the deprecation schedule, component (master or node group) should be upgraded. --- The ` + "`" + `kms_provider` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `KMS key ID. ---`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_kubernetes_node_group",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Kubernetes Node Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "node_group_id",
					Description: `(Optional) ID of a specific Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of a specific Kubernetes node group. ~>`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the Kubernetes cluster that this node group belongs to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs assigned to the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The Kubernetes node group creation timestamp.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `Template used to create compute instances in this Kubernetes node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scale_policy",
					Description: `Scale policy of the node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "allocation_policy",
					Description: `This argument specify subnets (zones), that will be used by node group compute instances. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "instance_group_id",
					Description: `ID of instance group that is used to manage this Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `Information about maintenance policy for this Kubernetes node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "node_labels",
					Description: `A set of key/value label pairs, that are assigned to all the nodes of this Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "node_taints",
					Description: `A list of Kubernetes taints, that are applied to all the nodes of this Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "allowed_unsafe_sysctls",
					Description: `A list of allowed unsafe sysctl parameters for this node group. For more details see [documentation](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/).`,
				},
				resource.Attribute{
					Name:        "version_info",
					Description: `Information about Kubernetes node group version. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "deploy_policy",
					Description: `Deploy policy of the node group. The structure is documented below. --- The ` + "`" + `instance_template` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `The ID of the hardware platform configuration for the instance.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Boolean flag, when true, NAT for node group instances is enabled.`,
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
					Name:        "boot_disk",
					Description: `The specifications for boot disks that will be attached to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduling_policy",
					Description: `The scheduling policy for the instances in node group. The structure is documented below. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the disk in GB. Allowed minimal size: 64 GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The disk type. --- The ` + "`" + `scheduling_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `Specifies if the instance is preemptible. Defaults to false. --- The ` + "`" + `scale_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fixed_scale",
					Description: `Scale policy for a fixed scale node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "auto_scale",
					Description: `Scale policy for an autoscaled node group. The structure is documented below. --- The ` + "`" + `fixed_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the node group. --- The ` + "`" + `auto_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Minimum number of instances in the node group.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Maximum number of instances in the node group.`,
				},
				resource.Attribute{
					Name:        "initial",
					Description: `Initial number of instances in the node group. --- The ` + "`" + `allocation_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Repeated field, that specify subnets (zones), that will be used by node group compute instances. The structure is documented below. --- The ` + "`" + `location` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of the availability zone where for one compute instance in node group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet, that will be used by one compute instance in node group. Subnet specified by ` + "`" + `subnet_id` + "`" + ` should be allocated in zone specified by 'zone' argument --- The ` + "`" + `maintenance_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_upgrade",
					Description: `Boolean flag.`,
				},
				resource.Attribute{
					Name:        "auto_repair",
					Description: `Boolean flag.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `Set of day intervals, when maintenance is allowed for this node group. When omitted, it defaults to any time. Weekly maintenance policy expands to one element, with only two fields set: ` + "`" + `start_time` + "`" + `, ` + "`" + `duration` + "`" + `, and ` + "`" + `day` + "`" + ` field omitted. Daily maintenance policy expands to list of elements, with all fields set, that specify time interval for selected days. Only one interval is possible for any week day. Some days can be omitted, when there is no allowed interval for maintenance specified. --- The ` + "`" + `version_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "current_version",
					Description: `Current Kubernetes version, major.minor (e.g. 1.15).`,
				},
				resource.Attribute{
					Name:        "new_revision_available",
					Description: `True/false flag. Newer revisions may include Kubernetes patches (e.g 1.15.1 -> 1.15.2) as well as some internal component updates - new features or bug fixes in yandex-specific components either on the master or nodes.`,
				},
				resource.Attribute{
					Name:        "new_revision_summary",
					Description: `Human readable description of the changes to be applied when updating to the latest revision. Empty if new_revision_available is false.`,
				},
				resource.Attribute{
					Name:        "version_deprecated",
					Description: `True/false flag. The current version is on the deprecation schedule, component (master or node group) should be upgraded. --- The ` + "`" + `deploy_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "max_expansion",
					Description: `The maximum number of instances that can be temporarily allocated above the group's target size during the update.`,
				},
				resource.Attribute{
					Name:        "max_unavailable",
					Description: `The maximum number of running instances that can be taken offline during update. ---`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the Kubernetes cluster that this node group belongs to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs assigned to the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The Kubernetes node group creation timestamp.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `Template used to create compute instances in this Kubernetes node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scale_policy",
					Description: `Scale policy of the node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "allocation_policy",
					Description: `This argument specify subnets (zones), that will be used by node group compute instances. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "instance_group_id",
					Description: `ID of instance group that is used to manage this Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `Information about maintenance policy for this Kubernetes node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "node_labels",
					Description: `A set of key/value label pairs, that are assigned to all the nodes of this Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "node_taints",
					Description: `A list of Kubernetes taints, that are applied to all the nodes of this Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "allowed_unsafe_sysctls",
					Description: `A list of allowed unsafe sysctl parameters for this node group. For more details see [documentation](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/).`,
				},
				resource.Attribute{
					Name:        "version_info",
					Description: `Information about Kubernetes node group version. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "deploy_policy",
					Description: `Deploy policy of the node group. The structure is documented below. --- The ` + "`" + `instance_template` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `The ID of the hardware platform configuration for the instance.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Boolean flag, when true, NAT for node group instances is enabled.`,
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
					Name:        "boot_disk",
					Description: `The specifications for boot disks that will be attached to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduling_policy",
					Description: `The scheduling policy for the instances in node group. The structure is documented below. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the disk in GB. Allowed minimal size: 64 GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The disk type. --- The ` + "`" + `scheduling_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `Specifies if the instance is preemptible. Defaults to false. --- The ` + "`" + `scale_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fixed_scale",
					Description: `Scale policy for a fixed scale node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "auto_scale",
					Description: `Scale policy for an autoscaled node group. The structure is documented below. --- The ` + "`" + `fixed_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the node group. --- The ` + "`" + `auto_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Minimum number of instances in the node group.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Maximum number of instances in the node group.`,
				},
				resource.Attribute{
					Name:        "initial",
					Description: `Initial number of instances in the node group. --- The ` + "`" + `allocation_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Repeated field, that specify subnets (zones), that will be used by node group compute instances. The structure is documented below. --- The ` + "`" + `location` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of the availability zone where for one compute instance in node group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet, that will be used by one compute instance in node group. Subnet specified by ` + "`" + `subnet_id` + "`" + ` should be allocated in zone specified by 'zone' argument --- The ` + "`" + `maintenance_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_upgrade",
					Description: `Boolean flag.`,
				},
				resource.Attribute{
					Name:        "auto_repair",
					Description: `Boolean flag.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `Set of day intervals, when maintenance is allowed for this node group. When omitted, it defaults to any time. Weekly maintenance policy expands to one element, with only two fields set: ` + "`" + `start_time` + "`" + `, ` + "`" + `duration` + "`" + `, and ` + "`" + `day` + "`" + ` field omitted. Daily maintenance policy expands to list of elements, with all fields set, that specify time interval for selected days. Only one interval is possible for any week day. Some days can be omitted, when there is no allowed interval for maintenance specified. --- The ` + "`" + `version_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "current_version",
					Description: `Current Kubernetes version, major.minor (e.g. 1.15).`,
				},
				resource.Attribute{
					Name:        "new_revision_available",
					Description: `True/false flag. Newer revisions may include Kubernetes patches (e.g 1.15.1 -> 1.15.2) as well as some internal component updates - new features or bug fixes in yandex-specific components either on the master or nodes.`,
				},
				resource.Attribute{
					Name:        "new_revision_summary",
					Description: `Human readable description of the changes to be applied when updating to the latest revision. Empty if new_revision_available is false.`,
				},
				resource.Attribute{
					Name:        "version_deprecated",
					Description: `True/false flag. The current version is on the deprecation schedule, component (master or node group) should be upgraded. --- The ` + "`" + `deploy_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "max_expansion",
					Description: `The maximum number of instances that can be temporarily allocated above the group's target size during the update.`,
				},
				resource.Attribute{
					Name:        "max_unavailable",
					Description: `The maximum number of running instances that can be taken offline during update. ---`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_lb_network_load_balancer",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Load Balancer network load balancer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) - Name of the network load balancer. ~>`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the network load balancer.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this network load balancer.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region where the network load balancer resides.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the network load balancer.`,
				},
				resource.Attribute{
					Name:        "attached_target_group",
					Description: `An attached target group is a group of targets that is attached to a load balancer. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "listener",
					Description: `Listener specification that will be used by a network load balancer. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this network load balancer. --- The ` + "`" + `attached_target_group` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `ID of the target group that attached to the network load balancer.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.name",
					Description: `Name of the health check.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.interval",
					Description: `The interval between health checks.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.timeout",
					Description: `Timeout for a target to return a response for the health check.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.unhealthy_threshold",
					Description: `Number of failed health checks before changing the status to ` + "`" + `UNHEALTHY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.healthy_threshold",
					Description: `Number of successful health checks required in order to set the ` + "`" + `HEALTHY` + "`" + ` status for the target.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.tcp_options.0.port",
					Description: `Port to use for TCP health checks.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.http_options.0.port",
					Description: `Port to use for HTTP health checks.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.http_options.0.path",
					Description: `URL path to use for HTTP health checks. The ` + "`" + `listener` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port for incoming traffic.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol for incoming traffic.`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `Port of a target.`,
				},
				resource.Attribute{
					Name:        "external_address_spec.0.address",
					Description: `External IP address of a listener.`,
				},
				resource.Attribute{
					Name:        "external_address_spec.0.ip_version",
					Description: `IP version of the external addresses.`,
				},
				resource.Attribute{
					Name:        "internal_address_spec.0.subnet_id",
					Description: `Subnet ID to which the internal IP address belongs`,
				},
				resource.Attribute{
					Name:        "internal_address_spec.0.address",
					Description: `Internal IP address of a listener.`,
				},
				resource.Attribute{
					Name:        "internal_address_spec.0.ip_version",
					Description: `IP version of the internal addresses. [Load Balancer Network Load Balancers]: https://cloud.yandex.com/docs/load-balancer/concepts/`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the network load balancer.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this network load balancer.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region where the network load balancer resides.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the network load balancer.`,
				},
				resource.Attribute{
					Name:        "attached_target_group",
					Description: `An attached target group is a group of targets that is attached to a load balancer. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "listener",
					Description: `Listener specification that will be used by a network load balancer. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this network load balancer. --- The ` + "`" + `attached_target_group` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `ID of the target group that attached to the network load balancer.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.name",
					Description: `Name of the health check.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.interval",
					Description: `The interval between health checks.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.timeout",
					Description: `Timeout for a target to return a response for the health check.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.unhealthy_threshold",
					Description: `Number of failed health checks before changing the status to ` + "`" + `UNHEALTHY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.healthy_threshold",
					Description: `Number of successful health checks required in order to set the ` + "`" + `HEALTHY` + "`" + ` status for the target.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.tcp_options.0.port",
					Description: `Port to use for TCP health checks.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.http_options.0.port",
					Description: `Port to use for HTTP health checks.`,
				},
				resource.Attribute{
					Name:        "healthcheck.0.http_options.0.path",
					Description: `URL path to use for HTTP health checks. The ` + "`" + `listener` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port for incoming traffic.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol for incoming traffic.`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `Port of a target.`,
				},
				resource.Attribute{
					Name:        "external_address_spec.0.address",
					Description: `External IP address of a listener.`,
				},
				resource.Attribute{
					Name:        "external_address_spec.0.ip_version",
					Description: `IP version of the external addresses.`,
				},
				resource.Attribute{
					Name:        "internal_address_spec.0.subnet_id",
					Description: `Subnet ID to which the internal IP address belongs`,
				},
				resource.Attribute{
					Name:        "internal_address_spec.0.address",
					Description: `Internal IP address of a listener.`,
				},
				resource.Attribute{
					Name:        "internal_address_spec.0.ip_version",
					Description: `IP version of the internal addresses. [Load Balancer Network Load Balancers]: https://cloud.yandex.com/docs/load-balancer/concepts/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_lb_target_group",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Load Balancer target group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) - Name of the Target Group. ~>`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the target group.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this target group.`,
				},
				resource.Attribute{
					Name:        "target.0.address",
					Description: `IP address of the target.`,
				},
				resource.Attribute{
					Name:        "target.0.subnet_id",
					Description: `ID of the subnet that targets are connected to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this target group. [Load Balancer Target Groups]: https://cloud.yandex.com/docs/load-balancer/concepts/target-resources`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the target group.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this target group.`,
				},
				resource.Attribute{
					Name:        "target.0.address",
					Description: `IP address of the target.`,
				},
				resource.Attribute{
					Name:        "target.0.subnet_id",
					Description: `ID of the subnet that targets are connected to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this target group. [Load Balancer Target Groups]: https://cloud.yandex.com/docs/load-balancer/concepts/target-resources`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_clickhouse_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Managed ClickHouse cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) The ID of the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the ClickHouse cluster. ~>`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network, to which the ClickHouse cluster belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Deployment environment of the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "clickhouse",
					Description: `Configuration of the ClickHouse subcluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `A user of the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `A database of the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A host of the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `Time to start the daily backup, in the UTC timezone. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Access policy to the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "zookeeper",
					Description: `Configuration of the ZooKeeper subcluster. The structure is documented below. The ` + "`" + `clickhouse` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the ClickHouse subcluster. The structure is documented below. The ` + "`" + `zookeeper` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the ZooKeeper subcluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `The ID of the preset for computational resources available to a ClickHouse or ZooKeeper host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a ClickHouse or ZooKeeper host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `Type of the storage of ClickHouse or ZooKeeper hosts. The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The name of the database that the permission grants access to. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the host.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the host to be deployed.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The availability zone where the ClickHouse host will be created.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "shard_name",
					Description: `The name of the shard to which the host belongs.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `Sets whether the host should get a public IP address on creation. The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The minute at which backup will be started. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "web_sql",
					Description: `Allow access for DataLens.`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `Allow access for Web SQL.`,
				},
				resource.Attribute{
					Name:        "metrika",
					Description: `Allow access for Yandex.Metrika.`,
				},
				resource.Attribute{
					Name:        "serverless",
					Description: `Allow access for Serverless.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network, to which the ClickHouse cluster belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Deployment environment of the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "clickhouse",
					Description: `Configuration of the ClickHouse subcluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `A user of the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `A database of the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A host of the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `Time to start the daily backup, in the UTC timezone. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Access policy to the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "zookeeper",
					Description: `Configuration of the ZooKeeper subcluster. The structure is documented below. The ` + "`" + `clickhouse` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the ClickHouse subcluster. The structure is documented below. The ` + "`" + `zookeeper` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the ZooKeeper subcluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `The ID of the preset for computational resources available to a ClickHouse or ZooKeeper host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a ClickHouse or ZooKeeper host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `Type of the storage of ClickHouse or ZooKeeper hosts. The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The name of the database that the permission grants access to. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the host.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the host to be deployed.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The availability zone where the ClickHouse host will be created.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "shard_name",
					Description: `The name of the shard to which the host belongs.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `Sets whether the host should get a public IP address on creation. The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The minute at which backup will be started. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "web_sql",
					Description: `Allow access for DataLens.`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `Allow access for Web SQL.`,
				},
				resource.Attribute{
					Name:        "metrika",
					Description: `Allow access for Yandex.Metrika.`,
				},
				resource.Attribute{
					Name:        "serverless",
					Description: `Allow access for Serverless.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_mongodb_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Managed MongoDB cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) The ID of the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the MongoDB cluster. ~>`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network, to which the MongoDB cluster belongs.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Deployment environment of the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "sharded",
					Description: `MongoDB Cluster mode enabled/disabled.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A host of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `Configuration of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `A user of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `A database of the MongoDB cluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `The ID of the preset for computational resources available to a host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `The ID of the storage type. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts/storage) The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully qualified domain name of the host.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The availability zone where the MongoDB host will be created.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `The role of the cluster (either PRIMARY or SECONDARY).`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `The health of the host.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `Has assigned public IP.`,
				},
				resource.Attribute{
					Name:        "shard_name",
					Description: `The name of the shard to which the host belongs.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type of mongo demon which runs on this host (mongod, mongos or monogcfg). The ` + "`" + `cluster_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of MongoDB (either 4.2, 4.0 or 3.6).`,
				},
				resource.Attribute{
					Name:        "feature_compatibility_version",
					Description: `Feature compatibility version of MongoDB.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `Time to start the daily backup, in the UTC timezone. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Access policy to MongoDB cluster. The structure is documented below. The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The minute at which backup will be started. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `Shows whether cluster has access to data lens. The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The name of the database that the permission grants access to.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional) List of strings. The roles of the user in this database. For more information see [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts/users-and-roles). The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network, to which the MongoDB cluster belongs.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Deployment environment of the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "sharded",
					Description: `MongoDB Cluster mode enabled/disabled.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A host of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `Configuration of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `A user of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `A database of the MongoDB cluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `The ID of the preset for computational resources available to a host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `The ID of the storage type. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts/storage) The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully qualified domain name of the host.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The availability zone where the MongoDB host will be created.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `The role of the cluster (either PRIMARY or SECONDARY).`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `The health of the host.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `Has assigned public IP.`,
				},
				resource.Attribute{
					Name:        "shard_name",
					Description: `The name of the shard to which the host belongs.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type of mongo demon which runs on this host (mongod, mongos or monogcfg). The ` + "`" + `cluster_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of MongoDB (either 4.2, 4.0 or 3.6).`,
				},
				resource.Attribute{
					Name:        "feature_compatibility_version",
					Description: `Feature compatibility version of MongoDB.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `Time to start the daily backup, in the UTC timezone. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Access policy to MongoDB cluster. The structure is documented below. The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The minute at which backup will be started. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `Shows whether cluster has access to data lens. The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The name of the database that the permission grants access to.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional) List of strings. The roles of the user in this database. For more information see [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts/users-and-roles). The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_mysql_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Managed MySQL cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) The ID of the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the MySQL cluster. ~>`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network, to which the MySQL cluster belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Deployment environment of the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the MySQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `A user of the MySQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `A database of the MySQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A host of the MySQL cluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `The ID of the preset for computational resources available to a MySQL host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mysql/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a MySQL host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `Type of the storage for MySQL hosts. The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The minute at which backup will be started. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `Allow access for Web SQL. The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The name of the database that the permission grants access to.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List user's roles in the database. Allowed roles: ` + "`" + `ALL` + "`" + `,` + "`" + `ALTER` + "`" + `,` + "`" + `ALTER_ROUTINE` + "`" + `,` + "`" + `CREATE` + "`" + `,` + "`" + `CREATE_ROUTINE` + "`" + `,` + "`" + `CREATE_TEMPORARY_TABLES` + "`" + `, ` + "`" + `CREATE_VIEW` + "`" + `,` + "`" + `DELETE` + "`" + `,` + "`" + `DROP` + "`" + `,` + "`" + `EVENT` + "`" + `,` + "`" + `EXECUTE` + "`" + `,` + "`" + `INDEX` + "`" + `,` + "`" + `INSERT` + "`" + `,` + "`" + `LOCK_TABLES` + "`" + `,` + "`" + `SELECT` + "`" + `,` + "`" + `SHOW_VIEW` + "`" + `,` + "`" + `TRIGGER` + "`" + `,` + "`" + `UPDATE` + "`" + `. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the host.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The availability zone where the MySQL host will be created.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `Sets whether the host should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network, to which the MySQL cluster belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Deployment environment of the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the MySQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `A user of the MySQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `A database of the MySQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A host of the MySQL cluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `The ID of the preset for computational resources available to a MySQL host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mysql/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a MySQL host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `Type of the storage for MySQL hosts. The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The minute at which backup will be started. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `Allow access for Web SQL. The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The name of the database that the permission grants access to.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List user's roles in the database. Allowed roles: ` + "`" + `ALL` + "`" + `,` + "`" + `ALTER` + "`" + `,` + "`" + `ALTER_ROUTINE` + "`" + `,` + "`" + `CREATE` + "`" + `,` + "`" + `CREATE_ROUTINE` + "`" + `,` + "`" + `CREATE_TEMPORARY_TABLES` + "`" + `, ` + "`" + `CREATE_VIEW` + "`" + `,` + "`" + `DELETE` + "`" + `,` + "`" + `DROP` + "`" + `,` + "`" + `EVENT` + "`" + `,` + "`" + `EXECUTE` + "`" + `,` + "`" + `INDEX` + "`" + `,` + "`" + `INSERT` + "`" + `,` + "`" + `LOCK_TABLES` + "`" + `,` + "`" + `SELECT` + "`" + `,` + "`" + `SHOW_VIEW` + "`" + `,` + "`" + `TRIGGER` + "`" + `,` + "`" + `UPDATE` + "`" + `. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the host.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The availability zone where the MySQL host will be created.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `Sets whether the host should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_postgresql_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Managed PostgreSQL cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) The ID of the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the PostgreSQL cluster. ~>`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network, to which the PostgreSQL cluster belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp of cluster creation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Deployment environment of the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Configuration of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `A user of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `A database of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A host of the PostgreSQL cluster. The structure is documented below. The ` + "`" + `config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "autofailover",
					Description: `Configuration setting which enables/disables autofailover in cluster.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "pooler_config",
					Description: `Configuration of the connection pooler. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `Time to start the daily backup, in the UTC timezone. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Access policy to the PostgreSQL cluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `The ID of the preset for computational resources available to a PostgreSQL host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-postgresql/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a PostgreSQL host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `Type of the storage for PostgreSQL hosts. The ` + "`" + `pooler_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "pooling_mode",
					Description: `Mode that the connection pooler is working in. See descriptions of all modes in the [documentation for PgBouncer](https://pgbouncer.github.io/usage).`,
				},
				resource.Attribute{
					Name:        "pool_discard",
					Description: `Setting ` + "`" + `server_reset_query_always` + "`" + ` parameter in PgBouncer. The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The minute at which backup will be started. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `Allow access for [Yandex DataLens](https://cloud.yandex.com/services/datalens). The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Set of permissions granted to the user. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `User's ability to login.`,
				},
				resource.Attribute{
					Name:        "grants",
					Description: `List of the user's grants. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The name of the database that the permission grants access to. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Name of the user assigned as the owner of the database.`,
				},
				resource.Attribute{
					Name:        "lc_collate",
					Description: `POSIX locale for string sorting order. Forbidden to change in an existing database.`,
				},
				resource.Attribute{
					Name:        "lc_type",
					Description: `POSIX locale for character classification. Forbidden to change in an existing database.`,
				},
				resource.Attribute{
					Name:        "extension",
					Description: `Set of database extensions. The structure is documented below The ` + "`" + `extension` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the database extension. For more information on available extensions see [the official documentation](https://cloud.yandex.com/docs/managed-postgresql/operations/cluster-extensions).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the extension. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the host.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The availability zone where the PostgreSQL host will be created.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `Sets whether the host should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network, to which the PostgreSQL cluster belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp of cluster creation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Deployment environment of the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Configuration of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `A user of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `A database of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A host of the PostgreSQL cluster. The structure is documented below. The ` + "`" + `config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "autofailover",
					Description: `Configuration setting which enables/disables autofailover in cluster.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "pooler_config",
					Description: `Configuration of the connection pooler. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `Time to start the daily backup, in the UTC timezone. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Access policy to the PostgreSQL cluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `The ID of the preset for computational resources available to a PostgreSQL host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-postgresql/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a PostgreSQL host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `Type of the storage for PostgreSQL hosts. The ` + "`" + `pooler_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "pooling_mode",
					Description: `Mode that the connection pooler is working in. See descriptions of all modes in the [documentation for PgBouncer](https://pgbouncer.github.io/usage).`,
				},
				resource.Attribute{
					Name:        "pool_discard",
					Description: `Setting ` + "`" + `server_reset_query_always` + "`" + ` parameter in PgBouncer. The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The minute at which backup will be started. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `Allow access for [Yandex DataLens](https://cloud.yandex.com/services/datalens). The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Set of permissions granted to the user. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `User's ability to login.`,
				},
				resource.Attribute{
					Name:        "grants",
					Description: `List of the user's grants. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The name of the database that the permission grants access to. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Name of the user assigned as the owner of the database.`,
				},
				resource.Attribute{
					Name:        "lc_collate",
					Description: `POSIX locale for string sorting order. Forbidden to change in an existing database.`,
				},
				resource.Attribute{
					Name:        "lc_type",
					Description: `POSIX locale for character classification. Forbidden to change in an existing database.`,
				},
				resource.Attribute{
					Name:        "extension",
					Description: `Set of database extensions. The structure is documented below The ` + "`" + `extension` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the database extension. For more information on available extensions see [the official documentation](https://cloud.yandex.com/docs/managed-postgresql/operations/cluster-extensions).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the extension. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the host.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The availability zone where the PostgreSQL host will be created.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `Sets whether the host should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_redis_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex Managed Redis cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) The ID of the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Redis cluster. ~>`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network, to which the Redis cluster belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Deployment environment of the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Configuration of the Redis cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the Redis cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A host of the Redis cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "sharded",
					Description: `Redis Cluster mode enabled/disabled. The ` + "`" + `config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Close the connection after a client is idle for N seconds.`,
				},
				resource.Attribute{
					Name:        "maxmemory_policy",
					Description: `Redis key eviction policy for a dataset that reaches maximum memory. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `The ID of the preset for computational resources available to a host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-redis/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a host, in gigabytes. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The availability zone where the Redis host will be created.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "shard_name",
					Description: `The name of the shard to which the host belongs.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the host.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network, to which the Redis cluster belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Deployment environment of the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Configuration of the Redis cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Resources allocated to hosts of the Redis cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A host of the Redis cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "sharded",
					Description: `Redis Cluster mode enabled/disabled. The ` + "`" + `config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Close the connection after a client is idle for N seconds.`,
				},
				resource.Attribute{
					Name:        "maxmemory_policy",
					Description: `Redis key eviction policy for a dataset that reaches maximum memory. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `The ID of the preset for computational resources available to a host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-redis/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of the storage available to a host, in gigabytes. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The availability zone where the Redis host will be created.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "shard_name",
					Description: `The name of the shard to which the host belongs.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the host.`,
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
					Name:        "cloud_id",
					Description: `(Optional) Cloud that the resource belongs to. If value is omitted, the default provider cloud is used. ## Attributes Reference The following attributes are exported:`,
				},
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
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the network.`,
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
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the route table.`,
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
			Type:             "yandex_vpc_security_group",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Yandex VPC Security Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ## Attributes Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the security group.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network this security group belongs to.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder this security group belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this security group.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `A list of ingress rules. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `A list of egress rules. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this security group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this security group. --- The ` + "`" + `ingress` + "`" + ` and ` + "`" + `egress` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the rule.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `One of ` + "`" + `ANY` + "`" + `, ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + `, ` + "`" + `IPV6_ICMP` + "`" + ` or protocol number.`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `Minimum port number.`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `Maximum port number.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number (if applied to a single port).`,
				},
				resource.Attribute{
					Name:        "v4_cidr_blocks",
					Description: `The blocks of IPv4 addresses for this rule.`,
				},
				resource.Attribute{
					Name:        "v6_cidr_blocks",
					Description: `The blocks of IPv6 addresses for this rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the security group.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of the network this security group belongs to.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `ID of the folder this security group belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this security group.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `A list of ingress rules. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `A list of egress rules. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this security group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this security group. --- The ` + "`" + `ingress` + "`" + ` and ` + "`" + `egress` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the rule.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels to assign to this rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `One of ` + "`" + `ANY` + "`" + `, ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + `, ` + "`" + `IPV6_ICMP` + "`" + ` or protocol number.`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `Minimum port number.`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `Maximum port number.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number (if applied to a single port).`,
				},
				resource.Attribute{
					Name:        "v4_cidr_blocks",
					Description: `The blocks of IPv4 addresses for this rule.`,
				},
				resource.Attribute{
					Name:        "v6_cidr_blocks",
					Description: `The blocks of IPv6 addresses for this rule.`,
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
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the subnet.`,
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

		"yandex_compute_disk":             0,
		"yandex_compute_image":            1,
		"yandex_compute_instance":         2,
		"yandex_compute_instance_group":   3,
		"yandex_compute_snapshot":         4,
		"yandex_container_registry":       5,
		"yandex_dataproc_cluster":         6,
		"yandex_function":                 7,
		"yandex_function_trigger":         8,
		"yandex_iam_policy":               9,
		"yandex_iam_role":                 10,
		"yandex_iam_service_account":      11,
		"yandex_iam_user":                 12,
		"yandex_iot_core":                 13,
		"yandex_iot_registry":             14,
		"yandex_kubernetes_cluster":       15,
		"yandex_kubernetes_node_group":    16,
		"yandex_lb_network_load_balancer": 17,
		"yandex_lb_target_group":          18,
		"yandex_mdb_clickhouse_cluster":   19,
		"yandex_mdb_mongodb_cluster":      20,
		"yandex_mdb_mysql_cluster":        21,
		"yandex_mdb_postgresql_cluster":   22,
		"yandex_mdb_redis_cluster":        23,
		"yandex_resourcemanager_cloud":    24,
		"yandex_resourcemanager_folder":   25,
		"yandex_vpc_network":              26,
		"yandex_vpc_route_table":          27,
		"yandex_vpc_security_group":       28,
		"yandex_vpc_subnet":               29,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
