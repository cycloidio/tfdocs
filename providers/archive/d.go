package archive

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "archive_file",
			Category:         "Data Sources",
			ShortDescription: `Generates an archive from content, a file, or directory of files.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of archive to generate. NOTE: ` + "`" + `zip` + "`" + ` is supported.`,
				},
				resource.Attribute{
					Name:        "output_path",
					Description: `(Required) The output of the archive file.`,
				},
				resource.Attribute{
					Name:        "source_content",
					Description: `(Optional) Add only this content to the archive with ` + "`" + `source_content_filename` + "`" + ` as the filename.`,
				},
				resource.Attribute{
					Name:        "source_content_filename",
					Description: `(Optional) Set this as the filename when using ` + "`" + `source_content` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_file",
					Description: `(Optional) Package this file into the archive.`,
				},
				resource.Attribute{
					Name:        "source_dir",
					Description: `(Optional) Package entire contents of this directory into the archive.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Specifies attributes of a single source file to include into the archive.`,
				},
				resource.Attribute{
					Name:        "excludes",
					Description: `(Optional) Specify files to ignore when reading the ` + "`" + `source_dir` + "`" + `. The ` + "`" + `source` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) Add this content to the archive with ` + "`" + `filename` + "`" + ` as the filename.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Required) Set this as the filename when declaring a ` + "`" + `source` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_size",
					Description: `The size of the output archive file.`,
				},
				resource.Attribute{
					Name:        "output_sha",
					Description: `The SHA1 checksum of output archive file.`,
				},
				resource.Attribute{
					Name:        "output_base64sha256",
					Description: `The base64-encoded SHA256 checksum of output archive file.`,
				},
				resource.Attribute{
					Name:        "output_md5",
					Description: `The MD5 checksum of output archive file.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_size",
					Description: `The size of the output archive file.`,
				},
				resource.Attribute{
					Name:        "output_sha",
					Description: `The SHA1 checksum of output archive file.`,
				},
				resource.Attribute{
					Name:        "output_base64sha256",
					Description: `The base64-encoded SHA256 checksum of output archive file.`,
				},
				resource.Attribute{
					Name:        "output_md5",
					Description: `The MD5 checksum of output archive file.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"archive_file": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
