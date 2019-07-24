package cloudstack

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_template",
			Category:         "Data Sources",
			ShortDescription: `Get informations on a Cloudstack template.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_filter",
					Description: `(Required) The template filter. Possible values are ` + "`" + `featured` + "`" + `, ` + "`" + `self` + "`" + `, ` + "`" + `selfexecutable` + "`" + `, ` + "`" + `sharedexecutable` + "`" + `, ` + "`" + `executable` + "`" + ` and ` + "`" + `community` + "`" + ` (see the Cloudstack API`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) One or more name/value pairs to filter off of. You can apply filters on any exported attributes. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The template ID.`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `The account name to which the template belongs.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date this template was created.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `The template display text.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `The format of the template.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor on which the templates runs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the template.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The template ID.`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `The account name to which the template belongs.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date this template was created.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `The template display text.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `The format of the template.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor on which the templates runs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the template.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"cloudstack_template": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
