package zpa

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_access_policy_client_types",
			Category:         "Policy Set Controller",
			ShortDescription: `Get information about all client types for the specified customer.`,
			Description: `

Use the **zpa_access_policy_client_types** data source to get information about all client types for the specified customer in the Zscaler Private Access cloud. This data source can be optionally used when defining the following policy types:
    - ` + "`" + `` + "`" + `zpa_policy_access_rule` + "`" + `` + "`" + `
    - ` + "`" + `` + "`" + `zpa_policy_timeout_rule` + "`" + `` + "`" + `
    - ` + "`" + `` + "`" + `zpa_policy_forwarding_rule` + "`" + `` + "`" + `
    - ` + "`" + `` + "`" + `zpa_policy_isolation_rule` + "`" + `` + "`" + `
    - ` + "`" + `` + "`" + `zpa_policy_inspection_rule` + "`" + `` + "`" + `

The ` + "`" + `` + "`" + `object_type` + "`" + `` + "`" + ` attribute must be defined as "CLIENT_TYPE" in the policy operand condition. To learn more see the To learn more see the [Getting Details of All Client Types](https://help.zscaler.com/zpa/configuring-access-policies-using-api#getClientTypes)

-> **NOTE** By Default the ZPA provider will return all client types

-> **NOTE** When defining a ` + "`" + `` + "`" + `zpa_policy_isolation_rule` + "`" + `` + "`" + ` policy the ` + "`" + `` + "`" + `object_type` + "`" + `` + "`" + ` "CLIENT_TYPE" is mandatory and ` + "`" + `` + "`" + `zpn_client_type_exporter` + "`" + `` + "`" + ` is the only supported value.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"client",
				"types",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_access_policy_platforms",
			Category:         "Policy Set Controller",
			ShortDescription: `Get information about all platforms for the specified customer.`,
			Description: `

Use the **zpa_access_policy_platforms** data source to get information about all platforms for the specified customer in the Zscaler Private Access cloud. This data source can be optionally used when defining the following policy types:
    - ` + "`" + `` + "`" + `zpa_policy_access_rule` + "`" + `` + "`" + `
    - ` + "`" + `` + "`" + `zpa_policy_timeout_rule` + "`" + `` + "`" + `
    - ` + "`" + `` + "`" + `zpa_policy_forwarding_rule` + "`" + `` + "`" + `
    - ` + "`" + `` + "`" + `zpa_policy_isolation_rule` + "`" + `` + "`" + `
    - ` + "`" + `` + "`" + `zpa_policy_inspection_rule` + "`" + `` + "`" + `

