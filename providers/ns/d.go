package ns

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ns_app_connection",
			Category:         "Data Sources",
			ShortDescription: `Data source to configure a connection to another nullstone workspace through a capability's application.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of nullstone connection.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of nullstone module to make connection.`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `By default, if this connection has not been configured, this causes an error. Set to true to disable. (Default: ` + "`" + `false` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "via",
					Description: `Name of connection to satisfy this connection through. Typically, this is set to ` + "`" + `data.ns_connection.other.name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `This refers to the workspace in nullstone. This follows the form ` + "`" + `{stack_id}/{block_id}/{env_id}` + "`" + `. - ` + "`" + `outputs` + "`" + ` - An object containing every root-level output in the remote state. This attribute is interchangeable for ` + "`" + `data.terraform_remote_state.outputs` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns_app_env",
			Category:         "Data Sources",
			ShortDescription: `Data source to read information about an application in a specific environment.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `The version configured in nullstone for the application in the specific environment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns_connection",
			Category:         "Data Sources",
			ShortDescription: `Data source to configure a connection to another nullstone workspace.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of nullstone connection.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of nullstone module to make connection.`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `By default, if this connection has not been configured, this causes an error. Set to true to disable. (Default: ` + "`" + `false` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "via",
					Description: `Name of connection to satisfy this connection through. Typically, this is set to ` + "`" + `data.ns_connection.other.name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `This refers to the workspace in nullstone. This follows the form ` + "`" + `{stack_id}/{block_id}/{env_id}` + "`" + `. - ` + "`" + `outputs` + "`" + ` - An object containing every root-level output in the remote state. This attribute is interchangeable for ` + "`" + `data.terraform_remote_state.outputs` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source to read a nullstone domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dns_name",
					Description: `The domain name that has been configured for this domain. An example would be ` + "`" + `nullstone.io` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dns_name",
					Description: `The domain name that has been configured for this domain. An example would be ` + "`" + `nullstone.io` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns_subdomain",
			Category:         "Data Sources",
			ShortDescription: `Data source to read a nullstone subdomain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the subdomain.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The subdomain name that has been configured for this domain. An example would be ` + "`" + `api` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the subdomain.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The subdomain name that has been configured for this domain. An example would be ` + "`" + `api` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns_workspace",
			Category:         "Data Sources",
			ShortDescription: `Data source to configure module based on current nullstone workspace.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The fully qualified workspace ID. This follows the form ` + "`" + `{stack_id}/{block_id}/{env_id}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "stack_id",
					Description: `Workspace stack ID. (Environment variable: ` + "`" + `NULLSTONE_STACK_ID` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "stack_name",
					Description: `Workspace stack name. (Environment variable: ` + "`" + `NULLSTONE_STACK_NAME` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "block_id",
					Description: `Workspace block ID. (Environment variable: ` + "`" + `NULLSTONE_BLOCK_ID` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "block_name",
					Description: `Workspace block name. (Environment variable: ` + "`" + `NULLSTONE_BLOCK_NAME` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "block_ref",
					Description: `Workspace block reference. Unique name used for constructing resource names. (Environment variable: ` + "`" + `NULLSTONE_BLOCK_REF` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "env_id",
					Description: `Workspace environment ID. (Environment variable: ` + "`" + `NULLSTONE_ENV_ID` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "env_name",
					Description: `Workspace environment name. (Environment variable: ` + "`" + `NULLSTONE_ENV_NAME` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `Use ` + "`" + `id` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `Use ` + "`" + `stack_name` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Use ` + "`" + `env_name` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "block",
					Description: `Use ` + "`" + `block_name` + "`" + ` instead.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ns_app_connection": 0,
		"ns_app_env":        1,
		"ns_connection":     2,
		"ns_domain":         3,
		"ns_subdomain":      4,
		"ns_workspace":      5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
