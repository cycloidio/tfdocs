package artifactory

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "artifactory_access_token",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"access",
				"token",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username or subject for the token. A non-admin can only specify their own username. Admins can specify any existing username, or a new name for a temporary token. Temporary tokens require ` + "`" + `groups` + "`" + ` to be set.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The end date which the token is valid until, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the token is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Valid time units are "s", "m", "h".`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of groups. The token is granted access based on the permissions of the groups. Specify ` + "`" + `["`,
				},
				resource.Attribute{
					Name:        "admin_token",
					Description: `(Optional) Specify the ` + "`" + `instance_id` + "`" + ` in this block to grant this token admin privileges. This can only be created when the authenticated user is an admin. ` + "`" + `admin_token` + "`" + ` cannot be specified with ` + "`" + `groups` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "refreshable",
					Description: `(Optional) Is this token refreshable? Defaults to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `(Optional) A space-separate list of the other Artifactory instances or services that should accept this token identified by their Artifactory Service IDs. You may set ` + "`" + `"jfrt@`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `Returns the access token to authenciate to Artifactory`,
				},
				resource.Attribute{
					Name:        "refresh_token",
					Description: `Returns the refresh token when ` + "`" + `refreshable` + "`" + ` is true, or an empty string when ` + "`" + `refreshable` + "`" + ` is false. ## References - https://www.jfrog.com/confluence/display/ACC1X/Access+Tokens - https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateToken ## Import Artifactory`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_anonymous_user",
			Category:         "User",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"user",
				"anonymous",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Username for anonymous user. This is only for ensuring resource schema is valid for Terraform. This is not meant to be set or updated in the HCL. ## Import Anonymous user can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_anonymous_user.anonymous-user anonymous ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_api_key",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"api",
				"key",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_key",
					Description: `The API key. ## Import A user's API key can be imported using any identifier, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_api_key.test import ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifact_property_webhook",
			Category:         "Webhook",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"webhook",
				"artifact",
				"property",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Webhook description. Max length 1000 characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Status of webhook. Default to 'true'.`,
				},
				resource.Attribute{
					Name:        "event_types",
					Description: `(Required) List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "added", "deleted".`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Required) Specifies where the webhook will be applied on which repositories.`,
				},
				resource.Attribute{
					Name:        "any_local",
					Description: `(Required) Trigger on any local repo.`,
				},
				resource.Attribute{
					Name:        "any_remote",
					Description: `(Required) Trigger on any remote repo.`,
				},
				resource.Attribute{
					Name:        "repo_keys",
					Description: `(Required) Trigger on this list of repo keys.`,
				},
				resource.Attribute{
					Name:        "include_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "exclude_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Secret authentication token that will be sent to the configured URL`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy key from Artifactory Proxies setting.`,
				},
				resource.Attribute{
					Name:        "custom_http_headers",
					Description: `(Optional) Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifact_webhook",
			Category:         "Webhook",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"webhook",
				"artifact",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Webhook description. Max length 1000 characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Status of webhook. Default to 'true'.`,
				},
				resource.Attribute{
					Name:        "event_types",
					Description: `(Required) List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "deployed", "deleted", "moved", "copied", "cached".`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Required) Specifies where the webhook will be applied on which repositories.`,
				},
				resource.Attribute{
					Name:        "any_local",
					Description: `(Required) Trigger on any local repo.`,
				},
				resource.Attribute{
					Name:        "any_remote",
					Description: `(Required) Trigger on any remote repo.`,
				},
				resource.Attribute{
					Name:        "repo_keys",
					Description: `(Required) Trigger on this list of repo keys.`,
				},
				resource.Attribute{
					Name:        "include_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "exclude_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Secret authentication token that will be sent to the configured URL.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy key from Artifactory Proxies setting.`,
				},
				resource.Attribute{
					Name:        "custom_http_headers",
					Description: `(Optional) Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_release_bundle_webhook",
			Category:         "Webhook",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"webhook",
				"release",
				"bundle",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Webhook description. Max length 1000 characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Status of webhook. Default to 'true'`,
				},
				resource.Attribute{
					Name:        "event_types",
					Description: `(Required) List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "received", "delete_started", "delete_completed", "delete_failed"`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Required) Specifies where the webhook will be applied on which repositories.`,
				},
				resource.Attribute{
					Name:        "any_release_bundle",
					Description: `(Required) Trigger on any release bundle`,
				},
				resource.Attribute{
					Name:        "registered_release_bundle_names",
					Description: `(Required) Trigger on this list of release bundle names`,
				},
				resource.Attribute{
					Name:        "include_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "exclude_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Secret authentication token that will be sent to the configured URL`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy key from Artifactory Proxies setting`,
				},
				resource.Attribute{
					Name:        "custom_http_headers",
					Description: `(Optional) Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_backup",
			Category:         "Configuration",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"configuration",
				"backup",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_build_webhook",
			Category:         "Webhook",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"webhook",
				"build",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Webhook description. Max length 1000 characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Status of webhook. Default to 'true'.`,
				},
				resource.Attribute{
					Name:        "event_types",
					Description: `(Required) List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "uploaded", "deleted", "promoted".`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Required) Specifies where the webhook will be applied on which repositories.`,
				},
				resource.Attribute{
					Name:        "any_build",
					Description: `(Required) Trigger on any build.`,
				},
				resource.Attribute{
					Name:        "selected_builds",
					Description: `(Required) Trigger on this list of build names.`,
				},
				resource.Attribute{
					Name:        "include_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "exclude_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Secret authentication token that will be sent to the configured URL.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy key from Artifactory Proxies setting.`,
				},
				resource.Attribute{
					Name:        "custom_http_headers",
					Description: `(Optional) Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_certificate",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alias",
					Description: `(Required) Name of certificate.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) PEM-encoded client certificate and private key. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `SHA256 fingerprint of the certificate.`,
				},
				resource.Attribute{
					Name:        "issued_by",
					Description: `Name of the certificate authority that issued the certificate.`,
				},
				resource.Attribute{
					Name:        "issued_on",
					Description: `The time & date when the certificate is valid from.`,
				},
				resource.Attribute{
					Name:        "issued_to",
					Description: `Name of whom the certificate has been issued to.`,
				},
				resource.Attribute{
					Name:        "valid_until",
					Description: `The time & date when the certificate expires. ## Import Certificates can be imported using their alias, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_certificate.my-cert my-cert ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fingerprint",
					Description: `SHA256 fingerprint of the certificate.`,
				},
				resource.Attribute{
					Name:        "issued_by",
					Description: `Name of the certificate authority that issued the certificate.`,
				},
				resource.Attribute{
					Name:        "issued_on",
					Description: `The time & date when the certificate is valid from.`,
				},
				resource.Attribute{
					Name:        "issued_to",
					Description: `Name of whom the certificate has been issued to.`,
				},
				resource.Attribute{
					Name:        "valid_until",
					Description: `The time & date when the certificate expires. ## Import Certificates can be imported using their alias, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_certificate.my-cert my-cert ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_distribution_webhook",
			Category:         "Webhook",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"webhook",
				"distribution",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Webhook description. Max length 1000 characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Status of webhook. Default to 'true'.`,
				},
				resource.Attribute{
					Name:        "event_types",
					Description: `(Required) List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "distribute_started", "distribute_completed", "distribute_aborted", "distribute_failed", "delete_started", "delete_completed", "delete_failed"`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Required) Specifies where the webhook will be applied on which repositories.`,
				},
				resource.Attribute{
					Name:        "any_release_bundle",
					Description: `(Required) Trigger on any release bundle.`,
				},
				resource.Attribute{
					Name:        "registered_release_bundle_names",
					Description: `(Required) Trigger on this list of release bundle names.`,
				},
				resource.Attribute{
					Name:        "include_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "exclude_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Secret authentication token that will be sent to the configured URL.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy key from Artifactory Proxies setting.`,
				},
				resource.Attribute{
					Name:        "custom_http_headers",
					Description: `(Optional) Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_docker_webhook",
			Category:         "Webhook",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"webhook",
				"docker",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Webhook description. Max length 1000 characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Status of webhook. Default to 'true'.`,
				},
				resource.Attribute{
					Name:        "event_types",
					Description: `(Required) List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "pushed", "deleted", "promoted".`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Required) Specifies where the webhook will be applied on which repositories.`,
				},
				resource.Attribute{
					Name:        "any_local",
					Description: `(Required) Trigger on any local repo.`,
				},
				resource.Attribute{
					Name:        "any_remote",
					Description: `(Required) Trigger on any remote repo.`,
				},
				resource.Attribute{
					Name:        "repo_keys",
					Description: `(Required) Trigger on this list of repo keys.`,
				},
				resource.Attribute{
					Name:        "include_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "exclude_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Secret authentication token that will be sent to the configured URL.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy key from Artifactory Proxies setting.`,
				},
				resource.Attribute{
					Name:        "custom_http_headers",
					Description: `(Optional) Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_alpine_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"alpine",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_alpine_repository.terraform-federated-test-alpine-repo terraform-federated-test-alpine-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_bower_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"bower",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_bower_repository.terraform-federated-test-bower-repo terraform-federated-test-bower-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_cargo_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"cargo",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_cargo_repository.terraform-federated-test-cargo-repo terraform-federated-test-cargo-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_chef_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"chef",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_chef_repository.terraform-federated-test-chef-repo terraform-federated-test-chef-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_cocoapods_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"cocoapods",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_cocoapods_repository.terraform-federated-test-cocoapods-repo terraform-federated-test-cocoapods-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_composer_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"composer",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_composer_repository.terraform-federated-test-composer-repo terraform-federated-test-composer-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_conan_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"conan",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import resource "artifactory_federated_conan_repository" "terraform-federated-test-conan-repo" { .terraform-federated-test-conan-repo terraform-federated-test-conan-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_conda_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"conda",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_conda_repository.terraform-federated-test-conda-repo terraform-federated-test-conda-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_cran_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"cran",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_cran_repository.terraform-federated-test-cran-repo terraform-federated-test-cran-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_debian_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"debian",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_debian_repository.terraform-federated-test-debian-repo terraform-federated-test-debian-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_docker_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"docker",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_docker_repository.terraform-federated-test-docker-repo terraform-federated-test-docker-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_gems_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"gems",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_gem_repository.terraform-federated-test-gem-repo terraform-federated-test-gem-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_generic_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"generic",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_generic_repository.terraform-federated-test-generic-repo terraform-federated-test-generic-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_gitlfs_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"gitlfs",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_gitlfs_repository.terraform-federated-test-gitlfs-repo terraform-federated-test-gitlfs-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_go_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"go",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_go_repository.terraform-federated-test-go-repo terraform-federated-test-go-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_gradle_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"gradle",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_gradle_repository.terraform-federated-test-gradle-repo terraform-federated-test-gradle-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_helm_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"helm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_helm_repository.terraform-federated-test-helm-repo terraform-federated-test-helm-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_ivy_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"ivy",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_ivy_repository.terraform-federated-test-ivy-repo terraform-federated-test-ivy-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_maven_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"maven",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_maven_repository.terraform-federated-test-maven-repo terraform-federated-test-maven-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_npm_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"npm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_npm_repository.terraform-federated-test-npm-repo terraform-federated-test-npm-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_nuget_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"nuget",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_nuget_repository.terraform-federated-test-nuget-repo terraform-federated-test-nuget-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_opkg_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"opkg",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_opkg_repository.terraform-federated-test-opkg-repo terraform-federated-test-opkg-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_puppet_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"puppet",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_puppet_repository.terraform-federated-test-puppet-repo terraform-federated-test-puppet-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_pypi_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"pypi",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_pypi_repository.terraform-federated-test-pypi-repo terraform-federated-test-pypi-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_rpm_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"rpm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_rpm_repository.terraform-federated-test-rpm-repo terraform-federated-test-rpm-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_sbt_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"sbt",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_sbt_repository.terraform-federated-test-sbt-repo terraform-federated-test-sbt-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_vagrant_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"vagrant",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically. ## Import Federated repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_federated_vagrant_repository.terraform-federated-test-vagrant-repo terraform-federated-test-vagrant-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_general_security",
			Category:         "Configuration",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"configuration",
				"general",
				"security",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable_anonoymous_access",
					Description: `(Optional) Enable anonymous access. Default value is ` + "`" + `false` + "`" + `. ## Import Current general security settings can be imported using ` + "`" + `security` + "`" + ` as the ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_general_security.security security ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_group",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_keypair",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"keypair",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pair_name",
					Description: `(Required) A unique identifier for the Key Pair record.`,
				},
				resource.Attribute{
					Name:        "pair_type",
					Description: `(Required) Key Pair type. Supported types - GPG and RSA.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `(Required) Will be used as a filename when retrieving the public key via REST API.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required, Sensitive) - Private key. PEM format will be validated.`,
				},
				resource.Attribute{
					Name:        "passphrase",
					Description: `(Optional) Passphrase will be used to decrypt the private key. Validated server side.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) Public key. PEM format will be validated.`,
				},
				resource.Attribute{
					Name:        "unavailable",
					Description: `(Computed) Unknown usage. Returned in the json payload and cannot be set. Artifactory REST API call Get Key Pair doesn't return keys ` + "`" + `private_key` + "`" + ` and ` + "`" + `passphrase` + "`" + `, but consumes these keys in the POST call. The meta-argument ` + "`" + `lifecycle` + "`" + ` used here to make Provider ignore the changes for these two keys in the Terraform state. ## Import Keypair can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_keypair.my-keypair my-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_ldap_group_setting",
			Category:         "Configuration",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"configuration",
				"ldap",
				"group",
				"setting",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_ldap_setting",
			Category:         "Configuration",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"configuration",
				"ldap",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_user_to_access_profile",
					Description: `(Optional) When set, users created after logging in using LDAP will be able to access their profile page. Default value is ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `(Optional) Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.`,
				},
				resource.Attribute{
					Name:        "project_environments",
					Description: `(Optional) Project environment for assigning this repository to. Allow values: "DEV" or "PROD"`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `(Optional) List of artifact patterns to include when evaluating artifact requests in the form of x/y/`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `(Optional) List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/`,
				},
				resource.Attribute{
					Name:        "repo_layout_ref",
					Description: `(Optional) Sets the layout that the repository should use for storing and identifying modules. A recommended layout that corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.`,
				},
				resource.Attribute{
					Name:        "blacked_out",
					Description: `(Optional, Default: false) When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.`,
				},
				resource.Attribute{
					Name:        "xray_index",
					Description: `(Optional, Default: false) Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.`,
				},
				resource.Attribute{
					Name:        "priority_resolution",
					Description: `(Optional, Default: false) Setting repositories with priority will cause metadata to be merged only from repositories set with this field`,
				},
				resource.Attribute{
					Name:        "property_sets",
					Description: `(Optional) List of property set name`,
				},
				resource.Attribute{
					Name:        "archive_browsing_enabled",
					Description: `(Optional) When set, you may view content such as HTML or Javadoc files directly from Artifactory.\nThis may not be safe and therefore requires strict content moderation to prevent malicious users from uploading content that may compromise security (e.g., cross-site scripting attacks).`,
				},
				resource.Attribute{
					Name:        "download_direct",
					Description: `(Optional) When set, download requests to this repository will redirect the client to download the artifact directly from the cloud storage provider. Available in Enterprise+ and Edge licenses only.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_alpine_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"alpine",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) The RSA key to be used to sign alpine indices. Artifactory REST API call Get Key Pair doesn't return keys ` + "`" + `private_key` + "`" + ` and ` + "`" + `passphrase` + "`" + `, but consumes these keys in the POST call. The meta-argument ` + "`" + `lifecycle` + "`" + ` used here to make Provider ignore the changes for these two keys in the Terraform state. ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_alpine_repository.terraform-local-test-alpine-repo-basic terraform-local-test-alpine-repo-basic ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_bower_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"bower",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_bower_repository.terraform-local-test-bower-repo terraform-local-test-bower-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_cargo_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"cargo",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "anonymous_access",
					Description: `(Optional) Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is ` + "`" + `false` + "`" + `. ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_cargo_repository.terraform-local-test-cargo-repo-basic terraform-local-test-cargo-repo-basic ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_chef_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"chef",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_chef_repository.terraform-local-test-chef-repo terraform-local-test-chef-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_cocoapods_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"cocoapods",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_cocoapods_repository.terraform-local-test-cocoapods-repo terraform-local-test-cocoapods-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_composer_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"composer",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_composer_repository.terraform-local-test-composer-repo terraform-local-test-composer-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_conan_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"conan",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_conan_repository.terraform-local-test-conan-repo terraform-local-test-conan-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_conda_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"conda",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_conda_repository.terraform-local-test-conda-repo terraform-local-test-conda-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_cran_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"cran",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_cran_repository.terraform-local-test-cran-repo terraform-local-test-cran-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_debian_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"debian",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) The primary RSA key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `(Optional) The secondary RSA key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "index_compression_formats",
					Description: `(Optional) The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension) and XZ (.xz extension).`,
				},
				resource.Attribute{
					Name:        "trivial_layout",
					Description: `(Optional) When set, the repository will use the deprecated trivial layout. Artifactory REST API call Get Key Pair doesn't return keys ` + "`" + `private_key` + "`" + ` and ` + "`" + `passphrase` + "`" + `, but consumes these keys in the POST call. The meta-argument ` + "`" + `lifecycle` + "`" + ` used here to make Provider ignore the changes for these two keys in the Terraform state. ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_debian_repository.my-debian-repo my-debian-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_docker_v1_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"docker",
				"v1",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_docker_v1_repository.foo foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_docker_v2_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"docker",
				"v2",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "block_pushing_schema1",
					Description: `(Optional) When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.`,
				},
				resource.Attribute{
					Name:        "tag_retention",
					Description: `(Optional) If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2.`,
				},
				resource.Attribute{
					Name:        "max_unique_tags",
					Description: `(Optional) The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only applies to manifest v2. ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_docker_v2_repository.foo foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_gems_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"gems",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_gems_repository.terraform-local-test-gems-repo terraform-local-test-gems-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_generic_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"generic",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters. ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_generic_repository.terraform-local-test-generic-repo terraform-local-test-generic-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_gitlfs_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"gitlfs",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_gitlfs_repository.terraform-local-test-gitlfs-repo terraform-local-test-gitlfs-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_go_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"go",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_go_repository.terraform-local-test-go-repo terraform-local-test-go-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_gradle_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"gradle",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "checksum_policy_type",
					Description: `(Optional) Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are ` + "`" + `client-checksums` + "`" + ` and ` + "`" + `generated-checksums` + "`" + `. For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).`,
				},
				resource.Attribute{
					Name:        "snapshot_version_behavior",
					Description: `(Optional) Specifies the naming convention for Maven SNAPSHOT versions. The options are -`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `(Optional) The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional) If set, Artifactory allows you to deploy release artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional) If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional) By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository. ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_gradle_repository.terraform-local-test-gradle-repo-basic terraform-local-test-gradle-repo-basic ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_helm_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"helm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_helm_repository.terraform-local-test-helm-repo terraform-local-test-helm-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_ivy_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"ivy",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_ivy_repository.terraform-local-test-ivy-repo terraform-local-test-ivy-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_maven_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"maven",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "checksum_policy_type",
					Description: `(Optional) Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are ` + "`" + `client-checksums` + "`" + ` and ` + "`" + `generated-checksums` + "`" + `. For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).`,
				},
				resource.Attribute{
					Name:        "snapshot_version_behavior",
					Description: `(Optional) Specifies the naming convention for Maven SNAPSHOT versions. The options are -`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `(Optional) The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional) If set, Artifactory allows you to deploy release artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional) If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional) By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository. ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_maven_repository.terraform-local-test-maven-repo-basic terraform-local-test-maven-repo-basic ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_npm_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"npm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_npm_repository.terraform-local-test-npm-repo terraform-local-test-npm-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_nuget_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"nuget",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `(Optional) The maximum number of unique snapshots of a single artifact to store Once the number of snapshots exceeds this setting, older versions are removed A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `(Optional) Force basic authentication credentials in order to use this repository. Default is ` + "`" + `false` + "`" + `. ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_nuget_repository.terraform-local-test-nuget-repo-basic terraform-local-test-nuget-repo-basic ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_opkg_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"opkg",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_opkg_repository.terraform-local-test-opkg-repo terraform-local-test-opkg-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_pub_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"pub",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_pub_repository.terraform-local-test-pub-repo terraform-local-test-pub-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_puppet_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"puppet",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_puppet_repository.terraform-local-test-puppet-repo terraform-local-test-puppet-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_pypi_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"pypi",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_pypi_repository.terraform-local-test-pypi-repo terraform-local-test-pypi-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_rpm_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"rpm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "yum_root_depth",
					Description: `(Optional) The depth, relative to the repository's root folder, where RPM metadata is created. This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "calculate_yum_metadata",
					Description: `(Optional) Default: false.`,
				},
				resource.Attribute{
					Name:        "enable_file_lists_indexing",
					Description: `(Optional) Default: false.`,
				},
				resource.Attribute{
					Name:        "yum_group_file_names",
					Description: `(Optional) A comma separated list of XML file names containing RPM group component definitions. Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically generating a gzipped version of the group files, if required. Default is empty string.`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) The primary GPG key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `(Optional) The secondary GPG key to be used to sign packages. ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_rpm_repository.terraform-local-test-rpm-repo-basic terraform-local-test-rpm-repo-basic ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_sbt_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"sbt",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_sbt_repository.terraform-local-test-sbt-repo terraform-local-test-sbt-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_vagrant_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"vagrant",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Local repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_local_vagrant_repository.terraform-local-test-vagrant-repo terraform-local-test-vagrant-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_managed_user",
			Category:         "User",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"user",
				"managed",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Username for user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email for user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for the user.`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `(Optional) When enabled, this user is an administrator with all the ensuing privileges. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "profile_updatable",
					Description: `(Optional) When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_ui_access",
					Description: `(Optional) When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internal_password_disabled",
					Description: `(Optional) When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of groups this user is a part of. ## Import Users can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_managed_user.test-user myusername ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_oauth_settings",
			Category:         "Configuration",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"configuration",
				"oauth",
				"settings",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_permission_target",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"permission",
				"target",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of permission.`,
				},
				resource.Attribute{
					Name:        "repo",
					Description: `(Optional) Repository permission configuration.`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `(Optional) Pattern of artifacts to include.`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `(Optional) Pattern of artifacts to exclude.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) List of repositories this permission target is applicable for.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: ``,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) Users this permission target applies for.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Groups this permission applies for.`,
				},
				resource.Attribute{
					Name:        "build",
					Description: `(Optional) As for repo but for artifactory-build-info permssions. ## Permissions The provider supports the following ` + "`" + `permission` + "`" + ` enums:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `matches ` + "`" + `Read` + "`" + ` permissions.`,
				},
				resource.Attribute{
					Name:        "write",
					Description: `matches ` + "`" + ` Deploy / Cache / Create` + "`" + ` permissions.`,
				},
				resource.Attribute{
					Name:        "annotate",
					Description: `matches ` + "`" + `Annotate` + "`" + ` permissions.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `matches ` + "`" + `Delete / Overwrite` + "`" + ` permissions.`,
				},
				resource.Attribute{
					Name:        "manage",
					Description: `matches ` + "`" + `Manage` + "`" + ` permissions. ## Import Permission targets can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_permission_target.terraform-test-permission mypermission ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_pull_replication",
			Category:         "Replication",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"replication",
				"pull",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repo_key",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "cron_exp",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "enable_event_replication",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) Required for local repository, but not needed for remote repository.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Required for local repository, but not needed for remote repository.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Required for local repository, but not needed for remote repository.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_deletes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_properties",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_statistics",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "path_prefix",
					Description: `(Optional) ## Import Pull replication config can be imported using its repo key, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_pull_replication.foo-rep repository-key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_push_replication",
			Category:         "Replication",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"replication",
				"push",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repo_key",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "cron_exp",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "enable_event_replication",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "replications",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "socket_timeout_millis",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_deletes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_properties",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_statistics",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "path_prefix",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy key from Artifactory Proxies setting. ## Import Push replication configs can be imported using their repo key, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_push_replication.foo-rep provider_test_source ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_release_bundle_webhook",
			Category:         "Webhook",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"webhook",
				"release",
				"bundle",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Webhook description. Max length 1000 characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Status of webhook. Default to 'true'.`,
				},
				resource.Attribute{
					Name:        "event_types",
					Description: `(Required) List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "created", "signed", "deleted".`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Required) Specifies where the webhook will be applied on which repositories.`,
				},
				resource.Attribute{
					Name:        "any_release_bundle",
					Description: `(Required) Trigger on any release bundle.`,
				},
				resource.Attribute{
					Name:        "registered_release_bundle_names",
					Description: `(Required) Trigger on this list of release bundle names.`,
				},
				resource.Attribute{
					Name:        "include_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "exclude_patterns",
					Description: `(Optional) Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Secret authentication token that will be sent to the configured URL.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy key from Artifactory Proxies setting.`,
				},
				resource.Attribute{
					Name:        "custom_http_headers",
					Description: `(Optional) Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `(Optional) Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.`,
				},
				resource.Attribute{
					Name:        "project_environments",
					Description: `(Optional) Project environment for assigning this repository to. Allow values: "DEV" or "PROD".`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy key from Artifactory Proxies settings.`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `(Optional) List of artifact patterns to include when evaluating artifact requests in the form of x/y/\`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `(Optional) List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/`,
				},
				resource.Attribute{
					Name:        "repo_layout_ref",
					Description: `(Optional) Sets the layout that the repository should use for storing and identifying modules. A recommended layout that corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.`,
				},
				resource.Attribute{
					Name:        "remote_repo_layout_ref",
					Description: `(Optional) Repository layout key for the remote layout mapping.`,
				},
				resource.Attribute{
					Name:        "hard_fail",
					Description: `(Optional) When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to communicate with this repository.`,
				},
				resource.Attribute{
					Name:        "offline",
					Description: `(Optional) If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.`,
				},
				resource.Attribute{
					Name:        "blacked_out",
					Description: `(Optional) (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact resolution.`,
				},
				resource.Attribute{
					Name:        "xray_index",
					Description: `(Optional, Default: false) Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.`,
				},
				resource.Attribute{
					Name:        "store_artifacts_locally",
					Description: `(Optional) When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory servers.`,
				},
				resource.Attribute{
					Name:        "socket_timeout_millis",
					Description: `(Optional) Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network operation is considered a retrieval failure.`,
				},
				resource.Attribute{
					Name:        "local_address",
					Description: `(Optional) The local address to be used when creating connections. Useful for specifying the interface to use on systems with multiple network interfaces.`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: 7200) The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.`,
				},
				resource.Attribute{
					Name:        "failed_retrieval_cache_period_secs",
					Description: `(Optional) This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function.`,
				},
				resource.Attribute{
					Name:        "missed_cache_period_seconds",
					Description: `(Optional) The number of seconds to cache artifact retrieval misses (artifact not found). A value of 0 indicates no caching.`,
				},
				resource.Attribute{
					Name:        "unused_artifacts_cleanup_period_enabled",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "unused_artifacts_cleanup_period_hours",
					Description: `(Optional) The number of hours to wait before an artifact is deemed "unused" and eligible for cleanup from the repository. A value of 0 means automatic cleanup of cached artifacts is disabled.`,
				},
				resource.Attribute{
					Name:        "assumed_offline_period_secs",
					Description: `(Optional, Default: 300) The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time, an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed offline. Default to 300.`,
				},
				resource.Attribute{
					Name:        "share_configuration",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "synchronize_properties",
					Description: `(Optional) When set, remote artifacts are fetched along with their properties.`,
				},
				resource.Attribute{
					Name:        "block_mismatching_mime_types",
					Description: `(Optional) Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources, HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked, Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.`,
				},
				resource.Attribute{
					Name:        "mismatching_mime_types_override_list",
					Description: `(Optional) The set of mime types that should override the block_mismatching_mime_types setting. Eg: "application/json,application/xml". Default value is empty.`,
				},
				resource.Attribute{
					Name:        "property_sets",
					Description: `(Optional) List of property set names.`,
				},
				resource.Attribute{
					Name:        "allow_any_host_auth",
					Description: `(Optional) Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to any other host.`,
				},
				resource.Attribute{
					Name:        "enable_cookie_management",
					Description: `(Optional) Enables cookie management if the remote repository uses cookies to manage client state.`,
				},
				resource.Attribute{
					Name:        "bypass_head_requests",
					Description: `(Optional) Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources, HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked, Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.`,
				},
				resource.Attribute{
					Name:        "priority_resolution",
					Description: `(Optional) Setting repositories with priority will cause metadata to be merged only from repositories set with this field.`,
				},
				resource.Attribute{
					Name:        "client_tls_certificate",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "content_synchronisation",
					Description: `(Optional) Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is 'false'.`,
				},
				resource.Attribute{
					Name:        "statistics_enabled",
					Description: `(Optional) If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is 'false'.`,
				},
				resource.Attribute{
					Name:        "properties_enabled",
					Description: `(Optional) If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is 'false'.`,
				},
				resource.Attribute{
					Name:        "source_origin_absence_detection",
					Description: `(Optional) If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is 'false'.`,
				},
				resource.Attribute{
					Name:        "propagate_query_params",
					Description: `(Optional, Default: false) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.`,
				},
				resource.Attribute{
					Name:        "list_remote_folder_items",
					Description: `(Optional, Default: false) Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of the 'Retrieval Cache Period'. This field exists in the API but not in the UI.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_alpine_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"alpine",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_alpine_repository.my-remote-alpine my-remote-alpine ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_bower_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"bower",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is "GITHUB". Possible values are: "GITHUB", "BITBUCKET", "OLDSTASH", "STASH", "ARTIFACTORY", "CUSTOM".`,
				},
				resource.Attribute{
					Name:        "vcs_git_download_url",
					Description: `(Optional) This attribute is used when vcs_git_provider is set to 'CUSTOM'. Provided URL will be used as proxy.`,
				},
				resource.Attribute{
					Name:        "bower_registry_url",
					Description: `(Optional) Proxy remote Bower repository. Default value is "https://registry.bower.io". ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_bower_repository.my-remote-bower my-remote-bower ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_cargo_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"cargo",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "anonymous_access",
					Description: `(Required) Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option.`,
				},
				resource.Attribute{
					Name:        "git_registry_url",
					Description: `(Optional) This is the index url, expected to be a git repository. Default value is "https://github.com/rust-lang/crates.io-index". ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_cargo_repository.my-remote-cargo my-remote-cargo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_chef_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"chef",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_chef_repository.my-remote-chef my-remote-chef ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_cocoapods_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"cocoapods",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is "GITHUB". Possible values are: "GITHUB", "BITBUCKET", "OLDSTASH", "STASH", "ARTIFACTORY", "CUSTOM".`,
				},
				resource.Attribute{
					Name:        "vcs_git_download_url",
					Description: `(Optional) This attribute is used when vcs_git_provider is set to 'CUSTOM'. Provided URL will be used as proxy.`,
				},
				resource.Attribute{
					Name:        "pods_specs_repo_url",
					Description: `(Optional) Proxy remote CocoaPods Specs repositories. Default value is "https://github.com/CocoaPods/Specs". ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_cocoapods_repository.my-remote-cocoapods my-remote-cocoapods ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_composer_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"composer",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is "GITHUB". Possible values are: "GITHUB", "BITBUCKET", "OLDSTASH", "STASH", "ARTIFACTORY", "CUSTOM".`,
				},
				resource.Attribute{
					Name:        "vcs_git_download_url",
					Description: `(Optional) This attribute is used when vcs_git_provider is set to 'CUSTOM'. Provided URL will be used as proxy.`,
				},
				resource.Attribute{
					Name:        "composer_registry_url",
					Description: `(Optional) Proxy remote Composer repository. Default value is "https://packagist.org". ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_composer_repository.my-remote-composer my-remote-composer ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_conan_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"conan",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_conan_repository.my-remote-conan my-remote-conan ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_conda_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"conda",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_conda_repository.my-remote-conda my-remote-conda ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_cran_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"cran",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_cran_repository.my-remote-cran my-remote-cran ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_debian_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"debian",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_debian_repository.my-remote-debian my-remote-debian ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_docker_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"docker",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "block_pushing_schema1",
					Description: `(Optional) When set, Artifactory will block the pulling of Docker images with manifest v2 schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1 that exist in the cache.`,
				},
				resource.Attribute{
					Name:        "enable_token_authentication",
					Description: `(Optional) Enable token (Bearer) based authentication.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) Also known as 'Foreign Layers Caching' on the UI.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) An allow list of Ant-style path patterns that determine which remote VCS. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_docker_repository.my-remote-docker my-remote-docker ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_gems_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"gems",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_gems_repository.my-remote-gems my-remote-gems ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_generic_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"generic",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_generic_repository.my-remote-generic my-remote-generic ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_gitlfs_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"gitlfs",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_gitlfs_repository.my-remote-gitlfs my-remote-gitlfs ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_go_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"go",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is "ARTIFACTORY". ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_go_repository.my-remote-go my-remote-go ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_gradle_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"gradle",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: false) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: false) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: true) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: true) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: true) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to 'true'.`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: false) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: 'generate-if-absent') Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are 'generate-if-absent', 'fail', 'ignore-and-generate', and 'pass-thru'. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_gradle_repository.gradle-remote gradle-remote ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_helm_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"helm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "helm_charts_base_url",
					Description: `(Optional) No documentation is available. Hopefully you know what this means.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) When set, external dependencies are rewritten.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_ivy_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"ivy",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: false) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: false) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: true) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: true) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: true) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to 'true'.`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: false) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: 'generate-if-absent') Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are 'generate-if-absent', 'fail', 'ignore-and-generate', and 'pass-thru'. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_ivy_repository.ivy-remote ivy-remote ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_maven_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"maven",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: false) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: false) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: true) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: true) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: true) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to 'true'.`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: false) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: 'generate-if-absent') Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are 'generate-if-absent', 'fail', 'ignore-and-generate', and 'pass-thru'. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_maven_repository.maven-remote maven-remote ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_npm_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"npm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_npm_repository.npm-remote npm-remote ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_nuget_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"nuget",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "feed_context_path",
					Description: `(Optional) When proxying a remote NuGet repository, customize feed resource location using this attribute. Default value is 'api/v2'.`,
				},
				resource.Attribute{
					Name:        "download_context_path",
					Description: `(Optional) The context path prefix through which NuGet downloads are served. For example, the NuGet Gallery download URL is 'https://nuget.org/api/v2/package', so the repository URL should be configured as 'https://nuget.org' and the download context path should be configured as 'api/v2/package'. Default value is 'api/v2/package'.`,
				},
				resource.Attribute{
					Name:        "v3_feed_url",
					Description: `(Optional) The URL to the NuGet v3 feed. Default value is 'https://api.nuget.org/v3/index.json'.`,
				},
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `(Optional) Force basic authentication credentials in order to use this repository. Default value is 'false'. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_nuget_repository.my-remote-nuget my-remote-nuget ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_opkg_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"opkg",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_opkg_repository.my-remote-opkg my-remote-opkg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_p2_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"p2",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_p2_repository.my-remote-p2 my-remote-p2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_pub_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"pub",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repository URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_pub_repository.my-remote-pub my-remote-pub ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_puppet_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"puppet",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_puppet_repository.my-remote-puppet my-remote-puppet ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_pypi_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"pypi",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "pypi_registry_url",
					Description: `(Optional) To configure the remote repo to proxy public external PyPI repository, or a PyPI repository hosted on another Artifactory server. See JFrog Pypi documentation [here](https://www.jfrog.com/confluence/display/JFROG/PyPI+Repositories) for the usage details. Default value is 'https://pypi.org'.`,
				},
				resource.Attribute{
					Name:        "pypi_repository_suffix",
					Description: `(Optional) Usually should be left as a default for 'simple', unless the remote is a PyPI server that has custom registry suffix, like +simple in DevPI. Default value is 'simple'. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_pypi_repository.pypi-remote pypi-remote ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_rpm_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"rpm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_rpm_repository.my-remote-rpm my-remote-rpm ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_sbt_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"sbt",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: false) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: false) - When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: true) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: true) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: true) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to 'true'.`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: false) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: 'generate-if-absent') Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are 'generate-if-absent', 'fail', 'ignore-and-generate', and 'pass-thru'. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_sbt_repository.sbt-remote sbt-remote ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_vcs_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"vcs",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The remote repo URL.`,
				},
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub, Bitbucket, Stash, a remote Artifactory instance or a custom Git repository. Allowed values are: 'GITHUB', 'BITBUCKET', 'OLDSTASH', 'STASH', 'ARTIFACTORY', 'CUSTOM'. Default value is 'GITHUB'`,
				},
				resource.Attribute{
					Name:        "vcs_git_download_url",
					Description: `(Optional) This attribute is used when vcs_git_provider is set to 'CUSTOM'. Provided URL will be used as proxy.`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `(Optional) The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up. ## Import Remote repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_remote_go_repository.my-remote-vcs my-remote-vcs ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_replication_config",
			Category:         "Replication",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"replication",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repo_key",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "cron_exp",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "enable_event_replication",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "replications",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "socket_timeout_millis",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Requires password encryption to be turned off ` + "`" + `POST /api/system/decrypt` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_deletes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_properties",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_statistics",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "path_prefix",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy key from Artifactory Proxies setting ## Import Replication configs can be imported using their repo key, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_replication_config.foo-rep provider_test_source ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_saml_settings",
			Category:         "Configuration",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"configuration",
				"saml",
				"settings",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_single_replication_config",
			Category:         "Replication",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"replication",
				"single",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repo_key",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "cron_exp",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "enable_event_replication",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "socket_timeout_millis",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Requires password encryption to be turned off ` + "`" + `POST /api/system/decrypt` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_deletes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_properties",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sync_statistics",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "path_prefix",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy key from Artifactory Proxies setting. ## Import Replication configs can be imported using their repo key, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_single_replication_config.foo-rep repository-key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_unmanaged_user",
			Category:         "User",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"user",
				"unmanaged",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Username for user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email for user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for the user. When omitted, a random password is generated using the following password policy: 10 characters with 1 digit, 1 symbol, with upper and lower case letters.`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `(Optional) When enabled, this user is an administrator with all the ensuing privileges. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "profile_updatable",
					Description: `(Optional) When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_ui_access",
					Description: `(Optional) When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internal_password_disabled",
					Description: `(Optional) When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of groups this user is a part of. ## Import Users can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_unmanaged_user.test-user myusername ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_user",
			Category:         "User",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Username for user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email for user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for the user. When omitted, a random password is generated using the following password policy: 10 characters with 1 digit, 1 symbol, with upper and lower case letters.`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `(Optional) When enabled, this user is an administrator with all the ensuing privileges. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "profile_updatable",
					Description: `(Optional) When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_ui_access",
					Description: `(Optional) When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internal_password_disabled",
					Description: `(Optional) When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of groups this user is a part of. - Note: If "groups" attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically. ## Import Users can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_user.test-user myusername ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository. The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `(Optional) Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.`,
				},
				resource.Attribute{
					Name:        "project_environments",
					Description: `(Optional) Project environment for assigning this repository to. Allow values: "DEV" or "PROD"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `(Optional) List of artifact patterns to include when evaluating artifact requests in the form of x/y/\`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `(Optional) List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/`,
				},
				resource.Attribute{
					Name:        "repo_layout_ref",
					Description: `(Optional) Repository layout key for the virtual repository.`,
				},
				resource.Attribute{
					Name:        "artifactory_requests_can_retrieve_remote_artifacts",
					Description: `(Optional, Default: false) Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by another Artifactory instance.`,
				},
				resource.Attribute{
					Name:        "default_deployment_repo",
					Description: `(Optional) Default repository to deploy artifacts.`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: 7200) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching. Default: 7200 seconds. ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_generic_repository.foo-generic foo-generic ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_alpine_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"alpine",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) Primary keypair used to sign artifacts. Default value is empty. ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_alpine_repository.foo-alpine foo-alpine ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_bower_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"bower",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) When set, external dependencies are rewritten. Default value is false.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_remote_repo",
					Description: `(Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_chef_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"chef",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_chef_repository.foo-chef foo-chef ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_composer_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"composer",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_composer_repository.foo-composer foo-composer ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_conan_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"conan",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_conan_repository.foo-conan foo-conan ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_conda_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"conda",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_conda_repository.foo-conda foo-conda ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_cran_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"cran",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_cran_repository.foo-cran foo-cran ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_debian_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"debian",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) Primary keypair used to sign artifacts. Default is empty.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `(Optional) Secondary keypair used to sign artifacts. Default is empty.`,
				},
				resource.Attribute{
					Name:        "optional_index_compression_formats",
					Description: `(Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are 'bz2','lzma' and 'xz'. Default value is 'bz2'.`,
				},
				resource.Attribute{
					Name:        "debian_default_architectures",
					Description: `(Optional) Specifying architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386. ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_debian_repository.foo-debian foo-debian ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_docker_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"docker",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_docker_repository.foo-docker foo-docker ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_gems_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"gems",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_gems_repository.foo-gems foo-gems ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_generic_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"generic",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_generic_repository.foo-generic foo-generic ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_gitlfs_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"gitlfs",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_gitlfs_repository.foo-gitlfs foo-gitlfs ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_go_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"go",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list. When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) 'go-import' Allow List on the UI. ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_go_repository.baz-go baz-go ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_gradle_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"gradle",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault. - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not. - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The keypair used to sign artifacts. ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_gradle_repository.foo-gradle foo-gradle ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_helm_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"helm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "use_namespaces",
					Description: `(Optional) From Artifactory 7.24.1 (SaaS Version), you can explicitly state a specific aggregated local or remote repository to fetch from a virtual by assigning namespaces to local and remote repositories. See the documentation [here](https://www.jfrog.com/confluence/display/JFROG/Kubernetes+Helm+Chart+Repositories#KubernetesHelmChartRepositories-NamespaceSupportforHelmVirtualRepositories). Default is 'false'. ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_helm_repository.foo-helm-virtual foo-helm-virtual ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_ivy_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"ivy",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault. - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not. - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The keypair used to sign artifacts. ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_ivy_repository.foo-ivy foo-ivy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_maven_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"maven",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) One of: ` + "`" + `"discard_active_reference", "discard_any_reference", "nothing"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "force_maven_authentication",
					Description: `(Optional) Forces authentication when fetching from remote repos. ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_maven_repository.maven-virt-repo maven-virt-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_npm_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"npm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_npm_repository.foo-npm foo-npm ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_nuget_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"nuget",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `(Optional) If set, user authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This is also enforced when aggregated repositories support anonymous requests. Default is ` + "`" + `false` + "`" + `. ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_nuget_repository.foo-nuget foo-nuget ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_p2_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"p2",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_p2_repository.foo-p2 foo-p2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_pub_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"pub",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_pub_repository.foo-pub foo-pub ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_puppet_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"puppet",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_puppet_repository.foo-puppet foo-puppet ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_pypi_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"pypi",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_pypi_repository.foo-pypi foo-pypi ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_rpm_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"rpm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) The primary GPG key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `(Optional) The secondary GPG key to be used to sign packages. Artifactory REST API call Get Key Pair doesn't return keys ` + "`" + `private_key` + "`" + ` and ` + "`" + `passphrase` + "`" + `, but consumes these keys in the POST call. The meta-argument ` + "`" + `lifecycle` + "`" + ` used here to make Provider ignore the changes for these two keys in the Terraform state. ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_rpm_repository.foo-rpm-virtual foo-rpm-virtual ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_sbt_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"sbt",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The effective list of actual repositories included in this virtual repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault. - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not. - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The keypair used to sign artifacts. ## Import Virtual repositories can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import artifactory_virtual_sbt_repository.foo-sbt foo-sbt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"artifactory_access_token":                       0,
		"artifactory_anonymous_user":                     1,
		"artifactory_api_key":                            2,
		"artifactory_artifact_property_webhook":          3,
		"artifactory_artifact_webhook":                   4,
		"artifactory_artifactory_release_bundle_webhook": 5,
		"artifactory_backup":                             6,
		"artifactory_build_webhook":                      7,
		"artifactory_certificate":                        8,
		"artifactory_distribution_webhook":               9,
		"artifactory_docker_webhook":                     10,
		"artifactory_federated_alpine_repository":        11,
		"artifactory_federated_bower_repository":         12,
		"artifactory_federated_cargo_repository":         13,
		"artifactory_federated_chef_repository":          14,
		"artifactory_federated_cocoapods_repository":     15,
		"artifactory_federated_composer_repository":      16,
		"artifactory_federated_conan_repository":         17,
		"artifactory_federated_conda_repository":         18,
		"artifactory_federated_cran_repository":          19,
		"artifactory_federated_debian_repository":        20,
		"artifactory_federated_docker_repository":        21,
		"artifactory_federated_gems_repository":          22,
		"artifactory_federated_generic_repository":       23,
		"artifactory_federated_gitlfs_repository":        24,
		"artifactory_federated_go_repository":            25,
		"artifactory_federated_gradle_repository":        26,
		"artifactory_federated_helm_repository":          27,
		"artifactory_federated_ivy_repository":           28,
		"artifactory_federated_maven_repository":         29,
		"artifactory_federated_npm_repository":           30,
		"artifactory_federated_nuget_repository":         31,
		"artifactory_federated_opkg_repository":          32,
		"artifactory_federated_puppet_repository":        33,
		"artifactory_federated_pypi_repository":          34,
		"artifactory_federated_rpm_repository":           35,
		"artifactory_federated_sbt_repository":           36,
		"artifactory_federated_vagrant_repository":       37,
		"artifactory_general_security":                   38,
		"artifactory_group":                              39,
		"artifactory_keypair":                            40,
		"artifactory_ldap_group_setting":                 41,
		"artifactory_ldap_setting":                       42,
		"artifactory_local":                              43,
		"artifactory_local_alpine_repository":            44,
		"artifactory_local_bower_repository":             45,
		"artifactory_local_cargo_repository":             46,
		"artifactory_local_chef_repository":              47,
		"artifactory_local_cocoapods_repository":         48,
		"artifactory_local_composer_repository":          49,
		"artifactory_local_conan_repository":             50,
		"artifactory_local_conda_repository":             51,
		"artifactory_local_cran_repository":              52,
		"artifactory_local_debian_repository":            53,
		"artifactory_local_docker_v1_repository":         54,
		"artifactory_local_docker_v2_repository":         55,
		"artifactory_local_gems_repository":              56,
		"artifactory_local_generic_repository":           57,
		"artifactory_local_gitlfs_repository":            58,
		"artifactory_local_go_repository":                59,
		"artifactory_local_gradle_repository":            60,
		"artifactory_local_helm_repository":              61,
		"artifactory_local_ivy_repository":               62,
		"artifactory_local_maven_repository":             63,
		"artifactory_local_npm_repository":               64,
		"artifactory_local_nuget_repository":             65,
		"artifactory_local_opkg_repository":              66,
		"artifactory_local_pub_repository":               67,
		"artifactory_local_puppet_repository":            68,
		"artifactory_local_pypi_repository":              69,
		"artifactory_local_rpm_repository":               70,
		"artifactory_local_sbt_repository":               71,
		"artifactory_local_vagrant_repository":           72,
		"artifactory_managed_user":                       73,
		"artifactory_oauth_settings":                     74,
		"artifactory_permission_target":                  75,
		"artifactory_pull_replication":                   76,
		"artifactory_push_replication":                   77,
		"artifactory_release_bundle_webhook":             78,
		"artifactory_remote":                             79,
		"artifactory_remote_alpine_repository":           80,
		"artifactory_remote_bower_repository":            81,
		"artifactory_remote_cargo_repository":            82,
		"artifactory_remote_chef_repository":             83,
		"artifactory_remote_cocoapods_repository":        84,
		"artifactory_remote_composer_repository":         85,
		"artifactory_remote_conan_repository":            86,
		"artifactory_remote_conda_repository":            87,
		"artifactory_remote_cran_repository":             88,
		"artifactory_remote_debian_repository":           89,
		"artifactory_remote_docker_repository":           90,
		"artifactory_remote_gems_repository":             91,
		"artifactory_remote_generic_repository":          92,
		"artifactory_remote_gitlfs_repository":           93,
		"artifactory_remote_go_repository":               94,
		"artifactory_remote_gradle_repository":           95,
		"artifactory_remote_helm_repository":             96,
		"artifactory_remote_ivy_repository":              97,
		"artifactory_remote_maven_repository":            98,
		"artifactory_remote_npm_repository":              99,
		"artifactory_remote_nuget_repository":            100,
		"artifactory_remote_opkg_repository":             101,
		"artifactory_remote_p2_repository":               102,
		"artifactory_remote_pub_repository":              103,
		"artifactory_remote_puppet_repository":           104,
		"artifactory_remote_pypi_repository":             105,
		"artifactory_remote_rpm_repository":              106,
		"artifactory_remote_sbt_repository":              107,
		"artifactory_remote_vcs_repository":              108,
		"artifactory_replication_config":                 109,
		"artifactory_saml_settings":                      110,
		"artifactory_single_replication_config":          111,
		"artifactory_unmanaged_user":                     112,
		"artifactory_user":                               113,
		"artifactory_virtual":                            114,
		"artifactory_virtual_alpine_repository":          115,
		"artifactory_virtual_bower_repository":           116,
		"artifactory_virtual_chef_repository":            117,
		"artifactory_virtual_composer_repository":        118,
		"artifactory_virtual_conan_repository":           119,
		"artifactory_virtual_conda_repository":           120,
		"artifactory_virtual_cran_repository":            121,
		"artifactory_virtual_debian_repository":          122,
		"artifactory_virtual_docker_repository":          123,
		"artifactory_virtual_gems_repository":            124,
		"artifactory_virtual_generic_repository":         125,
		"artifactory_virtual_gitlfs_repository":          126,
		"artifactory_virtual_go_repository":              127,
		"artifactory_virtual_gradle_repository":          128,
		"artifactory_virtual_helm_repository":            129,
		"artifactory_virtual_ivy_repository":             130,
		"artifactory_virtual_maven_repository":           131,
		"artifactory_virtual_npm_repository":             132,
		"artifactory_virtual_nuget_repository":           133,
		"artifactory_virtual_p2_repository":              134,
		"artifactory_virtual_pub_repository":             135,
		"artifactory_virtual_puppet_repository":          136,
		"artifactory_virtual_pypi_repository":            137,
		"artifactory_virtual_rpm_repository":             138,
		"artifactory_virtual_sbt_repository":             139,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
