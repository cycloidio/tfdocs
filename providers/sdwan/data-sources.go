package sdwan

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "sdwan_device_template",
			Category:         "Data Sources",
			ShortDescription: `Data source for SD-WAN Device Templates`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Unique Device Template Name ## Common Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_description",
					Description: `Long Description of Device Template`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Type of device supported by Device Template`,
				},
				resource.Attribute{
					Name:        "device_role",
					Description: `Device role for the Device Template`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Configuration type for Device Template`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `Flag indicating whether Device Template is factory default or not`,
				},
				resource.Attribute{
					Name:        "template_class",
					Description: `Template Class type for Device Template`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Device Template id for Device Template ## Attribute Reference for Device Template created from Feature Templates ##`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `policyId for Device Template`,
				},
				resource.Attribute{
					Name:        "feature_template_uid_range",
					Description: `feature_template_uid_range for Device Template`,
				},
				resource.Attribute{
					Name:        "connection_preference_required",
					Description: `connectionPreferenceRequired flag for Device Template`,
				},
				resource.Attribute{
					Name:        "connection_preference",
					Description: `connectionPreference flag for Device Template`,
				},
				resource.Attribute{
					Name:        "general_templates",
					Description: `List of Feature Templates and thier Sub Templates thourgh which Device Template is created (see [below for nested schema](#nestedblock--general_templates)) <a id="nestedblock--general_templates"></a> ## Nested Schema for general_templates`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Template Id of Feature Template`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Template Type of Feature Template`,
				},
				resource.Attribute{
					Name:        "sub_templates",
					Description: `List of Sub Templates associated with feature Template (see [below for nested schema](#nestedblock--sub_templates)) <a id="nestedblock--sub_templates"></a> ## Nested Schema for sub_templates`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Template Id of Feature Template`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Template Type of Feature Template ## Attribute Reference for Device Template created from CLI ##`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who updated Device Template last`,
				},
				resource.Attribute{
					Name:        "template_configuration",
					Description: `Template Configuration for Device Template`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Time when Device Template was created`,
				},
				resource.Attribute{
					Name:        "rid",
					Description: `Request ID for Device Template`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `Feature for Device Template`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `User who created Device Template`,
				},
				resource.Attribute{
					Name:        "last_updated_on",
					Description: `Time when Device Template was last updated`,
				},
				resource.Attribute{
					Name:        "template_attached",
					Description: `Number of templates attached to Device Template`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `policyId for Device Template`,
				},
				resource.Attribute{
					Name:        "feature_template_uid_range",
					Description: `feature_template_uid_range for Device Template`,
				},
				resource.Attribute{
					Name:        "connection_preference_required",
					Description: `connectionPreferenceRequired flag for Device Template`,
				},
				resource.Attribute{
					Name:        "connection_preference",
					Description: `connectionPreference flag for Device Template`,
				},
				resource.Attribute{
					Name:        "general_templates",
					Description: `List of Feature Templates and thier Sub Templates thourgh which Device Template is created (see [below for nested schema](#nestedblock--general_templates)) <a id="nestedblock--general_templates"></a> ## Nested Schema for general_templates`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Template Id of Feature Template`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Template Type of Feature Template`,
				},
				resource.Attribute{
					Name:        "sub_templates",
					Description: `List of Sub Templates associated with feature Template (see [below for nested schema](#nestedblock--sub_templates)) <a id="nestedblock--sub_templates"></a> ## Nested Schema for sub_templates`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Template Id of Feature Template`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Template Type of Feature Template ## Attribute Reference for Device Template created from CLI ##`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who updated Device Template last`,
				},
				resource.Attribute{
					Name:        "template_configuration",
					Description: `Template Configuration for Device Template`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Time when Device Template was created`,
				},
				resource.Attribute{
					Name:        "rid",
					Description: `Request ID for Device Template`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `Feature for Device Template`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `User who created Device Template`,
				},
				resource.Attribute{
					Name:        "last_updated_on",
					Description: `Time when Device Template was last updated`,
				},
				resource.Attribute{
					Name:        "template_attached",
					Description: `Number of templates attached to Device Template`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_device_uuids",
			Category:         "Data Sources",
			ShortDescription: `Data source for SD-WAN Device UUID List`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_devices_attachment",
			Category:         "Data Sources",
			ShortDescription: `Data Source for SD-WAN Devices Attachment`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_template_id",
					Description: `(Required) Unique Device Template ID ## Common Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "devices_list",
					Description: `List of Devices attached to the Device Template (see [below for nested schema](#nestedblock--devices_list)) <a id="nestedblock--devices_list"></a> ## Nested Schema for devices_list`,
				},
				resource.Attribute{
					Name:        "chassis_number",
					Description: `Chassis number of the Device attached to the Device Template`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_feature_template_ids",
			Category:         "Data Sources",
			ShortDescription: `Data source for SD-WAN Feature Templates List`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_type",
					Description: `(Required) Type of the Feature Template ## Common Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `List of Templates on vManage server with the given Template Type (see [below for nested schema](#nestedblock--template)) <a id="nestedblock--template"></a> ## Nested Schema for template`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `Name of the Feature Template`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Template ID of the Feature Template`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_ntp_feature_template",
			Category:         "Data Sources",
			ShortDescription: `Data source for SD-WAN NTP Feature Template`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Unique NTP Type Feature Template Name ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_description",
					Description: `Long Description of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Type of device which supports NTP Feature Template`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Type of template which supports NTP Feature Template`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `Minimum Version for the NTP Feature template`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `Configuration for NTP type Feature Template, (see [below for nested schema](#nestedblock--template_definition))`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `Boolean flag indicating whether NTP type Feature Template is factory default or not`,
				},
				resource.Attribute{
					Name:        "attached_masters_count",
					Description: `Number of Device Templates attached to NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "devices_attached",
					Description: `Number of Devices attached to NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who updated the NTP type Feature Template latest`,
				},
				resource.Attribute{
					Name:        "last_updated_on",
					Description: `Time for the latest Update of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "g_template_class",
					Description: `Template Class type for NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Unique ID for NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Type of configuration for NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Time of creation of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "rid",
					Description: `rid for the NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `Feature for the NTP type Feature Template <a id="nestedblock--template_definition"></a> ## Nested Schema for template_definition`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Server configuration for the NTP type Feature Template, (see [below for nested schema](#nestedblock--server))`,
				},
				resource.Attribute{
					Name:        "master",
					Description: `(Optional) Master configuration for the Cisco NTP type Feature Template, see [below for nested schema](#nestedblock--master))`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Authentication configuration for the NTP type Feature Template, (see [below for nested schema](#nestedblock--authentication))`,
				},
				resource.Attribute{
					Name:        "trusted",
					Description: `Trusted key for the NTP type Feature Template <a id="nestedblock--server"></a> ## Nested Schema for server`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `IP address of an NTP server, or a DNS server that knows how to reach the NTP server`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key Id for configuration of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Vpn ID for configuration of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version number of the NTP protocol software for configuration of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "source_interface",
					Description: `Name of a specific interface to use for outgoing NTP packets`,
				},
				resource.Attribute{
					Name:        "prefer",
					Description: `It is a boolean flag, set true if multiple NTP servers are at the same stratum level and you want one to be preferred <a id="nestedblock--master"></a> ## Nested Schema for master`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Boolean flag indicating whether master is enabled for configuration of Cisco NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "stratum",
					Description: `(Optional) Stratum for configuration of Cisco NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Source for configuration of Cisco NTP type Feature Template <a id="nestedblock--authentication"></a> ## Nested Schema for authentication`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `MD5 authentication key ID for configuration of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Key value for configuration of NTP type Feature Template`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_description",
					Description: `Long Description of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Type of device which supports NTP Feature Template`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Type of template which supports NTP Feature Template`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `Minimum Version for the NTP Feature template`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `Configuration for NTP type Feature Template, (see [below for nested schema](#nestedblock--template_definition))`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `Boolean flag indicating whether NTP type Feature Template is factory default or not`,
				},
				resource.Attribute{
					Name:        "attached_masters_count",
					Description: `Number of Device Templates attached to NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "devices_attached",
					Description: `Number of Devices attached to NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who updated the NTP type Feature Template latest`,
				},
				resource.Attribute{
					Name:        "last_updated_on",
					Description: `Time for the latest Update of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "g_template_class",
					Description: `Template Class type for NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Unique ID for NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Type of configuration for NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Time of creation of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "rid",
					Description: `rid for the NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `Feature for the NTP type Feature Template <a id="nestedblock--template_definition"></a> ## Nested Schema for template_definition`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Server configuration for the NTP type Feature Template, (see [below for nested schema](#nestedblock--server))`,
				},
				resource.Attribute{
					Name:        "master",
					Description: `(Optional) Master configuration for the Cisco NTP type Feature Template, see [below for nested schema](#nestedblock--master))`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Authentication configuration for the NTP type Feature Template, (see [below for nested schema](#nestedblock--authentication))`,
				},
				resource.Attribute{
					Name:        "trusted",
					Description: `Trusted key for the NTP type Feature Template <a id="nestedblock--server"></a> ## Nested Schema for server`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `IP address of an NTP server, or a DNS server that knows how to reach the NTP server`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key Id for configuration of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Vpn ID for configuration of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version number of the NTP protocol software for configuration of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "source_interface",
					Description: `Name of a specific interface to use for outgoing NTP packets`,
				},
				resource.Attribute{
					Name:        "prefer",
					Description: `It is a boolean flag, set true if multiple NTP servers are at the same stratum level and you want one to be preferred <a id="nestedblock--master"></a> ## Nested Schema for master`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Boolean flag indicating whether master is enabled for configuration of Cisco NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "stratum",
					Description: `(Optional) Stratum for configuration of Cisco NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Source for configuration of Cisco NTP type Feature Template <a id="nestedblock--authentication"></a> ## Nested Schema for authentication`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `MD5 authentication key ID for configuration of NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Key value for configuration of NTP type Feature Template`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_system_feature_template",
			Category:         "Data Sources",
			ShortDescription: `Data source for SD-WAN System Feature Template`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Unique System Type Feature Template Name ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_description",
					Description: `Long Description of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Type of device which supports Feature Template`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `Minimum Version for the Feature template`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `Configuration for System type Feature Template, (see [below for nested schema](#nestedblock--template_definition))`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `Boolean flag indicating whether System type Feature Template is factory default or not`,
				},
				resource.Attribute{
					Name:        "attached_masters_count",
					Description: `Number of Device Templates attached to System type Feature Template`,
				},
				resource.Attribute{
					Name:        "devices_attached",
					Description: `Number of Devices attached to System type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who updated the System type Feature Template latest`,
				},
				resource.Attribute{
					Name:        "last_updated_on",
					Description: `Time for the latest Update of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "g_template_class",
					Description: `Template Class type for System type Feature Template`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Unique ID for System type Feature Template`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Type of configuration for System type Feature Template`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Time of creation of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "rid",
					Description: `rid for the System type Feature Template`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `Feature for the System type Feature Template <a id="nestedblock--template_definition"></a> ## Nested Schema for template_definition`,
				},
				resource.Attribute{
					Name:        "system_basic",
					Description: `Basic configuration for the System type Feature Template, (see [below for nested schema](#nestedblock--system_basic))`,
				},
				resource.Attribute{
					Name:        "system_gps",
					Description: `GPS configuration for the System type Feature Template, (see [below for nested schema](#nestedblock--system_gps))`,
				},
				resource.Attribute{
					Name:        "system_trackers",
					Description: `Trackers for the System type Feature Template, (see [below for nested schema](#nestedblock--system_trackers))`,
				},
				resource.Attribute{
					Name:        "system_advanced",
					Description: `Advanced configuration for the System type Feature Template, (see [below for nested schema](#nestedblock--system_advanced)) <a id="nestedblock--system_basic"></a> ## Nested Schema for system_basic`,
				},
				resource.Attribute{
					Name:        "overlay_id",
					Description: `Overlay ID for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Timezone for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "device_groups",
					Description: `Device Groups for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "controller_groups",
					Description: `Controller Group List for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "console_baud_rate",
					Description: `Console Baud Rate for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "max_omp_sessions",
					Description: `Maximum OMP Sessions for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tcp_optimization_enabled",
					Description: `Boolean flag indicating whether TCP Optimization is enabled for configuration of System type Feature Template <a id="nestedblock--system_gps"></a> ## Nested Schema for system_gps`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `Latitude for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `Longitude for configuration of System type Feature Template <a id="nestedblock--system_trackers"></a> ## Nested Schema for system_trackers`,
				},
				resource.Attribute{
					Name:        "tracker_name",
					Description: `Name of Tracker for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_type",
					Description: `Type of Tracker for configuration of Cisco System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_endpoint_type",
					Description: `Type of Tracker Endpoint for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_endpoint_value",
					Description: `Value of Tracker Endpoint for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_threshold",
					Description: `Threshold of Tracker for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_interval",
					Description: `Interval of Tracker for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_multiplier",
					Description: `Multiplier of Tracker for configuration of System type Feature Template <a id="nestedblock--system_advanced"></a> ## Nested Schema for system_advanced`,
				},
				resource.Attribute{
					Name:        "control_session_pps",
					Description: `Policer Rate(pps) for Control Sessions for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "system_tunnel_mtu",
					Description: `MTU of System's internal DTLS Tunnel for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "port_hop",
					Description: `Boolean flag to indicate whether Port Hopping is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "port_offset",
					Description: `TLOC Port Offset for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "dns_cache_timeout",
					Description: `DNS Cache Timeout(minutes) for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_transport",
					Description: `Boolean flag to indicate whether tracking of transport is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "vbond_local",
					Description: `Boolean flag to indicate whether the local device is set as vBond for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_interface_tag",
					Description: `Interface tracking for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "multicast_buffer_percent",
					Description: `Percentage(5-100) of buffer multicast packets can consume for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "usb_controller",
					Description: `Boolean flag to indicate whether external USB Controller on the device is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_default-gateway",
					Description: `Boolean flag to indicate whether default Gateway Tracking on the device is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "host_policer_pps",
					Description: `Rate(1000-25000) at which to police packets bound to the control plane for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "icmp_error_pps",
					Description: `Rate at which to police ICMP error messages for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "allow_same_site_tunnels",
					Description: `Boolean flag to indicate whether tunnels are allowed to be formed between vEdges in the same site for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "route_consistency_check",
					Description: `Boolean flag to indicate whether route consistency check is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "admin_tech_on_failure",
					Description: `Boolean flag to indicate whether the collection of admin-tech before reboot due to daemon failure is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle CLI timeout(minutes) for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "eco_friendly_mode",
					Description: `Boolean flag to indicate whether vEdge Cloud router can run in the eco-friendly mode for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "on_demand_enable",
					Description: `Boolean flag to indicate whether On-Demand Tunnel is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "on_demand_idle_timeout",
					Description: `Idle timeout for On-demand Tunnels for configuration of System type Feature Template`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_description",
					Description: `Long Description of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Type of device which supports Feature Template`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `Minimum Version for the Feature template`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `Configuration for System type Feature Template, (see [below for nested schema](#nestedblock--template_definition))`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `Boolean flag indicating whether System type Feature Template is factory default or not`,
				},
				resource.Attribute{
					Name:        "attached_masters_count",
					Description: `Number of Device Templates attached to System type Feature Template`,
				},
				resource.Attribute{
					Name:        "devices_attached",
					Description: `Number of Devices attached to System type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who updated the System type Feature Template latest`,
				},
				resource.Attribute{
					Name:        "last_updated_on",
					Description: `Time for the latest Update of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "g_template_class",
					Description: `Template Class type for System type Feature Template`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Unique ID for System type Feature Template`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Type of configuration for System type Feature Template`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Time of creation of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "rid",
					Description: `rid for the System type Feature Template`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `Feature for the System type Feature Template <a id="nestedblock--template_definition"></a> ## Nested Schema for template_definition`,
				},
				resource.Attribute{
					Name:        "system_basic",
					Description: `Basic configuration for the System type Feature Template, (see [below for nested schema](#nestedblock--system_basic))`,
				},
				resource.Attribute{
					Name:        "system_gps",
					Description: `GPS configuration for the System type Feature Template, (see [below for nested schema](#nestedblock--system_gps))`,
				},
				resource.Attribute{
					Name:        "system_trackers",
					Description: `Trackers for the System type Feature Template, (see [below for nested schema](#nestedblock--system_trackers))`,
				},
				resource.Attribute{
					Name:        "system_advanced",
					Description: `Advanced configuration for the System type Feature Template, (see [below for nested schema](#nestedblock--system_advanced)) <a id="nestedblock--system_basic"></a> ## Nested Schema for system_basic`,
				},
				resource.Attribute{
					Name:        "overlay_id",
					Description: `Overlay ID for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Timezone for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "device_groups",
					Description: `Device Groups for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "controller_groups",
					Description: `Controller Group List for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "console_baud_rate",
					Description: `Console Baud Rate for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "max_omp_sessions",
					Description: `Maximum OMP Sessions for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tcp_optimization_enabled",
					Description: `Boolean flag indicating whether TCP Optimization is enabled for configuration of System type Feature Template <a id="nestedblock--system_gps"></a> ## Nested Schema for system_gps`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `Latitude for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `Longitude for configuration of System type Feature Template <a id="nestedblock--system_trackers"></a> ## Nested Schema for system_trackers`,
				},
				resource.Attribute{
					Name:        "tracker_name",
					Description: `Name of Tracker for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_type",
					Description: `Type of Tracker for configuration of Cisco System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_endpoint_type",
					Description: `Type of Tracker Endpoint for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_endpoint_value",
					Description: `Value of Tracker Endpoint for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_threshold",
					Description: `Threshold of Tracker for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_interval",
					Description: `Interval of Tracker for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_multiplier",
					Description: `Multiplier of Tracker for configuration of System type Feature Template <a id="nestedblock--system_advanced"></a> ## Nested Schema for system_advanced`,
				},
				resource.Attribute{
					Name:        "control_session_pps",
					Description: `Policer Rate(pps) for Control Sessions for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "system_tunnel_mtu",
					Description: `MTU of System's internal DTLS Tunnel for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "port_hop",
					Description: `Boolean flag to indicate whether Port Hopping is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "port_offset",
					Description: `TLOC Port Offset for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "dns_cache_timeout",
					Description: `DNS Cache Timeout(minutes) for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_transport",
					Description: `Boolean flag to indicate whether tracking of transport is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "vbond_local",
					Description: `Boolean flag to indicate whether the local device is set as vBond for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_interface_tag",
					Description: `Interface tracking for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "multicast_buffer_percent",
					Description: `Percentage(5-100) of buffer multicast packets can consume for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "usb_controller",
					Description: `Boolean flag to indicate whether external USB Controller on the device is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_default-gateway",
					Description: `Boolean flag to indicate whether default Gateway Tracking on the device is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "host_policer_pps",
					Description: `Rate(1000-25000) at which to police packets bound to the control plane for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "icmp_error_pps",
					Description: `Rate at which to police ICMP error messages for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "allow_same_site_tunnels",
					Description: `Boolean flag to indicate whether tunnels are allowed to be formed between vEdges in the same site for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "route_consistency_check",
					Description: `Boolean flag to indicate whether route consistency check is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "admin_tech_on_failure",
					Description: `Boolean flag to indicate whether the collection of admin-tech before reboot due to daemon failure is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle CLI timeout(minutes) for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "eco_friendly_mode",
					Description: `Boolean flag to indicate whether vEdge Cloud router can run in the eco-friendly mode for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "on_demand_enable",
					Description: `Boolean flag to indicate whether On-Demand Tunnel is enabled for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "on_demand_idle_timeout",
					Description: `Idle timeout for On-demand Tunnels for configuration of System type Feature Template`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_vpn_feature_template",
			Category:         "Data Sources",
			ShortDescription: `Data source for SD-WAN VPN Feature Template`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Unique VPN Type Feature Template Name ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Type of VPN Feature Template`,
				},
				resource.Attribute{
					Name:        "template_description",
					Description: `Long Description of VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Type of device which supports Feature Template`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `Minimum Version for the Feature template`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `Configuration for VPN type Feature Template, (see [below for nested schema](#nestedblock--template_definition))`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `Boolean flag indicating whether VPN type Feature Template is factory default or not`,
				},
				resource.Attribute{
					Name:        "attached_masters_count",
					Description: `Number of Device Templates attached to VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "devices_attached",
					Description: `Number of Devices attached to VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who updated the VPN type Feature Template latest`,
				},
				resource.Attribute{
					Name:        "last_updated_on",
					Description: `Time for the latest Update of VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "g_template_class",
					Description: `Template Class type for VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Unique ID for VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Type of configuration for VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Time of creation of VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "rid",
					Description: `Request id for the VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `Feature for the VPN type Feature Template <a id="nestedblock--template_definition"></a> ## Nested Schema for template_definition`,
				},
				resource.Attribute{
					Name:        "vpn_basic",
					Description: `Basic configuration for the VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_basic))`,
				},
				resource.Attribute{
					Name:        "vpn_dns",
					Description: `DNS configuration for the VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_dns))`,
				},
				resource.Attribute{
					Name:        "ipv4_route",
					Description: `IPv4 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv4_route))`,
				},
				resource.Attribute{
					Name:        "ipv6_route",
					Description: `IPv6 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv6_route))`,
				},
				resource.Attribute{
					Name:        "te_service_enabled",
					Description: `Flag indicating TE service is enabled or not for the VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "service_route",
					Description: `Service route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--service_route))`,
				},
				resource.Attribute{
					Name:        "nat64",
					Description: `NAT64 for the VPN type Feature Template, (see [below for nested schema](#nestedblock--nat64)) <a id="nestedblock--vpn_basic"></a> ## Nested Schema for vpn_basic`,
				},
				resource.Attribute{
					Name:        "vpn_id",
					Description: `Numeric identifier of the VPN for VPN type of Feature Template`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of VPN for VPN type of Feature Template`,
				},
				resource.Attribute{
					Name:        "ecmp_hash_key",
					Description: `Boolean flag indicating whether ECMP hash key is enabled or not <a id="nestedblock--vpn_dns"></a> ## Nested Schema for vpn_dns`,
				},
				resource.Attribute{
					Name:        "primary_dns_ipv4",
					Description: `Primary IPv4 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "secondary_dns_ipv4",
					Description: `Secondary IPv4 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "primary_dns_ipv6",
					Description: `Primary IPv6 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "secondary_dns_ipv4",
					Description: `Secondary IPv6 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "vpn_host",
					Description: `VPN host list for DNS server (see [below for nested schema](#nestedblock--vpn_host)) <a id="nestedblock--ipv4_route"></a> ## Nested Schema for ipv4_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `IPv4 prefix or address for IPv4 route`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `Gateway Type for next hop to reach the static route`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `Next hop list (see [below for nested schema](#nestedblock--next_hop)) when gateway_type is ` + "`" + `next-hop` + "`" + ``,
				},
				resource.Attribute{
					Name:        "null0_distance",
					Description: `Null0 distance when gateway_type is ` + "`" + `null0` + "`" + ` <a id="nestedblock--ipv6_route"></a> ## Nested Schema for ipv6_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `IPv6 prefix or address for IPv6 route`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `Gateway Type for next hop to reach the static route`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `Next hop list (see [below for nested schema](#nestedblock--next_hop)) when gateway_type is ` + "`" + `next-hop` + "`" + ``,
				},
				resource.Attribute{
					Name:        "null0_distance",
					Description: `Null0 distance when gateway_type is ` + "`" + `null0` + "`" + ` <a id="nestedblock--service_route"></a> ## Nested Schema for service_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `IPv6 or IPv4 prefix or address for service route`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `Type of the service <a id="nestedblock--nat64"></a> ## Nested Schema for nat64`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `NAT64 pool name`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `Starting IPv4 address of NAT64 pool`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `Ending IPv4 address of NAT64 pool`,
				},
				resource.Attribute{
					Name:        "overload",
					Description: `Flag indicating NAT64 pool overload <a id="nestedblock--vpn_host"></a> ## Nested Schema for vpn_host`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Host name of the DNS server`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `List of ip addresses associated with the hostname <a id="nestedblock--next_hop"></a> ## Nested Schema for next_hop`,
				},
				resource.Attribute{
					Name:        "next_hop_address",
					Description: `IP address of the nexthop`,
				},
				resource.Attribute{
					Name:        "next_hop_distance",
					Description: `Distance of the nexthop`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_type",
					Description: `Type of VPN Feature Template`,
				},
				resource.Attribute{
					Name:        "template_description",
					Description: `Long Description of VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Type of device which supports Feature Template`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `Minimum Version for the Feature template`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `Configuration for VPN type Feature Template, (see [below for nested schema](#nestedblock--template_definition))`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `Boolean flag indicating whether VPN type Feature Template is factory default or not`,
				},
				resource.Attribute{
					Name:        "attached_masters_count",
					Description: `Number of Device Templates attached to VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "devices_attached",
					Description: `Number of Devices attached to VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who updated the VPN type Feature Template latest`,
				},
				resource.Attribute{
					Name:        "last_updated_on",
					Description: `Time for the latest Update of VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "g_template_class",
					Description: `Template Class type for VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Unique ID for VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Type of configuration for VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Time of creation of VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "rid",
					Description: `Request id for the VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `Feature for the VPN type Feature Template <a id="nestedblock--template_definition"></a> ## Nested Schema for template_definition`,
				},
				resource.Attribute{
					Name:        "vpn_basic",
					Description: `Basic configuration for the VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_basic))`,
				},
				resource.Attribute{
					Name:        "vpn_dns",
					Description: `DNS configuration for the VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_dns))`,
				},
				resource.Attribute{
					Name:        "ipv4_route",
					Description: `IPv4 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv4_route))`,
				},
				resource.Attribute{
					Name:        "ipv6_route",
					Description: `IPv6 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv6_route))`,
				},
				resource.Attribute{
					Name:        "te_service_enabled",
					Description: `Flag indicating TE service is enabled or not for the VPN type Feature Template`,
				},
				resource.Attribute{
					Name:        "service_route",
					Description: `Service route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--service_route))`,
				},
				resource.Attribute{
					Name:        "nat64",
					Description: `NAT64 for the VPN type Feature Template, (see [below for nested schema](#nestedblock--nat64)) <a id="nestedblock--vpn_basic"></a> ## Nested Schema for vpn_basic`,
				},
				resource.Attribute{
					Name:        "vpn_id",
					Description: `Numeric identifier of the VPN for VPN type of Feature Template`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of VPN for VPN type of Feature Template`,
				},
				resource.Attribute{
					Name:        "ecmp_hash_key",
					Description: `Boolean flag indicating whether ECMP hash key is enabled or not <a id="nestedblock--vpn_dns"></a> ## Nested Schema for vpn_dns`,
				},
				resource.Attribute{
					Name:        "primary_dns_ipv4",
					Description: `Primary IPv4 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "secondary_dns_ipv4",
					Description: `Secondary IPv4 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "primary_dns_ipv6",
					Description: `Primary IPv6 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "secondary_dns_ipv4",
					Description: `Secondary IPv6 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "vpn_host",
					Description: `VPN host list for DNS server (see [below for nested schema](#nestedblock--vpn_host)) <a id="nestedblock--ipv4_route"></a> ## Nested Schema for ipv4_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `IPv4 prefix or address for IPv4 route`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `Gateway Type for next hop to reach the static route`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `Next hop list (see [below for nested schema](#nestedblock--next_hop)) when gateway_type is ` + "`" + `next-hop` + "`" + ``,
				},
				resource.Attribute{
					Name:        "null0_distance",
					Description: `Null0 distance when gateway_type is ` + "`" + `null0` + "`" + ` <a id="nestedblock--ipv6_route"></a> ## Nested Schema for ipv6_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `IPv6 prefix or address for IPv6 route`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `Gateway Type for next hop to reach the static route`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `Next hop list (see [below for nested schema](#nestedblock--next_hop)) when gateway_type is ` + "`" + `next-hop` + "`" + ``,
				},
				resource.Attribute{
					Name:        "null0_distance",
					Description: `Null0 distance when gateway_type is ` + "`" + `null0` + "`" + ` <a id="nestedblock--service_route"></a> ## Nested Schema for service_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `IPv6 or IPv4 prefix or address for service route`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `Type of the service <a id="nestedblock--nat64"></a> ## Nested Schema for nat64`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `NAT64 pool name`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `Starting IPv4 address of NAT64 pool`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `Ending IPv4 address of NAT64 pool`,
				},
				resource.Attribute{
					Name:        "overload",
					Description: `Flag indicating NAT64 pool overload <a id="nestedblock--vpn_host"></a> ## Nested Schema for vpn_host`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Host name of the DNS server`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `List of ip addresses associated with the hostname <a id="nestedblock--next_hop"></a> ## Nested Schema for next_hop`,
				},
				resource.Attribute{
					Name:        "next_hop_address",
					Description: `IP address of the nexthop`,
				},
				resource.Attribute{
					Name:        "next_hop_distance",
					Description: `Distance of the nexthop`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_vpn_interface_feature_template",
			Category:         "Data Sources",
			ShortDescription: `Data source for SD-WAN VPN interface Feature Template`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Unique VPN interface Type Feature Template Name ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "template_description",
					Description: `Long Description of VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Type of device which supports VPN interface Feature Template`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Template type for the VPN Interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `Minimum Version for the VPN interface Feature template`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `Configuration for VPN interface type Feature Template, (see [below for nested schema](#nestedblock--template_definition))`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `Boolean flag indicating whether VPN interface type Feature Template is factory default or not`,
				},
				resource.Attribute{
					Name:        "attached_masters_count",
					Description: `Number of Device Templates attached to VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "devices_attached",
					Description: `Number of Devices attached to VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who updated the VPN interface type Feature Template latest`,
				},
				resource.Attribute{
					Name:        "last_updated_on",
					Description: `Time for the latest Update of VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "g_template_class",
					Description: `Template Class type for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Unique ID for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Type of configuration for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Time of creation of VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "rid",
					Description: `rid for the VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `Feature for the VPN interface type Feature Template <a id="nestedblock--template_definition"></a> ## Nested Schema for template_definition`,
				},
				resource.Attribute{
					Name:        "vpn_interface_basic",
					Description: `Basic configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_basic))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_tunnel",
					Description: `Tunnel configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_tunnel))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_nat",
					Description: `Tunnel configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_nat))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_vrrp",
					Description: `VRRP(Virtual Router Redundancy Protocol) configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_vrrp))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_acl_qos",
					Description: `ACL(Apply Access Lists) and QoS configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_acl_qos))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_arp",
					Description: `ARP(Address Resolution Protocol) configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_arp))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_8021x",
					Description: `IEEE 802.1X Authentication configuration for the Vedge VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_8021x))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_trustsec",
					Description: `TrustSec configuration for the Cisco VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_trustsec))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_advanced",
					Description: `Advanced configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_advanced)) <a id="nestedblock--vpn_interface_basic"></a> ## Nested Schema for vpn_interface_basic`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `Shutdown flag for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Interface Description for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `IPv4 configuration for VPN interface type Feature Template, (see [below for nested schema](#nestedblock--ipv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `IPv6 configuration for VPN interface type Feature Template, (see [below for nested schema](#nestedblock--ipv6))`,
				},
				resource.Attribute{
					Name:        "block_non_source_ip",
					Description: `Flag indicating whether forwarded traffic matches with interface's IP prefix range for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "bandwidth_upstream",
					Description: `Value of bandwidth above which to generate notifications in case of transmitted traffic for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "bandwidth_downstream",
					Description: `Value of bandwidth above which to generate notifications in case of received traffic for VPN interface type Feature Template <a id="nestedblock--vpn_interface_tunnel"></a> ## Nested Schema for vpn_interface_tunnel`,
				},
				resource.Attribute{
					Name:        "per_tunnel_qos",
					Description: `Flag indicating whether per_tunnel_qos is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "per_tunnel_qos_aggregator",
					Description: `Flag indicating whether per_tunnel_qos_aggregator is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "tunnels_bandwidth_percent",
					Description: `Value of tunnel bandwith percentage for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `TLOC color value for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "restrict",
					Description: `Restrict flag for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Value of Groups in tunnle for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "border",
					Description: `Flag indicating whether Border is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "control_connection",
					Description: `Flag indicating TLOC connection is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "maximum_control_connections",
					Description: `Maximum number of vSamart controllers that WAN tunnel interface can connect for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "vbond_as_stun_server",
					Description: `Flag indicating whether Session Traversal Utilities for NAT(STUN) is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "exclude_controller_group_list",
					Description: `List of vSmart controllers that the tunnel interface is not allowed to connect for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "vmanage_connection_preference",
					Description: `Flag indicating preference for using a tunnel interface to exchange control traffic with the vManage NMS for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "port_hop",
					Description: `Flag indicating whether Port Hop is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "low_bandwidth_link",
					Description: `Flag indicating whether tunnle is set as a low-bandwidth link for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "allow_service",
					Description: `Nested block for allow or disallow various services for VPN interface type Feature Template(see [below for nested schema](#nestedblock--allow_service))`,
				},
				resource.Attribute{
					Name:        "encapsulation",
					Description: `Encapsulation configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--encapsulation)) <a id="nestedblock--vpn_interface_nat"></a> ## Nested Schema for vpn_interface_nat`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `NAT IPv4 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--natv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `NAT IPv6 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--natv6)) <a id="nestedblock--vpn_interface_vrrp"></a> ## Nested Schema for vpn_interface_vrrp`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `VRRP IPv4 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--vrrpv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `VRRP IPv6 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--vrrpv6)) <a id="nestedblock--vpn_interface_acl_qos"></a> ## Nested Schema for vpn_interface_acl_qos`,
				},
				resource.Attribute{
					Name:        "adapt_period",
					Description: `Value of Adapt Period for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_min",
					Description: `Minimum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_max",
					Description: `Maximum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_default",
					Description: `Default value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_min",
					Description: `Minimum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_max",
					Description: `Maximum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_default",
					Description: `Default value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate",
					Description: `Value of aggregate traffic transmission rate on the VPN interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "qos_map",
					Description: `Name of the QoS map to apply to packets being transmitted out the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "rewrite_rule",
					Description: `Name of the rewrite rule to apply on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4_ingress_access_list",
					Description: `Name of the access list to apply to IPv4 packets being received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4_egress_access_list",
					Description: `Name of the access list to apply to IPv4 packets being transmitted on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv6_ingress_access_list",
					Description: `Name of the access list to apply to IPv6 packets being received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv6_egress_access_list",
					Description: `Name of the access list to apply to IPv6 packets being transmitted on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ingress_policer_name",
					Description: `Name of the policer to apply to packets received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "egress_policer_name",
					Description: `Name of the policer to apply to packets being transmitted on the interface for Vedge VPN interface type Feature Template <a id="nestedblock--vpn_interface_arp"></a> ## Nested Schema for vpn_interface_arp`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address for the ARP entry in dotted decimal notation or as a fully qualified host name for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `MAC address in colon-separated hexadecimal notation for VPN interface type Feature Template <a id="nestedblock--vpn_interface_8021x"></a> ## Nested Schema for vpn_interface_8021x`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `The tag of the RADIUS server to use for 802.1X authentication for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "account_interim_interval",
					Description: `Value of how often to send 802.1X interim accounting updates to the RADIUS server for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nas_identifier",
					Description: `The NAS identifier of the local router for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nas_ip",
					Description: `The NAS IP address of the local router for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "wake_on_lan",
					Description: `The flag that enable client to be powered up when the router receives an Ethernet magic packet frame for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "control_direction",
					Description: `Direction type of packets for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "reauthentication",
					Description: `Value indicating how often to reauthenticate 802.1X clients for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "inactivity",
					Description: `Value indicating how long to wait before revoking an 802.1X client's network access for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "host_mode",
					Description: `Value indicating whether an 802.1X interface grants access to a single client or to multiple clients for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "advanced_options",
					Description: `Advance configuration of vpn_interface_8021x for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--advanced_options)) <a id="nestedblock--vpn_interface_trustsec"></a> ## Nested Schema for vpn_interface_trustsec`,
				},
				resource.Attribute{
					Name:        "enable_sgt_propagation",
					Description: `Flag indicating whether SGT propagation is enabled or not for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "propagate",
					Description: `Flag indicating whether propagate is enabled or not for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "security_group_tag",
					Description: `Value of security group tag for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "trusted",
					Description: `Flag indicating whether trusted is enable or not for Cisco VPN interface type Feature Template <a id="nestedblock--vpn_interface_advanced"></a> ## Nested Schema for vpn_interface_advanced`,
				},
				resource.Attribute{
					Name:        "duplex",
					Description: `Value of duplex to run interface in auto, full, or half mode for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `Value of MAC address to associate with the interface, in colon-separated hexadecimal notation for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ip_mtu",
					Description: `Value of MAC address to associate with the interface, in colon-separated hexadecimal notation for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "pmtu_discovery",
					Description: `Flag indicating whether path MTU is enable or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "flow_control",
					Description: `Value of bidirectional flow control settings for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `Maximum segment size (MSS) of TPC SYN packets passing through the router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Value of speed to use when the remote end of the connection does not support autonegotiation for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "clear_dont_fragment",
					Description: `Flag indicating Don't Fragment (DF) bit in the IPv4 packet header for packets being transmitted out the interface is clar or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "static_ingess_qos",
					Description: `Queue number to use for incoming traffic for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "arp_timeout",
					Description: `Value indicating how long it takes for a dynamically learned ARP entry to time out for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "autonegotiation",
					Description: `Flag indicating autonegotiation mode for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "tloc_extension",
					Description: `Name of a physical interface on the same router that connects to the WAN transport for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "power_over_ethernet",
					Description: `Flag indicating whether PoE(Power over Ethernet) is enabled or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "load_interval",
					Description: `Value of load interval for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker",
					Description: `Name of a tracker to track the status of transport interfaces that connect to the internet for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "icmp_redirect_disable",
					Description: `Flag indicating whether ICMP redirect disable or not on interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "gre_tunnel_source_ip",
					Description: `IP address of the extended WAN interface`,
				},
				resource.Attribute{
					Name:        "xconnect",
					Description: `The name of a physical interface on the same router that connects to the WAN transport for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ip_directed_broadcast",
					Description: `Flag indicate whether directed IP braodcast is enabled or not for VPN interface type Feature Template <a id="nestedblock--ipv4"></a> ## Nested Schema for Basic ipv4`,
				},
				resource.Attribute{
					Name:        "primary_address",
					Description: `Value of IPv4 primary address for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "secondary_address",
					Description: `Value of IPv4 secondary address for VPN interface type Feature Template(see [below for nested schema](#nestedblock--secondary_addressv4))`,
				},
				resource.Attribute{
					Name:        "dhcp_distance",
					Description: `Value of IPv4 DHCP distance for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp_helper",
					Description: `Value of IPv4 DHCP Helper for VPN interface type Feature Template <a id="nestedblock--ipv6"></a> ## Nested Schema for Basic ipv6`,
				},
				resource.Attribute{
					Name:        "primary_address",
					Description: `Value of IPv6 primary address for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "secondary_address",
					Description: `Value of IPv6 secondary address for VPN interface type Feature Template(see [below for nested schema](#nestedblock--secondary_addressv6))`,
				},
				resource.Attribute{
					Name:        "dhcp_distance",
					Description: `Value of IPv6 DHCP distance for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp_helper",
					Description: `Value of IPv6 DHCP Helper for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp_rapid_commit",
					Description: `Value of IPv6 DHCP rapid commit enabled or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp_helper",
					Description: `Value of IPv6 DHCP rapid commit enabled or not for VPN interface type Feature Template(see [below for nested schema](#nestedblock--dhcp_helper)) <a id="nestedblock--allow_service"></a> ## Nested Schema for Tunnel allow_service`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `Flag indicating whether All serivce is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "bgp",
					Description: `Flag indicating whether BGP is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Flag indicating whether DHCP is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `Flag indicating whether DNS is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "icmp",
					Description: `Flag indicating whether ICMP is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "netconf",
					Description: `Flag indicating whether NETCONF is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ntp",
					Description: `Flag indicating whether NTP is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `Flag indicating whether OSPF is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ssh",
					Description: `Flag indicating whether SSH is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "stun",
					Description: `Flag indicating whether STUN is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "https",
					Description: `Flag indicating whether HTTPS is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "snmp",
					Description: `Flag indicating whether SNMP is enabled or not for Cisco VPN interface type Feature Template <a id="nestedblock--encapsulation"></a> ## Nested Schema for Tunnel encapsulation`,
				},
				resource.Attribute{
					Name:        "gre_preference",
					Description: `Value of GRE preference for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "gre_weight",
					Description: `Value of GRE weight for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipsec_preference",
					Description: `Value of IPsec preference for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipsec_weight",
					Description: `Value of IPsec weight for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "carrier",
					Description: `Carrier name or private network identifier to associate with the tunnel for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "bind_loopback_tunnel",
					Description: `Name of a physical interface to bind to a loopback interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_resort_circuit",
					Description: `Flag indicating whether tunnel interface is used as the circuit of last resort for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `Value of hold time for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nat_refresh_interval",
					Description: `Value of interval between NAT refresh packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `Value of interval between Hello packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "hello_tolerance",
					Description: `Value of the time to wait for a Hello packet on a DTLS or TLS WAN transport connection before declaring that transport tunnel to be down for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "gre_tunnel_destination_ip",
					Description: `Value of GRE tunnel destination IP for VPN interface type Feature Template <a id="nestedblock--natv4"></a> ## Nested Schema for NAT ipv4`,
				},
				resource.Attribute{
					Name:        "nat_type",
					Description: `NAT type for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "refresh_mode",
					Description: `Refresh mode for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "log_nat_flow_creations_or_deletions",
					Description: `Flag indicating whether log nat flow creations or deletions are enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `The time when NAT translations over UDP sessions time out for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `The time when NAT translations over TCP sessions time out for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "block_icmp",
					Description: `Flag indicating whether to block icmp or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "respond_to_ping",
					Description: `Flag indicating whether to respond to ping or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nat_pool_range_start",
					Description: `Value of starting IP address for the NAT pool for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nat_pool_range_end",
					Description: `Value of closing IP address for the NAT pool for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "overload",
					Description: `Flag indicating whether to enable per-port translation or not for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nat_inside_source_loopback_interface",
					Description: `Value of NAT inside source loopback interface for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "port_forward",
					Description: `Configure port forwarding rule for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--port_forward))`,
				},
				resource.Attribute{
					Name:        "static_nat",
					Description: `Configure static nat for VPN interface type Feature Template(see [below for nested schema](#nestedblock--static_nat)) <a id="nestedblock--natv6"></a> ## Nested Schema for NAT ipv6`,
				},
				resource.Attribute{
					Name:        "nat64",
					Description: `Flag indicating whether to enable nat64 or not for VPN interface type Feature Template <a id="nestedblock--vrrpv4"></a> ## Nested Schema for VRRP ipv4`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_omp",
					Description: `Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_prefix_list",
					Description: `(Required) Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Value of IPv4 address of the virtual router for VPN interface type Feature Template <a id="nestedblock--vrrpv6"></a> ## Nested Schema for VRRP ipv6`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_omp",
					Description: `Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_prefix_list",
					Description: `Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "link_local_ipv6_address",
					Description: `Value of IPv6 address of the virtual router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "global_ipv6_prefix",
					Description: `Value of global IPv6 prefix of the virtual router for VPN interface type Feature Template <a id="nestedblock--advanced_options"></a> ## Nested Schema for 8021.X advanced_options`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `Configure VLAN for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--vlan))`,
				},
				resource.Attribute{
					Name:        "dynamic_authentication_server",
					Description: `Configure Dynamic Authentication Server for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--dynamic_authentication_server))`,
				},
				resource.Attribute{
					Name:        "mac_authentication_bypass",
					Description: `Configure MAC Authentication Bypass for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--mac_authentication_bypass))`,
				},
				resource.Attribute{
					Name:        "request_attributes",
					Description: `Configure Request Attributes for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--request_attributes)) <a id="nestedblock--secondary_addressv4"></a> ## Nested Schema for Basic IPv4 secondary_address`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Value of IPv4 prefix for VPN interface type Feature Template <a id="nestedblock--secondary_addressv6"></a> ## Nested Schema for Basic IPv6 secondary_address`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Value of IPv6 prefix for VPN interface type Feature Template <a id="nestedblock--dhcp_helper"></a> ## Nested Schema for Basic IPv6 dhcp_helper`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Value of IPv6 address for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Value of VPN id for VPN interface type Feature Template <a id="nestedblock--port_forward"></a> ## Nested Schema for NAT ipv4 port_forward`,
				},
				resource.Attribute{
					Name:        "port_start_range",
					Description: `Value of port starting range for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "port_end_range",
					Description: `Value of port ending range for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `VPN id for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Required) Private IP address for Vedge VPN interface type Feature Template <a id="nestedblock--static_nat"></a> ## Nested Schema for NAT ipv4 static_nat`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Value of the inside local address as source IP address for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "translated_source_ip",
					Description: `Value of the inside global address as the translated source IP address for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "source_vpn_id",
					Description: `Configures Source VPN ID, which is the Service VPN in which the host resides for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "static_nat_direction",
					Description: `Set the direction in which to perform network address translation for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `Value of source port for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "translate_port",
					Description: `Value of translate port for Vedge VPN interface type Feature Template <a id="nestedblock--vlan"></a> ## Nested Schema for vlan`,
				},
				resource.Attribute{
					Name:        "authentication_fail_vlan",
					Description: `Value of authentication fail VLAN for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "guest_vlan",
					Description: `Value of guest VLAN for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "authentication_reject_vlan",
					Description: `Value of authentication reject VLAN for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "default_vlan",
					Description: `Value of default VLAN for Vedge VPN interface type Feature Template <a id="nestedblock--dynamic_authentication_server"></a> ## Nested Schema for dynamic_authentication_server`,
				},
				resource.Attribute{
					Name:        "das_port",
					Description: `Value of Dynamic Authentication Server port for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "client",
					Description: `Value of Client IP address for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Secret key Value for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "time_window",
					Description: `Time window value for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "required_timestamp",
					Description: `Flag indicating whether to enable required timestamp or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Value of VPN ID for Vedge VPN interface type Feature Template <a id="nestedblock--mac_authentication_bypass"></a> ## Nested Schema for mac_authentication_bypass`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Flag indicating whether to enable server or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "mac_authentication_bypass_entries",
					Description: `List of MAC Authenticaion bypass entries for Vedge VPN interface type Feature Template <a id="nestedblock--request_attributes"></a> ## Nested Schema for request_attributes`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Configure Authentication for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--authentication))`,
				},
				resource.Attribute{
					Name:        "accounting",
					Description: `Configure Accounting for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--accounting)) ## Nested Schema for request_attributes authentication`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "syntax_choice",
					Description: `Syntax Choice for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value for Vedge VPN interface type Feature Template <a id="nestedblock--accounting"></a> ## Nested Schema for request_attributes accounting`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "syntax_choice",
					Description: `Syntax Choice for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value for Vedge VPN interface type Feature Template`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_description",
					Description: `Long Description of VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Type of device which supports VPN interface Feature Template`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Template type for the VPN Interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `Minimum Version for the VPN interface Feature template`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `Configuration for VPN interface type Feature Template, (see [below for nested schema](#nestedblock--template_definition))`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `Boolean flag indicating whether VPN interface type Feature Template is factory default or not`,
				},
				resource.Attribute{
					Name:        "attached_masters_count",
					Description: `Number of Device Templates attached to VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "devices_attached",
					Description: `Number of Devices attached to VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who updated the VPN interface type Feature Template latest`,
				},
				resource.Attribute{
					Name:        "last_updated_on",
					Description: `Time for the latest Update of VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "g_template_class",
					Description: `Template Class type for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Unique ID for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Type of configuration for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Time of creation of VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "rid",
					Description: `rid for the VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `Feature for the VPN interface type Feature Template <a id="nestedblock--template_definition"></a> ## Nested Schema for template_definition`,
				},
				resource.Attribute{
					Name:        "vpn_interface_basic",
					Description: `Basic configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_basic))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_tunnel",
					Description: `Tunnel configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_tunnel))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_nat",
					Description: `Tunnel configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_nat))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_vrrp",
					Description: `VRRP(Virtual Router Redundancy Protocol) configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_vrrp))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_acl_qos",
					Description: `ACL(Apply Access Lists) and QoS configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_acl_qos))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_arp",
					Description: `ARP(Address Resolution Protocol) configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_arp))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_8021x",
					Description: `IEEE 802.1X Authentication configuration for the Vedge VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_8021x))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_trustsec",
					Description: `TrustSec configuration for the Cisco VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_trustsec))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_advanced",
					Description: `Advanced configuration for the VPN interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_advanced)) <a id="nestedblock--vpn_interface_basic"></a> ## Nested Schema for vpn_interface_basic`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `Shutdown flag for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Interface Description for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `IPv4 configuration for VPN interface type Feature Template, (see [below for nested schema](#nestedblock--ipv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `IPv6 configuration for VPN interface type Feature Template, (see [below for nested schema](#nestedblock--ipv6))`,
				},
				resource.Attribute{
					Name:        "block_non_source_ip",
					Description: `Flag indicating whether forwarded traffic matches with interface's IP prefix range for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "bandwidth_upstream",
					Description: `Value of bandwidth above which to generate notifications in case of transmitted traffic for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "bandwidth_downstream",
					Description: `Value of bandwidth above which to generate notifications in case of received traffic for VPN interface type Feature Template <a id="nestedblock--vpn_interface_tunnel"></a> ## Nested Schema for vpn_interface_tunnel`,
				},
				resource.Attribute{
					Name:        "per_tunnel_qos",
					Description: `Flag indicating whether per_tunnel_qos is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "per_tunnel_qos_aggregator",
					Description: `Flag indicating whether per_tunnel_qos_aggregator is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "tunnels_bandwidth_percent",
					Description: `Value of tunnel bandwith percentage for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `TLOC color value for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "restrict",
					Description: `Restrict flag for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Value of Groups in tunnle for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "border",
					Description: `Flag indicating whether Border is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "control_connection",
					Description: `Flag indicating TLOC connection is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "maximum_control_connections",
					Description: `Maximum number of vSamart controllers that WAN tunnel interface can connect for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "vbond_as_stun_server",
					Description: `Flag indicating whether Session Traversal Utilities for NAT(STUN) is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "exclude_controller_group_list",
					Description: `List of vSmart controllers that the tunnel interface is not allowed to connect for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "vmanage_connection_preference",
					Description: `Flag indicating preference for using a tunnel interface to exchange control traffic with the vManage NMS for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "port_hop",
					Description: `Flag indicating whether Port Hop is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "low_bandwidth_link",
					Description: `Flag indicating whether tunnle is set as a low-bandwidth link for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "allow_service",
					Description: `Nested block for allow or disallow various services for VPN interface type Feature Template(see [below for nested schema](#nestedblock--allow_service))`,
				},
				resource.Attribute{
					Name:        "encapsulation",
					Description: `Encapsulation configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--encapsulation)) <a id="nestedblock--vpn_interface_nat"></a> ## Nested Schema for vpn_interface_nat`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `NAT IPv4 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--natv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `NAT IPv6 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--natv6)) <a id="nestedblock--vpn_interface_vrrp"></a> ## Nested Schema for vpn_interface_vrrp`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `VRRP IPv4 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--vrrpv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `VRRP IPv6 configuration for VPN interface type Feature Template(see [below for nested schema](#nestedblock--vrrpv6)) <a id="nestedblock--vpn_interface_acl_qos"></a> ## Nested Schema for vpn_interface_acl_qos`,
				},
				resource.Attribute{
					Name:        "adapt_period",
					Description: `Value of Adapt Period for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_min",
					Description: `Minimum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_max",
					Description: `Maximum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_default",
					Description: `Default value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_min",
					Description: `Minimum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_max",
					Description: `Maximum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_default",
					Description: `Default value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "shaping_rate",
					Description: `Value of aggregate traffic transmission rate on the VPN interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "qos_map",
					Description: `Name of the QoS map to apply to packets being transmitted out the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "rewrite_rule",
					Description: `Name of the rewrite rule to apply on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4_ingress_access_list",
					Description: `Name of the access list to apply to IPv4 packets being received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4_egress_access_list",
					Description: `Name of the access list to apply to IPv4 packets being transmitted on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv6_ingress_access_list",
					Description: `Name of the access list to apply to IPv6 packets being received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv6_egress_access_list",
					Description: `Name of the access list to apply to IPv6 packets being transmitted on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ingress_policer_name",
					Description: `Name of the policer to apply to packets received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "egress_policer_name",
					Description: `Name of the policer to apply to packets being transmitted on the interface for Vedge VPN interface type Feature Template <a id="nestedblock--vpn_interface_arp"></a> ## Nested Schema for vpn_interface_arp`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address for the ARP entry in dotted decimal notation or as a fully qualified host name for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `MAC address in colon-separated hexadecimal notation for VPN interface type Feature Template <a id="nestedblock--vpn_interface_8021x"></a> ## Nested Schema for vpn_interface_8021x`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `The tag of the RADIUS server to use for 802.1X authentication for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "account_interim_interval",
					Description: `Value of how often to send 802.1X interim accounting updates to the RADIUS server for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nas_identifier",
					Description: `The NAS identifier of the local router for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nas_ip",
					Description: `The NAS IP address of the local router for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "wake_on_lan",
					Description: `The flag that enable client to be powered up when the router receives an Ethernet magic packet frame for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "control_direction",
					Description: `Direction type of packets for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "reauthentication",
					Description: `Value indicating how often to reauthenticate 802.1X clients for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "inactivity",
					Description: `Value indicating how long to wait before revoking an 802.1X client's network access for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "host_mode",
					Description: `Value indicating whether an 802.1X interface grants access to a single client or to multiple clients for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "advanced_options",
					Description: `Advance configuration of vpn_interface_8021x for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--advanced_options)) <a id="nestedblock--vpn_interface_trustsec"></a> ## Nested Schema for vpn_interface_trustsec`,
				},
				resource.Attribute{
					Name:        "enable_sgt_propagation",
					Description: `Flag indicating whether SGT propagation is enabled or not for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "propagate",
					Description: `Flag indicating whether propagate is enabled or not for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "security_group_tag",
					Description: `Value of security group tag for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "trusted",
					Description: `Flag indicating whether trusted is enable or not for Cisco VPN interface type Feature Template <a id="nestedblock--vpn_interface_advanced"></a> ## Nested Schema for vpn_interface_advanced`,
				},
				resource.Attribute{
					Name:        "duplex",
					Description: `Value of duplex to run interface in auto, full, or half mode for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `Value of MAC address to associate with the interface, in colon-separated hexadecimal notation for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ip_mtu",
					Description: `Value of MAC address to associate with the interface, in colon-separated hexadecimal notation for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "pmtu_discovery",
					Description: `Flag indicating whether path MTU is enable or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "flow_control",
					Description: `Value of bidirectional flow control settings for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `Maximum segment size (MSS) of TPC SYN packets passing through the router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Value of speed to use when the remote end of the connection does not support autonegotiation for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "clear_dont_fragment",
					Description: `Flag indicating Don't Fragment (DF) bit in the IPv4 packet header for packets being transmitted out the interface is clar or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "static_ingess_qos",
					Description: `Queue number to use for incoming traffic for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "arp_timeout",
					Description: `Value indicating how long it takes for a dynamically learned ARP entry to time out for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "autonegotiation",
					Description: `Flag indicating autonegotiation mode for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "tloc_extension",
					Description: `Name of a physical interface on the same router that connects to the WAN transport for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "power_over_ethernet",
					Description: `Flag indicating whether PoE(Power over Ethernet) is enabled or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "load_interval",
					Description: `Value of load interval for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker",
					Description: `Name of a tracker to track the status of transport interfaces that connect to the internet for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "icmp_redirect_disable",
					Description: `Flag indicating whether ICMP redirect disable or not on interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "gre_tunnel_source_ip",
					Description: `IP address of the extended WAN interface`,
				},
				resource.Attribute{
					Name:        "xconnect",
					Description: `The name of a physical interface on the same router that connects to the WAN transport for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ip_directed_broadcast",
					Description: `Flag indicate whether directed IP braodcast is enabled or not for VPN interface type Feature Template <a id="nestedblock--ipv4"></a> ## Nested Schema for Basic ipv4`,
				},
				resource.Attribute{
					Name:        "primary_address",
					Description: `Value of IPv4 primary address for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "secondary_address",
					Description: `Value of IPv4 secondary address for VPN interface type Feature Template(see [below for nested schema](#nestedblock--secondary_addressv4))`,
				},
				resource.Attribute{
					Name:        "dhcp_distance",
					Description: `Value of IPv4 DHCP distance for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp_helper",
					Description: `Value of IPv4 DHCP Helper for VPN interface type Feature Template <a id="nestedblock--ipv6"></a> ## Nested Schema for Basic ipv6`,
				},
				resource.Attribute{
					Name:        "primary_address",
					Description: `Value of IPv6 primary address for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "secondary_address",
					Description: `Value of IPv6 secondary address for VPN interface type Feature Template(see [below for nested schema](#nestedblock--secondary_addressv6))`,
				},
				resource.Attribute{
					Name:        "dhcp_distance",
					Description: `Value of IPv6 DHCP distance for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp_helper",
					Description: `Value of IPv6 DHCP Helper for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp_rapid_commit",
					Description: `Value of IPv6 DHCP rapid commit enabled or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp_helper",
					Description: `Value of IPv6 DHCP rapid commit enabled or not for VPN interface type Feature Template(see [below for nested schema](#nestedblock--dhcp_helper)) <a id="nestedblock--allow_service"></a> ## Nested Schema for Tunnel allow_service`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `Flag indicating whether All serivce is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "bgp",
					Description: `Flag indicating whether BGP is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Flag indicating whether DHCP is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `Flag indicating whether DNS is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "icmp",
					Description: `Flag indicating whether ICMP is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "netconf",
					Description: `Flag indicating whether NETCONF is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ntp",
					Description: `Flag indicating whether NTP is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `Flag indicating whether OSPF is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ssh",
					Description: `Flag indicating whether SSH is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "stun",
					Description: `Flag indicating whether STUN is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "https",
					Description: `Flag indicating whether HTTPS is enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "snmp",
					Description: `Flag indicating whether SNMP is enabled or not for Cisco VPN interface type Feature Template <a id="nestedblock--encapsulation"></a> ## Nested Schema for Tunnel encapsulation`,
				},
				resource.Attribute{
					Name:        "gre_preference",
					Description: `Value of GRE preference for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "gre_weight",
					Description: `Value of GRE weight for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipsec_preference",
					Description: `Value of IPsec preference for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipsec_weight",
					Description: `Value of IPsec weight for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "carrier",
					Description: `Carrier name or private network identifier to associate with the tunnel for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "bind_loopback_tunnel",
					Description: `Name of a physical interface to bind to a loopback interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_resort_circuit",
					Description: `Flag indicating whether tunnel interface is used as the circuit of last resort for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `Value of hold time for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nat_refresh_interval",
					Description: `Value of interval between NAT refresh packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `Value of interval between Hello packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "hello_tolerance",
					Description: `Value of the time to wait for a Hello packet on a DTLS or TLS WAN transport connection before declaring that transport tunnel to be down for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "gre_tunnel_destination_ip",
					Description: `Value of GRE tunnel destination IP for VPN interface type Feature Template <a id="nestedblock--natv4"></a> ## Nested Schema for NAT ipv4`,
				},
				resource.Attribute{
					Name:        "nat_type",
					Description: `NAT type for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "refresh_mode",
					Description: `Refresh mode for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "log_nat_flow_creations_or_deletions",
					Description: `Flag indicating whether log nat flow creations or deletions are enabled or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `The time when NAT translations over UDP sessions time out for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `The time when NAT translations over TCP sessions time out for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "block_icmp",
					Description: `Flag indicating whether to block icmp or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "respond_to_ping",
					Description: `Flag indicating whether to respond to ping or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nat_pool_range_start",
					Description: `Value of starting IP address for the NAT pool for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nat_pool_range_end",
					Description: `Value of closing IP address for the NAT pool for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "overload",
					Description: `Flag indicating whether to enable per-port translation or not for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nat_inside_source_loopback_interface",
					Description: `Value of NAT inside source loopback interface for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "port_forward",
					Description: `Configure port forwarding rule for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--port_forward))`,
				},
				resource.Attribute{
					Name:        "static_nat",
					Description: `Configure static nat for VPN interface type Feature Template(see [below for nested schema](#nestedblock--static_nat)) <a id="nestedblock--natv6"></a> ## Nested Schema for NAT ipv6`,
				},
				resource.Attribute{
					Name:        "nat64",
					Description: `Flag indicating whether to enable nat64 or not for VPN interface type Feature Template <a id="nestedblock--vrrpv4"></a> ## Nested Schema for VRRP ipv4`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_omp",
					Description: `Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_prefix_list",
					Description: `(Required) Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Value of IPv4 address of the virtual router for VPN interface type Feature Template <a id="nestedblock--vrrpv6"></a> ## Nested Schema for VRRP ipv6`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_omp",
					Description: `Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "track_prefix_list",
					Description: `Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "link_local_ipv6_address",
					Description: `Value of IPv6 address of the virtual router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "global_ipv6_prefix",
					Description: `Value of global IPv6 prefix of the virtual router for VPN interface type Feature Template <a id="nestedblock--advanced_options"></a> ## Nested Schema for 8021.X advanced_options`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `Configure VLAN for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--vlan))`,
				},
				resource.Attribute{
					Name:        "dynamic_authentication_server",
					Description: `Configure Dynamic Authentication Server for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--dynamic_authentication_server))`,
				},
				resource.Attribute{
					Name:        "mac_authentication_bypass",
					Description: `Configure MAC Authentication Bypass for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--mac_authentication_bypass))`,
				},
				resource.Attribute{
					Name:        "request_attributes",
					Description: `Configure Request Attributes for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--request_attributes)) <a id="nestedblock--secondary_addressv4"></a> ## Nested Schema for Basic IPv4 secondary_address`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Value of IPv4 prefix for VPN interface type Feature Template <a id="nestedblock--secondary_addressv6"></a> ## Nested Schema for Basic IPv6 secondary_address`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Value of IPv6 prefix for VPN interface type Feature Template <a id="nestedblock--dhcp_helper"></a> ## Nested Schema for Basic IPv6 dhcp_helper`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Value of IPv6 address for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Value of VPN id for VPN interface type Feature Template <a id="nestedblock--port_forward"></a> ## Nested Schema for NAT ipv4 port_forward`,
				},
				resource.Attribute{
					Name:        "port_start_range",
					Description: `Value of port starting range for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "port_end_range",
					Description: `Value of port ending range for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `VPN id for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Required) Private IP address for Vedge VPN interface type Feature Template <a id="nestedblock--static_nat"></a> ## Nested Schema for NAT ipv4 static_nat`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Value of the inside local address as source IP address for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "translated_source_ip",
					Description: `Value of the inside global address as the translated source IP address for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "source_vpn_id",
					Description: `Configures Source VPN ID, which is the Service VPN in which the host resides for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "static_nat_direction",
					Description: `Set the direction in which to perform network address translation for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `Value of source port for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "translate_port",
					Description: `Value of translate port for Vedge VPN interface type Feature Template <a id="nestedblock--vlan"></a> ## Nested Schema for vlan`,
				},
				resource.Attribute{
					Name:        "authentication_fail_vlan",
					Description: `Value of authentication fail VLAN for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "guest_vlan",
					Description: `Value of guest VLAN for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "authentication_reject_vlan",
					Description: `Value of authentication reject VLAN for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "default_vlan",
					Description: `Value of default VLAN for Vedge VPN interface type Feature Template <a id="nestedblock--dynamic_authentication_server"></a> ## Nested Schema for dynamic_authentication_server`,
				},
				resource.Attribute{
					Name:        "das_port",
					Description: `Value of Dynamic Authentication Server port for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "client",
					Description: `Value of Client IP address for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Secret key Value for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "time_window",
					Description: `Time window value for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "required_timestamp",
					Description: `Flag indicating whether to enable required timestamp or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Value of VPN ID for Vedge VPN interface type Feature Template <a id="nestedblock--mac_authentication_bypass"></a> ## Nested Schema for mac_authentication_bypass`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Flag indicating whether to enable server or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "mac_authentication_bypass_entries",
					Description: `List of MAC Authenticaion bypass entries for Vedge VPN interface type Feature Template <a id="nestedblock--request_attributes"></a> ## Nested Schema for request_attributes`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Configure Authentication for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--authentication))`,
				},
				resource.Attribute{
					Name:        "accounting",
					Description: `Configure Accounting for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--accounting)) ## Nested Schema for request_attributes authentication`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "syntax_choice",
					Description: `Syntax Choice for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value for Vedge VPN interface type Feature Template <a id="nestedblock--accounting"></a> ## Nested Schema for request_attributes accounting`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "syntax_choice",
					Description: `Syntax Choice for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value for Vedge VPN interface type Feature Template`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"sdwan_device_template":                0,
		"sdwan_device_uuids":                   1,
		"sdwan_devices_attachment":             2,
		"sdwan_feature_template_ids":           3,
		"sdwan_ntp_feature_template":           4,
		"sdwan_system_feature_template":        5,
		"sdwan_vpn_feature_template":           6,
		"sdwan_vpn_interface_feature_template": 7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
