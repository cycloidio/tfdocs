package cloudinit

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudinit_config",
			Category:         "Data Sources",
			ShortDescription: `Renders a multi-part cloud-init config from source files.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gzip",
					Description: `(Optional) Specify whether or not to gzip the rendered output. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "base64_encode",
					Description: `(Optional) Base64 encoding of the rendered output. Defaults to ` + "`" + `true` + "`" + `, and cannot be disabled if ` + "`" + `gzip` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "part",
					Description: `(Required) A nested block type which adds a file to the generated cloud-init configuration. Use multiple ` + "`" + `part` + "`" + ` blocks to specify multiple files, which will be included in order of declaration in the final MIME document. Each ` + "`" + `part` + "`" + ` block expects the following arguments:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) Body content for the part.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Optional) A filename to report in the header for the part.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) A MIME-style content type to report in the header for the part.`,
				},
				resource.Attribute{
					Name:        "merge_type",
					Description: `(Optional) A value for the ` + "`" + `X-Merge-Type` + "`" + ` header of the part, to control [cloud-init merging behavior](https://cloudinit.readthedocs.io/en/latest/topics/merging.html). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rendered",
					Description: `The final rendered multi-part cloud-init config.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rendered",
					Description: `The final rendered multi-part cloud-init config.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"cloudinit_config": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
