package bigip

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_datagroup",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_datagroup data source`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"datagroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the datagroup`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Required) partition of the datagroup ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Data Group type (string, ip, integer)"`,
				},
				resource.Attribute{
					Name:        "record",
					Description: `Specifies record of type (string/ip/integer)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The Data Group type (string, ip, integer)"`,
				},
				resource.Attribute{
					Name:        "record",
					Description: `Specifies record of type (string/ip/integer)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_irule",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_irule data source`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"irule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the irule`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Required) partition of the ltm irule ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "irule",
					Description: `Irule configured on bigip`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of irule configured on bigip with full path`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `Bigip partition in which rule is configured`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "irule",
					Description: `Irule configured on bigip`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of irule configured on bigip with full path`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `Bigip partition in which rule is configured`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_monitor",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_monitor data source`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"monitor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the ltm monitor`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Required) partition of the ltm monitor ## Attributes Reference Additionally, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `id will be full path name of ltm monitor.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown.`,
				},
				resource.Attribute{
					Name:        "ip_dscp",
					Description: `Displays the differentiated services code point (DSCP). DSCP is a 6-bit value in the Differentiated Services (DS) field of the IP header.`,
				},
				resource.Attribute{
					Name:        "manual_resume",
					Description: `Displays whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `Instructs the system to mark the target resource down when the test is successful.`,
				},
				resource.Attribute{
					Name:        "transparent",
					Description: `Displays whether the monitor operates in transparent mode.`,
				},
				resource.Attribute{
					Name:        "adaptive_limit",
					Description: `Displays whether adaptive response time monitoring is enabled for this monitor.`,
				},
				resource.Attribute{
					Name:        "adaptive",
					Description: `Displays whether adaptive response time monitoring is enabled for this monitor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "destination",
					Description: `id will be full path name of ltm monitor.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown.`,
				},
				resource.Attribute{
					Name:        "ip_dscp",
					Description: `Displays the differentiated services code point (DSCP). DSCP is a 6-bit value in the Differentiated Services (DS) field of the IP header.`,
				},
				resource.Attribute{
					Name:        "manual_resume",
					Description: `Displays whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `Instructs the system to mark the target resource down when the test is successful.`,
				},
				resource.Attribute{
					Name:        "transparent",
					Description: `Displays whether the monitor operates in transparent mode.`,
				},
				resource.Attribute{
					Name:        "adaptive_limit",
					Description: `Displays whether adaptive response time monitoring is enabled for this monitor.`,
				},
				resource.Attribute{
					Name:        "adaptive",
					Description: `Displays whether adaptive response time monitoring is enabled for this monitor.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_node",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_node data source`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"node",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the node.`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Required) partition of the node. ## Attributes Reference Additionally, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the node.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `User defined description of the node.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `Node connection limit.`,
				},
				resource.Attribute{
					Name:        "dynamic_ratio",
					Description: `The dynamic ratio number for the node.`,
				},
				resource.Attribute{
					Name:        "full_path",
					Description: `Full path of the node (partition and name)`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `Specifies the health monitors the system currently uses to monitor this node.`,
				},
				resource.Attribute{
					Name:        "rate_limit",
					Description: `Node rate limit.`,
				},
				resource.Attribute{
					Name:        "ratio",
					Description: `Node ratio weight.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the node. The ` + "`" + `fqdn` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `The FQDN node's address family.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The amount of time before sending the next DNS query.`,
				},
				resource.Attribute{
					Name:        "downinterval",
					Description: `The number of attempts to resolve a domain name.`,
				},
				resource.Attribute{
					Name:        "autopopulate",
					Description: `Specifies if the node should scale to the IP address set returned by DNS.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `The address of the node.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `User defined description of the node.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `Node connection limit.`,
				},
				resource.Attribute{
					Name:        "dynamic_ratio",
					Description: `The dynamic ratio number for the node.`,
				},
				resource.Attribute{
					Name:        "full_path",
					Description: `Full path of the node (partition and name)`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `Specifies the health monitors the system currently uses to monitor this node.`,
				},
				resource.Attribute{
					Name:        "rate_limit",
					Description: `Node rate limit.`,
				},
				resource.Attribute{
					Name:        "ratio",
					Description: `Node ratio weight.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the node. The ` + "`" + `fqdn` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `The FQDN node's address family.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The amount of time before sending the next DNS query.`,
				},
				resource.Attribute{
					Name:        "downinterval",
					Description: `The number of attempts to resolve a domain name.`,
				},
				resource.Attribute{
					Name:        "autopopulate",
					Description: `Specifies if the node should scale to the IP address set returned by DNS.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_policy",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_policy data source`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the policy which includes partion ( /partition/policy-name ) ## Attributes Reference Additionally, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `Specifies the match strategy.`,
				},
				resource.Attribute{
					Name:        "requires",
					Description: `Specifies the protocol.`,
				},
				resource.Attribute{
					Name:        "controls",
					Description: `Specifies the controls.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `Rules defined in the policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `Specifies the match strategy.`,
				},
				resource.Attribute{
					Name:        "requires",
					Description: `Specifies the protocol.`,
				},
				resource.Attribute{
					Name:        "controls",
					Description: `Specifies the controls.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `Rules defined in the policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ltm_pool",
			Category:         "Local Traffic Manager(LTM)",
			ShortDescription: `Provides details about bigip_ltm_pool data source`,
			Description:      ``,
			Keywords: []string{
				"local",
				"traffic",
				"manager",
				"ltm",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the ltm monitor`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Required) partition of the ltm monitor ## Attributes Reference Additionally, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "full_path",
					Description: `Full path to the pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "full_path",
					Description: `Full path to the pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_ssl_certificate",
			Category:         "System",
			ShortDescription: `Provides details about bigip_ssl_certificate data source`,
			Description:      ``,
			Keywords: []string{
				"system",
				"ssl",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the ssl_certificate`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Required) partition of the ltm ssl_certificate ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of ssl_certificate configured on bigip with full path`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `Bigip partition in which ssl-certificate is configured`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of ssl_certificate configured on bigip with full path`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `Bigip partition in which ssl-certificate is configured`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_vwan_config",
			Category:         "vWAN",
			ShortDescription: `Provides details about bigip_vwan_config data source`,
			Description:      ``,
			Keywords: []string{
				"vwan",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "azure_vwan_resourcegroup",
					Description: `(Required, type ` + "`" + `string` + "`" + `) Name of the Azure vWAN resource group`,
				},
				resource.Attribute{
					Name:        "azure_vwan_name",
					Description: `(Required,type ` + "`" + `string` + "`" + `) Name of the Azure vWAN Name`,
				},
				resource.Attribute{
					Name:        "azure_vwan_vpnsite",
					Description: `(Required,type ` + "`" + `string` + "`" + `) Name of the Azure vWAN VPN site from which configuration to be download ## Pre-required Environment Settings:`,
				},
				resource.Attribute{
					Name:        "AZURE_CLIENT_ID",
					Description: `(Required) Set this environment variable with the Azure app client ID to use.`,
				},
				resource.Attribute{
					Name:        "AZURE_CLIENT_SECRET",
					Description: `(Required) Set this environment variable with the Azure app secret to use.`,
				},
				resource.Attribute{
					Name:        "AZURE_SUBSCRIPTION_ID",
					Description: `(Required) Set this environment variable with the Azure subscription ID to use.`,
				},
				resource.Attribute{
					Name:        "AZURE_TENANT_ID",
					Description: `(Required) Set this environment variable with the Tenant ID to which to authenticate.`,
				},
				resource.Attribute{
					Name:        "STORAGE_ACCOUNT_NAME",
					Description: `(Required) Set this environment variable with the storage account for download config.`,
				},
				resource.Attribute{
					Name:        "STORAGE_ACCOUNT_KEY",
					Description: `(Required) Specifies the storage account key to authenticate,set this Environment variable with account key value. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "bigip_gw_ip",
					Description: `(type ` + "`" + `string` + "`" + `) provides IP address of BIGIP G/W for IPSec Endpoint.`,
				},
				resource.Attribute{
					Name:        "preshared_key",
					Description: `(type ` + "`" + `string` + "`" + `) provides pre-shared-key used for IPSec Tunnel creation.`,
				},
				resource.Attribute{
					Name:        "hub_address_space",
					Description: `(type ` + "`" + `string` + "`" + `) Provides IP Address space used on vWAN Hub.`,
				},
				resource.Attribute{
					Name:        "hub_connected_subnets",
					Description: `(type ` + "`" + `list` + "`" + `) Provides Subnets connected to vWAN Hub.`,
				},
				resource.Attribute{
					Name:        "vwan_gw_address",
					Description: `(type ` + "`" + `list` + "`" + `) Provides vWAN Gateway Address for IPSec End point`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bigip_gw_ip",
					Description: `(type ` + "`" + `string` + "`" + `) provides IP address of BIGIP G/W for IPSec Endpoint.`,
				},
				resource.Attribute{
					Name:        "preshared_key",
					Description: `(type ` + "`" + `string` + "`" + `) provides pre-shared-key used for IPSec Tunnel creation.`,
				},
				resource.Attribute{
					Name:        "hub_address_space",
					Description: `(type ` + "`" + `string` + "`" + `) Provides IP Address space used on vWAN Hub.`,
				},
				resource.Attribute{
					Name:        "hub_connected_subnets",
					Description: `(type ` + "`" + `list` + "`" + `) Provides Subnets connected to vWAN Hub.`,
				},
				resource.Attribute{
					Name:        "vwan_gw_address",
					Description: `(type ` + "`" + `list` + "`" + `) Provides vWAN Gateway Address for IPSec End point`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"bigip_bigip_ltm_datagroup":   0,
		"bigip_bigip_ltm_irule":       1,
		"bigip_bigip_ltm_monitor":     2,
		"bigip_bigip_ltm_node":        3,
		"bigip_bigip_ltm_policy":      4,
		"bigip_bigip_ltm_pool":        5,
		"bigip_bigip_ssl_certificate": 6,
		"bigip_bigip_vwan_config":     7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
