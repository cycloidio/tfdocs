package yandex

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_disk",
			Category:         "Yandex Compute Service Resources",
			ShortDescription: `Persistent disks are durable storage devices that function similarly to the physical disks in a desktop or a server.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"service",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the disk. Provide this property when you create a resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the disk. Provide this property when you create a resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the disk belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to assign to this disk. A list of key/value pairs.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Availability zone where the disk will reside.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the ` + "`" + `image_id` + "`" + ` or ` + "`" + `snapshot_id` + "`" + ` parameter, or specify it alone to create an empty persistent disk. If you specify this field along with ` + "`" + `image_id` + "`" + ` or ` + "`" + `snapshot_id` + "`" + `, the size value must not be less than the size of the source image or the size of the snapshot.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of disk to create. Provide this when creating a disk. One of ` + "`" + `network-hdd` + "`" + ` (default) or ` + "`" + `network-nvme` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The source image to use for disk creation.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The source snapshot to use for disk creation. ~>`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the disk.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the disk. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 5 minutes. ## Import A disk can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_disk.default disk_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of the disk.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the disk. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 5 minutes. ## Import A disk can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_disk.default disk_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_image",
			Category:         "Yandex Compute Service Resources",
			ShortDescription: `Creates a VM image for the Yandex Compute service from an existing tarball.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"service",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the disk.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the image. Provide this property when you create a resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the image.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `(Optional) The name of the image family to which this image belongs.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `(Optional) Minimum size in GB of the disk that will be created from this image.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional) Operating system type that is contained in the image. Possible values: "LINUX", "WINDOWS".`,
				},
				resource.Attribute{
					Name:        "source_family",
					Description: `(Optional) The name of the family to use as the source of the new image. The ID of the latest image is taken from the "standard-images" folder. Changing the family forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "source_image",
					Description: `(Optional) The ID of an existing image to use as the source of the image. Changing this ID forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "source_snapshot",
					Description: `(Optional) The ID of a snapshot to use as the source of the image. Changing this ID forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "source_disk",
					Description: `(Optional) The ID of a disk to use as the source of the image. Changing this ID forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "source_url",
					Description: `(Optional) The URL to use as the source of the image. Changing this URL forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "product_ids",
					Description: `(Optional) License IDs that indicate which licenses are attached to this image. ~>`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image, specified in GB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the image. ## Timeouts ` + "`" + `yandex_compute_image` + "`" + ` provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default 5 minutes - ` + "`" + `update` + "`" + ` - Default 5 minutes - ` + "`" + `delete` + "`" + ` - Default 5 minutes ## Import A VM image can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_image.web-image image_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image, specified in GB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the image. ## Timeouts ` + "`" + `yandex_compute_image` + "`" + ` provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default 5 minutes - ` + "`" + `update` + "`" + ` - Default 5 minutes - ` + "`" + `delete` + "`" + ` - Default 5 minutes ## Import A VM image can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_image.web-image image_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_instance",
			Category:         "Yandex Compute Service Resources",
			ShortDescription: `Manages a VM instance resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"service",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Compute resources that are allocated for the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `(Required) The boot disk for the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Required) Networks to attach to the instance. This can be specified multiple times. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Resource name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the instance.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The availability zone where the virtual machine will be created. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Host name for the instance. This field is used to generate the instance ` + "`" + `fqdn` + "`" + ` value. The host name must be unique within the network and region. If not specified, the host name will be equal to ` + "`" + `id` + "`" + ` of the instance and ` + "`" + `fqdn` + "`" + ` will be ` + "`" + `<id>.auto.internal` + "`" + `. Otherwise FQDN will be ` + "`" + `<hostname>.<region_id>.internal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the instance.`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `(Optional) The type of virtual machine to create. The default is 'standard-v1'.`,
				},
				resource.Attribute{
					Name:        "secondary_disk",
					Description: `(Optional) A list of disks to attach to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduling_policy",
					Description: `(Optional) Scheduling policy configuration. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Optional) ID of the service account authorized for this instance.`,
				},
				resource.Attribute{
					Name:        "allow_stopping_for_update",
					Description: `(Optional) If true, allows Terraform to stop the instance in order to update its properties. If you try to update a property that requires stopping the instance without setting this field, the update will fail. --- The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `(Required) CPU cores for the instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size in GB.`,
				},
				resource.Attribute{
					Name:        "core_fraction",
					Description: `(Optional) If provided, specifies baseline performance for a core as a percent. The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `(Optional) Defines whether the disk will be auto-deleted when the instance is deleted. The default value is ` + "`" + `True` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) Name that can be used to access an attached disk.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Type of access to the disk resource. By default, a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `(Optional) The ID of the existing disk (such as those managed by ` + "`" + `yandex_compute_disk` + "`" + `) to attach as a boot disk.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `(Optional) Parameters for a new disk that will be created alongside the new instance. Either ` + "`" + `initialize_params` + "`" + ` or ` + "`" + `disk_id` + "`" + ` must be set. The structure is documented below. ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the boot disk.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) A disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) A snapshot to initialize this disk from. ~>`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) ID of the subnet to attach this interface to. The subnet must exist in the same zone where this instance will be created.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The private IP address to assign to the instance. If empty, the address will be automatically assigned from the specified subnet.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) If true, allocate an IPv6 address for the interface. The address will be automatically assigned from the specified subnet.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) The private IPv6 address to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `(Optional) Provide a public address, for instance, to access the internet over NAT. The ` + "`" + `secondary_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `(Required) ID of the disk that is attached to the instance.`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `(Optional) Whether the disk is auto-deleted when the instance is deleted. The default value is false.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) Name that can be used to access an attached disk under ` + "`" + `/dev/disk/by-id/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Type of access to the disk resource. By default, a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode. The ` + "`" + `scheduling_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) Specifies if the instance is preemptible. Defaults to false. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified DNS name of this instance.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.ip_address",
					Description: `The internal IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.nat_ip_address",
					Description: `The external IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this instance.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the instance. ## Import Instances can be imported using the ` + "`" + `ID` + "`" + ` of an instance, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_instance.default instance_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified DNS name of this instance.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.ip_address",
					Description: `The internal IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.nat_ip_address",
					Description: `The external IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this instance.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the instance. ## Import Instances can be imported using the ` + "`" + `ID` + "`" + ` of an instance, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_instance.default instance_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_instance_group",
			Category:         "Yandex Compute Service Resources",
			ShortDescription: `Manages an Instance group resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"service",
				"instance",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Required) The ID of the folder that the resources belong to.`,
				},
				resource.Attribute{
					Name:        "scale_policy",
					Description: `(Required) The scaling policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "deploy_policy",
					Description: `(Required) The deployment policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) The ID of the service account authorized for this instance group.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `(Required) The template for creating new instances. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "allocation_policy",
					Description: `(Required) The allocation policy of the instance group by zone and region. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the instance group.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Optional) Health check specifications. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Optional) Load balancing specifications. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the instance group.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the instance group. --- The ` + "`" + `load_balancer` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_group_name",
					Description: `(Optional) The name of the target group.`,
				},
				resource.Attribute{
					Name:        "target_group_description",
					Description: `(Optional) A description of the target group.`,
				},
				resource.Attribute{
					Name:        "target_group_labels",
					Description: `(Optional) A set of key/value label pairs. --- The ` + "`" + `health_check` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The interval to wait between health checks in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The length of time to wait for a response before the health check times out in seconds.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) The number of successful health checks before the managed instance is declared healthy.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) The number of failed health checks before the managed instance is declared unhealthy.`,
				},
				resource.Attribute{
					Name:        "tcp_options",
					Description: `(Optional) TCP check options. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "http_options",
					Description: `(Optional) HTTP check options. The structure is documented below. --- The ` + "`" + `http_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port used for HTTP health checks.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The URL path used for health check requests. --- The ` + "`" + `tcp_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port used for TCP health checks. --- The ` + "`" + `allocation_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `(Required) A list of availability zones. --- The ` + "`" + `instance_template` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `(Required) Boot disk specifications for the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Compute resource specifications for the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Required) Network specifications for the instance. This can be used multiple times for adding multiple interfaces. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "scheduling_policy",
					Description: `(Optional) The scheduling policy configuration. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the instance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A set of metadata key/value pairs to make available from within the instance.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `(Optional) The ID of the hardware platform configuration for the instance. The default is 'standard-v1'.`,
				},
				resource.Attribute{
					Name:        "secondary_disk",
					Description: `(Optional) A list of disks to attach to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Optional) The ID of the service account authorized for this instance. --- The ` + "`" + `secondary_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The access mode to the disk resource. By default a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `(Required) Parameters used for creating a disk alongside the instance. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) This value can be used to reference the device under ` + "`" + `/dev/disk/by-id/` + "`" + `. --- The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The snapshot to initialize this disk from. ~>`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) Specifies if the instance is preemptible. Defaults to false. --- The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The ID of the subnets to attach this interface to.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `(Optional) A public address that can be used to access the internet over NAT. --- The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) The memory size in GB.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `(Required) The number of CPU cores for the instance. - - -`,
				},
				resource.Attribute{
					Name:        "core_fraction",
					Description: `(Optional) If provided, specifies baseline core performance as a percent. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The access mode to the disk resource. By default a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `(Required) Parameters for creating a disk alongside the instance. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) This value can be used to reference the device under ` + "`" + `/dev/disk/by-id/` + "`" + `. --- The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The snapshot to initialize this disk from. ~>`,
				},
				resource.Attribute{
					Name:        "max_unavailable",
					Description: `(Required) The maximum number of running instances that can be taken offline (stopped or deleted) at the same time during the update process.`,
				},
				resource.Attribute{
					Name:        "max_expansion",
					Description: `(Required) The maximum number of instances that can be temporarily allocated above the group's target size during the update process. - - -`,
				},
				resource.Attribute{
					Name:        "max_deleting",
					Description: `(Optional) The maximum number of instances that can be deleted at the same time.`,
				},
				resource.Attribute{
					Name:        "max_creating",
					Description: `(Optional) The maximum number of instances that can be created at the same time.`,
				},
				resource.Attribute{
					Name:        "startup_duration",
					Description: `(Optional) The amount of time in seconds to allow for an instance to start. Instance will be considered up and running (and start receiving traffic) only after the startup_duration has elapsed and all health checks are passed. --- The ` + "`" + `scale_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fixed_scale",
					Description: `(Optional) The fixed scaling policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "auto_scale",
					Description: `(Optional) The auto scaling policy of the instance group. The structure is documented below. ~>`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The number of instances in the instance group. --- The ` + "`" + `auto_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "initial_size",
					Description: `(Required) The initial number of instances in the instance group.`,
				},
				resource.Attribute{
					Name:        "measurement_duration",
					Description: `(Required) The amount of time, in seconds, that metrics are averaged for. If the average value at the end of the interval is higher than the ` + "`" + `cpu_utilization_target` + "`" + `, the instance group will increase the number of virtual machines in the group.`,
				},
				resource.Attribute{
					Name:        "cpu_utilization_target",
					Description: `(Required) Target CPU load level.`,
				},
				resource.Attribute{
					Name:        "min_zone_size",
					Description: `(Optional) The minimum number of virtual machines in a single availability zone.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional) The maximum number of virtual machines in the group.`,
				},
				resource.Attribute{
					Name:        "warmup_duration",
					Description: `(Optional) The warm-up time of the virtual machine, in seconds. During this time, traffic is fed to the virtual machine, but load metrics are not taken into account.`,
				},
				resource.Attribute{
					Name:        "stabilization_duration",
					Description: `(Optional) The minimum time interval, in seconds, to monitor the load before an instance group can reduce the number of virtual machines in the group. During this time, the group will not decrease even if the average load falls below the value of ` + "`" + `cpu_utilization_target` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_rule",
					Description: `(Optional) A list of custom rules. The structure is documented below. --- The ` + "`" + `custom_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `(Required) Rule type: ` + "`" + `UTILIZATION` + "`" + ` - This type means that the metric applies to one instance. First, Instance Groups calculates the average metric value for each instance, then averages the values for instances in one availability zone. This type of metric must have the ` + "`" + `instance_id` + "`" + ` label. ` + "`" + `WORKLOAD` + "`" + ` - This type means that the metric applies to instances in one availability zone. This type of metric must have the ` + "`" + `zone_id` + "`" + ` label.`,
				},
				resource.Attribute{
					Name:        "metric_type",
					Description: `(Required) Metric type, ` + "`" + `GAUGE` + "`" + ` or ` + "`" + `COUNTER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) The name of metric.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) Target metric value level. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The instance group creation timestamp.`,
				},
				resource.Attribute{
					Name:        "load_balancer.0.target_group_id",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "load_balancer.0.status_message",
					Description: `The status message of the target group. The ` + "`" + `instances` + "`" + ` block supports:`,
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
					Description: `An array with the network interfaces attached to the managed instance. --- The ` + "`" + `network_interface` + "`" + ` block supports:`,
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
					Description: `The IP version for the public address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The instance group creation timestamp.`,
				},
				resource.Attribute{
					Name:        "load_balancer.0.target_group_id",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "load_balancer.0.status_message",
					Description: `The status message of the target group. The ` + "`" + `instances` + "`" + ` block supports:`,
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
					Description: `An array with the network interfaces attached to the managed instance. --- The ` + "`" + `network_interface` + "`" + ` block supports:`,
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
					Description: `The IP version for the public address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_snapshot",
			Category:         "Yandex Compute Service Resources",
			ShortDescription: `Creates a new snapshot of a disk.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"service",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_disk_id",
					Description: `(Required) ID of the disk to create a snapshot from. - - -`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name for the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the snapshot. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Size of the disk when the snapshot was created, specified in GB.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Size of the snapshot, specified in GB.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the snapshot. ## Import A snapshot can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_snapshot.disk-snapshot shapshot_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disk_size",
					Description: `Size of the disk when the snapshot was created, specified in GB.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Size of the snapshot, specified in GB.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the snapshot. ## Import A snapshot can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_snapshot.disk-snapshot shapshot_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_cr_container_registry",
			Category:         "Yandex Container Registry Resources",
			ShortDescription: `Creates a new container registry.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"cr",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name of the registry.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the registry. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the registry.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the registry. ## Import A registry can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_container_registry.default registry_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Status of the registry.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the registry. ## Import A registry can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_container_registry.default registry_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the service account. Can be updated without creating a new resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the service account.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) ID of the folder that the service account will be created in. Defaults to the provider folder configuration. ## Import A service account can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_iam_service_account.sa account_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_api_key",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IAM service account API key.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"api",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) ID of the service account to an API key for. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the key.`,
				},
				resource.Attribute{
					Name:        "pgp_key",
					Description: `(Optional) An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form ` + "`" + `keybase:keybaseusername` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The secret key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_secret_key",
					Description: `The encrypted secret key, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the secret key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_key",
					Description: `The secret key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_secret_key",
					Description: `The encrypted secret key, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the secret key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_iam_binding",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a single IAM binding for a Yandex IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) The service account ID to apply a binding to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `yandex_iam_service_account_iam_binding` + "`" + ` can be used per role.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_iam_member",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a single member for a single IAM binding for a Yandex IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) The service account ID to apply a policy to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `yandex_iam_service_account_iam_binding` + "`" + ` can be used per role.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) Identity that will be granted the privilege in ` + "`" + `role` + "`" + `. Entry can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_iam_policy",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of the IAM policy for a Yandex IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) The service account ID to apply a policy to.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `yandex_iam_service_account_iam_binding` + "`" + ` can be used per role.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `yandex_iam_service_account_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `yandex_iam_policy` + "`" + ` data source. ## Import Service account IAM policy resources can be imported using the service account ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_iam_service_account_iam_policy.admin-account-iam service_account_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_key",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IAM service account key.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) ID of the service account to create a pair for. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the key pair.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The output format of the keys. ` + "`" + `PEM_FILE` + "`" + ` is the default format.`,
				},
				resource.Attribute{
					Name:        "key_algorithm",
					Description: `(Optional) The algorithm used to generate the key. ` + "`" + `RSA_2048` + "`" + ` is the default algorithm. Valid values are listed in the [API reference](https://cloud.yandex.com/docs/iam/api-ref/Key).`,
				},
				resource.Attribute{
					Name:        "pgp_key",
					Description: `(Optional) An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a keybase username in the form ` + "`" + `keybase:keybaseusername` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The public key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_private_key",
					Description: `The encrypted private key, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the private key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_key",
					Description: `The public key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_private_key",
					Description: `The encrypted private key, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the private key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_static_access_key",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IAM service account static access key.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"static",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) ID of the service account which is used to get a static key. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the service account static key.`,
				},
				resource.Attribute{
					Name:        "pgp_key",
					Description: `(Optional) An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form ` + "`" + `keybase:keybaseusername` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `ID of the static access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Private part of generated static access key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_secret_key",
					Description: `The encrypted secret, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the secret key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_key",
					Description: `ID of the static access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Private part of generated static access key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_secret_key",
					Description: `The encrypted secret, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the secret key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_kubernetes_cluster",
			Category:         "Yandex Managed Service for Kubernetes Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"kubernetes",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of a specific Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the Kubernetes cluster belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the cluster network.`,
				},
				resource.Attribute{
					Name:        "cluster_ipv4_range",
					Description: `(Optional) CIDR block. IP range for allocating pod addresses. It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be set up for this CIDR blocks in node subnets.`,
				},
				resource.Attribute{
					Name:        "service_ipv4_range",
					Description: `(Optional) CIDR block. IP range Kubernetes service Kubernetes cluster IP addresses will be allocated from. It should not overlap with any subnet in the network the Kubernetes cluster located in.`,
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
					Description: `IP allocation policy of the Kubernetes cluster. The structure is documented below. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Computed) ID of a new Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed)Status of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `(Computed) Health of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Computed) The Kubernetes cluster creation timestamp. --- The ` + "`" + `master` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) (Computed) Version of Kubernetes that will be used for master.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) (Computed) Boolean flag. When ` + "`" + `true` + "`" + `, Kubernetes master will have visible ipv4 address.`,
				},
				resource.Attribute{
					Name:        "zonal",
					Description: `(Optional) Initialize parameters for Zonal Master (one node master). The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "regional",
					Description: `(Optional) Initialize parameters for Zonal Master (one node master). The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "version_info",
					Description: `(Computed) Information about cluster version. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "internal_v4_address",
					Description: `(Computed) An IPv4 internal network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "external_v4_address",
					Description: `(Computed) An IPv4 external network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "internal_v4_endpoint",
					Description: `(Computed) Internal endpoint that can be used to connect to the master from cloud networks.`,
				},
				resource.Attribute{
					Name:        "external_v4_endpoint",
					Description: `(Computed) External endpoint that can be used to access Kubernetes cluster API from the internet (outside of the cloud).`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `(Computed) PEM-encoded public certificate that is the root of trust for the Kubernetes cluster. --- The ` + "`" + `zonal` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) ID of the availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet. If no ID is specified, and there only one subnet in specified zone, an address in this subnet will be allocated. --- The ` + "`" + `regional` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Array of locations, where master will be allocated. The structure is documented below. --- The ` + "`" + `location` + "`" + ` block supports repeated values:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) ID of the availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet. --- The ` + "`" + `version_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "current_version",
					Description: `Current Kubernetes version, major.minor (e.g. 1.15).`,
				},
				resource.Attribute{
					Name:        "new_revision_available",
					Description: `Boolean flag. Newer revisions may include Kubernetes patches (e.g 1.15.1 -> 1.15.2) as well as some internal component updates - new features or bug fixes in yandex-specific components either on the master or nodes.`,
				},
				resource.Attribute{
					Name:        "new_revision_summary",
					Description: `Human readable description of the changes to be applied when updating to the latest revision. Empty if new_revision_available is false.`,
				},
				resource.Attribute{
					Name:        "version_deprecated",
					Description: `Boolean flag. The current version is on the deprecation schedule, component (master or node group) should be upgraded. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 15 minute. - ` + "`" + `update` + "`" + ` - Default is 5 minute. - ` + "`" + `delete` + "`" + ` - Default is 5 minute. ## Import A Managed Kubernetes cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_kubernetes_cluster.default cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Computed) ID of a new Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed)Status of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `(Computed) Health of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Computed) The Kubernetes cluster creation timestamp. --- The ` + "`" + `master` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) (Computed) Version of Kubernetes that will be used for master.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) (Computed) Boolean flag. When ` + "`" + `true` + "`" + `, Kubernetes master will have visible ipv4 address.`,
				},
				resource.Attribute{
					Name:        "zonal",
					Description: `(Optional) Initialize parameters for Zonal Master (one node master). The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "regional",
					Description: `(Optional) Initialize parameters for Zonal Master (one node master). The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "version_info",
					Description: `(Computed) Information about cluster version. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "internal_v4_address",
					Description: `(Computed) An IPv4 internal network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "external_v4_address",
					Description: `(Computed) An IPv4 external network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "internal_v4_endpoint",
					Description: `(Computed) Internal endpoint that can be used to connect to the master from cloud networks.`,
				},
				resource.Attribute{
					Name:        "external_v4_endpoint",
					Description: `(Computed) External endpoint that can be used to access Kubernetes cluster API from the internet (outside of the cloud).`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `(Computed) PEM-encoded public certificate that is the root of trust for the Kubernetes cluster. --- The ` + "`" + `zonal` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) ID of the availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet. If no ID is specified, and there only one subnet in specified zone, an address in this subnet will be allocated. --- The ` + "`" + `regional` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Array of locations, where master will be allocated. The structure is documented below. --- The ` + "`" + `location` + "`" + ` block supports repeated values:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) ID of the availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet. --- The ` + "`" + `version_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "current_version",
					Description: `Current Kubernetes version, major.minor (e.g. 1.15).`,
				},
				resource.Attribute{
					Name:        "new_revision_available",
					Description: `Boolean flag. Newer revisions may include Kubernetes patches (e.g 1.15.1 -> 1.15.2) as well as some internal component updates - new features or bug fixes in yandex-specific components either on the master or nodes.`,
				},
				resource.Attribute{
					Name:        "new_revision_summary",
					Description: `Human readable description of the changes to be applied when updating to the latest revision. Empty if new_revision_available is false.`,
				},
				resource.Attribute{
					Name:        "version_deprecated",
					Description: `Boolean flag. The current version is on the deprecation schedule, component (master or node group) should be upgraded. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 15 minute. - ` + "`" + `update` + "`" + ` - Default is 5 minute. - ` + "`" + `delete` + "`" + ` - Default is 5 minute. ## Import A Managed Kubernetes cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_kubernetes_cluster.default cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_kubernetes_node_group",
			Category:         "Yandex Managed Service for Kubernetes Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"kubernetes",
				"node",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the Kubernetes cluster that this node group belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of a specific Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs assigned to the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of Kubernetes that will be used for Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `(Required) Template used to create compute instances in this Kubernetes node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scale_policy",
					Description: `(Required) Scale policy of the node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "allocation_policy",
					Description: `This argument specify subnets (zones), that will be used by node group compute instances. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "instance_group_id",
					Description: `(Computed) ID of instance group that is used to manage this Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `(Computed) Information about maintenance policy for this Kubernetes node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "version_info",
					Description: `(Computed) Information about Kubernetes node group version. The structure is documented below. --- The ` + "`" + `instance_template` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `The ID of the hardware platform configuration for the node group compute instances.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Boolean flag, enables NAT for node group compute instances.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The set of metadata ` + "`" + `key:value` + "`" + ` pairs assigned to this instance template. This includes custom metadata and predefined keys.`,
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
					Description: `The scheduling policy for the instances in node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) Status of the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Computed) The Kubernetes node group creation timestamp. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
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
					Description: `The fixed scaling policy of the instance group. The structure is documented below. --- The ` + "`" + `fixed_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the node group. --- The ` + "`" + `allocation_policy` + "`" + ` block supports:`,
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
					Description: `True/false flag. The current version is on the deprecation schedule, component (master or node group) should be upgraded. ## Import A Yandex Kubernetes Node Group can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_kubernetes_node_group.default node_group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_lb_network_load_balancer",
			Category:         "Yandex Load Balancer Resources",
			ShortDescription: `A network load balancer is used to evenly distribute the load across cloud resources.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the network load balancer. Provided by the client when the network load balancer is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the network load balancer. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder to which the resource belongs. If omitted, the provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to assign to this network load balancer. A list of key/value pairs.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Optional) ID of the availability zone where the network load balancer resides. The default is 'ru-central1'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.`,
				},
				resource.Attribute{
					Name:        "attached_target_group",
					Description: `(Optional) An AttachedTargetGroup resource. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "listener",
					Description: `(Optional) Listener specification that will be used by a network load balancer. The structure is documented below. --- The ` + "`" + `attached_target_group` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `(Required) ID of the target group.`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `(Required) A HealthCheck resource. The structure is documented below. --- The ` + "`" + `healthcheck` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the health check. The name must be unique for each target group that attached to a single load balancer.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The interval between health checks. The default is 2 seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout for a target to return a response for the health check. The default is 1 second.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) Number of failed health checks before changing the status to ` + "`" + `UNHEALTHY` + "`" + `. The default is 2.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) Number of successful health checks required in order to set the ` + "`" + `HEALTHY` + "`" + ` status for the target.`,
				},
				resource.Attribute{
					Name:        "http_options",
					Description: `(Optional) Options for HTTP health check. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tcp_options",
					Description: `(Optional) Options for TCP health check. The structure is documented below. ~>`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port to use for HTTP health checks.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) URL path to set for health checking requests for every target in the target group. For example ` + "`" + `/ping` + "`" + `. The default path is ` + "`" + `/` + "`" + `. --- The ` + "`" + `tcp_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port to use for TCP health checks. --- The ` + "`" + `listener` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the listener. The name must be unique for each listener on a single load balancer.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port for incoming traffic.`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `(Optional) Port of a target. The default is the same as listener's port.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol for incoming traffic. Only tcp network load balancers are currently available.`,
				},
				resource.Attribute{
					Name:        "external_address_spec",
					Description: `(Optional) External IP address specification. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "internal_address_spec",
					Description: `(Optional) Internal IP address specification. The structure is documented below. ~>`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) External IP address for a listener. IP address will be allocated if it wasn't been set.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version of the external addresses that the load balancer works with. Must be one of ipv4 or ipv6. The default is ipv4. --- The ` + "`" + `internal_address_spec` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) ID of the subnet to which the internal IP address belongs.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) Internal IP address for a listener. IP address will be allocated if it wasn't been set.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version of the internal addresses that the load balancer works with. Must be one of ipv4 or ipv6. The default is ipv4. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network load balancer.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The network load balancer creation timestamp. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minute. - ` + "`" + `update` + "`" + ` - Default is 5 minute. - ` + "`" + `delete` + "`" + ` - Default is 5 minute. ## Import A network load balancer can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_lb_network_load_balancer.default network_load_balancer_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network load balancer.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The network load balancer creation timestamp. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minute. - ` + "`" + `update` + "`" + ` - Default is 5 minute. - ` + "`" + `delete` + "`" + ` - Default is 5 minute. ## Import A network load balancer can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_lb_network_load_balancer.default network_load_balancer_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_lb_target_group",
			Category:         "Yandex Load Balancer Resources",
			ShortDescription: `A load balancer distributes the load across cloud resources that are combined into a target group.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"target",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the target group. Provided by the client when the target group is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the target group. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder to which the resource belongs. If omitted, the provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to assign to this target group. A list of key/value pairs.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Optional) ID of the availability zone where the target group resides. The default is 'ru-central1'.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) A Target resource. The structure is documented below. --- The ` + "`" + `target` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) IP address of the target.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) ID of the subnet that targets are connected to. All targets in the target group must be connected to the same subnet within a single availability zone. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The target group creation timestamp. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minute. - ` + "`" + `update` + "`" + ` - Default is 5 minute. - ` + "`" + `delete` + "`" + ` - Default is 5 minute. ## Import A target group can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_lb_target_group.default target_group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The target group creation timestamp. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minute. - ` + "`" + `update` + "`" + ` - Default is 5 minute. - ` + "`" + `delete` + "`" + ` - Default is 5 minute. ## Import A target group can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_lb_target_group.default target_group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_clickhouse_cluster",
			Category:         "Yandex Managed Service for Database Resources",
			ShortDescription: `Manages a ClickHouse cluster within Yandex.Cloud.`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"database",
				"mdb",
				"clickhouse",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the ClickHouse cluster. Provided by the client when the cluster is created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network, to which the ClickHouse cluster belongs.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Required) Deployment environment of the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "clickhouse",
					Description: `(Required) Configuration of the ClickHouse subcluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) A user of the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) A database of the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) A host of the ClickHouse cluster. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the ClickHouse server software.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `(Optional) Time to start the daily backup, in the UTC timezone. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Access policy to the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "zookeeper",
					Description: `(Optional) Configuration of the ZooKeeper subcluster. The structure is documented below. - - - The ` + "`" + `clickhouse` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Resources allocated to hosts of the ClickHouse subcluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `(Required) The ID of the preset for computational resources available to a ClickHouse host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) Volume of the storage available to a ClickHouse host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `(Required) Type of the storage of ClickHouse hosts. The ` + "`" + `zookeeper` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) Resources allocated to hosts of the ZooKeeper subcluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `(Optional) The ID of the preset for computational resources available to a ZooKeeper host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) Volume of the storage available to a ZooKeeper host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `(Optional) Type of the storage of ZooKeeper hosts. The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Optional) Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) The name of the database that the permission grants access to. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Computed) The fully qualified domain name of the host.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the host to be deployed.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The availability zone where the ClickHouse host will be created.`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `(Optional) The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `(Optional) The minute at which backup will be started. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "web_sql",
					Description: `(Optional) Allow access for DataLens.`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `(Optional) Allow access for Web SQL. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_clickhouse_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_clickhouse_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_mongodb_cluster",
			Category:         "Yandex Managed Service for Database Resources",
			ShortDescription: `Manages a MongoDB cluster within Yandex.Cloud.`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"database",
				"mdb",
				"mongodb",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the MongoDB cluster. Provided by the client when the cluster is created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network, to which the MongoDB cluster belongs.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Required) Deployment environment of the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `(Required) Configuration of the MongoDB subcluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) A user of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) A database of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) A host of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Resources allocated to hosts of the MongoDB cluster. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the MongoDB server software.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Access policy to the MongoDB cluster. The structure is documented below. - - - The ` + "`" + `cluster_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version of MongoDB (either 4.2, 4.0 or 3.6).`,
				},
				resource.Attribute{
					Name:        "feature_compatibility_version",
					Description: `(Optional) Feature compatibility version of MongoDB. If not provided version is taken.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `(Optional) Time to start the daily backup, in the UTC timezone. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Shows whether cluster has access to data lens. The structure is documented below. The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `(Optional) The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `(Optional) The minute at which backup will be started. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `(Required) The ID of the preset for computational resources available to a MongoDB host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) Volume of the storage available to a MongoDB host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `(Required) Type of the storage of MongoDB hosts. The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Optional) Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) The name of the database that the permission grants access to. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The fully qualified domain name of the host. Computed on server side.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The availability zone where the MongoDB host will be created.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The role of the cluster (either PRIMARY or SECONDARY).`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `(Computed) The health of the host.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "shard_name",
					Description: `(Optional) The name of the shard to which the host belongs.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) type of mongo daemon which runs on this host (mongod, mongos or monogcfg). Defaults to mongod. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `(Optional) Allow access for DataLens. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
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
					Name:        "cluster_id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "sharded",
					Description: `MongoDB Cluster mode enabled/disabled. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_mongodb_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
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
					Name:        "cluster_id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "sharded",
					Description: `MongoDB Cluster mode enabled/disabled. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_mongodb_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_redis_cluster",
			Category:         "Yandex Managed Service for Database Resources",
			ShortDescription: `Manages a Redis cluster within Yandex.Cloud.`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"database",
				"mdb",
				"redis",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Redis cluster. Provided by the client when the cluster is created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network, to which the Redis cluster belongs.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Required) Deployment environment of the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required) Configuration of the Redis cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Resources allocated to hosts of the Redis cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) A host of the Redis cluster. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "sharded",
					Description: `(Optional) Redis Cluster mode enabled/disabled. - - - The ` + "`" + `config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Close the connection after a client is idle for N seconds.`,
				},
				resource.Attribute{
					Name:        "maxmemory_policy",
					Description: `(Optional) Redis key eviction policy for a dataset that reaches maximum memory. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `(Required) The ID of the preset for computational resources available to a host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-redis/concepts).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) Volume of the storage available to a host, in gigabytes. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The availability zone where the Redis host will be created.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_redis_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_redis_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_cloud_iam_binding",
			Category:         "Yandex Resource Manager Resources",
			ShortDescription: `Allows management of a single IAM binding for a Yandex Resource Manager cloud.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"resourcemanager",
				"cloud",
				"iam",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_id",
					Description: `(Required) ID of the cloud to attach the policy to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be assigned. Only one ` + "`" + `yandex_resourcemanager_cloud_iam_binding` + "`" + ` can be used per role.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) An array of identities that will be granted the privilege in the ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_cloud_iam_member",
			Category:         "Yandex Resource Manager Resources",
			ShortDescription: `Allows management of a single member for a single IAM binding on a Yandex Resource Manager cloud.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"resourcemanager",
				"cloud",
				"iam",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_id",
					Description: `(Required) ID of the cloud to attach a policy to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be assigned.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The identity that will be granted the privilege that is specified in the ` + "`" + `role` + "`" + ` field. This field can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_folder_iam_binding",
			Category:         "Yandex Resource Manager Resources",
			ShortDescription: `Allows management of a single IAM binding for a Yandex Resource Manager folder.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"resourcemanager",
				"folder",
				"iam",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Required) ID of the folder to attach a policy to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be assigned. Only one ` + "`" + `yandex_resourcemanager_folder_iam_binding` + "`" + ` can be used per role.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) An array of identities that will be granted the privilege that is specified in the ` + "`" + `role` + "`" + ` field. Each entry can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_folder_iam_member",
			Category:         "Yandex Resource Manager Resources",
			ShortDescription: `Allows management of a single member for a single IAM binding for a Yandex Resource Manager folder.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"resourcemanager",
				"folder",
				"iam",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Required) ID of the folder to attach a policy to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be assigned.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The identity that will be granted the privilege that is specified in the ` + "`" + `role` + "`" + ` field. This field can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_folder_iam_policy",
			Category:         "Yandex Resource Manager Resources",
			ShortDescription: `Allows management of the IAM policy for a Yandex Resource Manager folder.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"resourcemanager",
				"folder",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Required) ID of the folder that the policy is attached to.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required) The ` + "`" + `yandex_iam_policy` + "`" + ` data source that represents the IAM policy that will be applied to the folder. This policy overrides any existing policy applied to the folder.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_storage_bucket",
			Category:         "Yandex Object Storage",
			ShortDescription: `Allows management of a Yandex.Cloud Storage Bucket.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Optional, Forces new resource) The name of the bucket. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "bucket_prefix",
					Description: `(Optional, Forces new resource) Creates a unique bucket name beginning with the specified prefix. Conflicts with ` + "`" + `bucket` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) The access key to use when applying changes. If omitted, ` + "`" + `storage_access_key` + "`" + ` specified in provider config is used.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The secret key to use when applying changes. If omitted, ` + "`" + `storage_secret_key` + "`" + ` specified in provider config is used.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to ` + "`" + `private` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) A website object (documented below).`,
				},
				resource.Attribute{
					Name:        "cors_rule",
					Description: `(Optional) A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/cors/) (documented below). The ` + "`" + `website` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `(Required) Storage returns this index document when requests are made to the root domain or any of the subfolders.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `(Optional) An absolute path to the document to return in case of a 4XX error. The ` + "`" + `CORS` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. ## Import Storage bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_storage_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. ## Import Storage bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_storage_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_storage_object",
			Category:         "Yandex Object Storage",
			ShortDescription: `Allows management of a Yandex.Cloud Storage Object.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the containing bucket.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The name of the object once it is in the bucket.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional, conflicts with ` + "`" + `content` + "`" + ` and ` + "`" + `content_base64` + "`" + `) The path to a file that will be read and uploaded as raw bytes for the object content.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional, conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `content_base64` + "`" + `) Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `(Optional, conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `content` + "`" + `) Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the ` + "`" + `gzipbase64` + "`" + ` function with small text strings. For larger objects, use ` + "`" + `source` + "`" + ` to stream the content from a disk file.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) The access key to use when applying changes. If omitted, ` + "`" + `storage_access_key` + "`" + ` specified in config is used.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The secret key to use when applying changes. If omitted, ` + "`" + `storage_secret_key` + "`" + ` specified in config is used.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to ` + "`" + `private` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_vpc_network",
			Category:         "Yandex VPC Resources",
			ShortDescription: `Manages a network within Yandex.Cloud.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the network. Provided by the client when the network is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to apply to this network. A list of key/value pairs. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key. ## Import A network can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_network.default network_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key. ## Import A network can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_network.default network_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_vpc_route_table",
			Category:         "Yandex VPC Resources",
			ShortDescription: `A VPC route table is a virtual version of the traditional route table on router device.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"route",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network this route table belongs to. - - -`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the route table. Provided by the client when the route table is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the route table. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder to which the resource belongs. If omitted, the provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to assign to this route table. A list of key/value pairs.`,
				},
				resource.Attribute{
					Name:        "static_route",
					Description: `(Optional) A list of static route records for the route table. The structure is documented below. The ` + "`" + `static_route` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "destination_prefix",
					Description: `Route prefix in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "next_hop_address",
					Description: `Address of the next hop. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the route table. ## Import A route table can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_route_table.default route_table_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the route table. ## Import A route table can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_route_table.default route_table_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_vpc_subnet",
			Category:         "Yandex VPC Resources",
			ShortDescription: `A VPC network is a virtual version of the traditional physical networks that exist within and between physical data centers.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.`,
				},
				resource.Attribute{
					Name:        "v4_cidr_blocks",
					Description: `(Required) A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the subnet. For example, 10.0.0.0/22 or 192.168.0.0/16. Blocks of addresses must be unique and non-overlapping within a network. Minimum subnet size is /28, and maximum subnet size is /16. Only IPv4 is supported.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) Name of the Yandex.Cloud zone for this subnet. - - -`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the subnet. Provided by the client when the subnet is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the subnet. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder to which the resource belongs. If omitted, the provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to assign to this subnet. A list of key/value pairs.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional) The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this subnet. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "v6_cidr_blocks",
					Description: `An optional list of blocks of IPv6 addresses that are owned by this subnet. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 3 minute. - ` + "`" + `update` + "`" + ` - Default is 3 minute. - ` + "`" + `delete` + "`" + ` - Default is 3 minute. ## Import A subnet can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_subnet.default subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "v6_cidr_blocks",
					Description: `An optional list of blocks of IPv6 addresses that are owned by this subnet. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 3 minute. - ` + "`" + `update` + "`" + ` - Default is 3 minute. - ` + "`" + `delete` + "`" + ` - Default is 3 minute. ## Import A subnet can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_subnet.default subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"yandex_compute_disk":                          0,
		"yandex_compute_image":                         1,
		"yandex_compute_instance":                      2,
		"yandex_compute_instance_group":                3,
		"yandex_compute_snapshot":                      4,
		"yandex_cr_container_registry":                 5,
		"yandex_iam_service_account":                   6,
		"yandex_iam_service_account_api_key":           7,
		"yandex_iam_service_account_iam_binding":       8,
		"yandex_iam_service_account_iam_member":        9,
		"yandex_iam_service_account_iam_policy":        10,
		"yandex_iam_service_account_key":               11,
		"yandex_iam_service_account_static_access_key": 12,
		"yandex_kubernetes_cluster":                    13,
		"yandex_kubernetes_node_group":                 14,
		"yandex_lb_network_load_balancer":              15,
		"yandex_lb_target_group":                       16,
		"yandex_mdb_clickhouse_cluster":                17,
		"yandex_mdb_mongodb_cluster":                   18,
		"yandex_mdb_redis_cluster":                     19,
		"yandex_resourcemanager_cloud_iam_binding":     20,
		"yandex_resourcemanager_cloud_iam_member":      21,
		"yandex_resourcemanager_folder_iam_binding":    22,
		"yandex_resourcemanager_folder_iam_member":     23,
		"yandex_resourcemanager_folder_iam_policy":     24,
		"yandex_storage_bucket":                        25,
		"yandex_storage_object":                        26,
		"yandex_vpc_network":                           27,
		"yandex_vpc_route_table":                       28,
		"yandex_vpc_subnet":                            29,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
