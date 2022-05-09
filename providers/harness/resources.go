package harness

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "harness_add_user_to_group",
			Category:         "Resources",
			ShortDescription: `Resource for adding a user to a group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_application",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Harness application`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_application_gitsync",
			Category:         "Resources",
			ShortDescription: `Resource for configuring application git sync.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_cloudprovider_aws",
			Category:         "Resources",
			ShortDescription: `Resource for creating an AWS cloud provider. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_cloudprovider_azure",
			Category:         "Resources",
			ShortDescription: `Resource for creating an Azure cloud provider. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_cloudprovider_datacenter",
			Category:         "Resources",
			ShortDescription: `Resource for creating a physical data center cloud provider. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_cloudprovider_gcp",
			Category:         "Resources",
			ShortDescription: `Resource for creating a GCP cloud provider. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_cloudprovider_kubernetes",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Kubernetes cloud provider. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_cloudprovider_spot",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Spot cloud provider. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_cloudprovider_tanzu",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Tanzu cloud provider. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_delegate_approval",
			Category:         "Resources",
			ShortDescription: `Resource for approving or rejecting delegates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_encrypted_text",
			Category:         "Resources",
			ShortDescription: `Resource for creating an encrypted text secret`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_environment",
			Category:         "Resources",
			ShortDescription: `Resource for creating an environment`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_git_connector",
			Category:         "Resources",
			ShortDescription: `Resource for creating a git connector`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_infrastructure_definition",
			Category:         "Resources",
			ShortDescription: `Resource for creating am infrastructure definition. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_appdynamics",
			Category:         "Resources",
			ShortDescription: `Resource for creating an App Dynamics connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_artifactory",
			Category:         "Resources",
			ShortDescription: `Resource for creating an Artifactory connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_aws",
			Category:         "Resources",
			ShortDescription: `Resource for creating an AWS connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_aws_secret_manager",
			Category:         "Resources",
			ShortDescription: `Resource for creating an AWS Secret Manager connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_awscc",
			Category:         "Resources",
			ShortDescription: `Resource for creating an AWS Cloud Cost connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_awskms",
			Category:         "Resources",
			ShortDescription: `Resource for creating an AWS KMS connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_bitbucket",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Bitbucket connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_datadog",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Datadog connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_docker",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Docker connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_dynatrace",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Dynatrace connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_gcp",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Gcp connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_git",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Git connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_github",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Github connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_gitlab",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Gitlab connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_helm",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Helm connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_jira",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Jira connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_kubernetes",
			Category:         "Resources",
			ShortDescription: `Resource for creating a K8s connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_newrelic",
			Category:         "Resources",
			ShortDescription: `Resource for creating a New Relic connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_nexus",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Nexus connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_pagerduty",
			Category:         "Resources",
			ShortDescription: `Resource for creating a PagerDuty connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_prometheus",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Prometheus connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_splunk",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Splunk connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_connector_sumologic",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Sumologic connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_environment",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Harness environment.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_organization",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Harness organization.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_pipeline",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Harness pipeline.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_project",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Harness project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_platform_service",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Harness project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_service_ami",
			Category:         "Resources",
			ShortDescription: `Resource for creating an AMI service. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_service_aws_codedeploy",
			Category:         "Resources",
			ShortDescription: `Resource for creating an AWS CodeDeploy service. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_service_aws_lambda",
			Category:         "Resources",
			ShortDescription: `Resource for creating an AWS Lambda service. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_service_ecs",
			Category:         "Resources",
			ShortDescription: `Resource for creating an AWS ECS service. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_service_helm",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Kubernetes Helm service. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_service_kubernetes",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Kubernetes service. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_service_ssh",
			Category:         "Resources",
			ShortDescription: `Resource for creating an SSH service. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_service_tanzu",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Tanzu (PCF) service. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_service_winrm",
			Category:         "Resources",
			ShortDescription: `Resource for creating an WinRM service. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_ssh_credential",
			Category:         "Resources",
			ShortDescription: `Resource for creating an encrypted text secret`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_user",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Harness user`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_user_group",
			Category:         "Resources",
			ShortDescription: `Resource for creating a Harness user group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_user_group_permissions",
			Category:         "Resources",
			ShortDescription: `Resource for adding permissions to an existing Harness user group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "harness_yaml_config",
			Category:         "Resources",
			ShortDescription: `Resource for creating a raw YAML configuration in Harness. Note: This works for all objects EXCEPT application objects. This resource uses the config-as-code API's. When updating the name or path of this resource you should typically also set the create_before_destroy = true lifecycle setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"harness_add_user_to_group":                     0,
		"harness_application":                           1,
		"harness_application_gitsync":                   2,
		"harness_cloudprovider_aws":                     3,
		"harness_cloudprovider_azure":                   4,
		"harness_cloudprovider_datacenter":              5,
		"harness_cloudprovider_gcp":                     6,
		"harness_cloudprovider_kubernetes":              7,
		"harness_cloudprovider_spot":                    8,
		"harness_cloudprovider_tanzu":                   9,
		"harness_delegate_approval":                     10,
		"harness_encrypted_text":                        11,
		"harness_environment":                           12,
		"harness_git_connector":                         13,
		"harness_infrastructure_definition":             14,
		"harness_platform_connector_appdynamics":        15,
		"harness_platform_connector_artifactory":        16,
		"harness_platform_connector_aws":                17,
		"harness_platform_connector_aws_secret_manager": 18,
		"harness_platform_connector_awscc":              19,
		"harness_platform_connector_awskms":             20,
		"harness_platform_connector_bitbucket":          21,
		"harness_platform_connector_datadog":            22,
		"harness_platform_connector_docker":             23,
		"harness_platform_connector_dynatrace":          24,
		"harness_platform_connector_gcp":                25,
		"harness_platform_connector_git":                26,
		"harness_platform_connector_github":             27,
		"harness_platform_connector_gitlab":             28,
		"harness_platform_connector_helm":               29,
		"harness_platform_connector_jira":               30,
		"harness_platform_connector_kubernetes":         31,
		"harness_platform_connector_newrelic":           32,
		"harness_platform_connector_nexus":              33,
		"harness_platform_connector_pagerduty":          34,
		"harness_platform_connector_prometheus":         35,
		"harness_platform_connector_splunk":             36,
		"harness_platform_connector_sumologic":          37,
		"harness_platform_environment":                  38,
		"harness_platform_organization":                 39,
		"harness_platform_pipeline":                     40,
		"harness_platform_project":                      41,
		"harness_platform_service":                      42,
		"harness_service_ami":                           43,
		"harness_service_aws_codedeploy":                44,
		"harness_service_aws_lambda":                    45,
		"harness_service_ecs":                           46,
		"harness_service_helm":                          47,
		"harness_service_kubernetes":                    48,
		"harness_service_ssh":                           49,
		"harness_service_tanzu":                         50,
		"harness_service_winrm":                         51,
		"harness_ssh_credential":                        52,
		"harness_user":                                  53,
		"harness_user_group":                            54,
		"harness_user_group_permissions":                55,
		"harness_yaml_config":                           56,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
