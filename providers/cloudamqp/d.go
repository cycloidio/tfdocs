package cloudamqp

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_alarm",
			Category:         "Data Sources",
			ShortDescription: `Get information on pre-defined or created alarms.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_credentials",
			Category:         "Data Sources",
			ShortDescription: `Get credentials information`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_instance",
			Category:         "Data Sources",
			ShortDescription: `Get information about an already created CloudAMQP instance`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_nodes",
			Category:         "Data Sources",
			ShortDescription: `Get information about the node(s) in the CloudAMQP instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_notification",
			Category:         "Data Sources",
			ShortDescription: `Get information on pre-defined or created recipients.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_plugins",
			Category:         "Data Sources",
			ShortDescription: `Get information installed and available plugins.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_plugins_community",
			Category:         "Data Sources",
			ShortDescription: `Get information about available community plugins.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudamqp_vpc_info",
			Category:         "Data Sources",
			ShortDescription: `Get information about VPC hosted in AWS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"cloudamqp_alarm":             0,
		"cloudamqp_credentials":       1,
		"cloudamqp_instance":          2,
		"cloudamqp_nodes":             3,
		"cloudamqp_notification":      4,
		"cloudamqp_plugins":           5,
		"cloudamqp_plugins_community": 6,
		"cloudamqp_vpc_info":          7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
