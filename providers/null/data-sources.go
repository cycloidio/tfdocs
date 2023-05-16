package null

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "null_data_source",
			Category:         "Data Sources",
			ShortDescription: `The null_data_source data source implements the standard data source lifecycle but does not interact with any external APIs. Historically, the null_data_source was typically used to construct intermediate values to re-use elsewhere in configuration. The same can now be achieved using locals https://developer.hashicorp.com/terraform/language/values/locals or the terraform_data resource type https://developer.hashicorp.com/terraform/language/resources/terraform-data in Terraform 1.4 and later.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"null_data_source": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
