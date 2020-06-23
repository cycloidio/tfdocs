package digitalocean

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_account",
			Category:         "Data Sources",
			ShortDescription: `Get information on your DigitalOcean account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get information on a certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of certificate. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_container_registry",
			Category:         "Data Sources",
			ShortDescription: `Get information on a container registry.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the container registry. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container registry`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container registry`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information on a database cluster resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database cluster. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database cluster.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the database cluster.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database engine used by the cluster (ex. ` + "`" + `pg` + "`" + ` for PostreSQL).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Engine version used by the cluster (ex. ` + "`" + `11` + "`" + ` for PostgreSQL 11).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Database droplet size associated with the cluster (ex. ` + "`" + `db-s-1vcpu-1gb` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `DigitalOcean region where the cluster will reside.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of nodes that will be included in the cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `Defines when the automatic maintenance should be performed for the database cluster.`,
				},
				resource.Attribute{
					Name:        "private_network_uuid",
					Description: `The ID of the VPC where the database cluster is located.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Database cluster's hostname.`,
				},
				resource.Attribute{
					Name:        "private_host",
					Description: `Same as ` + "`" + `host` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Network port that the database cluster is listening on.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The full URI for connecting to the database cluster.`,
				},
				resource.Attribute{
					Name:        "private_uri",
					Description: `Same as ` + "`" + `uri` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `Name of the cluster's default database.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Username for the cluster's default user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the cluster's default user. ` + "`" + `maintenance_window` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `The day of the week on which to apply maintenance updates.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `The hour in UTC at which maintenance updates will be applied in 24 hour format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database cluster.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the database cluster.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database engine used by the cluster (ex. ` + "`" + `pg` + "`" + ` for PostreSQL).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Engine version used by the cluster (ex. ` + "`" + `11` + "`" + ` for PostgreSQL 11).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Database droplet size associated with the cluster (ex. ` + "`" + `db-s-1vcpu-1gb` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `DigitalOcean region where the cluster will reside.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of nodes that will be included in the cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `Defines when the automatic maintenance should be performed for the database cluster.`,
				},
				resource.Attribute{
					Name:        "private_network_uuid",
					Description: `The ID of the VPC where the database cluster is located.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Database cluster's hostname.`,
				},
				resource.Attribute{
					Name:        "private_host",
					Description: `Same as ` + "`" + `host` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Network port that the database cluster is listening on.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The full URI for connecting to the database cluster.`,
				},
				resource.Attribute{
					Name:        "private_uri",
					Description: `Same as ` + "`" + `uri` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `Name of the cluster's default database.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Username for the cluster's default user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the cluster's default user. ` + "`" + `maintenance_window` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `The day of the week on which to apply maintenance updates.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `The hour in UTC at which maintenance updates will be applied in 24 hour format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_domain",
			Category:         "Data Sources",
			ShortDescription: `Get information on a domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the domain. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the domain`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the domain`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_droplet",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Droplet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the Droplet`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Droplet.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag applied to the Droplet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the Droplet`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the Droplet is running in.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The Droplet image ID or slug.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The unique slug that indentifies the type of Droplet.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The size of the Droplets disk in GB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of the Droplets virtual CPUs.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of the Droplets memory in MB.`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `Droplet hourly price.`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `Droplet monthly price.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Droplet.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the Droplet is locked.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The Droplets public IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv6_address_private",
					Description: `The Droplets private IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The Droplets public IPv4 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address_private",
					Description: `The Droplets private IPv4 address`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `Whether backups are enabled.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Whether IPv6 is enabled.`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `Whether private networks are enabled.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `Whether monitoring agent is installed.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `List of the IDs of each volumes attached to the Droplet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the Droplet.`,
				},
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `The ID of the VPC where the Droplet is located.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the Droplet`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the Droplet is running in.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The Droplet image ID or slug.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The unique slug that indentifies the type of Droplet.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The size of the Droplets disk in GB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of the Droplets virtual CPUs.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of the Droplets memory in MB.`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `Droplet hourly price.`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `Droplet monthly price.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Droplet.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the Droplet is locked.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The Droplets public IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv6_address_private",
					Description: `The Droplets private IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The Droplets public IPv4 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address_private",
					Description: `The Droplets private IPv4 address`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `Whether backups are enabled.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Whether IPv6 is enabled.`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `Whether private networks are enabled.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `Whether monitoring agent is installed.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `List of the IDs of each volumes attached to the Droplet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the Droplet.`,
				},
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `The ID of the VPC where the Droplet is located.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_droplet_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Get information about a DigitalOcean Droplet snapshot.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Droplet snapshot.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the Droplet snapshot list returned by DigitalOcean. This allows more advanced filtering not supported from the DigitalOcean API. This filtering is done locally on what DigitalOcean returns.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) A "slug" representing a DigitalOcean region (e.g. ` + "`" + `nyc1` + "`" + `). If set, only Droplet snapshots available in the region will be returned.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent Droplet snapshot. ~>`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the Droplet snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a Droplet to be created based on this Droplet snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the Droplet snapshot is available.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `The ID of the Droplet from which the Droplet snapshot originated.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the Droplet snapshot in gigabytes.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the Droplet snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a Droplet to be created based on this Droplet snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the Droplet snapshot is available.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `The ID of the Droplet from which the Droplet snapshot originated.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the Droplet snapshot in gigabytes.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_droplets",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information on Droplets.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the Droplets by this key. This may be one of '` + "`" + `backups` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `id` + "`" + `, ` + "`" + `image` + "`" + `, ` + "`" + `ipv4_address` + "`" + `, ` + "`" + `ipv4_address_private` + "`" + `, ` + "`" + `ipv6` + "`" + `, ` + "`" + `ipv6_address` + "`" + `, ` + "`" + `ipv6_address_private` + "`" + `, ` + "`" + `locked` + "`" + `, ` + "`" + `memory` + "`" + `, ` + "`" + `monitoring` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `price_hourly` + "`" + `, ` + "`" + `price_monthly` + "`" + `, ` + "`" + `private_networking` + "`" + `, ` + "`" + `region` + "`" + `, ` + "`" + `size` + "`" + `, ` + "`" + `status` + "`" + `, ` + "`" + `tags` + "`" + `, ` + "`" + `urn` + "`" + `, ` + "`" + `vcpus` + "`" + `, ` + "`" + `volume_ids` + "`" + `, or ` + "`" + `vpc_uuid` + "`" + `'.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves Droplets where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the Droplets by this key. This may be one of ` + "`" + `backups` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `id` + "`" + `, ` + "`" + `image` + "`" + `, ` + "`" + `ipv4_address` + "`" + `, ` + "`" + `ipv4_address_private` + "`" + `, ` + "`" + `ipv6` + "`" + `, ` + "`" + `ipv6_address` + "`" + `, ` + "`" + `ipv6_address_private` + "`" + `, ` + "`" + `locked` + "`" + `, ` + "`" + `memory` + "`" + `, ` + "`" + `monitoring` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `price_hourly` + "`" + `, ` + "`" + `price_monthly` + "`" + `, ` + "`" + `private_networking` + "`" + `, ` + "`" + `region` + "`" + `, ` + "`" + `size` + "`" + `, ` + "`" + `status` + "`" + `, ` + "`" + `urn` + "`" + `, ` + "`" + `vcpus` + "`" + `, or ` + "`" + `vpc_uuid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "droplets",
					Description: `A list of Droplets satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each Droplet has the following attributes: - ` + "`" + `id` + "`" + ` - The ID of the Droplet. - ` + "`" + `urn` + "`" + ` - The uniform resource name of the Droplet - ` + "`" + `region` + "`" + ` - The region the Droplet is running in. - ` + "`" + `image` + "`" + ` - The Droplet image ID or slug. - ` + "`" + `size` + "`" + ` - The unique slug that identifies the type of Droplet. - ` + "`" + `disk` + "`" + ` - The size of the Droplet's disk in GB. - ` + "`" + `vcpus` + "`" + ` - The number of the Droplet's virtual CPUs. - ` + "`" + `memory` + "`" + ` - The amount of the Droplet's memory in MB. - ` + "`" + `price_hourly` + "`" + ` - Droplet hourly price. - ` + "`" + `price_monthly` + "`" + ` - Droplet monthly price. - ` + "`" + `status` + "`" + ` - The status of the Droplet. - ` + "`" + `locked` + "`" + ` - Whether the Droplet is locked. - ` + "`" + `ipv6_address` + "`" + ` - The Droplet's public IPv6 address - ` + "`" + `ipv6_address_private` + "`" + ` - The Droplet's private IPv6 address - ` + "`" + `ipv4_address` + "`" + ` - The Droplet's public IPv4 address - ` + "`" + `ipv4_address_private` + "`" + ` - The Droplet's private IPv4 address - ` + "`" + `backups` + "`" + ` - Whether backups are enabled. - ` + "`" + `ipv6` + "`" + ` - Whether IPv6 is enabled. - ` + "`" + `private_networking` + "`" + ` - Whether private networks are enabled. - ` + "`" + `monitoring` + "`" + ` - Whether monitoring agent is installed. - ` + "`" + `volume_ids` + "`" + ` - List of the IDs of each volumes attached to the Droplet. - ` + "`" + `tags` + "`" + ` - A list of the tags associated to the Droplet. - ` + "`" + `vpc_uuid` + "`" + ` - The ID of the VPC where the Droplet is located.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "droplets",
					Description: `A list of Droplets satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each Droplet has the following attributes: - ` + "`" + `id` + "`" + ` - The ID of the Droplet. - ` + "`" + `urn` + "`" + ` - The uniform resource name of the Droplet - ` + "`" + `region` + "`" + ` - The region the Droplet is running in. - ` + "`" + `image` + "`" + ` - The Droplet image ID or slug. - ` + "`" + `size` + "`" + ` - The unique slug that identifies the type of Droplet. - ` + "`" + `disk` + "`" + ` - The size of the Droplet's disk in GB. - ` + "`" + `vcpus` + "`" + ` - The number of the Droplet's virtual CPUs. - ` + "`" + `memory` + "`" + ` - The amount of the Droplet's memory in MB. - ` + "`" + `price_hourly` + "`" + ` - Droplet hourly price. - ` + "`" + `price_monthly` + "`" + ` - Droplet monthly price. - ` + "`" + `status` + "`" + ` - The status of the Droplet. - ` + "`" + `locked` + "`" + ` - Whether the Droplet is locked. - ` + "`" + `ipv6_address` + "`" + ` - The Droplet's public IPv6 address - ` + "`" + `ipv6_address_private` + "`" + ` - The Droplet's private IPv6 address - ` + "`" + `ipv4_address` + "`" + ` - The Droplet's public IPv4 address - ` + "`" + `ipv4_address_private` + "`" + ` - The Droplet's private IPv4 address - ` + "`" + `backups` + "`" + ` - Whether backups are enabled. - ` + "`" + `ipv6` + "`" + ` - Whether IPv6 is enabled. - ` + "`" + `private_networking` + "`" + ` - Whether private networks are enabled. - ` + "`" + `monitoring` + "`" + ` - Whether monitoring agent is installed. - ` + "`" + `volume_ids` + "`" + ` - List of the IDs of each volumes attached to the Droplet. - ` + "`" + `tags` + "`" + ` - A list of the tags associated to the Droplet. - ` + "`" + `vpc_uuid` + "`" + ` - The ID of the VPC where the Droplet is located.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_floating_ip",
			Category:         "Data Sources",
			ShortDescription: `Get information on a floating IP.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The allocated IP address of the specific floating IP to retrieve. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_image",
			Category:         "Data Sources",
			ShortDescription: `Get information on a DigitalOcean image.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the image`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the image.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug of the official image. If ` + "`" + `name` + "`" + ` is specified, you may also specify:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Restrict the search to one of the following categories of images: - ` + "`" + `all` + "`" + ` - All images (whether public or private) - ` + "`" + `applications` + "`" + ` - One-click applications - ` + "`" + `distributions` + "`" + ` - Distributions - ` + "`" + `user` + "`" + ` - (Default) User (private) images ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `The name of the distribution of the OS of the image.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `Is image a public image or not. Public images represent Linux distributions or One-Click Applications, while non-public images represent snapshots and backups and are only available within your account.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The id of the image (legacy parameter).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "distribution",
					Description: `The name of the distribution of the OS of the image.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `Is image a public image or not. Public images represent Linux distributions or One-Click Applications, while non-public images represent snapshots and backups and are only available within your account.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The id of the image (legacy parameter).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_images",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about DigitalOcean images (public and private).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the images by this key. This may be one of ` + "`" + `distribution` + "`" + `, ` + "`" + `error_message` + "`" + `, ` + "`" + `id` + "`" + `, ` + "`" + `image` + "`" + `, ` + "`" + `min_disk_size` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `private` + "`" + `, ` + "`" + `regions` + "`" + `, ` + "`" + `size_gigabytes` + "`" + `, ` + "`" + `slug` + "`" + `, ` + "`" + `status` + "`" + `, ` + "`" + `tags` + "`" + `, or ` + "`" + `type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves images where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the images by this key. This may be one of ` + "`" + `distribution` + "`" + `, ` + "`" + `error_message` + "`" + `, ` + "`" + `id` + "`" + `, ` + "`" + `image` + "`" + `, ` + "`" + `min_disk_size` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `private` + "`" + `, ` + "`" + `size_gigabytes` + "`" + `, ` + "`" + `slug` + "`" + `, ` + "`" + `status` + "`" + `, or ` + "`" + `type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `A set of images satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each image has the following attributes: - ` + "`" + `slug` + "`" + `: Unique text identifier of the image. - ` + "`" + `id` + "`" + `: The ID of the image. - ` + "`" + `name` + "`" + `: The name of the image. - ` + "`" + `type` + "`" + `: Type of the image. - ` + "`" + `distribution` + "`" + ` - The name of the distribution of the OS of the image. - ` + "`" + `min_disk_size` + "`" + `: The minimum 'disk' required for the image. - ` + "`" + `size_gigabytes` + "`" + `: The size of the image in GB. - ` + "`" + `private` + "`" + ` - Is image a public image or not. Public images represent Linux distributions or One-Click Applications, while non-public images represent snapshots and backups and are only available within your account. - ` + "`" + `regions` + "`" + `: A set of the regions that the image is available in. - ` + "`" + `tags` + "`" + `: A set of tags applied to the image - ` + "`" + `created` + "`" + `: When the image was created - ` + "`" + `status` + "`" + `: Current status of the image - ` + "`" + `error_message` + "`" + `: Any applicable error message pertaining to the image - ` + "`" + `image` + "`" + ` - The id of the image (legacy parameter).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `A set of images satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each image has the following attributes: - ` + "`" + `slug` + "`" + `: Unique text identifier of the image. - ` + "`" + `id` + "`" + `: The ID of the image. - ` + "`" + `name` + "`" + `: The name of the image. - ` + "`" + `type` + "`" + `: Type of the image. - ` + "`" + `distribution` + "`" + ` - The name of the distribution of the OS of the image. - ` + "`" + `min_disk_size` + "`" + `: The minimum 'disk' required for the image. - ` + "`" + `size_gigabytes` + "`" + `: The size of the image in GB. - ` + "`" + `private` + "`" + ` - Is image a public image or not. Public images represent Linux distributions or One-Click Applications, while non-public images represent snapshots and backups and are only available within your account. - ` + "`" + `regions` + "`" + `: A set of the regions that the image is available in. - ` + "`" + `tags` + "`" + `: A set of tags applied to the image - ` + "`" + `created` + "`" + `: When the image was created - ` + "`" + `status` + "`" + `: Current status of the image - ` + "`" + `error_message` + "`" + `: Any applicable error message pertaining to the image - ` + "`" + `image` + "`" + ` - The id of the image (legacy parameter).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_kubernetes_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information on a DigitalOcean Kubernetes cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of Kubernetes cluster. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID that can be used to identify and reference a Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The slug identifier for the region where the Kubernetes cluster is located.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The slug identifier for the version of Kubernetes used for the cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tag names to be applied to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `The ID of the VPC where the Kubernetes cluster is located.`,
				},
				resource.Attribute{
					Name:        "cluster_subnet",
					Description: `The range of IP addresses in the overlay network of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "service_subnet",
					Description: `The range of assignable IP addresses for services running in the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The public IPv4 address of the Kubernetes master node.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The base URL of the API server on the Kubernetes master node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the Kubernetes cluster was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date and time when the Kubernetes cluster was last updated.`,
				},
				resource.Attribute{
					Name:        "kube_config.0",
					Description: `A representation of the Kubernetes cluster's kubeconfig with the following attributes: - ` + "`" + `raw_config` + "`" + ` - The full contents of the Kubernetes cluster's kubeconfig file. - ` + "`" + `host` + "`" + ` - The URL of the API server on the Kubernetes master node. - ` + "`" + `cluster_ca_certificate` + "`" + ` - The base64 encoded public certificate for the cluster's certificate authority. - ` + "`" + `token` + "`" + ` - The DigitalOcean API access token used by clients to access the cluster. - ` + "`" + `client_key` + "`" + ` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster. - ` + "`" + `client_certificate` + "`" + ` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster. - ` + "`" + `expires_at` + "`" + ` - The date and time when the credentials will expire and need to be regenerated.`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `A list of node pools associated with the cluster. Each node pool exports the following attributes: - ` + "`" + `id` + "`" + ` - The unique ID that can be used to identify and reference the node pool. - ` + "`" + `name` + "`" + ` - The name of the node pool. - ` + "`" + `size` + "`" + ` - The slug identifier for the type of Droplet used as workers in the node pool. - ` + "`" + `node_count` + "`" + ` - The number of Droplet instances in the node pool. - ` + "`" + `actual_node_count` + "`" + ` - The actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled. - ` + "`" + `auto_scale` + "`" + ` - A boolean indicating whether auto-scaling is enabled on the node pool. - ` + "`" + `min_nodes` + "`" + ` - If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to. - ` + "`" + `max_nodes` + "`" + ` - If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to. - ` + "`" + `tags` + "`" + ` - A list of tag names applied to the node pool. - ` + "`" + `labels` + "`" + ` - A map of key/value pairs applied to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/). - ` + "`" + `nodes` + "`" + ` - A list of nodes in the pool. Each node exports the following attributes: + ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. + ` + "`" + `name` + "`" + ` - The auto-generated name for the node. + ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. + ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. + ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID that can be used to identify and reference a Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The slug identifier for the region where the Kubernetes cluster is located.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The slug identifier for the version of Kubernetes used for the cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tag names to be applied to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `The ID of the VPC where the Kubernetes cluster is located.`,
				},
				resource.Attribute{
					Name:        "cluster_subnet",
					Description: `The range of IP addresses in the overlay network of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "service_subnet",
					Description: `The range of assignable IP addresses for services running in the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The public IPv4 address of the Kubernetes master node.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The base URL of the API server on the Kubernetes master node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the Kubernetes cluster was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date and time when the Kubernetes cluster was last updated.`,
				},
				resource.Attribute{
					Name:        "kube_config.0",
					Description: `A representation of the Kubernetes cluster's kubeconfig with the following attributes: - ` + "`" + `raw_config` + "`" + ` - The full contents of the Kubernetes cluster's kubeconfig file. - ` + "`" + `host` + "`" + ` - The URL of the API server on the Kubernetes master node. - ` + "`" + `cluster_ca_certificate` + "`" + ` - The base64 encoded public certificate for the cluster's certificate authority. - ` + "`" + `token` + "`" + ` - The DigitalOcean API access token used by clients to access the cluster. - ` + "`" + `client_key` + "`" + ` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster. - ` + "`" + `client_certificate` + "`" + ` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster. - ` + "`" + `expires_at` + "`" + ` - The date and time when the credentials will expire and need to be regenerated.`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `A list of node pools associated with the cluster. Each node pool exports the following attributes: - ` + "`" + `id` + "`" + ` - The unique ID that can be used to identify and reference the node pool. - ` + "`" + `name` + "`" + ` - The name of the node pool. - ` + "`" + `size` + "`" + ` - The slug identifier for the type of Droplet used as workers in the node pool. - ` + "`" + `node_count` + "`" + ` - The number of Droplet instances in the node pool. - ` + "`" + `actual_node_count` + "`" + ` - The actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled. - ` + "`" + `auto_scale` + "`" + ` - A boolean indicating whether auto-scaling is enabled on the node pool. - ` + "`" + `min_nodes` + "`" + ` - If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to. - ` + "`" + `max_nodes` + "`" + ` - If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to. - ` + "`" + `tags` + "`" + ` - A list of tag names applied to the node pool. - ` + "`" + `labels` + "`" + ` - A map of key/value pairs applied to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/). - ` + "`" + `nodes` + "`" + ` - A list of nodes in the pool. Each node exports the following attributes: + ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. + ` + "`" + `name` + "`" + ` - The auto-generated name for the node. + ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. + ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. + ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_kubernetes_versions",
			Category:         "Data Sources",
			ShortDescription: `Get available DigitalOcean Kubernetes versions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "version_prefix",
					Description: `(Optional) If provided, Terraform will only return versions that match the string prefix. For example, ` + "`" + `1.15.` + "`" + ` will match all 1.15.x series releases. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "valid_versions",
					Description: `A list of available versions.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The most recent version available.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "valid_versions",
					Description: `A list of available versions.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The most recent version available.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_loadbalancer",
			Category:         "Data Sources",
			ShortDescription: `Get information on a loadbalancer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of load balancer.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the Load Balancer ## Attributes Reference See the [Load Balancer Resource](/docs/providers/do/r/loadbalancer.html) for details on the returned attributes - they are identical.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_project",
			Category:         "Data Sources",
			ShortDescription: `Get information on a DigitalOcean project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) the ID of the project to retrieve`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) the name of the project to retrieve. The data source will raise an error if more than one project has the provided name or if no project has that name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the project`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `The purpose of the project, (Default: "Web Application")`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The environment of the project's resources. The possible values are: ` + "`" + `Development` + "`" + `, ` + "`" + `Staging` + "`" + `, ` + "`" + `Production` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `A set of uniform resource names (URNs) for the resources associated with the project`,
				},
				resource.Attribute{
					Name:        "owner_uuid",
					Description: `The unique universal identifier of the project owner.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the project owner.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the project was created, (ISO8601)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date and time when the project was last updated, (ISO8601)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the project`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `The purpose of the project, (Default: "Web Application")`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The environment of the project's resources. The possible values are: ` + "`" + `Development` + "`" + `, ` + "`" + `Staging` + "`" + `, ` + "`" + `Production` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `A set of uniform resource names (URNs) for the resources associated with the project`,
				},
				resource.Attribute{
					Name:        "owner_uuid",
					Description: `The unique universal identifier of the project owner.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the project owner.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the project was created, (ISO8601)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date and time when the project was last updated, (ISO8601)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_projects",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about DigitalOcean projects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the projects by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `purpose` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `environment` + "`" + `, or ` + "`" + `is_default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves projects where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the projects by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `purpose` + "`" + `, ` + "`" + `description` + "`" + `, or ` + "`" + `environment` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "projects",
					Description: `A set of projects satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each project has the following attributes: - ` + "`" + `id` + "`" + ` - The ID of the project - ` + "`" + `name` + "`" + ` - The name of the project - ` + "`" + `description` + "`" + ` - The description of the project - ` + "`" + `purpose` + "`" + ` - The purpose of the project (Default: "Web Application") - ` + "`" + `environment` + "`" + ` - The environment of the project's resources. The possible values are: ` + "`" + `Development` + "`" + `, ` + "`" + `Staging` + "`" + `, ` + "`" + `Production` + "`" + `. - ` + "`" + `resources` + "`" + ` - A set of uniform resource names (URNs) for the resources associated with the project - ` + "`" + `owner_uuid` + "`" + ` - The unique universal identifier of the project owner - ` + "`" + `owner_id` + "`" + ` - The ID of the project owner - ` + "`" + `created_at` + "`" + ` - The date and time when the project was created, (ISO8601) - ` + "`" + `updated_at` + "`" + ` - The date and time when the project was last updated, (ISO8601)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "projects",
					Description: `A set of projects satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each project has the following attributes: - ` + "`" + `id` + "`" + ` - The ID of the project - ` + "`" + `name` + "`" + ` - The name of the project - ` + "`" + `description` + "`" + ` - The description of the project - ` + "`" + `purpose` + "`" + ` - The purpose of the project (Default: "Web Application") - ` + "`" + `environment` + "`" + ` - The environment of the project's resources. The possible values are: ` + "`" + `Development` + "`" + `, ` + "`" + `Staging` + "`" + `, ` + "`" + `Production` + "`" + `. - ` + "`" + `resources` + "`" + ` - A set of uniform resource names (URNs) for the resources associated with the project - ` + "`" + `owner_uuid` + "`" + ` - The unique universal identifier of the project owner - ` + "`" + `owner_id` + "`" + ` - The ID of the project owner - ` + "`" + `created_at` + "`" + ` - The date and time when the project was created, (ISO8601) - ` + "`" + `updated_at` + "`" + ` - The date and time when the project was last updated, (ISO8601)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_record",
			Category:         "Data Sources",
			ShortDescription: `Get information on a DNS record.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain name of the record. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_region",
			Category:         "Data Sources",
			ShortDescription: `Get information on a DigitalOcean region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) A human-readable string that is used as a unique identifier for each region. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `A human-readable string that is used as a unique identifier for each region.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The display name of the region.`,
				},
				resource.Attribute{
					Name:        "available",
					Description: `A boolean value that represents whether new Droplets can be created in this region.`,
				},
				resource.Attribute{
					Name:        "sizes",
					Description: `A set of identifying slugs for the Droplet sizes available in this region.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `A set of features available in this region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `A human-readable string that is used as a unique identifier for each region.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The display name of the region.`,
				},
				resource.Attribute{
					Name:        "available",
					Description: `A boolean value that represents whether new Droplets can be created in this region.`,
				},
				resource.Attribute{
					Name:        "sizes",
					Description: `A set of identifying slugs for the Droplet sizes available in this region.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `A set of features available in this region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_regions",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about DigitalOcean regions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the regions by this key. This may be one of ` + "`" + `slug` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `available` + "`" + `, ` + "`" + `features` + "`" + `, or ` + "`" + `sizes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves regions where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the regions by this key. This may be one of ` + "`" + `slug` + "`" + `, ` + "`" + `name` + "`" + `, or ` + "`" + `available` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A set of regions satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each region has the following attributes: - ` + "`" + `slug` + "`" + ` - A human-readable string that is used as a unique identifier for each region. - ` + "`" + `name` + "`" + ` - The display name of the region. - ` + "`" + `available` + "`" + ` - A boolean value that represents whether new Droplets can be created in this region. - ` + "`" + `sizes` + "`" + ` - A set of identifying slugs for the Droplet sizes available in this region. - ` + "`" + `features` + "`" + ` - A set of features available in this region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `A set of regions satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each region has the following attributes: - ` + "`" + `slug` + "`" + ` - A human-readable string that is used as a unique identifier for each region. - ` + "`" + `name` + "`" + ` - The display name of the region. - ` + "`" + `available` + "`" + ` - A boolean value that represents whether new Droplets can be created in this region. - ` + "`" + `sizes` + "`" + ` - A set of identifying slugs for the Droplet sizes available in this region. - ` + "`" + `features` + "`" + ` - A set of features available in this region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_sizes",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information on supported Droplet sizes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the sizes by this key. This may be one of ` + "`" + `slug` + "`" + `, ` + "`" + `regions` + "`" + `, ` + "`" + `memory` + "`" + `, ` + "`" + `vcpus` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `transfer` + "`" + `, ` + "`" + `price_monthly` + "`" + `, ` + "`" + `price_hourly` + "`" + `, or ` + "`" + `available` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Only retrieves images which keys has value that matches one of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the sizes by this key. This may be one of ` + "`" + `slug` + "`" + `, ` + "`" + `memory` + "`" + `, ` + "`" + `vcpus` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `transfer` + "`" + `, ` + "`" + `price_monthly` + "`" + `, or ` + "`" + `price_hourly` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `A human-readable string that is used to uniquely identify each size.`,
				},
				resource.Attribute{
					Name:        "available",
					Description: `This represents whether new Droplets can be created with this size.`,
				},
				resource.Attribute{
					Name:        "transfer",
					Description: `The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `The monthly cost of Droplets created in this size if they are kept for an entire month. The value is measured in US dollars.`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `The hourly cost of Droplets created in this size as measured hourly. The value is measured in US dollars.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of RAM allocated to Droplets created of this size. The value is measured in megabytes.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of CPUs allocated to Droplets of this size.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The amount of disk space set aside for Droplets of this size. The value is measured in gigabytes.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of region slugs where Droplets can be created in this size.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `A human-readable string that is used to uniquely identify each size.`,
				},
				resource.Attribute{
					Name:        "available",
					Description: `This represents whether new Droplets can be created with this size.`,
				},
				resource.Attribute{
					Name:        "transfer",
					Description: `The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `The monthly cost of Droplets created in this size if they are kept for an entire month. The value is measured in US dollars.`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `The hourly cost of Droplets created in this size as measured hourly. The value is measured in US dollars.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of RAM allocated to Droplets created of this size. The value is measured in megabytes.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of CPUs allocated to Droplets of this size.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The amount of disk space set aside for Droplets of this size. The value is measured in gigabytes.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of region slugs where Droplets can be created in this size.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_bucket",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Spaces bucket.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Spaces bucket.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The slug of the region where the bucket is stored. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Spaces bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The slug of the region where the bucket is stored.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the bucket`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Spaces bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The slug of the region where the bucket is stored.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the bucket`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_bucket_object",
			Category:         "Data Sources",
			ShortDescription: `Provides metadata and optionally content of a Spaces object`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to read the object from.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The slug of the region where the bucket is stored.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The full path to the object inside the bucket`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Optional) Specific version ID of the object returned (defaults to latest version) ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in Spaces`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Spaces stores the value of this header in the object metadata. ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in Spaces`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Spaces stores the value of this header in the object metadata. ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_bucket_objects",
			Category:         "Data Sources",
			ShortDescription: `Returns keys and metadata of Spaces objects`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Lists object keys in this Spaces bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The slug of the region where the bucket is stored.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Limits results to object keys with this prefix (Default: none)`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) A character used to group keys (Default: none)`,
				},
				resource.Attribute{
					Name:        "encoding_type",
					Description: `(Optional) Encodes keys using this method (Default: none; besides none, only "url" can be used)`,
				},
				resource.Attribute{
					Name:        "max_keys",
					Description: `(Optional) Maximum object keys to return (Default: 1000) ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `List of strings representing object keys`,
				},
				resource.Attribute{
					Name:        "common_prefixes",
					Description: `List of any keys between ` + "`" + `prefix` + "`" + ` and the next occurrence of ` + "`" + `delimiter` + "`" + ` (i.e., similar to subdirectories of the ` + "`" + `prefix` + "`" + ` "directory"); the list is only returned when you specify ` + "`" + `delimiter` + "`" + ``,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `List of strings representing object owner IDs`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keys",
					Description: `List of strings representing object keys`,
				},
				resource.Attribute{
					Name:        "common_prefixes",
					Description: `List of any keys between ` + "`" + `prefix` + "`" + ` and the next occurrence of ` + "`" + `delimiter` + "`" + ` (i.e., similar to subdirectories of the ` + "`" + `prefix` + "`" + ` "directory"); the list is only returned when you specify ` + "`" + `delimiter` + "`" + ``,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `List of strings representing object owner IDs`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_buckets",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information on Spaces buckets.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the images by this key. This may be one of ` + "`" + `bucket_domain_name` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `region` + "`" + `, or ` + "`" + `urn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves images where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the images by this key. This may be one of ` + "`" + `bucket_domain_name` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `region` + "`" + `, or ` + "`" + `urn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "buckets",
					Description: `A list of Spaces buckets satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each bucket has the following attributes: - ` + "`" + `name` + "`" + ` - The name of the Spaces bucket - ` + "`" + `region` + "`" + ` - The slug of the region where the bucket is stored. - ` + "`" + `urn` + "`" + ` - The uniform resource name of the bucket - ` + "`" + `bucket_domain_name` + "`" + ` - The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "buckets",
					Description: `A list of Spaces buckets satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each bucket has the following attributes: - ` + "`" + `name` + "`" + ` - The name of the Spaces bucket - ` + "`" + `region` + "`" + ` - The slug of the region where the bucket is stored. - ` + "`" + `urn` + "`" + ` - The uniform resource name of the bucket - ` + "`" + `bucket_domain_name` + "`" + ` - The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Get information on a ssh key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ssh key. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_tag",
			Category:         "Data Sources",
			ShortDescription: `Get information on a tag.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tag. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "total_resource_count",
					Description: `A count of the total number of resources that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "droplets_count",
					Description: `A count of the Droplets the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "images_count",
					Description: `A count of the images that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volumes_count",
					Description: `A count of the volumes that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volume_snapshots_count",
					Description: `A count of the volume snapshots that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "databases_count",
					Description: `A count of the database clusters that the tag is applied to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total_resource_count",
					Description: `A count of the total number of resources that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "droplets_count",
					Description: `A count of the Droplets the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "images_count",
					Description: `A count of the images that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volumes_count",
					Description: `A count of the volumes that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volume_snapshots_count",
					Description: `A count of the volume snapshots that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "databases_count",
					Description: `A count of the database clusters that the tag is applied to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_tags",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information on tags.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the tags by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `total_resource_count` + "`" + `, ` + "`" + `droplets_count` + "`" + `, ` + "`" + `images_count` + "`" + `, ` + "`" + `volumes_count` + "`" + `, ` + "`" + `volume_snapshots_count` + "`" + `, or ` + "`" + `databases_count` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Only retrieves tags which keys has value that matches one of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the tags by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `total_resource_count` + "`" + `, ` + "`" + `droplets_count` + "`" + `, ` + "`" + `images_count` + "`" + `, ` + "`" + `volumes_count` + "`" + `, ` + "`" + `volume_snapshots_count` + "`" + `, or ` + "`" + `databases_count` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference The following attributes are exported for each tag:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
				},
				resource.Attribute{
					Name:        "total_resource_count",
					Description: `A count of the total number of resources that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "droplets_count",
					Description: `A count of the Droplets the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "images_count",
					Description: `A count of the images that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volumes_count",
					Description: `A count of the volumes that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volume_snapshots_count",
					Description: `A count of the volume snapshots that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "databases_count",
					Description: `A count of the database clusters that the tag is applied to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
				},
				resource.Attribute{
					Name:        "total_resource_count",
					Description: `A count of the total number of resources that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "droplets_count",
					Description: `A count of the Droplets the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "images_count",
					Description: `A count of the images that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volumes_count",
					Description: `A count of the volumes that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volume_snapshots_count",
					Description: `A count of the volume snapshots that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "databases_count",
					Description: `A count of the database clusters that the tag is applied to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_volume",
			Category:         "Data Sources",
			ShortDescription: `Get information on a volume.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of block storage volume.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region the block storage volume is provisioned in. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the block storage volume in GiB.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Text describing a block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_type",
					Description: `Filesystem type currently in-use on the block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_label",
					Description: `Filesystem label currently in-use on the block storage volume.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `A list of associated Droplet ids.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the Volume.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `The size of the block storage volume in GiB.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Text describing a block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_type",
					Description: `Filesystem type currently in-use on the block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_label",
					Description: `Filesystem label currently in-use on the block storage volume.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `A list of associated Droplet ids.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the Volume.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_volume_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Get information about a DigitalOcean volume snapshot.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the volume snapshot.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the volume snapshot list returned by DigitalOcean. This allows more advanced filtering not supported from the DigitalOcean API. This filtering is done locally on what DigitalOcean returns.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) A "slug" representing a DigitalOcean region (e.g. ` + "`" + `nyc1` + "`" + `). If set, only volume snapshots available in the region will be returned.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent volume snapshot. ~>`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the volume snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a volume to be created based on this volume snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume from which the volume snapshot originated.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the volume snapshot in gigabytes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the volume snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the volume snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a volume to be created based on this volume snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume from which the volume snapshot originated.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the volume snapshot in gigabytes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the volume snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_vpc",
			Category:         "Data Sources",
			ShortDescription: `Get information about a VPC.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of an existing VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of an existing VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The DigitalOcean region slug for the VPC's location. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The DigitalOcean region slug for the VPC's location.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free-form text field describing the VPC.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The range of IP addresses for the VPC in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name (URN) for the VPC.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `A boolean indicating whether or not the VPC is the default one for the region.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time of when the VPC was created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The DigitalOcean region slug for the VPC's location.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free-form text field describing the VPC.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The range of IP addresses for the VPC in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name (URN) for the VPC.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `A boolean indicating whether or not the VPC is the default one for the region.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time of when the VPC was created.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"digitalocean_account":               0,
		"digitalocean_certificate":           1,
		"digitalocean_container_registry":    2,
		"digitalocean_database_cluster":      3,
		"digitalocean_domain":                4,
		"digitalocean_droplet":               5,
		"digitalocean_droplet_snapshot":      6,
		"digitalocean_droplets":              7,
		"digitalocean_floating_ip":           8,
		"digitalocean_image":                 9,
		"digitalocean_images":                10,
		"digitalocean_kubernetes_cluster":    11,
		"digitalocean_kubernetes_versions":   12,
		"digitalocean_loadbalancer":          13,
		"digitalocean_project":               14,
		"digitalocean_projects":              15,
		"digitalocean_record":                16,
		"digitalocean_region":                17,
		"digitalocean_regions":               18,
		"digitalocean_sizes":                 19,
		"digitalocean_spaces_bucket":         20,
		"digitalocean_spaces_bucket_object":  21,
		"digitalocean_spaces_bucket_objects": 22,
		"digitalocean_spaces_buckets":        23,
		"digitalocean_ssh_key":               24,
		"digitalocean_tag":                   25,
		"digitalocean_tags":                  26,
		"digitalocean_volume":                27,
		"digitalocean_volume_snapshot":       28,
		"digitalocean_vpc":                   29,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
