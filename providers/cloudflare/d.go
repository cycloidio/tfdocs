package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Get information on Cloudflare IP ranges.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Argument{},
			Attributes:       []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_zones",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Cloudflare Zones.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Argument{},
			Attributes:       []resource.Argument{},
		},
	}

	dataSourcesMap = map[string]Resource{

		"cloudflare_ip_ranges": 0,
		"cloudflare_zones":     1,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
