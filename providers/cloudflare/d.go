package cloudflare

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Get information on Cloudflare IP ranges.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_waf_groups",
			Category:         "Data Sources",
			ShortDescription: `List available Cloudflare WAF Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_waf_packages",
			Category:         "Data Sources",
			ShortDescription: `List available Cloudflare WAF Packages.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_waf_rules",
			Category:         "Data Sources",
			ShortDescription: `List available Cloudflare WAF Rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_zones",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Cloudflare Zones.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"cloudflare_ip_ranges":    0,
		"cloudflare_waf_groups":   1,
		"cloudflare_waf_packages": 2,
		"cloudflare_waf_rules":    3,
		"cloudflare_zones":        4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
