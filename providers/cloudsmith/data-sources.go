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
			Type:             "cloudsmith_organization",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug identifies the organization in URIs. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `Country in which the organization is based.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO 8601 timestamp at which the organization was created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The city/town/area in which the organization is based.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the organization.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the organization in URIs.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the organization. It will never change once a organization has been created.`,
				},
				resource.Attribute{
					Name:        "tagline",
					Description: `A short public description for the organization.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "country",
					Description: `Country in which the organization is based.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO 8601 timestamp at which the organization was created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The city/town/area in which the organization is based.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the organization.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the organization in URIs.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the organization. It will never change once a organization has been created.`,
				},
				resource.Attribute{
					Name:        "tagline",
					Description: `A short public description for the organization.`,
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
					Description: `(Required) Namespace (or organization) to which the repository belongs.`,
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
					Name:        "contextual_auth_realm",
					Description: `If checked, missing credentials for this repository where basic authentication is required shall present an enriched value in the 'WWW-Authenticate' header containing the namespace and repository. This can be useful for tooling such as SBT where the authentication realm is used to distinguish and disambiguate credentials.`,
				},
				resource.Attribute{
					Name:        "copy_own",
					Description: `If checked, users can copy any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "copy_packages",
					Description: `This defines the minimum level of privilege required for a user to copy packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific copy setting.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO 8601 timestamp at which the repository was created.`,
				},
				resource.Attribute{
					Name:        "default_privilege",
					Description: `This defines the default level of privilege that all of your organization members have for this`,
				},
				resource.Attribute{
					Name:        "deleted_at",
					Description: `ISO 8601 timestamp at which the repository was deleted. repository. This does not include collaborators, but applies to any member of the org regardless of their own membership role (i.e. it applies to owners, managers and members). Be careful if setting this to admin, because any member will be able to change settings.`,
				},
				resource.Attribute{
					Name:        "delete_own",
					Description: `If checked, users can delete any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "delete_packages",
					Description: `This defines the minimum level of privilege required for a user to delete packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific delete setting.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the repository's purpose/contents.`,
				},
				resource.Attribute{
					Name:        "docker_refresh_tokens_enabled",
					Description: `If checked, refresh tokens will be issued in addition to access tokens for Docker authentication. This allows unlimited extension of the lifetime of access tokens.`,
				},
				resource.Attribute{
					Name:        "index_files",
					Description: `When ` + "`" + `true` + "`" + `, package indexing is enabled for this repository.`,
				},
				resource.Attribute{
					Name:        "is_open_source",
					Description: `True if this repository is open source.`,
				},
				resource.Attribute{
					Name:        "is_private",
					Description: `True if this repository is private.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `True if this repository is public.`,
				},
				resource.Attribute{
					Name:        "move_own",
					Description: `If checked, users can move any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "move_packages",
					Description: `This defines the minimum level of privilege required for a user to move packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific move setting.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the repository.`,
				},
				resource.Attribute{
					Name:        "namespace_url",
					Description: `API endpoint to where data about this namespace can be retrieved.`,
				},
				resource.Attribute{
					Name:        "proxy_npmjs",
					Description: `If checked, Npm packages that are not in the repository when requested by clients will automatically be proxied from the public npmjs.org registry. If there is at least one version for a package, others will not be proxied.`,
				},
				resource.Attribute{
					Name:        "proxy_pypi",
					Description: `If checked, Python packages that are not in the repository when requested by clients will automatically be proxied from the public pypi.python.org registry. If there is at least one version for a package, others will not be proxied.`,
				},
				resource.Attribute{
					Name:        "raw_package_index_enabled",
					Description: `If checked, HTML and JSON indexes will be generated that list all available raw packages in the repository.`,
				},
				resource.Attribute{
					Name:        "raw_package_index_signatures_enabled",
					Description: `If checked, the HTML and JSON indexes will display raw package GPG signatures alongside the index packages.`,
				},
				resource.Attribute{
					Name:        "replace_packages",
					Description: `This defines the minimum level of privilege required for a user to republish packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific republish setting. Please note that the user still requires the privilege to delete packages that will be replaced by the new package; otherwise the republish will fail.`,
				},
				resource.Attribute{
					Name:        "replace_packages_by_default",
					Description: `If checked, uploaded packages will overwrite/replace any others with the same attributes (e.g. same version) by default. This only applies if the user has the required privilege for the republishing AND has the required privilege to delete existing packages that they don't own.`,
				},
				resource.Attribute{
					Name:        "repository_type",
					Description: `The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Public repositories are free to use on all plans and visible to all Cloudsmith users.`,
				},
				resource.Attribute{
					Name:        "resync_own",
					Description: `If checked, users can resync any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "resync_packages",
					Description: `This defines the minimum level of privilege required for a user to resync packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific resync setting.`,
				},
				resource.Attribute{
					Name:        "scan_own",
					Description: `If checked, users can scan any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "scan_packages",
					Description: `This defines the minimum level of privilege required for a user to scan packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific scan setting.`,
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
					Name:        "show_setup_all",
					Description: `If checked, the Set Me Up help for all formats will always be shown, even if you don't have packages of that type uploaded. Otherwise, help will only be shown for packages that are in the repository. For example, if you have uploaded only NuGet packages, then the Set Me Up help for NuGet packages will be shown only.`,
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
				resource.Attribute{
					Name:        "strict_npm_validation",
					Description: `If checked, npm packages will be validated strictly to ensure the package matches specifcation. You can turn this off if you have packages that are old or otherwise mildly off-spec, but we can't guarantee the packages will work with npm-cli or other tooling correctly. Turn off at your own risk!`,
				},
				resource.Attribute{
					Name:        "use_debian_labels",
					Description: `If checked, a 'Label' field will be present in Debian-based repositories. It will contain a string that identifies the entitlement token used to authenticate the repository, in the form of 'source=t-'; or 'source=none' if no token was used. You can use this to help with pinning.`,
				},
				resource.Attribute{
					Name:        "use_default_cargo_upstream",
					Description: `If checked, dependencies of uploaded Cargo crates which do not set an explicit value for \"registry\" will be assumed to be available from crates.io. If unchecked, dependencies with unspecified \"registry\" values will be assumed to be available in the registry being uploaded to. Uncheck this if you want to ensure that dependencies are only ever installed from Cloudsmith unless explicitly specified as belong to another registry.`,
				},
				resource.Attribute{
					Name:        "use_noarch_packages",
					Description: `If checked, noarch packages (if supported) are enabled in installations/configurations. A noarch package is one that is not tied to specific system architecture (like i686).`,
				},
				resource.Attribute{
					Name:        "use_source_packages",
					Description: `If checked, source packages (if supported) are enabled in installations/configurations. A source package is one that contains source code rather than built binaries.`,
				},
				resource.Attribute{
					Name:        "use_vulnerability_scanning",
					Description: `If checked, vulnerability scanning will be enabled for all supported packages within this repository.`,
				},
				resource.Attribute{
					Name:        "user_entitlements_enabled",
					Description: `If checked, users can use and manage their own user-specific entitlement token for the repository (if private). Otherwise, user-specific entitlements are disabled for all users.`,
				},
				resource.Attribute{
					Name:        "view_statistics",
					Description: `This defines the minimum level of privilege required for a user to view repository statistics, to include entitlement-based usage, if applicable. If a user does not have the permission, they won't be able to view any statistics, either via the UI, API or CLI.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cdn_url",
					Description: `Base URL from which packages and other artifacts are downloaded.`,
				},
				resource.Attribute{
					Name:        "contextual_auth_realm",
					Description: `If checked, missing credentials for this repository where basic authentication is required shall present an enriched value in the 'WWW-Authenticate' header containing the namespace and repository. This can be useful for tooling such as SBT where the authentication realm is used to distinguish and disambiguate credentials.`,
				},
				resource.Attribute{
					Name:        "copy_own",
					Description: `If checked, users can copy any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "copy_packages",
					Description: `This defines the minimum level of privilege required for a user to copy packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific copy setting.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO 8601 timestamp at which the repository was created.`,
				},
				resource.Attribute{
					Name:        "default_privilege",
					Description: `This defines the default level of privilege that all of your organization members have for this`,
				},
				resource.Attribute{
					Name:        "deleted_at",
					Description: `ISO 8601 timestamp at which the repository was deleted. repository. This does not include collaborators, but applies to any member of the org regardless of their own membership role (i.e. it applies to owners, managers and members). Be careful if setting this to admin, because any member will be able to change settings.`,
				},
				resource.Attribute{
					Name:        "delete_own",
					Description: `If checked, users can delete any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "delete_packages",
					Description: `This defines the minimum level of privilege required for a user to delete packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific delete setting.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the repository's purpose/contents.`,
				},
				resource.Attribute{
					Name:        "docker_refresh_tokens_enabled",
					Description: `If checked, refresh tokens will be issued in addition to access tokens for Docker authentication. This allows unlimited extension of the lifetime of access tokens.`,
				},
				resource.Attribute{
					Name:        "index_files",
					Description: `When ` + "`" + `true` + "`" + `, package indexing is enabled for this repository.`,
				},
				resource.Attribute{
					Name:        "is_open_source",
					Description: `True if this repository is open source.`,
				},
				resource.Attribute{
					Name:        "is_private",
					Description: `True if this repository is private.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `True if this repository is public.`,
				},
				resource.Attribute{
					Name:        "move_own",
					Description: `If checked, users can move any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "move_packages",
					Description: `This defines the minimum level of privilege required for a user to move packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific move setting.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the repository.`,
				},
				resource.Attribute{
					Name:        "namespace_url",
					Description: `API endpoint to where data about this namespace can be retrieved.`,
				},
				resource.Attribute{
					Name:        "proxy_npmjs",
					Description: `If checked, Npm packages that are not in the repository when requested by clients will automatically be proxied from the public npmjs.org registry. If there is at least one version for a package, others will not be proxied.`,
				},
				resource.Attribute{
					Name:        "proxy_pypi",
					Description: `If checked, Python packages that are not in the repository when requested by clients will automatically be proxied from the public pypi.python.org registry. If there is at least one version for a package, others will not be proxied.`,
				},
				resource.Attribute{
					Name:        "raw_package_index_enabled",
					Description: `If checked, HTML and JSON indexes will be generated that list all available raw packages in the repository.`,
				},
				resource.Attribute{
					Name:        "raw_package_index_signatures_enabled",
					Description: `If checked, the HTML and JSON indexes will display raw package GPG signatures alongside the index packages.`,
				},
				resource.Attribute{
					Name:        "replace_packages",
					Description: `This defines the minimum level of privilege required for a user to republish packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific republish setting. Please note that the user still requires the privilege to delete packages that will be replaced by the new package; otherwise the republish will fail.`,
				},
				resource.Attribute{
					Name:        "replace_packages_by_default",
					Description: `If checked, uploaded packages will overwrite/replace any others with the same attributes (e.g. same version) by default. This only applies if the user has the required privilege for the republishing AND has the required privilege to delete existing packages that they don't own.`,
				},
				resource.Attribute{
					Name:        "repository_type",
					Description: `The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Public repositories are free to use on all plans and visible to all Cloudsmith users.`,
				},
				resource.Attribute{
					Name:        "resync_own",
					Description: `If checked, users can resync any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "resync_packages",
					Description: `This defines the minimum level of privilege required for a user to resync packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific resync setting.`,
				},
				resource.Attribute{
					Name:        "scan_own",
					Description: `If checked, users can scan any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "scan_packages",
					Description: `This defines the minimum level of privilege required for a user to scan packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific scan setting.`,
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
					Name:        "show_setup_all",
					Description: `If checked, the Set Me Up help for all formats will always be shown, even if you don't have packages of that type uploaded. Otherwise, help will only be shown for packages that are in the repository. For example, if you have uploaded only NuGet packages, then the Set Me Up help for NuGet packages will be shown only.`,
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
				resource.Attribute{
					Name:        "strict_npm_validation",
					Description: `If checked, npm packages will be validated strictly to ensure the package matches specifcation. You can turn this off if you have packages that are old or otherwise mildly off-spec, but we can't guarantee the packages will work with npm-cli or other tooling correctly. Turn off at your own risk!`,
				},
				resource.Attribute{
					Name:        "use_debian_labels",
					Description: `If checked, a 'Label' field will be present in Debian-based repositories. It will contain a string that identifies the entitlement token used to authenticate the repository, in the form of 'source=t-'; or 'source=none' if no token was used. You can use this to help with pinning.`,
				},
				resource.Attribute{
					Name:        "use_default_cargo_upstream",
					Description: `If checked, dependencies of uploaded Cargo crates which do not set an explicit value for \"registry\" will be assumed to be available from crates.io. If unchecked, dependencies with unspecified \"registry\" values will be assumed to be available in the registry being uploaded to. Uncheck this if you want to ensure that dependencies are only ever installed from Cloudsmith unless explicitly specified as belong to another registry.`,
				},
				resource.Attribute{
					Name:        "use_noarch_packages",
					Description: `If checked, noarch packages (if supported) are enabled in installations/configurations. A noarch package is one that is not tied to specific system architecture (like i686).`,
				},
				resource.Attribute{
					Name:        "use_source_packages",
					Description: `If checked, source packages (if supported) are enabled in installations/configurations. A source package is one that contains source code rather than built binaries.`,
				},
				resource.Attribute{
					Name:        "use_vulnerability_scanning",
					Description: `If checked, vulnerability scanning will be enabled for all supported packages within this repository.`,
				},
				resource.Attribute{
					Name:        "user_entitlements_enabled",
					Description: `If checked, users can use and manage their own user-specific entitlement token for the repository (if private). Otherwise, user-specific entitlements are disabled for all users.`,
				},
				resource.Attribute{
					Name:        "view_statistics",
					Description: `This defines the minimum level of privilege required for a user to view repository statistics, to include entitlement-based usage, if applicable. If a user does not have the permission, they won't be able to view any statistics, either via the UI, API or CLI.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"cloudsmith_namespace":    0,
		"cloudsmith_organization": 1,
		"cloudsmith_package_list": 2,
		"cloudsmith_repository":   3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
