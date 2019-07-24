package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Arguments: []resource.Argument{
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
					Name:        "network_interface.0.address",
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
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified DNS name of this instance.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.address",
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
			Arguments: []resource.Argument{
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
					Description: `(Optional) The fixed scaling policy of the instance group. The structure is documented below. --- The ` + "`" + `fixed_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The number of instances in the instance group. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The instance group creation timestamp.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The instance group creation timestamp.`,
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
			Arguments: []resource.Argument{
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
					Description: `Creation timestamp of the snapshot.`,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `Creation timestamp of the snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account",
			Category:         "Yandex IAM Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"service",
				"account",
			},
			Arguments: []resource.Argument{
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
					Description: `(Optional) ID of the folder that the service account will be created in. Defaults to the provider folder configuration. ## Import Service accounts can be imported using their IDs, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_iam_service_account.my_sa service_account_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_iam_binding",
			Category:         "Yandex IAM Resources",
			ShortDescription: `Allows management of a single IAM binding for a Yandex IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"service",
				"account",
				"binding",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_iam_member",
			Category:         "Yandex IAM Resources",
			ShortDescription: `Allows management of a single member for a single IAM binding for a Yandex IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"service",
				"account",
				"member",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_iam_policy",
			Category:         "Yandex IAM Resources",
			ShortDescription: `Allows management of the IAM policy for a Yandex IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"service",
				"account",
				"policy",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_static_access_key",
			Category:         "Yandex IAM Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IAM service account static access key.`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"service",
				"account",
				"static",
				"access",
				"key",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) ID of the service account which is used to get a static key. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the service account static key. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `ID of the static access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Private part of generated static access key.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key. [Yandex Object Storage]: https://cloud.yandex.com/docs/storage/`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "access_key",
					Description: `ID of the static access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Private part of generated static access key.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key. [Yandex Object Storage]: https://cloud.yandex.com/docs/storage/`,
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Required) ID of the folder that the policy is attached to.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required) The ` + "`" + `yandex_iam_policy` + "`" + ` data source that represents the IAM policy that will be applied to the folder. This policy overrides any existing policy applied to the folder.`,
				},
			},
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Arguments: []resource.Argument{
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
					Description: `(Optional) The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this subnet.`,
				},
				resource.Attribute{
					Name:        "v6_cidr_blocks",
					Description: `(Optional) An optional list of blocks of IPv6 addresses that are owned by this subnet. ~>`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"yandex_compute_disk":                          0,
		"yandex_compute_image":                         1,
		"yandex_compute_instance":                      2,
		"yandex_compute_instance_group":                3,
		"yandex_compute_snapshot":                      4,
		"yandex_iam_service_account":                   5,
		"yandex_iam_service_account_iam_binding":       6,
		"yandex_iam_service_account_iam_member":        7,
		"yandex_iam_service_account_iam_policy":        8,
		"yandex_iam_service_account_static_access_key": 9,
		"yandex_resourcemanager_cloud_iam_binding":     10,
		"yandex_resourcemanager_cloud_iam_member":      11,
		"yandex_resourcemanager_folder_iam_binding":    12,
		"yandex_resourcemanager_folder_iam_member":     13,
		"yandex_resourcemanager_folder_iam_policy":     14,
		"yandex_vpc_network":                           15,
		"yandex_vpc_route_table":                       16,
		"yandex_vpc_subnet":                            17,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
