package volterra

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "volterra_volterra_address_allocator",
			Category:         "Data Sources",
			ShortDescription: `"Data source for volterra_address_allocator resource"`,
			Description: `

volterra_address_allocator is used to get mode of the allocator object and allocation map. Mode determines if the allocation is limited to VERs within the local site / cluster (local) or if the allocation is across different sites (global).


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the tunnel`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace where the tunnel is configure. Example ` + "`" + `system` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode of the address allocator. Expected values are ` + "`" + `LOCAL` + "`" + ` ` + "`" + `GLOBAL_PER_SITE_NODE` + "`" + ` (` + "`" + `string` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "allocation_map",
					Description: `Map of per site node allocation (` + "`" + `map(string)` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mode",
					Description: `Mode of the address allocator. Expected values are ` + "`" + `LOCAL` + "`" + ` ` + "`" + `GLOBAL_PER_SITE_NODE` + "`" + ` (` + "`" + `string` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "allocation_map",
					Description: `Map of per site node allocation (` + "`" + `map(string)` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "volterra_volterra_http_loadbalancer_state",
			Category:         "Data Sources",
			ShortDescription: `"Data source for volterra_http_loadbalancer_state resource"`,
			Description: `

volterra_http_loadbalancer_state is used to get the state information related to a http loadbalancer. This includes cname, ip_address and auto_cert info related state


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "volterra_volterra_namespace",
			Category:         "Data Sources",
			ShortDescription: `"Data source for namespace resource"`,
			Description: `

Namespace creates logical independent workspace within a tenant. Data Source reads the namespace object and gets values like tenant_name, id of the namespace object.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the namespace to be queried ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the namespace returned`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `Tenant name configured in volterra`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the namespace returned`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `Tenant name configured in volterra`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "volterra_volterra_parse_aws_cgw_configuration",
			Category:         "Data Sources",
			ShortDescription: `"Data source for volterra_parse_aws_cgw_configuration resource"`,
			Description: `

Parses the aws gw configuration provided and gets all the info in a unsorted manner.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "customer_gateway_configuration",
					Description: `(Required) Customer gateway configuration received (in XML format) after creating vpn connection object on aws ## Attribute Reference ` + "`" + `tunnel1_address` + "`" + ` - Public IP address of the first VPN tunnel. ` + "`" + `tunnel1_cgw_inside_address` + "`" + ` - Customer Gateway Side inside address of the first VPN tunnel. ` + "`" + `tunnel1_vgw_inside_address` + "`" + ` - AWS VGW side inside address of the first VPN tunnel. ` + "`" + `tunnel1_preshared_key` + "`" + ` - Preshared key of the first VPN tunnel. ` + "`" + `tunnel1_bgp_asn` + "`" + ` - BGP asn number of the first VPN tunnel. ` + "`" + `tunnel1_bgp_holdtime` + "`" + ` - The bgp holdtime of the first VPN tunnel. ` + "`" + `tunnel2_address` + "`" + ` - Public IP address of the second VPN tunnel. ` + "`" + `tunnel2_cgw_inside_address` + "`" + ` - Customer Gateway Side inside address of the second VPN tunnel. ` + "`" + `tunnel2_vgw_inside_address` + "`" + ` - AWS VGW side inside address of the second VPN tunnel. ` + "`" + `tunnel2_preshared_key` + "`" + ` - Preshared key of the second VPN tunnel. ` + "`" + `tunnel2_bgp_asn` + "`" + ` - BGP asn number of the second VPN tunnel. ` + "`" + `tunnel2_bgp_holdtime` + "`" + ` - BGP holdtime of the second VPN tunnel.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "volterra_volterra_tunnel",
			Category:         "Data Sources",
			ShortDescription: `"Data source for volterra_tunnel resource"`,
			Description: `

volterra_tunnel is a way to get tunnel information


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the tunnel`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace where the tunnel is configure. Example ` + "`" + `system` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the tunnel created`,
				},
				resource.Attribute{
					Name:        "ipsec_psk_b64",
					Description: `If tunnel type is IPSEC_PSK, it will return the ipsec pre shared key on base64 format`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `Type of the tunnel created`,
				},
				resource.Attribute{
					Name:        "ipsec_psk_b64",
					Description: `If tunnel type is IPSEC_PSK, it will return the ipsec pre shared key on base64 format`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "volterra_volterra_virtual_host_dns_info",
			Category:         "Data Sources",
			ShortDescription: `"Data source for virtual_host_dns_info resource"`,
			Description: `

volterra_virtual_host_dns_info is a way to get DNS information like dns cname or IP address for a given virtual host


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the virtual host whose dns info must be queried ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `Host name to be used for the virtual host`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `Tenant name configured in volterra`,
				},
				resource.Attribute{
					Name:        "dns_info",
					Description: `DNS information for this virtual host`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address associated with the virtual host`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "host_name",
					Description: `Host name to be used for the virtual host`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `Tenant name configured in volterra`,
				},
				resource.Attribute{
					Name:        "dns_info",
					Description: `DNS information for this virtual host`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address associated with the virtual host`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"volterra_volterra_address_allocator":           0,
		"volterra_volterra_http_loadbalancer_state":     1,
		"volterra_volterra_namespace":                   2,
		"volterra_volterra_parse_aws_cgw_configuration": 3,
		"volterra_volterra_tunnel":                      4,
		"volterra_volterra_virtual_host_dns_info":       5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
