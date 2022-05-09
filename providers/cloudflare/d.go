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
			ShortDescription: `Get information on a Cloudflare Access Identity Provider by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The account for which to look for an Access Identity Provider. Conflicts with ` + "`" + `zone_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The Zone's ID. Conflicts with ` + "`" + `account_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Access Identity Provider name to search for. ## Attributes Reference - ` + "`" + `id` + "`" + ` - Access Identity Provider ID - ` + "`" + `name` + "`" + ` - Access Identity Provider Name - ` + "`" + `type` + "`" + ` - Access Identity Provider Type [access_identity_provider_guide]: https://developers.cloudflare.com/cloudflare-one/identity/idp-integration`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_account_roles",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Cloudflare Account Roles.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The account for which to list the roles. ## Attributes Reference - ` + "`" + `roles` + "`" + ` - A list of roles object. See below for nested attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_api_token_permission_groups",
			Category:         "Data Sources",
			ShortDescription: `List available API Token Permission Group IDs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudflare_devices",
			Category:         "Data Sources",
			ShortDescription: `Get information on Cloudflare Devices.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The account for which to list the devices. ## Attributes Reference - ` + "`" + `devices` + "`" + ` - A list of device object. See below for nested attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
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
			Type:             "cloudflare_origin_ca_root_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get Cloudflare Origin CA root certificate.`,
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

		"cloudflare_access_identity_provider":    0,
		"cloudflare_account_roles":               1,
		"cloudflare_api_token_permission_groups": 2,
		"cloudflare_devices":                     3,
		"cloudflare_ip_ranges":                   4,
		"cloudflare_origin_ca_root_certificate":  5,
		"cloudflare_waf_groups":                  6,
		"cloudflare_waf_packages":                7,
		"cloudflare_waf_rules":                   8,
		"cloudflare_zones":                       9,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
