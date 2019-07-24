package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fastly_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly IP ranges.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Argument{},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `The lexically ordered list of CIDR blocks. [1]: https://docs.fastly.com/guides/securing-communications/accessing-fastlys-ip-ranges`,
				},
			},
		},
	}

	dataSourcesMap = map[string]Resource{

		"fastly_ip_ranges": 0,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
