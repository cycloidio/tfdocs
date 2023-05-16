package susepubliccloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "susepubliccloud_susepubliccloud_image_ids",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud",
					Description: `(Required) Name of the target cloud to use. Valid values: ` + "`" + `amazon` + "`" + `, ` + "`" + `google` + "`" + `, ` + "`" + `microsoft` + "`" + ` and ` + "`" + `oracle` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) One of the known regions in the cloud framework. Use the region identifiers as the provider describes them, for example ` + "`" + `us-east-1` + "`" + ` in Amazon EC2, or ` + "`" + `East US 2` + "`" + ` in Microsoft Azure.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Defaults to ` + "`" + `active` + "`" + `) State of the image. Valid values: ` + "`" + `active` + "`" + `, ` + "`" + `inactive` + "`" + `, ` + "`" + `deprecated` + "`" + `. Note well: the ` + "`" + `deleted` + "`" + ` state isn't accepted by the data source because these images would not be usable by terraform.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the images list returned by the remote API managed by SUSE.`,
				},
				resource.Attribute{
					Name:        "sort_ascending",
					Description: `(Defaults to ` + "`" + `false` + "`" + `) Used to sort by publication time.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"susepubliccloud_susepubliccloud_image_ids": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
