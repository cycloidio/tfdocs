package genesyscloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_architect_datatable",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Architect Datatables`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_architect_datatable_row",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Architect Datatable Row`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_architect_emergencygroup",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Architect Emergency Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_architect_ivr",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud IVR config`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_architect_schedulegroups",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Architect Schedule Groups`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_architect_schedules",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Architect Schedules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_architect_user_prompt",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud User Audio Prompt`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_auth_division",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Authorization Division`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_auth_role",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Authorization Role`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_flow",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Flow NOTE: This component is currently in beta. If you are interested in participating in the beta, please contact Becky Powell (becky.powell@genesys.com) to be added to the beta. If you attempt to use this resource and you are not part of the beta program, your flow deploy will fail with HTTP status 403, no permissions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_group",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Directory Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_group_roles",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Group Roles maintains group role assignments.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_idp_adfs",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Single Sign-on ADFS Identity Provider. See this page for detailed configuration instructions: https://help.mypurecloud.com/articles/add-microsoft-adfs-single-sign-provider/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_idp_generic",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Single Sign-on Generic Identity Provider. See this page for detailed configuration instructions: https://help.mypurecloud.com/articles/add-a-generic-single-sign-on-provider/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_idp_gsuite",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Single Sign-on GSuite Identity Provider. See this page for detailed configuration instructions: https://help.mypurecloud.com/articles/add-google-g-suite-single-sign-provider/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_idp_okta",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Single Sign-on Okta Identity Provider. See this page for detailed configuration instructions: https://help.mypurecloud.com/articles/add-okta-as-a-single-sign-on-provider/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_idp_onelogin",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Single Sign-on OneLogin Identity Provider. See this page for detailed configuration instructions: https://help.mypurecloud.com/articles/add-onelogin-as-single-sign-on-provider/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_idp_ping",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Single Sign-on Ping Identity Provider. See this page for detailed configuration instructions: https://help.mypurecloud.com/articles/add-ping-identity-single-sign-provider/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_idp_salesforce",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Single Sign-on Salesforce Identity Provider. See this page for detailed configuration instructions: https://help.mypurecloud.com/articles/add-salesforce-as-a-single-sign-on-provider/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_integration",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Integration`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_integration_action",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Integration Actions. See this page for detailed information on configuring Actions: https://help.mypurecloud.com/articles/add-configuration-custom-actions-integrations/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_integration_credential",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Credential`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_location",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Location`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_oauth_client",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud OAuth Clients. See this page for detailed configuration information: https://help.mypurecloud.com/articles/create-an-oauth-client/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_processautomation_trigger",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Process Automation Trigger NOTE: This component is currently in beta. If you wish to use this provider make sure your client has the correct permissions`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_quality_forms_evaluation",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Evaluation Forms`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_routing_email_domain",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Routing Email Domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_routing_email_route",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Routing Email Domain Route`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_routing_language",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Routing Language`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_routing_queue",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Routing Queue`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_routing_skill",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Routing Skill`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_routing_utilization",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Org-wide Routing Utilization Settings.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_routing_wrapupcode",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Routing Wrapup Code`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_telephony_providers_edges_did_pool",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud DID Pool`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_telephony_providers_edges_edge_group",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Edge Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_telephony_providers_edges_extension_pool",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Extension Pool`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_telephony_providers_edges_phone",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Phone`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_telephony_providers_edges_phonebasesettings",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Phone Base Settings`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_telephony_providers_edges_site",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Site`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_telephony_providers_edges_trunk",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Trunk. Created by assigning a trunk base settings to an edge or edge group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_telephony_providers_edges_trunkbasesettings",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Trunk Base Settings`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_tf_export",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Resource to export Terraform config and (optionally) tfstate files to a local directory. The config file is named 'genesyscloud.tf.json' or 'genesyscloud.tf', and the state file is named 'terraform.tfstate'.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_user",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud User`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_user_roles",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud User Roles maintains user role assignments.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_webdeployments_configuration",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Web Deployment Configuration`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_webdeployments_deployment",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Web Deployment`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_widget_deployment",
			Category:         "Resources",
			ShortDescription: `Genesys Cloud Widget Deployment`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"genesyscloud_architect_datatable":                         0,
		"genesyscloud_architect_datatable_row":                     1,
		"genesyscloud_architect_emergencygroup":                    2,
		"genesyscloud_architect_ivr":                               3,
		"genesyscloud_architect_schedulegroups":                    4,
		"genesyscloud_architect_schedules":                         5,
		"genesyscloud_architect_user_prompt":                       6,
		"genesyscloud_auth_division":                               7,
		"genesyscloud_auth_role":                                   8,
		"genesyscloud_flow":                                        9,
		"genesyscloud_group":                                       10,
		"genesyscloud_group_roles":                                 11,
		"genesyscloud_idp_adfs":                                    12,
		"genesyscloud_idp_generic":                                 13,
		"genesyscloud_idp_gsuite":                                  14,
		"genesyscloud_idp_okta":                                    15,
		"genesyscloud_idp_onelogin":                                16,
		"genesyscloud_idp_ping":                                    17,
		"genesyscloud_idp_salesforce":                              18,
		"genesyscloud_integration":                                 19,
		"genesyscloud_integration_action":                          20,
		"genesyscloud_integration_credential":                      21,
		"genesyscloud_location":                                    22,
		"genesyscloud_oauth_client":                                23,
		"genesyscloud_processautomation_trigger":                   24,
		"genesyscloud_quality_forms_evaluation":                    25,
		"genesyscloud_routing_email_domain":                        26,
		"genesyscloud_routing_email_route":                         27,
		"genesyscloud_routing_language":                            28,
		"genesyscloud_routing_queue":                               29,
		"genesyscloud_routing_skill":                               30,
		"genesyscloud_routing_utilization":                         31,
		"genesyscloud_routing_wrapupcode":                          32,
		"genesyscloud_telephony_providers_edges_did_pool":          33,
		"genesyscloud_telephony_providers_edges_edge_group":        34,
		"genesyscloud_telephony_providers_edges_extension_pool":    35,
		"genesyscloud_telephony_providers_edges_phone":             36,
		"genesyscloud_telephony_providers_edges_phonebasesettings": 37,
		"genesyscloud_telephony_providers_edges_site":              38,
		"genesyscloud_telephony_providers_edges_trunk":             39,
		"genesyscloud_telephony_providers_edges_trunkbasesettings": 40,
		"genesyscloud_tf_export":                                   41,
		"genesyscloud_user":                                        42,
		"genesyscloud_user_roles":                                  43,
		"genesyscloud_webdeployments_configuration":                44,
		"genesyscloud_webdeployments_deployment":                   45,
		"genesyscloud_widget_deployment":                           46,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
