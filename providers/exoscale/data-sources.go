package exoscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "exoscale_affinity",
			Category:         "Deprecated",
			ShortDescription: `Provides information about an Anti-Affinity Group.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"affinity",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_anti_affinity_group",
			Category:         "Data Sources",
			ShortDescription: `Provides information about an Anti-Affinity Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instances",
					Description: `A list of Compute instance IDs belonging to the Anti-Affinity Group. [aag-doc]: https://community.exoscale.com/documentation/compute/anti-affinity-groups/ [r-compute]: ../resources/compute`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_compute",
			Category:         "Deprecated",
			ShortDescription: `Provides information about a Compute.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"compute",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `Creation date of the Compute instance.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Name of the zone.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `Name of the template.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Current size of the Compute instance.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Size of the Compute instance disk.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of cpu the Compute instance is running with.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory allocated for the Compute instance.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the Compute instance.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Public IPv4 address of the Compute instance.`,
				},
				resource.Attribute{
					Name:        "ip6_address",
					Description: `Public IPv6 address of the Compute instance (if IPv6 is enabled).`,
				},
				resource.Attribute{
					Name:        "private_network_ip_addresses",
					Description: `List of Compute private IP addresses (in managed Private Networks only). [compute-doc]: https://www.exoscale.com/compute/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_compute_instance",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Compute instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "anti_affinity_group_ids",
					Description: `A list of [Anti-Affinity Group][r-anti_affinity_group] IDs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation date of the Compute instance.`,
				},
				resource.Attribute{
					Name:        "deploy_target_id",
					Description: `A Deploy Target ID.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The Compute instance disk size in GiB.`,
				},
				resource.Attribute{
					Name:        "elastic_ip_ids",
					Description: `A list of [Elastic IP][r-elastic_ip] IDs attached to the Compute instance.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The IPv6 address of the Compute instance main network interface.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Whether IPv6 is enabled on the Compute instance.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of key/value labels.`,
				},
				resource.Attribute{
					Name:        "manager_id",
					Description: `The ID of the Compute instance manager, if any.`,
				},
				resource.Attribute{
					Name:        "manager_type",
					Description: `The type of Compute instance manager, if any.`,
				},
				resource.Attribute{
					Name:        "private_network_ids",
					Description: `A list of [Private Network][r-private_network] IDs attached to the Compute instance.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `The IPv4 address of the Compute instance's main network interface.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `A list of [Security Group][r-security_group] IDs attached to the Compute instance.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `The name of the [SSH key pair][sshkeypair] installed to the Compute instance's user account during creation.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Compute instance.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The ID of the instance [template][template] used when creating the Compute instance.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Compute instance [type][type].`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `A [cloud-init][cloudinit] configuration. [cloudinit]: http://cloudinit.readthedocs.io/en/latest/ [compute-doc]: https://community.exoscale.com/documentation/compute/ [r-anti_affinity_group]: ../resources/anti_affinity_group [r-elastic_ip]: ../resources/elastic_ip [r-private_network]: ../resources/private_network [r-security_group]: ../resources/security_group [sshkeypair-doc]: https://community.exoscale.com/documentation/compute/ssh-keypairs/ [template]: https://www.exoscale.com/templates/ [type]: https://www.exoscale.com/pricing/#/compute/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_compute_ipaddress",
			Category:         "Deprecated",
			ShortDescription: `Provides information about a Compute template.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"compute",
				"ipaddress",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_compute_template",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Compute template.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `Username to use to log into a Compute Instance based on this template [r-compute]: ../resources/compute [templates]: https://www.exoscale.com/templates/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_domain",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain [exo-dns]: https://www.exoscale.com/dns/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_domain_record",
			Category:         "Data Sources",
			ShortDescription: `Provides information about an Exoscale DNS domain record.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `The content of the domain record.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `The priority of the domain record. [exo-dns]: https://www.exoscale.com/dns/ [r-domain]: ../resources/domain [r-domain_record]: ../resources/domain_record`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_elastic_ip",
			Category:         "Data Sources",
			ShortDescription: `Provides information about an Elastic IP.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Elastic IP.`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `A health checking configuration for managed Elastic IPs. Structure is documented below. The ` + "`" + `healthcheck` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The health checking mode.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The health checking port.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The health checking URI.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The health checking interval in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The time in seconds before considering a healthcheck probing failed.`,
				},
				resource.Attribute{
					Name:        "strikes_ok",
					Description: `The number of successful attempts before considering a managed Elastic IP target healthy.`,
				},
				resource.Attribute{
					Name:        "strikes_fail",
					Description: `The number of failed attempts before considering a managed Elastic IP target unhealthy.`,
				},
				resource.Attribute{
					Name:        "tls_sni",
					Description: `The health checking server name to present with SNI in ` + "`" + `https` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "tls_skip_verify",
					Description: `Disable TLS certificate verification for health checking in ` + "`" + `https` + "`" + ` mode. [eip-doc]: https://community.exoscale.com/documentation/compute/eip/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_network",
			Category:         "Deprecated",
			ShortDescription: `Provides information about a Private Network.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"network",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Private Network`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `The first address of IP range used by the DHCP service to automatically assign (for`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `The last address of the IP range used by the DHCP service (for`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The netmask defining the IP network allowed for the static lease (for`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_nlb",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Network Load Balancer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NLB.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the NLB.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation date of the NLB.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The public IP address of the NLB. [nlb-doc]: https://community.exoscale.com/documentation/compute/network-load-balancer/ [r-instance_pool]: ../resources/instance_pool [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_private_network",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Private Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Private Network.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `The first address of IP range used by the DHCP service to automatically assign (for`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `The last address of the IP range used by the DHCP service (for`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The netmask defining the IP network allowed for the static lease (for`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_security_group",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Security Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"exoscale_affinity":            0,
		"exoscale_anti_affinity_group": 1,
		"exoscale_compute":             2,
		"exoscale_compute_instance":    3,
		"exoscale_compute_ipaddress":   4,
		"exoscale_compute_template":    5,
		"exoscale_domain":              6,
		"exoscale_domain_record":       7,
		"exoscale_elastic_ip":          8,
		"exoscale_network":             9,
		"exoscale_nlb":                 10,
		"exoscale_private_network":     11,
		"exoscale_security_group":      12,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
