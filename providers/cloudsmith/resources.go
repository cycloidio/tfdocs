package cloudsmith

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_entitlement",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "is_active",
					Description: `(Optional) If enabled, the token will allow downloads based on configured restrictions (if any).`,
				},
				resource.Attribute{
					Name:        "limit_date_range_from",
					Description: `(Optional) The starting date/time the token is allowed to be used from.`,
				},
				resource.Attribute{
					Name:        "limit_date_range_to",
					Description: `(Optional) The ending date/time the token is allowed to be used until.`,
				},
				resource.Attribute{
					Name:        "limit_num_clients",
					Description: `(Optional) The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.`,
				},
				resource.Attribute{
					Name:        "limit_num_downloads",
					Description: `(Optional) The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.`,
				},
				resource.Attribute{
					Name:        "limit_package_query",
					Description: `(Optional) The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata.`,
				},
				resource.Attribute{
					Name:        "limit_path_query",
					Description: `(Optional) The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A descriptive name for the entitlement.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace to which this entitlement belongs.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Repository to which this entitlement belongs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) The literal value of the token to be created. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "is_active",
					Description: `If enabled, the token will allow downloads based on configured restrictions (if any).`,
				},
				resource.Attribute{
					Name:        "limit_date_range_from",
					Description: `The starting date/time the token is allowed to be used from.`,
				},
				resource.Attribute{
					Name:        "limit_date_range_to",
					Description: `The ending date/time the token is allowed to be used until.`,
				},
				resource.Attribute{
					Name:        "limit_num_clients",
					Description: `The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.`,
				},
				resource.Attribute{
					Name:        "limit_num_downloads",
					Description: `The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.`,
				},
				resource.Attribute{
					Name:        "limit_package_query",
					Description: `The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata.`,
				},
				resource.Attribute{
					Name:        "limit_path_query",
					Description: `The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the entitlement.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace to which this entitlement belongs.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `Repository to which this entitlement belongs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The literal value of the token to be created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "is_active",
					Description: `If enabled, the token will allow downloads based on configured restrictions (if any).`,
				},
				resource.Attribute{
					Name:        "limit_date_range_from",
					Description: `The starting date/time the token is allowed to be used from.`,
				},
				resource.Attribute{
					Name:        "limit_date_range_to",
					Description: `The ending date/time the token is allowed to be used until.`,
				},
				resource.Attribute{
					Name:        "limit_num_clients",
					Description: `The maximum number of unique clients allowed for the token. Please note that since clients are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.`,
				},
				resource.Attribute{
					Name:        "limit_num_downloads",
					Description: `The maximum number of downloads allowed for the token. Please note that since downloads are calculated asynchronously (after the download happens), the limit may not be imposed immediately but at a later point.`,
				},
				resource.Attribute{
					Name:        "limit_package_query",
					Description: `The package-based search query to apply to restrict downloads to. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. This will still allow access to non-package files, such as metadata.`,
				},
				resource.Attribute{
					Name:        "limit_path_query",
					Description: `The path-based search query to apply to restrict downloads to. This supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. The path evaluated does not include the domain name, the namespace, the entitlement code used, the package format, etc. and it always starts with a forward slash.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the entitlement.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace to which this entitlement belongs.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `Repository to which this entitlement belongs.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The literal value of the token to be created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_repository",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the repository's purpose/contents.`,
				},
				resource.Attribute{
					Name:        "index_files",
					Description: `(Optional) If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A descriptive name for the repository.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace to which this repository belongs.`,
				},
				resource.Attribute{
					Name:        "repository_type",
					Description: `(Optional) The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Public repositories are free to use on all plans and visible to all Cloudsmith users.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Optional) The slug identifies the repository in URIs.`,
				},
				resource.Attribute{
					Name:        "storage_region",
					Description: `(Optional) The Cloudsmith region in which package files are stored.`,
				},
				resource.Attribute{
					Name:        "wait_for_deletion",
					Description: `(Optional) If true, terraform will wait for a repository to be permanently deleted before finishing. ## Attribute Reference`,
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
					Description: `ISO 8601 timestamp at which the repository was deleted (repositories are soft deleted temporarily to allow cancelling).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the repository's purpose/contents.`,
				},
				resource.Attribute{
					Name:        "index_files",
					Description: `If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the repository.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace to which this repository belongs.`,
				},
				resource.Attribute{
					Name:        "namespace_url",
					Description: `API endpoint where data about this namespace can be retrieved.`,
				},
				resource.Attribute{
					Name:        "repository_type",
					Description: `The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Public repositories are free to use on all plans and visible to all Cloudsmith users.`,
				},
				resource.Attribute{
					Name:        "self_html_url",
					Description: `Website URL for this repository.`,
				},
				resource.Attribute{
					Name:        "self_url",
					Description: `API endpoint where data about this repository can be retrieved.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the repository in URIs.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the repository. It will never change once a repository has been created.`,
				},
				resource.Attribute{
					Name:        "storage_region",
					Description: `The Cloudsmith region in which package files are stored.`,
				},
				resource.Attribute{
					Name:        "wait_for_deletion",
					Description: `If true, terraform will wait for a repository to be permanently deleted before finishing.`,
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
					Description: `ISO 8601 timestamp at which the repository was deleted (repositories are soft deleted temporarily to allow cancelling).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the repository's purpose/contents.`,
				},
				resource.Attribute{
					Name:        "index_files",
					Description: `If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the repository.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace to which this repository belongs.`,
				},
				resource.Attribute{
					Name:        "namespace_url",
					Description: `API endpoint where data about this namespace can be retrieved.`,
				},
				resource.Attribute{
					Name:        "repository_type",
					Description: `The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Public repositories are free to use on all plans and visible to all Cloudsmith users.`,
				},
				resource.Attribute{
					Name:        "self_html_url",
					Description: `Website URL for this repository.`,
				},
				resource.Attribute{
					Name:        "self_url",
					Description: `API endpoint where data about this repository can be retrieved.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the repository in URIs.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the repository. It will never change once a repository has been created.`,
				},
				resource.Attribute{
					Name:        "storage_region",
					Description: `The Cloudsmith region in which package files are stored.`,
				},
				resource.Attribute{
					Name:        "wait_for_deletion",
					Description: `If true, terraform will wait for a repository to be permanently deleted before finishing.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"cloudsmith_entitlement": 0,
		"cloudsmith_repository":  1,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
