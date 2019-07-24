package oneandone

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "oneandone_baremetal",
			Category:         "Resources",
			ShortDescription: `Creates and manages 1&1 Baremetal Server.`,
			Description:      ``,
			Keywords: []string{
				"baremetal",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) Location of desired 1and1 datacenter. Can be ` + "`" + `DE` + "`" + `, ` + "`" + `GB` + "`" + `, ` + "`" + `US` + "`" + ` or ` + "`" + `ES` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the server`,
				},
				resource.Attribute{
					Name:        "firewall_policy_id",
					Description: `(Optional) ID of firewall policy`,
				},
				resource.Attribute{
					Name:        "baremetal_model_id",
					Description: `(Required) ID of a baremetal model`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP address for the server`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Optional) ID of the load balancer`,
				},
				resource.Attribute{
					Name:        "monitoring_policy_id",
					Description: `(Optional) ID of monitoring policy`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Desired password.`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional) Path to private ssh key`,
				},
				resource.Attribute{
					Name:        "ssh_key_public",
					Description: `(Optional) The public key data in OpenSSH authorized_keys format. IPs (` + "`" + `ips` + "`" + `) expose the following attributes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the attached IP`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Computed) The IP`,
				},
				resource.Attribute{
					Name:        "firewall_policy_id",
					Description: `(Computed) The attached firewall policy`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oneandone_block_storage",
			Category:         "Resources",
			ShortDescription: `Creates and manages 1&1 Block Storage.`,
			Description:      ``,
			Keywords: []string{
				"block",
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) Location of desired 1and1 datacenter, where the block storage will be created. Can be ` + "`" + `DE` + "`" + `, ` + "`" + `GB` + "`" + `, ` + "`" + `US` + "`" + ` or ` + "`" + `ES` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the block storage`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the storage`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Optional) ID of the server that the block storage will be attached to`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Size of the block storage (` + "`" + `min: 20, max: 500, multipleOf: 10` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oneandone_firewall_policy",
			Category:         "Resources",
			ShortDescription: `Creates and manages 1&1 Firewall Policy.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the VPN`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPN. Firewall Policy Rules (` + "`" + `rules` + "`" + `) support the follwing:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol for the rule. Allowed values are ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `TCP/UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` and ` + "`" + `IPSEC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_from",
					Description: `(Optional) Defines the start range of the allowed port`,
				},
				resource.Attribute{
					Name:        "port_to",
					Description: `(Optional) Defines the end range of the allowed port`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `(Optional) Only traffic directed to the respective IP address`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oneandone_loadbalancer",
			Category:         "Resources",
			ShortDescription: `Creates and manages 1&1 Load Balancer.`,
			Description:      ``,
			Keywords: []string{
				"loadbalancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the load balancer`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) Balancing procedure Can be ` + "`" + `ROUND_ROBIN` + "`" + ` or ` + "`" + `LEAST_CONNECTIONS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) Location of desired 1and1 datacenter. Can be ` + "`" + `DE` + "`" + `, ` + "`" + `GB` + "`" + `, ` + "`" + `US` + "`" + ` or ` + "`" + `ES` + "`" + ``,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `(Optional) True/false defines whether persistence should be turned on/off`,
				},
				resource.Attribute{
					Name:        "persistence_time",
					Description: `(Optional) Persistence duration in seconds`,
				},
				resource.Attribute{
					Name:        "health_check_test",
					Description: `(Optional) Can be ` + "`" + `TCP` + "`" + ` or` + "`" + `ICMP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "health_check_path_parser",
					Description: `(Optional) Loadbalancer rules (` + "`" + `rules` + "`" + `) support the following`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol for the rule. Allowed values are ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `TCP/UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` and ` + "`" + `IPSEC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_balancer",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "port_server",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `(Required)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oneandone_monitoring_policy",
			Category:         "Resources",
			ShortDescription: `Creates and manages 1&1 Monitoring Policy.`,
			Description:      ``,
			Keywords: []string{
				"monitoring",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPN.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the VPN`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Email address to which notifications monitoring system will send`,
				},
				resource.Attribute{
					Name:        "email_notification",
					Description: `(Required) If set true email will be sent.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port number.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol of the port. Allowed values are ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `TCP/UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` and ` + "`" + `IPSEC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "alert_if",
					Description: `(Required) Condition for the alert to be issued. Monitoring Policy Ports (` + "`" + `processes` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "email_notification",
					Description: `(Required) If set true email will be sent.`,
				},
				resource.Attribute{
					Name:        "process",
					Description: `(Required) Process name.`,
				},
				resource.Attribute{
					Name:        "alert_if",
					Description: `(Required) Condition for the alert to be issued.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oneandone_private_network",
			Category:         "Resources",
			ShortDescription: `Creates and manages 1&1 Private Network.`,
			Description:      ``,
			Keywords: []string{
				"private",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) Location of desired 1and1 datacenter. Can be ` + "`" + `DE` + "`" + `, ` + "`" + `GB` + "`" + `, ` + "`" + `US` + "`" + ` or ` + "`" + `ES` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the shared storage`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the private network`,
				},
				resource.Attribute{
					Name:        "network_address",
					Description: `(Optional) Network address for the private network`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `(Optional) Subnet mask for the private network`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oneandone_public_ip",
			Category:         "Resources",
			ShortDescription: `Creates and manages 1&1 Public IP.`,
			Description:      ``,
			Keywords: []string{
				"public",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_type",
					Description: `(Required) IP type. Can be ` + "`" + `IPV4` + "`" + ` or ` + "`" + `IPV6` + "`" + ``,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) Location of desired 1and1 datacenter. Can be ` + "`" + `DE` + "`" + `, ` + "`" + `GB` + "`" + `, ` + "`" + `US` + "`" + ` or ` + "`" + `ES` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Computed) The IP address.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oneandone_server",
			Category:         "Resources",
			ShortDescription: `Creates and manages 1&1 Server.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) Location of desired 1and1 datacenter. Can be ` + "`" + `DE` + "`" + `, ` + "`" + `GB` + "`" + `, ` + "`" + `US` + "`" + ` or ` + "`" + `ES` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the server`,
				},
				resource.Attribute{
					Name:        "firewall_policy_id",
					Description: `(Optional) ID of firewall policy`,
				},
				resource.Attribute{
					Name:        "fixed_instance_size",
					Description: `(Optional) ID of a fixed instance size`,
				},
				resource.Attribute{
					Name:        "hdds",
					Description: `(Optional) List of HDDs. One HDD must be main.`,
				},
				resource.Attribute{
					Name:        "*is_main",
					Description: `(Optional) Indicates if HDD is to be used as main hard disk of the server`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP address for the server`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Optional) ID of the load balancer`,
				},
				resource.Attribute{
					Name:        "monitoring_policy_id",
					Description: `(Optional) ID of monitoring policy`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Desired password.`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional) Path to private ssh key`,
				},
				resource.Attribute{
					Name:        "ssh_key_public",
					Description: `(Optional) The public key data in OpenSSH authorized_keys format.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the attached IP`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Computed) The IP`,
				},
				resource.Attribute{
					Name:        "firewall_policy_id",
					Description: `(Computed) The attached firewall policy`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oneandone_shared_storage",
			Category:         "Resources",
			ShortDescription: `Creates and manages 1&1 Shared Storage.`,
			Description:      ``,
			Keywords: []string{
				"shared",
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the storage`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) Location of desired 1and1 datacenter. Can be ` + "`" + `DE` + "`" + `, ` + "`" + `GB` + "`" + `, ` + "`" + `US` + "`" + ` or ` + "`" + `ES` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the shared storage`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Size of the shared storage`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the server`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `(Required) Access rights to be assigned to the server. Can be ` + "`" + `RW` + "`" + ` or ` + "`" + `R` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oneandone_ssh_key",
			Category:         "Resources",
			ShortDescription: `Creates and manages 1&1 SSH Key.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the ssh key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the storage`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) Public key to import. If not given, new SSH key pair will be created and the private key is returned in the response`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oneandone_vpn",
			Category:         "Resources",
			ShortDescription: `Creates and manages 1&1 VPN.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) Location of desired 1and1 datacenter. Can be ` + "`" + `DE` + "`" + `, ` + "`" + `GB` + "`" + `, ` + "`" + `US` + "`" + ` or ` + "`" + `ES` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPN`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "download_path",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Optional)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"oneandone_baremetal":         0,
		"oneandone_block_storage":     1,
		"oneandone_firewall_policy":   2,
		"oneandone_loadbalancer":      3,
		"oneandone_monitoring_policy": 4,
		"oneandone_private_network":   5,
		"oneandone_public_ip":         6,
		"oneandone_server":            7,
		"oneandone_shared_storage":    8,
		"oneandone_ssh_key":           9,
		"oneandone_vpn":               10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
