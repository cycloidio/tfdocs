package akamai

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_advanced_settings_logging",
			Category:         "Application Security",
			ShortDescription: `AdvancedSettingsLogging`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"advanced",
				"settings",
				"logging",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The configuration ID.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Optional) The ID of the security policy to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of information about the logging settings.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the logging settings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of information about the logging settings.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the logging settings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_advanced_settings_prefetch",
			Category:         "Application Security",
			ShortDescription: `AdvancedSettingsPrefetch`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"advanced",
				"settings",
				"prefetch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The configuration ID.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of information about the prefetch request settings.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the prefetch request settings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of information about the prefetch request settings.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the prefetch request settings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_api_endpoints",
			Category:         "Application Security",
			ShortDescription: `ApiEndpoints`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"api",
				"endpoints",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The configuration ID.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Optional) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `(Optional) The name of a specific endpoint. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id_list",
					Description: `A list of IDs of the API endpoints.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of information about the API endpoints.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and name of the API endpoints.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id_list",
					Description: `A list of IDs of the API endpoints.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of information about the API endpoints.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and name of the API endpoints.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_api_request_constraints",
			Category:         "Application Security",
			ShortDescription: `ApiRequestConstraints`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"api",
				"request",
				"constraints",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The configuration ID to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `(Optional) The ID of a specific API for which to retrieve constraint information. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of information about the APIs and their constraints and actions.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the APIs and their constraints and actions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of information about the APIs and their constraints and actions.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the APIs and their constraints and actions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_attack_group_actions",
			Category:         "Application Security",
			ShortDescription: `Attack Group Actions`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"attack",
				"group",
				"actions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "attack_group",
					Description: `(Optional) The attack group to use. If not supplied, information about all attack groups will be returned. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The attack group action for the attack group if one was specified: ` + "`" + `alert` + "`" + `, ` + "`" + `deny` + "`" + `, or ` + "`" + `none` + "`" + `. If the action is none, the attack group is inactive in the security policy.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ` + "`" + `action` + "`" + ` and ` + "`" + `group` + "`" + ` name for each attack group.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `The attack group information in JSON format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `The attack group action for the attack group if one was specified: ` + "`" + `alert` + "`" + `, ` + "`" + `deny` + "`" + `, or ` + "`" + `none` + "`" + `. If the action is none, the attack group is inactive in the security policy.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ` + "`" + `action` + "`" + ` and ` + "`" + `group` + "`" + ` name for each attack group.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `The attack group information in JSON format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_attack_group_condition_exception",
			Category:         "Application Security",
			ShortDescription: `Attack Group Condition Exception`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"attack",
				"group",
				"condition",
				"exception",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "attack_group",
					Description: `(Required) The attack group to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the condition and exception information.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `The condition and exception information in JSON format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the condition and exception information.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `The condition and exception information in JSON format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_bypass_network_lists",
			Category:         "Application Security",
			ShortDescription: `BypassNetworkLists`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"bypass",
				"network",
				"lists",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The configuration ID to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bypass_network_list",
					Description: `A list of strings containing the network list IDs.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of information about the bypass network lists.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the bypass network list information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bypass_network_list",
					Description: `A list of strings containing the network list IDs.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of information about the bypass network lists.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the bypass network list information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_configuration",
			Category:         "Application Security",
			ShortDescription: `Configuration`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"configuration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of a specific security configuration. If not supplied, information about all security configurations is returned. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `The ID of the specified security configuration. Returned only if ` + "`" + `name` + "`" + ` was specified.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the following information about all available security configurations: config_id, name, latest_version, version_active_in_staging, and version_active_in_production.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The last version of the specified security configuration created. Returned only if ` + "`" + `name` + "`" + ` was specified.`,
				},
				resource.Attribute{
					Name:        "staging_version",
					Description: `The version of the specified security configuration currently active in staging. Returned only if ` + "`" + `name` + "`" + ` was specified.`,
				},
				resource.Attribute{
					Name:        "production_version",
					Description: `The version of the specified security configuration currently active in production. Returned only if ` + "`" + `name` + "`" + ` was specified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `The ID of the specified security configuration. Returned only if ` + "`" + `name` + "`" + ` was specified.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the following information about all available security configurations: config_id, name, latest_version, version_active_in_staging, and version_active_in_production.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The last version of the specified security configuration created. Returned only if ` + "`" + `name` + "`" + ` was specified.`,
				},
				resource.Attribute{
					Name:        "staging_version",
					Description: `The version of the specified security configuration currently active in staging. Returned only if ` + "`" + `name` + "`" + ` was specified.`,
				},
				resource.Attribute{
					Name:        "production_version",
					Description: `The version of the specified security configuration currently active in production. Returned only if ` + "`" + `name` + "`" + ` was specified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_configuration_version",
			Category:         "Application Security",
			ShortDescription: `ConfigurationVersion`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"configuration",
				"version",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The version number of the security configuration to use. If not supplied, information about all versions of the specified security configuration is returned. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The last version of the security configuration created.`,
				},
				resource.Attribute{
					Name:        "staging_status",
					Description: `The status of the specified version in staging: "Active", "Inactive", or "Deactivated". Returned only if ` + "`" + `version` + "`" + ` was specified.`,
				},
				resource.Attribute{
					Name:        "production_status",
					Description: `The status of the specified version in production: "Active", "Inactive", or "Deactivated". Returned only if ` + "`" + `version` + "`" + ` was specified.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the following information about all versions of the security configuration: version number, staging status, and production status.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "latest_version",
					Description: `The last version of the security configuration created.`,
				},
				resource.Attribute{
					Name:        "staging_status",
					Description: `The status of the specified version in staging: "Active", "Inactive", or "Deactivated". Returned only if ` + "`" + `version` + "`" + ` was specified.`,
				},
				resource.Attribute{
					Name:        "production_status",
					Description: `The status of the specified version in production: "Active", "Inactive", or "Deactivated". Returned only if ` + "`" + `version` + "`" + ` was specified.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the following information about all versions of the security configuration: version number, staging status, and production status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_contracts_groups",
			Category:         "Application Security",
			ShortDescription: `ContractsGroups`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"contracts",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "contractid",
					Description: `(Optional) The ID of a contract for which to retrieve information.`,
				},
				resource.Attribute{
					Name:        "groupid",
					Description: `(Optional) The ID of a group for which to retrieve information. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the contract and group information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the contract and group information.`,
				},
				resource.Attribute{
					Name:        "default_contractid",
					Description: `The default contract ID for the specified contract and group.`,
				},
				resource.Attribute{
					Name:        "default_groupid",
					Description: `The default group ID for the specified contract and group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "contractid",
					Description: `(Optional) The ID of a contract for which to retrieve information.`,
				},
				resource.Attribute{
					Name:        "groupid",
					Description: `(Optional) The ID of a group for which to retrieve information. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the contract and group information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the contract and group information.`,
				},
				resource.Attribute{
					Name:        "default_contractid",
					Description: `The default contract ID for the specified contract and group.`,
				},
				resource.Attribute{
					Name:        "default_groupid",
					Description: `The default group ID for the specified contract and group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_custom_deny",
			Category:         "Application Security",
			ShortDescription: `CustomDeny`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"custom",
				"deny",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The configuration ID to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration to use.`,
				},
				resource.Attribute{
					Name:        "custom_deny_id",
					Description: `(Optional) The ID of a specific custom deny action. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the custom deny action information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the custom deny action information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the custom deny action information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the custom deny action information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_custom_rule_actions",
			Category:         "Application Security",
			ShortDescription: `Custom Rule Actions`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"custom",
				"rule",
				"actions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use`,
				},
				resource.Attribute{
					Name:        "custom_rule_id",
					Description: `(Optional) A specific custom rule for which to retrieve information. If not supplied, information about all custom rules will be returned. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID, name, and action of all custom rules, or of the specific custom rule, associated with the specified security configuration, version and security policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID, name, and action of all custom rules, or of the specific custom rule, associated with the specified security configuration, version and security policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_custom_rules",
			Category:         "Application Security",
			ShortDescription: `CustomRules`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"custom",
				"rules",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "custom_rule_id",
					Description: `(Optional) The ID of a specific custom rule to use. If not supplied, information about all custom rules associated with the given security configuration will be returned. ## Attributes Reference In addition to the argument above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and name of the custom rule(s).`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted display of the custom rule information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and name of the custom rule(s).`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted display of the custom rule information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_eval_hostnames",
			Category:         "Application Security",
			ShortDescription: `EvalHostnames`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"eval",
				"hostnames",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `A list of the evaluation hostnames.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the evaluation hostnames.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the evaluation hostnames.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hostnames",
					Description: `A list of the evaluation hostnames.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the evaluation hostnames.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the evaluation hostnames.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_eval_rule_actions",
			Category:         "Application Security",
			ShortDescription: `Eval Rule Actions`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"eval",
				"rule",
				"actions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) The ID of a specific rule. If not supplied, information about all eval rules will be returned. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the ID and action for all rules in the security policy.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted display of the ID and action for all rules in the security policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action configured for the given rule if a ` + "`" + `rule_id` + "`" + ` was specified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the ID and action for all rules in the security policy.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted display of the ID and action for all rules in the security policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action configured for the given rule if a ` + "`" + `rule_id` + "`" + ` was specified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_eval_rule_condition_exception",
			Category:         "Application Security",
			ShortDescription: `KRS Eval Rule Condition-Exception`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"eval",
				"rule",
				"condition",
				"exception",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) The ID of the rule to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing boolean values indicating whether conditions and exceptions are present.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the condition and exception information for the specified rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing boolean values indicating whether conditions and exceptions are present.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the condition and exception information for the specified rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_export_configuration",
			Category:         "Application Security",
			ShortDescription: `ExportConfiguration`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"export",
				"configuration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "search",
					Description: `(Optional) A bracket-delimited list of quoted strings specifying the types of information to be retrieved and made available for display in the ` + "`" + `output_text` + "`" + ` format. The following types are available:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `The complete set of information about the specified security configuration version, in JSON format. This includes the types available for the ` + "`" + `search` + "`" + ` parameter, plus several additional fields such as createDate and createdBy.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the types of data specified in the ` + "`" + `search` + "`" + ` parameter. Included only if the ` + "`" + `search` + "`" + ` parameter specifies at least one type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `The complete set of information about the specified security configuration version, in JSON format. This includes the types available for the ` + "`" + `search` + "`" + ` parameter, plus several additional fields such as createDate and createdBy.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the types of data specified in the ` + "`" + `search` + "`" + ` parameter. Included only if the ` + "`" + `search` + "`" + ` parameter specifies at least one type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_failover_hostnames",
			Category:         "Application Security",
			ShortDescription: `FailoverHostnames`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"failover",
				"hostnames",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use. ## Attributes Reference In addition to the argument above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `A list of the failover hostnames.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the failover hostnames.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the failover hostnames.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hostnames",
					Description: `A list of the failover hostnames.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the failover hostnames.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the failover hostnames.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_hostname_coverage",
			Category:         "Application Security",
			ShortDescription: `HostnameCoverage`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"hostname",
				"coverage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the hostname coverage information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the hostname coverage information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the hostname coverage information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the hostname coverage information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_hostname_coverage_match_targets",
			Category:         "Application Security",
			ShortDescription: `ApiHostnameCoverageMatchTargets`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"hostname",
				"coverage",
				"match",
				"targets",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The hostname for which to retrieve information. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the coverage information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the coverage information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the coverage information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the coverage information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_hostname_coverage_overlapping",
			Category:         "Application Security",
			ShortDescription: `ApiHostnameCoverageOverlapping`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"hostname",
				"coverage",
				"overlapping",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The hostname for which to retrieve information. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the overlap information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the overlap information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the overlap information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the overlap information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_ip_geo",
			Category:         "Application Security",
			ShortDescription: `IP/Geo`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"ip",
				"geo",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Optional) The ID of the security policy to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The mode used for IP/Geo firewall blocking: ` + "`" + `block` + "`" + ` to block specific IPs, geographies or network lists, or ` + "`" + `allow` + "`" + ` to allow specific IPs or geographies to be let through while blocking the rest.`,
				},
				resource.Attribute{
					Name:        "geo_network_lists",
					Description: `The network lists to be blocked or allowed geographically.`,
				},
				resource.Attribute{
					Name:        "ip_network_lists",
					Description: `The network lists to be blocked or allowd by IP address.`,
				},
				resource.Attribute{
					Name:        "exception_ip_network_lists",
					Description: `The network lists to be allowed regardless of ` + "`" + `mode` + "`" + `, ` + "`" + `geo_network_lists` + "`" + `, and ` + "`" + `ip_network_lists` + "`" + ` parameters.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the IP/Geo firewall settings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mode",
					Description: `The mode used for IP/Geo firewall blocking: ` + "`" + `block` + "`" + ` to block specific IPs, geographies or network lists, or ` + "`" + `allow` + "`" + ` to allow specific IPs or geographies to be let through while blocking the rest.`,
				},
				resource.Attribute{
					Name:        "geo_network_lists",
					Description: `The network lists to be blocked or allowed geographically.`,
				},
				resource.Attribute{
					Name:        "ip_network_lists",
					Description: `The network lists to be blocked or allowd by IP address.`,
				},
				resource.Attribute{
					Name:        "exception_ip_network_lists",
					Description: `The network lists to be allowed regardless of ` + "`" + `mode` + "`" + `, ` + "`" + `geo_network_lists` + "`" + `, and ` + "`" + `ip_network_lists` + "`" + ` parameters.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the IP/Geo firewall settings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_match_targets",
			Category:         "Application Security",
			ShortDescription: `MatchTargets`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"match",
				"targets",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "match_target_id",
					Description: `(Optional) The ID of the match target to use. If not supplied, information about all match targets is returned. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and Policy ID of all match targets associated with the specified security configuration and version, or of the specific match target if ` + "`" + `match_target_id` + "`" + ` was supplied.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and Policy ID of all match targets associated with the specified security configuration and version, or of the specific match target if ` + "`" + `match_target_id` + "`" + ` was supplied.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_penalty_box",
			Category:         "Application Security",
			ShortDescription: `Penalty Box`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"penalty",
				"box",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action for the penalty box: ` + "`" + `alert` + "`" + `, ` + "`" + `deny` + "`" + `, or ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether penalty box protection is enabled.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the ` + "`" + `action` + "`" + ` and ` + "`" + `enabled` + "`" + ` information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `The action for the penalty box: ` + "`" + `alert` + "`" + `, ` + "`" + `deny` + "`" + `, or ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether penalty box protection is enabled.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the ` + "`" + `action` + "`" + ` and ` + "`" + `enabled` + "`" + ` information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_rate_policies",
			Category:         "Application Security",
			ShortDescription: `Rate Policies`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"rate",
				"policies",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "rate_policy_id",
					Description: `(Optional) The ID of the rate policy to use. If this parameter is not supplied, information about all rate policies will be returned. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and name of all rate policies associated with the specified security configuration version.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the rate policy information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and name of all rate policies associated with the specified security configuration version.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the rate policy information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_rate_policy_actions",
			Category:         "Application Security",
			ShortDescription: `Rate Policy Actions`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"rate",
				"policy",
				"actions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "rate_policy_id",
					Description: `(Optional) The ID of the rate policy to use. If not supplied, information about all rate policies will be returned. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID IPv4Action and IPv6Action of the indicated security policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID IPv4Action and IPv6Action of the indicated security policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_reputation_profile_actions",
			Category:         "Application Security",
			ShortDescription: `Reputation Profile Actions`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"reputation",
				"profile",
				"actions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) THe ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "reputation_profile_id",
					Description: `(Optional) The ID of a given reputation profile. If not supplied, information about all reputation profiles is returned. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action that the specified reputation profile or profiles take when triggered.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted display of the specified reputation profile action information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the specified reputation profile action information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `The action that the specified reputation profile or profiles take when triggered.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted display of the specified reputation profile action information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the specified reputation profile action information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_reputation_profile_analysis",
			Category:         "Application Security",
			ShortDescription: `Reputation Profile Analysis`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"reputation",
				"profile",
				"analysis",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The configuration ID to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the reputation analysis settings.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the reputation analysis settings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the reputation analysis settings.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the reputation analysis settings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_reputation_profiles",
			Category:         "Application Security",
			ShortDescription: `Reputation Profiles`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"reputation",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "reputation_profile_id",
					Description: `(Optional) The ID of a given reputation profile. If not supplied, information about all reputation profiles is returned. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the details about the indicated reputation profile or profiles.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted display of the details about the indicated reputation profile or profiles.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the details about the indicated reputation profile or profiles.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted display of the details about the indicated reputation profile or profiles.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_rule_actions",
			Category:         "Application Security",
			ShortDescription: `KRS Rule Actions`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"rule",
				"actions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) The ID of a specific rule. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the ID and action for all rules in the security policy.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted display of the ID and action for all rules in the security policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action configured for the given rule if a ` + "`" + `rule_id` + "`" + ` was specified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the ID and action for all rules in the security policy.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted display of the ID and action for all rules in the security policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action configured for the given rule if a ` + "`" + `rule_id` + "`" + ` was specified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_rule_condition_exception",
			Category:         "Application Security",
			ShortDescription: `KRS Rule Condition-Exception`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"rule",
				"condition",
				"exception",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) The ID of the rule to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing boolean values indicating whether conditions and exceptions are present`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the condition and exception information for the specified rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing boolean values indicating whether conditions and exceptions are present`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the condition and exception information for the specified rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_rule_upgrade_details",
			Category:         "Application Security",
			ShortDescription: `Rule Upgrade Details`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"rule",
				"upgrade",
				"details",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing changes (additions and deletions) to the rules for the specified security policy.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the changes (additions and deletions) to the rules for the specified security policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing changes (additions and deletions) to the rules for the specified security policy.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the changes (additions and deletions) to the rules for the specified security policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_security_policy",
			Category:         "Application Security",
			ShortDescription: `SecurityPolicy`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "policy_list",
					Description: `A list of the IDs of all security policies.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and name of all security policies.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The ID of the security policy. Included only if ` + "`" + `name` + "`" + ` was specified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_list",
					Description: `A list of the IDs of all security policies.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and name of all security policies.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The ID of the security policy. Included only if ` + "`" + `name` + "`" + ` was specified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_security_policy_protections",
			Category:         "Application Security",
			ShortDescription: `Security Policy Protections`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"policy",
				"protections",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "apply_application_layer_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether application layer controls are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_network_layer_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether network layer controls are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_rate_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether rate controls are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_reputation_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether reputation controls are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_botman_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether botman controls are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_api_constraints",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether API constraints are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_slow_post_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether slow post controls are in effect.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `a JSON-formatted list showing the status of the protection settings`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `a tabular display showing the status of the protection settings`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "apply_application_layer_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether application layer controls are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_network_layer_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether network layer controls are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_rate_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether rate controls are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_reputation_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether reputation controls are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_botman_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether botman controls are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_api_constraints",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether API constraints are in effect.`,
				},
				resource.Attribute{
					Name:        "apply_slow_post_controls",
					Description: `` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, indicating whether slow post controls are in effect.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `a JSON-formatted list showing the status of the protection settings`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `a tabular display showing the status of the protection settings`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_selectable_hostnames",
			Category:         "Application Security",
			ShortDescription: `SelectableHostnames`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"selectable",
				"hostnames",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Optional) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "contractid",
					Description: `(Optional) The ID of the contract to use.`,
				},
				resource.Attribute{
					Name:        "groupid",
					Description: `(Optional) The ID of the group to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `The list of selectable hostnames.`,
				},
				resource.Attribute{
					Name:        "hostnames_json",
					Description: `The list of selectable hostnames in json format.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the selectable hostnames showing the name and config_id of the security configuration under which the host is protected in production, or '-' if the host is not protected in production.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hostnames",
					Description: `The list of selectable hostnames.`,
				},
				resource.Attribute{
					Name:        "hostnames_json",
					Description: `The list of selectable hostnames in json format.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the selectable hostnames showing the name and config_id of the security configuration under which the host is protected in production, or '-' if the host is not protected in production.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_selected_hostnames",
			Category:         "Application Security",
			ShortDescription: `SelectedHostnames`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"selected",
				"hostnames",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `The list of selected hostnames.`,
				},
				resource.Attribute{
					Name:        "hostnames_json",
					Description: `The list of selected hostnames in JSON format.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the selected hostnames.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hostnames",
					Description: `The list of selected hostnames.`,
				},
				resource.Attribute{
					Name:        "hostnames_json",
					Description: `The list of selected hostnames in JSON format.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the selected hostnames.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_siem_definitions",
			Category:         "Application Security",
			ShortDescription: `SiemDefinitions`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"siem",
				"definitions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the SIEM version information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and name of each SIEM version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the SIEM version information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the ID and name of each SIEM version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_siem_settings",
			Category:         "Application Security",
			ShortDescription: `SiemSettijgs`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"siem",
				"settings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the SIEM setting information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the SIEM setting information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the SIEM setting information.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the SIEM setting information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_slow_post",
			Category:         "Application Security",
			ShortDescription: `Slow Post`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"slow",
				"post",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display including the following columns:`,
				},
				resource.Attribute{
					Name:        "ACTION",
					Description: `The action that the rule should trigger (either ` + "`" + `alert` + "`" + ` or ` + "`" + `abort` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "SLOW_RATE_THRESHOLD RATE",
					Description: `The average rate in bytes per second over the period specified by ` + "`" + `period` + "`" + ` before the specified ` + "`" + `action` + "`" + ` is triggered.`,
				},
				resource.Attribute{
					Name:        "SLOW_RATE_THRESHOLD PERIOD",
					Description: `The length in seconds of the period during which the server should accept a request before determining whether a POST request is too slow.`,
				},
				resource.Attribute{
					Name:        "DURATION_THRESHOLD TIMEOUT",
					Description: `The time in seconds before the first eight kilobytes of the POST body must be received to avoid triggering the specified ` + "`" + `action` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display including the following columns:`,
				},
				resource.Attribute{
					Name:        "ACTION",
					Description: `The action that the rule should trigger (either ` + "`" + `alert` + "`" + ` or ` + "`" + `abort` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "SLOW_RATE_THRESHOLD RATE",
					Description: `The average rate in bytes per second over the period specified by ` + "`" + `period` + "`" + ` before the specified ` + "`" + `action` + "`" + ` is triggered.`,
				},
				resource.Attribute{
					Name:        "SLOW_RATE_THRESHOLD PERIOD",
					Description: `The length in seconds of the period during which the server should accept a request before determining whether a POST request is too slow.`,
				},
				resource.Attribute{
					Name:        "DURATION_THRESHOLD TIMEOUT",
					Description: `The time in seconds before the first eight kilobytes of the POST body must be received to avoid triggering the specified ` + "`" + `action` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_version_notes",
			Category:         "Application Security",
			ShortDescription: `VersionNotes`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"version",
				"notes",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The configuration ID to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list showing the version notes.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the version notes.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list showing the version notes.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the version notes.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_waf_mode",
			Category:         "Application Security",
			ShortDescription: `WAF Mode`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"waf",
				"mode",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security policy to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The security policy mode, either ` + "`" + `KRS` + "`" + ` (update manually) or ` + "`" + `AAG` + "`" + ` (update automatically),`,
				},
				resource.Attribute{
					Name:        "current_ruleset",
					Description: `The current rule set version and the ISO 8601 date the rule set version was introduced; this date acts like a version number.`,
				},
				resource.Attribute{
					Name:        "eval_status",
					Description: `Whether the evaluation mode is enabled or disabled."`,
				},
				resource.Attribute{
					Name:        "eval_ruleset",
					Description: `The evaluation rule set version and the ISO 8601 date the evaluation starts.`,
				},
				resource.Attribute{
					Name:        "eval_expiration_date",
					Description: `The ISO 8601 time stamp when the evaluation is expiring. This value only appears when ` + "`" + `eval` + "`" + ` is set to "enabled".`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the mode information.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the mode information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mode",
					Description: `The security policy mode, either ` + "`" + `KRS` + "`" + ` (update manually) or ` + "`" + `AAG` + "`" + ` (update automatically),`,
				},
				resource.Attribute{
					Name:        "current_ruleset",
					Description: `The current rule set version and the ISO 8601 date the rule set version was introduced; this date acts like a version number.`,
				},
				resource.Attribute{
					Name:        "eval_status",
					Description: `Whether the evaluation mode is enabled or disabled."`,
				},
				resource.Attribute{
					Name:        "eval_ruleset",
					Description: `The evaluation rule set version and the ISO 8601 date the evaluation starts.`,
				},
				resource.Attribute{
					Name:        "eval_expiration_date",
					Description: `The ISO 8601 time stamp when the evaluation is expiring. This value only appears when ` + "`" + `eval` + "`" + ` is set to "enabled".`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the mode information.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `A JSON-formatted list of the mode information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_authorities_set",
			Category:         "DNS",
			ShortDescription: `DNS Authorities Set`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"authorities",
				"set",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_contract",
			Category:         "Common",
			ShortDescription: `Contract`,
			Description:      ``,
			Keywords: []string{
				"common",
				"contract",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_contracts",
			Category:         "Common",
			ShortDescription: `Property contracts`,
			Description:      ``,
			Keywords: []string{
				"common",
				"contracts",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_cp_code",
			Category:         "Provisioning",
			ShortDescription: `CP Code`,
			Description:      ``,
			Keywords: []string{
				"provisioning",
				"cp",
				"code",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_group",
			Category:         "Common",
			ShortDescription: `Group`,
			Description:      ``,
			Keywords: []string{
				"common",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_groups",
			Category:         "Common",
			ShortDescription: `Property groups`,
			Description:      ``,
			Keywords: []string{
				"common",
				"groups",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_default_datacenter",
			Category:         "Global Traffic Management",
			ShortDescription: `Default data center`,
			Description:      ``,
			Keywords: []string{
				"global",
				"traffic",
				"management",
				"gtm",
				"default",
				"datacenter",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_iam_contact_types",
			Category:         "IAM",
			ShortDescription: `IAM Contact Types`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"contact",
				"types",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_iam_countries",
			Category:         "IAM",
			ShortDescription: `IAM Countries`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"countries",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_iam_groups",
			Category:         "IAM",
			ShortDescription: `IAM Groups`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"groups",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_iam_roles",
			Category:         "IAM",
			ShortDescription: `IAM Roles`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"roles",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_iam_states",
			Category:         "IAM",
			ShortDescription: `IAM States`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"states",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_iam_supported_langs",
			Category:         "IAM",
			ShortDescription: `IAM Supported Languages`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"supported",
				"langs",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_iam_timeout_policies",
			Category:         "IAM",
			ShortDescription: `IAM Timeout Policies`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"timeout",
				"policies",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_iam_timezones",
			Category:         "IAM",
			ShortDescription: `IAM Timeout Policies`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"timezones",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_properties",
			Category:         "Provisioning",
			ShortDescription: `Properties`,
			Description:      ``,
			Keywords: []string{
				"provisioning",
				"properties",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property_products",
			Category:         "Provisioning",
			ShortDescription: `Property products`,
			Description:      ``,
			Keywords: []string{
				"provisioning",
				"property",
				"products",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property_rule_formats",
			Category:         "Provisioning",
			ShortDescription: `Properties rule formats`,
			Description:      ``,
			Keywords: []string{
				"provisioning",
				"property",
				"rule",
				"formats",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property_rules",
			Category:         "Provisioning",
			ShortDescription: `Property rule tree`,
			Description:      ``,
			Keywords: []string{
				"provisioning",
				"property",
				"rules",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property_rules_template",
			Category:         "Provisioning",
			ShortDescription: `Property Rules Template`,
			Description:      ``,
			Keywords: []string{
				"provisioning",
				"property",
				"rules",
				"template",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"akamai_appsec_advanced_settings_logging":        0,
		"akamai_appsec_advanced_settings_prefetch":       1,
		"akamai_appsec_api_endpoints":                    2,
		"akamai_appsec_api_request_constraints":          3,
		"akamai_appsec_attack_group_actions":             4,
		"akamai_appsec_attack_group_condition_exception": 5,
		"akamai_appsec_bypass_network_lists":             6,
		"akamai_appsec_configuration":                    7,
		"akamai_appsec_configuration_version":            8,
		"akamai_appsec_contracts_groups":                 9,
		"akamai_appsec_custom_deny":                      10,
		"akamai_appsec_custom_rule_actions":              11,
		"akamai_appsec_custom_rules":                     12,
		"akamai_appsec_eval_hostnames":                   13,
		"akamai_appsec_eval_rule_actions":                14,
		"akamai_appsec_eval_rule_condition_exception":    15,
		"akamai_appsec_export_configuration":             16,
		"akamai_appsec_failover_hostnames":               17,
		"akamai_appsec_hostname_coverage":                18,
		"akamai_appsec_hostname_coverage_match_targets":  19,
		"akamai_appsec_hostname_coverage_overlapping":    20,
		"akamai_appsec_ip_geo":                           21,
		"akamai_appsec_match_targets":                    22,
		"akamai_appsec_penalty_box":                      23,
		"akamai_appsec_rate_policies":                    24,
		"akamai_appsec_rate_policy_actions":              25,
		"akamai_appsec_reputation_profile_actions":       26,
		"akamai_appsec_reputation_profile_analysis":      27,
		"akamai_appsec_reputation_profiles":              28,
		"akamai_appsec_rule_actions":                     29,
		"akamai_appsec_rule_condition_exception":         30,
		"akamai_appsec_rule_upgrade_details":             31,
		"akamai_appsec_security_policy":                  32,
		"akamai_appsec_security_policy_protections":      33,
		"akamai_appsec_selectable_hostnames":             34,
		"akamai_appsec_selected_hostnames":               35,
		"akamai_appsec_siem_definitions":                 36,
		"akamai_appsec_siem_settings":                    37,
		"akamai_appsec_slow_post":                        38,
		"akamai_appsec_version_notes":                    39,
		"akamai_appsec_waf_mode":                         40,
		"akamai_authorities_set":                         41,
		"akamai_contract":                                42,
		"akamai_contracts":                               43,
		"akamai_cp_code":                                 44,
		"akamai_group":                                   45,
		"akamai_groups":                                  46,
		"akamai_gtm_default_datacenter":                  47,
		"akamai_iam_contact_types":                       48,
		"akamai_iam_countries":                           49,
		"akamai_iam_groups":                              50,
		"akamai_iam_roles":                               51,
		"akamai_iam_states":                              52,
		"akamai_iam_supported_langs":                     53,
		"akamai_iam_timeout_policies":                    54,
		"akamai_iam_timezones":                           55,
		"akamai_properties":                              56,
		"akamai_property_products":                       57,
		"akamai_property_rule_formats":                   58,
		"akamai_property_rules":                          59,
		"akamai_property_rules_template":                 60,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
