package codefresh

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "codefresh_account",
			Category:         "Data Sources",
			ShortDescription: `This data source retrieves an account by _id or name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_context",
			Category:         "Data Sources",
			ShortDescription: `This data source allows to retrieve information on any defined context.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_current_account",
			Category:         "Data Sources",
			ShortDescription: `Returns the current account (owner of the token) and its users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_idps",
			Category:         "Data Sources",
			ShortDescription: `This data source retrieves all Identity Providers (IdPs) in the system.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_pipelines",
			Category:         "Data Sources",
			ShortDescription: `This resource retrives all pipelines belonging to the current user, which can be optionally filtered by the name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_registry",
			Category:         "Data Sources",
			ShortDescription: `This data source allows retrieving information on any existing registry.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_step_types",
			Category:         "Data Sources",
			ShortDescription: `This data source allows to retrieve the published versions of step-types.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_team",
			Category:         "Data Sources",
			ShortDescription: `This data source retrieves a team by its ID or name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_user",
			Category:         "Data Sources",
			ShortDescription: `This data source retrieves a user by email.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_users",
			Category:         "Data Sources",
			ShortDescription: `This data source retrieves all users in the system.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"codefresh_account":         0,
		"codefresh_context":         1,
		"codefresh_current_account": 2,
		"codefresh_idps":            3,
		"codefresh_pipelines":       4,
		"codefresh_registry":        5,
		"codefresh_step_types":      6,
		"codefresh_team":            7,
		"codefresh_user":            8,
		"codefresh_users":           9,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
