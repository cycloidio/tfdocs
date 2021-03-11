package ah

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_cloud_images",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results by specified key and direction. The structure of the block is documented below. --- The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `distribution` + "`" + `, ` + "`" + `version` + "`" + `, ` + "`" + `architecture` + "`" + `, ` + "`" + `slug` + "`" + ``,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. The ` + "`" + `sort` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `distribution` + "`" + `, ` + "`" + `version` + "`" + `, ` + "`" + `architecture` + "`" + `, ` + "`" + `slug` + "`" + ``,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Sort direction of the results. Can be one of: ` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `. Default option is ` + "`" + `desc` + "`" + `. --- ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `A list of Images that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Image.`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `Name of the Image Distribution.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Distribution version.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `Distribution architecture.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Slug of the Image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `A list of Images that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Image.`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `Name of the Image Distribution.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Distribution version.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `Distribution architecture.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Slug of the Image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_cloud_server_products",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results by specified key and direction. The structure of the block is documented below. --- The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `slug` + "`" + `, ` + "`" + `price` + "`" + `, ` + "`" + `currency` + "`" + `, ` + "`" + `vcpu` + "`" + `, ` + "`" + `ram` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `available_on_trial` + "`" + ``,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. The ` + "`" + `sort` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `slug` + "`" + `, ` + "`" + `price` + "`" + `, ` + "`" + `currency` + "`" + `, ` + "`" + `vcpu` + "`" + `, ` + "`" + `ram` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `available_on_trial` + "`" + ``,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Sort direction of the results. Can be one of: ` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `. Default option is ` + "`" + `desc` + "`" + `. --- ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "products",
					Description: `A list of Products that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Product.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Slug of the Product.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Monthly price of the Product.`,
				},
				resource.Attribute{
					Name:        "currency",
					Description: `Currency for the price.`,
				},
				resource.Attribute{
					Name:        "vcpu",
					Description: `Number of vCPUs.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `RAM in GiB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "available_on_trial",
					Description: `Boolean flag indicating whether the Product is available on trial.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "products",
					Description: `A list of Products that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Product.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Slug of the Product.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Monthly price of the Product.`,
				},
				resource.Attribute{
					Name:        "currency",
					Description: `Currency for the price.`,
				},
				resource.Attribute{
					Name:        "vcpu",
					Description: `Number of vCPUs.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `RAM in GiB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "available_on_trial",
					Description: `Boolean flag indicating whether the Product is available on trial.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_cloud_server_snapshots_and_backups",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results by specified key and direction. The structure of the block is documented below. --- The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `cloud_server_id` + "`" + `, ` + "`" + `cloud_server_name` + "`" + `, ` + "`" + `state` + "`" + `, ` + "`" + `size` + "`" + `, ` + "`" + `type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. The ` + "`" + `sort` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `cloud_server_id` + "`" + `, ` + "`" + `cloud_server_name` + "`" + `, ` + "`" + `state` + "`" + `, ` + "`" + `size` + "`" + `, ` + "`" + `type` + "`" + `,` + "`" + `created_at` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Sort direction of the results. Can be one of: ` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `. Default option is ` + "`" + `desc` + "`" + `. --- ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshots_and_backups",
					Description: `A list of snapshots and backups that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Snapshot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "cloud_server_id",
					Description: `Cloud Server ID a Snapshot was created from.`,
				},
				resource.Attribute{
					Name:        "cloud_server_name",
					Description: `Cloud Server Name snapshot was created for.`,
				},
				resource.Attribute{
					Name:        "cloud_server_deleted",
					Description: `Boolen flag indicating whether the original Cloud Server was deleted.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Snapshot.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Snapshot size, in GB`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type. Can be ` + "`" + `snapshot` + "`" + ` (for manual snapshots) or ` + "`" + `backup` + "`" + ` (for automatic backups).`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshots_and_backups",
					Description: `A list of snapshots and backups that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Snapshot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "cloud_server_id",
					Description: `Cloud Server ID a Snapshot was created from.`,
				},
				resource.Attribute{
					Name:        "cloud_server_name",
					Description: `Cloud Server Name snapshot was created for.`,
				},
				resource.Attribute{
					Name:        "cloud_server_deleted",
					Description: `Boolen flag indicating whether the original Cloud Server was deleted.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Snapshot.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Snapshot size, in GB`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type. Can be ` + "`" + `snapshot` + "`" + ` (for manual snapshots) or ` + "`" + `backup` + "`" + ` (for automatic backups).`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_cloud_servers",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results by specified key and direction. The structure of the block is documented below. --- The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `state` + "`" + `, ` + "`" + `vcpu` + "`" + `, ` + "`" + `ram` + "`" + `, ` + "`" + `disk` + "`" + ``,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. The ` + "`" + `sort` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `state` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `vcpu` + "`" + `, ` + "`" + `ram` + "`" + `, ` + "`" + `disk` + "`" + ``,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Sort direction of the results. Can be one of: ` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `. Default option is ` + "`" + `desc` + "`" + `. --- ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cloud_servers",
					Description: `A list of Cloud Servers that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `Datacenter slug of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `Product slug that indentifies the product type of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "vcpu",
					Description: `Number of vCPUs on the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `RAM of the server in MiB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disk size of the server in GB.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The Cloud Server Image ID or Snapshot / Auto Backup ID the server was created from.`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `Boolean indicating whether backups are enabled for the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "use_password",
					Description: `Boolean indicating whether the Cloud Server was created with a password generated.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Array of Public and Anycast IP addresses assigned to the the Cloud Server. The structure of the block is documented below.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `Array of Volume IDs attached to the server.`,
				},
				resource.Attribute{
					Name:        "private_networks",
					Description: `Array of Private Networks connected to the server. The structure of the block is documented below. --- The ` + "`" + `ips` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Public IP address assigned to the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP address type. Can be either ` + "`" + `public` + "`" + ` or ` + "`" + `anycast` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Boolean indicating a Primary IP flag.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `Reverse DNS assigned to the IP address.`,
				},
				resource.Attribute{
					Name:        "assignment_id",
					Description: `ID of the IP Address Assignment. The ` + "`" + `private_networks` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Private Network ID.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Private network IP address of the Cloud Server within the network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_servers",
					Description: `A list of Cloud Servers that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `Datacenter slug of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `Product slug that indentifies the product type of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "vcpu",
					Description: `Number of vCPUs on the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `RAM of the server in MiB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disk size of the server in GB.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The Cloud Server Image ID or Snapshot / Auto Backup ID the server was created from.`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `Boolean indicating whether backups are enabled for the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "use_password",
					Description: `Boolean indicating whether the Cloud Server was created with a password generated.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Array of Public and Anycast IP addresses assigned to the the Cloud Server. The structure of the block is documented below.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `Array of Volume IDs attached to the server.`,
				},
				resource.Attribute{
					Name:        "private_networks",
					Description: `Array of Private Networks connected to the server. The structure of the block is documented below. --- The ` + "`" + `ips` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Public IP address assigned to the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP address type. Can be either ` + "`" + `public` + "`" + ` or ` + "`" + `anycast` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Boolean indicating a Primary IP flag.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `Reverse DNS assigned to the IP address.`,
				},
				resource.Attribute{
					Name:        "assignment_id",
					Description: `ID of the IP Address Assignment. The ` + "`" + `private_networks` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Private Network ID.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Private network IP address of the Cloud Server within the network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_datacenters",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results by specified key and direction. The structure of the block is documented below. --- The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `slug` + "`" + `, ` + "`" + `full_name` + "`" + `, ` + "`" + `region_id` + "`" + `, ` + "`" + `region_name` + "`" + `, ` + "`" + `region_country_code` + "`" + ``,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. The ` + "`" + `sort` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `slug` + "`" + `, ` + "`" + `full_name` + "`" + `, ` + "`" + `region_id` + "`" + `, ` + "`" + `region_name` + "`" + `, ` + "`" + `region_country_code` + "`" + ``,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Sort direction of the results. Can be one of: ` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `. Default option is ` + "`" + `desc` + "`" + `. --- ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `A list of Datacenters that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datacenter.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Datacenter name.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Datacenter slug.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `Datacenter full name.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Datacenter region ID.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Datacenter region name.`,
				},
				resource.Attribute{
					Name:        "region_country_code",
					Description: `Datacenter region country code.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenters",
					Description: `A list of Datacenters that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datacenter.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Datacenter name.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Datacenter slug.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `Datacenter full name.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Datacenter region ID.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Datacenter region name.`,
				},
				resource.Attribute{
					Name:        "region_country_code",
					Description: `Datacenter region country code.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_ips",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results by specified key and direction. The structure of the block is documented below. --- The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `ip_address` + "`" + `, ` + "`" + `type` + "`" + `, ` + "`" + `datacenter` + "`" + `, ` + "`" + `reverse_dns` + "`" + `, ` + "`" + `cloud_server_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. The ` + "`" + `sort` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `ip_address` + "`" + `, ` + "`" + `type` + "`" + `, ` + "`" + `datacenter` + "`" + `, ` + "`" + `reverse_dns` + "`" + `, ` + "`" + `cloud_server_id` + "`" + `, ` + "`" + `created_at` + "`" + ``,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Sort direction of the results. Can be one of: ` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `. Default option is ` + "`" + `desc` + "`" + `. --- ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `A list of IP addresses that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of IP address.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address value.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP address type. Can be either ` + "`" + `public` + "`" + ` or ` + "`" + `anycast` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `Datacenter Slug where IP addresses is allocated (returned only if ` + "`" + `type="public"` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `Reverse DNS assigned to the IP address.`,
				},
				resource.Attribute{
					Name:        "cloud_server_ids",
					Description: `List of Cloud Server IDs the IP addresses is assigned to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the IP address.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Boolean for the Primary IP flag. Only IPs of ` + "`" + `public` + "`" + ` type will have this flag, can contain a value only if IP is assigned to a server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ips",
					Description: `A list of IP addresses that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of IP address.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address value.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP address type. Can be either ` + "`" + `public` + "`" + ` or ` + "`" + `anycast` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `Datacenter Slug where IP addresses is allocated (returned only if ` + "`" + `type="public"` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `Reverse DNS assigned to the IP address.`,
				},
				resource.Attribute{
					Name:        "cloud_server_ids",
					Description: `List of Cloud Server IDs the IP addresses is assigned to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the IP address.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Boolean for the Primary IP flag. Only IPs of ` + "`" + `public` + "`" + ` type will have this flag, can contain a value only if IP is assigned to a server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_private_networks",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results by specified key and direction. The structure of the block is documented below. --- The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `ip_range` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `cloud_server_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. The ` + "`" + `sort` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `ip_range` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `cloud_server_id` + "`" + `, ` + "`" + `created_at` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Sort direction of the results. Can be one of: ` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `. Default option is ` + "`" + `desc` + "`" + `. --- ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "private_networks",
					Description: `A list of Private Networks that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Private Network.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `Private Network IP range in CIDR format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Private Network.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Private Network.`,
				},
				resource.Attribute{
					Name:        "cloud_servers",
					Description: `List of Cloud Servers the Private Network is connected to. The structure of the block is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Private Network. --- The ` + "`" + `cloud_servers` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cloud Server ID.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Private network IP address of the Cloud Server within the network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "private_networks",
					Description: `A list of Private Networks that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Private Network.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `Private Network IP range in CIDR format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Private Network.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Private Network.`,
				},
				resource.Attribute{
					Name:        "cloud_servers",
					Description: `List of Cloud Servers the Private Network is connected to. The structure of the block is documented below.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Private Network. --- The ` + "`" + `cloud_servers` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cloud Server ID.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Private network IP address of the Cloud Server within the network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_ssh_keys",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results by specified key and direction. The structure of the block is documented below. --- The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `fingerprint` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. The ` + "`" + `sort` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `fingerprint` + "`" + `, ` + "`" + `created_at` + "`" + ``,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Sort direction of the results. Can be one of: ` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `. Default option is ` + "`" + `desc` + "`" + `. --- ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `A list of Private Networks that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SSH key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SSH key name.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Public key.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of the SSH key.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the SSH key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `A list of Private Networks that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SSH key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SSH key name.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Public key.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of the SSH key.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the SSH key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_volume_products",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results by specified key and direction. The structure of the block is documented below. --- The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `slug` + "`" + `, ` + "`" + `price` + "`" + `, ` + "`" + `currency` + "`" + `, ` + "`" + `min_size` + "`" + `, ` + "`" + `max_size` + "`" + `, ` + "`" + `datacenter_id` + "`" + `, ` + "`" + `datacenter_name` + "`" + `, ` + "`" + `datacenter_slug` + "`" + `, ` + "`" + `datacenter_full_name` + "`" + ``,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. The ` + "`" + `sort` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `slug` + "`" + `, ` + "`" + `price` + "`" + `, ` + "`" + `currency` + "`" + `, ` + "`" + `min_size` + "`" + `, ` + "`" + `max_size` + "`" + `, ` + "`" + `datacenter_id` + "`" + `, ` + "`" + `datacenter_name` + "`" + `, ` + "`" + `datacenter_slug` + "`" + `, ` + "`" + `datacenter_full_name` + "`" + ``,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Sort direction of the results. Can be one of: ` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `. Default option is ` + "`" + `desc` + "`" + `. --- ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "products",
					Description: `A list of Products that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Product.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Slug of the Product.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the Product (per GB/month).`,
				},
				resource.Attribute{
					Name:        "currency",
					Description: `Currency for the price.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `Minimum size available for Volume creation in GB.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `Maximum size available for Volume creation in GB.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datacenter.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Datacenter name.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Datacenter slug.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `Datacenter full name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "products",
					Description: `A list of Products that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Product.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Slug of the Product.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the Product (per GB/month).`,
				},
				resource.Attribute{
					Name:        "currency",
					Description: `Currency for the price.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `Minimum size available for Volume creation in GB.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `Maximum size available for Volume creation in GB.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datacenter.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Datacenter name.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Datacenter slug.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `Datacenter full name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_volumes",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results by specified key and direction. The structure of the block is documented below. --- The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `state` + "`" + `, ` + "`" + `product_id` + "`" + `, ` + "`" + `size` + "`" + `, ` + "`" + `file_system` + "`" + `, ` + "`" + `cloud_server_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. The ` + "`" + `sort` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the results by specified key. Can be one of: ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `state` + "`" + `, ` + "`" + `product_id` + "`" + `, ` + "`" + `size` + "`" + `, ` + "`" + `file_system` + "`" + `, ` + "`" + `created_at` + "`" + ``,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Sort direction of the results. Can be one of: ` + "`" + `asc` + "`" + `, ` + "`" + `desc` + "`" + `. Default option is ` + "`" + `desc` + "`" + `. --- ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `A list of Volumes that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Volume name.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Volume.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `Product Slug that indentifies the product type of Volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Volume size in GB.`,
				},
				resource.Attribute{
					Name:        "file_system",
					Description: `File system formatting option selected on volume creation. Can be one of: ` + "`" + `ext4` + "`" + `, ` + "`" + `btrfs` + "`" + `, ` + "`" + `xfs` + "`" + `, or empty.`,
				},
				resource.Attribute{
					Name:        "cloud_server_id",
					Description: `Cloud Server ID the Volumes is attached to, if attached.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Volume.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "volumes",
					Description: `A list of Volumes that satisfy the search criteria.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Volume name.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Volume.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `Product Slug that indentifies the product type of Volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Volume size in GB.`,
				},
				resource.Attribute{
					Name:        "file_system",
					Description: `File system formatting option selected on volume creation. Can be one of: ` + "`" + `ext4` + "`" + `, ` + "`" + `btrfs` + "`" + `, ` + "`" + `xfs` + "`" + `, or empty.`,
				},
				resource.Attribute{
					Name:        "cloud_server_id",
					Description: `Cloud Server ID the Volumes is attached to, if attached.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Volume.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ah_ah_cloud_images":                       0,
		"ah_ah_cloud_server_products":              1,
		"ah_ah_cloud_server_snapshots_and_backups": 2,
		"ah_ah_cloud_servers":                      3,
		"ah_ah_datacenters":                        4,
		"ah_ah_ips":                                5,
		"ah_ah_private_networks":                   6,
		"ah_ah_ssh_keys":                           7,
		"ah_ah_volume_products":                    8,
		"ah_ah_volumes":                            9,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
