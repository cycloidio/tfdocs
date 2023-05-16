package octopusdeploy

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_aws_account",
			Category:         "Accounts",
			ShortDescription: `This resource manages AWS accounts in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"aws",
				"account",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_aws_elastic_container_registry",
			Category:         "Feeds",
			ShortDescription: `This resource manages an AWS Elastic Container Registry in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"feeds",
				"aws",
				"elastic",
				"container",
				"registry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_azure_cloud_service_deployment_target",
			Category:         "Deployment Targets",
			ShortDescription: `This resource manages Azure cloud service deployment targets in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"targets",
				"azure",
				"cloud",
				"service",
				"target",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_azure_service_fabric_cluster_deployment_target",
			Category:         "Deployment Targets",
			ShortDescription: `This resource manages Azure service fabric cluster deployment targets in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"targets",
				"azure",
				"service",
				"fabric",
				"cluster",
				"target",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_azure_service_principal",
			Category:         "Accounts",
			ShortDescription: `This resource manages Azure service principal accounts in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"azure",
				"service",
				"principal",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_azure_subscription_account",
			Category:         "Accounts",
			ShortDescription: `This resource manages Azure subscription accounts in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"azure",
				"subscription",
				"account",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_azure_web_app_deployment_target",
			Category:         "Deployment Targets",
			ShortDescription: `This resource manages Azure web app deployment targets in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"targets",
				"azure",
				"web",
				"app",
				"target",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_certificate",
			Category:         "Resources",
			ShortDescription: `This resource manages certificates in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_channel",
			Category:         "Resources",
			ShortDescription: `This resource manages channels in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_cloud_region_deployment_target",
			Category:         "Deployment Targets",
			ShortDescription: `This resource manages cloud region deployment targets in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"targets",
				"cloud",
				"region",
				"target",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_deployment_process",
			Category:         "Resources",
			ShortDescription: `This resource manages deployment processes in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_docker_container_registry",
			Category:         "Feeds",
			ShortDescription: `This resource manages a Docker Container Registry in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"feeds",
				"docker",
				"container",
				"registry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_dynamic_worker_pool",
			Category:         "Worker Pools",
			ShortDescription: `This resource manages dynamic worker pools in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"worker",
				"pools",
				"dynamic",
				"pool",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_environment",
			Category:         "Resources",
			ShortDescription: `This resource manages environments in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_gcp_account",
			Category:         "Accounts",
			ShortDescription: `This resource manages GCP accounts in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"gcp",
				"account",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_git_credential",
			Category:         "Resources",
			ShortDescription: `This resource manages Git credentials in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_github_repository_feed",
			Category:         "Feeds",
			ShortDescription: `This resource manages a GitHub repository feed in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"feeds",
				"github",
				"repository",
				"feed",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_helm_feed",
			Category:         "Feeds",
			ShortDescription: `This resource manages a Helm feed in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"feeds",
				"helm",
				"feed",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_kubernetes_cluster_deployment_target",
			Category:         "Deployment Targets",
			ShortDescription: `This resource manages Kubernets cluster deployment targets in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"targets",
				"kubernetes",
				"cluster",
				"target",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_library_variable_set",
			Category:         "Resources",
			ShortDescription: `This resource manages library variable sets in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_lifecycle",
			Category:         "Resources",
			ShortDescription: `This resource manages lifecycles in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_listening_tentacle_deployment_target",
			Category:         "Deployment Targets",
			ShortDescription: `This resource manages listening tentacle deployment targets in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"targets",
				"listening",
				"tentacle",
				"target",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_machine_policy",
			Category:         "Resources",
			ShortDescription: `This resource manages machine policies in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_maven_feed",
			Category:         "Feeds",
			ShortDescription: `This resource manages a Maven feed in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"feeds",
				"maven",
				"feed",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_nuget_feed",
			Category:         "Feeds",
			ShortDescription: `This resource manages a NuGet feed in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"feeds",
				"nuget",
				"feed",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_offline_package_drop_deployment_target",
			Category:         "Deployment Targets",
			ShortDescription: `This resource manages offline package drop deployment targets in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"targets",
				"offline",
				"package",
				"drop",
				"target",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_polling_tentacle_deployment_target",
			Category:         "Deployment Targets",
			ShortDescription: `This resource manages polling tentacle deployment targets in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"targets",
				"polling",
				"tentacle",
				"target",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_project",
			Category:         "Resources",
			ShortDescription: `This resource manages projects in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_project_deployment_target_trigger",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_project_group",
			Category:         "Resources",
			ShortDescription: `This resource manages project groups in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_runbook",
			Category:         "Resources",
			ShortDescription: `This resource manages runbooks in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_runbook_process",
			Category:         "Resources",
			ShortDescription: `This resource manages runbook processes in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_scoped_user_role",
			Category:         "Resources",
			ShortDescription: `This resource manages scoped user roles in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_script_module",
			Category:         "Resources",
			ShortDescription: `This resource manages script modules in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_space",
			Category:         "Resources",
			ShortDescription: `This resource manages spaces in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_ssh_connection_deployment_target",
			Category:         "Deployment Targets",
			ShortDescription: `This resource manages SSH connection deployment targets in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"targets",
				"ssh",
				"connection",
				"target",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_ssh_key_account",
			Category:         "Accounts",
			ShortDescription: `This resource manages SSH key accounts in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"ssh",
				"key",
				"account",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_static_worker_pool",
			Category:         "Worker Pools",
			ShortDescription: `This resource manages static worker pools in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"worker",
				"pools",
				"static",
				"pool",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_tag",
			Category:         "Resources",
			ShortDescription: `This resource manages tags in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_tag_set",
			Category:         "Resources",
			ShortDescription: `This resource manages tag sets in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_team",
			Category:         "Resources",
			ShortDescription: `This resource manages teams in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_tenant",
			Category:         "Resources",
			ShortDescription: `This resource manages tenants in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_tenant_common_variable",
			Category:         "Resources",
			ShortDescription: `This resource manages tenant common variables in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_tenant_project_variable",
			Category:         "Resources",
			ShortDescription: `This resource manages tenant project variables in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_token_account",
			Category:         "Accounts",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"token",
				"account",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_user",
			Category:         "Resources",
			ShortDescription: `This resource manages users in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_user_role",
			Category:         "Resources",
			ShortDescription: `This resource manages user roles in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_username_password_account",
			Category:         "Accounts",
			ShortDescription: `This resource manages username-password accounts in Octopus Deploy.`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"username",
				"password",
				"account",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "octopusdeploy_variable",
			Category:         "Resources",
			ShortDescription: `This resource manages variables in Octopus Deploy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"octopusdeploy_aws_account":                                    0,
		"octopusdeploy_aws_elastic_container_registry":                 1,
		"octopusdeploy_azure_cloud_service_deployment_target":          2,
		"octopusdeploy_azure_service_fabric_cluster_deployment_target": 3,
		"octopusdeploy_azure_service_principal":                        4,
		"octopusdeploy_azure_subscription_account":                     5,
		"octopusdeploy_azure_web_app_deployment_target":                6,
		"octopusdeploy_certificate":                                    7,
		"octopusdeploy_channel":                                        8,
		"octopusdeploy_cloud_region_deployment_target":                 9,
		"octopusdeploy_deployment_process":                             10,
		"octopusdeploy_docker_container_registry":                      11,
		"octopusdeploy_dynamic_worker_pool":                            12,
		"octopusdeploy_environment":                                    13,
		"octopusdeploy_gcp_account":                                    14,
		"octopusdeploy_git_credential":                                 15,
		"octopusdeploy_github_repository_feed":                         16,
		"octopusdeploy_helm_feed":                                      17,
		"octopusdeploy_kubernetes_cluster_deployment_target":           18,
		"octopusdeploy_library_variable_set":                           19,
		"octopusdeploy_lifecycle":                                      20,
		"octopusdeploy_listening_tentacle_deployment_target":           21,
		"octopusdeploy_machine_policy":                                 22,
		"octopusdeploy_maven_feed":                                     23,
		"octopusdeploy_nuget_feed":                                     24,
		"octopusdeploy_offline_package_drop_deployment_target":         25,
		"octopusdeploy_polling_tentacle_deployment_target":             26,
		"octopusdeploy_project":                                        27,
		"octopusdeploy_project_deployment_target_trigger":              28,
		"octopusdeploy_project_group":                                  29,
		"octopusdeploy_runbook":                                        30,
		"octopusdeploy_runbook_process":                                31,
		"octopusdeploy_scoped_user_role":                               32,
		"octopusdeploy_script_module":                                  33,
		"octopusdeploy_space":                                          34,
		"octopusdeploy_ssh_connection_deployment_target":               35,
		"octopusdeploy_ssh_key_account":                                36,
		"octopusdeploy_static_worker_pool":                             37,
		"octopusdeploy_tag":                                            38,
		"octopusdeploy_tag_set":                                        39,
		"octopusdeploy_team":                                           40,
		"octopusdeploy_tenant":                                         41,
		"octopusdeploy_tenant_common_variable":                         42,
		"octopusdeploy_tenant_project_variable":                        43,
		"octopusdeploy_token_account":                                  44,
		"octopusdeploy_user":                                           45,
		"octopusdeploy_user_role":                                      46,
		"octopusdeploy_username_password_account":                      47,
		"octopusdeploy_variable":                                       48,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
