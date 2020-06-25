package helm

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "helm_release",
			Category:         "Helm Provider",
			ShortDescription: ``,
			Description: `

A Release is an instance of a chart running in a Kubernetes cluster.
A Chart is a Helm package. It contains all of the resource definitions necessary to run an application, tool, or service inside of a Kubernetes cluster.

` + "`" + `helm_release` + "`" + ` describes the desired status of a chart in a kubernetes cluster.

`,
			Keywords: []string{
				"provider",
				"release",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Release name.`,
				},
				resource.Attribute{
					Name:        "chart",
					Description: `(Required) Chart name to be installed. A path may be used.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Optional) Repository URL where to locate the requested chart.`,
				},
				resource.Attribute{
					Name:        "repository_key_file",
					Description: `(Optional) The repositories cert key file`,
				},
				resource.Attribute{
					Name:        "repository_cert_file",
					Description: `(Optional) The repositories cert file`,
				},
				resource.Attribute{
					Name:        "repository_ca_file",
					Description: `(Optional) The Repositories CA File.`,
				},
				resource.Attribute{
					Name:        "repository_username",
					Description: `(Optional) Username for HTTP basic authentication against the repository.`,
				},
				resource.Attribute{
					Name:        "repository_password",
					Description: `(Optional) Password for HTTP basic authentication against the repository.`,
				},
				resource.Attribute{
					Name:        "devel",
					Description: `(Optional) Use chart development versions, too. Equivalent to version '>0.0.0-0'. If version is set, this is ignored.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Specify the exact chart version to install. If this is not specified, the latest version is installed.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to install the release into. Defaults to ` + "`" + `default` + "`" + ``,
				},
				resource.Attribute{
					Name:        "verify",
					Description: `(Optional) Verify the package before installing it. Defaults to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "keyring",
					Description: `(Optional) Location of public keys used for verification. Used only if ` + "`" + `verify` + "`" + ` is true. Defaults to ` + "`" + `/.gnupg/pubring.gpg` + "`" + ` in the location set by ` + "`" + `home` + "`" + ``,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Time in seconds to wait for any individual kubernetes operation (like Jobs for hooks). Defaults to ` + "`" + `300` + "`" + ` seconds.`,
				},
				resource.Attribute{
					Name:        "disable_webhooks",
					Description: `(Optional) Prevent hooks from running. Defauts to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "reuse_values",
					Description: `(Optional) When upgrading, reuse the last release's values and merge in any overrides. If 'reset_values' is specified, this is ignored. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reset_values",
					Description: `(Optional) When upgrading, reset the values to the ones built into the chart. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_update",
					Description: `(Optional) Force resource update through delete/recreate if needed. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "recreate_pods",
					Description: `(Optional) Perform pods restart during upgrade/rollback. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cleanup_on_fail",
					Description: `(Optional) Allow deletion of new resources created in this upgrade when upgrade fails. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_history",
					Description: `(Optional) Maximum number of release versions stored per release. Defaults to ` + "`" + `0` + "`" + ` (no limit).`,
				},
				resource.Attribute{
					Name:        "atomic",
					Description: `(Optional) If set, installation process purges chart on fail. The wait flag will be set automatically if atomic is used. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "skip_crds",
					Description: `(Optional) If set, no CRDs will be installed. By default, CRDs are installed if not already present. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "render_subchart_notes",
					Description: `(Optional) If set, render subchart notes along with the parent. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_openapi_validation",
					Description: `(Optional) If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "wait",
					Description: `(Optional) Will wait until all resources are in a ready state before marking the release as successful. It will wait for as long as ` + "`" + `timeout` + "`" + `. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) List of values in raw yaml to pass to helm. Values will be merged, in order, as Helm does with multiple ` + "`" + `-f` + "`" + ` options.`,
				},
				resource.Attribute{
					Name:        "set",
					Description: `(Optional) Value block with custom values to be merged with the values yaml.`,
				},
				resource.Attribute{
					Name:        "set_sensitive",
					Description: `(Optional) Value block with custom sensitive values to be merged with the values yaml that won't be exposed in the plan's diff.`,
				},
				resource.Attribute{
					Name:        "set_string",
					Description: `(Optional) Value block with custom STRING values to be merged with the values yaml.`,
				},
				resource.Attribute{
					Name:        "dependency_update",
					Description: `(Optional) Runs helm dependency update before installing the chart. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) Re-use the given name, even if that name is already used. This is unsafe in production. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Set release description attribute (visible in the history).`,
				},
				resource.Attribute{
					Name:        "postrender",
					Description: `(Optional) Configure a command to run after helm renders the manifest which can alter the manifest contents.`,
				},
				resource.Attribute{
					Name:        "lint",
					Description: `(Optional) Run the helm chart linter during the plan. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_namespace",
					Description: `(Optional) Create the namespace if it does not yet exist. Defaults to ` + "`" + `false` + "`" + `. The ` + "`" + `set` + "`" + `, ` + "`" + `set_sensitive` + "`" + ` and ` + "`" + `set_strings` + "`" + ` blocks support:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) full name of the variable to be set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) value of the variable to be set. The ` + "`" + `postrender` + "`" + ` block supports a single attribute:`,
				},
				resource.Attribute{
					Name:        "binary_path",
					Description: `(Required) relative or full path to command binary. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Block status of the deployed release. The ` + "`" + `metadata` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "chart",
					Description: `The name of the chart.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name is the name of the release.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace is the kubernetes namespace of the release.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Version is an int32 which represents the version of the release.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the release.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `A SemVer 2 conformant version string of the chart.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `The compounded values from ` + "`" + `values` + "`" + ` and ` + "`" + `set`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `Block status of the deployed release. The ` + "`" + `metadata` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "chart",
					Description: `The name of the chart.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name is the name of the release.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace is the kubernetes namespace of the release.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Version is an int32 which represents the version of the release.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the release.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `A SemVer 2 conformant version string of the chart.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `The compounded values from ` + "`" + `values` + "`" + ` and ` + "`" + `set`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"helm_release": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
