package cloudflare

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_access_identity_provider",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to lookup a single Access Identity Provider https://developers.cloudflare.com/cloudflare-one/identity/idp-integration by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_account_roles",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_accounts",
			Category:         "Data Sources",
			ShortDescription: `Data source for looking up Cloudflare Accounts.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_api_token_permission_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to look up API Token Permission Groups https://developers.cloudflare.com/api/tokens/create/permissions. Commonly used as references within cloudflare_token resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_devices",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get the IP ranges https://www.cloudflare.com/ips/ of Cloudflare network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_list",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to lookup a List https://developers.cloudflare.com/api/operations/lists-get-lists.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_lists",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to lookup Lists https://developers.cloudflare.com/api/operations/lists-get-lists.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_load_balancer_pools",
			Category:         "Data Sources",
			ShortDescription: `A datasource to find Load Balancer Pools.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_origin_ca_root_certificate",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_record",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_rulesets",
			Category:         "Data Sources",
			ShortDescription: `Use this datasource to lookup Rulesets in an account or zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_zone",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_zone_dnssec",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to look up Zone DNSSEC settings.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_zones",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to look up Zone results for use in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"cloudflare_access_identity_provider":    0,
		"cloudflare_account_roles":               1,
		"cloudflare_accounts":                    2,
		"cloudflare_api_token_permission_groups": 3,
		"cloudflare_devices":                     4,
		"cloudflare_ip_ranges":                   5,
		"cloudflare_list":                        6,
		"cloudflare_lists":                       7,
		"cloudflare_load_balancer_pools":         8,
		"cloudflare_origin_ca_root_certificate":  9,
		"cloudflare_record":                      10,
		"cloudflare_rulesets":                    11,
		"cloudflare_zone":                        12,
		"cloudflare_zone_dnssec":                 13,
		"cloudflare_zones":                       14,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