The ` + "`" + `` + "`" + `object_type` + "`" + `` + "`" + ` attribute must be defined as "PLATFORM" in the policy operand condition. To learn more see the To learn more see the [Getting Platform Types for a Customer](https://help.zscaler.com/zpa/configuring-access-policies-using-api#getPlatformTypes)

-> **NOTE** By Default the ZPA provider will return all platform types

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"platforms",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_app_connector_controller",
			Category:         "App Connector Controller",
			ShortDescription: `Get information about ZPA App Connector in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_app_connector_controller** data source to get information about a app connector created in the Zscaler Private Access cloud. This data source can then be referenced in an App Connector Group.

`,
			Keywords: []string{
				"app",
				"connector",
				"controller",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the App Connector Group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Name of the App Connector Group. ## Attribute Reference In addition to all arguments above, the following attributes are exported: The following values are ignored in PUT/POST calls. Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Computed) - Longitude of the App Connector. Integer or decimal. With values in the range of ` + "`" + `-180` + "`" + ` to ` + "`" + `180` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Computed) - Whether this App Connector is enabled or not. Default value: ` + "`" + `true` + "`" + `. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "longitude",
					Description: `(Computed) - Longitude of the App Connector. Integer or decimal. With values in the range of ` + "`" + `-180` + "`" + ` to ` + "`" + `180` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Computed) - Whether this App Connector is enabled or not. Default value: ` + "`" + `true` + "`" + `. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_app_connector_group",
			Category:         "App Connector Group",
			ShortDescription: `Get information about ZPA App Connector Group in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_app_connector_group** data source to get information about a app connector group in the Zscaler Private Access cloud. This data source can then be referenced in an App Connector Group. This data source can then be referenced in the following resources:

* Create a server group
* Provisioning Key
* Access policy rule

`,
			Keywords: []string{
				"app",
				"connector",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the App Connector Group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the App Connector Group. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(String) Whether this App Connector Group is enabled or not. Default value: ` + "`" + `true` + "`" + `. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(String) Latitude of the App Connector Group. Integer or decimal. With values in the range of ` + "`" + `-90` + "`" + ` to ` + "`" + `90` + "`" + ``,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(String) Longitude of the App Connector Group. Integer or decimal. With values in the range of ` + "`" + `-180` + "`" + ` to ` + "`" + `180` + "`" + ``,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(String) Location of the App Connector Group.`,
				},
				resource.Attribute{
					Name:        "city_country",
					Description: `(String) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "upgrade_day",
					Description: `(String) App Connectors in this group will attempt to update to a newer version of the software during this specified day`,
				},
				resource.Attribute{
					Name:        "upgrade_time_in_secs",
					Description: `(String) App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: ` + "`" + `66600` + "`" + `. Integer in seconds (i.e., ` + "`" + `-66600` + "`" + `). The integer should be greater than or equal to ` + "`" + `0` + "`" + ` and less than ` + "`" + `86400` + "`" + `, in ` + "`" + `15` + "`" + ` minute intervals`,
				},
				resource.Attribute{
					Name:        "override_version_profile",
					Description: `(bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: ` + "`" + `false` + "`" + ` Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "version_profile_id",
					Description: `(String) ID of the version profile. Exported values are:`,
				},
				resource.Attribute{
					Name:        "version_profile_name",
					Description: `(String) Exported values are:`,
				},
				resource.Attribute{
					Name:        "version_profile_visibility_scope",
					Description: `(String) Exported values are:`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "dns_query_type",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "geo_location_id",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "use_in_dr_mode",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "pra_enabled",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `(String) Whether this App Connector Group is enabled or not. Default value: ` + "`" + `true` + "`" + `. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(String) Latitude of the App Connector Group. Integer or decimal. With values in the range of ` + "`" + `-90` + "`" + ` to ` + "`" + `90` + "`" + ``,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(String) Longitude of the App Connector Group. Integer or decimal. With values in the range of ` + "`" + `-180` + "`" + ` to ` + "`" + `180` + "`" + ``,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(String) Location of the App Connector Group.`,
				},
				resource.Attribute{
					Name:        "city_country",
					Description: `(String) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "upgrade_day",
					Description: `(String) App Connectors in this group will attempt to update to a newer version of the software during this specified day`,
				},
				resource.Attribute{
					Name:        "upgrade_time_in_secs",
					Description: `(String) App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: ` + "`" + `66600` + "`" + `. Integer in seconds (i.e., ` + "`" + `-66600` + "`" + `). The integer should be greater than or equal to ` + "`" + `0` + "`" + ` and less than ` + "`" + `86400` + "`" + `, in ` + "`" + `15` + "`" + ` minute intervals`,
				},
				resource.Attribute{
					Name:        "override_version_profile",
					Description: `(bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: ` + "`" + `false` + "`" + ` Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "version_profile_id",
					Description: `(String) ID of the version profile. Exported values are:`,
				},
				resource.Attribute{
					Name:        "version_profile_name",
					Description: `(String) Exported values are:`,
				},
				resource.Attribute{
					Name:        "version_profile_visibility_scope",
					Description: `(String) Exported values are:`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "dns_query_type",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "geo_location_id",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "use_in_dr_mode",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "pra_enabled",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_application_segment",
			Category:         "Application Segment",
			ShortDescription: `Get information about ZPA Application Segment in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_application_segment** data source to get information about a application segment created in the Zscaler Private Access cloud. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.

`,
			Keywords: []string{
				"application",
				"segment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the application.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(Optional) Indicates whether users can bypass ZPA to access applications. Default: ` + "`" + `NEVER` + "`" + `. Supported values: ` + "`" + `ALWAYS` + "`" + `, ` + "`" + `NEVER` + "`" + `, ` + "`" + `ON_NET` + "`" + `. The value ` + "`" + `NEVER` + "`" + ` indicates the use of the client forwarding policy.`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(Boolean) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors. Default: true. Boolean: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_checktype",
					Description: `(String) Whether health reporting for the app is Continuous or On Access. Supported values: ` + "`" + `NONE` + "`" + `, ` + "`" + `ON_ACCESS` + "`" + `, ` + "`" + `CONTINUOUS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(String) Whether Double Encryption is enabled or disabled for the app. Default: false. Boolean: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Boolean) Whether this application is enabled or not. Default: false. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "default_idle_timeout",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "default_max_age",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "segment_group_name",
					Description: `(string)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_application_segment_browser_access",
			Category:         "Application Segment",
			ShortDescription: `Get information about ZPA Browser Access Application Segment in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_application_segment_browser_access** data source to get information about a browser access application segment created in the Zscaler Private Access cloud. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.

`,
			Keywords: []string{
				"application",
				"segment",
				"browser",
				"access",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This field defines the name of the server.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) This field defines the id of the application server. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(string) Indicates whether users can bypass ZPA to access applications. Default: ` + "`" + `NEVER` + "`" + `. Supported values: ` + "`" + `ALWAYS` + "`" + `, ` + "`" + `NEVER` + "`" + `, ` + "`" + `ON_NET` + "`" + `. The value ` + "`" + `NEVER` + "`" + ` indicates the use of the client forwarding policy.`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors. Default: true. Boolean: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_checktype",
					Description: `(string) Whether health reporting for the app is Continuous or On Access. Supported values: ` + "`" + `NONE` + "`" + `, ` + "`" + `ON_ACCESS` + "`" + `, ` + "`" + `CONTINUOUS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(string) Whether Double Encryption is enabled or disabled for the app. Default: false. Boolean: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Boolean) Whether this application is enabled or not. Default: false. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(string) UDP port ranges used to access the app. ->`,
				},
				resource.Attribute{
					Name:        "tcp_port_range",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_range",
					Description: `(string) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_idle_timeout",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_max_age",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "segment_group_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "allow_options",
					Description: `(bool)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(string) Indicates whether users can bypass ZPA to access applications. Default: ` + "`" + `NEVER` + "`" + `. Supported values: ` + "`" + `ALWAYS` + "`" + `, ` + "`" + `NEVER` + "`" + `, ` + "`" + `ON_NET` + "`" + `. The value ` + "`" + `NEVER` + "`" + ` indicates the use of the client forwarding policy.`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors. Default: true. Boolean: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_checktype",
					Description: `(string) Whether health reporting for the app is Continuous or On Access. Supported values: ` + "`" + `NONE` + "`" + `, ` + "`" + `ON_ACCESS` + "`" + `, ` + "`" + `CONTINUOUS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(string) Whether Double Encryption is enabled or disabled for the app. Default: false. Boolean: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Boolean) Whether this application is enabled or not. Default: false. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(string) UDP port ranges used to access the app. ->`,
				},
				resource.Attribute{
					Name:        "tcp_port_range",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_range",
					Description: `(string) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_idle_timeout",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_max_age",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "segment_group_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "allow_options",
					Description: `(bool)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_application_segment_inspection",
			Category:         "Application Segment",
			ShortDescription: `Get information about ZPA Application Segment for Inspection.`,
			Description: `

Use the **zpa_application_segment_inspection** data source to get information about an inspection application segment in the Zscaler Private Access cloud. This resource can then be referenced in a ZPA access inspection policy. This resource supports ZPA Inspection for both ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + `.

`,
			Keywords: []string{
				"application",
				"segment",
				"inspection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Inspection Application Segment to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the Inspection Application Segment to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(string) List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `(string) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "id:",
					Description: `(string) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(String) Segment Group IDs`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "modifiedby",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(string) UDP port ranges used to access the app. ->`,
				},
				resource.Attribute{
					Name:        "tcp_port_range",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_range",
					Description: `(string) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(string) Indicates whether users can bypass ZPA to access applications.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(bool) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this application is enabled or not.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(string) Whether health reporting for the app is Continuous or On Access. Supported values: ` + "`" + `NONE` + "`" + `, ` + "`" + `ON_ACCESS` + "`" + `, ` + "`" + `CONTINUOUS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "icmp_access_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "inspection_apps",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "app_id:",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name:",
					Description: `(string) Name of the Inspection Application`,
				},
				resource.Attribute{
					Name:        "description:",
					Description: `(string) Description of the Inspection Application`,
				},
				resource.Attribute{
					Name:        "domain:",
					Description: `(string) Domain name of the inspection application`,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(string) TCP/UDP Port for ZPA Inspection.`,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(string) Protocol for the Inspection Application. Supported values: ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(string) - ID of the signing certificate. This field is required if the applicationProtocol is set to ` + "`" + `HTTPS` + "`" + `. The certificateId is not supported if the applicationProtocol is set to ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(string) - Parameter required when ` + "`" + `application_protocol` + "`" + ` is of type ` + "`" + `HTTPS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this application is enabled or not`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_names",
					Description: `(string) List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `(string) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "id:",
					Description: `(string) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(String) Segment Group IDs`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "modifiedby",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(string) UDP port ranges used to access the app. ->`,
				},
				resource.Attribute{
					Name:        "tcp_port_range",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_range",
					Description: `(string) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(string) Indicates whether users can bypass ZPA to access applications.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(bool) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this application is enabled or not.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(string) Whether health reporting for the app is Continuous or On Access. Supported values: ` + "`" + `NONE` + "`" + `, ` + "`" + `ON_ACCESS` + "`" + `, ` + "`" + `CONTINUOUS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "icmp_access_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "inspection_apps",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "app_id:",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name:",
					Description: `(string) Name of the Inspection Application`,
				},
				resource.Attribute{
					Name:        "description:",
					Description: `(string) Description of the Inspection Application`,
				},
				resource.Attribute{
					Name:        "domain:",
					Description: `(string) Domain name of the inspection application`,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(string) TCP/UDP Port for ZPA Inspection.`,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(string) Protocol for the Inspection Application. Supported values: ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(string) - ID of the signing certificate. This field is required if the applicationProtocol is set to ` + "`" + `HTTPS` + "`" + `. The certificateId is not supported if the applicationProtocol is set to ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(string) - Parameter required when ` + "`" + `application_protocol` + "`" + ` is of type ` + "`" + `HTTPS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this application is enabled or not`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_application_segment_pra",
			Category:         "Application Segment",
			ShortDescription: `Get information about ZPA Application Segment for Privileged Remote Access.`,
			Description: `

Use the **zpa_application_segment_pra** data source to get information about an application segment for Privileged Remote Access in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule, access policy timeout rule, access policy client forwarding rule and inspection policy. This resource supports Privileged Remote Access for both ` + "`" + `RDP` + "`" + ` and ` + "`" + `SSH` + "`" + `.

`,
			Keywords: []string{
				"application",
				"segment",
				"pra",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the PRA Application Segment to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(string) List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `(string) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "id:",
					Description: `(string) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(String) Segment Group IDs`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "modifiedby",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(string) UDP port ranges used to access the app. ->`,
				},
				resource.Attribute{
					Name:        "tcp_port_range",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_range",
					Description: `(string) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(string) Indicates whether users can bypass ZPA to access applications.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(bool) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this application is enabled or not.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(string) Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.`,
				},
				resource.Attribute{
					Name:        "icmp_access_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "sra_apps",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "app_id:",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name:",
					Description: `(string) Name of the Privileged Remote Access`,
				},
				resource.Attribute{
					Name:        "description:",
					Description: `(string) Description of the Privileged Remote Access`,
				},
				resource.Attribute{
					Name:        "domain:",
					Description: `(string) Domain name of the Privileged Remote Access`,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(string) Port for the Privileged Remote Accessvalues: ` + "`" + `RDP` + "`" + ` and ` + "`" + `SSH` + "`" + ``,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(string) Protocol for the Privileged Remote Access. Supported values: ` + "`" + `RDP` + "`" + ` and ` + "`" + `SSH` + "`" + ``,
				},
				resource.Attribute{
					Name:        "connection_security",
					Description: `(string) - Parameter required when ` + "`" + `application_protocol` + "`" + ` is of type ` + "`" + `RDP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this application is enabled or not`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_names",
					Description: `(string) List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `(string) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "id:",
					Description: `(string) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(String) Segment Group IDs`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "modifiedby",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(string) UDP port ranges used to access the app. ->`,
				},
				resource.Attribute{
					Name:        "tcp_port_range",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_range",
					Description: `(string) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(string) Indicates whether users can bypass ZPA to access applications.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(bool) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this application is enabled or not.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(string) Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.`,
				},
				resource.Attribute{
					Name:        "icmp_access_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "sra_apps",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "app_id:",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name:",
					Description: `(string) Name of the Privileged Remote Access`,
				},
				resource.Attribute{
					Name:        "description:",
					Description: `(string) Description of the Privileged Remote Access`,
				},
				resource.Attribute{
					Name:        "domain:",
					Description: `(string) Domain name of the Privileged Remote Access`,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(string) Port for the Privileged Remote Accessvalues: ` + "`" + `RDP` + "`" + ` and ` + "`" + `SSH` + "`" + ``,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(string) Protocol for the Privileged Remote Access. Supported values: ` + "`" + `RDP` + "`" + ` and ` + "`" + `SSH` + "`" + ``,
				},
				resource.Attribute{
					Name:        "connection_security",
					Description: `(string) - Parameter required when ` + "`" + `application_protocol` + "`" + ` is of type ` + "`" + `RDP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this application is enabled or not`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_application_server",
			Category:         "Application Server",
			ShortDescription: `Get information about ZPA Application Server in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_application_server** data source to get information about an application server created in the Zscaler Private Access cloud. This data source must be used in the following circumstances:

1. Server Group (When Dynamic Discovery is set to false)

`,
			Keywords: []string{
				"application",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This field defines the name of the server.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) This field defines the id of the application server. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) This field defines the description of the server.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(string) This field defines the domain or IP address of the server.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) This field defines the status of the server.`,
				},
				resource.Attribute{
					Name:        "app_server_group_ids",
					Description: `(Set of String) This field defines the list of server groups IDs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(string) This field defines the description of the server.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(string) This field defines the domain or IP address of the server.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) This field defines the status of the server.`,
				},
				resource.Attribute{
					Name:        "app_server_group_ids",
					Description: `(Set of String) This field defines the list of server groups IDs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_ba_certificate",
			Category:         "Browser Access Certificate",
			ShortDescription: `Get information about ZPA Browser Access Certificate in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_ba_certificate** data source to get information about a browser access certificate created in the Zscaler Private Access cloud. This data source is required when creating a browser access application segment resource.

`,
			Keywords: []string{
				"browser",
				"access",
				"certificate",
				"ba",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the browser access certificate to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the browser access certificate to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cert_chain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(string) The certificate text is in PEM format.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_to",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modifiedby",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "san",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "serial_no",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_from_in_epochsec",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_to_in_epochsec",
					Description: `(string) :warning: Notice that certificate and public_keys are omitted from the output.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cert_chain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(string) The certificate text is in PEM format.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_to",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modifiedby",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "san",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "serial_no",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_from_in_epochsec",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_to_in_epochsec",
					Description: `(string) :warning: Notice that certificate and public_keys are omitted from the output.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_browser_access",
			Category:         "Application Segment",
			ShortDescription: `Get information about ZPA Browser Access Application Segment in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_browser_access** data source to get information about a browser access application segment created in the Zscaler Private Access cloud. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.

`,
			Keywords: []string{
				"application",
				"segment",
				"browser",
				"access",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This field defines the name of the server.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) This field defines the id of the application server. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(string) Indicates whether users can bypass ZPA to access applications. Default: ` + "`" + `NEVER` + "`" + `. Supported values: ` + "`" + `ALWAYS` + "`" + `, ` + "`" + `NEVER` + "`" + `, ` + "`" + `ON_NET` + "`" + `. The value ` + "`" + `NEVER` + "`" + ` indicates the use of the client forwarding policy.`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors. Default: true. Boolean: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_checktype",
					Description: `(string) Whether health reporting for the app is Continuous or On Access. Supported values: ` + "`" + `NONE` + "`" + `, ` + "`" + `ON_ACCESS` + "`" + `, ` + "`" + `CONTINUOUS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(string) Whether Double Encryption is enabled or disabled for the app. Default: false. Boolean: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Boolean) Whether this application is enabled or not. Default: false. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(string) UDP port ranges used to access the app. ->`,
				},
				resource.Attribute{
					Name:        "tcp_port_range",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_range",
					Description: `(string) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_idle_timeout",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_max_age",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "segment_group_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "allow_options",
					Description: `(bool)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(string) Indicates whether users can bypass ZPA to access applications. Default: ` + "`" + `NEVER` + "`" + `. Supported values: ` + "`" + `ALWAYS` + "`" + `, ` + "`" + `NEVER` + "`" + `, ` + "`" + `ON_NET` + "`" + `. The value ` + "`" + `NEVER` + "`" + ` indicates the use of the client forwarding policy.`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors. Default: true. Boolean: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_checktype",
					Description: `(string) Whether health reporting for the app is Continuous or On Access. Supported values: ` + "`" + `NONE` + "`" + `, ` + "`" + `ON_ACCESS` + "`" + `, ` + "`" + `CONTINUOUS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(string) Whether Double Encryption is enabled or disabled for the app. Default: false. Boolean: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Boolean) Whether this application is enabled or not. Default: false. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(string) UDP port ranges used to access the app. ->`,
				},
				resource.Attribute{
					Name:        "tcp_port_range",
					Description: `(string) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_range",
					Description: `(string) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_idle_timeout",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_max_age",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "segment_group_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "allow_options",
					Description: `(bool)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_cloud_connector_group",
			Category:         "Cloud Connector Group",
			ShortDescription: `Get information about ZPA Cloud Connector Group in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_cloud_connector_group** data source to get information about a cloud connector group created from the Zscaler Private Access cloud. This data source can then be referenced within an Access Policy rule

~> **NOTE:** A Cloud Connector Group resource is created in the Zscaler Cloud Connector cloud and replicated to the ZPA cloud. This resource can then be referenced in a Access Policy Rule where the Object Type = ` + "`" + `CLOUD_CONNECTOR_GROUP` + "`" + ` is being used.

`,
			Keywords: []string{
				"cloud",
				"connector",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This field defines the name of the cloud connector group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) This field defines the id of the cloud connector group. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_customer_version_profile",
			Category:         "Customer Version Profile",
			ShortDescription: `Get information about all customer version profile details.`,
			Description: `

Use the **zpa_customer_version_profile** data source to get information about all customer version profiles from the Zscaler Private Access cloud. This data source can be associated with an App Connector Group within the parameter ` + "`" + `version_profile_id` + "`" + ` or ` + "`" + `version_profile_name` + "`" + `

The customer version profile IDs are:

* ` + "`" + `Default` + "`" + ` = ` + "`" + `0` + "`" + `
* ` + "`" + `Previous Default` + "`" + ` = ` + "`" + `1` + "`" + `
* ` + "`" + `New Release` + "`" + ` = ` + "`" + `2` + "`" + `

`,
			Keywords: []string{
				"customer",
				"version",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the enrollment certificate to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the enrollment certificate to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "allow_signing",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(string) The certificate text is in PEM format.`,
				},
				resource.Attribute{
					Name:        "client_cert_type",
					Description: `(string) Returned values are:`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_to",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "parent_cert_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "parent_cert_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "cert_chain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "serial_no",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_from_in_epoch_sec",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_to_in_epochsec",
					Description: `(string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_signing",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(string) The certificate text is in PEM format.`,
				},
				resource.Attribute{
					Name:        "client_cert_type",
					Description: `(string) Returned values are:`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_to",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "parent_cert_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "parent_cert_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "cert_chain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "serial_no",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_from_in_epoch_sec",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_to_in_epochsec",
					Description: `(string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_enrollment_cert",
			Category:         "Enrollment Certificate",
			ShortDescription: `Get information about all configured enrollment certificate details.`,
			Description: `

Use the **zpa_enrollment_cert** data source to get information about all configured enrollment certificate details created in the Zscaler Private Access cloud. This data source is required when creating provisioning key resources.

`,
			Keywords: []string{
				"enrollment",
				"certificate",
				"cert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the enrollment certificate to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the enrollment certificate to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "allow_signing",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(string) The certificate text is in PEM format.`,
				},
				resource.Attribute{
					Name:        "client_cert_type",
					Description: `(string) Returned values are:`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_to",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "parent_cert_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "parent_cert_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "cert_chain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "serial_no",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_from_in_epoch_sec",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_to_in_epochsec",
					Description: `(string) :warning: Notice that certificate, public and private key information are omitted from the output.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_signing",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(string) The certificate text is in PEM format.`,
				},
				resource.Attribute{
					Name:        "client_cert_type",
					Description: `(string) Returned values are:`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "issued_to",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "parent_cert_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "parent_cert_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "cert_chain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "serial_no",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_from_in_epoch_sec",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "valid_to_in_epochsec",
					Description: `(string) :warning: Notice that certificate, public and private key information are omitted from the output.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_idp_controller",
			Category:         "Identity Provider",
			ShortDescription: `Get information about an Identity Provider in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_idp_controller** data source to get information about an Identity Provider created in the Zscaler Private Access cloud. This data source is required when creating:

1. Access policy Rules
2. Access policy timeout rules
3. Access policy forwarding rules
4. Access policy inspection rules
5. Access policy isolation rules

`,
			Keywords: []string{
				"identity",
				"provider",
				"idp",
				"controller",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Identity Provider (IdP) to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The name of the Identity Provider (IdP) to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_inspection_all_predefined_controls",
			Category:         "Inspection",
			ShortDescription: `Get information about all Inspection Predefined Control by group name in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_inspection_all_predefined_controls** data source to get information about all OWASP predefined control and prefedined control version by group name. The ` + "`" + `Preprocessors` + "`" + ` predefined control is the default predefined control, This data source is always required, when creating an inspection profile.

`,
			Keywords: []string{
				"inspection",
				"all",
				"predefined",
				"controls",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) The name of the predefined control.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version of the predefined control, the default is: ` + "`" + `OWASP_CRS/3.3.0` + "`" + ` ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "control_group",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "control_number",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(string) Returned values: ` + "`" + `WEBSOCKET_PREDEFINED` + "`" + `, ` + "`" + `WEBSOCKET_CUSTOM` + "`" + `, ` + "`" + `ZSCALER` + "`" + `, ` + "`" + `CUSTOM` + "`" + `, ` + "`" + `PREDEFINED` + "`" + ``,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_action_value",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(string) Returned values: ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `, ` + "`" + `FTP` + "`" + `, ` + "`" + `RDP` + "`" + `, ` + "`" + `SSH` + "`" + `, ` + "`" + `WEBSOCKET` + "`" + ``,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "control_group",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "control_number",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(string) Returned values: ` + "`" + `WEBSOCKET_PREDEFINED` + "`" + `, ` + "`" + `WEBSOCKET_CUSTOM` + "`" + `, ` + "`" + `ZSCALER` + "`" + `, ` + "`" + `CUSTOM` + "`" + `, ` + "`" + `PREDEFINED` + "`" + ``,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_action_value",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(string) Returned values: ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `, ` + "`" + `FTP` + "`" + `, ` + "`" + `RDP` + "`" + `, ` + "`" + `SSH` + "`" + `, ` + "`" + `WEBSOCKET` + "`" + ``,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_inspection_custom_control",
			Category:         "Inspection",
			ShortDescription: `Get information about an Inspection Custom Control in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_inspection_custom_controls** data source to get information about an inspection custom control. This data source can be associated with an inspection profile.

`,
			Keywords: []string{
				"inspection",
				"custom",
				"control",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the predefined control.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version of the predefined control, the default is: ` + "`" + `OWASP_CRS/3.3.0` + "`" + ` ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(string) The performed action. Returned values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(string) Denotes the action`,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(string) Name of the inspection profile`,
				},
				resource.Attribute{
					Name:        "control_number",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(string) Returned values: ` + "`" + `WEBSOCKET_PREDEFINED` + "`" + `, ` + "`" + `WEBSOCKET_CUSTOM` + "`" + `, ` + "`" + `ZSCALER` + "`" + `, ` + "`" + `CUSTOM` + "`" + `, ` + "`" + `PREDEFINED` + "`" + ``,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(string) The performed action. Returned values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_action_value",
					Description: `(string) This is used to provide the redirect URL if the default action is set to ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the custom control`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string) Name of the custom control`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(string) OWASP Predefined Paranoia Level.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(string) Returned values: ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `, ` + "`" + `FTP` + "`" + `, ` + "`" + `RDP` + "`" + `, ` + "`" + `SSH` + "`" + `, ` + "`" + `WEBSOCKET` + "`" + ``,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(string) Severity of the control number. Returned values: ` + "`" + `CRITICAL` + "`" + `, ` + "`" + `ERROR` + "`" + `, ` + "`" + `WARNING` + "`" + `, ` + "`" + `INFO` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(string) Rules to be applied to the request or response type`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(string) Rules of the custom controls applied as conditions ` + "`" + `JSON` + "`" + ``,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "lhs",
					Description: `(string) Signifies the key for the object type Supported values: ` + "`" + `SIZE` + "`" + `, ` + "`" + `VALUE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(string) If lhs is set to SIZE, then the user may pass one of the following: ` + "`" + `EQ: Equals` + "`" + `, ` + "`" + `LE: Less than or equal to` + "`" + `, ` + "`" + `GE: Greater than or equal to` + "`" + `. If the lhs is set to ` + "`" + `VALUE` + "`" + `, then the user may pass one of the following: ` + "`" + `CONTAINS` + "`" + `, ` + "`" + `STARTS_WITH` + "`" + `, ` + "`" + `ENDS_WITH` + "`" + `, ` + "`" + `RX` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rhs",
					Description: `(string) Denotes the value for the given object type. Its value depends on the key. If rules.type is set to REQUEST_METHOD, the conditions.rhs field must have one of the following values: ` + "`" + `GET` + "`" + `,` + "`" + `HEAD` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `OPTIONS` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `TRACE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(string) Name of the rules. If rules.type is set to ` + "`" + `REQUEST_HEADERS` + "`" + `, ` + "`" + `REQUEST_COOKIES` + "`" + `, or ` + "`" + `RESPONSE_HEADERS` + "`" + `, the rules.name field is required.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(string) Type value for the rules`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(string) The performed action. Returned values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(string) Denotes the action`,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(string) Name of the inspection profile`,
				},
				resource.Attribute{
					Name:        "control_number",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(string) Returned values: ` + "`" + `WEBSOCKET_PREDEFINED` + "`" + `, ` + "`" + `WEBSOCKET_CUSTOM` + "`" + `, ` + "`" + `ZSCALER` + "`" + `, ` + "`" + `CUSTOM` + "`" + `, ` + "`" + `PREDEFINED` + "`" + ``,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(string) The performed action. Returned values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_action_value",
					Description: `(string) This is used to provide the redirect URL if the default action is set to ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the custom control`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string) Name of the custom control`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(string) OWASP Predefined Paranoia Level.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(string) Returned values: ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `, ` + "`" + `FTP` + "`" + `, ` + "`" + `RDP` + "`" + `, ` + "`" + `SSH` + "`" + `, ` + "`" + `WEBSOCKET` + "`" + ``,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(string) Severity of the control number. Returned values: ` + "`" + `CRITICAL` + "`" + `, ` + "`" + `ERROR` + "`" + `, ` + "`" + `WARNING` + "`" + `, ` + "`" + `INFO` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(string) Rules to be applied to the request or response type`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(string) Rules of the custom controls applied as conditions ` + "`" + `JSON` + "`" + ``,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "lhs",
					Description: `(string) Signifies the key for the object type Supported values: ` + "`" + `SIZE` + "`" + `, ` + "`" + `VALUE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(string) If lhs is set to SIZE, then the user may pass one of the following: ` + "`" + `EQ: Equals` + "`" + `, ` + "`" + `LE: Less than or equal to` + "`" + `, ` + "`" + `GE: Greater than or equal to` + "`" + `. If the lhs is set to ` + "`" + `VALUE` + "`" + `, then the user may pass one of the following: ` + "`" + `CONTAINS` + "`" + `, ` + "`" + `STARTS_WITH` + "`" + `, ` + "`" + `ENDS_WITH` + "`" + `, ` + "`" + `RX` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rhs",
					Description: `(string) Denotes the value for the given object type. Its value depends on the key. If rules.type is set to REQUEST_METHOD, the conditions.rhs field must have one of the following values: ` + "`" + `GET` + "`" + `,` + "`" + `HEAD` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `OPTIONS` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `TRACE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(string) Name of the rules. If rules.type is set to ` + "`" + `REQUEST_HEADERS` + "`" + `, ` + "`" + `REQUEST_COOKIES` + "`" + `, or ` + "`" + `RESPONSE_HEADERS` + "`" + `, the rules.name field is required.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(string) Type value for the rules`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_inspection_predefined_controls",
			Category:         "Inspection",
			ShortDescription: `Get information about an Inspection Predefined Control in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_inspection_predefined_controls** data source to get information about an OWASP predefined control and prefedined control version. This data source is required when creating an inspection profile.

`,
			Keywords: []string{
				"inspection",
				"predefined",
				"controls",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the predefined control.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version of the predefined control, the default is: ` + "`" + `OWASP_CRS/3.3.0` + "`" + ` ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "control_group",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "control_number",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "default_action_value",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Computed)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "control_group",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "control_number",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "default_action_value",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Computed)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_inspection_profile",
			Category:         "Inspection",
			ShortDescription: `Get information about an Inspection Profile in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_inspection_profile** data source to get information about an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.

`,
			Keywords: []string{
				"inspection",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This field defines the name of the inspection profile.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) This field defines the id of the inspection profile. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the inspection profile.`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(string) OWASP Predefined Paranoia Level. Range: [1-4], inclusive`,
				},
				resource.Attribute{
					Name:        "predefined_controls",
					Description: `(string) The predefined controls`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) ID of the predefined control`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(string) The action of the predefined control. Supported values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(string) Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to ` + "`" + `REDIRECT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_controls",
					Description: `(string) Types for custom controls`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(string) Rules of the custom controls applied as conditions ` + "`" + `JSON` + "`" + ``,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "lhs",
					Description: `(string) Signifies the key for the object type Supported values: ` + "`" + `SIZE` + "`" + `, ` + "`" + `VALUE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(string) If lhs is set to SIZE, then the user may pass one of the following: ` + "`" + `EQ: Equals` + "`" + `, ` + "`" + `LE: Less than or equal to` + "`" + `, ` + "`" + `GE: Greater than or equal to` + "`" + `. If the lhs is set to ` + "`" + `VALUE` + "`" + `, then the user may pass one of the following: ` + "`" + `CONTAINS` + "`" + `, ` + "`" + `STARTS_WITH` + "`" + `, ` + "`" + `ENDS_WITH` + "`" + `, ` + "`" + `RX` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rhs",
					Description: `(string) Denotes the value for the given object type. Its value depends on the key. If rules.type is set to REQUEST_METHOD, the conditions.rhs field must have one of the following values: ` + "`" + `GET` + "`" + `,` + "`" + `HEAD` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `OPTIONS` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `TRACE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(string) Name of the rules. If rules.type is set to ` + "`" + `REQUEST_HEADERS` + "`" + `, ` + "`" + `REQUEST_COOKIES` + "`" + `, or ` + "`" + `RESPONSE_HEADERS` + "`" + `, the rules.name field is required.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(string) Type value for the rules`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(string) The version of the predefined control, the default is: ` + "`" + `OWASP_CRS/3.3.0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(string) Name of the inspection profile`,
				},
				resource.Attribute{
					Name:        "common_global_override_actions_config",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "controls_info",
					Description: `(string) Types for custom controls`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(string) Control types. Supported Values: ` + "`" + `WEBSOCKET_PREDEFINED` + "`" + `, ` + "`" + `WEBSOCKET_CUSTOM` + "`" + `, ` + "`" + `CUSTOM` + "`" + `, ` + "`" + `PREDEFINED` + "`" + `, ` + "`" + `ZSCALER` + "`" + ``,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(string) Control information counts ` + "`" + `Long` + "`" + ``,
				},
				resource.Attribute{
					Name:        "web_socket_controls",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) ID of the predefined control`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(string) The action of the predefined control. Supported values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(string) Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to ` + "`" + `REDIRECT` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(string) Description of the inspection profile.`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(string) OWASP Predefined Paranoia Level. Range: [1-4], inclusive`,
				},
				resource.Attribute{
					Name:        "predefined_controls",
					Description: `(string) The predefined controls`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) ID of the predefined control`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(string) The action of the predefined control. Supported values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(string) Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to ` + "`" + `REDIRECT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_controls",
					Description: `(string) Types for custom controls`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(string) Rules of the custom controls applied as conditions ` + "`" + `JSON` + "`" + ``,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "lhs",
					Description: `(string) Signifies the key for the object type Supported values: ` + "`" + `SIZE` + "`" + `, ` + "`" + `VALUE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(string) If lhs is set to SIZE, then the user may pass one of the following: ` + "`" + `EQ: Equals` + "`" + `, ` + "`" + `LE: Less than or equal to` + "`" + `, ` + "`" + `GE: Greater than or equal to` + "`" + `. If the lhs is set to ` + "`" + `VALUE` + "`" + `, then the user may pass one of the following: ` + "`" + `CONTAINS` + "`" + `, ` + "`" + `STARTS_WITH` + "`" + `, ` + "`" + `ENDS_WITH` + "`" + `, ` + "`" + `RX` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rhs",
					Description: `(string) Denotes the value for the given object type. Its value depends on the key. If rules.type is set to REQUEST_METHOD, the conditions.rhs field must have one of the following values: ` + "`" + `GET` + "`" + `,` + "`" + `HEAD` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `OPTIONS` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `TRACE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(string) Name of the rules. If rules.type is set to ` + "`" + `REQUEST_HEADERS` + "`" + `, ` + "`" + `REQUEST_COOKIES` + "`" + `, or ` + "`" + `RESPONSE_HEADERS` + "`" + `, the rules.name field is required.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(string) Type value for the rules`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(string) The version of the predefined control, the default is: ` + "`" + `OWASP_CRS/3.3.0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(string) Name of the inspection profile`,
				},
				resource.Attribute{
					Name:        "common_global_override_actions_config",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "controls_info",
					Description: `(string) Types for custom controls`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(string) Control types. Supported Values: ` + "`" + `WEBSOCKET_PREDEFINED` + "`" + `, ` + "`" + `WEBSOCKET_CUSTOM` + "`" + `, ` + "`" + `CUSTOM` + "`" + `, ` + "`" + `PREDEFINED` + "`" + `, ` + "`" + `ZSCALER` + "`" + ``,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(string) Control information counts ` + "`" + `Long` + "`" + ``,
				},
				resource.Attribute{
					Name:        "web_socket_controls",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) ID of the predefined control`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(string) The action of the predefined control. Supported values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(string) Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to ` + "`" + `REDIRECT` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_isolation_profile",
			Category:         "Isolation",
			ShortDescription: `Get information about an Isolation Profile in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_isolation_profile** data source to get information about an isolation profile in the Zscaler Private Access cloud. This data source is required when configuring an isolation policy rule resource

`,
			Keywords: []string{
				"isolation",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This field defines the name of the isolation profile.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) This field defines the id of the isolation profile. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "isolation_profile_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "isolation_tenant_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "isolation_url",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "isolation_profile_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "isolation_tenant_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "isolation_url",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_lss_config_client_types",
			Category:         "Log Streaming (LSS)",
			ShortDescription: `Get information about all LSS client type details.`,
			Description: `

Use the **zpa_lss_config_client_types** data source to get information about all LSS client types in the Zscaler Private Access cloud. This data source is required when the defining a policy rule resource for an object type as ` + "`" + `CLIENT_TYPE` + "`" + ` parameter in the LSS Config Controller resource is set. To learn more see the To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSClientTypes)

-> **NOTE** By Default the ZPA provider will return all client types

`,
			Keywords: []string{
				"log",
				"streaming",
				"lss",
				"config",
				"client",
				"types",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_lss_config_controller",
			Category:         "Log Streaming (LSS)",
			ShortDescription: `Get information about Log Streaming (LSS) configuration Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_lss_config_controller** data source to get information about a Log Streaming (LSS) configuration resource created in the Zscaler Private Access.

`,
			Keywords: []string{
				"log",
				"streaming",
				"lss",
				"config",
				"controller",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This field defines the name of the log streaming resource.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) This field defines the name of the log streaming resource. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "audit_message",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "lss_host",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "lss_port",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "source_log_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "connector_groups",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "audit_message",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "lss_host",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "lss_port",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "source_log_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "connector_groups",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_lss_config_log_type_formats",
			Category:         "Log Streaming (LSS)",
			ShortDescription: `Get information about all all LSS log type format details.`,
			Description: `

Use the **zpa_lss_config_log_type_formats** data source to get information about all LSS log type formats in the Zscaler Private Access cloud. This data source is required when creating an LSS Config Controller resource.

`,
			Keywords: []string{
				"log",
				"streaming",
				"lss",
				"config",
				"type",
				"formats",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "log_type",
					Description: `(Required) The type of log to be exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_lss_config_status_codes",
			Category:         "Log Streaming (LSS)",
			ShortDescription: `Get information about all LSS status codes details.`,
			Description: `

Use the **zpa_lss_config_status_codes** data source to get information about all LSS status codes in the Zscaler Private Access cloud. This data source is required when the ` + "`" + `filter` + "`" + ` parameter in the LSS Config Controller resource is set. To learn more see the [Getting Details of All LSS Status Codes](https://help.zscaler.com/zpa/log-streaming-service-configuration-use-cases#GettingLSSStatusCodes)

-> **NOTE** By Default the ZPA provider will return all status codes

`,
			Keywords: []string{
				"log",
				"streaming",
				"lss",
				"config",
				"status",
				"codes",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_machine_group",
			Category:         "Machine Group",
			ShortDescription: `Get information about Machine Groups in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_machine_group** data source to get information about a machine group created in the Zscaler Private Access cloud. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.

`,
			Keywords: []string{
				"machine",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the machine group to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the machine group to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_type",
			Category:         "Policy Set Controller",
			ShortDescription: `Get information about Policy Set ID in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_policy_type** data source to get information about an a ` + "`" + `` + "`" + `policy_set_id` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `policy_type` + "`" + `` + "`" + `. This data source is required when creating:

1. Access policy Rules
2. Access policy timeout rules
3. Access policy forwarding rules
4. Access policy inspection rules

~> **NOTE** The parameters ` + "`" + `` + "`" + `policy_set_id` + "`" + `` + "`" + ` is required in all circumstances and is exported when checking for the policy_type parameter. The policy_type value is used for differentiating the policy types, in the request endpoint. The supported values are:

* ` + "`" + `` + "`" + `ACCESS_POLICY/GLOBAL_POLICY` + "`" + `` + "`" + `
* ` + "`" + `` + "`" + `TIMEOUT_POLICY/REAUTH_POLICY` + "`" + `` + "`" + `
* ` + "`" + `` + "`" + `BYPASS_POLICY/CLIENT_FORWARDING_POLICY` + "`" + `` + "`" + `
* ` + "`" + `` + "`" + `INSPECTION_POLICY` + "`" + `` + "`" + `

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"type",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Optional) The value for differentiating the policy types.`,
				},
				resource.Attribute{
					Name:        "policy_set_id",
					Description: `(Required) The ID of the global policy set. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_posture_profile",
			Category:         "Posture Profile",
			ShortDescription: `Get information about Posture Profile in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_posture_profile** data source to get information about a posture profile created in the Zscaler Private Access Mobile Portal. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.

`,
			Keywords: []string{
				"posture",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the posture profile to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the posture profile to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "master_customer_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "posture_udid",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zscaler_cloud",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zscaler_customer_id",
					Description: `(string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "master_customer_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "posture_udid",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zscaler_cloud",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zscaler_customer_id",
					Description: `(string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_provisioning_key",
			Category:         "Provisioning Key",
			ShortDescription: `Get information about Provisioning Key in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_provisioning_key** data source to get information about a provisioning key in the Zscaler Private Access portal or via API. This data source can be referenced in the following ZPA resources:

* App Connector Groups
* Service Edge Groups

-> **NOTE** The ` + "`" + `` + "`" + `association_type` + "`" + `` + "`" + ` parameter is required in order to distinguish between ` + "`" + `` + "`" + `CONNECTOR_GRP` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `SERVICE_EDGE_GRP` + "`" + `` + "`" + `

`,
			Keywords: []string{
				"provisioning",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the provisioning key.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the provisioning key to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "expiration_in_epoch_sec",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_acl",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "max_usage",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "provisioning_key",
					Description: `(string) Ignored in PUT/POST calls.`,
				},
				resource.Attribute{
					Name:        "enrollment_cert_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enrollment_cert_name",
					Description: `(string) Applicable only for GET calls, ignored in PUT/POST calls.`,
				},
				resource.Attribute{
					Name:        "ui_config",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "usage_count",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zcomponent_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zcomponent_name",
					Description: `(string) Applicable only for GET calls, ignored in PUT/POST calls. :warning: Notice that certificate and public_keys are omitted from the output.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "expiration_in_epoch_sec",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_acl",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "max_usage",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "provisioning_key",
					Description: `(string) Ignored in PUT/POST calls.`,
				},
				resource.Attribute{
					Name:        "enrollment_cert_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enrollment_cert_name",
					Description: `(string) Applicable only for GET calls, ignored in PUT/POST calls.`,
				},
				resource.Attribute{
					Name:        "ui_config",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "usage_count",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zcomponent_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zcomponent_name",
					Description: `(string) Applicable only for GET calls, ignored in PUT/POST calls. :warning: Notice that certificate and public_keys are omitted from the output.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_saml_attribute",
			Category:         "SAML Attributes",
			ShortDescription: `Get information about SAML attributes from an Identity Provider (IdP) in the Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_saml_attribute** data source to get information about a SAML Attributes from an Identity Provider (IdP). This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.

`,
			Keywords: []string{
				"saml",
				"attributes",
				"attribute",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the saml attribute to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the machine group to be exported.`,
				},
				resource.Attribute{
					Name:        "idp_name",
					Description: `(Optional) The name of the IdP corresponding to the SAML attribute. ->`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(string) The ID of the IdP corresponding to the SAML attribute.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "saml_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "user_attribute",
					Description: `(string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(string) The ID of the IdP corresponding to the SAML attribute.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "saml_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "user_attribute",
					Description: `(string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_scim_attribute_header",
			Category:         "SCIM Attribute Header",
			ShortDescription: `Get information about SCIM attributes from an Identity Provider (IdP) in the Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_scim_attribute_header** data source to get information about a SCIM attribute from an Identity Provider (IdP). This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Inspection Policy.

`,
			Keywords: []string{
				"scim",
				"attribute",
				"header",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the scim attribute header to be exported.`,
				},
				resource.Attribute{
					Name:        "idp_name",
					Description: `(Required) The name of the scim attribute header that must be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "canonical_values",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "data_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(string) The ID of the IdP corresponding to the SAML attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "canonical_values",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "data_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(string) The ID of the IdP corresponding to the SAML attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_scim_groups",
			Category:         "SCIM Groups",
			ShortDescription: `Get information about SCIM Group from an Identity Provider (IdP) in the Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_scim_groups** data source to get information about a SCIM Group from an Identity Provider (IdP). This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.

`,
			Keywords: []string{
				"scim",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name. The name of the scim group to be exported.`,
				},
				resource.Attribute{
					Name:        "idp_name",
					Description: `(Required) Name. The name of the IdP where the scim group must be exported from. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(string) The ID of the IdP corresponding to the SAML attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(string) The ID of the IdP corresponding to the SAML attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_segment_group",
			Category:         "Segment Group",
			ShortDescription: `Get information about Segment Groups in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_segment_group** data source to get information about a machine group created in the Zscaler Private Access cloud. This data source can then be referenced in an application segment or Access Policy rule.

-> **NOTE:** Segment Groups association is only supported in an Access Policy rule.

`,
			Keywords: []string{
				"segment",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the segment group to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the segment group to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "policy_migrated",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "tcp_keep_alive_enabled",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_idle_timeout",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_max_age",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "dynamic_discovery",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "policy_migrated",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "tcp_keep_alive_enabled",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_idle_timeout",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "default_max_age",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `(Computed)`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "dynamic_discovery",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_server_group",
			Category:         "Server Group",
			ShortDescription: `Get information about Server Groups in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_server_group** data source to get information about a server group created in the Zscaler Private Access cloud. This data source can then be referenced in an application segment, application server and Access Policy rule.

`,
			Keywords: []string{
				"server",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the server group to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the server group to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) This field is the description of the server group.`,
				},
				resource.Attribute{
					Name:        "dynamic_discovery",
					Description: `(bool) This field controls dynamic discovery of the servers.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) This field defines if the server group is enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config_space",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) This field is the description of the server group.`,
				},
				resource.Attribute{
					Name:        "dynamic_discovery",
					Description: `(bool) This field controls dynamic discovery of the servers.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) This field defines if the server group is enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(bool)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_service_edge_controller",
			Category:         "Service Edge Controller",
			ShortDescription: `Get information about Service Edge Controller in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_service_edge_controller** data source to get information about a service edge controller in the Zscaler Private Access cloud. This data source can then be referenced in a Service Edge Group and Provisioning Key.

`,
			Keywords: []string{
				"service",
				"edge",
				"controller",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the service edge controller to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the service edge controllerto be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this Service Edge Controller is enabled or not. Default value: ` + "`" + `true` + "`" + `. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(string) Latitude of the Service Edge Controller. Integer or decimal. With values in the range of ` + "`" + `-90` + "`" + ` to ` + "`" + `90` + "`" + ``,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(string) Longitude of the Service Edge Controller. Integer or decimal. With values in the range of ` + "`" + `-180` + "`" + ` to ` + "`" + `180` + "`" + ``,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(string) Location of the Service Edge Controller.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this Service Edge Controller is enabled or not. Default value: ` + "`" + `true` + "`" + `. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(string) Latitude of the Service Edge Controller. Integer or decimal. With values in the range of ` + "`" + `-90` + "`" + ` to ` + "`" + `90` + "`" + ``,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(string) Longitude of the Service Edge Controller. Integer or decimal. With values in the range of ` + "`" + `-180` + "`" + ` to ` + "`" + `180` + "`" + ``,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(string) Location of the Service Edge Controller.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_service_edge_group",
			Category:         "Service Edge Group",
			ShortDescription: `Get information about ZPA Service Edge Group in Zscaler Private Access cloud.`,
			Description: `

Use the **zpa_service_edge_group** data source to get information about a service edge group in the Zscaler Private Access cloud. This data source can then be referenced in an App Connector Group. This data source can then be referenced in the following resources:

* Create a server group
* Provisioning Key
* Access policy rule

`,
			Keywords: []string{
				"service",
				"edge",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the service edge group to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the service edge group to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this App Connector Group is enabled or not. Default value: ` + "`" + `true` + "`" + `. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "city_country",
					Description: `(string) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "geo_location_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(string) Latitude of the Service Edge Group. Integer or decimal. With values in the range of ` + "`" + `-90` + "`" + ` to ` + "`" + `90` + "`" + ``,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(string) Longitude of the Service Edge Group.Integer or decimal. With values in the range of ` + "`" + `-180` + "`" + ` to ` + "`" + `180` + "`" + ``,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(string) Location of the Service Edge Group.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "upgrade_day",
					Description: `(string) App Connectors in this group will attempt to update to a newer version of the software during this specified day`,
				},
				resource.Attribute{
					Name:        "upgrade_time_in_secs",
					Description: `(string) App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: ` + "`" + `66600` + "`" + `. Integer in seconds (i.e., ` + "`" + `-66600` + "`" + `). The integer should be greater than or equal to ` + "`" + `0` + "`" + ` and less than ` + "`" + `86400` + "`" + `, in ` + "`" + `15` + "`" + ` minute intervals`,
				},
				resource.Attribute{
					Name:        "override_version_profile",
					Description: `(bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: ` + "`" + `false` + "`" + ` Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "version_profile_id",
					Description: `(String) ID of the version profile. Exported values are:`,
				},
				resource.Attribute{
					Name:        "version_profile_name",
					Description: `(String) Exported values are:`,
				},
				resource.Attribute{
					Name:        "version_profile_visibility_scope",
					Description: `(string) Exported values are:`,
				},
				resource.Attribute{
					Name:        "service_edges",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "trusted_networks",
					Description: `(string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether this App Connector Group is enabled or not. Default value: ` + "`" + `true` + "`" + `. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "city_country",
					Description: `(string) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "geo_location_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(string) Latitude of the Service Edge Group. Integer or decimal. With values in the range of ` + "`" + `-90` + "`" + ` to ` + "`" + `90` + "`" + ``,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(string) Longitude of the Service Edge Group.Integer or decimal. With values in the range of ` + "`" + `-180` + "`" + ` to ` + "`" + `180` + "`" + ``,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(string) Location of the Service Edge Group.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "upgrade_day",
					Description: `(string) App Connectors in this group will attempt to update to a newer version of the software during this specified day`,
				},
				resource.Attribute{
					Name:        "upgrade_time_in_secs",
					Description: `(string) App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: ` + "`" + `66600` + "`" + `. Integer in seconds (i.e., ` + "`" + `-66600` + "`" + `). The integer should be greater than or equal to ` + "`" + `0` + "`" + ` and less than ` + "`" + `86400` + "`" + `, in ` + "`" + `15` + "`" + ` minute intervals`,
				},
				resource.Attribute{
					Name:        "override_version_profile",
					Description: `(bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: ` + "`" + `false` + "`" + ` Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "version_profile_id",
					Description: `(String) ID of the version profile. Exported values are:`,
				},
				resource.Attribute{
					Name:        "version_profile_name",
					Description: `(String) Exported values are:`,
				},
				resource.Attribute{
					Name:        "version_profile_visibility_scope",
					Description: `(string) Exported values are:`,
				},
				resource.Attribute{
					Name:        "service_edges",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "trusted_networks",
					Description: `(string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_trusted_network",
			Category:         "Trusted Network",
			ShortDescription: `Get information about Trusted Network in Zscaler Private Access cloud.`,
			Description: `

The **zpa_trusted_network** data source to get information about a trusted network created in the Zscaler Private Access Mobile Portal. This data source can then be referenced within the following resources:

1. Access Policy
2. Forwarding Policy
3. Inspection Policy
4. Isolation Policy
5. Service Edge Group.

`,
			Keywords: []string{
				"trusted",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the posture profile to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the posture profile to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "master_customer_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zscaler_cloud",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zscaler_customer_id",
					Description: `(string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "master_customer_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zscaler_cloud",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "zscaler_customer_id",
					Description: `(string)`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"zpa_zpa_access_policy_client_types":         0,
		"zpa_zpa_access_policy_platforms":            1,
		"zpa_zpa_app_connector_controller":           2,
		"zpa_zpa_app_connector_group":                3,
		"zpa_zpa_application_segment":                4,
		"zpa_zpa_application_segment_browser_access": 5,
		"zpa_zpa_application_segment_inspection":     6,
		"zpa_zpa_application_segment_pra":            7,
		"zpa_zpa_application_server":                 8,
		"zpa_zpa_ba_certificate":                     9,
		"zpa_zpa_browser_access":                     10,
		"zpa_zpa_cloud_connector_group":              11,
		"zpa_zpa_customer_version_profile":           12,
		"zpa_zpa_enrollment_cert":                    13,
		"zpa_zpa_idp_controller":                     14,
		"zpa_zpa_inspection_all_predefined_controls": 15,
		"zpa_zpa_inspection_custom_control":          16,
		"zpa_zpa_inspection_predefined_controls":     17,
		"zpa_zpa_inspection_profile":                 18,
		"zpa_zpa_isolation_profile":                  19,
		"zpa_zpa_lss_config_client_types":            20,
		"zpa_zpa_lss_config_controller":              21,
		"zpa_zpa_lss_config_log_type_formats":        22,
		"zpa_zpa_lss_config_status_codes":            23,
		"zpa_zpa_machine_group":                      24,
		"zpa_zpa_policy_type":                        25,
		"zpa_zpa_posture_profile":                    26,
		"zpa_zpa_provisioning_key":                   27,
		"zpa_zpa_saml_attribute":                     28,
		"zpa_zpa_scim_attribute_header":              29,
		"zpa_zpa_scim_groups":                        30,
		"zpa_zpa_segment_group":                      31,
		"zpa_zpa_server_group":                       32,
		"zpa_zpa_service_edge_controller":            33,
		"zpa_zpa_service_edge_group":                 34,
		"zpa_zpa_trusted_network":                    35,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
