package incapsula

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "incapsula_acl_security_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Incapsula ACL Security Rule resource.`,
			Description:      ``,
			Keywords: []string{
				"acl",
				"security",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Numeric identifier of the site to operate on.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) The id of the acl, e.g api.acl.blacklisted_ips. Options are ` + "`" + `api.acl.blacklisted_countries` + "`" + `, ` + "`" + `api.acl.blacklisted_urls` + "`" + `, ` + "`" + `api.acl.blacklisted_ips` + "`" + `, and ` + "`" + `api.acl.whitelisted_ips` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "continents",
					Description: `(Optional) A comma separated list of continent codes.`,
				},
				resource.Attribute{
					Name:        "countries",
					Description: `(Optional) A comma separated list of country codes.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `(Optional) A comma separated list of IPs or IP ranges, e.g: ` + "`" + `192.168.1.1` + "`" + `, ` + "`" + `192.168.1.1-192.168.1.100` + "`" + ` or ` + "`" + `192.168.1.1/24` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `(Optional) A comma separated list of resource paths.`,
				},
				resource.Attribute{
					Name:        "url_patterns",
					Description: `(Optional) The patterns should be in accordance with the matching urls sent by the urls parameter. Options are ` + "`" + `CONTAINS` + "`" + `, ` + "`" + `EQUALS` + "`" + `, ` + "`" + `PREFIX` + "`" + `, ` + "`" + `SUFFIX` + "`" + `, ` + "`" + `NOT_EQUALS` + "`" + `, ` + "`" + `NOT_CONTAIN` + "`" + `, ` + "`" + `NOT_PREFIX` + "`" + `,and ` + "`" + `NOT_SUFFIX` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the ACL security rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the ACL security rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "incapsula_custom_certificate",
			Category:         "Resources",
			ShortDescription: `Provides a Incapsula Custom Certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"custom",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Numeric identifier of the site to operate on.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The certificate file in base64 format. You can use the Terraform HCL ` + "`" + `file` + "`" + ` directive to pull in the contents from a file. You can also inline the certificate in the configuration.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) The private key of the certificate in base64 format. Optional in case of PFX certificate file format.`,
				},
				resource.Attribute{
					Name:        "passphrase",
					Description: `(Optional) The passphrase used to protect your SSL certificate. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `At the moment, only one active certificate can be stored. This exported value is always set as ` + "`" + `12345` + "`" + `. This will be augmented in future versions of the API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `At the moment, only one active certificate can be stored. This exported value is always set as ` + "`" + `12345` + "`" + `. This will be augmented in future versions of the API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "incapsula_data_center",
			Category:         "Resources",
			ShortDescription: `Provides a Incapsula Data Center resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"center",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Numeric identifier of the site to operate on.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The new data center's name.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Required) The server's address. Possible values: IP, CNAME.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Enables the data center.`,
				},
				resource.Attribute{
					Name:        "is_standby",
					Description: `(Optional) Defines the data center as standby for failover.`,
				},
				resource.Attribute{
					Name:        "is_content",
					Description: `(Optional) The data center will be available for specific resources (Forward Delivery Rules). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the data center.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the data center.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "incapsula_data_center_server",
			Category:         "Resources",
			ShortDescription: `Provides a Incapsula Data Center Server resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dc_id",
					Description: `(Required) Numeric identifier of the data center server to operate on.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Numeric identifier of the site to operate on.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Optional) The server's address.`,
				},
				resource.Attribute{
					Name:        "is_standby",
					Description: `(Optional) Set the server as Active (P0) or Standby (P1).`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Enables the data center server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the data center server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the data center server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "incapsula_incap_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Incapsula Incap Rule resource.`,
			Description:      ``,
			Keywords: []string{
				"incap",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Numeric identifier of the site to operate on.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Rule name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Rule action. See the detailed descriptions in the API documentation. Possible values: ` + "`" + `RULE_ACTION_REDIRECT` + "`" + `, ` + "`" + `RULE_ACTION_SIMPLIFIED_REDIRECT` + "`" + `, ` + "`" + `RULE_ACTION_REWRITE_URL` + "`" + `, ` + "`" + `RULE_ACTION_REWRITE_HEADER` + "`" + `, ` + "`" + `RULE_ACTION_REWRITE_COOKIE` + "`" + `, ` + "`" + `RULE_ACTION_DELETE_HEADER` + "`" + `, ` + "`" + `RULE_ACTION_DELETE_COOKIE` + "`" + `, ` + "`" + `RULE_ACTION_RESPONSE_REWRITE_HEADER` + "`" + `, ` + "`" + `RULE_ACTION_RESPONSE_DELETE_HEADER` + "`" + `, ` + "`" + `RULE_ACTION_RESPONSE_REWRITE_RESPONSE_CODE` + "`" + `, ` + "`" + `RULE_ACTION_FORWARD_TO_DC` + "`" + `, ` + "`" + `RULE_ACTION_ALERT` + "`" + `, ` + "`" + `RULE_ACTION_BLOCK` + "`" + `, ` + "`" + `RULE_ACTION_BLOCK_USER` + "`" + `, ` + "`" + `RULE_ACTION_BLOCK_IP` + "`" + `, ` + "`" + `RULE_ACTION_RETRY` + "`" + `, ` + "`" + `RULE_ACTION_INTRUSIVE_HTML` + "`" + `, ` + "`" + `RULE_ACTION_CAPTCHA` + "`" + `, ` + "`" + `RULE_ACTION_RATE` + "`" + `, ` + "`" + `RULE_ACTION_CUSTOM_ERROR_RESPONSE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) The filter defines the conditions that trigger the rule action. For action ` + "`" + `RULE_ACTION_SIMPLIFIED_REDIRECT` + "`" + ` filter is not relevant. For other actions, if left empty, the rule is always run.`,
				},
				resource.Attribute{
					Name:        "response_code",
					Description: `(Optional) For ` + "`" + `RULE_ACTION_REDIRECT` + "`" + ` or ` + "`" + `RULE_ACTION_SIMPLIFIED_REDIRECT` + "`" + ` rule's response code, valid values are ` + "`" + `302` + "`" + `, ` + "`" + `301` + "`" + `, ` + "`" + `303` + "`" + `, ` + "`" + `307` + "`" + `, ` + "`" + `308` + "`" + `. For ` + "`" + `RULE_ACTION_RESPONSE_REWRITE_RESPONSE_CODE` + "`" + ` rule's response code, valid values are all 3-digits numbers. For ` + "`" + `RULE_ACTION_CUSTOM_ERROR_RESPONSE` + "`" + `, valid values are ` + "`" + `400` + "`" + `, ` + "`" + `401` + "`" + `, ` + "`" + `402` + "`" + `, ` + "`" + `403` + "`" + `, ` + "`" + `404` + "`" + `, ` + "`" + `405` + "`" + `, ` + "`" + `406` + "`" + `, ` + "`" + `407` + "`" + `, ` + "`" + `408` + "`" + `, ` + "`" + `409` + "`" + `, ` + "`" + `410` + "`" + `, ` + "`" + `411` + "`" + `, ` + "`" + `412` + "`" + `, ` + "`" + `413` + "`" + `, ` + "`" + `414` + "`" + `, ` + "`" + `415` + "`" + `, ` + "`" + `416` + "`" + `, ` + "`" + `417` + "`" + `, ` + "`" + `419` + "`" + `, ` + "`" + `420` + "`" + `, ` + "`" + `422` + "`" + `, ` + "`" + `423` + "`" + `, ` + "`" + `424` + "`" + `, ` + "`" + `500` + "`" + `, ` + "`" + `501` + "`" + `, ` + "`" + `502` + "`" + `, ` + "`" + `503` + "`" + `, ` + "`" + `504` + "`" + `, ` + "`" + `505` + "`" + `, ` + "`" + `507` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "add_missing",
					Description: `(Optional) Add cookie or header if it doesn't exist (Rewrite cookie rule only).`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Optional) Pattern to rewrite. For ` + "`" + `RULE_ACTION_REWRITE_URL` + "`" + ` - Url to rewrite. For ` + "`" + `RULE_ACTION_REWRITE_HEADER` + "`" + ` and ` + "`" + `RULE_ACTION_RESPONSE_REWRITE_HEADER` + "`" + ` - Header value to rewrite. For ` + "`" + `RULE_ACTION_REWRITE_COOKIE` + "`" + ` - Cookie value to rewrite.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Optional) Pattern to change to. ` + "`" + `RULE_ACTION_REWRITE_URL` + "`" + ` - Url to change to. ` + "`" + `RULE_ACTION_REWRITE_HEADER` + "`" + ` and ` + "`" + `RULE_ACTION_RESPONSE_REWRITE_HEADER` + "`" + ` - Header value to change to. ` + "`" + `RULE_ACTION_REWRITE_COOKIE` + "`" + ` - Cookie value to change to.`,
				},
				resource.Attribute{
					Name:        "rewrite_name",
					Description: `(Optional) Name of cookie or header to rewrite. Applies only for ` + "`" + `RULE_ACTION_REWRITE_COOKIE` + "`" + `, ` + "`" + `RULE_ACTION_REWRITE_HEADER` + "`" + ` and ` + "`" + `RULE_ACTION_RESPONSE_REWRITE_HEADER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `(Optional) Data center to forward request to. Applies only for ` + "`" + `RULE_ACTION_FORWARD_TO_DC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rate_context",
					Description: `(Optional) The context of the rate counter. Possible values ` + "`" + `IP` + "`" + ` or ` + "`" + `Session` + "`" + `. Applies only to rules using ` + "`" + `RULE_ACTION_RATE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rate_interval",
					Description: `(Optional) The interval in seconds of the rate counter. Possible values is a multiple of ` + "`" + `10` + "`" + `; minimum ` + "`" + `10` + "`" + ` and maximum ` + "`" + `300` + "`" + `. Applies only to rules using ` + "`" + `RULE_ACTION_RATE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "error_type",
					Description: `(Optional) The error that triggers the rule. ` + "`" + `error.type.all` + "`" + ` triggers the rule regardless of the error type. Applies only for ` + "`" + `RULE_ACTION_CUSTOM_ERROR_RESPONSE` + "`" + `. Possible values: ` + "`" + `error.type.all` + "`" + `, ` + "`" + `error.type.connection_timeout` + "`" + `, ` + "`" + `error.type.access_denied` + "`" + `, ` + "`" + `error.type.parse_req_error` + "`" + `, ` + "`" + `error.type.parse_resp_error` + "`" + `, ` + "`" + `error.type.connection_failed` + "`" + `, ` + "`" + `error.type.deny_and_retry` + "`" + `, ` + "`" + `error.type.ssl_failed` + "`" + `, ` + "`" + `error.type.deny_and_captcha` + "`" + `, ` + "`" + `error.type.2fa_required` + "`" + `, ` + "`" + `error.type.no_ssl_config` + "`" + `, ` + "`" + `error.type.no_ipv6_config` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "error_response_format",
					Description: `(Optional) The format of the given error response in the error_response_data field. Applies only for ` + "`" + `RULE_ACTION_CUSTOM_ERROR_RESPONSE` + "`" + `. Possible values: ` + "`" + `json` + "`" + `, ` + "`" + `xml` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "error_response_data",
					Description: `(Optional) The response returned when the request matches the filter and is blocked. Applies only for ` + "`" + `RULE_ACTION_CUSTOM_ERROR_RESPONSE` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the Incap Rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the Incap Rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "incapsula_site",
			Category:         "Resources",
			ShortDescription: `Provides a Incapsula Site resource.`,
			Description:      ``,
			Keywords: []string{
				"site",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The fully qualified domain name of the site. For example: www.example.com, hello.example.com.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The account to operate on. If not specified, operation will be performed on the account identified by the authentication parameters.`,
				},
				resource.Attribute{
					Name:        "send_site_setup_emails",
					Description: `(Optional) If this value is false, end users will not get emails about the add site process such as DNS instructions and SSL setup.`,
				},
				resource.Attribute{
					Name:        "site_ip",
					Description: `(Optional) The web server IP/CNAME.`,
				},
				resource.Attribute{
					Name:        "force_ssl",
					Description: `(Optional) Force SSL. This option is only available for sites with manually configured IP/CNAME and for specific accounts.`,
				},
				resource.Attribute{
					Name:        "log_level",
					Description: `(Optional) Log level. Available only for Enterprise Plan customers that purchased the Logs Integration SKU. Sets the log reporting level for the site. Options are ` + "`" + `full` + "`" + `, ` + "`" + `security` + "`" + `, ` + "`" + `none` + "`" + `, and ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logs_account_id",
					Description: `(Optional) Account where logs should be stored. Available only for Enterprise Plan customers that purchased the Logs Integration SKU. Numeric identifier of the account that purchased the logs integration SKU and which collects the logs. If not specified, operation will be performed on the account identified by the authentication parameters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the site.`,
				},
				resource.Attribute{
					Name:        "site_creation_date",
					Description: `Numeric representation of the site creation date.`,
				},
				resource.Attribute{
					Name:        "dns_cname_record_name",
					Description: `The CNAME record name.`,
				},
				resource.Attribute{
					Name:        "dns_cname_record_value",
					Description: `The CNAME record value.`,
				},
				resource.Attribute{
					Name:        "dns_a_record_name",
					Description: `The A record name.`,
				},
				resource.Attribute{
					Name:        "dns_a_record_value",
					Description: `The A record value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the site.`,
				},
				resource.Attribute{
					Name:        "site_creation_date",
					Description: `Numeric representation of the site creation date.`,
				},
				resource.Attribute{
					Name:        "dns_cname_record_name",
					Description: `The CNAME record name.`,
				},
				resource.Attribute{
					Name:        "dns_cname_record_value",
					Description: `The CNAME record value.`,
				},
				resource.Attribute{
					Name:        "dns_a_record_name",
					Description: `The A record name.`,
				},
				resource.Attribute{
					Name:        "dns_a_record_value",
					Description: `The A record value.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "incapsula_waf_security_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Incapsula WAF Security Rule resource.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"security",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Numeric identifier of the site to operate on.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) The identifier of the WAF rule, e.g api.threats.cross_site_scripting.`,
				},
				resource.Attribute{
					Name:        "security_rule_action",
					Description: `(Optional) The action that should be taken when a threat is detected, for example: api.threats.action.block_ip. See above examples for ` + "`" + `rule_id` + "`" + ` and ` + "`" + `action` + "`" + ` combinations.`,
				},
				resource.Attribute{
					Name:        "activation_mode",
					Description: `(Optional) The mode of activation for ddos on a site. Possible values: off, auto, on.`,
				},
				resource.Attribute{
					Name:        "ddos_traffic_threshold",
					Description: `(Optional) Consider site to be under DDoS if the request rate is above this threshold. The valid values are 10, 20, 50, 100, 200, 500, 750, 1000, 2000, 3000, 4000, 5000.`,
				},
				resource.Attribute{
					Name:        "block_bad_bots",
					Description: `(Optional) Whether or not to block bad bots. Possible values: true, false.`,
				},
				resource.Attribute{
					Name:        "challenge_suspected_bots",
					Description: `(Optional) Whether or not to send a challenge to clients that are suspected to be bad bots (CAPTCHA for example). Possible values: true, false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the Incap Rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier in the API for the Incap Rule.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"incapsula_acl_security_rule":  0,
		"incapsula_custom_certificate": 1,
		"incapsula_data_center":        2,
		"incapsula_data_center_server": 3,
		"incapsula_incap_rule":         4,
		"incapsula_site":               5,
		"incapsula_waf_security_rule":  6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
