package gocache

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "gocache_gocache_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `The lexically ordered list of all CIDR blocks.`,
				},
				resource.Attribute{
					Name:        "ipv4_cidr_blocks",
					Description: `The lexically ordered list of only the IPv4 CIDR blocks.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"gocache_gocache_ip_ranges": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
