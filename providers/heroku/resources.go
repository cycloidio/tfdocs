package heroku

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "heroku_account_feature",
			Category:         "Resources",
			ShortDescription: `Provides a resource to create and manage User Features on Heroku.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the account feature`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Enable or disable the account feature ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Comprised of acount email & feature name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of account feature`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of account feature ## Import Existing account features can be imported using a combination of the account email (the email address tied to the Heroku API key) and the feature name. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_account_feature.example_metrics name@example.com:metrics-request-volume ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Comprised of acount email & feature name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of account feature`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of account feature ## Import Existing account features can be imported using a combination of the account email (the email address tied to the Heroku API key) and the feature name. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_account_feature.example_metrics name@example.com:metrics-request-volume ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_addon",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku Add-On resource. These can be attach services to a Heroku app.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) The addon to add.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) Optional plan configuration.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Globally unique name of the add-on. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the add-on`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The add-on name`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The plan name`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `The ID of the plan provider`,
				},
				resource.Attribute{
					Name:        "config_vars",
					Description: `The Configuration variables of the add-on`,
				},
				resource.Attribute{
					Name:        "config_var_values",
					Description: `A sensitive map of the add-on's configuration variables. Upon add-on creation, these values will be up-to-date, while the app's own ` + "`" + `config_vars` + "`" + ` require another Terraform refresh cycle to be updated. Useful when an output contains an add-on config var value, or when a configuration needs to operate on a new add-on during an apply. ## Import Addons can be imported using the Addon ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_addon.foobar 12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the add-on`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The add-on name`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The plan name`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `The ID of the plan provider`,
				},
				resource.Attribute{
					Name:        "config_vars",
					Description: `The Configuration variables of the add-on`,
				},
				resource.Attribute{
					Name:        "config_var_values",
					Description: `A sensitive map of the add-on's configuration variables. Upon add-on creation, these values will be up-to-date, while the app's own ` + "`" + `config_vars` + "`" + ` require another Terraform refresh cycle to be updated. Useful when an output contains an add-on config var value, or when a configuration needs to operate on a new add-on during an apply. ## Import Addons can be imported using the Addon ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_addon.foobar 12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_addon_attachment",
			Category:         "Resources",
			ShortDescription: `Attaches a Heroku Addon Resource to an additional Heroku App`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "addon_id",
					Description: `(Required) The ID of the existing Heroku Addon to attach.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A friendly name for the Heroku Addon Attachment.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace value for the Heroku Addon Attachment. This can be used to configure the behaviour of the attachment. See [Heroku Platform API Reference](https://devcenter.heroku.com/articles/platform-api-reference#add-on-attachment-create) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the add-on attachment ## Import Addons can be imported using the unique Addon Attachment ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_addon_attachment.foobar 01234567-89ab-cdef-0123-456789abcdef ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the add-on attachment ## Import Addons can be imported using the unique Addon Attachment ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_addon_attachment.foobar 01234567-89ab-cdef-0123-456789abcdef ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_app",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku App resource. This can be used to create and manage applications on Heroku.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the application. In Heroku, this is also the unique ID, so it must be unique and have a minimum of 3 characters.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region that the app should be deployed in.`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `(Optional) The application stack is what platform to run the application in.`,
				},
				resource.Attribute{
					Name:        "buildpacks",
					Description: `(Optional) Buildpack names or URLs for the application. Buildpacks configured externally won't be altered if this is not present.`,
				},
				resource.Attribute{
					Name:        "space",
					Description: `(Optional) The name of a private space to create the app in.`,
				},
				resource.Attribute{
					Name:        "internal_routing",
					Description: `(Optional) If true, the application will be routable only internally in a private space. This option is only available for apps that also specify ` + "`" + `space` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) A block that can be specified once to define Heroku Team settings for this app. The fields for this block are documented below.`,
				},
				resource.Attribute{
					Name:        "acm",
					Description: `(Optional) The flag representing Automated Certificate Management for the app. The ` + "`" + `organization` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID (UUID) of the app.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the app.`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `The application stack is what platform to run the application in.`,
				},
				resource.Attribute{
					Name:        "space",
					Description: `The private space the app should run in.`,
				},
				resource.Attribute{
					Name:        "internal_routing",
					Description: `Whether internal routing is enabled the private space app.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region that the app should be deployed in.`,
				},
				resource.Attribute{
					Name:        "git_url",
					Description: `The Git URL for the application. This is used for deploying new versions of the app.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `The web (HTTP) URL that the application can be accessed at by default.`,
				},
				resource.Attribute{
					Name:        "heroku_hostname",
					Description: `A hostname for the Heroku application, suitable for pointing DNS records.`,
				},
				resource.Attribute{
					Name:        "all_config_vars",
					Description: `A map of all configuration variables that exist for the app, containing both those set by Terraform and those set externally. (These are treated as "sensitive" so that their values are redacted in console output.) This attribute is not set in state if the ` + "`" + `provider` + "`" + ` attribute ` + "`" + `set_app_all_config_vars_in_state` + "`" + ` is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The unique UUID of the Heroku app.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID (UUID) of the app.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the app.`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `The application stack is what platform to run the application in.`,
				},
				resource.Attribute{
					Name:        "space",
					Description: `The private space the app should run in.`,
				},
				resource.Attribute{
					Name:        "internal_routing",
					Description: `Whether internal routing is enabled the private space app.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region that the app should be deployed in.`,
				},
				resource.Attribute{
					Name:        "git_url",
					Description: `The Git URL for the application. This is used for deploying new versions of the app.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `The web (HTTP) URL that the application can be accessed at by default.`,
				},
				resource.Attribute{
					Name:        "heroku_hostname",
					Description: `A hostname for the Heroku application, suitable for pointing DNS records.`,
				},
				resource.Attribute{
					Name:        "all_config_vars",
					Description: `A map of all configuration variables that exist for the app, containing both those set by Terraform and those set externally. (These are treated as "sensitive" so that their values are redacted in console output.) This attribute is not set in state if the ` + "`" + `provider` + "`" + ` attribute ` + "`" + `set_app_all_config_vars_in_state` + "`" + ` is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The unique UUID of the Heroku app.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_app_config_association",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku App Config Association resource, making it possible set, update, and remove Heroku app config vars`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "vars",
					Description: `Map of config vars that can be output in plaintext.`,
				},
				resource.Attribute{
					Name:        "sensitive_vars",
					Description: `This is the same as ` + "`" + `vars` + "`" + `. The main difference between the two attributes is ` + "`" + `sensitive_vars` + "`" + ` outputs are redacted on-screen and replaced by a <sensitive> placeholder, following a terraform plan or apply. It is recommended to put private keys, passwords, etc in this argument. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app config association. ## Import This resource defines two config var attributes with one of them used for masking any sensitive/secret variables during a ` + "`" + `terraform plan|apply` + "`" + ` in a CI build, terminal, etc. This 'sensitive' distinction for config vars is unique to this provider and not a built-in feature of the Heroku Platform API. Therefore, it will not be possible to import this resource. However, it is safe to define the resource in your configuration file and execute a ` + "`" + `terraform apply` + "`" + ` as the end result is ` + "`" + `noop` + "`" + ` when the config vars already exist on the remote resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app config association. ## Import This resource defines two config var attributes with one of them used for masking any sensitive/secret variables during a ` + "`" + `terraform plan|apply` + "`" + ` in a CI build, terminal, etc. This 'sensitive' distinction for config vars is unique to this provider and not a built-in feature of the Heroku Platform API. Therefore, it will not be possible to import this resource. However, it is safe to define the resource in your configuration file and execute a ` + "`" + `terraform apply` + "`" + ` as the end result is ` + "`" + `noop` + "`" + ` when the config vars already exist on the remote resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_app_feature",
			Category:         "Resources",
			ShortDescription: `Provides a resource to create and manage App Features on Heroku.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the App Feature to manage.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether to enable or disable the App Feature. The default value is true. ## Import App features can be imported using the combination of the application name, a colon, and the feature's name. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_app_feature.log-runtime-metrics foobar:log-runtime-metrics ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_app_release",
			Category:         "Resources",
			ShortDescription: `Provides the ability to deploy a heroku release to an application`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "slug_id",
					Description: `unique identifier of slug`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description of changes in this release ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app release ## Import The most recent app release can be imported using the application name. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_app_release.foobar-release foobar ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app release ## Import The most recent app release can be imported using the application name. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_app_release.foobar-release foobar ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_app_webhook",
			Category:         "Resources",
			ShortDescription: `Provides the ability to manage application webhooks`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Required) The webhook level (either ` + "`" + `notify` + "`" + ` or ` + "`" + `sync` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Optional plan configuration.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `(Required) List of events to deliver to the webhook.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Value used to sign webhook payloads. Once set, this value cannot be fetched from the Heroku API, but it can be updated.`,
				},
				resource.Attribute{
					Name:        "authorization",
					Description: `(Optional) Values used in ` + "`" + `Authorization` + "`" + ` header. Once set, this value cannot be fetched from the Heroku API, but it can be updated. ## Importing Existing webhooks can be imported using the combination of the application name or id, a colon, and the webhook name or id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_app_webhook.foobar_release foobar:b85d9224-310b-409b-891e-c903f5a40568 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_build",
			Category:         "Resources",
			ShortDescription: `"Deploy to Heroku" for Terraform. Provides the ability to build & release code from a local or remote source, making it possible to launch apps directly from a Terraform config`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "buildpacks",
					Description: `List of buildpack GitHub URLs`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) A block that specifies the source code to build & release:`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `SHA256 hash of the tarball archive to verify its integrity, example: ` + "`" + `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855` + "`" + ``,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required unless ` + "`" + `source.url` + "`" + ` is set) Local path to the source directory or tarball archive for the app`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required unless ` + "`" + `source.path` + "`" + ` is set) ` + "`" + `https` + "`" + ` location of the source archive for the app`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Use to track what version of your source originated this build. If you are creating builds from git-versioned source code, for example, the commit hash, or release tag would be a good value to use for the version parameter. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The ID of the build`,
				},
				resource.Attribute{
					Name:        "output_stream_url",
					Description: `URL that [streams the log output from the build](https://devcenter.heroku.com/articles/build-and-release-using-the-api#streaming-build-output)`,
				},
				resource.Attribute{
					Name:        "release_id",
					Description: `The Heroku app release created with a build's slug`,
				},
				resource.Attribute{
					Name:        "slug_id",
					Description: `The Heroku slug created by a build`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `Name or ID of the [Heroku stack](https://devcenter.heroku.com/articles/stack)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of a build. Possible values are ` + "`" + `pending` + "`" + `, ` + "`" + `successful` + "`" + ` and ` + "`" + `failed` + "`" + ``,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Heroku account that created a build`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `The ID of the build`,
				},
				resource.Attribute{
					Name:        "output_stream_url",
					Description: `URL that [streams the log output from the build](https://devcenter.heroku.com/articles/build-and-release-using-the-api#streaming-build-output)`,
				},
				resource.Attribute{
					Name:        "release_id",
					Description: `The Heroku app release created with a build's slug`,
				},
				resource.Attribute{
					Name:        "slug_id",
					Description: `The Heroku slug created by a build`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `Name or ID of the [Heroku stack](https://devcenter.heroku.com/articles/stack)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of a build. Possible values are ` + "`" + `pending` + "`" + `, ` + "`" + `successful` + "`" + ` and ` + "`" + `failed` + "`" + ``,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Heroku account that created a build`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_cert",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku SSL certificate resource to manage a certificate for a Heroku app.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `(Required) The certificate chain to add`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The private key for a given certificate chain ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the add-on`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME for the SSL endpoint`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSL certificate ## Importing When importing a Heroku cert resource, the ID must be built using the app name colon the unique ID from the Heroku API. For an app named ` + "`" + `production-api` + "`" + ` with a certificate ID of ` + "`" + `b85d9224-310b-409b-891e-c903f5a40568` + "`" + `, you would import it as: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_cert.production_api production-api:b85d9224-310b-409b-891e-c903f5a40568 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the add-on`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME for the SSL endpoint`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSL certificate ## Importing When importing a Heroku cert resource, the ID must be built using the app name colon the unique ID from the Heroku API. For an app named ` + "`" + `production-api` + "`" + ` with a certificate ID of ` + "`" + `b85d9224-310b-409b-891e-c903f5a40568` + "`" + `, you would import it as: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_cert.production_api production-api:b85d9224-310b-409b-891e-c903f5a40568 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_collaborator",
			Category:         "Resources",
			ShortDescription: `Provides the ability to add/remove collaborators from applications that are not owned by a team`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email address of the collaborator ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the collaborator ## Import Collaborators can be imported using the combination of the application name, a colon, and the collaborator's email address For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_collaborator.foobar-collaborator foobar_app:collaborator@foobar.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the collaborator ## Import Collaborators can be imported using the combination of the application name, a colon, and the collaborator's email address For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_collaborator.foobar-collaborator foobar_app:collaborator@foobar.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_config",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku Config resource, making it possible to define variables that can be used in other Heroku terraform resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vars",
					Description: `Map of vars that are can be outputted in plaintext.`,
				},
				resource.Attribute{
					Name:        "sensitive_vars",
					Description: `This is the same as ` + "`" + `vars` + "`" + `. The main difference between the two attributes is ` + "`" + `sensitive_vars` + "`" + ` outputs are redacted on-screen and replaced by a <sensitive> placeholder, following a terraform ` + "`" + `plan` + "`" + ` or ` + "`" + `apply` + "`" + `. It is recommended to put private keys, passwords, etc in this argument. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the config. ## Import The ` + "`" + `heroku_config` + "`" + ` resource is a meta-resource, managed only within Terraform state. It does not exist as a native Heroku resource. Therefore, it is not possible to import an existing ` + "`" + `heroku_config` + "`" + ` configuration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the config. ## Import The ` + "`" + `heroku_config` + "`" + ` resource is a meta-resource, managed only within Terraform state. It does not exist as a native Heroku resource. Therefore, it is not possible to import an existing ` + "`" + `heroku_config` + "`" + ` configuration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_domain",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku App resource. This can be used to create and manage applications on Heroku.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name) For apps with ACM enabled (automated certificate management):`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The hostname to setup via ACM. For apps with ` + "`" + `heroku_ssl` + "`" + ` (SNI Endpoint) resources (manual certificate management):`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Must match common name or a subject alternative name of certificate in the ` + "`" + `heroku_ssl` + "`" + ` resource references by ` + "`" + `sni_endpoint_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sni_endpoint_id",
					Description: `(Required) The ID of the ` + "`" + `heroku_ssl` + "`" + ` resource to associate the domain with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain record.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME traffic should route to. ## Importing When importing a Heroku domain resource, the ID is specified ` + "`" + `APP_NAME:DOMAIN_IDENTIFIER` + "`" + `, where the domain can be identified either with the UUID from the Heroku API or the domain name. For an app named ` + "`" + `test-app` + "`" + ` with a domain name of ` + "`" + `terraform.example.com` + "`" + `, you could import it with: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_domain.default test-app:terraform.example.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain record.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME traffic should route to. ## Importing When importing a Heroku domain resource, the ID is specified ` + "`" + `APP_NAME:DOMAIN_IDENTIFIER` + "`" + `, where the domain can be identified either with the UUID from the Heroku API or the domain name. For an app named ` + "`" + `test-app` + "`" + ` with a domain name of ` + "`" + `terraform.example.com` + "`" + `, you could import it with: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_domain.default test-app:terraform.example.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_drain",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku Drain resource. This can be used to create and manage Log Drains on Heroku.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) The URL for Heroku to drain your logs to. Either ` + "`" + `url` + "`" + ` or ` + "`" + `sensitive_url` + "`" + ` must be defined.`,
				},
				resource.Attribute{
					Name:        "sensitive_url",
					Description: `(Optional) The URL for Heroku to drain your logs to. The main difference between ` + "`" + `sensitive_url` + "`" + ` and ` + "`" + `url` + "`" + ` is ` + "`" + `sensitive_url` + "`" + ` outputs are redacted, with <sensitive> displayed in place of their value during a ` + "`" + `terraform apply` + "`" + ` or ` + "`" + `terraform refresh` + "`" + `. Either ` + "`" + `url` + "`" + ` or ` + "`" + `sensitive_url` + "`" + ` must be defined. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The unique token for your created drain. ## Importing When importing a Heroku drain resource, the ID must be built using the app name colon the unique ID from the Heroku API. For an app named ` + "`" + `production-api` + "`" + ` with a drain ID of ` + "`" + `b85d9224-310b-409b-891e-c903f5a40568` + "`" + ` and the ` + "`" + `url` + "`" + ` attribute value defined for the resource, you would import it as: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_drain.production_api production-api:b85d9224-310b-409b-891e-c903f5a40568 ` + "`" + `` + "`" + `` + "`" + ` When importing a Heroku drain resource, the ID must be built using the app name colon the unique ID from the Heroku API. For an app named ` + "`" + `production-api` + "`" + ` with a drain ID of ` + "`" + `b85d9224-310b-409b-891e-c903f5a40568` + "`" + ` and the ` + "`" + `sensitive_url` + "`" + ` attribute value defined for the resource, you would import it as: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_drain.production_api production-api:b85d9224-310b-409b-891e-c903f5a40568:sensitive ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "token",
					Description: `The unique token for your created drain. ## Importing When importing a Heroku drain resource, the ID must be built using the app name colon the unique ID from the Heroku API. For an app named ` + "`" + `production-api` + "`" + ` with a drain ID of ` + "`" + `b85d9224-310b-409b-891e-c903f5a40568` + "`" + ` and the ` + "`" + `url` + "`" + ` attribute value defined for the resource, you would import it as: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_drain.production_api production-api:b85d9224-310b-409b-891e-c903f5a40568 ` + "`" + `` + "`" + `` + "`" + ` When importing a Heroku drain resource, the ID must be built using the app name colon the unique ID from the Heroku API. For an app named ` + "`" + `production-api` + "`" + ` with a drain ID of ` + "`" + `b85d9224-310b-409b-891e-c903f5a40568` + "`" + ` and the ` + "`" + `sensitive_url` + "`" + ` attribute value defined for the resource, you would import it as: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_drain.production_api production-api:b85d9224-310b-409b-891e-c903f5a40568:sensitive ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_formation",
			Category:         "Resources",
			ShortDescription: `Provides the ability to update the formation of a heroku app that has a running dyno.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) type of process such as "web"`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Required) number of processes to maintain`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) dyno size (Example: “standard-1X”). Capitalization does not matter. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the formation ## Import Existing formations can be imported using the combination of the application name, a colon, and the formation's type. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_formation.foobar-web foobar:web ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the formation ## Import Existing formations can be imported using the combination of the application name, a colon, and the formation's type. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_formation.foobar-web foobar:web ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_pipeline",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku Pipeline resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the pipeline.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Required) The owner of the pipeline. This block as the following required attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique identifier (UUID) of a pipeline owner.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of pipeline owner. Can be either ` + "`" + `user` + "`" + ` or ` + "`" + `team` + "`" + `. Regarding the ` + "`" + `owner` + "`" + ` attribute block, please note the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the pipeline.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the pipeline. ## Import Pipelines can be imported using the Pipeline ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_pipeline.foobar 12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the pipeline.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the pipeline. ## Import Pipelines can be imported using the Pipeline ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_pipeline.foobar 12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_pipeline_config_var",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku Pipeline Config Var resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pipeline_id",
					Description: `(Required) The UUID of an existing pipeline.`,
				},
				resource.Attribute{
					Name:        "pipeline_stage",
					Description: `(Required) The pipeline's stage. Supported values are ` + "`" + `test` + "`" + ` & ` + "`" + `review` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vars",
					Description: `Map of config vars that can be output in plaintext.`,
				},
				resource.Attribute{
					Name:        "sensitive_vars",
					Description: `This is the same as ` + "`" + `vars` + "`" + `. The main difference between the two attributes is ` + "`" + `sensitive_vars` + "`" + ` outputs are redacted on-screen and replaced by a ` + "`" + `<sensitive>` + "`" + ` placeholder, following a terraform ` + "`" + `plan` + "`" + ` or ` + "`" + `apply` + "`" + `. It is recommended to put private keys, passwords, etc in this argument. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "all_vars",
					Description: `All vars of a pipeline stage. This is marked ` + "`" + `sensitive` + "`" + ` so that ` + "`" + `sensitive_vars` + "`" + ` do not leak in the console/logs. ## Import This resource defines two config var attributes with one of them used for masking any sensitive/secret variables during a ` + "`" + `terraform plan|apply` + "`" + ` in a CI build, terminal, etc. This 'sensitive' distinction for config vars is unique to this provider and not a built-in feature of the Heroku Platform API. Therefore, it will not be possible to import this resource. However, it is safe to define the resource in your configuration file and execute a ` + "`" + `terraform apply` + "`" + ` as the end result is ` + "`" + `noop` + "`" + ` when the config vars already exist on the remote resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "all_vars",
					Description: `All vars of a pipeline stage. This is marked ` + "`" + `sensitive` + "`" + ` so that ` + "`" + `sensitive_vars` + "`" + ` do not leak in the console/logs. ## Import This resource defines two config var attributes with one of them used for masking any sensitive/secret variables during a ` + "`" + `terraform plan|apply` + "`" + ` in a CI build, terminal, etc. This 'sensitive' distinction for config vars is unique to this provider and not a built-in feature of the Heroku Platform API. Therefore, it will not be possible to import this resource. However, it is safe to define the resource in your configuration file and execute a ` + "`" + `terraform apply` + "`" + ` as the end result is ` + "`" + `noop` + "`" + ` when the config vars already exist on the remote resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_pipeline_coupling",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku Pipeline Coupling resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "pipeline",
					Description: `(Required) The ID of the pipeline to add this app to.`,
				},
				resource.Attribute{
					Name:        "stage",
					Description: `(Required) The stage to couple this app to. Must be one of ` + "`" + `review` + "`" + `, ` + "`" + `development` + "`" + `, ` + "`" + `staging` + "`" + `, or ` + "`" + `production` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this pipeline coupling. ## Import Pipeline couplings can be imported using the Pipeline coupling ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_pipeline_coupling.foobar 12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this pipeline coupling. ## Import Pipeline couplings can be imported using the Pipeline coupling ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_pipeline_coupling.foobar 12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_review_app_config",
			Category:         "Resources",
			ShortDescription: `Provides a resource for configuring review apps.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pipeline_id",
					Description: `(Required) The UUID of an existing pipeline.`,
				},
				resource.Attribute{
					Name:        "org_repo",
					Description: `(Required) The full_name of the repository that you want to enable review-apps from. Example ` + "`" + `heroku/homebrew-brew` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deploy_target",
					Description: `(Required) Provides a key/value pair to specify whether to use a common runtime or a private space.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Unique identifier of deploy target.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of deploy target. Must be either ` + "`" + `space` + "`" + ` or ` + "`" + `region` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "automatic_review_apps",
					Description: `(Optional) If true, this will trigger the creation of review apps when pull-requests are opened in the repo. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "base_name",
					Description: `(Optional) A unique prefix that will be used to create review app names.`,
				},
				resource.Attribute{
					Name:        "destroy_stale_apps",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, this will trigger automatic deletion of review apps when they’re stale. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "stale_days",
					Description: `(Optional) Destroy stale review apps automatically after these many days without any deploys. Must be set with ` + "`" + `destroy_stale_apps` + "`" + ` and value needs to be between ` + "`" + `1` + "`" + ` and ` + "`" + `30` + "`" + ` inclusive.`,
				},
				resource.Attribute{
					Name:        "wait_for_ci",
					Description: `(Optional) If true, review apps will only be created when CI passes. Defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported: ` + "`" + `repo_id` + "`" + ` - ID of the Github repository used for review apps. ## Import An Existing review app config using the combination of the pipeline UUID and the Github organization/repository separated by a colon. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import heroku_review_app_config.foobar afd193fb-7c5a-4d8f-afad-2388f4e6049d:heroku/homebrew-brew ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_slug",
			Category:         "Resources",
			ShortDescription: `Provides the ability to create & upload a slug (archive of executable code) to an app, making it possible to launch apps directly from a Terraform config`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "buildpack_provided_description",
					Description: `Description of language or app framework, ` + "`" + `"Ruby/Rack"` + "`" + `; displayed as the app's language in the Heroku Dashboard`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `Hash of the slug for verifying its integrity, auto-generated from contents of ` + "`" + `file_path` + "`" + ` or ` + "`" + `file_url` + "`" + `, ` + "`" + `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855` + "`" + ``,
				},
				resource.Attribute{
					Name:        "commit",
					Description: `Identification of the code with your version control system (eg: SHA of the git HEAD), ` + "`" + `"60883d9e8947a57e04dc9124f25df004866a2051"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "commit_description",
					Description: `Description of the provided commit`,
				},
				resource.Attribute{
					Name:        "file_path",
					Description: `(Required unless ` + "`" + `file_url` + "`" + ` is set) Local path to a slug archive, ` + "`" + `"slugs/current.tgz"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "file_url",
					Description: `(Required unless ` + "`" + `file_path` + "`" + ` is set)`,
				},
				resource.Attribute{
					Name:        "process_types",
					Description: `(Required) Map of [processes to launch on Heroku Dynos](https://devcenter.heroku.com/articles/process-model)`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `Name or ID of the [Heroku stack](https://devcenter.heroku.com/articles/stack) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the slug`,
				},
				resource.Attribute{
					Name:        "app",
					Description: `The ID or unique name of the Heroku app`,
				},
				resource.Attribute{
					Name:        "blob",
					Description: `Slug archive (compressed tar of executable code)`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `HTTP method to upload the archive`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Pre-signed, expiring URL to upload the archive`,
				},
				resource.Attribute{
					Name:        "buildpack_provided_description",
					Description: `Description of language or app framework, ` + "`" + `"Ruby/Rack"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `Hash of the slug for verifying its integrity, auto-generated from contents of ` + "`" + `file_path` + "`" + ` or ` + "`" + `file_url` + "`" + ``,
				},
				resource.Attribute{
					Name:        "commit",
					Description: `Identification of the code with your version control system (eg: SHA of the git HEAD), ` + "`" + `"60883d9e8947a57e04dc9124f25df004866a2051"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "commit_description",
					Description: `Description of the provided commit`,
				},
				resource.Attribute{
					Name:        "process_types",
					Description: `Map of [processes to launch on Heroku Dynos](https://devcenter.heroku.com/articles/process-model)`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Slug archive filesize in bytes`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `[Heroku stack](https://devcenter.heroku.com/articles/stack) name`,
				},
				resource.Attribute{
					Name:        "stack_id",
					Description: `[Heroku stack](https://devcenter.heroku.com/articles/stack) ID ## Import Existing slugs can be imported using the combination of the application name, a colon, and the slug ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_slug.foobar bazbux:4f1db8ef-ed5c-4c42-a3d6-3c28262d5abc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the slug`,
				},
				resource.Attribute{
					Name:        "app",
					Description: `The ID or unique name of the Heroku app`,
				},
				resource.Attribute{
					Name:        "blob",
					Description: `Slug archive (compressed tar of executable code)`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `HTTP method to upload the archive`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Pre-signed, expiring URL to upload the archive`,
				},
				resource.Attribute{
					Name:        "buildpack_provided_description",
					Description: `Description of language or app framework, ` + "`" + `"Ruby/Rack"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `Hash of the slug for verifying its integrity, auto-generated from contents of ` + "`" + `file_path` + "`" + ` or ` + "`" + `file_url` + "`" + ``,
				},
				resource.Attribute{
					Name:        "commit",
					Description: `Identification of the code with your version control system (eg: SHA of the git HEAD), ` + "`" + `"60883d9e8947a57e04dc9124f25df004866a2051"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "commit_description",
					Description: `Description of the provided commit`,
				},
				resource.Attribute{
					Name:        "process_types",
					Description: `Map of [processes to launch on Heroku Dynos](https://devcenter.heroku.com/articles/process-model)`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Slug archive filesize in bytes`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `[Heroku stack](https://devcenter.heroku.com/articles/stack) name`,
				},
				resource.Attribute{
					Name:        "stack_id",
					Description: `[Heroku stack](https://devcenter.heroku.com/articles/stack) ID ## Import Existing slugs can be imported using the combination of the application name, a colon, and the slug ID. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_slug.foobar bazbux:4f1db8ef-ed5c-4c42-a3d6-3c28262d5abc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_space",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku Space resource for running apps in isolated, highly available, secure app execution environments.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Private Space.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) The name of the Heroku Team which will own the Private Space.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) The RFC-1918 CIDR the Private Space will use. It must be a /16 in 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16`,
				},
				resource.Attribute{
					Name:        "data_cidr",
					Description: `(Optional) The RFC-1918 CIDR that the Private Space will use for the Heroku-managed peering connection that’s automatically created when using Heroku Data add-ons. It must be between a /16 and a /20`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) provision in a specific [Private Spaces region](https://devcenter.heroku.com/articles/regions#viewing-available-regions).`,
				},
				resource.Attribute{
					Name:        "shield",
					Description: `(Optional) provision as a [Shield Private Space](https://devcenter.heroku.com/articles/private-spaces#shield-private-spaces). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID (UUID) of the space.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The space's name.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The space's Heroku Team.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The space's region.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The space's CIDR.`,
				},
				resource.Attribute{
					Name:        "data_cidr",
					Description: `The space's Data CIDR.`,
				},
				resource.Attribute{
					Name:        "outbound_ips",
					Description: `The space's stable outbound [NAT IPs](https://devcenter.heroku.com/articles/platform-api-reference#space-network-address-translation). ## Import Spaces can be imported using the space ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_space.foobar MySpace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID (UUID) of the space.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The space's name.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The space's Heroku Team.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The space's region.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The space's CIDR.`,
				},
				resource.Attribute{
					Name:        "data_cidr",
					Description: `The space's Data CIDR.`,
				},
				resource.Attribute{
					Name:        "outbound_ips",
					Description: `The space's stable outbound [NAT IPs](https://devcenter.heroku.com/articles/platform-api-reference#space-network-address-translation). ## Import Spaces can be imported using the space ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_space.foobar MySpace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_space_app_access",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku Space App Access resource for managing permissions within the Private Space.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "space",
					Description: `(Required) The ID of the Private Space.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The email of the existing Heroku Team member.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) The permissions to grant the team member for the Private Space. Currently ` + "`" + `create_apps` + "`" + ` is the only supported permission. If not provided the member will have no permissions to the space. Members with admin role will always have ` + "`" + `create_apps` + "`" + ` permissions, which cannot be removed. ## Importing Existing permissions can be imported using the combination of the Private Space name, a colon, and the member email. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_space_app_access.member1 my-space:member1@foobar.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_space_inbound_ruleset",
			Category:         "Resources",
			ShortDescription: `Provides a resource for managing inbound rulesets for Heroku Private Spaces.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "space",
					Description: `(Required) The ID of the space.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) At least one ` + "`" + `rule` + "`" + ` block. Rules are documented below. A ` + "`" + `rule` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action to apply this rule to. Must be one of ` + "`" + `allow` + "`" + ` or ` + "`" + `deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) A CIDR block source for the rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the inbound ruleset.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the inbound ruleset.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_space_peering_connection_accepter",
			Category:         "Resources",
			ShortDescription: `Provides a resource for accepting VPC peering requests to Heroku Private Spaces.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "space",
					Description: `(Required) The ID of the space.`,
				},
				resource.Attribute{
					Name:        "vpc_peering_connection_id",
					Description: `(Required) The peering connection request ID. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the peering connection request.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the peering connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of the peering connection request.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the peering connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_space_vpn_connection",
			Category:         "Resources",
			ShortDescription: `Create a VPN connection between a network and a Heroku Private Space.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "space",
					Description: `(Required) The ID of the Heroku Private Space where the VPN connection will be established.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Required) The public IP address of the VPN endpoint on the network where the VPN connection will be established.`,
				},
				resource.Attribute{
					Name:        "routable_cidrs",
					Description: `(Required) A list of IPv4 CIDR blocks used by the network where the VPN connection will be established. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "space_cidr_block",
					Description: `The CIDR block for the Heroku Private Space. The network where the VPN will be established should be configured to route traffic destined for this CIDR block over the VPN link.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `The IKE version used to setup the IPsec tunnel.`,
				},
				resource.Attribute{
					Name:        "tunnels",
					Description: `Details about each VPN tunnel endpoint.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The public IP address of the tunnel.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `The pre-shared IPSec secret for the tunnel.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "space_cidr_block",
					Description: `The CIDR block for the Heroku Private Space. The network where the VPN will be established should be configured to route traffic destined for this CIDR block over the VPN link.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `The IKE version used to setup the IPsec tunnel.`,
				},
				resource.Attribute{
					Name:        "tunnels",
					Description: `Details about each VPN tunnel endpoint.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The public IP address of the tunnel.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `The pre-shared IPSec secret for the tunnel.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_ssl",
			Category:         "Resources",
			ShortDescription: `Provides a Heroku SSL certificate resource to manage a certificate for a Heroku app.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `(Required) The certificate chain to add.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) The private key for a given certificate chain. You`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSL certificate`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSL certificate ## Importing An existing SSL resource can be imported using a composite value of the app name and certificate UUID separated by a colon. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_ssl.production_api production-api:b85d9224-310b-409b-891e-c903f5a40568 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSL certificate`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSL certificate ## Importing An existing SSL resource can be imported using a composite value of the app name and certificate UUID separated by a colon. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_ssl.production_api production-api:b85d9224-310b-409b-891e-c903f5a40568 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_team_collaborator",
			Category:         "Resources",
			ShortDescription: `Provides the ability to add/edit/remove team collaborators from team applications`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) Heroku app ID (do not use app name)`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email address of the team collaborator`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) List of permissions that will be granted to the team collaborator. The order in which individual permissions are set here does not matter. Please [visit this link](https://devcenter.heroku.com/articles/app-permissions) for more information on available permissions. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team collaborator ## Import Team Collaborators can be imported using the combination of the team application name, a colon, and the collaborator's email address For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_team_collaborator.foobar-collaborator foobar_app:collaborator@foobar.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team collaborator ## Import Team Collaborators can be imported using the combination of the team application name, a colon, and the collaborator's email address For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_team_collaborator.foobar-collaborator foobar_app:collaborator@foobar.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_team_member",
			Category:         "Resources",
			ShortDescription: `Provides the ability to manage members of a Heroku team`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team",
					Description: `(Required) The name of the Heroku Team.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email address of the member`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role to assign the member. See [the API docs](https://devcenter.heroku.com/articles/platform-api-reference#team-member) for available options. ## Import Team members can be imported using the combination of the team application name, a colon, and the member's email address. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import heroku_team_member.foobar-member my-team-foobar:some-user@example.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"heroku_account_feature":                   0,
		"heroku_addon":                             1,
		"heroku_addon_attachment":                  2,
		"heroku_app":                               3,
		"heroku_app_config_association":            4,
		"heroku_app_feature":                       5,
		"heroku_app_release":                       6,
		"heroku_app_webhook":                       7,
		"heroku_build":                             8,
		"heroku_cert":                              9,
		"heroku_collaborator":                      10,
		"heroku_config":                            11,
		"heroku_domain":                            12,
		"heroku_drain":                             13,
		"heroku_formation":                         14,
		"heroku_pipeline":                          15,
		"heroku_pipeline_config_var":               16,
		"heroku_pipeline_coupling":                 17,
		"heroku_review_app_config":                 18,
		"heroku_slug":                              19,
		"heroku_space":                             20,
		"heroku_space_app_access":                  21,
		"heroku_space_inbound_ruleset":             22,
		"heroku_space_peering_connection_accepter": 23,
		"heroku_space_vpn_connection":              24,
		"heroku_ssl":                               25,
		"heroku_team_collaborator":                 26,
		"heroku_team_member":                       27,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
