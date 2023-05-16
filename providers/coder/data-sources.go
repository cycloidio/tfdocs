package coder

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "coder_git_auth",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to require users to authenticate with a Git provider prior to workspace creation. This can be used to perform an authenticated git clone in startup scripts.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coder_parameter",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to configure editable options for workspaces.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coder_provisioner",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information about the Coder provisioner.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coder_workspace",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information for the active workspace build.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"coder_git_auth":    0,
		"coder_parameter":   1,
		"coder_provisioner": 2,
		"coder_workspace":   3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
