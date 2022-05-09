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
			Category:         "Deprecated",
			ShortDescription: `Provides an Exoscale Anti-Affinity Group resource.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
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
					Description: `The IDs of the Compute instance resources member of the Anti-Affinity Group. ## Import An existing Anti-Affinity Group can be imported as a resource by name or ID: ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_affinity.mygroup mygroup # By ID $ terraform import exoscale_affinity.mygroup eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` [aag-doc]: https://community.exoscale.com/documentation/compute/anti-affinity-groups/ [r-anti-affinity-group]: ../resources/anti_affinity_group`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_anti_affinity_group",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Anti-Affinity Group resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Anti-Affinity Group. ## Import An existing Anti-Affinity Group can be imported as a resource its ID: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_anti_affinity_group.my-group eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` [aag-doc]: https://community.exoscale.com/documentation/compute/anti-affinity-groups/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_compute",
			Category:         "Deprecated",
			ShortDescription: `Provides an Exoscale Compute instance resource.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
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
			Type:             "exoscale_compute_instance",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Compute instance resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation date of the Compute instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Compute instance.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The IPv6 address of the Compute instance main network interface.`,
				},
				resource.Attribute{
					Name:        "private_network_ids",
					Description: `(Deprecated) A list of [Private Network][r-private_network] IDs attached to the Compute instance. Attached network interfaces can be set using the ` + "`" + `network_interface` + "`" + ` block argument. Usage of the [` + "`" + `exoscale_private_network` + "`" + `][d-private_network] data source is recommended.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `The IPv4 address of the Compute instance's main network interface. ## ` + "`" + `remote-exec` + "`" + ` provisioner usage If you wish to log to a ` + "`" + `exoscale_compute_instance` + "`" + ` resource using the [` + "`" + `remote-exec` + "`" + `][remote-exec] provisioner, make sure to explicity set the SSH ` + "`" + `user` + "`" + ` setting to connect to the instance to the actual template username returned by the [` + "`" + `exoscale_compute_template` + "`" + `][compute_template] data source: ` + "`" + `` + "`" + `` + "`" + `hcl data "exoscale_compute_template" "ubuntu" { zone = "ch-gva-2" name = "Linux Ubuntu 20.04 LTS 64-bit" } resource "exoscale_compute_instance" "mymachine" { # ... provisioner "remote-exec" { connection { type = "ssh" host = self.ip_address user = data.exoscale_compute_template.ubuntu.username } } } ` + "`" + `` + "`" + `` + "`" + ` ## Import An existing Compute instance can be imported as a resource by ` + "`" + `<ID>@<ZONE>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_compute_instance.my-instance eb556678-ec59-4be6-8c54-0406ae0f6da6@ch-gva-2 ` + "`" + `` + "`" + `` + "`" + ` [cloudinit]: http://cloudinit.readthedocs.io/en/latest/ [compute-doc]: https://community.exoscale.com/documentation/compute/ [d-anti_affinity_group]: ../data-sources/anti_affinity_group [d-compute_template]: ../data-sources/compute_template [d-elastic_ip]: ../data-sources/elastic_ip [d-private_network]: ../data-sources/private_network [d-security_group]: ../data-sources/security_group [r-anti_affinity_group]: ../resources/anti_affinity_group [r-elastic_ip]: ../resources/elastic_ip [r-private_network]: ../resources/private_network [r-security_group]: ../resources/security_group [remote-exec]: https://www.terraform.io/docs/provisioners/remote-exec.html [sshkeypair-doc]: https://community.exoscale.com/documentation/compute/ssh-keypairs/ [template]: https://www.exoscale.com/templates/ [type]: https://www.exoscale.com/pricing/#/compute/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_database",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale database service resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation date of the database service.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The disk size of the database service.`,
				},
				resource.Attribute{
					Name:        "node_cpus",
					Description: `The number of CPUs of the database service.`,
				},
				resource.Attribute{
					Name:        "node_memory",
					Description: `The amount of memory of the database service.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `The number of nodes of the database service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the database service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the database service.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date of the latest database service update.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The database service connection URI. ## Import An existing database service can be imported as a resource by specifying ` + "`" + `NAME@ZONE` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_database.example my-database@de-fra-1 ` + "`" + `` + "`" + `` + "`" + ` [dbaas-doc]: https://community.exoscale.com/documentation/dbaas/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_domain",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale DNS Domain resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
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
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale DNS domain record resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `The DNS domain record's`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_elastic_ip",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Elastic IP address.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `The Elastic IP address. ## Import An existing Elastic IP can be imported as a resource by specifying ` + "`" + `ID@ZONE` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_elastic_ip.web eb556678-ec59-4be6-8c54-0406ae0f6da6@ch-dk-2 ` + "`" + `` + "`" + `` + "`" + ` [eip-doc]: https://community.exoscale.com/documentation/compute/eip/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_instance_pool",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Instance Pool resource.`,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "instance_type",
					Description: `(Required) The managed Compute instances [type][type] (format: ` + "`" + `FAMILY.SIZE` + "`" + `, e.g. ` + "`" + `standard.medium` + "`" + `, ` + "`" + `memory.huge` + "`" + `).`,
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
					Name:        "ipv6",
					Description: `Enable IPv6 on managed Compute instances (default: ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "instance_prefix",
					Description: `The string to add as prefix to managed Compute instances name (default: ` + "`" + `pool` + "`" + `).`,
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
					Description: `A Deploy Target ID.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of key/value labels. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The list of Instance Pool members. The ` + "`" + `instances` + "`" + ` items contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the compute instance.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The IPv6 address of the compute instance's main network interface.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the compute instance.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `The IPv4 address of the compute instance's main network interface. ## Import An existing Instance Pool can be imported as a resource by ` + "`" + `<ID>@<ZONE>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_instance_pool.example eb556678-ec59-4be6-8c54-0406ae0f6da6@de-fra-1 ` + "`" + `` + "`" + `` + "`" + ` [cloudinit]: http://cloudinit.readthedocs.io/en/latest/ [d-compute_template]: ../data-sources/compute_template [eip-doc]: https://community.exoscale.com/documentation/compute/eip/ [privnet-doc]: https://community.exoscale.com/documentation/compute/private-networks/ [r-affinity]: ../resources/affinity [r-security_group]: ../resources/security_group [sshkeypair]: https://community.exoscale.com/documentation/compute/ssh-keypairs/ [template]: https://www.exoscale.com/templates/ [type]: https://www.exoscale.com/pricing/#/compute/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instances",
					Description: `The list of Instance Pool members. The ` + "`" + `instances` + "`" + ` items contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the compute instance.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The IPv6 address of the compute instance's main network interface.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the compute instance.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `The IPv4 address of the compute instance's main network interface. ## Import An existing Instance Pool can be imported as a resource by ` + "`" + `<ID>@<ZONE>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_instance_pool.example eb556678-ec59-4be6-8c54-0406ae0f6da6@de-fra-1 ` + "`" + `` + "`" + `` + "`" + ` [cloudinit]: http://cloudinit.readthedocs.io/en/latest/ [d-compute_template]: ../data-sources/compute_template [eip-doc]: https://community.exoscale.com/documentation/compute/eip/ [privnet-doc]: https://community.exoscale.com/documentation/compute/private-networks/ [r-affinity]: ../resources/affinity [r-security_group]: ../resources/security_group [sshkeypair]: https://community.exoscale.com/documentation/compute/ssh-keypairs/ [template]: https://www.exoscale.com/templates/ [type]: https://www.exoscale.com/pricing/#/compute/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_ipaddress",
			Category:         "Deprecated",
			ShortDescription: `Provides an Exoscale Elastic IP address.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"ipaddress",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `The Elastic IP address. ## Import An existing Elastic IP can be imported as a resource by address or ID: ` + "`" + `` + "`" + `` + "`" + `console # By address $ terraform import exoscale_ipaddress.myip 159.100.251.224 # By ID $ terraform import exoscale_ipaddress.myip eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` [eip-doc]: https://community.exoscale.com/documentation/compute/eip/ [r-secondary_ipaddress]: ../resources/secondary_ipaddress [r-elastic-ip]: ../resources/elastic_ip [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_network",
			Category:         "Deprecated",
			ShortDescription: `Provides an Exoscale Private Network.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"network",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Network. ## Import An existing Private Network can be imported as a resource by name or ID: ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_network.net myprivnet # By ID $ terraform import exoscale_network.net 04fb76a2-6d22-49be-8da7-f2a5a0b902e1 ` + "`" + `` + "`" + `` + "`" + ` [r-nic]: ../resources/nic [r-private-network]: ../resources/private_network [privnet-doc]: https://community.exoscale.com/documentation/compute/private-networks/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_nic",
			Category:         "Deprecated",
			ShortDescription: `Provides an Exoscale Compute instance Private Network Interface (NIC).`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
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
					Description: `The physical address (MAC) of the Compute instance NIC. ## Import This resource is automatically imported when importing an ` + "`" + `exoscale_compute` + "`" + ` resource. [privnet-doc]: https://community.exoscale.com/documentation/compute/private-networks/ [r-compute]: ../resources/compute [r-network]: ../resources/network`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_nlb",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Network Load Balancer resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
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
					Description: `The list of the NLB service names. ## Import An existing NLB can be imported as a resource by ` + "`" + `<ID>@<ZONE>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_nlb.example eb556678-ec59-4be6-8c54-0406ae0f6da6@de-fra-1 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_nlb_service",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Network Load Balancer service resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NLB service. ## Import An existing NLB service can be imported as a resource by ` + "`" + `<NLB-ID>/<SERVICE-ID>@<ZONE>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_nlb_service.example eb556678-ec59-4be6-8c54-0406ae0f6da6/9ecc6b8b-73d4-4211-8ced-f7f29bb79524@de-fra-1 ` + "`" + `` + "`" + `` + "`" + ` [r-nlb]: ../resources/nlb [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_private_network",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Private Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Network. ## Import An existing Private Network can be imported as a resource by specifying ` + "`" + `ID@ZONE` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_private_network.net 04fb76a2-6d22-49be-8da7-f2a5a0b902e1@ch-gva-2 ` + "`" + `` + "`" + `` + "`" + ` [privnet-doc]: https://community.exoscale.com/documentation/compute/private-networks/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_secondary_ipaddress",
			Category:         "Deprecated",
			ShortDescription: `Provides an Exoscale resource for assigning an existing Elastic IP to a Compute instance.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
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
					Description: `The ID of the Network the Compute instance NIC is attached to. ## Import This resource is automatically imported when importing an ` + "`" + `exoscale_compute` + "`" + ` resource. [r-compute]: ../resources/compute [r-compute-instance]: ../resources/compute_instance [r-ipaddress]: ../resources/ipaddress`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_security_group",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Security Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Security Group. ## Import An existing Security Group can be imported as a resource by its ID: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_security_group.http eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_security_group_rule",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Security Group Rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Security Group rule. ## Import An existing Security Group rule can be imported as a resource by ` + "`" + `<SECURITY-GROUP-ID>/<SECURITY-GROUP-RULE-ID>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_security_group_rule.http eb556678-ec59-4be6-8c54-0406ae0f6da6/846831cb-a0fc-454b-9abd-cb526559fcf9 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_security_group_rules",
			Category:         "Deprecated",
			ShortDescription: `Provides a resource for assigning multiple rules to an existing Exoscale Security Group.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
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
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
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
					Description: `The list of [SKS Nodepools][r-sks_nodepool] (IDs) attached to the SKS cluster. ## Import An existing SKS cluster can be imported as a resource by specifying ` + "`" + `ID@ZONE` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_sks_cluster.example eb556678-ec59-4be6-8c54-0406ae0f6da6@de-fra-1 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_sks_kubeconfig",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale SKS cluster Kubeconfig.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `The generated Kubeconfig for use to interact with the SKS cluster. ## Automatic Renewal This resource considers its instances to have been deleted after either their validity period ends or the early renewal period is reached. At this time, applying the Terraform configuration will cause a new Kubeconfig to be generated for the instance. Therefore in a development environment with frequent deployments, it may be convenient to set a relatively-short expiration time and use early renewal to automatically provision a new Kubeconfig when the current one is about to expire. The creation of a new Kubeconfig may of course cause dependent resources to be updated or replaced, depending on the lifecycle rules applying to those resources. [sks-doc]: https://community.exoscale.com/documentation/sks/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_sks_nodepool",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale SKS Nodepool resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
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
					Description: `The Kubernetes version of the SKS Nodepool members. ## Import An existing SKS Nodepool can be imported as a resource by ` + "`" + `<CLUSTER-ID>/<NODEPOOL-ID>@<ZONE>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_sks_nodepool.ci-builders eb556678-ec59-4be6-8c54-0406ae0f6da6/8c08b92a-e673-47c7-866e-dc009f620a82@de-fra-1 ` + "`" + `` + "`" + `` + "`" + ` [k8s-taints]: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ [r-sks_cluster]: ../resources/sks_cluster [sks-doc]: https://community.exoscale.com/documentation/sks/ [type]: https://www.exoscale.com/pricing/#/compute/ [zone]: https://www.exoscale.com/datacenters/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale SSH Key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The unique identifier of the SSH Key. ## Import An existing SSH Key can be imported as a resource by name: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_ssh_key.my-key my-key ` + "`" + `` + "`" + `` + "`" + ` [ssh-keys-doc]: https://community.exoscale.com/documentation/compute/ssh-keys/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_ssh_keypair",
			Category:         "Deprecated",
			ShortDescription: `Provides an Exoscale SSH Keypair.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
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
					Description: `The SSH private key generated if no public key was provided. ## Import An existing SSH Keypair can be imported as a resource by name: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_ssh_keypair.mykey my-key ` + "`" + `` + "`" + `` + "`" + ` [ssh-keypairs-doc]: https://community.exoscale.com/documentation/compute/ssh-keypairs/ [r-ssh-key]: ../resources/ssh_key`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"exoscale_affinity":             0,
		"exoscale_anti_affinity_group":  1,
		"exoscale_compute":              2,
		"exoscale_compute_instance":     3,
		"exoscale_database":             4,
		"exoscale_domain":               5,
		"exoscale_domain_record":        6,
		"exoscale_elastic_ip":           7,
		"exoscale_instance_pool":        8,
		"exoscale_ipaddress":            9,
		"exoscale_network":              10,
		"exoscale_nic":                  11,
		"exoscale_nlb":                  12,
		"exoscale_nlb_service":          13,
		"exoscale_private_network":      14,
		"exoscale_secondary_ipaddress":  15,
		"exoscale_security_group":       16,
		"exoscale_security_group_rule":  17,
		"exoscale_security_group_rules": 18,
		"exoscale_sks_cluster":          19,
		"exoscale_sks_kubeconfig":       20,
		"exoscale_sks_nodepool":         21,
		"exoscale_ssh_key":              22,
		"exoscale_ssh_keypair":          23,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
