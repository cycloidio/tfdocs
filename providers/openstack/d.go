package openstack

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "openstack_blockstorage_availability_zones_v3",
			Category:         "Data Sources",
			ShortDescription: `Get a list of Block Storage availability zones from OpenStack`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the Block Storage client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The ` + "`" + `state` + "`" + ` of the availability zones to match. Can either be ` + "`" + `available` + "`" + ` or ` + "`" + `unavailable` + "`" + `. Default is ` + "`" + `available` + "`" + `. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to hash of the returned zone list. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_blockstorage_quotaset_v3",
			Category:         "Data Sources",
			ShortDescription: `Get information on a BlockStorage Quotaset v3 of a project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Blockstorage client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project to retrieve the quotaset. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `The number of volumes that are allowed.`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `The number of snapshots that are allowed.`,
				},
				resource.Attribute{
					Name:        "gigabytes",
					Description: `The size (GB) of volumes and snapshots that are allowed.`,
				},
				resource.Attribute{
					Name:        "per_volume_gigabytes",
					Description: `The size (GB) of volumes that are allowed for each volume.`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `The number of backups that are allowed.`,
				},
				resource.Attribute{
					Name:        "backup_gigabytes",
					Description: `The size (GB) of backups that are allowed.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `The number of groups that are allowed.`,
				},
				resource.Attribute{
					Name:        "volume_type_quota",
					Description: `Map with gigabytes_{volume_type}, snapshots_{volume_type}, volumes_{volume_type} for each volume type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `The number of volumes that are allowed.`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `The number of snapshots that are allowed.`,
				},
				resource.Attribute{
					Name:        "gigabytes",
					Description: `The size (GB) of volumes and snapshots that are allowed.`,
				},
				resource.Attribute{
					Name:        "per_volume_gigabytes",
					Description: `The size (GB) of volumes that are allowed for each volume.`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `The number of backups that are allowed.`,
				},
				resource.Attribute{
					Name:        "backup_gigabytes",
					Description: `The size (GB) of backups that are allowed.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `The number of groups that are allowed.`,
				},
				resource.Attribute{
					Name:        "volume_type_quota",
					Description: `Map with gigabytes_{volume_type}, snapshots_{volume_type}, volumes_{volume_type} for each volume type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_blockstorage_snapshot_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Snapshot.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Block Storage client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Optional) The ID of the snapshot's volume.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) Pick the most recently created snapshot if there are multiple results. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The snapshot's description.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The snapshot's metadata.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The snapshot's description.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The snapshot's metadata.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_blockstorage_snapshot_v3",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Snapshot.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Block Storage client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Optional) The ID of the snapshot's volume.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) Pick the most recently created snapshot if there are multiple results. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The snapshot's description.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The snapshot's metadata.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The snapshot's description.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The snapshot's metadata.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_blockstorage_volume_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Volume.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Block Storage client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the volume.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the volume.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs associated with the volume. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found volume. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume.`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `Indicates if the volume is bootable.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume in GBs.`,
				},
				resource.Attribute{
					Name:        "source_volume_id",
					Description: `The ID of the volume from which the current volume was created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume.`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `Indicates if the volume is bootable.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume in GBs.`,
				},
				resource.Attribute{
					Name:        "source_volume_id",
					Description: `The ID of the volume from which the current volume was created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_blockstorage_volume_v3",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Volume.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Block Storage client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the volume.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the volume.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs associated with the volume. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found volume. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume.`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `Indicates if the volume is bootable.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume in GBs.`,
				},
				resource.Attribute{
					Name:        "source_volume_id",
					Description: `The ID of the volume from which the current volume was created.`,
				},
				resource.Attribute{
					Name:        "multiattach",
					Description: `Indicates if the volume can be attached to more then one server.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The OpenStack host on which the volume is located.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume.`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `Indicates if the volume is bootable.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume in GBs.`,
				},
				resource.Attribute{
					Name:        "source_volume_id",
					Description: `The ID of the volume from which the current volume was created.`,
				},
				resource.Attribute{
					Name:        "multiattach",
					Description: `Indicates if the volume can be attached to more then one server.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The OpenStack host on which the volume is located.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_availability_zones_v2",
			Category:         "Data Sources",
			ShortDescription: `Get a list of availability zones from OpenStack`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The ` + "`" + `region` + "`" + ` to fetch availability zones from, defaults to the provider's ` + "`" + `region` + "`" + ``,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The ` + "`" + `state` + "`" + ` of the availability zones to match, default ("available"). ## Attributes Reference ` + "`" + `id` + "`" + ` is set to hash of the returned zone list. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_flavor_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Flavor.`,
			Description:      ``,
			Icon:             "flavor-gray.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Optional) The ID of the flavor. Conflicts with the ` + "`" + `name` + "`" + `, ` + "`" + `min_ram` + "`" + ` and ` + "`" + `min_disk` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the flavor. Conflicts with the ` + "`" + `flavor_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_ram",
					Description: `(Optional) The minimum amount of RAM (in megabytes). Conflicts with the ` + "`" + `flavor_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `(Optional) The exact amount of RAM (in megabytes).`,
				},
				resource.Attribute{
					Name:        "min_disk",
					Description: `(Optional) The minimum amount of disk (in gigabytes). Conflicts with the ` + "`" + `flavor_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional) The exact amount of disk (in gigabytes).`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `(Optional) The amount of VCPUs.`,
				},
				resource.Attribute{
					Name:        "swap",
					Description: `(Optional) The amount of swap (in gigabytes).`,
				},
				resource.Attribute{
					Name:        "rx_tx_factor",
					Description: `(Optional) The ` + "`" + `rx_tx_factor` + "`" + ` of the flavor.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) The flavor visibility. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found flavor. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "extra_specs",
					Description: `Key/Value pairs of metadata for the flavor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "extra_specs",
					Description: `Key/Value pairs of metadata for the flavor.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_keypair_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Keypair.`,
			Description:      ``,
			Icon:             "keypair-gray.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the keypair. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the OpenSSH key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The OpenSSH-formatted public key of the keypair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the OpenSSH key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The OpenSSH-formatted public key of the keypair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_limits_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Compute Limits of a project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project to retrieve the limits. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_total_cores",
					Description: `The number of allowed server cores for the tenant.`,
				},
				resource.Attribute{
					Name:        "max_image_meta",
					Description: `The number of allowed metadata items for each image. Starting from version 2.39 this field is dropped from ‘os-limits’ response, because ‘image-metadata’ proxy API was deprecated. Available until version 2.38.`,
				},
				resource.Attribute{
					Name:        "max_server_meta",
					Description: `The number of allowed server groups for the tenant.`,
				},
				resource.Attribute{
					Name:        "max_personality",
					Description: `The number of allowed injected files for the tenant. Available until version 2.56.`,
				},
				resource.Attribute{
					Name:        "max_personality_size",
					Description: `The number of allowed bytes of content for each injected file. Available until version 2.56.`,
				},
				resource.Attribute{
					Name:        "max_total_keypairs",
					Description: `The number of allowed key pairs for the user.`,
				},
				resource.Attribute{
					Name:        "max_security_groups",
					Description: `The number of allowed security groups for the tenant. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "max_security_group_rules",
					Description: `The number of allowed rules for each security group. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "max_server_groups",
					Description: `The number of allowed server groups for the tenant.`,
				},
				resource.Attribute{
					Name:        "max_server_group_members",
					Description: `The number of allowed members for each server group.`,
				},
				resource.Attribute{
					Name:        "max_total_floating_ips",
					Description: `The number of allowed floating IP addresses for each tenant. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "max_total_instances",
					Description: `The number of allowed servers for the tenant.`,
				},
				resource.Attribute{
					Name:        "max_total_ram_size",
					Description: `The number of allowed floating IP addresses for the tenant. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "total_cores_used",
					Description: `The number of used server cores in the tenant.`,
				},
				resource.Attribute{
					Name:        "total_instances_used",
					Description: `The number of used server cores in the tenant.`,
				},
				resource.Attribute{
					Name:        "total_floating_ips_used",
					Description: `The number of used floating IP addresses in the tenant.`,
				},
				resource.Attribute{
					Name:        "total_ram_used",
					Description: `The amount of used server RAM in the tenant.`,
				},
				resource.Attribute{
					Name:        "total_security_groups_used",
					Description: `The number of used security groups in the tenant. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "total_server_groups_used",
					Description: `The number of used server groups in each tenant.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_total_cores",
					Description: `The number of allowed server cores for the tenant.`,
				},
				resource.Attribute{
					Name:        "max_image_meta",
					Description: `The number of allowed metadata items for each image. Starting from version 2.39 this field is dropped from ‘os-limits’ response, because ‘image-metadata’ proxy API was deprecated. Available until version 2.38.`,
				},
				resource.Attribute{
					Name:        "max_server_meta",
					Description: `The number of allowed server groups for the tenant.`,
				},
				resource.Attribute{
					Name:        "max_personality",
					Description: `The number of allowed injected files for the tenant. Available until version 2.56.`,
				},
				resource.Attribute{
					Name:        "max_personality_size",
					Description: `The number of allowed bytes of content for each injected file. Available until version 2.56.`,
				},
				resource.Attribute{
					Name:        "max_total_keypairs",
					Description: `The number of allowed key pairs for the user.`,
				},
				resource.Attribute{
					Name:        "max_security_groups",
					Description: `The number of allowed security groups for the tenant. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "max_security_group_rules",
					Description: `The number of allowed rules for each security group. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "max_server_groups",
					Description: `The number of allowed server groups for the tenant.`,
				},
				resource.Attribute{
					Name:        "max_server_group_members",
					Description: `The number of allowed members for each server group.`,
				},
				resource.Attribute{
					Name:        "max_total_floating_ips",
					Description: `The number of allowed floating IP addresses for each tenant. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "max_total_instances",
					Description: `The number of allowed servers for the tenant.`,
				},
				resource.Attribute{
					Name:        "max_total_ram_size",
					Description: `The number of allowed floating IP addresses for the tenant. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "total_cores_used",
					Description: `The number of used server cores in the tenant.`,
				},
				resource.Attribute{
					Name:        "total_instances_used",
					Description: `The number of used server cores in the tenant.`,
				},
				resource.Attribute{
					Name:        "total_floating_ips_used",
					Description: `The number of used floating IP addresses in the tenant.`,
				},
				resource.Attribute{
					Name:        "total_ram_used",
					Description: `The amount of used server RAM in the tenant.`,
				},
				resource.Attribute{
					Name:        "total_security_groups_used",
					Description: `The number of used security groups in the tenant. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "total_server_groups_used",
					Description: `The number of used server groups in each tenant.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_quotaset_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Compute Quotaset of a project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project to retrieve the quotaset. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The number of allowed server cores.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The number of allowed servers.`,
				},
				resource.Attribute{
					Name:        "key_pairs",
					Description: `The number of allowed key pairs for each user.`,
				},
				resource.Attribute{
					Name:        "metadata_items",
					Description: `The number of allowed metadata items for each server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of allowed server RAM, in MiB.`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `The number of allowed server groups.`,
				},
				resource.Attribute{
					Name:        "server_group_members",
					Description: `The number of allowed members for each server group.`,
				},
				resource.Attribute{
					Name:        "fixed_ips",
					Description: `The number of allowed fixed IP addresses. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "floating_ips",
					Description: `The number of allowed floating IP addresses. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "security_group_rules",
					Description: `The number of allowed rules for each security group. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `The number of allowed security groups. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "injected_file_content_bytes",
					Description: `The number of allowed bytes of content for each injected file. Available until version 2.56.`,
				},
				resource.Attribute{
					Name:        "injected_file_path_bytes",
					Description: `The number of allowed bytes for each injected file path. Available until version 2.56.`,
				},
				resource.Attribute{
					Name:        "injected_files",
					Description: `The number of allowed injected files. Available until version 2.56.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The number of allowed server cores.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The number of allowed servers.`,
				},
				resource.Attribute{
					Name:        "key_pairs",
					Description: `The number of allowed key pairs for each user.`,
				},
				resource.Attribute{
					Name:        "metadata_items",
					Description: `The number of allowed metadata items for each server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of allowed server RAM, in MiB.`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `The number of allowed server groups.`,
				},
				resource.Attribute{
					Name:        "server_group_members",
					Description: `The number of allowed members for each server group.`,
				},
				resource.Attribute{
					Name:        "fixed_ips",
					Description: `The number of allowed fixed IP addresses. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "floating_ips",
					Description: `The number of allowed floating IP addresses. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "security_group_rules",
					Description: `The number of allowed rules for each security group. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `The number of allowed security groups. Available until version 2.35.`,
				},
				resource.Attribute{
					Name:        "injected_file_content_bytes",
					Description: `The number of allowed bytes of content for each injected file. Available until version 2.56.`,
				},
				resource.Attribute{
					Name:        "injected_file_path_bytes",
					Description: `The number of allowed bytes for each injected file path. Available until version 2.56.`,
				},
				resource.Attribute{
					Name:        "injected_files",
					Description: `The number of allowed injected files. Available until version 2.56.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_containerinfra_cluster_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Magnum cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V1 Container Infra client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the cluster. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found cluster. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project of the cluster.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user of the cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which cluster was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which cluster was updated.`,
				},
				resource.Attribute{
					Name:        "api_address",
					Description: `COE API address.`,
				},
				resource.Attribute{
					Name:        "coe_version",
					Description: `COE software version.`,
				},
				resource.Attribute{
					Name:        "cluster_template_id",
					Description: `The UUID of the V1 Container Infra cluster template.`,
				},
				resource.Attribute{
					Name:        "create_timeout",
					Description: `The timeout (in minutes) for creating the cluster.`,
				},
				resource.Attribute{
					Name:        "discovery_url",
					Description: `The URL used for cluster node discovery.`,
				},
				resource.Attribute{
					Name:        "docker_volume_size",
					Description: `The size (in GB) of the Docker volume.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `The flavor for the nodes of the cluster.`,
				},
				resource.Attribute{
					Name:        "master_flavor",
					Description: `The flavor for the master nodes.`,
				},
				resource.Attribute{
					Name:        "keypair",
					Description: `The name of the Compute service SSH keypair.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of key value pairs representing additional properties of the cluster.`,
				},
				resource.Attribute{
					Name:        "master_count",
					Description: `The number of master nodes for the cluster.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The number of nodes for the cluster.`,
				},
				resource.Attribute{
					Name:        "fixed_network",
					Description: `The fixed network that is attached to the cluster.`,
				},
				resource.Attribute{
					Name:        "fixed_subnet",
					Description: `The fixed subnet that is attached to the cluster.`,
				},
				resource.Attribute{
					Name:        "master_addresses",
					Description: `IP addresses of the master node of the cluster.`,
				},
				resource.Attribute{
					Name:        "node_addresses",
					Description: `IP addresses of the node of the cluster.`,
				},
				resource.Attribute{
					Name:        "stack_id",
					Description: `UUID of the Orchestration service stack.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project of the cluster.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user of the cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which cluster was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which cluster was updated.`,
				},
				resource.Attribute{
					Name:        "api_address",
					Description: `COE API address.`,
				},
				resource.Attribute{
					Name:        "coe_version",
					Description: `COE software version.`,
				},
				resource.Attribute{
					Name:        "cluster_template_id",
					Description: `The UUID of the V1 Container Infra cluster template.`,
				},
				resource.Attribute{
					Name:        "create_timeout",
					Description: `The timeout (in minutes) for creating the cluster.`,
				},
				resource.Attribute{
					Name:        "discovery_url",
					Description: `The URL used for cluster node discovery.`,
				},
				resource.Attribute{
					Name:        "docker_volume_size",
					Description: `The size (in GB) of the Docker volume.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `The flavor for the nodes of the cluster.`,
				},
				resource.Attribute{
					Name:        "master_flavor",
					Description: `The flavor for the master nodes.`,
				},
				resource.Attribute{
					Name:        "keypair",
					Description: `The name of the Compute service SSH keypair.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of key value pairs representing additional properties of the cluster.`,
				},
				resource.Attribute{
					Name:        "master_count",
					Description: `The number of master nodes for the cluster.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The number of nodes for the cluster.`,
				},
				resource.Attribute{
					Name:        "fixed_network",
					Description: `The fixed network that is attached to the cluster.`,
				},
				resource.Attribute{
					Name:        "fixed_subnet",
					Description: `The fixed subnet that is attached to the cluster.`,
				},
				resource.Attribute{
					Name:        "master_addresses",
					Description: `IP addresses of the master node of the cluster.`,
				},
				resource.Attribute{
					Name:        "node_addresses",
					Description: `IP addresses of the node of the cluster.`,
				},
				resource.Attribute{
					Name:        "stack_id",
					Description: `UUID of the Orchestration service stack.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_containerinfra_clustertemplate_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Magnum cluster template.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V1 Container Infra client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the cluster template. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found cluster template. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project of the cluster template.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user of the cluster template.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which cluster template was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which cluster template was updated.`,
				},
				resource.Attribute{
					Name:        "apiserver_port",
					Description: `The API server port for the Container Orchestration Engine for this cluster template.`,
				},
				resource.Attribute{
					Name:        "coe",
					Description: `The Container Orchestration Engine for this cluster template.`,
				},
				resource.Attribute{
					Name:        "cluster_distro",
					Description: `The distro for the cluster (fedora-atomic, coreos, etc.).`,
				},
				resource.Attribute{
					Name:        "dns_nameserver",
					Description: `Address of the DNS nameserver that is used in nodes of the cluster.`,
				},
				resource.Attribute{
					Name:        "docker_storage_driver",
					Description: `Docker storage driver. Changing this updates the Docker storage driver of the existing cluster template.`,
				},
				resource.Attribute{
					Name:        "docker_volume_size",
					Description: `The size (in GB) of the Docker volume.`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `The ID of the external network that will be used for the cluster.`,
				},
				resource.Attribute{
					Name:        "fixed_network",
					Description: `The fixed network that will be attached to the cluster.`,
				},
				resource.Attribute{
					Name:        "fixed_subnet",
					Description: `=The fixed subnet that will be attached to the cluster.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `The flavor for the nodes of the cluster.`,
				},
				resource.Attribute{
					Name:        "master_flavor",
					Description: `The flavor for the master nodes.`,
				},
				resource.Attribute{
					Name:        "floating_ip_enabled",
					Description: `Indicates whether created cluster should create IP floating IP for every node or not.`,
				},
				resource.Attribute{
					Name:        "http_proxy",
					Description: `The address of a proxy for receiving all HTTP requests and relay them.`,
				},
				resource.Attribute{
					Name:        "https_proxy",
					Description: `The address of a proxy for receiving all HTTPS requests and relay them.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The reference to an image that is used for nodes of the cluster.`,
				},
				resource.Attribute{
					Name:        "insecure_registry",
					Description: `The insecure registry URL for the cluster template.`,
				},
				resource.Attribute{
					Name:        "keypair_id",
					Description: `The name of the Compute service SSH keypair.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of key value pairs representing additional properties of the cluster template.`,
				},
				resource.Attribute{
					Name:        "master_lb_enabled",
					Description: `Indicates whether created cluster should has a loadbalancer for master nodes or not.`,
				},
				resource.Attribute{
					Name:        "network_driver",
					Description: `The name of the driver for the container network.`,
				},
				resource.Attribute{
					Name:        "no_proxy",
					Description: `A comma-separated list of IP addresses that shouldn't be used in the cluster.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates whether cluster template should be public.`,
				},
				resource.Attribute{
					Name:        "registry_enabled",
					Description: `Indicates whether Docker registry is enabled in the cluster.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `The server type for the cluster template.`,
				},
				resource.Attribute{
					Name:        "tls_disabled",
					Description: `Indicates whether the TLS should be disabled in the cluster.`,
				},
				resource.Attribute{
					Name:        "volume_driver",
					Description: `The name of the driver that is used for the volumes of the cluster nodes.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project of the cluster template.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user of the cluster template.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which cluster template was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which cluster template was updated.`,
				},
				resource.Attribute{
					Name:        "apiserver_port",
					Description: `The API server port for the Container Orchestration Engine for this cluster template.`,
				},
				resource.Attribute{
					Name:        "coe",
					Description: `The Container Orchestration Engine for this cluster template.`,
				},
				resource.Attribute{
					Name:        "cluster_distro",
					Description: `The distro for the cluster (fedora-atomic, coreos, etc.).`,
				},
				resource.Attribute{
					Name:        "dns_nameserver",
					Description: `Address of the DNS nameserver that is used in nodes of the cluster.`,
				},
				resource.Attribute{
					Name:        "docker_storage_driver",
					Description: `Docker storage driver. Changing this updates the Docker storage driver of the existing cluster template.`,
				},
				resource.Attribute{
					Name:        "docker_volume_size",
					Description: `The size (in GB) of the Docker volume.`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `The ID of the external network that will be used for the cluster.`,
				},
				resource.Attribute{
					Name:        "fixed_network",
					Description: `The fixed network that will be attached to the cluster.`,
				},
				resource.Attribute{
					Name:        "fixed_subnet",
					Description: `=The fixed subnet that will be attached to the cluster.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `The flavor for the nodes of the cluster.`,
				},
				resource.Attribute{
					Name:        "master_flavor",
					Description: `The flavor for the master nodes.`,
				},
				resource.Attribute{
					Name:        "floating_ip_enabled",
					Description: `Indicates whether created cluster should create IP floating IP for every node or not.`,
				},
				resource.Attribute{
					Name:        "http_proxy",
					Description: `The address of a proxy for receiving all HTTP requests and relay them.`,
				},
				resource.Attribute{
					Name:        "https_proxy",
					Description: `The address of a proxy for receiving all HTTPS requests and relay them.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The reference to an image that is used for nodes of the cluster.`,
				},
				resource.Attribute{
					Name:        "insecure_registry",
					Description: `The insecure registry URL for the cluster template.`,
				},
				resource.Attribute{
					Name:        "keypair_id",
					Description: `The name of the Compute service SSH keypair.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of key value pairs representing additional properties of the cluster template.`,
				},
				resource.Attribute{
					Name:        "master_lb_enabled",
					Description: `Indicates whether created cluster should has a loadbalancer for master nodes or not.`,
				},
				resource.Attribute{
					Name:        "network_driver",
					Description: `The name of the driver for the container network.`,
				},
				resource.Attribute{
					Name:        "no_proxy",
					Description: `A comma-separated list of IP addresses that shouldn't be used in the cluster.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates whether cluster template should be public.`,
				},
				resource.Attribute{
					Name:        "registry_enabled",
					Description: `Indicates whether Docker registry is enabled in the cluster.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `The server type for the cluster template.`,
				},
				resource.Attribute{
					Name:        "tls_disabled",
					Description: `Indicates whether the TLS should be disabled in the cluster.`,
				},
				resource.Attribute{
					Name:        "volume_driver",
					Description: `The name of the driver that is used for the volumes of the cluster nodes.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_containerinfra_nodegroup_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Magnum node group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V1 Container Infra client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The name of the OpenStack Magnum cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the node group. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found node group. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project of the node group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which the node group was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which the node group was updated.`,
				},
				resource.Attribute{
					Name:        "docker_volume_size",
					Description: `The size (in GB) of the Docker volume.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of key value pairs representing additional properties of the node group.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `The role of the node group.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The number of nodes for the node group.`,
				},
				resource.Attribute{
					Name:        "min_node_count",
					Description: `The minimum number of nodes for the node group.`,
				},
				resource.Attribute{
					Name:        "max_node_count",
					Description: `The maximum number of nodes for the node group.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The reference to an image that is used for nodes of the node group.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `The flavor for the nodes of the node group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project of the node group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which the node group was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which the node group was updated.`,
				},
				resource.Attribute{
					Name:        "docker_volume_size",
					Description: `The size (in GB) of the Docker volume.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of key value pairs representing additional properties of the node group.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `The role of the node group.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The number of nodes for the node group.`,
				},
				resource.Attribute{
					Name:        "min_node_count",
					Description: `The minimum number of nodes for the node group.`,
				},
				resource.Attribute{
					Name:        "max_node_count",
					Description: `The maximum number of nodes for the node group.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The reference to an image that is used for nodes of the node group.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `The flavor for the nodes of the node group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_dns_zone_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack DNS Zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 DNS client. A DNS client is needed to retrieve zone ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the zone.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The ID of the project the DNS zone is obtained from, sets ` + "`" + `X-Auth-Sudo-Tenant-ID` + "`" + ` header (requires an assigned user role in target project)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the zone.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The email contact for the zone record.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The zone's status.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the zone.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the zone. Can either be ` + "`" + `PRIMARY` + "`" + ` or ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "all_projects",
					Description: `(Optional) Try to obtain zone ID by listing all projects (requires admin role by default, depends on your policy configuration) ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found zone. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
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
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `Attributes of the DNS Service scheduler.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `An array of master DNS servers. When ` + "`" + `type` + "`" + ` is ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the zone was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time the zone was last updated.`,
				},
				resource.Attribute{
					Name:        "transferred_at",
					Description: `The time the zone was transferred.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the zone.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial number of the zone.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `The ID of the pool hosting the zone.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project ID that owns the zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
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
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `Attributes of the DNS Service scheduler.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `An array of master DNS servers. When ` + "`" + `type` + "`" + ` is ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the zone was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time the zone was last updated.`,
				},
				resource.Attribute{
					Name:        "transferred_at",
					Description: `The time the zone was transferred.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the zone.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial number of the zone.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `The ID of the pool hosting the zone.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project ID that owns the zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_fw_policy_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Firewall Policy.`,
			Description:      ``,
			Icon:             "firewall-gray.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve firewall policy ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) The ID of the firewall policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the firewall policy.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the firewall policy. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the firewall policy.`,
				},
				resource.Attribute{
					Name:        "audited",
					Description: `The audit status of the firewall policy.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `The sharing status of the firewall policy.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `The array of one or more firewall rules that comprise the policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the firewall policy.`,
				},
				resource.Attribute{
					Name:        "audited",
					Description: `The audit status of the firewall policy.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `The sharing status of the firewall policy.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `The array of one or more firewall rules that comprise the policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_auth_scope_v3",
			Category:         "Data Sources",
			ShortDescription: `Get authentication information from the current authenticated scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the scope. This is an arbitrary name which is only used as a unique identifier so an actual token isn't used as the ID.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Identity client. A Identity client is needed to retrieve tokens IDs. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the name given to the scope. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The username of the scope.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user ID the of the scope.`,
				},
				resource.Attribute{
					Name:        "user_domain_name",
					Description: `The domain name of the user.`,
				},
				resource.Attribute{
					Name:        "user_domain_id",
					Description: `The domain ID of the user.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name of the scope.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The domain ID of the scope.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `The project name of the scope.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project ID of the scope.`,
				},
				resource.Attribute{
					Name:        "project_domain_name",
					Description: `The domain name of the project.`,
				},
				resource.Attribute{
					Name:        "project_domain_id",
					Description: `The domain ID of the project.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles in the current scope. See reference below.`,
				},
				resource.Attribute{
					Name:        "service_catalog",
					Description: `A list of service catalog entries returned with the token. The ` + "`" + `roles` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The ID of the role.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `The name of the role. The ` + "`" + `service_catalog` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the service.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `A list of endpoints for the service. The ` + "`" + `endpoints` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the endpoint.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the endpoint.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region ID of the endpoint.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `The interface of the endpoint.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the endpoint.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_name",
					Description: `The username of the scope.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user ID the of the scope.`,
				},
				resource.Attribute{
					Name:        "user_domain_name",
					Description: `The domain name of the user.`,
				},
				resource.Attribute{
					Name:        "user_domain_id",
					Description: `The domain ID of the user.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name of the scope.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The domain ID of the scope.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `The project name of the scope.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project ID of the scope.`,
				},
				resource.Attribute{
					Name:        "project_domain_name",
					Description: `The domain name of the project.`,
				},
				resource.Attribute{
					Name:        "project_domain_id",
					Description: `The domain ID of the project.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles in the current scope. See reference below.`,
				},
				resource.Attribute{
					Name:        "service_catalog",
					Description: `A list of service catalog entries returned with the token. The ` + "`" + `roles` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The ID of the role.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `The name of the role. The ` + "`" + `service_catalog` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the service.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `A list of endpoints for the service. The ` + "`" + `endpoints` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the endpoint.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the endpoint.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region ID of the endpoint.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `The interface of the endpoint.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the endpoint.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_endpoint_v3",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Endpoint.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Keystone client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the endpoint.`,
				},
				resource.Attribute{
					Name:        "endpoint_region",
					Description: `(Optional) The region the endpoint is assigned to. The ` + "`" + `region` + "`" + ` and ` + "`" + `endpoint_region` + "`" + ` can be different.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Optional) The service id this endpoint belongs to.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The service name of the endpoint.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Optional) The service type of the endpoint.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) The endpoint interface. Valid values are ` + "`" + `public` + "`" + `, ` + "`" + `internal` + "`" + `, and ` + "`" + `admin` + "`" + `. Default value is ` + "`" + `public` + "`" + ` ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found endpoint. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "endpoint_region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The endpoint URL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "endpoint_region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The endpoint URL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_group_v3",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the group.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain the group belongs to.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Keystone client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found group. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_project_v3",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain this project belongs to.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the project is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_domain",
					Description: `(Optional) Whether this project is a domain. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the project.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) The parent of this project.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags for the project. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found project. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the project.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the project is located in.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the project.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the project is located in.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_role_v3",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Role.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the role.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain the role belongs to.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Keystone client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found role. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_service_v3",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Keystone client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The service name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The service type.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The service status. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found service. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The service description.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The service description.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_user_v3",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack User.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_project_id",
					Description: `(Optional) The default project this user belongs to.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain this user belongs to.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the user is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(Optional) The identity provider ID of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the user.`,
				},
				resource.Attribute{
					Name:        "password_expires_at",
					Description: `(Optional) Query for expired passwords. See the [OpenStack API docs](https://developer.openstack.org/api-ref/identity/v3/#list-users) for more information on the query format.`,
				},
				resource.Attribute{
					Name:        "protocol_id",
					Description: `(Optional) The protocol ID of the user.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `(Optional) The unique ID of the user. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "default_project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password_expires_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the user is located in.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "default_project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password_expires_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the user is located in.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_images_image_ids_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Openstack Image IDs`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Glance client. A Glance client is needed to create an Image that can be used with a compute instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "member_status",
					Description: `(Optional) The status of the image. Must be one of "accepted", "pending", "rejected", or "all".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the image. Cannot be used simultaneously with ` + "`" + `name_regex` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) The regular expressian of the name of the image. Cannot be used simultaneously with ` + "`" + `name` + "`" + `. Unlike filtering by ` + "`" + `name` + "`" + ` the ` + "`" + `name_regex` + "`" + ` filtering does by client on the result of OpenStack search query.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) The owner (UUID) of the image.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) a map of key/value pairs to match an image with. All specified properties must be matched. Unlike other options filtering by ` + "`" + `properties` + "`" + ` does by client on the result of OpenStack search query.`,
				},
				resource.Attribute{
					Name:        "size_min",
					Description: `(Optional) The minimum size (in bytes) of the image to return.`,
				},
				resource.Attribute{
					Name:        "size_max",
					Description: `(Optional) The maximum size (in bytes) of the image to return.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sorts the response by one or more attribute and sort direction combinations. You can also set multiple sort keys and directions. Default direction is ` + "`" + `desc` + "`" + `. Use the comma (,) character to separate multiple values. For example expression ` + "`" + `sort = "name:asc,status"` + "`" + ` sorts ascending by name and descending by status. ` + "`" + `sort` + "`" + ` cannot be used simultaneously with ` + "`" + `sort_key` + "`" + `. If both are present in a configuration then only ` + "`" + `sort` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "sort_direction",
					Description: `(Optional) Order the results in either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. Can be applied only with ` + "`" + `sort_key` + "`" + `. Defaults to ` + "`" + `asc` + "`" + ``,
				},
				resource.Attribute{
					Name:        "sort_key",
					Description: `(Optional) Sort images based on a certain key. Defaults to ` + "`" + `name` + "`" + `. ` + "`" + `sort_key` + "`" + ` cannot be used simultaneously with ` + "`" + `sort` + "`" + `. If both are present in a configuration then only ` + "`" + `sort` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Search for images with a specific tag.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The visibility of the image. Must be one of "public", "private", "community", or "shared". Defaults to "private". ## Attributes Reference ` + "`" + `ids` + "`" + ` is set to the list of Openstack Image IDs.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_images_image_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Image.`,
			Description:      ``,
			Icon:             "image-gray.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Glance client. A Glance client is needed to create an Image that can be used with a compute instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the image.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) The owner (UUID) of the image.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) a map of key/value pairs to match an image with. All specified properties must be matched. Unlike other options filtering by ` + "`" + `properties` + "`" + ` does by client on the result of OpenStack search query. Filtering is applied if server responce contains at least 2 images. In case there is only one image the ` + "`" + `properties` + "`" + ` ignores.`,
				},
				resource.Attribute{
					Name:        "size_min",
					Description: `(Optional) The minimum size (in bytes) of the image to return.`,
				},
				resource.Attribute{
					Name:        "size_max",
					Description: `(Optional) The maximum size (in bytes) of the image to return.`,
				},
				resource.Attribute{
					Name:        "sort_direction",
					Description: `(Optional) Order the results in either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sort_key",
					Description: `(Optional) Sort images based on a certain key. Defaults to ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Search for images with a specific tag.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The visibility of the image. Must be one of "public", "private", "community", or "shared". Defaults to "private".`,
				},
				resource.Attribute{
					Name:        "hidden",
					Description: `(Optional) Whether or not the image is hidden from public list.`,
				},
				resource.Attribute{
					Name:        "member_status",
					Description: `(Optional) The status of the image. Must be one of "accepted", "pending", "rejected", or "all". ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found image. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `The minimum amount of disk space required to use the image.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `The minimum amount of ram required to use the image.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `Freeform information about the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `Whether or not the image is protected.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size of the image (in bytes).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags list of the image.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the image was last updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `The minimum amount of disk space required to use the image.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `The minimum amount of ram required to use the image.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `Freeform information about the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `Whether or not the image is protected.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size of the image (in bytes).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags list of the image.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the image was last updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_keymanager_container_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on a V1 Barbican container resource within OpenStack.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V1 KeyManager client. A KeyManager client is needed to fetch a container. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Container name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "container_ref",
					Description: `The container reference / where to find the container.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The container type.`,
				},
				resource.Attribute{
					Name:        "secret_refs",
					Description: `A set of dictionaries containing references to secrets. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The creator of the container.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the container.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the container was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the container was last updated.`,
				},
				resource.Attribute{
					Name:        "consumers",
					Description: `The list of the container consumers. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The list of ACLs assigned to a container. The ` + "`" + `read` + "`" + ` structure is described below. The ` + "`" + `secret_refs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `The secret reference / where to find the secret, URL. The ` + "`" + `consumers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the consumer.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The consumer URL. The ` + "`" + `acl` + "`" + ` ` + "`" + `read` + "`" + ` attribute supports:`,
				},
				resource.Attribute{
					Name:        "project_access",
					Description: `Whether the container is accessible project wide.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `The list of user IDs, which are allowed to access the container, when ` + "`" + `project_access` + "`" + ` is set to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the container ACL was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the container ACL was last updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "container_ref",
					Description: `The container reference / where to find the container.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The container type.`,
				},
				resource.Attribute{
					Name:        "secret_refs",
					Description: `A set of dictionaries containing references to secrets. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The creator of the container.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the container.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the container was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the container was last updated.`,
				},
				resource.Attribute{
					Name:        "consumers",
					Description: `The list of the container consumers. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The list of ACLs assigned to a container. The ` + "`" + `read` + "`" + ` structure is described below. The ` + "`" + `secret_refs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `The secret reference / where to find the secret, URL. The ` + "`" + `consumers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the consumer.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The consumer URL. The ` + "`" + `acl` + "`" + ` ` + "`" + `read` + "`" + ` attribute supports:`,
				},
				resource.Attribute{
					Name:        "project_access",
					Description: `Whether the container is accessible project wide.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `The list of user IDs, which are allowed to access the container, when ` + "`" + `project_access` + "`" + ` is set to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the container ACL was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the container ACL was last updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_keymanager_secret_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on a V1 Barbican secret resource within OpenStack.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V1 KeyManager client. A KeyManager client is needed to fetch a secret. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Secret name.`,
				},
				resource.Attribute{
					Name:        "bit_length",
					Description: `(Optional) The Secret bit length.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Optional) The Secret algorithm.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The Secret mode.`,
				},
				resource.Attribute{
					Name:        "secret_type",
					Description: `(Optional) The Secret type. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).`,
				},
				resource.Attribute{
					Name:        "acl_only",
					Description: `(Optional) Select the Secret with an ACL that contains the user. Project scope is ignored. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expiration_filter",
					Description: `(Optional) Date filter to select the Secret with expiration matching the specified criteria. See Date Filters below for more detail.`,
				},
				resource.Attribute{
					Name:        "created_at_filter",
					Description: `(Optional) Date filter to select the Secret with created matching the specified criteria. See Date Filters below for more detail.`,
				},
				resource.Attribute{
					Name:        "updated_at_filter",
					Description: `(Optional) Date filter to select the Secret with updated matching the specified criteria. See Date Filters below for more detail. ## Date Filters The values for the ` + "`" + `expiration_filter` + "`" + `, ` + "`" + `created_at_filter` + "`" + `, and ` + "`" + `updated_at_filter` + "`" + ` parameters are comma-separated lists of time stamps in RFC3339 format. The time stamps can be prefixed with any of these comparison operators:`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `The secret reference / where to find the secret.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bit_length",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "secret_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "acl_only",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expiration_filter",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at_filter",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "updated_at_filter",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "payload",
					Description: `The secret payload.`,
				},
				resource.Attribute{
					Name:        "payload_content_type",
					Description: `The Secret content type.`,
				},
				resource.Attribute{
					Name:        "payload_content_encoding",
					Description: `The Secret encoding.`,
				},
				resource.Attribute{
					Name:        "content_types",
					Description: `The map of the content types, assigned on the secret.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The creator of the secret.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the secret.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The date the secret will expire.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the secret was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the secret was last updated.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The map of metadata, assigned on the secret, which has been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The list of ACLs assigned to a secret. The ` + "`" + `read` + "`" + ` structure is described below. The ` + "`" + `acl` + "`" + ` ` + "`" + `read` + "`" + ` attribute supports:`,
				},
				resource.Attribute{
					Name:        "project_access",
					Description: `Whether the secret is accessible project wide.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `The list of user IDs, which are allowed to access the secret, when ` + "`" + `project_access` + "`" + ` is set to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the secret ACL was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the secret ACL was last updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_ref",
					Description: `The secret reference / where to find the secret.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bit_length",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "secret_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "acl_only",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expiration_filter",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at_filter",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "updated_at_filter",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "payload",
					Description: `The secret payload.`,
				},
				resource.Attribute{
					Name:        "payload_content_type",
					Description: `The Secret content type.`,
				},
				resource.Attribute{
					Name:        "payload_content_encoding",
					Description: `The Secret encoding.`,
				},
				resource.Attribute{
					Name:        "content_types",
					Description: `The map of the content types, assigned on the secret.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The creator of the secret.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the secret.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The date the secret will expire.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the secret was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the secret was last updated.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The map of metadata, assigned on the secret, which has been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The list of ACLs assigned to a secret. The ` + "`" + `read` + "`" + ` structure is described below. The ` + "`" + `acl` + "`" + ` ` + "`" + `read` + "`" + ` attribute supports:`,
				},
				resource.Attribute{
					Name:        "project_access",
					Description: `Whether the secret is accessible project wide.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `The list of user IDs, which are allowed to access the secret, when ` + "`" + `project_access` + "`" + ` is set to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the secret ACL was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the secret ACL was last updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_addressscope_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Address Scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve address-scopes. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the address-scope.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Indicates whether this address-scope is shared across all projects.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the address-scope. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found address-scope. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_floatingip_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Floating IP.`,
			Description:      ``,
			Icon:             "floatingip-gray.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve floating IP ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the floating IP.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) The IP address of the floating IP.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Optional) The name of the pool from which the floating IP belongs to.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) The ID of the port the floating IP is attached.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status of the floating IP (ACTIVE/DOWN).`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) The specific IP address of the internal port which should be associated with the floating IP.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of floating IP tags to filter.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the floating IP. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found floating IP. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `A set of string tags applied on the floating IP.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The floating IP DNS name. Available, when Neutron DNS extension is enabled.`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `The floating IP DNS domain. Available, when Neutron DNS extension is enabled.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "all_tags",
					Description: `A set of string tags applied on the floating IP.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The floating IP DNS name. Available, when Neutron DNS extension is enabled.`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `The floating IP DNS domain. Available, when Neutron DNS extension is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_network_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Network.`,
			Description:      ``,
			Icon:             "network-gray.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve networks ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the network.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the network.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `(Optional) The external routing facility of the network.`,
				},
				resource.Attribute{
					Name:        "matching_subnet_cidr",
					Description: `(Optional) The CIDR of a subnet within the network.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the network.`,
				},
				resource.Attribute{
					Name:        "availability_zone_hints",
					Description: `(Optional) The availability zone candidates for the network.`,
				},
				resource.Attribute{
					Name:        "transparent_vlan",
					Description: `(Optional) The VLAN transparent attribute for the network.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of network tags to filter.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The network MTU to filter. Available, when Neutron ` + "`" + `net-mtu` + "`" + ` extension is enabled. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found network. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `The administrative state of the network.`,
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
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Specifies whether the network resource can be accessed by any tenant or not.`,
				},
				resource.Attribute{
					Name:        "availability_zone_hints",
					Description: `The availability zone candidates for the network.`,
				},
				resource.Attribute{
					Name:        "transparent_vlan",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `The network DNS domain. Available, when Neutron DNS extension is enabled`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `A list of subnet IDs belonging to the network.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `The administrative state of the network.`,
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
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Specifies whether the network resource can be accessed by any tenant or not.`,
				},
				resource.Attribute{
					Name:        "availability_zone_hints",
					Description: `The availability zone candidates for the network.`,
				},
				resource.Attribute{
					Name:        "transparent_vlan",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `The network DNS domain. Available, when Neutron DNS extension is enabled`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `A list of subnet IDs belonging to the network.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_port_ids_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Openstack Port IDs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve port ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the port.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the port.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the port.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the port.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network the port belongs to.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `(Optional) The device owner of the port.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The MAC address of the port.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) The ID of the device the port belongs to.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) The port IP address filter.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the port.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) The list of port security group IDs to filter.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of port tags to filter.`,
				},
				resource.Attribute{
					Name:        "sort_key",
					Description: `(Optional) Sort ports based on a certain key. Defaults to none.`,
				},
				resource.Attribute{
					Name:        "sort_direction",
					Description: `(Optional) Order the results in either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. Defaults to none. ## Attributes Reference ` + "`" + `ids` + "`" + ` is set to the list of Openstack Port IDs.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_port_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information of an OpenStack Port.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve port ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the port.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) The ID of the port.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the port.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the port.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the port.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network the port belongs to.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `(Optional) The device owner of the port.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The MAC address of the port.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) The ID of the device the port belongs to.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) The port IP address filter.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the port.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) The list of port security group IDs to filter.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of port tags to filter.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `(Optional) The port DNS name to filter. Available, when Neutron DNS extension is enabled. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found port. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
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
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `An IP/MAC Address pair of additional IP addresses that can be active on this port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The set of security group IDs applied on the port.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the port.`,
				},
				resource.Attribute{
					Name:        "extra_dhcp_option",
					Description: `An extra DHCP option configured on the port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "binding",
					Description: `The port binding information. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_assignment",
					Description: `The list of maps representing port DNS assignments. The ` + "`" + `allowed_address_pairs` + "`" + ` attribute has fields below:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The additional IP address.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The additional MAC address. The ` + "`" + `extra_dhcp_option` + "`" + ` attribute has fields below:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DHCP option.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of the DHCP option.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `IP protocol version The ` + "`" + `binding` + "`" + ` attribute has fields below:`,
				},
				resource.Attribute{
					Name:        "host_id",
					Description: `The ID of the host, which has the allocatee port.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `A JSON string containing the binding profile information.`,
				},
				resource.Attribute{
					Name:        "vnic_type",
					Description: `VNIC type for the port.`,
				},
				resource.Attribute{
					Name:        "vif_details",
					Description: `A map of JSON strings containing additional details for this specific binding.`,
				},
				resource.Attribute{
					Name:        "vif_type",
					Description: `The VNIC type of the port binding.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
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
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `An IP/MAC Address pair of additional IP addresses that can be active on this port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The set of security group IDs applied on the port.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the port.`,
				},
				resource.Attribute{
					Name:        "extra_dhcp_option",
					Description: `An extra DHCP option configured on the port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "binding",
					Description: `The port binding information. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_assignment",
					Description: `The list of maps representing port DNS assignments. The ` + "`" + `allowed_address_pairs` + "`" + ` attribute has fields below:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The additional IP address.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The additional MAC address. The ` + "`" + `extra_dhcp_option` + "`" + ` attribute has fields below:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DHCP option.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of the DHCP option.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `IP protocol version The ` + "`" + `binding` + "`" + ` attribute has fields below:`,
				},
				resource.Attribute{
					Name:        "host_id",
					Description: `The ID of the host, which has the allocatee port.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `A JSON string containing the binding profile information.`,
				},
				resource.Attribute{
					Name:        "vnic_type",
					Description: `VNIC type for the port.`,
				},
				resource.Attribute{
					Name:        "vif_details",
					Description: `A map of JSON strings containing additional details for this specific binding.`,
				},
				resource.Attribute{
					Name:        "vif_type",
					Description: `The VNIC type of the port binding.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_qos_bandwidth_limit_rule_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack QoS Bandwidth limit rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `(Required) The QoS policy reference.`,
				},
				resource.Attribute{
					Name:        "max_kbps",
					Description: `(Optional) The maximum kilobits per second of a QoS bandwidth limit rule.`,
				},
				resource.Attribute{
					Name:        "max_burst_kbps",
					Description: `(Optional) The maximum burst size in kilobits of a QoS bandwidth limit rule.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) The direction of traffic. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ` + "`" + `qos_policy_id/bandwidth_limit_rule_id` + "`" + ` format of the found QoS bandwidth limit rule. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_kbps",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_burst_kbps",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_kbps",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_burst_kbps",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_qos_dscp_marking_rule_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack QoS DSCP marking rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `(Required) The QoS policy reference.`,
				},
				resource.Attribute{
					Name:        "dscp_mark",
					Description: `(Optional) The value of a DSCP mark. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ` + "`" + `qos_policy_id/dscp_marking_rule_id` + "`" + ` format of the found QoS DSCP marking rule. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dscp_mark",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dscp_mark",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_qos_minimum_bandwidth_rule_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack QoS minimum bandwidth rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `(Required) The QoS policy reference.`,
				},
				resource.Attribute{
					Name:        "min_kbps",
					Description: `(Optional) The value of a minimum kbps bandwidth. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ` + "`" + `qos_policy_id/minimum_bandwidth_rule_id` + "`" + ` format of the found QoS minimum bandwidth rule. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_kbps",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_kbps",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_qos_policy_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack QoS Policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to retrieve a QoS policy ID. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the QoS policy.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the QoS policy.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether this QoS policy is shared across all projects.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the QoS policy.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Whether the QoS policy is default policy or not.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of QoS policy tags to filter. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found QoS policy. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which QoS policy was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which QoS policy was created.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "revision_number",
					Description: `The revision number of the QoS policy.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the QoS policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which QoS policy was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which QoS policy was created.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "revision_number",
					Description: `The revision number of the QoS policy.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the QoS policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_quota_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on a NEtworking Quota of a project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Network client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project to retrieve the quota. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floatingip",
					Description: `The number of allowed floating ips.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The number of allowed networks.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The number of allowed ports.`,
				},
				resource.Attribute{
					Name:        "rbac_policy",
					Description: `The number of allowed rbac policies.`,
				},
				resource.Attribute{
					Name:        "router",
					Description: `The amount of allowed routers.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `The number of allowed security groups.`,
				},
				resource.Attribute{
					Name:        "security_group_rule",
					Description: `The number of allowed security group rules.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The number of allowed subnets.`,
				},
				resource.Attribute{
					Name:        "subnetpool-",
					Description: `The number of allowed subnet pools.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floatingip",
					Description: `The number of allowed floating ips.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The number of allowed networks.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The number of allowed ports.`,
				},
				resource.Attribute{
					Name:        "rbac_policy",
					Description: `The number of allowed rbac policies.`,
				},
				resource.Attribute{
					Name:        "router",
					Description: `The amount of allowed routers.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `The number of allowed security groups.`,
				},
				resource.Attribute{
					Name:        "security_group_rule",
					Description: `The number of allowed security group rules.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The number of allowed subnets.`,
				},
				resource.Attribute{
					Name:        "subnetpool-",
					Description: `The number of allowed subnet pools.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_router_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Floating IP.`,
			Description:      ``,
			Icon:             "router-gray.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve router ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Optional) The UUID of the router resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the router.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the router.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the router (must be "true" or "false" if provided).`,
				},
				resource.Attribute{
					Name:        "distributed",
					Description: `(Optional) Indicates whether or not to get a distributed router.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the router (ACTIVE/DOWN).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of router tags to filter.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the router. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found router. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `The value that points out if the Source NAT is enabled on the router.`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `The network UUID of an external gateway for the router.`,
				},
				resource.Attribute{
					Name:        "availability_zone_hints",
					Description: `The availability zone that is used to make router resources highly available.`,
				},
				resource.Attribute{
					Name:        "external_fixed_ip",
					Description: `The external fixed IPs of the router. The ` + "`" + `external_fixed_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address to set on the router.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the router.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "enable_snat",
					Description: `The value that points out if the Source NAT is enabled on the router.`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `The network UUID of an external gateway for the router.`,
				},
				resource.Attribute{
					Name:        "availability_zone_hints",
					Description: `The availability zone that is used to make router resources highly available.`,
				},
				resource.Attribute{
					Name:        "external_fixed_ip",
					Description: `The external fixed IPs of the router. The ` + "`" + `external_fixed_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address to set on the router.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the router.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_secgroup_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Security Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve security groups ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "secgroup_id",
					Description: `(Optional) The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description the the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of security group tags to filter.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found security group. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the security group.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the security group.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_subnet_ids_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Openstack Subnet IDs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve subnet ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp_enabled",
					Description: `(Optional) If the subnet has DHCP enabled.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network the subnet belongs to.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the subnet.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) The IP version of the subnet (either 4 or 6).`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) The IP of the subnet's gateway.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) The CIDR of the subnet.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_mode",
					Description: `(Optional) The IPv6 address mode. Valid values are ` + "`" + `dhcpv6-stateful` + "`" + `, ` + "`" + `dhcpv6-stateless` + "`" + `, or ` + "`" + `slaac` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv6_ra_mode",
					Description: `(Optional) The IPv6 Router Advertisement mode. Valid values are ` + "`" + `dhcpv6-stateful` + "`" + `, ` + "`" + `dhcpv6-stateless` + "`" + `, or ` + "`" + `slaac` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnetpool_id",
					Description: `(Optional) The ID of the subnetpool associated with the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of subnet tags to filter.`,
				},
				resource.Attribute{
					Name:        "sort_key",
					Description: `(Optional) Sort subnets based on a certain key. Defaults to none.`,
				},
				resource.Attribute{
					Name:        "sort_direction",
					Description: `(Optional) Order the results in either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. Defaults to none. ## Attributes Reference ` + "`" + `ids` + "`" + ` is set to the list of Openstack Subnet IDs.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_subnet_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Subnet.`,
			Description:      ``,
			Icon:             "network-gray.svg",
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve subnet ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp_enabled",
					Description: `(Optional) If the subnet has DHCP enabled.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network the subnet belongs to.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the subnet.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) The IP version of the subnet (either 4 or 6).`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) The IP of the subnet's gateway.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) The CIDR of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_mode",
					Description: `(Optional) The IPv6 address mode. Valid values are ` + "`" + `dhcpv6-stateful` + "`" + `, ` + "`" + `dhcpv6-stateless` + "`" + `, or ` + "`" + `slaac` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv6_ra_mode",
					Description: `(Optional) The IPv6 Router Advertisement mode. Valid values are ` + "`" + `dhcpv6-stateful` + "`" + `, ` + "`" + `dhcpv6-stateless` + "`" + `, or ` + "`" + `slaac` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnetpool_id",
					Description: `(Optional) The ID of the subnetpool associated with the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of subnet tags to filter. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found subnet. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `Allocation pools of the subnet.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `Whether the subnet has DHCP enabled or not.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `DNS Nameservers of the subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `Host Routes of the subnet.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `A set of string tags applied on the subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `Allocation pools of the subnet.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `Whether the subnet has DHCP enabled or not.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `DNS Nameservers of the subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `Host Routes of the subnet.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `A set of string tags applied on the subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_subnetpool_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OpenStack Subnetpool.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to retrieve a subnetpool id. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subnetpool.`,
				},
				resource.Attribute{
					Name:        "default_quota",
					Description: `(Optional) The per-project quota on the prefix space that can be allocated from the subnetpool for project subnets.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the subnetpool.`,
				},
				resource.Attribute{
					Name:        "prefixes",
					Description: `(Optional) A list of subnet prefixes that are assigned to the subnetpool.`,
				},
				resource.Attribute{
					Name:        "default_prefixlen",
					Description: `(Optional) The size of the subnetpool default prefix length.`,
				},
				resource.Attribute{
					Name:        "min_prefixlen",
					Description: `(Optional) The size of the subnetpool min prefix length.`,
				},
				resource.Attribute{
					Name:        "max_prefixlen",
					Description: `(Optional) The size of the subnetpool max prefix length.`,
				},
				resource.Attribute{
					Name:        "address_scope_id",
					Description: `(Optional) The Neutron address scope that subnetpools is assigned to.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `The IP protocol version.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether this subnetpool is shared across all projects.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the subnetpool.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Whether the subnetpool is default subnetpool or not.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of subnetpool tags to filter. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found subnetpool. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_quota",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which subnetpool was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which subnetpool was created.`,
				},
				resource.Attribute{
					Name:        "prefixes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address_scope_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "revision_number",
					Description: `The revision number of the subnetpool.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the subnetpool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_quota",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which subnetpool was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which subnetpool was created.`,
				},
				resource.Attribute{
					Name:        "prefixes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address_scope_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "revision_number",
					Description: `The revision number of the subnetpool.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the subnetpool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_trunk_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information of an OpenStack Trunk.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve trunk ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the trunk.`,
				},
				resource.Attribute{
					Name:        "trunk_id",
					Description: `(Optional) The ID of the trunk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the trunk.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the trunk.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) The ID of the trunk parent port.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the trunk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the trunk.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of trunk tags to filter. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found trunk. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the trunk.`,
				},
				resource.Attribute{
					Name:        "sub_port",
					Description: `The set of the trunk subports. The structure of each subport is described below. The ` + "`" + `sub_port` + "`" + ` attribute has fields below:`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `The ID of the trunk subport.`,
				},
				resource.Attribute{
					Name:        "segmentation_type",
					Description: `The segmenation tecnology used, e.g., "vlan".`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `The numeric id of the subport segment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "all_tags",
					Description: `The set of string tags applied on the trunk.`,
				},
				resource.Attribute{
					Name:        "sub_port",
					Description: `The set of the trunk subports. The structure of each subport is described below. The ` + "`" + `sub_port` + "`" + ` attribute has fields below:`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `The ID of the trunk subport.`,
				},
				resource.Attribute{
					Name:        "segmentation_type",
					Description: `The segmenation tecnology used, e.g., "vlan".`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `The numeric id of the subport segment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_sharedfilesystem_availability_zones_v2",
			Category:         "Data Sources",
			ShortDescription: `Get a list of Shared File System availability zones from OpenStack`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Shared File System client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to hash of the returned zone list. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_sharedfilesystem_share_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Shared File System share.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the share.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the share.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the share.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The UUID of the share's base snapshot.`,
				},
				resource.Attribute{
					Name:        "share_network_id",
					Description: `(Optional) The UUID of the share's share network.`,
				},
				resource.Attribute{
					Name:        "export_location_path",
					Description: `(Optional) The export location path of the share. Available since Manila API version 2.35.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) One or more metadata key and value pairs as a dictionary of strings.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) A share status filter. A valid value is ` + "`" + `creating` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `available` + "`" + `, ` + "`" + `deleting` + "`" + `, ` + "`" + `error_deleting` + "`" + `, ` + "`" + `manage_starting` + "`" + `, ` + "`" + `manage_error` + "`" + `, ` + "`" + `unmanage_starting` + "`" + `, ` + "`" + `unmanage_error` + "`" + `, ` + "`" + `unmanaged` + "`" + `, ` + "`" + `extending` + "`" + `, ` + "`" + `extending_error` + "`" + `, ` + "`" + `shrinking` + "`" + `, ` + "`" + `shrinking_error` + "`" + `, or ` + "`" + `shrinking_possible_data_loss_error` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) The level of visibility for the share. length. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found share. In addition, the following attributes are exported:`,
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
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "share_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "export_location_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to obtain the V2 Shared File System client.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The share availability zone.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `The share protocol.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The share size, in GBs.`,
				},
				resource.Attribute{
					Name:        "export_locations",
					Description: `A list of export locations. For example, when a share server has more than one network interface, it can have multiple export locations.`,
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
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "share_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "export_location_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to obtain the V2 Shared File System client.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The share availability zone.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `The share protocol.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The share size, in GBs.`,
				},
				resource.Attribute{
					Name:        "export_locations",
					Description: `A list of export locations. For example, when a share server has more than one network interface, it can have multiple export locations.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_sharedfilesystem_sharenetwork_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Shared File System share network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Shared File System client. A Shared File System client is needed to read a share network. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the share network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description of the share network.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the share network.`,
				},
				resource.Attribute{
					Name:        "neutron_net_id",
					Description: `(Optional) The neutron network UUID of the share network.`,
				},
				resource.Attribute{
					Name:        "neutron_subnet_id",
					Description: `(Optional) The neutron subnet UUID of the share network.`,
				},
				resource.Attribute{
					Name:        "security_service_id",
					Description: `(Optional) The security service IDs associated with the share network.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional) The share network type. Can either be VLAN, VXLAN, GRE, or flat.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `(Optional) The share network segmentation ID.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) The share network CIDR.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) The IP version of the share network. Can either be 4 or 6. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found share network . In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The owner of the Share Network.`,
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
					Name:        "neutron_net_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "neutron_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_service_ids",
					Description: `The list of security service IDs associated with the share network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The owner of the Share Network.`,
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
					Name:        "neutron_net_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "neutron_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_service_ids",
					Description: `The list of security service IDs associated with the share network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_sharedfilesystem_snapshot_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Shared File System snapshot.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Shared File System client.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) A snapshot status filter. A valid value is ` + "`" + `available` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `creating` + "`" + `, ` + "`" + `deleting` + "`" + `, ` + "`" + `manage_starting` + "`" + `, ` + "`" + `manage_error` + "`" + `, ` + "`" + `unmanage_starting` + "`" + `, ` + "`" + `unmanage_error` + "`" + ` or ` + "`" + `error_deleting` + "`" + `. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found snapshot. In addition, the following attributes are exported:`,
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
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The snapshot size, in GBs.`,
				},
				resource.Attribute{
					Name:        "share_id",
					Description: `The UUID of the source share that was used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `The file system protocol of a share snapshot.`,
				},
				resource.Attribute{
					Name:        "share_size",
					Description: `The share snapshot size, in GBs.`,
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
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The snapshot size, in GBs.`,
				},
				resource.Attribute{
					Name:        "share_id",
					Description: `The UUID of the source share that was used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `The file system protocol of a share snapshot.`,
				},
				resource.Attribute{
					Name:        "share_size",
					Description: `The share snapshot size, in GBs.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"openstack_blockstorage_availability_zones_v3":       0,
		"openstack_blockstorage_quotaset_v3":                 1,
		"openstack_blockstorage_snapshot_v2":                 2,
		"openstack_blockstorage_snapshot_v3":                 3,
		"openstack_blockstorage_volume_v2":                   4,
		"openstack_blockstorage_volume_v3":                   5,
		"openstack_compute_availability_zones_v2":            6,
		"openstack_compute_flavor_v2":                        7,
		"openstack_compute_keypair_v2":                       8,
		"openstack_compute_limits_v2":                        9,
		"openstack_compute_quotaset_v2":                      10,
		"openstack_containerinfra_cluster_v1":                11,
		"openstack_containerinfra_clustertemplate_v1":        12,
		"openstack_containerinfra_nodegroup_v1":              13,
		"openstack_dns_zone_v2":                              14,
		"openstack_fw_policy_v1":                             15,
		"openstack_identity_auth_scope_v3":                   16,
		"openstack_identity_endpoint_v3":                     17,
		"openstack_identity_group_v3":                        18,
		"openstack_identity_project_v3":                      19,
		"openstack_identity_role_v3":                         20,
		"openstack_identity_service_v3":                      21,
		"openstack_identity_user_v3":                         22,
		"openstack_images_image_ids_v2":                      23,
		"openstack_images_image_v2":                          24,
		"openstack_keymanager_container_v1":                  25,
		"openstack_keymanager_secret_v1":                     26,
		"openstack_networking_addressscope_v2":               27,
		"openstack_networking_floatingip_v2":                 28,
		"openstack_networking_network_v2":                    29,
		"openstack_networking_port_ids_v2":                   30,
		"openstack_networking_port_v2":                       31,
		"openstack_networking_qos_bandwidth_limit_rule_v2":   32,
		"openstack_networking_qos_dscp_marking_rule_v2":      33,
		"openstack_networking_qos_minimum_bandwidth_rule_v2": 34,
		"openstack_networking_qos_policy_v2":                 35,
		"openstack_networking_quota_v2":                      36,
		"openstack_networking_router_v2":                     37,
		"openstack_networking_secgroup_v2":                   38,
		"openstack_networking_subnet_ids_v2":                 39,
		"openstack_networking_subnet_v2":                     40,
		"openstack_networking_subnetpool_v2":                 41,
		"openstack_networking_trunk_v2":                      42,
		"openstack_sharedfilesystem_availability_zones_v2":   43,
		"openstack_sharedfilesystem_share_v2":                44,
		"openstack_sharedfilesystem_sharenetwork_v2":         45,
		"openstack_sharedfilesystem_snapshot_v2":             46,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
