package artifactory

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_file",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Name of the repository where the file is stored.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path to the file within the repository.`,
				},
				resource.Attribute{
					Name:        "output_path",
					Description: `(Required) The local path the file should be downloaded to.`,
				},
				resource.Attribute{
					Name:        "force_overwrite",
					Description: `(Optional) If set to true, an existing file in the output_path will be overwritten. Default: false`,
				},
				resource.Attribute{
					Name:        "path_is_aliased",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory if the file exists. More details in the [official documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RetrieveLatestArtifact) ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The time & date when the file was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user who created the file.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The time & date when the file was last modified.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `The user who last modified the file.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The time & date when the file was last updated.`,
				},
				resource.Attribute{
					Name:        "mimetype",
					Description: `The mimetype of the file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the file.`,
				},
				resource.Attribute{
					Name:        "download_uri",
					Description: `The URI that can be used to download the file.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `MD5 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha1",
					Description: `SHA1 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `SHA256 checksum of the file.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `The time & date when the file was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user who created the file.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The time & date when the file was last modified.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `The user who last modified the file.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The time & date when the file was last updated.`,
				},
				resource.Attribute{
					Name:        "mimetype",
					Description: `The mimetype of the file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the file.`,
				},
				resource.Attribute{
					Name:        "download_uri",
					Description: `The URI that can be used to download the file.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `MD5 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha1",
					Description: `SHA1 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `SHA256 checksum of the file.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_fileinfo",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Name of the repository where the file is stored.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path to the file within the repository. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The time & date when the file was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user who created the file.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The time & date when the file was last modified.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `The user who last modified the file.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The time & date when the file was last updated.`,
				},
				resource.Attribute{
					Name:        "mimetype",
					Description: `The mimetype of the file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the file.`,
				},
				resource.Attribute{
					Name:        "download_uri",
					Description: `The URI that can be used to download the file.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `MD5 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha1",
					Description: `SHA1 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `SHA256 checksum of the file.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `The time & date when the file was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user who created the file.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The time & date when the file was last modified.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `The user who last modified the file.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The time & date when the file was last updated.`,
				},
				resource.Attribute{
					Name:        "mimetype",
					Description: `The mimetype of the file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the file.`,
				},
				resource.Attribute{
					Name:        "download_uri",
					Description: `The URI that can be used to download the file.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `MD5 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha1",
					Description: `SHA1 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `SHA256 checksum of the file.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"artifactory_artifactory_file":     0,
		"artifactory_artifactory_fileinfo": 1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
