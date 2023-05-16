package fortiadc

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_ca_certificateupload",
			Category:         "FortiADC Resources",
			ShortDescription: `Upload fortiadc ca.`,
			Description:      ``,
			Keywords: []string{
				"ca",
				"certificateupload",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `cert file path.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_certificate_crlupload",
			Category:         "FortiADC Resources",
			ShortDescription: `Upload fortiadc certificate crl.`,
			Description:      ``,
			Keywords: []string{
				"certificate",
				"crlupload",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `cert file path.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_certificate_intmed_caupload",
			Category:         "FortiADC Resources",
			ShortDescription: `Upload fortiadc intermediate ca.`,
			Description:      ``,
			Keywords: []string{
				"certificate",
				"intmed",
				"caupload",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `cert file path.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_config_sync_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc config sync list.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"config",
				"sync",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `config sync list name.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `comment for sync-list.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `filename.`,
				},
				resource.Attribute{
					Name:        "server_ip",
					Description: `ip address of server.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `peer admin password.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `synchronization type. Valid values: 8192:share-resource, 1024:security, 32:log, 1:system, 4:networking, 128:link-load-balance, 8:load-balance, 512:global-load-balance, 16384:user .`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `port of server. (1,65535) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Config Sync List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_config_sync_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Config Sync List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_config_sync_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_firewall_nat_snat",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc snat.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"firewall",
				"nat",
				"snat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `The name of the snat.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(en|dis)able snat status. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "out_interface",
					Description: `Output interface name.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `Source network.`,
				},
				resource.Attribute{
					Name:        "traffic_group",
					Description: `traffic group name.`,
				},
				resource.Attribute{
					Name:        "trans_to_ip_end",
					Description: `Trans-to-pool ip max.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `Destination network.`,
				},
				resource.Attribute{
					Name:        "trans_to_type",
					Description: `Trans to type ip or ip pool or no-nat. Valid values: 1:pool, 0:ip, 2:no-nat .`,
				},
				resource.Attribute{
					Name:        "trans_to_ip_start",
					Description: `Trans-to-pool ip min.`,
				},
				resource.Attribute{
					Name:        "trans_to_ip",
					Description: `Trans-to ip. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Firewall Nat Snat can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_firewall_nat_snat.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Firewall Nat Snat can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_firewall_nat_snat.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_firewall_vip",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc virtual IP.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"firewall",
				"vip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `The name of the virtual IP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(en|dis)able static nat status. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "mappedip_max",
					Description: `max mapped IP address.`,
				},
				resource.Attribute{
					Name:        "mappedport_min",
					Description: `min mapped port .`,
				},
				resource.Attribute{
					Name:        "no_nat",
					Description: `(en|dis)able static nat no-nat support. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "proto",
					Description: `protocol type. Valid values: 17:udp, 6:tcp .`,
				},
				resource.Attribute{
					Name:        "traffic_group",
					Description: `traffic group name.`,
				},
				resource.Attribute{
					Name:        "extif",
					Description: `external interface name.`,
				},
				resource.Attribute{
					Name:        "mappedip_min",
					Description: `min mapped IP address.`,
				},
				resource.Attribute{
					Name:        "mappedport_max",
					Description: `max mapped port .`,
				},
				resource.Attribute{
					Name:        "extport",
					Description: `external port.`,
				},
				resource.Attribute{
					Name:        "portforward",
					Description: `(en|dis)able port forwarding. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "extip",
					Description: `external IP address. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Firewall Vip can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_firewall_vip.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Firewall Vip can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_firewall_vip.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_allowlist",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure geography ip allow list..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"allowlist",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name of allow list..`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `simple description.. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Allowlist can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_allowlist.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Allowlist can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_allowlist.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_auth_policy",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc authentication policy.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"auth",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Auth Policy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_auth_policy.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Auth Policy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_auth_policy.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_caching",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance caching info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"caching",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `caching name.`,
				},
				resource.Attribute{
					Name:        "max_age",
					Description: `max age of cache entries. (60,86400)`,
				},
				resource.Attribute{
					Name:        "max_cache_size",
					Description: `max cache size.`,
				},
				resource.Attribute{
					Name:        "max_entries",
					Description: `max cache entries. (1,262144)`,
				},
				resource.Attribute{
					Name:        "max_object_size",
					Description: `maximum object size. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Caching can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_caching.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Caching can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_caching.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_caching_child_dyn_cache_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance caching info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"caching",
				"child",
				"dyn",
				"cache",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `dynamic cache rule id.`,
				},
				resource.Attribute{
					Name:        "invalid_uri",
					Description: `invalid uri pattern.`,
				},
				resource.Attribute{
					Name:        "age",
					Description: `cache age out time. (1,86400)`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `dynamic cache uri pattern. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Caching Child Dyn Cache List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_caching_child_dyn_cache_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Caching Child Dyn Cache List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_caching_child_dyn_cache_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_caching_child_uri_exclude_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance caching info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"caching",
				"child",
				"uri",
				"exclude",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `uri list id.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `uri string. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Caching Child Uri Exclude List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_caching_child_uri_exclude_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Caching Child Uri Exclude List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_caching_child_uri_exclude_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_captcha_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance captcha profile.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"captcha",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `captcha profile name.`,
				},
				resource.Attribute{
					Name:        "max_block_period",
					Description: `User blocked time after verification failure, unit second. (10,864000)`,
				},
				resource.Attribute{
					Name:        "custom_page",
					Description: `If use customize captcha page file.. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "max_valid_period",
					Description: `Maximum valid time for verified users, unit second. (60,864000)`,
				},
				resource.Attribute{
					Name:        "max_picture_changes",
					Description: `Maximum number of times the user can change the captcha picture. (1,100)`,
				},
				resource.Attribute{
					Name:        "max_verify_period",
					Description: `Maximum user verification process, unit second. (20,86400)`,
				},
				resource.Attribute{
					Name:        "max_retry_times",
					Description: `The maximum number of times a user can try. (1,100)`,
				},
				resource.Attribute{
					Name:        "vpath",
					Description: `virtual path of captcha service.`,
				},
				resource.Attribute{
					Name:        "custom_page_file",
					Description: `Customize captcha page package filename.`,
				},
				resource.Attribute{
					Name:        "pic_difficulty",
					Description: `Verification picture difficulty. Valid values: 1:hard, 0:easy . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Captcha Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_captcha_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Captcha Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_captcha_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_certificate_caching",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance certificate-caching info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"certificate",
				"caching",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `certificate-caching name.`,
				},
				resource.Attribute{
					Name:        "max_entries",
					Description: `max certificate cache entries. (1,1000000)`,
				},
				resource.Attribute{
					Name:        "max_certificate_cache_size",
					Description: `max certificate cache size. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Certificate Caching can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_certificate_caching.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Certificate Caching can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_certificate_caching.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_client_ssl_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance client-ssl-profile info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"client",
				"ssl",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `profile name.`,
				},
				resource.Attribute{
					Name:        "ssl_allowed_versions",
					Description: `SSL/TLS versions allowed. Valid values: 8:tlsv1.1, 2:sslv3, 4:tlsv1.0, 32:tlsv1.3, 16:tlsv1.2 .`,
				},
				resource.Attribute{
					Name:        "ssl_renegotiate_size",
					Description: `size (MB) of application data transmitted over SSL when ADC initiate renegotiation, set to 0 to disable ADC initiate renegotiation by size. (0,2147483647)`,
				},
				resource.Attribute{
					Name:        "backend_certificate_verify",
					Description: `backend certificate verify for SSL forward proxy.`,
				},
				resource.Attribute{
					Name:        "backend_customized_ssl_ciphers",
					Description: `backend SSL user-customized ciphers for SSL forward proxy (concatenate multiple ciphers by :).`,
				},
				resource.Attribute{
					Name:        "ssl_renegotiation",
					Description: `(en|dis)able renegotiation. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "customized_ssl_ciphers_flag",
					Description: `(en|dis)able SSL user-customized ciphers. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "backend_customized_ssl_ciphers_flag",
					Description: `(en|dis)able backend SSL user-customized ciphers for SSL forward proxy. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "ssl_secure_renegotiation",
					Description: `set secure renegotiation. Valid values: 1:request, 3:require-strict, 2:require .`,
				},
				resource.Attribute{
					Name:        "http_forward_client_certificate_header",
					Description: `forward client certificate header (https profile only).`,
				},
				resource.Attribute{
					Name:        "supported_groups",
					Description: `ssl Supported Groups. Valid values: 1:secp256r1, 0:x25519, 3:secp521r1, 2:x448, 5:ffdhe2048, 4:secp384r1, 7:ffdhe4096, 6:ffdhe3072, 9:ffdhe8192, 8:ffdhe6144 .`,
				},
				resource.Attribute{
					Name:        "http_forward_client_certificate",
					Description: `foward client certificate (https profile only). Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "use_tls_tickets",
					Description: `(en|dis)able TLS tickets support (When SSL session ID based persistence is used in a VS, this flag is ignored and tls ticket is disabled). Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "client_sni_required",
					Description: `client SNI required. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "backend_ssl_ciphers",
					Description: `backend ssl ciphers including Perfect Forward Secrecy for SSL forward proxy. Valid values: 32768:, 4194304:Perfect Forward Secrecy cipher, 1099511627776:Perfect Forward Secrecy cipher, 131072:Perfect Forward Secrecy cipher, 262144:Perfect Forward Secrecy cipher, 64:Perfect Forward Secrecy cipher, 65536:, 256:Perfect Forward Secrecy cipher, 33554432:, 8192:Perfect Forward Secrecy cipher, 16777216:, 536870912:Perfect Forward Secrecy cipher, 1048576:Perfect Forward Secrecy cipher, 1:Perfect Forward Secrecy cipher, 2097152:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 16384:, 17179869184:Non Encryption Ciphers, 2199023255552:Perfect Forward Secrecy cipher, 4096:Perfect Forward Secrecy cipher, 134217728:, 137438953472:Perfect Forward Secrecy cipher, 274877906944:Perfect Forward Secrecy cipher, 8388608:, 67108864:Perfect Forward Secrecy cipher, 524288:Perfect Forward Secrecy cipher, 549755813888:Perfect Forward Secrecy cipher, 128:Perfect Forward Secrecy cipher, 2147483648:, 1024:Perfect Forward Secrecy cipher, 268435456:, 34359738368:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher, 32:Perfect Forward Secrecy cipher, 2048:Perfect Forward Secrecy cipher, 1073741824:Perfect Forward Secrecy cipher, 512:Perfect Forward Secrecy cipher, 68719476736:Perfect Forward Secrecy cipher .`,
				},
				resource.Attribute{
					Name:        "client_certificate_verify",
					Description: `certificate verify name.`,
				},
				resource.Attribute{
					Name:        "backend_ssl_sni_forward",
					Description: `(en|dis)able backend SSL SNI forward for SSL forward proxy. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "client_certificate_verify_mode",
					Description: `certificate verify option. Valid values: 1:required, 2:optional .`,
				},
				resource.Attribute{
					Name:        "local_certificate_group",
					Description: `local certificate group name.`,
				},
				resource.Attribute{
					Name:        "backend_ssl_allowed_versions",
					Description: `backend SSL/TLS versions allowed for SSL forward proxy. Valid values: 8:tlsv1.1, 2:sslv3, 4:tlsv1.0, 32:tlsv1.3, 16:tlsv1.2 .`,
				},
				resource.Attribute{
					Name:        "ssl_dh_param_size",
					Description: `The maximum size of the Diffie-Hellman parameters used for generating the enphemeral/temporary Diffie-Hellman key for DHE key exchange. Default value is 1024 bits.. Valid values: 1024:1024bit, 4096:4096bit, 2048:2048bit .`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers_tlsv13",
					Description: `tlsv1.3 ciphers. Valid values: 1:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher .`,
				},
				resource.Attribute{
					Name:        "backend_ciphers_tlsv13",
					Description: `backend tlsv1.3 ciphers. Valid values: 1:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher .`,
				},
				resource.Attribute{
					Name:        "ssl_renegotiate_period",
					Description: `period (s|m|h) (by default in seconds) of ADC initiate renegotiation, set to 0 to disable periodical renegotiation.`,
				},
				resource.Attribute{
					Name:        "forward_proxy_certificate_caching",
					Description: `certificate caching name.`,
				},
				resource.Attribute{
					Name:        "forward_proxy_local_signing_ca",
					Description: `local signing CA for SSL forward proxy.`,
				},
				resource.Attribute{
					Name:        "customized_ssl_ciphers",
					Description: `SSL user-customized ciphers (concatenate multiple ciphers by :).`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `ssl ciphers including Perfect Forward Secrecy. Valid values: 32768:, 4194304:Perfect Forward Secrecy cipher, 1099511627776:Perfect Forward Secrecy cipher, 131072:Perfect Forward Secrecy cipher, 262144:Perfect Forward Secrecy cipher, 64:Perfect Forward Secrecy cipher, 65536:, 256:Perfect Forward Secrecy cipher, 33554432:, 8192:Perfect Forward Secrecy cipher, 16777216:, 536870912:Perfect Forward Secrecy cipher, 1048576:Perfect Forward Secrecy cipher, 1:Perfect Forward Secrecy cipher, 2097152:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 16384:, 17179869184:Non Encryption Ciphers, 2199023255552:Perfect Forward Secrecy cipher, 4096:Perfect Forward Secrecy cipher, 134217728:, 137438953472:Perfect Forward Secrecy cipher, 274877906944:Perfect Forward Secrecy cipher, 8388608:, 67108864:Perfect Forward Secrecy cipher, 524288:Perfect Forward Secrecy cipher, 549755813888:Perfect Forward Secrecy cipher, 128:Perfect Forward Secrecy cipher, 2147483648:, 1024:Perfect Forward Secrecy cipher, 268435456:, 34359738368:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher, 32:Perfect Forward Secrecy cipher, 2048:Perfect Forward Secrecy cipher, 1073741824:Perfect Forward Secrecy cipher, 512:Perfect Forward Secrecy cipher, 68719476736:Perfect Forward Secrecy cipher .`,
				},
				resource.Attribute{
					Name:        "reject_ocsp_stapling_with_missing_nextupdate",
					Description: `reject OCSP stapling with missing nextupdate. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "ssl_dynamic_record_sizing",
					Description: `Adjust the SSL record size dynamically according to the state of the connection. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "rfc7919_comply",
					Description: `(en|dis)able SSL RFC7919 Comply. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "backend_ssl_ocsp_stapling_support",
					Description: `(en|dis)able backend SSL OCSP stapling support for SSL forward proxy. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "ssl_session_cache_flag",
					Description: `(en|dis)able SSL session cache. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "forward_proxy",
					Description: `ssl forward proxy mode (https profile only). Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "ssl_renegotiation_interval",
					Description: `minimum renegotiation interval time (s|m|h) (by default in seconds), set to -1 to disable client-initiated renegotiation.`,
				},
				resource.Attribute{
					Name:        "forward_proxy_intermediate_ca_group",
					Description: `intermediate ca group for SSL forward proxy. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Client Ssl Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_client_ssl_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Client Ssl Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_client_ssl_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_clone_pool",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance clone pool info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"clone",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `pool name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Clone Pool can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_clone_pool.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Clone Pool can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_clone_pool.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_clone_pool_child_pool_member",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance clone pool info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"clone",
				"pool",
				"child",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `server name.`,
				},
				resource.Attribute{
					Name:        "dst_mac",
					Description: `destination hardware address.`,
				},
				resource.Attribute{
					Name:        "src_mac",
					Description: `source hardware address.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `ip address of server.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `interface name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `server port. (0,65535)`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `mode of update. Valid values: 1:mirror-dst-mac-update, 0:mirror-interface, 3:mirror-src-dst-mac-update, 2:mirror-src-mac-update, 4:mirror-ip-update . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Clone Pool Child Pool Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_clone_pool_child_pool_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Clone Pool Child Pool Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_clone_pool_child_pool_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_compression",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance compression info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"compression",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `compression name.`,
				},
				resource.Attribute{
					Name:        "max_cpu_usage",
					Description: `max cpu usage. (1,100)`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `compression level. (1,9)`,
				},
				resource.Attribute{
					Name:        "uri_list_type",
					Description: `uri list type. Valid values: 1:include, 2:exclude .`,
				},
				resource.Attribute{
					Name:        "gzip_window_size",
					Description: `gzip window size in Kbytes. Valid values: 11:8, 10:4, 13:32, 12:16, 15:128, 14:64, 9:2, 8:1 .`,
				},
				resource.Attribute{
					Name:        "gzip_memory_level",
					Description: `gzip memory level in Kbytes. Valid values: 1:1, 3:4, 2:2, 5:16, 4:8, 7:64, 6:32, 9:256, 8:128 .`,
				},
				resource.Attribute{
					Name:        "cpu_limit",
					Description: `(en|dis)able cpu limit. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "min_content_length",
					Description: `minimum content length. (1,2147483647)`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `compression method. Valid values: 1:gzip, 3:all, 2:deflate . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Compression can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_compression.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Compression can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_compression.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_compression_child_content_types",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance compression info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"compression",
				"child",
				"content",
				"types",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `content type id.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `content type string. Valid values: 10:custom, 1:application/soap+xml, 3:text/html, 2:application/xml, 5:text/css, 4:text/plain, 7:application/javascript, 6:application/x-javascript, 9:text/xml, 8:text/javascript .`,
				},
				resource.Attribute{
					Name:        "custom_content_type",
					Description: `custom content type. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Compression Child Content Types can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_compression_child_content_types.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Compression Child Content Types can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_compression_child_content_types.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_compression_child_uri_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance compression info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"compression",
				"child",
				"uri",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `uri list id.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `uri string. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Compression Child Uri List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_compression_child_uri_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Compression Child Uri List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_compression_child_uri_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_content_rewriting",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance content rewriting info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"content",
				"rewriting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `content rewriting name.`,
				},
				resource.Attribute{
					Name:        "redirect",
					Description: `redirect content.`,
				},
				resource.Attribute{
					Name:        "header_name",
					Description: `http header name.`,
				},
				resource.Attribute{
					Name:        "referer_content",
					Description: `referer content.`,
				},
				resource.Attribute{
					Name:        "referer_status",
					Description: `replace referer header content. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `content rewriting comments.`,
				},
				resource.Attribute{
					Name:        "host_status",
					Description: `replace host header content. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "url_status",
					Description: `replace URL. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "action_type",
					Description: `action type. Valid values: 1:request, 2:response .`,
				},
				resource.Attribute{
					Name:        "header_value",
					Description: `http header value.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `action. Valid values: 1:rewrite_http_header, 3:send-403-forbidden, 2:redirect, 5:add_http_header, 4:rewrite_http_location, 6:delete_http_header .`,
				},
				resource.Attribute{
					Name:        "url_content",
					Description: `URL content.`,
				},
				resource.Attribute{
					Name:        "host_content",
					Description: `host content.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `location content. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Content Rewriting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_content_rewriting.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Content Rewriting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_content_rewriting.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_content_rewriting_child_match_condition",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance content rewriting info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"content",
				"rewriting",
				"child",
				"match",
				"condition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `match condition id.`,
				},
				resource.Attribute{
					Name:        "ignorecase",
					Description: `ignore case. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `reverse match. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `match object. Valid values: 1:http-host-header, 3:http-referer-header, 2:http-request-url, 5:http-location-header, 4:ip-source-address .`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `match content.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `match type. Valid values: 1:string, 2:regular-expression . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Content Rewriting Child Match Condition can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_content_rewriting_child_match_condition.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Content Rewriting Child Match Condition can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_content_rewriting_child_match_condition.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_content_routing",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance content routing info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"content",
				"routing",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `content routing name.`,
				},
				resource.Attribute{
					Name:        "source_pool_list",
					Description: `ip pool name.`,
				},
				resource.Attribute{
					Name:        "schedule_list",
					Description: `enable/disable schedule list. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `persistence name.`,
				},
				resource.Attribute{
					Name:        "connection_pool",
					Description: `connection pool name.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `ipv4 subnet address.`,
				},
				resource.Attribute{
					Name:        "packet_fwd_method",
					Description: `packet forwarding method. Valid values: 0:inherit, 3:FullNAT .`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `content routing comments.`,
				},
				resource.Attribute{
					Name:        "connection_pool_inherit",
					Description: `connection pool inherit. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "ip6",
					Description: `ipv6 subnet address.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `content-routing type. Valid values: 1:l4-content-routing, 2:l7-content-routing .`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `pool name.`,
				},
				resource.Attribute{
					Name:        "persistence_inherit",
					Description: `persistence inherit. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "schedule_pool_list",
					Description: `schedule pool name.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `method name.`,
				},
				resource.Attribute{
					Name:        "method_inherit",
					Description: `method inherit. Valid values: enable/disable. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Content Routing can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_content_routing.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Content Routing can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_content_routing.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_content_routing_child_match_condition",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance content routing info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"content",
				"routing",
				"child",
				"match",
				"condition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `match condition id.`,
				},
				resource.Attribute{
					Name:        "ignorecase",
					Description: `ignore case. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `reverse match. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `match object. Valid values: 1:http-host-header, 3:http-referer-header, 2:http-request-url, 4:ip-source-address, 6:sni .`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `match content.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `match type. Valid values: 1:string, 2:regular-expression . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Content Routing Child Match Condition can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_content_routing_child_match_condition.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Content Routing Child Match Condition can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_content_routing_child_match_condition.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_decompression",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance decompression info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"decompression",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `decompression name.`,
				},
				resource.Attribute{
					Name:        "max_cpu_usage",
					Description: `max cpu usage. (1,100)`,
				},
				resource.Attribute{
					Name:        "uri_list_type",
					Description: `uri list type. Valid values: 1:include, 2:exclude .`,
				},
				resource.Attribute{
					Name:        "cpu_limit",
					Description: `(en|dis)able cpu limit. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `decompression method. Valid values: 1:gzip, 3:all, 2:deflate . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Decompression can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_decompression.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Decompression can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_decompression.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_decompression_child_content_types",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance decompression info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"decompression",
				"child",
				"content",
				"types",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `content type id.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `content type string. Valid values: 10:custom, 1:application/soap+xml, 3:text/html, 2:application/xml, 5:text/css, 4:text/plain, 7:application/javascript, 6:application/x-javascript, 9:text/xml, 8:text/javascript .`,
				},
				resource.Attribute{
					Name:        "custom_content_type",
					Description: `custom content type. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Decompression Child Content Types can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_decompression_child_content_types.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Decompression Child Content Types can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_decompression_child_content_types.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_decompression_child_uri_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance decompression info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"decompression",
				"child",
				"uri",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `uri list id.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `uri string. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Decompression Child Uri List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_decompression_child_uri_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Decompression Child Uri List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_decompression_child_uri_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_geoip_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure geography ip block list..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"geoip",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name of block list..`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `enable/disable blocking log.. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `default action. Valid values: 1:pass, 2:deny, 5:send-403-forbidden, 4:redirect .`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `High, Medium or Low. Valid values: 1:low, 3:high, 2:medium . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Geoip List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_geoip_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Geoip List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_geoip_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_http2_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance HTTP/2 profile.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"http2",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Profile name.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_stream",
					Description: `Maximum number of concurrent stream. (1,200)`,
				},
				resource.Attribute{
					Name:        "max_frame_size",
					Description: `Maximum size of frame. (16384,131072)`,
				},
				resource.Attribute{
					Name:        "priority_mode",
					Description: `Priority of stream mode. Valid values: 0:best-effort .`,
				},
				resource.Attribute{
					Name:        "ssl_constraint",
					Description: `SSL constraint check. Valid values: 1:enable, 0:disable .`,
				},
				resource.Attribute{
					Name:        "max_receive_window",
					Description: `Maximum size of receive window. (16384,524288)`,
				},
				resource.Attribute{
					Name:        "header_table_size",
					Description: `size of header table for HPACK. (4096,65536)`,
				},
				resource.Attribute{
					Name:        "upgrade_mode",
					Description: `Protocol upgrade to HTTP/2 mode. Valid values: 0:upgradeable .`,
				},
				resource.Attribute{
					Name:        "max_header_list_size",
					Description: `Maximum size of header list. (4096,262144) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Http2 Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_http2_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Http2 Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_http2_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_ippool",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance ip pool info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"ippool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `pool name.`,
				},
				resource.Attribute{
					Name:        "pool_type",
					Description: `address type. Valid values: 1:ipv4, 2:ipv6 .`,
				},
				resource.Attribute{
					Name:        "ip_end",
					Description: `IP max.`,
				},
				resource.Attribute{
					Name:        "ip6_end",
					Description: `IPv6 max.`,
				},
				resource.Attribute{
					Name:        "ip_start",
					Description: `IP min.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `interface name.`,
				},
				resource.Attribute{
					Name:        "ip6_start",
					Description: `IPv6 min. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Ippool can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_ippool.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Ippool can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_ippool.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_ippool_child_node_member",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance ip pool info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"ippool",
				"child",
				"node",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `pool name.`,
				},
				resource.Attribute{
					Name:        "ip6_max",
					Description: `IPv6 max.`,
				},
				resource.Attribute{
					Name:        "ip6_min",
					Description: `IPv6 min.`,
				},
				resource.Attribute{
					Name:        "ha_node_num",
					Description: `ha-node id. (0,7)`,
				},
				resource.Attribute{
					Name:        "ip_min",
					Description: `IP min.`,
				},
				resource.Attribute{
					Name:        "pool_type",
					Description: `address type. Valid values: 1:ipv4, 2:ipv6 .`,
				},
				resource.Attribute{
					Name:        "ip_max",
					Description: `IP max.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `interface name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Ippool Child Node Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_ippool_child_node_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Ippool Child Node Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_ippool_child_node_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_l2_exception_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance layer2 exception list.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"l2",
				"exception",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `layer2 exception list description.`,
				},
				resource.Attribute{
					Name:        "web_filter_profile",
					Description: `layer2 exception web filter profile. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance L2 Exception List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_l2_exception_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance L2 Exception List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_l2_exception_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_l2_exception_list_child_member",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance layer2 exception list.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"l2",
				"exception",
				"list",
				"child",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `id.`,
				},
				resource.Attribute{
					Name:        "ip_netmask",
					Description: `ip netmask.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `comments.`,
				},
				resource.Attribute{
					Name:        "host_pattern",
					Description: `host wildcard pattern.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `match type: host wildcard pattern/ip netmask. Valid values: 1:match type ip netmask, 0:match type host wildcard pattern . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance L2 Exception List Child Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_l2_exception_list_child_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance L2 Exception List Child Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_l2_exception_list_child_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_method",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance method info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"method",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `method table name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `method type. Valid values: 1:round-robin, 3:fastest-response, 2:least-connection, 5:full-uri-hash, 4:uri-hash, 7:host-domain-hash, 6:host-hash, 9:dynamic-load, 8:dest-ip-hash . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Method can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_method.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Method can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_method.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_pagespeed",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance pagespeed.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"pagespeed",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `pagespeed name.`,
				},
				resource.Attribute{
					Name:        "pagespeed_profile_id",
					Description: `pagespeed profile name.`,
				},
				resource.Attribute{
					Name:        "file_cache_max_inode",
					Description: `Maximum file cache inode. (1,100000)`,
				},
				resource.Attribute{
					Name:        "file_cache_max_size",
					Description: `Maximum file cache size (1M ~512M). (1,512) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pagespeed can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pagespeed.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pagespeed can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pagespeed.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_pagespeed_child_page_control",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance pagespeed.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"pagespeed",
				"child",
				"page",
				"control",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `ID.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `include/exclue type. Valid values: 1:include, 2:exclude .`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `full uri wildcards, such as http(s)://`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pagespeed Child Page Control can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pagespeed_child_page_control.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pagespeed Child Page Control can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pagespeed_child_page_control.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_pagespeed_child_resource_control",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance pagespeed.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"pagespeed",
				"child",
				"resource",
				"control",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `ID.`,
				},
				resource.Attribute{
					Name:        "rewrite_domain",
					Description: `fetch domain string, such as http://www.example.com.`,
				},
				resource.Attribute{
					Name:        "fetch_domain",
					Description: `fetch domain string, such as http://www.example.com.`,
				},
				resource.Attribute{
					Name:        "origin_domain_pattern",
					Description: `origin domain wildcards, such as (http(s)://)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pagespeed Child Resource Control can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pagespeed_child_resource_control.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pagespeed Child Resource Control can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pagespeed_child_resource_control.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_pagespeed_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc pagespeed profile.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"pagespeed",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `profile name.`,
				},
				resource.Attribute{
					Name:        "move_css_to_head",
					Description: `move css elements above script tags. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "jpeg_sampling",
					Description: `reduces the color sampling of jpeg images to 4:2:0. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `enable/disable image optimizer. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "resize_image",
					Description: `resizes images when the corresponding img tag specifies a smaller width and height. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "html",
					Description: `enable/disable html optimizer. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "combine_css",
					Description: `combine multiple CSS elements into one. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "max_combine_css_byte",
					Description: `maximum combined css bytes. (1,1048576)`,
				},
				resource.Attribute{
					Name:        "css",
					Description: `enable/disable css optimizer. Valid values: enable/disable. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pagespeed Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pagespeed_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pagespeed Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pagespeed_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_persistence",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance persistence info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"persistence",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `persistence table name.`,
				},
				resource.Attribute{
					Name:        "match_across_servers",
					Description: `Match Across Servers. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "keyvalue_relation",
					Description: `relation to store key value. Valid values: 1:OR, 0:AND .`,
				},
				resource.Attribute{
					Name:        "cookie_secure",
					Description: `Enable or disable Secure cookie attribute. Default disable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "cookie_httponly",
					Description: `Enable or disable HttpOnly cookie attribute. Default disable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "sess_kw_type",
					Description: `keyword type of session cookie name. Valid values: 1:PHPSESSID, 0:auto, 3:CFID+CFTOKEN, 2:JSESSIONID, 5:custom, 4:ASP.NET_SessionId .`,
				},
				resource.Attribute{
					Name:        "avpno",
					Description: `avpno.`,
				},
				resource.Attribute{
					Name:        "ipv6_maskbits",
					Description: `IPv6 netmask bits. (1,128)`,
				},
				resource.Attribute{
					Name:        "cookie_custom_attr",
					Description: `Enable or disable custom cookie attribute. Default disable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `persistence type. Valid values: 11:radius-attribute, 10:embedded-cookie, 13:sip-call-id, 12:ssl-session-id, 15:passive-cookie, 14:rdp-cookie, 17:iso8583-bitmap, 1:source-address, 3:hash-source-address-port, 2:consistent-hash-ip, 5:hash-http-request, 4:hash-http-header, 7:persistent-cookie, 6:hash-cookie, 9:rewrite-cookie, 8:insert-cookie .`,
				},
				resource.Attribute{
					Name:        "iso8583_bitmap_relation",
					Description: `bitmap relation to trigger persistence. Valid values: 1:OR, 0:AND .`,
				},
				resource.Attribute{
					Name:        "ipv4_maskbits",
					Description: `IPv4 netmask bits. (1,32)`,
				},
				resource.Attribute{
					Name:        "keyword",
					Description: `persistence keyword.`,
				},
				resource.Attribute{
					Name:        "cookie_domain",
					Description: `Set the Domain cookie attribute value. Default is empty.`,
				},
				resource.Attribute{
					Name:        "cookie_custom_attr_val",
					Description: `Set the custom cookie attribute value. Default is empty.`,
				},
				resource.Attribute{
					Name:        "override_connection_limit",
					Description: `Override Connection Limit. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "cookie_samesite",
					Description: `Set the SameSite cookie attribute: nothing|None|Lax|Strict. Default is nothing. Valid values: 1:Secure should be set when using samesite none, 0:nothing, 3:strict, 2:lax .`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `persistence entry timeout, unit: second, default 300. (1,86400)`,
				},
				resource.Attribute{
					Name:        "radius_attribute_relation",
					Description: `radius attribute relation. Valid values: 1:OR, 0:AND . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Persistence can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_persistence.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Persistence can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_persistence.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_persistence_child_iso8583_bitmap",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance persistence info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"persistence",
				"child",
				"iso8583",
				"bitmap",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `iso8583 bitmap id.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `iso8583 bitmap type. Valid values: 67:67-ext-payment-code, 68:68-recv-inst-cy-code, 69:69-settle-inst-cy-code, 80:80-inquiries-num, 24:24-function-code, 20:20-PAN-ext-cy-code, 21:21-fwd-inst-cy-code, 22:22-POS-data, 23:23-card-seq-num, 42:42-card-accept-id-code, 40:40-serv-restrict-code, 41:41-card-accept-term-id, 3:3-process-code, 2:2-primary-acct-num, 7:7-date-time-trans, 77:77-debits-reversal-num, 76:76-debits-num, 75:75-credits-reversal-num, 74:74-credits-num, 73:73-date-action, 79:79-transfer-reversal-num, 78:78-tranfer-num, 11:11-sys-trace-audit-num, 12:12-date-tm-loc-trans, 19:19-acq-inst-cy-code, 18:18-merchant-type, 37:37-retrieval-ref-num, 52:52-PIN-data, 33:33-fwd-inst-id-code, 32:32-acq-inst-id-code . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Persistence Child Iso8583 Bitmap can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_persistence_child_iso8583_bitmap.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Persistence Child Iso8583 Bitmap can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_persistence_child_iso8583_bitmap.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_persistence_child_radius_attribute",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance persistence info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"persistence",
				"child",
				"radius",
				"attribute",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `radius attribute id.`,
				},
				resource.Attribute{
					Name:        "vendor_id",
					Description: `vendor id, 0 means will use entire attribute as input of persistence.`,
				},
				resource.Attribute{
					Name:        "vendor_type",
					Description: `vendor type, 0 means will use entire attribute as input of persistence. (0,255)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `radius attribute type. Valid values: 50:50-acct-multi-session-id, 60:60-chap-challenge, 61:61-nas-port-type, 62:62-port-limit, 63:63-login-lat-port, 24:24-state, 26:26-vendor-specific, 48:48-acct-output-packets, 49:49-acct-terminate-cause, 46:46-acct-session-time, 47:47-acct-input-packets, 44:44-acct-session-id, 45:45-acct-authentic, 42:42-acct-input-octets, 43:43-acct-output-octets, 40:40-acct-status-type, 41:41-acct-delay-time, 1:1-user-name, 5:5-nas-port, 4:4-nas-ip-address, 7:7-framed-protocol, 6:6-service-type, 9:9-framed-ip-netmask, 8:8-framed-ip-address, 13:13-framed-compression, 12:12-framed-mtu, 14:14-login-ip-host, 19:19-callback-number, 32:32-nas-identifier, 31:31-calling-station-id, 30:30-called-station-id, 51:51-acct-link-count, 36:36-login-lat-group, 35:35-login-lat-node, 34:34-login-lat-service, 33:33-proxy-state . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Persistence Child Radius Attribute can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_persistence_child_radius_attribute.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Persistence Child Radius Attribute can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_persistence_child_radius_attribute.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_pool",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance pool info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `pool name.`,
				},
				resource.Attribute{
					Name:        "pool_type",
					Description: `address type. Valid values: 1:ipv4, 2:ipv6 .`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `health check control. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "health_check_relationship",
					Description: `health check relationship. Valid values: 1:OR, 0:AND .`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `service.`,
				},
				resource.Attribute{
					Name:        "sdn_connector",
					Description: `sdn connector.`,
				},
				resource.Attribute{
					Name:        "sdn_addr_private",
					Description: `use private node ip. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "direct_route_ip",
					Description: `direct route ip.`,
				},
				resource.Attribute{
					Name:        "direct_route_mode",
					Description: `direct route mode. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "health_check_list",
					Description: `health check list.`,
				},
				resource.Attribute{
					Name:        "direct_route_ip6",
					Description: `direct route ip6.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `pool type. Valid values: 1:static, 2:dynamic .`,
				},
				resource.Attribute{
					Name:        "rs_profile",
					Description: `real server SSL profile name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pool can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pool.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pool can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pool.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_pool_child_pool_member",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance pool info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"pool",
				"child",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `pool member id.`,
				},
				resource.Attribute{
					Name:        "proxy_protocol",
					Description: `pool member support proxy protocol version. Valid values: 1:v1, 0:none, 2:v2 .`,
				},
				resource.Attribute{
					Name:        "hc_status",
					Description: `pool member hc status.`,
				},
				resource.Attribute{
					Name:        "connlimit",
					Description: `connection limit. (0,1048576)`,
				},
				resource.Attribute{
					Name:        "recover",
					Description: `pool member recover time. (0,86400)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `pool member service port. (0,65535)`,
				},
				resource.Attribute{
					Name:        "rs_profile_inherit",
					Description: `real server SSL profile inherit. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "warmrate",
					Description: `pool member warm up rate. (1,86400)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `pool member weight. (1,256)`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `pool member server name.`,
				},
				resource.Attribute{
					Name:        "mssql_read_only",
					Description: `set as read only mssql server. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "m_health_check_relationship",
					Description: `health check relationship. Valid values: 1:OR, 0:AND .`,
				},
				resource.Attribute{
					Name:        "mysql_read_only",
					Description: `set as read only mysql server. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(en|dis)able/maintain pool member. Valid values: 1:enable, 0:disable, 2:maintain .`,
				},
				resource.Attribute{
					Name:        "real_server_id",
					Description: `real server.`,
				},
				resource.Attribute{
					Name:        "health_check_inherit",
					Description: `health check inherit. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "address6",
					Description: `ipv6 address of interface.`,
				},
				resource.Attribute{
					Name:        "mysql_group_id",
					Description: `pool member mysql group id.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `(en|dis)able SSL. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `host content that modify to.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `pool member cookie.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `ip address of interface.`,
				},
				resource.Attribute{
					Name:        "rs_profile",
					Description: `real server SSL profile name.`,
				},
				resource.Attribute{
					Name:        "warmup",
					Description: `pool member warm up time. (0,86400)`,
				},
				resource.Attribute{
					Name:        "m_health_check",
					Description: `health check control. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "connection_rate_limit",
					Description: `pool member connection rate limit(0 - disable), only work for L4VS. (0,86400)`,
				},
				resource.Attribute{
					Name:        "m_health_check_list",
					Description: `health check list.`,
				},
				resource.Attribute{
					Name:        "modify_host",
					Description: `modify the Host HTTP header. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "backup",
					Description: `set as backup server. Valid values: enable/disable. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pool Child Pool Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pool_child_pool_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Pool Child Pool Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_pool_child_pool_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `profile name.`,
				},
				resource.Attribute{
					Name:        "msg_encode_type",
					Description: `encode type of message, default ascii. Valid values: 1:ascii, 2:binary .`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `geography ip allowlist..`,
				},
				resource.Attribute{
					Name:        "ram_caching",
					Description: `caching name.`,
				},
				resource.Attribute{
					Name:        "timeout_tcp_session",
					Description: `timeout tcp session. (1,86400)`,
				},
				resource.Attribute{
					Name:        "dns_authenticate_flag",
					Description: `authenticate client by redirecting UDP query to TCP. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "dns_max_query_length",
					Description: `maximum query length. (256,4096)`,
				},
				resource.Attribute{
					Name:        "decompression",
					Description: `decompression name.`,
				},
				resource.Attribute{
					Name:        "client_protocol",
					Description: `transport protocol used. Valid values: 1:udp, 2:tcp .`,
				},
				resource.Attribute{
					Name:        "intermediate_ca_group",
					Description: `intermediate ca group name.`,
				},
				resource.Attribute{
					Name:        "vendor_id",
					Description: `vendor id.`,
				},
				resource.Attribute{
					Name:        "dns_cache_ageout_time",
					Description: `cache ageout time(seconds). (1,65535)`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `use ssl to support https. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "timeout_ip_session",
					Description: `timeout ip session. (1,86400)`,
				},
				resource.Attribute{
					Name:        "smtp_disable_command",
					Description: `smtp disable command. Valid values: 1:expn, 2:turn, 4:vrfy .`,
				},
				resource.Attribute{
					Name:        "server_keepalive",
					Description: `wait for a new server side request in a limited time. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `ssl ciphers.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `product name.`,
				},
				resource.Attribute{
					Name:        "failed_server_type",
					Description: `action when server side fail. Valid values: 1:drop, 2:send .`,
				},
				resource.Attribute{
					Name:        "forward_client_certificate_header",
					Description: `foward client certificate header.`,
				},
				resource.Attribute{
					Name:        "tune_bufsize",
					Description: `buffer size of a session. (128,1048576)`,
				},
				resource.Attribute{
					Name:        "local_cert",
					Description: `local cert name.`,
				},
				resource.Attribute{
					Name:        "client_address",
					Description: `use client address to connect to pool. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "customized_ssl_ciphers_flag",
					Description: `(en|dis)able SSL user-customized ciphers. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "smtp_disable_command_status",
					Description: `smtp disable command status. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "client_keepalive",
					Description: `wait for a new client side request in a limited time. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "failed_server_str",
					Description: `reply status string when server side fail.`,
				},
				resource.Attribute{
					Name:        "timeout_radius_session",
					Description: `timeout of radius session. (1,3600)`,
				},
				resource.Attribute{
					Name:        "dns_malform_query_action",
					Description: `action on malform query. Valid values: 1:Drop the query., 2:Forward the query. .`,
				},
				resource.Attribute{
					Name:        "client_sni_required",
					Description: `client SNI required. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "http_keepalive_timeout",
					Description: `maximum allowed time to wait for a new HTTP request to appear. (1,86400)`,
				},
				resource.Attribute{
					Name:        "sip_dlg_timeout",
					Description: `maximum allowed timeout for a dialog. (1,300)`,
				},
				resource.Attribute{
					Name:        "client_timeout",
					Description: `maximum inactivity time on the client side. (1,86400)`,
				},
				resource.Attribute{
					Name:        "mysql_mode",
					Description: `mysql mode. Valid values: 0:single-primary, 2:sharding .`,
				},
				resource.Attribute{
					Name:        "insert_client_ip",
					Description: `insert client source IP in SIP header of request. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "opt_header_length",
					Description: `length of optional header before MTI, including the length-indicator, range 0-32. (0,32)`,
				},
				resource.Attribute{
					Name:        "http_request_timeout",
					Description: `maximum allowed time to wait for a complete HTTP request. (1,86400)`,
				},
				resource.Attribute{
					Name:        "http_mode",
					Description: `operating mode for http(s). Valid values: 1:The server-facing connection is closed after the response., 3:All requests and responses are processed and connection stays keep-alived on both sides., 2:Process only the first request and response. .`,
				},
				resource.Attribute{
					Name:        "forward_client_certificate",
					Description: `foward client certificate. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "dns_cache_response_type",
					Description: `cache response type. Valid values: 1:all-records, 2:round-robin .`,
				},
				resource.Attribute{
					Name:        "max_http_headers",
					Description: `Max HTTP headers limit, notice: even enlarge this limit you may also meet parse failed because the buffer size limit.. (10,10000)`,
				},
				resource.Attribute{
					Name:        "origin_realm",
					Description: `origin realm.`,
				},
				resource.Attribute{
					Name:        "geoip_list",
					Description: `geography ip block list..`,
				},
				resource.Attribute{
					Name:        "server_keepalive_timeout",
					Description: `maximum allowed time to wait for a new server side request to appear. (5,300)`,
				},
				resource.Attribute{
					Name:        "client_ssl",
					Description: `client ssl. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "compression",
					Description: `compression name.`,
				},
				resource.Attribute{
					Name:        "response_half_closed_request",
					Description: `If enable ADC will continue serve the request in half closed connection until response complete.. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "dynamic_auth",
					Description: `enable/disable RADIUS dynamic authorization (CoA, Disconnect messages). Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "server_max_size",
					Description: `maximum connections on the server side. (1,30000)`,
				},
				resource.Attribute{
					Name:        "ip_reputation",
					Description: `use IP Reputation. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `protocol type. Valid values: 20:explicit_http, 21:l7-tcp, 22:l7-udp, 1:tcp, 3:http, 2:udp, 5:radius, 4:ftp, 7:https, 6:tcps, 9:sip, 8:turbohttp, 11:ip, 10:rdp, 13:smtp, 12:dns, 15:rtmp, 14:rtsp, 17:diameter, 16:mysql, 19:mssql, 18:iso8583 .`,
				},
				resource.Attribute{
					Name:        "failed_client_str",
					Description: `reply status string when client side fail.`,
				},
				resource.Attribute{
					Name:        "timeout_udp_session",
					Description: `timeout udp session. (1,86400)`,
				},
				resource.Attribute{
					Name:        "local_cert_group",
					Description: `local cert group name.`,
				},
				resource.Attribute{
					Name:        "http2_profile",
					Description: `HTTP/2 profile.`,
				},
				resource.Attribute{
					Name:        "customized_ssl_ciphers",
					Description: `SSL user-customized ciphers (concatenate multiple ciphers by :).`,
				},
				resource.Attribute{
					Name:        "starttls_active_mode",
					Description: `starttls active mode. Valid values: 1:allow, 3:none, 2:require .`,
				},
				resource.Attribute{
					Name:        "server_age",
					Description: `maximum inactivity time on the server side. (1,86400)`,
				},
				resource.Attribute{
					Name:        "smtp_domain_name",
					Description: `domain name.`,
				},
				resource.Attribute{
					Name:        "dns_cache_size",
					Description: `max cache size(M bytes). (1,100)`,
				},
				resource.Attribute{
					Name:        "dns_cache_flag",
					Description: `cache. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "new_ssl_ciphers_long",
					Description: `ssl ciphers including Perfect Forward Secrecy. Valid values: 32768:, 4194304:Perfect Forward Secrecy cipher, 131072:Perfect Forward Secrecy cipher, 262144:Perfect Forward Secrecy cipher, 64:Perfect Forward Secrecy cipher, 65536:, 256:Perfect Forward Secrecy cipher, 33554432:, 8192:Perfect Forward Secrecy cipher, 16777216:, 536870912:Perfect Forward Secrecy cipher, 1048576:Perfect Forward Secrecy cipher, 1:Perfect Forward Secrecy cipher, 2097152:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 16384:, 17179869184:Non Encryption Ciphers, 4096:Perfect Forward Secrecy cipher, 134217728:, 8388608:, 67108864:Perfect Forward Secrecy cipher, 524288:Perfect Forward Secrecy cipher, 128:Perfect Forward Secrecy cipher, 2147483648:, 1024:Perfect Forward Secrecy cipher, 268435456:, 16:Perfect Forward Secrecy cipher, 32:Perfect Forward Secrecy cipher, 2048:Perfect Forward Secrecy cipher, 1073741824:Perfect Forward Secrecy cipher, 512:Perfect Forward Secrecy cipher .`,
				},
				resource.Attribute{
					Name:        "http_x_forwarded_for_header",
					Description: `change X-Forwarded-For header name.`,
				},
				resource.Attribute{
					Name:        "idle_time",
					Description: `idle time. (60,3600)`,
				},
				resource.Attribute{
					Name:        "max_header_size",
					Description: `max header size. (16,65536)`,
				},
				resource.Attribute{
					Name:        "deploy_mode",
					Description: `deploy mode. Valid values: 1:Proxy mode, 2:Direct Server Return mode (DSR mode does not support TCP traffic) .`,
				},
				resource.Attribute{
					Name:        "server_protocol",
					Description: `transport protocol used for server side. Valid values: 1:udp, 2:tcp .`,
				},
				resource.Attribute{
					Name:        "origin_host",
					Description: `origin host.`,
				},
				resource.Attribute{
					Name:        "media_address",
					Description: `media address.`,
				},
				resource.Attribute{
					Name:        "server_close_propagation",
					Description: `server close propagation. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "timeout_tcp_session_after_fin",
					Description: `timeout tcp session_after_FIN. (1,86400)`,
				},
				resource.Attribute{
					Name:        "sip_max_size",
					Description: `maximum allowed size of a message. (1,65535)`,
				},
				resource.Attribute{
					Name:        "ssl_proxy",
					Description: `ssl proxy mode. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "use_tls_tickets",
					Description: `(en|dis)able SSL tls-tickets support. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "dns_cache_entry_size",
					Description: `maximum cache entry size. (256,4096)`,
				},
				resource.Attribute{
					Name:        "queue_timeout",
					Description: `maximum time to wait in the queue for a connection slot to be free. (1,86400)`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `maximum time to wait for a connection attempt to a server to succeed. (1,86400)`,
				},
				resource.Attribute{
					Name:        "opt_trailer_hex",
					Description: `hex string of optional trailer, maximum length 16, i.e. 8 bytes in binary.`,
				},
				resource.Attribute{
					Name:        "length_indicator_shift",
					Description: `bytes to shift from the beginning of payload to read length value, range 0-32. (0,32)`,
				},
				resource.Attribute{
					Name:        "ip_reputation_redirect",
					Description: `redirect url for IP Reputation.`,
				},
				resource.Attribute{
					Name:        "length_indicator_size",
					Description: `total bytes reading to calculate length, range 0-8. (0,8)`,
				},
				resource.Attribute{
					Name:        "http_x_forwarded_for",
					Description: `insert X-Forwarded-For header to request. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "allow_ssl_versions",
					Description: `SSL/TLS versions allowed. Valid values: 8:tlsv1.1, 2:sslv3, 4:tlsv1.0, 16:tlsv1.2 .`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `use source port to connect to pool. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `idle timeout. (1,86400)`,
				},
				resource.Attribute{
					Name:        "security_mode",
					Description: `security mode. Valid values: 1:Use explicit(AUTH SSL/TLS) security mode., 0:No SSL security., 2:Use implicit security mode. .`,
				},
				resource.Attribute{
					Name:        "http_send_timeout",
					Description: `the timeout(in second) of HTTP send out all the buffered data. (0,86400)`,
				},
				resource.Attribute{
					Name:        "dynamic_auth_port",
					Description: `dynamic auth port. (1,65535)`,
				},
				resource.Attribute{
					Name:        "ssl_algorithm",
					Description: `ssl encryption algorithm. Valid values: 1:All algorithms., 3:High algorithms., 2:High and medium algorithms. .`,
				},
				resource.Attribute{
					Name:        "failed_client_type",
					Description: `action when client side fail. Valid values: 1:drop, 2:send .`,
				},
				resource.Attribute{
					Name:        "stateless",
					Description: `UDP stateless. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "server_timeout",
					Description: `maximum inactivity time on the server side. (1,86400)`,
				},
				resource.Attribute{
					Name:        "cert_verify",
					Description: `cert verify name.`,
				},
				resource.Attribute{
					Name:        "geoip_redirect",
					Description: `redirect url for IP geography.`,
				},
				resource.Attribute{
					Name:        "length_indicator_type",
					Description: `encode type of length indicator, default binary. Valid values: 1:binary, 3:decimal-str, 2:BCD, 4:hex-str . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_client_request_header_erase",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"client",
				"request",
				"header",
				"erase",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `header erase id.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `header erase type. Valid values: 1:first, 2:all .`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `field identify which header to be erased. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Client Request Header Erase can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_client_request_header_erase.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Client Request Header Erase can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_client_request_header_erase.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_client_request_header_insert",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"client",
				"request",
				"header",
				"insert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `header insert id.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `header insert type. Valid values: 1:insert-if-not-exist, 3:append-if-not-exist, 2:insert-always, 4:append-always .`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `field:value pair string to be inserted. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Client Request Header Insert can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_client_request_header_insert.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Client Request Header Insert can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_client_request_header_insert.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_client_response_header_erase",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"client",
				"response",
				"header",
				"erase",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `header erase id.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `header erase type. Valid values: 1:first, 2:all .`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `field identify which header to be erased. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Client Response Header Erase can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_client_response_header_erase.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Client Response Header Erase can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_client_response_header_erase.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_client_response_header_insert",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"client",
				"response",
				"header",
				"insert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `header insert id.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `header insert type. Valid values: 1:insert-if-not-exist, 3:append-if-not-exist, 2:insert-always, 4:append-always .`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `field:value pair string to be inserted. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Client Response Header Insert can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_client_response_header_insert.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Client Response Header Insert can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_client_response_header_insert.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_mssql_user_password",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"mssql",
				"user",
				"password",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `mssql user password id.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `password. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Mssql User Password can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_mssql_user_password.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Mssql User Password can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_mssql_user_password.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_mysql_rule",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"mysql",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `mysql rule id.`,
				},
				resource.Attribute{
					Name:        "client_ip_list",
					Description: `rule client ip list.`,
				},
				resource.Attribute{
					Name:        "user_list",
					Description: `rule user list.`,
				},
				resource.Attribute{
					Name:        "db_list",
					Description: `rule database list.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type. Valid values: 1:secondary, 0:primary .`,
				},
				resource.Attribute{
					Name:        "table_list",
					Description: `rule table list.`,
				},
				resource.Attribute{
					Name:        "sql_list",
					Description: `rule sql statement list. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Mysql Rule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_mysql_rule.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Mysql Rule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_mysql_rule.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_mysql_sharding",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"mysql",
				"sharding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `mysql sharding id.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `sharding database.`,
				},
				resource.Attribute{
					Name:        "group_list",
					Description: `sharding group list, for range type, such as: 0:1-10 1:11-20, for hash type, such as: 0 1.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `sharding table column.`,
				},
				resource.Attribute{
					Name:        "table",
					Description: `sharding table.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type. Valid values: 1:hash, 0:range . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Mysql Sharding can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_mysql_sharding.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Mysql Sharding can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_mysql_sharding.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_mysql_user_password",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"mysql",
				"user",
				"password",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `mysql user password id.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `password. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Mysql User Password can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_mysql_user_password.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Mysql User Password can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_mysql_user_password.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_server_request_header_erase",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"server",
				"request",
				"header",
				"erase",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `header erase id.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `header erase type. Valid values: 1:first, 2:all .`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `field identify which header to be erased. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Server Request Header Erase can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_server_request_header_erase.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Server Request Header Erase can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_server_request_header_erase.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_server_request_header_insert",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"server",
				"request",
				"header",
				"insert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `header insert id.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `header insert type. Valid values: 1:insert-if-not-exist, 3:append-if-not-exist, 2:insert-always, 4:append-always .`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `field:value pair string to be inserted. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Server Request Header Insert can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_server_request_header_insert.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Server Request Header Insert can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_server_request_header_insert.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_server_response_header_erase",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"server",
				"response",
				"header",
				"erase",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `header erase id.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `header erase type. Valid values: 1:first, 2:all .`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `field identify which header to be erased. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Server Response Header Erase can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_server_response_header_erase.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Server Response Header Erase can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_server_response_header_erase.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_profile_child_server_response_header_insert",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance profile help info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"profile",
				"child",
				"server",
				"response",
				"header",
				"insert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `header insert id.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `header insert type. Valid values: 1:insert-if-not-exist, 3:append-if-not-exist, 2:insert-always, 4:append-always .`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `field:value pair string to be inserted. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Server Response Header Insert can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_server_response_header_insert.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Profile Child Server Response Header Insert can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_profile_child_server_response_header_insert.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_real_server",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance real server info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"real",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `real server name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(en|dis)able/maintain real server. Valid values: 1:enable, 0:disable, 2:maintain .`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `real server type. Valid values: 1:static, 3:dynamic_manual, 2:dynamic_auto .`,
				},
				resource.Attribute{
					Name:        "sdn_connector",
					Description: `sdn connector.`,
				},
				resource.Attribute{
					Name:        "sdn_addr_private",
					Description: `use private node ip. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name.`,
				},
				resource.Attribute{
					Name:        "address6",
					Description: `ipv6 address of interface.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `instances filter.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `ip address of interface.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `ip or fqdn. Valid values: 1:fqdn, 0:ip . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Real Server can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_real_server.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Real Server can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_real_server.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_real_server_ssl_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance real-server-ssl-profile info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"real",
				"server",
				"ssl",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `real-server-ssl-profile table name.`,
				},
				resource.Attribute{
					Name:        "new_ssl_ciphers_long",
					Description: `ssl ciphers including Perfect Forward Secrecy. Valid values: 32768:, 4194304:Perfect Forward Secrecy cipher, 1099511627776:Perfect Forward Secrecy cipher, 131072:Perfect Forward Secrecy cipher, 262144:Perfect Forward Secrecy cipher, 64:Perfect Forward Secrecy cipher, 65536:, 256:Perfect Forward Secrecy cipher, 33554432:, 8192:Perfect Forward Secrecy cipher, 16777216:, 536870912:Perfect Forward Secrecy cipher, 1048576:Perfect Forward Secrecy cipher, 1:Perfect Forward Secrecy cipher, 2097152:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 16384:, 17179869184:Non Encryption Ciphers, 2199023255552:Perfect Forward Secrecy cipher, 4096:Perfect Forward Secrecy cipher, 134217728:, 137438953472:Perfect Forward Secrecy cipher, 274877906944:Perfect Forward Secrecy cipher, 8388608:, 67108864:Perfect Forward Secrecy cipher, 524288:Perfect Forward Secrecy cipher, 549755813888:Perfect Forward Secrecy cipher, 128:Perfect Forward Secrecy cipher, 2147483648:, 1024:Perfect Forward Secrecy cipher, 268435456:, 34359738368:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher, 32:Perfect Forward Secrecy cipher, 2048:Perfect Forward Secrecy cipher, 1073741824:Perfect Forward Secrecy cipher, 512:Perfect Forward Secrecy cipher, 68719476736:Perfect Forward Secrecy cipher .`,
				},
				resource.Attribute{
					Name:        "local_cert",
					Description: `Local certificate reference.`,
				},
				resource.Attribute{
					Name:        "server_ocsp_stapling",
					Description: `Enable/disable real-server OCSP stapling support. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "customized_ssl_ciphers_flag",
					Description: `(en|dis)able SSL user-customized ciphers. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "secure_renegotiation",
					Description: `set backend secure renegotiation. Valid values: 1:request, 3:require-strict, 2:require .`,
				},
				resource.Attribute{
					Name:        "supported_groups",
					Description: `ssl Supported Groups. Valid values: 1:secp256r1, 0:x25519, 3:secp521r1, 2:x448, 5:ffdhe2048, 4:secp384r1, 7:ffdhe4096, 6:ffdhe3072, 9:ffdhe8192, 8:ffdhe6144 .`,
				},
				resource.Attribute{
					Name:        "renegotiation",
					Description: `(en|dis)able backend renegotiation. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "renegotiation_deny_action",
					Description: `set renegotiation deny action. Valid values: 1:ignore, 2:terminate .`,
				},
				resource.Attribute{
					Name:        "renegotiate_period",
					Description: `period (s|m|h) (by default in seconds) of ADC initiate renegotiation, set to 0 to disable periodical renegotiation.`,
				},
				resource.Attribute{
					Name:        "allow_ssl_versions",
					Description: `SSL/TLS versions allowed. Valid values: 8:tlsv1.1, 2:sslv3, 4:tlsv1.0, 32:tlsv1.3, 16:tlsv1.2 .`,
				},
				resource.Attribute{
					Name:        "sni_forward_flag",
					Description: `(en|dis)able SSL SNI forward. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `(en|dis)able SSL. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "sni",
					Description: `Set SSL SNI value.`,
				},
				resource.Attribute{
					Name:        "cert_verify",
					Description: `server cert verify name.`,
				},
				resource.Attribute{
					Name:        "customized_ssl_ciphers",
					Description: `SSL customized ciphers (concatenate multiple ciphers by :).`,
				},
				resource.Attribute{
					Name:        "rfc7919_comply",
					Description: `(en|dis)able SSL RFC7919 Comply. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "renegotiate_size",
					Description: `size (MB) of application data transmitted over SSL when ADC initiate renegotiation, set to 0 to disable ADC initiate renegotiation by size. (0,2147483647)`,
				},
				resource.Attribute{
					Name:        "ciphers_tlsv13",
					Description: `tlsv1.3 ciphers. Valid values: 1:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher .`,
				},
				resource.Attribute{
					Name:        "tls_ticket_flag",
					Description: `(en|dis)able SSL session reuse using tls-ticket. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "session_reuse_limit",
					Description: `SSL session reuse limit (0 - disable). (0,1048576)`,
				},
				resource.Attribute{
					Name:        "session_reuse_flag",
					Description: `(en|dis)able SSL session reuse. Valid values: enable/disable. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Real Server Ssl Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_real_server_ssl_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Real Server Ssl Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_real_server_ssl_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_schedule_pool",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance schedule pool info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"schedule",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `pool name.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `schedule. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Schedule Pool can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_schedule_pool.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Schedule Pool can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_schedule_pool.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_virtual_server",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc load-balance virtual-server info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"virtual",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `virtual-server name.`,
				},
				resource.Attribute{
					Name:        "scripting_list",
					Description: `virtual server scripting list.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `enable/maintain/disable virtual server. Valid values: 1:enable, 0:disable, 2:maintain .`,
				},
				resource.Attribute{
					Name:        "packet_fwd_method",
					Description: `packet forwarding method. Valid values: 1:direct_routing, 3:FullNAT, 2:NAT, 5:NAT64, 4:NAT46, 6:tunneling .`,
				},
				resource.Attribute{
					Name:        "warmrate",
					Description: `virtual server warm up rate. (1,86400)`,
				},
				resource.Attribute{
					Name:        "http2https",
					Description: `enable/disable redirect HTTP request to HTTPS. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `pool name.`,
				},
				resource.Attribute{
					Name:        "error_page",
					Description: `error-page name.`,
				},
				resource.Attribute{
					Name:        "fortiview",
					Description: `enable/disable fortiview. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "ztna_profile",
					Description: `ZTNA profile name.`,
				},
				resource.Attribute{
					Name:        "auth_policy",
					Description: `authentication policy name.`,
				},
				resource.Attribute{
					Name:        "adfs_published_service",
					Description: `AD FS published service.`,
				},
				resource.Attribute{
					Name:        "content_rewriting",
					Description: `content rewriting. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "http2https_port",
					Description: `HTTP service port list for redirecting HTTP to HTTPS.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `virtual server domainname on one click gslb server.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `virtual server protocol numbers.`,
				},
				resource.Attribute{
					Name:        "error_msg",
					Description: `error message.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `virtual server service port.`,
				},
				resource.Attribute{
					Name:        "scripting_flag",
					Description: `enable/disable virtual server scripting. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `virtual server comments.`,
				},
				resource.Attribute{
					Name:        "content_routing_list",
					Description: `content routing list.`,
				},
				resource.Attribute{
					Name:        "wccp",
					Description: `enable/disable redirect HTTP/HTTPS request to WCCP client. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "public_ip_type",
					Description: `address type. Valid values: 1:ipv4, 2:ipv6 .`,
				},
				resource.Attribute{
					Name:        "one_click_gslb_server",
					Description: `enable/disable setting one click gslb server. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "pagespeed",
					Description: `virtual server pagespeed.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `connection-limit. (0,100000000)`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `virtual server hostname on one click gslb server.`,
				},
				resource.Attribute{
					Name:        "ssl_mirror_intf",
					Description: `interface list to mirror.`,
				},
				resource.Attribute{
					Name:        "source_pool_list",
					Description: `ip pool name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `virtual-server service type. Valid values: 1:l4-load-balance, 3:l2-load-balance, 2:l7-load-balance .`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `method name.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `persistence name.`,
				},
				resource.Attribute{
					Name:        "dos_profile",
					Description: `DoS profile name.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `profile name.`,
				},
				resource.Attribute{
					Name:        "content_rewriting_list",
					Description: `content rewriting list.`,
				},
				resource.Attribute{
					Name:        "l2_exception_list",
					Description: `layer2 exception list.`,
				},
				resource.Attribute{
					Name:        "ips_profile",
					Description: `IPS profile name.`,
				},
				resource.Attribute{
					Name:        "traffic_group",
					Description: `traffic group name.`,
				},
				resource.Attribute{
					Name:        "stream_scripting_list",
					Description: `virtual server scripting list.`,
				},
				resource.Attribute{
					Name:        "address6",
					Description: `ipv6 address of virtual server.`,
				},
				resource.Attribute{
					Name:        "clone_pool",
					Description: `clone pool name.`,
				},
				resource.Attribute{
					Name:        "use_azure_lb_backend_ip",
					Description: `use azure lb backend ip. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `public ip address of virtual server.`,
				},
				resource.Attribute{
					Name:        "clone_traffic_type",
					Description: `the traffic type to be cloned. Valid values: 1:client-side, 0:both-sides, 2:server-side .`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `ip address of virtual server.`,
				},
				resource.Attribute{
					Name:        "alone",
					Description: `enable/disable alone mode. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "stream_scripting_flag",
					Description: `enable/disable virtual server stream scripting. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "client_ssl_profile",
					Description: `client SSL profile (needed for https tcps smtp ftp).`,
				},
				resource.Attribute{
					Name:        "schedule_pool_list",
					Description: `schedule pool name.`,
				},
				resource.Attribute{
					Name:        "ssl_mirror",
					Description: `enable/disable SSL mirror. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "schedule_list",
					Description: `enable/disable schedule list. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "warmup",
					Description: `virtual server warm up time. (0,86400)`,
				},
				resource.Attribute{
					Name:        "addr_type",
					Description: `address type. Valid values: 1:ipv4, 2:ipv6 .`,
				},
				resource.Attribute{
					Name:        "connection_pool",
					Description: `connection pool name.`,
				},
				resource.Attribute{
					Name:        "azure_lb_backend",
					Description: `azure lb backend name.`,
				},
				resource.Attribute{
					Name:        "connection_rate_limit",
					Description: `virtual server connection rate limit(0 - disable). (0,86400)`,
				},
				resource.Attribute{
					Name:        "trans_rate_limit",
					Description: `virtual server transactions rate limit. (0,1048567)`,
				},
				resource.Attribute{
					Name:        "traffic_log",
					Description: `enable/disable traffic log. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `antivirus profile name.`,
				},
				resource.Attribute{
					Name:        "public_ip6",
					Description: `public ip address of virtual server.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `interface name.`,
				},
				resource.Attribute{
					Name:        "captcha_profile",
					Description: `Captcha profile.`,
				},
				resource.Attribute{
					Name:        "waf_profile",
					Description: `web application firewall profile name.`,
				},
				resource.Attribute{
					Name:        "content_routing",
					Description: `content routing. Valid values: enable/disable. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Virtual Server can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_virtual_server.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Virtual Server can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_virtual_server.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_web_category",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Web Category.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"web",
				"category",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Web Category Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Web Category Description.`,
				},
				resource.Attribute{
					Name:        "fadc_id",
					Description: `Web Category Id. (0,32) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Web Category can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_web_category.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Web Category can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_web_category.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_web_filter_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Web Filter Profile.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"web",
				"filter",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Web Filter Profile Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Web Filter Profile Description. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Web Filter Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_web_filter_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Web Filter Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_web_filter_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_web_filter_profile_child_category_members",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Web Filter Profile.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"web",
				"filter",
				"profile",
				"child",
				"category",
				"members",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Web Filter Profile Category Member Id.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Web Categroy / Web Sub Category. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Web Filter Profile Child Category Members can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_web_filter_profile_child_category_members.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Web Filter Profile Child Category Members can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_web_filter_profile_child_category_members.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_load_balance_web_sub_category",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Web Sub Category.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"load",
				"balance",
				"web",
				"sub",
				"category",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Web Sub Category Name.`,
				},
				resource.Attribute{
					Name:        "fadc_id",
					Description: `Web Sub Category Id. (0,255)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Web Category Description. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Web Sub Category can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_web_sub_category.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Load Balance Web Sub Category can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_load_balance_web_sub_category.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_access_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure ipv4 access list..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"access",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name of ipv4 access list..`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Comments.. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Access List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_access_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Access List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_access_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_access_list6",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure ipv6 access list..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"access",
				"list6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name of ipv6 access list..`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Comments.. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Access List6 can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_access_list6.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Access List6 can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_access_list6.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_access_list6_child_rule",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure ipv6 access list..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"access",
				"list6",
				"child",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `id..`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action.. Valid values: 1:permit, 2:deny .`,
				},
				resource.Attribute{
					Name:        "prefix6",
					Description: `Prefix6 to define regular filter criteria.. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Access List6 Child Rule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_access_list6_child_rule.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Access List6 Child Rule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_access_list6_child_rule.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_access_list_child_rule",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure ipv4 access list..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"access",
				"list",
				"child",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `id..`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action.. Valid values: 1:permit, 2:deny .`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Prefix to define regular filter criteria.. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Access List Child Rule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_access_list_child_rule.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Access List Child Rule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_access_list_child_rule.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_bgp",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure BGP..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"bgp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "redistribute_ospf",
					Description: `Enable/Disable redistribute ospf. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `Router ID..`,
				},
				resource.Attribute{
					Name:        "redistribute_connected",
					Description: `Enable/Disable redistribute connected. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "deterministic_med",
					Description: `Enable/disable enforce deterministic comparison of MED.. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "redistribute_connected6",
					Description: `Enable/Disable redistribute connected6. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "as",
					Description: `Router AS number..`,
				},
				resource.Attribute{
					Name:        "redistribute_static6",
					Description: `Enable/Disable redistribute static6.. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "always_compare_med",
					Description: `Enable/disable always compare MED.. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "bestpath_cmp_routerid",
					Description: `Enable/disable compare router ID for identical EBGP paths.. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "redistribute_static",
					Description: `Enable/Disable redistribute static.. Valid values: enable/disable. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Bgp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_bgp.labelname RouterBgp ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Bgp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_bgp.labelname RouterBgp ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_bgp_child_ha_router_id_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure BGP..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"bgp",
				"child",
				"ha",
				"id",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Router-id list entry ID..`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `Node ID of HA Node. (0,7)`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `Router-id.. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Bgp Child Ha Router Id List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_bgp_child_ha_router_id_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Bgp Child Ha Router Id List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_bgp_child_ha_router_id_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_bgp_child_neighbor",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure BGP..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"bgp",
				"child",
				"neighbor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Neighbor entry ID..`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of neighbor..`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Default weight for routes from this neighbor.. (0,65535)`,
				},
				resource.Attribute{
					Name:        "distribute_list_out6",
					Description: `Filter for IP updates to this neighbor..`,
				},
				resource.Attribute{
					Name:        "distribute_list_out",
					Description: `Filter for IP updates to this neighbor..`,
				},
				resource.Attribute{
					Name:        "prefix_list_in6",
					Description: `IP Inbound filter for updates from this neighbor..`,
				},
				resource.Attribute{
					Name:        "prefix_list_out",
					Description: `IP Outbound filter for updates to this neighbor..`,
				},
				resource.Attribute{
					Name:        "prefix_list_out6",
					Description: `IP Outbound filter for updates to this neighbor..`,
				},
				resource.Attribute{
					Name:        "prefix_list_in",
					Description: `IP Inbound filter for updates from this neighbor..`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `address type. Valid values: 1:ipv4, 2:ipv6 .`,
				},
				resource.Attribute{
					Name:        "distribute_list_in",
					Description: `Filter for IP updates from this neighbor..`,
				},
				resource.Attribute{
					Name:        "distribute_list_in6",
					Description: `Filter for IP updates from this neighbor..`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface..`,
				},
				resource.Attribute{
					Name:        "ip6",
					Description: `IPv6 address of neighbor..`,
				},
				resource.Attribute{
					Name:        "remote_as",
					Description: `AS number of neighbor..`,
				},
				resource.Attribute{
					Name:        "holdtime",
					Description: `Number of seconds to mark peer as dead.. (0,65535)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port.. (0,65535)`,
				},
				resource.Attribute{
					Name:        "keepalive",
					Description: `Frequency to send keep alive requests.. (0,65535) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Bgp Child Neighbor can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_bgp_child_neighbor.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Bgp Child Neighbor can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_bgp_child_neighbor.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_bgp_child_network",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure BGP..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"bgp",
				"child",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Network entry ID..`,
				},
				resource.Attribute{
					Name:        "prefix6",
					Description: `Network6 prefix..`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Network prefix..`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `address type. Valid values: 1:ipv4, 2:ipv6 . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Bgp Child Network can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_bgp_child_network.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Bgp Child Network can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_bgp_child_network.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_isp",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc route isp.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"isp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `The number of route isp.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `The name of ISP address group.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway for this route. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Isp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_isp.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Isp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_isp.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_md5_ospf",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure md5 authentication list for ospf..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"md5",
				"ospf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name of md5 list.. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Md5 Ospf can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_md5_ospf.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Md5 Ospf can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_md5_ospf.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_md5_ospf_child_md5_member",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure md5 authentication list for ospf..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"md5",
				"ospf",
				"child",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `id..`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `md5 key. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Md5 Ospf Child Md5 Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_md5_ospf_child_md5_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Md5 Ospf Child Md5 Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_md5_ospf_child_md5_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_ospf",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure OSPF..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"ospf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "redistribute_connected",
					Description: `Enable/Disable redistribute connected. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "default_information_metric",
					Description: `Default information metric.. (-1,16777214)`,
				},
				resource.Attribute{
					Name:        "redistribute_static_metric_type",
					Description: `Redistribute static metric type setting.. Valid values: 1:1, 2:2 .`,
				},
				resource.Attribute{
					Name:        "default_information_metric_type",
					Description: `Default information metric type.. Valid values: 1:1, 2:2 .`,
				},
				resource.Attribute{
					Name:        "redistribute_connected_metric",
					Description: `Redistribute connected metric setting.. (-1,16777214)`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `router-id must be set..`,
				},
				resource.Attribute{
					Name:        "default_information_originate",
					Description: `Area type setting.. Valid values: 1:enable, 0:disable, 2:always .`,
				},
				resource.Attribute{
					Name:        "redistribute_static_metric",
					Description: `Redistribute static metric setting.. (-1,16777214)`,
				},
				resource.Attribute{
					Name:        "redistribute_static",
					Description: `Enable/Disable redistribute static.. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "default_metric",
					Description: `Default metric of redistribute routes.. (0,16777214)`,
				},
				resource.Attribute{
					Name:        "redistribute_connected_metric_type",
					Description: `Redistribute connected metric type setting.. Valid values: 1:1, 2:2 .`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Distance of the route.. (1,255) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Ospf can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_ospf.labelname RouterOspf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Ospf can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_ospf.labelname RouterOspf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_ospf_child_area",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure OSPF..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"ospf",
				"child",
				"area",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Area entry IP address..`,
				},
				resource.Attribute{
					Name:        "stub_type",
					Description: `Stub type setting.. Valid values: 1:no-summary, 0:summary .`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Authentication type.. Valid values: 1:text, 0:none, 2:md5 .`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Area type setting.. Valid values: 1:stub, 0:regular . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Ospf Child Area can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_ospf_child_area.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Ospf Child Area can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_ospf_child_area.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_ospf_child_ha_router_id_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure OSPF..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"ospf",
				"child",
				"ha",
				"id",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Number of Router-ids in list.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `Node ID of HA Node. (0,7)`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `router-id must be set.. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Ospf Child Ha Router Id List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_ospf_child_ha_router_id_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Ospf Child Ha Router Id List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_ospf_child_ha_router_id_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_ospf_child_network",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure OSPF..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"ospf",
				"child",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Network entry ID..`,
				},
				resource.Attribute{
					Name:        "area_id",
					Description: `Attach the network to area.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `prefix. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Ospf Child Network can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_ospf_child_network.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Ospf Child Network can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_ospf_child_network.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_ospf_child_ospf_interface",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure OSPF..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"ospf",
				"child",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Interface entry name..`,
				},
				resource.Attribute{
					Name:        "transmit_delay",
					Description: `Transmit delay.. (1,65535)`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `Retransmit interval.. (3,65535)`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `Hello interval.. (1,65535)`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Authentication text key..`,
				},
				resource.Attribute{
					Name:        "mtu_ignore",
					Description: `Enable/disable ignore MTU.. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority.. (0,255)`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Authentication type.. Valid values: 1:text, 0:none, 2:md5 .`,
				},
				resource.Attribute{
					Name:        "dead_interval",
					Description: `Dead interval.. (1,65535)`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Configuration interface name..`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `Authentication md5 key list..`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Network type.. Valid values: 1:point-to-point, 0:broadcast, 2:point-to-multipoint .`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `Cost of the interface.. (0,65535) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Ospf Child Ospf Interface can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_ospf_child_ospf_interface.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Ospf Child Ospf Interface can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_ospf_child_ospf_interface.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_policy",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc route policy.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `The number of route policy.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `Destination for this route.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Source for this route policy.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway for this route. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Policy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_policy.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Policy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_policy.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_prefix_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure ipv4 prefix list..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"prefix",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name of ipv4 prefix list..`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Comments.. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Prefix List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_prefix_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Prefix List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_prefix_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_prefix_list6",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure ipv6 prefix list..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"prefix",
				"list6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name of ipv6 prefix list..`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Comments.. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Prefix List6 can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_prefix_list6.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Prefix List6 can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_prefix_list6.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_prefix_list6_child_rule",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure ipv6 prefix list..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"prefix",
				"list6",
				"child",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `id..`,
				},
				resource.Attribute{
					Name:        "le",
					Description: `Maximum prefix length to be matched..`,
				},
				resource.Attribute{
					Name:        "prefix6",
					Description: `Prefix6 to define regular filter criteria..`,
				},
				resource.Attribute{
					Name:        "ge",
					Description: `Minimum prefix length to be matched..`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action.. Valid values: 1:permit, 2:deny . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Prefix List6 Child Rule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_prefix_list6_child_rule.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Prefix List6 Child Rule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_prefix_list6_child_rule.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_prefix_list_child_rule",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Configure ipv4 prefix list..`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"prefix",
				"list",
				"child",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `id..`,
				},
				resource.Attribute{
					Name:        "le",
					Description: `Maximum prefix length to be matched..`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Prefix to define regular filter criteria..`,
				},
				resource.Attribute{
					Name:        "ge",
					Description: `Minimum prefix length to be matched..`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action.. Valid values: 1:permit, 2:deny . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Prefix List Child Rule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_prefix_list_child_rule.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Prefix List Child Rule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_prefix_list_child_rule.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_router_static",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc route configuration.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"router",
				"static",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `The number of the route.`,
				},
				resource.Attribute{
					Name:        "gw",
					Description: `Gateway for this route.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `distance for this route. (1,255)`,
				},
				resource.Attribute{
					Name:        "dest",
					Description: `Destination for this route. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Static can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_static.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Router Static can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_router_static.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_security_antivirus_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc security antivirus profile info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"security",
				"antivirus",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `antivirus profile name.`,
				},
				resource.Attribute{
					Name:        "uncomp_nest_limit",
					Description: `maximum number of allowed levels of compression nesting that attempt to decompress, range is 2 - 100, default 2. (2,100)`,
				},
				resource.Attribute{
					Name:        "streaming_content_bypass",
					Description: `allow streaming content to be bypassed rather than buffered, default enable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "emulator",
					Description: `enable/disable Win32 emulator, default disable to improve throughput. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "scan_bzip2",
					Description: `enable to scan archives using BZIP2 algorithm, default disable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "av_block_log",
					Description: `enable/disable logging files that are blocked by antivirus, default enable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `antivirus profile comments.`,
				},
				resource.Attribute{
					Name:        "analytics_db",
					Description: `use signature database from FortiSandbox to supplement the AV signature databases, default disable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "oversize_limit",
					Description: `maximum in-memory file size that will be scanned in kilobytes (KB), range is 1 - 12000000 KB, default 1024 KB. (1,12000000)`,
				},
				resource.Attribute{
					Name:        "av_virus_log",
					Description: `enable/disable logging for antivirus scanning, default enable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "analytics_max_upload",
					Description: `maximum upload size to FortiSandbox in kilobytes (KB), range is 1 - 2048 KB, default 1024 KB. (1,2048)`,
				},
				resource.Attribute{
					Name:        "uncomp_size_limit",
					Description: `maximum size in megabytes (MB) of memory buffer that use to temporarily decompress, range is 1 - 2000 MB, default 2 MB. (1,2000)`,
				},
				resource.Attribute{
					Name:        "oversize",
					Description: `bypass means bypass oversize files, log means logging oversize files, block means block oversize files, default bypass. Valid values: 1:log, 0:bypass, 2:block .`,
				},
				resource.Attribute{
					Name:        "ftgd_analytics",
					Description: `control which files are uploaded to FortiSandbox, would affect the destination of quarantine, default disable. Valid values: 1:suspicious, 0:disable, 2:all .`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `avmonitor means scanning and logging, quarantine means scanning, logging and quarantine infected files, default avmonitor. Valid values: 1:quarantine, 0:avmonitor . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Security Antivirus Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_security_antivirus_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Security Antivirus Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_security_antivirus_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_security_dos_dos_protection_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc DoS protect profile.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"security",
				"dos",
				"protection",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "tcp_slow_data_limit",
					Description: `tcp slowdata limit.`,
				},
				resource.Attribute{
					Name:        "http_conn_limit",
					Description: `HTTP connection limit.`,
				},
				resource.Attribute{
					Name:        "tcp_conn_limit",
					Description: `tcp connection access limit.`,
				},
				resource.Attribute{
					Name:        "http_req_limit",
					Description: `HTTP request limit.`,
				},
				resource.Attribute{
					Name:        "http_access_limit",
					Description: `HTTP access limit. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Security Dos Dos Protection Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_security_dos_dos_protection_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Security Dos Dos Protection Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_security_dos_dos_protection_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_security_ips_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc security IPS profile.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"security",
				"ips",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "ips_profile_name",
					Description: `IPS profile name.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `ips profile comments. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Security Ips Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_security_ips_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Security Ips Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_security_ips_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_security_waf_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc web application firewall profile info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"security",
				"waf",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `profile name.`,
				},
				resource.Attribute{
					Name:        "data_leak_prevention_name",
					Description: `Data leak prevention.`,
				},
				resource.Attribute{
					Name:        "cors_protection",
					Description: `CORS protection.`,
				},
				resource.Attribute{
					Name:        "http_protocol_constraint",
					Description: `HTTP protocol constraint.`,
				},
				resource.Attribute{
					Name:        "http_header_security_name",
					Description: `HTTP Header Security.`,
				},
				resource.Attribute{
					Name:        "threshold_based_detection",
					Description: `Threshold Based Detection.`,
				},
				resource.Attribute{
					Name:        "biometrics_based_detection",
					Description: `Biometrics Based Detection.`,
				},
				resource.Attribute{
					Name:        "exception_name",
					Description: `Exceptions.`,
				},
				resource.Attribute{
					Name:        "xml_validation_name",
					Description: `XML validation.`,
				},
				resource.Attribute{
					Name:        "csrf_protection",
					Description: `CSRF protection.`,
				},
				resource.Attribute{
					Name:        "cookie_security",
					Description: `Cookie Security.`,
				},
				resource.Attribute{
					Name:        "brute_force_login_name",
					Description: `Brute force login.`,
				},
				resource.Attribute{
					Name:        "credential_stuffing_defense",
					Description: `Credential stuffing defense.`,
				},
				resource.Attribute{
					Name:        "heuristic_sql_xss_injection_detection",
					Description: `Heuristic sql injection and xss detect.`,
				},
				resource.Attribute{
					Name:        "web_attack_signature",
					Description: `Web attack signature.`,
				},
				resource.Attribute{
					Name:        "rule_match_record",
					Description: `record matched rules content enable/disable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "openapi_validation_name",
					Description: `OPENAPI validation.`,
				},
				resource.Attribute{
					Name:        "url_protection",
					Description: `URL protection.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `waf profile description.`,
				},
				resource.Attribute{
					Name:        "api_gateway_policy_name",
					Description: `API Gateway.`,
				},
				resource.Attribute{
					Name:        "bot_detection_name",
					Description: `Bot detection.`,
				},
				resource.Attribute{
					Name:        "json_validation_name",
					Description: `JSON validation.`,
				},
				resource.Attribute{
					Name:        "advanced_protection_name",
					Description: `Advanced protection profile.`,
				},
				resource.Attribute{
					Name:        "input_validation_policy_name",
					Description: `input validation policy. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Security Waf Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_security_waf_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Security Waf Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_security_waf_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_security_ztna_profile",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc security ZTNA profile.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"security",
				"ztna",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `profile name.`,
				},
				resource.Attribute{
					Name:        "log_status",
					Description: `log status. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `comments. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Security Ztna Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_security_ztna_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import Security Ztna Profile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_security_ztna_profile.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_address",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc system address.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "ip_max",
					Description: `IP max.`,
				},
				resource.Attribute{
					Name:        "ip_netmask",
					Description: `IP/netmask.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type. Valid values: 1:ip-netmask, 2:ip-range .`,
				},
				resource.Attribute{
					Name:        "ip_min",
					Description: `IP min. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Address can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_address.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Address can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_address.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_address6",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc system IPv6 address.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"address6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "ip6_network",
					Description: `IPv6/Prefix.`,
				},
				resource.Attribute{
					Name:        "ip6_max",
					Description: `IP6 max.`,
				},
				resource.Attribute{
					Name:        "ip6_min",
					Description: `IPv6 min.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type. Valid values: 1:ip6-network, 2:ip6-range . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Address6 can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_address6.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Address6 can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_address6.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_addrgrp",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc address group info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"addrgrp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `address group name.`,
				},
				resource.Attribute{
					Name:        "member_list",
					Description: `member list. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Addrgrp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_addrgrp.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Addrgrp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_addrgrp.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_addrgrp6",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc address6 group info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"addrgrp6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `address6 group name.`,
				},
				resource.Attribute{
					Name:        "member_list",
					Description: `member list. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Addrgrp6 can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_addrgrp6.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Addrgrp6 can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_addrgrp6.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_auto_backup",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Auto configuration backup.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"auto",
				"backup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `username for local server.`,
				},
				resource.Attribute{
					Name:        "overwrite_config",
					Description: `overwrite previous configuration when backup space is full. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "config_password",
					Description: `password for config.`,
				},
				resource.Attribute{
					Name:        "scheduled_backup_status",
					Description: `enable/disable auto backup. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "scheduled_backup_frequency",
					Description: `scheduled backup frequency. Valid values: 1:every day, 0:time interval, 2:every week .`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `save configuration to ADC disk or local server. Valid values: 1:sftp, 0:disk .`,
				},
				resource.Attribute{
					Name:        "scheduled_backup_time",
					Description: `hour and minute, hh: 0-23, mm: {00|15|30|45}.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `local server ip address.`,
				},
				resource.Attribute{
					Name:        "scheduled_backup_day",
					Description: `day of the week: Sunday to Saturday. Valid values: 1:Monday, 0:Sunday, 3:Wednesday, 2:Tuesday, 5:Friday, 4:Thursday, 6:Saturday .`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `local directory.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `password for local server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `local server port. (1,65535) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Auto Backup can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_auto_backup.labelname SystemAutoBackup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Auto Backup can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_auto_backup.labelname SystemAutoBackup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_azure_lb_backend_ip",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc azure lb backend ip list.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"azure",
				"lb",
				"backend",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `backend ip. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Azure Lb Backend Ip can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_azure_lb_backend_ip.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Azure Lb Backend Ip can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_azure_lb_backend_ip.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_certificate_ca_group",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc CA certificate group.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"certificate",
				"ca",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `ca certificate group name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Ca Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_ca_group.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Ca Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_ca_group.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_certificate_ca_group_child_group_member",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc CA certificate group.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"certificate",
				"ca",
				"group",
				"child",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `CA certificate group member id.`,
				},
				resource.Attribute{
					Name:        "ca",
					Description: `CA certificate reference. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Ca Group Child Group Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_ca_group_child_group_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Ca Group Child Group Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_ca_group_child_group_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_certificate_certificate_verify",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Certificate Verify.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"certificate",
				"verify",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `certificate verify name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Certificate Verify can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_certificate_verify.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Certificate Verify can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_certificate_verify.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_certificate_crl",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc certificate revokation list.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"certificate",
				"crl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `certificate name.`,
				},
				resource.Attribute{
					Name:        "ldap_server",
					Description: `LDAP server.`,
				},
				resource.Attribute{
					Name:        "crldp_status",
					Description: `Enable/disable certificate revocation list distribution point. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "use_ca_id",
					Description: `Enable/Disable Message of CRL in SCEP. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "update_ahead_time",
					Description: `CRL update start ahead time before nextupd (by default in seconds, at least 1h).`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `CRL update interval in case of no change or failure (by default in seconds, at least 5m).`,
				},
				resource.Attribute{
					Name:        "ca_id",
					Description: `Message type of CRL update via SCEP. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Crl can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_crl.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Crl can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_crl.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_certificate_intermediate_ca",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Intermediate CA certificate.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"certificate",
				"intermediate",
				"ca",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `certificate name.`,
				},
				resource.Attribute{
					Name:        "scep_url",
					Description: `scep url.`,
				},
				resource.Attribute{
					Name:        "ca_id",
					Description: `Specify the CA Identifier.`,
				},
				resource.Attribute{
					Name:        "host_header",
					Description: `host geader. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Intermediate Ca can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_intermediate_ca.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Intermediate Ca can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_intermediate_ca.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_certificate_intermediate_ca_group",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Intermediate CA certificate group.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"certificate",
				"intermediate",
				"ca",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Intermediate CA certificate group name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Intermediate Ca Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_intermediate_ca_group.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Intermediate Ca Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_intermediate_ca_group.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_certificate_intermediate_ca_group_child_group_member",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Intermediate CA certificate group.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"certificate",
				"intermediate",
				"ca",
				"group",
				"child",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Intermediate CA certificate group member id.`,
				},
				resource.Attribute{
					Name:        "intmed_ca",
					Description: `Intermediate CA certificate reference. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Intermediate Ca Group Child Group Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_intermediate_ca_group_child_group_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Intermediate Ca Group Child Group Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_intermediate_ca_group_child_group_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_certificate_local",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc local certificate configuration.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"certificate",
				"local",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `certificate name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Local can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_local.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Local can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_local.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_certificate_local_cert_group",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Local certificate group.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"certificate",
				"local",
				"cert",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Local certificate group name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Local Cert Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_local_cert_group.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Local Cert Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_local_cert_group.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_certificate_local_cert_group_child_group_member",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Local certificate group.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"certificate",
				"local",
				"cert",
				"group",
				"child",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `Local certificate group member id.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `As default certificate. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "local_cert",
					Description: `Local certificate reference.`,
				},
				resource.Attribute{
					Name:        "ocsp_stapling",
					Description: `OCSP stapling reference.`,
				},
				resource.Attribute{
					Name:        "extra_intermediate_cag",
					Description: `Intermediate CA group reference.`,
				},
				resource.Attribute{
					Name:        "extra_local_cert",
					Description: `Local certificate reference.`,
				},
				resource.Attribute{
					Name:        "extra_ocsp_stapling",
					Description: `OCSP stapling reference.`,
				},
				resource.Attribute{
					Name:        "intermediate_cag",
					Description: `Intermediate CA group reference. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Local Cert Group Child Group Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_local_cert_group_child_group_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Local Cert Group Child Group Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_local_cert_group_child_group_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_certificate_ocsp",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc ocsp setup.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"certificate",
				"ocsp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `ocsp name.`,
				},
				resource.Attribute{
					Name:        "reject_ocsp_response_with_missing_nextupdate",
					Description: `reject OCSP response with missing nextupdate. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "max_age",
					Description: `OCSP response thisupd max-age (in second, set to -1 to disable max-age check). (-1,2147483647)`,
				},
				resource.Attribute{
					Name:        "nonce_check",
					Description: `enable/disable OCSP response nonce check. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "tunnel_username",
					Description: `username for web proxy authentication.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of OCSP server.`,
				},
				resource.Attribute{
					Name:        "tunnel_status",
					Description: `enable/disable web proxy tunneling for OCSP. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "verify_others",
					Description: `verify OCSP response signing certificates using trusted certificates or CA chain. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "accept_trusted_root_ca",
					Description: `accept trusted root CA or not. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "caching_extra_max_age_check",
					Description: `OCSP caching thisupd extra max-age check (in second, set to -1 to disable it). (-1,2147483647)`,
				},
				resource.Attribute{
					Name:        "tunnel_password",
					Description: `password for web proxy authentication.`,
				},
				resource.Attribute{
					Name:        "caching_expire_ahead_time",
					Description: `OCSP caching nextupd ahead time (in second, set to -1 to disable it). (-1,2147483647)`,
				},
				resource.Attribute{
					Name:        "tunnel_port",
					Description: `web proxy port. (1,65535)`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `OCSP inquery timeout (in millisecond). (1,2147483647)`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `enable/disable OCSP result caching. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `CA group reference for verify OCSP signing certificate.`,
				},
				resource.Attribute{
					Name:        "tunnel_ip",
					Description: `web proxy ip address.`,
				},
				resource.Attribute{
					Name:        "criteria_check",
					Description: `enable/disable OCSP signing certificate verify against the OCSP issuer criteria. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "remote_certificates",
					Description: `remote certificates reference. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Ocsp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_ocsp.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Certificate Ocsp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_certificate_ocsp.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_dns",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc Domain Name Server.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"dns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "ip_second",
					Description: `Secondary DNS.`,
				},
				resource.Attribute{
					Name:        "ip_primary",
					Description: `Primary DNS. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Dns can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_dns.labelname SystemDns ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Dns can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_dns.labelname SystemDns ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_global",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc system global configuration.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"global",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "vdom_admin",
					Description: `enable/disable . Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "telnet_port",
					Description: `the port number of the telnet service. (1,65535)`,
				},
				resource.Attribute{
					Name:        "pre_login_banner",
					Description: `enable/disable pre-login-banner. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "https_server_cert",
					Description: `appliance's local default certificate.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `appliance's host name.`,
				},
				resource.Attribute{
					Name:        "config_sync_enable",
					Description: `config sync enable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "vdom_mode",
					Description: `set vdom mode. Valid values: 1:share-network, 0:independent-network .`,
				},
				resource.Attribute{
					Name:        "sys_global_language",
					Description: `Global GUI display language. Valid values: 1:english, 3:japanese, 2:chinese-simplified, 5:spanish, 7:portuguese, 6:chinese-traditional .`,
				},
				resource.Attribute{
					Name:        "gui_log",
					Description: `Enable/disable menu log on Web GUI. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "ssh_port",
					Description: `the port number of the ssh service. (1,65535)`,
				},
				resource.Attribute{
					Name:        "config_sync_port",
					Description: `config sync port. (1,65535)`,
				},
				resource.Attribute{
					Name:        "gui_device_latitude",
					Description: `latitude between (-90, 90).`,
				},
				resource.Attribute{
					Name:        "gui_router",
					Description: `Enable/disable menu router on Web GUI. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "default_intermediate_ca_group",
					Description: `appliance's default ssl certificate chain.`,
				},
				resource.Attribute{
					Name:        "admin_idle_timeout",
					Description: `the idle time-out for system administration. (1,480)`,
				},
				resource.Attribute{
					Name:        "https_port",
					Description: `the port number of the https service. (1,65535)`,
				},
				resource.Attribute{
					Name:        "https_redirect",
					Description: `HTTPS redirect enable/disable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "gui_system",
					Description: `Enable/disable menu system on Web GUI. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "http_port",
					Description: `the port number of the http service. (1,65535)`,
				},
				resource.Attribute{
					Name:        "use_default_hostname",
					Description: `use-default-hostname enable/disable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "gui_device_longtitude",
					Description: `longitude between (-180, 180). ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Global can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_global.labelname SystemGlobal ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Global can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_global.labelname SystemGlobal ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_ha",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc HA parameters configuration.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"ha",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mgmt_interface",
					Description: `Management interface.`,
				},
				resource.Attribute{
					Name:        "local_node_id",
					Description: `local node id (0-7). (0,7)`,
				},
				resource.Attribute{
					Name:        "peer_address",
					Description: `Fort unicast peer ip address.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `group name.`,
				},
				resource.Attribute{
					Name:        "sync_l4_persistent",
					Description: `sync l4 persistent. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "monitor_enable",
					Description: `remote ip monitor. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "hbdev",
					Description: `heartbeat port.`,
				},
				resource.Attribute{
					Name:        "mgmt_ip_allowaccess",
					Description: `Allow access with the management ip address. Valid values: 16:http, 32:telnet, 1:https, 2:ping, 4:ssh, 8:snmp .`,
				},
				resource.Attribute{
					Name:        "config_priority",
					Description: `ha config priority (0-255). (0,255)`,
				},
				resource.Attribute{
					Name:        "hb_lost_threshold",
					Description: `lost heartbeat threshold (1-60). (1,60)`,
				},
				resource.Attribute{
					Name:        "override",
					Description: `override on resurge. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "groupid",
					Description: `group id (0-31). (0,31)`,
				},
				resource.Attribute{
					Name:        "node_list",
					Description: `node id. Valid values: 1:1, 0:0, 3:3, 2:2, 5:5, 4:4, 7:7, 6:6 .`,
				},
				resource.Attribute{
					Name:        "local_address",
					Description: `For unicast local listening ip address.`,
				},
				resource.Attribute{
					Name:        "mgmt_ip",
					Description: `IP address of management interface.`,
				},
				resource.Attribute{
					Name:        "sync_l4_connection",
					Description: `sync l4 connection. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "datadev",
					Description: `data port.`,
				},
				resource.Attribute{
					Name:        "arp_num",
					Description: `num for sending (1-60). (1,60)`,
				},
				resource.Attribute{
					Name:        "failover_hold_time",
					Description: `failover hold time for remote ip (60-86400 sec). (60,86400)`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `heartbeat interval (1-20 (100`,
				},
				resource.Attribute{
					Name:        "arp_interval",
					Description: `interval for sending arp (1-20 sec). (1,20)`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `high availability mode. Valid values: 1:active-passive, 0:standalone, 3:active-active-vrrp, 2:active-active .`,
				},
				resource.Attribute{
					Name:        "mgmt_status",
					Description: `Management status enable/disable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "interface_list",
					Description: `interface list to track.`,
				},
				resource.Attribute{
					Name:        "sync_l7_persistent",
					Description: `sync http persistent. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "hbtype",
					Description: `heartbeat type. Valid values: 1:broadcast, 0:multicast, 2:unicast .`,
				},
				resource.Attribute{
					Name:        "failover_threshold",
					Description: `failover threshold for remote ip (1-64). (1,64)`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `ha priority (0-9). (0,9) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Ha can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_ha.labelname SystemHa ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Ha can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_ha.labelname SystemHa ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_health_check",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc system health-check info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"health",
				"check",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `health-check name.`,
				},
				resource.Attribute{
					Name:        "counter_value",
					Description: `snmp counter value: range(0, 2147483647).`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `the service name for connecting oracle db.`,
				},
				resource.Attribute{
					Name:        "basedn",
					Description: `baseDN.`,
				},
				resource.Attribute{
					Name:        "pwd_type",
					Description: `RADIUS password type. Valid values: 1:chap-password, 0:user-password .`,
				},
				resource.Attribute{
					Name:        "dest_addr6",
					Description: `dest-addr6.`,
				},
				resource.Attribute{
					Name:        "mssql_row",
					Description: `row.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `disk threshold for snmp. (1,99)`,
				},
				resource.Attribute{
					Name:        "row",
					Description: `the row for sql searching.`,
				},
				resource.Attribute{
					Name:        "http_version",
					Description: `http version. Valid values: 1:http_1.0, 2:http_1.1 .`,
				},
				resource.Attribute{
					Name:        "passive",
					Description: `FTP passive enable/disable. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "mssql_column",
					Description: `column.`,
				},
				resource.Attribute{
					Name:        "status_code",
					Description: `status-code.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `Mail folder name.`,
				},
				resource.Attribute{
					Name:        "addr_type",
					Description: `address type. Valid values: 1:ipv4, 2:ipv6 .`,
				},
				resource.Attribute{
					Name:        "mssql_receive_string",
					Description: `mssql receive string.`,
				},
				resource.Attribute{
					Name:        "oracle_receive_string",
					Description: `the receive-string for connecting oracle db.`,
				},
				resource.Attribute{
					Name:        "host_addr",
					Description: `host-addr.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `password.`,
				},
				resource.Attribute{
					Name:        "sip_request_type",
					Description: `type of SIP request. Valid values: 1:register, 2:options .`,
				},
				resource.Attribute{
					Name:        "remote_password",
					Description: `remote-password.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `timeout. (1,3600)`,
				},
				resource.Attribute{
					Name:        "acct_appid",
					Description: `diameter acct application id.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `port. (0,65535)`,
				},
				resource.Attribute{
					Name:        "allow_ssl_version",
					Description: `SSL/TLS versions allowed. Valid values: 8:tlsv1.1, 2:sslv3, 4:tlsv1.0, 32:tlsv1.3, 16:tlsv1.2 .`,
				},
				resource.Attribute{
					Name:        "remote_port",
					Description: `remote-port. (0,65535)`,
				},
				resource.Attribute{
					Name:        "connect_string",
					Description: `the connect data for connecting oracle db.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `domain-name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `snmp version. Valid values: 1:v1, 2:v2c .`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `product name.`,
				},
				resource.Attribute{
					Name:        "nas_ip",
					Description: `RADIUS NAS IP address.`,
				},
				resource.Attribute{
					Name:        "binddn",
					Description: `bindDN.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `mem threshold for snmp. (1,99)`,
				},
				resource.Attribute{
					Name:        "host_addr6",
					Description: `host-addr6.`,
				},
				resource.Attribute{
					Name:        "http_extra_string",
					Description: `additional-string.`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `the column for sql seaching.`,
				},
				resource.Attribute{
					Name:        "http_connect",
					Description: `http connect type. Valid values: 1:no_connect, 3:remote_connect, 2:local_connect .`,
				},
				resource.Attribute{
					Name:        "auth_appid",
					Description: `diameter auth application id.`,
				},
				resource.Attribute{
					Name:        "agent_type",
					Description: `snmp agent type. Valid values: 1:UCD, 2:WIN2000 .`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `filter.`,
				},
				resource.Attribute{
					Name:        "remote_host",
					Description: `remote-host.`,
				},
				resource.Attribute{
					Name:        "connect_type",
					Description: `the type of oracle connect data. Valid values: 1:service_name, 3:connect_string, 2:sid .`,
				},
				resource.Attribute{
					Name:        "method_type",
					Description: `method type. Valid values: 1:http_get, 2:http_head .`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `ssl ciphers including Perfect Forward Secrecy. Valid values: 32768:, 4194304:Perfect Forward Secrecy cipher, 1099511627776:Perfect Forward Secrecy cipher, 131072:Perfect Forward Secrecy cipher, 262144:Perfect Forward Secrecy cipher, 64:Perfect Forward Secrecy cipher, 65536:, 256:Perfect Forward Secrecy cipher, 33554432:, 8192:Perfect Forward Secrecy cipher, 16777216:, 536870912:Perfect Forward Secrecy cipher, 1048576:Perfect Forward Secrecy cipher, 1:Perfect Forward Secrecy cipher, 2097152:Perfect Forward Secrecy cipher, 2:Perfect Forward Secrecy cipher, 4:Perfect Forward Secrecy cipher, 8:Perfect Forward Secrecy cipher, 16384:, 17179869184:Non Encryption Ciphers, 2199023255552:Perfect Forward Secrecy cipher, 4096:Perfect Forward Secrecy cipher, 134217728:, 137438953472:Perfect Forward Secrecy cipher, 274877906944:Perfect Forward Secrecy cipher, 8388608:, 67108864:Perfect Forward Secrecy cipher, 524288:Perfect Forward Secrecy cipher, 549755813888:Perfect Forward Secrecy cipher, 128:Perfect Forward Secrecy cipher, 2147483648:, 1024:Perfect Forward Secrecy cipher, 268435456:, 34359738368:Perfect Forward Secrecy cipher, 16:Perfect Forward Secrecy cipher, 32:Perfect Forward Secrecy cipher, 2048:Perfect Forward Secrecy cipher, 1073741824:Perfect Forward Secrecy cipher, 512:Perfect Forward Secrecy cipher, 68719476736:Perfect Forward Secrecy cipher .`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `snmp community.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `health-check script.`,
				},
				resource.Attribute{
					Name:        "dest_addr",
					Description: `dest-addr.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `health-check type. Valid values: 24:mysql, 25:diameter, 26:oracle, 27:script, 20:sip, 21:sip-tcp, 22:snmp-custom, 23:rtsp, 28:ldap, 29:mssql, 1:icmp, 3:tcp, 2:tcp-echo, 5:https, 4:http, 7:radius, 6:dns, 9:pop3, 8:smtp, 11:radacct, 10:imap4, 12:ftp, 15:tcpssl, 14:tcphalf, 17:ssh, 16:snmp, 19:udp, 18:l2-detection .`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `username.`,
				},
				resource.Attribute{
					Name:        "rtsp_method_type",
					Description: `type of RTSP method. Valid values: 1:options, 2:describe .`,
				},
				resource.Attribute{
					Name:        "mssql_send_string",
					Description: `mssql send string.`,
				},
				resource.Attribute{
					Name:        "string_value",
					Description: `snmp string value.`,
				},
				resource.Attribute{
					Name:        "down_retry",
					Description: `retry. (1,10)`,
				},
				resource.Attribute{
					Name:        "origin_realm",
					Description: `origin realm.`,
				},
				resource.Attribute{
					Name:        "origin_host",
					Description: `origin host.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `RADIUS Secret Key.`,
				},
				resource.Attribute{
					Name:        "compare_type",
					Description: `snmp compare type. Valid values: 1:equal, 3:greater, 2:less .`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `match-type. Valid values: 1:match_string, 3:match_all, 2:match_status .`,
				},
				resource.Attribute{
					Name:        "attribute",
					Description: `attribute.`,
				},
				resource.Attribute{
					Name:        "local_cert",
					Description: `Local certificate reference.`,
				},
				resource.Attribute{
					Name:        "mysql_server_type",
					Description: `type of MySQL server. Valid values: 1:primary, 2:secondary .`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `File path on FTP server.`,
				},
				resource.Attribute{
					Name:        "dest_addr_type",
					Description: `destination address type. Valid values: 1:ipv4, 2:ipv6 .`,
				},
				resource.Attribute{
					Name:        "cpu_weight",
					Description: `snmp cpu weight. (0,200)`,
				},
				resource.Attribute{
					Name:        "host_ip_addr",
					Description: `host ip addr.`,
				},
				resource.Attribute{
					Name:        "host_ip6_addr",
					Description: `host ip6 addr.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `hostname.`,
				},
				resource.Attribute{
					Name:        "mem_weight",
					Description: `snmp mem weight. (0,200)`,
				},
				resource.Attribute{
					Name:        "sid",
					Description: `the sid for connecting oracle db.`,
				},
				resource.Attribute{
					Name:        "disk_weight",
					Description: `snmp disk weight. (0,200)`,
				},
				resource.Attribute{
					Name:        "receive_string",
					Description: `receive-string.`,
				},
				resource.Attribute{
					Name:        "oid",
					Description: `snmp oid.`,
				},
				resource.Attribute{
					Name:        "value_type",
					Description: `snmp value type. Valid values: 2:ASN_INTEGER, 4:ASN_OCTET_STR, 71:ASN_UINTEGER, 65:ASN_COUNTER .`,
				},
				resource.Attribute{
					Name:        "radius_reject",
					Description: `enbale radius reject. Valid values: 1:disable, 2:enable .`,
				},
				resource.Attribute{
					Name:        "rtsp_describe_url",
					Description: `URL for DESCRIBE method.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `datebase name of the mssql server.`,
				},
				resource.Attribute{
					Name:        "oracle_send_string",
					Description: `the send string for connecting oracle db.`,
				},
				resource.Attribute{
					Name:        "send_string",
					Description: `send-string.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `interval. (1,3600)`,
				},
				resource.Attribute{
					Name:        "up_retry",
					Description: `up-retry. (1,10)`,
				},
				resource.Attribute{
					Name:        "remote_username",
					Description: `remote-username.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `cpu threshold for snmp. (1,99)`,
				},
				resource.Attribute{
					Name:        "vendor_id",
					Description: `vendor id. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Health Check can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_health_check.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Health Check can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_health_check.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_health_check_script",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc the script of health-check.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"health",
				"check",
				"script",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `scripting name.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `scripting content. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Health Check Script can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_health_check_script.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Health Check Script can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_health_check_script.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_interface",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc interface configuration.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `interface name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `interface status. Valid values: 16:up .`,
				},
				resource.Attribute{
					Name:        "aggregate_mode",
					Description: `aggregate mode. Valid values: 0:balance-rr, 3:broadcast, 2:balance-xor, 5:balance-tlb, 4:802.3ad, 6:balance-alb .`,
				},
				resource.Attribute{
					Name:        "dhcp_gw_override",
					Description: `enable/disable to override gateway. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "secondary_ip",
					Description: `enable/disable to set secondary ip. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `floating ip address of interface.`,
				},
				resource.Attribute{
					Name:        "floating",
					Description: `enable/disable to set floating ip. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `speed. Valid values: 1:10half, 0:auto, 3:100half, 2:10full, 5:1000half, 4:100full, 7:10000full, 6:1000full, 9:100000full, 8:40000full .`,
				},
				resource.Attribute{
					Name:        "redundant_member",
					Description: `redundant/aggregate interface slaves.`,
				},
				resource.Attribute{
					Name:        "dhcp_ip_overlap",
					Description: `dhcp_ip_is_overlap.`,
				},
				resource.Attribute{
					Name:        "dhcp_gw_distance",
					Description: `distance for dhcp_gateway. (1,255)`,
				},
				resource.Attribute{
					Name:        "wccp",
					Description: `enable/disable WCCP. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "default_gw",
					Description: `retrieve default gateway from pppoe server. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "trust_ip",
					Description: `enable/disable to set trust ip. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "vlanid",
					Description: `Vlan ID. (1,4094)`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `vdom name.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Allow access with the interface. Valid values: 16:http, 32:telnet, 1:https, 2:ping, 4:ssh, 8:snmp .`,
				},
				resource.Attribute{
					Name:        "dns_server_override",
					Description: `use pppoe server DNS or not. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `interface type. Valid values: 10:vdom-link, 1:vlan, 0:physical, 2:aggregate, 7:soft-switch, 6:loopback, 9:nvgre, 8:vxlan .`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `ip address of interface.`,
				},
				resource.Attribute{
					Name:        "traffic_group",
					Description: `traffic group name.`,
				},
				resource.Attribute{
					Name:        "disc_retry_timeout",
					Description: `pppoe discovery retry timeout. (1,255)`,
				},
				resource.Attribute{
					Name:        "dhcp_gateway",
					Description: `gateway.`,
				},
				resource.Attribute{
					Name:        "ip6",
					Description: `ipv6 address of interface.`,
				},
				resource.Attribute{
					Name:        "dedicate_to_management",
					Description: `dedicate to managemnet interface. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `physical interface name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `pppoe account password.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `maximum transportation unit. (68,9000)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `pppoe account user name.`,
				},
				resource.Attribute{
					Name:        "floating_ip6",
					Description: `floating ipv6 address of interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `static,pppoe or dhcp. Valid values: 1:dhcp, 0:static, 2:pppoe .`,
				},
				resource.Attribute{
					Name:        "aggregate_algorithm",
					Description: `aggregate interface slave selection algorithm. Valid values: 1:layer3-4, 0:layer2, 2:layer2-3 . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Interface can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_interface.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Interface can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_interface.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_interface_child_ha_node_ip_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc interface configuration.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"interface",
				"child",
				"ha",
				"node",
				"ip",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "fadc_id",
					Description: `Number of IPs in list.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `Node ID of HA Node. (0,7)`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Allow access with the IP. Valid values: 16:http, 32:telnet, 1:https, 2:ping, 4:ssh, 8:snmp . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Interface Child Ha Node Ip List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_interface_child_ha_node_ip_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Interface Child Ha Node Ip List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_interface_child_ha_node_ip_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_interface_child_secondary_ip_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc interface configuration.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"interface",
				"child",
				"secondary",
				"ip",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "fadc_id",
					Description: `secondary ip number.`,
				},
				resource.Attribute{
					Name:        "traffic_group",
					Description: `traffic group name.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `secondary ip address of interface.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `floating ip address of interface.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Scondary Allow access with the interface. Valid values: 16:http, 32:telnet, 1:https, 2:ping, 4:ssh, 8:snmp .`,
				},
				resource.Attribute{
					Name:        "floating",
					Description: `enable/disable to set floating ip. Valid values: enable/disable. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Interface Child Secondary Ip List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_interface_child_secondary_ip_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Interface Child Secondary Ip List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_interface_child_secondary_ip_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_interface_child_trust_ip_list",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc interface configuration.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"interface",
				"child",
				"trust",
				"ip",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `trust ip name.`,
				},
				resource.Attribute{
					Name:        "ip6_max",
					Description: `end trust IP6.`,
				},
				resource.Attribute{
					Name:        "ip6_min",
					Description: `start trust IP6.`,
				},
				resource.Attribute{
					Name:        "ip_max",
					Description: `end trust ip.`,
				},
				resource.Attribute{
					Name:        "ip_netmask",
					Description: `trust ip/netmask.`,
				},
				resource.Attribute{
					Name:        "ip6_netmask",
					Description: `trust ip6/netmask.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `trust ip type. Valid values: 1:ip-netmask, 3:ip6-netmask, 2:ip-range, 4:ip6-range .`,
				},
				resource.Attribute{
					Name:        "ip_min",
					Description: `start trust IP. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Interface Child Trust Ip List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_interface_child_trust_ip_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Interface Child Trust Ip List can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_interface_child_trust_ip_list.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_isp_addr",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc isp address group.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"isp",
				"addr",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `The name of the addr isp.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of isp-addr. Valid values: 1:predef, 2:restored, 4:userdef . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Isp Addr can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_isp_addr.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Isp Addr can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_isp_addr.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_isp_addr_child_address",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc isp address group.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"isp",
				"addr",
				"child",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `entry id.`,
				},
				resource.Attribute{
					Name:        "province",
					Description: `province name.`,
				},
				resource.Attribute{
					Name:        "ip_netmask",
					Description: `IP netmask. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Isp Addr Child Address can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_isp_addr_child_address.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Isp Addr Child Address can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_isp_addr_child_address.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_isp_province",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc isp provice info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"isp",
				"province",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `The name of province. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Isp Province can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_isp_province.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Isp Province can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_isp_province.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_mailserver",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc mail server configuration.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"mailserver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `username for SMTP authentication.`,
				},
				resource.Attribute{
					Name:        "auth",
					Description: `enable/disable SMTP authentication. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `SMTP server address.`,
				},
				resource.Attribute{
					Name:        "security",
					Description: `set security option. Valid values: 1:starttls, 2:none .`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `password for SMTP authentication.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `SMTP server port. (1,65535) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Mailserver can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_mailserver.labelname SystemMailserver ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Mailserver can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_mailserver.labelname SystemMailserver ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_overlay_tunnel",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc overlay tunnel.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"overlay",
				"tunnel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `overlay tunnel name.`,
				},
				resource.Attribute{
					Name:        "dstip",
					Description: `the IPv4 addresses of the remote VTEPs.`,
				},
				resource.Attribute{
					Name:        "tunnel_type",
					Description: `interface type. Valid values: 1:nvgre, 0:vxlan .`,
				},
				resource.Attribute{
					Name:        "vni",
					Description: `VXLAN Network Identifier. (1,16777215)`,
				},
				resource.Attribute{
					Name:        "mttl",
					Description: `TTL of multicast packet. (1,255)`,
				},
				resource.Attribute{
					Name:        "ipversion",
					Description: `the ip type used in destination-ip-addresses. Valid values: 1:ipv4-multicast, 0:ipv4-unicast .`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `the outgoing interface.`,
				},
				resource.Attribute{
					Name:        "vsid",
					Description: `Virtual Subnet Identifier. (1,16777215)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `the VxLAN UDP port in destination. (1,65535) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Overlay Tunnel can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_overlay_tunnel.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Overlay Tunnel can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_overlay_tunnel.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_overlay_tunnel_child_remote_host",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc overlay tunnel.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"overlay",
				"tunnel",
				"child",
				"remote",
				"host",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `host id.`,
				},
				resource.Attribute{
					Name:        "vtep",
					Description: `the IPv4 address of virtual tunnel endpoint.`,
				},
				resource.Attribute{
					Name:        "host_mac",
					Description: `the MAC address of the remote host. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Overlay Tunnel Child Remote Host can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_overlay_tunnel_child_remote_host.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Overlay Tunnel Child Remote Host can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_overlay_tunnel_child_remote_host.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_schedule_group",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc schedule group info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"schedule",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `schedule group name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Schedule Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_schedule_group.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Schedule Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_schedule_group.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_schedule_group_child_schedule_member",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc schedule group info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"schedule",
				"group",
				"child",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "pkey",
					Description: `The parent key.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `schedule member name.`,
				},
				resource.Attribute{
					Name:        "startdate",
					Description: `start date: YYYY/MM/DD.`,
				},
				resource.Attribute{
					Name:        "enddate",
					Description: `end date: YYYY/MM/DD.`,
				},
				resource.Attribute{
					Name:        "day_of_month",
					Description: `monthly day. (1,31)`,
				},
				resource.Attribute{
					Name:        "day_of_week",
					Description: `weekly day. Valid values: 1:monday, 0:sunday, 3:wednesday, 2:tuesday, 5:friday, 4:thursday, 6:saturday .`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `member type. Valid values: 1:monthly-recurring, 0:one-time, 3:daily-recurring, 2:weekly-recurring, 4:hourly-recurring .`,
				},
				resource.Attribute{
					Name:        "starttime_of_startdate",
					Description: `start time of start date: HH:MM.`,
				},
				resource.Attribute{
					Name:        "endtime_of_enddate",
					Description: `end time of end date: HH:MM. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Schedule Group Child Schedule Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_schedule_group_child_schedule_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Schedule Group Child Schedule Member can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_schedule_group_child_schedule_member.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_scripting",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc system scripting.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"scripting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `scripting name.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `scripting. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Scripting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_scripting.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Scripting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_scripting.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_sdn_connector",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc sdn connector.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"sdn",
				"connector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `sdn connector name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `enable/disable sdn connector. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "oci_region_type",
					Description: `sdn connector oci region type. Valid values: 1:GOVERNMENT_REGION, 0:COMMERCIAL_REGION .`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `server port. (1,65535)`,
				},
				resource.Attribute{
					Name:        "oci_cert",
					Description: `sdn connector oci cert.`,
				},
				resource.Attribute{
					Name:        "sap_ms_http_port",
					Description: `sap message server http port. (1,65535)`,
				},
				resource.Attribute{
					Name:        "aws_secretkey",
					Description: `sdn connector secret token.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `interval (seconds). (30,3600)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `sdn connector type. Valid values: 10:kubernetes, 13:sap, 2:aws, 7:oci .`,
				},
				resource.Attribute{
					Name:        "sap_icm_http_port",
					Description: `sap icm http port. (1,65535)`,
				},
				resource.Attribute{
					Name:        "secret_token",
					Description: `sdn connector secret token.`,
				},
				resource.Attribute{
					Name:        "oci_tenant_id",
					Description: `sdn connector oci tenant id.`,
				},
				resource.Attribute{
					Name:        "oci_region",
					Description: `sdn connector oci region.`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `sdn connector aws region.`,
				},
				resource.Attribute{
					Name:        "oci_user_id",
					Description: `sdn connector oci user id.`,
				},
				resource.Attribute{
					Name:        "aws_accesskey",
					Description: `sdn connector secret token.`,
				},
				resource.Attribute{
					Name:        "use_metadata_iam",
					Description: `enable/disable using iam role. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "sap_dns_suffix",
					Description: `dns name for sap system.`,
				},
				resource.Attribute{
					Name:        "oci_ha_status",
					Description: `enable/disable use HA service. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "sap_password",
					Description: `sap soap sid admin password.`,
				},
				resource.Attribute{
					Name:        "sap_sidadm",
					Description: `sap soap sid admin.`,
				},
				resource.Attribute{
					Name:        "oci_compartment_id",
					Description: `sdn connector oci compartment id.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `server ip/FQDN/host(host is only for sap). ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Sdn Connector can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_sdn_connector.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Sdn Connector can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_sdn_connector.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_service",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc system service.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "destination_port_max",
					Description: `destination port max. (1,65535)`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `protocol number. (1,255)`,
				},
				resource.Attribute{
					Name:        "source_port_min",
					Description: `source port min. (1,65535)`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `protocol type. Valid values: 1:ip, 3:tcp, 2:icmp, 5:tcp-and-udp, 4:udp, 6:sctp .`,
				},
				resource.Attribute{
					Name:        "destination_port_min",
					Description: `destination port min. (1,65535)`,
				},
				resource.Attribute{
					Name:        "specify_source_port",
					Description: `specify source port or not. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "source_port_max",
					Description: `source port max. (1,65535) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Service can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_service.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Service can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_service.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_servicegrp",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc service group info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"servicegrp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `service group name.`,
				},
				resource.Attribute{
					Name:        "member_list",
					Description: `member list. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Servicegrp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_servicegrp.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Servicegrp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_servicegrp.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_stream_scripting",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc system stream scripting.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"stream",
				"scripting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `stream scripting name.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `scripting. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Stream Scripting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_stream_scripting.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Stream Scripting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_stream_scripting.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_time_manual",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc set date and time.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"time",
				"manual",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "tz",
					Description: `time zone. (0,71)`,
				},
				resource.Attribute{
					Name:        "dst",
					Description: `enable daylight saving time. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "ntpsync",
					Description: `enable/disable synchronization with NTP server. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "syncinterval",
					Description: `synchronization time interval. (1,1440)`,
				},
				resource.Attribute{
					Name:        "ntpserver",
					Description: `IP address or hostname of NTP server.`,
				},
				resource.Attribute{
					Name:        "year",
					Description: `year (int)`,
				},
				resource.Attribute{
					Name:        "month",
					Description: `month (int)`,
				},
				resource.Attribute{
					Name:        "mday",
					Description: `mday (int)`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `hour (int)`,
				},
				resource.Attribute{
					Name:        "minute",
					Description: `minute (int)`,
				},
				resource.Attribute{
					Name:        "second",
					Description: `second (int) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Time Manual can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_time_manual.labelname SystemTimeManual ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Time Manual can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_time_manual.labelname SystemTimeManual ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_traffic_group",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc traffic group parameters configuration.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"traffic",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `traffic group name.`,
				},
				resource.Attribute{
					Name:        "network_failover",
					Description: `network failover track. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "preempt",
					Description: `preempt mode. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "failover_order",
					Description: `node id. (0,7) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Traffic Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_traffic_group.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Traffic Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_traffic_group.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_system_vdom_link",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc system vdom-link.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"system",
				"vdom",
				"link",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `vdom-link name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type of vdom-link. Valid values: 1:ethernet, 0:ppp . ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Vdom Link can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_vdom_link.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import System Vdom Link can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_system_vdom_link.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_upload_captcha_cust",
			Category:         "FortiADC Resources",
			ShortDescription: `Upload fortiadc captcha.`,
			Description:      ``,
			Keywords: []string{
				"upload",
				"captcha",
				"cust",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "pic_difficulty",
					Description: `Verification picture difficulty. Valid values: 1:hard, 0:easy .`,
				},
				resource.Attribute{
					Name:        "max_block_period",
					Description: `User blocked time after verification failure, unit second. (10,864000)`,
				},
				resource.Attribute{
					Name:        "custom_page",
					Description: `If use customize captcha page file.. Valid values: enable/disable.`,
				},
				resource.Attribute{
					Name:        "max_valid_period",
					Description: `Maximum valid time for verified users, unit second. (60,864000)`,
				},
				resource.Attribute{
					Name:        "max_picture_changes",
					Description: `Maximum number of times the user can change the captcha picture. (1,100)`,
				},
				resource.Attribute{
					Name:        "max_verify_period",
					Description: `Maximum user verification process, unit second. (20,86400)`,
				},
				resource.Attribute{
					Name:        "max_retry_times",
					Description: `The maximum number of times a user can try. (1,100)`,
				},
				resource.Attribute{
					Name:        "vpath",
					Description: `virtual path for error page.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Create or update (true/false).`,
				},
				resource.Attribute{
					Name:        "errorpagefile",
					Description: `captcha file path.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_upload_error_page",
			Category:         "FortiADC Resources",
			ShortDescription: `Upload fortiadc error page.`,
			Description:      ``,
			Keywords: []string{
				"upload",
				"error",
				"page",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mkey",
					Description: `name.`,
				},
				resource.Attribute{
					Name:        "vpath",
					Description: `virtual path for error page.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Create or update (true/false).`,
				},
				resource.Attribute{
					Name:        "errorpagefile",
					Description: `error page file path.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_user_adfs_publish",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc adfs publish info.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"user",
				"adfs",
				"publish",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdom",
					Description: `Specifies the vdom to which the data source will be applied when the FortiADC unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.`,
				},
				resource.Attribute{
					Name:        "mkey",
					Description: `adfs-proxy name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `enable/disable adfs publish. Valid values: 1:enable, 0:disable .`,
				},
				resource.Attribute{
					Name:        "backend_server_url",
					Description: `Backend Server URL.`,
				},
				resource.Attribute{
					Name:        "external_url",
					Description: `External URL.`,
				},
				resource.Attribute{
					Name:        "preauth",
					Description: `preauth method. Valid values: 1:ADFS, 0:Pass-through .`,
				},
				resource.Attribute{
					Name:        "relying_party",
					Description: `Relying Party.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `AD FS proxy. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import User Adfs Publish can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_user_adfs_publish.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}. ## Import User Adfs Publish can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fortiadc_user_adfs_publish.labelname {{mkey}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortiadc_vdom",
			Category:         "FortiADC Data Source FortiADC Resources",
			ShortDescription: `Configure fortiadc route configuration.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"source",
				"vdom",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mkey",
					Description: `The vdom name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{mkey}}.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"fortiadc_ca_certificateupload":                                        0,
		"fortiadc_certificate_crlupload":                                       1,
		"fortiadc_certificate_intmed_caupload":                                 2,
		"fortiadc_config_sync_list":                                            3,
		"fortiadc_firewall_nat_snat":                                           4,
		"fortiadc_firewall_vip":                                                5,
		"fortiadc_load_balance_allowlist":                                      6,
		"fortiadc_load_balance_auth_policy":                                    7,
		"fortiadc_load_balance_caching":                                        8,
		"fortiadc_load_balance_caching_child_dyn_cache_list":                   9,
		"fortiadc_load_balance_caching_child_uri_exclude_list":                 10,
		"fortiadc_load_balance_captcha_profile":                                11,
		"fortiadc_load_balance_certificate_caching":                            12,
		"fortiadc_load_balance_client_ssl_profile":                             13,
		"fortiadc_load_balance_clone_pool":                                     14,
		"fortiadc_load_balance_clone_pool_child_pool_member":                   15,
		"fortiadc_load_balance_compression":                                    16,
		"fortiadc_load_balance_compression_child_content_types":                17,
		"fortiadc_load_balance_compression_child_uri_list":                     18,
		"fortiadc_load_balance_content_rewriting":                              19,
		"fortiadc_load_balance_content_rewriting_child_match_condition":        20,
		"fortiadc_load_balance_content_routing":                                21,
		"fortiadc_load_balance_content_routing_child_match_condition":          22,
		"fortiadc_load_balance_decompression":                                  23,
		"fortiadc_load_balance_decompression_child_content_types":              24,
		"fortiadc_load_balance_decompression_child_uri_list":                   25,
		"fortiadc_load_balance_geoip_list":                                     26,
		"fortiadc_load_balance_http2_profile":                                  27,
		"fortiadc_load_balance_ippool":                                         28,
		"fortiadc_load_balance_ippool_child_node_member":                       29,
		"fortiadc_load_balance_l2_exception_list":                              30,
		"fortiadc_load_balance_l2_exception_list_child_member":                 31,
		"fortiadc_load_balance_method":                                         32,
		"fortiadc_load_balance_pagespeed":                                      33,
		"fortiadc_load_balance_pagespeed_child_page_control":                   34,
		"fortiadc_load_balance_pagespeed_child_resource_control":               35,
		"fortiadc_load_balance_pagespeed_profile":                              36,
		"fortiadc_load_balance_persistence":                                    37,
		"fortiadc_load_balance_persistence_child_iso8583_bitmap":               38,
		"fortiadc_load_balance_persistence_child_radius_attribute":             39,
		"fortiadc_load_balance_pool":                                           40,
		"fortiadc_load_balance_pool_child_pool_member":                         41,
		"fortiadc_load_balance_profile":                                        42,
		"fortiadc_load_balance_profile_child_client_request_header_erase":      43,
		"fortiadc_load_balance_profile_child_client_request_header_insert":     44,
		"fortiadc_load_balance_profile_child_client_response_header_erase":     45,
		"fortiadc_load_balance_profile_child_client_response_header_insert":    46,
		"fortiadc_load_balance_profile_child_mssql_user_password":              47,
		"fortiadc_load_balance_profile_child_mysql_rule":                       48,
		"fortiadc_load_balance_profile_child_mysql_sharding":                   49,
		"fortiadc_load_balance_profile_child_mysql_user_password":              50,
		"fortiadc_load_balance_profile_child_server_request_header_erase":      51,
		"fortiadc_load_balance_profile_child_server_request_header_insert":     52,
		"fortiadc_load_balance_profile_child_server_response_header_erase":     53,
		"fortiadc_load_balance_profile_child_server_response_header_insert":    54,
		"fortiadc_load_balance_real_server":                                    55,
		"fortiadc_load_balance_real_server_ssl_profile":                        56,
		"fortiadc_load_balance_schedule_pool":                                  57,
		"fortiadc_load_balance_virtual_server":                                 58,
		"fortiadc_load_balance_web_category":                                   59,
		"fortiadc_load_balance_web_filter_profile":                             60,
		"fortiadc_load_balance_web_filter_profile_child_category_members":      61,
		"fortiadc_load_balance_web_sub_category":                               62,
		"fortiadc_router_access_list":                                          63,
		"fortiadc_router_access_list6":                                         64,
		"fortiadc_router_access_list6_child_rule":                              65,
		"fortiadc_router_access_list_child_rule":                               66,
		"fortiadc_router_bgp":                                                  67,
		"fortiadc_router_bgp_child_ha_router_id_list":                          68,
		"fortiadc_router_bgp_child_neighbor":                                   69,
		"fortiadc_router_bgp_child_network":                                    70,
		"fortiadc_router_isp":                                                  71,
		"fortiadc_router_md5_ospf":                                             72,
		"fortiadc_router_md5_ospf_child_md5_member":                            73,
		"fortiadc_router_ospf":                                                 74,
		"fortiadc_router_ospf_child_area":                                      75,
		"fortiadc_router_ospf_child_ha_router_id_list":                         76,
		"fortiadc_router_ospf_child_network":                                   77,
		"fortiadc_router_ospf_child_ospf_interface":                            78,
		"fortiadc_router_policy":                                               79,
		"fortiadc_router_prefix_list":                                          80,
		"fortiadc_router_prefix_list6":                                         81,
		"fortiadc_router_prefix_list6_child_rule":                              82,
		"fortiadc_router_prefix_list_child_rule":                               83,
		"fortiadc_router_static":                                               84,
		"fortiadc_security_antivirus_profile":                                  85,
		"fortiadc_security_dos_dos_protection_profile":                         86,
		"fortiadc_security_ips_profile":                                        87,
		"fortiadc_security_waf_profile":                                        88,
		"fortiadc_security_ztna_profile":                                       89,
		"fortiadc_system_address":                                              90,
		"fortiadc_system_address6":                                             91,
		"fortiadc_system_addrgrp":                                              92,
		"fortiadc_system_addrgrp6":                                             93,
		"fortiadc_system_auto_backup":                                          94,
		"fortiadc_system_azure_lb_backend_ip":                                  95,
		"fortiadc_system_certificate_ca_group":                                 96,
		"fortiadc_system_certificate_ca_group_child_group_member":              97,
		"fortiadc_system_certificate_certificate_verify":                       98,
		"fortiadc_system_certificate_crl":                                      99,
		"fortiadc_system_certificate_intermediate_ca":                          100,
		"fortiadc_system_certificate_intermediate_ca_group":                    101,
		"fortiadc_system_certificate_intermediate_ca_group_child_group_member": 102,
		"fortiadc_system_certificate_local":                                    103,
		"fortiadc_system_certificate_local_cert_group":                         104,
		"fortiadc_system_certificate_local_cert_group_child_group_member":      105,
		"fortiadc_system_certificate_ocsp":                                     106,
		"fortiadc_system_dns":                                                  107,
		"fortiadc_system_global":                                               108,
		"fortiadc_system_ha":                                                   109,
		"fortiadc_system_health_check":                                         110,
		"fortiadc_system_health_check_script":                                  111,
		"fortiadc_system_interface":                                            112,
		"fortiadc_system_interface_child_ha_node_ip_list":                      113,
		"fortiadc_system_interface_child_secondary_ip_list":                    114,
		"fortiadc_system_interface_child_trust_ip_list":                        115,
		"fortiadc_system_isp_addr":                                             116,
		"fortiadc_system_isp_addr_child_address":                               117,
		"fortiadc_system_isp_province":                                         118,
		"fortiadc_system_mailserver":                                           119,
		"fortiadc_system_overlay_tunnel":                                       120,
		"fortiadc_system_overlay_tunnel_child_remote_host":                     121,
		"fortiadc_system_schedule_group":                                       122,
		"fortiadc_system_schedule_group_child_schedule_member":                 123,
		"fortiadc_system_scripting":                                            124,
		"fortiadc_system_sdn_connector":                                        125,
		"fortiadc_system_service":                                              126,
		"fortiadc_system_servicegrp":                                           127,
		"fortiadc_system_stream_scripting":                                     128,
		"fortiadc_system_time_manual":                                          129,
		"fortiadc_system_traffic_group":                                        130,
		"fortiadc_system_vdom_link":                                            131,
		"fortiadc_upload_captcha_cust":                                         132,
		"fortiadc_upload_error_page":                                           133,
		"fortiadc_user_adfs_publish":                                           134,
		"fortiadc_vdom":                                                        135,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
