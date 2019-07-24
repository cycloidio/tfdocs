package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "template_cloudinit_config",
			Category:         "Data Sources",
			ShortDescription: `Renders a multi-part cloud-init config from source files.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "rendered",
					Description: `The final rendered multi-part cloud-init config.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "template_file",
			Category:         "Data Sources",
			ShortDescription: `Renders a template from a file.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The contents of the template, as a string using [Terraform template syntax](/docs/configuration/expressions.html#string-templates). Use [the ` + "`" + `file` + "`" + ` function](/docs/configuration/functions/file.html) to load the template source from a separate file on disk.`,
				},
				resource.Attribute{
					Name:        "vars",
					Description: `(Optional) Variables for interpolation within the template. Note that variables must all be primitives. Direct references to lists or maps will cause a validation error. Earlier versions of ` + "`" + `template_file` + "`" + ` accepted another argument ` + "`" + `filename` + "`" + ` as an alternative to ` + "`" + `template` + "`" + `. This has now been removed. Use the ` + "`" + `template` + "`" + ` argument with the ` + "`" + `file` + "`" + ` function to get the same effect. ## Template Syntax The ` + "`" + `template` + "`" + ` argument is processed as [Terraform template syntax](/docs/configuration/expressions.html#string-templates). However, this provider has its own copy of the template engine embedded in it, separate from Terraform itself, and so which features are available are decided based on what Terraform version the provider was compiled against, and not on which Terraform version you are running. For more consistent results, Terraform 0.12 has a built in function [` + "`" + `templatefile` + "`" + `](/docs/configuration/functions/templatefile.html) which serves the same purpose as this data source. Use that function instead if you are using Terraform 0.12 or later. Its template and expression capabilities will always match the version of Terraform you are using. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vars",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rendered",
					Description: `The final rendered template.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vars",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rendered",
					Description: `The final rendered template.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]Resource{

		"template_cloudinit_config": 0,
		"template_file":             1,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
