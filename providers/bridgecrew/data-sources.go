package bridgecrew

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_apitokens",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform apitokens. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/apitokenlistbyuserid>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_apitokens_customer",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all of your Bridgecrew platform apitokens. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/apitokenlistbycustomername>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_errors",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform errors. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getgitblameauthors>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_incidents",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform incidents. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getincidents>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_incidents_info",
			Category:         "Data Sources",
			ShortDescription: `Gets all the info and counters of the incidents and violations. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getinfo>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_incidents_preset",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform incidents presets and counters. More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getpresets>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_integrations",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform integrations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_organisation",
			Category:         "Data Sources",
			ShortDescription: `Get your Bridgecrew Organisation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_policies",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all of your custom policies in the Bridgecrew platform. Details on this API call are available here <https://docs.bridgecrew.io/reference/getcustompoliciestable>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_repositories",
			Category:         "Data Sources",
			ShortDescription: `Gets a list of all your repositories managed by the Bridgecrew platform <https://docs.bridgecrew.io/reference/getrepositories>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_repository_branches",
			Category:         "Data Sources",
			ShortDescription: `Supplies the details on the branches under analysis for a given repository <https://docs.bridgecrew.io/reference/getbranches>.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_suppressions",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your policies suppressions in the Bridgecrew platform <https://docs.bridgecrew.io/reference/getsuppressions>`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_users",
			Category:         "Data Sources",
			ShortDescription: `Get a list of all your Bridgecrew platform users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"bridgecrew_apitokens":           0,
		"bridgecrew_apitokens_customer":  1,
		"bridgecrew_errors":              2,
		"bridgecrew_incidents":           3,
		"bridgecrew_incidents_info":      4,
		"bridgecrew_incidents_preset":    5,
		"bridgecrew_integrations":        6,
		"bridgecrew_organisation":        7,
		"bridgecrew_policies":            8,
		"bridgecrew_repositories":        9,
		"bridgecrew_repository_branches": 10,
		"bridgecrew_suppressions":        11,
		"bridgecrew_users":               12,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
