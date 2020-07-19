package cloudamqp

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_alarm",
			Category:         "Resources",
			ShortDescription: `Creates and manages alarms to trigger notifications.`,
			Description:      ``,
			Keywords: []string{
				"alarm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "value_threshold",
					Description: `(Optional) The value to trigger the alarm for.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_instance",
			Category:         "Resources",
			ShortDescription: `Creates and manages a Rabbit MQ instance within CloudAMQP.`,
			Description:      ``,
			Keywords: []string{
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rmq_version",
					Description: `(Optional) The Rabbit MQ version. Default set to current loaded default value in CloudAMQP API.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_integration_log",
			Category:         "Resources",
			ShortDescription: `Creates and manages third party log integration for a CloudAMQP instance.`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"log",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional) AWS secret access key.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_integration_metric",
			Category:         "Resources",
			ShortDescription: `Creates and manages third party metrics integration for a CloudAMQP instance.`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"metric",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_notification",
			Category:         "Resources",
			ShortDescription: `Creates and manages recipients to receive alarm notifications.`,
			Description:      ``,
			Keywords: []string{
				"notification",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The CloudAMQP instance ID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_plugin",
			Category:         "Resources",
			ShortDescription: `Enable and disable Rabbit MQ plugin.`,
			Description:      ``,
			Keywords: []string{
				"plugin",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The CloudAMQP instance ID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_plugin_community",
			Category:         "Resources",
			ShortDescription: `Install or uninstall community plugin.`,
			Description:      ``,
			Keywords: []string{
				"plugin",
				"community",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The CloudAMQP instance ID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_security_firewall",
			Category:         "Resources",
			ShortDescription: `Configure and manage firewall rules`,
			Description:      ``,
			Keywords: []string{
				"security",
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The CloudAMQP instance ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description name of the rule. e.g. Default. Supported services:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_vpc_peering",
			Category:         "Resources",
			ShortDescription: `Accepting VPC peering request from an AWS accepter.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"peering",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The CloudAMQP instance ID.`,
				},
				resource.Attribute{
					Name:        "peering_id",
					Description: `(Required) Peering identifier created by AW peering request. ## Depedency This resource depends on CloudAMQP instance identifier, ` + "`" + `cloudamqp_instance.instance.id` + "`" + `. ## Import ` + "`" + `cloudamqp_vpc_peering` + "`" + ` can be imported using the CloudAMQP instance identifier. ` + "`" + `terraform import cloudamqp_vpc_peering.<resource_name> <instance_id>` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"cloudamqp_alarm":              0,
		"cloudamqp_instance":           1,
		"cloudamqp_integration_log":    2,
		"cloudamqp_integration_metric": 3,
		"cloudamqp_notification":       4,
		"cloudamqp_plugin":             5,
		"cloudamqp_plugin_community":   6,
		"cloudamqp_security_firewall":  7,
		"cloudamqp_vpc_peering":        8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
