package zpa

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_app_connector_group",
			Category:         "App Connector Group",
			ShortDescription: `Creates and manages ZPA App Connector Groups.`,
			Description: `

The **zpa_app_connector_group** resource creates a and manages app connector groups in the Zscaler Private Access (ZPA) cloud. This resource can then be associated with the following resources: server groups, log receivers and access policies.

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
					Name:        "enabled",
					Description: `(Optional) Whether this App Connector Group is enabled or not. Default value: ` + "`" + `true` + "`" + `. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Required) Latitude of the App Connector Group. Integer or decimal. With values in the range of ` + "`" + `-90` + "`" + ` to ` + "`" + `90` + "`" + ``,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Required) Longitude of the App Connector Group. Integer or decimal. With values in the range of ` + "`" + `-180` + "`" + ` to ` + "`" + `180` + "`" + ``,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Location of the App Connector Group. i.e ` + "`" + `` + "`" + `"San Jose, CA, USA"` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "city_country",
					Description: `(Optional) Whether Double Encryption is enabled or disabled for the app. i.e ` + "`" + `` + "`" + `"San Jose, US"` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "upgrade_day",
					Description: `(Optional) App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ` + "`" + `` + "`" + `SUNDAY` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "upgrade_time_in_secs",
					Description: `(Optional) App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: ` + "`" + `66600` + "`" + `. Integer in seconds (i.e., ` + "`" + `-66600` + "`" + `). The integer should be greater than or equal to ` + "`" + `0` + "`" + ` and less than ` + "`" + `86400` + "`" + `, in ` + "`" + `15` + "`" + ` minute intervals`,
				},
				resource.Attribute{
					Name:        "override_version_profile",
					Description: `(Optional) Whether the default version profile of the App Connector Group is applied or overridden. Default: ` + "`" + `false` + "`" + ` Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "version_profile_id",
					Description: `(Optional) ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:`,
				},
				resource.Attribute{
					Name:        "version_profile_name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "version_profile_visibility_scope",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(Optional) i.e ` + "`" + `` + "`" + `"US"` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `"CA"` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns_query_type",
					Description: `(Optional) Supported values are:`,
				},
				resource.Attribute{
					Name:        "tcp_quick_ack_app",
					Description: `(Optional) Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tcp_quick_ack_assistant",
					Description: `(Optional) Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tcp_quick_ack_read_assistant",
					Description: `(Optional) Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "use_in_dr_mode",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "pra_enabled",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ` ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Group Role Assignment. ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) App Connector Group can be imported by using ` + "`" + `<APP CONNECTOR GROUP ID>` + "`" + ` or ` + "`" + `<APP CONNECTOR GROUP NAME>` + "`" + `as the import ID. ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_app_connector_group.example <app_connector_group_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_app_connector_group.example <app_connector_group_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Group Role Assignment. ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) App Connector Group can be imported by using ` + "`" + `<APP CONNECTOR GROUP ID>` + "`" + ` or ` + "`" + `<APP CONNECTOR GROUP NAME>` + "`" + `as the import ID. ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_app_connector_group.example <app_connector_group_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_app_connector_group.example <app_connector_group_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_application_segment",
			Category:         "Application Segment",
			ShortDescription: `Creates and manages ZPA Application Segments.`,
			Description: `

The **zpa_application_segment** resource creates an application segment in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule, access policy timeout rule or access policy client forwarding rule.

`,
			Keywords: []string{
				"application",
				"segment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name. The name of the App Connector Group to be exported.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Required) List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `(Required) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(Required) List of Segment Group IDs`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(Required) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(Required) UDP port ranges used to access the app. ->`,
				},
				resource.Attribute{
					Name:        "tcp_port_range",
					Description: `(Required) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_range",
					Description: `(Required) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(Optional) Indicates whether users can bypass ZPA to access applications.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(Optional) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this application is enabled or not.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(Optional) Whether health reporting for the app is Continuous or On Access. Supported values: ` + "`" + `NONE` + "`" + `, ` + "`" + `ON_ACCESS` + "`" + `, ` + "`" + `CONTINUOUS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "icmp_access_type",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.`,
				},
				resource.Attribute{
					Name:        "log_features",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "use_in_dr_mode",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_incomplete_dr_config",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ` ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Application Segment can be imported by using ` + "`" + `<APPLICATION SEGMENT ID>` + "`" + ` or ` + "`" + `<APPLICATION SEGMENT NAME>` + "`" + ` as the import ID. ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment.example <application_segment_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment.example <application_segment_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(Optional) Indicates whether users can bypass ZPA to access applications.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(Optional) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this application is enabled or not.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(Optional) Whether health reporting for the app is Continuous or On Access. Supported values: ` + "`" + `NONE` + "`" + `, ` + "`" + `ON_ACCESS` + "`" + `, ` + "`" + `CONTINUOUS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "icmp_access_type",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.`,
				},
				resource.Attribute{
					Name:        "log_features",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "use_in_dr_mode",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_incomplete_dr_config",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ` ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Application Segment can be imported by using ` + "`" + `<APPLICATION SEGMENT ID>` + "`" + ` or ` + "`" + `<APPLICATION SEGMENT NAME>` + "`" + ` as the import ID. ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment.example <application_segment_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment.example <application_segment_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_application_segment_browser_access",
			Category:         "Application Segment",
			ShortDescription: `Creates and manages ZPA Browser Access Application Segment.`,
			Description: `

The **zpa_application_segment_browser_access** creates and manages a browser access application segment resource in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule, access policy timeout rule or access policy client forwarding rule.

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
					Description: `(Required) Name of the application.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Required) List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(Required) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(Required) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `(Required) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(Required) List of Segment Group IDs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Name of BA app.`,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(Required) - Port for the BA app.`,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(Required) - Protocol for the BA app. Supported values: ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Required) - ID of the BA certificate. Refer to the data source documentation for [` + "`" + `zpa_ba_certificate` + "`" + `](https://github.com/zscaler/terraform-provider-zpa/blob/master/docs/data-sources/zpa_ba_certificate.md)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) - Domain name or IP address of the BA app.`,
				},
				resource.Attribute{
					Name:        "allow_options",
					Description: `(Optional) - If you want ZPA to forward unauthenticated HTTP preflight OPTIONS requests from the browser to the app.. Supported values: ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + ` ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(Optional) - Name of the BA certificate. Refer to the data source documentation for [` + "`" + `zpa_ba_certificate` + "`" + `](https://github.com/zscaler/terraform-provider-zpa/blob/master/docs/data-sources/zpa_ba_certificate.md)`,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "use_in_dr_mode",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_incomplete_dr_config",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ` ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(Optional) - Name of the BA certificate. Refer to the data source documentation for [` + "`" + `zpa_ba_certificate` + "`" + `](https://github.com/zscaler/terraform-provider-zpa/blob/master/docs/data-sources/zpa_ba_certificate.md)`,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "use_in_dr_mode",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_incomplete_dr_config",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ` ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_application_segment_inspection",
			Category:         "Application Segment",
			ShortDescription: `Creates and manages ZPA Application Segment for Inspection.`,
			Description: `

The **zpa_application_segment_inspection** resource creates an inspection application segment in the Zscaler Private Access cloud. This resource can then be referenced in an access policy inspection rule. This resource supports Inspection for both ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + `.

`,
			Keywords: []string{
				"application",
				"segment",
				"inspection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name. The name of the App Connector Group to be exported.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Required) List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `(Required) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(Required) List of Segment Group IDs`,
				},
				resource.Attribute{
					Name:        "common_apps_dto",
					Description: `(Required) List of applications (e.g., Inspection, Browser Access or Privileged Remote Access)`,
				},
				resource.Attribute{
					Name:        "apps_config:",
					Description: `(Required) List of applications to be configured`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Inspection Application Segment.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain name of the Inspection Application Segment.`,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(Required) Protocol for the Inspection Application Segment.. Supported values: ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(Required) Port for the Inspection Application Segment.`,
				},
				resource.Attribute{
					Name:        "app_types",
					Description: `(Required) Indicates the type of application as inspection. Supported value: ` + "`" + `INSPECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(string) - ID of the signing certificate. This field is required if the applicationProtocol is set to ` + "`" + `HTTPS` + "`" + `. The certificateId is not supported if the applicationProtocol is set to ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this application is enabled or not`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(Required) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(Required) UDP port ranges used to access the app. ->`,
				},
				resource.Attribute{
					Name:        "tcp_port_range",
					Description: `(Required) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_range",
					Description: `(Required) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(Optional) Indicates whether users can bypass ZPA to access applications.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(Optional) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this application is enabled or not.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(Optional) Whether health reporting for the app is Continuous or On Access. Supported values: ` + "`" + `NONE` + "`" + `, ` + "`" + `ON_ACCESS` + "`" + `, ` + "`" + `CONTINUOUS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "icmp_access_type",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "use_in_dr_mode",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_incomplete_dr_config",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ` ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Inspection Application Segment can be imported by using ` + "`" + `<APPLICATION SEGMENT ID>` + "`" + ` or ` + "`" + `<APPLICATION SEGMENT NAME>` + "`" + ` as the import ID. ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment_inspection.example <application_segment_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment_inspection.example <application_segment_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(Optional) Indicates whether users can bypass ZPA to access applications.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(Optional) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this application is enabled or not.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(Optional) Whether health reporting for the app is Continuous or On Access. Supported values: ` + "`" + `NONE` + "`" + `, ` + "`" + `ON_ACCESS` + "`" + `, ` + "`" + `CONTINUOUS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "icmp_access_type",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "use_in_dr_mode",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_incomplete_dr_config",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ` ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Inspection Application Segment can be imported by using ` + "`" + `<APPLICATION SEGMENT ID>` + "`" + ` or ` + "`" + `<APPLICATION SEGMENT NAME>` + "`" + ` as the import ID. ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment_inspection.example <application_segment_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment_inspection.example <application_segment_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_application_segment_pra",
			Category:         "Application Segment",
			ShortDescription: `Creates and manages ZPA Application Segment for Privileged Remote Access.`,
			Description: `

The **zpa_application_segment_pra** resource creates an application segment for Privileged Remote Access in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule, access policy timeout rule, access policy client forwarding rule and inspection policy. This resource supports Privileged Remote Access for both ` + "`" + `RDP` + "`" + ` and ` + "`" + `SSH` + "`" + `.

`,
			Keywords: []string{
				"application",
				"segment",
				"pra",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name. The name of the App Connector Group to be exported.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Required) List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `(Required) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(Required) List of Segment Group IDs`,
				},
				resource.Attribute{
					Name:        "common_apps_dto",
					Description: `(Required) List of applications (e.g., Inspection, Browser Access or Privileged Remote Access)`,
				},
				resource.Attribute{
					Name:        "apps_config:",
					Description: `(Required) List of applications to be configured`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Privileged Remote Access`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain name of the Privileged Remote Access`,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(Required) Protocol for the Privileged Remote Access. Supported values: ` + "`" + `RDP` + "`" + ` and ` + "`" + `SSH` + "`" + ``,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(Required) Port for the Privileged Remote Access`,
				},
				resource.Attribute{
					Name:        "app_types",
					Description: `(Required) Indicates the type of application as Privileged Remote Access. Supported value: ` + "`" + `SECURE_REMOTE_ACCESS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "connection_security",
					Description: `(Required) - Parameter required when ` + "`" + `application_protocol` + "`" + ` is of type ` + "`" + `RDP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this application is enabled or not`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(Required) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(Required) UDP port ranges used to access the app. ->`,
				},
				resource.Attribute{
					Name:        "tcp_port_range",
					Description: `(Required) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_range",
					Description: `(Required) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(Optional) Indicates whether users can bypass ZPA to access applications.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(Optional) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this application is enabled or not.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(Optional) Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.`,
				},
				resource.Attribute{
					Name:        "icmp_access_type",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "use_in_dr_mode",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_incomplete_dr_config",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ` ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Application Segment can be imported by using ` + "`" + `<APPLICATION SEGMENT ID>` + "`" + ` or ` + "`" + `<APPLICATION SEGMENT NAME>` + "`" + ` as the import ID. ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment_pra.example <application_segment_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment_pra.example <application_segment_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the application.`,
				},
				resource.Attribute{
					Name:        "bypass_type",
					Description: `(Optional) Indicates whether users can bypass ZPA to access applications.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "double_encrypt",
					Description: `(Optional) Whether Double Encryption is enabled or disabled for the app.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this application is enabled or not.`,
				},
				resource.Attribute{
					Name:        "health_reporting",
					Description: `(Optional) Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.`,
				},
				resource.Attribute{
					Name:        "icmp_access_type",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ip_anchored",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "is_cname_enabled",
					Description: `(Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.`,
				},
				resource.Attribute{
					Name:        "passive_health_enabled",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "use_in_dr_mode",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_incomplete_dr_config",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "select_connector_close_to_app",
					Description: `(Optional) Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ` ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Application Segment can be imported by using ` + "`" + `<APPLICATION SEGMENT ID>` + "`" + ` or ` + "`" + `<APPLICATION SEGMENT NAME>` + "`" + ` as the import ID. ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment_pra.example <application_segment_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_segment_pra.example <application_segment_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_application_server",
			Category:         "Application Server",
			ShortDescription: `Creates and manages ZPA Application Servers.`,
			Description: ` (Resource)

The **zpa_application_server** resource creates an application server in the Zscaler Private Access cloud. This resource can then be referenced in a server group.

`,
			Keywords: []string{
				"application",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name. The name of the application server to be exported.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Address. The address of the application server to be exported. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "app_server_group_ids",
					Description: `(Optional) This field defines the list of server group IDs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) This field defines the description of the server.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) This field defines the status of the server.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(Optional) ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Application Server can be imported by using ` + "`" + `<APPLICATION SERVER ID>` + "`" + ` or ` + "`" + `<APPLICATION SERVER NAME>` + "`" + ` as the import ID For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_server.example <application_server_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_server.example <application_server_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "app_server_group_ids",
					Description: `(Optional) This field defines the list of server group IDs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) This field defines the description of the server.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) This field defines the status of the server.`,
				},
				resource.Attribute{
					Name:        "config_space",
					Description: `(Optional) ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Application Server can be imported by using ` + "`" + `<APPLICATION SERVER ID>` + "`" + ` or ` + "`" + `<APPLICATION SERVER NAME>` + "`" + ` as the import ID For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_server.example <application_server_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_application_server.example <application_server_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_browser_access",
			Category:         "Application Segment",
			ShortDescription: `Creates and manages ZPA Browser Access Application Segment.`,
			Description: `

The **zpa_browser_access** creates and manages a browser access application segment resource in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule, access policy timeout rule or access policy client forwarding rule.

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
					Description: `(Required) Name of the application.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Required) List of domains and IPs.`,
				},
				resource.Attribute{
					Name:        "tcp_port_ranges",
					Description: `(Required) TCP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "udp_port_ranges",
					Description: `(Required) UDP port ranges used to access the app.`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `(Required) List of Server Group IDs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "segment_group_id",
					Description: `(Required) List of Segment Group IDs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Name of BA app.`,
				},
				resource.Attribute{
					Name:        "application_port",
					Description: `(Required) - Port for the BA app.`,
				},
				resource.Attribute{
					Name:        "application_protocol",
					Description: `(Required) - Protocol for the BA app. Supported values: ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Required) - ID of the BA certificate. Refer to the data source documentation for [` + "`" + `zpa_ba_certificate` + "`" + `](https://github.com/zscaler/terraform-provider-zpa/blob/master/docs/data-sources/zpa_ba_certificate.md)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) - Domain name or IP address of the BA app.`,
				},
				resource.Attribute{
					Name:        "allow_options",
					Description: `(Optional) - If you want ZPA to forward unauthenticated HTTP preflight OPTIONS requests from the browser to the app.. Supported values: ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + ` ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(Optional) - Name of the BA certificate. Refer to the data source documentation for [` + "`" + `zpa_ba_certificate` + "`" + `](https://github.com/zscaler/terraform-provider-zpa/blob/master/docs/data-sources/zpa_ba_certificate.md) ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(Optional) - Name of the BA certificate. Refer to the data source documentation for [` + "`" + `zpa_ba_certificate` + "`" + `](https://github.com/zscaler/terraform-provider-zpa/blob/master/docs/data-sources/zpa_ba_certificate.md) ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_inspection_custom_control",
			Category:         "Inspection",
			ShortDescription: `Creates and manages Inspection Custom Control in Zscaler Private Access cloud.`,
			Description: `

The **zpa_inspection_custom_controls** resource creates an inspection custom control. This resource can then be referenced in an inspection profile resource.

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
					Description: `(Required) The version of the predefined control, the default is: ` + "`" + `OWASP_CRS/3.3.0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The performed action. Supported values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(Required) Denotes the action`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the custom control`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(Required) OWASP Predefined Paranoia Level.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(string) Returned values: ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `, ` + "`" + `FTP` + "`" + `, ` + "`" + `RDP` + "`" + `, ` + "`" + `SSH` + "`" + `, ` + "`" + `WEBSOCKET` + "`" + ``,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Required) Severity of the control number. Supported values: ` + "`" + `CRITICAL` + "`" + `, ` + "`" + `ERROR` + "`" + `, ` + "`" + `WARNING` + "`" + `, ` + "`" + `INFO` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Rules to be applied to the request or response type`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Required) Rules of the custom controls applied as conditions ` + "`" + `JSON` + "`" + ``,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "lhs",
					Description: `(Required) Signifies the key for the object type Supported values: ` + "`" + `SIZE` + "`" + `, ` + "`" + `VALUE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(Required) If lhs is set to SIZE, then the user may pass one of the following: ` + "`" + `EQ: Equals` + "`" + `, ` + "`" + `LE: Less than or equal to` + "`" + `, ` + "`" + `GE: Greater than or equal to` + "`" + `. If the lhs is set to ` + "`" + `VALUE` + "`" + `, then the user may pass one of the following: ` + "`" + `CONTAINS` + "`" + `, ` + "`" + `STARTS_WITH` + "`" + `, ` + "`" + `ENDS_WITH` + "`" + `, ` + "`" + `RX` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rhs",
					Description: `(Required) Denotes the value for the given object type. Its value depends on the key. If rules.type is set to REQUEST_METHOD, the conditions.rhs field must have one of the following values: ` + "`" + `GET` + "`" + `,` + "`" + `HEAD` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `OPTIONS` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `TRACE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Required) Name of the rules. If rules.type is set to ` + "`" + `REQUEST_HEADERS` + "`" + `, ` + "`" + `REQUEST_COOKIES` + "`" + `, or ` + "`" + `RESPONSE_HEADERS` + "`" + `, the rules.name field is required.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type value for the rules ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the custom control`,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(Optional) Name of the inspection profile`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(string) Returned values: ` + "`" + `WEBSOCKET_PREDEFINED` + "`" + `, ` + "`" + `WEBSOCKET_CUSTOM` + "`" + `, ` + "`" + `ZSCALER` + "`" + `, ` + "`" + `CUSTOM` + "`" + `, ` + "`" + `PREDEFINED` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(Required) The performed action. Supported values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_action_value",
					Description: `(Optional) This is used to provide the redirect URL if the default action is set to ` + "`" + `REDIRECT` + "`" + ` ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the custom control`,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(Optional) Name of the inspection profile`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(string) Returned values: ` + "`" + `WEBSOCKET_PREDEFINED` + "`" + `, ` + "`" + `WEBSOCKET_CUSTOM` + "`" + `, ` + "`" + `ZSCALER` + "`" + `, ` + "`" + `CUSTOM` + "`" + `, ` + "`" + `PREDEFINED` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(Required) The performed action. Supported values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_action_value",
					Description: `(Optional) This is used to provide the redirect URL if the default action is set to ` + "`" + `REDIRECT` + "`" + ` ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_inspection_profile",
			Category:         "Inspection",
			ShortDescription: `Creates and manages Inspection Profile in Zscaler Private Access cloud.`,
			Description: `

The  **zpa_inspection_profile** resource creates an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.

`,
			Keywords: []string{
				"inspection",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the inspection profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the inspection profile.`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(Required) OWASP Predefined Paranoia Level. Range: [1-4], inclusive`,
				},
				resource.Attribute{
					Name:        "predefined_controls",
					Description: `(Required) The predefined controls. The default predefined control ` + "`" + `Preprocessors` + "`" + ` are mandatory and injected in the request by default. Individual ` + "`" + `predefined_controls` + "`" + ` can be set by using the data source ` + "`" + `data_source_zpa_predefined_controls` + "`" + ` or by group using the data source ` + "`" + `zpa_inspection_all_predefined_controls` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the predefined control`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action of the predefined control. Supported values: ` + "`" + `PASS` + "`" + `, ` + "`" + `BLOCK` + "`" + ` and ` + "`" + `REDIRECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `(Required) Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to ` + "`" + `REDIRECT` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "custom_controls",
					Description: `(Optional) Types for custom controls`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional) Rules of the custom controls applied as conditions ` + "`" + `JSON` + "`" + ``,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "lhs",
					Description: `(Optional) Signifies the key for the object type Supported values: ` + "`" + `SIZE` + "`" + `, ` + "`" + `VALUE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(Optional) If lhs is set to SIZE, then the user may pass one of the following: ` + "`" + `EQ: Equals` + "`" + `, ` + "`" + `LE: Less than or equal to` + "`" + `, ` + "`" + `GE: Greater than or equal to` + "`" + `. If the lhs is set to ` + "`" + `VALUE` + "`" + `, then the user may pass one of the following: ` + "`" + `CONTAINS` + "`" + `, ` + "`" + `STARTS_WITH` + "`" + `, ` + "`" + `ENDS_WITH` + "`" + `, ` + "`" + `RX` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rhs",
					Description: `(Optional) Denotes the value for the given object type. Its value depends on the key. If rules.type is set to REQUEST_METHOD, the conditions.rhs field must have one of the following values: ` + "`" + `GET` + "`" + `,` + "`" + `HEAD` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `OPTIONS` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `TRACE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) Name of the rules. If rules.type is set to ` + "`" + `REQUEST_HEADERS` + "`" + `, ` + "`" + `REQUEST_COOKIES` + "`" + `, or ` + "`" + `RESPONSE_HEADERS` + "`" + `, the rules.name field is required.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type value for the rules`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The version of the predefined control, the default is: ` + "`" + `OWASP_CRS/3.3.0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(Optional) Name of the inspection profile`,
				},
				resource.Attribute{
					Name:        "common_global_override_actions_config",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "controls_info",
					Description: `(Optional) Types for custom controls`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(string) Control types. Supported Values: ` + "`" + `WEBSOCKET_PREDEFINED` + "`" + `, ` + "`" + `WEBSOCKET_CUSTOM` + "`" + `, ` + "`" + `CUSTOM` + "`" + `, ` + "`" + `PREDEFINED` + "`" + `, ` + "`" + `ZSCALER` + "`" + ``,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional) Control information counts ` + "`" + `Long` + "`" + ``,
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
					Description: `(string) Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to ` + "`" + `REDIRECT` + "`" + `. ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "custom_controls",
					Description: `(Optional) Types for custom controls`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional) Rules of the custom controls applied as conditions ` + "`" + `JSON` + "`" + ``,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "lhs",
					Description: `(Optional) Signifies the key for the object type Supported values: ` + "`" + `SIZE` + "`" + `, ` + "`" + `VALUE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(Optional) If lhs is set to SIZE, then the user may pass one of the following: ` + "`" + `EQ: Equals` + "`" + `, ` + "`" + `LE: Less than or equal to` + "`" + `, ` + "`" + `GE: Greater than or equal to` + "`" + `. If the lhs is set to ` + "`" + `VALUE` + "`" + `, then the user may pass one of the following: ` + "`" + `CONTAINS` + "`" + `, ` + "`" + `STARTS_WITH` + "`" + `, ` + "`" + `ENDS_WITH` + "`" + `, ` + "`" + `RX` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rhs",
					Description: `(Optional) Denotes the value for the given object type. Its value depends on the key. If rules.type is set to REQUEST_METHOD, the conditions.rhs field must have one of the following values: ` + "`" + `GET` + "`" + `,` + "`" + `HEAD` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `OPTIONS` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `TRACE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) Name of the rules. If rules.type is set to ` + "`" + `REQUEST_HEADERS` + "`" + `, ` + "`" + `REQUEST_COOKIES` + "`" + `, or ` + "`" + `RESPONSE_HEADERS` + "`" + `, the rules.name field is required.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type value for the rules`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The version of the predefined control, the default is: ` + "`" + `OWASP_CRS/3.3.0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "associated_inspection_profile_names",
					Description: `(Optional) Name of the inspection profile`,
				},
				resource.Attribute{
					Name:        "common_global_override_actions_config",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "controls_info",
					Description: `(Optional) Types for custom controls`,
				},
				resource.Attribute{
					Name:        "control_type",
					Description: `(string) Control types. Supported Values: ` + "`" + `WEBSOCKET_PREDEFINED` + "`" + `, ` + "`" + `WEBSOCKET_CUSTOM` + "`" + `, ` + "`" + `CUSTOM` + "`" + `, ` + "`" + `PREDEFINED` + "`" + `, ` + "`" + `ZSCALER` + "`" + ``,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional) Control information counts ` + "`" + `Long` + "`" + ``,
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
					Description: `(string) Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to ` + "`" + `REDIRECT` + "`" + `. ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_lss_config_controller",
			Category:         "Log Streaming (LSS)",
			ShortDescription: `Creates and manages ZPA LSS Configuration.`,
			Description: `

The **zpa_lss_config_controller** resource creates and manages Log Streaming Service (LSS) in the Zscaler Private Access cloud.

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
					Name:        "config",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Required) The format of the LSS resource. The supported formats are: ` + "`" + `JSON` + "`" + `, ` + "`" + `CSV` + "`" + `, and ` + "`" + `TSV` + "`" + ``,
				},
				resource.Attribute{
					Name:        "lss_host",
					Description: `(Required) The IP or FQDN of the SIEM (Log Receiver) where logs will be forwarded to.`,
				},
				resource.Attribute{
					Name:        "lss_port",
					Description: `(Required) The destination port of the SIEM (Log Receiver) where logs will be forwarded to.`,
				},
				resource.Attribute{
					Name:        "source_log_type",
					Description: `(Required) Refer to the log type documentation`,
				},
				resource.Attribute{
					Name:        "connector_groups",
					Description: `(Required) - ` + "`" + `id` + "`" + ` - (Required) - App Connector Group ID(s) where logs will be forwarded to. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "use_tls",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "source_log_type",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "connector_groups",
					Description: `(Required) - ` + "`" + `id` + "`" + ` - (Required) - App Connector Group ID(s) where logs will be forwarded to.`,
				},
				resource.Attribute{
					Name:        "policy_rule_resource",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "audit_message",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "custom_msg",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "connector_groups",
					Description: `(Optional) - ` + "`" + `id` + "`" + ` - (Optional) - App Connector Group ID(s) where logs will be forwarded to.`,
				},
				resource.Attribute{
					Name:        "app_server_groups",
					Description: `(Optional) - ` + "`" + `id` + "`" + ` - (Optional) - Server Group ID(s).`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "use_tls",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "source_log_type",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "connector_groups",
					Description: `(Required) - ` + "`" + `id` + "`" + ` - (Required) - App Connector Group ID(s) where logs will be forwarded to.`,
				},
				resource.Attribute{
					Name:        "policy_rule_resource",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "audit_message",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "custom_msg",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "connector_groups",
					Description: `(Optional) - ` + "`" + `id` + "`" + ` - (Optional) - App Connector Group ID(s) where logs will be forwarded to.`,
				},
				resource.Attribute{
					Name:        "app_server_groups",
					Description: `(Optional) - ` + "`" + `id` + "`" + ` - (Optional) - Server Group ID(s).`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_forwarding_rule",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Access Forwarding Rule.`,
			Description: `

The **zpa_policy_forwarding_rule** resource creates a policy forwarding access rule in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"forwarding",
				"rule",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) This is for providing the rule action.`,
				},
				resource.Attribute{
					Name:        "custom_msg",
					Description: `(Optional) This is for providing a customer message for the user.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) This is the description of the access policy rule.`,
				},
				resource.Attribute{
					Name:        "rule_order",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_inspection_rule",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Access Inspection Rule.`,
			Description: `

The **zpa_policy_inspection_rule** resource creates a policy inspection access rule in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"inspection",
				"rule",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) This is for providing the rule action.`,
				},
				resource.Attribute{
					Name:        "action_id",
					Description: `(Optional) This field defines the description of the server.`,
				},
				resource.Attribute{
					Name:        "bypass_default_rule",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "custom_msg",
					Description: `(Optional) This is for providing a customer message for the user.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) This is the description of the access policy rule.`,
				},
				resource.Attribute{
					Name:        "rule_order",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_isolation_rule",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Access Isolation Rule.`,
			Description: `

The **zpa_policy_isolation_rule** resource creates a policy isolation access rule in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"isolation",
				"rule",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) This is for providing the rule action.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) This is the description of the access policy rule.`,
				},
				resource.Attribute{
					Name:        "rule_order",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_rule",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Access Rule.`,
			Description: `

The **zpa_policy_access_rule** resource creates and manages policy access rule in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"rule",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of an app connector group resource`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of a server group resource ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Policy access rule can be imported by using ` + "`" + `<POLICY ACCESS RULE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_policy_access_rule.example <policy_access_rule_id> ` + "`" + `` + "`" + `` + "`" + ` ## LHS and RHS Values | Object Type | LHS| RHS |----------|-----------|---------- | [APP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment) | "id" | <application_segment_ID> | | [APP_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_segment_group) | "id" | <segment_group_ID> | | [CLIENT_TYPE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment_browser_access) | "id" | zpn_client_type_zappl or zpn_client_type_exporter | | [EDGE_CONNECTOR_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_cloud_connector_group) | "id" | <edge_connector_ID> | | [IDP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_idp_controller) | "id" | <identity_provider_ID> | | [MACHINE_GRP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_machine_group) | "id" | <machine_group_ID> | | [POSTURE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_posture_profile) | <posture_udid> | "true" / "false" | | [TRUSTED_NETWORK](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_trusted_network) | <network_id> | "true" | | [SAML](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_saml_attribute) | <saml_attribute_id> | <Attribute_value_to_match> | | [SCIM](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_attribute_header) | <scim_attribute_id> | <Attribute_value_to_match> | | [SCIM_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_groups) | <scim_group_attribute_id> | <Attribute_value_to_match> |`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_rule_application_segment",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Access Rule with Application Segment conditions.`,
			Description: `

The **zpa_policy_access_rule** resource creates and manages a policy access rule with Application Segment conditions in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"rule",
				"application",
				"segment",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of an app connector group resource`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of a server group resource ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Policy Access Rule for Browser Access can be imported by using` + "`" + `<POLICY ACCESS RULE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_policy_access_rule.example <policy_access_rule_id> ` + "`" + `` + "`" + `` + "`" + ` ## LHS and RHS Values LHS and RHS values differ based on object types. Refer to the following table: | Object Type | LHS| RHS |----------|-----------|---------- | [APP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment) | "id" | <application_segment_ID> | | [APP_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_segment_group) | "id" | <segment_group_ID> | | [CLIENT_TYPE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment_browser_access) | "id" | zpn_client_type_zappl or zpn_client_type_exporter | | [EDGE_CONNECTOR_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_cloud_connector_group) | "id" | <edge_connector_ID> | | [IDP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_idp_controller) | "id" | <identity_provider_ID> | | [MACHINE_GRP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_machine_group) | "id" | <machine_group_ID> | | [POSTURE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_posture_profile) | <posture_udid> | "true" / "false" | | [TRUSTED_NETWORK](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_trusted_network) | <network_id> | "true" | | [SAML](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_saml_attribute) | <saml_attribute_id> | <Attribute_value_to_match> | | [SCIM](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_attribute_header) | <scim_attribute_id> | <Attribute_value_to_match> | | [SCIM_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_groups) | <scim_group_attribute_id> | <Attribute_value_to_match> |`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_rule_browser_access",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Access Rule for Browser Access.`,
			Description: `

The **zpa_policy_access_rule** resource creates and manages policy access rule to support Browser Access  in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"rule",
				"browser",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "zpn_client_type_exporter",
					Description: `Used to support/enforce access via browser access tp application segments.`,
				},
				resource.Attribute{
					Name:        "zpn_client_type_zappl",
					Description: `Used to support/enforce access via Zscaler Client Connector to application segments.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) A list of IDs of an app connector group resources`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) A list of IDs of a server group resources ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Policy Access Rule for Browser Access can be imported by using` + "`" + `<POLICY ACCESS RULE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_policy_access_rule.example <policy_access_rule_id> ` + "`" + `` + "`" + `` + "`" + ` ## LHS and RHS Values LHS and RHS values differ based on object types. Refer to the following table: | Object Type | LHS| RHS |----------|-----------|---------- | [APP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment) | "id" | <application_segment_ID> | | [APP_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_segment_group) | "id" | <segment_group_ID> | | [CLIENT_TYPE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment_browser_access) | "id" | zpn_client_type_zappl or zpn_client_type_exporter | | [EDGE_CONNECTOR_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_cloud_connector_group) | "id" | <edge_connector_ID> | | [IDP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_idp_controller) | "id" | <identity_provider_ID> | | [MACHINE_GRP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_machine_group) | "id" | <machine_group_ID> | | [POSTURE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_posture_profile) | <posture_udid> | "true" / "false" | | [TRUSTED_NETWORK](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_trusted_network) | <network_id> | "true" | | [SAML](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_saml_attribute) | <saml_attribute_id> | <Attribute_value_to_match> | | [SCIM](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_attribute_header) | <scim_attribute_id> | <Attribute_value_to_match> | | [SCIM_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_groups) | <scim_group_attribute_id> | <Attribute_value_to_match> |`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_rule_posture_profile",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Access Rule with Posture Profile conditions.`,
			Description: `

The **zpa_policy_access_rule** resource creates and manages a policy access rule with Posture Profile conditions in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"rule",
				"posture",
				"profile",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of an app connector group resource`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of a server group resource ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Policy Access Rule for Browser Access can be imported by using` + "`" + `<POLICY ACCESS RULE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_policy_access_rule.example <policy_access_rule_id> ` + "`" + `` + "`" + `` + "`" + ` ## LHS and RHS Values LHS and RHS values differ based on object types. Refer to the following table: | Object Type | LHS| RHS |----------|-----------|---------- | [APP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment) | "id" | <application_segment_ID> | | [APP_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_segment_group) | "id" | <segment_group_ID> | | [CLIENT_TYPE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment_browser_access) | "id" | zpn_client_type_zappl or zpn_client_type_exporter | | [EDGE_CONNECTOR_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_cloud_connector_group) | "id" | <edge_connector_ID> | | [IDP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_idp_controller) | "id" | <identity_provider_ID> | | [MACHINE_GRP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_machine_group) | "id" | <machine_group_ID> | | [POSTURE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_posture_profile) | <posture_udid> | "true" / "false" | | [TRUSTED_NETWORK](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_trusted_network) | <network_id> | "true" | | [SAML](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_saml_attribute) | <saml_attribute_id> | <Attribute_value_to_match> | | [SCIM](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_attribute_header) | <scim_attribute_id> | <Attribute_value_to_match> | | [SCIM_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_groups) | <scim_group_attribute_id> | <Attribute_value_to_match> |`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_rule_saml",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Access Rule with SAML Attribute conditions.`,
			Description: `

The **zpa_policy_access_rule** resource creates and manages a policy access rule with SAML attribute conditions in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"rule",
				"saml",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of an app connector group resource`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of a server group resource ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Policy Access Rule for Browser Access can be imported by using` + "`" + `<POLICY ACCESS RULE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_policy_access_rule.example <policy_access_rule_id> ` + "`" + `` + "`" + `` + "`" + ` ## LHS and RHS Values LHS and RHS values differ based on object types. Refer to the following table: | Object Type | LHS| RHS |----------|-----------|---------- | [APP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment) | "id" | <application_segment_ID> | | [APP_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_segment_group) | "id" | <segment_group_ID> | | [CLIENT_TYPE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment_browser_access) | "id" | zpn_client_type_zappl or zpn_client_type_exporter | | [EDGE_CONNECTOR_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_cloud_connector_group) | "id" | <edge_connector_ID> | | [IDP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_idp_controller) | "id" | <identity_provider_ID> | | [MACHINE_GRP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_machine_group) | "id" | <machine_group_ID> | | [POSTURE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_posture_profile) | <posture_udid> | "true" / "false" | | [TRUSTED_NETWORK](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_trusted_network) | <network_id> | "true" | | [SAML](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_saml_attribute) | <saml_attribute_id> | <Attribute_value_to_match> | | [SCIM](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_attribute_header) | <scim_attribute_id> | <Attribute_value_to_match> | | [SCIM_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_groups) | <scim_group_attribute_id> | <Attribute_value_to_match> |`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_rule_scim_attribute",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Access Rule with SCIM Attribute Header conditions.`,
			Description: `

The **zpa_policy_access_rule** resource creates and manages a policy access rule with SCIM attribute header conditions in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"rule",
				"scim",
				"attribute",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of an app connector group resource`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of a server group resource ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Policy Access Rule for Browser Access can be imported by using` + "`" + `<POLICY ACCESS RULE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_policy_access_rule.example <policy_access_rule_id> ` + "`" + `` + "`" + `` + "`" + ` ## LHS and RHS Values LHS and RHS values differ based on object types. Refer to the following table: | Object Type | LHS| RHS |----------|-----------|---------- | [APP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment) | "id" | <application_segment_ID> | | [APP_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_segment_group) | "id" | <segment_group_ID> | | [CLIENT_TYPE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment_browser_access) | "id" | zpn_client_type_zappl or zpn_client_type_exporter | | [EDGE_CONNECTOR_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_cloud_connector_group) | "id" | <edge_connector_ID> | | [IDP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_idp_controller) | "id" | <identity_provider_ID> | | [MACHINE_GRP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_machine_group) | "id" | <machine_group_ID> | | [POSTURE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_posture_profile) | <posture_udid> | "true" / "false" | | [TRUSTED_NETWORK](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_trusted_network) | <network_id> | "true" | | [SAML](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_saml_attribute) | <saml_attribute_id> | <Attribute_value_to_match> | | [SCIM](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_attribute_header) | <scim_attribute_id> | <Attribute_value_to_match> | | [SCIM_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_groups) | <scim_group_attribute_id> | <Attribute_value_to_match> |`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_rule_scim_group",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Access Rule with SCIM Group conditions.`,
			Description: `

The **zpa_policy_access_rule** resource creates and manages a policy access rule with SCIM Group conditions in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"rule",
				"scim",
				"group",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of an app connector group resource`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of a server group resource ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Policy Access Rule for Browser Access can be imported by using` + "`" + `<POLICY ACCESS RULE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_policy_access_rule.example <policy_access_rule_id> ` + "`" + `` + "`" + `` + "`" + ` ## LHS and RHS Values LHS and RHS values differ based on object types. Refer to the following table: | Object Type | LHS| RHS |----------|-----------|---------- | [APP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment) | "id" | <application_segment_ID> | | [APP_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_segment_group) | "id" | <segment_group_ID> | | [CLIENT_TYPE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment_browser_access) | "id" | zpn_client_type_zappl or zpn_client_type_exporter | | [EDGE_CONNECTOR_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_cloud_connector_group) | "id" | <edge_connector_ID> | | [IDP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_idp_controller) | "id" | <identity_provider_ID> | | [MACHINE_GRP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_machine_group) | "id" | <machine_group_ID> | | [POSTURE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_posture_profile) | <posture_udid> | "true" / "false" | | [TRUSTED_NETWORK](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_trusted_network) | <network_id> | "true" | | [SAML](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_saml_attribute) | <saml_attribute_id> | <Attribute_value_to_match> | | [SCIM](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_attribute_header) | <scim_attribute_id> | <Attribute_value_to_match> | | [SCIM_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_groups) | <scim_group_attribute_id> | <Attribute_value_to_match> |`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_rule_trusted_networks",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Access Rule with Trusted Networks conditions.`,
			Description: `

The **zpa_policy_access_rule** resource creates and manages a policy access rule with Trusted Networks conditions in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"rule",
				"trusted",
				"networks",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of an app connector group resource`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of a server group resource ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Policy Access Rule for Browser Access can be imported by using` + "`" + `<POLICY ACCESS RULE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_policy_access_rule.example <policy_access_rule_id> ` + "`" + `` + "`" + `` + "`" + ` ## LHS and RHS Values LHS and RHS values differ based on object types. Refer to the following table: | Object Type | LHS| RHS |----------|-----------|---------- | [APP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment) | "id" | <application_segment_ID> | | [APP_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_segment_group) | "id" | <segment_group_ID> | | [CLIENT_TYPE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/resources/zpa_application_segment_browser_access) | "id" | zpn_client_type_zappl or zpn_client_type_exporter | | [EDGE_CONNECTOR_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_cloud_connector_group) | "id" | <edge_connector_ID> | | [IDP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_idp_controller) | "id" | <identity_provider_ID> | | [MACHINE_GRP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_machine_group) | "id" | <machine_group_ID> | | [POSTURE](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_posture_profile) | <posture_udid> | "true" / "false" | | [TRUSTED_NETWORK](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_trusted_network) | <network_id> | "true" | | [SAML](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_saml_attribute) | <saml_attribute_id> | <Attribute_value_to_match> | | [SCIM](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_attribute_header) | <scim_attribute_id> | <Attribute_value_to_match> | | [SCIM_GROUP](https://registry.terraform.io/providers/zscaler/zpa/latest/docs/data-sources/zpa_scim_groups) | <scim_group_attribute_id> | <Attribute_value_to_match> |`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_policy_access_timeout_rule",
			Category:         "Policy Set Controller",
			ShortDescription: `Creates and manages ZPA Policy Timeout Access Rule.`,
			Description: `

The **zpa_policy_timeout_rule** resource creates a policy timeout rule in the Zscaler Private Access cloud.

`,
			Keywords: []string{
				"policy",
				"set",
				"controller",
				"access",
				"timeout",
				"rule",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Supported values: ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_provisioning_key",
			Category:         "Provisioning Key",
			ShortDescription: `Creates and manages ZPA Provisioning Key for Service Edge and/or App Connector Groups.`,
			Description: `

The **zpa_provisioning_key** resource provides creates a provisioning key in the Zscaler Private Access portal. This resource can then be referenced in the following ZPA resources:

1. App Connector Groups
2. Service Edge Groups

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
					Name:        "max_usage",
					Description: `(Required) The maximum number of instances where this provisioning key can be used for enrolling an App Connector or Service Edge.`,
				},
				resource.Attribute{
					Name:        "enrollment_cert_id",
					Description: `(Required) ID of the enrollment certificate that can be used for this provisioning key. ` + "`" + `ID` + "`" + ` of the existing enrollment certificate that has the private key`,
				},
				resource.Attribute{
					Name:        "zcomponent_id",
					Description: `(Required) ID of the existing App Connector or Service Edge Group.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_segment_group",
			Category:         "Segment Group",
			ShortDescription: `Creates and manages ZPA Segment Group resource`,
			Description: `

The **zpa_segment_group** resource creates a segment group in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule or application segment resource.

`,
			Keywords: []string{
				"segment",
				"group",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the segment group. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_server_group",
			Category:         "Server Group",
			ShortDescription: `Creates and manages ZPA Server Group resource`,
			Description: `

The **zpa_server_group** resource creates a server group in the Zscaler Private Access cloud. This resource can then be referenced in an application segment or application server resource.

`,
			Keywords: []string{
				"server",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zpa_zpa_service_edge_group",
			Category:         "Service Edge Group",
			ShortDescription: `Creates and manages ZPA Service Edge Group details.`,
			Description: `

The **zpa_service_edge_group** resource creates a service edge group in the Zscaler Private Access cloud. This resource can then be referenced in a service edge connector.

`,
			Keywords: []string{
				"service",
				"edge",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Service Edge Group.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Required) Latitude for the Service Edge Group. Integer or decimal with values in the range of ` + "`" + `-90` + "`" + ` to ` + "`" + `90` + "`" + ``,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Required) Longitude for the Service Edge Group. Integer or decimal with values in the range of ` + "`" + `-180` + "`" + ` to ` + "`" + `180` + "`" + ``,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Location for the Service Edge Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Service Edge Group.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this Service Edge Group is enabled or not. Default value: ` + "`" + `true` + "`" + ` Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "city_country",
					Description: `(Optional) This field controls dynamic discovery of the servers.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(Optional) This field is an array of app-connector-id only.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) Enable or disable public access for the Service Edge Group. Default value: ` + "`" + `false` + "`" + ` Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "override_version_profile",
					Description: `(Optional) Whether the default version profile of the App Connector Group is applied or overridden. Default: ` + "`" + `false` + "`" + ` Supported values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "version_profile_id",
					Description: `(Optional) ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:`,
				},
				resource.Attribute{
					Name:        "version_profile_name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "service_edges",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "trusted_networks",
					Description: `(Optional) Trusted networks for this Service Edge Group. List of trusted network objects`,
				},
				resource.Attribute{
					Name:        "upgrade_day",
					Description: `(Optional) Service Edges in this group will attempt to update to a newer version of the software during this specified day. Default value: ` + "`" + `SUNDAY` + "`" + ` List of valid days (i.e., Sunday, Monday)`,
				},
				resource.Attribute{
					Name:        "upgrade_time_in_secs",
					Description: `(Optional) Service Edges in this group will attempt to update to a newer version of the software during this specified time. Default value: ` + "`" + `66600` + "`" + ` Integer in seconds (i..e, 66600). The integer must be greater than or equal to 0 and less than ` + "`" + `86400` + "`" + `, in ` + "`" + `15` + "`" + ` minute intervals ## Import Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Service Edge Group can be imported; use ` + "`" + `<SERVER EDGE GROUP ID>` + "`" + ` or ` + "`" + `<SERVER EDGE GROUP NAME>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_service_edge_group.example <service_edge_group_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_service_edge_group.example <service_edge_group_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"zpa_zpa_app_connector_group":                    0,
		"zpa_zpa_application_segment":                    1,
		"zpa_zpa_application_segment_browser_access":     2,
		"zpa_zpa_application_segment_inspection":         3,
		"zpa_zpa_application_segment_pra":                4,
		"zpa_zpa_application_server":                     5,
		"zpa_zpa_browser_access":                         6,
		"zpa_zpa_inspection_custom_control":              7,
		"zpa_zpa_inspection_profile":                     8,
		"zpa_zpa_lss_config_controller":                  9,
		"zpa_zpa_policy_access_forwarding_rule":          10,
		"zpa_zpa_policy_access_inspection_rule":          11,
		"zpa_zpa_policy_access_isolation_rule":           12,
		"zpa_zpa_policy_access_rule":                     13,
		"zpa_zpa_policy_access_rule_application_segment": 14,
		"zpa_zpa_policy_access_rule_browser_access":      15,
		"zpa_zpa_policy_access_rule_posture_profile":     16,
		"zpa_zpa_policy_access_rule_saml":                17,
		"zpa_zpa_policy_access_rule_scim_attribute":      18,
		"zpa_zpa_policy_access_rule_scim_group":          19,
		"zpa_zpa_policy_access_rule_trusted_networks":    20,
		"zpa_zpa_policy_access_timeout_rule":             21,
		"zpa_zpa_provisioning_key":                       22,
		"zpa_zpa_segment_group":                          23,
		"zpa_zpa_server_group":                           24,
		"zpa_zpa_service_edge_group":                     25,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
