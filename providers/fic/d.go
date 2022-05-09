package fic

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_switch_v1",
			Category:         "Data Sources",
			ShortDescription: `Get a V1 Switch information within Flexible InterConnect.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Alias name of switch.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `(Optional) Area name.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Location(Data center) name.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `(Required) Port type, 1G or 10G. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of switch.`,
				},
				resource.Attribute{
					Name:        "number_of_available_vlans",
					Description: `Number of available VLANs.`,
				},
				resource.Attribute{
					Name:        "vlan_ranges",
					Description: `List of available VLAN ranges.`,
				},
				resource.Attribute{
					Name:        "vlan_ranges/start",
					Description: `Start number of VLAN range.`,
				},
				resource.Attribute{
					Name:        "vlan_ranges/end",
					Description: `End number of VLAN range.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of switch.`,
				},
				resource.Attribute{
					Name:        "number_of_available_vlans",
					Description: `Number of available VLANs.`,
				},
				resource.Attribute{
					Name:        "vlan_ranges",
					Description: `List of available VLAN ranges.`,
				},
				resource.Attribute{
					Name:        "vlan_ranges/start",
					Description: `Start number of VLAN range.`,
				},
				resource.Attribute{
					Name:        "vlan_ranges/end",
					Description: `End number of VLAN range.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"fic_eri_switch_v1": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
