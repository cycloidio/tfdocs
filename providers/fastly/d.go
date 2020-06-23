package fastly

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fastly_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly IP ranges.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `The lexically ordered list of ipv4 CIDR blocks.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_blocks",
					Description: `The lexically ordered list of ipv6 CIDR blocks. [1]: https://docs.fastly.com/guides/securing-communications/accessing-fastlys-ip-ranges`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"fastly_ip_ranges": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
