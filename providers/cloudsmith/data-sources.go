package cloudsmith

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_namespace",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug identifies the namespace in URIs. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the namespace.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the namespace in URIs.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the namespace. It will never change once a namespace has been created.`,
				},
				resource.Attribute{
					Name:        "type_name",
					Description: `Is this a user or an organization namespace?.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the namespace.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the namespace in URIs.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the namespace. It will never change once a namespace has been created.`,
				},
				resource.Attribute{
					Name:        "type_name",
					Description: `Is this a user or an organization namespace?.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_package_list",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace to which the packages belong.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Repository ` + "`" + `slug_perm` + "`" + ` to which the packages belong.`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `(Optional) A list of Cloudsmith search filters (e.g ` + "`" + `format:docker` + "`" + `, ` + "`" + `name:^foo` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, only the most recent package resolved will be returned. ## Attribute Reference All of the argument attributes are also exported as result attributes. The following attribute is additionally exported:`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `A list of ` + "`" + `package` + "`" + ` entries as discovered by the data source.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "packages",
					Description: `A list of ` + "`" + `package` + "`" + ` entries as discovered by the data source.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_repository",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace to which the repository belongs.`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `(Required) An identifier used to resolve this repository. This can be the repository ` + "`" + `slug` + "`" + `, or ` + "`" + `slug_perm` + "`" + `. ## Attribute Reference All of the argument attributes are also exported as result attributes. Additionally, the following attributes are also exported:`,
				},
				resource.Attribute{
					Name:        "cdn_url",
					Description: `Base URL from which packages and other artifacts are downloaded.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO 8601 timestamp at which the repository was created.`,
				},
				resource.Attribute{
					Name:        "deleted_at",
					Description: `ISO 8601 timestamp at which the repository was deleted.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the repository's purpose/contents.`,
				},
				resource.Attribute{
					Name:        "index_files",
					Description: `When ` + "`" + `true` + "`" + `, package indexing is enabled for this repository.`,
				},
				resource.Attribute{
					Name:        "namespace_url",
					Description: `API endpoint to where data about this namespace can be retrieved.`,
				},
				resource.Attribute{
					Name:        "repository_type",
					Description: `A string describing the type of repository (Private, Public, Open-Source)`,
				},
				resource.Attribute{
					Name:        "self_html_url",
					Description: `The Cloudsmith web URL for this repository.`,
				},
				resource.Attribute{
					Name:        "self_url",
					Description: `The Cloudsmith API endpoint for this repository.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the repository in URIs.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The internal immutable identifier for this repository.`,
				},
				resource.Attribute{
					Name:        "storage_region",
					Description: `The Cloudsmith region in which package files are stored.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cdn_url",
					Description: `Base URL from which packages and other artifacts are downloaded.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO 8601 timestamp at which the repository was created.`,
				},
				resource.Attribute{
					Name:        "deleted_at",
					Description: `ISO 8601 timestamp at which the repository was deleted.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the repository's purpose/contents.`,
				},
				resource.Attribute{
					Name:        "index_files",
					Description: `When ` + "`" + `true` + "`" + `, package indexing is enabled for this repository.`,
				},
				resource.Attribute{
					Name:        "namespace_url",
					Description: `API endpoint to where data about this namespace can be retrieved.`,
				},
				resource.Attribute{
					Name:        "repository_type",
					Description: `A string describing the type of repository (Private, Public, Open-Source)`,
				},
				resource.Attribute{
					Name:        "self_html_url",
					Description: `The Cloudsmith web URL for this repository.`,
				},
				resource.Attribute{
					Name:        "self_url",
					Description: `The Cloudsmith API endpoint for this repository.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the repository in URIs.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The internal immutable identifier for this repository.`,
				},
				resource.Attribute{
					Name:        "storage_region",
					Description: `The Cloudsmith region in which package files are stored.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"cloudsmith_namespace":    0,
		"cloudsmith_package_list": 1,
		"cloudsmith_repository":   2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
