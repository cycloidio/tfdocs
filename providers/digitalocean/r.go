package digitalocean

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_cdn",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean CDN Endpoint resource.`,
			Description:      ``,
			Keywords: []string{
				"cdn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "origin",
					Description: `(Required) The fully qualified domain name, (FQDN) for a Space.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live for the CDN Endpoint, in seconds. Default is 3600 seconds.`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `(Optional) The fully qualified domain name (FQDN) of the custom subdomain used with the CDN Endpoint. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a CDN Endpoint.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `The fully qualified domain name, (FQDN) of a space referenced by the CDN Endpoint.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The fully qualified domain name (FQDN) from which the CDN-backed content is served.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the CDN Endpoint was created.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The time to live for the CDN Endpoint, in seconds.`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `The fully qualified domain name (FQDN) of the custom subdomain used with the CDN Endpoint. ## Import CDN Endpoints can be imported using the CDN ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_cdn.mycdn fb06ad00-351f-45c8-b5eb-13523c438661 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a CDN Endpoint.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `The fully qualified domain name, (FQDN) of a space referenced by the CDN Endpoint.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The fully qualified domain name (FQDN) from which the CDN-backed content is served.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the CDN Endpoint was created.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The time to live for the CDN Endpoint, in seconds.`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `The fully qualified domain name (FQDN) of the custom subdomain used with the CDN Endpoint. ## Import CDN Endpoints can be imported using the CDN ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_cdn.mycdn fb06ad00-351f-45c8-b5eb-13523c438661 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_certificate",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean Certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the certificate for identification.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of certificate to provision. Can be either ` + "`" + `custom` + "`" + ` or ` + "`" + `lets_encrypt` + "`" + `. Defaults to ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) The contents of a PEM-formatted private-key corresponding to the SSL certificate. Only valid when type is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "leaf_certificate",
					Description: `(Optional) The contents of a PEM-formatted public TLS certificate. Only valid when type is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `(Optional) The full PEM-formatted trust chain between the certificate authority's certificate and your domain's TLS certificate. Only valid when type is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `(Optional) List of fully qualified domain names (FQDNs) for which the certificate will be issued. The domains must be managed using DigitalOcean's DNS. Only valid when type is ` + "`" + `lets_encrypt` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the certificate`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the certificate`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The expiration date of the certificate`,
				},
				resource.Attribute{
					Name:        "sha1_fingerprint",
					Description: `The SHA-1 fingerprint of the certificate ## Import Certificates can be imported using the certificate ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_certificate.mycertificate 892071a0-bb95-49bc-8021-3afd67a210bf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the certificate`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the certificate`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The expiration date of the certificate`,
				},
				resource.Attribute{
					Name:        "sha1_fingerprint",
					Description: `The SHA-1 fingerprint of the certificate ## Import Certificates can be imported using the certificate ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_certificate.mycertificate 892071a0-bb95-49bc-8021-3afd67a210bf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_cluster",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean database cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database cluster.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) Database engine used by the cluster (ex. ` + "`" + `pg` + "`" + ` for PostreSQL).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Engine version used by the cluster (ex. ` + "`" + `11` + "`" + ` for PostgreSQL 11).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Database droplet size associated with the cluster (ex. ` + "`" + `db-s-1vcpu-1gb` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) DigitalOcean region where the cluster will reside.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Required) Number of nodes that will be included in the cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `(Optional) Defines when the automatic maintenance should be performed for the database cluster. ` + "`" + `maintenance_window` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `(Required) The day of the week on which to apply maintenance updates.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `(Required) The hour in UTC at which maintenance updates will be applied in 24 hour format. ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database cluster.`,
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
					Description: `Username for the cluster's default user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the cluster's default user. ## Import Database clusters can be imported using the ` + "`" + `id` + "`" + ` returned from DigitalOcean, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_cluster.mycluster 245bcfd0-7f31-4ce6-a2bc-475a116cca97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database cluster.`,
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
					Description: `Username for the cluster's default user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the cluster's default user. ## Import Database clusters can be imported using the ` + "`" + `id` + "`" + ` returned from DigitalOcean, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_cluster.mycluster 245bcfd0-7f31-4ce6-a2bc-475a116cca97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_domain",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean domain resource.`,
			Description:      ``,
			Keywords: []string{
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the domain`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address of the domain. If specified, this IP is used to created an initial A record for the domain. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the domain`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the domain ## Import Domains can be imported using the ` + "`" + `domain name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_domain.mydomain mytestdomain.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the domain`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the domain ## Import Domains can be imported using the ` + "`" + `domain name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_domain.mydomain mytestdomain.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_droplet",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean Droplet resource. This can be used to create, modify, and delete Droplets. Droplets also support provisioning.`,
			Description:      ``,
			Keywords: []string{
				"droplet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image",
					Description: `(Required) The Droplet image ID or slug.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Droplet name.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region to start in.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The unique slug that indentifies the type of Droplet. You can find a list of available slugs on [DigitalOcean API documentation](https://developers.digitalocean.com/documentation/v2/#list-all-sizes).`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `(Optional) Boolean controlling if backups are made. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional) Boolean controlling whether monitoring agent is installed. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) Boolean controlling if IPv6 is enabled. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `(Optional) Boolean controlling if private networks are enabled. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `(Optional) A list of SSH IDs or fingerprints to enable in the format ` + "`" + `[12345, 123456]` + "`" + `. To retrieve this info, use a tool such as ` + "`" + `curl` + "`" + ` with the [DigitalOcean API](https://developers.digitalocean.com/documentation/v2/#ssh-keys), to retrieve them.`,
				},
				resource.Attribute{
					Name:        "resize_disk",
					Description: `(Optional) Boolean controlling whether to increase the disk size when resizing a Droplet. It defaults to ` + "`" + `true` + "`" + `. When set to ` + "`" + `false` + "`" + `, only the Droplet's RAM and CPU will be resized.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of the tags to be applied to this Droplet.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Droplet`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the Droplet`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the Droplet`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The image of the Droplet`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Is IPv6 enabled`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The IPv4 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address_private",
					Description: `The private networking IPv4 address`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Is the Droplet locked`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `Is private networking enabled`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `Droplet hourly price`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `Droplet monthly price`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The instance size`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The size of the instance's disk in GB`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of the instance's virtual CPUs`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Droplet`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags associated with the Droplet`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `A list of the attached block storage volumes ## Import Droplets can be imported using the Droplet ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_droplet.mydroplet 100823 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Droplet`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the Droplet`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the Droplet`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The image of the Droplet`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Is IPv6 enabled`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The IPv4 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address_private",
					Description: `The private networking IPv4 address`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Is the Droplet locked`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `Is private networking enabled`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `Droplet hourly price`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `Droplet monthly price`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The instance size`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The size of the instance's disk in GB`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of the instance's virtual CPUs`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Droplet`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags associated with the Droplet`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `A list of the attached block storage volumes ## Import Droplets can be imported using the Droplet ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_droplet.mydroplet 100823 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_droplet_snapshot",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean Droplet snapshot resource.`,
			Description:      ``,
			Keywords: []string{
				"droplet",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the Droplet snapshot.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `(Required) The ID of the Droplet from which the snapshot will be taken. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the Droplet snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a Droplet to be created based on this snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the droplet snapshot is available.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the Droplet snapshot in gigabytes. ## Import Droplet Snapshots can be imported using the ` + "`" + `snapshot id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_droplet_snapshot.mysnapshot 123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the Droplet snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a Droplet to be created based on this snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the droplet snapshot is available.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the Droplet snapshot in gigabytes. ## Import Droplet Snapshots can be imported using the ` + "`" + `snapshot id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_droplet_snapshot.mysnapshot 123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_firewall",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean Cloud Firewall resource. This can be used to create, modify, and delete Firewalls.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Firewall name`,
				},
				resource.Attribute{
					Name:        "inbound_rule",
					Description: `(Optional) The inbound access rule block for the Firewall. The ` + "`" + `inbound_rule` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "outbound_rule",
					Description: `(Optional) The outbound access rule block for the Firewall. The ` + "`" + `outbound_rule` + "`" + ` block is documented below. ` + "`" + `inbound_rule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The type of traffic to be allowed. This may be one of "tcp", "udp", or "icmp".`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional) The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "1-65535" to open all ports for a protocol. Required for when protocol is ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Optional) An array of strings containing the IPv4 addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_droplet_ids",
					Description: `(Optional) An array containing the IDs of the Droplets from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_tags",
					Description: `(Optional) An array containing the names of Tags corresponding to groups of Droplets from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_load_balancer_uids",
					Description: `(Optional) An array containing the IDs of the Load Balancers from which the inbound traffic will be accepted. ` + "`" + `outbound_rule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The type of traffic to be allowed. This may be one of "tcp", "udp", or "icmp".`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional) The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "1-65535" to open all ports for a protocol. Required for when protocol is ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Optional) An array of strings containing the IPv4 addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the outbound traffic will be allowed.`,
				},
				resource.Attribute{
					Name:        "destination_droplet_ids",
					Description: `(Optional) An array containing the IDs of the Droplets to which the outbound traffic will be allowed.`,
				},
				resource.Attribute{
					Name:        "destination_tags",
					Description: `(Optional) An array containing the names of Tags corresponding to groups of Droplets to which the outbound traffic will be allowed. traffic.`,
				},
				resource.Attribute{
					Name:        "destination_load_balancer_uids",
					Description: `(Optional) An array containing the IDs of the Load Balancers to which the outbound traffic will be allowed. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Firewall.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `A status string indicating the current state of the Firewall. This can be "waiting", "succeeded", or "failed".`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the Firewall was created.`,
				},
				resource.Attribute{
					Name:        "pending_changes",
					Description: `An list of object containing the fields, "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Firewall.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `The list of the IDs of the Droplets assigned to the Firewall.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The names of the Tags assigned to the Firewall.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound access rule block for the Firewall.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound access rule block for the Firewall. ## Import Firewalls can be imported using the firewall ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_firewall.myfirewall b8ecd2ab-2267-4a5e-8692-cbf1d32583e3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Firewall.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `A status string indicating the current state of the Firewall. This can be "waiting", "succeeded", or "failed".`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the Firewall was created.`,
				},
				resource.Attribute{
					Name:        "pending_changes",
					Description: `An list of object containing the fields, "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Firewall.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `The list of the IDs of the Droplets assigned to the Firewall.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The names of the Tags assigned to the Firewall.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound access rule block for the Firewall.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound access rule block for the Firewall. ## Import Firewalls can be imported using the firewall ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_firewall.myfirewall b8ecd2ab-2267-4a5e-8692-cbf1d32583e3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_floating_ip",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean Floating IP resource.`,
			Description:      ``,
			Keywords: []string{
				"floating",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region that the Floating IP is reserved to.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `(Optional) The ID of Droplet that the Floating IP will be assigned to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address of the resource`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the floating ip ## Import Floating IPs can be imported using the ` + "`" + `ip` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_floating_ip.myip 192.168.0.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address of the resource`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the floating ip ## Import Floating IPs can be imported using the ` + "`" + `ip` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_floating_ip.myip 192.168.0.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_floating_ip_assignment",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean resource for assigning an existing floating IP to a Droplet.`,
			Description:      ``,
			Keywords: []string{
				"floating",
				"ip",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The Floating IP to assign to the Droplet.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `(Optional) The ID of Droplet that the Floating IP will be assigned to.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_kubernetes_cluster",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean Kubernetes cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The slug identifier for the region where the Kubernetes cluster will be created.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The slug identifier for the version of Kubernetes used for the cluster.`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `(Required) A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the ` + "`" + `digitalocean_kubernetes_node_pool` + "`" + ` resource. The following arguments may be specified: - ` + "`" + `name` + "`" + ` - (Required) A name for the node pool. - ` + "`" + `size` + "`" + ` - (Required) The slug identifier for the type of Droplet to be used as workers in the node pool. - ` + "`" + `node_count` + "`" + ` - (Required) The number of Droplet instances in the node pool. - ` + "`" + `tags` + "`" + ` - (Optional) A list of tag names to be applied to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tag names to be applied to the Kubernetes cluster. ## Attributes Reference In addition to the arguments listed above, the following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Kubernetes cluster.`,
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
					Description: `In addition to the arguments provided, these additional attributes about the cluster's default node pool are exported: - ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node pool. - ` + "`" + `nodes` + "`" + ` - A list of nodes in the pool. Each node exports the following attributes: + ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. + ` + "`" + `name` + "`" + ` - The auto-generated name for the node. + ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. + ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. + ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated. ## Import Kubernetes clusters can not be imported at this time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Kubernetes cluster.`,
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
					Description: `In addition to the arguments provided, these additional attributes about the cluster's default node pool are exported: - ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node pool. - ` + "`" + `nodes` + "`" + ` - A list of nodes in the pool. Each node exports the following attributes: + ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. + ` + "`" + `name` + "`" + ` - The auto-generated name for the node. + ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. + ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. + ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated. ## Import Kubernetes clusters can not be imported at this time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_kubernetes_node_pool",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean Kubernetes node pool.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"node",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the Kubernetes cluster to which the node pool is associated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the node pool.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The slug identifier for the type of Droplet to be used as workers in the node pool.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Required) The number of Droplet instances in the node pool.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tag names to be applied to the Kubernetes cluster. ## Attributes Reference In addition to the arguments listed above, the following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference the node pool.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `A list of nodes in the pool. Each node exports the following attributes: - ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. - ` + "`" + `name` + "`" + ` - The auto-generated name for the node. - ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. - ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. - ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated. ## Import Kubernetes node pools can not be imported at this time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference the node pool.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `A list of nodes in the pool. Each node exports the following attributes: - ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. - ` + "`" + `name` + "`" + ` - The auto-generated name for the node. - ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. - ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. - ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated. ## Import Kubernetes node pools can not be imported at this time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_loadbalancer",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean Load Balancer resource. This can be used to create, modify, and delete Load Balancers.`,
			Description:      ``,
			Keywords: []string{
				"loadbalancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Load Balancer name`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region to start in`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Optional) The load balancing algorithm used to determine which backend Droplet will be selected by a client. It must be either ` + "`" + `round_robin` + "`" + ` or ` + "`" + `least_connections` + "`" + `. The default value is ` + "`" + `round_robin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "forwarding_rule",
					Description: `(Required) A list of ` + "`" + `forwarding_rule` + "`" + ` to be assigned to the Load Balancer. The ` + "`" + `forwarding_rule` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `(Optional) A ` + "`" + `healthcheck` + "`" + ` block to be assigned to the Load Balancer. The ` + "`" + `healthcheck` + "`" + ` block is documented below. Only 1 healthcheck is allowed.`,
				},
				resource.Attribute{
					Name:        "sticky_sessions",
					Description: `(Optional) A ` + "`" + `sticky_sessions` + "`" + ` block to be assigned to the Load Balancer. The ` + "`" + `sticky_sessions` + "`" + ` block is documented below. Only 1 sticky_sessions block is allowed.`,
				},
				resource.Attribute{
					Name:        "redirect_http_to_https",
					Description: `(Optional) A boolean value indicating whether HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_proxy_protocol",
					Description: `(Optional) A boolean value indicating whether PROXY Protocol should be used to pass information from connecting client requests to the backend service. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "entry_protocol",
					Description: `(Required) The protocol used for traffic to the Load Balancer. The possible values are: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, or ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "entry_port",
					Description: `(Required) An integer representing the port on which the Load Balancer instance will listen.`,
				},
				resource.Attribute{
					Name:        "target_protocol",
					Description: `(Required) The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, or ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `(Required) An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) The ID of the TLS certificate to be used for SSL termination.`,
				},
				resource.Attribute{
					Name:        "tls_passthrough",
					Description: `(Optional) A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is ` + "`" + `false` + "`" + `. ` + "`" + `sticky_sessions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are ` + "`" + `cookies` + "`" + ` or ` + "`" + `none` + "`" + `. If not specified, the default value is ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Optional) The name to be used for the cookie sent to the client. This attribute is required when using ` + "`" + `cookies` + "`" + ` for the sticky sessions type.`,
				},
				resource.Attribute{
					Name:        "cookie_ttl_seconds",
					Description: `(Optional) The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using ` + "`" + `cookies` + "`" + ` for the sticky sessions type. ` + "`" + `healthcheck` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol used for health checks sent to the backend Droplets. The possible values are ` + "`" + `http` + "`" + ` or ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) An integer representing the port on the backend Droplets on which the health check will attempt a connection.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path on the backend Droplets to which the Load Balancer instance will send a request.`,
				},
				resource.Attribute{
					Name:        "check_interval_seconds",
					Description: `(Optional) The number of seconds between between two consecutive health checks. If not specified, the default value is ` + "`" + `10` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_timeout_seconds",
					Description: `(Optional) The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool. If not specified, the default value is ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool. If not specified, the default value is ` + "`" + `5` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the Load Balancer ## Import Load Balancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_loadbalancer.myloadbalancer 4de7ac8b-495b-4884-9a69-1050c6793cd6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the Load Balancer ## Import Load Balancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_loadbalancer.myloadbalancer 4de7ac8b-495b-4884-9a69-1050c6793cd6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_project",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean Project resource.`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Project`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) the description of the project`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `(Optional) the purpose of the project, (Default: "Web Application")`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) the environment of the project's resources. The possible values are: ` + "`" + `Development` + "`" + `, ` + "`" + `Staging` + "`" + `, ` + "`" + `Production` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `a list of uniform resource names (URNs) for the resources associated with the project ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the project`,
				},
				resource.Attribute{
					Name:        "owner_uuid",
					Description: `the unique universal identifier of the project owner.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `the id of the project owner.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `the date and time when the project was created, (ISO8601)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `the date and time when the project was last updated, (ISO8601)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the project`,
				},
				resource.Attribute{
					Name:        "owner_uuid",
					Description: `the unique universal identifier of the project owner.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `the id of the project owner.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `the date and time when the project was created, (ISO8601)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `the date and time when the project was last updated, (ISO8601)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_record",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean DNS record resource.`,
			Description:      ``,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of record. Must be one of ` + "`" + `A` + "`" + `, ` + "`" + `AAAA` + "`" + `, ` + "`" + `CAA` + "`" + `, ` + "`" + `CNAME` + "`" + `, ` + "`" + `MX` + "`" + `, ` + "`" + `NS` + "`" + `, ` + "`" + `TXT` + "`" + `, or ` + "`" + `SRV` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain to add the record to.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the record.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port of the record. Only valid when type is ` + "`" + `SRV` + "`" + `. Must be between 1 and 65535.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the record. Only valid when type is ` + "`" + `MX` + "`" + ` or ` + "`" + `SRV` + "`" + `. Must be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The weight of the record. Only valid when type is ` + "`" + `SRV` + "`" + `. Must be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live for the record, in seconds. Must be at least 0.`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `(Optional) The flags of the record. Only valid when type is ` + "`" + `CAA` + "`" + `. Must be between 0 and 255.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) The tag of the record. Only valid when type is ` + "`" + `CAA` + "`" + `. Must be one of ` + "`" + `issue` + "`" + `, ` + "`" + `issuewild` + "`" + `, or ` + "`" + `iodef` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the record ## Import Records can be imported using the domain name and record ` + "`" + `id` + "`" + ` when joined with a comma. See the following example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_record.example_record example.com,12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the record ## Import Records can be imported using the domain name and record ` + "`" + `id` + "`" + ` when joined with a comma. See the following example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_record.example_record example.com,12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_bucket",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean Spaces Bucket resource.`,
			Description:      ``,
			Keywords: []string{
				"spaces",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the bucket resides (Defaults to ` + "`" + `nyc3` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Canned ACL applied on bucket creation (` + "`" + `private` + "`" + ` or ` + "`" + `public-read` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `Unless ` + "`" + `true` + "`" + `, the bucket will only be destroyed if empty (Defaults to ` + "`" + `false` + "`" + `) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the bucket`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The name of the region`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com) ## Import Buckets can be imported using the ` + "`" + `region` + "`" + ` and ` + "`" + `name` + "`" + ` attributes (delimited by a comma): ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_spaces_bucket.foobar ` + "`" + `region` + "`" + `,` + "`" + `name` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the bucket`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The name of the region`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com) ## Import Buckets can be imported using the ` + "`" + `region` + "`" + ` and ` + "`" + `name` + "`" + ` attributes (delimited by a comma): ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_spaces_bucket.foobar ` + "`" + `region` + "`" + `,` + "`" + `name` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean SSH key resource.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SSH key for identification`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The public key. If this is a file, it can be read using the file interpolation function ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key ## Import SSH Keys can be imported using the ` + "`" + `ssh key id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_ssh_key.mykey 263654 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key ## Import SSH Keys can be imported using the ` + "`" + `ssh key id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_ssh_key.mykey 263654 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_tag",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean Tag resource.`,
			Description:      ``,
			Keywords: []string{
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tag ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the tag`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag ## Import Tags can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_tag.mytag tagname ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the tag`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag ## Import Tags can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_tag.mytag tagname ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_volume",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean volume resource.`,
			Description:      ``,
			Keywords: []string{
				"volume",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region that the block storage volume will be created in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the block storage volume in GiB. If updated, can only be expanded.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A free-form text field up to a limit of 1024 bytes to describe a block storage volume.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limitied on creation to that of the referenced snapshot`,
				},
				resource.Attribute{
					Name:        "initial_filesystem_type",
					Description: `(Optional) Initial filesystem type (` + "`" + `xfs` + "`" + ` or ` + "`" + `ext4` + "`" + `) for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "initial_filesystem_label",
					Description: `(Optional) Initial filesystem label for the block storage volume. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_type",
					Description: `Filesystem type (` + "`" + `xfs` + "`" + ` or ` + "`" + `ext4` + "`" + `) for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_label",
					Description: `Filesystem label for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `A list of associated droplet ids. ## Import Volumes can be imported using the ` + "`" + `volume id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_volume.volume 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_type",
					Description: `Filesystem type (` + "`" + `xfs` + "`" + ` or ` + "`" + `ext4` + "`" + `) for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_label",
					Description: `Filesystem label for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `A list of associated droplet ids. ## Import Volumes can be imported using the ` + "`" + `volume id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_volume.volume 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_volume_attachment",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean volume attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"volume",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "droplet_id",
					Description: `(Required) ID of the Droplet to attach the volume to.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) ID of the Volume to be attached to the Droplet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the volume attachment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the volume attachment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_volume_snapshot",
			Category:         "Resources",
			ShortDescription: `Provides a DigitalOcean volume snapshot resource.`,
			Description:      ``,
			Keywords: []string{
				"volume",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the volume snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) The ID of the volume from which the volume snapshot originated. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "size",
					Description: `The billable size of the volume snapshot in gigabytes. ## Import Volume Snapshots can be imported using the ` + "`" + `snapshot id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_volume_snapshot.snapshot 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "size",
					Description: `The billable size of the volume snapshot in gigabytes. ## Import Volume Snapshots can be imported using the ` + "`" + `snapshot id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_volume_snapshot.snapshot 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"digitalocean_cdn":                    0,
		"digitalocean_certificate":            1,
		"digitalocean_database_cluster":       2,
		"digitalocean_domain":                 3,
		"digitalocean_droplet":                4,
		"digitalocean_droplet_snapshot":       5,
		"digitalocean_firewall":               6,
		"digitalocean_floating_ip":            7,
		"digitalocean_floating_ip_assignment": 8,
		"digitalocean_kubernetes_cluster":     9,
		"digitalocean_kubernetes_node_pool":   10,
		"digitalocean_loadbalancer":           11,
		"digitalocean_project":                12,
		"digitalocean_record":                 13,
		"digitalocean_spaces_bucket":          14,
		"digitalocean_ssh_key":                15,
		"digitalocean_tag":                    16,
		"digitalocean_volume":                 17,
		"digitalocean_volume_attachment":      18,
		"digitalocean_volume_snapshot":        19,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
