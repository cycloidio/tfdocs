package opc

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_image_list_entry",
			Category:         "Data Sources",
			ShortDescription: `Gets information about the configuration of an Image List Entry within an Oracle Cloud Infrastructure Compute Classic domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_list",
					Description: `(Required) - The name of the image list to lookup.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) - The version (integer) of the Image List to use.`,
				},
				resource.Attribute{
					Name:        "entry",
					Description: `(Optional) - Which machine image to use. See [Entry](#entry) below for more details ## Entry The ` + "`" + `entry` + "`" + ` argument is fully optional when configuring the Data Source. If specified, however, the returned array of machine images will have a length of 1, and only contain the desired image. Thus, if "my_image_list" is an image list that contains the following images: ` + "`" + `` + "`" + `` + "`" + ` ["image1", "image2", "image3", "image4", "image5"] ` + "`" + `` + "`" + `` + "`" + ` Then specifing an ` + "`" + `entry` + "`" + ` of ` + "`" + `3` + "`" + `, the returned ` + "`" + `machine_images` + "`" + ` array will have a sigle element: ` + "`" + `"image3"` + "`" + `. If ` + "`" + `entry` + "`" + ` was omitted, or set to ` + "`" + `0` + "`" + `, the returned ` + "`" + `machine_images` + "`" + ` array will contain all of the images for that image list version. If the supplied ` + "`" + `entry` + "`" + ` value is invalid for the image list, Terraform will exit with an error, that the desired image list entry was not found. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `Array of DNS servers for the interface.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `JSON object of all of the image list's attributes`,
				},
				resource.Attribute{
					Name:        "machine_images",
					Description: `An array of machine images as strings`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The URI of the image list`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dns",
					Description: `Array of DNS servers for the interface.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `JSON object of all of the image list's attributes`,
				},
				resource.Attribute{
					Name:        "machine_images",
					Description: `An array of machine images as strings`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The URI of the image list`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_ip_address_reservation",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing IP Network IP address reservation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ip address reservation. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ip_address_pool",
					Description: `The IP address pool from which the IP address is allocated.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the ip address reservation.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of tags that applied to the IP address reservation.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The reserved IPv4 Public IP address.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier of the ip address reservation`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address_pool",
					Description: `The IP address pool from which the IP address is allocated.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the ip address reservation.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of tags that applied to the IP address reservation.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The reserved IPv4 Public IP address.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier of the ip address reservation`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_ip_reservation",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing IP reservation for the Shared Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the IP Reservation. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The reserved IPv4 Public IP address.`,
				},
				resource.Attribute{
					Name:        "permanent",
					Description: `Whether the IP address remains reserved even when it is no longer associated with an instance.`,
				},
				resource.Attribute{
					Name:        "parent_pool",
					Description: `The pool from which to allocate the IP address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of tags applied to the IP reservation.`,
				},
				resource.Attribute{
					Name:        "used",
					Description: `indicates that the IP reservation is associated with an instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip",
					Description: `The reserved IPv4 Public IP address.`,
				},
				resource.Attribute{
					Name:        "permanent",
					Description: `Whether the IP address remains reserved even when it is no longer associated with an instance.`,
				},
				resource.Attribute{
					Name:        "parent_pool",
					Description: `The pool from which to allocate the IP address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of tags applied to the IP reservation.`,
				},
				resource.Attribute{
					Name:        "used",
					Description: `indicates that the IP reservation is associated with an instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_machine_image",
			Category:         "Data Sources",
			ShortDescription: `Gets information about the configuration of an Machine Image in a Compute Classic identity domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account",
					Description: `(Required) The two part name of the compute object storage account in the format ` + "`" + `/Compute-{identity_domain}/cloud_storage` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Machine Image. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `The name of the Machine Image .tar.gz file in the ` + "`" + `compute_images` + "`" + ` storage container.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the Machine Image.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `An optional JSON object of arbitrary attributes to be made available to the instance. These are user-defined tags. After defining attributes, you can view them from within an instance at http://192.0.0.192/`,
				},
				resource.Attribute{
					Name:        "error_reason",
					Description: `Description of the state of the machine image if there is an error.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `Dictionary of hypervisor-specific attributes.`,
				},
				resource.Attribute{
					Name:        "image_format",
					Description: `The format of the image.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The OS platform of the image.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the uploaded machine image.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Machine Image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "file",
					Description: `The name of the Machine Image .tar.gz file in the ` + "`" + `compute_images` + "`" + ` storage container.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the Machine Image.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `An optional JSON object of arbitrary attributes to be made available to the instance. These are user-defined tags. After defining attributes, you can view them from within an instance at http://192.0.0.192/`,
				},
				resource.Attribute{
					Name:        "error_reason",
					Description: `Description of the state of the machine image if there is an error.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `Dictionary of hypervisor-specific attributes.`,
				},
				resource.Attribute{
					Name:        "image_format",
					Description: `The format of the image.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The OS platform of the image.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the uploaded machine image.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Machine Image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_network_interface",
			Category:         "Data Sources",
			ShortDescription: `Gets information about the configuration of an instance's network interface`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dns",
					Description: `Array of DNS servers for the interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP Address assigned to the interface.`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `The IP Network assigned to the interface.`,
				},
				resource.Attribute{
					Name:        "is_default_gateway",
					Description: `Whether or not the the interface is the default gateway.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address of the interface.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `The model of the NIC card used.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Array of name servers for the interface.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `The IP Reservation (in IP Networks) associated with the interface.`,
				},
				resource.Attribute{
					Name:        "search_domains",
					Description: `The search domains that are sent through DHCP as option 119.`,
				},
				resource.Attribute{
					Name:        "sec_lists",
					Description: `The security lists the interface is added to.`,
				},
				resource.Attribute{
					Name:        "shared_network",
					Description: `Whether or not the interface is inside the Shared Network or an IP Network.`,
				},
				resource.Attribute{
					Name:        "vnic",
					Description: `The name of the vNIC created for the IP Network.`,
				},
				resource.Attribute{
					Name:        "vnic_sets",
					Description: `The array of vNIC Sets the interface was added to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dns",
					Description: `Array of DNS servers for the interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP Address assigned to the interface.`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `The IP Network assigned to the interface.`,
				},
				resource.Attribute{
					Name:        "is_default_gateway",
					Description: `Whether or not the the interface is the default gateway.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address of the interface.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `The model of the NIC card used.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Array of name servers for the interface.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `The IP Reservation (in IP Networks) associated with the interface.`,
				},
				resource.Attribute{
					Name:        "search_domains",
					Description: `The search domains that are sent through DHCP as option 119.`,
				},
				resource.Attribute{
					Name:        "sec_lists",
					Description: `The security lists the interface is added to.`,
				},
				resource.Attribute{
					Name:        "shared_network",
					Description: `Whether or not the interface is inside the Shared Network or an IP Network.`,
				},
				resource.Attribute{
					Name:        "vnic",
					Description: `The name of the vNIC created for the IP Network.`,
				},
				resource.Attribute{
					Name:        "vnic_sets",
					Description: `The array of vNIC Sets the interface was added to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing SSH key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique (within this identity domain) name of the SSH key. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The public SSH key.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether or not the key is enabled.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `The public SSH key.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether or not the key is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_storage_volume_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Gets information about the configuration of a storage volume snapshot`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account",
					Description: `Account of the snapshot.`,
				},
				resource.Attribute{
					Name:        "collocated",
					Description: `Boolean specifying whether the snapshot is collocated or remote.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the storage volume snapshot.`,
				},
				resource.Attribute{
					Name:        "machine_image_name",
					Description: `The name of the machine image that's used in the boot volume from which this snapshot is taken.`,
				},
				resource.Attribute{
					Name:        "parent_volume_bootable",
					Description: `Boolean specifying whether or not the snapshot's parent volume was bootable.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `Where the snapshot is stored, whether collocated, or in the Oracle Storage Cloud Service instance.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The OS platform this snapshot is compatible with`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "snapshot_timestamp",
					Description: `Timestamp of the storage snapshot, generated by storage server. The snapshot will contain data written to the original volume before this time.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The Oracle ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "start_timestamp",
					Description: `Timestamp when the snapshot was started.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "status_detail",
					Description: `Details about the latest state of the storage volume snapshot.`,
				},
				resource.Attribute{
					Name:        "status_timestamp",
					Description: `Indicates the time that the current view of the storage volume snapshot was generated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Comma-separated strings that tag the storage volume.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `Uniform Resource Identifier`,
				},
				resource.Attribute{
					Name:        "volume_name",
					Description: `The name of the storage volume that the snapshot was created from`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account",
					Description: `Account of the snapshot.`,
				},
				resource.Attribute{
					Name:        "collocated",
					Description: `Boolean specifying whether the snapshot is collocated or remote.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the storage volume snapshot.`,
				},
				resource.Attribute{
					Name:        "machine_image_name",
					Description: `The name of the machine image that's used in the boot volume from which this snapshot is taken.`,
				},
				resource.Attribute{
					Name:        "parent_volume_bootable",
					Description: `Boolean specifying whether or not the snapshot's parent volume was bootable.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `Where the snapshot is stored, whether collocated, or in the Oracle Storage Cloud Service instance.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The OS platform this snapshot is compatible with`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "snapshot_timestamp",
					Description: `Timestamp of the storage snapshot, generated by storage server. The snapshot will contain data written to the original volume before this time.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The Oracle ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "start_timestamp",
					Description: `Timestamp when the snapshot was started.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "status_detail",
					Description: `Details about the latest state of the storage volume snapshot.`,
				},
				resource.Attribute{
					Name:        "status_timestamp",
					Description: `Indicates the time that the current view of the storage volume snapshot was generated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Comma-separated strings that tag the storage volume.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `Uniform Resource Identifier`,
				},
				resource.Attribute{
					Name:        "volume_name",
					Description: `The name of the storage volume that the snapshot was created from`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_vnic",
			Category:         "Data Sources",
			ShortDescription: `Gets information about the configuration of a Virtual NIC.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"opc_compute_image_list_entry":        0,
		"opc_compute_ip_address_reservation":  1,
		"opc_compute_ip_reservation":          2,
		"opc_compute_machine_image":           3,
		"opc_compute_network_interface":       4,
		"opc_compute_ssh_key":                 5,
		"opc_compute_storage_volume_snapshot": 6,
		"opc_compute_vnic":                    7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
