package digitalocean

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_account",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_app",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) The ID of the app to retrieve information about. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The date and time of when the app was created.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `A DigitalOcean App spec describing the app. A spec can contain multiple components. A ` + "`" + `service` + "`" + ` can contain:`,
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
					Description: `The instance size to use for this component.`,
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
					Description: `A Git repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo_clone_url` + "`" + ` - The clone URL of the repo. - ` + "`" + `branch` + "`" + ` - The name of the branch to use.`,
				},
				resource.Attribute{
					Name:        "github",
					Description: `A GitHub repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + ` or ` + "`" + `gitlab` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "gitlab",
					Description: `A Gitlab repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
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
					Name:        "route",
					Description: `An HTTP paths that should be routed to this component. - ` + "`" + `path` + "`" + ` - Paths must start with ` + "`" + `/` + "`" + ` and must be unique within the app. - ` + "`" + `preserve_path_prefix` + "`" + ` - An optional flag to preserve the path that is forwarded to the backend service.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `A health check to determine the availability of this component. - ` + "`" + `http_path` + "`" + ` - The route path used for the HTTP health check ping. - ` + "`" + `initial_delay_seconds` + "`" + ` - The number of seconds to wait before beginning health checks. - ` + "`" + `period_seconds` + "`" + ` - The number of seconds to wait between health checks. - ` + "`" + `timeout_seconds` + "`" + ` - The number of seconds after which the check times out. - ` + "`" + `success_threshold` + "`" + ` - The number of successful health checks before considered healthy. - ` + "`" + `failure_threshold` + "`" + ` - The number of failed health checks before considered unhealthy. A ` + "`" + `static_site` + "`" + ` can contain:`,
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
					Description: `A Git repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo_clone_url` + "`" + ` - The clone URL of the repo. - ` + "`" + `branch` + "`" + ` - The name of the branch to use.`,
				},
				resource.Attribute{
					Name:        "github",
					Description: `A GitHub repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + ` or ` + "`" + `gitlab` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "gitlab",
					Description: `A Gitlab repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Describes an environment variable made available to an app competent. - ` + "`" + `key` + "`" + ` - The name of the environment variable. - ` + "`" + `value` + "`" + ` - The value of the environment variable. - ` + "`" + `scope` + "`" + ` - The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default). - ` + "`" + `type` + "`" + ` - The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `An HTTP paths that should be routed to this component. - ` + "`" + `path` + "`" + ` - Paths must start with ` + "`" + `/` + "`" + ` and must be unique within the app. - ` + "`" + `preserve_path_prefix` + "`" + ` - An optional flag to preserve the path that is forwarded to the backend service. A ` + "`" + `worker` + "`" + ` can contain:`,
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
					Description: `The instance size to use for this component.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The amount of instances that this component should be scaled to.`,
				},
				resource.Attribute{
					Name:        "git",
					Description: `A Git repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo_clone_url` + "`" + ` - The clone URL of the repo. - ` + "`" + `branch` + "`" + ` - The name of the branch to use.`,
				},
				resource.Attribute{
					Name:        "github",
					Description: `A GitHub repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + ` or ` + "`" + `gitlab` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "gitlab",
					Description: `A Gitlab repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `An image to use as the component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `registry_type` + "`" + ` - The registry type. One of ` + "`" + `DOCR` + "`" + ` (DigitalOcean container registry) or ` + "`" + `DOCKER_HUB` + "`" + `. - ` + "`" + `registry` + "`" + ` - The registry name. Must be left empty for the ` + "`" + `DOCR` + "`" + ` registry type. Required for the ` + "`" + `DOCKER_HUB` + "`" + ` registry type. - ` + "`" + `repository` + "`" + ` - The repository name. - ` + "`" + `tag` + "`" + ` - The repository tag. Defaults to ` + "`" + `latest` + "`" + ` if not provided. - ` + "`" + `deploy_on_push` + "`" + ` - Configures automatically deploying images pushed to DOCR. - ` + "`" + `enabled` + "`" + ` - Whether to automatically deploy images pushed to DOCR.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Describes an environment variable made available to an app competent. - ` + "`" + `key` + "`" + ` - The name of the environment variable. - ` + "`" + `value` + "`" + ` - The value of the environment variable. - ` + "`" + `scope` + "`" + ` - The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default). - ` + "`" + `type` + "`" + ` - The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `. A ` + "`" + `job` + "`" + ` can contain:`,
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
					Description: `The instance size to use for this component.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The amount of instances that this component should be scaled to.`,
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
					Name:        "image",
					Description: `An image to use as the component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `registry_type` + "`" + ` - The registry type. One of ` + "`" + `DOCR` + "`" + ` (DigitalOcean container registry) or ` + "`" + `DOCKER_HUB` + "`" + `. - ` + "`" + `registry` + "`" + ` - The registry name. Must be left empty for the ` + "`" + `DOCR` + "`" + ` registry type. Required for the ` + "`" + `DOCKER_HUB` + "`" + ` registry type. - ` + "`" + `repository` + "`" + ` - The repository name. - ` + "`" + `tag` + "`" + ` - The repository tag. Defaults to ` + "`" + `latest` + "`" + ` if not provided. - ` + "`" + `deploy_on_push` + "`" + ` - Configures automatically deploying images pushed to DOCR. - ` + "`" + `enabled` + "`" + ` - Whether to automatically deploy images pushed to DOCR.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Describes an environment variable made available to an app competent. - ` + "`" + `key` + "`" + ` - The name of the environment variable. - ` + "`" + `value` + "`" + ` - The value of the environment variable. - ` + "`" + `scope` + "`" + ` - The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default). - ` + "`" + `type` + "`" + ` - The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `. A ` + "`" + `function` + "`" + ` component can contain:`,
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
					Name:        "route",
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
					Description: `The name of the MySQL or PostgreSQL user to configure.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The date and time of when the app was created.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `A DigitalOcean App spec describing the app. A spec can contain multiple components. A ` + "`" + `service` + "`" + ` can contain:`,
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
					Description: `The instance size to use for this component.`,
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
					Description: `A Git repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo_clone_url` + "`" + ` - The clone URL of the repo. - ` + "`" + `branch` + "`" + ` - The name of the branch to use.`,
				},
				resource.Attribute{
					Name:        "github",
					Description: `A GitHub repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + ` or ` + "`" + `gitlab` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "gitlab",
					Description: `A Gitlab repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
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
					Name:        "route",
					Description: `An HTTP paths that should be routed to this component. - ` + "`" + `path` + "`" + ` - Paths must start with ` + "`" + `/` + "`" + ` and must be unique within the app. - ` + "`" + `preserve_path_prefix` + "`" + ` - An optional flag to preserve the path that is forwarded to the backend service.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `A health check to determine the availability of this component. - ` + "`" + `http_path` + "`" + ` - The route path used for the HTTP health check ping. - ` + "`" + `initial_delay_seconds` + "`" + ` - The number of seconds to wait before beginning health checks. - ` + "`" + `period_seconds` + "`" + ` - The number of seconds to wait between health checks. - ` + "`" + `timeout_seconds` + "`" + ` - The number of seconds after which the check times out. - ` + "`" + `success_threshold` + "`" + ` - The number of successful health checks before considered healthy. - ` + "`" + `failure_threshold` + "`" + ` - The number of failed health checks before considered unhealthy. A ` + "`" + `static_site` + "`" + ` can contain:`,
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
					Description: `A Git repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo_clone_url` + "`" + ` - The clone URL of the repo. - ` + "`" + `branch` + "`" + ` - The name of the branch to use.`,
				},
				resource.Attribute{
					Name:        "github",
					Description: `A GitHub repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + ` or ` + "`" + `gitlab` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "gitlab",
					Description: `A Gitlab repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Describes an environment variable made available to an app competent. - ` + "`" + `key` + "`" + ` - The name of the environment variable. - ` + "`" + `value` + "`" + ` - The value of the environment variable. - ` + "`" + `scope` + "`" + ` - The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default). - ` + "`" + `type` + "`" + ` - The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `An HTTP paths that should be routed to this component. - ` + "`" + `path` + "`" + ` - Paths must start with ` + "`" + `/` + "`" + ` and must be unique within the app. - ` + "`" + `preserve_path_prefix` + "`" + ` - An optional flag to preserve the path that is forwarded to the backend service. A ` + "`" + `worker` + "`" + ` can contain:`,
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
					Description: `The instance size to use for this component.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The amount of instances that this component should be scaled to.`,
				},
				resource.Attribute{
					Name:        "git",
					Description: `A Git repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo_clone_url` + "`" + ` - The clone URL of the repo. - ` + "`" + `branch` + "`" + ` - The name of the branch to use.`,
				},
				resource.Attribute{
					Name:        "github",
					Description: `A GitHub repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + ` or ` + "`" + `gitlab` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "gitlab",
					Description: `A Gitlab repo to use as component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `repo` + "`" + ` - The name of the repo in the format ` + "`" + `owner/repo` + "`" + `. - ` + "`" + `branch` + "`" + ` - The name of the branch to use. - ` + "`" + `deploy_on_push` + "`" + ` - Whether to automatically deploy new commits made to the repo.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `An image to use as the component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `registry_type` + "`" + ` - The registry type. One of ` + "`" + `DOCR` + "`" + ` (DigitalOcean container registry) or ` + "`" + `DOCKER_HUB` + "`" + `. - ` + "`" + `registry` + "`" + ` - The registry name. Must be left empty for the ` + "`" + `DOCR` + "`" + ` registry type. Required for the ` + "`" + `DOCKER_HUB` + "`" + ` registry type. - ` + "`" + `repository` + "`" + ` - The repository name. - ` + "`" + `tag` + "`" + ` - The repository tag. Defaults to ` + "`" + `latest` + "`" + ` if not provided. - ` + "`" + `deploy_on_push` + "`" + ` - Configures automatically deploying images pushed to DOCR. - ` + "`" + `enabled` + "`" + ` - Whether to automatically deploy images pushed to DOCR.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Describes an environment variable made available to an app competent. - ` + "`" + `key` + "`" + ` - The name of the environment variable. - ` + "`" + `value` + "`" + ` - The value of the environment variable. - ` + "`" + `scope` + "`" + ` - The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default). - ` + "`" + `type` + "`" + ` - The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `. A ` + "`" + `job` + "`" + ` can contain:`,
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
					Description: `The instance size to use for this component.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The amount of instances that this component should be scaled to.`,
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
					Name:        "image",
					Description: `An image to use as the component's source. Only one of ` + "`" + `git` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `gitlab` + "`" + `, or ` + "`" + `image` + "`" + ` may be set. - ` + "`" + `registry_type` + "`" + ` - The registry type. One of ` + "`" + `DOCR` + "`" + ` (DigitalOcean container registry) or ` + "`" + `DOCKER_HUB` + "`" + `. - ` + "`" + `registry` + "`" + ` - The registry name. Must be left empty for the ` + "`" + `DOCR` + "`" + ` registry type. Required for the ` + "`" + `DOCKER_HUB` + "`" + ` registry type. - ` + "`" + `repository` + "`" + ` - The repository name. - ` + "`" + `tag` + "`" + ` - The repository tag. Defaults to ` + "`" + `latest` + "`" + ` if not provided. - ` + "`" + `deploy_on_push` + "`" + ` - Configures automatically deploying images pushed to DOCR. - ` + "`" + `enabled` + "`" + ` - Whether to automatically deploy images pushed to DOCR.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Describes an environment variable made available to an app competent. - ` + "`" + `key` + "`" + ` - The name of the environment variable. - ` + "`" + `value` + "`" + ` - The value of the environment variable. - ` + "`" + `scope` + "`" + ` - The visibility scope of the environment variable. One of ` + "`" + `RUN_TIME` + "`" + `, ` + "`" + `BUILD_TIME` + "`" + `, or ` + "`" + `RUN_AND_BUILD_TIME` + "`" + ` (default). - ` + "`" + `type` + "`" + ` - The type of the environment variable, ` + "`" + `GENERAL` + "`" + ` or ` + "`" + `SECRET` + "`" + `. A ` + "`" + `function` + "`" + ` component can contain:`,
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
					Name:        "route",
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
					Description: `The name of the MySQL or PostgreSQL user to configure.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_certificate",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of certificate. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_container_registry",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the container registry. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The date and time when the registry was created`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The date and time when the registry was created`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_ca",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the source database cluster. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The CA certificate used to secure database connections decoded to a string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate",
					Description: `The CA certificate used to secure database connections decoded to a string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_cluster",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database cluster. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "engine",
					Description: `Database engine used by the cluster (ex. ` + "`" + `pg` + "`" + ` for PostreSQL).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Engine version used by the cluster (ex. ` + "`" + `11` + "`" + ` for PostgreSQL 11).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Database droplet size associated with the cluster (ex. ` + "`" + `db-s-1vcpu-1gb` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `DigitalOcean region where the cluster will reside.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of nodes that will be included in the cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `Defines when the automatic maintenance should be performed for the database cluster.`,
				},
				resource.Attribute{
					Name:        "private_network_uuid",
					Description: `The ID of the VPC where the database cluster is located.`,
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
					Description: `Password for the cluster's default user.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project that the database cluster is assigned to. ` + "`" + `maintenance_window` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `The day of the week on which to apply maintenance updates.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `The hour in UTC at which maintenance updates will be applied in 24 hour format.`,
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
					Name:        "engine",
					Description: `Database engine used by the cluster (ex. ` + "`" + `pg` + "`" + ` for PostreSQL).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Engine version used by the cluster (ex. ` + "`" + `11` + "`" + ` for PostgreSQL 11).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Database droplet size associated with the cluster (ex. ` + "`" + `db-s-1vcpu-1gb` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `DigitalOcean region where the cluster will reside.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of nodes that will be included in the cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `Defines when the automatic maintenance should be performed for the database cluster.`,
				},
				resource.Attribute{
					Name:        "private_network_uuid",
					Description: `The ID of the VPC where the database cluster is located.`,
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
					Description: `Password for the cluster's default user.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project that the database cluster is assigned to. ` + "`" + `maintenance_window` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `The day of the week on which to apply maintenance updates.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `The hour in UTC at which maintenance updates will be applied in 24 hour format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_replica",
			Category:         "Data Sources",
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
					Description: `(Required) The name for the database replica. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database replica.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the database replica.`,
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
					Name:        "tags",
					Description: `A list of tag names to be applied to the database replica.`,
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
					Description: `Password for the replica's default user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the database replica.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the database replica.`,
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
					Name:        "tags",
					Description: `A list of tag names to be applied to the database replica.`,
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
					Description: `Password for the replica's default user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_database_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the database cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database user. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `The role of the database user. The value will be either ` + "`" + `primary` + "`" + ` or ` + "`" + `normal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the database user. This will not be set for MongoDB users.`,
				},
				resource.Attribute{
					Name:        "mysql_auth_plugin",
					Description: `The authentication method of the MySQL user. The value will be ` + "`" + `mysql_native_password` + "`" + ` or ` + "`" + `caching_sha2_password` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role",
					Description: `The role of the database user. The value will be either ` + "`" + `primary` + "`" + ` or ` + "`" + `normal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the database user. This will not be set for MongoDB users.`,
				},
				resource.Attribute{
					Name:        "mysql_auth_plugin",
					Description: `The authentication method of the MySQL user. The value will be ` + "`" + `mysql_native_password` + "`" + ` or ` + "`" + `caching_sha2_password` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_domain",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the domain. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the domain`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the domain`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_domains",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the domains by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `urn` + "`" + `, and ` + "`" + `ttl` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves domains where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) One of ` + "`" + `exact` + "`" + ` (default), ` + "`" + `re` + "`" + `, or ` + "`" + `substring` + "`" + `. For string-typed fields, specify ` + "`" + `re` + "`" + ` to match by using the ` + "`" + `values` + "`" + ` as regular expressions, or specify ` + "`" + `substring` + "`" + ` to match by treating the ` + "`" + `values` + "`" + ` as substrings to find within the string field.`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to require that a field match all of the ` + "`" + `values` + "`" + ` instead of just one or more of them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure that all of the ` + "`" + `values` + "`" + ` are present in the list or set. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the domains by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `urn` + "`" + `, and ` + "`" + `ttl` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `A list of domains satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each domain has the following attributes: - ` + "`" + `name` + "`" + ` - (Required) The name of the domain. - ` + "`" + `ttl` + "`" + `- The TTL of the domain. - ` + "`" + `urn` + "`" + ` - The uniform resource name of the domain`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domains",
					Description: `A list of domains satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each domain has the following attributes: - ` + "`" + `name` + "`" + ` - (Required) The name of the domain. - ` + "`" + `ttl` + "`" + `- The TTL of the domain. - ` + "`" + `urn` + "`" + ` - The uniform resource name of the domain`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_droplet",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the Droplet`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Droplet.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag applied to the Droplet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the Droplet`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the Droplet is running in.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The Droplet image ID or slug.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The unique slug that indentifies the type of Droplet.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The size of the Droplets disk in GB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of the Droplets virtual CPUs.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of the Droplets memory in MB.`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `Droplet hourly price.`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `Droplet monthly price.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Droplet.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the Droplet is locked.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The Droplets public IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv6_address_private",
					Description: `The Droplets private IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The Droplets public IPv4 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address_private",
					Description: `The Droplets private IPv4 address`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `Whether backups are enabled.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Whether IPv6 is enabled.`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `Whether private networks are enabled.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `Whether monitoring agent is installed.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `List of the IDs of each volumes attached to the Droplet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the Droplet.`,
				},
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `The ID of the VPC where the Droplet is located.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the Droplet`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the Droplet is running in.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The Droplet image ID or slug.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The unique slug that indentifies the type of Droplet.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The size of the Droplets disk in GB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of the Droplets virtual CPUs.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of the Droplets memory in MB.`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `Droplet hourly price.`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `Droplet monthly price.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Droplet.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the Droplet is locked.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The Droplets public IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv6_address_private",
					Description: `The Droplets private IPv6 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `The Droplets public IPv4 address`,
				},
				resource.Attribute{
					Name:        "ipv4_address_private",
					Description: `The Droplets private IPv4 address`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `Whether backups are enabled.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Whether IPv6 is enabled.`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `Whether private networks are enabled.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `Whether monitoring agent is installed.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `List of the IDs of each volumes attached to the Droplet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the Droplet.`,
				},
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `The ID of the VPC where the Droplet is located.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_droplet_snapshot",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Droplet snapshot.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the Droplet snapshot list returned by DigitalOcean. This allows more advanced filtering not supported from the DigitalOcean API. This filtering is done locally on what DigitalOcean returns.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) A "slug" representing a DigitalOcean region (e.g. ` + "`" + `nyc1` + "`" + `). If set, only Droplet snapshots available in the region will be returned.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent Droplet snapshot. ~>`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the Droplet snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a Droplet to be created based on this Droplet snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the Droplet snapshot is available.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `The ID of the Droplet from which the Droplet snapshot originated.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the Droplet snapshot in gigabytes.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time the Droplet snapshot was created.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `The minimum size in gigabytes required for a Droplet to be created based on this Droplet snapshot.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of DigitalOcean region "slugs" indicating where the Droplet snapshot is available.`,
				},
				resource.Attribute{
					Name:        "droplet_id",
					Description: `The ID of the Droplet from which the Droplet snapshot originated.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the Droplet snapshot in gigabytes.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_droplets",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the Droplets by this key. This may be one of ` + "`" + `backups` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `id` + "`" + `, ` + "`" + `image` + "`" + `, ` + "`" + `ipv4_address` + "`" + `, ` + "`" + `ipv4_address_private` + "`" + `, ` + "`" + `ipv6` + "`" + `, ` + "`" + `ipv6_address` + "`" + `, ` + "`" + `ipv6_address_private` + "`" + `, ` + "`" + `locked` + "`" + `, ` + "`" + `memory` + "`" + `, ` + "`" + `monitoring` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `price_hourly` + "`" + `, ` + "`" + `price_monthly` + "`" + `, ` + "`" + `private_networking` + "`" + `, ` + "`" + `region` + "`" + `, ` + "`" + `size` + "`" + `, ` + "`" + `status` + "`" + `, ` + "`" + `tags` + "`" + `, ` + "`" + `urn` + "`" + `, ` + "`" + `vcpus` + "`" + `, ` + "`" + `volume_ids` + "`" + `, or ` + "`" + `vpc_uuid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves Droplets where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) One of ` + "`" + `exact` + "`" + ` (default), ` + "`" + `re` + "`" + `, or ` + "`" + `substring` + "`" + `. For string-typed fields, specify ` + "`" + `re` + "`" + ` to match by using the ` + "`" + `values` + "`" + ` as regular expressions, or specify ` + "`" + `substring` + "`" + ` to match by treating the ` + "`" + `values` + "`" + ` as substrings to find within the string field.`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to require that a field match all of the ` + "`" + `values` + "`" + ` instead of just one or more of them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure that all of the ` + "`" + `values` + "`" + ` are present in the list or set. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the Droplets by this key. This may be one of ` + "`" + `backups` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `id` + "`" + `, ` + "`" + `image` + "`" + `, ` + "`" + `ipv4_address` + "`" + `, ` + "`" + `ipv4_address_private` + "`" + `, ` + "`" + `ipv6` + "`" + `, ` + "`" + `ipv6_address` + "`" + `, ` + "`" + `ipv6_address_private` + "`" + `, ` + "`" + `locked` + "`" + `, ` + "`" + `memory` + "`" + `, ` + "`" + `monitoring` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `price_hourly` + "`" + `, ` + "`" + `price_monthly` + "`" + `, ` + "`" + `private_networking` + "`" + `, ` + "`" + `region` + "`" + `, ` + "`" + `size` + "`" + `, ` + "`" + `status` + "`" + `, ` + "`" + `urn` + "`" + `, ` + "`" + `vcpus` + "`" + `, or ` + "`" + `vpc_uuid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "droplets",
					Description: `A list of Droplets satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each Droplet has the following attributes: - ` + "`" + `id` + "`" + ` - The ID of the Droplet. - ` + "`" + `urn` + "`" + ` - The uniform resource name of the Droplet - ` + "`" + `region` + "`" + ` - The region the Droplet is running in. - ` + "`" + `image` + "`" + ` - The Droplet image ID or slug. - ` + "`" + `size` + "`" + ` - The unique slug that identifies the type of Droplet. - ` + "`" + `disk` + "`" + ` - The size of the Droplet's disk in GB. - ` + "`" + `vcpus` + "`" + ` - The number of the Droplet's virtual CPUs. - ` + "`" + `memory` + "`" + ` - The amount of the Droplet's memory in MB. - ` + "`" + `price_hourly` + "`" + ` - Droplet hourly price. - ` + "`" + `price_monthly` + "`" + ` - Droplet monthly price. - ` + "`" + `status` + "`" + ` - The status of the Droplet. - ` + "`" + `locked` + "`" + ` - Whether the Droplet is locked. - ` + "`" + `ipv6_address` + "`" + ` - The Droplet's public IPv6 address - ` + "`" + `ipv6_address_private` + "`" + ` - The Droplet's private IPv6 address - ` + "`" + `ipv4_address` + "`" + ` - The Droplet's public IPv4 address - ` + "`" + `ipv4_address_private` + "`" + ` - The Droplet's private IPv4 address - ` + "`" + `backups` + "`" + ` - Whether backups are enabled. - ` + "`" + `ipv6` + "`" + ` - Whether IPv6 is enabled. - ` + "`" + `private_networking` + "`" + ` - Whether private networks are enabled. - ` + "`" + `monitoring` + "`" + ` - Whether monitoring agent is installed. - ` + "`" + `volume_ids` + "`" + ` - List of the IDs of each volumes attached to the Droplet. - ` + "`" + `tags` + "`" + ` - A list of the tags associated to the Droplet. - ` + "`" + `vpc_uuid` + "`" + ` - The ID of the VPC where the Droplet is located.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "droplets",
					Description: `A list of Droplets satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each Droplet has the following attributes: - ` + "`" + `id` + "`" + ` - The ID of the Droplet. - ` + "`" + `urn` + "`" + ` - The uniform resource name of the Droplet - ` + "`" + `region` + "`" + ` - The region the Droplet is running in. - ` + "`" + `image` + "`" + ` - The Droplet image ID or slug. - ` + "`" + `size` + "`" + ` - The unique slug that identifies the type of Droplet. - ` + "`" + `disk` + "`" + ` - The size of the Droplet's disk in GB. - ` + "`" + `vcpus` + "`" + ` - The number of the Droplet's virtual CPUs. - ` + "`" + `memory` + "`" + ` - The amount of the Droplet's memory in MB. - ` + "`" + `price_hourly` + "`" + ` - Droplet hourly price. - ` + "`" + `price_monthly` + "`" + ` - Droplet monthly price. - ` + "`" + `status` + "`" + ` - The status of the Droplet. - ` + "`" + `locked` + "`" + ` - Whether the Droplet is locked. - ` + "`" + `ipv6_address` + "`" + ` - The Droplet's public IPv6 address - ` + "`" + `ipv6_address_private` + "`" + ` - The Droplet's private IPv6 address - ` + "`" + `ipv4_address` + "`" + ` - The Droplet's public IPv4 address - ` + "`" + `ipv4_address_private` + "`" + ` - The Droplet's private IPv4 address - ` + "`" + `backups` + "`" + ` - Whether backups are enabled. - ` + "`" + `ipv6` + "`" + ` - Whether IPv6 is enabled. - ` + "`" + `private_networking` + "`" + ` - Whether private networks are enabled. - ` + "`" + `monitoring` + "`" + ` - Whether monitoring agent is installed. - ` + "`" + `volume_ids` + "`" + ` - List of the IDs of each volumes attached to the Droplet. - ` + "`" + `tags` + "`" + ` - A list of the tags associated to the Droplet. - ` + "`" + `vpc_uuid` + "`" + ` - The ID of the VPC where the Droplet is located.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_firewall",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "firewall_id",
					Description: `(Required) The ID of the firewall to retrieve information about. ## Attributes Reference The following attributes are exported:`,
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
					Description: `A set of object containing the fields, ` + "`" + `droplet_id` + "`" + `, ` + "`" + `removing` + "`" + `, and ` + "`" + `status` + "`" + `. It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.`,
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
					Name:        "inbound_rules",
					Description: `The inbound access rule block for the Firewall.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound access rule block for the Firewall. ` + "`" + `inbound_rule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The type of traffic to be allowed. This may be one of "tcp", "udp", or "icmp".`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "1-65535" to open all ports for a protocol. Required for when protocol is ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `An array of strings containing the IPv4 addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_droplet_ids",
					Description: `An array containing the IDs of the Droplets from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_tags",
					Description: `A set of names of Tags corresponding to group of Droplets from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_load_balancer_uids",
					Description: `An array containing the IDs of the Load Balancers from which the inbound traffic will be accepted. ` + "`" + `outbound_rule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The type of traffic to be allowed. This may be one of "tcp", "udp", or "icmp".`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "1-65535" to open all ports for a protocol. Required for when protocol is ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `An array of strings containing the IPv4 addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the outbound traffic will be allowed.`,
				},
				resource.Attribute{
					Name:        "destination_droplet_ids",
					Description: `An array containing the IDs of the Droplets to which the outbound traffic will be allowed.`,
				},
				resource.Attribute{
					Name:        "destination_tags",
					Description: `An array containing the names of Tags corresponding to groups of Droplets to which the outbound traffic will be allowed. traffic.`,
				},
				resource.Attribute{
					Name:        "destination_load_balancer_uids",
					Description: `An array containing the IDs of the Load Balancers to which the outbound traffic will be allowed.`,
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
					Description: `A set of object containing the fields, ` + "`" + `droplet_id` + "`" + `, ` + "`" + `removing` + "`" + `, and ` + "`" + `status` + "`" + `. It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.`,
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
					Name:        "inbound_rules",
					Description: `The inbound access rule block for the Firewall.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound access rule block for the Firewall. ` + "`" + `inbound_rule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The type of traffic to be allowed. This may be one of "tcp", "udp", or "icmp".`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "1-65535" to open all ports for a protocol. Required for when protocol is ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `An array of strings containing the IPv4 addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_droplet_ids",
					Description: `An array containing the IDs of the Droplets from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_tags",
					Description: `A set of names of Tags corresponding to group of Droplets from which the inbound traffic will be accepted.`,
				},
				resource.Attribute{
					Name:        "source_load_balancer_uids",
					Description: `An array containing the IDs of the Load Balancers from which the inbound traffic will be accepted. ` + "`" + `outbound_rule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The type of traffic to be allowed. This may be one of "tcp", "udp", or "icmp".`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. "8000-9000"), or "1-65535" to open all ports for a protocol. Required for when protocol is ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `An array of strings containing the IPv4 addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the outbound traffic will be allowed.`,
				},
				resource.Attribute{
					Name:        "destination_droplet_ids",
					Description: `An array containing the IDs of the Droplets to which the outbound traffic will be allowed.`,
				},
				resource.Attribute{
					Name:        "destination_tags",
					Description: `An array containing the names of Tags corresponding to groups of Droplets to which the outbound traffic will be allowed. traffic.`,
				},
				resource.Attribute{
					Name:        "destination_load_balancer_uids",
					Description: `An array containing the IDs of the Load Balancers to which the outbound traffic will be allowed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_floating_ip",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The allocated IP address of the specific floating IP to retrieve. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_image",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the image`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the image.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug of the official image. If ` + "`" + `name` + "`" + ` is specified, you may also specify:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Restrict the search to one of the following categories of images: - ` + "`" + `all` + "`" + ` - All images (whether public or private) - ` + "`" + `applications` + "`" + ` - One-click applications - ` + "`" + `distributions` + "`" + ` - Distributions - ` + "`" + `user` + "`" + ` - (Default) User (private) images ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `The name of the distribution of the OS of the image.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `Is image a public image or not. Public images represent Linux distributions or One-Click Applications, while non-public images represent snapshots and backups and are only available within your account.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The id of the image (legacy parameter).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "distribution",
					Description: `The name of the distribution of the OS of the image.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `Is image a public image or not. Public images represent Linux distributions or One-Click Applications, while non-public images represent snapshots and backups and are only available within your account.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The id of the image (legacy parameter).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_images",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the images by this key. This may be one of ` + "`" + `distribution` + "`" + `, ` + "`" + `error_message` + "`" + `, ` + "`" + `id` + "`" + `, ` + "`" + `image` + "`" + `, ` + "`" + `min_disk_size` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `private` + "`" + `, ` + "`" + `regions` + "`" + `, ` + "`" + `size_gigabytes` + "`" + `, ` + "`" + `slug` + "`" + `, ` + "`" + `status` + "`" + `, ` + "`" + `tags` + "`" + `, or ` + "`" + `type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves images where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) One of ` + "`" + `exact` + "`" + ` (default), ` + "`" + `re` + "`" + `, or ` + "`" + `substring` + "`" + `. For string-typed fields, specify ` + "`" + `re` + "`" + ` to match by using the ` + "`" + `values` + "`" + ` as regular expressions, or specify ` + "`" + `substring` + "`" + ` to match by treating the ` + "`" + `values` + "`" + ` as substrings to find within the string field.`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to require that a field match all of the ` + "`" + `values` + "`" + ` instead of just one or more of them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure that all of the ` + "`" + `values` + "`" + ` are present in the list or set. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the images by this key. This may be one of ` + "`" + `distribution` + "`" + `, ` + "`" + `error_message` + "`" + `, ` + "`" + `id` + "`" + `, ` + "`" + `image` + "`" + `, ` + "`" + `min_disk_size` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `private` + "`" + `, ` + "`" + `size_gigabytes` + "`" + `, ` + "`" + `slug` + "`" + `, ` + "`" + `status` + "`" + `, or ` + "`" + `type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `A set of images satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each image has the following attributes: - ` + "`" + `slug` + "`" + `: Unique text identifier of the image. - ` + "`" + `id` + "`" + `: The ID of the image. - ` + "`" + `name` + "`" + `: The name of the image. - ` + "`" + `type` + "`" + `: Type of the image. - ` + "`" + `distribution` + "`" + ` - The name of the distribution of the OS of the image. - ` + "`" + `min_disk_size` + "`" + `: The minimum 'disk' required for the image. - ` + "`" + `size_gigabytes` + "`" + `: The size of the image in GB. - ` + "`" + `private` + "`" + ` - Is image a public image or not. Public images represent Linux distributions or One-Click Applications, while non-public images represent snapshots and backups and are only available within your account. - ` + "`" + `regions` + "`" + `: A set of the regions that the image is available in. - ` + "`" + `tags` + "`" + `: A set of tags applied to the image - ` + "`" + `created` + "`" + `: When the image was created - ` + "`" + `status` + "`" + `: Current status of the image - ` + "`" + `error_message` + "`" + `: Any applicable error message pertaining to the image - ` + "`" + `image` + "`" + ` - The id of the image (legacy parameter).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `A set of images satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each image has the following attributes: - ` + "`" + `slug` + "`" + `: Unique text identifier of the image. - ` + "`" + `id` + "`" + `: The ID of the image. - ` + "`" + `name` + "`" + `: The name of the image. - ` + "`" + `type` + "`" + `: Type of the image. - ` + "`" + `distribution` + "`" + ` - The name of the distribution of the OS of the image. - ` + "`" + `min_disk_size` + "`" + `: The minimum 'disk' required for the image. - ` + "`" + `size_gigabytes` + "`" + `: The size of the image in GB. - ` + "`" + `private` + "`" + ` - Is image a public image or not. Public images represent Linux distributions or One-Click Applications, while non-public images represent snapshots and backups and are only available within your account. - ` + "`" + `regions` + "`" + `: A set of the regions that the image is available in. - ` + "`" + `tags` + "`" + `: A set of tags applied to the image - ` + "`" + `created` + "`" + `: When the image was created - ` + "`" + `status` + "`" + `: Current status of the image - ` + "`" + `error_message` + "`" + `: Any applicable error message pertaining to the image - ` + "`" + `image` + "`" + ` - The id of the image (legacy parameter).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_kubernetes_cluster",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of Kubernetes cluster. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID that can be used to identify and reference a Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The slug identifier for the region where the Kubernetes cluster is located.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The slug identifier for the version of Kubernetes used for the cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tag names to be applied to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `The ID of the VPC where the Kubernetes cluster is located.`,
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
					Description: `The public IPv4 address of the Kubernetes master node.`,
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
					Name:        "maintenance_policy",
					Description: `The maintenance policy of the Kubernetes cluster. Digital Ocean has a default maintenancen window. - ` + "`" + `day` + "`" + ` - The day for the service window of the Kubernetes cluster. - ` + "`" + `duration` + "`" + ` - The duration of the operation. - ` + "`" + `start_time` + "`" + ` - The start time of the upgrade operation.`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `A list of node pools associated with the cluster. Each node pool exports the following attributes: - ` + "`" + `id` + "`" + ` - The unique ID that can be used to identify and reference the node pool. - ` + "`" + `name` + "`" + ` - The name of the node pool. - ` + "`" + `size` + "`" + ` - The slug identifier for the type of Droplet used as workers in the node pool. - ` + "`" + `node_count` + "`" + ` - The number of Droplet instances in the node pool. - ` + "`" + `actual_node_count` + "`" + ` - The actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled. - ` + "`" + `auto_scale` + "`" + ` - A boolean indicating whether auto-scaling is enabled on the node pool. - ` + "`" + `min_nodes` + "`" + ` - If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to. - ` + "`" + `max_nodes` + "`" + ` - If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to. - ` + "`" + `tags` + "`" + ` - A list of tag names applied to the node pool. - ` + "`" + `labels` + "`" + ` - A map of key/value pairs applied to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/). - ` + "`" + `nodes` + "`" + ` - A list of nodes in the pool. Each node exports the following attributes: + ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. + ` + "`" + `name` + "`" + ` - The auto-generated name for the node. + ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. + ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. + ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated. - ` + "`" + `taint` + "`" + ` - A list of taints applied to all nodes in the pool. Each taint exports the following attributes: + ` + "`" + `key` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. + ` + "`" + `value` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. + ` + "`" + `effect` + "`" + ` - How the node reacts to pods that it won't tolerate. Available effect values are: "NoSchedule", "PreferNoSchedule", "NoExecute".`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name (URN) for the Kubernetes cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID that can be used to identify and reference a Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The slug identifier for the region where the Kubernetes cluster is located.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The slug identifier for the version of Kubernetes used for the cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tag names to be applied to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_uuid",
					Description: `The ID of the VPC where the Kubernetes cluster is located.`,
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
					Description: `The public IPv4 address of the Kubernetes master node.`,
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
					Name:        "maintenance_policy",
					Description: `The maintenance policy of the Kubernetes cluster. Digital Ocean has a default maintenancen window. - ` + "`" + `day` + "`" + ` - The day for the service window of the Kubernetes cluster. - ` + "`" + `duration` + "`" + ` - The duration of the operation. - ` + "`" + `start_time` + "`" + ` - The start time of the upgrade operation.`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `A list of node pools associated with the cluster. Each node pool exports the following attributes: - ` + "`" + `id` + "`" + ` - The unique ID that can be used to identify and reference the node pool. - ` + "`" + `name` + "`" + ` - The name of the node pool. - ` + "`" + `size` + "`" + ` - The slug identifier for the type of Droplet used as workers in the node pool. - ` + "`" + `node_count` + "`" + ` - The number of Droplet instances in the node pool. - ` + "`" + `actual_node_count` + "`" + ` - The actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled. - ` + "`" + `auto_scale` + "`" + ` - A boolean indicating whether auto-scaling is enabled on the node pool. - ` + "`" + `min_nodes` + "`" + ` - If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to. - ` + "`" + `max_nodes` + "`" + ` - If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to. - ` + "`" + `tags` + "`" + ` - A list of tag names applied to the node pool. - ` + "`" + `labels` + "`" + ` - A map of key/value pairs applied to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/). - ` + "`" + `nodes` + "`" + ` - A list of nodes in the pool. Each node exports the following attributes: + ` + "`" + `id` + "`" + ` - A unique ID that can be used to identify and reference the node. + ` + "`" + `name` + "`" + ` - The auto-generated name for the node. + ` + "`" + `status` + "`" + ` - A string indicating the current status of the individual node. + ` + "`" + `created_at` + "`" + ` - The date and time when the node was created. + ` + "`" + `updated_at` + "`" + ` - The date and time when the node was last updated. - ` + "`" + `taint` + "`" + ` - A list of taints applied to all nodes in the pool. Each taint exports the following attributes: + ` + "`" + `key` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. + ` + "`" + `value` + "`" + ` - An arbitrary string. The "key" and "value" fields of the "taint" object form a key-value pair. + ` + "`" + `effect` + "`" + ` - How the node reacts to pods that it won't tolerate. Available effect values are: "NoSchedule", "PreferNoSchedule", "NoExecute".`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name (URN) for the Kubernetes cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_kubernetes_versions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "version_prefix",
					Description: `(Optional) If provided, Terraform will only return versions that match the string prefix. For example, ` + "`" + `1.15.` + "`" + ` will match all 1.15.x series releases. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "valid_versions",
					Description: `A list of available versions.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The most recent version available.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "valid_versions",
					Description: `A list of available versions.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The most recent version available.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_loadbalancer",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of load balancer.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of load balancer.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name for the Load Balancer ## Attributes Reference See the [Load Balancer Resource](/providers/digitalocean/digitalocean/latest/docs/resources/loadbalancer) for details on the returned attributes - they are identical.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_project",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) the ID of the project to retrieve`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) the name of the project to retrieve. The data source will raise an error if more than one project has the provided name or if no project has that name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the project`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `The purpose of the project, (Default: "Web Application")`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The environment of the project's resources. The possible values are: ` + "`" + `Development` + "`" + `, ` + "`" + `Staging` + "`" + `, ` + "`" + `Production` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `A set of uniform resource names (URNs) for the resources associated with the project`,
				},
				resource.Attribute{
					Name:        "owner_uuid",
					Description: `The unique universal identifier of the project owner.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the project owner.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the project was created, (ISO8601)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date and time when the project was last updated, (ISO8601)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the project`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `The purpose of the project, (Default: "Web Application")`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The environment of the project's resources. The possible values are: ` + "`" + `Development` + "`" + `, ` + "`" + `Staging` + "`" + `, ` + "`" + `Production` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `A set of uniform resource names (URNs) for the resources associated with the project`,
				},
				resource.Attribute{
					Name:        "owner_uuid",
					Description: `The unique universal identifier of the project owner.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the project owner.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date and time when the project was created, (ISO8601)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date and time when the project was last updated, (ISO8601)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_projects",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the projects by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `purpose` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `environment` + "`" + `, or ` + "`" + `is_default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves projects where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) One of ` + "`" + `exact` + "`" + ` (default), ` + "`" + `re` + "`" + `, or ` + "`" + `substring` + "`" + `. For string-typed fields, specify ` + "`" + `re` + "`" + ` to match by using the ` + "`" + `values` + "`" + ` as regular expressions, or specify ` + "`" + `substring` + "`" + ` to match by treating the ` + "`" + `values` + "`" + ` as substrings to find within the string field.`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to require that a field match all of the ` + "`" + `values` + "`" + ` instead of just one or more of them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure that all of the ` + "`" + `values` + "`" + ` are present in the list or set. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the projects by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `purpose` + "`" + `, ` + "`" + `description` + "`" + `, or ` + "`" + `environment` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "projects",
					Description: `A set of projects satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each project has the following attributes: - ` + "`" + `id` + "`" + ` - The ID of the project - ` + "`" + `name` + "`" + ` - The name of the project - ` + "`" + `description` + "`" + ` - The description of the project - ` + "`" + `purpose` + "`" + ` - The purpose of the project (Default: "Web Application") - ` + "`" + `environment` + "`" + ` - The environment of the project's resources. The possible values are: ` + "`" + `Development` + "`" + `, ` + "`" + `Staging` + "`" + `, ` + "`" + `Production` + "`" + `. - ` + "`" + `resources` + "`" + ` - A set of uniform resource names (URNs) for the resources associated with the project - ` + "`" + `owner_uuid` + "`" + ` - The unique universal identifier of the project owner - ` + "`" + `owner_id` + "`" + ` - The ID of the project owner - ` + "`" + `created_at` + "`" + ` - The date and time when the project was created, (ISO8601) - ` + "`" + `updated_at` + "`" + ` - The date and time when the project was last updated, (ISO8601)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "projects",
					Description: `A set of projects satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each project has the following attributes: - ` + "`" + `id` + "`" + ` - The ID of the project - ` + "`" + `name` + "`" + ` - The name of the project - ` + "`" + `description` + "`" + ` - The description of the project - ` + "`" + `purpose` + "`" + ` - The purpose of the project (Default: "Web Application") - ` + "`" + `environment` + "`" + ` - The environment of the project's resources. The possible values are: ` + "`" + `Development` + "`" + `, ` + "`" + `Staging` + "`" + `, ` + "`" + `Production` + "`" + `. - ` + "`" + `resources` + "`" + ` - A set of uniform resource names (URNs) for the resources associated with the project - ` + "`" + `owner_uuid` + "`" + ` - The unique universal identifier of the project owner - ` + "`" + `owner_id` + "`" + ` - The ID of the project owner - ` + "`" + `created_at` + "`" + ` - The date and time when the project was created, (ISO8601) - ` + "`" + `updated_at` + "`" + ` - The date and time when the project was last updated, (ISO8601)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_record",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain name of the record. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_records",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain name to search for DNS records`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the DNS records by this key. This may be one of ` + "`" + `domain` + "`" + `, ` + "`" + `flags` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `port` + "`" + `, ` + "`" + `priority` + "`" + `, ` + "`" + `tag` + "`" + `, ` + "`" + `ttl` + "`" + `, ` + "`" + `type` + "`" + `, ` + "`" + `value` + "`" + `, or ` + "`" + `weight` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves DNS records where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) One of ` + "`" + `exact` + "`" + ` (default), ` + "`" + `re` + "`" + `, or ` + "`" + `substring` + "`" + `. For string-typed fields, specify ` + "`" + `re` + "`" + ` to match by using the ` + "`" + `values` + "`" + ` as regular expressions, or specify ` + "`" + `substring` + "`" + ` to match by treating the ` + "`" + `values` + "`" + ` as substrings to find within the string field.`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to require that a field match all of the ` + "`" + `values` + "`" + ` instead of just one or more of them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure that all of the ` + "`" + `values` + "`" + ` are present in the list or set. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the DNS records by this key. This may be one of ` + "`" + `domain` + "`" + `, ` + "`" + `flags` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `port` + "`" + `, ` + "`" + `priority` + "`" + `, ` + "`" + `tag` + "`" + `, ` + "`" + `ttl` + "`" + `, ` + "`" + `type` + "`" + `, ` + "`" + `value` + "`" + `, or ` + "`" + `weight` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_region",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) A human-readable string that is used as a unique identifier for each region. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `A human-readable string that is used as a unique identifier for each region.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The display name of the region.`,
				},
				resource.Attribute{
					Name:        "available",
					Description: `A boolean value that represents whether new Droplets can be created in this region.`,
				},
				resource.Attribute{
					Name:        "sizes",
					Description: `A set of identifying slugs for the Droplet sizes available in this region.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `A set of features available in this region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `A human-readable string that is used as a unique identifier for each region.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The display name of the region.`,
				},
				resource.Attribute{
					Name:        "available",
					Description: `A boolean value that represents whether new Droplets can be created in this region.`,
				},
				resource.Attribute{
					Name:        "sizes",
					Description: `A set of identifying slugs for the Droplet sizes available in this region.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `A set of features available in this region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_regions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the regions by this key. This may be one of ` + "`" + `slug` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `available` + "`" + `, ` + "`" + `features` + "`" + `, or ` + "`" + `sizes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves regions where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) One of ` + "`" + `exact` + "`" + ` (default), ` + "`" + `re` + "`" + `, or ` + "`" + `substring` + "`" + `. For string-typed fields, specify ` + "`" + `re` + "`" + ` to match by using the ` + "`" + `values` + "`" + ` as regular expressions, or specify ` + "`" + `substring` + "`" + ` to match by treating the ` + "`" + `values` + "`" + ` as substrings to find within the string field.`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to require that a field match all of the ` + "`" + `values` + "`" + ` instead of just one or more of them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure that all of the ` + "`" + `values` + "`" + ` are present in the list or set. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the regions by this key. This may be one of ` + "`" + `slug` + "`" + `, ` + "`" + `name` + "`" + `, or ` + "`" + `available` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A set of regions satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each region has the following attributes: - ` + "`" + `slug` + "`" + ` - A human-readable string that is used as a unique identifier for each region. - ` + "`" + `name` + "`" + ` - The display name of the region. - ` + "`" + `available` + "`" + ` - A boolean value that represents whether new Droplets can be created in this region. - ` + "`" + `sizes` + "`" + ` - A set of identifying slugs for the Droplet sizes available in this region. - ` + "`" + `features` + "`" + ` - A set of features available in this region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `A set of regions satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each region has the following attributes: - ` + "`" + `slug` + "`" + ` - A human-readable string that is used as a unique identifier for each region. - ` + "`" + `name` + "`" + ` - The display name of the region. - ` + "`" + `available` + "`" + ` - A boolean value that represents whether new Droplets can be created in this region. - ` + "`" + `sizes` + "`" + ` - A set of identifying slugs for the Droplet sizes available in this region. - ` + "`" + `features` + "`" + ` - A set of features available in this region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_reserved_ip",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The allocated IP address of the specific reserved IP to retrieve. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_sizes",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the sizes by this key. This may be one of ` + "`" + `slug` + "`" + `, ` + "`" + `regions` + "`" + `, ` + "`" + `memory` + "`" + `, ` + "`" + `vcpus` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `transfer` + "`" + `, ` + "`" + `price_monthly` + "`" + `, ` + "`" + `price_hourly` + "`" + `, or ` + "`" + `available` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Only retrieves sizes which keys has value that matches one of the values provided here.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) One of ` + "`" + `exact` + "`" + ` (default), ` + "`" + `re` + "`" + `, or ` + "`" + `substring` + "`" + `. For string-typed fields, specify ` + "`" + `re` + "`" + ` to match by using the ` + "`" + `values` + "`" + ` as regular expressions, or specify ` + "`" + `substring` + "`" + ` to match by treating the ` + "`" + `values` + "`" + ` as substrings to find within the string field.`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to require that a field match all of the ` + "`" + `values` + "`" + ` instead of just one or more of them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure that all of the ` + "`" + `values` + "`" + ` are present in the list or set. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the sizes by this key. This may be one of ` + "`" + `slug` + "`" + `, ` + "`" + `memory` + "`" + `, ` + "`" + `vcpus` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `transfer` + "`" + `, ` + "`" + `price_monthly` + "`" + `, or ` + "`" + `price_hourly` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `A human-readable string that is used to uniquely identify each size.`,
				},
				resource.Attribute{
					Name:        "available",
					Description: `This represents whether new Droplets can be created with this size.`,
				},
				resource.Attribute{
					Name:        "transfer",
					Description: `The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `The monthly cost of Droplets created in this size if they are kept for an entire month. The value is measured in US dollars.`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `The hourly cost of Droplets created in this size as measured hourly. The value is measured in US dollars.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of RAM allocated to Droplets created of this size. The value is measured in megabytes.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of CPUs allocated to Droplets of this size.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The amount of disk space set aside for Droplets of this size. The value is measured in gigabytes.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of region slugs where Droplets can be created in this size.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `A human-readable string that is used to uniquely identify each size.`,
				},
				resource.Attribute{
					Name:        "available",
					Description: `This represents whether new Droplets can be created with this size.`,
				},
				resource.Attribute{
					Name:        "transfer",
					Description: `The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.`,
				},
				resource.Attribute{
					Name:        "price_monthly",
					Description: `The monthly cost of Droplets created in this size if they are kept for an entire month. The value is measured in US dollars.`,
				},
				resource.Attribute{
					Name:        "price_hourly",
					Description: `The hourly cost of Droplets created in this size as measured hourly. The value is measured in US dollars.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of RAM allocated to Droplets created of this size. The value is measured in megabytes.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of CPUs allocated to Droplets of this size.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The amount of disk space set aside for Droplets of this size. The value is measured in gigabytes.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of region slugs where Droplets can be created in this size.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_bucket",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Spaces bucket.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The slug of the region where the bucket is stored. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Spaces bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The slug of the region where the bucket is stored.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the bucket`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The FQDN of the bucket without the bucket name (e.g. nyc3.digitaloceanspaces.com)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Spaces bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The slug of the region where the bucket is stored.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `The uniform resource name of the bucket`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The FQDN of the bucket without the bucket name (e.g. nyc3.digitaloceanspaces.com)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_bucket_object",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to read the object from.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The slug of the region where the bucket is stored.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The full path to the object inside the bucket`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Optional) Specific version ID of the object returned (defaults to latest version) ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in Spaces`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Spaces stores the value of this header in the object metadata. ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in Spaces`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Spaces stores the value of this header in the object metadata. ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_bucket_objects",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Lists object keys in this Spaces bucket`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The slug of the region where the bucket is stored.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Limits results to object keys with this prefix (Default: none)`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) A character used to group keys (Default: none)`,
				},
				resource.Attribute{
					Name:        "encoding_type",
					Description: `(Optional) Encodes keys using this method (Default: none; besides none, only "url" can be used)`,
				},
				resource.Attribute{
					Name:        "max_keys",
					Description: `(Optional) Maximum object keys to return (Default: 1000) ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `List of strings representing object keys`,
				},
				resource.Attribute{
					Name:        "common_prefixes",
					Description: `List of any keys between ` + "`" + `prefix` + "`" + ` and the next occurrence of ` + "`" + `delimiter` + "`" + ` (i.e., similar to subdirectories of the ` + "`" + `prefix` + "`" + ` "directory"); the list is only returned when you specify ` + "`" + `delimiter` + "`" + ``,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `List of strings representing object owner IDs`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keys",
					Description: `List of strings representing object keys`,
				},
				resource.Attribute{
					Name:        "common_prefixes",
					Description: `List of any keys between ` + "`" + `prefix` + "`" + ` and the next occurrence of ` + "`" + `delimiter` + "`" + ` (i.e., similar to subdirectories of the ` + "`" + `prefix` + "`" + ` "directory"); the list is only returned when you specify ` + "`" + `delimiter` + "`" + ``,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `List of strings representing object owner IDs`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_spaces_buckets",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the images by this key. This may be one of ` + "`" + `bucket_domain_name` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `region` + "`" + `, or ` + "`" + `urn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves Spaces buckets where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) One of ` + "`" + `exact` + "`" + ` (default), ` + "`" + `re` + "`" + `, or ` + "`" + `substring` + "`" + `. For string-typed fields, specify ` + "`" + `re` + "`" + ` to match by using the ` + "`" + `values` + "`" + ` as regular expressions, or specify ` + "`" + `substring` + "`" + ` to match by treating the ` + "`" + `values` + "`" + ` as substrings to find within the string field.`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to require that a field match all of the ` + "`" + `values` + "`" + ` instead of just one or more of them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure that all of the ` + "`" + `values` + "`" + ` are present in the list or set. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the images by this key. This may be one of ` + "`" + `bucket_domain_name` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `region` + "`" + `, or ` + "`" + `urn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "buckets",
					Description: `A list of Spaces buckets satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each bucket has the following attributes: - ` + "`" + `name` + "`" + ` - The name of the Spaces bucket - ` + "`" + `region` + "`" + ` - The slug of the region where the bucket is stored. - ` + "`" + `urn` + "`" + ` - The uniform resource name of the bucket - ` + "`" + `bucket_domain_name` + "`" + ` - The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com) - ` + "`" + `endpoint` + "`" + ` - The FQDN of the bucket without the bucket name (e.g. nyc3.digitaloceanspaces.com)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "buckets",
					Description: `A list of Spaces buckets satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each bucket has the following attributes: - ` + "`" + `name` + "`" + ` - The name of the Spaces bucket - ` + "`" + `region` + "`" + ` - The slug of the region where the bucket is stored. - ` + "`" + `urn` + "`" + ` - The uniform resource name of the bucket - ` + "`" + `bucket_domain_name` + "`" + ` - The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com) - ` + "`" + `endpoint` + "`" + ` - The FQDN of the bucket without the bucket name (e.g. nyc3.digitaloceanspaces.com)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_ssh_key",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ssh key. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_ssh_keys",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the SSH Keys by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `public_key` + "`" + `, or ` + "`" + `fingerprint` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the key field. Only retrieves SSH keys where the key field matches one or more of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the SSH Keys by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `public_key` + "`" + `, or ` + "`" + `fingerprint` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `A list of SSH Keys. Each SSH Key has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ssh key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `A list of SSH Keys. Each SSH Key has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ssh key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_tag",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tag. ## Attributes Reference The following attributes are exported:`,
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
					Description: `A count of the database clusters that the tag is applied to.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `A count of the database clusters that the tag is applied to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_tags",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the tags by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `total_resource_count` + "`" + `, ` + "`" + `droplets_count` + "`" + `, ` + "`" + `images_count` + "`" + `, ` + "`" + `volumes_count` + "`" + `, ` + "`" + `volume_snapshots_count` + "`" + `, or ` + "`" + `databases_count` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Only retrieves tags which keys has value that matches one of the values provided here.`,
				},
				resource.Attribute{
					Name:        "match_by",
					Description: `(Optional) One of ` + "`" + `exact` + "`" + ` (default), ` + "`" + `re` + "`" + `, or ` + "`" + `substring` + "`" + `. For string-typed fields, specify ` + "`" + `re` + "`" + ` to match by using the ` + "`" + `values` + "`" + ` as regular expressions, or specify ` + "`" + `substring` + "`" + ` to match by treating the ` + "`" + `values` + "`" + ` as substrings to find within the string field.`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to require that a field match all of the ` + "`" + `values` + "`" + ` instead of just one or more of them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure that all of the ` + "`" + `values` + "`" + ` are present in the list or set. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the tags by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `total_resource_count` + "`" + `, ` + "`" + `droplets_count` + "`" + `, ` + "`" + `images_count` + "`" + `, ` + "`" + `volumes_count` + "`" + `, ` + "`" + `volume_snapshots_count` + "`" + `, or ` + "`" + `databases_count` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference The following attributes are exported for each tag:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
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
					Description: `A count of the database clusters that the tag is applied to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
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
					Description: `A count of the database clusters that the tag is applied to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_volume",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of block storage volume.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region the block storage volume is provisioned in. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the block storage volume in GiB.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Text describing a block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_type",
					Description: `Filesystem type currently in-use on the block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_label",
					Description: `Filesystem label currently in-use on the block storage volume.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `A list of associated Droplet ids.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the Volume.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `The size of the block storage volume in GiB.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Text describing a block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_type",
					Description: `Filesystem type currently in-use on the block storage volume.`,
				},
				resource.Attribute{
					Name:        "filesystem_label",
					Description: `Filesystem label currently in-use on the block storage volume.`,
				},
				resource.Attribute{
					Name:        "droplet_ids",
					Description: `A list of associated Droplet ids.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the Volume.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_volume_snapshot",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the volume snapshot.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the volume snapshot list returned by DigitalOcean. This allows more advanced filtering not supported from the DigitalOcean API. This filtering is done locally on what DigitalOcean returns.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) A "slug" representing a DigitalOcean region (e.g. ` + "`" + `nyc1` + "`" + `). If set, only volume snapshots available in the region will be returned.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent volume snapshot. ~>`,
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
					Name:        "volume_id",
					Description: `The ID of the volume from which the volume snapshot originated.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the volume snapshot in gigabytes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the volume snapshot.`,
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
					Name:        "volume_id",
					Description: `The ID of the volume from which the volume snapshot originated.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The billable size of the volume snapshot in gigabytes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of the tags associated to the volume snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "digitalocean_vpc",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of an existing VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of an existing VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The DigitalOcean region slug for the VPC's location. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The DigitalOcean region slug for the VPC's location.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free-form text field describing the VPC.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The range of IP addresses for the VPC in CIDR notation.`,
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
					Description: `The date and time of when the VPC was created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The DigitalOcean region slug for the VPC's location.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free-form text field describing the VPC.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The range of IP addresses for the VPC in CIDR notation.`,
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
					Description: `The date and time of when the VPC was created.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"digitalocean_account":               0,
		"digitalocean_app":                   1,
		"digitalocean_certificate":           2,
		"digitalocean_container_registry":    3,
		"digitalocean_database_ca":           4,
		"digitalocean_database_cluster":      5,
		"digitalocean_database_replica":      6,
		"digitalocean_database_user":         7,
		"digitalocean_domain":                8,
		"digitalocean_domains":               9,
		"digitalocean_droplet":               10,
		"digitalocean_droplet_snapshot":      11,
		"digitalocean_droplets":              12,
		"digitalocean_firewall":              13,
		"digitalocean_floating_ip":           14,
		"digitalocean_image":                 15,
		"digitalocean_images":                16,
		"digitalocean_kubernetes_cluster":    17,
		"digitalocean_kubernetes_versions":   18,
		"digitalocean_loadbalancer":          19,
		"digitalocean_project":               20,
		"digitalocean_projects":              21,
		"digitalocean_record":                22,
		"digitalocean_records":               23,
		"digitalocean_region":                24,
		"digitalocean_regions":               25,
		"digitalocean_reserved_ip":           26,
		"digitalocean_sizes":                 27,
		"digitalocean_spaces_bucket":         28,
		"digitalocean_spaces_bucket_object":  29,
		"digitalocean_spaces_bucket_objects": 30,
		"digitalocean_spaces_buckets":        31,
		"digitalocean_ssh_key":               32,
		"digitalocean_ssh_keys":              33,
		"digitalocean_tag":                   34,
		"digitalocean_tags":                  35,
		"digitalocean_volume":                36,
		"digitalocean_volume_snapshot":       37,
		"digitalocean_vpc":                   38,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
