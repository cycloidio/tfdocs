package lacework

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "lacework_agent_access_token",
			Category:         "Data Sources",
			ShortDescription: `Lookup agent access token.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The agent access token name. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The agent access token.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "token",
					Description: `The agent access token.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_api_token",
			Category:         "Data Sources",
			ShortDescription: `Generate API access token.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "token",
					Description: `An API access token.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"lacework_agent_access_token": 0,
		"lacework_api_token":          1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
