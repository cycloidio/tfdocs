package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "file",
			Category:         "Resources",
			ShortDescription: `Generates a local file from content.`,
			Description:      ``,
			Keywords: []string{
				"file",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) The content of file to create. Conflicts with ` + "`" + `sensitive_content` + "`" + ` and ` + "`" + `content_base64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sensitive_content",
					Description: `(Optional) The content of file to create. Will not be displayed in diffs. Conflicts with ` + "`" + `content` + "`" + ` and ` + "`" + `content_base64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `(Optional) The base64 encoded content of the file to create. Use this when dealing with binary data. Conflicts with ` + "`" + `content` + "`" + ` and ` + "`" + `sensitive_content` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Required) The path of the file to create. Any required parent directories will be created automatically, and any existing file with the given name will be overwritten.`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"file": 0,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
