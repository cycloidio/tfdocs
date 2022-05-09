package local

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "local_file",
			Category:         "Data Sources",
			ShortDescription: `Reads a file from the local filesystem.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filename",
					Description: `(Required) Path to the file that will be read. The data source will return an error if the file does not exist. ## Attributes Exported The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Raw content of the file that was read, assumed by Terraform to be UTF-8 encoded string. Files that do not contain UTF-8 text will have invalid UTF-8 sequences in ` + "`" + `content` + "`" + ` replaced with the Unicode replacement character.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `Base64 encoded version of the file content. Use this when dealing with binary data.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "local_sensitive_file",
			Category:         "Data Sources",
			ShortDescription: `Reads a file that contains sensitive data, from the local filesystem.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filename",
					Description: `(Required) Path to the file that will be read. The data source will return an error if the file does not exist. ## Attributes Exported The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Raw content of the file that was read, assumed by Terraform to be UTF-8 encoded string. Files that do not contain UTF-8 text will have invalid UTF-8 sequences in ` + "`" + `content` + "`" + ` replaced with the Unicode replacement character.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `Base64 encoded version of the file content. Use this when dealing with binary data.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"local_file":           0,
		"local_sensitive_file": 1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
