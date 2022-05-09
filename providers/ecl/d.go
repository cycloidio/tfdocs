package ecl

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ecl_baremetal_availability_zone_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Baremetal Availability Zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Optional) The name of the availability zone. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the zone_name of the found availability zone. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_baremetal_flavor_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Baremetal Flavor.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the flavor.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `(Optional) The exact amount of RAM (in megabytes).`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `(Optional) The amount of VCPUs.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional) The exact amount of disk (in gigabytes). ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the name of the found flavor. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_baremetal_keypair_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Baremetal Keypair.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The public_key of the keypair.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `(Optional) The fingerprint of the keypair. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the name of the found availability zone. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_compute_flavor_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Flavor.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the flavor.`,
				},
				resource.Attribute{
					Name:        "min_ram",
					Description: `(Optional) The minimum amount of RAM (in megabytes).`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `(Optional) The exact amount of RAM (in megabytes).`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `(Optional) The amount of VCPUs.`,
				},
				resource.Attribute{
					Name:        "min_disk",
					Description: `(Optional) The minimum amount of disk (in gigabytes).`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional) The exact amount of disk (in gigabytes).`,
				},
				resource.Attribute{
					Name:        "swap",
					Description: `(Optional) The amount of swap (in gigabytes).`,
				},
				resource.Attribute{
					Name:        "rx_tx_factor",
					Description: `(Optional) The ` + "`" + `rx_tx_factor` + "`" + ` of the flavor. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found flavor. In addition, the following attributes are exported:`,
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
					Name:        "min_ram",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_disk",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "swap",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rx_tx_factor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Whether the flavor is public or private.`,
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
					Name:        "min_ram",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_disk",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "swap",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rx_tx_factor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Whether the flavor is public or private.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_compute_keypair_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Keypair.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
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
					Name:        "public_key",
					Description: `The OpenSSH-formatted public key of the keypair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_dns_zone_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) Domain name of the zone.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) DNS Name for the zone.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Optional) ID for the pool hosting this zone.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID for the project(tenant) that owns the zone.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) e-mail for the zone. Used in SOA records for the zone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the zone.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of the zone.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of zone. PRIMARY is controlled by ECL2.0 DNS, SECONDARY zones are slaved from another DNS Server. Defaults to PRIMARY.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL (Time to Live) for the zone.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the zone.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Optional) Current serial number for the zone.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Optional) Date / Time when zone was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `(Optional) Date / Time when zone last updated.`,
				},
				resource.Attribute{
					Name:        "transferred_at",
					Description: `(Optional) For secondary zones. The last time an update was retrieved from the master servers.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `(Optional) For secondary zones. The servers to slave from to get DNS information. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
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
					Name:        "type",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns zero.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns 1.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns zero.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transferred_at",
					Description: `(Optional) See Argument Reference above. This parameter is not currently supported. It always returns null.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Optional) See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
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
					Name:        "type",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns zero.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns 1.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns zero.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transferred_at",
					Description: `(Optional) See Argument Reference above. This parameter is not currently supported. It always returns null.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Optional) See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `See Argument Reference above. This parameter is not currently supported. It always returns an empty string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_imagestorages_image_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Image.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "member_status",
					Description: `(Optional) Only show images with the specified member status. Must be one of "queued", "saving", "active", "killed", "deleted", "pending_delete".`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the image as a string.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) Shows images shared with me by the specified owner, where the owner is indicated by project ID.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) a map of key/value pairs to match an image with. All specified properties must be matched.`,
				},
				resource.Attribute{
					Name:        "size_max",
					Description: `(Optional) Value of the maximum size of the image in bytes.`,
				},
				resource.Attribute{
					Name:        "size_min",
					Description: `(Optional) Value of the minimum size of the image in bytes.`,
				},
				resource.Attribute{
					Name:        "sort_direction",
					Description: `(Optional) Sort direction. Must be one of "desc", "asc".`,
				},
				resource.Attribute{
					Name:        "sort_key",
					Description: `(Optional) Sort key.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Image tag.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) Image visibility. Must be one of "public", "private", "shared". ## Attributes Reference The following attributes are exported: ` + "`" + `id` + "`" + ` is set to the ID of the found image. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `md5 hash of image contents.`,
				},
				resource.Attribute{
					Name:        "container_format",
					Description: `Format of the container.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time of image registration.`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `Format of the disk.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `URL for the virtual machine image file.`,
				},
				resource.Attribute{
					Name:        "member_status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The location metadata.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `Amount of disk space (in GB) required to boot image.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `Amount of ram (in MB) required to boot image.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peroperties",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `If true, image will not be deletable.`,
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
					Name:        "size_max",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size_min",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sort_direction",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sort_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date and time of the last image modification.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `md5 hash of image contents.`,
				},
				resource.Attribute{
					Name:        "container_format",
					Description: `Format of the container.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time of image registration.`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `Format of the disk.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `URL for the virtual machine image file.`,
				},
				resource.Attribute{
					Name:        "member_status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The location metadata.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `Amount of disk space (in GB) required to boot image.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `Amount of ram (in MB) required to boot image.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peroperties",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `If true, image will not be deletable.`,
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
					Name:        "size_max",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size_min",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sort_direction",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sort_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date and time of the last image modification.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_common_function_gateway_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Common Function Gateway.`,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Optional) Common Function Pool instantiated by this Gateway.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) ID of automatically created network connected to this Common Function Gateway.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of automatically created subnet connected to this Common Function Gateway (using link-local address).`,
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
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_common_function_pool_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Common Function Pool.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Common Function Pool resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Common Function Pool resource.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique ID of the Common Function Pool resource. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_fic_gateway_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud FIC Gateway.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the FIC Gateway resource.`,
				},
				resource.Attribute{
					Name:        "fic_service_id",
					Description: `(Optional) FIC Service ID of the FIC Gateway resource.`,
				},
				resource.Attribute{
					Name:        "fic_gateway_id",
					Description: `(Optional) Unique ID of the FIC Gateway resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the FIC Gateway resource.`,
				},
				resource.Attribute{
					Name:        "qos_option_id",
					Description: `(Optional) QoS Option ID of the FIC Gateway resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of the FIC Gateway resource.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID). ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found fic gateway. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fic_service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fic_gateway_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_option_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fic_service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fic_gateway_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_option_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_gateway_interface_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Gateway interface.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "aws_gw_id",
					Description: `(Optional) AWS Gateway to which this port is connected.`,
				},
				resource.Attribute{
					Name:        "azure_gw_id",
					Description: `(Optional) Azure Gateway to which this port is connected.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Gateway Interface resource.`,
				},
				resource.Attribute{
					Name:        "gateway_interface_id",
					Description: `(Optional) Unique ID of the Gateway Interface resource.`,
				},
				resource.Attribute{
					Name:        "gcp_gw_id",
					Description: `(Optional) GCP Gateway to which this port is connected.`,
				},
				resource.Attribute{
					Name:        "gw_vipv4",
					Description: `(Optional) IP version 4 address to be assigned virtual router on VRRP.`,
				},
				resource.Attribute{
					Name:        "gw_vipv6",
					Description: `(Optional) IP version 6 address to be assigned virtual router on VRRP.`,
				},
				resource.Attribute{
					Name:        "interdc_gw_id",
					Description: `(Optional) Inter DC Gateway to which this port is connected.`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `(Optional) Internet GW to which this port is connected.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Gateway Interface resource.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) Netmask for IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) Network connected to this interface.`,
				},
				resource.Attribute{
					Name:        "primary_ipv4",
					Description: `(Optional) IP version 4 address to be assigned to primary device on VRRP.`,
				},
				resource.Attribute{
					Name:        "primary_ipv6",
					Description: `(Optional) IP version 6 address to be assigned to primary device on VRRP.`,
				},
				resource.Attribute{
					Name:        "secondary_ipv4",
					Description: `(Optional) IP version 4 address to be assigned to secondary device on VRRP.`,
				},
				resource.Attribute{
					Name:        "secondary_ipv6",
					Description: `(Optional) IP version 6 address to be assigned to secondary device on VRRP.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Optional) Service type for this interface. Must be one of "aws", "azure", "gcp", "interdc", "internet" and "vpn".`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The Gateway Interface status.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID).`,
				},
				resource.Attribute{
					Name:        "vpn_gw_id",
					Description: `(Optional) VPN Gateway to which this port is connected.`,
				},
				resource.Attribute{
					Name:        "vrid",
					Description: `(Optional) VRRP Group ID for this GW Interface. ## Attributes Reference The following attributes are exported: ` + "`" + `id` + "`" + ` is set to the ID of the found gateway interface. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "aws_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "azure_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above .`,
				},
				resource.Attribute{
					Name:        "gcp_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gw_vipv4",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gw_vipv6",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interdc_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "primary_ipv4",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "primary_ipv6",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "secondary_ipv4",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "secondary_ipv6",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpn_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vrid",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "aws_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "azure_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above .`,
				},
				resource.Attribute{
					Name:        "gcp_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gw_vipv4",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gw_vipv6",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interdc_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "primary_ipv4",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "primary_ipv6",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "secondary_ipv4",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "secondary_ipv6",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpn_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vrid",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_internet_gateway_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Internet gateway.`,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "internet_gateway_id",
					Description: `(Optional) Unique ID of the Internet Gateway resource.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `(Optional) Internet Service instantiated by this Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Internet Gateway resource.`,
				},
				resource.Attribute{
					Name:        "qos_option_id",
					Description: `(Optional) Quality of Service options selected for this Gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The Internet Gateway status.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID). ## Attributes Reference The following attributes are exported: ` + "`" + `id` + "`" + ` is set to the ID of the found internet gateway. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_option_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_option_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_load_balancer_plan_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Load Balancer Plan.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Load Balancer Plan.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the Load Balancer Plan is enable or not.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique ID of the Load Balancer Plan.`,
				},
				resource.Attribute{
					Name:        "maximum_syslog_servers",
					Description: `(Optional) Maximum number of syslog servers.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `(Optional) Model of load balancer. The ` + "`" + `model` + "`" + ` object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Load Balancer Plan.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `(Optional) Load Balancer Type.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version name. The ` + "`" + `model` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "edition",
					Description: `(Optional) Edition of Load Balancer Plan.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Bandwidth of Load Balancer Plan. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found Load Balancer Plan. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maximum_syslog_servers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maximum_syslog_servers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_network_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the network.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network.`,
				},
				resource.Attribute{
					Name:        "matching_subnet_cidr",
					Description: `(Optional) The CIDR of a subnet within the network.`,
				},
				resource.Attribute{
					Name:        "plane",
					Description: `(Optional) The plane of the network. Allowed values are "data" and "storage".`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `The administrative state of the network.`,
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
					Name:        "network_id",
					Description: `ID of network.`,
				},
				resource.Attribute{
					Name:        "matching_subnet_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "plane",
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
					Description: `The subnets of the network.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The network tags.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `The administrative state of the network.`,
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
					Name:        "network_id",
					Description: `ID of network.`,
				},
				resource.Attribute{
					Name:        "matching_subnet_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "plane",
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
					Description: `The subnets of the network.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The network tags.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_port_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Port.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Port description.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) The Id of device (i.e physical port id for bare-metal).`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `(Optional) The name of the port owner`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The MAC address of the port.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Port name.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of network this port belongs to.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) Port unique id.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `(Optional) The segmentation ID used for this port (i.e. for vlan type it is vlan tag)`,
				},
				resource.Attribute{
					Name:        "segmentation_type",
					Description: `(Optional) The segmentation type used for this port (i.e. vlan) ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found port. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `The administrative state of the port.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `An IP/MAC Address pair of additional IP addresses that can be active on this port. The structure is described below.`,
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
					Description: `List of the port IP address`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "managed_by_service",
					Description: `Set to true if only admin can modify it. Normal user has only read access.`,
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
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "segmentation_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the port.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Port tags.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The owner name of port.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `The administrative state of the port.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `An IP/MAC Address pair of additional IP addresses that can be active on this port. The structure is described below.`,
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
					Description: `List of the port IP address`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "managed_by_service",
					Description: `Set to true if only admin can modify it. Normal user has only read access.`,
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
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "segmentation_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the port.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Port tags.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The owner name of port.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_public_ip_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Public ip.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) The IP address of the block (assigned automatically).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Public IPs.`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `(Optional) Unique ID of the Public IPs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Public IPs.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `(Optional) Unique ID of the Public IPs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Public IP status.`,
				},
				resource.Attribute{
					Name:        "submask_length",
					Description: `(Optional) Specifies the size of the block by the length of its subnetwork mask length.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID). ## Attributes Reference The following attributes are exported: ` + "`" + `id` + "`" + ` is set to the ID of the found public ip. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
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
					Name:        "internet_gw_id",
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
					Name:        "submask_length",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
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
					Name:        "internet_gw_id",
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
					Name:        "submask_length",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_static_route_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Static route.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "aws_gw_id",
					Description: `(Optional) AWS Gateway on which this static route will be set.`,
				},
				resource.Attribute{
					Name:        "azure_gw_id",
					Description: `(Optional) Azure Gateway on which this static route will be set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Static Route resource.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) CIDR this static route points to.`,
				},
				resource.Attribute{
					Name:        "gcp_gw_id",
					Description: `(Optional) GCP Gateway on which this static route will be set.`,
				},
				resource.Attribute{
					Name:        "interdc_gw_id",
					Description: `(Optional) Inter DC Gateway on which this static route will be set.`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `(Optional) Internet Gateway on which this static route will be set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Static Route resource.`,
				},
				resource.Attribute{
					Name:        "nexthop",
					Description: `(Optional) Next Hop address for specified CIDR.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Optional) Service type for this route. Must be one of "aws", "azure", "gcp", "interdc", "internet" and "vpn".`,
				},
				resource.Attribute{
					Name:        "static_route_id",
					Description: `(Optional) Unique ID of the Static Route resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Static Route status.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID).`,
				},
				resource.Attribute{
					Name:        "vpn_gw_id",
					Description: `(Optional) VPN Gateway on which this static route will be set. ## Attributes Reference The following attributes are exported: ` + "`" + `id` + "`" + ` is set to the ID of the found static route. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "aws_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "azure_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above .`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `See Argument Reference above .`,
				},
				resource.Attribute{
					Name:        "gcp_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interdc_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nexthop",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpn_gw_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "aws_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "azure_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above .`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `See Argument Reference above .`,
				},
				resource.Attribute{
					Name:        "gcp_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interdc_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internet_gw_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nexthop",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpn_gw_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_network_subnet_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) CIDR representing IP range for this subnet, based on IP version. You can omit this option if you are creating a subnet from a subnet pool.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Subnet description.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Default gateway used by devices in this subnet. Leaving this blank and not setting ` + "`" + `no_gateway` + "`" + ` will cause a default gateway of ` + "`" + `.1` + "`" + ` to be used. Changing this updates the gateway IP of the existing subnet.`,
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
					Name:        "region",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of subnet. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found subnet. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `An array of sub-ranges of CIDR available for dynamic allocation to ports.`,
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
					Description: `List of subnet dns name servers.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `The administrative state of the network.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `An array of routes that should be used by devices with IPs from this subnet`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `IP version. In Enterprise Cloud service this parameter is fixed as 4.`,
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
					Name:        "ntp_servers",
					Description: `List of ntp servers.`,
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
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Subnet tags.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The owner of the subnet. Required if admin wants to create a subnet for another tenant. Changing this creates a new subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `An array of sub-ranges of CIDR available for dynamic allocation to ports.`,
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
					Description: `List of subnet dns name servers.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `The administrative state of the network.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `An array of routes that should be used by devices with IPs from this subnet`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `IP version. In Enterprise Cloud service this parameter is fixed as 4.`,
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
					Name:        "ntp_servers",
					Description: `List of ntp servers.`,
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
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Subnet tags.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The owner of the subnet. Required if admin wants to create a subnet for another tenant. Changing this creates a new subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_sss_tenant_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Tenant.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_name",
					Description: `(Required) Name of new tenant. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for this tenant.`,
				},
				resource.Attribute{
					Name:        "tenant_region",
					Description: `Region this tenant belongs to.`,
				},
				resource.Attribute{
					Name:        "contract_id",
					Description: `Contract which new tenant belongs to.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Tenant created time.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `ID of the tenant.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description for this tenant.`,
				},
				resource.Attribute{
					Name:        "tenant_region",
					Description: `Region this tenant belongs to.`,
				},
				resource.Attribute{
					Name:        "contract_id",
					Description: `Contract which new tenant belongs to.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Tenant created time.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `ID of the tenant.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_storage_virtualstorage_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Virtual Storage.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_storage_id",
					Description: `(Optional) ID of Virtual Storage.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of Virtual Storage. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Virtual Storage.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID(UUID) for network to be connected to the Virtual Storage.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID(UUID) for subnet to be connected to the Virtual Storage.`,
				},
				resource.Attribute{
					Name:        "volume_type_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_addr_pool",
					Description: `IP address pool which specifies IP address range used by the Virtual Storage. The ip_addr_pool structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `List of static routes to be set to this Virtual Storage. The host_routes structure is documented below.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of Virtual Storage.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `ID(UUID) for network to be connected to the Virtual Storage.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID(UUID) for subnet to be connected to the Virtual Storage.`,
				},
				resource.Attribute{
					Name:        "volume_type_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_addr_pool",
					Description: `IP address pool which specifies IP address range used by the Virtual Storage. The ip_addr_pool structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `List of static routes to be set to this Virtual Storage. The host_routes structure is documented below.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_storage_volume_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Volume.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Optional) ID of Volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of Volume. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of volume in Gigabyte.`,
				},
				resource.Attribute{
					Name:        "iops_per_gb",
					Description: `Provisioned IOPS/GB for volume.`,
				},
				resource.Attribute{
					Name:        "throughput",
					Description: `Throughput for volume.`,
				},
				resource.Attribute{
					Name:        "initiator_iqns",
					Description: `List of initiator IQN who can access to this volume.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of volume.`,
				},
				resource.Attribute{
					Name:        "virtual_storage_id",
					Description: `Virtual Storage ID (UUID) which this volume belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of Volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of volume in Gigabyte.`,
				},
				resource.Attribute{
					Name:        "iops_per_gb",
					Description: `Provisioned IOPS/GB for volume.`,
				},
				resource.Attribute{
					Name:        "throughput",
					Description: `Throughput for volume.`,
				},
				resource.Attribute{
					Name:        "initiator_iqns",
					Description: `List of initiator IQN who can access to this volume.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of volume.`,
				},
				resource.Attribute{
					Name:        "virtual_storage_id",
					Description: `Virtual Storage ID (UUID) which this volume belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_storage_volumetype_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Cloud Volume Type.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "volume_type_id",
					Description: `(Optional) ID of Volume Type.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of Volume Type. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "extra_specs",
					Description: `Includes available_volume_size, and available_iops_per_gb or available_throughput. The extra_specs structure is documented below. The ` + "`" + `extra_specs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "available_volume_size",
					Description: `List of available volume sizes for the volume type.`,
				},
				resource.Attribute{
					Name:        "available_iops_per_gb",
					Description: `List of available IOPS/GB values for the volume type.`,
				},
				resource.Attribute{
					Name:        "available_throughput",
					Description: `List of available throughput (MByte/s) values for the volume type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "extra_specs",
					Description: `Includes available_volume_size, and available_iops_per_gb or available_throughput. The extra_specs structure is documented below. The ` + "`" + `extra_specs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "available_volume_size",
					Description: `List of available volume sizes for the volume type.`,
				},
				resource.Attribute{
					Name:        "available_iops_per_gb",
					Description: `List of available IOPS/GB values for the volume type.`,
				},
				resource.Attribute{
					Name:        "available_throughput",
					Description: `List of available throughput (MByte/s) values for the volume type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ecl_vna_appliance_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Enterprise Virtual Network Appliance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Virtual Network Appliance.`,
				},
				resource.Attribute{
					Name:        "virtual_network_appliance_id",
					Description: `(Optional) ID of the Virtual Network Appliance Plan.`,
				},
				resource.Attribute{
					Name:        "appliance_type",
					Description: `(Optional) Appliance type of Virtual Network Appliance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Virtual Network Appliance.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) vailability Zone, this can be referred to using Virtual Server (Nova)'s list availability zones.`,
				},
				resource.Attribute{
					Name:        "os_monitoring_status",
					Description: `(Optional) OS Monitoring Status.`,
				},
				resource.Attribute{
					Name:        "os_login_status",
					Description: `(Optional) OS Login Status.`,
				},
				resource.Attribute{
					Name:        "vm_status",
					Description: `(Optional) VM Status.`,
				},
				resource.Attribute{
					Name:        "operation_status",
					Description: `(Optional) Operation Status.`,
				},
				resource.Attribute{
					Name:        "virtual_network_appliance_plan_id",
					Description: `(Optional) ID of the Virtual Network Appliance Plan.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Tenant ID of the owner (UUID).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the Virtual Network Appliance. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "default_gateway",
					Description: `IP address of default gateway.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "virtual_network_appliance_plan_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/name",
					Description: `Name of the interface.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/description",
					Description: `Description of the interface.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/network_id",
					Description: `The ID of network this interface belongs to.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/tags",
					Description: `Tags of the interface.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_fixed_ips/ip_address",
					Description: `The IP address assign to interface within subnet.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_fixed_ips/subnet_id",
					Description: `The subnet ID assign to interface.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_allowed_address_pairs/ip_address",
					Description: `The IP address of allowed address pairs.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_allowed_address_pairs/mac_address",
					Description: `The MAC address of allowed address pairs.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_allowed_address_pairs/type",
					Description: `Type of allowed address pairs.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_allowed_address_pairs/vrid",
					Description: `VRID of allowed address pairs.`,
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
					Name:        "default_gateway",
					Description: `IP address of default gateway.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "virtual_network_appliance_plan_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/name",
					Description: `Name of the interface.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/description",
					Description: `Description of the interface.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/network_id",
					Description: `The ID of network this interface belongs to.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_info/tags",
					Description: `Tags of the interface.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_fixed_ips/ip_address",
					Description: `The IP address assign to interface within subnet.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_fixed_ips/subnet_id",
					Description: `The subnet ID assign to interface.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_allowed_address_pairs/ip_address",
					Description: `The IP address of allowed address pairs.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_allowed_address_pairs/mac_address",
					Description: `The MAC address of allowed address pairs.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_allowed_address_pairs/type",
					Description: `Type of allowed address pairs.`,
				},
				resource.Attribute{
					Name:        "interface_[slot number]_allowed_address_pairs/vrid",
					Description: `VRID of allowed address pairs.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ecl_baremetal_availability_zone_v2":     0,
		"ecl_baremetal_flavor_v2":                1,
		"ecl_baremetal_keypair_v2":               2,
		"ecl_compute_flavor_v2":                  3,
		"ecl_compute_keypair_v2":                 4,
		"ecl_dns_zone_v2":                        5,
		"ecl_imagestorages_image_v2":             6,
		"ecl_network_common_function_gateway_v2": 7,
		"ecl_network_common_function_pool_v2":    8,
		"ecl_network_fic_gateway_v2":             9,
		"ecl_network_gateway_interface_v2":       10,
		"ecl_network_internet_gateway_v2":        11,
		"ecl_network_load_balancer_plan_v2":      12,
		"ecl_network_network_v2":                 13,
		"ecl_network_port_v2":                    14,
		"ecl_network_public_ip_v2":               15,
		"ecl_network_static_route_v2":            16,
		"ecl_network_subnet_v2":                  17,
		"ecl_sss_tenant_v1":                      18,
		"ecl_storage_virtualstorage_v2":          19,
		"ecl_storage_volume_v1":                  20,
		"ecl_storage_volumetype_v1":              21,
		"ecl_vna_appliance_v1":                   22,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
