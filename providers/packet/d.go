package packet

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "packet_operating_system",
			Category:         "Data Sources",
			ShortDescription: `Get a Packet operating system image`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "distro",
					Description: `(Optional) Name of the OS distribution.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name or part of the name of the distribution. Case insensitive.`,
				},
				resource.Attribute{
					Name:        "provisionable_on",
					Description: `(Optional) Plan name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the distribution ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Operating system slug`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Operating system slug (same as ` + "`" + `id` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Operating system slug`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Operating system slug (same as ` + "`" + `id` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "packet_precreated_ip_block",
			Category:         "Data Sources",
			ShortDescription: `Load automatically created IP blocks from your Packet project`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) ID of the project where the searched block should be.`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Required) 4 or 6, depending on which block you are looking for.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(Required) Whether to look for public or private block.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional) Whether to look for global block. Default is false for backward compatibility.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) Facility of the searched block. (Optional) Only allowed for non-global blocks. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `CIDR notation of the looked up block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `CIDR notation of the looked up block.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "packet_spot_market_price",
			Category:         "Data Sources",
			ShortDescription: `Get a Packet Spot Market Price`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "facility",
					Description: `(Required) Name of the facility.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Name of the plan. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Current spot market price for given plan in given facility.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "price",
					Description: `Current spot market price for given plan in given facility.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"packet_operating_system":    0,
		"packet_precreated_ip_block": 1,
		"packet_spot_market_price":   2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
