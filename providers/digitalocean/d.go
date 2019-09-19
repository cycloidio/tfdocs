package digitalocean

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

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
					Name:        "host",
					Description: `Database cluster's hostname.`,
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
					Name:        "database",
					Description: `Name of the cluster's default database.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Username for the cluster's default user. ` + "`" + `maintenance_window` + "`" + ` supports the following:`,
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
					Name:        "host",
					Description: `Database cluster's hostname.`,
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
					Name:        "database",
					Description: `Name of the cluster's default database.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Username for the cluster's default user. ` + "`" + `maintenance_window` + "`" + ` supports the following:`,
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
					Name:        "name",
					Description: `(Required) The name of Droplet. ## Attributes Reference The following attributes are exported:`,
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
			ShortDescription: `Get information on an snapshot.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the private image.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Optional) The slug of the official image. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The id of the image.`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `The name of the distribution of the OS of the image.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `Is image a public image or not. Public images represent Linux distributions or One-Click Applications, while non-public images represent snapshots and backups and are only available within your account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "image",
					Description: `The id of the image.`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `The name of the distribution of the OS of the image.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `Is image a public image or not. Public images represent Linux distributions or One-Click Applications, while non-public images represent snapshots and backups and are only available within your account.`,
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
					Description: `A representation of the Kubernetes cluster's kubeconfig with the following attributes: - ` + "`" + `raw_config` + "`" + ` - The full contents of the Kubernetes cluster's kubeconfig file. - ` + "`" + `host` + "`" + ` - The URL of the API server on the Kubernetes master node. - ` + "`" + `client_key` + "`" + ` - The base64 encoded private key used by clients to access the cluster. - ` + "`" + `client_certificate` + "`" + ` - The base64 encoded public certificate used by clients to access the cluster. - ` + "`" + `cluster_ca_certificate` + "`" + ` - The base64 encoded public certificate for the cluster's certificate authority.`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `A list of node pools associated with the cluster. Each node pool exports the following attributes: - ` + "`" + `id` + "`" + ` - The unique ID that can be used to identify and reference the node pool. - ` + "`" + `name` + "`" + ` - The name of the node pool. - ` + "`" + `size` + "`" + ` - The slug identifier for the type of Droplet used as workers in the node pool. - ` + "`" + `node_count` + "`" + ` - The number of Droplet instances in the node pool. - ` + "`" + `tags` + "`" + ` - A list of tag names applied to the node pool. - ` + "`" + `nodes` + "`" + ` - A list of nodes in the pool. Each node exports the following attributes: + ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. + ` + "`" + `name` + "`" + ` - The auto-generated name for the node. + ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. + ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. + ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated.`,
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
					Description: `A representation of the Kubernetes cluster's kubeconfig with the following attributes: - ` + "`" + `raw_config` + "`" + ` - The full contents of the Kubernetes cluster's kubeconfig file. - ` + "`" + `host` + "`" + ` - The URL of the API server on the Kubernetes master node. - ` + "`" + `client_key` + "`" + ` - The base64 encoded private key used by clients to access the cluster. - ` + "`" + `client_certificate` + "`" + ` - The base64 encoded public certificate used by clients to access the cluster. - ` + "`" + `cluster_ca_certificate` + "`" + ` - The base64 encoded public certificate for the cluster's certificate authority.`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `A list of node pools associated with the cluster. Each node pool exports the following attributes: - ` + "`" + `id` + "`" + ` - The unique ID that can be used to identify and reference the node pool. - ` + "`" + `name` + "`" + ` - The name of the node pool. - ` + "`" + `size` + "`" + ` - The slug identifier for the type of Droplet used as workers in the node pool. - ` + "`" + `node_count` + "`" + ` - The number of Droplet instances in the node pool. - ` + "`" + `tags` + "`" + ` - A list of tag names applied to the node pool. - ` + "`" + `nodes` + "`" + ` - A list of nodes in the pool. Each node exports the following attributes: + ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. + ` + "`" + `name` + "`" + ` - The auto-generated name for the node. + ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. + ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. + ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated.`,
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
			},
			Attributes: []resource.Attribute{},
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
			},
		},
	}

	dataSourcesMap = map[string]int{

		"digitalocean_certificate":        0,
		"digitalocean_database_cluster":   1,
		"digitalocean_domain":             2,
		"digitalocean_droplet":            3,
		"digitalocean_droplet_snapshot":   4,
		"digitalocean_floating_ip":        5,
		"digitalocean_image":              6,
		"digitalocean_kubernetes_cluster": 7,
		"digitalocean_loadbalancer":       8,
		"digitalocean_record":             9,
		"digitalocean_ssh_key":            10,
		"digitalocean_tag":                11,
		"digitalocean_volume":             12,
		"digitalocean_volume_snapshot":    13,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
