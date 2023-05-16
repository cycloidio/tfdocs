package bigip

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_fast_aws_service_discovery",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details for AWS Service discovery config`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"fast",
				"aws",
				"service",
				"discovery",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tag_key",
					Description: `(` + "`" + `Required` + "`" + `,type ` + "`" + `string` + "`" + `) The tag key associated with the node to add to this pool.`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `(` + "`" + `Required` + "`" + `,type ` + "`" + `string` + "`" + `) The tag value associated with the node to add to this pool.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `int` + "`" + `)Port to be used for AWS service discovery,default ` + "`" + `80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_realm",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Specifies whether to look for public or private IP addresses,default ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "undetectable_action",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Action to take when node cannot be detected,default ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "credential_update",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `bool` + "`" + `) Specifies whether you are updating your credentials,default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `) AWS region in which ADC is running,default Empty string.`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the ` + "`" + `aws_secret_access_key` + "`" + ` field)`,
				},
				resource.Attribute{
					Name:        "aws_secret_access_key",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the ` + "`" + `aws_secret_access_key` + "`" + ` field)`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)AWS externalID field.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `) Assume a role (also requires the ` + "`" + `external_id` + "`" + ` field)`,
				},
				resource.Attribute{
					Name:        "minimum_monitors",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Member is down when fewer than minimum monitors report it healthy.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Update interval for service discovery. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "aws_sd_json",
					Description: `The JSON for AWS service discovery block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "aws_sd_json",
					Description: `The JSON for AWS service discovery block.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_fast_azure_service_discovery",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details for Azure Service discovery config`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"fast",
				"azure",
				"service",
				"discovery",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_group",
					Description: `(` + "`" + `Required` + "`" + `,type ` + "`" + `string` + "`" + `) Azure Resource Group name.`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(` + "`" + `Required` + "`" + `,type ` + "`" + `string` + "`" + `) Azure subscription ID.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `int` + "`" + `)Port to be used for Azure service discovery,default ` + "`" + `80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(` + "`" + `Required` + "`" + `,type ` + "`" + `string` + "`" + `) The tag key associated with the node to add to this pool.`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `(` + "`" + `Required` + "`" + `,type ` + "`" + `string` + "`" + `) The tag value associated with the node to add to this pool.`,
				},
				resource.Attribute{
					Name:        "address_realm",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Specifies whether to look for public or private IP addresses,default ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "undetectable_action",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Action to take when node cannot be detected,default ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "credential_update",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `bool` + "`" + `) Specifies whether you are updating your credentials,default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minimum_monitors",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Member is down when fewer than minimum monitors report it healthy.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Update interval for service discovery. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "azure_sd_json",
					Description: `The JSON for Azure service discovery block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "azure_sd_json",
					Description: `The JSON for Azure service discovery block.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_fast_consul_service_discovery",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details for Hashicorp Consul Service discovery config`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"fast",
				"consul",
				"service",
				"discovery",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uri",
					Description: `(` + "`" + `Required` + "`" + `,type ` + "`" + `string` + "`" + `) The location of the node data.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `int` + "`" + `)Port to be used for AWS service discovery,default ` + "`" + `80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_realm",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Specifies whether to look for public or private IP addresses,default ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "undetectable_action",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Action to take when node cannot be detected,default ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "credential_update",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `bool` + "`" + `) Specifies whether you are updating your credentials,default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "encoded_token",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `) Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.`,
				},
				resource.Attribute{
					Name:        "jmes_path_query",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Custom JMESPath Query.`,
				},
				resource.Attribute{
					Name:        "reject_unauthorized",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `bool` + "`" + `)If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.`,
				},
				resource.Attribute{
					Name:        "trust_ca",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)CA Bundle to validate server certificates.`,
				},
				resource.Attribute{
					Name:        "minimum_monitors",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Member is down when fewer than minimum monitors report it healthy.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Update interval for service discovery. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "consul_sd_json",
					Description: `The JSON for Hashicorp Consul service discovery block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "consul_sd_json",
					Description: `The JSON for Hashicorp Consul service discovery block.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_fast_gce_service_discovery",
			Category:         "F5 Automation Tool Chain(ATC)",
			ShortDescription: `Provides details for GCE Service discovery config`,
			Description:      ``,
			Keywords: []string{
				"f5",
				"automation",
				"tool",
				"chain",
				"atc",
				"fast",
				"gce",
				"service",
				"discovery",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tag_key",
					Description: `(` + "`" + `Required` + "`" + `,type ` + "`" + `string` + "`" + `) The tag key associated with the node to add to this pool.`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `(` + "`" + `Required` + "`" + `,type ` + "`" + `string` + "`" + `) The tag value associated with the node to add to this pool.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(` + "`" + `Required` + "`" + `,type ` + "`" + `string` + "`" + `) GCE region in which ADC is running.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `int` + "`" + `)Port to be used for AWS service discovery,default ` + "`" + `80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_realm",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Specifies whether to look for public or private IP addresses,default ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "undetectable_action",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Action to take when node cannot be detected,default ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "credential_update",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `bool` + "`" + `) Specifies whether you are updating your credentials,default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "encoded_credentials",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Base 64 encoded service account credentials JSON.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)For Google Cloud Engine (GCE) only: The ID of the project in which the members are located.`,
				},
				resource.Attribute{
					Name:        "minimum_monitors",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Member is down when fewer than minimum monitors report it healthy.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `(` + "`" + `optional` + "`" + `,type ` + "`" + `string` + "`" + `)Update interval for service discovery. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "gce_sd_json",
					Description: `The JSON for GCE service discovery block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gce_sd_json",
					Description: `The JSON for GCE service discovery block.`,
				},
			},
		},
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
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_waf_entity_parameters",
			Category:         "Web Application Firewall(WAF)",
			ShortDescription: `Provides details entity parameters associated with a policy`,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"entity",
				"parameters",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Entity Parameter.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Entity Parameter.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies whether the parameter is an explicit or a wildcard attribute.`,
				},
				resource.Attribute{
					Name:        "value_type",
					Description: `Specify the valid type for the value of the attribute.`,
				},
				resource.Attribute{
					Name:        "allow_empty_type",
					Description: `Determines whether an empty value is allowed for a parameter.`,
				},
				resource.Attribute{
					Name:        "allow_repeated_parameter_name",
					Description: `Determines whether multiple parameter instances with the same name are allowed in one request.`,
				},
				resource.Attribute{
					Name:        "attack_signatures_check",
					Description: `Determines whether attack signatures and threat campaigns must be detected in a parameter's value.`,
				},
				resource.Attribute{
					Name:        "check_max_value_length",
					Description: `Determines whether a parameter has a restricted maximum length for value.`,
				},
				resource.Attribute{
					Name:        "check_min_value_length",
					Description: `Determines whether a parameter has a restricted minimum length for value.`,
				},
				resource.Attribute{
					Name:        "data_type",
					Description: `Specifies data type of parameter's value.`,
				},
				resource.Attribute{
					Name:        "enable_regular_expression",
					Description: `Determines whether the parameter value includes the pattern defined in regularExpression.`,
				},
				resource.Attribute{
					Name:        "is_base64",
					Description: `Determines whether a parameter’s value contains a Base64 encoded string.`,
				},
				resource.Attribute{
					Name:        "is_cookie",
					Description: `Determines whether a parameter is located in the value of Cookie header.`,
				},
				resource.Attribute{
					Name:        "is_header",
					Description: `Determines whether a parameter is located in headers as one of the headers.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `Specifies whether the parameter is associated with a URL, a flow, or neither.`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `Determines whether a parameter must exist in the request.`,
				},
				resource.Attribute{
					Name:        "metachars_on_parameter_value_check",
					Description: `Determines whether disallowed metacharacters must be detected in a parameter’s value.`,
				},
				resource.Attribute{
					Name:        "parameter_location",
					Description: `Specifies location of parameter in request.`,
				},
				resource.Attribute{
					Name:        "perform_staging",
					Description: `Determines the staging state of a parameter.`,
				},
				resource.Attribute{
					Name:        "sensitive_parameter",
					Description: `Determines whether a parameter is sensitive and must be not visible in logs nor in the user interface.`,
				},
				resource.Attribute{
					Name:        "signature_overrides_disable",
					Description: `List of Attack Signature Ids which are disabled for this particular parameter.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `` + "`" + `url` + "`" + ` block will provide options to be used for binding urls to parameter entity.See [url](#url) below for more details. ### url The ` + "`" + `url` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required , ` + "`" + `string` + "`" + `) name of url block attribute`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required , ` + "`" + `string` + "`" + `) Unique ID of a URL with a protocol type and name. Select a Method for the URL to create an API endpoint: URL + Method`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required , ` + "`" + `string` + "`" + `) Specifies whether the protocol for the URL is HTTP or HTTPS.The available options are : ["http","https"]`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required , ` + "`" + `string` + "`" + `) Determines the type of the name attribute. Only when setting the type to wildcard will the special wildcard characters in the name be interpreted as such. The available types are : ["explicit","wildcard"] ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `JSON string representing the WAF Entity Parameter declaration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `JSON string representing the WAF Entity Parameter declaration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_waf_entity_url",
			Category:         "Web Application Firewall(WAF)",
			ShortDescription: `Provides details and exports bigip_waf_entity_url data source`,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"entity",
				"url",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) WAF entity URL name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the URL.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Specifies whether the parameter is an 'explicit' or a 'wildcard' attribute. Default is: wildcard.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Specifies whether the protocol for the URL is 'http' or 'https'. Default is: http.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) Select a Method for the URL to create an API endpoint. Default is :`,
				},
				resource.Attribute{
					Name:        "perform_staging",
					Description: `(Optional) If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.`,
				},
				resource.Attribute{
					Name:        "signature_overrides_disable",
					Description: `(Optional) List of Attack Signature Ids which are disabled for this particular URL.`,
				},
				resource.Attribute{
					Name:        "method_overrides",
					Description: `(Optional) A list of methods that are allowed or disallowed for a specific URL.`,
				},
				resource.Attribute{
					Name:        "allow",
					Description: `(Required) Specifies that the system allows or disallows a method for this URL`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) Specifies an HTTP method. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `Json string representing created WAF entity URL declaration in JSON format`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `Json string representing created WAF entity URL declaration in JSON format`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_waf_pb_suggestions",
			Category:         "Web Application Firewall(WAF)",
			ShortDescription: `Provides details and exports bigip_waf_pb_suggestions data source`,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"pb",
				"suggestions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) WAF policy name from which PB suggestions should be exported.`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Required) Partition on which WAF policy is located.`,
				},
				resource.Attribute{
					Name:        "minimum_learning_score",
					Description: `(Required) The minimum learning score for suggestions. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `System generated id of the WAF policy`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `Json string representing exported PB suggestions ready to be used in WAF policy declaration`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `System generated id of the WAF policy`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `Json string representing exported PB suggestions ready to be used in WAF policy declaration`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_waf_policy",
			Category:         "Web Application Firewall(WAF)",
			ShortDescription: `Provides details about deployed waf policy using its ID`,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) ID of the WAF policy deployed in the BIG-IP. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "policy_json",
					Description: `Exported WAF policy JSON`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_json",
					Description: `Exported WAF policy JSON`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_bigip_waf_signatures",
			Category:         "Web Application Firewall(WAF)",
			ShortDescription: `Provides details about system installed bigip_waf_signatures data source`,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"signatures",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "signature_id",
					Description: `(Required) ID of the signature in the BIG-IP WAF database. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the signature as configured on the system.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the signature.`,
				},
				resource.Attribute{
					Name:        "system_signature_id",
					Description: `System generated ID of the signature.`,
				},
				resource.Attribute{
					Name:        "signature_id",
					Description: `ID of the signature in the database.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the signature.`,
				},
				resource.Attribute{
					Name:        "accuracy",
					Description: `The relative detection accuracy of the signature.`,
				},
				resource.Attribute{
					Name:        "risk",
					Description: `The relative risk level of the attack that matches this signature.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the signature as configured on the system.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the signature.`,
				},
				resource.Attribute{
					Name:        "system_signature_id",
					Description: `System generated ID of the signature.`,
				},
				resource.Attribute{
					Name:        "signature_id",
					Description: `ID of the signature in the database.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the signature.`,
				},
				resource.Attribute{
					Name:        "accuracy",
					Description: `The relative detection accuracy of the signature.`,
				},
				resource.Attribute{
					Name:        "risk",
					Description: `The relative risk level of the attack that matches this signature.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"bigip_bigip_fast_aws_service_discovery":    0,
		"bigip_bigip_fast_azure_service_discovery":  1,
		"bigip_bigip_fast_consul_service_discovery": 2,
		"bigip_bigip_fast_gce_service_discovery":    3,
		"bigip_bigip_ltm_datagroup":                 4,
		"bigip_bigip_ltm_irule":                     5,
		"bigip_bigip_ltm_monitor":                   6,
		"bigip_bigip_ltm_node":                      7,
		"bigip_bigip_ltm_policy":                    8,
		"bigip_bigip_ltm_pool":                      9,
		"bigip_bigip_ssl_certificate":               10,
		"bigip_bigip_vwan_config":                   11,
		"bigip_bigip_waf_entity_parameters":         12,
		"bigip_bigip_waf_entity_url":                13,
		"bigip_bigip_waf_pb_suggestions":            14,
		"bigip_bigip_waf_policy":                    15,
		"bigip_bigip_waf_signatures":                16,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
