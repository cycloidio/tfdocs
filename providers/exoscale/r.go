package exoscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "exoscale_affinity",
			Category:         "Resources Data",
			ShortDescription: `Provides an Exoscale Anti-Affinity Group resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"affinity",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Anti-Affinity Group.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_ids",
					Description: `The IDs of the Compute instance resources member of the Anti-Affinity Group. ## Import An existing Anti-Affinity Group can be imported as a resource by name or ID: ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_affinity.mygroup mygroup # By ID $ terraform import exoscale_affinity.mygroup eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` [aag-doc]: https://community.exoscale.com/documentation/compute/anti-affinity-groups/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_compute",
			Category:         "Resources Data",
			ShortDescription: `Provides an Exoscale Compute instance resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"compute",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Compute instance.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The user to use to connect to the Compute instance with SSH. If you've referenced a`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The initial Compute instance password and/or encrypted password.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address of the Compute instance main network interface.`,
				},
				resource.Attribute{
					Name:        "ip6_address",
					Description: `The IPv6 address of the Compute instance main network interface. ## ` + "`" + `remote-exec` + "`" + ` provisioner usage If you wish to log to a ` + "`" + `exoscale_compute` + "`" + ` resource using the [` + "`" + `remote-exec` + "`" + `][remote-exec] provisioner, make sure to explicity set the SSH ` + "`" + `user` + "`" + ` setting to connect to the instance to the actual template username returned by the [` + "`" + `exoscale_compute_template` + "`" + `][compute_template] data source: ` + "`" + `` + "`" + `` + "`" + `hcl data "exoscale_compute_template" "ubuntu" { zone = "ch-gva-2" name = "Linux Ubuntu 18.04 LTS 64-bit" } resource "exoscale_compute" "mymachine" { zone = "ch-gva-2" display_name = "mymachine" template_id = data.exoscale_compute_template.ubuntu.id size = "Medium" disk_size = 10 key_pair = "me@mymachine" state = "Running" provisioner "remote-exec" { connection { type = "ssh" host = self.ip_address user = data.exoscale_compute_template.ubuntu.username } } } ` + "`" + `` + "`" + `` + "`" + ` ## Import An existing Compute instance can be imported as a resource by name or ID: ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_compute.vm1 vm1 # By ID $ terraform import exoscale_compute.vm1 eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_domain",
			Category:         "Resources Data",
			ShortDescription: `Provides an Exoscale DNS Domain resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"domain",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "token",
					Description: `A security token that can be used as an alternative way to manage DNS Domains via the Exoscale API.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the DNS Domain.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Boolean indicating that the DNS Domain has automatic renewal enabled.`,
				},
				resource.Attribute{
					Name:        "expires_on",
					Description: `The date of expiration of the DNS Domain, if known. ## Import An existing DNS Domain can be imported as a resource by name: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_domain.example example.net ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_domain_record",
			Category:         "Resources Data",
			ShortDescription: `Provides an Exoscale DNS domain record resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"domain",
				"record",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `The DNS domain record's`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_instance_pool",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Instance Pool resource.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of the [zone][zone] to deploy the Instance Pool into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Instance Pool.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Required) The ID of the instance [template][template] to use when creating Compute instances. Usage of the [` + "`" + `compute_template` + "`" + `][d-compute_template] data source is recommended.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The number of Compute instance members the Instance Pool manages.`,
				},
				resource.Attribute{
					Name:        "service_offering",
					Description: `(Required) The managed Compute instances [size][size], e.g. ` + "`" + `tiny` + "`" + `, ` + "`" + `small` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `large` + "`" + ` etc.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The managed Compute instances disk size.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Instance Pool.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `A [cloud-init][cloudinit] configuration to apply when creating Compute instances. Whenever possible don't base64-encode neither gzip it yourself, as this will be automatically taken care of on your behalf by the provider.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `The name of the [SSH key pair][sshkeypair] to install when creating Compute instances.`,
				},
				resource.Attribute{
					Name:        "instance_prefix",
					Description: `The string to add as prefix to managed Compute instances name (default ` + "`" + `pool` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "affinity_group_ids",
					Description: `A list of [Anti-Affinity Group][r-affinity] IDs (at creation time only).`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `A list of [Security Group][r-security_group] IDs (at creation time only).`,
				},
				resource.Attribute{
					Name:        "network_ids",
					Description: `A list of [Private Network][privnet-doc] IDs.`,
				},
				resource.Attribute{
					Name:        "elastic_ip_ids",
					Description: `A list of [Elastic IP][eip-doc] IDs.`,
				},
				resource.Attribute{
					Name:        "deploy_target_id",
					Description: `A Deploy Target ID. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_ipaddress",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Elastic IP address.`,
			Description:      ``,
			Keywords: []string{
				"ipaddress",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `The Elastic IP address. ## Import An existing Elastic IP can be imported as a resource by address or ID: ` + "`" + `` + "`" + `` + "`" + `console # By address $ terraform import exoscale_ipaddress.myip 159.100.251.224 # By ID $ terraform import exoscale_ipaddress.myip eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` [eip-doc]: https://community.exoscale.com/documentation/compute/eip/ [r-secondary_ipaddress]: secondary_ipaddress.html [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_network",
			Category:         "Resources Data",
			ShortDescription: `Provides an Exoscale Private Network.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"network",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Network. ## Import An existing Private Network can be imported as a resource by name or ID: ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_network.net myprivnet # By ID $ terraform import exoscale_network.net 04fb76a2-6d22-49be-8da7-f2a5a0b902e1 ` + "`" + `` + "`" + `` + "`" + ` [r-nic]: nic.html [privnet-doc]: https://community.exoscale.com/documentation/compute/private-networks/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_nic",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Compute instance Private Network Interface (NIC).`,
			Description:      ``,
			Keywords: []string{
				"nic",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Compute instance NIC.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The physical address (MAC) of the Compute instance NIC. ## Import This resource is automatically imported when importing an ` + "`" + `exoscale_compute` + "`" + ` resource. [privnet-doc]: https://community.exoscale.com/documentation/compute/private-networks/ [r-compute]: compute.html [r-network]: network.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_nlb",
			Category:         "Resources Data",
			ShortDescription: `Provides an Exoscale Network Load Balancer resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"nlb",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NLB.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The public IP address of the NLB.`,
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
					Name:        "services",
					Description: `The list of the NLB service names. ## Import An existing NLB can be imported as a resource by ID: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_nlb.website eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_nlb_service",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Network Load Balancer service resource.`,
			Description:      ``,
			Keywords: []string{
				"nlb",
				"service",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NLB service. ## Import An existing NLB service can be imported as a resource by ID. ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_nlb_service.website 9ecc6b8b-73d4-4211-8ced-f7f29bb79524 ` + "`" + `` + "`" + `` + "`" + ` [r-nlb]: nlb.html [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_secondary_ipaddress",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale resource for assigning an existing Elastic IP to a Compute instance.`,
			Description:      ``,
			Keywords: []string{
				"secondary",
				"ipaddress",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The ID of the Network the Compute instance NIC is attached to. ## Import This resource is automatically imported when importing an ` + "`" + `exoscale_compute` + "`" + ` resource. [r-compute]: compute.html [r-ipaddress]: ipaddress.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_security_group",
			Category:         "Resources Data",
			ShortDescription: `Provides an Exoscale Security Group.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Security Group. ## Import An existing Security Group can be imported as a resource by name or ID: ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_security_group.http http # By ID $ terraform import exoscale_security_group.http eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_security_group_rule",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Security Group Rule.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
				"rule",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_security_group_rules",
			Category:         "Resources",
			ShortDescription: `Provides a resource for assigning multiple rules to an existing Exoscale Security Group.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
				"rules",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_sks_cluster",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale SKS cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"sks",
				"cluster",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SKS cluster.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The Kubernetes public API endpoint of the SKS cluster.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the SKS cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation date of the SKS cluster.`,
				},
				resource.Attribute{
					Name:        "nodepools",
					Description: `The list of [SKS Nodepools][r-sks_nodepool] (IDs) attached to the SKS cluster. ## Import An existing SKS cluster can be imported as a resource by ID: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_sks_cluster.prod eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_sks_nodepool",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale SKS Nodepool resource.`,
			Description:      ``,
			Keywords: []string{
				"sks",
				"nodepool",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SKS Nodepool.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the SKS Nodepool.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation date of the SKS Nodepool.`,
				},
				resource.Attribute{
					Name:        "instance_pool_id",
					Description: `The ID of the Instance Pool managed by the SKS Nodepool.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The ID of the Compute instance template used by the SKS Nodepool members.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Kubernetes version of the SKS Nodepool members. ## Import An existing SKS Nodepool can be imported as a resource by ID: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_sks_nodepool.ci-builders eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` [r-sks_cluster]: sks_cluster.html [sks-doc]: https://community.exoscale.com/documentation/sks/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_ssh_keypair",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale SSH Keypair.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"keypair",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The unique identifier of the SSH Keypair.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The SSH private key generated if no public key was provided. ## Import An existing SSH Keypair can be imported as a resource by name: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_ssh_keypair.mykey my-key ` + "`" + `` + "`" + `` + "`" + ` [ssh-keypairs-doc]: https://community.exoscale.com/documentation/compute/ssh-keypairs/`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"exoscale_affinity":             0,
		"exoscale_compute":              1,
		"exoscale_domain":               2,
		"exoscale_domain_record":        3,
		"exoscale_instance_pool":        4,
		"exoscale_ipaddress":            5,
		"exoscale_network":              6,
		"exoscale_nic":                  7,
		"exoscale_nlb":                  8,
		"exoscale_nlb_service":          9,
		"exoscale_secondary_ipaddress":  10,
		"exoscale_security_group":       11,
		"exoscale_security_group_rule":  12,
		"exoscale_security_group_rules": 13,
		"exoscale_sks_cluster":          14,
		"exoscale_sks_nodepool":         15,
		"exoscale_ssh_keypair":          16,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
