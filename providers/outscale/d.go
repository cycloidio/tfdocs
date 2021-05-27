package outscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "access_key",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific access key.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "access_key_ids",
					Description: `(Optional) The IDs of the access keys.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the access keys (` + "`" + `ACTIVE` + "`" + ` \| ` + "`" + `INACTIVE` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the access key.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time of the last modification of the access key.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the access key.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time of the last modification of the access key.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "access_keys",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about access keys.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "access_key_ids",
					Description: `(Optional) The IDs of the access keys.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the access keys (` + "`" + `ACTIVE` + "`" + ` \| ` + "`" + `INACTIVE` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_keys",
					Description: `A list of access keys.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the access key.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time of the last modification of the access key.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_keys",
					Description: `A list of access keys.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the access key.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time of the last modification of the access key.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "client_gateway",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific client gateway.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "bgp_asns",
					Description: `(Optional) The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_ids",
					Description: `(Optional) The IDs of the client gateways.`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of communication tunnels used by the client gateways (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `(Optional) The public IPv4 addresses of the client gateways.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the client gateways (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the client gateways.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the client gateways.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the client gateways, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `An Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of communication tunnel used by the client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IPv4 address of the client gateway (must be a fixed address into a NATed network).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the client gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the client gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `An Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of communication tunnel used by the client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IPv4 address of the client gateway (must be a fixed address into a NATed network).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the client gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the client gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "client_gateways",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about client gateways.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "bgp_asns",
					Description: `(Optional) The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_ids",
					Description: `(Optional) The IDs of the client gateways.`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of communication tunnels used by the client gateways (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `(Optional) The public IPv4 addresses of the client gateways.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the client gateways (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the client gateways.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the client gateways.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the client gateways, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "client_gateways",
					Description: `Information about one or more client gateways.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `An Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of communication tunnel used by the client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IPv4 address of the client gateway (must be a fixed address into a NATed network).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the client gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the client gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_gateways",
					Description: `Information about one or more client gateways.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `An Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of communication tunnel used by the client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IPv4 address of the client gateway (must be a fixed address into a NATed network).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the client gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the client gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dhcp_option",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific DHCP option.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, lists all default DHCP options set. If ` + "`" + `false` + "`" + `, lists all non-default DHCP options set.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_ids",
					Description: `(Optional) The IDs of the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `(Optional) The domain name servers used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Optional) The domain names used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `(Optional) The Network Time Protocol (NTP) servers used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the DHCP options sets, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If ` + "`" + `true` + "`" + `, the DHCP options set is a default one. If ` + "`" + `false` + "`" + `, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_name",
					Description: `The name of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IP addresses for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IP addresses for the NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "default",
					Description: `If ` + "`" + `true` + "`" + `, the DHCP options set is a default one. If ` + "`" + `false` + "`" + `, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_name",
					Description: `The name of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IP addresses for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IP addresses for the NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dhcp_options",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about DHCP options.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, lists all default DHCP options set. If ` + "`" + `false` + "`" + `, lists all non-default DHCP options set.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_ids",
					Description: `(Optional) The IDs of the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `(Optional) The domain name servers used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Optional) The domain names used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `(Optional) The Network Time Protocol (NTP) servers used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the DHCP options sets, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dhcp_options_sets",
					Description: `Information about one or more DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If ` + "`" + `true` + "`" + `, the DHCP options set is a default one. If ` + "`" + `false` + "`" + `, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_name",
					Description: `The name of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IP addresses for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IP addresses for the NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_sets",
					Description: `Information about one or more DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If ` + "`" + `true` + "`" + `, the DHCP options set is a default one. If ` + "`" + `false` + "`" + `, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_name",
					Description: `The name of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IP addresses for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IP addresses for the NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexible_gpu",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific flexible GPU.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) Indicates whether the fGPU is deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_ids",
					Description: `(Optional) One or more IDs of fGPUs.`,
				},
				resource.Attribute{
					Name:        "generations",
					Description: `(Optional) The processor generations that the fGPUs are compatible with.`,
				},
				resource.Attribute{
					Name:        "model_names",
					Description: `(Optional) One or more models of fGPUs. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the fGPUs (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The Subregions where the fGPUs are located.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "flexible_gpus",
					Description: `Information about one or more fGPUs.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `The ID of the fGPU.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The compatible processor generation.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the fGPU (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion where the fGPU is located.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the fGPU is attached to, if any.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "flexible_gpus",
					Description: `Information about one or more fGPUs.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `The ID of the fGPU.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The compatible processor generation.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the fGPU (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion where the fGPU is located.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the fGPU is attached to, if any.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexible_gpu_catalog",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific flexible GPU catalog.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "flexible_gpu_catalog",
					Description: `Information about one or more fGPUs available in the public catalog.`,
				},
				resource.Attribute{
					Name:        "generations",
					Description: `The generations of VMs that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "max_cpu",
					Description: `The maximum number of VM vCores that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "max_ram",
					Description: `The maximum amount of VM memory that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU.`,
				},
				resource.Attribute{
					Name:        "v_ram",
					Description: `The amount of video RAM (VRAM) of the fGPU.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "flexible_gpu_catalog",
					Description: `Information about one or more fGPUs available in the public catalog.`,
				},
				resource.Attribute{
					Name:        "generations",
					Description: `The generations of VMs that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "max_cpu",
					Description: `The maximum number of VM vCores that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "max_ram",
					Description: `The maximum amount of VM memory that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU.`,
				},
				resource.Attribute{
					Name:        "v_ram",
					Description: `The amount of video RAM (VRAM) of the fGPU.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexible_gpus",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about flexible GPUs.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) Indicates whether the fGPU is deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_ids",
					Description: `(Optional) One or more IDs of fGPUs.`,
				},
				resource.Attribute{
					Name:        "generations",
					Description: `(Optional) The processor generations that the fGPUs are compatible with.`,
				},
				resource.Attribute{
					Name:        "model_names",
					Description: `(Optional) One or more models of fGPUs. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the fGPUs (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The Subregions where the fGPUs are located.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "flexible_gpus",
					Description: `Information about one or more fGPUs.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `The ID of the fGPU.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The compatible processor generation.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the fGPU (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion where the fGPU is located.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the fGPU is attached to, if any.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "flexible_gpus",
					Description: `Information about one or more fGPUs.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `The ID of the fGPU.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The compatible processor generation.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the fGPU (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion where the fGPU is located.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the fGPU is attached to, if any.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "image",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific image.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "account_aliases",
					Description: `(Optional) The account aliases of the owners of the OMIs.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the OMIs. By default, all the OMIs for which you have launch permissions are described.`,
				},
				resource.Attribute{
					Name:        "architectures",
					Description: `(Optional) The architectures of the OMIs (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_delete_on_vm_deletion",
					Description: `(Optional) Indicates whether the block device mapping is deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_device_names",
					Description: `(Optional) The device names for the volumes.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_snapshot_ids",
					Description: `(Optional) The IDs of the snapshots used to create the volumes.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_volume_sizes",
					Description: `(Optional) The sizes of the volumes, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_volume_types",
					Description: `(Optional) The types of volumes (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the OMIs, provided when they were created.`,
				},
				resource.Attribute{
					Name:        "file_locations",
					Description: `(Optional) The locations where the OMI files are stored on Object Storage Unit (OSU).`,
				},
				resource.Attribute{
					Name:        "image_ids",
					Description: `(Optional) The IDs of the OMIs.`,
				},
				resource.Attribute{
					Name:        "image_names",
					Description: `(Optional) The names of the OMIs, provided when they were created.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch_account_ids",
					Description: `(Optional) The account IDs of the users who have launch permissions for the OMIs.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch_global_permission",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, lists all public OMIs. If ` + "`" + `false` + "`" + `, lists all private OMIs.`,
				},
				resource.Attribute{
					Name:        "root_device_names",
					Description: `(Optional) The device names of the root devices (for example, ` + "`" + `/dev/sda1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "root_device_types",
					Description: `(Optional) The types of root device used by the OMIs (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the OMIs (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the OMIs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the OMIs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the OMIs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "virtualization_types",
					Description: `(Optional) The virtualization types (always ` + "`" + `hvm` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `Information about one or more OMIs.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the OMI (by default, ` + "`" + `i386` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `Set to ` + "`" + `true` + "`" + ` by default, which means that the volume is deleted when the VM is terminated. If set to ` + "`" + `false` + "`" + `, the volume is not deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (ephemeralN).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time at which the OMI was created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location where the OMI file is stored on Object Storage Unit (OSU).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of the OMI.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If ` + "`" + `true` + "`" + `, the resource is public. If ` + "`" + `false` + "`" + `, the resource is private.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the OMI (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_comment",
					Description: `Information about the change of state.`,
				},
				resource.Attribute{
					Name:        "state_code",
					Description: `The code of the change of state.`,
				},
				resource.Attribute{
					Name:        "state_message",
					Description: `A message explaining the change of state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the OMI.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `Information about one or more OMIs.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the OMI (by default, ` + "`" + `i386` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `Set to ` + "`" + `true` + "`" + ` by default, which means that the volume is deleted when the VM is terminated. If set to ` + "`" + `false` + "`" + `, the volume is not deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (ephemeralN).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time at which the OMI was created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location where the OMI file is stored on Object Storage Unit (OSU).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of the OMI.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If ` + "`" + `true` + "`" + `, the resource is public. If ` + "`" + `false` + "`" + `, the resource is private.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the OMI (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_comment",
					Description: `Information about the change of state.`,
				},
				resource.Attribute{
					Name:        "state_code",
					Description: `The code of the change of state.`,
				},
				resource.Attribute{
					Name:        "state_message",
					Description: `A message explaining the change of state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the OMI.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "images",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about images.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "account_aliases",
					Description: `(Optional) The account aliases of the owners of the OMIs.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the OMIs. By default, all the OMIs for which you have launch permissions are described.`,
				},
				resource.Attribute{
					Name:        "architectures",
					Description: `(Optional) The architectures of the OMIs (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_delete_on_vm_deletion",
					Description: `(Optional) Indicates whether the block device mapping is deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_device_names",
					Description: `(Optional) The device names for the volumes.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_snapshot_ids",
					Description: `(Optional) The IDs of the snapshots used to create the volumes.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_volume_sizes",
					Description: `(Optional) The sizes of the volumes, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_volume_types",
					Description: `(Optional) The types of volumes (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the OMIs, provided when they were created.`,
				},
				resource.Attribute{
					Name:        "file_locations",
					Description: `(Optional) The locations where the OMI files are stored on Object Storage Unit (OSU).`,
				},
				resource.Attribute{
					Name:        "image_ids",
					Description: `(Optional) The IDs of the OMIs.`,
				},
				resource.Attribute{
					Name:        "image_names",
					Description: `(Optional) The names of the OMIs, provided when they were created.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch_account_ids",
					Description: `(Optional) The account IDs of the users who have launch permissions for the OMIs.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch_global_permission",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, lists all public OMIs. If ` + "`" + `false` + "`" + `, lists all private OMIs.`,
				},
				resource.Attribute{
					Name:        "root_device_names",
					Description: `(Optional) The device names of the root devices (for example, ` + "`" + `/dev/sda1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "root_device_types",
					Description: `(Optional) The types of root device used by the OMIs (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the OMIs (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the OMIs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the OMIs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the OMIs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "virtualization_types",
					Description: `(Optional) The virtualization types (always ` + "`" + `hvm` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `Information about one or more OMIs.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the OMI (by default, ` + "`" + `i386` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `Set to ` + "`" + `true` + "`" + ` by default, which means that the volume is deleted when the VM is terminated. If set to ` + "`" + `false` + "`" + `, the volume is not deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (ephemeralN).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time at which the OMI was created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location where the OMI file is stored on Object Storage Unit (OSU).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of the OMI.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If ` + "`" + `true` + "`" + `, the resource is public. If ` + "`" + `false` + "`" + `, the resource is private.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the OMI (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_comment",
					Description: `Information about the change of state.`,
				},
				resource.Attribute{
					Name:        "state_code",
					Description: `The code of the change of state.`,
				},
				resource.Attribute{
					Name:        "state_message",
					Description: `A message explaining the change of state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the OMI.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `Information about one or more OMIs.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the OMI (by default, ` + "`" + `i386` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `Set to ` + "`" + `true` + "`" + ` by default, which means that the volume is deleted when the VM is terminated. If set to ` + "`" + `false` + "`" + `, the volume is not deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (ephemeralN).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time at which the OMI was created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location where the OMI file is stored on Object Storage Unit (OSU).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of the OMI.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If ` + "`" + `true` + "`" + `, the resource is public. If ` + "`" + `false` + "`" + `, the resource is private.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the OMI (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_comment",
					Description: `Information about the change of state.`,
				},
				resource.Attribute{
					Name:        "state_code",
					Description: `The code of the change of state.`,
				},
				resource.Attribute{
					Name:        "state_message",
					Description: `A message explaining the change of state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the OMI.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "internet_service",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific Internet service.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "internet_service_ids",
					Description: `(Optional) The IDs of the Internet services.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Internet services.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Internet services.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the Internet services, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "internet_services",
					Description: `Information about one or more Internet services.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net attached to the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Net to the Internet service (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "internet_services",
					Description: `Information about one or more Internet services.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net attached to the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Net to the Internet service (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "internet_services",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Internet services.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "internet_service_ids",
					Description: `(Optional) The IDs of the Internet services.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Internet services.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Internet services.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the Internet services, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "internet_services",
					Description: `Information about one or more Internet services.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net attached to the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Net to the Internet service (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "internet_services",
					Description: `Information about one or more Internet services.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net attached to the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Net to the Internet service (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "keypair",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific keypair.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprints",
					Description: `(Optional) The fingerprints of the keypairs.`,
				},
				resource.Attribute{
					Name:        "keypair_names",
					Description: `(Optional) The names of the keypairs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "keypairs",
					Description: `Information about one or more keypairs.`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `If you create a keypair, the SHA-1 digest of the DER encoded private key.<br /> If you import a keypair, the MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keypairs",
					Description: `Information about one or more keypairs.`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `If you create a keypair, the SHA-1 digest of the DER encoded private key.<br /> If you import a keypair, the MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "keypairs",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about keypairs.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprints",
					Description: `(Optional) The fingerprints of the keypairs.`,
				},
				resource.Attribute{
					Name:        "keypair_names",
					Description: `(Optional) The names of the keypairs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "keypairs",
					Description: `Information about one or more keypairs.`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `If you create a keypair, the SHA-1 digest of the DER encoded private key.<br /> If you import a keypair, the MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keypairs",
					Description: `Information about one or more keypairs.`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `If you create a keypair, the SHA-1 digest of the DER encoded private key.<br /> If you import a keypair, the MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "load_balancer",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific load balancer.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "load_balancer_names",
					Description: `(Optional) The names of the load balancers. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "load_balancers",
					Description: `Information about one or more load balancers.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If ` + "`" + `true` + "`" + `, access logs are enabled for your load balancer. If ` + "`" + `false` + "`" + `, they are not. If you set this to ` + "`" + `true` + "`" + ` in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the Object Storage Unit (OSU) bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your Object Storage Unit (OSU) bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the Object Storage Unit (OSU) bucket, in minutes. This value can be either 5 or 60 (by default, 60).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path for HTTP or HTTPS requests.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between 1 and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP address.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP address.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The IDs of the Subnets for the load balancer.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `One or more names of Subregions for the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancers",
					Description: `Information about one or more load balancers.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If ` + "`" + `true` + "`" + `, access logs are enabled for your load balancer. If ` + "`" + `false` + "`" + `, they are not. If you set this to ` + "`" + `true` + "`" + ` in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the Object Storage Unit (OSU) bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your Object Storage Unit (OSU) bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the Object Storage Unit (OSU) bucket, in minutes. This value can be either 5 or 60 (by default, 60).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path for HTTP or HTTPS requests.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between 1 and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP address.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP address.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The IDs of the Subnets for the load balancer.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `One or more names of Subregions for the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "load_balancer_listener_rule",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific load balancer listener rule.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "listener_rule_names",
					Description: `(Optional) The names of the listener rules. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_rules",
					Description: `The list of the rules to describe.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_rule_id",
					Description: `The ID of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~"'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `The IDs of the backend VMs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_rules",
					Description: `The list of the rules to describe.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_rule_id",
					Description: `The ID of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~"'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `The IDs of the backend VMs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "load_balancer_listener_rules",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about load balancer listener rules.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "listener_rule_names",
					Description: `(Optional) The names of the listener rules. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_rules",
					Description: `The list of the rules to describe.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_rule_id",
					Description: `The ID of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~"'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `The IDs of the backend VMs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_rules",
					Description: `The list of the rules to describe.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_rule_id",
					Description: `The ID of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~"'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `The IDs of the backend VMs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "load_balancer_vm_health",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about VM health of a specific load balancer.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `(Optional) One or more IDs of back-end VMs.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Required) The name of the load balancer. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "backend_vm_health",
					Description: `Information about the health of one or more back-end VMs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the state of the back-end VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the back-end VM (` + "`" + `InService` + "`" + ` \| ` + "`" + `OutOfService` + "`" + ` \| ` + "`" + `Unknown` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `Information about the cause of ` + "`" + `OutOfService` + "`" + ` VMs.<br /> Specifically, whether the cause is Elastic Load Balancing or the VM (` + "`" + `ELB` + "`" + ` \| ` + "`" + `Instance` + "`" + ` \| ` + "`" + `N/A` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the back-end VM.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "backend_vm_health",
					Description: `Information about the health of one or more back-end VMs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the state of the back-end VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the back-end VM (` + "`" + `InService` + "`" + ` \| ` + "`" + `OutOfService` + "`" + ` \| ` + "`" + `Unknown` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `Information about the cause of ` + "`" + `OutOfService` + "`" + ` VMs.<br /> Specifically, whether the cause is Elastic Load Balancing or the VM (` + "`" + `ELB` + "`" + ` \| ` + "`" + `Instance` + "`" + ` \| ` + "`" + `N/A` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the back-end VM.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "load_balancers",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about load balancers.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "load_balancer_names",
					Description: `(Optional) The names of the load balancers. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "load_balancers",
					Description: `Information about one or more load balancers.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If ` + "`" + `true` + "`" + `, access logs are enabled for your load balancer. If ` + "`" + `false` + "`" + `, they are not. If you set this to ` + "`" + `true` + "`" + ` in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the Object Storage Unit (OSU) bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your Object Storage Unit (OSU) bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the Object Storage Unit (OSU) bucket, in minutes. This value can be either 5 or 60 (by default, 60).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path for HTTP or HTTPS requests.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between 1 and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP address.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP address.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The IDs of the Subnets for the load balancer.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `One or more names of Subregions for the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancers",
					Description: `Information about one or more load balancers.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If ` + "`" + `true` + "`" + `, access logs are enabled for your load balancer. If ` + "`" + `false` + "`" + `, they are not. If you set this to ` + "`" + `true` + "`" + ` in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the Object Storage Unit (OSU) bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your Object Storage Unit (OSU) bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the Object Storage Unit (OSU) bucket, in minutes. This value can be either 5 or 60 (by default, 60).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path for HTTP or HTTPS requests.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between 1 and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP address.<br /> If ` + "`" + `LoadBalancerType` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP address.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The IDs of the Subnets for the load balancer.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `One or more names of Subregions for the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nat_service",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific NAT service.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "nat_service_ids",
					Description: `(Optional) The IDs of the NAT services.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets in which the NAT services are.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the NAT services (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets in which the NAT services are.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the NAT services.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the NAT services.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the NAT services, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nat_services",
					Description: `Information about one or more NAT services.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of the NAT service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about the External IP address or addresses (EIPs) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NAT service (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nat_services",
					Description: `Information about one or more NAT services.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of the NAT service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about the External IP address or addresses (EIPs) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NAT service (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nat_services",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about NAT services.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "nat_service_ids",
					Description: `(Optional) The IDs of the NAT services.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets in which the NAT services are.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the NAT services (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets in which the NAT services are.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the NAT services.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the NAT services.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the NAT services, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nat_services",
					Description: `Information about one or more NAT services.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of the NAT service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about the External IP address or addresses (EIPs) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NAT service (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nat_services",
					Description: `Information about one or more NAT services.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of the NAT service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about the External IP address or addresses (EIPs) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NAT service (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific Net.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_ids",
					Description: `(Optional) The IDs of the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) The IP ranges for the Nets, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Nets (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Nets.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Nets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the Nets, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nets",
					Description: `Information about the described Nets.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nets",
					Description: `Information about the described Nets.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net_access_point",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific Net access point.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "net_access_point_ids",
					Description: `(Optional) The IDs of the Net access points.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets.`,
				},
				resource.Attribute{
					Name:        "service_names",
					Description: `(Optional) The names of the services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Net access points (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Net access points.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Net access points.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the Net access points, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_access_points",
					Description: `One or more Net access points.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `The ID of the route tables associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net access point (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_access_points",
					Description: `One or more Net access points.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `The ID of the route tables associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net access point (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net_access_point_services",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Net access point services.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `(Optional) The IDs of the services.`,
				},
				resource.Attribute{
					Name:        "service_names",
					Description: `(Optional) The names of the services. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The names of the services you can use for Net access points.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `The list of network prefixes used by the service, in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "services",
					Description: `The names of the services you can use for Net access points.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `The list of network prefixes used by the service, in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net_access_points",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Net access points.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "net_access_point_ids",
					Description: `(Optional) The IDs of the Net access points.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets.`,
				},
				resource.Attribute{
					Name:        "service_names",
					Description: `(Optional) The names of the services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Net access points (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Net access points.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Net access points.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the Net access points, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_access_points",
					Description: `One or more Net access points.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `The ID of the route tables associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net access point (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_access_points",
					Description: `One or more Net access points.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `The ID of the route tables associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net access point (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net_attributes",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Net attributes.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "net_id",
					Description: `(Optional) The ID of the Net. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` | ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net (` + "`" + `default` + "`" + ` | ` + "`" + `dedicated` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` | ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net (` + "`" + `default` + "`" + ` | ` + "`" + `dedicated` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net_peering",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific Net peering.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "accepter_net_account_ids",
					Description: `(Optional) The account IDs of the owners of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "accepter_net_ip_ranges",
					Description: `(Optional) The IP ranges of the peer Nets, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "accepter_net_net_ids",
					Description: `(Optional) The IDs of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "net_peering_ids",
					Description: `(Optional) The IDs of the Net peering connections.`,
				},
				resource.Attribute{
					Name:        "source_net_account_ids",
					Description: `(Optional) The account IDs of the owners of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "source_net_ip_ranges",
					Description: `(Optional) The IP ranges of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "source_net_net_ids",
					Description: `(Optional) The IDs of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "state_messages",
					Description: `(Optional) Additional information about the states of the Net peering connections.`,
				},
				resource.Attribute{
					Name:        "state_names",
					Description: `(Optional) The states of the Net peering connections (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Net peering connections.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Net peering connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the Net peering connections, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_peerings",
					Description: `Information about one or more Net peering connections.`,
				},
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering connection (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_peerings",
					Description: `Information about one or more Net peering connections.`,
				},
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering connection (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "net_peerings",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Net peerings.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "accepter_net_account_ids",
					Description: `(Optional) The account IDs of the owners of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "accepter_net_ip_ranges",
					Description: `(Optional) The IP ranges of the peer Nets, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "accepter_net_net_ids",
					Description: `(Optional) The IDs of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "net_peering_ids",
					Description: `(Optional) The IDs of the Net peering connections.`,
				},
				resource.Attribute{
					Name:        "source_net_account_ids",
					Description: `(Optional) The account IDs of the owners of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "source_net_ip_ranges",
					Description: `(Optional) The IP ranges of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "source_net_net_ids",
					Description: `(Optional) The IDs of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "state_messages",
					Description: `(Optional) Additional information about the states of the Net peering connections.`,
				},
				resource.Attribute{
					Name:        "state_names",
					Description: `(Optional) The states of the Net peering connections (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Net peering connections.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Net peering connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the Net peering connections, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_peerings",
					Description: `Information about one or more Net peering connections.`,
				},
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering connection (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_peerings",
					Description: `Information about one or more Net peering connections.`,
				},
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering connection (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nets",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Nets.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_ids",
					Description: `(Optional) The IDs of the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) The IP ranges for the Nets, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Nets (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Nets.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Nets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the Nets, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nets",
					Description: `Information about the described Nets.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nets",
					Description: `Information about the described Nets.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nic",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific NIC.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "link_nic_sort_numbers",
					Description: `(Optional) The device numbers the NICs are attached to.`,
				},
				resource.Attribute{
					Name:        "link_nic_vm_ids",
					Description: `(Optional) The IDs of the VMs the NICs are attached to.`,
				},
				resource.Attribute{
					Name:        "nic_ids",
					Description: `(Optional) The IDs of the NICs.`,
				},
				resource.Attribute{
					Name:        "private_ips_private_ips",
					Description: `(Optional) The private IP addresses of the NICs.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets for the NICs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `Information about one or more NICs.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If ` + "`" + `true` + "`" + `, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the NIC is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nics",
					Description: `Information about one or more NICs.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If ` + "`" + `true` + "`" + `, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the NIC is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nics",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about NICs.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "link_nic_sort_numbers",
					Description: `(Optional) The device numbers the NICs are attached to.`,
				},
				resource.Attribute{
					Name:        "link_nic_vm_ids",
					Description: `(Optional) The IDs of the VMs the NICs are attached to.`,
				},
				resource.Attribute{
					Name:        "nic_ids",
					Description: `(Optional) The IDs of the NICs.`,
				},
				resource.Attribute{
					Name:        "private_ips_private_ips",
					Description: `(Optional) The private IP addresses of the NICs.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets for the NICs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `Information about one or more NICs.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If ` + "`" + `true` + "`" + `, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the NIC is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nics",
					Description: `Information about one or more NICs.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If ` + "`" + `true` + "`" + `, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the NIC is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "public_ip",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific public IP.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_ids",
					Description: `(Optional) The IDs representing the associations of EIPs with VMs or NICs.`,
				},
				resource.Attribute{
					Name:        "nic_account_ids",
					Description: `(Optional) The account IDs of the owners of the NICs.`,
				},
				resource.Attribute{
					Name:        "nic_ids",
					Description: `(Optional) The IDs of the NICs.`,
				},
				resource.Attribute{
					Name:        "placements",
					Description: `(Optional) Whether the EIPs are for use in the public Cloud or in a Net.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) The private IP addresses associated with the EIPs.`,
				},
				resource.Attribute{
					Name:        "public_ip_ids",
					Description: `(Optional) The IDs of the External IP addresses (EIPs).`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `(Optional) The External IP addresses (EIPs).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the EIPs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the EIPs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the EIPs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) The IDs of the VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about one or more EIPs.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the EIP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the External IP (EIP) is associated with (if any).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about one or more EIPs.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the EIP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the External IP (EIP) is associated with (if any).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "public_ips",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about public IPs.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_ids",
					Description: `(Optional) The IDs representing the associations of EIPs with VMs or NICs.`,
				},
				resource.Attribute{
					Name:        "nic_account_ids",
					Description: `(Optional) The account IDs of the owners of the NICs.`,
				},
				resource.Attribute{
					Name:        "nic_ids",
					Description: `(Optional) The IDs of the NICs.`,
				},
				resource.Attribute{
					Name:        "placements",
					Description: `(Optional) Whether the EIPs are for use in the public Cloud or in a Net.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) The private IP addresses associated with the EIPs.`,
				},
				resource.Attribute{
					Name:        "public_ip_ids",
					Description: `(Optional) The IDs of the External IP addresses (EIPs).`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `(Optional) The External IP addresses (EIPs).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the EIPs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the EIPs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the EIPs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) The IDs of the VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about one or more EIPs.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the EIP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the External IP (EIP) is associated with (if any).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about one or more EIPs.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the EIP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the EIP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the EIP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the EIP.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the External IP (EIP) is associated with (if any).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "regions",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Regions.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `Information about one or more Regions.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The hostname of the gateway to access the Region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The administrative name of the Region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `Information about one or more Regions.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The hostname of the gateway to access the Region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The administrative name of the Region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "route_table",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific route table.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "link_route_table_link_route_table_ids",
					Description: `(Optional) The IDs of the associations between the route tables and the Subnets.`,
				},
				resource.Attribute{
					Name:        "link_subnet_ids",
					Description: `(Optional) The IDs of the Subnets involved in the associations.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets for the route tables.`,
				},
				resource.Attribute{
					Name:        "route_creation_methods",
					Description: `(Optional) The methods used to create a route.`,
				},
				resource.Attribute{
					Name:        "route_destination_ip_ranges",
					Description: `(Optional) The IP ranges specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_gateway_ids",
					Description: `(Optional) The IDs of the gateways specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_nat_service_ids",
					Description: `(Optional) The IDs of the NAT services specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_net_peering_ids",
					Description: `(Optional) The IDs of the Net peering connections specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_states",
					Description: `(Optional) The states of routes in the route tables (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `(Optional) The IDs of the route tables.`,
				},
				resource.Attribute{
					Name:        "route_vm_ids",
					Description: `(Optional) The IDs of the VMs specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the route tables.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the route tables.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the route tables, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "route_tables",
					Description: `Information about one or more route tables.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If ` + "`" + `true` + "`" + `, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the 3DS OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "route_tables",
					Description: `Information about one or more route tables.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If ` + "`" + `true` + "`" + `, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the 3DS OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "route_tables",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about route tables.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "link_route_table_link_route_table_ids",
					Description: `(Optional) The IDs of the associations between the route tables and the Subnets.`,
				},
				resource.Attribute{
					Name:        "link_subnet_ids",
					Description: `(Optional) The IDs of the Subnets involved in the associations.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets for the route tables.`,
				},
				resource.Attribute{
					Name:        "route_creation_methods",
					Description: `(Optional) The methods used to create a route.`,
				},
				resource.Attribute{
					Name:        "route_destination_ip_ranges",
					Description: `(Optional) The IP ranges specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_gateway_ids",
					Description: `(Optional) The IDs of the gateways specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_nat_service_ids",
					Description: `(Optional) The IDs of the NAT services specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_net_peering_ids",
					Description: `(Optional) The IDs of the Net peering connections specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_states",
					Description: `(Optional) The states of routes in the route tables (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `(Optional) The IDs of the route tables.`,
				},
				resource.Attribute{
					Name:        "route_vm_ids",
					Description: `(Optional) The IDs of the VMs specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the route tables.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the route tables.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the route tables, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "route_tables",
					Description: `Information about one or more route tables.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If ` + "`" + `true` + "`" + `, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the 3DS OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "route_tables",
					Description: `Information about one or more route tables.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If ` + "`" + `true` + "`" + `, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the 3DS OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering connection.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (` + "`" + `active` + "`" + ` \| ` + "`" + `blackhole` + "`" + `). The ` + "`" + `blackhole` + "`" + ` state indicates that the target of the route is not available.`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "security_group",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific security group.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the security groups.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) The IDs of the security groups.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `(Optional) The names of the security groups.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the security groups.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the security groups.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the security groups, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Information about one or more security groups.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_groups",
					Description: `Information about one or more security groups.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "security_groups",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about security groups.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the security groups.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) The IDs of the security groups.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `(Optional) The names of the security groups.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the security groups.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the security groups.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the security groups, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Information about one or more security groups.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_groups",
					Description: `Information about one or more security groups.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `) or protocol number. By default, ` + "`" + `-1` + "`" + `, which means all protocols.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "snapshot",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific snapshot.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "account_aliases",
					Description: `(Optional) The account aliases of the owners of the snapshots.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the snapshots.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the snapshots.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume_account_ids",
					Description: `(Optional) The account IDs of one or more users who have permissions to create volumes.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume_global_permission",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, lists all public volumes. If ` + "`" + `false` + "`" + `, lists all private volumes.`,
				},
				resource.Attribute{
					Name:        "progresses",
					Description: `(Optional) The progresses of the snapshots, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `(Optional) The IDs of the snapshots.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the snapshots (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the snapshots.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the snapshots.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the snapshots, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `(Optional) The IDs of the volumes used to create the snapshots.`,
				},
				resource.Attribute{
					Name:        "volume_sizes",
					Description: `(Optional) The sizes of the volumes used to create the snapshots, in gibibytes (GiB). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `Information about one or more snapshots and their permissions.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If ` + "`" + `true` + "`" + `, the resource is public. If ` + "`" + `false` + "`" + `, the resource is private.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshots",
					Description: `Information about one or more snapshots and their permissions.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If ` + "`" + `true` + "`" + `, the resource is public. If ` + "`" + `false` + "`" + `, the resource is private.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "snapshots",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about snapshots.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "account_aliases",
					Description: `(Optional) The account aliases of the owners of the snapshots.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the snapshots.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the snapshots.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume_account_ids",
					Description: `(Optional) The account IDs of one or more users who have permissions to create volumes.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume_global_permission",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, lists all public volumes. If ` + "`" + `false` + "`" + `, lists all private volumes.`,
				},
				resource.Attribute{
					Name:        "progresses",
					Description: `(Optional) The progresses of the snapshots, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `(Optional) The IDs of the snapshots.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the snapshots (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the snapshots.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the snapshots.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the snapshots, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `(Optional) The IDs of the volumes used to create the snapshots.`,
				},
				resource.Attribute{
					Name:        "volume_sizes",
					Description: `(Optional) The sizes of the volumes used to create the snapshots, in gibibytes (GiB). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `Information about one or more snapshots and their permissions.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If ` + "`" + `true` + "`" + `, the resource is public. If ` + "`" + `false` + "`" + `, the resource is private.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshots",
					Description: `Information about one or more snapshots and their permissions.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If ` + "`" + `true` + "`" + `, the resource is public. If ` + "`" + `false` + "`" + `, the resource is private.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "subnet",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific Subnet.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "available_ips_counts",
					Description: `(Optional) The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) The IP ranges in the Subnets, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets in which the Subnets are.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Subnets (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions in which the Subnets are located. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `Information about one or more Subnets.`,
				},
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IP addresses in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If ` + "`" + `true` + "`" + `, a public IP address is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion in which the Subnet is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Subnet.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subnets",
					Description: `Information about one or more Subnets.`,
				},
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IP addresses in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If ` + "`" + `true` + "`" + `, a public IP address is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion in which the Subnet is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Subnet.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "subnets",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Subnets.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "available_ips_counts",
					Description: `(Optional) The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) The IP ranges in the Subnets, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets in which the Subnets are.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Subnets (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions in which the Subnets are located. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `Information about one or more Subnets.`,
				},
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IP addresses in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If ` + "`" + `true` + "`" + `, a public IP address is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion in which the Subnet is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Subnet.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subnets",
					Description: `Information about one or more Subnets.`,
				},
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IP addresses in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If ` + "`" + `true` + "`" + `, a public IP address is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion in which the Subnet is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Subnet.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "subregions",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about subregions.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "subregions",
					Description: `Information about one or more Subregions.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The name of the Region containing the Subregion.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subregion (` + "`" + `available` + "`" + ` \| ` + "`" + `information` + "`" + ` \| ` + "`" + `impaired` + "`" + ` \| ` + "`" + `unavailable` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subregions",
					Description: `Information about one or more Subregions.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The name of the Region containing the Subregion.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subregion (` + "`" + `available` + "`" + ` \| ` + "`" + `information` + "`" + ` \| ` + "`" + `impaired` + "`" + ` \| ` + "`" + `unavailable` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "virtual_gateway",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific virtual gateway.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of the virtual gateways (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "link_net_ids",
					Description: `(Optional) The IDs of the Nets the virtual gateways are attached to.`,
				},
				resource.Attribute{
					Name:        "link_states",
					Description: `(Optional) The current states of the attachments between the virtual gateways and the Nets (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the virtual gateways (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the virtual gateways, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_ids",
					Description: `(Optional) The IDs of the virtual gateways. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "net_to_virtual_gateway_links",
					Description: `The Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the virtual gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "net_to_virtual_gateway_links",
					Description: `The Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the virtual gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "virtual_gateways",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about virtual gateways.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of the virtual gateways (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "link_net_ids",
					Description: `(Optional) The IDs of the Nets the virtual gateways are attached to.`,
				},
				resource.Attribute{
					Name:        "link_states",
					Description: `(Optional) The current states of the attachments between the virtual gateways and the Nets (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the virtual gateways (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the virtual gateways, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_ids",
					Description: `(Optional) The IDs of the virtual gateways. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "virtual_gateways",
					Description: `Information about one or more virtual gateways.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "net_to_virtual_gateway_links",
					Description: `The Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the virtual gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_gateways",
					Description: `Information about one or more virtual gateways.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "net_to_virtual_gateway_links",
					Description: `The Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the virtual gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vm",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific virtual machine (VM).]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the VMs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the VMs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the VMs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `Information about one or more VMs.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Windows VM only) The administrator password of the VM. This password is encrypted with the keypair of the VM and encoded in Base64. You need to wait about 10 minutes after launching the VM to be able to retrieve this attribute. Once the password is ready, this attribute will appear in the Terraform state after the next refresh or apply command. Note also that after the first reboot of the VM, this attribute can no longer be retrieved. For more information on how to use this password to connect to the VM, see [Accessing a Windows Instance](https://wiki.outscale.net/display/EN/Accessing+a+Windows+Instance).`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the VM (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings_created",
					Description: `The block device mapping of the VM.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the created BSU volume.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `Set to ` + "`" + `true` + "`" + ` by default, which means that the volume is deleted when the VM is terminated. If set to ` + "`" + `false` + "`" + `, the volume is not deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "link_date",
					Description: `The time and date of attachment of the volume to the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `If ` + "`" + `true` + "`" + `, the VM is optimized for BSU I/O.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If ` + "`" + `true` + "`" + `, you cannot terminate the VM using Cockpit, the CLI or the API. If ` + "`" + `false` + "`" + `, you can.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the VMs (` + "`" + `ovm` + "`" + ` \| ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI used to create the VM.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, 0, 1, 2, and so on).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the VM is running.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Net only) The network interface cards (NICs) the VMs are attached to.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP address or addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If ` + "`" + `true` + "`" + `, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the NIC.`,
				},
				resource.Attribute{
					Name:        "os_family",
					Description: `Indicates the operating system (OS) of the VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `The performance of the VM (` + "`" + `medium` + "`" + ` \| ` + "`" + `high` + "`" + ` \| ` + "`" + `highest` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `Information about the placement of the VM.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` \| ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI used to create the VM (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, /dev/vda1).`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the VM (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more security groups associated with the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the VM.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VM.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded MIME user data.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `The VM behavior when you stop it. By default or if set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vms",
					Description: `Information about one or more VMs.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Windows VM only) The administrator password of the VM. This password is encrypted with the keypair of the VM and encoded in Base64. You need to wait about 10 minutes after launching the VM to be able to retrieve this attribute. Once the password is ready, this attribute will appear in the Terraform state after the next refresh or apply command. Note also that after the first reboot of the VM, this attribute can no longer be retrieved. For more information on how to use this password to connect to the VM, see [Accessing a Windows Instance](https://wiki.outscale.net/display/EN/Accessing+a+Windows+Instance).`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the VM (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings_created",
					Description: `The block device mapping of the VM.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the created BSU volume.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `Set to ` + "`" + `true` + "`" + ` by default, which means that the volume is deleted when the VM is terminated. If set to ` + "`" + `false` + "`" + `, the volume is not deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "link_date",
					Description: `The time and date of attachment of the volume to the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `If ` + "`" + `true` + "`" + `, the VM is optimized for BSU I/O.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If ` + "`" + `true` + "`" + `, you cannot terminate the VM using Cockpit, the CLI or the API. If ` + "`" + `false` + "`" + `, you can.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the VMs (` + "`" + `ovm` + "`" + ` \| ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI used to create the VM.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, 0, 1, 2, and so on).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the VM is running.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Net only) The network interface cards (NICs) the VMs are attached to.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP address or addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If ` + "`" + `true` + "`" + `, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the NIC.`,
				},
				resource.Attribute{
					Name:        "os_family",
					Description: `Indicates the operating system (OS) of the VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `The performance of the VM (` + "`" + `medium` + "`" + ` \| ` + "`" + `high` + "`" + ` \| ` + "`" + `highest` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `Information about the placement of the VM.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` \| ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI used to create the VM (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, /dev/vda1).`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the VM (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more security groups associated with the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the VM.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VM.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded MIME user data.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `The VM behavior when you stop it. By default or if set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vm_state",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific VM state.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all_vms",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, includes the status of all VMs. By default or if set to ` + "`" + `false` + "`" + `, only includes the status of running VMs.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions of the VMs.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs.`,
				},
				resource.Attribute{
					Name:        "vm_states",
					Description: `(Optional) The states of the VMs (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vm_states",
					Description: `Information about one or more VM states.`,
				},
				resource.Attribute{
					Name:        "maintenance_events",
					Description: `One or more scheduled events associated with the VM.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the event (` + "`" + `system-reboot` + "`" + ` \| ` + "`" + `system-maintenance` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The latest scheduled end time for the event.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `The earliest scheduled start time for the event.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vm_states",
					Description: `Information about one or more VM states.`,
				},
				resource.Attribute{
					Name:        "maintenance_events",
					Description: `One or more scheduled events associated with the VM.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the event (` + "`" + `system-reboot` + "`" + ` \| ` + "`" + `system-maintenance` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The latest scheduled end time for the event.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `The earliest scheduled start time for the event.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vm_states",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about VM states.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all_vms",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, includes the status of all VMs. By default or if set to ` + "`" + `false` + "`" + `, only includes the status of running VMs.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions of the VMs.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs.`,
				},
				resource.Attribute{
					Name:        "vm_states",
					Description: `(Optional) The states of the VMs (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vm_states",
					Description: `Information about one or more VM states.`,
				},
				resource.Attribute{
					Name:        "maintenance_events",
					Description: `One or more scheduled events associated with the VM.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the event (` + "`" + `system-reboot` + "`" + ` \| ` + "`" + `system-maintenance` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The latest scheduled end time for the event.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `The earliest scheduled start time for the event.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vm_states",
					Description: `Information about one or more VM states.`,
				},
				resource.Attribute{
					Name:        "maintenance_events",
					Description: `One or more scheduled events associated with the VM.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the event (` + "`" + `system-reboot` + "`" + ` \| ` + "`" + `system-maintenance` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The latest scheduled end time for the event.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `The earliest scheduled start time for the event.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vm_types",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about VM types.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `(Optional) Indicates whether the VM is optimized for BSU I/O.`,
				},
				resource.Attribute{
					Name:        "memory_sizes",
					Description: `(Optional) The amounts of memory, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "vcore_counts",
					Description: `(Optional) The numbers of vCores.`,
				},
				resource.Attribute{
					Name:        "vm_type_names",
					Description: `(Optional) The names of the VM types. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types).`,
				},
				resource.Attribute{
					Name:        "volume_counts",
					Description: `(Optional) The maximum number of ephemeral storage disks.`,
				},
				resource.Attribute{
					Name:        "volume_sizes",
					Description: `(Optional) The size of one ephemeral storage disk, in gibibytes (GiB). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vm_types",
					Description: `Information about one or more VM types.`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `Indicates whether the VM is optimized for BSU I/O.`,
				},
				resource.Attribute{
					Name:        "max_private_ips",
					Description: `The maximum number of private IP addresses per network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `The amount of memory, in gibibytes.`,
				},
				resource.Attribute{
					Name:        "vcore_count",
					Description: `The number of vCores.`,
				},
				resource.Attribute{
					Name:        "vm_type_name",
					Description: `The name of the VM type.`,
				},
				resource.Attribute{
					Name:        "volume_count",
					Description: `The maximum number of ephemeral storage disks.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of one ephemeral storage disk, in gibibytes (GiB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vm_types",
					Description: `Information about one or more VM types.`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `Indicates whether the VM is optimized for BSU I/O.`,
				},
				resource.Attribute{
					Name:        "max_private_ips",
					Description: `The maximum number of private IP addresses per network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `The amount of memory, in gibibytes.`,
				},
				resource.Attribute{
					Name:        "vcore_count",
					Description: `The number of vCores.`,
				},
				resource.Attribute{
					Name:        "vm_type_name",
					Description: `The name of the VM type.`,
				},
				resource.Attribute{
					Name:        "volume_count",
					Description: `The maximum number of ephemeral storage disks.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of one ephemeral storage disk, in gibibytes (GiB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vms",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about virtual machines (VMs).]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the VMs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the VMs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the VMs, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `Information about one or more VMs.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Windows VM only) The administrator password of the VM. This password is encrypted with the keypair of the VM and encoded in Base64. You need to wait about 10 minutes after launching the VM to be able to retrieve this attribute. Once the password is ready, this attribute will appear in the Terraform state after the next refresh or apply command. Note also that after the first reboot of the VM, this attribute can no longer be retrieved. For more information on how to use this password to connect to the VM, see [Accessing a Windows Instance](https://wiki.outscale.net/display/EN/Accessing+a+Windows+Instance).`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the VM (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings_created",
					Description: `The block device mapping of the VM.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the created BSU volume.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `Set to ` + "`" + `true` + "`" + ` by default, which means that the volume is deleted when the VM is terminated. If set to ` + "`" + `false` + "`" + `, the volume is not deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "link_date",
					Description: `The time and date of attachment of the volume to the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `If ` + "`" + `true` + "`" + `, the VM is optimized for BSU I/O.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If ` + "`" + `true` + "`" + `, you cannot terminate the VM using Cockpit, the CLI or the API. If ` + "`" + `false` + "`" + `, you can.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the VMs (` + "`" + `ovm` + "`" + ` \| ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI used to create the VM.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, 0, 1, 2, and so on).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the VM is running.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Net only) The network interface cards (NICs) the VMs are attached to.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP address or addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If ` + "`" + `true` + "`" + `, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the NIC.`,
				},
				resource.Attribute{
					Name:        "os_family",
					Description: `Indicates the operating system (OS) of the VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `The performance of the VM (` + "`" + `medium` + "`" + ` \| ` + "`" + `high` + "`" + ` \| ` + "`" + `highest` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `Information about the placement of the VM.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` \| ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI used to create the VM (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, /dev/vda1).`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the VM (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more security groups associated with the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the VM.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VM.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded MIME user data.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `The VM behavior when you stop it. By default or if set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vms",
					Description: `Information about one or more VMs.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Windows VM only) The administrator password of the VM. This password is encrypted with the keypair of the VM and encoded in Base64. You need to wait about 10 minutes after launching the VM to be able to retrieve this attribute. Once the password is ready, this attribute will appear in the Terraform state after the next refresh or apply command. Note also that after the first reboot of the VM, this attribute can no longer be retrieved. For more information on how to use this password to connect to the VM, see [Accessing a Windows Instance](https://wiki.outscale.net/display/EN/Accessing+a+Windows+Instance).`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the VM (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings_created",
					Description: `The block device mapping of the VM.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the created BSU volume.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `Set to ` + "`" + `true` + "`" + ` by default, which means that the volume is deleted when the VM is terminated. If set to ` + "`" + `false` + "`" + `, the volume is not deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "link_date",
					Description: `The time and date of attachment of the volume to the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `If ` + "`" + `true` + "`" + `, the VM is optimized for BSU I/O.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If ` + "`" + `true` + "`" + `, you cannot terminate the VM using Cockpit, the CLI or the API. If ` + "`" + `false` + "`" + `, you can.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the VMs (` + "`" + `ovm` + "`" + ` \| ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI used to create the VM.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, 0, 1, 2, and so on).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the VM is running.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Net only) The network interface cards (NICs) the VMs are attached to.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If ` + "`" + `true` + "`" + `, the source/destination check is enabled. If ` + "`" + `false` + "`" + `, it is disabled. This value must be ` + "`" + `false` + "`" + ` for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between 1 and 7, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP address or addresses of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If ` + "`" + `true` + "`" + `, the IP address is the primary private IP address of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the EIP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The External IP address (EIP) associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the NIC.`,
				},
				resource.Attribute{
					Name:        "os_family",
					Description: `Indicates the operating system (OS) of the VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `The performance of the VM (` + "`" + `medium` + "`" + ` \| ` + "`" + `high` + "`" + ` \| ` + "`" + `highest` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `Information about the placement of the VM.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` \| ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI used to create the VM (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, /dev/vda1).`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the VM (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more security groups associated with the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the VM.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VM.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded MIME user data.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `The VM behavior when you stop it. By default or if set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "volume",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific volume.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "creation_dates",
					Description: `(Optional) The dates and times at which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "link_volume_delete_on_vm_deletion",
					Description: `(Optional) Indicates whether the volumes are deleted when terminating the VMs.`,
				},
				resource.Attribute{
					Name:        "link_volume_device_names",
					Description: `(Optional) The VM device names.`,
				},
				resource.Attribute{
					Name:        "link_volume_link_dates",
					Description: `(Optional) The dates and times at which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "link_volume_link_states",
					Description: `(Optional) The attachment states of the volumes (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_volume_vm_ids",
					Description: `(Optional) One or more IDs of VMs.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `(Optional) The snapshots from which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions in which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the volumes.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the volumes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the volumes, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `(Optional) The IDs of the volumes.`,
				},
				resource.Attribute{
					Name:        "volume_sizes",
					Description: `(Optional) The sizes of the volumes, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "volume_states",
					Description: `(Optional) The states of the volumes (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "volume_types",
					Description: `(Optional) The types of the volumes (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `Information about one or more volumes.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS): For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS. For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot from which the volume was created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the volume was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the volume.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "volumes",
					Description: `Information about one or more volumes.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS): For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS. For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot from which the volume was created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the volume was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the volume.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "volumes",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about volumes.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "creation_dates",
					Description: `(Optional) The dates and times at which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "link_volume_delete_on_vm_deletion",
					Description: `(Optional) Indicates whether the volumes are deleted when terminating the VMs.`,
				},
				resource.Attribute{
					Name:        "link_volume_device_names",
					Description: `(Optional) The VM device names.`,
				},
				resource.Attribute{
					Name:        "link_volume_link_dates",
					Description: `(Optional) The dates and times at which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "link_volume_link_states",
					Description: `(Optional) The attachment states of the volumes (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_volume_vm_ids",
					Description: `(Optional) One or more IDs of VMs.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `(Optional) The snapshots from which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions in which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the volumes.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the volumes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combination of the tags associated with the volumes, in the following format: "Filters":{"Tags":["TAGKEY=TAGVALUE"]}.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `(Optional) The IDs of the volumes.`,
				},
				resource.Attribute{
					Name:        "volume_sizes",
					Description: `(Optional) The sizes of the volumes, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "volume_states",
					Description: `(Optional) The states of the volumes (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "volume_types",
					Description: `(Optional) The types of the volumes (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `Information about one or more volumes.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS): For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS. For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot from which the volume was created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the volume was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the volume.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "volumes",
					Description: `Information about one or more volumes.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS): For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS. For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If ` + "`" + `true` + "`" + `, the volume is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot from which the volume was created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the volume was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the volume.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vpn_connection",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a specific VPN connection.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "bgp_asns",
					Description: `(Optional) The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_ids",
					Description: `(Optional) The IDs of the client gateways.`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of the VPN connections (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "route_destination_ip_ranges",
					Description: `(Optional) The destination IP ranges.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the VPN connections (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `(Optional) If ` + "`" + `false` + "`" + `, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If ` + "`" + `true` + "`" + `, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_ids",
					Description: `(Optional) The IDs of the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_ids",
					Description: `(Optional) The IDs of the VPN connections. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `The configuration to apply to the client gateway to establish the VPN connection, in XML format.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway used on the client end of the connection.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection (always ` + "`" + `ipsec.1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Information about one or more static routes associated with the VPN connection, if any.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `The type of route (always ` + "`" + `static` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the static route (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VPN connection (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `If ` + "`" + `false` + "`" + `, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If ` + "`" + `true` + "`" + `, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VPN connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the 3DS OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `The configuration to apply to the client gateway to establish the VPN connection, in XML format.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway used on the client end of the connection.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection (always ` + "`" + `ipsec.1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Information about one or more static routes associated with the VPN connection, if any.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `The type of route (always ` + "`" + `static` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the static route (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VPN connection (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `If ` + "`" + `false` + "`" + `, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If ` + "`" + `true` + "`" + `, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VPN connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the 3DS OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vpn_connections",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about VPN connections.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `One or more filters.`,
				},
				resource.Attribute{
					Name:        "bgp_asns",
					Description: `(Optional) The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_ids",
					Description: `(Optional) The IDs of the client gateways.`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of the VPN connections (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "route_destination_ip_ranges",
					Description: `(Optional) The destination IP ranges.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the VPN connections (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `(Optional) If ` + "`" + `false` + "`" + `, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If ` + "`" + `true` + "`" + `, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_ids",
					Description: `(Optional) The IDs of the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_ids",
					Description: `(Optional) The IDs of the VPN connections. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpn_connections",
					Description: `Information about one or more VPN connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `The configuration to apply to the client gateway to establish the VPN connection, in XML format.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway used on the client end of the connection.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection (always ` + "`" + `ipsec.1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Information about one or more static routes associated with the VPN connection, if any.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `The type of route (always ` + "`" + `static` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the static route (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VPN connection (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `If ` + "`" + `false` + "`" + `, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If ` + "`" + `true` + "`" + `, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VPN connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the 3DS OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpn_connections",
					Description: `Information about one or more VPN connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `The configuration to apply to the client gateway to establish the VPN connection, in XML format.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway used on the client end of the connection.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection (always ` + "`" + `ipsec.1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Information about one or more static routes associated with the VPN connection, if any.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `The type of route (always ` + "`" + `static` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the static route (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VPN connection (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `If ` + "`" + `false` + "`" + `, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If ` + "`" + `true` + "`" + `, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VPN connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the 3DS OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"access_key":                   0,
		"access_keys":                  1,
		"client_gateway":               2,
		"client_gateways":              3,
		"dhcp_option":                  4,
		"dhcp_options":                 5,
		"flexible_gpu":                 6,
		"flexible_gpu_catalog":         7,
		"flexible_gpus":                8,
		"image":                        9,
		"images":                       10,
		"internet_service":             11,
		"internet_services":            12,
		"keypair":                      13,
		"keypairs":                     14,
		"load_balancer":                15,
		"load_balancer_listener_rule":  16,
		"load_balancer_listener_rules": 17,
		"load_balancer_vm_health":      18,
		"load_balancers":               19,
		"nat_service":                  20,
		"nat_services":                 21,
		"net":                          22,
		"net_access_point":             23,
		"net_access_point_services":    24,
		"net_access_points":            25,
		"net_attributes":               26,
		"net_peering":                  27,
		"net_peerings":                 28,
		"nets":                         29,
		"nic":                          30,
		"nics":                         31,
		"public_ip":                    32,
		"public_ips":                   33,
		"regions":                      34,
		"route_table":                  35,
		"route_tables":                 36,
		"security_group":               37,
		"security_groups":              38,
		"snapshot":                     39,
		"snapshots":                    40,
		"subnet":                       41,
		"subnets":                      42,
		"subregions":                   43,
		"virtual_gateway":              44,
		"virtual_gateways":             45,
		"vm":                           46,
		"vm_state":                     47,
		"vm_states":                    48,
		"vm_types":                     49,
		"vms":                          50,
		"volume":                       51,
		"volumes":                      52,
		"vpn_connection":               53,
		"vpn_connections":              54,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
