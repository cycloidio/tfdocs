package outscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "outscale_access_key",
			Category:         "Resources",
			ShortDescription: `[Manages an access key.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "expiration_date",
					Description: `(Optional) The date and time at which you want the access key to expire, in ISO 8601 format (for example, ` + "`" + `2017-06-14` + "`" + ` or ` + "`" + `2017-06-14T00:00:00Z` + "`" + `). To remove an existing expiration date, use the method without specifying this parameter.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The state for the access key (` + "`" + `ACTIVE` + "`" + ` | ` + "`" + `INACTIVE` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time (UTC) of creation of the access key.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date and time (UTC) at which the access key expires.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time (UTC) of the last modification of the access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The access key that enables you to send requests.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not). ## Import An access key can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_access_key.ImportedAccessKey ABCDEFGHIJ0123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time (UTC) of creation of the access key.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date and time (UTC) at which the access key expires.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time (UTC) of the last modification of the access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The access key that enables you to send requests.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not). ## Import An access key can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_access_key.ImportedAccessKey ABCDEFGHIJ0123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_api_access_policy",
			Category:         "Resources",
			ShortDescription: `[Manages the API access policy.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "max_access_key_expiration_seconds",
					Description: `(Required) The maximum possible lifetime for your access keys, in seconds (between ` + "`" + `0` + "`" + ` and ` + "`" + `3153600000` + "`" + `, both included). If set to ` + "`" + `O` + "`" + `, your access keys can have unlimited lifetimes, but a trusted session cannot be activated. Otherwise, all your access keys must have an expiration date. This value must be greater than the remaining lifetime of each access key of your account.`,
				},
				resource.Attribute{
					Name:        "require_trusted_env",
					Description: `(Required) If true, a trusted session is activated, provided that you specify the ` + "`" + `max_access_key_expiration_seconds` + "`" + ` parameter with a value greater than ` + "`" + `0` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "max_access_key_expiration_seconds",
					Description: `The maximum possible lifetime for your access keys, in seconds. If ` + "`" + `0` + "`" + `, your access keys can have unlimited lifetimes.`,
				},
				resource.Attribute{
					Name:        "require_trusted_env",
					Description: `If true, a trusted session is activated, allowing you to bypass Certificate Authorities (CAs) enforcement. For more information, see the ` + "`" + `ApiKeyAuth` + "`" + ` authentication scheme in the [Authentication](https://docs.outscale.com/api#authentication) section.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "max_access_key_expiration_seconds",
					Description: `The maximum possible lifetime for your access keys, in seconds. If ` + "`" + `0` + "`" + `, your access keys can have unlimited lifetimes.`,
				},
				resource.Attribute{
					Name:        "require_trusted_env",
					Description: `If true, a trusted session is activated, allowing you to bypass Certificate Authorities (CAs) enforcement. For more information, see the ` + "`" + `ApiKeyAuth` + "`" + ` authentication scheme in the [Authentication](https://docs.outscale.com/api#authentication) section.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_api_access_rule",
			Category:         "Resources",
			ShortDescription: `[Manages an API access rule.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ca_ids",
					Description: `(Optional) One or more IDs of Client Certificate Authorities (CAs).`,
				},
				resource.Attribute{
					Name:        "cns",
					Description: `(Optional) One or more Client Certificate Common Names (CNs). If this parameter is specified, you must also specify the ` + "`" + `ca_ids` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the API access rule.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) One or more IP addresses or CIDR blocks (for example, ` + "`" + `192.0.2.0/16` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "api_access_rule_id",
					Description: `The ID of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `One or more IDs of Client Certificate Authorities (CAs) used for the API access rule.`,
				},
				resource.Attribute{
					Name:        "cns",
					Description: `One or more Client Certificate Common Names (CNs).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges used for the API access rule, in CIDR notation (for example, ` + "`" + `192.0.2.0/16` + "`" + `). ## Import An API access rule can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_api_access_rule.ImportedAPIAccessRule "aar-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_access_rule_id",
					Description: `The ID of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `One or more IDs of Client Certificate Authorities (CAs) used for the API access rule.`,
				},
				resource.Attribute{
					Name:        "cns",
					Description: `One or more Client Certificate Common Names (CNs).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges used for the API access rule, in CIDR notation (for example, ` + "`" + `192.0.2.0/16` + "`" + `). ## Import An API access rule can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_api_access_rule.ImportedAPIAccessRule "aar-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_ca",
			Category:         "Resources",
			ShortDescription: `[Manages a Certificate Authority (CA).]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ca_pem",
					Description: `(Required) The CA in PEM format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the CA. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ca_fingerprint",
					Description: `The fingerprint of the CA.`,
				},
				resource.Attribute{
					Name:        "ca_id",
					Description: `The ID of the CA.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the CA. ## Import A CA can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_ca.ImportedCa ca-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ca_fingerprint",
					Description: `The fingerprint of the CA.`,
				},
				resource.Attribute{
					Name:        "ca_id",
					Description: `The ID of the CA.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the CA. ## Import A CA can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_ca.ImportedCa ca-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_client_gateway",
			Category:         "Resources",
			ShortDescription: `[Manages a client gateway.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `(Required) The Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet. This number must be between ` + "`" + `1` + "`" + ` and ` + "`" + `4294967295` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Required) The communication protocol used to establish tunnel with your client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Required) The public fixed IPv4 address of your client gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `The Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A client gateway can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_client_gateway.ImportedClientGateway cgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `The Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A client gateway can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_client_gateway.ImportedClientGateway cgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_dhcp_option",
			Category:         "Resources",
			ShortDescription: `[Manages a DHCP option.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `(Optional) The IPs of domain name servers. If no IPs are specified, the ` + "`" + `OutscaleProvidedDNS` + "`" + ` value is set by default. You must specify at least one of the following parameters: ` + "`" + `DomainName` + "`" + `, ` + "`" + `DomainNameServers` + "`" + `, or ` + "`" + `NtpServers` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) Specify a domain name (for example, MyCompany.com). You can specify only one domain name. You must specify at least one of the following parameters: ` + "`" + `DomainName` + "`" + `, ` + "`" + `DomainNameServers` + "`" + `, or ` + "`" + `NtpServers` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_servers",
					Description: `(Optional) The IPs of the log servers. You must specify at least one of the following parameters: ` + "`" + `domain_name` + "`" + `, ` + "`" + `domain_name_servers` + "`" + `, ` + "`" + `log_servers` + "`" + `, or ` + "`" + `ntp_servers` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `(Optional) The IPs of the Network Time Protocol (NTP) servers. You must specify at least one of the following parameters: ` + "`" + `DomainName` + "`" + `, ` + "`" + `DomainNameServers` + "`" + `, or ` + "`" + `NtpServers` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If true, the DHCP options set is a default one. If false, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IPs for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "log_servers",
					Description: `One or more IPs for the log servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IPs for the NTP servers.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import DHCP options can be imported using the DHCP option ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_dhcp_option.ImportedDhcpSet dopt-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "default",
					Description: `If true, the DHCP options set is a default one. If false, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IPs for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "log_servers",
					Description: `One or more IPs for the log servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IPs for the NTP servers.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import DHCP options can be imported using the DHCP option ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_dhcp_option.ImportedDhcpSet dopt-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_flexible_gpu",
			Category:         "Resources",
			ShortDescription: `[Manages a flexible GPU.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `(Optional) The processor generation that the fGPU must be compatible with. If not specified, the oldest possible processor generation is selected (as provided by [ReadFlexibleGpuCatalog](https://docs.outscale.com/api#readflexiblegpucatalog) for the specified model of fGPU).`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `(Required) The model of fGPU you want to allocate. For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `(Required) The Subregion in which you want to create the fGPU. ## Attribute Reference The following attributes are exported:`,
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
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html).`,
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
					Description: `The ID of the VM the fGPU is attached to, if any. ## Import A flexible GPU can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_flexible_gpu.imported_fgpu fgpu-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html).`,
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
					Description: `The ID of the VM the fGPU is attached to, if any. ## Import A flexible GPU can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_flexible_gpu.imported_fgpu fgpu-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_flexible_gpu_link",
			Category:         "Resources",
			ShortDescription: `[Manages a flexible GPU link.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `(Required) The ID of the fGPU you want to attach.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Required) The ID of the VM you want to attach the fGPU to. ## Attribute Reference No attribute is exported. ## Import A flexible GPU link can be imported using the flexible GPU ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_flexible_gpu_link.imported_link_fgpu fgpu-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_image",
			Category:         "Resources",
			ShortDescription: `[Manages an image.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "architecture",
					Description: `(Optional) The architecture of the OMI (by default, ` + "`" + `i386` + "`" + ` if you specified the ` + "`" + `file_location` + "`" + ` or ` + "`" + `root_device_name` + "`" + ` parameter).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `(Optional) One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + ` with a maximum performance ratio of 300 IOPS per gibibyte.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Optional) The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [About Volumes > Volume Types and IOPS](https://docs.outscale.com/en/userguide/About-Volumes.html#_volume_types_and_iops).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) The device name for the volume. For a root device, you must use ` + "`" + `/dev/sda1` + "`" + `. For other volumes, you must use ` + "`" + `/dev/sdX` + "`" + `, ` + "`" + `/dev/sdXX` + "`" + `, ` + "`" + `/dev/xvdX` + "`" + `, or ` + "`" + `/dev/xvdXX` + "`" + ` (where the first ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `b` + "`" + ` and ` + "`" + `z` + "`" + `, and the second ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `a` + "`" + ` and ` + "`" + `z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `(Optional) The name of the virtual device (` + "`" + `ephemeralN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the new OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `(Optional) The pre-signed URL of the OMI manifest file, or the full path to the OMI stored in a bucket. If you specify this parameter, a copy of the OMI is created in your account. You must specify only one of the following parameters: ` + "`" + `file_location` + "`" + `, ` + "`" + `root_device_name` + "`" + `, ` + "`" + `source_image_id` + "`" + ` or ` + "`" + `vm_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional) A unique name for the new OMI.<br /> Constraints: 3-128 alphanumeric characters, underscores (_), spaces ( ), parentheses (()), slashes (/), periods (.), or dashes (-).`,
				},
				resource.Attribute{
					Name:        "no_reboot",
					Description: `(Optional) If false, the VM shuts down before creating the OMI and then reboots. If true, the VM does not.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `(Optional) The name of the root device. You must specify only one of the following parameters: ` + "`" + `file_location` + "`" + `, ` + "`" + `root_device_name` + "`" + `, ` + "`" + `source_image_id` + "`" + ` or ` + "`" + `vm_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_image_id",
					Description: `(Optional) The ID of the OMI you want to copy. You must specify only one of the following parameters: ` + "`" + `file_location` + "`" + `, ` + "`" + `root_device_name` + "`" + `, ` + "`" + `source_image_id` + "`" + ` or ` + "`" + `vm_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_region_name",
					Description: `(Optional) The name of the source Region, which must be the same as the Region of your account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Optional) The ID of the VM from which you want to create the OMI. You must specify only one of the following parameters: ` + "`" + `file_location` + "`" + `, ` + "`" + `root_device_name` + "`" + `, ` + "`" + `source_image_id` + "`" + ` or ` + "`" + `vm_id` + "`" + `. ## Attribute Reference The following attributes are exported:`,
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
					Description: `By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + ` with a maximum performance ratio of 300 IOPS per gibibyte.`,
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
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [About Volumes > Volume Types and IOPS](https://docs.outscale.com/en/userguide/About-Volumes.html#_volume_types_and_iops).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The device name for the volume. For a root device, you must use ` + "`" + `/dev/sda1` + "`" + `. For other volumes, you must use ` + "`" + `/dev/sdX` + "`" + `, ` + "`" + `/dev/sdXX` + "`" + `, ` + "`" + `/dev/xvdX` + "`" + `, or ` + "`" + `/dev/xvdXX` + "`" + ` (where the first ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `b` + "`" + ` and ` + "`" + `z` + "`" + `, and the second ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `a` + "`" + ` and ` + "`" + `z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (` + "`" + `ephemeralN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the OMI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location of the bucket where the OMI files are stored.`,
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
					Description: `If true, the resource is public. If false, the resource is private.`,
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
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import An image can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_image.ImportedImage ami-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + ` with a maximum performance ratio of 300 IOPS per gibibyte.`,
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
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [About Volumes > Volume Types and IOPS](https://docs.outscale.com/en/userguide/About-Volumes.html#_volume_types_and_iops).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The device name for the volume. For a root device, you must use ` + "`" + `/dev/sda1` + "`" + `. For other volumes, you must use ` + "`" + `/dev/sdX` + "`" + `, ` + "`" + `/dev/sdXX` + "`" + `, ` + "`" + `/dev/xvdX` + "`" + `, or ` + "`" + `/dev/xvdXX` + "`" + ` (where the first ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `b` + "`" + ` and ` + "`" + `z` + "`" + `, and the second ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `a` + "`" + ` and ` + "`" + `z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (` + "`" + `ephemeralN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the OMI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location of the bucket where the OMI files are stored.`,
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
					Description: `If true, the resource is public. If false, the resource is private.`,
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
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import An image can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_image.ImportedImage ami-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_image_export_task",
			Category:         "Resources",
			ShortDescription: `[Manages an image export task.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) The ID of the OMI to export.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the OOS export task to create.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `(Optional) The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_api_key",
					Description: `Information about the OOS API key.`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `(Optional) The API key of the OOS account that enables you to access the bucket.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The secret key of the OOS account that enables you to access the bucket.`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `(Optional) The name of the OOS bucket where you want to export the object.`,
				},
				resource.Attribute{
					Name:        "osu_manifest_url",
					Description: `(Optional) The URL of the manifest file.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `(Optional) The prefix for the key of the OOS object.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `If the OMI export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI to be exported.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the OMI export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the OMI is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_manifest_url",
					Description: `The URL of the manifest file.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the image.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the OMI export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI export task (` + "`" + `pending/queued` + "`" + ` \| ` + "`" + `pending` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `cancelled` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the image export task.`,
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
					Name:        "task_id",
					Description: `The ID of the OMI export task.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "comment",
					Description: `If the OMI export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI to be exported.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the OMI export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the OMI is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_manifest_url",
					Description: `The URL of the manifest file.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the image.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the OMI export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI export task (` + "`" + `pending/queued` + "`" + ` \| ` + "`" + `pending` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `cancelled` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the image export task.`,
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
					Name:        "task_id",
					Description: `The ID of the OMI export task.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_image_launch_permission",
			Category:         "Resources",
			ShortDescription: `[Manages an image launch permission.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) The ID of the OMI you want to modify.`,
				},
				resource.Attribute{
					Name:        "permission_additions",
					Description: `(Optional) Information about the users to whom you want to give permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account ID of one or more users to whom you want to give permissions.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `(Optional) If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "permission_removals",
					Description: `(Optional) Information about the users from whom you want to remove permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account ID of one or more users from whom you want to remove permissions.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `(Optional) If true, the resource is public. If false, the resource is private. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
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
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
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
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_internet_service",
			Category:         "Resources",
			ShortDescription: `[Manages an Internet service.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
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
					Description: `The state of the attachment of the Internet service to the Net (always ` + "`" + `available` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import An internet service can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_internet_service.ImportedInternetService igw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The state of the attachment of the Internet service to the Net (always ` + "`" + `available` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import An internet service can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_internet_service.ImportedInternetService igw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_internet_service_link",
			Category:         "Resources",
			ShortDescription: `[Manages an Internet service link.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `(Required) The ID of the Internet service you want to attach.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net to which you want to attach the Internet service. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Internet service to the Net (always ` + "`" + `available` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import An internet service link can be imported using the internet service ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_internet_service_link.ImportedInternetServiceLink igw-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Internet service to the Net (always ` + "`" + `available` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import An internet service link can be imported using the internet service ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_internet_service_link.ImportedInternetServiceLink igw-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_keypair",
			Category:         "Resources",
			ShortDescription: `[Manages a keypair.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "keypair_name",
					Description: `(Required) A unique name for the keypair, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters).`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The public key. It must be Base64-encoded. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `The MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key. When saving the private key in a .rsa file, replace the ` + "`" + `\n` + "`" + ` escape sequences with line breaks. ## Import A keypair can be imported using its name. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_keypair.ImportedKeypair Name-of-the-Keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `The MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key. When saving the private key in a .rsa file, replace the ` + "`" + `\n` + "`" + ` escape sequences with line breaks. ## Import A keypair can be imported using its name. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_keypair.ImportedKeypair Name-of-the-Keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_load_balancer",
			Category:         "Resources",
			ShortDescription: `[Manages a load balancer.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listeners",
					Description: `(Required) One or more listeners to create.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `(Optional) The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `(Optional) The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `(Optional) The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `(Optional) The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `(Optional) The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Required) The unique name of the load balancer (32 alphanumeric or hyphen characters maximum, but cannot start or end with a hyphen).`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `(Optional) The type of load balancer: ` + "`" + `internet-facing` + "`" + ` or ` + "`" + `internal` + "`" + `. Use this parameter only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) (internet-facing only) The public IP you want to associate with the load balancer. If not specified, a public IP owned by 3DS OUTSCALE is associated.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) (Net only) One or more IDs of security groups you want to assign to the load balancer. If not specified, the default security group of the Net is assigned to the load balancer.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) (Net only) The ID of the Subnet in which you want to create the load balancer. Regardless of this Subnet, the load balancer can distribute traffic to all Subnets. This parameter is required in a Net.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) (public Cloud only) The Subregion in which you want to create the load balancer. Regardless of this Subregion, the load balancer can distribute traffic to all Subregions. This parameter is required in the public Cloud.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either ` + "`" + `5` + "`" + ` or ` + "`" + `60` + "`" + ` (by default, ` + "`" + `60` + "`" + `).`,
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
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
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
					Description: `The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
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
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
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
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(internet-facing only) The public IP associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "secured_cookies",
					Description: `Whether secure cookies are enabled for the load balancer.`,
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
					Description: `The ID of the Subnet in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `The ID of the Subregion in which the load balancer was created.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A load balancer can be imported using its name. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_load_balancer.ImportedLbu Name-of-the-Lbu ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either ` + "`" + `5` + "`" + ` or ` + "`" + `60` + "`" + ` (by default, ` + "`" + `60` + "`" + `).`,
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
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
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
					Description: `The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
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
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
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
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(internet-facing only) The public IP associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "secured_cookies",
					Description: `Whether secure cookies are enabled for the load balancer.`,
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
					Description: `The ID of the Subnet in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `The ID of the Subregion in which the load balancer was created.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A load balancer can be imported using its name. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_load_balancer.ImportedLbu Name-of-the-Lbu ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_load_balancer_attributes",
			Category:         "Resources",
			ShortDescription: `[Manages load balancer attributes.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `(Optional) The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `(Optional) The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `(Optional) The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either ` + "`" + `5` + "`" + ` or ` + "`" + `60` + "`" + ` (by default, ` + "`" + `60` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `(Optional) The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Required) The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `(Optional) The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included). This parameter is required if you want to update the server certificate.`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `(Optional) The name of the policy you want to enable for the listener.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `(Optional) The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns). If this parameter is specified, you must also specify the ` + "`" + `load_balancer_port` + "`" + ` parameter. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either ` + "`" + `5` + "`" + ` or ` + "`" + `60` + "`" + ` (by default, ` + "`" + `60` + "`" + `).`,
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
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
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
					Description: `The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
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
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
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
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP.`,
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
					Description: `The ID of the Subnet in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `The ID of the Subregion in which the load balancer was created.`,
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
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either ` + "`" + `5` + "`" + ` or ` + "`" + `60` + "`" + ` (by default, ` + "`" + `60` + "`" + `).`,
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
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
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
					Description: `The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
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
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
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
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP.`,
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
					Description: `The ID of the Subnet in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `The ID of the Subregion in which the load balancer was created.`,
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
			Type:             "outscale_load_balancer_listener_rule",
			Category:         "Resources",
			ShortDescription: `[Manages a load balancer listener rule.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_rule",
					Description: `Information about the listener rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `(Optional) A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `(Optional) A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `(Optional) A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "listener",
					Description: `Information about the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Optional) The name of the load balancer to which the listener is attached.`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `(Optional) The port of load balancer on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Required) The IDs of the backend VMs. ## Attribute Reference The following attributes are exported:`,
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
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].`,
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
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].`,
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
			Type:             "outscale_load_balancer_policy",
			Category:         "Resources",
			ShortDescription: `[Manages a load balancer policy.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Optional) The name of the application cookie used for stickiness. This parameter is required if you create a stickiness policy based on an application-generated cookie.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Required) The name of the load balancer for which you want to create a policy.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) The name of the policy. This name must be unique and consist of alphanumeric characters and dashes (-).`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Required) The type of stickiness policy you want to create: ` + "`" + `app` + "`" + ` or ` + "`" + `load_balancer` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either ` + "`" + `5` + "`" + ` or ` + "`" + `60` + "`" + ` (by default, ` + "`" + `60` + "`" + `).`,
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
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
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
					Description: `The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
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
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
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
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP.`,
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
					Description: `The ID of the Subnet in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `The ID of the Subregion in which the load balancer was created.`,
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
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either ` + "`" + `5` + "`" + ` or ` + "`" + `60` + "`" + ` (by default, ` + "`" + `60` + "`" + `).`,
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
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
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
					Description: `The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
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
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
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
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP.`,
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
					Description: `The ID of the Subnet in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `The ID of the Subregion in which the load balancer was created.`,
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
			Type:             "outscale_load_balancer_vms",
			Category:         "Resources",
			ShortDescription: `[Manages load balancer VMs.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `(Required) One or more IDs of back-end VMs.<br /> Specifying the same ID several times has no effect as each back-end VM has equal weight.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Required) The name of the load balancer. ## Attribute Reference No attribute is exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_nat_service",
			Category:         "Resources",
			ShortDescription: `[Manages a NAT service.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `(Required) The allocation ID of the public IP to associate with the NAT service.<br /> If the public IP is already associated with another resource, you must first disassociate it.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the Subnet in which you want to create the NAT service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
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
					Description: `Information about the public IP or IPs associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP associated with the NAT service.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A NAT service can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_nat_service.ImportedNatService nat-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Information about the public IP or IPs associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP associated with the NAT service.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A NAT service can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_nat_service.ImportedNatService nat-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net",
			Category:         "Resources",
			ShortDescription: `[Manages a Net.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Required) The IP range for the Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `(Optional) The tenancy options for the VMs (` + "`" + `default` + "`" + ` if a VM created in a Net can be launched with any tenancy, ` + "`" + `dedicated` + "`" + ` if it can be launched with dedicated tenancy VMs running on single-tenant hardware). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
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
					Description: `The VM tenancy in a Net. ## Import A Net can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_net.ImportedNet vpc-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
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
					Description: `The VM tenancy in a Net. ## Import A Net can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_net.ImportedNet vpc-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net_access_point",
			Category:         "Resources",
			ShortDescription: `[Manages a Net access point.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `(Optional) One or more IDs of route tables to use for the connection.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The name of the service (in the format ` + "`" + `com.outscale.region.service` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A Net access point can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_net_access_point.ImportedNetAccessPoint vpce-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A Net access point can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_net_access_point.ImportedNetAccessPoint vpce-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net_attributes",
			Category:         "Resources",
			ShortDescription: `[Manages the attributes of a Net.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `(Required) The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
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
					Description: `The VM tenancy in a Net. ## Import A Net attribute can be imported using the Net ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_net_attributes.ImportedNet vpc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
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
					Description: `The VM tenancy in a Net. ## Import A Net attribute can be imported using the Net ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_net_attributes.ImportedNet vpc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net_peering",
			Category:         "Resources",
			ShortDescription: `[Manages a Net peering.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accepter_net_id",
					Description: `(Required) The ID of the Net you want to connect with.`,
				},
				resource.Attribute{
					Name:        "source_net_id",
					Description: `(Required) The ID of the Net you send the peering request from.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
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
					Description: `The IP range for the accepter Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
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
					Description: `The IP range for the source Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A Net peering can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_net_peering.ImportedNetPeering pcx-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The IP range for the accepter Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
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
					Description: `The IP range for the source Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters. ## Import A Net peering can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_net_peering.ImportedNetPeering pcx-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net_peering_acceptation",
			Category:         "Resources",
			ShortDescription: `[Manages a Net peering acceptation.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `(Required) The ID of the Net peering you want to accept. ## Attribute Reference The following attributes are exported:`,
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
					Description: `The IP range for the accepter Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
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
					Description: `The IP range for the source Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering.`,
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
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
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
					Description: `The IP range for the source Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering.`,
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
			Type:             "outscale_nic",
			Category:         "Resources",
			ShortDescription: `[Manages a network interface card (NIC).]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the NIC.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) The primary private IP for the NIC.<br /> This IP must be within the IP range of the Subnet that you specify with the ` + "`" + `subnet_id` + "`" + ` attribute.<br /> If you do not specify this attribute, a random private IP is selected within the IP range of the Subnet.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `(Optional) If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) The private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the Subnet in which you want to create the NIC.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
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
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
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
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
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
					Description: `The private IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP of the NIC.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A NIC can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_nic.ImportedNic eni-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
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
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
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
					Description: `The private IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP of the NIC.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A NIC can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_nic.ImportedNic eni-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_nic_link",
			Category:         "Resources",
			ShortDescription: `[Manages a NIC link.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_number",
					Description: `(Required) The index of the VM device for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Required) The ID of the NIC you want to attach.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Required) The ID of the VM to which you want to attach the NIC. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC attachment. ## Import A NIC link can be imported using the NIC ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_nic_link.ImportedNicLink eni-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC attachment. ## Import A NIC link can be imported using the NIC ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_nic_link.ImportedNicLink eni-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_nic_private_ip",
			Category:         "Resources",
			ShortDescription: `[Manages a NIC's private IPs.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_relink",
					Description: `(Optional) If true, allows an IP that is already assigned to another NIC in the same Subnet to be assigned to the NIC you specified.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Required) The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) The secondary private IP or IPs you want to assign to the NIC within the IP range of the Subnet.`,
				},
				resource.Attribute{
					Name:        "secondary_private_ip_count",
					Description: `(Optional) The number of secondary private IPs to assign to the NIC. ## Attribute Reference No attribute is exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_public_ip",
			Category:         "Resources",
			ShortDescription: `[Manages a public IP.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the public IP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP associated with the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the public IP.`,
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
					Description: `The ID of the VM the public IP is associated with (if any). ## Import A public IP can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_public_ip.ImportedPublicIp eipalloc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the public IP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP associated with the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the public IP.`,
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
					Description: `The ID of the VM the public IP is associated with (if any). ## Import A public IP can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_public_ip.ImportedPublicIp eipalloc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_public_ip_link",
			Category:         "Resources",
			ShortDescription: `[Manages a public IP link.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_relink",
					Description: `(Optional) If true, allows the public IP to be associated with the VM or NIC that you specify even if it is already associated with another VM or NIC. If false, prevents the EIP from being associated with the VM or NIC that you specify if it is already associated with another VM or NIC. (By default, true in the public Cloud, false in a Net.)`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Optional) (Net only) The ID of the NIC. This parameter is required if the VM has more than one NIC attached. Otherwise, you need to specify the ` + "`" + `vm_id` + "`" + ` parameter instead. You cannot specify both parameters at the same time.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) (Net only) The primary or secondary private IP of the specified NIC. By default, the primary private IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `(Optional) The allocation ID of the public IP. This parameter is required unless you use the ` + "`" + `public_ip` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) The public IP. This parameter is required unless you use the ` + "`" + `public_ip_id` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Optional) The ID of the VM.<br />- In the public Cloud, this parameter is required.<br />- In a Net, this parameter is required if the VM has only one NIC. Otherwise, you need to specify the ` + "`" + `nic_id` + "`" + ` parameter instead. You cannot specify both parameters at the same time. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Net only) The ID representing the association of the public IP with the VM or the NIC. ## Import A public IP link can be imported using the public IP or the public IP link ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_public_ip_link.ImportedPublicIpLink eipassoc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Net only) The ID representing the association of the public IP with the VM or the NIC. ## Import A public IP link can be imported using the public IP or the public IP link ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_public_ip_link.ImportedPublicIpLink eipassoc-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_route",
			Category:         "Resources",
			ShortDescription: `[Manages a route.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "await_active_state",
					Description: `(Optional) By default or if set to true, waits for the route to be in the ` + "`" + `active` + "`" + ` state to declare its successful creation.<br />If false, the created route is in the ` + "`" + `active` + "`" + ` state if available, or the ` + "`" + `blackhole` + "`" + ` state if not available.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `(Required) The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `(Optional) The ID of an Internet service or virtual gateway attached to your Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `(Optional) The ID of a NAT service.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `(Optional) The ID of a Net peering.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Optional) The ID of a NIC.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required) The ID of the route table for which you want to create a route.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Optional) The ID of a NAT VM in your Net (attached to exactly one NIC). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "await_active_state",
					Description: `If true, the route is declared created when in the ` + "`" + `active` + "`" + ` state. If false, the route is created in the ` + "`" + `active` + "`" + ` state if available, or in the ` + "`" + `blackhole` + "`" + ` state if not available.`,
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
					Description: `If true, the route table is the main one.`,
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
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
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
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (always ` + "`" + `active` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A route can be imported using the route table ID and the destination IP range. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_routeImportedRoute rtb-12345678_10.0.0.0/0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "await_active_state",
					Description: `If true, the route is declared created when in the ` + "`" + `active` + "`" + ` state. If false, the route is created in the ` + "`" + `active` + "`" + ` state if available, or in the ` + "`" + `blackhole` + "`" + ` state if not available.`,
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
					Description: `If true, the route table is the main one.`,
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
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
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
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (always ` + "`" + `active` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A route can be imported using the route table ID and the destination IP range. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_routeImportedRoute rtb-12345678_10.0.0.0/0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_route_table",
			Category:         "Resources",
			ShortDescription: `[Manages a route table.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net for which you want to create a route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
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
					Description: `If true, the route table is the main one.`,
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
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
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
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (always ` + "`" + `active` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A route table can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_route_table.ImportedRouteTable rtb-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `If true, the route table is the main one.`,
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
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
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
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (always ` + "`" + `active` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A route table can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_route_table.ImportedRouteTable rtb-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_route_table_link",
			Category:         "Resources",
			ShortDescription: `[Manages a route table link.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required) The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the Subnet. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet. ## Import A route table link can be imported using the route table ID and the route table link ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_route_table_link.ImportedRouteTableLink rtb-12345678_rtbassoc-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet. ## Import A route table link can be imported using the route table ID and the route table link ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_route_table_link.ImportedRouteTableLink rtb-12345678_rtbassoc-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_security_group",
			Category:         "Resources",
			ShortDescription: `[Manages a security group.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the security group, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `(Optional) The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "remove_default_outbound_rule",
					Description: `(Optional) (Net only) By default or if set to false, the security group is created with a default outbound rule allowing all outbound flows. If set to true, the security group is created without a default outbound rule. For an existing security group, setting this parameter to true deletes the security group and creates a new one.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `(Required) The name of the security group.<br /> This name must not start with ` + "`" + `sg-` + "`" + `.</br> This name must be unique and contain between 1 and 255 ASCII characters. Accented letters are not allowed.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
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
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
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
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
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
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
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
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
				},
				resource.Attribute{
					Name:        "remove_default_outbound_rule",
					Description: `If false, the security group is created with a default outbound rule allowing all outbound flows. If true, the security group is created without a default outbound rule.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A security group can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_security_group.ImportedSecurityGroup sg-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
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
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
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
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
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
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
				},
				resource.Attribute{
					Name:        "remove_default_outbound_rule",
					Description: `If false, the security group is created with a default outbound rule allowing all outbound flows. If true, the security group is created without a default outbound rule.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A security group can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_security_group.ImportedSecurityGroup sg-87654321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_security_group_rule",
			Category:         "Resources",
			ShortDescription: `[Manages a security group rule.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "flow",
					Description: `(Required) The direction of the flow: ` + "`" + `Inbound` + "`" + ` or ` + "`" + `Outbound` + "`" + `. You can specify ` + "`" + `Outbound` + "`" + ` for Nets only.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `(Optional) The beginning of the port range for the TCP and UDP protocols, or an ICMP type number. If you specify this parameter, you cannot specify the ` + "`" + `rules` + "`" + ` parameter and its subparameters.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). If you specify this parameter, you cannot specify the ` + "`" + `rules` + "`" + ` parameter and its subparameters.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Optional) The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16). If you specify this parameter, you cannot specify the ` + "`" + `rules` + "`" + ` parameter and its subparameters.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional) Information about the security group rule to create. If you specify this parent parameter and its subparameters, you cannot specify the following parent parameters: ` + "`" + `from_port_range` + "`" + `, ` + "`" + `ip_protocol` + "`" + `, ` + "`" + `ip_range` + "`" + `, and ` + "`" + `to_port_range` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `(Optional) The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `(Optional) Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `(Optional) The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `(Optional) One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `(Optional) The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id_to_link",
					Description: `(Optional) The account ID of the owner of the security group for which you want to create a rule.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The ID of the security group for which you want to create a rule.`,
				},
				resource.Attribute{
					Name:        "security_group_name_to_link",
					Description: `(Optional) The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group.`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `(Optional) The end of the port range for the TCP and UDP protocols, or an ICMP code number. If you specify this parameter, you cannot specify the ` + "`" + `rules` + "`" + ` parameter and its subparameters. ## Attribute Reference The following attributes are exported:`,
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
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
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
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
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
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
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
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A security group rule can be imported using the following format: ` + "`" + `SecurityGroupId_Flow_IpProtocol_FromPortRange_ToPortRange_IpRange` + "`" + `. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_security_group_rule.ImportedRule sg-87654321_outbound_-1_-1_-1_10.0.0.0/16 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
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
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
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
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
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
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A security group rule can be imported using the following format: ` + "`" + `SecurityGroupId_Flow_IpProtocol_FromPortRange_ToPortRange_IpRange` + "`" + `. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_security_group_rule.ImportedRule sg-87654321_outbound_-1_-1_-1_10.0.0.0/16 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_server_certificate",
			Category:         "Resources",
			ShortDescription: `[Manages a server certificate.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `(Required) The PEM-encoded X509 certificate.`,
				},
				resource.Attribute{
					Name:        "chain",
					Description: `(Optional) The PEM-encoded intermediate certification authorities.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the certificate. Constraints: 1-128 alphanumeric characters, pluses (+), equals (=), commas (,), periods (.), at signs (@), minuses (-), or underscores (_).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to the server certificate, set to a slash (/) if not specified.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The PEM-encoded private key matching the certificate. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date at which the server certificate expires.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the server certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server certificate.`,
				},
				resource.Attribute{
					Name:        "orn",
					Description: `The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the server certificate.`,
				},
				resource.Attribute{
					Name:        "upload_date",
					Description: `The date at which the server certificate has been uploaded. ## Import A server certificate can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_server_certificate.ImportedServerCertificate 0123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date at which the server certificate expires.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the server certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server certificate.`,
				},
				resource.Attribute{
					Name:        "orn",
					Description: `The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the server certificate.`,
				},
				resource.Attribute{
					Name:        "upload_date",
					Description: `The date at which the server certificate has been uploaded. ## Import A server certificate can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_server_certificate.ImportedServerCertificate 0123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_snapshot",
			Category:         "Resources",
			ShortDescription: `[Manages a snapshot.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the snapshot.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `(Optional) (When importing) The pre-signed URL of the snapshot you want to import, or the normal URL of the snapshot if you have permission on the OOS bucket. For more information, see [Configuring a Pre-signed URL](https://docs.outscale.com/en/userguide/Configuring-a-Pre-signed-URL.html) or [Managing Access to Your Buckets and Objects](https://docs.outscale.com/en/userguide/Managing-Access-to-Your-Buckets-and-Objects.html).`,
				},
				resource.Attribute{
					Name:        "snapshot_size",
					Description: `(Optional) (When importing) The size of the snapshot you want to create in your account, in bytes. This size must be greater than or equal to the size of the original, uncompressed snapshot.`,
				},
				resource.Attribute{
					Name:        "source_region_name",
					Description: `(Optional) (When copying) The name of the source Region, which must be the same as the Region of your account.`,
				},
				resource.Attribute{
					Name:        "source_snapshot_id",
					Description: `(Optional) (When copying) The ID of the snapshot you want to copy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Optional) (When creating) The ID of the volume you want to create a snapshot of. ## Attribute Reference The following attributes are exported:`,
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
					Name:        "creation_date",
					Description: `The date and time of creation of the snapshot.`,
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
					Description: `If true, the resource is public. If false, the resource is private.`,
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
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB). ## Import A snapshot can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import terraform import outscale_snapshot.ImportedSnapshot snap-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the snapshot.`,
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
					Description: `If true, the resource is public. If false, the resource is private.`,
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
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB). ## Import A snapshot can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import terraform import outscale_snapshot.ImportedSnapshot snap-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_snapshot_attributes",
			Category:         "Resources",
			ShortDescription: `[Manages snapshot attributes.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "permissions_to_create_volume_additions",
					Description: `(Optional) Information about the users to whom you want to give permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account ID of one or more users to whom you want to give permissions.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `(Optional) If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume_removals",
					Description: `(Optional) Information about the users from whom you want to remove permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account ID of one or more users from whom you want to remove permissions.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `(Optional) If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Required) The ID of the snapshot. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_snapshot_export_task",
			Category:         "Resources",
			ShortDescription: `[Manages a snapshot export task.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the OOS export task to create.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `(Optional) The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_api_key",
					Description: `Information about the OOS API key.`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `(Optional) The API key of the OOS account that enables you to access the bucket.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The secret key of the OOS account that enables you to access the bucket.`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `(Optional) The name of the OOS bucket where you want to export the object.`,
				},
				resource.Attribute{
					Name:        "osu_manifest_url",
					Description: `(Optional) The URL of the manifest file.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `(Optional) The prefix for the key of the OOS object.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Required) The ID of the snapshot to export.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `If the snapshot export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the snapshot export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the snapshot is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the snapshot.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot to be exported.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot export task (` + "`" + `pending` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot export task.`,
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
					Name:        "task_id",
					Description: `The ID of the snapshot export task.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "comment",
					Description: `If the snapshot export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the snapshot export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the snapshot is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the snapshot.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot to be exported.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot export task (` + "`" + `pending` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot export task.`,
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
					Name:        "task_id",
					Description: `The ID of the snapshot export task.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_subnet",
			Category:         "Resources",
			ShortDescription: `[Manages a Subnet.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Required) The IP range in the Subnet, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).<br /> The IP range of the Subnet can be either the same as the Net one if you create only a single Subnet in this Net, or a subset of the Net one. In case of several Subnets in a Net, their IP ranges must not overlap. The smallest Subnet you can create uses a /29 netmask (eight IPs). For more information, see [About VPCs](https://docs.outscale.com/en/userguide/About-VPCs.html).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net for which you want to create a Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `(Optional) The name of the Subregion in which you want to create the Subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IPs in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If true, a public IP is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A subnet can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_subnet.ImportedSubnet subnet-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IPs in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If true, a public IP is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
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
					Description: `The value of the tag, between 0 and 255 characters. ## Import A subnet can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_subnet.ImportedSubnet subnet-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_virtual_gateway",
			Category:         "Resources",
			ShortDescription: `[Manages a virtual gateway.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Required) The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters. ## Attribute Reference The following attributes are exported:`,
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
					Description: `The ID of the virtual gateway. ## Import A virtual gateway can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_virtual_gateway.ImportedVirtualGateway vgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The ID of the virtual gateway. ## Import A virtual gateway can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_virtual_gateway.ImportedVirtualGateway vgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_virtual_gateway_link",
			Category:         "Resources",
			ShortDescription: `[Manages a virtual gateway link.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required) The ID of the Net to which you want to attach the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `(Required) The ID of the virtual gateway. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `). ## Import A virtual gateway link can be imported using its virtual gateway ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_virtual_gateway_link.ImportedVirtualGatewayLink vgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `). ## Import A virtual gateway link can be imported using its virtual gateway ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_virtual_gateway_link.ImportedVirtualGatewayLink vgw-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_virtual_gateway_route_propagation",
			Category:         "Resources",
			ShortDescription: `[Manages a virtual gateway route propagation.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable",
					Description: `(Required) If true, a virtual gateway can propagate routes to a specified route table of a Net. If false, the propagation is disabled.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required) The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `(Required) The ID of the virtual gateway. ## Attribute Reference The following attributes are exported:`,
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
					Description: `If true, the route table is the main one.`,
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
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
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
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (always ` + "`" + `active` + "`" + `).`,
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
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
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
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
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
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (always ` + "`" + `active` + "`" + `).`,
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
			Type:             "outscale_vm",
			Category:         "Resources",
			ShortDescription: `[Manages a virtual machine (VM).]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `(Optional) One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + ` with a maximum performance ratio of 300 IOPS per gibibyte.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Optional) The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [About Volumes > Volume Types and IOPS](https://docs.outscale.com/en/userguide/About-Volumes.html#_volume_types_and_iops).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) The device name for the volume. For a root device, you must use ` + "`" + `/dev/sda1` + "`" + `. For other volumes, you must use ` + "`" + `/dev/sdX` + "`" + `, ` + "`" + `/dev/sdXX` + "`" + `, ` + "`" + `/dev/xvdX` + "`" + `, or ` + "`" + `/dev/xvdXX` + "`" + ` (where the first ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `b` + "`" + ` and ` + "`" + `z` + "`" + `, and the second ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `a` + "`" + ` and ` + "`" + `z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "no_device",
					Description: `(Optional) Removes the device which is included in the block device mapping of the OMI.`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `(Optional) The name of the virtual device (` + "`" + `ephemeralN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `(Optional) A unique identifier which enables you to manage the idempotency.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `(Optional) If true, you cannot delete the VM unless you change this parameter back to false.`,
				},
				resource.Attribute{
					Name:        "get_admin_password",
					Description: `(Optional) (Windows VM only) If true, waits for the administrator password of the VM to become available in order to retrieve the VM. The password is exported to the ` + "`" + `admin_password` + "`" + ` attribute.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) The ID of the OMI used to create the VM. You can find the list of OMIs by calling the [ReadImages](https://docs.outscale.com/api#readimages) method.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `(Optional) The name of the keypair.`,
				},
				resource.Attribute{
					Name:        "nested_virtualization",
					Description: `(Optional) (dedicated tenancy only) If true, nested virtualization is enabled. If false, it is disabled.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Optional) One or more NICs. If you specify this parameter, you must not specify the ` + "`" + `subnet_id` + "`" + ` and ` + "`" + `subregion_name` + "`" + ` parameters. You also must define one NIC as the primary network interface of the VM with ` + "`" + `0` + "`" + ` as its device number.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) If true, the NIC is deleted when the VM is terminated. You can specify this parameter only for a new NIC. To modify this value for an existing NIC, see [UpdateNic](https://docs.outscale.com/api#updatenic).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the NIC, if you are creating a NIC when creating the VM.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `(Optional) The index of the VM device for the NIC attachment (between ` + "`" + `0` + "`" + ` and ` + "`" + `7` + "`" + `, both included). This parameter is required if you create a NIC when creating the VM.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Optional) The ID of the NIC, if you are attaching an existing NIC when creating a VM.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) One or more private IPs to assign to the NIC, if you create a NIC when creating a VM. Only one private IP can be the primary private IP.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `(Optional) If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) The private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "secondary_private_ip_count",
					Description: `(Optional) The number of secondary private IPs, if you create a NIC when creating a VM. This parameter cannot be specified if you specified more than one private IP in the ` + "`" + `private_ips` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) One or more IDs of security groups for the NIC, if you create a NIC when creating a VM.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of the Subnet for the NIC, if you create a NIC when creating a VM. This parameter is required if you create a NIC when creating the VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `(Optional) The performance of the VM (` + "`" + `medium` + "`" + ` | ` + "`" + `high` + "`" + ` | ` + "`" + `highest` + "`" + `). Updating this parameter will trigger a stop/start of the VM.`,
				},
				resource.Attribute{
					Name:        "placement_subregion_name",
					Description: `(Optional) The name of the Subregion where the VM is placed.`,
				},
				resource.Attribute{
					Name:        "placement_tenancy",
					Description: `(Optional) The tenancy of the VM (` + "`" + `default` + "`" + ` | ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) One or more private IPs of the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) One or more IDs of security group for the VMs.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `(Optional) One or more names of security groups for the VMs.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `running` + "`" + ` | ` + "`" + `stopped` + "`" + `). If set to ` + "`" + `stopped` + "`" + `, the VM is stopped regardless of the value of the ` + "`" + `vm_initiated_shutdown_behavior` + "`" + ` argument.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of the Subnet in which you want to create the VM. If you specify this parameter, you must not specify the ` + "`" + `nics` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Data or script used to add a specific configuration to the VM. It must be Base64-encoded, either directly or using the [base64encode](https://www.terraform.io/docs/configuration/functions/base64encode.html) Terraform function. For multiline strings, use a [heredoc syntax](https://www.terraform.io/docs/configuration/expressions.html#string-literals). Updating this parameter will trigger a stop/start of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `(Optional) The VM behavior when you stop it. By default or if set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is terminated.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `(Optional) The type of VM (` + "`" + `t2.small` + "`" + ` by default). Updating this parameter will trigger a stop/start of the VM.<br /> For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Windows VM only) The administrator password of the VM. This password is encrypted with the keypair you specified when launching the VM and encoded in Base64. You need to wait about 10 minutes after launching the VM to be able to retrieve this password.<br />If ` + "`" + `get_admin_password` + "`" + ` is false or not specified, the VM resource is created without the ` + "`" + `admin_password` + "`" + ` attribute. Once ` + "`" + `admin_password` + "`" + ` is available, it will appear in the Terraform state after the next`,
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
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
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
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If true, you cannot delete the VM unless you change this parameter back to false.`,
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
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, and so on).`,
				},
				resource.Attribute{
					Name:        "nested_virtualization",
					Description: `If true, nested virtualization is enabled. If false, it is disabled.`,
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
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
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
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
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
					Description: `The private IP or IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP.`,
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
					Name:        "placement_subregion_name",
					Description: `The name of the Subregion where the VM is placed.`,
				},
				resource.Attribute{
					Name:        "placement_tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` | ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP of the VM.`,
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
					Description: `The public IP of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, ` + "`" + `/dev/vda1` + "`" + `).`,
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
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
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
					Description: `The VM behavior when you stop it. If set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html). ## Import A VM can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_vm.ImportedVm i-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Windows VM only) The administrator password of the VM. This password is encrypted with the keypair you specified when launching the VM and encoded in Base64. You need to wait about 10 minutes after launching the VM to be able to retrieve this password.<br />If ` + "`" + `get_admin_password` + "`" + ` is false or not specified, the VM resource is created without the ` + "`" + `admin_password` + "`" + ` attribute. Once ` + "`" + `admin_password` + "`" + ` is available, it will appear in the Terraform state after the next`,
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
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
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
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If true, you cannot delete the VM unless you change this parameter back to false.`,
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
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, and so on).`,
				},
				resource.Attribute{
					Name:        "nested_virtualization",
					Description: `If true, nested virtualization is enabled. If false, it is disabled.`,
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
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
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
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
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
					Description: `The private IP or IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP.`,
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
					Name:        "placement_subregion_name",
					Description: `The name of the Subregion where the VM is placed.`,
				},
				resource.Attribute{
					Name:        "placement_tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` | ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP of the VM.`,
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
					Description: `The public IP of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, ` + "`" + `/dev/vda1` + "`" + `).`,
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
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
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
					Description: `The VM behavior when you stop it. If set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html). ## Import A VM can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_vm.ImportedVm i-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_volume",
			Category:         "Resources",
			ShortDescription: `[Manages a volume.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + ` with a maximum performance ratio of 300 IOPS per gibibyte.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the volume, in gibibytes (GiB). The maximum allowed size for a volume is 14901 GiB. This parameter is required if the volume is not created from a snapshot (` + "`" + `snapshot_id` + "`" + ` unspecified).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of the snapshot from which you want to create the volume.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `(Required) The Subregion in which you want to create the volume.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of volume you want to create (` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `standard` + "`" + `). If not specified, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [About Volumes > Volume Types and IOPS](https://docs.outscale.com/en/userguide/About-Volumes.html#_volume_types_and_iops). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the volume.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS):<br />- For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS.<br />- For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
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
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `). ## Import A volume can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_volume.ImportedVolume vol-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the volume.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS):<br />- For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS.<br />- For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
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
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `). ## Import A volume can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_volume.ImportedVolume vol-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_volumes_link",
			Category:         "Resources",
			ShortDescription: `[Manages a volume link.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_name",
					Description: `(Required) The name of the device. For a root device, you must use ` + "`" + `/dev/sda1` + "`" + `. For other volumes, you must use ` + "`" + `/dev/sdX` + "`" + `, ` + "`" + `/dev/sdXX` + "`" + `, ` + "`" + `/dev/xvdX` + "`" + `, or ` + "`" + `/dev/xvdXX` + "`" + ` (where the first ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `b` + "`" + ` and ` + "`" + `z` + "`" + `, and the second ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `a` + "`" + ` and ` + "`" + `z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `(Required) The ID of the VM you want to attach the volume to.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) The ID of the volume you want to attach. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` | ` + "`" + `detaching` + "`" + ` | ` + "`" + `attached` + "`" + ` | ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume. ## Import A volume link can be imported using a volume ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_volumes_link.ImportedVolumeLink vol-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` | ` + "`" + `detaching` + "`" + ` | ` + "`" + `attached` + "`" + ` | ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume. ## Import A volume link can be imported using a volume ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_volumes_link.ImportedVolumeLink vol-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_vpn_connection",
			Category:         "Resources",
			ShortDescription: `[Manages a VPN connection.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `(Required) The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Required) The type of VPN connection (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `(Optional) If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A tag to add to this resource. You can specify this argument several times.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `(Required) The ID of the virtual gateway. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `Example configuration for the client gateway.`,
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
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
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
					Description: `If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
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
					Name:        "vgw_telemetries",
					Description: `Information about the current state of one or more of the VPN tunnels.`,
				},
				resource.Attribute{
					Name:        "accepted_route_count",
					Description: `The number of routes accepted through BGP (Border Gateway Protocol) route exchanges.`,
				},
				resource.Attribute{
					Name:        "last_state_change_date",
					Description: `The date and time (UTC) of the latest state update.`,
				},
				resource.Attribute{
					Name:        "outside_ip_address",
					Description: `The IP on the OUTSCALE side of the tunnel.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the IPSEC tunnel (` + "`" + `UP` + "`" + ` \| ` + "`" + `DOWN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_description",
					Description: `A description of the current state of the tunnel.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection. ## Import A VPN connection can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_vpn_connection.ImportedVPN vpn-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `Example configuration for the client gateway.`,
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
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
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
					Description: `If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
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
					Name:        "vgw_telemetries",
					Description: `Information about the current state of one or more of the VPN tunnels.`,
				},
				resource.Attribute{
					Name:        "accepted_route_count",
					Description: `The number of routes accepted through BGP (Border Gateway Protocol) route exchanges.`,
				},
				resource.Attribute{
					Name:        "last_state_change_date",
					Description: `The date and time (UTC) of the latest state update.`,
				},
				resource.Attribute{
					Name:        "outside_ip_address",
					Description: `The IP on the OUTSCALE side of the tunnel.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the IPSEC tunnel (` + "`" + `UP` + "`" + ` \| ` + "`" + `DOWN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_description",
					Description: `A description of the current state of the tunnel.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection. ## Import A VPN connection can be imported using its ID. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_vpn_connection.ImportedVPN vpn-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_vpn_connection_route",
			Category:         "Resources",
			ShortDescription: `[Manages a VPN connection route.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `(Required) The network prefix of the route, in CIDR notation (for example, ` + "`" + `10.12.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `(Required) The ID of the target VPN connection of the static route. ## Attribute Reference No attribute is exported. ## Import A VPN connection route can be imported using the VPN connection ID and the route destination IP range. For example: ` + "`" + `` + "`" + `` + "`" + `console $ terraform import outscale_vpn_connection_route.ImportedRoute vpn-12345678_10.0.0.0/0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"outscale_access_key":                        0,
		"outscale_api_access_policy":                 1,
		"outscale_api_access_rule":                   2,
		"outscale_ca":                                3,
		"outscale_client_gateway":                    4,
		"outscale_dhcp_option":                       5,
		"outscale_flexible_gpu":                      6,
		"outscale_flexible_gpu_link":                 7,
		"outscale_image":                             8,
		"outscale_image_export_task":                 9,
		"outscale_image_launch_permission":           10,
		"outscale_internet_service":                  11,
		"outscale_internet_service_link":             12,
		"outscale_keypair":                           13,
		"outscale_load_balancer":                     14,
		"outscale_load_balancer_attributes":          15,
		"outscale_load_balancer_listener_rule":       16,
		"outscale_load_balancer_policy":              17,
		"outscale_load_balancer_vms":                 18,
		"outscale_nat_service":                       19,
		"outscale_net":                               20,
		"outscale_net_access_point":                  21,
		"outscale_net_attributes":                    22,
		"outscale_net_peering":                       23,
		"outscale_net_peering_acceptation":           24,
		"outscale_nic":                               25,
		"outscale_nic_link":                          26,
		"outscale_nic_private_ip":                    27,
		"outscale_public_ip":                         28,
		"outscale_public_ip_link":                    29,
		"outscale_route":                             30,
		"outscale_route_table":                       31,
		"outscale_route_table_link":                  32,
		"outscale_security_group":                    33,
		"outscale_security_group_rule":               34,
		"outscale_server_certificate":                35,
		"outscale_snapshot":                          36,
		"outscale_snapshot_attributes":               37,
		"outscale_snapshot_export_task":              38,
		"outscale_subnet":                            39,
		"outscale_virtual_gateway":                   40,
		"outscale_virtual_gateway_link":              41,
		"outscale_virtual_gateway_route_propagation": 42,
		"outscale_vm":                                43,
		"outscale_volume":                            44,
		"outscale_volumes_link":                      45,
		"outscale_vpn_connection":                    46,
		"outscale_vpn_connection_route":              47,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
