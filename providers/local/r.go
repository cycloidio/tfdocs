package local

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "local_file",
			Category:         "Resources",
			ShortDescription: `Generates a local file from content.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filename",
					Description: `(Required) The path to the file that will be created. Missing parent directories will be created. If the file already exists, it will be overridden with the given content.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Content to store in the file, expected to be an UTF-8 encoded string. Conflicts with ` + "`" + `sensitive_content` + "`" + `, ` + "`" + `content_base64` + "`" + ` and ` + "`" + `source` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sensitive_content",
					Description: `(Optional - Deprecated) Sensitive content to store in the file, expected to be an UTF-8 encoded string. Will not be displayed in diffs. Conflicts with ` + "`" + `content` + "`" + `, ` + "`" + `content_base64` + "`" + ` and ` + "`" + `source` + "`" + `. If in need to use _sensitive_ content, please use the [` + "`" + `local_sensitive_file` + "`" + `](./sensitive_file.html) resource instead.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `(Optional) Content to store in the file, expected to be binary encoded as base64 string. Conflicts with ` + "`" + `content` + "`" + `, ` + "`" + `sensitive_content` + "`" + ` and ` + "`" + `source` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Path to file to use as source for the one we are creating. Conflicts with ` + "`" + `content` + "`" + `, ` + "`" + `sensitive_content` + "`" + ` and ` + "`" + `content_base64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_permission",
					Description: `(Optional) Permissions to set for the output file, expressed as string in [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation). Default value is ` + "`" + `"0777"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory_permission",
					Description: `(Optional) Permissions to set for directories created, expressed as string in [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation). Default value is ` + "`" + `"0777"` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "local_sensitive_file",
			Category:         "Resources",
			ShortDescription: `Generates a local file with the given sensitive content.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filename",
					Description: `(Required) The path to the file that will be created. Missing parent directories will be created. If the file already exists, it will be overridden with the given content.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Sensitive content to store in the file, expected to be an UTF-8 encoded string. Conflicts with ` + "`" + `content_base64` + "`" + ` and ` + "`" + `source` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `(Optional) Sensitive content to store in the file, expected to be binary encoded as base64 string. Conflicts with ` + "`" + `content` + "`" + ` and ` + "`" + `source` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Path to file to use as source for the one we are creating. Conflicts with ` + "`" + `content` + "`" + ` and ` + "`" + `content_base64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_permission",
					Description: `(Optional) Permissions to set for the output file, expressed as string in [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation). Default value is ` + "`" + `"0700"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory_permission",
					Description: `(Optional) Permissions to set for directories created, expressed as string in [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation). Default value is ` + "`" + `"0700"` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"local_file":           0,
		"local_sensitive_file": 1,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
