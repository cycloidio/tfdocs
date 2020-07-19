package aci

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "aci_local_user",
			Category:         "Resources",
			ShortDescription: `Manages ACI Local User`,
			Description:      ``,
			Keywords: []string{
				"local",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object local_user.`,
				},
				resource.Attribute{
					Name:        "account_status",
					Description: `(Optional) local AAA user account status`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object local_user.`,
				},
				resource.Attribute{
					Name:        "cert_attribute",
					Description: `(Optional) cert_attribute for object local_user.`,
				},
				resource.Attribute{
					Name:        "clear_pwd_history",
					Description: `(Optional) clear password history of local user`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) email address of the local user`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) local user account expiration date`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `(Optional) enables local user account expiration`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Optional) first name of the local user`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Optional) last name of the local user`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object local_user.`,
				},
				resource.Attribute{
					Name:        "otpenable",
					Description: `(Optional) otpenable for object local_user.`,
				},
				resource.Attribute{
					Name:        "otpkey",
					Description: `(Optional) otpkey for object local_user.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `(Optional) phone number of the local user`,
				},
				resource.Attribute{
					Name:        "pwd",
					Description: `(Optional) system user password`,
				},
				resource.Attribute{
					Name:        "pwd_life_time",
					Description: `(Optional) lifetime of the local user password`,
				},
				resource.Attribute{
					Name:        "pwd_update_required",
					Description: `(Optional) pwd_update_required for object local_user.`,
				},
				resource.Attribute{
					Name:        "rbac_string",
					Description: `(Optional) rbac_string for object local_user.`,
				},
				resource.Attribute{
					Name:        "unix_user_id",
					Description: `(Optional) UNIX identifier of the local user ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Local User. ## Importing ## An existing Local User can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_local_user.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_x509_certificate",
			Category:         "Resources",
			ShortDescription: `Manages ACI X509 Certificate`,
			Description:      ``,
			Keywords: []string{
				"x509",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "local_user_dn",
					Description: `(Required) Distinguished name of parent LocalUser object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object x509_certificate.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object x509_certificate.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Optional) data from the user certificate`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object x509_certificate. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the X509 Certificate. ## Importing ## An existing X509 Certificate can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_x509_certificate.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cdp_interface_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI CDP Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"cdp",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cdp_interface_policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) administrative state`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cdp_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cdp_interface_policy. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the CDP Interface Policy. ## Importing ## An existing CDP Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cdp_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_applicationcontainer",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud Application container`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"applicationcontainer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_applicationcontainer.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_applicationcontainer.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_applicationcontainer. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Application container. ## Importing ## An existing Cloud Application container can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_applicationcontainer.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_aws_provider",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud AWS Provider`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"aws",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `(Optional) access_key_id for the AWS account provided in the account_id field.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) AWS account-id to manage with cloud APIC.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) email address of the local user.`,
				},
				resource.Attribute{
					Name:        "http_proxy",
					Description: `(Optional) http_proxy for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "is_account_in_org",
					Description: `(Optional) Flag to decide whether the account is in the organization or not.`,
				},
				resource.Attribute{
					Name:        "is_trusted",
					Description: `(Optional) Whether the account is trusted with Tenant infra account.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `(Optional) provider_id for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) which AWS region to manage.`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional) secret_access_key for the AWS account provided in the account_id field. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud AWS Provider. ## Importing ## An existing Cloud AWS Provider can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_aws_provider.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_autonomous_system_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI Autonomous System Profile`,
			Description:      ``,
			Keywords: []string{
				"autonomous",
				"system",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object autonomous_system_profile.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Optional) A number that uniquely identifies an autonomous system.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object autonomous_system_profile. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Autonomous System Profile. ## Importing ## An existing Autonomous System Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_autonomous_system_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_cidr_pool",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud CIDR Pool`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"cidr",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_context_profile_dn",
					Description: `(Required) Distinguished name of parent CloudContextProfile object.`,
				},
				resource.Attribute{
					Name:        "addr",
					Description: `(Required) CIDR IPv4 block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_cidr_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_cidr_pool.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Optional) Flag to specify whether CIDR is primary CIDR or not. Allowed values are "yes" and "no". Default is "no". Only one primary CIDR is supported under a cloud context profile. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud CIDR Pool. ## Importing ## An existing Cloud CIDR Pool can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_cidr_pool.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_context_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud Context Profile`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"context",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_context_profile.`,
				},
				resource.Attribute{
					Name:        "primary_cidr",
					Description: `(Required) Primary CIDR block of Cloud Context profile.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) AWS region in which profile is created.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_context_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_context_profile.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The specific type of the object or component. Allowed values are "regular" and "shadow". Default is "regular".`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_ctx_to_flow_log",
					Description: `(Optional) Relation to class cloudAwsFlowLogPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_to_ctx",
					Description: `(Required) Relation to class fvCtx. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_ctx_profile_to_region",
					Description: `(Optional) Relation to class cloudRegion. Cardinality - N_TO_ONE. Type - String.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_domain_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud Domain Profile`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"domain",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_domain_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_domain_profile.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Optional) site_id for object cloud_domain_profile. Allowed value range is "0" to "1000". Default is "0". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Domain Profile. ## Importing ## An existing Cloud Domain Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_domain_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_epg",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud EPg`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"epg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_applicationcontainer_dn",
					Description: `(Required) Distinguished name of parent CloudApplicationcontainer object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_epg.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_epg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object cloud_epg. Allowed value range is "0" to "512".`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings. Allowed values are "disabled" and "enabled". Default is "disabled".`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria. Allowed values are "All", "AtleastOne", "AtmostOne" and "None". Default values is "AtleastOne".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_epg.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication. Allowed values are "include" and "exclude". Default is "exclude".`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id. Allowed values are "unspecified", "level1", "level2", "level3", "level4", "level5" and "level6". Default is "unspecified.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_sec_inherited",
					Description: `(Optional) Relation to class fvEPg. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prov",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons_if",
					Description: `(Optional) Relation to class vzCPIf. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cust_qos_pol",
					Description: `(Optional) Relation to class qosCustomPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_cloud_epg_ctx",
					Description: `(Optional) Relation to class fvCtx. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prot_by",
					Description: `(Optional) Relation to class vzTaboo. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_intra_epg",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud EPg. ## Importing ## An existing Cloud EPg can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_epg.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_endpoint_selector",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud Endpoint Selector`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"endpoint",
				"selector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_epg_dn",
					Description: `(Required) Distinguished name of parent CloudEPg object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_endpoint_selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_endpoint_selector.`,
				},
				resource.Attribute{
					Name:        "match_expression",
					Description: `(Optional) Match expression for the endpoint selector to select EP on criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_endpoint_selector. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Endpoint Selector. ## Importing ## An existing Cloud Endpoint Selector can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_endpoint_selector.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_external_epg",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud External EPg`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"external",
				"epg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_applicationcontainer_dn",
					Description: `(Required) Distinguished name of parent CloudApplicationcontainer object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_external_epg.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_external_epg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object cloud_external_epg. Allowed value range is "0" to "512".`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings. Allowed values are "disabled" and "enabled". Default is "disabled".`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria. Allowed values are "All", "AtleastOne", "AtmostOne" and "None". Default values is "AtleastOne".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_external_epg.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication. Allowed values are "include" and "exclude". Default is "exclude".`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id. Allowed values are "unspecified", "level1", "level2", "level3", "level4", "level5" and "level6". Default is "unspecified.`,
				},
				resource.Attribute{
					Name:        "route_reachability",
					Description: `(Optional) Route reachability for this EPG. Allowed values are "unspecified", "inter-site", "internet" and "inter-site-ext".Default is "inter-site".`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_sec_inherited",
					Description: `(Optional) Relation to class fvEPg. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prov",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons_if",
					Description: `(Optional) Relation to class vzCPIf. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cust_qos_pol",
					Description: `(Optional) Relation to class qosCustomPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_cloud_epg_ctx",
					Description: `(Optional) Relation to class fvCtx. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prot_by",
					Description: `(Optional) Relation to class vzTaboo. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_intra_epg",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud External EPg. ## Importing ## An existing Cloud External EPg can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_external_epg.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_endpoint_selectorfor_external_epgs",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud Endpoint Selector for External EPgs`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"endpoint",
				"selectorfor",
				"external",
				"epgs",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_external_epg_dn",
					Description: `(Required) Distinguished name of parent CloudExternalEPg object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_endpoint_selectorfor_external_epgs.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_endpoint_selectorfor_external_epgs.`,
				},
				resource.Attribute{
					Name:        "is_shared",
					Description: `(Optional) For Selectors set the shared route control. Allowed values are "yes" and "no". Default value is "yes".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_endpoint_selectorfor_external_epgs.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) Subnet from which EP to select. Any valid CIDR block is allowed here. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Endpoint Selector for External EPgs. ## Importing ## An existing Cloud Endpoint Selector for External EPgs can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_endpoint_selectorfor_external_epgs.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_provider_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud Provider Profile`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"provider",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vendor",
					Description: `(Required) vendor of Object cloud_provider_profile. Currently only supported vendor is "aws".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_provider_profile. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Provider Profile. ## Importing ## An existing Cloud Provider Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_provider_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_providers_region",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud Providers Region`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"providers",
				"region",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_provider_profile_dn",
					Description: `(Required) Distinguished name of parent CloudProviderProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_providers_region.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_providers_region.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_providers_region. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Providers Region. ## Importing ## An existing Cloud Providers Region can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_providers_region.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_subnet",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud Subnet`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_cidr_pool_dn",
					Description: `(Required) Distinguished name of parent CloudCIDRPool object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) CIDR block of Object cloud_subnet.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_subnet.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The domain applicable to the capability. Allowed values are "public", "private" and "shared". Default is "private".`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `(Optional) The usage of the port. This property shows how the port is used. Allowed values are "user" and "infra-router". Default is "user".`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) [AWS Only] Availability zone where the subnet must be deployed. This property can carry both the actual zone or the ACI logical zone name. In the former case, driver directly uses the value of this property. In the latter case, Connector has to first resolve the mapping from ACI logical zone to actual AWS zone. This parameter is required in APIC v5.0 or higher.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_zone_attach",
					Description: `(Optional) Relation to class cloudZone. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_subnet_to_flow_log",
					Description: `(Optional) Relation to class cloudAwsFlowLogPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Subnet. ## Importing ## An existing Cloud Subnet can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_subnet.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_availability_zone",
			Category:         "Resources",
			ShortDescription: `Manages ACI Cloud Availability Zone`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"availability",
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_providers_region_dn",
					Description: `(Required) Distinguished name of parent CloudProvidersRegion object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_availability_zone. Should match the Availability zone name in AWS cloud.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_availability_zone.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_availability_zone. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Availability Zone. ## Importing ## An existing Cloud Availability Zone can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_availability_zone.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_configuration_export_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI Configuration Export Policy`,
			Description:      ``,
			Keywords: []string{
				"configuration",
				"export",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) admin state of the export policy`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) export data format`,
				},
				resource.Attribute{
					Name:        "include_secure_fields",
					Description: `(Optional) include_secure_fields for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "max_snapshot_count",
					Description: `(Optional) max_snapshot_count for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional) snapshot for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "target_dn",
					Description: `(Optional) target export object`,
				},
				resource.Attribute{
					Name:        "relation_config_rs_export_destination",
					Description: `(Optional) Relation to class fileRemotePath. Cardinality - ONE_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_trig_rs_triggerable",
					Description: `(Optional) Relation to class trigTriggerable. Cardinality - ONE_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_config_rs_remote_path",
					Description: `(Optional) Relation to class fileRemotePath. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_config_rs_export_scheduler",
					Description: `(Optional) Relation to class trigSchedP. Cardinality - ONE_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Configuration Export Policy. ## Importing ## An existing Configuration Export Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_configuration_export_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_configuration_import_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI Configuration Import Policy`,
			Description:      ``,
			Keywords: []string{
				"configuration",
				"import",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object configuration_import_policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) admin state of the import`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object configuration_import_policy.`,
				},
				resource.Attribute{
					Name:        "fail_on_decrypt_errors",
					Description: `(Optional) fail_on_decrypt_errors for object configuration_import_policy.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Optional) import file name`,
				},
				resource.Attribute{
					Name:        "import_mode",
					Description: `(Optional) data import mode`,
				},
				resource.Attribute{
					Name:        "import_type",
					Description: `(Optional) data import type`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object configuration_import_policy.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional) snapshot for object configuration_import_policy.`,
				},
				resource.Attribute{
					Name:        "relation_config_rs_import_source",
					Description: `(Optional) Relation to class fileRemotePath. Cardinality - ONE_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_trig_rs_triggerable",
					Description: `(Optional) Relation to class trigTriggerable. Cardinality - ONE_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_config_rs_remote_path",
					Description: `(Optional) Relation to class fileRemotePath. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Configuration Import Policy. ## Importing ## An existing Configuration Import Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_configuration_import_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epg_to_contract",
			Category:         "Resources",
			ShortDescription: `Manages ACI EPG to contract relationship.`,
			Description:      ``,
			Keywords: []string{
				"epg",
				"to",
				"contract",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_epg_dn",
					Description: `(Required) Distinguished name of Parent epg.`,
				},
				resource.Attribute{
					Name:        "contract_dn",
					Description: `(Required) Distinguished name of contract to attach.`,
				},
				resource.Attribute{
					Name:        "contract_type",
					Description: `(Required) Type of relationship. Allowed values are ` + "`" + `consumer` + "`" + ` and ` + "`" + `provider` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) Provider matching criteria.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) Priority of relation object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the provider/consumer contract.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the provider/consumer contract.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vpc_explicit_protection_group",
			Category:         "Resources",
			ShortDescription: `Manages ACI VPC Explicit Protection Group`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"explicit",
				"protection",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vpc_explicit_protection_group.`,
				},
				resource.Attribute{
					Name:        "switch1",
					Description: `(Required) Id of switch 1 to attach.`,
				},
				resource.Attribute{
					Name:        "switch2",
					Description: `(Required) Id of switch 2 to attach.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vpc_explicit_protection_group.`,
				},
				resource.Attribute{
					Name:        "vpc_explicit_protection_group_id",
					Description: `(Optional) explicit protection group ID`,
				},
				resource.Attribute{
					Name:        "vpc_domain_policy",
					Description: `(Optional) VPC domain policy name. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the VPC Explicit Protection Group. ## Importing ## An existing VPC Explicit Protection Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vpc_explicit_protection_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_node_block_firmware",
			Category:         "Resources",
			ShortDescription: `Manages ACI Node Block`,
			Description:      ``,
			Keywords: []string{
				"node",
				"block",
				"firmware",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "firmware_group_dn",
					Description: `(Required) Distinguished name of parent FirmwareGroup object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object node_block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object node_block.`,
				},
				resource.Attribute{
					Name:        "from_",
					Description: `(Optional) from`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object node_block.`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) to ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Node Block. ## Importing ## An existing Node Block can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_node_block_firmware.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fabric_node_member",
			Category:         "Resources",
			ShortDescription: `Manages ACI Fabric Node Member`,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"node",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) serial of Object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Fabric Node member.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "ext_pool_id",
					Description: `(Optional) ext_pool_id for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "fabric_id",
					Description: `(Optional) place holder for a value`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `(Optional) node id`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Optional) node_type for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "pod_id",
					Description: `(Optional) pod id`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) system role type`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Optional) serial number ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Fabric Node Member. ## Importing ## An existing Fabric Node Member can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_fabric_node_member.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fc_domain",
			Category:         "Resources",
			ShortDescription: `Manages ACI FC Domain`,
			Description:      ``,
			Keywords: []string{
				"fc",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object fc_domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fc_domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object fc_domain.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns",
					Description: `(Optional) Relation to class fvnsVlanInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fc_rs_vsan_ns",
					Description: `(Optional) Relation to class fvnsVsanInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fc_rs_vsan_attr",
					Description: `(Optional) Relation to class fcVsanAttrP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vip_addr_ns",
					Description: `(Optional) Relation to class fvnsAddrInst. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_dom_vxlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fc_rs_vsan_attr_def",
					Description: `(Optional) Relation to class fcVsanAttrP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fc_rs_vsan_ns_def",
					Description: `(Optional) Relation to class fvnsAVsanInstP. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the FC Domain. ## Importing ## An existing FC Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_fc_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_app_profileaci_interface_fc_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI Interface FC Policy`,
			Description:      ``,
			Keywords: []string{
				"app",
				"profileaci",
				"interface",
				"fc",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object interface_fc_policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object interface_fc_policy.`,
				},
				resource.Attribute{
					Name:        "automaxspeed",
					Description: `(Optional) automaxspeed for object interface_fc_policy.`,
				},
				resource.Attribute{
					Name:        "fill_pattern",
					Description: `(Optional) Fill Pattern for native FC ports. Allowed values are "ARBFF" and "IDLE". Default is "IDLE".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object interface_fc_policy.`,
				},
				resource.Attribute{
					Name:        "port_mode",
					Description: `(Optional) In which mode Ports should be used. Allowed values are "f" and "np". Default is "f".`,
				},
				resource.Attribute{
					Name:        "rx_bb_credit",
					Description: `(Optional) Receive buffer credits for native FC ports Range:(161 - 641). Default value is 64.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) cpu or port speed. All the supported values are unknown, auto, 4G, 8G, 16G, 32G. Default value is auto.`,
				},
				resource.Attribute{
					Name:        "trunk_mode",
					Description: `(Optional) Trunking on/off for native FC ports. Allowed values are "un-init", "trunk-off", "trunk-on" and "auto".Default value is "trunk-off". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Interface FC Policy. ## Importing ## An existing Interface FC Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_interface_fc_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_firmware_group",
			Category:         "Resources",
			ShortDescription: `Manages ACI Firmware Group`,
			Description:      ``,
			Keywords: []string{
				"firmware",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object firmware_group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object firmware_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object firmware_group.`,
				},
				resource.Attribute{
					Name:        "firmware_group_type",
					Description: `(Optional) component type`,
				},
				resource.Attribute{
					Name:        "relation_firmware_rs_fwgrpp",
					Description: `(Optional) Relation to class firmwareFwP. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Firmware Group. ## Importing ## An existing Firmware Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_firmware_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_firmware_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI Firmware Policy`,
			Description:      ``,
			Keywords: []string{
				"firmware",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object firmware_policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object firmware_policy.`,
				},
				resource.Attribute{
					Name:        "effective_on_reboot",
					Description: `(Optional) firmware version effective on reboot selection`,
				},
				resource.Attribute{
					Name:        "ignore_compat",
					Description: `(Optional) whether compatibility check required`,
				},
				resource.Attribute{
					Name:        "internal_label",
					Description: `(Optional) firmware label`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object firmware_policy.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) firmware version`,
				},
				resource.Attribute{
					Name:        "version_check_override",
					Description: `(Optional) version check override ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Firmware Policy. ## Importing ## An existing Firmware Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_firmware_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_firmware_download_task",
			Category:         "Resources",
			ShortDescription: `Manages ACI Firmware Download Task`,
			Description:      ``,
			Keywords: []string{
				"firmware",
				"download",
				"task",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "auth_pass",
					Description: `(Optional) authentication type`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) ospf authentication type specifier`,
				},
				resource.Attribute{
					Name:        "dnld_task_flip",
					Description: `(Optional) dnld_task_flip for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "identity_private_key_contents",
					Description: `(Optional) identity_private_key_contents for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "identity_private_key_passphrase",
					Description: `(Optional) identity_private_key_passphrase for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "identity_public_key_contents",
					Description: `(Optional) identity_public_key_contents for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "load_catalog_if_exists_and_newer",
					Description: `(Optional) tracks to load the contained catalog or newer`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) password/key string`,
				},
				resource.Attribute{
					Name:        "polling_interval",
					Description: `(Optional) polling interval`,
				},
				resource.Attribute{
					Name:        "proto",
					Description: `(Optional) download protocol`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) URL of image of source`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) username for source ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Firmware Download Task. ## Importing ## An existing Firmware Download Task can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_firmware_download_task.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_application_epg",
			Category:         "Resources",
			ShortDescription: `Manages ACI Application EPG`,
			Description:      ``,
			Keywords: []string{
				"application",
				"epg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_profile_dn",
					Description: `(Required) Distinguished name of parent ApplicationProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object application_epg.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object application_epg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object application_epg. Range: "0" - "512" .`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings. Allowed values are "disabled" and "enabled". Default is "disabled".`,
				},
				resource.Attribute{
					Name:        "fwd_ctrl",
					Description: `(Optional) Forwarding control at EPG level. Allowed values are "none" and "proxy-arp". Default is "none".`,
				},
				resource.Attribute{
					Name:        "has_mcast_source",
					Description: `(Optional) If the source for the EPG is multicast or not. Allowed values are "yes" and "no". Default values is "no".`,
				},
				resource.Attribute{
					Name:        "is_attr_based_epg",
					Description: `(Optional) If the EPG is attribute based or not. Allowed values are "yes" and "no". Default is "yes".`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria for EPG. Allowed values are "All", "AtleastOne", "AtmostOne", "None". Default is "AtleastOne".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object application_epg.`,
				},
				resource.Attribute{
					Name:        "pc_enf_pref",
					Description: `(Optional) The preferred policy control. Allowed values are "unenforced" and "enforced". Default is "unenforced".`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication. Allowed values are "exclude" and "include". Default is "exclude".`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id. Allowed values are "unspecified", "level1", "level2", "level3", "level4", "level5" and "level6". Default is "unspecified.`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `(Optional) shutdown for object application_epg. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd",
					Description: `(Optional) Relation to class fvBD. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cust_qos_pol",
					Description: `(Optional) Relation to class qosCustomPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_dom_att",
					Description: `(Optional) Relation to class infraDomP. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_fc_path_att",
					Description: `(Optional) Relation to class fabricPathEp. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prov",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_graph_def",
					Description: `(Optional) Relation to class vzGraphCont. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons_if",
					Description: `(Optional) Relation to class vzCPIf. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_sec_inherited",
					Description: `(Optional) Relation to class fvEPg. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_node_att",
					Description: `(Optional) Relation to class fabricNode. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_dpp_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prov_def",
					Description: `(Optional) Relation to class vzCtrctEPgCont. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_trust_ctrl",
					Description: `(Optional) Relation to class fhsTrustCtrlPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prot_by",
					Description: `(Optional) Relation to class vzTaboo. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_aepg_mon_pol",
					Description: `(Optional) Relation to class monEPGPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_intra_epg",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Application EPG. ## Importing ## An existing Application EPG can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_application_epg.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_application_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI Application Profile`,
			Description:      ``,
			Keywords: []string{
				"application",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object application_profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object application_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object application_profile.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) priority class id. Allowed values are "unspecified", "level1", "level2", "level3", "level4", "level5" and "level6". Default is "unspecified.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_ap_mon_pol",
					Description: `(Optional) Relation to class monEPGPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Application Profile. ## Importing ## An existing Application Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_application_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bridge_domain",
			Category:         "Resources",
			ShortDescription: `Manages ACI Bridge Domain`,
			Description:      ``,
			Keywords: []string{
				"bridge",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object bridge_domain.`,
				},
				resource.Attribute{
					Name:        "optimize_wan_bandwidth",
					Description: `(Optional) Flag to enable OptimizeWanBandwidth between sites. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object bridge_domain.`,
				},
				resource.Attribute{
					Name:        "arp_flood",
					Description: `(Optional) A property to specify whether ARP flooding is enabled. If flooding is disabled, unicast routing will be performed on the target IP address. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "ep_clear",
					Description: `(Optional) Represents the parameter used by the node (i.e. Leaf) to clear all EPs in all leaves for this BD. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "ep_move_detect_mode",
					Description: `(Optional) The End Point move detection option uses the Gratuitous Address Resolution Protocol (GARP). A gratuitous ARP is an ARP broadcast-type of packet that is used to verify that no other device on the network has the same IP address as the sending device.`,
				},
				resource.Attribute{
					Name:        "host_based_routing",
					Description: `(Optional) enables advertising host routes out of l3outs of this BD. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "intersite_bum_traffic_allow",
					Description: `(Optional) Control whether BUM traffic is allowed between sites .Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "intersite_l2_stretch",
					Description: `(Optional) Flag to enable l2Stretch between sites. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "ip_learning",
					Description: `(Optional) Endpoint Dataplane Learning.Allowed values are "yes" and "no". Default is "yes".`,
				},
				resource.Attribute{
					Name:        "ipv6_mcast_allow",
					Description: `(Optional) Flag to indicate multicast IpV6 is allowed or not.Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "limit_ip_learn_to_subnets",
					Description: `(Optional) Limits IP address learning to the bridge domain subnets only. Every BD can have multiple subnets associated with it. By default, all IPs are learned. Allowed values are "yes" and "no". Default is "yes".`,
				},
				resource.Attribute{
					Name:        "ll_addr",
					Description: `(Optional) override of system generated ipv6 link-local address.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) The MAC address of the bridge domain (BD) or switched virtual interface (SVI). Every BD by default takes the fabric-wide default MAC address. You can override that address with a different one. By default the BD will take a 00:22:BD:F8:19:FF mac address.`,
				},
				resource.Attribute{
					Name:        "mcast_allow",
					Description: `(Optional) Flag to indicate if multicast is enabled for IpV4 addresses. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "multi_dst_pkt_act",
					Description: `(Optional) The multiple destination forwarding method for L2 Multicast, Broadcast, and Link Layer traffic types. Allowed values are "bd-flood", "encap-flood" and "drop". Default is "bd-flood".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object bridge_domain.`,
				},
				resource.Attribute{
					Name:        "bridge_domain_type",
					Description: `(Optional) The specific type of the object or component. Allowed values are "regular" and "fc". Default is "regular".`,
				},
				resource.Attribute{
					Name:        "unicast_route",
					Description: `(Optional) The forwarding method based on predefined forwarding criteria (IP or MAC address). Allowed values are "yes" and "no". Default is "yes".`,
				},
				resource.Attribute{
					Name:        "unk_mac_ucast_act",
					Description: `(Optional) The forwarding method for unknown layer 2 destinations. Allowed values are "flood" and "proxy". Default is "proxy".`,
				},
				resource.Attribute{
					Name:        "unk_mcast_act",
					Description: `(Optional) The parameter used by the node (i.e. a leaf) for forwarding data for an unknown multicast destination. Allowed values are "flood" and "opt-flood". Default is "flood".`,
				},
				resource.Attribute{
					Name:        "v6unk_mcast_act",
					Description: `(Optional) v6unk_mcast_act for object bridge_domain.`,
				},
				resource.Attribute{
					Name:        "vmac",
					Description: `(Optional) Virtual MAC address of the BD/SVI. This is used when the BD is extended to multiple sites using l2 Outside. Only allowed values is "not-applicable".`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_to_profile",
					Description: `(Optional) Relation to class rtctrlProfile. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_mldsn",
					Description: `(Optional) Relation to class mldSnoopPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_abd_pol_mon_pol",
					Description: `(Optional) Relation to class monEPGPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_to_nd_p",
					Description: `(Optional) Relation to class ndIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_flood_to",
					Description: `(Optional) Relation to class vzFilter. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_to_fhs",
					Description: `(Optional) Relation to class fhsBDPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_to_relay_p",
					Description: `(Optional) Relation to class dhcpRelayP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_ctx",
					Description: `(Optional) Relation to class fvCtx. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_to_netflow_monitor_pol",
					Description: `(Optional) Relation to class netflowMonitorPol. Cardinality - N_TO_M. Type - Set of Map.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_igmpsn",
					Description: `(Optional) Relation to class igmpSnoopPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_to_ep_ret",
					Description: `(Optional) Relation to class fvEpRetPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_to_out",
					Description: `(Optional) Relation to class l3extOut. Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Bridge Domain. ## Importing ## An existing Bridge Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_bridge_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vrf",
			Category:         "Resources",
			ShortDescription: `Manages ACI VRF`,
			Description:      ``,
			Keywords: []string{
				"vrf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vrf.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation tags for object vrf.`,
				},
				resource.Attribute{
					Name:        "bd_enforced_enable",
					Description: `(Optional) Flag to enable/disable bd_enforced for VRF.Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "ip_data_plane_learning",
					Description: `(Optional) Flag to enable/disable ip-data-plane learning for VRF. Allowed values are "enabled" and "disabled". Default is "enabled".`,
				},
				resource.Attribute{
					Name:        "knw_mcast_act",
					Description: `(Optional) specifies if known multicast traffic is forwarded or not. Allowed values are "permit" and "deny". Default is "permit".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vrf.`,
				},
				resource.Attribute{
					Name:        "pc_enf_dir",
					Description: `(Optional) Policy Control Enforcement Direction. It is used for defining policy enforcement direction for the traffic coming to or from an L3Out. Egress and Ingress directions are wrt L3Out. Default will be Ingress. But on the existing L3Outs during upgrade it will get set to Egress so that right after upgrade behavior doesn't change for them. This also means that there is no special upgrade sequence needed for upgrading to the release introducing this feature. After upgrade user would have to change the property value to Ingress. Once changed, system will reprogram the rules and prefix entry. Rules will get removed from the egress leaf and will get installed on the ingress leaf. Actrl prefix entry, if not already, will get installed on the ingress leaf. This feature will be ignored for the following cases: 1. Golf: Gets applied at Ingress by design. 2. Transit Rules get applied at Ingress by design. 4. vzAny 5. Taboo. Allowed values are "egress" and "ingress". Default is "ingress".`,
				},
				resource.Attribute{
					Name:        "pc_enf_pref",
					Description: `(Optional) Determines if the fabric should enforce contract policies to allow routing and packet forwarding. Allowed values are "enforced" and "unenforced". Default is "enforced".`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_ospf_ctx_pol",
					Description: `(Optional) Relation to class ospfCtxPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_vrf_validation_pol",
					Description: `(Optional) Relation to class l3extVrfValidationPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_ctx_mcast_to",
					Description: `(Optional) Relation to class vzFilter. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_ctx_to_eigrp_ctx_af_pol",
					Description: `(Optional) Relation to class eigrpCtxAfPol. Cardinality - N_TO_M. Type - Set of Map.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_ctx_to_ospf_ctx_pol",
					Description: `(Optional) Relation to class ospfCtxPol. Cardinality - N_TO_M. Type - Set of Map.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_ctx_to_ep_ret",
					Description: `(Optional) Relation to class fvEpRetPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bgp_ctx_pol",
					Description: `(Optional) Relation to class bgpCtxPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_ctx_mon_pol",
					Description: `(Optional) Relation to class monEPGPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_ctx_to_ext_route_tag_pol",
					Description: `(Optional) Relation to class l3extRouteTagPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_ctx_to_bgp_ctx_af_pol",
					Description: `(Optional) Relation to class bgpCtxAfPol. Cardinality - N_TO_M. Type - Set of Map. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the VRF. ## Importing ## An existing VRF can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vrf.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_end_point_retention_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI End Point Retention Policy`,
			Description:      ``,
			Keywords: []string{
				"end",
				"point",
				"retention",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object end_point_retention_policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object end_point_retention_policy.`,
				},
				resource.Attribute{
					Name:        "bounce_age_intvl",
					Description: `(Optional) The aging interval for a bounce entry. When an endpoint (VM) migrates to another switch, the endpoint is marked as bouncing for the specified aging interval and is deleted afterwards. Allowed value range is "0" - "0xffff". Default is "630".`,
				},
				resource.Attribute{
					Name:        "bounce_trig",
					Description: `(Optional) Specifies whether to install the bounce entry by RARP flood or by COOP protocol. Allowed values are "rarp-flood" and "protocol". Default is "protocol".`,
				},
				resource.Attribute{
					Name:        "hold_intvl",
					Description: `(Optional) A time period during which new endpoint learn events will not be honored. This interval is triggered when the maximum endpoint move frequency is exceeded. Allowed value range is "5" - "0xffff". Default is "300".`,
				},
				resource.Attribute{
					Name:        "local_ep_age_intvl",
					Description: `(Optional) The aging interval for all local endpoints learned in this bridge domain. When 75% of the interval is reached, 3 ARP requests are sent to verify the existence of the endpoint. If no response is received, the endpoint is deleted. Allowed value range is "120" - "0xffff". Default is "900". "0" is treated as special value here. Providing interval as "0" is treated as infinite interval.`,
				},
				resource.Attribute{
					Name:        "move_freq",
					Description: `(Optional) A maximum allowed number of endpoint moves per second. If the move frequency is exceeded, the hold interval is triggered, and new endpoint learn events will not be honored until after the hold interval expires. Allowed value range is "0" - "0xffff". Default is "256".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object end_point_retention_policy.`,
				},
				resource.Attribute{
					Name:        "remote_ep_age_intvl",
					Description: `(Optional) The aging interval for all remote endpoints learned in this bridge domain.Allowed value range is "120" - "0xffff". Default is "900". "0" is treated as special value here. Providing interval as "0" is treated as infinite interval. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the End Point Retention Policy. ## Importing ## An existing End Point Retention Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_end_point_retention_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ranges",
			Category:         "Resources",
			ShortDescription: `Manages ACI Ranges`,
			Description:      ``,
			Keywords: []string{
				"ranges",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vlan_pool_dn",
					Description: `(Required) Distinguished name of parent VLANPool object.`,
				},
				resource.Attribute{
					Name:        "_from",
					Description: `(Required) _from of Object ranges.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Required) to of Object ranges.`,
				},
				resource.Attribute{
					Name:        "alloc_mode",
					Description: `(Optional) alloc_mode for object ranges.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object ranges.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Optional) encapsulation block start`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object ranges.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) system role type`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Optional) encapsulation block end ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Ranges. ## Importing ## An existing Ranges can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_ranges.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vlan_pool",
			Category:         "Resources",
			ShortDescription: `Manages ACI VLAN Pool`,
			Description:      ``,
			Keywords: []string{
				"vlan",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vlan_pool.`,
				},
				resource.Attribute{
					Name:        "allocMode",
					Description: `(Required) allocMode of Object vlan_pool.`,
				},
				resource.Attribute{
					Name:        "alloc_mode",
					Description: `(Optional) allocation mode`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vlan_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vlan_pool. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the VLAN Pool. ## Importing ## An existing VLAN Pool can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vlan_pool.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vsan_pool",
			Category:         "Resources",
			ShortDescription: `Manages ACI VSAN Pool`,
			Description:      ``,
			Keywords: []string{
				"vsan",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vsan_pool.`,
				},
				resource.Attribute{
					Name:        "allocMode",
					Description: `(Required) allocMode of Object vsan_pool.`,
				},
				resource.Attribute{
					Name:        "alloc_mode",
					Description: `(Optional) alloc_mode for object vsan_pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vsan_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vsan_pool. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the VSAN Pool. ## Importing ## An existing VSAN Pool can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vsan_pool.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vxlan_pool",
			Category:         "Resources",
			ShortDescription: `Manages ACI VXLAN Pool`,
			Description:      ``,
			Keywords: []string{
				"vxlan",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vxlan_pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vxlan_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vxlan_pool. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the VXLAN Pool. ## Importing ## An existing VXLAN Pool can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vxlan_pool.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epg_to_domain",
			Category:         "Resources",
			ShortDescription: `Manages ACI epg to Domain`,
			Description:      ``,
			Keywords: []string{
				"epg",
				"to",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_epg_dn",
					Description: `(Required) Distinguished name of parent ApplicationEPG object.`,
				},
				resource.Attribute{
					Name:        "t_dn",
					Description: `(Required) vmm domain instance.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object domain.`,
				},
				resource.Attribute{
					Name:        "binding_type",
					Description: `(Optional) binding_type for object domain.`,
				},
				resource.Attribute{
					Name:        "class_pref",
					Description: `(Optional) class_pref for object domain.`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) delimiter for object domain.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) port encapsulation`,
				},
				resource.Attribute{
					Name:        "encap_mode",
					Description: `(Optional) encap_mode for object domain.`,
				},
				resource.Attribute{
					Name:        "epg_cos",
					Description: `(Optional) epg_cos for object domain.`,
				},
				resource.Attribute{
					Name:        "epg_cos_pref",
					Description: `(Optional) epg_cos_pref for object domain.`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) determines when policies are pushed to cam`,
				},
				resource.Attribute{
					Name:        "lag_policy_name",
					Description: `(Optional) lag_policy_name for object domain.`,
				},
				resource.Attribute{
					Name:        "netflow_dir",
					Description: `(Optional) netflow_dir for object domain.`,
				},
				resource.Attribute{
					Name:        "netflow_pref",
					Description: `(Optional) netflow_pref for object domain.`,
				},
				resource.Attribute{
					Name:        "num_ports",
					Description: `(Optional) number of ports existing operationally in module`,
				},
				resource.Attribute{
					Name:        "port_allocation",
					Description: `(Optional) port_allocation for object domain.`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) primary_encap for object domain.`,
				},
				resource.Attribute{
					Name:        "primary_encap_inner",
					Description: `(Optional) primary_encap_inner for object domain.`,
				},
				resource.Attribute{
					Name:        "res_imedcy",
					Description: `(Optional) policy resolution`,
				},
				resource.Attribute{
					Name:        "secondary_encap_inner",
					Description: `(Optional) secondary_encap_inner for object domain.`,
				},
				resource.Attribute{
					Name:        "switching_mode",
					Description: `(Optional) switching_mode for object domain.`,
				},
				resource.Attribute{
					Name:        "vmm_allow_promiscuous",
					Description: `(Optional) allow_promiscuous for object vmm_security_policy.`,
				},
				resource.Attribute{
					Name:        "vmm_forged_transmits",
					Description: `(Optional) forged_transmits for object vmm_security_policy.`,
				},
				resource.Attribute{
					Name:        "vmm_mac_changes",
					Description: `(Optional) mac_changes for object vmm_security_policy. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Domain.`,
				},
				resource.Attribute{
					Name:        "vmm_id",
					Description: `which is set to the Dn of the VMM Security Policy. ## Importing ## An existing Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vmm_id",
					Description: `which is set to the Dn of the VMM Security Policy. ## Importing ## An existing Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epg_to_static_path",
			Category:         "Resources",
			ShortDescription: `Manages ACI Static Path`,
			Description:      ``,
			Keywords: []string{
				"epg",
				"to",
				"static",
				"path",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_epg_dn",
					Description: `(Required) Distinguished name of parent ApplicationEPG object.`,
				},
				resource.Attribute{
					Name:        "tDn",
					Description: `(Required) tDn of Object static_path.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object static_path.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) encapsulation`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) immediacy`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) mode of the static association with the path`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) primary_encap for object static_path. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Static Path. ## Importing ## An existing Static Path can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_static_path.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_subnet",
			Category:         "Resources",
			ShortDescription: `Manages ACI Subnet`,
			Description:      ``,
			Keywords: []string{
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bridge_domain_dn",
					Description: `(Required) Distinguished name of parent BridgeDomain object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP address and mask of the default gateway.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object subnet.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) The subnet control state. The control can be specific protocols applied to the subnet such as IGMP Snooping. Allowed values are "unspecified", "querier", "nd" and "no-default-gateway". Default is "nd".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object subnet.`,
				},
				resource.Attribute{
					Name:        "preferred",
					Description: `(Optional) Indicates if the subnet is preferred (primary) over the available alternatives. Only one preferred subnet is allowed. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The network visibility of the subnet. Allowed values are "private", "public" and "shared". Default is "private".`,
				},
				resource.Attribute{
					Name:        "virtual",
					Description: `(Optional) Treated as virtual IP address. Used in case of BD extended to multiple sites. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_subnet_to_out",
					Description: `(Optional) Relation to class l3extOut. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_nd_pfx_pol",
					Description: `(Optional) Relation to class ndPfxPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_subnet_to_profile",
					Description: `(Optional) Relation to class rtctrlProfile. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Subnet. ## Importing ## An existing Subnet can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_subnet.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tenant",
			Category:         "Resources",
			ShortDescription: `Manages ACI Tenant`,
			Description:      ``,
			Keywords: []string{
				"tenant",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object tenant.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object tenant.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object tenant.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_tn_deny_rule",
					Description: `(Optional) Relation to class vzFilter. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_tenant_mon_pol",
					Description: `(Optional) Relation to class monEPGPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Tenant. ## Importing ## An existing Tenant can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_tenant.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_pcvpc_interface_policy_group",
			Category:         "Resources",
			ShortDescription: `Manages ACI PC/VPC Interface Policy Group`,
			Description:      ``,
			Keywords: []string{
				"pcvpc",
				"interface",
				"policy",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object pcvpc_interface_policy_group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object pcvpc_interface_policy_group.`,
				},
				resource.Attribute{
					Name:        "lag_t",
					Description: `(Optional) The bundled ports group link aggregation type: port channel vs virtual port channel. Allowed values are "not-aggregated", "node" and "link". Default is "link".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object pcvpc_interface_policy_group.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_span_v_src_grp",
					Description: `(Optional) Relation to class spanVSrcGrp. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_acc_bndl_grp_to_aggr_if",
					Description: `(Optional) Relation to class pcAggrIf. Cardinality - ONE_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_stormctrl_if_pol",
					Description: `(Optional) Relation to class stormctrlIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_lldp_if_pol",
					Description: `(Optional) Relation to class lldpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_macsec_if_pol",
					Description: `(Optional) Relation to class macsecIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_dpp_if_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_h_if_pol",
					Description: `(Optional) Relation to class fabricHIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_netflow_monitor_pol",
					Description: `(Optional) Relation to class netflowMonitorPol. Cardinality - N_TO_M. Type - Set of Map.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_port_auth_pol",
					Description: `(Optional) Relation to class l2PortAuthPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_mcp_if_pol",
					Description: `(Optional) Relation to class mcpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_port_security_pol",
					Description: `(Optional) Relation to class l2PortSecurityPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_copp_if_pol",
					Description: `(Optional) Relation to class coppIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_span_v_dest_grp",
					Description: `(Optional) Relation to class spanVDestGrp. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_lacp_pol",
					Description: `(Optional) Relation to class lacpLagPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_cdp_if_pol",
					Description: `(Optional) Relation to class cdpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_pfc_if_pol",
					Description: `(Optional) Relation to class qosPfcIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_sd_if_pol",
					Description: `(Optional) Relation to class qosSdIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_mon_if_infra_pol",
					Description: `(Optional) Relation to class monInfraPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_fc_if_pol",
					Description: `(Optional) Relation to class fcIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_ingress_dpp_if_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_egress_dpp_if_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_if_pol",
					Description: `(Optional) Relation to class l2IfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_stp_if_pol",
					Description: `(Optional) Relation to class stpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_att_ent_p",
					Description: `(Optional) Relation to class infraAttEntityP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_inst_pol",
					Description: `(Optional) Relation to class l2InstPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the PC/VPC Interface Policy Group. ## Importing ## An existing PC/VPC Interface Policy Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_pcvpc_interface_policy_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_access_port_policy_group",
			Category:         "Resources",
			ShortDescription: `Manages ACI Leaf Access Port Policy Group`,
			Description:      ``,
			Keywords: []string{
				"leaf",
				"access",
				"port",
				"policy",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object leaf_access_port_policy_group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object leaf_access_port_policy_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object leaf_access_port_policy_group.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_span_v_src_grp",
					Description: `(Optional) Relation to class spanVSrcGrp. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_stormctrl_if_pol",
					Description: `(Optional) Relation to class stormctrlIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_poe_if_pol",
					Description: `(Optional) Relation to class poeIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_lldp_if_pol",
					Description: `(Optional) Relation to class lldpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_macsec_if_pol",
					Description: `(Optional) Relation to class macsecIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_dpp_if_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_h_if_pol",
					Description: `(Optional) Relation to class fabricHIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_netflow_monitor_pol",
					Description: `(Optional) Relation to class netflowMonitorPol. Cardinality - N_TO_M. Type - Set of Map.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_port_auth_pol",
					Description: `(Optional) Relation to class l2PortAuthPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_mcp_if_pol",
					Description: `(Optional) Relation to class mcpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_port_security_pol",
					Description: `(Optional) Relation to class l2PortSecurityPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_copp_if_pol",
					Description: `(Optional) Relation to class coppIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_span_v_dest_grp",
					Description: `(Optional) Relation to class spanVDestGrp. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_dwdm_if_pol",
					Description: `(Optional) Relation to class dwdmIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_pfc_if_pol",
					Description: `(Optional) Relation to class qosPfcIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_sd_if_pol",
					Description: `(Optional) Relation to class qosSdIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_mon_if_infra_pol",
					Description: `(Optional) Relation to class monInfraPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_fc_if_pol",
					Description: `(Optional) Relation to class fcIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_ingress_dpp_if_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_cdp_if_pol",
					Description: `(Optional) Relation to class cdpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_if_pol",
					Description: `(Optional) Relation to class l2IfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_stp_if_pol",
					Description: `(Optional) Relation to class stpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_egress_dpp_if_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_att_ent_p",
					Description: `(Optional) Relation to class infraAttEntityP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_inst_pol",
					Description: `(Optional) Relation to class l2InstPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Leaf Access Port Policy Group. ## Importing ## An existing Leaf Access Port Policy Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_leaf_access_port_policy_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_interface_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI Leaf Interface Profile`,
			Description:      ``,
			Keywords: []string{
				"leaf",
				"interface",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object leaf_interface_profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object leaf_interface_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object leaf_interface_profile. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Leaf Interface Profile. ## Importing ## An existing Leaf Interface Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_leaf_interface_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_attachable_access_entity_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI Attachable Access Entity Profile`,
			Description:      ``,
			Keywords: []string{
				"attachable",
				"access",
				"entity",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object attachable_access_entity_profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object attachable_access_entity_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object attachable_access_entity_profile.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_dom_p",
					Description: `(Optional) Relation to class infraADomP. Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Attachable Access Entity Profile. ## Importing ## An existing Attachable Access Entity Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_attachable_access_entity_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_port_selector",
			Category:         "Resources",
			ShortDescription: `Manages ACI Access Port Selector`,
			Description:      ``,
			Keywords: []string{
				"access",
				"port",
				"selector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "leaf_interface_profile_dn",
					Description: `(Required) Distinguished name of parent LeafInterfaceProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object access_port_selector.`,
				},
				resource.Attribute{
					Name:        "access_port_selector_type",
					Description: `(Required) The host port selector type.Allowed values are "ALL" and "range". Default is "ALL".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_port_selector.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object access_port_selector.`,
				},
				resource.Attribute{
					Name:        "access_port_selector_type",
					Description: `(Optional) host port selector type`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_acc_base_grp",
					Description: `(Optional) Relation to class infraAccBaseGrp. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Access Port Selector. ## Importing ## An existing Access Port Selector can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_access_port_selector.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_switch_association",
			Category:         "Resources",
			ShortDescription: `Manages ACI Switch Association`,
			Description:      ``,
			Keywords: []string{
				"switch",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "leaf_profile_dn",
					Description: `(Required) Distinguished name of parent LeafProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object switch_association.`,
				},
				resource.Attribute{
					Name:        "switch_association_type",
					Description: `(Required) switch_association_type of Object switch_association.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object switch_association.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object switch_association.`,
				},
				resource.Attribute{
					Name:        "switch_association_type",
					Description: `(Optional) leaf selector type`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_acc_node_p_grp",
					Description: `(Optional) Relation to class infraAccNodePGrp. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Switch Association. ## Importing ## An existing Switch Association can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_switch_association.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_node_block",
			Category:         "Resources",
			ShortDescription: `Manages ACI Node Block`,
			Description:      ``,
			Keywords: []string{
				"node",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "switch_association_dn",
					Description: `(Required) Distinguished name of parent SwitchAssociation object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object node_block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object node_block.`,
				},
				resource.Attribute{
					Name:        "from_",
					Description: `(Optional) from Node ID. Range from 101 to 110`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object node_block.`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) to node ID. Range from 101 to 110 ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Node Block. ## Importing ## An existing Node Block can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_node_block.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI Leaf Profile`,
			Description:      ``,
			Keywords: []string{
				"leaf",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object leaf_profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object leaf_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object leaf_profile.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_acc_card_p",
					Description: `(Optional) Relation to class infraAccCardP. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_acc_port_p",
					Description: `(Optional) Relation to class infraAccPortP. Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Leaf Profile. ## Importing ## An existing Leaf Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_leaf_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_port_block",
			Category:         "Resources",
			ShortDescription: `Manages ACI Access Port Block`,
			Description:      ``,
			Keywords: []string{
				"access",
				"port",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_port_selector_dn",
					Description: `(Required) Distinguished name of parent AccessPortSelector object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object access_port_block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_port_block.`,
				},
				resource.Attribute{
					Name:        "from_card",
					Description: `(Optional) The beginning (from-range) of the card range block for the leaf access port block.`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Optional) The beginning (from-range) of the port range block for the leaf access port block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object access_port_block.`,
				},
				resource.Attribute{
					Name:        "to_card",
					Description: `(Optional) The end (to-range) of the card range block for the leaf access port block.`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Optional) The end (to-range) of the port range block for the leaf access port block.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_acc_bndl_subgrp",
					Description: `(Optional) Relation to class infraAccBndlSubgrp. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Access Port Block. ## Importing ## An existing Access Port Block can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_access_port_block.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vlan_encapsulationfor_vxlan_traffic",
			Category:         "Resources",
			ShortDescription: `Manages ACI Vlan Encapsulation for Vxlan Traffic`,
			Description:      ``,
			Keywords: []string{
				"vlan",
				"encapsulationfor",
				"vxlan",
				"traffic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "attachable_access_entity_profile_dn",
					Description: `(Required) Distinguished name of parent AttachableAccessEntityProfile object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vlan_encapsulationfor_vxlan_traffic.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vlan_encapsulationfor_vxlan_traffic.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_func_to_epg",
					Description: `(Optional) Relation to class fvEPg. Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Vlan Encapsulation for Vxlan Traffic. ## Importing ## An existing Vlan Encapsulation for Vxlan Traffic can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vlan_encapsulationfor_vxlan_traffic.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_sub_port_block",
			Category:         "Resources",
			ShortDescription: `Manages ACI Access Sub Port Block`,
			Description:      ``,
			Keywords: []string{
				"access",
				"sub",
				"port",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_port_selector_dn",
					Description: `(Required) Distinguished name of parent AccessPortSelector object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object access_sub_port_block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_sub_port_block.`,
				},
				resource.Attribute{
					Name:        "from_card",
					Description: `(Optional) from card`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Optional) port block from port`,
				},
				resource.Attribute{
					Name:        "from_sub_port",
					Description: `(Optional) from_sub_port for object access_sub_port_block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object access_sub_port_block.`,
				},
				resource.Attribute{
					Name:        "to_card",
					Description: `(Optional) to card`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Optional) to port`,
				},
				resource.Attribute{
					Name:        "to_sub_port",
					Description: `(Optional) to_sub_port for object access_sub_port_block. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Access Sub Port Block. ## Importing ## An existing Access Sub Port Block can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_access_sub_port_block.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l2_interface_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI L2 Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"l2",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object l2_interface_policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object l2_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object l2_interface_policy.`,
				},
				resource.Attribute{
					Name:        "qinq",
					Description: `(Optional) Determines if QinQ is disabled or if the port should be considered a core or edge port.Allowed values are "disabled", "edgePort", "corePort" and "doubleQtagPort". Default is "disabled".`,
				},
				resource.Attribute{
					Name:        "vepa",
					Description: `(Optional) Determines if Virtual Ethernet Port Aggregator is disabled or enabled. Allowed values are "disabled" and "enabled". Default is "disabled".`,
				},
				resource.Attribute{
					Name:        "vlan_scope",
					Description: `(Optional) The scope of the VLAN. Allowed values are "global" and "portlocal". Default is "global". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the L2 Interface Policy. ## Importing ## An existing L2 Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_l2_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_port_security_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI Port Security Policy`,
			Description:      ``,
			Keywords: []string{
				"port",
				"security",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object port_security_policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object port_security_policy.`,
				},
				resource.Attribute{
					Name:        "maximum",
					Description: `(Optional) Port Security Maximum. Allowed value range is "0" - "12000". Default is "0".`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) bgp domain mode`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object port_security_policy.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) amount of time between authentication attempts. Allowed value range is "60" - "3600". Default is "60".`,
				},
				resource.Attribute{
					Name:        "violation",
					Description: `(Optional) Port Security Violation. default value is "protect". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Port Security Policy. ## Importing ## An existing Port Security Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_port_security_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_domain_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI L3 Domain Profile`,
			Description:      ``,
			Keywords: []string{
				"l3",
				"domain",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object l3_domain_profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object l3_domain_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object l3_domain_profile.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns",
					Description: `(Optional) Relation to class fvnsVlanInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vip_addr_ns",
					Description: `(Optional) Relation to class fvnsAddrInst. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_extnw_rs_out",
					Description: `(Optional) Relation to class infraAccGrp. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_dom_vxlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the L3 Domain Profile. ## Importing ## An existing L3 Domain Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_l3_domain_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_external_network_instance_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI External Network Instance Profile`,
			Description:      ``,
			Keywords: []string{
				"external",
				"network",
				"instance",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "l3_outside_dn",
					Description: `(Required) Distinguished name of parent L3Outside object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object external_network_instance_profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object external_network_instance_profile.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object external_network_instance_profile. Allowed value range is "0" - "512".`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings. Allowed values are "disabled" and "enabled". Default value is "disabled".`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria. Allowed values are "All", "AtleastOne", "AtmostOne" and "None". Default is "AtleastOne".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object external_network_instance_profile.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication. Allowed values are "include" and "exclude". Default is "exclude".`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The QoS priority class identifier. Allowed values are "unspecified", "level1", "level2", "level3", "level4", "level5" and "level6". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile. Allowed values are "CS0", "CS1", "AF11", "AF12", "AF13", "CS2", "AF21", "AF22", "AF23", "CS3", "AF31", "AF32", "AF33", "CS4", "AF41", "AF42", "AF43", "CS5", "VA", "EF", "CS6", "CS7" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_sec_inherited",
					Description: `(Optional) Relation to class fvEPg. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prov",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_l3_inst_p_to_dom_p",
					Description: `(Optional) Relation to class extnwDomP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_inst_p_to_nat_mapping_epg",
					Description: `(Optional) Relation to class fvAEPg. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons_if",
					Description: `(Optional) Relation to class vzCPIf. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cust_qos_pol",
					Description: `(Optional) Relation to class qosCustomPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_inst_p_to_profile",
					Description: `(Optional) Relation to class rtctrlProfile. Cardinality - N_TO_M. Type - Set of Map.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prot_by",
					Description: `(Optional) Relation to class vzTaboo. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_intra_epg",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the External Network Instance Profile. ## Importing ## An existing External Network Instance Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_external_network_instance_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_interface_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI Logical Interface Profile`,
			Description:      ``,
			Keywords: []string{
				"logical",
				"interface",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_node_profile_dn",
					Description: `(Required) Distinguished name of parent LogicalNodeProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object logical_interface_profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object logical_interface_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object logical_interface_profile.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id. Allowed values are "unspecified", "level1", "level2", "level3", "level4", "level5" and "level6". Default is "unspecified.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Specifies the color of a policy label. Allowed values are "black", "navy", "dark-blue", "medium-blue", "blue", "dark-green", "green", "teal", "dark-cyan", "deep-sky-blue", "dark-turquoise", "medium-spring-green", "lime", "spring-green", "aqua", "cyan", "midnight-blue", "dodger-blue", "light-sea-green", "forest-green", "sea-green", "dark-slate-gray", "lime-green", "medium-sea-green", "turquoise", "royal-blue", "steel-blue", "dark-slate-blue", "medium-turquoise", "indigo", "dark-olive-green", "cadet-blue", "cornflower-blue", "medium-aquamarine", "dim-gray", "slate-blue", "olive-drab", "slate-gray", "light-slate-gray", "medium-slate-blue", "lawn-green", "chartreuse", "aquamarine", "maroon", "purple", "olive", "gray", "sky-blue", "light-sky-blue", "blue-violet", "dark-red", "dark-magenta", "saddle-brown", "dark-sea-green", "light-green", "medium-purple", "dark-violet", "pale-green", "dark-orchid", "yellow-green", "sienna", "brown", "dark-gray", "light-blue", "green-yellow", "pale-turquoise", "light-steel-blue", "powder-blue", "fire-brick", "dark-goldenrod", "medium-orchid", "rosy-brown", "dark-khaki", "silver", "medium-violet-red", "indian-red", "peru", "chocolate", "tan", "light-gray", "thistle", "orchid", "goldenrod", "pale-violet-red", "crimson", "gainsboro", "plum", "burlywood", "light-cyan", "lavender", "dark-salmon", "violet", "pale-goldenrod", "light-coral", "khaki", "alice-blue", "honeydew", "azure", "sandy-brown", "wheat", "beige", "white-smoke", "mint-cream", "ghost-white", "salmon", "antique-white", "linen", "light-goldenrod-yellow", "old-lace", "red", "fuchsia", "magenta", "deep-pink", "orange-red", "tomato", "hot-pink", "coral", "dark-orange", "light-salmon", "orange", "light-pink", "pink", "gold", "peachpuff", "navajo-white", "moccasin", "bisque", "misty-rose", "blanched-almond", "papaya-whip", "lavender-blush", "seashell", "cornsilk", "lemon-chiffon", "floral-white", "snow", "yellow", "light-yellow", "ivory" and "white". Default is "black".`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_l_if_p_to_netflow_monitor_pol",
					Description: `(Optional) Relation to class netflowMonitorPol. Cardinality - N_TO_M. Type - Set of Map.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_path_l3_out_att",
					Description: `(Optional) Relation to class fabricPathEp. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_egress_qos_dpp_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_ingress_qos_dpp_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_l_if_p_cust_qos_pol",
					Description: `(Optional) Relation to class qosCustomPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_arp_if_pol",
					Description: `(Optional) Relation to class arpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_nd_if_pol",
					Description: `(Optional) Relation to class ndIfPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Logical Interface Profile. ## Importing ## An existing Logical Interface Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_logical_interface_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_node_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI Logical Node Profile`,
			Description:      ``,
			Keywords: []string{
				"logical",
				"node",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "l3_outside_dn",
					Description: `(Required) Distinguished name of parent L3Outside object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object logical_node_profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object logical_node_profile.`,
				},
				resource.Attribute{
					Name:        "config_issues",
					Description: `(Optional) Bitmask representation of the configuration issues found during the endpoint group deployment. Allowed values are "none", "node-path-misconfig", "routerid-not-changable-with-mcast" and "loopback-ip-missing". Default is "none".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object logical_node_profile.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Specifies the color of a policy label. Allowed values are "black", "navy", "dark-blue", "medium-blue", "blue", "dark-green", "green", "teal", "dark-cyan", "deep-sky-blue", "dark-turquoise", "medium-spring-green", "lime", "spring-green", "aqua", "cyan", "midnight-blue", "dodger-blue", "light-sea-green", "forest-green", "sea-green", "dark-slate-gray", "lime-green", "medium-sea-green", "turquoise", "royal-blue", "steel-blue", "dark-slate-blue", "medium-turquoise", "indigo", "dark-olive-green", "cadet-blue", "cornflower-blue", "medium-aquamarine", "dim-gray", "slate-blue", "olive-drab", "slate-gray", "light-slate-gray", "medium-slate-blue", "lawn-green", "chartreuse", "aquamarine", "maroon", "purple", "olive", "gray", "sky-blue", "light-sky-blue", "blue-violet", "dark-red", "dark-magenta", "saddle-brown", "dark-sea-green", "light-green", "medium-purple", "dark-violet", "pale-green", "dark-orchid", "yellow-green", "sienna", "brown", "dark-gray", "light-blue", "green-yellow", "pale-turquoise", "light-steel-blue", "powder-blue", "fire-brick", "dark-goldenrod", "medium-orchid", "rosy-brown", "dark-khaki", "silver", "medium-violet-red", "indian-red", "peru", "chocolate", "tan", "light-gray", "thistle", "orchid", "goldenrod", "pale-violet-red", "crimson", "gainsboro", "plum", "burlywood", "light-cyan", "lavender", "dark-salmon", "violet", "pale-goldenrod", "light-coral", "khaki", "alice-blue", "honeydew", "azure", "sandy-brown", "wheat", "beige", "white-smoke", "mint-cream", "ghost-white", "salmon", "antique-white", "linen", "light-goldenrod-yellow", "old-lace", "red", "fuchsia", "magenta", "deep-pink", "orange-red", "tomato", "hot-pink", "coral", "dark-orange", "light-salmon", "orange", "light-pink", "pink", "gold", "peachpuff", "navajo-white", "moccasin", "bisque", "misty-rose", "blanched-almond", "papaya-whip", "lavender-blush", "seashell", "cornsilk", "lemon-chiffon", "floral-white", "snow", "yellow", "light-yellow", "ivory" and "white". Default is "black".`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) Node level Dscp value. Allowed values are "CS0", "CS1", "AF11", "AF12", "AF13", "CS2", "AF21", "AF22", "AF23", "CS3", "AF31", "AF32", "AF33", "CS4", "AF41", "AF42", "AF43", "CS5", "VA", "EF", "CS6", "CS7" and "unspecified". Default is "unspecified". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Logical Node Profile. ## Importing ## An existing Logical Node Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_logical_node_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_outside",
			Category:         "Resources",
			ShortDescription: `Manages ACI L3 Outside`,
			Description:      ``,
			Keywords: []string{
				"l3",
				"outside",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object l3_outside.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object l3_outside.`,
				},
				resource.Attribute{
					Name:        "enforce_rtctrl",
					Description: `(Optional) enforce route control type. Allowed values are "import" and "export". Default is "export".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object l3_outside.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile. Allowed values are "CS0", "CS1", "AF11", "AF12", "AF13", "CS2", "AF21", "AF22", "AF23", "CS3", "AF31", "AF32", "AF33", "CS4", "AF41", "AF42", "AF43", "CS5", "VA", "EF", "CS6", "CS7" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_dampening_pol",
					Description: `(Optional) Relation to class rtctrlProfile. Cardinality - N_TO_M. Type - Set of Map.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_ectx",
					Description: `(Optional) Relation to class fvCtx. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_out_to_bd_public_subnet_holder",
					Description: `(Optional) Relation to class fvBDPublicSubnetHolder. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_interleak_pol",
					Description: `(Optional) Relation to class rtctrlProfile. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_l3_dom_att",
					Description: `(Optional) Relation to class extnwDomP. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the L3 Outside. ## Importing ## An existing L3 Outside can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_l3_outside.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_node_to_fabric_node",
			Category:         "Resources",
			ShortDescription: `Manages ACI Fabric Node`,
			Description:      ``,
			Keywords: []string{
				"logical",
				"node",
				"to",
				"fabric",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_node_profile_dn",
					Description: `(Required) Distinguished name of parent LogicalNodeProfile object.`,
				},
				resource.Attribute{
					Name:        "tDn",
					Description: `(Required) tDn of Object fabric_node.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fabric_node.`,
				},
				resource.Attribute{
					Name:        "config_issues",
					Description: `(Optional) configuration issues`,
				},
				resource.Attribute{
					Name:        "rtr_id",
					Description: `(Optional) router identifier`,
				},
				resource.Attribute{
					Name:        "rtr_id_loop_back",
					Description: `(Optional) ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Fabric Node. ## Importing ## An existing Fabric Node can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_fabric_node.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_ext_subnet",
			Category:         "Resources",
			ShortDescription: `Manages ACI Subnet`,
			Description:      ``,
			Keywords: []string{
				"l3",
				"ext",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "external_network_instance_profile_dn",
					Description: `(Required) Distinguished name of parent ExternalNetworkInstanceProfile object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) ip of Object subnet.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) Aggregate Routes for Subnet. Allowed values are "import-rtctrl", "export-rtctrl" and "shared-rtctrl".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object subnet.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The domain applicable to the capability. Allowed values are "import-rtctrl", "export-rtctrl", "import-security", "shared-security" and "shared-rtctrl". Default is "import-security".`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_subnet_to_profile",
					Description: `(Optional) Relation to class rtctrlProfile. Cardinality - N_TO_M. Type - Set of Map.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_subnet_to_rt_summ",
					Description: `(Optional) Relation to class rtsumARtSummPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Subnet. ## Importing ## An existing Subnet can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_subnet.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_lacp_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI LACP Policy`,
			Description:      ``,
			Keywords: []string{
				"lacp",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object lacp_policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object lacp_policy.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) LAG control properties. Allowed values are "symmetric-hash", "susp-individual", "graceful-conv", "load-defer" and "fast-sel-hot-stdby".`,
				},
				resource.Attribute{
					Name:        "max_links",
					Description: `(Optional) maximum number of links. Allowed value range is "11" - "161". Default is "16".`,
				},
				resource.Attribute{
					Name:        "min_links",
					Description: `(Optional) minimum number of links in port channel. Allowed value range is "11" - "161". Default is "1".`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) policy mode. Allowed values are "off", "active", "passive", "mac-pin" and "mac-pin-nicload". Default is "off".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object lacp_policy. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the LACP Policy. ## Importing ## An existing LACP Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_lacp_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_lldp_interface_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI LLDP Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"lldp",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object lldp_interface_policy.`,
				},
				resource.Attribute{
					Name:        "admin_rx_st",
					Description: `(Optional) admin receive state. Allowed values are "enabled" and "disabled". Default value is "enabled".`,
				},
				resource.Attribute{
					Name:        "admin_tx_st",
					Description: `(Optional) admin transmit state. Allowed values are "enabled" and "disabled". Default value is "enabled".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object lldp_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object lldp_interface_policy. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the LLDP Interface Policy. ## Importing ## An existing LLDP Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_lldp_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_pod_maintenance_group",
			Category:         "Resources",
			ShortDescription: `Manages ACI POD Maintenance Group`,
			Description:      ``,
			Keywords: []string{
				"pod",
				"maintenance",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object pod_maintenance_group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object pod_maintenance_group.`,
				},
				resource.Attribute{
					Name:        "fwtype",
					Description: `(Optional) fwtype for object pod_maintenance_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object pod_maintenance_group.`,
				},
				resource.Attribute{
					Name:        "pod_maintenance_group_type",
					Description: `(Optional) component type`,
				},
				resource.Attribute{
					Name:        "relation_maint_rs_mgrpp",
					Description: `(Optional) Relation to class maintMaintP. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the POD Maintenance Group. ## Importing ## An existing POD Maintenance Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_pod_maintenance_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_maintenance_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI Maintenance Policy`,
			Description:      ``,
			Keywords: []string{
				"maintenance",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object maintenance_policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) maintenance policy admin state`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object maintenance_policy.`,
				},
				resource.Attribute{
					Name:        "graceful",
					Description: `(Optional) graceful for object maintenance_policy.`,
				},
				resource.Attribute{
					Name:        "ignore_compat",
					Description: `(Optional) whether compatibility check required`,
				},
				resource.Attribute{
					Name:        "internal_label",
					Description: `(Optional) firmware label`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object maintenance_policy.`,
				},
				resource.Attribute{
					Name:        "notif_cond",
					Description: `(Optional) when to send notifications to the admin`,
				},
				resource.Attribute{
					Name:        "run_mode",
					Description: `(Optional) maintenance policy run mode`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) compatibility catalog version`,
				},
				resource.Attribute{
					Name:        "version_check_override",
					Description: `(Optional) version check override`,
				},
				resource.Attribute{
					Name:        "relation_maint_rs_pol_scheduler",
					Description: `(Optional) Relation to class trigSchedP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_maint_rs_pol_notif",
					Description: `(Optional) Relation to class maintUserNotif. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_trig_rs_triggerable",
					Description: `(Optional) Relation to class trigTriggerable. Cardinality - ONE_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Maintenance Policy. ## Importing ## An existing Maintenance Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_maintenance_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_miscabling_protocol_interface_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI Mis-cabling Protocol Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"miscabling",
				"protocol",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object miscabling_protocol_interface_policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) administrative state of the object or policy. Allowed values are "enabled" and "disabled". Default is "enabled".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object miscabling_protocol_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object miscabling_protocol_interface_policy. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Mis-cabling Protocol Interface Policy. ## Importing ## An existing Mis-cabling Protocol Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_miscabling_protocol_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ospf_interface_policy",
			Category:         "Resources",
			ShortDescription: `Manages ACI OSPF Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"ospf",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object ospf_interface_policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object ospf_interface_policy.`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `(Optional) The OSPF cost for the interface. The cost (also called metric) of an interface in OSPF is an indication of the overhead required to send packets across a certain interface. The cost of an interface is inversely proportional to the bandwidth of that interface. A higher bandwidth indicates a lower cost. There is more overhead (higher cost) and time delays involved in crossing a 56k serial line than crossing a 10M ethernet line. The formula used to calculate the cost is: cost= 10000 0000/bandwidth in bps For example, it will cost 10 EXP8/10 EXP7 = 10 to cross a 10M Ethernet line and will cost 10 EXP8/1544000 = 64 to cross a T1 line. By default, the cost of an interface is calculated based on the bandwidth; you can force the cost of an interface with the ip ospf cost value interface sub-configuration mode command. Allowed value range is "0" - "65535". Default is "unspecified(0)".`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) interface policy controls. Allowed values are "unspecified", "passive", "mtu-ignore", "advert-subnet" and "bfd". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "dead_intvl",
					Description: `(Optional) The interval between hello packets from a neighbor before the router declares the neighbor as down. This value must be the same for all networking devices on a specific network. Specifying a smaller dead interval (seconds) will give faster detection of a neighbor being down and improve convergence, but might cause more routing instability. Allowed value range is "1" - "65535". Default value is "40".`,
				},
				resource.Attribute{
					Name:        "hello_intvl",
					Description: `(Optional) The interval between hello packets that OSPF sends on the interface. Note that the smaller the hello interval, the faster topological changes will be detected, but more routing traffic will ensue. This value must be the same for all routers and access servers on a specific network. Allowed value range is "1" - "65535". Default value is "10".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object ospf_interface_policy.`,
				},
				resource.Attribute{
					Name:        "nw_t",
					Description: `(Optional) The OSPF interface policy network type. OSPF supports point-to-point and broadcast. Allowed values are "unspecified", "p2p" and "bcast". Default value is "unspecified".`,
				},
				resource.Attribute{
					Name:        "pfx_suppress",
					Description: `(Optional) pfx-suppression for object ospf_interface_policy. Allowed values are "inherit", "enable" and "disable". Default value is "inherit".`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The OSPF interface priority used to determine the designated router (DR) on a specific network. The router with the highest OSPF priority on a segment will become the DR for that segment. The same process is repeated for the backup designated router (BDR). In the case of a tie, the router with the highest RID will win. The default for the interface OSPF priority is one. Remember that the DR and BDR concepts are per multiaccess segment. Allowed value range is "0" - "255". Default value is "1".`,
				},
				resource.Attribute{
					Name:        "rexmit_intvl",
					Description: `(Optional) The interval between LSA retransmissions. The retransmit interval occurs while the router is waiting for an acknowledgement from the neighbor router that it received the LSA. If no acknowlegment is received at the end of the interval, then the LSA is resent. Allowed value range is "1" - "65535". Default value is "5".`,
				},
				resource.Attribute{
					Name:        "xmit_delay",
					Description: `(Optional) The delay time needed to send an LSA update packet. OSPF increments the LSA age time by the transmit delay amount before transmitting the LSA update. You should take into account the transmission and propagation delays for the interface when you set this value. Allowed value range is "1" - "450". Default is "1". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the OSPF Interface Policy. ## Importing ## An existing OSPF Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_ospf_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_physical_domain",
			Category:         "Resources",
			ShortDescription: `Manages ACI Physical Domain`,
			Description:      ``,
			Keywords: []string{
				"physical",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object physical_domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object physical_domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object physical_domain.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns",
					Description: `(Optional) Relation to class fvnsVlanInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vip_addr_ns",
					Description: `(Optional) Relation to class fvnsAddrInst. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_dom_vxlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Physical Domain. ## Importing ## An existing Physical Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_physical_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_rest",
			Category:         "Resources",
			ShortDescription: `Manages ACI Model Objects via REST API calls. Any Model Object that is not supported by provider can be created/managed using this resource.`,
			Description:      ``,
			Keywords: []string{
				"rest",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) ACI path where object should be created. Starting with api/node/mo/{parent-dn}(if applicable)/{rn of object}.json`,
				},
				resource.Attribute{
					Name:        "class_name",
					Description: `(Optional) Which class object is being created. (Make sure there is no colon in the classname )`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Map of key-value pairs those needed to be passed to the Model object as parameters. Make sure the key name matches the name with the object parameter in ACI.`,
				},
				resource.Attribute{
					Name:        "payload",
					Description: `(Optional) Freestyle JSON or YAML payload which can directly be passed to the REST endpoint added in path. Either of content or payload is required.`,
				},
				resource.Attribute{
					Name:        "dn",
					Description: `(Optional) Distinguished name of object being managed. <strong>NOTE:</strong> We don't set the Status field explicitly, as it creates an issue with the relation objects. If you have requirement to pass the status field, pass it in the content. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the object created by it. ## Importing ## This resource does not support import.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_action_rule_profile",
			Category:         "Resources",
			ShortDescription: `Manages ACI Action Rule Profile`,
			Description:      ``,
			Keywords: []string{
				"action",
				"rule",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object action_rule_profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object action_rule_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object action_rule_profile. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Action Rule Profile. ## Importing ## An existing Action Rule Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_action_rule_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_span_destination_group",
			Category:         "Resources",
			ShortDescription: `Manages ACI SPAN Destination Group`,
			Description:      ``,
			Keywords: []string{
				"span",
				"destination",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object span_destination_group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the SPAN Destination Group. ## Importing ## An existing SPAN Destination Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_span_destination_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_span_sourcedestination_group_match_label",
			Category:         "Resources",
			ShortDescription: `Manages ACI SPAN Source-destination Group Match Label`,
			Description:      ``,
			Keywords: []string{
				"span",
				"sourcedestination",
				"group",
				"match",
				"label",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "span_source_group_dn",
					Description: `(Required) Distinguished name of parent SPANSourceGroup object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object span_sourcedestination_group_match_label.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) label color ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the SPAN Source-destination Group Match Label. ## Importing ## An existing SPAN Source-destination Group Match Label can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_span_sourcedestination_group_match_label.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_span_source_group",
			Category:         "Resources",
			ShortDescription: `Manages ACI SPAN Source Group`,
			Description:      ``,
			Keywords: []string{
				"span",
				"source",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object span_source_group.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) administrative state of the object or policy`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "relation_span_rs_src_grp_to_filter_grp",
					Description: `(Optional) Relation to class spanFilterGrp. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the SPAN Source Group. ## Importing ## An existing SPAN Source Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_span_source_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_trigger_scheduler",
			Category:         "Resources",
			ShortDescription: `Manages ACI Trigger Scheduler`,
			Description:      ``,
			Keywords: []string{
				"trigger",
				"scheduler",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object trigger_scheduler.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object trigger_scheduler.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object trigger_scheduler. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Trigger Scheduler. ## Importing ## An existing Trigger Scheduler can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_trigger_scheduler.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vmm_domain",
			Category:         "Resources",
			ShortDescription: `Manages ACI VMM Domain`,
			Description:      ``,
			Keywords: []string{
				"vmm",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "provider_profile_dn",
					Description: `(Required) Distinguished name of parent ProviderProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) access_mode for object vmm_domain. Allowed values are "read-write" and "read-only". Default is "read-write".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "arp_learning",
					Description: `(Optional) Enable/Disable arp learning for AVS Domain. Allowed values are "enabled" and "disabled". Default value is "disabled".`,
				},
				resource.Attribute{
					Name:        "ave_time_out",
					Description: `(Optional) ave_time_out for object vmm_domain. Allowed value range is "101" - "3001".`,
				},
				resource.Attribute{
					Name:        "config_infra_pg",
					Description: `(Optional) Flag to enable config_infra_pg for object vmm_domain. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "ctrl_knob",
					Description: `(Optional) Type pf control knob to use. Allowed values are "none" and "epDpVerify".`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) delimiter for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "enable_ave",
					Description: `(Optional) Flag to enable ave for object vmm_domain. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "enable_tag",
					Description: `(Optional) Flag enable tagging for object vmm_domain. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "encap_mode",
					Description: `(Optional) The layer 2 encapsulation protocol to use with the virtual switch. Allowed values are "unknown", "vlan" and "vxlan". Default is "unknown".`,
				},
				resource.Attribute{
					Name:        "enf_pref",
					Description: `(Optional) The switching enforcement preference. This determines whether switches can be done within the virtual switch (Local Switching) or whether all switched traffic must go through the fabric (No Local Switching). Allowed values are "hw", "sw" and "unknown". Default is "hw".`,
				},
				resource.Attribute{
					Name:        "ep_inventory_type",
					Description: `(Optional) Determines which end point inventory_type to use for object vmm_domain. Allowed values are "none" and "on-link". Default is "on-link".`,
				},
				resource.Attribute{
					Name:        "ep_ret_time",
					Description: `(Optional) end point retention time for object vmm_domain. Allowed value range is "1" - "6001". Default value is "0".`,
				},
				resource.Attribute{
					Name:        "hv_avail_monitor",
					Description: `(Optional) Flag to enable hv_avail_monitor for object vmm_domain. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "mcast_addr",
					Description: `(Optional) The multicast address of the VMM domain profile.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The switch to be used for the domain profile. Allowed values are "default", "n1kv", "unknown", "ovs", "k8s", "rhev", "cf" and "openshift". Default is "default".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "pref_encap_mode",
					Description: `(Optional) The preferred encapsulation mode for object vmm_domain. Allowed values are "unspecified", "vlan" and "vxlan". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "relation_vmm_rs_pref_enhanced_lag_pol",
					Description: `(Optional) Relation to class lacpEnhancedLagPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns",
					Description: `(Optional) Relation to class fvnsVlanInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vmm_rs_dom_mcast_addr_ns",
					Description: `(Optional) Relation to class fvnsMcastAddrInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vmm_rs_default_cdp_if_pol",
					Description: `(Optional) Relation to class cdpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vmm_rs_default_lacp_lag_pol",
					Description: `(Optional) Relation to class lacpLagPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vip_addr_ns",
					Description: `(Optional) Relation to class fvnsAddrInst. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vmm_rs_default_lldp_if_pol",
					Description: `(Optional) Relation to class lldpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vmm_rs_default_stp_if_pol",
					Description: `(Optional) Relation to class stpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_dom_vxlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vmm_rs_default_fw_pol",
					Description: `(Optional) Relation to class nwsFwPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vmm_rs_default_l2_inst_pol",
					Description: `(Optional) Relation to class l2InstPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the VMM Domain. ## Importing ## An existing VMM Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vmm_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_any",
			Category:         "Resources",
			ShortDescription: `Manages ACI Any`,
			Description:      ``,
			Keywords: []string{
				"any",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `(Required) Distinguished name of parent VRF object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object any.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) Represents the provider label match criteria. Allowed values are "All", "None", "AtmostOne" and "AtleastOne". Default value is "AtleastOne".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object any.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPgs can be divided in a the context can be divided in two subgroups. Allowed values are "disabled" and "enabled". Default is "disabled".`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_any_to_cons",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_any_to_cons_if",
					Description: `(Optional) Relation to class vzCPIf. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_any_to_prov",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Any. ## Importing ## An existing Any can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_any.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_contract",
			Category:         "Resources",
			ShortDescription: `Manages ACI Contract`,
			Description:      ``,
			Keywords: []string{
				"contract",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object contract.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object contract.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) priority level of the service contract. Allowed values are "unspecified", "level1", "level2", "level3", "level4", "level5" and "level6". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile. Allowed values are "global", "tenant", "application-profile" and "context". Default is "context".`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile. Allowed values are "CS0", "CS1", "AF11", "AF12", "AF13", "CS2", "AF21", "AF22", "AF23", "CS3", "AF31", "AF32", "AF33", "CS4", "AF41", "AF42", "AF43", "CS5", "VA", "EF", "CS6", "CS7" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_graph_att",
					Description: `(Optional) Relation to class vnsAbsGraph. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) to manage filters from the contract resource. It has the attributes like filter_name, annotation, description and name_alias.`,
				},
				resource.Attribute{
					Name:        "filter.filter_name",
					Description: `(Required) Name of the filter object.`,
				},
				resource.Attribute{
					Name:        "filter.description",
					Description: `(Optional) Description for the filter object.`,
				},
				resource.Attribute{
					Name:        "filter.annotation",
					Description: `(Optional) Annotation for filter object.`,
				},
				resource.Attribute{
					Name:        "filter.name_alias",
					Description: `(Optional) Name alias for filter object.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry",
					Description: `(Optional) to manage filter entries for particular filter from the contract resource. It has following attributes.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.filter_entry_name",
					Description: `(Required) name of Object filter_entry.`,
				},
				resource.Attribute{
					Name:        "filter.filterentry.entry_annotation",
					Description: `(Optional) annotation for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.entry_description",
					Description: `(Optional) Description for the filter entry.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.apply_to_frag",
					Description: `(Optional) Flag to determine whether to apply changes to fragment. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.arp_opc",
					Description: `(Optional) open peripheral codes. Allowed values are "unspecified", "req" and "reply". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.d_from_port",
					Description: `(Optional) Destination From Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.d_to_port",
					Description: `(Optional) Destination To Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.ether_t",
					Description: `(Optional) ether type for the entry. Allowed values are "unspecified", "ipv4", "trill", "arp", "ipv6", "mpls_ucast", "mac_security", "fcoe" and "ip". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.icmpv4_t",
					Description: `(Optional) ICMPv4 message type; used when ip_protocol is icmp. Allowed values are "echo-rep", "dst-unreach", "src-quench", "echo", "time-exceeded" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.icmpv6_t",
					Description: `(Optional) ICMPv6 message type; used when ip_protocol is icmpv6. Allowed values are "unspecified", "dst-unreach", "time-exceeded", "echo-req", "echo-rep", "nbr-solicit", "nbr-advert" and "redirect". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.match_dscp",
					Description: `(Optional) The matching differentiated services code point (DSCP) of the path attached to the layer 3 outside profile. Allowed values are "CS0", "CS1", "AF11", "AF12", "AF13", "CS2", "AF21", "AF22", "AF23", "CS3", "AF31", "AF32", "AF33", "CS4", "AF41", "AF42", "AF43", "CS5", "VA", "EF", "CS6", "CS7" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.entry_name_alias",
					Description: `(Optional) name_alias for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.prot",
					Description: `(Optional) level 3 ip protocol. Allowed values are "unspecified", "icmp", "igmp", "tcp", "egp", "igp", "udp", "icmpv6", "eigrp", "ospfigp", "pim" and "l2tp". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.s_from_port",
					Description: `(Optional) Source From Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.s_to_port",
					Description: `(Optional) Source To Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.stateful",
					Description: `(Optional) Determines if entry is stateful or not. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.tcp_rules",
					Description: `(Optional) TCP Session Rules. Allowed values are "unspecified", "est", "syn", "ack", "fin" and "rst". Default is "unspecified". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Contract.`,
				},
				resource.Attribute{
					Name:        "filter.id",
					Description: `exports this attribute for filter object. Set to the Dn for the filter managed by the contract.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.id",
					Description: `exports this attribute for filter entry object of filter object. Set to the Dn for the filter entry managed by the contract. ## Importing ## An existing Contract can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_contract.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "filter.id",
					Description: `exports this attribute for filter object. Set to the Dn for the filter managed by the contract.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.id",
					Description: `exports this attribute for filter entry object of filter object. Set to the Dn for the filter entry managed by the contract. ## Importing ## An existing Contract can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_contract.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_imported_contract",
			Category:         "Resources",
			ShortDescription: `Manages ACI Imported Contract`,
			Description:      ``,
			Keywords: []string{
				"imported",
				"contract",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object imported_contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object imported_contract.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object imported_contract.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_if",
					Description: `(Optional) Relation to class vzACtrct. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Imported Contract. ## Importing ## An existing Imported Contract can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_imported_contract.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_filter_entry",
			Category:         "Resources",
			ShortDescription: `Manages ACI Filter Entry`,
			Description:      ``,
			Keywords: []string{
				"filter",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter_dn",
					Description: `(Required) Distinguished name of parent Filter object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object filter_entry.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "apply_to_frag",
					Description: `(Optional) Flag to determine whether to apply changes to fragment. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "arp_opc",
					Description: `(Optional) open peripheral codes. Allowed values are "unspecified", "req" and "reply". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "d_from_port",
					Description: `(Optional) Destination From Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "d_to_port",
					Description: `(Optional) Destination To Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "ether_t",
					Description: `(Optional) ether type for the entry. Allowed values are "unspecified", "ipv4", "trill", "arp", "ipv6", "mpls_ucast", "mac_security", "fcoe" and "ip". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "icmpv4_t",
					Description: `(Optional) ICMPv4 message type; used when ip_protocol is icmp. Allowed values are "echo-rep", "dst-unreach", "src-quench", "echo", "time-exceeded" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "icmpv6_t",
					Description: `(Optional) ICMPv6 message type; used when ip_protocol is icmpv6. Allowed values are "unspecified", "dst-unreach", "time-exceeded", "echo-req", "echo-rep", "nbr-solicit", "nbr-advert" and "redirect". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "match_dscp",
					Description: `(Optional) The matching differentiated services code point (DSCP) of the path attached to the layer 3 outside profile. Allowed values are "CS0", "CS1", "AF11", "AF12", "AF13", "CS2", "AF21", "AF22", "AF23", "CS3", "AF31", "AF32", "AF33", "CS4", "AF41", "AF42", "AF43", "CS5", "VA", "EF", "CS6", "CS7" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "prot",
					Description: `(Optional) level 3 ip protocol. Allowed values are "unspecified", "icmp", "igmp", "tcp", "egp", "igp", "udp", "icmpv6", "eigrp", "ospfigp", "pim" and "l2tp". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "s_from_port",
					Description: `(Optional) Source From Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "s_to_port",
					Description: `(Optional) Source To Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Optional) Determines if entry is stateful or not. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "tcp_rules",
					Description: `(Optional) TCP Session Rules. Allowed values are "unspecified", "est", "syn", "ack", "fin" and "rst". Default is "unspecified". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Filter Entry. ## Importing ## An existing Filter Entry can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_filter_entry.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_filter",
			Category:         "Resources",
			ShortDescription: `Manages ACI Filter`,
			Description:      ``,
			Keywords: []string{
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object filter.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object filter.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object filter.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_filt_graph_att",
					Description: `(Optional) Relation to class vnsInTerm. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_fwd_r_flt_p_att",
					Description: `(Optional) Relation to class vzAFilterableUnit. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_rev_r_flt_p_att",
					Description: `(Optional) Relation to class vzAFilterableUnit. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Filter. ## Importing ## An existing Filter can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_filter.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_contract_subject",
			Category:         "Resources",
			ShortDescription: `Manages ACI Contract Subject`,
			Description:      ``,
			Keywords: []string{
				"contract",
				"subject",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "contract_dn",
					Description: `(Required) Distinguished name of parent Contract object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object contract_subject.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object contract_subject.`,
				},
				resource.Attribute{
					Name:        "cons_match_t",
					Description: `(Optional) The subject match criteria across consumers. Allowed values are "All", "None", "AtmostOne" and "AtleastOne". Default value is "AtleastOne".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object contract_subject.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The priority level of a sub application running behind an endpoint group, such as an Exchange server. Allowed values are "unspecified", "level1", "level2", "level3", "level4", "level5" and "level6". Default is "unspecified.`,
				},
				resource.Attribute{
					Name:        "prov_match_t",
					Description: `(Optional) The subject match criteria across consumers. Allowed values are "All", "None", "AtmostOne" and "AtleastOne". Default value is "AtleastOne".`,
				},
				resource.Attribute{
					Name:        "rev_flt_ports",
					Description: `(Optional) enables filter to apply on ingress and egress traffic. Allowed values are "yes" and "no". Default is "yes".`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile. Allowed values are "CS0", "CS1", "AF11", "AF12", "AF13", "CS2", "AF21", "AF22", "AF23", "CS3", "AF31", "AF32", "AF33", "CS4", "AF41", "AF42", "AF43", "CS5", "VA", "EF", "CS6", "CS7" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_subj_graph_att",
					Description: `(Optional) Relation to class vnsAbsGraph. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_sdwan_pol",
					Description: `(Optional) Relation to class extdevSDWanSlaPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_subj_filt_att",
					Description: `(Optional) Relation to class vzFilter. Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Contract Subject. ## Importing ## An existing Contract Subject can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_contract_subject.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_taboo_contract",
			Category:         "Resources",
			ShortDescription: `Manages ACI Taboo Contract`,
			Description:      ``,
			Keywords: []string{
				"taboo",
				"contract",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object taboo_contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object taboo_contract.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object taboo_contract. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Taboo Contract. ## Importing ## An existing Taboo Contract can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_taboo_contract.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"aci_local_user":                               0,
		"aci_x509_certificate":                         1,
		"aci_cdp_interface_policy":                     2,
		"aci_cloud_applicationcontainer":               3,
		"aci_cloud_aws_provider":                       4,
		"aci_autonomous_system_profile":                5,
		"aci_cloud_cidr_pool":                          6,
		"aci_cloud_context_profile":                    7,
		"aci_cloud_domain_profile":                     8,
		"aci_cloud_epg":                                9,
		"aci_cloud_endpoint_selector":                  10,
		"aci_cloud_external_epg":                       11,
		"aci_cloud_endpoint_selectorfor_external_epgs": 12,
		"aci_cloud_provider_profile":                   13,
		"aci_cloud_providers_region":                   14,
		"aci_cloud_subnet":                             15,
		"aci_cloud_availability_zone":                  16,
		"aci_configuration_export_policy":              17,
		"aci_configuration_import_policy":              18,
		"aci_epg_to_contract":                          19,
		"aci_vpc_explicit_protection_group":            20,
		"aci_node_block_firmware":                      21,
		"aci_fabric_node_member":                       22,
		"aci_fc_domain":                                23,
		"aci_app_profileaci_interface_fc_policy":       24,
		"aci_firmware_group":                           25,
		"aci_firmware_policy":                          26,
		"aci_firmware_download_task":                   27,
		"aci_application_epg":                          28,
		"aci_application_profile":                      29,
		"aci_bridge_domain":                            30,
		"aci_vrf":                                      31,
		"aci_end_point_retention_policy":               32,
		"aci_ranges":                                   33,
		"aci_vlan_pool":                                34,
		"aci_vsan_pool":                                35,
		"aci_vxlan_pool":                               36,
		"aci_epg_to_domain":                            37,
		"aci_epg_to_static_path":                       38,
		"aci_subnet":                                   39,
		"aci_tenant":                                   40,
		"aci_pcvpc_interface_policy_group":             41,
		"aci_leaf_access_port_policy_group":            42,
		"aci_leaf_interface_profile":                   43,
		"aci_attachable_access_entity_profile":         44,
		"aci_access_port_selector":                     45,
		"aci_switch_association":                       46,
		"aci_node_block":                               47,
		"aci_leaf_profile":                             48,
		"aci_access_port_block":                        49,
		"aci_vlan_encapsulationfor_vxlan_traffic":      50,
		"aci_access_sub_port_block":                    51,
		"aci_l2_interface_policy":                      52,
		"aci_port_security_policy":                     53,
		"aci_l3_domain_profile":                        54,
		"aci_external_network_instance_profile":        55,
		"aci_logical_interface_profile":                56,
		"aci_logical_node_profile":                     57,
		"aci_l3_outside":                               58,
		"aci_logical_node_to_fabric_node":              59,
		"aci_l3_ext_subnet":                            60,
		"aci_lacp_policy":                              61,
		"aci_lldp_interface_policy":                    62,
		"aci_pod_maintenance_group":                    63,
		"aci_maintenance_policy":                       64,
		"aci_miscabling_protocol_interface_policy":     65,
		"aci_ospf_interface_policy":                    66,
		"aci_physical_domain":                          67,
		"aci_rest":                                     68,
		"aci_action_rule_profile":                      69,
		"aci_span_destination_group":                   70,
		"aci_span_sourcedestination_group_match_label": 71,
		"aci_span_source_group":                        72,
		"aci_trigger_scheduler":                        73,
		"aci_vmm_domain":                               74,
		"aci_any":                                      75,
		"aci_contract":                                 76,
		"aci_imported_contract":                        77,
		"aci_filter_entry":                             78,
		"aci_filter":                                   79,
		"aci_contract_subject":                         80,
		"aci_taboo_contract":                           81,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
