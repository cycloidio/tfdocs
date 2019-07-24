package logicmonitor

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "logicmonitor_collectors",
			Category:         "Data Sources",
			ShortDescription: `Use this datasource to get the ID of an available collector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The number of results to display. Max is 1000. Default is 50`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) The number of results to offset the displayed results by. Default is 0`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) The most recent collector installed that is online`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `(Optional) Filters the response according to the operator and value specified. Note that you can use`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `(Required if using filters) The name of filtered property. Currently the properties supported are ` + "`" + `hostname` + "`" + ` and ` + "`" + `description` + "`" + ``,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required if using filters) The type of operator. Currently the operators supported are ` + "`" + `:` + "`" + ` ` + "`" + `~` + "`" + ` ` + "`" + `!:` + "`" + ` ` + "`" + `!~` + "`" + ``,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required if using filters) The value of the filtered property.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logicmonitor_device_group",
			Category:         "Data Sources",
			ShortDescription: `Use this datasource to get the ID of an available device group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The number of results to display. Max is 1000. Default is 50`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) The number of results to offset the displayed results by. Default is 0`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `(Optional) Filters the response according to the operator and value specified. Note that you can use`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `(Required if using filters) The name of filtered property.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required if using filters) The type of operator.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required if using filters) The value of the filtered property. You can also do custom properties`,
				},
				resource.Attribute{
					Name:        "custom_property_name",
					Description: `(Required if using filters and custom properties) The name of filtered custom property.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required if using filters) The type of operator.`,
				},
				resource.Attribute{
					Name:        "custom_property_value",
					Description: `(Required if using filters and custom properties) The value of the filtered custom property.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"logicmonitor_collectors":   0,
		"logicmonitor_device_group": 1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
