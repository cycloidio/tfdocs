package akamai

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_activations",
			Category:         "Application Security",
			ShortDescription: `Activations`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"activations",
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
					Name:        "notification_emails",
					Description: `(Required) A bracketed, comma-separated list of email addresses that will be notified when the operation is complete.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The network in which the security configuration should be activated. If supplied, must be either STAGING or PRODUCTION. If not supplied, STAGING will be assumed.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `An optional text note describing this operation.`,
				},
				resource.Attribute{
					Name:        "activate",
					Description: `A boolean indicating whether to activate the specified configuration version. If not supplied, True is assumed. ## Attribute Reference In addition to the arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the operation. The following values are may be returned:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of the operation. The following values are may be returned:`,
				},
			},
		},
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
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Required) The logging settings to apply ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#puthttpheaderloggingforaconfiguration)).`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Optional) The ID of a specific security policy to which the logging settings should be applied. If not supplied, the indicated settings will be applied to all policies within the configuration. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "enable_app_layer",
					Description: `(Required) Whether to enable prefetch requests.`,
				},
				resource.Attribute{
					Name:        "all_extensions",
					Description: `(Required) Whether to enable prefetch requests for all extensions.`,
				},
				resource.Attribute{
					Name:        "enable_rate_controls",
					Description: `(Required) Whether to enable prefetch requests for rate controls.`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Required) The specific extensions for which to enable prefetch requests. If ` + "`" + `all_extensions` + "`" + ` is True, ` + "`" + `extensions` + "`" + ` must be an empty list. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Name:        "api_endpoint_id",
					Description: `(Optional) The ID of the API endpoint to use. If not supplied, the request constraint action will be updated for all APIs.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action to assign to API request constraints: either ` + "`" + `alert` + "`" + `, ` + "`" + `deny` + "`" + `, or ` + "`" + `none` + "`" + `. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the API request constraint information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display of the API request constraint information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_attack_group_action",
			Category:         "Application Security",
			ShortDescription: `Attack Group Action`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"attack",
				"group",
				"action",
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
					Description: `(Required) The ID of the attack group to use.`,
				},
				resource.Attribute{
					Name:        "attack_group_action",
					Description: `(Required) The action to be taken: ` + "`" + `alert` + "`" + ` to record the trigger of the event, ` + "`" + `deny` + "`" + ` to block the request, ` + "`" + `deny_custom_{custom_deny_id}` + "`" + ` to execute a custom deny action, or ` + "`" + `none` + "`" + ` to take no action. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_attack_group_condition_exception",
			Category:         "Application Security",
			ShortDescription: `Attack Group Condition & Exception`,
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
					Description: `The attack group to use.`,
				},
				resource.Attribute{
					Name:        "condition_exception",
					Description: `(Required) The name of a file containing a JSON-formatted description of the conditions and exceptions to use ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception)). ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required) The version number of the configuration to use.`,
				},
				resource.Attribute{
					Name:        "bypass_network_list",
					Description: `(Required) A list containing the IDs of the network lists to use. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the updated list of network list IDs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the updated list of network list IDs.`,
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
					Name:        "description",
					Description: `(Required) A description of the configuration.`,
				},
				resource.Attribute{
					Name:        "contract_id",
					Description: `(Required) The contract ID of the configuration.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The group ID of the configuration>`,
				},
				resource.Attribute{
					Name:        "host_names",
					Description: `(Required) The list of hostnames protected by this security configuration. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The latest version of the security configuration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The latest version of the security configuration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_configuration_clone",
			Category:         "Application Security",
			ShortDescription: `ConfigurationClone`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"configuration",
				"clone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to be applied to the new configuration.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description of the new configuration.`,
				},
				resource.Attribute{
					Name:        "create_from_config_id",
					Description: `(Required) The ID of the configuration to be cloned.`,
				},
				resource.Attribute{
					Name:        "create_from_version",
					Description: `(Required) The version number of the configuration to be cloned.`,
				},
				resource.Attribute{
					Name:        "contract_id",
					Description: `(Required) The contract id to use.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The group id to use.`,
				},
				resource.Attribute{
					Name:        "host_names",
					Description: `(Required) The hostnames to be protected under the new configuration. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `The ID of the newly created configuration.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version number of the newly created configuration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `The ID of the newly created configuration.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version number of the newly created configuration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_configuration_rename",
			Category:         "Application Security",
			ShortDescription: `ConfigurationRename`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"configuration",
				"rename",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to be renamed.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The new name to be given to the configuration.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description to be applied to the configuration. ## Attribute Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_configuration_version_clone",
			Category:         "Application Security",
			ShortDescription: `ConfigurationVersionClone`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"configuration",
				"version",
				"clone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "create_from_version",
					Description: `(Required) The version number of the security configuration to clone.`,
				},
				resource.Attribute{
					Name:        "rule_update",
					Description: `A boolean indicating whether to update the rules of the new version. If not supplied, False is assumed. ## Attribute Reference In addition to the arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The number of the cloned version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `The number of the cloned version.`,
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
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "custom_deny",
					Description: `(Required) The JSON-formatted definition of the custom deny action ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#63df3de3)). ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_custom_rule",
			Category:         "Application Security",
			ShortDescription: `Custom Rule`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"custom",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "custom_rule",
					Description: `(Required) The name of a JSON file containing a custom rule definition ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postcustomrules)). ## Attribute Reference In addition to the arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "custom_rule_id",
					Description: `The ID of the custom rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "custom_rule_id",
					Description: `The ID of the custom rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_custom_rule_action",
			Category:         "Application Security",
			ShortDescription: `Custom Rule Action`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"custom",
				"rule",
				"action",
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
					Description: `(Required) The security policy to use.`,
				},
				resource.Attribute{
					Name:        "custom_rule_id",
					Description: `(Required) The custom rule for which to apply the action.`,
				},
				resource.Attribute{
					Name:        "custom_rule_action",
					Description: `(Required) The action to take when the custom rule is invoked: ` + "`" + `alert` + "`" + ` to record the trigger event, ` + "`" + `deny` + "`" + ` to block the request, ` + "`" + `deny_custom_{custom_deny_id}` + "`" + ` to execute a custom deny action, or ` + "`" + `none` + "`" + ` to take no action. ## Attribute Reference In addition to the arguments above, the following attribute is exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_eval",
			Category:         "Application Security",
			ShortDescription: `Evaluation`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"eval",
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
					Name:        "eval_operation",
					Description: `(Required) The operation to perform: START, STOP, RESTART, UPDATE, or COMPLETE. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "evaluating_ruleset",
					Description: `The set of rules being evaluated.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date on which the evaluation period ends.`,
				},
				resource.Attribute{
					Name:        "current_ruleset",
					Description: `The set of rules currently in effect.`,
				},
				resource.Attribute{
					Name:        "eval_status",
					Description: `Either ` + "`" + `enabled` + "`" + ` if an evaluation is currently in progress (that is, if the ` + "`" + `eval_operation` + "`" + ` parameter was ` + "`" + `START` + "`" + `, ` + "`" + `RESTART` + "`" + `, or ` + "`" + `COMPLETE` + "`" + `) or ` + "`" + `disabled` + "`" + ` otherwise (that is, if the ` + "`" + `eval_operation` + "`" + ` parameter was ` + "`" + `STOP` + "`" + ` or ` + "`" + `UPDATE` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "evaluating_ruleset",
					Description: `The set of rules being evaluated.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date on which the evaluation period ends.`,
				},
				resource.Attribute{
					Name:        "current_ruleset",
					Description: `The set of rules currently in effect.`,
				},
				resource.Attribute{
					Name:        "eval_status",
					Description: `Either ` + "`" + `enabled` + "`" + ` if an evaluation is currently in progress (that is, if the ` + "`" + `eval_operation` + "`" + ` parameter was ` + "`" + `START` + "`" + `, ` + "`" + `RESTART` + "`" + `, or ` + "`" + `COMPLETE` + "`" + `) or ` + "`" + `disabled` + "`" + ` otherwise (that is, if the ` + "`" + `eval_operation` + "`" + ` parameter was ` + "`" + `STOP` + "`" + ` or ` + "`" + `UPDATE` + "`" + `).`,
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
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `(Required) A list of evaluation hostnames to be used for the specified configuration version. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_eval_protect_host",
			Category:         "Application Security",
			ShortDescription: `EvalProtectHost`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"eval",
				"protect",
				"host",
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
					Name:        "hostnames",
					Description: `(Required) The evaluation hostnames to be moved to active protection. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_eval_rule_action",
			Category:         "Application Security",
			ShortDescription: `Eval Rule Action`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"eval",
				"rule",
				"action",
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
					Description: `(Required) The ID of the rule being evaluated.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action to be taken: ` + "`" + `alert` + "`" + ` to record the trigger of the event, ` + "`" + `deny` + "`" + ` to block the request, ` + "`" + `deny_custom_{custom_deny_id}` + "`" + ` to execute a custom deny action, or ` + "`" + `none` + "`" + ` to take no action. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_eval_rule_condition_exception",
			Category:         "Application Security",
			ShortDescription: `Eval Rule Condition / Exception`,
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
					Description: `(Required) The ID of the eval rule to use.`,
				},
				resource.Attribute{
					Name:        "condition_exception",
					Description: `(Required) The name of a file containing a JSON-formatted description of the conditions and exceptions to use ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putevalconditionsexceptions)) ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_ip_geo",
			Category:         "Application Security",
			ShortDescription: `IP/Geo Firewall`,
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
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The mode to use for IP/Geo firewall blocking: ` + "`" + `block` + "`" + ` to block specific IPs, geographies or network lists, or ` + "`" + `allow` + "`" + ` to allow specific IPs or geographies to be let through while blocking the rest.`,
				},
				resource.Attribute{
					Name:        "geo_network_lists",
					Description: `(Optional) The network lists to be blocked or allowed geographically.`,
				},
				resource.Attribute{
					Name:        "ip_network_lists",
					Description: `(Optional) The network lists to be blocked or allowd by IP address.`,
				},
				resource.Attribute{
					Name:        "exception_ip_network_lists",
					Description: `(Required) The network lists to be allowed regardless of ` + "`" + `mode` + "`" + `, ` + "`" + `geo_network_lists` + "`" + `, and ` + "`" + `ip_network_lists` + "`" + ` parameters. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_match_target",
			Category:         "Application Security",
			ShortDescription: `MatchTarget`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"match",
				"target",
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
					Name:        "match_target",
					Description: `(Required) The name of a JSON file containing one or more match target definitions ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postmatchtargets)). ## Attribute Reference In addition to the arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "match_target_id",
					Description: `The ID of the match target.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "match_target_id",
					Description: `The ID of the match target.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_match_target_sequence",
			Category:         "Application Security",
			ShortDescription: `MatchTargetSequence`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"match",
				"target",
				"sequence",
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
					Name:        "match_target_sequence",
					Description: `(Required) The name of a JSON file containing the sequence of all match targets defined for the specified security configuration version ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsequence)). ## Attribute Reference In addition to the arguments above, the following attribute is exported:`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "penalty_box_protection",
					Description: `(Required) A boolean value indicating whether to enable penalty box protection.`,
				},
				resource.Attribute{
					Name:        "penalty_box_action",
					Description: `(Required) The action to take when penalty box protection is triggered: ` + "`" + `alert` + "`" + ` to record the trigger event, ` + "`" + `deny` + "`" + ` to block the request, ` + "`" + `deny_custom_{custom_deny_id}` + "`" + ` to execute a custom deny action, or ` + "`" + `none` + "`" + ` to take no action. Ignored if ` + "`" + `penalty_box_protection` + "`" + ` is set to ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_rate_policy",
			Category:         "Application Security",
			ShortDescription: `Rate Policy`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"rate",
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
					Name:        "rate_policy",
					Description: `(Required) The name of a file containing a JSON-formatted rate policy definition ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#57c65cbd)).`,
				},
				resource.Attribute{
					Name:        "rate_policy_id",
					Description: `(Optional) The ID of an existing rate policy to be modified. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rate_policy_id",
					Description: `The ID of the modified or newly created rate policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rate_policy_id",
					Description: `The ID of the modified or newly created rate policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_rate_policy_action",
			Category:         "Application Security",
			ShortDescription: `Rate Policy Action`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"rate",
				"policy",
				"action",
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
					Description: `(Required) The ID of the rate policy to use.`,
				},
				resource.Attribute{
					Name:        "ipv4_action",
					Description: `(Required) The ipv4 action to assign to this rate policy, either ` + "`" + `alert` + "`" + `, ` + "`" + `deny` + "`" + `, ` + "`" + `deny_custom_{custom_deny_id}` + "`" + `, or ` + "`" + `none` + "`" + `. If the action is none, the rate policy is inactive in the policy.`,
				},
				resource.Attribute{
					Name:        "ipv6_action",
					Description: `(Required) The ipv6 action to assign to this rate policy, either ` + "`" + `alert` + "`" + `, ` + "`" + `deny` + "`" + `, ` + "`" + `deny_custom_{custom_deny_id}` + "`" + `, or ` + "`" + `none` + "`" + `. If the action is none, the rate policy is inactive in the policy. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_rate_protection",
			Category:         "Application Security",
			ShortDescription: `Rate Protection`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"rate",
				"protection",
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
					Name:        "enabled",
					Description: `(Required) Whether to enable rate controls: either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current protection settings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current protection settings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_reputation_profile",
			Category:         "Application Security",
			ShortDescription: `Reputation Profile`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"reputation",
				"profile",
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
					Name:        "reputation_profile",
					Description: `(Required) The name of a file containing a JSON-formatted definition of the reputation profile. ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#postreputationprofiles)) ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "reputation_profile_id",
					Description: `The ID of the newly created or modified reputation profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "reputation_profile_id",
					Description: `The ID of the newly created or modified reputation profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_reputation_profile_action",
			Category:         "Application Security",
			ShortDescription: `Reputation Profile Action`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"reputation",
				"profile",
				"action",
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
					Name:        "reputation_profile_id",
					Description: `(Required) The ID of the reputation profile to use.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action to take when the specified reputation profile’s rule is triggered: ` + "`" + `alert` + "`" + ` to record the trigger event, ` + "`" + `deny` + "`" + ` to block the request, ` + "`" + `deny_custom_{custom_deny_id}` + "`" + ` to execute a custom deny action, or ` + "`" + `none` + "`" + ` to take no action. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `(Required) The ID of the security_policy_id to which the settings should be applied.`,
				},
				resource.Attribute{
					Name:        "forward_to_http_header",
					Description: `(Required) Whether to add client reputation details to requests forwarded to origin in an HTTP header.`,
				},
				resource.Attribute{
					Name:        "forward_shared_ip_to_http_header_siem",
					Description: `(Required) Whether to add value indicating that shared IPs are included in HTTP header and SIEM integration. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_reputation_protection",
			Category:         "Application Security",
			ShortDescription: `Reputation Protection`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"reputation",
				"protection",
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
					Name:        "enabled",
					Description: `(Required) Whether to enable reputation controls: either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current protection settings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current protection settings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_rule_action",
			Category:         "Application Security",
			ShortDescription: `Rule Action`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"rule",
				"action",
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
					Description: `(Required) The ID of the rule to use.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action to be taken: ` + "`" + `alert` + "`" + ` to record the trigger of the event, ` + "`" + `deny` + "`" + ` to block the request, ` + "`" + `deny_custom_{custom_deny_id}` + "`" + ` to execute a custom deny action, or ` + "`" + `none` + "`" + ` to take no action. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_rule_condition_exception",
			Category:         "Application Security",
			ShortDescription: `Rule Condition / Exception`,
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
					Description: `(Required) The ID of the rule to use.`,
				},
				resource.Attribute{
					Name:        "condition_exception",
					Description: `(Required) The name of a file containing a JSON-formatted description of the conditions and exceptions to use ([format](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putconditionexception)) ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_rule_upgrade",
			Category:         "Application Security",
			ShortDescription: `Rule Upgrade`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"rule",
				"upgrade",
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
					Name:        "current_ruleset",
					Description: `A string indicating the version number and release date of the current KRS rule set.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `A string indicating the current mode, either "KRS" or "AAG".`,
				},
				resource.Attribute{
					Name:        "eval_status",
					Description: `TBD`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "current_ruleset",
					Description: `A string indicating the version number and release date of the current KRS rule set.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `A string indicating the current mode, either "KRS" or "AAG".`,
				},
				resource.Attribute{
					Name:        "eval_status",
					Description: `TBD`,
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
					Description: `(Required) The configuration ID to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration to use.`,
				},
				resource.Attribute{
					Name:        "security_policy_name",
					Description: `(Required) The name of the new security policy.`,
				},
				resource.Attribute{
					Name:        "default_settings",
					Description: `(Optional) Whether the new policy should use the default settings. If not supplied, defaults to true. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `The ID of the newly created security policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `The ID of the newly created security policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_security_policy_clone",
			Category:         "Application Security",
			ShortDescription: `SecurityPolicyClone`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"policy",
				"clone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "create_from_version",
					Description: `(Required) The version number of the security configuration to clone.`,
				},
				resource.Attribute{
					Name:        "create_from_security_policy_id",
					Description: `(Required) The ID of the security policy to clone.`,
				},
				resource.Attribute{
					Name:        "security_policy_name",
					Description: `The name to be used for the new security policy. If not specified, the name will be autogenerated with the pattern 'clone from '.`,
				},
				resource.Attribute{
					Name:        "security_policy_prefix",
					Description: `The four-character policy prefix to be used for the new security policy. If not specified, the prefix will be autogenerated. ## Attribute Reference In addition to the arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `The ID of the new security policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_policy_id",
					Description: `The ID of the new security policy.`,
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
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "apply_application_layer_controls",
					Description: `(Required) Whether to enable application layer controls: either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "apply_network_layer_controls",
					Description: `(Required) Whether to enable network layer controls: either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "apply_rate_controls",
					Description: `(Required) Whether to enable rate controls: either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "apply_reputation_controls",
					Description: `(Required) Whether to enable reputation controls: either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "apply_botman_controls",
					Description: `(Required) Whether to enable botman controls: either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "apply_api_constraints",
					Description: `(Required) Whether to enable api constraints: either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "apply_slow_post_controls",
					Description: `(Required) Whether to enable slow post controls: either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current protection settings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current protection settings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_security_policy_rename",
			Category:         "Application Security",
			ShortDescription: `SecurityPolicyRename`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"policy",
				"rename",
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
					Description: `(Required) The ID of the security policy to be renamed.`,
				},
				resource.Attribute{
					Name:        "security_policy_name",
					Description: `(Required) The new name to be given to the security policy. ## Attribute Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required) The version number of the security configuration to use.`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `(Required) The list of hostnames to be applied, added or removed.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) A string specifying the interpretation of the ` + "`" + `hostnames` + "`" + ` parameter. Must be one of the following:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_siem_settings",
			Category:         "Application Security",
			ShortDescription: `SIEMSettings`,
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
					Description: `(Required) The configuration ID to use.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version number of the configuration to use.`,
				},
				resource.Attribute{
					Name:        "enable_siem",
					Description: `(Required) Whether you enabled SIEM in a security configuration version.`,
				},
				resource.Attribute{
					Name:        "enable_for_all_policies",
					Description: `(Required) Whether you enabled SIEM for all the security policies in the configuration version.`,
				},
				resource.Attribute{
					Name:        "enable_botman_siem",
					Description: `(Required) Whether you enabled SIEM for the Bot Manager events.`,
				},
				resource.Attribute{
					Name:        "siem_id",
					Description: `(Required) An integer that uniquely identifies the SIEM settings.`,
				},
				resource.Attribute{
					Name:        "security_policy_ids",
					Description: `(Required) The list of security policy identifiers for which to enable the SIEM integration. ## Attributes Reference In addition to the arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the updated SIEM integration settings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the updated SIEM integration settings.`,
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
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "slow_rate_action",
					Description: `(Required) The action that the rule should trigger (either ` + "`" + `alert` + "`" + ` or ` + "`" + `abort` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "slow_rate_threshold_rate",
					Description: `(Required) The average rate in bytes per second over the period specified by ` + "`" + `period` + "`" + ` before the specified ` + "`" + `action` + "`" + ` is triggered.`,
				},
				resource.Attribute{
					Name:        "slow_rate_threshold_period",
					Description: `(Required) The slow rate period value: the amount of time in seconds that the server should accept a request to determine whether a POST request is too slow.`,
				},
				resource.Attribute{
					Name:        "duration_threshold_timeout",
					Description: `(Required) The time in seconds before the first eight kilobytes of the POST body must be received to avoid triggering the specified ` + "`" + `action` + "`" + `. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_slowpost_protection",
			Category:         "Application Security",
			ShortDescription: `Slowpost Protection`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"slowpost",
				"protection",
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
					Name:        "enabled",
					Description: `(Required) Whether to enable slowpost controls: either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current protection settings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current protection settings.`,
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
					Description: `(Required) The version number of the configuration to use.`,
				},
				resource.Attribute{
					Name:        "version_notes",
					Description: `(Required) A string containing the version notes to be used. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the updated version notes.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the updated version notes.`,
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
					Description: `(Required) The ID of the security policy to use.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) "KRS" to update the rule sets manually, or "AAG" to have them update automatically. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "current_ruleset",
					Description: `The current rule set.`,
				},
				resource.Attribute{
					Name:        "eval_ruleset",
					Description: `The rule set being evaluated if any.`,
				},
				resource.Attribute{
					Name:        "eval_status",
					Description: `Either ` + "`" + `enabled` + "`" + ` if an evaluation is currently in progress, or ` + "`" + `disabled` + "`" + ` otherwise.`,
				},
				resource.Attribute{
					Name:        "eval_expiration_date",
					Description: `The date on which the evaluation period ends.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current rule set, WAF mode and evaluation status (` + "`" + `enabled` + "`" + ` if a rule set is currently being evaluated, ` + "`" + `disabled` + "`" + ` otherwise).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "current_ruleset",
					Description: `The current rule set.`,
				},
				resource.Attribute{
					Name:        "eval_ruleset",
					Description: `The rule set being evaluated if any.`,
				},
				resource.Attribute{
					Name:        "eval_status",
					Description: `Either ` + "`" + `enabled` + "`" + ` if an evaluation is currently in progress, or ` + "`" + `disabled` + "`" + ` otherwise.`,
				},
				resource.Attribute{
					Name:        "eval_expiration_date",
					Description: `The date on which the evaluation period ends.`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current rule set, WAF mode and evaluation status (` + "`" + `enabled` + "`" + ` if a rule set is currently being evaluated, ` + "`" + `disabled` + "`" + ` otherwise).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_appsec_waf_protection",
			Category:         "Application Security",
			ShortDescription: `WAF Protection`,
			Description:      ``,
			Keywords: []string{
				"application",
				"security",
				"appsec",
				"waf",
				"protection",
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
					Name:        "enabled",
					Description: `(Required) Whether to enable WAF controls: either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current protection settings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "output_text",
					Description: `A tabular display showing the current protection settings.`,
				},
			},
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
			Type:             "akamai_dns_record",
			Category:         "DNS",
			ShortDescription: `DNS record`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"record",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_dns_zone",
			Category:         "DNS",
			ShortDescription: `DNS Zone`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"zone",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_edge_hostname",
			Category:         "Provisioning",
			ShortDescription: `Edge Hostname`,
			Description:      ``,
			Keywords: []string{
				"provisioning",
				"edge",
				"hostname",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_asmap",
			Category:         "Global Traffic Management",
			ShortDescription: `GTM AS Map`,
			Description:      ``,
			Keywords: []string{
				"global",
				"traffic",
				"management",
				"gtm",
				"asmap",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_cidrmap",
			Category:         "Global Traffic Management",
			ShortDescription: `GTM Cidr Map`,
			Description:      ``,
			Keywords: []string{
				"global",
				"traffic",
				"management",
				"gtm",
				"cidrmap",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_datacenter",
			Category:         "Global Traffic Management",
			ShortDescription: `GTM Datacenter`,
			Description:      ``,
			Keywords: []string{
				"global",
				"traffic",
				"management",
				"gtm",
				"datacenter",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_domain",
			Category:         "Global Traffic Management",
			ShortDescription: `GTM Domain`,
			Description:      ``,
			Keywords: []string{
				"global",
				"traffic",
				"management",
				"gtm",
				"domain",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_geomap",
			Category:         "Global Traffic Management",
			ShortDescription: `GTM Geographic Map`,
			Description:      ``,
			Keywords: []string{
				"global",
				"traffic",
				"management",
				"gtm",
				"geomap",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_property",
			Category:         "Global Traffic Management",
			ShortDescription: `GTM Property`,
			Description:      ``,
			Keywords: []string{
				"global",
				"traffic",
				"management",
				"gtm",
				"property",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_resource",
			Category:         "Global Traffic Management",
			ShortDescription: `GTM Resource`,
			Description:      ``,
			Keywords: []string{
				"global",
				"traffic",
				"management",
				"gtm",
				"resource",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_iam_user",
			Category:         "IAM",
			ShortDescription: `Create user resources.`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_networklist_activations",
			Category:         "Network Lists",
			ShortDescription: `NetworkList Activations`,
			Description:      ``,
			Keywords: []string{
				"network",
				"lists",
				"networklist",
				"activations",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_list_id",
					Description: `(Required) The ID of the network list to be activated`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The network to be used, either ` + "`" + `STAGING` + "`" + ` or ` + "`" + `PRODUCTION` + "`" + `. If not supplied, defaults to ` + "`" + `STAGING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) A comment describing the activation.`,
				},
				resource.Attribute{
					Name:        "notification_emails",
					Description: `(Required) A bracketed, comma-separated list of email addresses that will be notified when the operation is complete. ## Attributes Reference In addition to the arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The string ` + "`" + `ACTIVATED` + "`" + ` if the activation was successful, or a string identifying the reason why the network list was not activated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The string ` + "`" + `ACTIVATED` + "`" + ` if the activation was successful, or a string identifying the reason why the network list was not activated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_networklist_description",
			Category:         "Network Lists",
			ShortDescription: `NetworkList Description`,
			Description:      ``,
			Keywords: []string{
				"network",
				"lists",
				"networklist",
				"description",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_list_id",
					Description: `(Required) The unique ID of the network list to use.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to be assigned to the network list.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description to be assigned to the network list. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_networklist_network_list",
			Category:         "Network Lists",
			ShortDescription: `NetworkList Network List`,
			Description:      ``,
			Keywords: []string{
				"network",
				"lists",
				"networklist",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to be assigned to the network list.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the network list; must be either "IP" or "GEO".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description to be assigned to the network list.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) A string specifying the interpretation of the ` + "`" + `list` + "`" + ` parameter. Must be one of the following:`,
				},
				resource.Attribute{
					Name:        "network_list_id",
					Description: `The ID of the network list.`,
				},
				resource.Attribute{
					Name:        "sync_point",
					Description: `An integer that identifies the current version of the network list; this value is incremented each time the list is modified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_list_id",
					Description: `The ID of the network list.`,
				},
				resource.Attribute{
					Name:        "sync_point",
					Description: `An integer that identifies the current version of the network list; this value is incremented each time the list is modified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_networklist_subscription",
			Category:         "Network Lists",
			ShortDescription: `NetworkList Subscription`,
			Description:      ``,
			Keywords: []string{
				"network",
				"lists",
				"networklist",
				"subscription",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_list",
					Description: `(Required) A list containing one or more IDs of the network lists to which the indicated email addresses should be subscribed.`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `(Required) A bracketed, comma-separated list of email addresses that will be notified of changes to any of the specified network lists. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property",
			Category:         "Provisioning",
			ShortDescription: `Create and update Akamai properties.`,
			Description:      ``,
			Keywords: []string{
				"provisioning",
				"property",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property_activation",
			Category:         "Provisioning",
			ShortDescription: `Property Activation`,
			Description:      ``,
			Keywords: []string{
				"provisioning",
				"property",
				"activation",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"akamai_appsec_activations":                      0,
		"akamai_appsec_advanced_settings_logging":        1,
		"akamai_appsec_advanced_settings_prefetch":       2,
		"akamai_appsec_api_request_constraints":          3,
		"akamai_appsec_attack_group_action":              4,
		"akamai_appsec_attack_group_condition_exception": 5,
		"akamai_appsec_bypass_network_lists":             6,
		"akamai_appsec_configuration":                    7,
		"akamai_appsec_configuration_clone":              8,
		"akamai_appsec_configuration_rename":             9,
		"akamai_appsec_configuration_version_clone":      10,
		"akamai_appsec_custom_deny":                      11,
		"akamai_appsec_custom_rule":                      12,
		"akamai_appsec_custom_rule_action":               13,
		"akamai_appsec_eval":                             14,
		"akamai_appsec_eval_hostnames":                   15,
		"akamai_appsec_eval_protect_host":                16,
		"akamai_appsec_eval_rule_action":                 17,
		"akamai_appsec_eval_rule_condition_exception":    18,
		"akamai_appsec_ip_geo":                           19,
		"akamai_appsec_match_target":                     20,
		"akamai_appsec_match_target_sequence":            21,
		"akamai_appsec_penalty_box":                      22,
		"akamai_appsec_rate_policy":                      23,
		"akamai_appsec_rate_policy_action":               24,
		"akamai_appsec_rate_protection":                  25,
		"akamai_appsec_reputation_profile":               26,
		"akamai_appsec_reputation_profile_action":        27,
		"akamai_appsec_reputation_profile_analysis":      28,
		"akamai_appsec_reputation_protection":            29,
		"akamai_appsec_rule_action":                      30,
		"akamai_appsec_rule_condition_exception":         31,
		"akamai_appsec_rule_upgrade":                     32,
		"akamai_appsec_security_policy":                  33,
		"akamai_appsec_security_policy_clone":            34,
		"akamai_appsec_security_policy_protections":      35,
		"akamai_appsec_security_policy_rename":           36,
		"akamai_appsec_selected_hostnames":               37,
		"akamai_appsec_siem_settings":                    38,
		"akamai_appsec_slow_post":                        39,
		"akamai_appsec_slowpost_protection":              40,
		"akamai_appsec_version_notes":                    41,
		"akamai_appsec_waf_mode":                         42,
		"akamai_appsec_waf_protection":                   43,
		"akamai_cp_code":                                 44,
		"akamai_dns_record":                              45,
		"akamai_dns_zone":                                46,
		"akamai_edge_hostname":                           47,
		"akamai_gtm_asmap":                               48,
		"akamai_gtm_cidrmap":                             49,
		"akamai_gtm_datacenter":                          50,
		"akamai_gtm_domain":                              51,
		"akamai_gtm_geomap":                              52,
		"akamai_gtm_property":                            53,
		"akamai_gtm_resource":                            54,
		"akamai_iam_user":                                55,
		"akamai_networklist_activations":                 56,
		"akamai_networklist_description":                 57,
		"akamai_networklist_network_list":                58,
		"akamai_networklist_subscription":                59,
		"akamai_property":                                60,
		"akamai_property_activation":                     61,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
