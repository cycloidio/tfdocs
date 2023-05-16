package digitalocean

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_app",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) A DigitalOcean App spec describing the app. - ` + "`" + `name` + "`" + ` - (Required) The name of the app. Must be unique across all apps in the same account. - ` + "`" + `region` + "`" + ` - The slug for the DigitalOcean data center region hosting the app. - ` + "`" + `domain` + "`" + ` - Describes a domain where the application will be made available.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The hostname for the domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The domain type, which can be one of the following: - ` + "`" + `DEFAULT` + "`" + `: The default .ondigitalocean.app domain assigned to this app. - ` + "`" + `PRIMARY` + "`" + `: The primary domain for this app that is displayed as the default in the control panel, used in bindable environment variables, and any other places that reference an app's live URL. Only one domain may be set as primary. - ` + "`" + `ALIAS` + "`" + `: A non-primary domain.`,
				},
				resource.Attribute{
					Name:        "wildcard",
					Description: `A boolean indicating whether the domain includes all sub-domains, in addition to the given domain.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `If the domain uses DigitalOcean DNS and you would like App Platform to automatically manage it for you, set this to the name of the domain on your account. - ` + "`" + `env` + "`" + ` - Describes an app-wide environment variable made available to all components.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The name of the environment variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the environment variable.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `. - ` + "`" + `alert` + "`" + ` - Describes an alert policy for the app.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The type of the alert to configure. Top-level app alert policies can be: ` + "`" + `DEPLOYMENT_FAILED` + "`" + `, ` + "`" + `DEPLOYMENT_LIVE` + "`" + `, ` + "`" + `DOMAIN_FAILED` + "`" + `, or ` + "`" + `DOMAIN_LIVE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Determines whether or not the alert is disabled (default: ` + "`" + `false` + "`" + `). A spec can contain multiple components. A ` + "`" + `service` + "`" + ` can contain:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the component`,
				},
				resource.Attribute{
					Name:        "build_command",
					Description: `An optional build command to run while building this component from source.`,
				},
				resource.Attribute{
					Name:        "dockerfile_path",
					Description: `The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.`,
				},
				resource.Attribute{
					Name:        "source_dir",
					Description: `An optional path to the working directory to use for the build.`,
				},
				resource.Attribute{
					Name:        "run_command",
					Description: `An optional run command to override the component's default.`,
				},
				resource.Attribute{
					Name:        "environment_slug",
					Description: `An environment slug describing the type of this app.`,
				},
				resource.Attribute{
					Name:        "instance_size_slug",
					Description: `The instance size to use for this component. This determines the plan (basic or professional) and the available CPU and memory. The list of available instance sizes can be [found with the API](https://docs.digitalocean.com/reference/api/api-reference/#operation/list_instance_sizes) or using the [doctl CLI](https://docs.digitalocean.com/reference/doctl/) (` + "`" + `doctl apps tier instance-size list` + "`" + `). Default: ` + "`" + `basic-xxs` + "`" + ``,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The amount of instances that this component should be scaled to.`,
				},
				resource.Attribute{
					Name:        "http_port",
					Description: `The internal port on which this service's run command will listen.`,
				},
				resource.Attribute{
					Name:        "internal_ports",
					Description: `A list of ports on which this service will listen for internal traffic.`,
				},
				resource.Attribute{
					Name:        "git",
					Description: `A Git repo to use as the component's source. The repository must be able to be cloned without authentication. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + ` or ` + "`" + `gitlab` + "`" + ` may be set - ` + "`" + `repo_clone_url` + "`" + ` - The clone URL of the repo. - ` + "`" + `branch` + "`" + ` - The name of the branch to use.`,
				},
				resource.Attribute{
					Name:        "github",
					Description: `A GitHub repo to use as the component's source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/github/install). Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "gitlab",
					Description: `A Gitlab repo to use as the component's source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/gitlab/install). Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `An image to use as the component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `registry_type` + "`" + ` - The registry type. One of ` + "`" + `DOCR` + "`" + ` (DigitalOcean container registry) or ` + "`" + `DOCKER_HUB` + "`" + `. - ` + "`" + `registry` + "`" + ` - The registry name. Must be left empty for the ` + "`" + `DOCR` + "`" + ` registry type. Required for the ` + "`" + `DOCKER_HUB` + "`" + ` registry type. - ` + "`" + `repository` + "`" + ` - The repository name. - ` + "`" + `tag` + "`" + ` - The repository tag. Defaults to ` + "`" + `latest` + "`" + ` if not provided. - ` + "`" + `deploy_on_push` + "`" + ` - Configures automatically deploying images pushed to DOCR. - ` + "`" + `enabled` + "`" + ` - Whether to automatically deploy images pushed to DOCR.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Describes an environment variable made available to an app competent. - ` + "`" + `key` + "`" + ` - The name of the environment variable. - ` + "`" + `value` + "`" + ` - The value of the environment variable. - ` + "`" + `scope` + "`" + ` - The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default). - ` + "`" + `type` + "`" + ` - The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `An HTTP paths that should be routed to this component. - ` + "`" + `path` + "`" + ` - Paths must start with ` + "`" + `/` + "`" + ` and must be unique within the app. - ` + "`" + `preserve_path_prefix` + "`" + ` - An optional flag to preserve the path that is forwarded to the backend service.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `A health check to determine the availability of this component. - ` + "`" + `http_path` + "`" + ` - The route path used for the HTTP health check ping. - ` + "`" + `initial_delay_seconds` + "`" + ` - The number of seconds to wait before beginning health checks. - ` + "`" + `period_seconds` + "`" + ` - The number of seconds to wait between health checks. - ` + "`" + `timeout_seconds` + "`" + ` - The number of seconds after which the check times out. - ` + "`" + `success_threshold` + "`" + ` - The number of successful health checks before considered healthy. - ` + "`" + `failure_threshold` + "`" + ` - The number of failed health checks before considered unhealthy.`,
				},
				resource.Attribute{
					Name:        "cors",
					Description: `The [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) policies of the app. - ` + "`" + `allow_origins` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` can be - ` + "`" + `exact` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` header will be set to the client's origin only if the client's origin exactly matches the value you provide. - ` + "`" + `prefix` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` header will be set to the client's origin if the beginning of the client's origin matches the value you provide. - ` + "`" + `regex` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` header will be set to the client's origin if the client’s origin matches the regex you provide, in [RE2 style syntax](https://github.com/google/re2/wiki/Syntax). - ` + "`" + `allow_headers` + "`" + ` - The set of allowed HTTP request headers. This configures the ` + "`" + `Access-Control-Allow-Headers` + "`" + ` header. - ` + "`" + `max_age` + "`" + ` - An optional duration specifying how long browsers can cache the results of a preflight request. This configures the Access-Control-Max-Age header. Example: ` + "`" + `5h30m` + "`" + `. - ` + "`" + `expose_headers` + "`" + ` - The set of HTTP response headers that browsers are allowed to access. This configures the ` + "`" + `Access-Control-Expose-Headers` + "`" + ` header. - ` + "`" + `allow_methods` + "`" + ` - The set of allowed HTTP methods. This configures the ` + "`" + `Access-Control-Allow-Methods` + "`" + ` header. - ` + "`" + `allow_credentials` + "`" + ` - Whether browsers should expose the response to the client-side JavaScript code when the request's credentials mode is ` + "`" + `include` + "`" + `. This configures the ` + "`" + `Access-Control-Allow-Credentials` + "`" + ` header.`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `Describes an alert policy for the component. - ` + "`" + `rule` + "`" + ` - The type of the alert to configure. Component app alert policies can be: ` + "`" + `CPU_UTILIZATION` + "`" + `, ` + "`" + `MEM_UTILIZATION` + "`" + `, or ` + "`" + `RESTART_COUNT` + "`" + `. - ` + "`" + `value` + "`" + ` - The threshold for the type of the warning. - ` + "`" + `operator` + "`" + ` - The operator to use. This is either of ` + "`" + `GREATER_THAN` + "`" + ` or ` + "`" + `LESS_THAN` + "`" + `. - ` + "`" + `window` + "`" + ` - The time before alerts should be triggered. This is may be one of: ` + "`" + `FIVE_MINUTES` + "`" + `, ` + "`" + `TEN_MINUTES` + "`" + `, ` + "`" + `THIRTY_MINUTES` + "`" + `, ` + "`" + `ONE_HOUR` + "`" + `. - ` + "`" + `disabled` + "`" + ` - Determines whether or not the alert is disabled (default: ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "log_destination",
					Description: `Describes a log forwarding destination. - ` + "`" + `name` + "`" + ` - Name of the log destination. Minimum length: 2. Maximum length: 42. - ` + "`" + `papertrail` + "`" + ` - Papertrail configuration. - ` + "`" + `endpoint` + "`" + ` - Papertrail syslog endpoint. - ` + "`" + `datadog` + "`" + ` - Datadog configuration. - ` + "`" + `endpoint` + "`" + ` - Datadog HTTP log intake endpoint. - ` + "`" + `api_key` + "`" + ` - Datadog API key. - ` + "`" + `logtail` + "`" + ` - Logtail configuration. - ` + "`" + `token` + "`" + ` - Logtail token. A ` + "`" + `static_site` + "`" + ` can contain:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the component.`,
				},
				resource.Attribute{
					Name:        "build_command",
					Description: `An optional build command to run while building this component from source.`,
				},
				resource.Attribute{
					Name:        "dockerfile_path",
					Description: `The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.`,
				},
				resource.Attribute{
					Name:        "source_dir",
					Description: `An optional path to the working directory to use for the build.`,
				},
				resource.Attribute{
					Name:        "environment_slug",
					Description: `An environment slug describing the type of this app.`,
				},
				resource.Attribute{
					Name:        "output_dir",
					Description: `An optional path to where the built assets will be located, relative to the build context. If not set, App Platform will automatically scan for these directory names: ` + "`" + `_static` + "`" + `, ` + "`" + `dist` + "`" + `, ` + "`" + `public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `The name of the index document to use when serving this static site.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `The name of the error document to use when serving this static site.`,
				},
				resource.Attribute{
					Name:        "catchall_document",
					Description: `The name of the document to use as the fallback for any requests to documents that are not found when serving this static site.`,
				},
				resource.Attribute{
					Name:        "git",
					Description: `A Git repo to use as the component's source. The repository must be able to be cloned without authentication. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + ` or ` + "`" + `gitlab` + "`" + ` may be set. - ` + "`" + `repo_clone_url` + "`" + ` - The clone URL of the repo. - ` + "`" + `branch` + "`" + ` - The name of the branch to use.`,
				},
				resource.Attribute{
					Name:        "github",
					Description: `A GitHub repo to use as the component's source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/github/install). Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "gitlab",
					Description: `A Gitlab repo to use as the component's source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/gitlab/install). Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Describes an environment variable made available to an app competent. - ` + "`" + `key` + "`" + ` - The name of the environment variable. - ` + "`" + `value` + "`" + ` - The value of the environment variable. - ` + "`" + `scope` + "`" + ` - The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default). - ` + "`" + `type` + "`" + ` - The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `An HTTP paths that should be routed to this component. - ` + "`" + `path` + "`" + ` - Paths must start with ` + "`" + `/` + "`" + ` and must be unique within the app. - ` + "`" + `preserve_path_prefix` + "`" + ` - An optional flag to preserve the path that is forwarded to the backend service.`,
				},
				resource.Attribute{
					Name:        "cors",
					Description: `The [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) policies of the app. - ` + "`" + `allow_origins` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` can be - ` + "`" + `exact` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` header will be set to the client's origin only if the client's origin exactly matches the value you provide. - ` + "`" + `prefix` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` header will be set to the client's origin if the beginning of the client's origin matches the value you provide. - ` + "`" + `regex` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` header will be set to the client's origin if the client’s origin matches the regex you provide, in [RE2 style syntax](https://github.com/google/re2/wiki/Syntax). - ` + "`" + `allow_headers` + "`" + ` - The set of allowed HTTP request headers. This configures the ` + "`" + `Access-Control-Allow-Headers` + "`" + ` header. - ` + "`" + `max_age` + "`" + ` - An optional duration specifying how long browsers can cache the results of a preflight request. This configures the Access-Control-Max-Age header. Example: ` + "`" + `5h30m` + "`" + `. - ` + "`" + `expose_headers` + "`" + ` - The set of HTTP response headers that browsers are allowed to access. This configures the ` + "`" + `Access-Control-Expose-Headers` + "`" + ` header. - ` + "`" + `allow_methods` + "`" + ` - The set of allowed HTTP methods. This configures the ` + "`" + `Access-Control-Allow-Methods` + "`" + ` header. - ` + "`" + `allow_credentials` + "`" + ` - Whether browsers should expose the response to the client-side JavaScript code when the request's credentials mode is ` + "`" + `include` + "`" + `. This configures the ` + "`" + `Access-Control-Allow-Credentials` + "`" + ` header. A ` + "`" + `worker` + "`" + ` can contain:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the component`,
				},
				resource.Attribute{
					Name:        "build_command",
					Description: `An optional build command to run while building this component from source.`,
				},
				resource.Attribute{
					Name:        "dockerfile_path",
					Description: `The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.`,
				},
				resource.Attribute{
					Name:        "source_dir",
					Description: `An optional path to the working directory to use for the build.`,
				},
				resource.Attribute{
					Name:        "run_command",
					Description: `An optional run command to override the component's default.`,
				},
				resource.Attribute{
					Name:        "environment_slug",
					Description: `An environment slug describing the type of this app.`,
				},
				resource.Attribute{
					Name:        "instance_size_slug",
					Description: `The instance size to use for this component. This determines the plan (basic or professional) and the available CPU and memory. The list of available instance sizes can be [found with the API](https://docs.digitalocean.com/reference/api/api-reference/#operation/list_instance_sizes) or using the [doctl CLI](https://docs.digitalocean.com/reference/doctl/) (` + "`" + `doctl apps tier instance-size list` + "`" + `). Default: ` + "`" + `basic-xxs` + "`" + ``,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The amount of instances that this component should be scaled to.`,
				},
				resource.Attribute{
					Name:        "git",
					Description: `A Git repo to use as the component's source. The repository must be able to be cloned without authentication. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + ` or ` + "`" + `gitlab` + "`" + ` may be set - ` + "`" + `repo_clone_url` + "`" + ` - The clone URL of the repo. - ` + "`" + `branch` + "`" + ` - The name of the branch to use.`,
				},
				resource.Attribute{
					Name:        "github",
					Description: `A GitHub repo to use as the component's source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/github/install). Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "gitlab",
					Description: `A Gitlab repo to use as the component's source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/gitlab/install). Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `An image to use as the component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `registry_type` + "`" + ` - The registry type. One of ` + "`" + `DOCR` + "`" + ` (DigitalOcean container registry) or ` + "`" + `DOCKER_HUB` + "`" + `. - ` + "`" + `registry` + "`" + ` - The registry name. Must be left empty for the ` + "`" + `DOCR` + "`" + ` registry type. Required for the ` + "`" + `DOCKER_HUB` + "`" + ` registry type. - ` + "`" + `repository` + "`" + ` - The repository name. - ` + "`" + `tag` + "`" + ` - The repository tag. Defaults to ` + "`" + `latest` + "`" + ` if not provided. - ` + "`" + `deploy_on_push` + "`" + ` - Configures automatically deploying images pushed to DOCR. - ` + "`" + `enabled` + "`" + ` - Whether to automatically deploy images pushed to DOCR.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Describes an environment variable made available to an app competent. - ` + "`" + `key` + "`" + ` - The name of the environment variable. - ` + "`" + `value` + "`" + ` - The value of the environment variable. - ` + "`" + `scope` + "`" + ` - The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default). - ` + "`" + `type` + "`" + ` - The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `Describes an alert policy for the component. - ` + "`" + `rule` + "`" + ` - The type of the alert to configure. Component app alert policies can be: ` + "`" + `CPU_UTILIZATION` + "`" + `, ` + "`" + `MEM_UTILIZATION` + "`" + `, or ` + "`" + `RESTART_COUNT` + "`" + `. - ` + "`" + `value` + "`" + ` - The threshold for the type of the warning. - ` + "`" + `operator` + "`" + ` - The operator to use. This is either of ` + "`" + `GREATER_THAN` + "`" + ` or ` + "`" + `LESS_THAN` + "`" + `. - ` + "`" + `window` + "`" + ` - The time before alerts should be triggered. This is may be one of: ` + "`" + `FIVE_MINUTES` + "`" + `, ` + "`" + `TEN_MINUTES` + "`" + `, ` + "`" + `THIRTY_MINUTES` + "`" + `, ` + "`" + `ONE_HOUR` + "`" + `. - ` + "`" + `disabled` + "`" + ` - Determines whether or not the alert is disabled (default: ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "log_destination",
					Description: `Describes a log forwarding destination. - ` + "`" + `name` + "`" + ` - Name of the log destination. Minimum length: 2. Maximum length: 42. - ` + "`" + `papertrail` + "`" + ` - Papertrail configuration. - ` + "`" + `endpoint` + "`" + ` - Papertrail syslog endpoint. - ` + "`" + `datadog` + "`" + ` - Datadog configuration. - ` + "`" + `endpoint` + "`" + ` - Datadog HTTP log intake endpoint. - ` + "`" + `api_key` + "`" + ` - Datadog API key. - ` + "`" + `logtail` + "`" + ` - Logtail configuration. - ` + "`" + `token` + "`" + ` - Logtail token. A ` + "`" + `job` + "`" + ` can contain:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the component`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The type of job and when it will be run during the deployment process. It may be one of: - ` + "`" + `UNSPECIFIED` + "`" + `: Default job type, will auto-complete to POST_DEPLOY kind. - ` + "`" + `PRE_DEPLOY` + "`" + `: Indicates a job that runs before an app deployment. - ` + "`" + `POST_DEPLOY` + "`" + `: Indicates a job that runs after an app deployment. - ` + "`" + `FAILED_DEPLOY` + "`" + `: Indicates a job that runs after a component fails to deploy.`,
				},
				resource.Attribute{
					Name:        "build_command",
					Description: `An optional build command to run while building this component from source.`,
				},
				resource.Attribute{
					Name:        "dockerfile_path",
					Description: `The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.`,
				},
				resource.Attribute{
					Name:        "source_dir",
					Description: `An optional path to the working directory to use for the build.`,
				},
				resource.Attribute{
					Name:        "run_command",
					Description: `An optional run command to override the component's default.`,
				},
				resource.Attribute{
					Name:        "environment_slug",
					Description: `An environment slug describing the type of this app.`,
				},
				resource.Attribute{
					Name:        "instance_size_slug",
					Description: `The instance size to use for this component. This determines the plan (basic or professional) and the available CPU and memory. The list of available instance sizes can be [found with the API](https://docs.digitalocean.com/reference/api/api-reference/#operation/list_instance_sizes) or using the [doctl CLI](https://docs.digitalocean.com/reference/doctl/) (` + "`" + `doctl apps tier instance-size list` + "`" + `). Default: ` + "`" + `basic-xxs` + "`" + ``,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The amount of instances that this component should be scaled to.`,
				},
				resource.Attribute{
					Name:        "git",
					Description: `A Git repo to use as the component's source. The repository must be able to be cloned without authentication. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + ` or ` + "`" + `gitlab` + "`" + ` may be set - ` + "`" + `repo_clone_url` + "`" + ` - The clone URL of the repo. - ` + "`" + `branch` + "`" + ` - The name of the branch to use.`,
				},
				resource.Attribute{
					Name:        "github",
					Description: `A GitHub repo to use as the component's source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/github/install). Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "gitlab",
					Description: `A Gitlab repo to use as the component's source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/gitlab/install). Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `An image to use as the component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `registry_type` + "`" + ` - The registry type. One of ` + "`" + `DOCR` + "`" + ` (DigitalOcean container registry) or ` + "`" + `DOCKER_HUB` + "`" + `. - ` + "`" + `registry` + "`" + ` - The registry name. Must be left empty for the ` + "`" + `DOCR` + "`" + ` registry type. Required for the ` + "`" + `DOCKER_HUB` + "`" + ` registry type. - ` + "`" + `repository` + "`" + ` - The repository name. - ` + "`" + `tag` + "`" + ` - The repository tag. Defaults to ` + "`" + `latest` + "`" + ` if not provided. - ` + "`" + `deploy_on_push` + "`" + ` - Configures automatically deploying images pushed to DOCR. - ` + "`" + `enabled` + "`" + ` - Whether to automatically deploy images pushed to DOCR.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Describes an environment variable made available to an app competent. - ` + "`" + `key` + "`" + ` - The name of the environment variable. - ` + "`" + `value` + "`" + ` - The value of the environment variable. - ` + "`" + `scope` + "`" + ` - The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default). - ` + "`" + `type` + "`" + ` - The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `Describes an alert policy for the component. - ` + "`" + `rule` + "`" + ` - The type of the alert to configure. Component app alert policies can be: ` + "`" + `CPU_UTILIZATION` + "`" + `, ` + "`" + `MEM_UTILIZATION` + "`" + `, or ` + "`" + `RESTART_COUNT` + "`" + `. - ` + "`" + `value` + "`" + ` - The threshold for the type of the warning. - ` + "`" + `operator` + "`" + ` - The operator to use. This is either of ` + "`" + `GREATER_THAN` + "`" + ` or ` + "`" + `LESS_THAN` + "`" + `. - ` + "`" + `window` + "`" + ` - The time before alerts should be triggered. This is may be one of: ` + "`" + `FIVE_MINUTES` + "`" + `, ` + "`" + `TEN_MINUTES` + "`" + `, ` + "`" + `THIRTY_MINUTES` + "`" + `, ` + "`" + `ONE_HOUR` + "`" + `. - ` + "`" + `disabled` + "`" + ` - Determines whether or not the alert is disabled (default: ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "log_destination",
					Description: `Describes a log forwarding destination. - ` + "`" + `name` + "`" + ` - Name of the log destination. Minimum length: 2. Maximum length: 42. - ` + "`" + `papertrail` + "`" + ` - Papertrail configuration. - ` + "`" + `endpoint` + "`" + ` - Papertrail syslog endpoint. - ` + "`" + `datadog` + "`" + ` - Datadog configuration. - ` + "`" + `endpoint` + "`" + ` - Datadog HTTP log intake endpoint. - ` + "`" + `api_key` + "`" + ` - Datadog API key. - ` + "`" + `logtail` + "`" + ` - Logtail configuration. - ` + "`" + `token` + "`" + ` - Logtail token. A ` + "`" + `function` + "`" + ` component can contain:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the component.`,
				},
				resource.Attribute{
					Name:        "source_dir",
					Description: `An optional path to the working directory to use for the build.`,
				},
				resource.Attribute{
					Name:        "git",
					Description: `A Git repo to use as the component's source. The repository must be able to be cloned without authentication. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + ` or ` + "`" + `gitlab` + "`" + ` may be set. - ` + "`" + `repo_clone_url` + "`" + ` - The clone URL of the repo. - ` + "`" + `branch` + "`" + ` - The name of the branch to use.`,
				},
				resource.Attribute{
					Name:        "github",
					Description: `A GitHub repo to use as the component's source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/github/install). Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "gitlab",
					Description: `A Gitlab repo to use as the component's source. DigitalOcean App Platform must have [access to the repository](https://cloud.digitalocean.com/apps/gitlab/install). Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Describes an environment variable made available to an app competent. - ` + "`" + `key` + "`" + ` - The name of the environment variable. - ` + "`" + `value` + "`" + ` - The value of the environment variable. - ` + "`" + `scope` + "`" + ` - The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default). - ` + "`" + `type` + "`" + ` - The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `An HTTP paths that should be routed to this component. - ` + "`" + `path` + "`" + ` - Paths must start with ` + "`" + `/` + "`" + ` and must be unique within the app. - ` + "`" + `preserve_path_prefix` + "`" + ` - An optional flag to preserve the path that is forwarded to the backend service.`,
				},
				resource.Attribute{
					Name:        "cors",
					Description: `The [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) policies of the app. - ` + "`" + `allow_origins` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` can be - ` + "`" + `exact` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` header will be set to the client's origin only if the client's origin exactly matches the value you provide. - ` + "`" + `prefix` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` header will be set to the client's origin if the beginning of the client's origin matches the value you provide. - ` + "`" + `regex` + "`" + ` - The ` + "`" + `Access-Control-Allow-Origin` + "`" + ` header will be set to the client's origin if the client’s origin matches the regex you provide, in [RE2 style syntax](https://github.com/google/re2/wiki/Syntax). - ` + "`" + `allow_headers` + "`" + ` - The set of allowed HTTP request headers. This configures the ` + "`" + `Access-Control-Allow-Headers` + "`" + ` header. - ` + "`" + `max_age` + "`" + ` - An optional duration specifying how long browsers can cache the results of a preflight request. This configures the Access-Control-Max-Age header. Example: ` + "`" + `5h30m` + "`" + `. - ` + "`" + `expose_headers` + "`" + ` - The set of HTTP response headers that browsers are allowed to access. This configures the ` + "`" + `Access-Control-Expose-Headers` + "`" + ` header. - ` + "`" + `allow_methods` + "`" + ` - The set of allowed HTTP methods. This configures the ` + "`" + `Access-Control-Allow-Methods` + "`" + ` header. - ` + "`" + `allow_credentials` + "`" + ` - Whether browsers should expose the response to the client-side JavaScript code when the request's credentials mode is ` + "`" + `include` + "`" + `. This configures the ` + "`" + `Access-Control-Allow-Credentials` + "`" + ` header.`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `Describes an alert policy for the component. - ` + "`" + `rule` + "`" + ` - The type of the alert to configure. Component app alert policies can be: ` + "`" + `CPU_UTILIZATION` + "`" + `, ` + "`" + `MEM_UTILIZATION` + "`" + `, or ` + "`" + `RESTART_COUNT` + "`" + `. - ` + "`" + `value` + "`" + ` - The threshold for the type of the warning. - ` + "`" + `operator` + "`" + ` - The operator to use. This is either of ` + "`" + `GREATER_THAN` + "`" + ` or ` + "`" + `LESS_THAN` + "`" + `. - ` + "`" + `window` + "`" + ` - The time before alerts should be triggered. This is may be one of: ` + "`" + `FIVE_MINUTES` + "`" + `, ` + "`" + `TEN_MINUTES` + "`" + `, ` + "`" + `THIRTY_MINUTES` + "`" + `, ` + "`" + `ONE_HOUR` + "`" + `. - ` + "`" + `disabled` + "`" + ` - Determines whether or not the alert is disabled (default: ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "log_destination",
					Description: `Describes a log forwarding destination. - ` + "`" + `name` + "`" + ` - Name of the log destination. Minimum length: 2. Maximum length: 42. - ` + "`" + `papertrail` + "`" + ` - Papertrail configuration. - ` + "`" + `endpoint` + "`" + ` - Papertrail syslog endpoint. - ` + "`" + `datadog` + "`" + ` - Datadog configuration. - ` + "`" + `endpoint` + "`" + ` - Datadog HTTP log intake endpoint. - ` + "`" + `api_key` + "`" + ` - Datadog API key. - ` + "`" + `logtail` + "`" + ` - Logtail configuration. - ` + "`" + `token` + "`" + ` - Logtail token. A ` + "`" + `database` + "`" + ` can contain:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the component.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The database engine to use (` + "`" + `MYSQL` + "`" + `, ` + "`" + `PG` + "`" + `, ` + "`" + `REDIS` + "`" + `, or ` + "`" + `MONGODB` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the database engine.`,
				},
				resource.Attribute{
					Name:        "production",
					Description: `Whether this is a production or dev database.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `The name of the underlying DigitalOcean DBaaS cluster. This is required for production databases. For dev databases, if ` + "`" + `cluster_name` + "`" + ` is not set, a new cluster will be provisioned.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `The name of the MySQL or PostgreSQL database to configure.`,
				},
				resource.Attribute{
					Name:        "db_user",
					Description: `The name of the MySQL or PostgreSQL user to configure. This resource supports [customized create timeouts](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts). The default timeout is 30 minutes. ## Attributes Reference In addition to the above attributes, the following are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app.`,
				},
				resource.Attribute{
					Name:        "default_ingress",
					Description: `The default URL to access the app.`,
				},
				resource.Attribute{
					Name:        "live_url",
					Description: `The live URL of the app.`,
				},
				resource.Attribute{
					Name:        "active_deployment_id",
					Description: `The ID the app's currently active deployment.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource identifier for the app.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date and time of when the app was last updated.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time of when the app was created. ## Import An app can be imported using its ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_app.myapp fb06ad00-351f-45c8-b5eb-13523c438661 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app.`,
				},
				resource.Attribute{
					Name:        "default_ingress",
					Description: `The default URL to access the app.`,
				},
				resource.Attribute{
					Name:        "live_url",
					Description: `The live URL of the app.`,
				},
				resource.Attribute{
					Name:        "active_deployment_id",
					Description: `The ID the app's currently active deployment.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource identifier for the app.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date and time of when the app was last updated.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time of when the app was created. ## Import An app can be imported using its ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_app.myapp fb06ad00-351f-45c8-b5eb-13523c438661 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_cdn",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "origin",
					Description: `(Required) The fully qualified domain name, (FQDN) for a Space.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live for the CDN Endpoint, in seconds. Default is 3600 seconds.`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `(Optional) The fully qualified domain name (FQDN) of the custom subdomain used with the CDN Endpoint. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a CDN Endpoint.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `The fully qualified domain name, (FQDN) of a space referenced by the CDN Endpoint.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The fully qualified domain name (FQDN) from which the CDN-backed content is served.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the CDN Endpoint was created.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The time to live for the CDN Endpoint, in seconds.`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `The fully qualified domain name (FQDN) of the custom subdomain used with the CDN Endpoint. ## Import CDN Endpoints can be imported using the CDN ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_cdn.mycdn fb06ad00-351f-45c8-b5eb-13523c438661 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a CDN Endpoint.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `The fully qualified domain name, (FQDN) of a space referenced by the CDN Endpoint.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The fully qualified domain name (FQDN) from which the CDN-backed content is served.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the CDN Endpoint was created.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The time to live for the CDN Endpoint, in seconds.`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `The fully qualified domain name (FQDN) of the custom subdomain used with the CDN Endpoint. ## Import CDN Endpoints can be imported using the CDN ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_cdn.mycdn fb06ad00-351f-45c8-b5eb-13523c438661 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_certificate",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the certificate for identification.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of certificate to provision. Can be either ` + "`" + `custom` + "`" + ` or ` + "`" + `lets_encrypt` + "`" + `. Defaults to ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) The contents of a PEM-formatted private-key corresponding to the SSL certificate. Only valid when type is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "leaf_certificate",
					Description: `(Optional) The contents of a PEM-formatted public TLS certificate. Only valid when type is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `(Optional) The full PEM-formatted trust chain between the certificate authority's certificate and your domain's TLS certificate. Only valid when type is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `(Optional) List of fully qualified domain names (FQDNs) for which the certificate will be issued. The domains must be managed using DigitalOcean's DNS. Only valid when type is ` + "`" + `lets_encrypt` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique name of the certificate`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the certificate`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the certificate`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The expiration date of the certificate`,
				},
				resource.Attribute{
					Name:        "sha1_fingerprint",
					Description: `The SHA-1 fingerprint of the certificate ## Import Certificates can be imported using the certificate ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_certificate.mycertificate cert-01 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique name of the certificate`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the certificate`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the certificate`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The expiration date of the certificate`,
				},
				resource.Attribute{
					Name:        "sha1_fingerprint",
					Description: `The SHA-1 fingerprint of the certificate ## Import Certificates can be imported using the certificate ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_certificate.mycertificate cert-01 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_container_registry",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the container_registry`,
				},
				resource.Attribute{
					Name:        "subscription_tier_slug",
					Description: `(Required) The slug identifier for the subscription tier to use (` + "`" + `starter` + "`" + `, ` + "`" + `basic` + "`" + `, or ` + "`" + `professional` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The slug identifier of for region where registry data will be stored. When not provided, a region will be selected automatically. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the container registry`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container registry`,
				},
				resource.Attribute{
					Name:        "subscription_tier_slug",
					Description: `The slug identifier for the subscription tier`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The slug identifier for the region`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The URL endpoint of the container registry. Ex: ` + "`" + `registry.digitalocean.com/my_registry` + "`" + ``,
				},
				resource.Attribute{
					Name:        "server_url",
					Description: `The domain of the container registry. Ex: ` + "`" + `registry.digitalocean.com` + "`" + ``,
				},
				resource.Attribute{
					Name:        "storage_usage_bytes",
					Description: `The amount of storage used in the registry in bytes.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the registry was created ## Import Container Registries can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_container_registry.myregistry registryname ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the container registry`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container registry`,
				},
				resource.Attribute{
					Name:        "subscription_tier_slug",
					Description: `The slug identifier for the subscription tier`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The slug identifier for the region`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The URL endpoint of the container registry. Ex: ` + "`" + `registry.digitalocean.com/my_registry` + "`" + ``,
				},
				resource.Attribute{
					Name:        "server_url",
					Description: `The domain of the container registry. Ex: ` + "`" + `registry.digitalocean.com` + "`" + ``,
				},
				resource.Attribute{
					Name:        "storage_usage_bytes",
					Description: `The amount of storage used in the registry in bytes.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the registry was created ## Import Container Registries can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_container_registry.myregistry registryname ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_container_registry_docker_credentials",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "registry_name",
					Description: `(Required) The name of the container registry.`,
				},
				resource.Attribute{
					Name:        "write",
					Description: `(Optional) Allow for write access to the container registry. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "expiry_seconds",
					Description: `(Optional) The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "registry_name",
					Description: `The name of the container registry`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "registry_name",
					Description: `The name of the container registry`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_custom_image",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the Custom Image.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) A URL from which the custom Linux virtual machine image may be retrieved.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Required) A list of regions. (Currently only one is supported).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An optional description for the image.`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `An optional distribution name for the image. Valid values are documented [here](https://docs.digitalocean.com/reference/api/api-reference/#operation/create_custom_image)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of optional tags for the image. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_cluster",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database cluster.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) Database engine used by the cluster (ex. ` + "`" + `pg` + "`" + ` for PostreSQL, ` + "`" + `mysql` + "`" + ` for MySQL, ` + "`" + `redis` + "`" + ` for Redis, or ` + "`" + `mongodb` + "`" + ` for MongoDB).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Database Droplet size associated with the cluster (ex. ` + "`" + `db-s-1vcpu-1gb` + "`" + `). See here for a [list of valid size slugs](https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) DigitalOcean region where the cluster will reside.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Required) Number of nodes that will be included in the cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Engine version used by the cluster (ex. ` + "`" + `14` + "`" + ` for PostgreSQL 14). When this value is changed, a call to the [Upgrade major Version for a Database](https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_update_major_version) API operation is made with the new version.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tag names to be applied to the database cluster.`,
				},
				resource.Attribute{
					Name:        "private_network_uuid",
					Description: `(Optional) The ID of the VPC where the database cluster will be located.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The ID of the project that the database cluster is assigned to. If excluded when creating a new database cluster, it will be assigned to your default project.`,
				},
				resource.Attribute{
					Name:        "eviction_policy",
					Description: `(Optional) A string specifying the eviction policy for a Redis cluster. Valid values are: ` + "`" + `noeviction` + "`" + `, ` + "`" + `allkeys_lru` + "`" + `, ` + "`" + `allkeys_random` + "`" + `, ` + "`" + `volatile_lru` + "`" + `, ` + "`" + `volatile_random` + "`" + `, or ` + "`" + `volatile_ttl` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sql_mode",
					Description: `(Optional) A comma separated string specifying the SQL modes for a MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `(Optional) Defines when the automatic maintenance should be performed for the database cluster. ` + "`" + `maintenance_window` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `(Required) The day of the week on which to apply maintenance updates.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `(Required) The hour in UTC at which maintenance updates will be applied in 24 hour format.`,
				},
				resource.Attribute{
					Name:        "backup_restore",
					Description: `(Optional) Create a new database cluster based on a backup of an existing cluster. ` + "`" + `backup_restore` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) The name of an existing database cluster from which the backup will be restored.`,
				},
				resource.Attribute{
					Name:        "backup_created_at",
					Description: `(Optional) The timestamp of an existing database cluster backup in ISO8601 combined date and time format. The most recent backup will be used if excluded. This resource supports [customized create timeouts](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts). The default timeout is 30 minutes. ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database cluster.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the database cluster.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Database cluster's hostname.`,
				},
				resource.Attribute{
					Name:        "private_host",
					Description: `Same as ` + "`" + `host` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Network port that the database cluster is listening on.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The full URI for connecting to the database cluster.`,
				},
				resource.Attribute{
					Name:        "private_uri",
					Description: `Same as ` + "`" + `uri` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `Name of the cluster's default database.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Username for the cluster's default user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the cluster's default user. ## Import Database clusters can be imported using the ` + "`" + `id` + "`" + ` returned from DigitalOcean, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_cluster.mycluster 245bcfd0-7f31-4ce6-a2bc-475a116cca97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database cluster.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the database cluster.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Database cluster's hostname.`,
				},
				resource.Attribute{
					Name:        "private_host",
					Description: `Same as ` + "`" + `host` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Network port that the database cluster is listening on.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The full URI for connecting to the database cluster.`,
				},
				resource.Attribute{
					Name:        "private_uri",
					Description: `Same as ` + "`" + `uri` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `Name of the cluster's default database.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Username for the cluster's default user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the cluster's default user. ## Import Database clusters can be imported using the ` + "`" + `id` + "`" + ` returned from DigitalOcean, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_cluster.mycluster 245bcfd0-7f31-4ce6-a2bc-475a116cca97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_connection_pool",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the source database cluster. Note: This must be a PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the database connection pool.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The desired size of the PGBouncer connection pool.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `(Required) The database for use with the connection pool.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) The name of the database user for use with the connection pool. When excluded, all sessions connect to the database as the inbound user. ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database connection pool.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The hostname used to connect to the database connection pool.`,
				},
				resource.Attribute{
					Name:        "private_host",
					Description: `Same as ` + "`" + `host` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Network port that the database connection pool is listening on.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The full URI for connecting to the database connection pool.`,
				},
				resource.Attribute{
					Name:        "private_uri",
					Description: `Same as ` + "`" + `uri` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the connection pool's user. ## Import Database connection pools can be imported using the ` + "`" + `id` + "`" + ` of the source database cluster and the ` + "`" + `name` + "`" + ` of the connection pool joined with a comma. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_connection_pool.pool-01 245bcfd0-7f31-4ce6-a2bc-475a116cca97,pool-01 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database connection pool.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The hostname used to connect to the database connection pool.`,
				},
				resource.Attribute{
					Name:        "private_host",
					Description: `Same as ` + "`" + `host` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Network port that the database connection pool is listening on.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The full URI for connecting to the database connection pool.`,
				},
				resource.Attribute{
					Name:        "private_uri",
					Description: `Same as ` + "`" + `uri` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the connection pool's user. ## Import Database connection pools can be imported using the ` + "`" + `id` + "`" + ` of the source database cluster and the ` + "`" + `name` + "`" + ` of the connection pool joined with a comma. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_connection_pool.pool-01 245bcfd0-7f31-4ce6-a2bc-475a116cca97,pool-01 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_db",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the original source database cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the database. ## Attributes Reference Only the above arguments are exported. ## Import Database can be imported using the ` + "`" + `id` + "`" + ` of the source database cluster and the ` + "`" + `name` + "`" + ` of the database joined with a comma. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_db.database-example 245bcfd0-7f31-4ce6-a2bc-475a116cca97,foobar ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_firewall",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the target database cluster.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) A rule specifying a resource allowed to access the database cluster. The following arguments must be specified: - ` + "`" + `type` + "`" + ` - (Required) The type of resource that the firewall rule allows to access the database cluster. The possible values are: ` + "`" + `droplet` + "`" + `, ` + "`" + `k8s` + "`" + `, ` + "`" + `ip_addr` + "`" + `, ` + "`" + `tag` + "`" + `, or ` + "`" + `app` + "`" + `. - ` + "`" + `value` + "`" + ` - (Required) The ID of the specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster. ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `A unique identifier for the firewall rule.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the firewall rule was created. ## Import Database firewalls can be imported using the ` + "`" + `id` + "`" + ` of the target database cluster For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_firewall.example-fw 5f55c6cd-863b-4907-99b8-7e09b0275d54 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `A unique identifier for the firewall rule.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the firewall rule was created. ## Import Database firewalls can be imported using the ` + "`" + `id` + "`" + ` of the target database cluster For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_firewall.example-fw 5f55c6cd-863b-4907-99b8-7e09b0275d54 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_replica",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the original source database cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the database replica.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Database Droplet size associated with the replica (ex. ` + "`" + `db-s-1vcpu-1gb` + "`" + `). Note that when resizing an existing replica, its size can only be increased. Decreasing its size is not supported.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) DigitalOcean region where the replica will reside.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tag names to be applied to the database replica.`,
				},
				resource.Attribute{
					Name:        "private_network_uuid",
					Description: `(Optional) The ID of the VPC where the database replica will be located. ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database replica created by Terraform.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the database replica. The uuid can be used to reference the database replica as the target database cluster in other resources. See example "Create firewall rule for database replica" above.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Database replica's hostname.`,
				},
				resource.Attribute{
					Name:        "private_host",
					Description: `Same as ` + "`" + `host` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Network port that the database replica is listening on.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The full URI for connecting to the database replica.`,
				},
				resource.Attribute{
					Name:        "private_uri",
					Description: `Same as ` + "`" + `uri` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `Name of the replica's default database.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Username for the replica's default user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the replica's default user. ## Import Database replicas can be imported using the ` + "`" + `id` + "`" + ` of the source database cluster and the ` + "`" + `name` + "`" + ` of the replica joined with a comma. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_replica.read-replica 245bcfd0-7f31-4ce6-a2bc-475a116cca97,read-replica ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database replica created by Terraform.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the database replica. The uuid can be used to reference the database replica as the target database cluster in other resources. See example "Create firewall rule for database replica" above.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Database replica's hostname.`,
				},
				resource.Attribute{
					Name:        "private_host",
					Description: `Same as ` + "`" + `host` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Network port that the database replica is listening on.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The full URI for connecting to the database replica.`,
				},
				resource.Attribute{
					Name:        "private_uri",
					Description: `Same as ` + "`" + `uri` + "`" + `, but only accessible from resources within the account and in the same region.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `Name of the replica's default database.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Username for the replica's default user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the replica's default user. ## Import Database replicas can be imported using the ` + "`" + `id` + "`" + ` of the source database cluster and the ` + "`" + `name` + "`" + ` of the replica joined with a comma. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_replica.read-replica 245bcfd0-7f31-4ce6-a2bc-475a116cca97,read-replica ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the original source database cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the database user.`,
				},
				resource.Attribute{
					Name:        "mysql_auth_plugin",
					Description: `(Optional) The authentication method to use for connections to the MySQL user account. The valid values are ` + "`" + `mysql_native_password` + "`" + ` or ` + "`" + `caching_sha2_password` + "`" + ` (this is the default). ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role for the database user. The value will be either "primary" or "normal".`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the database user. ## Import Database user can be imported using the ` + "`" + `id` + "`" + ` of the source database cluster and the ` + "`" + `name` + "`" + ` of the user joined with a comma. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_user.user-example 245bcfd0-7f31-4ce6-a2bc-475a116cca97,foobar ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role",
					Description: `Role for the database user. The value will be either "primary" or "normal".`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for the database user. ## Import Database user can be imported using the ` + "`" + `id` + "`" + ` of the source database cluster and the ` + "`" + `name` + "`" + ` of the user joined with a comma. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_database_user.user-example 245bcfd0-7f31-4ce6-a2bc-475a116cca97,foobar ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_domain",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the domain`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address of the domain. If specified, this IP is used to created an initial A record for the domain. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the domain`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the domain`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL value of the domain ## Import Domains can be imported using the ` + "`" + `domain name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_domain.mydomain mytestdomain.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the domain`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the domain`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL value of the domain ## Import Domains can be imported using the ` + "`" + `domain name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_domain.mydomain mytestdomain.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_droplet",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image",
					Description: `(Required) The Droplet image ID or slug. This could be either image ID or droplet snapshot ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Droplet name.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the Droplet will be created.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The unique slug that indentifies the type of Droplet. You can find a list of available slugs on [DigitalOcean API documentation](https://docs.digitalocean.com/reference/api/api-reference/#tag/Sizes).`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `(Optional) Boolean controlling if backups are made. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional) Boolean controlling whether monitoring agent is installed. Defaults to false. If set to ` + "`" + `true` + "`" + `, you can configure monitor alert policies [monitor alert resource](/providers/digitalocean/digitalocean/latest/docs/resources/monitor_alert)`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) Boolean controlling if IPv6 is enabled. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `(Optional) The ID of the VPC where the Droplet will be located.`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `(Optional) A list of SSH key IDs or fingerprints to enable in the format ` + "`" + `[12345, 123456]` + "`" + `. To retrieve this info, use the [DigitalOcean API](https://docs.digitalocean.com/reference/api/api-reference/#tag/SSH-Keys) or CLI (` + "`" + `doctl compute ssh-key list` + "`" + `). Once a Droplet is created keys can not be added or removed via this provider. Modifying this field will prompt you to destroy and recreate the Droplet.`,
				},
				resource.Attribute{
					Name:        "resize_disk",
					Description: `(Optional) Boolean controlling whether to increase the disk size when resizing a Droplet. It defaults to ` + "`" + `true` + "`" + `. When set to ` + "`" + `false` + "`" + `, only the Droplet's RAM and CPU will be resized.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of the tags to be applied to this Droplet.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Droplet`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the Droplet`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the Droplet`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The image of the Droplet`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Is IPv6 enabled`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The IPv4 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address_private",
					Description: `The private networking IPv4 address`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Is the Droplet locked`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `Is private networking enabled`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `Droplet hourly price`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `Droplet monthly price`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The instance size`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The size of the instance's disk in GB`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of the instance's virtual CPUs`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Droplet`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags associated with the Droplet`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `A list of the attached block storage volumes ## Import Droplets can be imported using the Droplet ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_droplet.mydroplet 100823 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Droplet`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the Droplet`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the Droplet`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The image of the Droplet`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Is IPv6 enabled`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The IPv4 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address_private",
					Description: `The private networking IPv4 address`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Is the Droplet locked`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `Is private networking enabled`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `Droplet hourly price`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `Droplet monthly price`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The instance size`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The size of the instance's disk in GB`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of the instance's virtual CPUs`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Droplet`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags associated with the Droplet`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `A list of the attached block storage volumes ## Import Droplets can be imported using the Droplet ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_droplet.mydroplet 100823 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_droplet_snapshot",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the Droplet snapshot.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `(Required) The ID of the Droplet from which the snapshot will be taken. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the Droplet snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a Droplet to be created based on this snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the droplet snapshot is available.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the Droplet snapshot in gigabytes. ## Import Droplet Snapshots can be imported using the ` + "`" + `snapshot id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_droplet_snapshot.mysnapshot 123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the Droplet snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a Droplet to be created based on this snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the droplet snapshot is available.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the Droplet snapshot in gigabytes. ## Import Droplet Snapshots can be imported using the ` + "`" + `snapshot id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_droplet_snapshot.mysnapshot 123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_firewall",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Firewall name`,
				},
				resource.Attribute{
					Name:        "inbound_rule",
					Description: `(Optional) The inbound access rule block for the Firewall. The ` + "`" + `inbound_rule` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "outbound_rule",
					Description: `(Optional) The outbound access rule block for the Firewall. The ` + "`" + `outbound_rule` + "`" + ` block is documented below. ` + "`" + `inbound_rule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The type of traffic to be allowed. This may be one of "tcp", "udp", or "icmp".`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional) The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "1-65535" to open all ports for a protocol. Required for when protocol is ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Optional) An array of strings containing the IPv4 addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_droplet_ids",
					Description: `(Optional) An array containing the IDs of the Droplets from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_tags",
					Description: `(Optional) An array containing the names of Tags corresponding to groups of Droplets from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_load_balancer_uids",
					Description: `(Optional) An array containing the IDs of the Load Balancers from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_kubernetes_ids",
					Description: `(Optional) An array containing the IDs of the Kubernetes clusters from which the inbound traffic will be accepted. ` + "`" + `outbound_rule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The type of traffic to be allowed. This may be one of "tcp", "udp", or "icmp".`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional) The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "1-65535" to open all ports for a protocol. Required for when protocol is ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Optional) An array of strings containing the IPv4 addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the outbound traffic will be allowed.`,
				},
				resource.Attribute{
					Name:        "destination_droplet_ids",
					Description: `(Optional) An array containing the IDs of the Droplets to which the outbound traffic will be allowed.`,
				},
				resource.Attribute{
					Name:        "destination_kubernetes_ids",
					Description: `(Optional) An array containing the IDs of the Kubernetes clusters to which the outbound traffic will be allowed.`,
				},
				resource.Attribute{
					Name:        "destination_tags",
					Description: `(Optional) An array containing the names of Tags corresponding to groups of Droplets to which the outbound traffic will be allowed.`,
				},
				resource.Attribute{
					Name:        "destination_load_balancer_uids",
					Description: `(Optional) An array containing the IDs of the Load Balancers to which the outbound traffic will be allowed. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Firewall.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `A status string indicating the current state of the Firewall. This can be "waiting", "succeeded", or "failed".`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the Firewall was created.`,
				},
				resource.Attribute{
					Name:        "pending_changes",
					Description: `An list of object containing the fields, "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Firewall.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `The list of the IDs of the Droplets assigned to the Firewall.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The names of the Tags assigned to the Firewall.`,
				},
				resource.Attribute{
					Name:        "inbound_rule",
					Description: `The inbound access rule block for the Firewall.`,
				},
				resource.Attribute{
					Name:        "outbound_rule",
					Description: `The outbound access rule block for the Firewall. ## Import Firewalls can be imported using the firewall ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_firewall.myfirewall b8ecd2ab-2267-4a5e-8692-cbf1d32583e3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Firewall.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `A status string indicating the current state of the Firewall. This can be "waiting", "succeeded", or "failed".`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the Firewall was created.`,
				},
				resource.Attribute{
					Name:        "pending_changes",
					Description: `An list of object containing the fields, "droplet_id", "removing", and "status". It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Firewall.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `The list of the IDs of the Droplets assigned to the Firewall.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The names of the Tags assigned to the Firewall.`,
				},
				resource.Attribute{
					Name:        "inbound_rule",
					Description: `The inbound access rule block for the Firewall.`,
				},
				resource.Attribute{
					Name:        "outbound_rule",
					Description: `The outbound access rule block for the Firewall. ## Import Firewalls can be imported using the firewall ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_firewall.myfirewall b8ecd2ab-2267-4a5e-8692-cbf1d32583e3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_floating_ip",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region that the Floating IP is reserved to.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `(Optional) The ID of Droplet that the Floating IP will be assigned to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address of the resource`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the floating ip ## Import Floating IPs can be imported using the ` + "`" + `ip` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_floating_ip.myip 192.168.0.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address of the resource`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the floating ip ## Import Floating IPs can be imported using the ` + "`" + `ip` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_floating_ip.myip 192.168.0.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_floating_ip_assignment",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The Floating IP to assign to the Droplet.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `(Optional) The ID of Droplet that the Floating IP will be assigned to. ## Import Floating IP assignments can be imported using the Floating IP itself and the ` + "`" + `id` + "`" + ` of the Droplet joined with a comma. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_floating_ip_assignment.foobar 192.0.2.1,123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_kubernetes_cluster",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The slug identifier for the region where the Kubernetes cluster will be created.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions ` + "`" + `doctl kubernetes options versions` + "`" + `. (`,
				},
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `(Optional) The ID of the VPC where the Kubernetes cluster will be located.`,
				},
				resource.Attribute{
					Name:        "auto_upgrade",
					Description: `(Optional) A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.`,
				},
				resource.Attribute{
					Name:        "surge_upgrade",
					Description: `(Optional) Enable/disable surge upgrades for a cluster. Default: false`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `(Optional) Enable/disable the high availability control plane for a cluster. High availability can only be set when creating a cluster. Any update will create a new cluster. Default: false`,
				},
				resource.Attribute{
					Name:        "registry_integration",
					Description: `(optional) Enables or disables the DigitalOcean container registry integration for the cluster. This requires that a container registry has first been created for the account. Default: false`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `(Required) A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the ` + "`" + `digitalocean_kubernetes_node_pool` + "`" + ` resource. The following arguments may be specified: - ` + "`" + `name` + "`" + ` - (Required) A name for the node pool. - ` + "`" + `size` + "`" + ` - (Required) The slug identifier for the type of Droplet to be used as workers in the node pool. - ` + "`" + `node_count` + "`" + ` - (Optional) The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value. - ` + "`" + `auto_scale` + "`" + ` - (Optional) Enable auto-scaling of the number of nodes in the node pool within the given min/max range. - ` + "`" + `min_nodes` + "`" + ` - (Optional) If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to. - ` + "`" + `max_nodes` + "`" + ` - (Optional) If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to. - ` + "`" + `tags` + "`" + ` - (Optional) A list of tag names applied to the node pool. - ` + "`" + `labels` + "`" + ` - (Optional) A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tag names to be applied to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `(Optional) A block representing the cluster's maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. ` + "`" + `auto_upgrade` + "`" + ` must be set to ` + "`" + `true` + "`" + ` for this to have an effect. - ` + "`" + `day` + "`" + ` - (Required) The day of the maintenance window policy. May be one of "monday" through "sunday", or "any" to indicate an arbitrary week day. - ` + "`" + `start_time` + "`" + ` (Required) The start time in UTC of the maintenance window policy in 24-hour clock format / HH:MM notation (e.g., 15:00). This resource supports [customized create timeouts](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts). The default timeout is 30 minutes. ## Attributes Reference In addition to the arguments listed above, the following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_subnet",
					Description: `The range of IP addresses in the overlay network of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "service_subnet",
					Description: `The range of assignable IP addresses for services running in the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The public IPv4 address of the Kubernetes master node. This will not be set if high availability is configured on the cluster (v1.21+)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The base URL of the API server on the Kubernetes master node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the Kubernetes cluster was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date and time when the Kubernetes cluster was last updated.`,
				},
				resource.Attribute{
					Name:        "auto_upgrade",
					Description: `A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.`,
				},
				resource.Attribute{
					Name:        "kube_config.0",
					Description: `A representation of the Kubernetes cluster's kubeconfig with the following attributes: - ` + "`" + `raw_config` + "`" + ` - The full contents of the Kubernetes cluster's kubeconfig file. - ` + "`" + `host` + "`" + ` - The URL of the API server on the Kubernetes master node. - ` + "`" + `cluster_ca_certificate` + "`" + ` - The base64 encoded public certificate for the cluster's certificate authority. - ` + "`" + `token` + "`" + ` - The DigitalOcean API access token used by clients to access the cluster. - ` + "`" + `client_key` + "`" + ` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster. - ` + "`" + `client_certificate` + "`" + ` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster. - ` + "`" + `expires_at` + "`" + ` - The date and time when the credentials will expire and need to be regenerated.`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `In addition to the arguments provided, these additional attributes about the cluster's default node pool are exported: - ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node pool. - ` + "`" + `actual_node_count` + "`" + ` - A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled. - ` + "`" + `nodes` + "`" + ` - A list of nodes in the pool. Each node exports the following attributes: + ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. + ` + "`" + `name` + "`" + ` - The auto-generated name for the node. + ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. + ` + "`" + `droplet_id` + "`" + ` - The id of the node's droplet + ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. + ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated. - ` + "`" + `taint` + "`" + ` - A block representing a taint applied to all nodes in the pool. Each taint exports the following attributes (taints must be unique by key and effect pair): + ` + "`" + `key` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. + ` + "`" + `value` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. + ` + "`" + `effect` + "`" + ` - How the node reacts to pods that it won't tolerate. Available effect values are: "NoSchedule", "PreferNoSchedule", "NoExecute".`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name (URN) for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `A block representing the cluster's maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. - ` + "`" + `day` + "`" + ` - The day of the maintenance window policy. May be one of "monday" through "sunday", or "any" to indicate an arbitrary week day. - ` + "`" + `duration` + "`" + ` A string denoting the duration of the service window, e.g., "04:00". - ` + "`" + `start_time` + "`" + ` The hour in UTC when maintenance updates will be applied, in 24 hour format (e.g. “16:00”). ## Import Before importing a Kubernetes cluster, the cluster's default node pool must be tagged with the ` + "`" + `terraform:default-node-pool` + "`" + ` tag. The provider will automatically add this tag if the cluster only has a single node pool. Clusters with more than one node pool, however, will require that you manually add the ` + "`" + `terraform:default-node-pool` + "`" + ` tag to the node pool that you intend to be the default node pool. Then the Kubernetes cluster and its default node pool can be imported using the cluster's ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_kubernetes_cluster.mycluster 1b8b2100-0e9f-4e8f-ad78-9eb578c2a0af ` + "`" + `` + "`" + `` + "`" + ` Additional node pools must be imported separately as ` + "`" + `digitalocean_kubernetes_cluster` + "`" + ` resources, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_kubernetes_node_pool.mynodepool 9d76f410-9284-4436-9633-4066852442c8 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_subnet",
					Description: `The range of IP addresses in the overlay network of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "service_subnet",
					Description: `The range of assignable IP addresses for services running in the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The public IPv4 address of the Kubernetes master node. This will not be set if high availability is configured on the cluster (v1.21+)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The base URL of the API server on the Kubernetes master node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the Kubernetes cluster was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date and time when the Kubernetes cluster was last updated.`,
				},
				resource.Attribute{
					Name:        "auto_upgrade",
					Description: `A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.`,
				},
				resource.Attribute{
					Name:        "kube_config.0",
					Description: `A representation of the Kubernetes cluster's kubeconfig with the following attributes: - ` + "`" + `raw_config` + "`" + ` - The full contents of the Kubernetes cluster's kubeconfig file. - ` + "`" + `host` + "`" + ` - The URL of the API server on the Kubernetes master node. - ` + "`" + `cluster_ca_certificate` + "`" + ` - The base64 encoded public certificate for the cluster's certificate authority. - ` + "`" + `token` + "`" + ` - The DigitalOcean API access token used by clients to access the cluster. - ` + "`" + `client_key` + "`" + ` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster. - ` + "`" + `client_certificate` + "`" + ` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster. - ` + "`" + `expires_at` + "`" + ` - The date and time when the credentials will expire and need to be regenerated.`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `In addition to the arguments provided, these additional attributes about the cluster's default node pool are exported: - ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node pool. - ` + "`" + `actual_node_count` + "`" + ` - A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled. - ` + "`" + `nodes` + "`" + ` - A list of nodes in the pool. Each node exports the following attributes: + ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. + ` + "`" + `name` + "`" + ` - The auto-generated name for the node. + ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. + ` + "`" + `droplet_id` + "`" + ` - The id of the node's droplet + ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. + ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated. - ` + "`" + `taint` + "`" + ` - A block representing a taint applied to all nodes in the pool. Each taint exports the following attributes (taints must be unique by key and effect pair): + ` + "`" + `key` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. + ` + "`" + `value` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. + ` + "`" + `effect` + "`" + ` - How the node reacts to pods that it won't tolerate. Available effect values are: "NoSchedule", "PreferNoSchedule", "NoExecute".`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name (URN) for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `A block representing the cluster's maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. - ` + "`" + `day` + "`" + ` - The day of the maintenance window policy. May be one of "monday" through "sunday", or "any" to indicate an arbitrary week day. - ` + "`" + `duration` + "`" + ` A string denoting the duration of the service window, e.g., "04:00". - ` + "`" + `start_time` + "`" + ` The hour in UTC when maintenance updates will be applied, in 24 hour format (e.g. “16:00”). ## Import Before importing a Kubernetes cluster, the cluster's default node pool must be tagged with the ` + "`" + `terraform:default-node-pool` + "`" + ` tag. The provider will automatically add this tag if the cluster only has a single node pool. Clusters with more than one node pool, however, will require that you manually add the ` + "`" + `terraform:default-node-pool` + "`" + ` tag to the node pool that you intend to be the default node pool. Then the Kubernetes cluster and its default node pool can be imported using the cluster's ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_kubernetes_cluster.mycluster 1b8b2100-0e9f-4e8f-ad78-9eb578c2a0af ` + "`" + `` + "`" + `` + "`" + ` Additional node pools must be imported separately as ` + "`" + `digitalocean_kubernetes_cluster` + "`" + ` resources, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_kubernetes_node_pool.mynodepool 9d76f410-9284-4436-9633-4066852442c8 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_kubernetes_node_pool",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the Kubernetes cluster to which the node pool is associated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the node pool.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The slug identifier for the type of Droplet to be used as workers in the node pool.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Optional) The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.`,
				},
				resource.Attribute{
					Name:        "auto_scale",
					Description: `(Optional) Enable auto-scaling of the number of nodes in the node pool within the given min/max range.`,
				},
				resource.Attribute{
					Name:        "min_nodes",
					Description: `(Optional) If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.`,
				},
				resource.Attribute{
					Name:        "max_nodes",
					Description: `(Optional) If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tag names to be applied to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).`,
				},
				resource.Attribute{
					Name:        "taint",
					Description: `(Optional) A list of taints applied to all nodes in the pool. This resource supports [customized create timeouts](https://www.terraform.io/docs/language/resources/syntax.html#operation-timeouts). The default timeout is 30 minutes. ## Attributes Reference In addition to the arguments listed above, the following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference the node pool.`,
				},
				resource.Attribute{
					Name:        "actual_node_count",
					Description: `A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `A list of nodes in the pool. Each node exports the following attributes: - ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. - ` + "`" + `name` + "`" + ` - The auto-generated name for the node. - ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. - ` + "`" + `droplet_id` + "`" + ` - The id of the node's droplet - ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. - ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated.`,
				},
				resource.Attribute{
					Name:        "taint",
					Description: `A list of taints applied to all nodes in the pool. Each taint exports the following attributes: - ` + "`" + `key` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. - ` + "`" + `value` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. - ` + "`" + `effect` + "`" + ` - How the node reacts to pods that it won't tolerate. Available effect values are: "NoSchedule", "PreferNoSchedule", "NoExecute". ## Import If you are importing an existing Kubernetes cluster with a single node pool, just import the cluster. Additional node pools can be imported by using their ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_kubernetes_node_pool.mynodepool 9d76f410-9284-4436-9633-4066852442c8 ` + "`" + `` + "`" + `` + "`" + ` Note: If the node pool has the ` + "`" + `terraform:default-node-pool` + "`" + ` tag, then it is a default node pool for an existing cluster. The provider will refuse to import the node pool in that case because the node pool is managed by the ` + "`" + `digitalocean_kubernetes_cluster` + "`" + ` resource and not by this ` + "`" + `digitalocean_kubernetes_node_pool` + "`" + ` resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference the node pool.`,
				},
				resource.Attribute{
					Name:        "actual_node_count",
					Description: `A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `A list of nodes in the pool. Each node exports the following attributes: - ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. - ` + "`" + `name` + "`" + ` - The auto-generated name for the node. - ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. - ` + "`" + `droplet_id` + "`" + ` - The id of the node's droplet - ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. - ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated.`,
				},
				resource.Attribute{
					Name:        "taint",
					Description: `A list of taints applied to all nodes in the pool. Each taint exports the following attributes: - ` + "`" + `key` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. - ` + "`" + `value` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. - ` + "`" + `effect` + "`" + ` - How the node reacts to pods that it won't tolerate. Available effect values are: "NoSchedule", "PreferNoSchedule", "NoExecute". ## Import If you are importing an existing Kubernetes cluster with a single node pool, just import the cluster. Additional node pools can be imported by using their ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_kubernetes_node_pool.mynodepool 9d76f410-9284-4436-9633-4066852442c8 ` + "`" + `` + "`" + `` + "`" + ` Note: If the node pool has the ` + "`" + `terraform:default-node-pool` + "`" + ` tag, then it is a default node pool for an existing cluster. The provider will refuse to import the node pool in that case because the node pool is managed by the ` + "`" + `digitalocean_kubernetes_cluster` + "`" + ` resource and not by this ` + "`" + `digitalocean_kubernetes_node_pool` + "`" + ` resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_loadbalancer",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Load Balancer name`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region to start in`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the Load Balancer. It must be either ` + "`" + `lb-small` + "`" + `, ` + "`" + `lb-medium` + "`" + `, or ` + "`" + `lb-large` + "`" + `. Defaults to ` + "`" + `lb-small` + "`" + `. Only one of ` + "`" + `size` + "`" + ` or ` + "`" + `size_unit` + "`" + ` may be provided.`,
				},
				resource.Attribute{
					Name:        "size_unit",
					Description: `(Optional) The size of the Load Balancer. It must be in the range (1, 100). Defaults to ` + "`" + `1` + "`" + `. Only one of ` + "`" + `size` + "`" + ` or ` + "`" + `size_unit` + "`" + ` may be provided.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Optional) The load balancing algorithm used to determine which backend Droplet will be selected by a client. It must be either ` + "`" + `round_robin` + "`" + ` or ` + "`" + `least_connections` + "`" + `. The default value is ` + "`" + `round_robin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "forwarding_rule",
					Description: `(Required) A list of ` + "`" + `forwarding_rule` + "`" + ` to be assigned to the Load Balancer. The ` + "`" + `forwarding_rule` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `(Optional) A ` + "`" + `healthcheck` + "`" + ` block to be assigned to the Load Balancer. The ` + "`" + `healthcheck` + "`" + ` block is documented below. Only 1 healthcheck is allowed.`,
				},
				resource.Attribute{
					Name:        "sticky_sessions",
					Description: `(Optional) A ` + "`" + `sticky_sessions` + "`" + ` block to be assigned to the Load Balancer. The ` + "`" + `sticky_sessions` + "`" + ` block is documented below. Only 1 sticky_sessions block is allowed.`,
				},
				resource.Attribute{
					Name:        "redirect_http_to_https",
					Description: `(Optional) A boolean value indicating whether HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_proxy_protocol",
					Description: `(Optional) A boolean value indicating whether PROXY Protocol should be used to pass information from connecting client requests to the backend service. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_backend_keepalive",
					Description: `(Optional) A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_idle_timeout_seconds",
					Description: `(Optional) Specifies the idle timeout for HTTPS connections on the load balancer in seconds.`,
				},
				resource.Attribute{
					Name:        "disable_lets_encrypt_dns_records",
					Description: `(Optional) A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project.`,
				},
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `(Optional) The ID of the VPC where the load balancer will be located.`,
				},
				resource.Attribute{
					Name:        "entry_protocol",
					Description: `(Required) The protocol used for traffic to the Load Balancer. The possible values are: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `http2` + "`" + `, ` + "`" + `http3` + "`" + `, ` + "`" + `tcp` + "`" + `, or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "entry_port",
					Description: `(Required) An integer representing the port on which the Load Balancer instance will listen.`,
				},
				resource.Attribute{
					Name:        "target_protocol",
					Description: `(Required) The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `http2` + "`" + `, ` + "`" + `tcp` + "`" + `, or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `(Required) An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(Optional) The unique name of the TLS certificate to be used for SSL termination.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tls_passthrough",
					Description: `(Optional) A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is ` + "`" + `false` + "`" + `. ` + "`" + `sticky_sessions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are ` + "`" + `cookies` + "`" + ` or ` + "`" + `none` + "`" + `. If not specified, the default value is ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Optional) The name to be used for the cookie sent to the client. This attribute is required when using ` + "`" + `cookies` + "`" + ` for the sticky sessions type.`,
				},
				resource.Attribute{
					Name:        "cookie_ttl_seconds",
					Description: `(Optional) The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using ` + "`" + `cookies` + "`" + ` for the sticky sessions type. ` + "`" + `healthcheck` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol used for health checks sent to the backend Droplets. The possible values are ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) An integer representing the port on the backend Droplets on which the health check will attempt a connection.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path on the backend Droplets to which the Load Balancer instance will send a request.`,
				},
				resource.Attribute{
					Name:        "check_interval_seconds",
					Description: `(Optional) The number of seconds between two consecutive health checks. If not specified, the default value is ` + "`" + `10` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_timeout_seconds",
					Description: `(Optional) The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool. If not specified, the default value is ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool. If not specified, the default value is ` + "`" + `5` + "`" + `. ` + "`" + `firewall` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "deny",
					Description: `(Optional) A list of strings describing deny rules. Must be colon delimited strings of the form ` + "`" + `{type}:{source}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "allow",
					Description: `(Optional) A list of strings describing allow rules. Must be colon delimited strings of the form ` + "`" + `{type}:{source}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the Load Balancer ## Import Load Balancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_loadbalancer.myloadbalancer 4de7ac8b-495b-4884-9a69-1050c6793cd6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the Load Balancer ## Import Load Balancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_loadbalancer.myloadbalancer 4de7ac8b-495b-4884-9a69-1050c6793cd6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_monitor_alert",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alerts",
					Description: `(Required) How to send notifications about the alerts. This is a list with one element, . Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can read more in [Slack's documentation](https://slack.com/intl/en-dk/help/articles/222386767-Manage-app-installation-settings-for-your-workspace)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the alert.`,
				},
				resource.Attribute{
					Name:        "compare",
					Description: `(Required) The comparison for ` + "`" + `value` + "`" + `. This may be either ` + "`" + `GreaterThan` + "`" + ` or ` + "`" + `LessThan` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the alert. This may be one of ` + "`" + `v1/insights/droplet/load_1` + "`" + `, ` + "`" + `v1/insights/droplet/load_5` + "`" + `, ` + "`" + `v1/insights/droplet/load_15` + "`" + `, ` + "`" + `v1/insights/droplet/memory_utilization_percent` + "`" + `, ` + "`" + `v1/insights/droplet/disk_utilization_percent` + "`" + `, ` + "`" + `v1/insights/droplet/cpu` + "`" + `, ` + "`" + `v1/insights/droplet/disk_read` + "`" + `, ` + "`" + `v1/insights/droplet/disk_write` + "`" + `, ` + "`" + `v1/insights/droplet/public_outbound_bandwidth` + "`" + `, ` + "`" + `v1/insights/droplet/public_inbound_bandwidth` + "`" + `, ` + "`" + `v1/insights/droplet/private_outbound_bandwidth` + "`" + `, ` + "`" + `v1/insights/droplet/private_inbound_bandwidth` + "`" + `, ` + "`" + `v1/insights/lbaas/avg_cpu_utilization_percent` + "`" + `, ` + "`" + `v1/insights/lbaas/connection_utilization_percent` + "`" + `, ` + "`" + `v1/insights/lbaas/droplet_health` + "`" + `, ` + "`" + `v1/insights/lbaas/tls_connections_per_second_utilization_percent` + "`" + `, ` + "`" + `v1/insights/lbaas/increase_in_http_error_rate_percentage_5xx` + "`" + `, ` + "`" + `v1/insights/lbaas/increase_in_http_error_rate_percentage_4xx` + "`" + `, ` + "`" + `v1/insights/lbaas/increase_in_http_error_rate_count_5xx` + "`" + `, ` + "`" + `v1/insights/lbaas/increase_in_http_error_rate_count_4xx` + "`" + `, ` + "`" + `v1/insights/lbaas/high_http_request_response_time` + "`" + `, ` + "`" + `v1/insights/lbaas/high_http_request_response_time_50p` + "`" + `, ` + "`" + `v1/insights/lbaas/high_http_request_response_time_95p` + "`" + `, ` + "`" + `v1/insights/lbaas/high_http_request_response_time_99p` + "`" + `, ` + "`" + `v1/dbaas/alerts/load_15_alerts` + "`" + `, ` + "`" + `v1/dbaas/alerts/cpu_alerts` + "`" + `, ` + "`" + `v1/dbaas/alerts/memory_utilization_alerts` + "`" + `, or ` + "`" + `v1/dbaas/alerts/disk_utilization_alerts` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) The status of the alert.`,
				},
				resource.Attribute{
					Name:        "entities",
					Description: `A list of IDs for the resources to which the alert policy applies.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags. When an included tag is added to a resource, the alert policy will apply to it.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number. DigitalOcean will show the correct unit in the web panel.`,
				},
				resource.Attribute{
					Name:        "window",
					Description: `(Required) The time frame of the alert. Either ` + "`" + `5m` + "`" + `, ` + "`" + `10m` + "`" + `, ` + "`" + `30m` + "`" + `, or ` + "`" + `1h` + "`" + `. ## Attributes Reference The following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The uuid of the alert.`,
				},
				resource.Attribute{
					Name:        "window",
					Description: `The time frame of the alert.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `The status of the alert.`,
				},
				resource.Attribute{
					Name:        "entities",
					Description: `The resources for which the alert policy applies`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the alert.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The percentage to start alerting at.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for the alert.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The percentage to start alerting at, e.g., 90.`,
				},
				resource.Attribute{
					Name:        "alerts",
					Description: `The notification policies of the alert policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the alert. ## Import Monitor alerts can be imported using the monitor alert ` + "`" + `uuid` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import digitalocean_monitor_alert.cpu_alert b8ecd2ab-2267-4a5e-8692-cbf1d32583e3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `The uuid of the alert.`,
				},
				resource.Attribute{
					Name:        "window",
					Description: `The time frame of the alert.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `The status of the alert.`,
				},
				resource.Attribute{
					Name:        "entities",
					Description: `The resources for which the alert policy applies`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the alert.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The percentage to start alerting at.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for the alert.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The percentage to start alerting at, e.g., 90.`,
				},
				resource.Attribute{
					Name:        "alerts",
					Description: `The notification policies of the alert policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the alert. ## Import Monitor alerts can be imported using the monitor alert ` + "`" + `uuid` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import digitalocean_monitor_alert.cpu_alert b8ecd2ab-2267-4a5e-8692-cbf1d32583e3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_project",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Project`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) the description of the project`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `(Optional) the purpose of the project, (Default: "Web Application")`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) the environment of the project's resources. The possible values are: ` + "`" + `Development` + "`" + `, ` + "`" + `Staging` + "`" + `, ` + "`" + `Production` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `a list of uniform resource names (URNs) for the resources associated with the project`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) a boolean indicating whether or not the project is the default project. (Default: "false") ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the project`,
				},
				resource.Attribute{
					Name:        "owner_uuid",
					Description: `the unique universal identifier of the project owner.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `the id of the project owner.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `the date and time when the project was created, (ISO8601)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `the date and time when the project was last updated, (ISO8601) ## Import Projects can be imported using the ` + "`" + `id` + "`" + ` returned from DigitalOcean, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_project.myproject 245bcfd0-7f31-4ce6-a2bc-475a116cca97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the project`,
				},
				resource.Attribute{
					Name:        "owner_uuid",
					Description: `the unique universal identifier of the project owner.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `the id of the project owner.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `the date and time when the project was created, (ISO8601)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `the date and time when the project was last updated, (ISO8601) ## Import Projects can be imported using the ` + "`" + `id` + "`" + ` returned from DigitalOcean, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_project.myproject 245bcfd0-7f31-4ce6-a2bc-475a116cca97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_project_resources",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) the ID of the project`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) a list of uniform resource names (URNs) for the resources associated with the project ## Attributes Reference No additional attributes are exported. ## Import Importing this resource is not supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_record",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of record. Must be one of ` + "`" + `A` + "`" + `, ` + "`" + `AAAA` + "`" + `, ` + "`" + `CAA` + "`" + `, ` + "`" + `CNAME` + "`" + `, ` + "`" + `MX` + "`" + `, ` + "`" + `NS` + "`" + `, ` + "`" + `TXT` + "`" + `, or ` + "`" + `SRV` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain to add the record to.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the record.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The hostname of the record. Use ` + "`" + `@` + "`" + ` for records on domain's name itself.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port of the record. Only valid when type is ` + "`" + `SRV` + "`" + `. Must be between 1 and 65535.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the record. Only valid when type is ` + "`" + `MX` + "`" + ` or ` + "`" + `SRV` + "`" + `. Must be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The weight of the record. Only valid when type is ` + "`" + `SRV` + "`" + `. Must be between 0 and 65535.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live for the record, in seconds. Must be at least 0. Defaults to 1800.`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `(Optional) The flags of the record. Only valid when type is ` + "`" + `CAA` + "`" + `. Must be between 0 and 255.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) The tag of the record. Only valid when type is ` + "`" + `CAA` + "`" + `. Must be one of ` + "`" + `issue` + "`" + `, ` + "`" + `issuewild` + "`" + `, or ` + "`" + `iodef` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the record ## Import Records can be imported using the domain name and record ` + "`" + `id` + "`" + ` when joined with a comma. See the following example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_record.example_record example.com,12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the record ## Import Records can be imported using the domain name and record ` + "`" + `id` + "`" + ` when joined with a comma. See the following example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_record.example_record example.com,12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_reserved_ip",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region that the reserved IP is reserved to.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `(Optional) The ID of Droplet that the reserved IP will be assigned to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address of the resource`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the reserved ip ## Import Reserved IPs can be imported using the ` + "`" + `ip` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_reserved_ip.myip 192.168.0.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address of the resource`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the reserved ip ## Import Reserved IPs can be imported using the ` + "`" + `ip` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_reserved_ip.myip 192.168.0.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_reserved_ip_assignment",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The reserved IP to assign to the Droplet.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `(Optional) The ID of Droplet that the reserved IP will be assigned to. ## Import Reserved IP assignments can be imported using the reserved IP itself and the ` + "`" + `id` + "`" + ` of the Droplet joined with a comma. For example: ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_reserved_ip_assignment.foobar 192.0.2.1,123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_bucket",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the bucket resides (Defaults to ` + "`" + `nyc3` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Canned ACL applied on bucket creation (` + "`" + `private` + "`" + ` or ` + "`" + `public-read` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "cors_rule",
					Description: `(Optional) A rule of Cross-Origin Resource Sharing (documented below).`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `(Optional) A configuration of object lifecycle management (documented below).`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `(Optional) A state of versioning (documented below)`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `Unless ` + "`" + `true` + "`" + `, the bucket will only be destroyed if empty (Defaults to ` + "`" + `false` + "`" + `) The ` + "`" + `cors_rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `(Optional) A list of headers that will be included in the CORS preflight request's ` + "`" + `Access-Control-Request-Headers` + "`" + `. A header may contain one wildcard (e.g. ` + "`" + `x-amz-`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `(Required) A list of HTTP methods (e.g. ` + "`" + `GET` + "`" + `) which are allowed from the specified origin.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `(Required) A list of hosts from which requests using the specified methods are allowed. A host may contain one wildcard (e.g. http://`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `(Optional) The time in seconds that browser can cache the response for a preflight request. The ` + "`" + `lifecycle_rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier for the rule.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Object key prefix identifying one or more objects to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Specifies lifecycle rule status.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) Specifies a time period after which applicable objects expire (documented below).`,
				},
				resource.Attribute{
					Name:        "noncurrent_version_expiration",
					Description: `(Optional) Specifies when non-current object versions expire (documented below). At least one of ` + "`" + `expiration` + "`" + ` or ` + "`" + `noncurrent_version_expiration` + "`" + ` must be specified. The ` + "`" + `expiration` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) Specifies the date/time after which you want applicable objects to expire. The argument uses RFC3339 format, e.g. "2020-03-22T15:03:55Z" or parts thereof e.g. "2019-02-28".`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional) Specifies the number of days after object creation when the applicable objects will expire.`,
				},
				resource.Attribute{
					Name:        "expired_object_delete_marker",
					Description: `(Optional) On a versioned bucket (versioning-enabled or versioning-suspended bucket), setting this to true directs Spaces to delete expired object delete markers. The ` + "`" + `noncurrent_version_expiration` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Required) Specifies the number of days after which an object's non-current versions expire. The ` + "`" + `versioning` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the bucket`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The name of the region`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The FQDN of the bucket without the bucket name (e.g. nyc3.digitaloceanspaces.com) ## Import Buckets can be imported using the ` + "`" + `region` + "`" + ` and ` + "`" + `name` + "`" + ` attributes (delimited by a comma): ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_spaces_bucket.foobar ` + "`" + `region` + "`" + `,` + "`" + `name` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the bucket`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The name of the region`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The FQDN of the bucket without the bucket name (e.g. nyc3.digitaloceanspaces.com) ## Import Buckets can be imported using the ` + "`" + `region` + "`" + ` and ` + "`" + `name` + "`" + ` attributes (delimited by a comma): ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_spaces_bucket.foobar ` + "`" + `region` + "`" + `,` + "`" + `name` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_bucket_object",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `The region where the bucket resides (Defaults to ` + "`" + `nyc3` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to put the file in.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The name of the object once it is in the bucket.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional, conflicts with ` + "`" + `content` + "`" + ` and ` + "`" + `content_base64` + "`" + `) The path to a file that will be read and uploaded as raw bytes for the object content.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional, conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `content_base64` + "`" + `) Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `(Optional, conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `content` + "`" + `) Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the ` + "`" + `gzipbase64` + "`" + ` function with small text strings. For larger objects, use ` + "`" + `source` + "`" + ` to stream the content from a disk file.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The canned ACL to apply. DigitalOcean supports "private" and "public-read". (Defaults to "private".)`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) Specifies presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `(Optional) The language the content is in e.g. en-US or en-GB.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.`,
				},
				resource.Attribute{
					Name:        "website_redirect",
					Description: `(Optional) Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Optional) Used to trigger updates. The only meaningful value is ` + "`" + `${filemd5("path/to/file")}` + "`" + ` (Terraform 0.11.12 or later) or ` + "`" + `${md5(file("path/to/file"))}` + "`" + ` (Terraform 0.11.11 or earlier).`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A mapping of keys/values to provision metadata (will be automatically prefixed by ` + "`" + `x-amz-meta-` + "`" + `, note that only lowercase label are currently supported by the AWS Go API).`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) Allow the object to be deleted by removing any legal hold on any object version. Default is ` + "`" + `false` + "`" + `. This value should be set to ` + "`" + `true` + "`" + ` only if the bucket has S3 object lock enabled. If no content is provided through ` + "`" + `source` + "`" + `, ` + "`" + `content` + "`" + ` or ` + "`" + `content_base64` + "`" + `, then the object will be empty. ->`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `the ETag generated for the object (an MD5 sum of the object content). The hash is an MD5 digest of the object data. For objects created by either the Multipart Upload or Part Copy operation, the hash is not an MD5 digest. More information on possible values can be found on [Common Response Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonResponseHeaders.html).`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object, if bucket versioning is enabled. ## Import Importing this resource is not supported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `the ETag generated for the object (an MD5 sum of the object content). The hash is an MD5 digest of the object data. For objects created by either the Multipart Upload or Part Copy operation, the hash is not an MD5 digest. More information on possible values can be found on [Common Response Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonResponseHeaders.html).`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object, if bucket versioning is enabled. ## Import Importing this resource is not supported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_bucket_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region where the bucket resides.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to which to apply the policy.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The text of the policy. ## Attributes Reference No additional attributes are exported. ## Import Bucket policies can be imported using the ` + "`" + `region` + "`" + ` and ` + "`" + `bucket` + "`" + ` attributes (delimited by a comma): ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_spaces_bucket_policy.foobar ` + "`" + `region` + "`" + `,` + "`" + `bucket` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_ssh_key",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SSH key for identification`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The public key. If this is a file, it can be read using the file interpolation function ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key ## Import SSH Keys can be imported using the ` + "`" + `ssh key id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_ssh_key.mykey 263654 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key ## Import SSH Keys can be imported using the ` + "`" + `ssh key id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_ssh_key.mykey 263654 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_tag",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tag ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the tag`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag`,
				},
				resource.Attribute{
					Name:        "total_resource_count",
					Description: `A count of the total number of resources that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "droplets_count",
					Description: `A count of the Droplets the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "images_count",
					Description: `A count of the images that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volumes_count",
					Description: `A count of the volumes that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volume_snapshots_count",
					Description: `A count of the volume snapshots that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "databases_count",
					Description: `A count of the database clusters that the tag is applied to. ## Import Tags can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_tag.mytag tagname ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the tag`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag`,
				},
				resource.Attribute{
					Name:        "total_resource_count",
					Description: `A count of the total number of resources that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "droplets_count",
					Description: `A count of the Droplets the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "images_count",
					Description: `A count of the images that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volumes_count",
					Description: `A count of the volumes that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "volume_snapshots_count",
					Description: `A count of the volume snapshots that the tag is applied to.`,
				},
				resource.Attribute{
					Name:        "databases_count",
					Description: `A count of the database clusters that the tag is applied to. ## Import Tags can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_tag.mytag tagname ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_uptime_alert",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "check_id",
					Description: `(Required) A unique identifier for a check`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly display name.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.`,
				},
				resource.Attribute{
					Name:        "comparison",
					Description: `The comparison operator used against the alert's threshold. Must be one of ` + "`" + `greater_than` + "`" + ` or ` + "`" + `less_than` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Period of time the threshold must be exceeded to trigger the alert. Must be one of ` + "`" + `2m` + "`" + `, ` + "`" + `3m` + "`" + `, ` + "`" + `5m` + "`" + `, ` + "`" + `10m` + "`" + `, ` + "`" + `15m` + "`" + `, ` + "`" + `30m` + "`" + ` or ` + "`" + `1h` + "`" + `. ` + "`" + `notifications` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `List of email addresses to sent notifications to.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the alert. ## Import Uptime checks can be imported using the uptime alert's ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import digitalocean_uptime_alert.target 5a4981aa-9653-4bd1-bef5-d6bff52042e4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the alert. ## Import Uptime checks can be imported using the uptime alert's ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import digitalocean_uptime_alert.target 5a4981aa-9653-4bd1-bef5-d6bff52042e4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_uptime_check",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly display name for the check.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The endpoint to perform healthchecks on.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of health check to perform: 'ping' 'http' 'https'.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `An array containing the selected regions to perform healthchecks from: "us_east", "us_west", "eu_west", "se_asia"`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `A boolean value indicating whether the check is enabled/disabled. ## Attributes Reference The following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the check. ## Import Uptime checks can be imported using the uptime check's ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import digitalocean_uptime_check.target 5a4981aa-9653-4bd1-bef5-d6bff52042e4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the check. ## Import Uptime checks can be imported using the uptime check's ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import digitalocean_uptime_check.target 5a4981aa-9653-4bd1-bef5-d6bff52042e4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_volume",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region that the block storage volume will be created in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the block storage volume in GiB. If updated, can only be expanded.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A free-form text field up to a limit of 1024 bytes to describe a block storage volume.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limitied on creation to that of the referenced snapshot`,
				},
				resource.Attribute{
					Name:        "initial_filesystem_type",
					Description: `(Optional) Initial filesystem type (` + "`" + `xfs` + "`" + ` or ` + "`" + `ext4` + "`" + `) for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "initial_filesystem_label",
					Description: `(Optional) Initial filesystem label for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of the tags to be applied to this Volume. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the volume.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the volume.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the volume.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of applied tags to the volume.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region that the volume is created in.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `A list of associated droplet ids.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the existing volume snapshot from which this volume was created from.`,
				},
				resource.Attribute{
					Name:        "filesystem_type",
					Description: `Filesystem type (` + "`" + `xfs` + "`" + ` or ` + "`" + `ext4` + "`" + `) for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_label",
					Description: `Filesystem label for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "initial_filesystem_type",
					Description: `Filesystem type (` + "`" + `xfs` + "`" + ` or ` + "`" + `ext4` + "`" + `) for the block storage volume when it was first created.`,
				},
				resource.Attribute{
					Name:        "initial_filesystem_label",
					Description: `Filesystem label for the block storage volume when it was first created. ## Import Volumes can be imported using the ` + "`" + `volume id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_volume.volume 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the volume.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the volume.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the volume.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of applied tags to the volume.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region that the volume is created in.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `A list of associated droplet ids.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the existing volume snapshot from which this volume was created from.`,
				},
				resource.Attribute{
					Name:        "filesystem_type",
					Description: `Filesystem type (` + "`" + `xfs` + "`" + ` or ` + "`" + `ext4` + "`" + `) for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_label",
					Description: `Filesystem label for the block storage volume.`,
				},
				resource.Attribute{
					Name:        "initial_filesystem_type",
					Description: `Filesystem type (` + "`" + `xfs` + "`" + ` or ` + "`" + `ext4` + "`" + `) for the block storage volume when it was first created.`,
				},
				resource.Attribute{
					Name:        "initial_filesystem_label",
					Description: `Filesystem label for the block storage volume when it was first created. ## Import Volumes can be imported using the ` + "`" + `volume id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_volume.volume 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_volume_attachment",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "droplet_id",
					Description: `(Required) ID of the Droplet to attach the volume to.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) ID of the Volume to be attached to the Droplet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the volume attachment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the volume attachment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_volume_snapshot",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the volume snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) The ID of the volume from which the volume snapshot originated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of the tags to be applied to this volume snapshot. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the volume snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a volume to be created based on this volume snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the volume snapshot in gigabytes. ## Import Volume Snapshots can be imported using the ` + "`" + `snapshot id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_volume_snapshot.snapshot 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the volume snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a volume to be created based on this volume snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the volume snapshot in gigabytes. ## Import Volume Snapshots can be imported using the ` + "`" + `snapshot id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_volume_snapshot.snapshot 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_vpc",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the VPC. Must be unique and contain alphanumeric characters, dashes, and periods only.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The DigitalOcean region slug for the VPC's location.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A free-form text field up to a limit of 255 characters to describe the VPC.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Optional) The range of IP addresses for the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be larger than ` + "`" + `/16` + "`" + ` or smaller than ` + "`" + `/24` + "`" + `. ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the VPC.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name (URN) for the VPC.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `A boolean indicating whether or not the VPC is the default one for the region.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time of when the VPC was created. ## Import A VPC can be imported using its ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_vpc.example 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the VPC.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name (URN) for the VPC.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `A boolean indicating whether or not the VPC is the default one for the region.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time of when the VPC was created. ## Import A VPC can be imported using its ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import digitalocean_vpc.example 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"digitalocean_app":                                   0,
		"digitalocean_cdn":                                   1,
		"digitalocean_certificate":                           2,
		"digitalocean_container_registry":                    3,
		"digitalocean_container_registry_docker_credentials": 4,
		"digitalocean_custom_image":                          5,
		"digitalocean_database_cluster":                      6,
		"digitalocean_database_connection_pool":              7,
		"digitalocean_database_db":                           8,
		"digitalocean_database_firewall":                     9,
		"digitalocean_database_replica":                      10,
		"digitalocean_database_user":                         11,
		"digitalocean_domain":                                12,
		"digitalocean_droplet":                               13,
		"digitalocean_droplet_snapshot":                      14,
		"digitalocean_firewall":                              15,
		"digitalocean_floating_ip":                           16,
		"digitalocean_floating_ip_assignment":                17,
		"digitalocean_kubernetes_cluster":                    18,
		"digitalocean_kubernetes_node_pool":                  19,
		"digitalocean_loadbalancer":                          20,
		"digitalocean_monitor_alert":                         21,
		"digitalocean_project":                               22,
		"digitalocean_project_resources":                     23,
		"digitalocean_record":                                24,
		"digitalocean_reserved_ip":                           25,
		"digitalocean_reserved_ip_assignment":                26,
		"digitalocean_spaces_bucket":                         27,
		"digitalocean_spaces_bucket_object":                  28,
		"digitalocean_spaces_bucket_policy":                  29,
		"digitalocean_ssh_key":                               30,
		"digitalocean_tag":                                   31,
		"digitalocean_uptime_alert":                          32,
		"digitalocean_uptime_check":                          33,
		"digitalocean_volume":                                34,
		"digitalocean_volume_attachment":                     35,
		"digitalocean_volume_snapshot":                       36,
		"digitalocean_vpc":                                   37,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
