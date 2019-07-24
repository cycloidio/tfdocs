package rundeck

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rundeck_acl_policy",
			Category:         "Resources",
			ShortDescription: `The rundeck_acl_policy resource allows Rundeck ACLs to be managed by Terraform.`,
			Description:      ``,
			Keywords: []string{
				"acl",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy. Must end with ` + "`" + `.aclpolicy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The name of the job, used to describe the job in the Rundeck UI.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rundeck_job",
			Category:         "Resources",
			ShortDescription: `The rundeck_job resource allows Rundeck jobs to be managed by Terraform.`,
			Description:      ``,
			Keywords: []string{
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the job, used to describe the job in the Rundeck UI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A longer description of the job, describing the job in the Rundeck UI.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `(Required) The name of the project that this job should belong to.`,
				},
				resource.Attribute{
					Name:        "execution_enabled",
					Description: `(Optional) If you want job execution to be enabled or disabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Optional) The name of a group within the project in which to place the job. Setting this creates collapsable subcategories within the Rundeck UI's project job index.`,
				},
				resource.Attribute{
					Name:        "log_level",
					Description: `(Optional) The log level that Rundeck should use for this job. Defaults to "INFO".`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Optional) The jobs schedule in Unix crontab format`,
				},
				resource.Attribute{
					Name:        "schedule_enabled",
					Description: `(Optional) Sets the job schedule to be enabled or disabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_concurrent_executions",
					Description: `(Optional) Boolean defining whether two or more executions of this job can run concurrently. The default is ` + "`" + `false` + "`" + `, meaning that jobs will only run sequentially.`,
				},
				resource.Attribute{
					Name:        "max_thread_count",
					Description: `(Optional) The maximum number of threads to use to execute this job, which controls on how many nodes the commands can be run simulateneously. Defaults to 1, meaning that the nodes will be visited sequentially.`,
				},
				resource.Attribute{
					Name:        "continue_on_error",
					Description: `(Optional) Boolean defining whether Rundeck will continue to run subsequent steps if any intermediate step fails. Defaults to ` + "`" + `false` + "`" + `, meaning that execution will stop and the execution will be considered to have failed.`,
				},
				resource.Attribute{
					Name:        "rank_attribute",
					Description: `(Optional) The name of the attribute that will be used to decide in which order the nodes will be visited while executing the job across multiple nodes.`,
				},
				resource.Attribute{
					Name:        "rank_order",
					Description: `(Optional) Keyword deciding which direction the nodes are sorted in terms of the chosen ` + "`" + `rank_attribute` + "`" + `. May be either "ascending" (the default) or "descending".`,
				},
				resource.Attribute{
					Name:        "node_filter_query",
					Description: `(Optional) A query string using [Rundeck's node filter language](http://rundeck.org/docs/manual/node-filters.html#node-filter-syntax) that defines which subset of the project's nodes will be used to execute this job.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The name of the plugin to use.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required) Map of arbitrary configuration properties for the selected plugin. ## Attributes Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the job.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the job.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rundeck_private_key",
			Category:         "Resources",
			ShortDescription: `The rundeck_private_key resource allows private keys to be stored in Rundeck's key store.`,
			Description:      ``,
			Keywords: []string{
				"private",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path within the key store where the key will be stored.`,
				},
				resource.Attribute{
					Name:        "key_material",
					Description: `(Required) The private key material to store, serialized in any way that is accepted by OpenSSH. The key material is hashed before it is stored in the state file, so sharing the resulting state will not disclose the private key contents. ## Attributes Reference Rundeck does not allow stored private keys to be retrieved via the API, so this resource does not export any attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rundeck_project",
			Category:         "Resources",
			ShortDescription: `The rundeck_project resource allows Rundeck projects to be managed by Terraform.`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project, used both in the UI and to uniquely identify the project. Must therefore be unique across a single Rundeck installation.`,
				},
				resource.Attribute{
					Name:        "resource_model_source",
					Description: `(Required) Nested block instructing Rundeck on how to determine the set of resources (nodes) for this project. The nested block structure is described below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the project, to be displayed in the Rundeck UI. Defaults to "Managed by Terraform".`,
				},
				resource.Attribute{
					Name:        "default_node_file_copier_plugin",
					Description: `(Optional) The name of a plugin to use to copy files onto nodes within this project. Defaults to ` + "`" + `jsch-scp` + "`" + `, which uses the "Secure Copy" protocol to send files over SSH.`,
				},
				resource.Attribute{
					Name:        "default_node_executor_plugin",
					Description: `(Optional) The name of a plugin to use to run commands on nodes within this project. Defaults to ` + "`" + `jsch-ssh` + "`" + `, which uses the SSH protocol to access the nodes.`,
				},
				resource.Attribute{
					Name:        "ssh_authentication_type",
					Description: `(Optional) When the SSH-based file copier and executor plugins are used, the type of SSH authentication to use. Defaults to ` + "`" + `privateKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssh_key_storage_path",
					Description: `(Optional) When the SSH-based file copier and executor plugins are used, the location within Rundeck's key store where the SSH private key can be found. Private keys can be uploaded to rundeck using the ` + "`" + `rundeck_private_key` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "ssh_key_file_path",
					Description: `(Optional) Like ` + "`" + `ssh_key_storage_path` + "`" + ` except that the key is read from the Rundeck server's local filesystem, rather than from the key store.`,
				},
				resource.Attribute{
					Name:        "extra_config",
					Description: `(Optional) Behind the scenes a Rundeck project is really an arbitrary set of key/value pairs. This map argument allows setting any configuration properties that aren't explicitly supported by the other arguments described above, but due to limitations of Terraform the key names must be written with slashes in place of dots. Do not use this argument to set properties that the above arguments set, or undefined behavior will result. ` + "`" + `resource_model_source` + "`" + ` blocks have the following nested arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The name of the resource model plugin to use.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required) Map of arbitrary configuration properties for the selected resource model plugin. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The unique name that identifies the project, as set in the arguments.`,
				},
				resource.Attribute{
					Name:        "ui_url",
					Description: `The URL of the index page for this project in the Rundeck UI.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The unique name that identifies the project, as set in the arguments.`,
				},
				resource.Attribute{
					Name:        "ui_url",
					Description: `The URL of the index page for this project in the Rundeck UI.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rundeck_public_key",
			Category:         "Resources",
			ShortDescription: `The rundeck_public_key resource allows public keys to be stored in Rundeck's key store.`,
			Description:      ``,
			Keywords: []string{
				"public",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "delete",
					Description: `(Computed) True if the key should be deleted when the resource is deleted. Defaults to true if key_material is provided in the configuration.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path within the key store where the key will be stored. By convention this path name normally ends with ".pub" and otherwise has the same name as the associated private key.`,
				},
				resource.Attribute{
					Name:        "key_material",
					Description: `(Optional) The public key string to store, serialized in any way that is accepted by OpenSSH. If this is not included, ` + "`" + `` + "`" + `key_material` + "`" + `` + "`" + ` becomes an attribute that can be used to read the already-existing key material in the Rundeck store. The key material is included inline as a string, which is consistent with the way a public key is provided to the ` + "`" + `aws_key_pair` + "`" + `, ` + "`" + `cloudstack_ssh_keypair` + "`" + `, ` + "`" + `digitalocean_ssh_key` + "`" + ` and ` + "`" + `openstack_compute_keypair_v2` + "`" + ` resources. This means the ` + "`" + `key_material` + "`" + ` argument can be populated from the interpolation of the ` + "`" + `public_key` + "`" + ` attribute of such a keypair resource, or vice-versa. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL at which the key material can be retrieved from the key store by other clients.`,
				},
				resource.Attribute{
					Name:        "key_material",
					Description: `If ` + "`" + `key_material` + "`" + ` is omitted in the configuration, it becomes an attribute that exposes the key material already stored at the given ` + "`" + `path` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `The URL at which the key material can be retrieved from the key store by other clients.`,
				},
				resource.Attribute{
					Name:        "key_material",
					Description: `If ` + "`" + `key_material` + "`" + ` is omitted in the configuration, it becomes an attribute that exposes the key material already stored at the given ` + "`" + `path` + "`" + `.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"rundeck_acl_policy":  0,
		"rundeck_job":         1,
		"rundeck_private_key": 2,
		"rundeck_project":     3,
		"rundeck_public_key":  4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
