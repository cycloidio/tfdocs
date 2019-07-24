package template

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "template_dir",
			Category:         "Resources",
			ShortDescription: `Renders a directory of templates.`,
			Description:      ``,
			Keywords: []string{
				"dir",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_dir",
					Description: `(Required) Path to the directory where the files to template reside.`,
				},
				resource.Attribute{
					Name:        "destination_dir",
					Description: `(Required) Path to the directory where the templated files will be written.`,
				},
				resource.Attribute{
					Name:        "vars",
					Description: `(Optional) Variables for interpolation within the template. Note that variables must all be primitives. Direct references to lists or maps will cause a validation error. Any required parent directories of ` + "`" + `destination_dir` + "`" + ` will be created automatically, and any pre-existing file or directory at that location will be deleted before template rendering begins. After rendering, this resource remembers the content of both the source and destination directories in the Terraform state, and will plan to recreate the output directory if any changes are detected during the plan phase. Note that it is _not_ safe to use the ` + "`" + `file` + "`" + ` interpolation function to read files created by this resource, since that function can be evaluated before the destination directory has been created or updated. It`,
				},
				resource.Attribute{
					Name:        "destination_dir",
					Description: `The destination directory given in configuration. Interpolate this attribute into other resource configurations to create a dependency to ensure that the destination directory is populated before another resource attempts to read it.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"template_dir": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
