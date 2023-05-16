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
					Description: `(Required) Namespace (or organization) to which this entitlement belongs.`,
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
					Description: `The literal value of the token to be created. ## Import This resource can be imported using the organization slug, the repository slug, and the entitlement slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_entitlement.my_entitlement my-organization.my-repository.3nt1lem3nT ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The literal value of the token to be created. ## Import This resource can be imported using the organization slug, the repository slug, and the entitlement slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_entitlement.my_entitlement my-organization.my-repository.3nt1lem3nT ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_license_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Organization to which the policy belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the license policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the license policy.`,
				},
				resource.Attribute{
					Name:        "spdx_identifiers",
					Description: `(Required) The licenses to deny.`,
				},
				resource.Attribute{
					Name:        "on_violation_quarantine",
					Description: `(Optional) On violation of the license policy, quarantine violating packages.`,
				},
				resource.Attribute{
					Name:        "allow_unknown_licenses",
					Description: `(Optional) Allow unknown licenses within the policy. ## Import This resource can be imported using the organization slug. ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_license_policy.my_policy my-organization ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
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
					Name:        "contextual_auth_realm",
					Description: `(Optional) If checked, missing credentials for this repository where basic authentication is required shall present an enriched value in the 'WWW-Authenticate' header containing the namespace and repository. This can be useful for tooling such as SBT where the authentication realm is used to distinguish and disambiguate credentials.`,
				},
				resource.Attribute{
					Name:        "copy_own",
					Description: `(Optional) If checked, users can copy any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "copy_packages",
					Description: `(Optional) This defines the minimum level of privilege required for a user to copy packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific copy setting.`,
				},
				resource.Attribute{
					Name:        "default_privilege",
					Description: `(Optional) This defines the default level of privilege that all of your organization members have for this`,
				},
				resource.Attribute{
					Name:        "deleted_at",
					Description: `ISO 8601 timestamp at which the repository was deleted. repository. This does not include collaborators, but applies to any member of the org regardless of their own membership role (i.e. it applies to owners, managers and members). Be careful if setting this to admin, because any member will be able to change settings.`,
				},
				resource.Attribute{
					Name:        "delete_own",
					Description: `(Optional) If checked, users can delete any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "delete_packages",
					Description: `(Optional) This defines the minimum level of privilege required for a user to delete packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific delete setting.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the repository's purpose/contents.`,
				},
				resource.Attribute{
					Name:        "docker_refresh_tokens_enabled",
					Description: `(Optional) If checked, refresh tokens will be issued in addition to access tokens for Docker authentication. This allows unlimited extension of the lifetime of access tokens.`,
				},
				resource.Attribute{
					Name:        "index_files",
					Description: `(Optional) If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted.`,
				},
				resource.Attribute{
					Name:        "move_own",
					Description: `(Optional) If checked, users can move any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "move_packages",
					Description: `(Optional) This defines the minimum level of privilege required for a user to move packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific move setting.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A descriptive name for the repository.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace (or organization) to which this repository belongs.`,
				},
				resource.Attribute{
					Name:        "proxy_npmjs",
					Description: `(Optional) If checked, Npm packages that are not in the repository when requested by clients will automatically be proxied from the public npmjs.org registry. If there is at least one version for a package, others will not be proxied.`,
				},
				resource.Attribute{
					Name:        "proxy_pypi",
					Description: `(Optional) If checked, Python packages that are not in the repository when requested by clients will automatically be proxied from the public pypi.python.org registry. If there is at least one version for a package, others will not be proxied.`,
				},
				resource.Attribute{
					Name:        "raw_package_index_enabled",
					Description: `(Optional) If checked, HTML and JSON indexes will be generated that list all available raw packages in the repository.`,
				},
				resource.Attribute{
					Name:        "raw_package_index_signatures_enabled",
					Description: `(Optional) If checked, the HTML and JSON indexes will display raw package GPG signatures alongside the index packages.`,
				},
				resource.Attribute{
					Name:        "replace_packages",
					Description: `(Optional) This defines the minimum level of privilege required for a user to republish packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific republish setting. Please note that the user still requires the privilege to delete packages that will be replaced by the new package; otherwise the republish will fail.`,
				},
				resource.Attribute{
					Name:        "replace_packages_by_default",
					Description: `(Optional) If checked, uploaded packages will overwrite/replace any others with the same attributes (e.g. same version) by default. This only applies if the user has the required privilege for the republishing AND has the required privilege to delete existing packages that they don't own.`,
				},
				resource.Attribute{
					Name:        "repository_type",
					Description: `(Optional) The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Public repositories are free to use on all plans and visible to all Cloudsmith users.`,
				},
				resource.Attribute{
					Name:        "resync_own",
					Description: `(Optional) If checked, users can resync any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "resync_packages",
					Description: `(Optional) This defines the minimum level of privilege required for a user to resync packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific resync setting.`,
				},
				resource.Attribute{
					Name:        "scan_own",
					Description: `(Optional) If checked, users can scan any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the 'Access Controls' section of the repository, and any inherited from the org.`,
				},
				resource.Attribute{
					Name:        "scan_packages",
					Description: `(Optional) This defines the minimum level of privilege required for a user to scan packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific scan setting.`,
				},
				resource.Attribute{
					Name:        "show_setup_all",
					Description: `(Optional) If checked, the Set Me Up help for all formats will always be shown, even if you don't have packages of that type uploaded. Otherwise, help will only be shown for packages that are in the repository. For example, if you have uploaded only NuGet packages, then the Set Me Up help for NuGet packages will be shown only.`,
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
					Name:        "strict_npm_validation",
					Description: `(Optional) If checked, npm packages will be validated strictly to ensure the package matches specifcation. You can turn this off if you have packages that are old or otherwise mildly off-spec, but we can't guarantee the packages will work with npm-cli or other tooling correctly. Turn off at your own risk!`,
				},
				resource.Attribute{
					Name:        "use_debian_labels",
					Description: `(Optional) If checked, a 'Label' field will be present in Debian-based repositories. It will contain a string that identifies the entitlement token used to authenticate the repository, in the form of 'source=t-'; or 'source=none' if no token was used. You can use this to help with pinning.`,
				},
				resource.Attribute{
					Name:        "use_default_cargo_upstream",
					Description: `(Optional) If checked, dependencies of uploaded Cargo crates which do not set an explicit value for \"registry\" will be assumed to be available from crates.io. If unchecked, dependencies with unspecified \"registry\" values will be assumed to be available in the registry being uploaded to. Uncheck this if you want to ensure that dependencies are only ever installed from Cloudsmith unless explicitly specified as belong to another registry.`,
				},
				resource.Attribute{
					Name:        "use_noarch_packages",
					Description: `(Optional) If checked, noarch packages (if supported) are enabled in installations/configurations. A noarch package is one that is not tied to specific system architecture (like i686).`,
				},
				resource.Attribute{
					Name:        "use_source_packages",
					Description: `(Optional) If checked, source packages (if supported) are enabled in installations/configurations. A source package is one that contains source code rather than built binaries.`,
				},
				resource.Attribute{
					Name:        "use_vulnerability_scanning",
					Description: `(Optional) If checked, vulnerability scanning will be enabled for all supported packages within this repository.`,
				},
				resource.Attribute{
					Name:        "user_entitlements_enabled",
					Description: `(Optional) If checked, users can use and manage their own user-specific entitlement token for the repository (if private). Otherwise, user-specific entitlements are disabled for all users.`,
				},
				resource.Attribute{
					Name:        "view_statistics",
					Description: `(Optional) This defines the minimum level of privilege required for a user to view repository statistics, to include entitlement-based usage, if applicable. If a user does not have the permission, they won't be able to view any statistics, either via the UI, API or CLI.`,
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
					Description: `This defines the minimum level of privilege required for a user to view repository statistics, to include entitlement-based usage, if applicable. If a user does not have the permission, they won't be able to view any statistics, either via the UI, API or CLI. ## Import This resource can be imported using the organization slug, and the repository slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_repository.my_repository my-organization.my-repository ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `This defines the minimum level of privilege required for a user to view repository statistics, to include entitlement-based usage, if applicable. If a user does not have the permission, they won't be able to view any statistics, either via the UI, API or CLI. ## Import This resource can be imported using the organization slug, and the repository slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_repository.my_repository my-organization.my-repository ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_repository_geo_ip_rules",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Organization to which the Repository belongs.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Repository to which these Geo/IP rules apply.`,
				},
				resource.Attribute{
					Name:        "cidr_allow",
					Description: `(Optional) The list of IP Addresses for which to allow access to the Repository, expressed in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "cidr_deny",
					Description: `(Optional) The list of IP Addresses for which to deny access to the Repository, expressed in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "country_code_allow",
					Description: `(Optional) The list of countries for which to allow access to the Repository, expressed in ISO 3166-1 country codes.`,
				},
				resource.Attribute{
					Name:        "country_code_deny",
					Description: `(Optional) The list of countries for which to deny access to the Repository, expressed in ISO 3166-1 country codes. ## Import This resource can be imported using the organization slug, and the repository slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_repository_geo_ip_rules.my_rules my-organization.my-repository ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_repository_privileges",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Organization to which this repository belongs.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Repository to which these privileges apply.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Variable number of blocks containing service accounts that should have repository privileges.`,
				},
				resource.Attribute{
					Name:        "privilege",
					Description: `(Required) The service's privilege level in the repository. Must be one of ` + "`" + `Admin` + "`" + `, ` + "`" + `Write` + "`" + `, or ` + "`" + `Read` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug/identifier of the service.`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Optional) Variable number of blocks containing teams that should have repository privileges.`,
				},
				resource.Attribute{
					Name:        "privilege",
					Description: `(Required) The team's privilege level in the repository. Must be one of ` + "`" + `Admin` + "`" + `, ` + "`" + `Write` + "`" + `, or ` + "`" + `Read` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug/identifier of the team.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) Variable number of blocks containing users that should have repository privileges.`,
				},
				resource.Attribute{
					Name:        "privilege",
					Description: `(Required) The user's privilege level in the repository. Must be one of ` + "`" + `Admin` + "`" + `, ` + "`" + `Write` + "`" + `, or ` + "`" + `Read` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug/identifier of the user. ## Import This resource can be imported using the organization slug, and the repository slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_repository_privileges.privs my-organization.my-repository ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_service",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the service's purpose.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A descriptive name for the service.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Organization to which this service belongs.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The service's role in the organization. If defined, must be one of ` + "`" + `Member` + "`" + ` or ` + "`" + `Manager` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Optional) Variable number of blocks containing team assignments for this service.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The service's role in the team. If defined, must be one of ` + "`" + `Member` + "`" + ` or ` + "`" + `Manager` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The team the service should be added to. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The service's API key.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the service in URIs or where a username is required. ## Import This resource can be imported using the organization slug, and the service slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_service.my_service my-organization.my-service ` + "`" + `` + "`" + `` + "`" + ` NOTE: It's not possible to retrieve a service's API key via the Cloudsmith API after creation, so when we import a service the key is unavailable. If the API key is needed for use within Terraform (to be passed to other resources) then the resource needs to be tainted and recreated (or otherwise created fresh within Terraform).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `The service's API key.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the service in URIs or where a username is required. ## Import This resource can be imported using the organization slug, and the service slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_service.my_service my-organization.my-service ` + "`" + `` + "`" + `` + "`" + ` NOTE: It's not possible to retrieve a service's API key via the Cloudsmith API after creation, so when we import a service the key is unavailable. If the API key is needed for use within Terraform (to be passed to other resources) then the resource needs to be tainted and recreated (or otherwise created fresh within Terraform).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_team",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the team's purpose.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A descriptive name for the team.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Organization to which this team belongs.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Optional) The slug identifies the team in URIs.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) Controls if the team is visible or hidden from non-members. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the team. It will never change once a team has been created. ## Import This resource can be imported using the organization slug, and the team slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_team.my_team my-organization.my-team ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the team. It will never change once a team has been created. ## Import This resource can be imported using the organization slug, and the team slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_team.my_team my-organization.my-team ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_vulnerability_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Organization to which the policy belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the vulnerability policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the vulnerability policy.`,
				},
				resource.Attribute{
					Name:        "min_severity",
					Description: `(Optional) The minimum severity level where a policy violation will be flagged.`,
				},
				resource.Attribute{
					Name:        "on_violation_quarantine",
					Description: `(Optional) On violation of the vulnerability policy, quarantine violating packages.`,
				},
				resource.Attribute{
					Name:        "allow_unknown_severity",
					Description: `(Optional) Allow an unknown severity level. ## Import This resource can be imported using the organization slug and the vulnerability policy slug_perm. ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_vulnerability_policy.my_policy my-organization ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_webhook",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "events",
					Description: `(Required) List of events for which this webhook will be fired.`,
				},
				resource.Attribute{
					Name:        "is_active",
					Description: `(Optional) If enabled, the webhook will trigger on subscribed events and send payloads to the configured target URL.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace (or organization) to which this webhook belongs.`,
				},
				resource.Attribute{
					Name:        "package_query",
					Description: `(Optional) The package-based search query for webhooks to fire. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. If a package does not match, the webhook will not fire.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Repository to which this webhook belongs.`,
				},
				resource.Attribute{
					Name:        "request_body_format",
					Description: `(Optional) The format of the payloads for webhook requests.`,
				},
				resource.Attribute{
					Name:        "request_body_template_format",
					Description: `(Optional) The format of the payloads for webhook requests.`,
				},
				resource.Attribute{
					Name:        "request_content_type",
					Description: `(Optional) The value that will be sent for the 'Content Type' header.`,
				},
				resource.Attribute{
					Name:        "secret_header",
					Description: `(Optional) The header to send the predefined secret in. This must be unique from existing headers or it won't be sent. You can use this as a form of authentication on the endpoint side.`,
				},
				resource.Attribute{
					Name:        "secret_value",
					Description: `(Optional) The value for the predefined secret (note: this is treated as a passphrase and is encrypted when we store it). You can use this as a form of authentication on the endpoint side.`,
				},
				resource.Attribute{
					Name:        "signature_key",
					Description: `(Optional) The value for the signature key - This is used to generate an HMAC-based hex digest of the request body, which we send as the X-Cloudsmith-Signature header so that you can ensure that the request wasn't modified by a malicious party (note: this is treated as a passphrase and is encrypted when we store it).`,
				},
				resource.Attribute{
					Name:        "target_url",
					Description: `(Required) The destination URL that webhook payloads will be POST'ed to.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) Variable number of blocks containing templates used to render webhook content before sending.`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `(Required) The event for which this template will be applied.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The contents of the template to be rendered.`,
				},
				resource.Attribute{
					Name:        "is_active",
					Description: `(Optional) If enabled, SSL certificates is verified when webhooks are sent. It's recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO 8601 timestamp at which the webhook was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user/account that created the webhook.`,
				},
				resource.Attribute{
					Name:        "disable_reason",
					Description: `Why this webhook has been disabled.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the webhook. It will never change once a webhook has been created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `ISO 8601 timestamp at which the webhook was updated.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user/account that updated the webhook. ## Import This resource can be imported using the organization slug, the repository slug, and the webhook slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_webhook.my_webhook my-organization.my-repository.w3bh0okS1uG ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO 8601 timestamp at which the webhook was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user/account that created the webhook.`,
				},
				resource.Attribute{
					Name:        "disable_reason",
					Description: `Why this webhook has been disabled.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the webhook. It will never change once a webhook has been created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `ISO 8601 timestamp at which the webhook was updated.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user/account that updated the webhook. ## Import This resource can be imported using the organization slug, the repository slug, and the webhook slug: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudsmith_webhook.my_webhook my-organization.my-repository.w3bh0okS1uG ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"cloudsmith_entitlement":             0,
		"cloudsmith_license_policy":          1,
		"cloudsmith_repository":              2,
		"cloudsmith_repository_geo_ip_rules": 3,
		"cloudsmith_repository_privileges":   4,
		"cloudsmith_service":                 5,
		"cloudsmith_team":                    6,
		"cloudsmith_vulnerability_policy":    7,
		"cloudsmith_webhook":                 8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
