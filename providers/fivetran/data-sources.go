package fivetran

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fivetran_connector",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This data source returns a connector object.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fivetran_connectors_metadata",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This data source returns all available source types within your Fivetran account. This data source makes it easier to display Fivetran connectors within your application because it provides metadata including the proper source name (‘Facebook Ad Account’ instead of facebook_ad_account), the source icon, and links to Fivetran resources. As we update source names and icons, that metadata will automatically update within this endpoint.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fivetran_destination",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This data source returns a destination object.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fivetran_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This data source returns a group object.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fivetran_group_connectors",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This data source returns a list of information about all connectors within a group in your Fivetran account.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fivetran_group_users",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This data source returns a list of information about all users within a group in your Fivetran account.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fivetran_groups",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This data source returns a list of all groups within your Fivetran account.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fivetran_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This data source returns a user object.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fivetran_users",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This data source returns a list of all users within your Fivetran account.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"fivetran_connector":           0,
		"fivetran_connectors_metadata": 1,
		"fivetran_destination":         2,
		"fivetran_group":               3,
		"fivetran_group_connectors":    4,
		"fivetran_group_users":         5,
		"fivetran_groups":              6,
		"fivetran_user":                7,
		"fivetran_users":               8,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
