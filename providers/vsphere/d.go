package vsphere

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vsphere_content_library_item",
			Category:         "Data Sources",
			ShortDescription: `Provides a vSphere cluster data source. This can be used to get the general attributes of a vSphere cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name or absolute path to the cluster.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of the datacenter the cluster is located in. This can be omitted if the search path used in ` + "`" + `name` + "`" + ` is an absolute path. For default datacenters, use the id attribute from an empty ` + "`" + `vsphere_datacenter` + "`" + ` data source. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ## Attribute Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_custom_attribute",
			Category:         "Data Sources",
			ShortDescription: `Provides a vSphere custom attribute data source. This can be used to reference custom attributes not managed in Terraform.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the custom attribute. ## Attribute Reference In addition to the ` + "`" + `id` + "`" + ` being exported, all of the fields that are available in the [` + "`" + `vsphere_custom_attribute` + "`" + ` resource][resource-custom-attribute] are also populated. See that page for further details.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_datacenter",
			Category:         "Data Sources",
			ShortDescription: `A data source that can be used to get the ID of a datacenter.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the datacenter. This can be a name or path. Can be omitted if there is only one datacenter in your inventory. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_datastore",
			Category:         "Data Sources",
			ShortDescription: `Provides a vSphere datastore data source. This can be used to get the general attributes of a vSphere datastore.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the datastore. This can be a name or path.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of the datacenter the datastore is located in. This can be omitted if the search path used in ` + "`" + `name` + "`" + ` is an absolute path. For default datacenters, use the id attribute from an empty ` + "`" + `vsphere_datacenter` + "`" + ` data source. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ## Attribute Reference Currently, the only exported attribute from this data source is ` + "`" + `id` + "`" + `, which represents the ID of the datastore that was looked up.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_datastore_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides a vSphere datastore cluster data source. This can be used to get the general attributes of a vSphere datastore cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name or absolute path to the datastore cluster.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of the datacenter the datastore cluster is located in. This can be omitted if the search path used in ` + "`" + `name` + "`" + ` is an absolute path. For default datacenters, use the id attribute from an empty ` + "`" + `vsphere_datacenter` + "`" + ` data source. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ## Attribute Reference Currently, the only exported attribute from this data source is ` + "`" + `id` + "`" + `, which represents the ID of the datastore cluster that was looked up.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_distributed_virtual_switch",
			Category:         "Data Sources",
			ShortDescription: `Provides a vSphere distributed virtual switch data source. This can be used to get select attributes of a DVS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the distributed virtual switch. This can be a name or path.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of the datacenter the DVS is located in. This can be omitted if the search path used in ` + "`" + `name` + "`" + ` is an absolute path. For default datacenters, use the id attribute from an empty ` + "`" + `vsphere_datacenter` + "`" + ` data source. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ## Attribute Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_host",
			Category:         "Data Sources",
			ShortDescription: `A data source that can be used to get the ID of a host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of a datacenter.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the host. This can be a name or path. Can be omitted if there is only one host in your inventory. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The [managed objectID][docs-about-morefs] of this host.`,
				},
				resource.Attribute{
					Name:        "resource_pool_id",
					Description: `The [managed object ID][docs-about-morefs] of the host's root resource pool. -> Note that the resource pool referenced by [` + "`" + `resource_pool_id` + "`" + `](#resource_pool_id) is dependent on the target host's state - if it's a standalone host, the resource pool will belong to the host only, however if it is a member of a cluster, the resource pool will be the root for the entire cluster. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The [managed objectID][docs-about-morefs] of this host.`,
				},
				resource.Attribute{
					Name:        "resource_pool_id",
					Description: `The [managed object ID][docs-about-morefs] of the host's root resource pool. -> Note that the resource pool referenced by [` + "`" + `resource_pool_id` + "`" + `](#resource_pool_id) is dependent on the target host's state - if it's a standalone host, the resource pool will belong to the host only, however if it is a member of a cluster, the resource pool will be the root for the entire cluster. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_host_thumbprint",
			Category:         "Data Sources",
			ShortDescription: `A data source that can be used to get the thumbprint of an ESXi host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The address of the ESXi host to retrieve the thumbprint from.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port to use connecting to the ESXi host. Default: 443`,
				},
				resource.Attribute{
					Name:        "insecure",
					Description: `(Optional) Boolean that can be set to true to disable SSL certificate verification. Default: false ## Attribute Reference The only exported attribute is ` + "`" + `id` + "`" + `, which is the thumbprint of the ESXi host.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_network",
			Category:         "Data Sources",
			ShortDescription: `Provides a vSphere network data source. This can be used to get the general attributes of a vSphere network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the network. This can be a name or path.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of the datacenter the network is located in. This can be omitted if the search path used in ` + "`" + `name` + "`" + ` is an absolute path. For default datacenters, use the id attribute from an empty ` + "`" + `vsphere_datacenter` + "`" + ` data source.`,
				},
				resource.Attribute{
					Name:        "distributed_virtual_switch_uuid",
					Description: `(Optional) For distributed port group type network objects, the ID of the distributed virtual switch the given port group belongs to. It is useful to differentiate port groups with same name using the Distributed virtual switch ID. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ## Attribute Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_ovf_vm_template",
			Category:         "Data Sources",
			ShortDescription: `A data source that can be used to extract the configuration of an OVF template`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the virtual machine to create.`,
				},
				resource.Attribute{
					Name:        "resource_pool_id",
					Description: `(Required) The ID of a resource pool to put the virtual machine in.`,
				},
				resource.Attribute{
					Name:        "host_system_id",
					Description: `(Required) The ID of an optional host system to pin the virtual machine to.`,
				},
				resource.Attribute{
					Name:        "datastore_id",
					Description: `(Required) The ID of the virtual machine's datastore. The virtual machine configuration is placed here, along with any virtual disks that are created without datastores.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Required) The name of the folder to locate the virtual machine in.`,
				},
				resource.Attribute{
					Name:        "local_ovf_path",
					Description: `(Optional) The absolute path to the ovf/ova file in the local system. While deploying from ovf, make sure the other necessary files like the .vmdk files are also in the same directory as the given ovf file.`,
				},
				resource.Attribute{
					Name:        "remote_ovf_url",
					Description: `(Optional) URL to the remote ovf/ova file to be deployed. ~>`,
				},
				resource.Attribute{
					Name:        "ip_allocation_policy",
					Description: `(Optional) The IP allocation policy.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol.`,
				},
				resource.Attribute{
					Name:        "disk_provisioning",
					Description: `(Optional) The disk provisioning. If set, all the disks in the deployed OVF will have the same specified disk type (accepted values {thin, flat, thick, sameAsSource}).`,
				},
				resource.Attribute{
					Name:        "deployment_option",
					Description: `(Optional) The key of the chosen deployment option. If empty, the default option is chosen.`,
				},
				resource.Attribute{
					Name:        "ovf_network_map",
					Description: `(Optional) The mapping of name of network identifiers from the ovf descriptor to network UUID in the VI infrastructure.`,
				},
				resource.Attribute{
					Name:        "allow_unverified_ssl_cert",
					Description: `(Optional) Allow unverified ssl certificates while deploying ovf/ova from url. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "num_cpus",
					Description: `The number of virtual processors to assign to this virtual machine.`,
				},
				resource.Attribute{
					Name:        "num_cores_per_socket",
					Description: `The number of cores to distribute amongst the CPUs in this virtual machine.`,
				},
				resource.Attribute{
					Name:        "cpu_hot_add_enabled",
					Description: `Allow CPUs to be added to this virtual machine while it is running.`,
				},
				resource.Attribute{
					Name:        "cpu_hot_remove_enabled",
					Description: `Allow CPUs to be added to this virtual machine while it is running.`,
				},
				resource.Attribute{
					Name:        "nested_hv_enabled",
					Description: `Enable nested hardware virtualization on this virtual machine, facilitating nested virtualization in the guest.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The size of the virtual machine's memory, in MB.`,
				},
				resource.Attribute{
					Name:        "memory_hot_add_enabled",
					Description: `Allow memory to be added to this virtual machine while it is running.`,
				},
				resource.Attribute{
					Name:        "swap_placement_policy",
					Description: `The swap file placement policy for this virtual machine.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `User-provided description of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "guest_id",
					Description: `The guest ID for the operating system`,
				},
				resource.Attribute{
					Name:        "alternate_guest_name",
					Description: `The guest name for the operating system .`,
				},
				resource.Attribute{
					Name:        "firmware",
					Description: `The firmware interface to use on the virtual machine.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "num_cpus",
					Description: `The number of virtual processors to assign to this virtual machine.`,
				},
				resource.Attribute{
					Name:        "num_cores_per_socket",
					Description: `The number of cores to distribute amongst the CPUs in this virtual machine.`,
				},
				resource.Attribute{
					Name:        "cpu_hot_add_enabled",
					Description: `Allow CPUs to be added to this virtual machine while it is running.`,
				},
				resource.Attribute{
					Name:        "cpu_hot_remove_enabled",
					Description: `Allow CPUs to be added to this virtual machine while it is running.`,
				},
				resource.Attribute{
					Name:        "nested_hv_enabled",
					Description: `Enable nested hardware virtualization on this virtual machine, facilitating nested virtualization in the guest.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The size of the virtual machine's memory, in MB.`,
				},
				resource.Attribute{
					Name:        "memory_hot_add_enabled",
					Description: `Allow memory to be added to this virtual machine while it is running.`,
				},
				resource.Attribute{
					Name:        "swap_placement_policy",
					Description: `The swap file placement policy for this virtual machine.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `User-provided description of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "guest_id",
					Description: `The guest ID for the operating system`,
				},
				resource.Attribute{
					Name:        "alternate_guest_name",
					Description: `The guest name for the operating system .`,
				},
				resource.Attribute{
					Name:        "firmware",
					Description: `The firmware interface to use on the virtual machine.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_resource_pool",
			Category:         "Data Sources",
			ShortDescription: `Provides a vSphere resource pool data source. This can be used to get the general attributes of a vSphere resource pool.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the resource pool. This can be a name or path. This is required when using vCenter.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of the datacenter the resource pool is located in. This can be omitted if the search path used in ` + "`" + `name` + "`" + ` is an absolute path. For default datacenters, use the id attribute from an empty ` + "`" + `vsphere_datacenter` + "`" + ` data source. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_tag",
			Category:         "Data Sources",
			ShortDescription: `Provides a vSphere tag data source. This can be used to reference tags not managed in Terraform.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tag.`,
				},
				resource.Attribute{
					Name:        "category_id",
					Description: `(Required) The ID of the tag category the tag is located in. ## Attribute Reference In addition to the ` + "`" + `id` + "`" + ` being exported, all of the fields that are available in the [` + "`" + `vsphere_tag` + "`" + ` resource][resource-tag] are also populated. See that page for further details.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_tag_category",
			Category:         "Data Sources",
			ShortDescription: `Provides a vSphere tag category data source. This can be used to reference tag categories not managed in Terraform.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tag category. ## Attribute Reference In addition to the ` + "`" + `id` + "`" + ` being exported, all of the fields that are available in the [` + "`" + `vsphere_tag_category` + "`" + ` resource][resource-tag-category] are also populated. See that page for further details.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_virtual_machine",
			Category:         "Data Sources",
			ShortDescription: `Provides a vSphere virtual machine data source. This can be used to get data from a virtual machine or template.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual machine. This can be a name or path.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of the datacenter the virtual machine is located in. This can be omitted if the search path used in ` + "`" + `name` + "`" + ` is an absolute path. For default datacenters, use the ` + "`" + `id` + "`" + ` attribute from an empty ` + "`" + `vsphere_datacenter` + "`" + ` data source.`,
				},
				resource.Attribute{
					Name:        "scsi_controller_scan_count",
					Description: `(Optional) The number of SCSI controllers to scan for disk attributes and controller types on. Default: ` + "`" + `1` + "`" + `. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the virtual machine or template.`,
				},
				resource.Attribute{
					Name:        "guest_id",
					Description: `The guest ID of the virtual machine or template.`,
				},
				resource.Attribute{
					Name:        "alternate_guest_name",
					Description: `The alternate guest name of the virtual machine when guest_id is a non-specific operating system, like ` + "`" + `otherGuest` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `The user-provided description of this virtual machine.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The size of the virtual machine's memory, in MB.`,
				},
				resource.Attribute{
					Name:        "num_cpus",
					Description: `The total number of virtual processor cores assigned to this virtual machine.`,
				},
				resource.Attribute{
					Name:        "num_cores_per_socket",
					Description: `The number of cores per socket for this virtual machine.`,
				},
				resource.Attribute{
					Name:        "firmware",
					Description: `The firmware interface that is used by this virtual machine. Can be either ` + "`" + `bios` + "`" + ` or ` + "`" + `EFI` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hardware_version",
					Description: `The hardware version number on this virtual machine.`,
				},
				resource.Attribute{
					Name:        "scsi_type",
					Description: `The common type of all SCSI controllers on this virtual machine. Will be one of ` + "`" + `lsilogic` + "`" + ` (LSI Logic Parallel), ` + "`" + `lsilogic-sas` + "`" + ` (LSI Logic SAS), ` + "`" + `pvscsi` + "`" + ` (VMware Paravirtual), ` + "`" + `buslogic` + "`" + ` (BusLogic), or ` + "`" + `mixed` + "`" + ` when there are multiple controller types. Only the first number of controllers defined by ` + "`" + `scsi_controller_scan_count` + "`" + ` are scanned.`,
				},
				resource.Attribute{
					Name:        "scsi_bus_sharing",
					Description: `Mode for sharing the SCSI bus. The modes are physicalSharing, virtualSharing, and noSharing. Only the first number of controllers defined by ` + "`" + `scsi_controller_scan_count` + "`" + ` are scanned.`,
				},
				resource.Attribute{
					Name:        "disks",
					Description: `Information about each of the disks on this virtual machine or template. These are sorted by bus and unit number so that they can be applied to a ` + "`" + `vsphere_virtual_machine` + "`" + ` resource in the order the resource expects while cloning. This is useful for discovering certain disk settings while performing a linked clone, as all settings that are output by this data source must be the same on the destination virtual machine as the source. Only the first number of controllers defined by ` + "`" + `scsi_controller_scan_count` + "`" + ` are scanned for disks. The sub-attributes are:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label for the disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the disk, in GIB.`,
				},
				resource.Attribute{
					Name:        "eagerly_scrub",
					Description: `Set to ` + "`" + `true` + "`" + ` if the disk has been eager zeroed.`,
				},
				resource.Attribute{
					Name:        "thin_provisioned",
					Description: `Set to ` + "`" + `true` + "`" + ` if the disk has been thin provisioned.`,
				},
				resource.Attribute{
					Name:        "unit_number",
					Description: `The disk number on the storage bus.`,
				},
				resource.Attribute{
					Name:        "network_interface_types",
					Description: `The network interface types for each network interface found on the virtual machine, in device bus order. Will be one of ` + "`" + `e1000` + "`" + `, ` + "`" + `e1000e` + "`" + `, ` + "`" + `pcnet32` + "`" + `, ` + "`" + `sriov` + "`" + `, ` + "`" + `vmxnet2` + "`" + `, or ` + "`" + `vmxnet3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "firmware",
					Description: `The firmware type for this virtual machine. Can be ` + "`" + `bios` + "`" + ` or ` + "`" + `efi` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "guest_ip_addresses",
					Description: `A list of IP addresses as reported by VMWare tools. ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the virtual machine or template.`,
				},
				resource.Attribute{
					Name:        "guest_id",
					Description: `The guest ID of the virtual machine or template.`,
				},
				resource.Attribute{
					Name:        "alternate_guest_name",
					Description: `The alternate guest name of the virtual machine when guest_id is a non-specific operating system, like ` + "`" + `otherGuest` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `The user-provided description of this virtual machine.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The size of the virtual machine's memory, in MB.`,
				},
				resource.Attribute{
					Name:        "num_cpus",
					Description: `The total number of virtual processor cores assigned to this virtual machine.`,
				},
				resource.Attribute{
					Name:        "num_cores_per_socket",
					Description: `The number of cores per socket for this virtual machine.`,
				},
				resource.Attribute{
					Name:        "firmware",
					Description: `The firmware interface that is used by this virtual machine. Can be either ` + "`" + `bios` + "`" + ` or ` + "`" + `EFI` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hardware_version",
					Description: `The hardware version number on this virtual machine.`,
				},
				resource.Attribute{
					Name:        "scsi_type",
					Description: `The common type of all SCSI controllers on this virtual machine. Will be one of ` + "`" + `lsilogic` + "`" + ` (LSI Logic Parallel), ` + "`" + `lsilogic-sas` + "`" + ` (LSI Logic SAS), ` + "`" + `pvscsi` + "`" + ` (VMware Paravirtual), ` + "`" + `buslogic` + "`" + ` (BusLogic), or ` + "`" + `mixed` + "`" + ` when there are multiple controller types. Only the first number of controllers defined by ` + "`" + `scsi_controller_scan_count` + "`" + ` are scanned.`,
				},
				resource.Attribute{
					Name:        "scsi_bus_sharing",
					Description: `Mode for sharing the SCSI bus. The modes are physicalSharing, virtualSharing, and noSharing. Only the first number of controllers defined by ` + "`" + `scsi_controller_scan_count` + "`" + ` are scanned.`,
				},
				resource.Attribute{
					Name:        "disks",
					Description: `Information about each of the disks on this virtual machine or template. These are sorted by bus and unit number so that they can be applied to a ` + "`" + `vsphere_virtual_machine` + "`" + ` resource in the order the resource expects while cloning. This is useful for discovering certain disk settings while performing a linked clone, as all settings that are output by this data source must be the same on the destination virtual machine as the source. Only the first number of controllers defined by ` + "`" + `scsi_controller_scan_count` + "`" + ` are scanned for disks. The sub-attributes are:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label for the disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the disk, in GIB.`,
				},
				resource.Attribute{
					Name:        "eagerly_scrub",
					Description: `Set to ` + "`" + `true` + "`" + ` if the disk has been eager zeroed.`,
				},
				resource.Attribute{
					Name:        "thin_provisioned",
					Description: `Set to ` + "`" + `true` + "`" + ` if the disk has been thin provisioned.`,
				},
				resource.Attribute{
					Name:        "unit_number",
					Description: `The disk number on the storage bus.`,
				},
				resource.Attribute{
					Name:        "network_interface_types",
					Description: `The network interface types for each network interface found on the virtual machine, in device bus order. Will be one of ` + "`" + `e1000` + "`" + `, ` + "`" + `e1000e` + "`" + `, ` + "`" + `pcnet32` + "`" + `, ` + "`" + `sriov` + "`" + `, ` + "`" + `vmxnet2` + "`" + `, or ` + "`" + `vmxnet3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "firmware",
					Description: `The firmware type for this virtual machine. Can be ` + "`" + `bios` + "`" + ` or ` + "`" + `efi` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "guest_ip_addresses",
					Description: `A list of IP addresses as reported by VMWare tools. ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_vmfs_disks",
			Category:         "Data Sources",
			ShortDescription: `A data source that can be used to discover storage devices that can be used for VMFS datastores.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "host_system_id",
					Description: `(Required) The [managed object ID][docs-about-morefs] of the host to look for disks on. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "rescan",
					Description: `(Optional) Whether or not to rescan storage adapters before searching for disks. This may lengthen the time it takes to perform the search. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A regular expression to filter the disks against. Only disks with canonical names that match will be included. ~>`,
				},
				resource.Attribute{
					Name:        "disks",
					Description: `A lexicographically sorted list of devices discovered by the operation, matching the supplied ` + "`" + `filter` + "`" + `, if provided.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disks",
					Description: `A lexicographically sorted list of devices discovered by the operation, matching the supplied ` + "`" + `filter` + "`" + `, if provided.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_role",
			Category:         "Data Sources",
			ShortDescription: `A data source that can be used to fetch details of a vsphere role given its name/label.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label of the role. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the role.`,
				},
				resource.Attribute{
					Name:        "role_privileges",
					Description: `The privileges associated with the role.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The display label of the role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the role.`,
				},
				resource.Attribute{
					Name:        "role_privileges",
					Description: `The privileges associated with the role.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The display label of the role.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"vsphere_content_library_item":       0,
		"vsphere_custom_attribute":           1,
		"vsphere_datacenter":                 2,
		"vsphere_datastore":                  3,
		"vsphere_datastore_cluster":          4,
		"vsphere_distributed_virtual_switch": 5,
		"vsphere_host":                       6,
		"vsphere_host_thumbprint":            7,
		"vsphere_network":                    8,
		"vsphere_ovf_vm_template":            9,
		"vsphere_resource_pool":              10,
		"vsphere_tag":                        11,
		"vsphere_tag_category":               12,
		"vsphere_virtual_machine":            13,
		"vsphere_vmfs_disks":                 14,
		"vsphere_role":                       15,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
