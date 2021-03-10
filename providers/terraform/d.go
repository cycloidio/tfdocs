package terraform

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "terraform_remote_state",
			Category:         "Data Sources",
			ShortDescription: `Accesses state meta data from a remote backend.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The remote backend to use.`,
				},
				resource.Attribute{
					Name:        "workspace",
					Description: `(Optional) The Terraform workspace to use.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) The configuration of the remote backend.`,
				},
				resource.Attribute{
					Name:        "defaults",
					Description: `(Optional) default value for outputs in case state file is empty or it does not have the output.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `See Argument Reference above. In addition, each output in the remote state appears as a top level attribute on the ` + "`" + `terraform_remote_state` + "`" + ` resource. ## Root Outputs Only Only the root level outputs from the remote state are accessible. Outputs from modules within the state cannot be accessed. If you want a module output to be accessible via a remote state, you must thread the output through to a root output. An example is shown below: ` + "`" + `` + "`" + `` + "`" + `hcl module "app" { source = "..." } output "app_value" { value = "${module.app.value}" } ` + "`" + `` + "`" + `` + "`" + ` In this example, the output ` + "`" + `value` + "`" + ` from the "app" module is available as "app_value". If this root level output hadn't been created, then a remote state resource wouldn't be able to access the ` + "`" + `value` + "`" + ` output on the module.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "backend",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `See Argument Reference above. In addition, each output in the remote state appears as a top level attribute on the ` + "`" + `terraform_remote_state` + "`" + ` resource. ## Root Outputs Only Only the root level outputs from the remote state are accessible. Outputs from modules within the state cannot be accessed. If you want a module output to be accessible via a remote state, you must thread the output through to a root output. An example is shown below: ` + "`" + `` + "`" + `` + "`" + `hcl module "app" { source = "..." } output "app_value" { value = "${module.app.value}" } ` + "`" + `` + "`" + `` + "`" + ` In this example, the output ` + "`" + `value` + "`" + ` from the "app" module is available as "app_value". If this root level output hadn't been created, then a remote state resource wouldn't be able to access the ` + "`" + `value` + "`" + ` output on the module.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"terraform_remote_state": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
