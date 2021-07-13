package outscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "access_key",
			Category:         "Resources",
			ShortDescription: `[Manages an access key.]`,
			Description:      ``,
			Keywords: []string{
				"access",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The state for the access key (` + "`" + `ACTIVE` + "`" + ` \| ` + "`" + `INACTIVE` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `Information about the secret access key.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the secret access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the secret access key.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time of the last modification of the secret access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The secret access key that enables you to send requests.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the secret access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not). ## Import An access key can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_access_key.ImportedAccessKey ABCDEFGHIJ0123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_key",
					Description: `Information about the secret access key.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the secret access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the secret access key.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time of the last modification of the secret access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The secret access key that enables you to send requests.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the secret access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not). ## Import An access key can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_access_key.ImportedAccessKey ABCDEFGHIJ0123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "client_gateway",
			Category:         "Resources",
			ShortDescription: `[Manages a client gateway.]`,
			Description:      ``,
			Keywords: []string{
				"client",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `(Required) An Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Required) The communication protocol used to establish tunnel with your client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Required) The public fixed IPv4 address of your client gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "client_gateway",
					Description: `Information about the client gateway.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `An Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of communication tunnel used by the client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IPv4 address of the client gateway (must be a fixed address into a NATed network).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the client gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the client gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A client gateway can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_client_gateway.ImportedClientGateway cgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_gateway",
					Description: `Information about the client gateway.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `An Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of communication tunnel used by the client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IPv4 address of the client gateway (must be a fixed address into a NATed network).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the client gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the client gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A client gateway can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_client_gateway.ImportedClientGateway cgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dhcp_option",
			Category:         "Resources",
			ShortDescription: `[Manages a DHCP option.]`,
			Description:      ``,
			Keywords: []string{
				"dhcp",
				"option",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `(Optional) The IP addresses of domain name servers. If no IP addresses are specified, the ` + "`" + `OutscaleProvidedDNS` + "`" + ` value is set by default.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) Specify a domain name (for example, MyCompany.com). You can specify only one domain name.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `(Optional) The IP addresses of the Network Time Protocol (NTP) servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set",
					Description: `Information about the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If true, the DHCP options set is a default one. If false, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_name",
					Description: `The name of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IP addresses for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IP addresses for the NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import DHCP options can be imported using the DHCP option ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_dhcp_option.ImportedDhcpSet dopt-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_set",
					Description: `Information about the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If true, the DHCP options set is a default one. If false, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_name",
					Description: `The name of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IP addresses for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IP addresses for the NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import DHCP options can be imported using the DHCP option ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_dhcp_option.ImportedDhcpSet dopt-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexible_gpu",
			Category:         "Resources",
			ShortDescription: `[Manages a flexible GPU.]`,
			Description:      ``,
			Keywords: []string{
				"flexible",
				"gpu",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `(Optional) The processor generation that the fGPU must be compatible with. If not specified, the oldest possible processor generation is selected (as provided by [ReadFlexibleGpuCatalog](https://docs.outscale.com/api#readflexiblegpucatalog) for the specified model of fGPU).`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `(Required) The model of fGPU you want to allocate. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `(Required) The Subregion in which you want to create the fGPU. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "flexible_gpu",
					Description: `Information about the flexible GPU (fGPU).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `The ID of the fGPU.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The compatible processor generation.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the fGPU (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion where the fGPU is located.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the fGPU is attached to, if any. ## Import A flexible GPU can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import outscale_flexible_gpu.imported_fgpu fgpu-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "flexible_gpu",
					Description: `Information about the flexible GPU (fGPU).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `The ID of the fGPU.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The compatible processor generation.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the fGPU (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion where the fGPU is located.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the fGPU is attached to, if any. ## Import A flexible GPU can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import outscale_flexible_gpu.imported_fgpu fgpu-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexible_gpu_link",
			Category:         "Resources",
			ShortDescription: `[Manages a flexible GPU link.]`,
			Description:      ``,
			Keywords: []string{
				"flexible",
				"gpu",
				"link",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `(Required) The ID of the fGPU you want to attach.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Required) The ID of the VM you want to attach the fGPU to. ## Attribute Reference No attribute is exported. ## Import A flexible GPU link can be imported using the flexible GPU ID. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import outscale_flexible_gpu_link.imported_link_fgpu fgpu-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "image",
			Category:         "Resources",
			ShortDescription: `[Manages an image.]`,
			Description:      ``,
			Keywords: []string{
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "architecture",
					Description: `(Optional) The architecture of the OMI (by default, ` + "`" + `i386` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `(Optional) One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Optional) The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) The name of the device.`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `(Optional) The name of the virtual device (ephemeralN).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the new OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `(Optional) The pre-signed URL of the OMI manifest file, or the full path to the OMI stored in a bucket. If you specify this parameter, a copy of the OMI is created in your account.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional) A unique name for the new OMI.<br /> Constraints: 3-128 alphanumeric characters, underscores (_), spaces ( ), parentheses (()), slashes (/), periods (.), or dashes (-).`,
				},
				resource.Attribute{
					Name:        "no_reboot",
					Description: `(Optional) If false, the VM shuts down before creating the OMI and then reboots. If true, the VM does not.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `(Optional) The name of the root device.`,
				},
				resource.Attribute{
					Name:        "source_image_id",
					Description: `(Optional) The ID of the OMI you want to copy.`,
				},
				resource.Attribute{
					Name:        "source_region_name",
					Description: `(Optional) The name of the source Region, which must be the same as the Region of your account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Optional) The ID of the VM from which you want to create the OMI. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `Information about the OMI.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the OMI (by default, ` + "`" + `i386` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (ephemeralN).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time at which the OMI was created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location of the bucket where the OMI files are stored.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of the OMI.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the OMI (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_comment",
					Description: `Information about the change of state.`,
				},
				resource.Attribute{
					Name:        "state_code",
					Description: `The code of the change of state.`,
				},
				resource.Attribute{
					Name:        "state_message",
					Description: `A message explaining the change of state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the OMI.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import An image can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_image.ImportedImage ami-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "image",
					Description: `Information about the OMI.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the OMI (by default, ` + "`" + `i386` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (ephemeralN).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time at which the OMI was created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location of the bucket where the OMI files are stored.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of the OMI.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the OMI (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_comment",
					Description: `Information about the change of state.`,
				},
				resource.Attribute{
					Name:        "state_code",
					Description: `The code of the change of state.`,
				},
				resource.Attribute{
					Name:        "state_message",
					Description: `A message explaining the change of state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the OMI.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import An image can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_image.ImportedImage ami-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "image_launch_permission",
			Category:         "Resources",
			ShortDescription: `[Manages an image launch permission.]`,
			Description:      ``,
			Keywords: []string{
				"image",
				"launch",
				"permission",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) The ID of the OMI you want to modify.`,
				},
				resource.Attribute{
					Name:        "permission_additions",
					Description: `(Optional) Information about the users you want to give permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `(Optional) If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account ID of one or more users you want to give permissions to.`,
				},
				resource.Attribute{
					Name:        "permission_removals",
					Description: `(Optional) Information about the users you want to remove permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `(Optional) If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account ID of one or more users you want to remove permissions from. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI you want to modify.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Information about the permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "accounts_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `A description of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI you want to modify.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Information about the permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "accounts_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "internet_service",
			Category:         "Resources",
			ShortDescription: `[Manages an Internet service.]`,
			Description:      ``,
			Keywords: []string{
				"internet",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Information about the Internet service.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net attached to the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Net to the Internet service (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import An internet service can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_internet_service.ImportedInternetService igw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "internet_service",
					Description: `Information about the Internet service.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net attached to the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Net to the Internet service (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import An internet service can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_internet_service.ImportedInternetService igw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "internet_service_link",
			Category:         "Resources",
			ShortDescription: `[Manages an Internet service link.]`,
			Description:      ``,
			Keywords: []string{
				"internet",
				"service",
				"link",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `(Required) The ID of the Internet service you want to attach.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net to which you want to attach the Internet service. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service you want to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Net to the Internet service (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service. ## Import An internet service link can be imported using the internet service ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_internet_service_link.ImportedInternetServiceLink igw-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service you want to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Net to the Internet service (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service. ## Import An internet service link can be imported using the internet service ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_internet_service_link.ImportedInternetServiceLink igw-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "keypair",
			Category:         "Resources",
			ShortDescription: `[Manages a keypair.]`,
			Description:      ``,
			Keywords: []string{
				"keypair",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "keypair_name",
					Description: `(Required) A unique name for the keypair, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters).`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The public key. It must be base64-encoded. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "keypair",
					Description: `Information about the created keypair.`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `The MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key. ## Import A keypair can be imported using its name. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_keypair.ImportedKeypair Name-of-the-Keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keypair",
					Description: `Information about the created keypair.`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `The MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key. ## Import A keypair can be imported using its name. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_keypair.ImportedKeypair Name-of-the-Keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "load_balancer",
			Category:         "Resources",
			ShortDescription: `[Manages a load balancer.]`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listeners",
					Description: `(Required) One or more listeners to create.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `(Optional) The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `(Optional) The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `(Optional) The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `(Optional) The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `(Optional) The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Required) The unique name of the load balancer (32 alphanumeric or hyphen characters maximum, but cannot start or end with a hyphen).`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `(Optional) The type of load balancer: ` + "`" + `internet-facing` + "`" + ` or ` + "`" + `internal` + "`" + `. Use this parameter only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) (Net only) One or more IDs of security groups you want to assign to the load balancer. If not specified, the default security group of the Net is assigned to the load balancer.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) One or more IDs of Subnets in your Net that you want to attach to the load balancer.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) One or more names of Subregions (currently, only one Subregion is supported). This parameter is not required if you create a load balancer in a Net. To create an internal load balancer, use the ` + "`" + `load_balancer_type` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) One or more tags assigned to the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `Information about the load balancer.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either 5 or 60 (by default, 60).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between 1 and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP address.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP address.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The IDs of the Subnets for the load balancer.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `One or more names of Subregions for the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A load balancer can be imported using its name. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_load_balancer.ImportedLbu Name-of-the-Lbu ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer",
					Description: `Information about the load balancer.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either 5 or 60 (by default, 60).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between 1 and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP address.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP address.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The IDs of the Subnets for the load balancer.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `One or more names of Subregions for the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A load balancer can be imported using its name. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_load_balancer.ImportedLbu Name-of-the-Lbu ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "load_balancer_attributes",
			Category:         "Resources",
			ShortDescription: `[Manages load balancer attributes.]`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"attributes",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `(Optional) The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `(Optional) The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `(Optional) The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either 5 or 60 (by default, 60).`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `(Optional) The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Required) The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `(Optional) The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included). This parameter is required if you want to update the server certificate.`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `(Optional) The name of the policy you want to enable for the listener.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `(Optional) The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat). If this parameter is specified, you must also specify the ` + "`" + `load_balancer_port` + "`" + ` parameter. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `Information about the load balancer.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either 5 or 60 (by default, 60).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between 1 and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP address.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP address.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The IDs of the Subnets for the load balancer.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `One or more names of Subregions for the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer",
					Description: `Information about the load balancer.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either 5 or 60 (by default, 60).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between 1 and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP address.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP address.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The IDs of the Subnets for the load balancer.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `One or more names of Subregions for the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "load_balancer_listener_rule",
			Category:         "Resources",
			ShortDescription: `[Manages a load balancer listener rule.]`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"listener",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_rule",
					Description: `Information about the listener rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `(Optional) A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `(Optional) A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `(Optional) A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "listener",
					Description: `Information about the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Optional) The name of the load balancer to which the listener is attached.`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `(Optional) The port of load balancer on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Required) The IDs of the backend VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_rule",
					Description: `Information about the listener rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_rule_id",
					Description: `The ID of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `The IDs of the backend VMs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_rule",
					Description: `Information about the listener rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_rule_id",
					Description: `The ID of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `The IDs of the backend VMs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "load_balancer_policy",
			Category:         "Resources",
			ShortDescription: `[Manages a load balancer policy.]`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Optional) The name of the application cookie used for stickiness. This parameter is required if you create a stickiness policy based on an application-generated cookie.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Required) The name of the load balancer for which you want to create a policy.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) The name of the policy. This name must be unique and consist of alphanumeric characters and dashes (-).`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Required) The type of stickiness policy you want to create: ` + "`" + `app` + "`" + ` or ` + "`" + `load_balancer` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `Information about the load balancer.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either 5 or 60 (by default, 60).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between 1 and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP address.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP address.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The IDs of the Subnets for the load balancer.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `One or more names of Subregions for the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer",
					Description: `Information about the load balancer.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either 5 or 60 (by default, 60).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between 1 and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP address.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP address.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The IDs of the Subnets for the load balancer.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `One or more names of Subregions for the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "load_balancer_vms",
			Category:         "Resources",
			ShortDescription: `[Manages load balancer VMs.]`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"vms",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `(Required) One or more IDs of back-end VMs.<br /> Specifying the same ID several times has no effect as each back-end VM has equal weight.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Required) The name of the load balancer. ## Attribute Reference No attribute is exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nat_service",
			Category:         "Resources",
			ShortDescription: `[Manages a NAT service.]`,
			Description:      ``,
			Keywords: []string{
				"nat",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `(Required) The allocation ID of the EIP to associate with the NAT service.<br /> If the EIP is already associated with another resource, you must first disassociate it.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the Subnet in which you want to create the NAT service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nat_service",
					Description: `Information about the NAT service.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of the NAT service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about the External IP address or addresses (EIPs) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NAT service (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A NAT service can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_nat_service.ImportedNatService nat-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nat_service",
					Description: `Information about the NAT service.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of the NAT service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about the External IP address or addresses (EIPs) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NAT service (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A NAT service can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_nat_service.ImportedNatService nat-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net",
			Category:         "Resources",
			ShortDescription: `[Manages a Net.]`,
			Description:      ``,
			Keywords: []string{
				"net",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Required) The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `(Optional) The tenancy options for the VMs (` + "`" + `default` + "`" + ` if a VM created in a Net can be launched with any tenancy, ` + "`" + `dedicated` + "`" + ` if it can be launched with dedicated tenancy VMs running on single-tenant hardware). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net",
					Description: `Information about the Net.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net. ## Import A Net can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_net.ImportedNet vpc-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net",
					Description: `Information about the Net.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net. ## Import A Net can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_net.ImportedNet vpc-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net_access_point",
			Category:         "Resources",
			ShortDescription: `[Manages a Net access point.]`,
			Description:      ``,
			Keywords: []string{
				"net",
				"access",
				"point",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `(Optional) One or more IDs of route tables to use for the connection.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The name of the service (in the format ` + "`" + `com.outscale.region.service` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_access_point",
					Description: `Information about the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `The ID of the route tables associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net access point (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A Net access point can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_net_access_point.ImportedNetAccessPoint vpce-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_access_point",
					Description: `Information about the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `The ID of the route tables associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net access point (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A Net access point can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_net_access_point.ImportedNetAccessPoint vpce-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net_attributes",
			Category:         "Resources",
			ShortDescription: `[Manages a Net attribute.]`,
			Description:      ``,
			Keywords: []string{
				"net",
				"attributes",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `(Required) The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net. ## Import A Net attribute can be imported using the Net ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_net_attributes.ImportedNet vpc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net. ## Import A Net attribute can be imported using the Net ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_net_attributes.ImportedNet vpc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net_peering",
			Category:         "Resources",
			ShortDescription: `[Manages a Net peering.]`,
			Description:      ``,
			Keywords: []string{
				"net",
				"peering",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accepter_net_id",
					Description: `(Required) The ID of the Net you want to connect with.`,
				},
				resource.Attribute{
					Name:        "source_net_id",
					Description: `(Required) The ID of the Net you send the peering request from.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_peering",
					Description: `Information about the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering connection (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A Net peering can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_net_peering.ImportedNetPeering pcx-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_peering",
					Description: `Information about the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering connection (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A Net peering can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_net_peering.ImportedNetPeering pcx-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net_peering_acceptation",
			Category:         "Resources",
			ShortDescription: `[Manages a Net peering acceptation.]`,
			Description:      ``,
			Keywords: []string{
				"net",
				"peering",
				"acceptation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `(Required) The ID of the Net peering connection you want to accept. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_peering",
					Description: `Information about the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering connection (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_peering",
					Description: `Information about the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering connection (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nic",
			Category:         "Resources",
			ShortDescription: `[Manages a network interface card (NIC).]`,
			Description:      ``,
			Keywords: []string{
				"nic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the NIC.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) The primary private IP address for the NIC.<br /> This IP address must be within the IP address range of the Subnet that you specify with the ` + "`" + `subnet_id` + "`" + ` attribute.<br /> If you do not specify this attribute, a random private IP address is selected within the IP address range of the Subnet.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `(Optional) If true, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) The private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the Subnet in which you want to create the NIC.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nic",
					Description: `Information about the NIC.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the NIC is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A NIC can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_nic.ImportedNic eni-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nic",
					Description: `Information about the NIC.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the NIC is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A NIC can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_nic.ImportedNic eni-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nic_link",
			Category:         "Resources",
			ShortDescription: `[Manages a NIC link.]`,
			Description:      ``,
			Keywords: []string{
				"nic",
				"link",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_number",
					Description: `(Required) The index of the VM device for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Required) The ID of the NIC you want to attach.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Required) The ID of the VM to which you want to attach the NIC. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC attachment. ## Import A NIC link can be imported using the NIC ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_nic_link.ImportedNicLink eni-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC attachment. ## Import A NIC link can be imported using the NIC ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_nic_link.ImportedNicLink eni-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nic_private_ip",
			Category:         "Resources",
			ShortDescription: `[Manages a NIC private IP.]`,
			Description:      ``,
			Keywords: []string{
				"nic",
				"private",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_relink",
					Description: `(Optional) If true, allows an IP address that is already assigned to another NIC in the same Subnet to be assigned to the NIC you specified.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Required) The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) The secondary private IP address or addresses you want to assign to the NIC within the IP address range of the Subnet.`,
				},
				resource.Attribute{
					Name:        "secondary_private_ip_count",
					Description: `(Optional) The number of secondary private IP addresses to assign to the NIC. ## Attribute Reference No attribute is exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "public_ip",
			Category:         "Resources",
			ShortDescription: `[Manages a public IP.]`,
			Description:      ``,
			Keywords: []string{
				"public",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Information about the public IP.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the EIP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the External IP (EIP) is associated with (if any). ## Import A public IP can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_public_ip.ImportedPublicIp eipalloc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_ip",
					Description: `Information about the public IP.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the EIP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the External IP (EIP) is associated with (if any). ## Import A public IP can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_public_ip.ImportedPublicIp eipalloc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "public_ip_link",
			Category:         "Resources",
			ShortDescription: `[Manages a public IP link.]`,
			Description:      ``,
			Keywords: []string{
				"public",
				"ip",
				"link",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_relink",
					Description: `(Optional) If true, allows the EIP to be associated with the VM or NIC that you specify even if it is already associated with another VM or NIC.<br /> If false, prevents the EIP from being associated with the VM or NIC that you specify if it is already associated with another VM or NIC.<br /> (By default, true in the public Cloud, false in a Net.)`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Optional) (Net only) The ID of the NIC. This parameter is required if the VM has more than one NIC attached. Otherwise, you need to specify the ` + "`" + `vm_id` + "`" + ` parameter instead. You cannot specify both parameters at the same time.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) (Net only) The primary or secondary private IP address of the specified NIC. By default, the primary private IP address.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `(Optional) The allocation ID of the EIP. In a Net, this parameter is required.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) The EIP. In the public Cloud, this parameter is required.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Optional) The ID of the VM.<br /> In the public Cloud, this parameter is required.<br /> In a Net, this parameter is required if the VM has only one NIC. Otherwise, you need to specify the ` + "`" + `nic_id` + "`" + ` parameter instead. You cannot specify both parameters at the same time. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Net only) The ID representing the association of the EIP with the VM or the NIC. ## Import A public IP link can be imported using the public IP or the public IP link ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_public_ip_link.ImportedPublicIpLink eipassoc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Net only) The ID representing the association of the EIP with the VM or the NIC. ## Import A public IP link can be imported using the public IP or the public IP link ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_public_ip_link.ImportedPublicIpLink eipassoc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "route",
			Category:         "Resources",
			ShortDescription: `[Manages a route.]`,
			Description:      ``,
			Keywords: []string{
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `(Required) The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `(Optional) The ID of an Internet service or virtual gateway attached to your Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `(Optional) The ID of a NAT service.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `(Optional) The ID of a Net peering connection.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Optional) The ID of a NIC.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required) The ID of the route table for which you want to create a route.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Optional) The ID of a NAT VM in your Net (attached to exactly one NIC). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `Information about the route table.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A route can be imported using the route table ID and the destination IP range. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_routeImportedRoute rtb-12345678_10.0.0.0/0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table",
					Description: `Information about the route table.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A route can be imported using the route table ID and the destination IP range. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_routeImportedRoute rtb-12345678_10.0.0.0/0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "route_table",
			Category:         "Resources",
			ShortDescription: `[Manages a route table.]`,
			Description:      ``,
			Keywords: []string{
				"route",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net for which you want to create a route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `Information about the route table.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A route table can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_route_table.ImportedRouteTable rtb-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table",
					Description: `Information about the route table.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A route table can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_route_table.ImportedRouteTable rtb-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "route_table_link",
			Category:         "Resources",
			ShortDescription: `[Manages a route table link.]`,
			Description:      ``,
			Keywords: []string{
				"route",
				"table",
				"link",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required) The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the Subnet. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet. ## Import A route table link can be imported using the route table ID and the route table link ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_route_table_link.ImportedRouteTableLink rtb-12345678_rtbassoc-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet. ## Import A route table link can be imported using the route table ID and the route table link ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_route_table_link.ImportedRouteTableLink rtb-12345678_rtbassoc-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "security_group",
			Category:         "Resources",
			ShortDescription: `[Manages a security group.]`,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the security group, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `(Optional) The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `(Required) The name of the security group.<br /> This name must not start with ` + "`" + `sg-` + "`" + `.</br> This name must be unique and contain between 1 and 255 ASCII characters. Accented letters are not allowed.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `Information about the security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A security group can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_security_group.ImportedSecurityGroup sg-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group",
					Description: `Information about the security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A security group can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_security_group.ImportedSecurityGroup sg-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "security_group_rule",
			Category:         "Resources",
			ShortDescription: `[Manages a security group rule.]`,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "flow",
					Description: `(Required) The direction of the flow: ` + "`" + `Inbound` + "`" + ` or ` + "`" + `Outbound` + "`" + `. You can specify ` + "`" + `Outbound` + "`" + ` for Nets only.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `(Optional) The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Optional) The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional) Information about the security group rule to create.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `(Optional) The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `(Optional) Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `(Optional) The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `(Optional) One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `(Optional) The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id_to_link",
					Description: `(Optional) The account ID of the owner of the security group for which you want to create a rule.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The ID of the security group for which you want to create a rule.`,
				},
				resource.Attribute{
					Name:        "security_group_name_to_link",
					Description: `(Optional) The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group.`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `(Optional) The end of the port range for the TCP and UDP protocols, or an ICMP type number. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `Information about the security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group",
					Description: `Information about the security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "snapshot",
			Category:         "Resources",
			ShortDescription: `[Manages a snapshot.]`,
			Description:      ``,
			Keywords: []string{
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the snapshot.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `(Optional) The pre-signed URL of the snapshot you want to import from the bucket.`,
				},
				resource.Attribute{
					Name:        "snapshot_size",
					Description: `(Optional) The size of the snapshot you want to create in your account, in bytes. This size must be greater than or equal to the size of the original, uncompressed snapshot.`,
				},
				resource.Attribute{
					Name:        "source_region_name",
					Description: `(Optional) The name of the source Region, which must be the same as the Region of your account.`,
				},
				resource.Attribute{
					Name:        "source_snapshot_id",
					Description: `(Optional) The ID of the snapshot you want to copy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Optional) The ID of the volume you want to create a snapshot of. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `Information about the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot",
					Description: `Information about the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "snapshot_attributes",
			Category:         "Resources",
			ShortDescription: `[Manages snapshot attributes.]`,
			Description:      ``,
			Keywords: []string{
				"snapshot",
				"attributes",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "permissions_to_create_volume_additions",
					Description: `(Optional) Information about the users you want to give permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account ID of one or more users you want to give permissions to.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `(Optional) If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume_removals",
					Description: `(Optional) Information about the users you want to remove permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account ID of one or more users you want to remove permissions from.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `(Optional) If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Required) The ID of the snapshot. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume_additions",
					Description: `Information about the permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "permissions_to_create_volume_additions",
					Description: `Information about the permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "subnet",
			Category:         "Resources",
			ShortDescription: `[Manages a Subnet.]`,
			Description:      ``,
			Keywords: []string{
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Required) The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net for which you want to create a Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `(Optional) The name of the Subregion in which you want to create the Subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Information about the Subnet.`,
				},
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IP addresses in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If true, a public IP address is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion in which the Subnet is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Subnet.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A subnet can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_subnet.ImportedSubnet subnet-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subnet",
					Description: `Information about the Subnet.`,
				},
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IP addresses in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If true, a public IP address is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion in which the Subnet is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Subnet.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A subnet can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_subnet.ImportedSubnet subnet-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "virtual_gateway",
			Category:         "Resources",
			ShortDescription: `[Manages a virtual gateway.]`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Required) The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "virtual_gateway",
					Description: `Information about the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "net_to_virtual_gateway_links",
					Description: `The Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the virtual gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway. ## Import A virtual gateway can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_virtual_gateway.ImportedVirtualGateway vgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_gateway",
					Description: `Information about the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "net_to_virtual_gateway_links",
					Description: `The Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the virtual gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway. ## Import A virtual gateway can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_virtual_gateway.ImportedVirtualGateway vgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "virtual_gateway_link",
			Category:         "Resources",
			ShortDescription: `[Manages a virtual gateway link.]`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"gateway",
				"link",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net to which you want to attach the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `(Required) The ID of the virtual gateway. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_to_virtual_gateway_link",
					Description: `Information about the attachment.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `). ## Import A virtual gateway link can be imported using its virtual gateway ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_virtual_gateway_link.ImportedVirtualGatewayLink vgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_to_virtual_gateway_link",
					Description: `Information about the attachment.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `). ## Import A virtual gateway link can be imported using its virtual gateway ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_virtual_gateway_link.ImportedVirtualGatewayLink vgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "virtual_gateway_route_propagation",
			Category:         "Resources",
			ShortDescription: `[Manages a virtual gateway route propagation.]`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"gateway",
				"route",
				"propagation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable",
					Description: `(Required) If true, a virtual gateway can propagate routes to a specified route table of a Net. If false, the propagation is disabled.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required) The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `(Required) The ID of the virtual gateway. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `Information about the route table.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table",
					Description: `Information about the route table.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vm",
			Category:         "Resources",
			ShortDescription: `[Manages a virtual machine (VM).]`,
			Description:      ``,
			Keywords: []string{
				"vm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `(Optional) One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Optional) The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) The name of the device.`,
				},
				resource.Attribute{
					Name:        "no_device",
					Description: `(Optional) Removes the device which is included in the block device mapping of the OMI.`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `(Optional) The name of the virtual device (ephemeralN).`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `(Optional) If true, the VM is created with optimized BSU I/O. Updating this parameter will trigger a stop/start of the VM.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `(Optional) A unique identifier which enables you to manage the idempotency.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `(Optional) If true, you cannot terminate the VM using Cockpit, the CLI or the API. If false, you can.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) The ID of the OMI used to create the VM. You can find the list of OMIs by calling the [ReadImages](https://docs.outscale.com/api#readimages) method.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `(Optional) The name of the keypair.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Optional) One or more NICs. If you specify this parameter, you must define one NIC as the primary network interface of the VM with ` + "`" + `0` + "`" + ` as its device number.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) By default or if set to true, the NIC is deleted when the VM is terminated. You can specify this parameter only for a new NIC. To modify this value for an existing NIC, see [UpdateNic](https://docs.outscale.com/api#updatenic).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the NIC, if you are creating a NIC when creating the VM.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `(Optional) The index of the VM device for the NIC attachment (between 0 and 7, both included). This parameter is required if you create a NIC when creating the VM.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Optional) The ID of the NIC, if you are attaching an existing NIC when creating a VM.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) One or more private IP addresses to assign to the NIC, if you create a NIC when creating a VM. Only one private IP address can be the primary private IP address.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `(Optional) If true, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) The private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "secondary_private_ip_count",
					Description: `(Optional) The number of secondary private IP addresses, if you create a NIC when creating a VM. This parameter cannot be specified if you specified more than one private IP address in the ` + "`" + `private_ips` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) One or more IDs of security groups for the NIC, if you acreate a NIC when creating a VM.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of the Subnet for the NIC, if you create a NIC when creating a VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `(Optional) The performance of the VM (` + "`" + `medium` + "`" + ` \| ` + "`" + `high` + "`" + ` \| ` + "`" + `highest` + "`" + `). Updating this parameter will trigger a stop/start of the VM.`,
				},
				resource.Attribute{
					Name:        "placement_subregion_name",
					Description: `(Optional) The name of the Subregion where the VM is placed.`,
				},
				resource.Attribute{
					Name:        "placement_tenancy",
					Description: `(Optional) The tenancy of the VM (` + "`" + `default` + "`" + ` | ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) One or more private IP addresses of the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) One or more IDs of security group for the VMs.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `(Optional) One or more names of security groups for the VMs.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of the Subnet in which you want to create the VM.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Data or script used to add a specific configuration to the VM. It must be base64-encoded, either directly or using the [base64encode](https://www.terraform.io/docs/configuration/functions/base64encode.html) Terraform function. For multiline strings, use a [heredoc syntax](https://www.terraform.io/docs/configuration/expressions.html#string-literals). Updating this parameter will trigger a stop/start of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `(Optional) The VM behavior when you stop it. By default or if set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is terminated.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `(Optional) The type of VM (` + "`" + `t2.small` + "`" + ` by default). Updating this parameter will trigger a stop/start of the VM.<br /> For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `Information about one or more created VMs.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Windows VM only) The administrator password of the VM. This password is encrypted with the keypair you specified when launching the VM and encoded in Base64. You need to wait about 10 minutes after launching the VM to be able to retrieve this password.<br /> If ` + "`" + `get_admin_password` + "`" + ` is false or not specified, the VM resource is created without the ` + "`" + `admin_password` + "`" + ` attribute. Once ` + "`" + `admin_password` + "`" + ` is available, it will appear in the Terraform state after the next`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the VM (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings_created",
					Description: `The block device mapping of the VM.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the created BSU volume.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "link_date",
					Description: `The time and date of attachment of the volume to the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `If true, the VM is optimized for BSU I/O.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If true, you cannot terminate the VM using Cockpit, the CLI or the API. If false, you can.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the VMs (` + "`" + `ovm` + "`" + ` \| ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI used to create the VM.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, 0, 1, 2, and so on).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the VM is running.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Net only) The network interface cards (NICs) the VMs are attached to.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP address or addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the NIC.`,
				},
				resource.Attribute{
					Name:        "os_family",
					Description: `Indicates the operating system (OS) of the VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `The performance of the VM (` + "`" + `medium` + "`" + ` \| ` + "`" + `high` + "`" + ` \| ` + "`" + `highest` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `Information about the placement of the VM.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` \| ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI used to create the VM (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, /dev/vda1).`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the VM (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more security groups associated with the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the VM.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VM.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded MIME user data.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `The VM behavior when you stop it. By default or if set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types). ## Import A VM can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_vm.ImportedVm i-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vms",
					Description: `Information about one or more created VMs.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Windows VM only) The administrator password of the VM. This password is encrypted with the keypair you specified when launching the VM and encoded in Base64. You need to wait about 10 minutes after launching the VM to be able to retrieve this password.<br /> If ` + "`" + `get_admin_password` + "`" + ` is false or not specified, the VM resource is created without the ` + "`" + `admin_password` + "`" + ` attribute. Once ` + "`" + `admin_password` + "`" + ` is available, it will appear in the Terraform state after the next`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the VM (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings_created",
					Description: `The block device mapping of the VM.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the created BSU volume.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "link_date",
					Description: `The time and date of attachment of the volume to the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `If true, the VM is optimized for BSU I/O.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If true, you cannot terminate the VM using Cockpit, the CLI or the API. If false, you can.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the VMs (` + "`" + `ovm` + "`" + ` \| ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI used to create the VM.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, 0, 1, 2, and so on).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the VM is running.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Net only) The network interface cards (NICs) the VMs are attached to.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP address or addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the NIC.`,
				},
				resource.Attribute{
					Name:        "os_family",
					Description: `Indicates the operating system (OS) of the VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `The performance of the VM (` + "`" + `medium` + "`" + ` \| ` + "`" + `high` + "`" + ` \| ` + "`" + `highest` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `Information about the placement of the VM.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` \| ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI used to create the VM (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, /dev/vda1).`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the VM (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more security groups associated with the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the VM.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VM.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded MIME user data.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `The VM behavior when you stop it. By default or if set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types). ## Import A VM can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_vm.ImportedVm i-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "volume",
			Category:         "Resources",
			ShortDescription: `[Manages a volume.]`,
			Description:      ``,
			Keywords: []string{
				"volume",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the volume, in gibibytes (GiB). The maximum allowed size for a volume is 14901 GiB. This parameter is required if the volume is not created from a snapshot (` + "`" + `snapshot_id` + "`" + ` unspecified).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of the snapshot from which you want to create the volume.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `(Required) The Subregion in which you want to create the volume.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of volume you want to create (` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `standard` + "`" + `). If not specified, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `Information about the volume.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS): For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS. For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot from which the volume was created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the volume was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the volume.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `). ## Import A volume can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_volume.ImportedVolume vol-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "volume",
					Description: `Information about the volume.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS): For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS. For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot from which the volume was created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the volume was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the volume.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `). ## Import A volume can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_volume.ImportedVolume vol-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vpn_connection",
			Category:         "Resources",
			ShortDescription: `[Manages a VPN connection.]`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `(Required) The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Required) The type of VPN connection (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `(Optional) If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags to add to this resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `(Required) The ID of the virtual gateway. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpn_connection",
					Description: `Information about a VPN connection.`,
				},
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `Example configuration for the client gateway.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway used on the client end of the connection.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection (always ` + "`" + `ipsec.1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Information about one or more static routes associated with the VPN connection, if any.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `The type of route (always ` + "`" + `static` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the static route (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VPN connection (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VPN connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection. ## Import A VPN connection can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_vpn_connection.ImportedVPN vpn-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpn_connection",
					Description: `Information about a VPN connection.`,
				},
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `Example configuration for the client gateway.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway used on the client end of the connection.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection (always ` + "`" + `ipsec.1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Information about one or more static routes associated with the VPN connection, if any.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `The type of route (always ` + "`" + `static` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the static route (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VPN connection (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VPN connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection. ## Import A VPN connection can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_vpn_connection.ImportedVPN vpn-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vpn_connection_route",
			Category:         "Resources",
			ShortDescription: `[Manages a VPN connection route.]`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"connection",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `(Required) The network prefix of the route, in CIDR notation (for example, 10.12.0.0/16).`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `(Required) The ID of the target VPN connection of the static route. ## Attribute Reference No attribute is exported. ## Import A VPN connection route can be imported using the VPN connection ID and the route destination IP range. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import outscale_vpn_connection_route.ImportedRoute vpn-12345678_10.0.0.0/0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"access_key":                        0,
		"client_gateway":                    1,
		"dhcp_option":                       2,
		"flexible_gpu":                      3,
		"flexible_gpu_link":                 4,
		"image":                             5,
		"image_launch_permission":           6,
		"internet_service":                  7,
		"internet_service_link":             8,
		"keypair":                           9,
		"load_balancer":                     10,
		"load_balancer_attributes":          11,
		"load_balancer_listener_rule":       12,
		"load_balancer_policy":              13,
		"load_balancer_vms":                 14,
		"nat_service":                       15,
		"net":                               16,
		"net_access_point":                  17,
		"net_attributes":                    18,
		"net_peering":                       19,
		"net_peering_acceptation":           20,
		"nic":                               21,
		"nic_link":                          22,
		"nic_private_ip":                    23,
		"public_ip":                         24,
		"public_ip_link":                    25,
		"route":                             26,
		"route_table":                       27,
		"route_table_link":                  28,
		"security_group":                    29,
		"security_group_rule":               30,
		"snapshot":                          31,
		"snapshot_attributes":               32,
		"subnet":                            33,
		"virtual_gateway":                   34,
		"virtual_gateway_link":              35,
		"virtual_gateway_route_propagation": 36,
		"vm":                                37,
		"volume":                            38,
		"vpn_connection":                    39,
		"vpn_connection_route":              40,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
