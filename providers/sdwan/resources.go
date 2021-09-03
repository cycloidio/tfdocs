package sdwan

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "sdwan_device_template",
			Category:         "Resources",
			ShortDescription: `Manages SD-WAN Device Templates`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) policyId for Device Template`,
				},
				resource.Attribute{
					Name:        "feature_template_uid_range",
					Description: `(Required) featureTemplateUidRange for Device Template`,
				},
				resource.Attribute{
					Name:        "general_templates",
					Description: `(Required) List of Feature Templates and thier Sub Templates thourgh which Device Template is created (should not be empty , (see [below for nested schema](#nestedblock--general_templates)))`,
				},
				resource.Attribute{
					Name:        "connection_preference_required",
					Description: `(Optional) connectionPreferenceRequired flag for Device Template, allowed values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "connection_preference",
					Description: `(Optional) connectionPreference flag for Device Template, allowed values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ` <a id="nestedblock--general_templates"></a> ## Nested Schema for general_templates ##`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Required) Template Id of Feature Template`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `(Required) Template Type of Feature Template, allowed values: "aaa", "bfd-vedge", "omp-vedge", "security-vedge", "system-vedge", "vpn-vedge", "cedge_aaa", "cisco_bfd", "cisco_omp", "cisco_security", "cisco_system", "cisco_vpn", "cedge_global" <br> <strong>NOTE</strong><br> ` + "`" + `aaa` + "`" + `, ` + "`" + `bfd-vedge` + "`" + `, ` + "`" + `omp-vedge` + "`" + `, ` + "`" + `security-vedge` + "`" + `, ` + "`" + `system-vedge` + "`" + `, ` + "`" + `vpn-vedge` + "`" + ` are required for vEdge Devices.<br> ` + "`" + `cedge_aaa` + "`" + `, ` + "`" + `cisco_bfd` + "`" + `, ` + "`" + `cisco_omp` + "`" + `, ` + "`" + `cisco_security` + "`" + `, ` + "`" + `cisco_system` + "`" + `, ` + "`" + `cisco_vpn` + "`" + `, ` + "`" + `cedge_global` + "`" + ` are required for cEdge Devices`,
				},
				resource.Attribute{
					Name:        "sub_templates",
					Description: `(Optional) List of Sub Templates associated with feature Template (see [below for nested schema](#nestedblock--sub_templates)) <a id="nestedblock--sub_templates"></a> ## Nested Schema for sub_templates ##`,
				},
				resource.Attribute{
					Name:        "sub_template_id",
					Description: `(Required) Template Id of Feature Template`,
				},
				resource.Attribute{
					Name:        "sub_template_type",
					Description: `(Required) Template Type of Feature Template, allowed values: "logging" (this can be allowed only with ` + "`" + `system-vedge` + "`" + ` type ` + "`" + `general_templates` + "`" + `), "ntp" (this can be allowed only with ` + "`" + `system-vedge` + "`" + ` type ` + "`" + `general_templates` + "`" + `), "vpn-vedge-interface" (this can be allowed only with` + "`" + `vpn-vedge` + "`" + ` type ` + "`" + `general_templates` + "`" + `), "cisco_logging" (this can be allowed only with ` + "`" + `cisco_system` + "`" + ` type ` + "`" + `general_templates` + "`" + `), "cisco_ntp"(this can be allowed only with ` + "`" + `"cisco_system"` + "`" + ` type ` + "`" + `general_templates` + "`" + `), "cisco_vpn_interface" (this can be allowed only with ` + "`" + `cisco_vpn` + "`" + ` type ` + "`" + `general_templates` + "`" + `) ## Argument Reference for Device Template created from CLI ##`,
				},
				resource.Attribute{
					Name:        "template_configuration",
					Description: `(Required) Template Configuration for Device Template ## Common Attributes ##`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Device Template id for Device Template`,
				},
				resource.Attribute{
					Name:        "template_class",
					Description: `Template Class type for Device Template ## Attributes for Device Template created from CLI ##`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who updated Device Template last`,
				},
				resource.Attribute{
					Name:        "last_updated_on",
					Description: `Time when Device Template was last updated`,
				},
				resource.Attribute{
					Name:        "rid",
					Description: `Request ID for Device Template`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Time when Device Template was created`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `User who created Device Template`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `Feature for Device Template`,
				},
				resource.Attribute{
					Name:        "template_attached",
					Description: `Number of templates attached to Device Template`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_devices_attachment",
			Category:         "Resources",
			ShortDescription: `Manages SD-WAN Devices Attachment`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Status of Device Attachment`,
				},
				resource.Attribute{
					Name:        "activity",
					Description: `List of Activities at the time of Device attachment`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_ntp_feature_template",
			Category:         "Resources",
			ShortDescription: `Manages SD-WAN NTP Type Feature Templates`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Unique NTP Type Feature Template Name, Must not include <, >, !, &, ", or white space; maximum 128 characters`,
				},
				resource.Attribute{
					Name:        "template_description",
					Description: `(Required) Long Description of NTP type Feature Template, Should not be empty`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Required) Type of device which supports NTP Feature Template, Allowed values for NTP: "vedge-1000", "vedge-2000", "vedge-cloud", "vedge-5000", "vedge-ISR1100-6G", "vedge-100-B", "vedge-ISR1100-4G", "vedge-100", "vedge-ISR1100-4GLTE", "vedge-100-WM", "vedge-100-M", "vedge-ISR1100X-6G", "vedge-ISR1100X-4G" Allowed values for Cisco NTP: "vedge-C8500-12X4QC", "vedge-C1111-4PLTEEA", "vedge-C1161-8P", "vedge-C1117-4PLTEEAW", "vedge-C1121X-8P", "vedge-C8200-1N-4T", "vedge-ISR-4331", "vedge-C1127X-8PMLTEP","vedge-C1117-4PMLTEEAWE", "vedge-ISR-4451-X", "vedge-C1113-8PLTEEA", "vedge-IR-1821", "vedge-ASR-1001-X", "vedge-ISR-4321","vedge-C1116-4PLTEEAWE", "vedge-C1109-4PLTE2P", "vedge-C1121-8P", "vedge-ASR-1002-HX", "vedge-C1111-8PLTEEAW", "vedge-C1112-8PWE", "vedge-C1101-4PLTEP", "vedge-ISR1100-4GLTENA-XE", "vedge-C1111-8PLTELA", "vedge-IR-1835", "vedge-C1121X-8PLTEP", "vedge-IR-1833", "vedge-C8300-1N1S-4T2X", "vedge-C1121-4P", "vedge-ISR-4351", "vedge-C1117-4PLTELA", "vedge-C1116-4PWE", "vedge-nfvis-C8200-UCPE", "vedge-C1113-8PM", "vedge-IR-1831", "vedge-C1127-8PLTEP","vedge-C1121-8PLTEPW", "vedge-C1113-8PW", "vedge-ASR-1001-HX", "vedge-C1128-8PLTEP", "vedge-C1113-8PLTEEAW", "vedge-C1117-4PW", "vedge-C1116-4P", "vedge-C1113-8PMLTEEA","vedge-C1112-8P", "vedge-ISR-4461", "vedge-C1116-4PLTEEA", "vedge-ISR-4221", "vedge-C1117-4PM", "vedge-C1113-8PLTELAWZ", "vedge-C1117-4PMWE", "vedge-C1109-2PLTEVZ", "vedge-C1113-8P", "vedge-C1117-4P", "vedge-nfvis-ENCS5400", "vedge-C8300-2N2S-6T", "vedge-C1127-8PMLTEP", "vedge-ESR-6300", "vedge-ISR-4221X", "vedge-ISR1100-4GLTEGB-XE", "vedge-C8500-12X", "vedge-C1109-2PLTEGB", "vedge-CSR-1000v", "vedge-C1113-8PLTEW", "vedge-C1121X-8PLTEPW", "vedge-ISR1100-6G-XE", "vedge-C1121-4PLTEP", "vedge-C1111-8PLTEEA", "vedge-C1117-4PLTEEA", "vedge-C1127X-8PLTEP", "vedge-C1109-2PLTEUS", "vedge-C1112-8PLTEEAWE", "vedge-C1161X-8P", "vedge-C8500L-8S4X", "vedge-C1111-8PW", "vedge-C1161X-8PLTEP", "vedge-C1101-4PLTEPW", "vedge-ISR1100X-4G-XE", "vedge-IR-1101", "vedge-C1111-4P", "vedge-C1111-4PW", "vedge-C1111-8P", "vedge-C1117-4PMLTEEA", "vedge-C1113-8PLTELA", "vedge-C1111-8PLTELAW", "vedge-C1161-8PLTEP", "vedge-ISR1100X-6G-XE", "vedge-ISR-4431", "vedge-C1101-4P", "vedge-C1109-4PLTE2PW", "vedge-C1113-8PMWE", "vedge-C1118-8P", "vedge-C1126-8PLTEP", "vedge-C8300-1N1S-6T", "vedge-C1121-8PLTEP", "vedge-C8300-2N2S-4T2X", "vedge-C1112-8PLTEEA", "vedge-C1111-4PLTELA", "vedge-ASR-1002-X", "vedge-C1111X-8P", "vedge-C1126X-8PLTEP", "vedge-ASR-1006-X", "vedge-C8000V", "vedge-ISR1100-4G-XE", "vedge-C1117-4PLTELAWZ", "vedge-ISRv"`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `(Required) Type of NTP Feature Template, Allowed values: "ntp", "cisco_ntp"`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `(Required) Minimum Version for the Feature template`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `(Required) Boolean flag indicating whether NTP type Feature Template is factory default or not, If we set it true we can not update or delete resource.`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `(Required) Configuration for NTP type Feature Template, (see [below for nested schema](#nestedblock--template_definition)) ## Attribute Reference ##`,
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
					Description: `(Optional) Server configuration for the NTP type Feature Template, see [below for nested schema](#nestedblock--server))`,
				},
				resource.Attribute{
					Name:        "master",
					Description: `(Optional) Master configuration for the Cisco NTP type Feature Template, see [below for nested schema](#nestedblock--master))`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) Authentication configuration for the NTP type Feature Template, see [below for nested schema](#nestedblock--authentication))`,
				},
				resource.Attribute{
					Name:        "trusted",
					Description: `(Optional) Enter the MD5 authentication key id to designate the key as trustworthy <a id="nestedblock--server"></a> ## Nested Schema for server`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) IP address of an NTP server, or a DNS server that knows how to reach the NTP server`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) MD5 key associated with the NTP server, to enable MD5 authentication, Range: [1, 65535]`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Required) Vpn ID for configuration of NTP type Feature Template, Range: [1, 65530]`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version number of the NTP protocol software, Range: [1, 4], Default: 4`,
				},
				resource.Attribute{
					Name:        "source_interface",
					Description: `(Optional) Name of a specific interface to use for outgoing NTP packets`,
				},
				resource.Attribute{
					Name:        "prefer",
					Description: `(Optional) It is a boolean flag, set true if multiple NTP servers are at the same stratum level and you want one to be preferred <a id="nestedblock--master"></a> ## Nested Schema for master`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Boolean flag indicating whether master is enabled for configuration of Cisco NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "stratum",
					Description: `(Optional) Stratum for configuration of Cisco NTP type Feature Template, Range: [1, 15]`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Source for configuration of Cisco NTP type Feature Template <a id="nestedblock--authentication"></a> ## Nested Schema for authentication`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) MD5 authentication key ID for configuration of NTP type Feature Template, Range: [1, 65535]`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) It is either a cleartext key or an AES-encrypted key for authentication of NTP server`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(Optional) Server configuration for the NTP type Feature Template, see [below for nested schema](#nestedblock--server))`,
				},
				resource.Attribute{
					Name:        "master",
					Description: `(Optional) Master configuration for the Cisco NTP type Feature Template, see [below for nested schema](#nestedblock--master))`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) Authentication configuration for the NTP type Feature Template, see [below for nested schema](#nestedblock--authentication))`,
				},
				resource.Attribute{
					Name:        "trusted",
					Description: `(Optional) Enter the MD5 authentication key id to designate the key as trustworthy <a id="nestedblock--server"></a> ## Nested Schema for server`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) IP address of an NTP server, or a DNS server that knows how to reach the NTP server`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) MD5 key associated with the NTP server, to enable MD5 authentication, Range: [1, 65535]`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Required) Vpn ID for configuration of NTP type Feature Template, Range: [1, 65530]`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version number of the NTP protocol software, Range: [1, 4], Default: 4`,
				},
				resource.Attribute{
					Name:        "source_interface",
					Description: `(Optional) Name of a specific interface to use for outgoing NTP packets`,
				},
				resource.Attribute{
					Name:        "prefer",
					Description: `(Optional) It is a boolean flag, set true if multiple NTP servers are at the same stratum level and you want one to be preferred <a id="nestedblock--master"></a> ## Nested Schema for master`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Boolean flag indicating whether master is enabled for configuration of Cisco NTP type Feature Template`,
				},
				resource.Attribute{
					Name:        "stratum",
					Description: `(Optional) Stratum for configuration of Cisco NTP type Feature Template, Range: [1, 15]`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Source for configuration of Cisco NTP type Feature Template <a id="nestedblock--authentication"></a> ## Nested Schema for authentication`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) MD5 authentication key ID for configuration of NTP type Feature Template, Range: [1, 65535]`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) It is either a cleartext key or an AES-encrypted key for authentication of NTP server`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_system_feature_template",
			Category:         "Resources",
			ShortDescription: `Manages SD-WAN System Type Feature Templates`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Unique System Type Feature Template Name, Must not include <, >, !, &, ", or white space; maximum 128 characters`,
				},
				resource.Attribute{
					Name:        "template_description",
					Description: `(Required) Long Description of System type Feature Template, String length: [1, 2048]`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Required) Type of device which supports Feature Template, Allowed values for template_type = "System": "vbond","vedge-100","vedge-100-B","vedge-100-M","vedge-100-WM","vedge-1000","vedge-2000","vedge-5000","vedge-cloud","vedge-ISR1100-4G","vedge-ISR1100-4GLTE","vedge-ISR1100-6G","vedge-ISR1100X-4G","vedge-ISR1100X-6G","vedge-nfvis-CSP-5444","vedge-nfvis-CSP-5456","vedge-nfvis-CSP2100","vedge-nfvis-CSP2100-X1","vedge-nfvis-CSP2100-X2","vedge-nfvis-UCSC-E","vedge-nfvis-UCSC-M5". Allowed values for template_type = "Cisco System" are: "vedge-C8500-12X4QC", "vedge-C1111-4PLTEEA", "vedge-C1161-8P", "vedge-C1117-4PLTEEAW", "vedge-C1121X-8P", "vedge-C8200-1N-4T", "vedge-ISR-4331", "vedge-ISRv", "cellular-gateway-CG522-E", "vedge-C1127X-8PMLTEP", "vedge-C1117-4PMLTEEAWE", "vedge-ISR-4451-X", "vedge-C1113-8PLTEEA", "vedge-IR-1821", "vedge-ASR-1001-X", "vedge-ISR-4321", "vedge-C1116-4PLTEEAWE", "vedge-C1109-4PLTE2P", "vedge-C1121-8P", "vedge-ASR-1002-HX", "cellular-gateway-CG418-E", "vedge-C1111-8PLTEEAW", "vedge-C1112-8PWE", "vedge-C1101-4PLTEP", "vedge-ISR1100-4GLTENA-XE", "vedge-C1111-8PLTELA", "vedge-IR-1835", "vedge-C1121X-8PLTEP", "vedge-IR-1833", "vedge-C8300-1N1S-4T2X", "vedge-C1121-4P", "vedge-ISR-4351", "vedge-C1117-4PLTELA", "vedge-C1116-4PWE", "vedge-nfvis-C8200-UCPE", "vedge-C1113-8PM", "vedge-IR-1831", "vedge-C1127-8PLTEP", "vedge-C1121-8PLTEPW", "vedge-C1113-8PW", "vedge-ASR-1001-HX", "vedge-C1128-8PLTEP", "vedge-C1113-8PLTEEAW", "vedge-C1117-4PW", "vedge-C1116-4P", "vedge-C1113-8PMLTEEA", "vedge-C1112-8P", "vedge-ISR-4461", "vedge-C1116-4PLTEEA", "vedge-ISR-4221", "vedge-C1117-4PM", "vedge-C1113-8PLTELAWZ", "vedge-C1117-4PMWE", "vedge-C1109-2PLTEVZ", "vedge-C1113-8P", "vedge-C1117-4P", "vedge-nfvis-ENCS5400", "vedge-C8300-2N2S-6T", "vedge-C1127-8PMLTEP", "vedge-ESR-6300", "vedge-ISR-4221X", "vedge-ISR1100-4GLTEGB-XE", "vedge-C8500-12X", "vedge-C1109-2PLTEGB", "vedge-CSR-1000v", "vedge-C1113-8PLTEW", "vedge-C1121X-8PLTEPW", "vedge-ISR1100-6G-XE", "vedge-C1121-4PLTEP", "vedge-C1111-8PLTEEA", "vedge-C1117-4PLTEEA", "vedge-C1127X-8PLTEP", "vedge-C1109-2PLTEUS", "vedge-C1112-8PLTEEAWE", "vedge-C1161X-8P", "vedge-C8500L-8S4X", "vedge-C1111-8PW", "vedge-C1161X-8PLTEP", "vedge-C1101-4PLTEPW", "vedge-ISR1100X-4G-XE", "vedge-IR-1101", "vedge-C1111-4P", "vedge-C1111-4PW", "vedge-C1111-8P", "vedge-C1117-4PMLTEEA", "vedge-C1113-8PLTELA", "vedge-C1111-8PLTELAW", "vedge-C1161-8PLTEP", "vedge-ISR1100X-6G-XE", "vedge-ISR-4431", "vedge-C1101-4P", "vedge-C1109-4PLTE2PW", "vedge-C1113-8PMWE", "vedge-C1118-8P", "vedge-C1126-8PLTEP", "vedge-C8300-1N1S-6T", "vedge-C1121-8PLTEP", "vedge-C8300-2N2S-4T2X", "vedge-C1112-8PLTEEA", "vedge-C1111-4PLTELA", "vedge-ASR-1002-X", "vedge-C1111X-8P", "vedge-C1126X-8PLTEP", "vedge-ASR-1006-X", "vedge-C8000V", "vedge-ISR1100-4G-XE", "vedge-C1117-4PLTELAWZ"`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `(Required) Minimum Version for the Feature template`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `(Required) Configuration for System type Feature Template, (see [below for nested schema](#nestedblock--template_definition))`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `(Required) Boolean flag indicating whether System type Feature Template is factory default or not ## Attribute Reference ##`,
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
					Description: `(Required) Basic configuration for the System type Feature Template, see [below for nested schema](#nestedblock--system_basic))`,
				},
				resource.Attribute{
					Name:        "system_gps",
					Description: `(Optional) GPS configuration for the System type Feature Template, see [below for nested schema](#nestedblock--system_gps))`,
				},
				resource.Attribute{
					Name:        "system_trackers",
					Description: `(Optional) Trackers for the System type Feature Template, see [below for nested schema](#nestedblock--system_trackers))`,
				},
				resource.Attribute{
					Name:        "system_advanced",
					Description: `(Optional) Advanced configuration for the System type Feature Template, see [below for nested schema](#nestedblock--system_advanced)) <a id="nestedblock--system_basic"></a> ## Nested Schema for system_basic`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) Timezone for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "overlay_id",
					Description: `(Optional) Overlay ID for configuration of System type Feature Template, Default: 1`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Location for configuration of System type Feature Template, String length: [1, 128]`,
				},
				resource.Attribute{
					Name:        "device_groups",
					Description: `(Optional) Device Groups for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "controller_groups",
					Description: `(Optional) Controller Group List for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "console_baud_rate",
					Description: `(Optional) Console Baud Rate for configuration of System type Feature Template, Allowed values: "1200","2400","4800","9600","19200","38400","57600","115200", Default: 9600`,
				},
				resource.Attribute{
					Name:        "max_omp_sessions",
					Description: `(Optional) Maximum OMP Sessions for configuration of System type Feature Template, Range: [0, 100], Default: 2`,
				},
				resource.Attribute{
					Name:        "tcp_optimization_enabled",
					Description: `(Optional) Boolean flag indicating whether TCP Optimization is enabled for configuration of System type Feature Template, Default: false <a id="nestedblock--system_gps"></a> ## Nested Schema for system_gps`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Optional) Latitude for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Optional) Longitude for configuration of System type Feature Template <a id="nestedblock--system_trackers"></a> ## Nested Schema for system_trackers`,
				},
				resource.Attribute{
					Name:        "tracker_name",
					Description: `(Required) Name of Tracker for configuration of System type Feature Template, Must not include <, >, !, &, ", or white space; maximum 128 characters`,
				},
				resource.Attribute{
					Name:        "tracker_endpoint_type",
					Description: `(Required) Type of Tracker Endpoint for configuration of System type Feature Template, Allowed values: "ip-address", "dns-name", "api-url"`,
				},
				resource.Attribute{
					Name:        "tracker_endpoint_value",
					Description: `(Required) Value of Tracker Endpoint for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_type",
					Description: `(Optional) Type of Tracker for configuration of Cisco System type Feature Template, Allowed Values: "interface", "static-route", Default: "interface"`,
				},
				resource.Attribute{
					Name:        "tracker_threshold",
					Description: `(Optional) Threshold of Tracker for configuration of System type Feature Template, Range: [100, 1000], Default: 300`,
				},
				resource.Attribute{
					Name:        "tracker_interval",
					Description: `(Optional) Interval of Tracker for configuration of System type Feature Template, Range: [10, 600], Default: 60`,
				},
				resource.Attribute{
					Name:        "tracker_multiplier",
					Description: `(Optional) Multiplier of Tracker for configuration of System type Feature Template, Range: [1, 10], Default: 3 <a id="nestedblock--system_advanced"></a> ## Nested Schema for system_advanced`,
				},
				resource.Attribute{
					Name:        "control_session_pps",
					Description: `(Optional) Policer Rate(pps) for Control Sessions for configuration of System type Feature Template, Range: [1, 65535], Default: 300`,
				},
				resource.Attribute{
					Name:        "system_tunnel_mtu",
					Description: `(Optional) MTU of System's internal DTLS Tunnel for configuration of System type Feature Template, Range: [500, 2000], Default: 1024`,
				},
				resource.Attribute{
					Name:        "port_hop",
					Description: `(Optional) Boolean flag to indicate whether Port Hopping is enabled for configuration of System type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "port_offset",
					Description: `(Optional) TLOC Port Offset for configuration of System type Feature Template, Range: [0, 19]`,
				},
				resource.Attribute{
					Name:        "dns_cache_timeout",
					Description: `(Optional) DNS Cache Timeout(minutes) for configuration of System type Feature Template, Range: [1, 30], Default: 2`,
				},
				resource.Attribute{
					Name:        "track_transport",
					Description: `(Optional) Boolean flag to indicate whether tracking of transport is enabled for configuration of System type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "vbond_local",
					Description: `(Optional) Boolean flag to indicate whether the local device is set as vBond for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "track_interface_tag",
					Description: `(Optional) Interface tracking for configuration of System type Feature Template, Range: [1, 4294967295]`,
				},
				resource.Attribute{
					Name:        "multicast_buffer_percent",
					Description: `(Optional) Percentage of buffer multicast packets can consume for configuration of System type Feature Template, Range: [5, 100]`,
				},
				resource.Attribute{
					Name:        "usb_controller",
					Description: `(Optional) Boolean flag to indicate whether external USB Controller on the device is enabled for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "track_default-gateway",
					Description: `(Optional) Boolean flag to indicate whether default Gateway Tracking on the device is enabled for configuration of System type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "host_policer_pps",
					Description: `(Optional) Rate(1000-25000) at which to police packets bound to the control plane for configuration of System type Feature Template, Range: [1000, 25000], Default: 20000`,
				},
				resource.Attribute{
					Name:        "icmp_error_pps",
					Description: `(Optional) Rate at which to police ICMP error messages for configuration of System type Feature Template, Range: [1, 200], Default: 100`,
				},
				resource.Attribute{
					Name:        "allow_same_site_tunnels",
					Description: `(Optional) Boolean flag to indicate whether tunnels are allowed to be formed between vEdges in the same site for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "route_consistency_check",
					Description: `(Optional) Boolean flag to indicate whether route consistency check is enabled for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "admin_tech_on_failure",
					Description: `(Optional) Boolean flag to indicate whether the collection of admin-tech before reboot due to daemon failure is enabled for configuration of System type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Idle CLI timeout(minutes) for configuration of System type Feature Template, Range: [0, 300]`,
				},
				resource.Attribute{
					Name:        "eco_friendly_mode",
					Description: `(Optional) Boolean flag to indicate whether vEdge Cloud router can run in the eco-friendly mode for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "on_demand_enable",
					Description: `(Optional) Boolean flag to indicate whether On-Demand Tunnel is enabled for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "on_demand_idle_timeout",
					Description: `(Optional) Idle timeout for On-demand Tunnels for configuration of System type Feature Template, Default: 10`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(Required) Basic configuration for the System type Feature Template, see [below for nested schema](#nestedblock--system_basic))`,
				},
				resource.Attribute{
					Name:        "system_gps",
					Description: `(Optional) GPS configuration for the System type Feature Template, see [below for nested schema](#nestedblock--system_gps))`,
				},
				resource.Attribute{
					Name:        "system_trackers",
					Description: `(Optional) Trackers for the System type Feature Template, see [below for nested schema](#nestedblock--system_trackers))`,
				},
				resource.Attribute{
					Name:        "system_advanced",
					Description: `(Optional) Advanced configuration for the System type Feature Template, see [below for nested schema](#nestedblock--system_advanced)) <a id="nestedblock--system_basic"></a> ## Nested Schema for system_basic`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) Timezone for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "overlay_id",
					Description: `(Optional) Overlay ID for configuration of System type Feature Template, Default: 1`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Location for configuration of System type Feature Template, String length: [1, 128]`,
				},
				resource.Attribute{
					Name:        "device_groups",
					Description: `(Optional) Device Groups for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "controller_groups",
					Description: `(Optional) Controller Group List for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "console_baud_rate",
					Description: `(Optional) Console Baud Rate for configuration of System type Feature Template, Allowed values: "1200","2400","4800","9600","19200","38400","57600","115200", Default: 9600`,
				},
				resource.Attribute{
					Name:        "max_omp_sessions",
					Description: `(Optional) Maximum OMP Sessions for configuration of System type Feature Template, Range: [0, 100], Default: 2`,
				},
				resource.Attribute{
					Name:        "tcp_optimization_enabled",
					Description: `(Optional) Boolean flag indicating whether TCP Optimization is enabled for configuration of System type Feature Template, Default: false <a id="nestedblock--system_gps"></a> ## Nested Schema for system_gps`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Optional) Latitude for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Optional) Longitude for configuration of System type Feature Template <a id="nestedblock--system_trackers"></a> ## Nested Schema for system_trackers`,
				},
				resource.Attribute{
					Name:        "tracker_name",
					Description: `(Required) Name of Tracker for configuration of System type Feature Template, Must not include <, >, !, &, ", or white space; maximum 128 characters`,
				},
				resource.Attribute{
					Name:        "tracker_endpoint_type",
					Description: `(Required) Type of Tracker Endpoint for configuration of System type Feature Template, Allowed values: "ip-address", "dns-name", "api-url"`,
				},
				resource.Attribute{
					Name:        "tracker_endpoint_value",
					Description: `(Required) Value of Tracker Endpoint for configuration of System type Feature Template`,
				},
				resource.Attribute{
					Name:        "tracker_type",
					Description: `(Optional) Type of Tracker for configuration of Cisco System type Feature Template, Allowed Values: "interface", "static-route", Default: "interface"`,
				},
				resource.Attribute{
					Name:        "tracker_threshold",
					Description: `(Optional) Threshold of Tracker for configuration of System type Feature Template, Range: [100, 1000], Default: 300`,
				},
				resource.Attribute{
					Name:        "tracker_interval",
					Description: `(Optional) Interval of Tracker for configuration of System type Feature Template, Range: [10, 600], Default: 60`,
				},
				resource.Attribute{
					Name:        "tracker_multiplier",
					Description: `(Optional) Multiplier of Tracker for configuration of System type Feature Template, Range: [1, 10], Default: 3 <a id="nestedblock--system_advanced"></a> ## Nested Schema for system_advanced`,
				},
				resource.Attribute{
					Name:        "control_session_pps",
					Description: `(Optional) Policer Rate(pps) for Control Sessions for configuration of System type Feature Template, Range: [1, 65535], Default: 300`,
				},
				resource.Attribute{
					Name:        "system_tunnel_mtu",
					Description: `(Optional) MTU of System's internal DTLS Tunnel for configuration of System type Feature Template, Range: [500, 2000], Default: 1024`,
				},
				resource.Attribute{
					Name:        "port_hop",
					Description: `(Optional) Boolean flag to indicate whether Port Hopping is enabled for configuration of System type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "port_offset",
					Description: `(Optional) TLOC Port Offset for configuration of System type Feature Template, Range: [0, 19]`,
				},
				resource.Attribute{
					Name:        "dns_cache_timeout",
					Description: `(Optional) DNS Cache Timeout(minutes) for configuration of System type Feature Template, Range: [1, 30], Default: 2`,
				},
				resource.Attribute{
					Name:        "track_transport",
					Description: `(Optional) Boolean flag to indicate whether tracking of transport is enabled for configuration of System type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "vbond_local",
					Description: `(Optional) Boolean flag to indicate whether the local device is set as vBond for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "track_interface_tag",
					Description: `(Optional) Interface tracking for configuration of System type Feature Template, Range: [1, 4294967295]`,
				},
				resource.Attribute{
					Name:        "multicast_buffer_percent",
					Description: `(Optional) Percentage of buffer multicast packets can consume for configuration of System type Feature Template, Range: [5, 100]`,
				},
				resource.Attribute{
					Name:        "usb_controller",
					Description: `(Optional) Boolean flag to indicate whether external USB Controller on the device is enabled for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "track_default-gateway",
					Description: `(Optional) Boolean flag to indicate whether default Gateway Tracking on the device is enabled for configuration of System type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "host_policer_pps",
					Description: `(Optional) Rate(1000-25000) at which to police packets bound to the control plane for configuration of System type Feature Template, Range: [1000, 25000], Default: 20000`,
				},
				resource.Attribute{
					Name:        "icmp_error_pps",
					Description: `(Optional) Rate at which to police ICMP error messages for configuration of System type Feature Template, Range: [1, 200], Default: 100`,
				},
				resource.Attribute{
					Name:        "allow_same_site_tunnels",
					Description: `(Optional) Boolean flag to indicate whether tunnels are allowed to be formed between vEdges in the same site for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "route_consistency_check",
					Description: `(Optional) Boolean flag to indicate whether route consistency check is enabled for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "admin_tech_on_failure",
					Description: `(Optional) Boolean flag to indicate whether the collection of admin-tech before reboot due to daemon failure is enabled for configuration of System type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Idle CLI timeout(minutes) for configuration of System type Feature Template, Range: [0, 300]`,
				},
				resource.Attribute{
					Name:        "eco_friendly_mode",
					Description: `(Optional) Boolean flag to indicate whether vEdge Cloud router can run in the eco-friendly mode for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "on_demand_enable",
					Description: `(Optional) Boolean flag to indicate whether On-Demand Tunnel is enabled for configuration of System type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "on_demand_idle_timeout",
					Description: `(Optional) Idle timeout for On-demand Tunnels for configuration of System type Feature Template, Default: 10`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_vpn_feature_template",
			Category:         "Resources",
			ShortDescription: `Manages SD-WAN VPN Type Feature Templates`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Unique VPN Type Feature Template Name, Must not include <, >, !, &, ", or white space; maximum 128 characters`,
				},
				resource.Attribute{
					Name:        "template_description",
					Description: `(Required) Long Description of System type Feature Template, Should not be empty`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `(Required) Template type for VPN type of Feature template, allowed values "vpn-vedge", "cisco_vpn"`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Required) Type of device which supports Feature Template, allowed values for template_type ` + "`" + `vpn-vedge` + "`" + ` "vedge-1000", "vedge-2000", "vedge-cloud", "vedge-5000", "vedge-ISR1100-6G", "vedge-100-B", "vedge-ISR1100-4G", "vedge-100", "vedge-ISR1100-4GLTE", "vedge-100-WM", "vbond", "vedge-100-M", "vedge-ISR1100X-6G", "vedge-ISR1100X-4G", allowed values for template_type ` + "`" + `cisco_vpn` + "`" + ` "vedge-C8500-12X4QC", "vedge-C1111-4PLTEEA", "vedge-C1161-8P", "vedge-C1117-4PLTEEAW", "vedge-C1121X-8P", "vedge-C8200-1N-4T", "vedge-ISR-4331", "vedge-ISRv", "vedge-C1127X-8PMLTEP", "vedge-C1117-4PMLTEEAWE", "vedge-ISR-4451-X", "vedge-C1113-8PLTEEA", "vedge-IR-1821", "vedge-ASR-1001-X", "vedge-ISR-4321", "vedge-C1116-4PLTEEAWE", "vedge-C1109-4PLTE2P", "vedge-C1121-8P", "vedge-ASR-1002-HX", "vedge-C1111-8PLTEEAW", "vedge-C1112-8PWE", "vedge-C1101-4PLTEP", "vedge-ISR1100-4GLTENA-XE", "vedge-C1111-8PLTELA", "vedge-IR-1835", "vedge-C1121X-8PLTEP", "vedge-IR-1833", "vedge-C8300-1N1S-4T2X", "vedge-C1121-4P", "vedge-ISR-4351", "vedge-C1117-4PLTELA", "vedge-C1116-4PWE", "vedge-nfvis-C8200-UCPE", "vedge-C1113-8PM", "vedge-IR-1831", "vedge-C1127-8PLTEP", "vedge-C1121-8PLTEPW", "vedge-C1113-8PW", "vedge-ASR-1001-HX", "vedge-C1128-8PLTEP", "vedge-C1113-8PLTEEAW", "vedge-C1117-4PW", "vedge-C1116-4P", "vedge-C1113-8PMLTEEA", "vedge-C1112-8P", "vedge-ISR-4461", "vedge-C1116-4PLTEEA", "vedge-ISR-4221", "vedge-C1117-4PM", "vedge-C1117-4PM", "vedge-C1113-8PLTELAWZ", "vedge-C1117-4PMWE", "vedge-C1109-2PLTEVZ", "vedge-C1113-8P", "vedge-C1117-4P", "vedge-nfvis-ENCS5400", "vedge-C8300-2N2S-6T", "vedge-C1127-8PMLTEP", "vedge-ESR-6300", "vedge-ISR-4221X" ,"vedge-ISR1100-4GLTEGB-XE vedge-C8500-12X", "vedge-C1109-2PLTEGB", "vedge-CSR-1000v", "vedge-C1113-8PLTEW", "vedge-C1121X-8PLTEPW", "vedge-ISR1100-6G-XE", "vedge-C1121-4PLTEP", "vedge-C1111-8PLTEEA", "vedge-C1117-4PLTEEA", "vedge-C1127X-8PLTEP", "vedge-C1109-2PLTEUS", "vedge-C1112-8PLTEEAWE", "vedge-C1161X-8P", "vedge-C8500L-8S4X", "vedge-C1111-8PW", "vedge-C1161X-8PLTEP", "vedge-C1101-4PLTEPW", "vedge-ISR1100X-4G-XE", "vedge-IR-1101", "vedge-C1111-4P", "vedge-C1111-4PW", "vedge-C1111-8P", "vedge-C1117-4PMLTEEA", "vedge-C1113-8PLTELA", "vedge-C1111-8PLTELAW", "vedge-C1161-8PLTEP", "vedge-ISR1100X-6G-XE", "vedge-ISR-4431", "vedge-C1101-4P", "vedge-C1109-4PLTE2PW", "vedge-C1113-8PMWE", "vedge-C1118-8P", "vedge-C1126-8PLTEP", "vedge-C8300-1N1S-6T", "vedge-C1121-8PLTEP", "vedge-C8300-2N2S-4T2X", "vedge-C1112-8PLTEEA", "vedge-C1111-4PLTELA", "vedge-ASR-1002-X", "vedge-C1111X-8P", "vedge-C1126X-8PLTEP", "vedge-ASR-1006-X", "vedge-C8000V", "vedge-ISR1100-4G-XE", "vedge-C1117-4PLTELAWZ"`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `(Required) Minimum Version for the Feature template`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `(Required) Configuration for VPN type Feature Template, (see [below for nested schema](#nestedblock--template_definition))`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `(Required) Boolean flag indicating whether VPN type Feature Template is factory default or not ## Attribute Reference ##`,
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
					Description: `(Required) Basic configuration for the VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_basic))`,
				},
				resource.Attribute{
					Name:        "vpn_dns",
					Description: `(Optional) DNS configuration for the VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_dns))`,
				},
				resource.Attribute{
					Name:        "ipv4_route",
					Description: `(Optional) IPv4 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv4_route))`,
				},
				resource.Attribute{
					Name:        "ipv6_route",
					Description: `(Optional) IPv6 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv6_route))`,
				},
				resource.Attribute{
					Name:        "te_service_enabled",
					Description: `(Optional) Flag indicating TE service is enabled or not for the VPN type Feature Template, can be enabled when ` + "`" + `vpn_id` + "`" + ` is 0, allowed values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `, default value: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "service_route",
					Description: `(Optional) Service route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--service_route))`,
				},
				resource.Attribute{
					Name:        "nat64",
					Description: `(Optional) NAT64 for the VPN type Feature Template, (see [below for nested schema](#nestedblock--nat64)), can be configured for ` + "`" + `template_type` + "`" + ` "vpn-vedge" <a id="nestedblock--vpn_basic"></a> ## Nested Schema for vpn_basic`,
				},
				resource.Attribute{
					Name:        "vpn_id",
					Description: `(Required) Numeric identifier of the VPN for VPN type of Feature Template, allowed values: 1, 512`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of VPN for VPN type of Feature Template`,
				},
				resource.Attribute{
					Name:        "ecmp_hash_key",
					Description: `(Optional) Boolean flag indicating whether ECMP hash key is enabled or not, allowed values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `, default value: ` + "`" + `false` + "`" + ` <a id="nestedblock--vpn_dns"></a> ## Nested Schema for vpn_dns`,
				},
				resource.Attribute{
					Name:        "primary_dns_ipv4",
					Description: `(Optional) Primary IPv4 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "secondary_dns_ipv4",
					Description: `(Optional) Secondary IPv4 address of DNS server for the VPN, can be set if ` + "`" + `primary_dns_ipv4` + "`" + ` is set`,
				},
				resource.Attribute{
					Name:        "primary_dns_ipv6",
					Description: `(Optional) Primary IPv6 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "secondary_dns_ipv6",
					Description: `(Optional) Secondary IPv6 address of DNS server for the VPN, can be set if ` + "`" + `primary_dns_ipv6` + "`" + ` is set`,
				},
				resource.Attribute{
					Name:        "vpn_host",
					Description: `(Optional) VPN host list for DNS server (see [below for nested schema](#nestedblock--vpn_host)) <a id="nestedblock--ipv4_route"></a> ## Nested Schema for ipv4_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) IPv4 prefix or address for IPv4 route`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `(Required) Gateway Type for next hop to reach the static route, allowed values: "next-hop", "dhcp", "null0"`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop list (see [below for nested schema](#nestedblock--next_hopv4)) when gateway_type is ` + "`" + `next-hop` + "`" + ``,
				},
				resource.Attribute{
					Name:        "null0_distance",
					Description: `(Optional) Null0 distance when gateway_type is ` + "`" + `null0` + "`" + `, range: [1 - 255] <a id="nestedblock--ipv6_route"></a> ## Nested Schema for ipv6_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) IPv6 prefix or address for IPv6 route`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `(Required) Gateway Type for next hop to reach the static route, allowed values: "next-hop", "null0"`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop list (see [below for nested schema](#nestedblock--next_hopv6)) when gateway_type is ` + "`" + `next-hop` + "`" + ``,
				},
				resource.Attribute{
					Name:        "null0_distance",
					Description: `(Optional) Null0 distance when gateway_type is ` + "`" + `null0` + "`" + `, cannot be set for ` + "`" + `template_type` + "`" + ` "cisco_vpn", range: [1 - 255] <a id="nestedblock--service_route"></a> ## Nested Schema for service_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) IPv4 prefix or address for service route`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Optional) Type of the service, allowed values for ` + "`" + `template_type` + "`" + ` "vpn-vedge": "sig", "FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "TE", allowed values for ` + "`" + `template_type` + "`" + ` "cisco_vpn": "FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "TE" <a id="nestedblock--nat64"></a> ## Nested Schema for nat64`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `(Required) NAT64 pool name`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) Starting IPv4 address of NAT64 pool`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) Ending IPv4 address of NAT64 pool`,
				},
				resource.Attribute{
					Name:        "overload",
					Description: `(Optional) Flag indicating NAT64 pool overload, allowed valus: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `, default values: ` + "`" + `false` + "`" + ` <a id="nestedblock--vpn_host"></a> ## Nested Schema for vpn_host`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Host name of the DNS server`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) List of ip addresses associated with the hostname, maximum length: 8 <a id="nestedblock--next_hopv4"></a> ## Nested Schema for next_hop`,
				},
				resource.Attribute{
					Name:        "next_hop_address",
					Description: `(Required) IPv4 address of the nexthop`,
				},
				resource.Attribute{
					Name:        "next_hop_distance",
					Description: `(Optional) Distance of the nexthop, range: [1 - 255] <a id="nestedblock--next_hopv6"></a> ## Nested Schema for next_hop`,
				},
				resource.Attribute{
					Name:        "next_hop_address",
					Description: `(Required) IPv6 address of the nexthop`,
				},
				resource.Attribute{
					Name:        "next_hop_distance",
					Description: `(Optional) Distance of the nexthop, range: [1 - 255]`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(Required) Basic configuration for the VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_basic))`,
				},
				resource.Attribute{
					Name:        "vpn_dns",
					Description: `(Optional) DNS configuration for the VPN type Feature Template, (see [below for nested schema](#nestedblock--vpn_dns))`,
				},
				resource.Attribute{
					Name:        "ipv4_route",
					Description: `(Optional) IPv4 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv4_route))`,
				},
				resource.Attribute{
					Name:        "ipv6_route",
					Description: `(Optional) IPv6 route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--ipv6_route))`,
				},
				resource.Attribute{
					Name:        "te_service_enabled",
					Description: `(Optional) Flag indicating TE service is enabled or not for the VPN type Feature Template, can be enabled when ` + "`" + `vpn_id` + "`" + ` is 0, allowed values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `, default value: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "service_route",
					Description: `(Optional) Service route for the VPN type Feature Template, (see [below for nested schema](#nestedblock--service_route))`,
				},
				resource.Attribute{
					Name:        "nat64",
					Description: `(Optional) NAT64 for the VPN type Feature Template, (see [below for nested schema](#nestedblock--nat64)), can be configured for ` + "`" + `template_type` + "`" + ` "vpn-vedge" <a id="nestedblock--vpn_basic"></a> ## Nested Schema for vpn_basic`,
				},
				resource.Attribute{
					Name:        "vpn_id",
					Description: `(Required) Numeric identifier of the VPN for VPN type of Feature Template, allowed values: 1, 512`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of VPN for VPN type of Feature Template`,
				},
				resource.Attribute{
					Name:        "ecmp_hash_key",
					Description: `(Optional) Boolean flag indicating whether ECMP hash key is enabled or not, allowed values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `, default value: ` + "`" + `false` + "`" + ` <a id="nestedblock--vpn_dns"></a> ## Nested Schema for vpn_dns`,
				},
				resource.Attribute{
					Name:        "primary_dns_ipv4",
					Description: `(Optional) Primary IPv4 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "secondary_dns_ipv4",
					Description: `(Optional) Secondary IPv4 address of DNS server for the VPN, can be set if ` + "`" + `primary_dns_ipv4` + "`" + ` is set`,
				},
				resource.Attribute{
					Name:        "primary_dns_ipv6",
					Description: `(Optional) Primary IPv6 address of DNS server for the VPN`,
				},
				resource.Attribute{
					Name:        "secondary_dns_ipv6",
					Description: `(Optional) Secondary IPv6 address of DNS server for the VPN, can be set if ` + "`" + `primary_dns_ipv6` + "`" + ` is set`,
				},
				resource.Attribute{
					Name:        "vpn_host",
					Description: `(Optional) VPN host list for DNS server (see [below for nested schema](#nestedblock--vpn_host)) <a id="nestedblock--ipv4_route"></a> ## Nested Schema for ipv4_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) IPv4 prefix or address for IPv4 route`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `(Required) Gateway Type for next hop to reach the static route, allowed values: "next-hop", "dhcp", "null0"`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop list (see [below for nested schema](#nestedblock--next_hopv4)) when gateway_type is ` + "`" + `next-hop` + "`" + ``,
				},
				resource.Attribute{
					Name:        "null0_distance",
					Description: `(Optional) Null0 distance when gateway_type is ` + "`" + `null0` + "`" + `, range: [1 - 255] <a id="nestedblock--ipv6_route"></a> ## Nested Schema for ipv6_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) IPv6 prefix or address for IPv6 route`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `(Required) Gateway Type for next hop to reach the static route, allowed values: "next-hop", "null0"`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop list (see [below for nested schema](#nestedblock--next_hopv6)) when gateway_type is ` + "`" + `next-hop` + "`" + ``,
				},
				resource.Attribute{
					Name:        "null0_distance",
					Description: `(Optional) Null0 distance when gateway_type is ` + "`" + `null0` + "`" + `, cannot be set for ` + "`" + `template_type` + "`" + ` "cisco_vpn", range: [1 - 255] <a id="nestedblock--service_route"></a> ## Nested Schema for service_route`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) IPv4 prefix or address for service route`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Optional) Type of the service, allowed values for ` + "`" + `template_type` + "`" + ` "vpn-vedge": "sig", "FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "TE", allowed values for ` + "`" + `template_type` + "`" + ` "cisco_vpn": "FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "TE" <a id="nestedblock--nat64"></a> ## Nested Schema for nat64`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `(Required) NAT64 pool name`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) Starting IPv4 address of NAT64 pool`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) Ending IPv4 address of NAT64 pool`,
				},
				resource.Attribute{
					Name:        "overload",
					Description: `(Optional) Flag indicating NAT64 pool overload, allowed valus: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `, default values: ` + "`" + `false` + "`" + ` <a id="nestedblock--vpn_host"></a> ## Nested Schema for vpn_host`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Host name of the DNS server`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) List of ip addresses associated with the hostname, maximum length: 8 <a id="nestedblock--next_hopv4"></a> ## Nested Schema for next_hop`,
				},
				resource.Attribute{
					Name:        "next_hop_address",
					Description: `(Required) IPv4 address of the nexthop`,
				},
				resource.Attribute{
					Name:        "next_hop_distance",
					Description: `(Optional) Distance of the nexthop, range: [1 - 255] <a id="nestedblock--next_hopv6"></a> ## Nested Schema for next_hop`,
				},
				resource.Attribute{
					Name:        "next_hop_address",
					Description: `(Required) IPv6 address of the nexthop`,
				},
				resource.Attribute{
					Name:        "next_hop_distance",
					Description: `(Optional) Distance of the nexthop, range: [1 - 255]`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sdwan_vpn_interface_feature_template",
			Category:         "Resources",
			ShortDescription: `Resource for SD-WAN VPN interface Feature Template`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Unique VPN Interface Type Feature Template Name, Must not include <, >, !, &, ", or white space; maximum 128 characters`,
				},
				resource.Attribute{
					Name:        "template_description",
					Description: `(Required) Long Description of VPN Interface type Feature Template, Should not be empty`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Required) Type of device which supports Feature Template, Allowed values for Vedge VPN Interface type Feature Template : "vedge-100", "vedge-100-WM", "vedge-ISR1100-4GLTE", "vedge-100-B", "vedge-1000", "vedge-ISR1100-6G", "vedge-100-M", "vedge-2000", "vedge-ISR1100X-4G", "vedge-ISR1100-4G", "vedge-5000","vedge-cloud", "vedge-ISR1100X-6G" Allowed values for Cisco VPN Interface type Feature Template : "vedge-C8500-12X4QC", "vedge-ISRv", "vedge-C1111-4PLTEEA", "vedge-C1161-8P", "vedge-C1117-4PLTEEAW", "vedge-C1121X-8P", "vedge-C8200-1N-4T", "vedge-ISR-4331", "vedge-C1127X-8PMLTEP", "vedge-C1117-4PMLTEEAWE", "vedge-ISR-4451-X", "vedge-C1113-8PLTEEA", "vedge-IR-1821", "vedge-ASR-1001-X", "vedge-ISR-4321", "vedge-C1116-4PLTEEAWE", "vedge-C1109-4PLTE2P", "vedge-C1121-8P", "vedge-ASR-1002-HX", "vedge-C1111-8PLTEEAW", "vedge-C1112-8PWE", "vedge-C1101-4PLTEP", "vedge-ISR1100-4GLTENA-XE", "vedge-C1111-8PLTELA", "vedge-IR-1835", "vedge-C1121X-8PLTEP", "vedge-IR-1833", "vedge-C8300-1N1S-4T2X", "vedge-C1121-4P", "vedge-ISR-4351", "vedge-C1117-4PLTELA", "vedge-C1116-4PWE", "vedge-nfvis-C8200-UCPE", "vedge-C1113-8PM", "vedge-IR-1831", "vedge-C1127-8PLTEP", "vedge-C1121-8PLTEPW", "vedge-C1113-8PW", "vedge-ASR-1001-HX", "vedge-C1128-8PLTEP", "vedge-C1113-8PLTEEAW", "vedge-C1117-4PW", "vedge-C1116-4P", "vedge-C1113-8PMLTEEA", "vedge-C1112-8P", "vedge-ISR-4461", "vedge-C1116-4PLTEEA", "vedge-ISR-4221", "vedge-C1117-4PM", "vedge-C1113-8PLTELAWZ", "vedge-C1117-4PMWE", "vedge-C1109-2PLTEVZ", "vedge-C1113-8P", "vedge-C1117-4P", "vedge-nfvis-ENCS5400", "vedge-C8300-2N2S-6T", "vedge-C1127-8PMLTEP", "vedge-ESR-6300", "vedge-ISR-4221X", vedge-ISR1100-4GLTEGB-XE", "vedge-C8500-12X", "vedge-C1109-2PLTEGB", "vedge-CSR-1000v", "vedge-C1113-8PLTEW", "vedge-C1121X-8PLTEPW", "vedge-ISR1100-6G-XE", "vedge-C1121-4PLTEP", "vedge-C1111-8PLTEEA", "vedge-C1117-4PLTEEA", "vedge-C1127X-8PLTEP", "vedge-C1109-2PLTEUS", vedge-C1112-8PLTEEAWE", "vedge-C1161X-8P", "vedge-C8500L-8S4X", "vedge-C1111-8PW", "vedge-C1161X-8PLTEP", "vedge-C1101-4PLTEPW", "vedge-ISR1100X-4G-XE", "vedge-IR-1101", "vedge-C1111-4P", "vedge-C1111-4PW", "vedge-C1111-8P", "vedge-C1117-4PMLTEEA", "vedge-C1113-8PLTELA", "vedge-C1111-8PLTELAW", "vedge-C1161-8PLTEP", "vedge-ISR1100X-6G-XE", "vedge-ISR-4431", "vedge-C1101-4P", "vedge-C1109-4PLTE2PW", "vedge-C1113-8PMWE", "vedge-C1118-8P", "vedge-C1126-8PLTEP", "vedge-C8300-1N1S-6T", "vedge-C1121-8PLTEP", "vedge-C8300-2N2S-4T2X", "vedge-C1112-8PLTEEA", "vedge-C1111-4PLTELA", "vedge-ASR-1002-X", "vedge-C1111X-8P", "vedge-C1126X-8PLTEP", "vedge-ASR-1006-X", "vedge-C8000V", "vedge-ISR1100-4G-XE", "vedge-C1117-4PLTELAWZ"`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `(Required) Template type for the VPN Interface type Feature Template, Allowed Value: "vpn-vedge-interface","cisco_vpn_interface"`,
				},
				resource.Attribute{
					Name:        "template_min_version",
					Description: `(Required) Minimum Version for the Feature template`,
				},
				resource.Attribute{
					Name:        "factory_default",
					Description: `(Required) Boolean flag indicating whether VPN Interface type Feature Template is factory default or not. If we set it true we can not update or delete resource.`,
				},
				resource.Attribute{
					Name:        "template_definition",
					Description: `(Required) Configuration for VPN Interface type Feature Template, (see [below for nested schema](#nestedblock--template_definition)) ## Attribute Reference ##`,
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
					Description: `(Required) Basic configuration for the VPN Interface type Feature Template, Max Items: 1, Min Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_basic))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_tunnel",
					Description: `(Optional) Tunnel configuration for the VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_tunnel))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_nat",
					Description: `(Optional) NAT(Network Address Translation) configuration for the VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_nat))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_vrrp",
					Description: `(Optional) VRRP(Virtual Router Redundancy Protocol) configuration for the VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_vrrp))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_acl_qos",
					Description: `(Optional) ACL(Apply Access Lists) and QoS configuration for the VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_acl_qos))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_arp",
					Description: `(Optional) ARP(Address Resolution Protocol) configuration for the VPN Interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_arp))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_8021x",
					Description: `(Optional) IEEE 802.1X Authentication configuration for the Vedge VPN Interface type Feature Template, Max Items: 1, If 'NAT' or 'TLOC Extension' is configured, we should not configure vpn_interface_8021x (see [below for nested schema](#nestedblock--vpn_interface_8021x))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_trustsec",
					Description: `(Optional) TrustSec configuration for the Cisco VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_trustsec))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_advanced",
					Description: `(Optional) Advanced configuration for the VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_advanced)) <a id="nestedblock--vpn_interface_basic"></a> ## Nested Schema for vpn_interface_basic`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `(Optional) Shutdown flag for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Interface Description for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) IPv4 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--ipv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) IPv6 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--ipv6))`,
				},
				resource.Attribute{
					Name:        "block_non_source_ip",
					Description: `(Optional) Flag indicating whether forwarded traffic matches with interface's IP prefix range for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "bandwidth_upstream",
					Description: `(Optional) Value of bandwidth above which to generate notifications in case of transmitted traffic for VPN interface type Feature Template, Range: [1 - 2147483647]`,
				},
				resource.Attribute{
					Name:        "bandwidth_downstream",
					Description: `(Optional) Value of bandwidth above which to generate notifications in case of received traffic for VPN interface type Feature Template, Range: [1 - 2147483647] <a id="nestedblock--vpn_interface_tunnel"></a> ## Nested Schema for vpn_interface_tunnel`,
				},
				resource.Attribute{
					Name:        "per_tunnel_qos",
					Description: `(Optional) Flag indicating whether per_tunnel_qos is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "per_tunnel_qos_aggregator",
					Description: `(Optional) Flag indicating whether per_tunnel_qos_aggregator is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "tunnels_bandwidth_percent",
					Description: `(Optional) Value of tunnel bandwith percentage for VPN interface type Feature Template, Range: [1 - 99]`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) TLOC color value for VPN interface type Feature Template, Allowed values: "3g", "biz-internet", "blue", "bronze", "custom1", "custom2", "custom3", "default", "gold", "green", "Ite", "metro-ethernet", "mpls", "private1", "private2", "private3", "private4", "private5", "private6", "public-internet", "red", "silver"`,
				},
				resource.Attribute{
					Name:        "restrict",
					Description: `(Optional) Restrict flag for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of Groups in tunnle for VPN interface type Feature Template, Range: [1 - 4294967295]`,
				},
				resource.Attribute{
					Name:        "border",
					Description: `(Optional) Flag indicating whether Border is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "control_connection",
					Description: `(Optional) Flag indicating TLOC connection is enabled or not for VPN interface type Feature Template, Allowed value: true, false`,
				},
				resource.Attribute{
					Name:        "maximum_control_connections",
					Description: `(Optional) Maximum number of vSamart controllers that WAN tunnel interface can connect for VPN interface type Feature Template, Range: [0 - 100]`,
				},
				resource.Attribute{
					Name:        "vbond_as_stun_server",
					Description: `(Optional) Flag indicating whether Session Traversal Utilities for NAT(STUN) is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "exclude_controller_group_list",
					Description: `(Optional) List of vSmart controllers that the tunnel interface is not allowed to connect for VPN interface type Feature Template, Range: [0 - 100]`,
				},
				resource.Attribute{
					Name:        "vmanage_connection_preference",
					Description: `(Optional) Flag indicating preference for using a tunnel interface to exchange control traffic with the vManage NMS for VPN interface type Feature Template, Range: [0 - 8], Default: 5`,
				},
				resource.Attribute{
					Name:        "port_hop",
					Description: `(Optional) Flag indicating whether Port Hop is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "low_bandwidth_link",
					Description: `(Optional) Flag indicating whether tunnle is set as a low-bandwidth link for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "allow_service",
					Description: `(Optional) Nested block for allow or disallow various services for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--allow_service))`,
				},
				resource.Attribute{
					Name:        "encapsulation",
					Description: `(Optional) Encapsulation configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--encapsulation)) <a id="nestedblock--vpn_interface_nat"></a> ## Nested Schema for vpn_interface_nat`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) NAT IPv4 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--natv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) NAT IPv6 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--natv6)) <a id="nestedblock--vpn_interface_vrrp"></a> ## Nested Schema for vpn_interface_vrrp`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) VRRP IPv4 configuration for VPN interface type Feature Template, Max Items for Cisco VPN Interface type Feature Template: 1, Max Items for Vedge VPN Interface type Feature Template: 5,(see [below for nested schema](#nestedblock--vrrpv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) VRRP IPv6 configuration for VPN interface type Feature Template, Max Items for VPN Interface type Feature Template: 1,(see [below for nested schema](#nestedblock--vrrpv6)) <a id="nestedblock--vpn_interface_acl_qos"></a> ## Nested Schema for vpn_interface_acl_qos`,
				},
				resource.Attribute{
					Name:        "adapt_period",
					Description: `(Optional) Value of Adapt Period for Cisco VPN interface type Feature Template, Range: [1 - 720]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_min",
					Description: `(Optional) Minimum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_max",
					Description: `(Optional) Maximum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_default",
					Description: `(Optional) Default value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_min",
					Description: `(Optional) Minimum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_max",
					Description: `(Optional) Maximum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_default",
					Description: `(Optional) Default value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate",
					Description: `(Optional) Value of aggregate traffic transmission rate on the VPN interface for VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "qos_map",
					Description: `(Optional) Name of the QoS map to apply to packets being transmitted out the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "rewrite_rule",
					Description: `(Optional) Name of the rewrite rule to apply on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4_ingress_access_list",
					Description: `(Optional) Name of the access list to apply to IPv4 packets being received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4_egress_access_list",
					Description: `(Optional) Name of the access list to apply to IPv4 packets being transmitted on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv6_ingress_access_list",
					Description: `(Optional) Name of the access list to apply to IPv6 packets being received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv6_egress_access_list",
					Description: `(Optional) Name of the access list to apply to IPv6 packets being transmitted on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ingress_policer_name",
					Description: `(Optional) Name of the policer to apply to packets received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "egress_policer_name",
					Description: `(Optional) Name of the policer to apply to packets being transmitted on the interface for Vedge VPN interface type Feature Template <a id="nestedblock--vpn_interface_arp"></a> ## Nested Schema for vpn_interface_arp`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) IP address for the ARP entry in dotted decimal notation or as a fully qualified host name for VPN interface type Feature Template, Allowed values: IPv4 address`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Required) MAC address in colon-separated hexadecimal notation for VPN interface type Feature Template, Allowed values: MAC address <a id="nestedblock--vpn_interface_8021x"></a> ## Nested Schema for vpn_interface_8021x`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `(Optional) List of tags of the RADIUS server to use for 802.1X authentication for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "account_interim_interval",
					Description: `(Optional) Value of how often to send 802.1X interim accounting updates to the RADIUS server for Vedge VPN interface type Feature Template, Range: [1 - 1440]`,
				},
				resource.Attribute{
					Name:        "nas_identifier",
					Description: `(Optional) The NAS identifier of the local router for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nas_ip",
					Description: `(Optional) The NAS IP address of the local router for Vedge VPN interface type Feature Template, Allowed values: IPv4 address`,
				},
				resource.Attribute{
					Name:        "wake_on_lan",
					Description: `(Optional) The flag that enable client to be powered up when the router receives an Ethernet magic packet frame for Vedge VPN interface type Feature Template, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "control_direction",
					Description: `(Optional) Direction type of packets for Vedge VPN interface type Feature Template, Allowed values: "in-and-out", "in-only", Default: "in-and-out"`,
				},
				resource.Attribute{
					Name:        "reauthentication",
					Description: `(Optional) Value indicating how often to reauthenticate 802.1X clients for Vedge VPN interface type Feature Template, Range: [0 - 1440], Default: 0`,
				},
				resource.Attribute{
					Name:        "inactivity",
					Description: `(Optional) Value indicating how long to wait before revoking an 802.1X client's network access for Vedge VPN interface type Feature Template, Range: [0 - 1440], Default: 60`,
				},
				resource.Attribute{
					Name:        "host_mode",
					Description: `(Optional) Value indicating whether an 802.1X interface grants access to a single client or to multiple clients for Vedge VPN interface type Feature Template, Allowed values: "multi-auth", "multi-host", "single-host", Default: "single-host"`,
				},
				resource.Attribute{
					Name:        "advanced_options",
					Description: `(Optional) Advance configuration of vpn_interface_8021x for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--advanced_options)) <a id="nestedblock--vpn_interface_trustsec"></a> ## Nested Schema for vpn_interface_trustsec`,
				},
				resource.Attribute{
					Name:        "enable_sgt_propagation",
					Description: `(Optional) Flag indicating whether SGT propagation is enabled or not for Cisco VPN interface type Feature Template, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "propagate",
					Description: `(Optional) Flag indicating whether propagate is enabled or not for Cisco VPN interface type Feature Template, Allowed value: true, false`,
				},
				resource.Attribute{
					Name:        "security_group_tag",
					Description: `(Optional) Value of security group tag for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "trusted",
					Description: `(Optional) Flag indicating whether trusted is enable or not for Cisco VPN interface type Feature Template, Allowed value: true, false Default: false <a id="nestedblock--vpn_interface_advanced"></a> ## Nested Schema for vpn_interface_advanced`,
				},
				resource.Attribute{
					Name:        "duplex",
					Description: `(Optional) Value of duplex to run interface in auto, full, or half mode for VPN interface type Feature Template, Allowed values for Cisco VPN Interface type feature template: "auto", "full", "half", Allowed values for Vedge VPN Interface type feature template: "full", "half"`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) Value of MAC address to associate with the interface, in colon-separated hexadecimal notation, Allowed values: MAC address`,
				},
				resource.Attribute{
					Name:        "ip_mtu",
					Description: `(Optional) Value of ip mtu to specify the maximum MTU size of packets on the interface, Range: [576 - 2000], Default: 1500`,
				},
				resource.Attribute{
					Name:        "pmtu_discovery",
					Description: `(Optional) Flag indicating whether path MTU is enable or not for Vedge VPN interface type Feature Template, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "flow_control",
					Description: `(Optional) Value of bidirectional flow control settings for Vedge VPN interface type Feature Template, Allowed values: "autoneg", "both", "egress", "ingress", "none", Default: "autoneg"`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `(Optional) Maximum segment size (MSS) of TPC SYN packets passing through the router, Range for Cisco VPN Interface : [500, 1460], Range for Vedge VPN Interface : [552 - 1460]`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) Value of speed to use when the remote end of the connection does not support autonegotiation, Allowed values for Cisco VPN Interface type feature template: "10", "100", "1000", "10000", "2500", Allowed values for Vedge VPN Interface type feature template: "10", "100", "1000", "10000"`,
				},
				resource.Attribute{
					Name:        "clear_dont_fragment",
					Description: `(Optional) Flag indicating Don't Fragment (DF) bit in the IPv4 packet header for packets being transmitted out the interface is clar or not, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "static_ingess_qos",
					Description: `(Optional) Queue number to use for incoming traffic for Vedge VPN interface type Feature Template, Range: [0 - 7]`,
				},
				resource.Attribute{
					Name:        "arp_timeout",
					Description: `(Optional) Value indicating how long it takes for a dynamically learned ARP entry to time out, Range for Cisco VPN Interface : [0 - 2147483], Range for Vedge VPN Interface : [0 - 2678400] Default: 1500`,
				},
				resource.Attribute{
					Name:        "autonegotiation",
					Description: `(Optional) Flag indicating autonegotiation mode, Allowed value: true, false Default: true`,
				},
				resource.Attribute{
					Name:        "tloc_extension",
					Description: `(Optional) Name of a physical interface on the same router that connects to the WAN transport`,
				},
				resource.Attribute{
					Name:        "power_over_ethernet",
					Description: `(Optional) Flag indicating whether PoE(Power over Ethernet) is enabled or not for Vedge VPN interface type Feature Template, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "load_interval",
					Description: `(Optional) Value of load interval for Cisco VPN interface type Feature Template, Range: [30 - 600], Allowed values: value can only be multiple of 30`,
				},
				resource.Attribute{
					Name:        "tracker",
					Description: `(Optional) List of a tracker name to track the status of transport interfaces that connect to the internet for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "icmp_redirect_disable",
					Description: `(Optional) Flag indicating whether ICMP redirect disable or not on interface for VPN interface type Feature Template, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "gre_tunnel_source_ip",
					Description: `(Optional) IP address of the extended WAN interface, Allowed values: IPv4 address`,
				},
				resource.Attribute{
					Name:        "xconnect",
					Description: `(Optional) The name of a physical interface on the same router that connects to the WAN transport for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ip_directed_broadcast",
					Description: `(Optional) Flag indicate whether directed IP braodcast is enabled or not for VPN interface type Feature Template, Allowed value: true, false Default: false <a id="nestedblock--ipv4"></a> ## Nested Schema for Basic ipv4`,
				},
				resource.Attribute{
					Name:        "primary_address",
					Description: `(Optional) Value of IPv4 prefix for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "secondary_address",
					Description: `(Optional) Value of IPv4 prefix for VPN interface type Feature Template, Max Items: 4(see [below for nested schema](#nestedblock--secondary_addressv4))`,
				},
				resource.Attribute{
					Name:        "dhcp_distance",
					Description: `(Optional) Value of IPv4 DHCP distance for VPN interface type Feature Template, Range: [1 - 65536]`,
				},
				resource.Attribute{
					Name:        "dhcp_helper",
					Description: `(Optional) List of IPv4 or IPv6 DHCP Helper for VPN interface type Feature Template, Allowed values: IPv4 or IPv6 address <a id="nestedblock--ipv6"></a> ## Nested Schema for Basic ipv6`,
				},
				resource.Attribute{
					Name:        "primary_address",
					Description: `(Optional) Value of IPv6 prefix for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "secondary_address",
					Description: `(Optional) Value of IPv6 prefix for VPN interface type Feature Template, Max Items: 2(see [below for nested schema](#nestedblock--secondary_addressv6))`,
				},
				resource.Attribute{
					Name:        "dhcp_distance",
					Description: `(Optional) Value of IPv6 DHCP distance for Vedge VPN interface type Feature Template, Range: [1 - 65536]`,
				},
				resource.Attribute{
					Name:        "dhcp_rapid_commit",
					Description: `(Optional) Value of IPv6 DHCP rapid commit enabled or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp_helper",
					Description: `(Optional) DHCP Helper value to designate the interface as a DHCP helper on a router for VPN interface type Feature Template, Max Items: 8(see [below for nested schema](#nestedblock--dhcp_helperV6)) <a id="nestedblock--allow_service"></a> ## Nested Schema for Tunnel allow_service`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Flag indicating whether All serivce is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "bgp",
					Description: `(Optional) Flag indicating whether BGP is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `(Optional) Flag indicating whether DHCP is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional) Flag indicating whether DNS is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "icmp",
					Description: `(Optional) Flag indicating whether ICMP is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "netconf",
					Description: `(Optional) Flag indicating whether NETCONF is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "ntp",
					Description: `(Optional) Flag indicating whether NTP is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `(Optional) Flag indicating whether OSPF is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "ssh",
					Description: `(Optional) Flag indicating whether SSH is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "https",
					Description: `(Optional) Flag indicating whether HTTPS is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "snmp",
					Description: `(Optional) Flag indicating whether SNMP is enabled or not for Cisco VPN interface type Feature Template, Default: false <a id="nestedblock--encapsulation"></a> ## Nested Schema for Tunnel encapsulation`,
				},
				resource.Attribute{
					Name:        "gre",
					Description: `(Optional) Flag indicating whether gre is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "gre_preference",
					Description: `(Optional) Value of GRE preference for VPN interface type Feature Template, Range: [0 - 4294967295]`,
				},
				resource.Attribute{
					Name:        "gre_weight",
					Description: `(Optional) Value of GRE weight for VPN interface type Feature Template, Range: [1 - 255], Default: 1`,
				},
				resource.Attribute{
					Name:        "ipsec",
					Description: `(Optional) Flag indicating whether ipsec is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "ipsec_preference",
					Description: `(Optional) Value of IPsec preference for VPN interface type Feature Template, Range: [0 - 4294967295]`,
				},
				resource.Attribute{
					Name:        "ipsec_weight",
					Description: `(Optional) Value of IPsec weight for VPN interface type Feature Template, Range: [1 - 255], Default: 1`,
				},
				resource.Attribute{
					Name:        "carrier",
					Description: `(Optional) Carrier name or private network identifier to associate with the tunnel for VPN interface type Feature Template, Allowed values: "carrier1", "carrier2", "carrier3", "carrier4", "carrier5", "carrier6", "carrier7", "carrier8", "default", Default: "default"`,
				},
				resource.Attribute{
					Name:        "bind_loopback_tunnel",
					Description: `(Optional) Name of a physical interface to bind to a loopback interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_resort_circuit",
					Description: `(Optional) Flag indicating whether tunnel interface is used as the circuit of last resort for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `(Optional) Value of hold time for Vedge VPN interface type Feature Template, Range: [100 - 10000], Default: 7000`,
				},
				resource.Attribute{
					Name:        "nat_refresh_interval",
					Description: `(Optional) Value of interval between NAT refresh packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template, Range: [1 - 60], Default: 5`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `(Optional) Value of interval between Hello packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template, Range: [10 - 600000], Default: 1000`,
				},
				resource.Attribute{
					Name:        "hello_tolerance",
					Description: `(Optional) Value of the time to wait for a Hello packet on a DTLS or TLS WAN transport connection before declaring that transport tunnel to be down for VPN interface type Feature Template, Range: [12 - 6000], Default: 12`,
				},
				resource.Attribute{
					Name:        "gre_tunnel_destination_ip",
					Description: `(Optional) Value of GRE tunnel destination IP for VPN interface type Feature Template, Allowed value: IPv4 Address <a id="nestedblock--natv4"></a> ## Nested Schema for NAT ipv4`,
				},
				resource.Attribute{
					Name:        "nat_type",
					Description: `(Optional) NAT type for Cisco VPN interface type Feature Template, Allowed value: "interface", "pool", "loopback" Default: "interface"`,
				},
				resource.Attribute{
					Name:        "refresh_mode",
					Description: `(Optional) Refresh mode for VPN interface type Feature Template, Allowed value: "bi-directional", "outbound", Default: "outbound"`,
				},
				resource.Attribute{
					Name:        "log_nat_flow_creations_or_deletions",
					Description: `(Optional) Flag indicating whether log nat flow creations or deletions are enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `(Optional) The time when NAT translations over UDP sessions time out for VPN interface type Feature Template, Range: [1 - 65536] minutes, Default: 1`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `(Optional) The time when NAT translations over TCP sessions time out for VPN interface type Feature Template, Range: [1 - 65536] minutes, Default: 60`,
				},
				resource.Attribute{
					Name:        "block_icmp",
					Description: `(Optional) Flag indicating whether to block icmp or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "respond_to_ping",
					Description: `(Optional) Flag indicating whether to respond to ping or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "nat_pool_range_start",
					Description: `(Optional) Value of starting IP address for the NAT pool for VPN interface type Feature Template, Allowed value: IPv4 Address`,
				},
				resource.Attribute{
					Name:        "nat_pool_range_end",
					Description: `(Optional) Value of closing IP address for the NAT pool for VPN interface type Feature Template, Allowed value: IPv4 Address`,
				},
				resource.Attribute{
					Name:        "nat_pool_prefix_length",
					Description: `(Optional) Value of NAT pool prefix length for Cisco VPN interface type Feature Template, Range: [1 - 65536], Default: 60`,
				},
				resource.Attribute{
					Name:        "overload",
					Description: `(Optional) Flag indicating whether to enable per-port translation or not for Cisco VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "nat_inside_source_loopback_interface",
					Description: `(Optional) Value of NAT inside source loopback interface for Cisco VPN interface type Feature Template, String length Range: [1 - 32]`,
				},
				resource.Attribute{
					Name:        "port_forward",
					Description: `(Optional) Configure port forwarding rule for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--port_forward))`,
				},
				resource.Attribute{
					Name:        "static_nat",
					Description: `(Optional) Configure static nat for VPN interface type Feature Template(see [below for nested schema](#nestedblock--static_nat)) <a id="nestedblock--natv6"></a> ## Nested Schema for NAT ipv6`,
				},
				resource.Attribute{
					Name:        "nat64",
					Description: `(Optional) Flag indicating whether to enable nat64 or not for VPN interface type Feature Template, Default: false <a id="nestedblock--vrrpv4"></a> ## Nested Schema for VRRP ipv4`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template, Range: [1 - 255]`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template, Range: [1 - 254], Default: 100`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `(Optional) Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template, Range for Vedge VPN interface: [1 - 3600] Seconds, Default: 1 Second, Range for Cisco VPN interface: [100 - 40950] Milliseconds, Default: 100 Milliseconds`,
				},
				resource.Attribute{
					Name:        "track_omp",
					Description: `(Optional) Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "track_prefix_list",
					Description: `(Optional) Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Value of IPv4 address of the virtual router for VPN interface type Feature Template, Allowed value: IPv4 Address <a id="nestedblock--vrrpv6"></a> ## Nested Schema for VRRP ipv6`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template, Range: [1 - 255]`,
				},
				resource.Attribute{
					Name:        "link_local_ipv6_address",
					Description: `(Required) Value of IPv6 address of the virtual router for VPN interface type Feature Template, Allowed value: IPv6 Address`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template, Range: [1 - 254], Default: 100`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `(Optional) Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template, Range for Vedge VPN interface: [1 - 3600] Seconds, Range for Cisco VPN interface: [1 - 40950] Milliseconds, Default: 1 Second`,
				},
				resource.Attribute{
					Name:        "track_omp",
					Description: `(Optional) Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "track_prefix_list",
					Description: `(Optional) Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "global_ipv6_prefix",
					Description: `(Optional) Value of global IPv6 prefix of the virtual router for VPN interface type Feature Template <a id="nestedblock--advanced_options"></a> ## Nested Schema for 8021.X advanced_options`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) Configure VLAN for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--vlan))`,
				},
				resource.Attribute{
					Name:        "dynamic_authentication_server",
					Description: `(Optional) Configure Dynamic Authentication Server for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--dynamic_authentication_server))`,
				},
				resource.Attribute{
					Name:        "mac_authentication_bypass",
					Description: `(Optional) Configure MAC Authentication Bypass for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--mac_authentication_bypass))`,
				},
				resource.Attribute{
					Name:        "request_attributes",
					Description: `(Optional) Configure Request Attributes for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--request_attributes)) <a id="nestedblock--secondary_addressv4"></a> ## Nested Schema for Basic IPv4 secondary_address`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Value of IPv4 prefix for VPN interface type Feature Template <a id="nestedblock--secondary_addressv6"></a> ## Nested Schema for Basic IPv6 secondary_address`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Value of IPv6 prefix for VPN interface type Feature Template <a id="nestedblock--dhcp_helperV6"></a> ## Nested Schema for Basic IPv6 dhcp_helper`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Value of IPv6 address for VPN interface type Feature Template, Allowed value: IPv6 Address`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Value of VPN id for VPN interface type Feature Template, Range: [1 - 65536] <a id="nestedblock--port_forward"></a> ## Nested Schema for NAT ipv4 port_forward`,
				},
				resource.Attribute{
					Name:        "port_start_range",
					Description: `(Required) Value of port starting range for Vedge VPN interface type Feature Template, Range: [0 - 65535]`,
				},
				resource.Attribute{
					Name:        "port_end_range",
					Description: `(Required) Value of port ending range for Vedge VPN interface type Feature Template, Range: [0 - 65535]`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol for Vedge VPN interface type Feature Template, Allowed value: "tcp", "udp"`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Required) VPN id for Vedge VPN interface type Feature Template, Range: [0 - 65535]`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Required) Private IP address for Vedge VPN interface type Feature Template, Allowed value: IPv4 Address <a id="nestedblock--static_nat"></a> ## Nested Schema for NAT ipv4 static_nat`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `(Required) Value of the inside local address as source IP address for VPN interface type Feature Template, Allowed value: IPv4 Address`,
				},
				resource.Attribute{
					Name:        "translated_source_ip",
					Description: `(Required) Value of the inside global address as the translated source IP address for VPN interface type Feature Template, Allowed value: IPv4 Address`,
				},
				resource.Attribute{
					Name:        "source_vpn_id",
					Description: `(Optional) Configures Source VPN ID, which is the Service VPN in which the host resides for VPN interface type Feature Template, Range: [0 - 65530], Default: 0`,
				},
				resource.Attribute{
					Name:        "static_nat_direction",
					Description: `(Optional) Set the direction in which to perform network address translation for VPN interface type Feature Template, Allowed value for Vedge VPN interface type Feature Template : "inside", "outside" Default: "inside" Allowed value for Cisco VPN interface type Feature Template : "inside"`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol for Vedge VPN interface type Feature Template, Allowed value: "tcp", "udp"`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) Value of source port for Vedge VPN interface type Feature Template, Range: [0 - 65535]`,
				},
				resource.Attribute{
					Name:        "translate_port",
					Description: `(Optional) Value of translate port for Vedge VPN interface type Feature Template, Range: [0 - 65535] <a id="nestedblock--vlan"></a> ## Nested Schema for vlan`,
				},
				resource.Attribute{
					Name:        "authentication_fail_vlan",
					Description: `(Optional) Value of authentication fail VLAN for Vedge VPN interface type Feature Template, Range: [1 - 65536]`,
				},
				resource.Attribute{
					Name:        "guest_vlan",
					Description: `(Optional) Value of guest VLAN for Vedge VPN interface type Feature Template, Range: [1 - 65536]`,
				},
				resource.Attribute{
					Name:        "authentication_reject_vlan",
					Description: `(Optional) Value of authentication reject VLAN for Vedge VPN interface type Feature Template, Range: [1 - 65536]`,
				},
				resource.Attribute{
					Name:        "default_vlan",
					Description: `(Optional) Value of default VLAN for Vedge VPN interface type Feature Template, Range: [1 - 65536] <a id="nestedblock--dynamic_authentication_server"></a> ## Nested Schema for dynamic_authentication_server`,
				},
				resource.Attribute{
					Name:        "das_port",
					Description: `(Optional) Value of Dynamic Authentication Server port for Vedge VPN interface type Feature Template, Range: [1 - 65535], Default: 3799,`,
				},
				resource.Attribute{
					Name:        "client",
					Description: `(Optional) Value of Client IP address for Vedge VPN interface type Feature Template, Allowed value: IPv4 Address`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) Secret key Value for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "time_window",
					Description: `(Optional) Time window value for Vedge VPN interface type Feature Template, Range: [10 - 1000] Seconds, Default: 300,`,
				},
				resource.Attribute{
					Name:        "required_timestamp",
					Description: `(Optional) Flag indicating whether to enable required timestamp or not for Vedge VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Value of VPN ID for Vedge VPN interface type Feature Template, Range: [0 - 65535], Default: 0 <a id="nestedblock--mac_authentication_bypass"></a> ## Nested Schema for mac_authentication_bypass`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) Flag indicating whether to enable server or not for Vedge VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "mac_authentication_bypass_entries",
					Description: `(Optional) List of MAC Authenticaion bypass entries for Vedge VPN interface type Feature Template, Max Items: 8 <a id="nestedblock--request_attributes"></a> ## Nested Schema for request_attributes`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) Configure Authentication for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--authentication))`,
				},
				resource.Attribute{
					Name:        "accounting",
					Description: `(Optional) Configure Accounting for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--accounting)) <a id="nestedblock--authentication"></a> ## Nested Schema for request_attributes authentication`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID for Vedge VPN interface type Feature Template, Range: [1 - 64]`,
				},
				resource.Attribute{
					Name:        "syntax_choice",
					Description: `(Required) Syntax Choice for Vedge VPN interface type Feature Template, Allowed value: "string", "integer", "octet"`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value for Vedge VPN interface type Feature Template <a id="nestedblock--accounting"></a> ## Nested Schema for request_attributes accounting`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID for Vedge VPN interface type Feature Template, Range: [1 - 64]`,
				},
				resource.Attribute{
					Name:        "syntax_choice",
					Description: `(Required) Syntax Choice for Vedge VPN interface type Feature Template, Allowed value: "string", "integer", "octet"`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value for Vedge VPN interface type Feature Template`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(Required) Basic configuration for the VPN Interface type Feature Template, Max Items: 1, Min Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_basic))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_tunnel",
					Description: `(Optional) Tunnel configuration for the VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_tunnel))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_nat",
					Description: `(Optional) NAT(Network Address Translation) configuration for the VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_nat))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_vrrp",
					Description: `(Optional) VRRP(Virtual Router Redundancy Protocol) configuration for the VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_vrrp))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_acl_qos",
					Description: `(Optional) ACL(Apply Access Lists) and QoS configuration for the VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_acl_qos))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_arp",
					Description: `(Optional) ARP(Address Resolution Protocol) configuration for the VPN Interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_arp))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_8021x",
					Description: `(Optional) IEEE 802.1X Authentication configuration for the Vedge VPN Interface type Feature Template, Max Items: 1, If 'NAT' or 'TLOC Extension' is configured, we should not configure vpn_interface_8021x (see [below for nested schema](#nestedblock--vpn_interface_8021x))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_trustsec",
					Description: `(Optional) TrustSec configuration for the Cisco VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_trustsec))`,
				},
				resource.Attribute{
					Name:        "vpn_interface_advanced",
					Description: `(Optional) Advanced configuration for the VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_advanced)) <a id="nestedblock--vpn_interface_basic"></a> ## Nested Schema for vpn_interface_basic`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `(Optional) Shutdown flag for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Interface Description for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) IPv4 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--ipv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) IPv6 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--ipv6))`,
				},
				resource.Attribute{
					Name:        "block_non_source_ip",
					Description: `(Optional) Flag indicating whether forwarded traffic matches with interface's IP prefix range for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "bandwidth_upstream",
					Description: `(Optional) Value of bandwidth above which to generate notifications in case of transmitted traffic for VPN interface type Feature Template, Range: [1 - 2147483647]`,
				},
				resource.Attribute{
					Name:        "bandwidth_downstream",
					Description: `(Optional) Value of bandwidth above which to generate notifications in case of received traffic for VPN interface type Feature Template, Range: [1 - 2147483647] <a id="nestedblock--vpn_interface_tunnel"></a> ## Nested Schema for vpn_interface_tunnel`,
				},
				resource.Attribute{
					Name:        "per_tunnel_qos",
					Description: `(Optional) Flag indicating whether per_tunnel_qos is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "per_tunnel_qos_aggregator",
					Description: `(Optional) Flag indicating whether per_tunnel_qos_aggregator is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "tunnels_bandwidth_percent",
					Description: `(Optional) Value of tunnel bandwith percentage for VPN interface type Feature Template, Range: [1 - 99]`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) TLOC color value for VPN interface type Feature Template, Allowed values: "3g", "biz-internet", "blue", "bronze", "custom1", "custom2", "custom3", "default", "gold", "green", "Ite", "metro-ethernet", "mpls", "private1", "private2", "private3", "private4", "private5", "private6", "public-internet", "red", "silver"`,
				},
				resource.Attribute{
					Name:        "restrict",
					Description: `(Optional) Restrict flag for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of Groups in tunnle for VPN interface type Feature Template, Range: [1 - 4294967295]`,
				},
				resource.Attribute{
					Name:        "border",
					Description: `(Optional) Flag indicating whether Border is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "control_connection",
					Description: `(Optional) Flag indicating TLOC connection is enabled or not for VPN interface type Feature Template, Allowed value: true, false`,
				},
				resource.Attribute{
					Name:        "maximum_control_connections",
					Description: `(Optional) Maximum number of vSamart controllers that WAN tunnel interface can connect for VPN interface type Feature Template, Range: [0 - 100]`,
				},
				resource.Attribute{
					Name:        "vbond_as_stun_server",
					Description: `(Optional) Flag indicating whether Session Traversal Utilities for NAT(STUN) is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "exclude_controller_group_list",
					Description: `(Optional) List of vSmart controllers that the tunnel interface is not allowed to connect for VPN interface type Feature Template, Range: [0 - 100]`,
				},
				resource.Attribute{
					Name:        "vmanage_connection_preference",
					Description: `(Optional) Flag indicating preference for using a tunnel interface to exchange control traffic with the vManage NMS for VPN interface type Feature Template, Range: [0 - 8], Default: 5`,
				},
				resource.Attribute{
					Name:        "port_hop",
					Description: `(Optional) Flag indicating whether Port Hop is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "low_bandwidth_link",
					Description: `(Optional) Flag indicating whether tunnle is set as a low-bandwidth link for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "allow_service",
					Description: `(Optional) Nested block for allow or disallow various services for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--allow_service))`,
				},
				resource.Attribute{
					Name:        "encapsulation",
					Description: `(Optional) Encapsulation configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--encapsulation)) <a id="nestedblock--vpn_interface_nat"></a> ## Nested Schema for vpn_interface_nat`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) NAT IPv4 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--natv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) NAT IPv6 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--natv6)) <a id="nestedblock--vpn_interface_vrrp"></a> ## Nested Schema for vpn_interface_vrrp`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) VRRP IPv4 configuration for VPN interface type Feature Template, Max Items for Cisco VPN Interface type Feature Template: 1, Max Items for Vedge VPN Interface type Feature Template: 5,(see [below for nested schema](#nestedblock--vrrpv4))`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) VRRP IPv6 configuration for VPN interface type Feature Template, Max Items for VPN Interface type Feature Template: 1,(see [below for nested schema](#nestedblock--vrrpv6)) <a id="nestedblock--vpn_interface_acl_qos"></a> ## Nested Schema for vpn_interface_acl_qos`,
				},
				resource.Attribute{
					Name:        "adapt_period",
					Description: `(Optional) Value of Adapt Period for Cisco VPN interface type Feature Template, Range: [1 - 720]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_min",
					Description: `(Optional) Minimum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_max",
					Description: `(Optional) Maximum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_upstream_default",
					Description: `(Optional) Default value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_min",
					Description: `(Optional) Minimum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_max",
					Description: `(Optional) Maximum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate_downstream_default",
					Description: `(Optional) Default value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "shaping_rate",
					Description: `(Optional) Value of aggregate traffic transmission rate on the VPN interface for VPN interface type Feature Template, Range: [8 - 100000000]`,
				},
				resource.Attribute{
					Name:        "qos_map",
					Description: `(Optional) Name of the QoS map to apply to packets being transmitted out the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "rewrite_rule",
					Description: `(Optional) Name of the rewrite rule to apply on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4_ingress_access_list",
					Description: `(Optional) Name of the access list to apply to IPv4 packets being received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv4_egress_access_list",
					Description: `(Optional) Name of the access list to apply to IPv4 packets being transmitted on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv6_ingress_access_list",
					Description: `(Optional) Name of the access list to apply to IPv6 packets being received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ipv6_egress_access_list",
					Description: `(Optional) Name of the access list to apply to IPv6 packets being transmitted on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ingress_policer_name",
					Description: `(Optional) Name of the policer to apply to packets received on the interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "egress_policer_name",
					Description: `(Optional) Name of the policer to apply to packets being transmitted on the interface for Vedge VPN interface type Feature Template <a id="nestedblock--vpn_interface_arp"></a> ## Nested Schema for vpn_interface_arp`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) IP address for the ARP entry in dotted decimal notation or as a fully qualified host name for VPN interface type Feature Template, Allowed values: IPv4 address`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Required) MAC address in colon-separated hexadecimal notation for VPN interface type Feature Template, Allowed values: MAC address <a id="nestedblock--vpn_interface_8021x"></a> ## Nested Schema for vpn_interface_8021x`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `(Optional) List of tags of the RADIUS server to use for 802.1X authentication for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "account_interim_interval",
					Description: `(Optional) Value of how often to send 802.1X interim accounting updates to the RADIUS server for Vedge VPN interface type Feature Template, Range: [1 - 1440]`,
				},
				resource.Attribute{
					Name:        "nas_identifier",
					Description: `(Optional) The NAS identifier of the local router for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "nas_ip",
					Description: `(Optional) The NAS IP address of the local router for Vedge VPN interface type Feature Template, Allowed values: IPv4 address`,
				},
				resource.Attribute{
					Name:        "wake_on_lan",
					Description: `(Optional) The flag that enable client to be powered up when the router receives an Ethernet magic packet frame for Vedge VPN interface type Feature Template, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "control_direction",
					Description: `(Optional) Direction type of packets for Vedge VPN interface type Feature Template, Allowed values: "in-and-out", "in-only", Default: "in-and-out"`,
				},
				resource.Attribute{
					Name:        "reauthentication",
					Description: `(Optional) Value indicating how often to reauthenticate 802.1X clients for Vedge VPN interface type Feature Template, Range: [0 - 1440], Default: 0`,
				},
				resource.Attribute{
					Name:        "inactivity",
					Description: `(Optional) Value indicating how long to wait before revoking an 802.1X client's network access for Vedge VPN interface type Feature Template, Range: [0 - 1440], Default: 60`,
				},
				resource.Attribute{
					Name:        "host_mode",
					Description: `(Optional) Value indicating whether an 802.1X interface grants access to a single client or to multiple clients for Vedge VPN interface type Feature Template, Allowed values: "multi-auth", "multi-host", "single-host", Default: "single-host"`,
				},
				resource.Attribute{
					Name:        "advanced_options",
					Description: `(Optional) Advance configuration of vpn_interface_8021x for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--advanced_options)) <a id="nestedblock--vpn_interface_trustsec"></a> ## Nested Schema for vpn_interface_trustsec`,
				},
				resource.Attribute{
					Name:        "enable_sgt_propagation",
					Description: `(Optional) Flag indicating whether SGT propagation is enabled or not for Cisco VPN interface type Feature Template, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "propagate",
					Description: `(Optional) Flag indicating whether propagate is enabled or not for Cisco VPN interface type Feature Template, Allowed value: true, false`,
				},
				resource.Attribute{
					Name:        "security_group_tag",
					Description: `(Optional) Value of security group tag for Cisco VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "trusted",
					Description: `(Optional) Flag indicating whether trusted is enable or not for Cisco VPN interface type Feature Template, Allowed value: true, false Default: false <a id="nestedblock--vpn_interface_advanced"></a> ## Nested Schema for vpn_interface_advanced`,
				},
				resource.Attribute{
					Name:        "duplex",
					Description: `(Optional) Value of duplex to run interface in auto, full, or half mode for VPN interface type Feature Template, Allowed values for Cisco VPN Interface type feature template: "auto", "full", "half", Allowed values for Vedge VPN Interface type feature template: "full", "half"`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) Value of MAC address to associate with the interface, in colon-separated hexadecimal notation, Allowed values: MAC address`,
				},
				resource.Attribute{
					Name:        "ip_mtu",
					Description: `(Optional) Value of ip mtu to specify the maximum MTU size of packets on the interface, Range: [576 - 2000], Default: 1500`,
				},
				resource.Attribute{
					Name:        "pmtu_discovery",
					Description: `(Optional) Flag indicating whether path MTU is enable or not for Vedge VPN interface type Feature Template, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "flow_control",
					Description: `(Optional) Value of bidirectional flow control settings for Vedge VPN interface type Feature Template, Allowed values: "autoneg", "both", "egress", "ingress", "none", Default: "autoneg"`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `(Optional) Maximum segment size (MSS) of TPC SYN packets passing through the router, Range for Cisco VPN Interface : [500, 1460], Range for Vedge VPN Interface : [552 - 1460]`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) Value of speed to use when the remote end of the connection does not support autonegotiation, Allowed values for Cisco VPN Interface type feature template: "10", "100", "1000", "10000", "2500", Allowed values for Vedge VPN Interface type feature template: "10", "100", "1000", "10000"`,
				},
				resource.Attribute{
					Name:        "clear_dont_fragment",
					Description: `(Optional) Flag indicating Don't Fragment (DF) bit in the IPv4 packet header for packets being transmitted out the interface is clar or not, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "static_ingess_qos",
					Description: `(Optional) Queue number to use for incoming traffic for Vedge VPN interface type Feature Template, Range: [0 - 7]`,
				},
				resource.Attribute{
					Name:        "arp_timeout",
					Description: `(Optional) Value indicating how long it takes for a dynamically learned ARP entry to time out, Range for Cisco VPN Interface : [0 - 2147483], Range for Vedge VPN Interface : [0 - 2678400] Default: 1500`,
				},
				resource.Attribute{
					Name:        "autonegotiation",
					Description: `(Optional) Flag indicating autonegotiation mode, Allowed value: true, false Default: true`,
				},
				resource.Attribute{
					Name:        "tloc_extension",
					Description: `(Optional) Name of a physical interface on the same router that connects to the WAN transport`,
				},
				resource.Attribute{
					Name:        "power_over_ethernet",
					Description: `(Optional) Flag indicating whether PoE(Power over Ethernet) is enabled or not for Vedge VPN interface type Feature Template, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "load_interval",
					Description: `(Optional) Value of load interval for Cisco VPN interface type Feature Template, Range: [30 - 600], Allowed values: value can only be multiple of 30`,
				},
				resource.Attribute{
					Name:        "tracker",
					Description: `(Optional) List of a tracker name to track the status of transport interfaces that connect to the internet for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "icmp_redirect_disable",
					Description: `(Optional) Flag indicating whether ICMP redirect disable or not on interface for VPN interface type Feature Template, Allowed value: true, false Default: false`,
				},
				resource.Attribute{
					Name:        "gre_tunnel_source_ip",
					Description: `(Optional) IP address of the extended WAN interface, Allowed values: IPv4 address`,
				},
				resource.Attribute{
					Name:        "xconnect",
					Description: `(Optional) The name of a physical interface on the same router that connects to the WAN transport for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ip_directed_broadcast",
					Description: `(Optional) Flag indicate whether directed IP braodcast is enabled or not for VPN interface type Feature Template, Allowed value: true, false Default: false <a id="nestedblock--ipv4"></a> ## Nested Schema for Basic ipv4`,
				},
				resource.Attribute{
					Name:        "primary_address",
					Description: `(Optional) Value of IPv4 prefix for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "secondary_address",
					Description: `(Optional) Value of IPv4 prefix for VPN interface type Feature Template, Max Items: 4(see [below for nested schema](#nestedblock--secondary_addressv4))`,
				},
				resource.Attribute{
					Name:        "dhcp_distance",
					Description: `(Optional) Value of IPv4 DHCP distance for VPN interface type Feature Template, Range: [1 - 65536]`,
				},
				resource.Attribute{
					Name:        "dhcp_helper",
					Description: `(Optional) List of IPv4 or IPv6 DHCP Helper for VPN interface type Feature Template, Allowed values: IPv4 or IPv6 address <a id="nestedblock--ipv6"></a> ## Nested Schema for Basic ipv6`,
				},
				resource.Attribute{
					Name:        "primary_address",
					Description: `(Optional) Value of IPv6 prefix for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "secondary_address",
					Description: `(Optional) Value of IPv6 prefix for VPN interface type Feature Template, Max Items: 2(see [below for nested schema](#nestedblock--secondary_addressv6))`,
				},
				resource.Attribute{
					Name:        "dhcp_distance",
					Description: `(Optional) Value of IPv6 DHCP distance for Vedge VPN interface type Feature Template, Range: [1 - 65536]`,
				},
				resource.Attribute{
					Name:        "dhcp_rapid_commit",
					Description: `(Optional) Value of IPv6 DHCP rapid commit enabled or not for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "dhcp_helper",
					Description: `(Optional) DHCP Helper value to designate the interface as a DHCP helper on a router for VPN interface type Feature Template, Max Items: 8(see [below for nested schema](#nestedblock--dhcp_helperV6)) <a id="nestedblock--allow_service"></a> ## Nested Schema for Tunnel allow_service`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Flag indicating whether All serivce is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "bgp",
					Description: `(Optional) Flag indicating whether BGP is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `(Optional) Flag indicating whether DHCP is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional) Flag indicating whether DNS is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "icmp",
					Description: `(Optional) Flag indicating whether ICMP is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "netconf",
					Description: `(Optional) Flag indicating whether NETCONF is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "ntp",
					Description: `(Optional) Flag indicating whether NTP is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `(Optional) Flag indicating whether OSPF is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "ssh",
					Description: `(Optional) Flag indicating whether SSH is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "https",
					Description: `(Optional) Flag indicating whether HTTPS is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "snmp",
					Description: `(Optional) Flag indicating whether SNMP is enabled or not for Cisco VPN interface type Feature Template, Default: false <a id="nestedblock--encapsulation"></a> ## Nested Schema for Tunnel encapsulation`,
				},
				resource.Attribute{
					Name:        "gre",
					Description: `(Optional) Flag indicating whether gre is enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "gre_preference",
					Description: `(Optional) Value of GRE preference for VPN interface type Feature Template, Range: [0 - 4294967295]`,
				},
				resource.Attribute{
					Name:        "gre_weight",
					Description: `(Optional) Value of GRE weight for VPN interface type Feature Template, Range: [1 - 255], Default: 1`,
				},
				resource.Attribute{
					Name:        "ipsec",
					Description: `(Optional) Flag indicating whether ipsec is enabled or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "ipsec_preference",
					Description: `(Optional) Value of IPsec preference for VPN interface type Feature Template, Range: [0 - 4294967295]`,
				},
				resource.Attribute{
					Name:        "ipsec_weight",
					Description: `(Optional) Value of IPsec weight for VPN interface type Feature Template, Range: [1 - 255], Default: 1`,
				},
				resource.Attribute{
					Name:        "carrier",
					Description: `(Optional) Carrier name or private network identifier to associate with the tunnel for VPN interface type Feature Template, Allowed values: "carrier1", "carrier2", "carrier3", "carrier4", "carrier5", "carrier6", "carrier7", "carrier8", "default", Default: "default"`,
				},
				resource.Attribute{
					Name:        "bind_loopback_tunnel",
					Description: `(Optional) Name of a physical interface to bind to a loopback interface for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "last_resort_circuit",
					Description: `(Optional) Flag indicating whether tunnel interface is used as the circuit of last resort for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `(Optional) Value of hold time for Vedge VPN interface type Feature Template, Range: [100 - 10000], Default: 7000`,
				},
				resource.Attribute{
					Name:        "nat_refresh_interval",
					Description: `(Optional) Value of interval between NAT refresh packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template, Range: [1 - 60], Default: 5`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `(Optional) Value of interval between Hello packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template, Range: [10 - 600000], Default: 1000`,
				},
				resource.Attribute{
					Name:        "hello_tolerance",
					Description: `(Optional) Value of the time to wait for a Hello packet on a DTLS or TLS WAN transport connection before declaring that transport tunnel to be down for VPN interface type Feature Template, Range: [12 - 6000], Default: 12`,
				},
				resource.Attribute{
					Name:        "gre_tunnel_destination_ip",
					Description: `(Optional) Value of GRE tunnel destination IP for VPN interface type Feature Template, Allowed value: IPv4 Address <a id="nestedblock--natv4"></a> ## Nested Schema for NAT ipv4`,
				},
				resource.Attribute{
					Name:        "nat_type",
					Description: `(Optional) NAT type for Cisco VPN interface type Feature Template, Allowed value: "interface", "pool", "loopback" Default: "interface"`,
				},
				resource.Attribute{
					Name:        "refresh_mode",
					Description: `(Optional) Refresh mode for VPN interface type Feature Template, Allowed value: "bi-directional", "outbound", Default: "outbound"`,
				},
				resource.Attribute{
					Name:        "log_nat_flow_creations_or_deletions",
					Description: `(Optional) Flag indicating whether log nat flow creations or deletions are enabled or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `(Optional) The time when NAT translations over UDP sessions time out for VPN interface type Feature Template, Range: [1 - 65536] minutes, Default: 1`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `(Optional) The time when NAT translations over TCP sessions time out for VPN interface type Feature Template, Range: [1 - 65536] minutes, Default: 60`,
				},
				resource.Attribute{
					Name:        "block_icmp",
					Description: `(Optional) Flag indicating whether to block icmp or not for VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "respond_to_ping",
					Description: `(Optional) Flag indicating whether to respond to ping or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "nat_pool_range_start",
					Description: `(Optional) Value of starting IP address for the NAT pool for VPN interface type Feature Template, Allowed value: IPv4 Address`,
				},
				resource.Attribute{
					Name:        "nat_pool_range_end",
					Description: `(Optional) Value of closing IP address for the NAT pool for VPN interface type Feature Template, Allowed value: IPv4 Address`,
				},
				resource.Attribute{
					Name:        "nat_pool_prefix_length",
					Description: `(Optional) Value of NAT pool prefix length for Cisco VPN interface type Feature Template, Range: [1 - 65536], Default: 60`,
				},
				resource.Attribute{
					Name:        "overload",
					Description: `(Optional) Flag indicating whether to enable per-port translation or not for Cisco VPN interface type Feature Template, Default: true`,
				},
				resource.Attribute{
					Name:        "nat_inside_source_loopback_interface",
					Description: `(Optional) Value of NAT inside source loopback interface for Cisco VPN interface type Feature Template, String length Range: [1 - 32]`,
				},
				resource.Attribute{
					Name:        "port_forward",
					Description: `(Optional) Configure port forwarding rule for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--port_forward))`,
				},
				resource.Attribute{
					Name:        "static_nat",
					Description: `(Optional) Configure static nat for VPN interface type Feature Template(see [below for nested schema](#nestedblock--static_nat)) <a id="nestedblock--natv6"></a> ## Nested Schema for NAT ipv6`,
				},
				resource.Attribute{
					Name:        "nat64",
					Description: `(Optional) Flag indicating whether to enable nat64 or not for VPN interface type Feature Template, Default: false <a id="nestedblock--vrrpv4"></a> ## Nested Schema for VRRP ipv4`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template, Range: [1 - 255]`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template, Range: [1 - 254], Default: 100`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `(Optional) Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template, Range for Vedge VPN interface: [1 - 3600] Seconds, Default: 1 Second, Range for Cisco VPN interface: [100 - 40950] Milliseconds, Default: 100 Milliseconds`,
				},
				resource.Attribute{
					Name:        "track_omp",
					Description: `(Optional) Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "track_prefix_list",
					Description: `(Optional) Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Value of IPv4 address of the virtual router for VPN interface type Feature Template, Allowed value: IPv4 Address <a id="nestedblock--vrrpv6"></a> ## Nested Schema for VRRP ipv6`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template, Range: [1 - 255]`,
				},
				resource.Attribute{
					Name:        "link_local_ipv6_address",
					Description: `(Required) Value of IPv6 address of the virtual router for VPN interface type Feature Template, Allowed value: IPv6 Address`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template, Range: [1 - 254], Default: 100`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `(Optional) Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template, Range for Vedge VPN interface: [1 - 3600] Seconds, Range for Cisco VPN interface: [1 - 40950] Milliseconds, Default: 1 Second`,
				},
				resource.Attribute{
					Name:        "track_omp",
					Description: `(Optional) Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "track_prefix_list",
					Description: `(Optional) Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "global_ipv6_prefix",
					Description: `(Optional) Value of global IPv6 prefix of the virtual router for VPN interface type Feature Template <a id="nestedblock--advanced_options"></a> ## Nested Schema for 8021.X advanced_options`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) Configure VLAN for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--vlan))`,
				},
				resource.Attribute{
					Name:        "dynamic_authentication_server",
					Description: `(Optional) Configure Dynamic Authentication Server for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--dynamic_authentication_server))`,
				},
				resource.Attribute{
					Name:        "mac_authentication_bypass",
					Description: `(Optional) Configure MAC Authentication Bypass for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--mac_authentication_bypass))`,
				},
				resource.Attribute{
					Name:        "request_attributes",
					Description: `(Optional) Configure Request Attributes for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--request_attributes)) <a id="nestedblock--secondary_addressv4"></a> ## Nested Schema for Basic IPv4 secondary_address`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Value of IPv4 prefix for VPN interface type Feature Template <a id="nestedblock--secondary_addressv6"></a> ## Nested Schema for Basic IPv6 secondary_address`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Value of IPv6 prefix for VPN interface type Feature Template <a id="nestedblock--dhcp_helperV6"></a> ## Nested Schema for Basic IPv6 dhcp_helper`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Value of IPv6 address for VPN interface type Feature Template, Allowed value: IPv6 Address`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Value of VPN id for VPN interface type Feature Template, Range: [1 - 65536] <a id="nestedblock--port_forward"></a> ## Nested Schema for NAT ipv4 port_forward`,
				},
				resource.Attribute{
					Name:        "port_start_range",
					Description: `(Required) Value of port starting range for Vedge VPN interface type Feature Template, Range: [0 - 65535]`,
				},
				resource.Attribute{
					Name:        "port_end_range",
					Description: `(Required) Value of port ending range for Vedge VPN interface type Feature Template, Range: [0 - 65535]`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol for Vedge VPN interface type Feature Template, Allowed value: "tcp", "udp"`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Required) VPN id for Vedge VPN interface type Feature Template, Range: [0 - 65535]`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Required) Private IP address for Vedge VPN interface type Feature Template, Allowed value: IPv4 Address <a id="nestedblock--static_nat"></a> ## Nested Schema for NAT ipv4 static_nat`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `(Required) Value of the inside local address as source IP address for VPN interface type Feature Template, Allowed value: IPv4 Address`,
				},
				resource.Attribute{
					Name:        "translated_source_ip",
					Description: `(Required) Value of the inside global address as the translated source IP address for VPN interface type Feature Template, Allowed value: IPv4 Address`,
				},
				resource.Attribute{
					Name:        "source_vpn_id",
					Description: `(Optional) Configures Source VPN ID, which is the Service VPN in which the host resides for VPN interface type Feature Template, Range: [0 - 65530], Default: 0`,
				},
				resource.Attribute{
					Name:        "static_nat_direction",
					Description: `(Optional) Set the direction in which to perform network address translation for VPN interface type Feature Template, Allowed value for Vedge VPN interface type Feature Template : "inside", "outside" Default: "inside" Allowed value for Cisco VPN interface type Feature Template : "inside"`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol for Vedge VPN interface type Feature Template, Allowed value: "tcp", "udp"`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) Value of source port for Vedge VPN interface type Feature Template, Range: [0 - 65535]`,
				},
				resource.Attribute{
					Name:        "translate_port",
					Description: `(Optional) Value of translate port for Vedge VPN interface type Feature Template, Range: [0 - 65535] <a id="nestedblock--vlan"></a> ## Nested Schema for vlan`,
				},
				resource.Attribute{
					Name:        "authentication_fail_vlan",
					Description: `(Optional) Value of authentication fail VLAN for Vedge VPN interface type Feature Template, Range: [1 - 65536]`,
				},
				resource.Attribute{
					Name:        "guest_vlan",
					Description: `(Optional) Value of guest VLAN for Vedge VPN interface type Feature Template, Range: [1 - 65536]`,
				},
				resource.Attribute{
					Name:        "authentication_reject_vlan",
					Description: `(Optional) Value of authentication reject VLAN for Vedge VPN interface type Feature Template, Range: [1 - 65536]`,
				},
				resource.Attribute{
					Name:        "default_vlan",
					Description: `(Optional) Value of default VLAN for Vedge VPN interface type Feature Template, Range: [1 - 65536] <a id="nestedblock--dynamic_authentication_server"></a> ## Nested Schema for dynamic_authentication_server`,
				},
				resource.Attribute{
					Name:        "das_port",
					Description: `(Optional) Value of Dynamic Authentication Server port for Vedge VPN interface type Feature Template, Range: [1 - 65535], Default: 3799,`,
				},
				resource.Attribute{
					Name:        "client",
					Description: `(Optional) Value of Client IP address for Vedge VPN interface type Feature Template, Allowed value: IPv4 Address`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) Secret key Value for Vedge VPN interface type Feature Template`,
				},
				resource.Attribute{
					Name:        "time_window",
					Description: `(Optional) Time window value for Vedge VPN interface type Feature Template, Range: [10 - 1000] Seconds, Default: 300,`,
				},
				resource.Attribute{
					Name:        "required_timestamp",
					Description: `(Optional) Flag indicating whether to enable required timestamp or not for Vedge VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Value of VPN ID for Vedge VPN interface type Feature Template, Range: [0 - 65535], Default: 0 <a id="nestedblock--mac_authentication_bypass"></a> ## Nested Schema for mac_authentication_bypass`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) Flag indicating whether to enable server or not for Vedge VPN interface type Feature Template, Default: false`,
				},
				resource.Attribute{
					Name:        "mac_authentication_bypass_entries",
					Description: `(Optional) List of MAC Authenticaion bypass entries for Vedge VPN interface type Feature Template, Max Items: 8 <a id="nestedblock--request_attributes"></a> ## Nested Schema for request_attributes`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) Configure Authentication for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--authentication))`,
				},
				resource.Attribute{
					Name:        "accounting",
					Description: `(Optional) Configure Accounting for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--accounting)) <a id="nestedblock--authentication"></a> ## Nested Schema for request_attributes authentication`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID for Vedge VPN interface type Feature Template, Range: [1 - 64]`,
				},
				resource.Attribute{
					Name:        "syntax_choice",
					Description: `(Required) Syntax Choice for Vedge VPN interface type Feature Template, Allowed value: "string", "integer", "octet"`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value for Vedge VPN interface type Feature Template <a id="nestedblock--accounting"></a> ## Nested Schema for request_attributes accounting`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID for Vedge VPN interface type Feature Template, Range: [1 - 64]`,
				},
				resource.Attribute{
					Name:        "syntax_choice",
					Description: `(Required) Syntax Choice for Vedge VPN interface type Feature Template, Allowed value: "string", "integer", "octet"`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value for Vedge VPN interface type Feature Template`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"sdwan_device_template":                0,
		"sdwan_devices_attachment":             1,
		"sdwan_ntp_feature_template":           2,
		"sdwan_system_feature_template":        3,
		"sdwan_vpn_feature_template":           4,
		"sdwan_vpn_interface_feature_template": 5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
