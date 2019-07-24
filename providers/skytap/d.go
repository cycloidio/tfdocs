package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "skytap_project",
			Category:         "Data Sources",
			ShortDescription: `Get information on a project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of project. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skytap_template",
			Category:         "Data Sources",
			ShortDescription: `Get information on a template.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A regular expression on the name of a template.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) Use the most recently created template from the returned list. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	dataSourcesMap = map[string]Resource{

		"skytap_project":  0,
		"skytap_template": 1,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
