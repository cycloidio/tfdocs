package local

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "file",
			Category:         "Data Sources",
			ShortDescription: `Reads a file from the local filesystem.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filename",
					Description: `(Required) The path to the file that will be read. The data source will return an error if the file does not exist. ## Attributes Exported The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `The raw content of the file that was read.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `The base64 encoded version of the file content (use this when dealing with binary data). The content of the file must be valid UTF-8 due to Terraform's assumptions about string encoding. Files that do not contain UTF-8 text will have invalid UTF-8 sequences replaced with the Unicode replacement character.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"file": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
