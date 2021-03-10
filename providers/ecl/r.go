package ecl

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ecl_baremetal_keypair_v2",
			Category:         "Baremetal Resources",
			ShortDescription: `Manages a baremetal v2 keypair resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"baremetal",
				"keypair",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the keypair. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) A pre-generated OpenSSH-formatted public key. Changing this creates a new keypair. If a public key is not specified, then a public/private key pair will be automatically generated. If a pair is created, then destroying this resource means you will lose access to that keypair forever. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The generated private key when no public key is specified.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the public key. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_baremetal_keypair_v2.test-keypair my-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The generated private key when no public key is specified.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the public key. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_baremetal_keypair_v2.test-keypair my-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_baremetal_server_v2",
			Category:         "Baremetal Resources",
			ShortDescription: `Manages a baremetal v2 server resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"baremetal",
				"server",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The image ID of the desired image for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional) The name of the desired image for the server. If ` + "`" + `image_id` + "`" + ` is empty, this argument will be used to get the image's id. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Optional; Required if ` + "`" + `flavor_name` + "`" + ` is empty) The flavor ID of the desired flavor for the server. Changing this resizes the existing server.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `(Optional; Required if ` + "`" + `flavor_id` + "`" + ` is empty) The name of the desired flavor for the server. Changing this resizes the existing server.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone in which to create the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Required) An array of one or more networks to attach to the server. The ` + "`" + `networks` + "`" + ` object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the server. Changing this updates the existing server metadata.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The name of a key pair to put on the server. The key pair must already be created and associated with the tenant's account. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "admin_pass",
					Description: `(Optional) The admin user password for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "raid_arrays",
					Description: `(Optional) The raid array information for the server. The ` + "`" + `raid_arrays` + "`" + ` object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "lvm_volume_groups",
					Description: `(Optional) The lvm volume group information for the server. The ` + "`" + `lvm_volume_groups` + "`" + ` object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "filesystems",
					Description: `(Optional) The filesystem information for the server. The ` + "`" + `filesystems` + "`" + ` object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "personality",
					Description: `(Optional) File path and contents to inject into the server. The ` + "`" + `personality` + "`" + ` object structure is documented below. Changing this creates a new server. The ` + "`" + `networks` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `port` + "`" + ` is provided) The network UUID to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` is provided) The port UUID of a network to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) Specifies a fixed IPv4 address to be used on this network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "plane",
					Description: `(Optional) The port UUID of a network to attach to the server. Changing this creates a new server. The ` + "`" + `raid_arrays` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "primary_storage",
					Description: `(Optional) Primary storage flag. At least one storage shoul be primary. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `(Optional) Partition information. The ` + "`" + `partitions` + "`" + ` object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "raid_card_hardware_id",
					Description: `(Optional) Raid card hardware ID. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "disk_hardware_ids",
					Description: `(Optional) List of disk hardware ID. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "raid_level",
					Description: `(Optional) Raid level. Changing this creates a new server. The ` + "`" + `partitions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "lvm",
					Description: `(Optional) LVM flag. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Partition size. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "partition_label",
					Description: `(Optional) Partition label. Changing this creates a new server. The ` + "`" + `lvm_volume_groups` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "vg_label",
					Description: `(Optional) Volume group label. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "physical_volume_partition_labels",
					Description: `(Optional) List of physical volume partition label. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "logical_volumes",
					Description: `(Optional) Logical volume information. The ` + "`" + `logical_volumes` + "`" + ` object structure is documented below. Changing this creates a new server. The ` + "`" + `logical_volumes` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Logical volume size. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "lv_label",
					Description: `(Optional) Logical volume label. Changing this creates a new server. The ` + "`" + `filesystems` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) Filesystem label. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "mount_point",
					Description: `(Optional) Mount point. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type. Changing this creates a new server. The ` + "`" + `personality` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) File path. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "contents",
					Description: `(Optional) Contents. Changing this creates a new server. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_compute_instance_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 Instance resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"instance",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional; Required if ` + "`" + `image_name` + "`" + ` is empty and not booting from a volume. Do not specify if booting from a volume.) The image ID of the desired image for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional; Required if ` + "`" + `image_id` + "`" + ` is empty and not booting from a volume. Do not specify if booting from a volume.) The name of the desired image for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Optional; Required if ` + "`" + `flavor_name` + "`" + ` is empty) The flavor ID of the desired flavor for the server. Changing this resizes the existing server.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `(Optional; Required if ` + "`" + `flavor_id` + "`" + ` is empty) The name of the desired flavor for the server. Changing this resizes the existing server.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "config_drive",
					Description: `(Optional) If true is specified, a configuration drive will be mounted to enable metadata injection in the server. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone in which to create the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) An array of one or more networks to attach to the instance. The network object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the instance. Changing this updates the existing server metadata.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The name of a key pair to put on the server. The key pair must already be created and associated with the tenant's account. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "block_device",
					Description: `(Optional) Configuration of block devices. The block_device structure is documented below. Changing this creates a new server. You can specify multiple block devices which will create an instance with multiple disks. This configuration is very flexible, so please see the above examples for more information.`,
				},
				resource.Attribute{
					Name:        "stop_before_destroy",
					Description: `(Optional) Whether to try stop instance gracefully before destroying it, thus giving chance for guest OS daemons to stop correctly. If instance doesn't stop within timeout, it will be destroyed anyway.`,
				},
				resource.Attribute{
					Name:        "power_state",
					Description: `(Optional) Provide the VM state. Only 'active' and 'shutoff' are supported values.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `port` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The network UUID to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `port` + "`" + ` is provided) The human-readable name of the network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The port UUID of a network to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `(Optional) Specifies a fixed IPv4 address to be used on this network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "access_network",
					Description: `(Optional) Specifies if this network should be used for provisioning access. Accepts true or false. Defaults to false. The ` + "`" + `block_device` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) The source type of the device. Must be one of "blank", "image", "volume", or "snapshot". Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `source_type` + "`" + ` is set to ` + "`" + `"blank"` + "`" + ` ) The UUID of the image, volume, or snapshot. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume to create (in gigabytes). Required in the following combinations: source=image and destination=volume, source=blank and destination=local, and source=blank and destination=volume. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "destination_type",
					Description: `(Optional) The type that gets created. Possible values are "volume" and "local". Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "boot_index",
					Description: `(Optional) The boot index of the volume. It defaults to 0. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `(Optional) Delete the volume / block device upon termination of the instance. Defaults to false. Changing this creates a new server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v4",
					Description: `The Fixed IPv4 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "network/access_network",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_ip_v4",
					Description: `The first detected Fixed IPv4 address.`,
				},
				resource.Attribute{
					Name:        "all_metadata",
					Description: `Contains all instance metadata, even metadata not set by Terraform.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v4",
					Description: `The Fixed IPv4 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "network/access_network",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_ip_v4",
					Description: `The first detected Fixed IPv4 address.`,
				},
				resource.Attribute{
					Name:        "all_metadata",
					Description: `Contains all instance metadata, even metadata not set by Terraform.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_compute_keypair_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 keypair resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"keypair",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the keypair. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) A pre-generated OpenSSH-formatted public key. Changing this creates a new keypair. If a public key is not specified, then a public/private key pair will be automatically generated. If a pair is created, then destroying this resource means you will lose access to that keypair forever. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The generated private key when no public key is specified.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the public key. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The generated private key when no public key is specified.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the public key. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_compute_volume_attach_v2",
			Category:         "Compute Resources",
			ShortDescription: `Attaches a Compute Volume to an Instance.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"volume",
				"attach",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) The ID of the Instance to attach the Volume to.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) The ID of the Volume to attach to an Instance.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Optional) The device of the volume attachment (ex: ` + "`" + `/dev/vdc` + "`" + `). _NOTE_: Being able to specify a device is dependent upon the hypervisor in use. There is a chance that the device specified in Terraform will not be the same device the hypervisor chose. If this happens, Terraform will wish to update the device upon subsequent applying which will cause the volume to be detached and reattached indefinitely. Please use with caution. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_compute_volume_attach_v2.volume_attach_1 <attach_id>/<instance_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_compute_volume_attach_v2.volume_attach_1 <attach_id>/<instance_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_compute_volume_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 volume resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"volume",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the volume to create (in gigabytes). Changing this creates a new volume. User can choice following volume sizes. 1, 15, 40, 80, 100, 300, 500, 1024, 2048, 3072, 4096`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the volume. Changing this updates the volume's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the volume. Changing this updates the volume's description.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone for the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to associate with the volume. Changing this updates the existing volume metadata.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The image ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of volume to create. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "source_replica",
					Description: `(Optional) The volume ID to replicate with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_compute_volume_v2.volume_1 <volume-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_compute_volume_v2.volume_1 <volume-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_dedicated_hypervisor_license_v1",
			Category:         "Dedicated Hypervisor Resources",
			ShortDescription: `Manages a dedicated hypervisor v1 license resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"dedicated",
				"hypervisor",
				"license",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "license_type",
					Description: `(Required) Name of your Guest Image license type as a string. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of the license.`,
				},
				resource.Attribute{
					Name:        "assigned_from",
					Description: `Date the license assigned from.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Expiration date for the license. ## Import Dedicated hypervisor licenses can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_dedicated_hypervisor_license_v1.license_1 0801a388-68e8-4e41-9158-73571117c915 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `Key of the license.`,
				},
				resource.Attribute{
					Name:        "assigned_from",
					Description: `Date the license assigned from.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Expiration date for the license. ## Import Dedicated hypervisor licenses can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_dedicated_hypervisor_license_v1.license_1 0801a388-68e8-4e41-9158-73571117c915 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_dedicated_hypervisor_server_v1",
			Category:         "Dedicated Hypervisor Resources",
			ShortDescription: `Manages a dedicated hypervisor v1 server resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"dedicated",
				"hypervisor",
				"server",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of your Dedicated Hypervisor/Baremetal server as a string.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of your Dedicated Hypervisor server as a string. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "image_ref",
					Description: `(Required) The image reference for the desired image for your Dedicated Hypervisor server. Specify as an UUID or full URL. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "flavor_ref",
					Description: `(Required) The flavor reference for the desired flavor for your Dedicated Hypervisor server. Specify as an UUID or full URL. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone name in which to launch the server. If omit this parameter, target availability_zone is random. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Required) An array of networks to attach to the server. The ` + "`" + `networks` + "`" + ` object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key and value pairs. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "admin_pass",
					Description: `(Optional) Password for the administrator. Changing this creates a new server. The ` + "`" + `networks` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `port` + "`" + ` is provided) The network UUID to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` is provided) The port UUID of a network to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) Specifies a fixed IPv4 address to be used on this network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "plane",
					Description: `(Required) The traffic type of a network to attach to the server. ` + "`" + `data` + "`" + ` and ` + "`" + `storage` + "`" + ` are supported. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `(Required) The segmentation ID of a network to attach to the server. This value is integer, no less than 4 and no more than 4093. Changing this creates a new server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "baremetal_server_id",
					Description: `The UUID of created baremetal server. ## Import Dedicated hypervisor servers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_dedicated_hypervisor_server_v1.server_1 f42dbc37-4642-4628-8b47-50bf95d8fdd5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "baremetal_server_id",
					Description: `The UUID of created baremetal server. ## Import Dedicated hypervisor servers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_dedicated_hypervisor_server_v1.server_1 f42dbc37-4642-4628-8b47-50bf95d8fdd5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_dns_recordset_v2",
			Category:         "DNS Resources",
			ShortDescription: `Manages a V2 recordset resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"recordset",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) Zone ID for the recordset.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) DNS Name for the recordset.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the recordset.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) RRTYPE of the recordset. Valid Values: A | AAAA | MX | CNAME | SRV | SPF | TXT | PTR | NS`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) TTL (Time to Live) for the recordset.`,
				},
				resource.Attribute{
					Name:        "record",
					Description: `(Required) Data for the recordset. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "record",
					Description: `See Argument Reference above. ## Import RecordSet can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_dns_recordset_v2.recordset_1 <recordset-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "record",
					Description: `See Argument Reference above. ## Import RecordSet can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_dns_recordset_v2.recordset_1 <recordset-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_dns_zone_v2",
			Category:         "DNS Resources",
			ShortDescription: `Manages a V2 zone resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"zone",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for this zone.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) E-mail for the zone. Used in SOA records for the zone. This parameter is not currently supported. Even if you set this parameter, it will be ignored.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `(Optional) For secondary zones. The servers to slave from to get DNS information.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) DNS Name for the zone.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL (Time to Live) for the zone. This parameter is not currently supported. Even if you set this parameter, it will be ignored.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of zone. PRIMARY is controlled by ECL2.0 DNS, SECONDARY zones are slaved from another DNS Server. Defaults to PRIMARY. This parameter is not currently supported. Even if you set this parameter, it will be ignored. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns zero.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string. ## Import Zone can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_dns_zone_v2.zone_1 <zone-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns zero.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string. ## Import Zone can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_dns_zone_v2.zone_1 <zone-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_imagestorages_image_v2",
			Category:         "Images Resources",
			ShortDescription: `Manages a V2 image resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"images",
				"imagestorages",
				"image",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "container_format",
					Description: `(Required) Format of the container. Must be "bare".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Internet Gateway resource.`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `(Required) Format of the disk. Must be one of "raw", "qcow2", "iso".`,
				},
				resource.Attribute{
					Name:        "license_switch",
					Description: `(Optional) Switch destination of the license type. Must be one of "WindowsServer_2012R2_Standard_64bit_ComLicense", "WindowsServer_2012_Standard_64bit_ComLicense", "WindowsServer_2008R2_Enterprise_64bit_ComLicense", "WindowsServer_2008R2_Standard_64bit_ComLicense", "WindowsServer_2008_Enterprise_64bit_ComLicense", "WindowsServer_2008_Standard_64bit_ComLicense", "Red_Hat_Enterprise_Linux_6_64bit_BYOL".`,
				},
				resource.Attribute{
					Name:        "local_file_path",
					Description: `(Required) This is the filepath of the raw image file that will be uploaded to Glance.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `(Optional) Amount of disk space (in GB) required to boot image. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `(Optional) Amount of ram (in MB) required to boot image. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Descriptive name for the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `(Optional) If true, image will not be deletable. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) String related to the image.`,
				},
				resource.Attribute{
					Name:        "verify_checksum",
					Description: `(Optional) If false, the checksum will not be verified once the image is finished uploading. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) Scope of image accessibility. Must be one of "public", "private". Defaults to "private".`,
				},
				resource.Attribute{
					Name:        "peroperties",
					Description: `(Optional) A map of key/value pairs to set freeform information about an image. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `URL for the virtual machine image file.`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `md5 hash of image contents.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time of image registration.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The location metadata.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Owner of the image.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `URL for schema of the virtual machine image.`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `Size of image file in bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the image.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date and time of the last image modification. ## Import Images can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_imagestorages_image_v2.image_1 Temp_Terraform_AccTest ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `URL for the virtual machine image file.`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `md5 hash of image contents.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time of image registration.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The location metadata.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Owner of the image.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `URL for schema of the virtual machine image.`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `Size of image file in bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the image.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date and time of the last image modification. ## Import Images can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_imagestorages_image_v2.image_1 Temp_Terraform_AccTest ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_imagestorages_member_accepter_v2",
			Category:         "Images Resources",
			ShortDescription: `Manages a V2 Image member resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"images",
				"imagestorages",
				"member",
				"accepter",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "image_member_id",
					Description: `(Required) An identifier for the image and member. You can refer it from ID of member resource. The format is "${image_id}/${member_id}", where member_id is accepter project ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) The status of this image member. Must be one of "pending", "accepted", "rejected". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_imagestorages_member_v2",
			Category:         "Images Resources",
			ShortDescription: `Manages a V2 Image member resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"images",
				"imagestorages",
				"member",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) An identifier for the image.`,
				},
				resource.Attribute{
					Name:        "member_id",
					Description: `(Required) An identifier for the image member (projectID). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time of image member creation.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `URL for schema of the member.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this image member.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date and time of last modification of image member.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time of image member creation.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `URL for schema of the member.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this image member.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date and time of last modification of image member.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_common_function_gateway_v2",
			Category:         "Network Resources",
			ShortDescription: `Manages a V2 common function gateway resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"common",
				"function",
				"gateway",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Common Function Gateway resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Common Function Gateway resource.`,
				},
				resource.Attribute{
					Name:        "common_function_pool_id",
					Description: `(Required) Common Function Pool instantiated by this Gateway.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "common_function_pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of automatically created network connected to this Common Function Gateway.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of automatically created subnet connected to this Common Function Gateway (using link-local address). ## Import Common Function Gateway can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_common_function_gateway_v2.common_function_gateway_1 <common-function-gateway-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "common_function_pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID of automatically created network connected to this Common Function Gateway.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of automatically created subnet connected to this Common Function Gateway (using link-local address). ## Import Common Function Gateway can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_common_function_gateway_v2.common_function_gateway_1 <common-function-gateway-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_gateway_interface_v2",
			Category:         "Network Resources",
			ShortDescription: `Manages a V2 gateway interface resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"gateway",
				"interface",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "aws_gw_id",
					Description: `(Optional) AWS Gateway to which this port is connected. Conflicts with "azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "azure_gw_id",
					Description: `(Optional) Azure Gateway to which this port is connected. Conflicts with "aws_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Gateway Interface resource.`,
				},
				resource.Attribute{
					Name:        "fic_gw_id",
					Description: `(Optional) FIC Gateway to which this port is connected. Conflicts with "aws_gw_id", "azure_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "gcp_gw_id",
					Description: `(Optional) GCP Gateway to which this port is connected. Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "gw_vipv4",
					Description: `(Required) IP version 4 address to be assigned virtual router on VRRP.`,
				},
				resource.Attribute{
					Name:        "interdc_gw_id",
					Description: `(Optional) Inter DC Gateway to which this port is connected. Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "internet_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `(Optional) Internet GW to which this port is connected. Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Gateway Interface resource.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Required) Netmask for IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) Network connected to this interface.`,
				},
				resource.Attribute{
					Name:        "primary_ipv4",
					Description: `(Required) IP version 4 address to be assigned to primary device on VRRP.`,
				},
				resource.Attribute{
					Name:        "secondary_ipv4",
					Description: `(Required) IP version 4 address to be assigned to secondary device on VRRP.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Required) Service type for this interface. Must be one of "aws", "azure", "fic", "gcp", "interdc", "internet" and "vpn".`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID).`,
				},
				resource.Attribute{
					Name:        "vpn_gw_id",
					Description: `(Optional) VPN Gateway to which this port is connected. Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id" and "internet_gw_id".`,
				},
				resource.Attribute{
					Name:        "vrid",
					Description: `(Required) VRRP Group ID for this GW Interface. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gw_vipv6",
					Description: `IP version 6 address to be assigned virtual router on VRRP.`,
				},
				resource.Attribute{
					Name:        "primary_ipv6",
					Description: `IP version 6 address to be assigned to primary device on VRRP.`,
				},
				resource.Attribute{
					Name:        "secondary_ipv6",
					Description: `IP version 6 address to be assigned to secondary device on VRRP. ## Import Gateway interfaces can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_gateway_interface_v2.gateway_interface_1 12610e1b-f675-437b-8b1a-f4d19f92421e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gw_vipv6",
					Description: `IP version 6 address to be assigned virtual router on VRRP.`,
				},
				resource.Attribute{
					Name:        "primary_ipv6",
					Description: `IP version 6 address to be assigned to primary device on VRRP.`,
				},
				resource.Attribute{
					Name:        "secondary_ipv6",
					Description: `IP version 6 address to be assigned to secondary device on VRRP. ## Import Gateway interfaces can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_gateway_interface_v2.gateway_interface_1 12610e1b-f675-437b-8b1a-f4d19f92421e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_internet_gateway_v2",
			Category:         "Network Resources",
			ShortDescription: `Manages a V2 internet gateway resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"internet",
				"gateway",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Internet Gateway resource.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `(Required) Internet Service instantiated by Internet gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Internet Gateway resource.`,
				},
				resource.Attribute{
					Name:        "qos_option_id",
					Description: `(Required) Quality of Service options selected for Internet gateway.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Internet gateways can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_internet_gateway_v2.internet_gateway_1 Terraform_Test_Internet_Gateway_01 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Internet gateways can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_internet_gateway_v2.internet_gateway_1 Terraform_Test_Internet_Gateway_01 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_load_balancer_v2",
			Category:         "Network Resources",
			ShortDescription: `Manages a V2 Load Balancer resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"load",
				"balancer",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone in which to create the Load Balancer. Changing this creates a new Load Balancer.`,
				},
				resource.Attribute{
					Name:        "default_gateway",
					Description: `(Optional) IP address of default gateway. The default gateway IP address must be in the network connected to the Load Balancer Interface defined as the argument ` + "`" + `interfaces` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Load Balancer description.`,
				},
				resource.Attribute{
					Name:        "load_balancer_plan_id",
					Description: `(Required) The UUID of Load Balancer Plan uses by the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) An array of connected interfaces in Load Balancer. The ` + "`" + `interfaces` + "`" + ` object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "syslog_servers",
					Description: `(Optional) An array of running syslog servers included in Load Balancer. The ` + "`" + `syslog_servers` + "`" + ` object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the Load Balancer. Required if admin wants to create a Load Balancer for another tenant. Changing this creates a new Load Balancer. The ` + "`" + `interfaces` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Load Balancer Interface description.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The physical IP address associated with the interface. The IP address must be in the network specified as the argument ` + "`" + `network_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Load Balancer Interface.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The UUID of the network associated with the interface.`,
				},
				resource.Attribute{
					Name:        "slot_number",
					Description: `(Required) The slot number of interface.`,
				},
				resource.Attribute{
					Name:        "virtual_ip_address",
					Description: `(Optional; Required if ` + "`" + `virtual_ip_properties` + "`" + ` is not empty) The virtual IP address associated with the interface. The IP address must be in the network specified as the argument ` + "`" + `network_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "virtual_ip_properties",
					Description: `(Optional; Required if ` + "`" + `virtual_ip_address` + "`" + ` is not empty) Properties used for virtual IP address. The ` + "`" + `virtual_ip_properties` + "`" + ` object structure is documented below. The ` + "`" + `virtual_ip_properties` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Redundancy Protocol. Must be "vrrp".`,
				},
				resource.Attribute{
					Name:        "vrid",
					Description: `(Required) VRRP group identifier. This value is integer, no less than 1 and no more than 255. The ` + "`" + `syslog_servers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "acl_logging",
					Description: `(Optional) Should syslog record acl info. Must be one of "ENABLED" and "DISABLED".`,
				},
				resource.Attribute{
					Name:        "appflow_logging",
					Description: `(Optional) Should syslog record appflow info. Must be one of "ENABLED" and "DISABLED".`,
				},
				resource.Attribute{
					Name:        "date_format",
					Description: `(Optional) Date format utilized by syslog. Must be one of "DDMMYYYY", "MMDDYYYY" and "YYYYMMDD".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Load Balancer Syslog Server description.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) IP address of syslog server. The syslog server IP address must be in the network connected to the Load Balancer Interface defined as the argument ` + "`" + `interfaces` + "`" + `. Changing this creates a new syslog server.`,
				},
				resource.Attribute{
					Name:        "log_facility",
					Description: `(Optional) Log facility for syslog. Must be one of "LOCAL0", "LOCAL1", "LOCAL2", "LOCAL3", "LOCAL4", "LOCAL5", "LOCAL6" and "LOCAL7".`,
				},
				resource.Attribute{
					Name:        "log_level",
					Description: `(Optional) Valid elements for log_level are "ALERT", "CRITICAL", "EMERGENCY", "INFORMATIONAL", "NOTICE", "ALL", "DEBUG", "ERROR", "NONE", "WARNING". ` + "`" + `log_level` + "`" + ` value can be assigned combining multiple elements as "ALERT|CRITICAL|EMERGENCY". Caution: Can not combine "ALL" or "NONE" with the others.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Load Balancer Syslog Server. Changing this creates a new syslog server.`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `(Optional) The port number of syslog server. This value is integer, no less than 1 and no more than 65535. Changing this creates a new syslog server.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of syslog server. This value is integer, no less than 0 and no more than 255.`,
				},
				resource.Attribute{
					Name:        "tcp_logging",
					Description: `(Optional) Should syslog record tcp protocol info. Must be one of "NONE" and "ALL".`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the syslog server. Required if admin wants to create a syslog server for another tenant. Changing this creates a new syslog server.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional) Time zone utilized by syslog. Must be one of "GMT_TIME" and "LOCAL_TIME".`,
				},
				resource.Attribute{
					Name:        "transport_type",
					Description: `(Optional) Protocol for syslog transport. Must be "UDP".`,
				},
				resource.Attribute{
					Name:        "user_configurable_log_messages",
					Description: `(Optional) Can user configure log messages. Must be one of "YES" and "NO". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Load Balancer unique ID.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `Admin’s password placeholder.`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `Username with admin access to Load Balancer VM instance.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `An array of connected interfaces in Load Balancer. This excludes disconnected interfaces. The ` + "`" + `interfaces` + "`" + ` object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of Load Balancer.`,
				},
				resource.Attribute{
					Name:        "syslog_servers",
					Description: `An array of running syslog servers included in Load Balancer. The ` + "`" + `syslog_servers` + "`" + ` object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_password",
					Description: `User’s password placeholder.`,
				},
				resource.Attribute{
					Name:        "user_username",
					Description: `Username with user access to Load Balancer VM instance. The ` + "`" + `interfaces` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Load Balancer Interface unique ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of Load Balancer Interface. The ` + "`" + `syslog_servers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Load Balancer Syslog Server unique ID.`,
				},
				resource.Attribute{
					Name:        "acl_logging",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "appflow_logging",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "date_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "log_facility",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "log_level",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of Load Balancer Syslog Server.`,
				},
				resource.Attribute{
					Name:        "tcp_logging",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transport_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_configurable_log_messages",
					Description: `See Argument Reference above. ## Import Load Balancer can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_load_balancer_v2.load_balancer_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Load Balancer unique ID.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `Admin’s password placeholder.`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `Username with admin access to Load Balancer VM instance.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `An array of connected interfaces in Load Balancer. This excludes disconnected interfaces. The ` + "`" + `interfaces` + "`" + ` object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of Load Balancer.`,
				},
				resource.Attribute{
					Name:        "syslog_servers",
					Description: `An array of running syslog servers included in Load Balancer. The ` + "`" + `syslog_servers` + "`" + ` object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_password",
					Description: `User’s password placeholder.`,
				},
				resource.Attribute{
					Name:        "user_username",
					Description: `Username with user access to Load Balancer VM instance. The ` + "`" + `interfaces` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Load Balancer Interface unique ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of Load Balancer Interface. The ` + "`" + `syslog_servers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Load Balancer Syslog Server unique ID.`,
				},
				resource.Attribute{
					Name:        "acl_logging",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "appflow_logging",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "date_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "log_facility",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "log_level",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of Load Balancer Syslog Server.`,
				},
				resource.Attribute{
					Name:        "tcp_logging",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transport_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_configurable_log_messages",
					Description: `See Argument Reference above. ## Import Load Balancer can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_load_balancer_v2.load_balancer_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_network_v2",
			Category:         "Network Resources",
			ShortDescription: `Manages a V2 network resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network. Acceptable values are "true" and "false". Changing this value updates the state of the existing network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Network description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the network. Changing this updates the name of the existing network.`,
				},
				resource.Attribute{
					Name:        "plane",
					Description: `(Optional) The plane of the network. Allowed values are "data" and "storage". Changing this creates a new network.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Network tags.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the network. Required if admin wants to create a network for another tenant. Changing this creates a new network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "plane",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The network status.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The associated subnets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "plane",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The network status.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The associated subnets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_port_v2",
			Category:         "Network Resources",
			ShortDescription: `Manages a V2 port resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"port",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the port (must be "true" or "false" if provided). Changing this updates the ` + "`" + `admin_state_up` + "`" + ` of an existing port.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `(Optional) An IP/MAC Address pair of additional IP addresses that can be active on this port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Port description.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) The ID of the device attached to the port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `(Optional) The device owner of the Port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional - Conflicts with ` + "`" + `no_fixed_ip` + "`" + `) An array of desired IPs for this port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) Specify a specific MAC address for the port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the port. Changing this updates the ` + "`" + `name` + "`" + ` of an existing port.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The ID of the network to attach the port to. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "no_fixed_ip",
					Description: `(Optional - Conflicts with ` + "`" + `fixed_ip` + "`" + `) Create a port with no fixed IP address. This will also remove any fixed IPs previously set on a port. ` + "`" + `true` + "`" + ` is the only valid value for this argument.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `(Optional) The segmentation ID used for this port. User can specify this value only in case segmentation type is "vlan".`,
				},
				resource.Attribute{
					Name:        "segmentation_type",
					Description: `(Optional) The segmentation type used for this port. User can use "vlan" of "flat" as this argument`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Port tags.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the Port. Required if admin wants to create a port for another tenant. Changing this creates a new port. The ` + "`" + `fixed_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet in which to allocate IP address for this port.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address desired in the subnet for this port. If you don't specify ` + "`" + `ip_address` + "`" + `, an available IP address from the specified subnet will be allocated to this port. This field will not be populated if it is left blank or omitted. To retrieve the assigned IP address, use the ` + "`" + `all_fixed_ips` + "`" + ` attribute.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The additional IP address.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The additional MAC address. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status for the Port.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see Enterprise Cloud 2.0 Knowledge Center documents for further information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status for the Port.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see Enterprise Cloud 2.0 Knowledge Center documents for further information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_public_ip_v2",
			Category:         "Network Resources",
			ShortDescription: `Manages a V2 public ip resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"public",
				"ip",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Public IP resource.`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `(Required) Internet Gateway the block will be assigned to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Public IP resource.`,
				},
				resource.Attribute{
					Name:        "submask_length",
					Description: `(Required) Specifies the size of the block by the length of its subnetwork mask length.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The IP address of the block (assigned automatically).`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Pulic ips can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_public_ip_v2.public_ip_1 Terraform_Test_Public_IP_01 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The IP address of the block (assigned automatically).`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Pulic ips can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_public_ip_v2.public_ip_1 Terraform_Test_Public_IP_01 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_static_route_v2",
			Category:         "Network Resources",
			ShortDescription: `Manages a V2 static route resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"static",
				"route",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "aws_gw_id",
					Description: `(Optional) AWS Gateway on which this static route will be set. Conflicts with "azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "azure_gw_id",
					Description: `(Optional) Azure Gateway on which this static route will be set. Conflicts with "aws_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Static Route resource.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) CIDR this static route points to.`,
				},
				resource.Attribute{
					Name:        "fic_gw_id",
					Description: `(Optional) FIC Gateway on which this static route will be set. Conflicts with "aws_gw_id", "azure_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "gcp_gw_id",
					Description: `(Optional) GCP Gateway on which this static route will be set. Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "interdc_gw_id", "internet_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "interdc_gw_id",
					Description: `(Optional) Inter DC Gateway on which this static route will be set. Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "internet_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `(Optional) Internet Gateway on which this static route will be set. Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id" and "vpn_gw_id".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Static Route resource.`,
				},
				resource.Attribute{
					Name:        "nexthop",
					Description: `(Required) Next Hop address for specified CIDR.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Required) Service type for this route. Must be one of "aws", "azure", "gcp", "interdc", "internet" and "vpn".`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID).`,
				},
				resource.Attribute{
					Name:        "vpn_gw_id",
					Description: `(Optional) VPN Gateway on which this static route will be set. Conflicts with "aws_gw_id", "azure_gw_id", "fic_gw_id", "gcp_gw_id", "interdc_gw_id" and "internet_gw_id". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Static routes can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_static_route_v2.static_route_1 Terraform_Test_Static_Route_01 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Static routes can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_static_route_v2.static_route_1 Terraform_Test_Static_Route_01 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_subnet_v2",
			Category:         "Network Resources",
			ShortDescription: `Manages a V2 subnet resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"subnet",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `(Optional) An array of sub-ranges of CIDR available for dynamic allocation to ports. The allocation_pool object structure is documented below. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) CIDR representing IP range for this subnet, based on IP version. You can omit this option if you are creating a subnet from a subnet pool.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Subnet description.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `(Optional) List of subnet dns name servers.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) The administrative state of the network. Acceptable values are "true" and "false". Changing this value enables or disables the DHCP capabilities of the existing subnet. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `(Optional) An array of routes that should be used by devices with IPs from this subnet (not including local subnet route). The host_route object structure is documented below. Changing this updates the host routes for the existing subnet.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version. In Enterprise Cloud service this parameter is fixed as 4.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Default gateway used by devices in this subnet. Leaving this blank and not setting ` + "`" + `no_gateway` + "`" + ` will cause a default gateway of ` + "`" + `.1` + "`" + ` to be used. Changing this updates the gateway IP of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subnet. Changing this updates the name of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The UUID of the parent network. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `(Optional) List of ntp servers.`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `(Optional) Do not set a gateway IP on this subnet. Changing this removes or adds a default gateway IP of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Subnet tags.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the subnet. Required if admin wants to create a subnet for another tenant. Changing this creates a new subnet. The ` + "`" + `allocation_pools` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The starting address.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) The ending address. The ` + "`" + `host_routes` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `(Required) The destination CIDR.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) The next hop in the route. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_mode",
					Description: `Address mode for IPv6 (not supported).`,
				},
				resource.Attribute{
					Name:        "ipv6_ra_mode",
					Description: `IPv6 router advertisement mode (not supported).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `True if gateway_ip is Nil`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Hidden Subnet status.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Subnet can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_mode",
					Description: `Address mode for IPv6 (not supported).`,
				},
				resource.Attribute{
					Name:        "ipv6_ra_mode",
					Description: `IPv6 router advertisement mode (not supported).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `True if gateway_ip is Nil`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Hidden Subnet status.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Subnet can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_network_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "provider_connectivity_tenant_connection_request_v2",
			Category:         "Provider Connectivity Resources",
			ShortDescription: `Manages a v2 Tenant Connection Request resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"provider",
				"connectivity",
				"tenant",
				"connection",
				"request",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_id_other",
					Description: `(Required) The owner tenant of network.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) Network unique id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of tenant_connection_request.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of tenant_connection_request.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) tenant_connection_request tags. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `tenant_connection_request unique ID.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the owner.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of tenant_connection_request.`,
				},
				resource.Attribute{
					Name:        "approval_request_id",
					Description: `SSS approval_request ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `tenant_connection_request unique ID.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the owner.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of tenant_connection_request.`,
				},
				resource.Attribute{
					Name:        "approval_request_id",
					Description: `SSS approval_request ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_provider_connectivity_tenant_connection_v2",
			Category:         "Provider Connectivity Resources",
			ShortDescription: `Manages a v2 Tenant Connection resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"provider",
				"connectivity",
				"tenant",
				"connection",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of tenant_connection.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of tenant_connection.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) tenant_connection tags.`,
				},
				resource.Attribute{
					Name:        "tenant_connection_request_id",
					Description: `(Required) ID of the tenant_connection_request.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Required) device type. (ECL::Compute::Server/ECL::Baremetal::Server/ECL::VirtualNetworkAppliance::VSRX)`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) ID of the device of the device_type.`,
				},
				resource.Attribute{
					Name:        "device_interface_id",
					Description: `(Optional) ID of the interface of the device. Required for device_type in ECL::Baremetal::Server, ECL::VirtualNetworkAppliance::VSRX. For device_type: ECL::Baremetal::Server, network_physical_port_id should be used. For ECL::VirtualNetworkAppliance::VSRX, interfaces.interface_<slot_number> should be used.`,
				},
				resource.Attribute{
					Name:        "attachment_opts_compute",
					Description: `(Optional) Additional options for tenant_connection. The ` + "`" + `attachment_opts_compute` + "`" + ` object structure is documented below. It is proxied to create a connection for the Compute Server resource.`,
				},
				resource.Attribute{
					Name:        "attachment_opts_baremetal",
					Description: `(Optional) Additional options for tenant_connection. The ` + "`" + `attachment_opts_baremetal` + "`" + ` object structure is documented below. It is proxied to create a connection for the Baremetal Server resource.`,
				},
				resource.Attribute{
					Name:        "attachment_opts_vna",
					Description: `(Optional) Additional options for tenant_connection. The ` + "`" + `attachment_opts_vna` + "`" + ` object structure is documented below. It is proxied to create a connection for the Virtual Network Appliance resource. The ` + "`" + `attachment_opts_compute` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fixed_ips",
					Description: `(Optional) Array of IP address assignment objects, attached to port.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address assigned to port.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of subnet from which IP address is allocated.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `(Optional) Array of Allowed address pairs.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address assigned to port for Allowed address pairs.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) MAC address assigned to port for Allowed address pairs. The ` + "`" + `attachment_opts_baremetal` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "segmentation_type",
					Description: `(Optional) Segmentation type used for port. Only valid for device_type = ECL::Baremetal::Server. (flat/vlan)`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `(Optional) Segmentation id used for port. Only valid for device_type = ECL::Baremetal::Server.`,
				},
				resource.Attribute{
					Name:        "fixed_ips",
					Description: `(Optional) Array of IP address assignment objects, attached to port.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address assigned to port.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of subnet from which IP address is allocated.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `(Optional) Array of Allowed address pairs.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address assigned to port for Allowed address pairs.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) MAC address assigned to port for Allowed address pairs. The ` + "`" + `attachment_opts_vna` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fixed_ips",
					Description: `(Optional) Array of IP address assignment objects, attached to port.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address assign to Interface within subnet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `tenant_connection unique ID.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the owner.`,
				},
				resource.Attribute{
					Name:        "tenant_id_other",
					Description: `The owner tenant of network.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `Network unique id.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Port unique id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of tenant_connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `tenant_connection unique ID.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the owner.`,
				},
				resource.Attribute{
					Name:        "tenant_id_other",
					Description: `The owner tenant of network.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `Network unique id.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Port unique id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of tenant_connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_rca_user_v1",
			Category:         "RCA Resources",
			ShortDescription: `Manages a rca v1 user resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"rca",
				"user",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of VPN connection. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User’s name of VPN connection.`,
				},
				resource.Attribute{
					Name:        "vpn_endpoints",
					Description: `List of VPN endpoint user can connect.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `URL of VPN endpoint.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of VPN endpoint. ## Import RCA users can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_rca_user_v1.user_1 8bbe05d4bec747189e0dab81e486969f-1005 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `User’s name of VPN connection.`,
				},
				resource.Attribute{
					Name:        "vpn_endpoints",
					Description: `List of VPN endpoint user can connect.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `URL of VPN endpoint.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of VPN endpoint. ## Import RCA users can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_rca_user_v1.user_1 8bbe05d4bec747189e0dab81e486969f-1005 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_security_host_based_v1",
			Category:         "Security Resources",
			ShortDescription: `Manages a V1 Host Based Security resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"host",
				"based",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Tenant ID of the owner (UUID).`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `(Optional) Messages are displayed in Japanese or English depending on this value. ja: Japanese, en: English. Default value is "en".`,
				},
				resource.Attribute{
					Name:        "service_order_service",
					Description: `(Required) Requested menu. Set "Managed Anti-Virus", "Managed Virtual Patch" or "Managed Host-based Security Package" to this field.`,
				},
				resource.Attribute{
					Name:        "max_agent_value",
					Description: `(Required) Set maximum quantity of Agent usage.`,
				},
				resource.Attribute{
					Name:        "mail_address",
					Description: `(Required) Contact e-mail address.`,
				},
				resource.Attribute{
					Name:        "dsm_langress",
					Description: `(Required) This value is used for language of Deep Security Manager. ja: Japanese, en: English.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Required) Set "Asia/Tokyo" for JST or "Etc/GMT" for UTC.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_security_network_based_device_ha_v1",
			Category:         "Security Resources",
			ShortDescription: `Manages a V1 Network Based Device(HA) resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"network",
				"based",
				"device",
				"ha",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Tenant ID of the owner (UUID).`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `(Optional) Messages are displayed in Japanese or English depending on this value. ja: Japanese, en: English. Default value is "en".`,
				},
				resource.Attribute{
					Name:        "operating_mode",
					Description: `(Required) Set "FW" or "UTM" to this value.`,
				},
				resource.Attribute{
					Name:        "license_kind",
					Description: `(Required) Set "02" or "08" as FW/UTM plan.`,
				},
				resource.Attribute{
					Name:        "host_1_az_group",
					Description: `(Required) Set availability zone for HA Host 1.`,
				},
				resource.Attribute{
					Name:        "host_2_az_group",
					Description: `(Required) Set availability zone for HA Host 2.`,
				},
				resource.Attribute{
					Name:        "ha_link_1",
					Description: `(Required) Set HA line information for HA Host 1.`,
				},
				resource.Attribute{
					Name:        "ha_link_2",
					Description: `(Required) Set HA line information for HA Host 2.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Set port information. Both ` + "`" + `ha_link_1` + "`" + ` and ` + "`" + `ha_link_2` + "`" + ` block support:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) Set the Network ID to be used for HA line.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Set the Subnet ID to be used for HA line.`,
				},
				resource.Attribute{
					Name:        "host_1_ip_address",
					Description: `(Required) Set IPv4 address for as HA Host 1 in this HA link.`,
				},
				resource.Attribute{
					Name:        "host_2_ip_address",
					Description: `(Required) Set IPv4 address for as HA Host 2 in this HA link. The ` + "`" + `port` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Required) Set "true" to enable the port "false" to disable the port.`,
				},
				resource.Attribute{
					Name:        "host_1_ip_address",
					Description: `(Required in case enabling the port) IP Address of the port for HA Host 1.`,
				},
				resource.Attribute{
					Name:        "host_1_ip_address_prefix",
					Description: `(Required in case enabling the port) IP Address prefix of the port for HA Host 1.`,
				},
				resource.Attribute{
					Name:        "host_2_ip_address",
					Description: `(Required in case enabling the port) IP Address of the port for HA Host 2.`,
				},
				resource.Attribute{
					Name:        "host_2_ip_address_prefix",
					Description: `(Required in case enabling the port) IP Address prefix of the port for HA Host 2.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required in case enabling the port) Network ID to which the port is associated.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required in case enabling the port) UUID Subnet IDto which the port is associated.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Required in case enabling the port) MTU value in the configuration of the port.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Required in case enabling the port) Comments for the port.`,
				},
				resource.Attribute{
					Name:        "enable_ping",
					Description: `(Required in case enabling the port) Set "true" to enable ping response, "false" to disable ping response. Note: Type of this value is "string".`,
				},
				resource.Attribute{
					Name:        "vrrp_grp_id",
					Description: `(Required in case enabling the port) VRRP Group ID. This value must be in the range of "1" to "100". Note: Type of this value is "string".`,
				},
				resource.Attribute{
					Name:        "vrrp_id",
					Description: `(Required in case enabling the port) VRRP ID. This value must be in the range of "1" to "100". Note: Type of this value is "string".`,
				},
				resource.Attribute{
					Name:        "vrrp_ip_address",
					Description: `(Required in case enabling the port) VRRP IP Address of this port.`,
				},
				resource.Attribute{
					Name:        "preempt",
					Description: `(Required in case enabling the port) "true" to enable preempt option, "false" to disable preempt option. Note: Type of this value is "string". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "port/host_1_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/host/1_ip_address_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/host_2_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/host/2_ip_address_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/comment",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "port/host_1_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/host/1_ip_address_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/host_2_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/host/2_ip_address_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/comment",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_security_network_based_device_single_v1",
			Category:         "Security Resources",
			ShortDescription: `Manages a V1 Network Based Device(Single) resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"network",
				"based",
				"device",
				"single",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Tenant ID of the owner (UUID).`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `(Optional) Messages are displayed in Japanese or English depending on this value. ja: Japanese, en: English. Default value is "en".`,
				},
				resource.Attribute{
					Name:        "operating_mode",
					Description: `(Required) Set "FW" or "UTM" to this value.`,
				},
				resource.Attribute{
					Name:        "license_kind",
					Description: `(Required) Set "02" or "08" as FW/UTM plan.`,
				},
				resource.Attribute{
					Name:        "az_group",
					Description: `(Required) Set availability zone.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Set port information. The ` + "`" + `port` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Required) Set "true" to enable the port "false" to disable the port.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required in case enabling the port) IP Address of the port.`,
				},
				resource.Attribute{
					Name:        "ip_address_prefix",
					Description: `(Required in case enabling the port) IP Address prefix of the port.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required in case enabling the port) Network ID to which the port is associated.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required in case enabling the port) UUID Subnet IDto which the port is associated.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Required in case enabling the port) MTU value in the configuration of the port.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Required in case enabling the port) Comments for the port. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "port/ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/ip_address_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/comment",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "port/ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/ip_address_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/comment",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_security_network_based_waf_single_v1",
			Category:         "Security Resources",
			ShortDescription: `Manages a V1 Network Based WAF(Single) resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"network",
				"based",
				"waf",
				"single",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Tenant ID of the owner (UUID).`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `(Optional) Messages are displayed in Japanese or English depending on this value. ja: Japanese, en: English. Default value is "en".`,
				},
				resource.Attribute{
					Name:        "license_kind",
					Description: `(Required) Set "02" or "04" or "08" as WAF plan.`,
				},
				resource.Attribute{
					Name:        "az_group",
					Description: `(Required) Set availability zone.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Set port information. The ` + "`" + `port` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Required) Set "true" to enable the port "false" to disable the port.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required in case enabling the port) IP Address of the port.`,
				},
				resource.Attribute{
					Name:        "ip_address_prefix",
					Description: `(Required in case enabling the port) IP Address prefix of the port`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required in case enabling the port) Network ID to which the port is associated.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required in case enabling the port) UUID Subnet IDto which the port is associated.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Required in case enabling the port) MTU value in the configuration of the port.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Required in case enabling the port) Comments for the port. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "port/ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/ip_address_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/comment",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "port/ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/ip_address_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port/comment",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_sss_approval_request_v1",
			Category:         "SSS Resources",
			ShortDescription: `Manages a V2 Approval Request resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"sss",
				"approval",
				"request",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "request_id",
					Description: `(Required) tenant_connection_request unique ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Modify status of the approval request.(approved/denied/cancelled) Request’s status can only change from ‘registerd’. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "external_request_id",
					Description: `External reqeust ID. (internal use)`,
				},
				resource.Attribute{
					Name:        "approver_type",
					Description: `Type of appover id. (tenant/tenant_owner/contract/contract_owner/user)`,
				},
				resource.Attribute{
					Name:        "approver_id",
					Description: `Resource ID of approver type. (e.g. tenant id, ecidXXXXXXXXXX, econXXXXXXXXX)`,
				},
				resource.Attribute{
					Name:        "request_user_id",
					Description: `User ID of request’s sender.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service name.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `Action to service.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service name.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region name.`,
				},
				resource.Attribute{
					Name:        "api_path",
					Description: `API Path of the resource to approve.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `Method name.`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `Body of the resource to approve.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `Description list. (contains lang and text)`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `Type of language. (e.g. en, ja)`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Text of description.`,
				},
				resource.Attribute{
					Name:        "request_user",
					Description: `API executing user is sender or not.`,
				},
				resource.Attribute{
					Name:        "approver",
					Description: `API executing user can approve request or not.`,
				},
				resource.Attribute{
					Name:        "approval_deadline",
					Description: `Time of response deadline.`,
				},
				resource.Attribute{
					Name:        "approval_expire",
					Description: `Time of approval expiration.`,
				},
				resource.Attribute{
					Name:        "registered_time",
					Description: `Registered time of approval request.`,
				},
				resource.Attribute{
					Name:        "updated_time",
					Description: `Updated time of approval request’s status.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_request_id",
					Description: `External reqeust ID. (internal use)`,
				},
				resource.Attribute{
					Name:        "approver_type",
					Description: `Type of appover id. (tenant/tenant_owner/contract/contract_owner/user)`,
				},
				resource.Attribute{
					Name:        "approver_id",
					Description: `Resource ID of approver type. (e.g. tenant id, ecidXXXXXXXXXX, econXXXXXXXXX)`,
				},
				resource.Attribute{
					Name:        "request_user_id",
					Description: `User ID of request’s sender.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service name.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `Action to service.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service name.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region name.`,
				},
				resource.Attribute{
					Name:        "api_path",
					Description: `API Path of the resource to approve.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `Method name.`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `Body of the resource to approve.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `Description list. (contains lang and text)`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `Type of language. (e.g. en, ja)`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Text of description.`,
				},
				resource.Attribute{
					Name:        "request_user",
					Description: `API executing user is sender or not.`,
				},
				resource.Attribute{
					Name:        "approver",
					Description: `API executing user can approve request or not.`,
				},
				resource.Attribute{
					Name:        "approval_deadline",
					Description: `Time of response deadline.`,
				},
				resource.Attribute{
					Name:        "approval_expire",
					Description: `Time of approval expiration.`,
				},
				resource.Attribute{
					Name:        "registered_time",
					Description: `Registered time of approval request.`,
				},
				resource.Attribute{
					Name:        "updated_time",
					Description: `Updated time of approval request’s status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_sss_tenant_v1",
			Category:         "SSS Resources",
			ShortDescription: `Manages a V1 tenant resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"sss",
				"tenant",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_name",
					Description: `(Required) Name of new tenant. This name need to be unique globally.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description for this tenant.`,
				},
				resource.Attribute{
					Name:        "tenant_region",
					Description: `(Required) Region this tenant belongs to.`,
				},
				resource.Attribute{
					Name:        "contract_id",
					Description: `(Optional) Contract which new tenant belongs to. If this parameter is not designated, API user's contract implicitly designated. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "contract_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Tenant created time.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `ID of the tenant. ## Import Tenant can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_sss_tenant_v1.tenant_1 <tenant-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "contract_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Tenant created time.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `ID of the tenant. ## Import Tenant can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_sss_tenant_v1.tenant_1 <tenant-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_sss_user_v1",
			Category:         "SSS Resources",
			ShortDescription: `Manages a V1 user resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"sss",
				"user",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "login_id",
					Description: `(Required) Login id of new user.`,
				},
				resource.Attribute{
					Name:        "mail_address",
					Description: `(Required) Mail address of new user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Initial password of new user. If this parameter is not designated, random initial password is generated and applied to new user.`,
				},
				resource.Attribute{
					Name:        "notify_password",
					Description: `(Optional) If this flag is set 'true', notification email will be sent to new user's email address. Even this parameter is optional, you must specify this in case "Creation". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "login_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mail_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `login id of the user. When this contract is tied with icp, this parameter is fixed {email}_{user_id}`,
				},
				resource.Attribute{
					Name:        "contract_owner",
					Description: `If this user is the Super user in this contract, true. If not, false`,
				},
				resource.Attribute{
					Name:        "keystone_name",
					Description: `This user’s API key for keystone authentication`,
				},
				resource.Attribute{
					Name:        "keystone_password",
					Description: `This user’s API secret for keystone authentication`,
				},
				resource.Attribute{
					Name:        "keystone_endpoint",
					Description: `Keystone address this user can use to get token for SSS API request`,
				},
				resource.Attribute{
					Name:        "sss_endpoint",
					Description: `SSS endpoint recommended for this user`,
				},
				resource.Attribute{
					Name:        "contract_id",
					Description: `Contract ID which this user belongs. Contract id format is econ[0-9]{10}`,
				},
				resource.Attribute{
					Name:        "login_integration",
					Description: `If this user's contract is tied with NTT Communications business portal, 'icp' is shown`,
				},
				resource.Attribute{
					Name:        "external_reference_id",
					Description: `External system oriented contract id. If this user's contract is NTT Communications, customer number with 15 numbers will be shown`,
				},
				resource.Attribute{
					Name:        "brand_id",
					Description: `Brand ID which this user belongs. (ex. ecl2)`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Created time of user. ## Import Tenant can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_sss_user_v1.user_1 <user-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "login_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mail_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `login id of the user. When this contract is tied with icp, this parameter is fixed {email}_{user_id}`,
				},
				resource.Attribute{
					Name:        "contract_owner",
					Description: `If this user is the Super user in this contract, true. If not, false`,
				},
				resource.Attribute{
					Name:        "keystone_name",
					Description: `This user’s API key for keystone authentication`,
				},
				resource.Attribute{
					Name:        "keystone_password",
					Description: `This user’s API secret for keystone authentication`,
				},
				resource.Attribute{
					Name:        "keystone_endpoint",
					Description: `Keystone address this user can use to get token for SSS API request`,
				},
				resource.Attribute{
					Name:        "sss_endpoint",
					Description: `SSS endpoint recommended for this user`,
				},
				resource.Attribute{
					Name:        "contract_id",
					Description: `Contract ID which this user belongs. Contract id format is econ[0-9]{10}`,
				},
				resource.Attribute{
					Name:        "login_integration",
					Description: `If this user's contract is tied with NTT Communications business portal, 'icp' is shown`,
				},
				resource.Attribute{
					Name:        "external_reference_id",
					Description: `External system oriented contract id. If this user's contract is NTT Communications, customer number with 15 numbers will be shown`,
				},
				resource.Attribute{
					Name:        "brand_id",
					Description: `Brand ID which this user belongs. (ex. ecl2)`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Created time of user. ## Import Tenant can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_sss_user_v1.user_1 <user-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_storage_virtualstorage_v1",
			Category:         "Storage Resources",
			ShortDescription: `Manages a V1 virtual storage resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"virtualstorage",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Virtual Storage.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of Virtual Storage.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID(UUID) for network to be connected to the Virtual Storage.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) ID(UUID) for subnet to be connected to the Virtual Storage.`,
				},
				resource.Attribute{
					Name:        "volume_type_id",
					Description: `(Optional) ID of volume type used for this Virtual Storage (UUID). User must specify either volume_type_id or volume_type_name. This parameter conflicts with ` + "`" + `volume_type_name` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "volume_type_name",
					Description: `(Optional) Name of volume type used for this Virtual Storage. User must specify either volume_type_id or volume_type_name. This parameter conflicts with ` + "`" + `volume_type_id` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "ip_addr_pool",
					Description: `(Required) IP address pool which specifies IP address range used by the Virtual Storage. The ip_addr_pool structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `(Optional) List of static routes to be set to this Virtual Storage. The host_routes structure is documented below. The ` + "`" + `ip_addr_pool` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) Start IP address of the ip_addr_pool.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) End IP address of the ip_addr_pool. The ` + "`" + `host_routes` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination CIDR of this routing.`,
				},
				resource.Attribute{
					Name:        "nexthop",
					Description: `(Required) Nexthop IP address of this routing. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "volume_type_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `Error message of Virtual Storage. ## Import Virtual Storage can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_storage_virtualstorage_v1.virtual_storage_1 <virtual-storage-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "volume_type_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `Error message of Virtual Storage. ## Import Virtual Storage can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_storage_virtualstorage_v1.virtual_storage_1 <virtual-storage-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_storage_volume_v1",
			Category:         "Storage Resources",
			ShortDescription: `Manages a V1 volume resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"volume",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of volume.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Size of volume in Gigabyte. User can choice following volume sizes depending on storage service type. Block storage service: 100, 250, 500, 1000, 2000, 4000, 8000, 12000 File storage premium service: 256, 512 File storage standard service: 1024, 2048, 3072, 4096, 5120, 10240, 15360, 20480, 25600, 30720, 35840, 40960, 46080, 51200, 56320, 61440, 66560, 71680, 81920, 87040, 92160, 102400`,
				},
				resource.Attribute{
					Name:        "iops_per_gb",
					Description: `(Optional) Provisioned IOPS/GB for volume. User can specify this parameter only in case block storage service.`,
				},
				resource.Attribute{
					Name:        "throughput",
					Description: `(Optional) Throughput for volume. User can specify this parameter only in case file storage standard service.`,
				},
				resource.Attribute{
					Name:        "initiator_iqns",
					Description: `(Optional) List of initiator IQN who can access to this volume. User can specify this parameter only in case block storage service.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone of volume.`,
				},
				resource.Attribute{
					Name:        "virtual_storage_id",
					Description: `(Required) Virtual Storage ID (UUID) which this volume belongs. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `Error message of Volume. ## Import Volume can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_storage_volume_v1.volume_1 <volume-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `Error message of Volume. ## Import Volume can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_storage_volume_v1.volume_1 <volume-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_vna_appliance_v1",
			Category:         "Virtual Network Appliance Resources",
			ShortDescription: `Manages a V1 Virtual Network Appliance resource within Enterprise Cloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"network",
				"appliance",
				"vna",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Virtual Network Appliance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Virtual Network Appliance.`,
				},
				resource.Attribute{
					Name:        "default_gateway",
					Description: `(Optional) IP address of default gateway.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability Zone, this can be referred to using Virtual Server (Nova)'s list availability zones.`,
				},
				resource.Attribute{
					Name:        "virtual_network_appliance_plan_id",
					Description: `(Required) ID of the Virtual Network Appliance Plan.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the Virtual Network Appliance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the interface.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the interface.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The ID of network this interface belongs to.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the interface.`,
				},
				resource.Attribute{
					Name:        "fixed_ips",
					Description: `(Optional) List of fixesIP addresses of interface. Each element of fixed_ips is documented below. The ` + "`" + `interface_[slot number]_fixed_ips` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The IP address assign to interface within subnet. The ` + "`" + `interface_[slot number]_allowed_address_pairs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The IP address of allowed address pairs.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The MAC address of allowed address pairs. In case allowed address pair type is "vrrp", you must specify blank string as mac_address.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of allowed address pairs. You can use ""(blak string) or "vrrp" as this argument.`,
				},
				resource.Attribute{
					Name:        "vrid",
					Description: `(Required) VRID of allowed address pairs. Even though type of this parameter is integer in actual API specification, You need to specify this argument by string, like, "null", "0", "255". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/updatable",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interfaces/[slot number]_fixed_ips/[index]/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interfaces/[slot number]_allowed_address_pairs/[index]/mac_address",
					Description: `See Argument Reference above. ## Import Virtual Network Appliance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_vna_appliance_v1.appliance_1 <appliance-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/updatable",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interfaces/[slot number]_fixed_ips/[index]/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interfaces/[slot number]_allowed_address_pairs/[index]/mac_address",
					Description: `See Argument Reference above. ## Import Virtual Network Appliance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ecl_vna_appliance_v1.appliance_1 <appliance-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ecl_baremetal_keypair_v2":                           0,
		"ecl_baremetal_server_v2":                            1,
		"ecl_compute_instance_v2":                            2,
		"ecl_compute_keypair_v2":                             3,
		"ecl_compute_volume_attach_v2":                       4,
		"ecl_compute_volume_v2":                              5,
		"ecl_dedicated_hypervisor_license_v1":                6,
		"ecl_dedicated_hypervisor_server_v1":                 7,
		"ecl_dns_recordset_v2":                               8,
		"ecl_dns_zone_v2":                                    9,
		"ecl_imagestorages_image_v2":                         10,
		"ecl_imagestorages_member_accepter_v2":               11,
		"ecl_imagestorages_member_v2":                        12,
		"ecl_network_common_function_gateway_v2":             13,
		"ecl_network_gateway_interface_v2":                   14,
		"ecl_network_internet_gateway_v2":                    15,
		"ecl_network_load_balancer_v2":                       16,
		"ecl_network_network_v2":                             17,
		"ecl_network_port_v2":                                18,
		"ecl_network_public_ip_v2":                           19,
		"ecl_network_static_route_v2":                        20,
		"ecl_network_subnet_v2":                              21,
		"provider_connectivity_tenant_connection_request_v2": 22,
		"ecl_provider_connectivity_tenant_connection_v2":     23,
		"ecl_rca_user_v1":                                    24,
		"ecl_security_host_based_v1":                         25,
		"ecl_security_network_based_device_ha_v1":            26,
		"ecl_security_network_based_device_single_v1":        27,
		"ecl_security_network_based_waf_single_v1":           28,
		"ecl_sss_approval_request_v1":                        29,
		"ecl_sss_tenant_v1":                                  30,
		"ecl_sss_user_v1":                                    31,
		"ecl_storage_virtualstorage_v1":                      32,
		"ecl_storage_volume_v1":                              33,
		"ecl_vna_appliance_v1":                               34,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
