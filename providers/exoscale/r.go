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
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Anti-Affinity Group resource.`,
			Description:      ``,
			Keywords: []string{
				"affinity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Anti-Affinity Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free-form text describing the Anti-Affinity Group purpose.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Anti-Affinity Group (` + "`" + `host anti-affinity` + "`" + ` is the only supported value). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Anti-Affinity Group.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_ids",
					Description: `The IDs of the Compute instance resources member of the Anti-Affinity Group. ## Import An existing Anti-Affinity Group can be imported as a resource by name or ID: ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_affinity.mygroup mygroup # By ID $ terraform import exoscale_affinity.mygroup eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Anti-Affinity Group.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_ids",
					Description: `The IDs of the Compute instance resources member of the Anti-Affinity Group. ## Import An existing Anti-Affinity Group can be imported as a resource by name or ID: ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_affinity.mygroup mygroup # By ID $ terraform import exoscale_affinity.mygroup eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ``,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of the [zone][zone] to deploy the Compute instance into.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The name of the Compute instance [template][template]. Only`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Required) The ID of the Compute instance [template][template]. Usage of the [` + "`" + `compute_template` + "`" + `][compute_template] data source is recommended.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The Compute instance [size][size], e.g. ` + "`" + `Tiny` + "`" + `, ` + "`" + `Small` + "`" + `, ` + "`" + `Medium` + "`" + `, ` + "`" + `Large` + "`" + ` etc.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) The Compute instance root disk size in GiB (at least ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The displayed name of the Compute instance. Note: if the ` + "`" + `hostname` + "`" + ` attribute is not set, this attribute is also used to set the OS'`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The Compute instance hostname, must contain only alphanumeric and hyphen ("-") characters. If neither ` + "`" + `display_name` + "`" + ` or ` + "`" + `hostname` + "`" + ` attributes are set, a random value will be generated automatically server-side. Note: updating this attribute's value requires to reboot the instance.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `The name of the [SSH key pair][sshkeypair] to be installed.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `A [cloud-init][cloudinit] configuration. Whenever possible don't base64-encode neither gzip it yourself, as this will be automatically taken care of on your behalf by the provider.`,
				},
				resource.Attribute{
					Name:        "keyboard",
					Description: `The keyboard layout configuration (at creation time only). Supported values are: ` + "`" + `de` + "`" + `, ` + "`" + `de-ch` + "`" + `, ` + "`" + `es` + "`" + `, ` + "`" + `fi` + "`" + `, ` + "`" + `fr` + "`" + `, ` + "`" + `fr-be` + "`" + `, ` + "`" + `fr-ch` + "`" + `, ` + "`" + `is` + "`" + `, ` + "`" + `it` + "`" + `, ` + "`" + `jp` + "`" + `, ` + "`" + `nl-be` + "`" + `, ` + "`" + `no` + "`" + `, ` + "`" + `pt` + "`" + `, ` + "`" + `uk` + "`" + `, ` + "`" + `us` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Compute instance, e.g. ` + "`" + `Running` + "`" + ` or ` + "`" + `Stopped` + "`" + ``,
				},
				resource.Attribute{
					Name:        "affinity_groups",
					Description: `A list of [Anti-Affinity Group][aag] names (at creation time only; conflicts with ` + "`" + `affinity_group_ids` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "affinity_group_ids",
					Description: `A list of [Anti-Affinity Group][aag] IDs (at creation time only; conflicts with ` + "`" + `affinity_groups` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A list of [Security Group][sg] names (conflicts with ` + "`" + `security_group_ids` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `A list of [Security Group][sg] IDs (conflicts with ` + "`" + `security_groups` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ip4",
					Description: `Boolean controlling if IPv4 is enabled (only supported value is ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ip6",
					Description: `Boolean controlling if IPv6 is enabled.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of tags (key/value). To remove all tags, set attribute to ` + "`" + `tags = {}` + "`" + `. [template]: https://www.exoscale.com/templates/ [zone]: https://www.exoscale.com/datacenters/ [size]: https://www.exoscale.com/pricing/#/compute/ [sshkeypair]: https://community.exoscale.com/documentation/compute/ssh-keypairs/ [cloudinit]: http://cloudinit.readthedocs.io/en/latest/ [aag]: affinity.html [sg]: security_group.html [compute_template]: ../d/compute_template.html ## Attributes Reference The following attributes are exported:`,
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
					Description: `The IPv6 address of the Compute instance main network interface. [compute_template]: ../d/compute_template.html ## ` + "`" + `remote-exec` + "`" + ` provisioner usage If you wish to log to a ` + "`" + `exoscale_compute` + "`" + ` resource using the [` + "`" + `remote-exec` + "`" + `][rexec] provisioner, make sure to explicity set the SSH ` + "`" + `user` + "`" + ` setting to connect to the instance to the actual template username returned by the [` + "`" + `exoscale_compute_template` + "`" + `][compute_template] data source: ` + "`" + `` + "`" + `` + "`" + `hcl data "exoscale_compute_template" "ubuntu" { zone = "ch-gva-2" name = "Linux Ubuntu 18.04 LTS 64-bit" } resource "exoscale_compute" "mymachine" { zone = "ch-gva-2" display_name = "mymachine" template_id = data.exoscale_compute_template.ubuntu.id size = "Medium" disk_size = 10 key_pair = "me@mymachine" state = "Running" provisioner "remote-exec" { connection { type = "ssh" host = self.ip_address user = data.exoscale_compute_template.ubuntu.username } } } ` + "`" + `` + "`" + `` + "`" + ` [rexec]: https://www.terraform.io/docs/provisioners/remote-exec.html [compute_template]: ../d/compute_template.html ## Import An existing Compute instance can be imported as a resource by name or ID. Importing a Compute instance imports the ` + "`" + `exoscale_compute` + "`" + ` resource as well as related [` + "`" + `exoscale_secondary_ipaddress` + "`" + `][secip] and [` + "`" + `exoscale_nic` + "`" + `][nic] resources. [secip]: secondary_ipaddress.html [nic]: nic.html ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_compute.vm1 vm1 # By ID $ terraform import exoscale_compute.vm1 eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The IPv6 address of the Compute instance main network interface. [compute_template]: ../d/compute_template.html ## ` + "`" + `remote-exec` + "`" + ` provisioner usage If you wish to log to a ` + "`" + `exoscale_compute` + "`" + ` resource using the [` + "`" + `remote-exec` + "`" + `][rexec] provisioner, make sure to explicity set the SSH ` + "`" + `user` + "`" + ` setting to connect to the instance to the actual template username returned by the [` + "`" + `exoscale_compute_template` + "`" + `][compute_template] data source: ` + "`" + `` + "`" + `` + "`" + `hcl data "exoscale_compute_template" "ubuntu" { zone = "ch-gva-2" name = "Linux Ubuntu 18.04 LTS 64-bit" } resource "exoscale_compute" "mymachine" { zone = "ch-gva-2" display_name = "mymachine" template_id = data.exoscale_compute_template.ubuntu.id size = "Medium" disk_size = 10 key_pair = "me@mymachine" state = "Running" provisioner "remote-exec" { connection { type = "ssh" host = self.ip_address user = data.exoscale_compute_template.ubuntu.username } } } ` + "`" + `` + "`" + `` + "`" + ` [rexec]: https://www.terraform.io/docs/provisioners/remote-exec.html [compute_template]: ../d/compute_template.html ## Import An existing Compute instance can be imported as a resource by name or ID. Importing a Compute instance imports the ` + "`" + `exoscale_compute` + "`" + ` resource as well as related [` + "`" + `exoscale_secondary_ipaddress` + "`" + `][secip] and [` + "`" + `exoscale_nic` + "`" + `][nic] resources. [secip]: secondary_ipaddress.html [nic]: nic.html ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_compute.vm1 vm1 # By ID $ terraform import exoscale_compute.vm1 eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ``,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the DNS Domain. ## Attributes Reference The following attributes are exported:`,
				},
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
			ShortDescription: `Provides an Exoscale DNS Domain Record resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"domain",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The name of the [` + "`" + `exoscale_domain` + "`" + `][domain] to create the record into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the DNS Domain Record.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) The type of the DNS Domain Record. Supported values are: ` + "`" + `A` + "`" + `, ` + "`" + `AAAA` + "`" + `, ` + "`" + `ALIAS` + "`" + `, ` + "`" + `CAA` + "`" + `, ` + "`" + `CNAME` + "`" + `, ` + "`" + `HINFO` + "`" + `, ` + "`" + `MX` + "`" + `, ` + "`" + `NAPTR` + "`" + `, ` + "`" + `NS` + "`" + `, ` + "`" + `POOL` + "`" + `, ` + "`" + `SPF` + "`" + `, ` + "`" + `SRV` + "`" + `, ` + "`" + `SSHFP` + "`" + `, ` + "`" + `TXT` + "`" + `, ` + "`" + `URL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) The value of the DNS Domain Record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The [Time To Live][ttl] of the DNS Domain Record.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `The priority of the DNS Domain Record (for types that support it). [domain]: domain.html [ttl]: https://en.wikipedia.org/wiki/Time_to_live ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The DNS Domain Record's`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `The DNS Domain Record's`,
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
					Description: `(Required) (Required) The ID of the instance [template][template] to use when creating Compute instances. Usage of the [` + "`" + `compute_template` + "`" + `][compute_template] data source is recommended.`,
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
					Name:        "security_group_ids",
					Description: `A list of [Security Group][sg] IDs.`,
				},
				resource.Attribute{
					Name:        "network_ids",
					Description: `A list of [Private Network][net] IDs. [template]: https://www.exoscale.com/templates/ [zone]: https://www.exoscale.com/datacenters/ [size]: https://www.exoscale.com/pricing/#/compute/ [sshkeypair]: https://community.exoscale.com/documentation/compute/ssh-keypairs/ [cloudinit]: http://cloudinit.readthedocs.io/en/latest/ [compute_template]: ../d/compute_template.html [net]: https://community.exoscale.com/documentation/compute/private-networks/ ## Import An existing Instance Pool can be imported as a resource by name or ID. Importing an Instance Pool imports the ` + "`" + `exoscale_instance_pool` + "`" + ` resource. ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_instance_pool.pool mypool # By ID $ terraform import exoscale_instance_pool.pool eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ``,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of the [zone][zone] to create the Elastic IP into.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Elastic IP.`,
				},
				resource.Attribute{
					Name:        "healthcheck_mode",
					Description: `The healthcheck probing mode (must be either ` + "`" + `tcp` + "`" + ` or ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "healthcheck_port",
					Description: `The healthcheck service port to probe (must be between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "healthcheck_path",
					Description: `The healthcheck probe HTTP request path (must be specified in ` + "`" + `http` + "`" + ` mode).`,
				},
				resource.Attribute{
					Name:        "healthcheck_interval",
					Description: `The healthcheck probing interval in seconds (must be between ` + "`" + `5` + "`" + ` and ` + "`" + `300` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "healthcheck_timeout",
					Description: `The time in seconds before considering a healthcheck probing failed (must be between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "healthcheck_strikes_ok",
					Description: `The number of successful healthcheck probes before considering the target healthy (must be between ` + "`" + `1` + "`" + ` and ` + "`" + `20` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "healthcheck_strikes_fail",
					Description: `The number of unsuccessful healthcheck probes before considering the target unhealthy (must be between ` + "`" + `1` + "`" + ` and ` + "`" + `20` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of tags (key/value). To remove all tags, set attribute to ` + "`" + `tags = {}` + "`" + `. [zone]: https://www.exoscale.com/datacenters/ ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The Elastic IP address. ## Import An existing Elastic IP can be imported as a resource by address or ID: ` + "`" + `` + "`" + `` + "`" + `console # By address $ terraform import exoscale_ipaddress.myip 159.100.251.224 # By ID $ terraform import exoscale_ipaddress.myip eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `The Elastic IP address. ## Import An existing Elastic IP can be imported as a resource by address or ID: ` + "`" + `` + "`" + `` + "`" + `console # By address $ terraform import exoscale_ipaddress.myip 159.100.251.224 # By ID $ terraform import exoscale_ipaddress.myip eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_network",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Private Network.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of the [zone][zone] to create the Private Network into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Private Network.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `A free-form text describing the Private Network purpose.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `The first address of IP range used by the DHCP service to automatically assign. Required for`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `The last address of the IP range used by the DHCP service. Required for`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The netmask defining the IP network allowed for the static lease (see ` + "`" + `exoscale_nic` + "`" + ` resource). Required for`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of tags (key/value). To remove all tags, set attribute to ` + "`" + `tags = {}` + "`" + `. [zone]: https://www.exoscale.com/datacenters/ ## Import An existing Private Network can be imported as a resource by name or ID: ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_network.net myprivnet # By ID $ terraform import exoscale_network.net 04fb76a2-6d22-49be-8da7-f2a5a0b902e1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "compute_id",
					Description: `(Required) The [Compute instance][compute] ID.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The [Private Network][privnet] ID.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address to request as static DHCP lease if the NIC is attached to a`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The physical address (MAC) of the Compute instance NIC. ## Import This resource is automatically imported when importing an ` + "`" + `exoscale_compute` + "`" + ` resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mac_address",
					Description: `The physical address (MAC) of the Compute instance NIC. ## Import This resource is automatically imported when importing an ` + "`" + `exoscale_compute` + "`" + ` resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_nlb",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Network Load Balancer resource.`,
			Description:      ``,
			Keywords: []string{
				"nlb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of the [zone][zone] to deploy the NLB into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the NLB.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NLB. [zone]: https://www.exoscale.com/datacenters/ ## Import An existing NLB can be imported as a resource by ID. Importing a NLB imports the ` + "`" + `exoscale_nlb` + "`" + ` resource. ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_nlb.website eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "nlb_id",
					Description: `(Required) The ID of the NLB to attach the service.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of the [zone][zone] used by the NLB.`,
				},
				resource.Attribute{
					Name:        "instance_pool_id",
					Description: `(Required) The ID of the Instance Pool to forward network traffic to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the NLB service.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port of the NLB service.`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `(Required) The port to forward network traffic to on target instances.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol (tcp/udp).`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `The strategy (round-robin/source-hash).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NLB service.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The healthcheck port.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The healthcheck mode (tcp/http).`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The healthcheck URI, must be set only if ` + "`" + `mode` + "`" + ` is ` + "`" + `http` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The healthcheck interval in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The healthcheck timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `The healthcheck retries. [zone]: https://www.exoscale.com/datacenters/ ## Import An existing NLB service can be imported as a resource by ID. Importing a NLB service imports the ` + "`" + `exoscale_nlb_service` + "`" + ` resource. ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_nlb_service.website 9ecc6b8b-73d4-4211-8ced-f7f29bb79524 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "compute_id",
					Description: `(Required) The ID of the [Compute instance][compute].`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The [Elastic IP][eip] address to assign. [compute]: compute.html [eip]: ipaddress.html ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The ID of the Network the Compute instance NIC is attached to. ## Import This resource is automatically imported when importing an ` + "`" + `exoscale_compute` + "`" + ` resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The ID of the Network the Compute instance NIC is attached to. ## Import This resource is automatically imported when importing an ` + "`" + `exoscale_compute` + "`" + ` resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "exoscale_security_group",
			Category:         "Resources",
			ShortDescription: `Provides an Exoscale Security Group.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Security Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free-form text describing the Anti-Affinity Group purpose.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of tags (key/value). To remove all tags, set attribute to ` + "`" + `tags = {}` + "`" + `. ## Import An existing Security Group can be imported as a resource by name or ID: ` + "`" + `` + "`" + `` + "`" + `console # By name $ terraform import exoscale_security_group.http http # By ID $ terraform import exoscale_security_group.http eb556678-ec59-4be6-8c54-0406ae0f6da6 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group",
					Description: `(Required) The Security Group name the rule applies to.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The Security Group ID the rule applies to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The traffic direction to match (` + "`" + `INGRESS` + "`" + ` or ` + "`" + `EGRESS` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The network protocol to match. Supported values are: ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + `, ` + "`" + `ICMPv6` + "`" + `, ` + "`" + `AH` + "`" + `, ` + "`" + `ESP` + "`" + `, ` + "`" + `GRE` + "`" + `, ` + "`" + `IPIP` + "`" + ` and ` + "`" + `ALL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free-form text describing the Security Group Rule purpose.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `A source (for ingress)/destination (for egress) IP subnet to match (conflicts with ` + "`" + `user_security_group` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "user_security_group_id",
					Description: `A source (for ingress)/destination (for egress) Security Group ID to match (conflicts with ` + "`" + `cidr` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "user_security_group",
					Description: `A source (for ingress)/destination (for egress) Security Group name to match (conflicts with ` + "`" + `cidr` + "`" + `). [icmp]: https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `The name of the Security Group the rule applies to.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the Security Group the rule applies to.`,
				},
				resource.Attribute{
					Name:        "user_security_group",
					Description: `The name of the source (for ingress)/destination (for egress) Security Group to match. ## Import This resource is automatically imported when importing an ` + "`" + `exoscale_security_group` + "`" + ` resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group",
					Description: `The name of the Security Group the rule applies to.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the Security Group the rule applies to.`,
				},
				resource.Attribute{
					Name:        "user_security_group",
					Description: `The name of the source (for ingress)/destination (for egress) Security Group to match. ## Import This resource is automatically imported when importing an ` + "`" + `exoscale_security_group` + "`" + ` resource.`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group",
					Description: `(Required) The Security Group name the rules apply to.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The Security Group ID the rules apply to. ` + "`" + `egress` + "`" + ` and ` + "`" + `ingress` + "`" + ` support the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The network protocol to match. Supported values are: ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + `, ` + "`" + `ICMPv6` + "`" + `, ` + "`" + `AH` + "`" + `, ` + "`" + `ESP` + "`" + `, ` + "`" + `GRE` + "`" + `, ` + "`" + `IPIP` + "`" + ` and ` + "`" + `ALL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free-form text describing the Security Group Rule purpose.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `A list of ports or port ranges (` + "`" + `start_port-end_port` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "cidr_list",
					Description: `A list of source (for ingress)/destination (for egress) IP subnet to match (conflicts with ` + "`" + `user_security_group` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "user_security_group_list",
					Description: `A source (for ingress)/destination (for egress) of the traffic identified by a security group [icmp]: https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `The name of the Security Group the rules apply to.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the Security Group the rules apply to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group",
					Description: `The name of the Security Group the rules apply to.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the Security Group the rules apply to.`,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SSH Keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `A SSH public key that will be copied into the instances at`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The unique identifier of the SSH Keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The SSH public key generated if none was provided.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The SSH private key generated if no public key was provided. ## Import An existing SSH Keypair can be imported as a resource by name: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_ssh_keypair.mykey my-key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The unique identifier of the SSH Keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The SSH public key generated if none was provided.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The SSH private key generated if no public key was provided. ## Import An existing SSH Keypair can be imported as a resource by name: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import exoscale_ssh_keypair.mykey my-key ` + "`" + `` + "`" + `` + "`" + ``,
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
		"exoscale_ssh_keypair":          14,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
