package rancher2

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rancher2_app",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 app resource. This can be used to deploy apps within Rancher v2 projects.`,
			Description:      ``,
			Keywords: []string{
				"app",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Required) Catalog name of the app. If modified, app will be upgraded. For use scoped catalogs:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required/ForceNew) The name of the app (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required/ForceNew) The project id where the app will be installed (string)`,
				},
				resource.Attribute{
					Name:        "target_namespace",
					Description: `(Required/ForceNew) The namespace name where the app will be installed (string)`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template name of the app. If modified, app will be upgraded (string)`,
				},
				resource.Attribute{
					Name:        "answers",
					Description: `(Optional) Answers for the app template. If modified, app will be upgraded (map)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional/Computed) Description for the app (string)`,
				},
				resource.Attribute{
					Name:        "force_upgrade",
					Description: `(Optional) Force app upgrade (string)`,
				},
				resource.Attribute{
					Name:        "revision_id",
					Description: `(Optional/Computed) Current revision id for the app. If modified, If this argument is provided or modified, app will be rollbacked to ` + "`" + `revision_id` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `(Optional/Computed) Template version of the app. If modified, app will be upgraded. Default: ` + "`" + `latest` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "values_yaml",
					Description: `(Optional) values.yaml base64 encoded file content for the app template. If modified, app will be upgraded (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for App object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for App object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Computed) The url of the app template on a catalog (string) ## Timeouts ` + "`" + `rancher2_app` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating apps. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for app modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting apps. ## Import Apps can be imported using the app ID in the format ` + "`" + `<project_id>:<app_name>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_app.foo <project_id>:<app_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Computed) The url of the app template on a catalog (string) ## Timeouts ` + "`" + `rancher2_app` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating apps. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for app modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting apps. ## Import Apps can be imported using the app ID in the format ` + "`" + `<project_id>:<app_name>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_app.foo <project_id>:<app_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_auth_config_adfs",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Auth Config ADFS resource. This can be used to configure and enable Auth Config ADFS for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"config",
				"adfs",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name_field",
					Description: `(Required) ADFS display name field (string)`,
				},
				resource.Attribute{
					Name:        "groups_field",
					Description: `(Required) ADFS group field (string)`,
				},
				resource.Attribute{
					Name:        "idp_metadata_content",
					Description: `(Required/Sensitive) ADFS IDP metadata content (string)`,
				},
				resource.Attribute{
					Name:        "rancher_api_host",
					Description: `(Required) Rancher url. Schema needs to be specified, ` + "`" + `https://<RANCHER_API_HOST>` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "sp_cert",
					Description: `(Required/Sensitive) ADFS SP cert (string)`,
				},
				resource.Attribute{
					Name:        "sp_key",
					Description: `(Required/Sensitive) ADFS SP key (string)`,
				},
				resource.Attribute{
					Name:        "uid_field",
					Description: `(Required) ADFS UID field (string)`,
				},
				resource.Attribute{
					Name:        "user_name_field",
					Description: `(Required) ADFS user name field (string)`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) Access mode for auth. ` + "`" + `required` + "`" + `, ` + "`" + `restricted` + "`" + `, ` + "`" + `unrestricted` + "`" + ` are supported. Default ` + "`" + `unrestricted` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "allowed_principal_ids",
					Description: `(Optional) Allowed principal ids for auth. Required if ` + "`" + `access_mode` + "`" + ` is ` + "`" + `required` + "`" + ` or ` + "`" + `restricted` + "`" + `. Ex: ` + "`" + `adfs_user://<USER_ID>` + "`" + ` ` + "`" + `adfs_group://<GROUP_ID>` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable auth config provider. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_auth_config_activedirectory",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Auth Config ActiveDirectory resource. This can be used to configure and enable Auth Config ActiveDirectory for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"config",
				"activedirectory",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "servers",
					Description: `(Required) ActiveDirectory servers list (list)`,
				},
				resource.Attribute{
					Name:        "service_account_username",
					Description: `(Required/Sensitive) Service account DN for access ActiveDirectory service (string)`,
				},
				resource.Attribute{
					Name:        "service_account_password",
					Description: `(Required/Sensitive) Service account password for access ActiveDirectory service (string)`,
				},
				resource.Attribute{
					Name:        "user_search_base",
					Description: `(Required) User search base DN (string)`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) Access mode for auth. ` + "`" + `required` + "`" + `, ` + "`" + `restricted` + "`" + `, ` + "`" + `unrestricted` + "`" + ` are supported. Default ` + "`" + `unrestricted` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "allowed_principal_ids",
					Description: `(Optional) Allowed principal ids for auth. Required if ` + "`" + `access_mode` + "`" + ` is ` + "`" + `required` + "`" + ` or ` + "`" + `restricted` + "`" + `. Ex: ` + "`" + `activedirectory_user://<DN>` + "`" + ` ` + "`" + `activedirectory_group://<DN>` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) CA certificate for TLS if selfsigned (string)`,
				},
				resource.Attribute{
					Name:        "connection_timeout",
					Description: `(Optional) ActiveDirectory connection timeout. Default ` + "`" + `5000` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "default_login_domain",
					Description: `(Optional) ActiveDirectory defult lgoin domain (string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable auth config provider. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "group_dn_attribute",
					Description: `(Optional/Computed) Group DN attribute. Default ` + "`" + `distinguishedName` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_member_mapping_attribute",
					Description: `(Optional/Computed) Group member mapping attribute. Default ` + "`" + `member` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_member_user_attribute",
					Description: `(Optional/Computed) Group member user attribute. Default ` + "`" + `distinguishedName` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_name_attribute",
					Description: `(Optional/Computed) Group name attribute. Default ` + "`" + `name` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_object_class",
					Description: `(Optional/Computed) Group object class. Default ` + "`" + `group` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_search_attribute",
					Description: `(Optional/Computed) Group search attribute. Default ` + "`" + `sAMAccountName` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_search_base",
					Description: `(Optional/Computed) Group search base (string)`,
				},
				resource.Attribute{
					Name:        "group_search_filter",
					Description: `(Optional/Computed) Group search filter (string)`,
				},
				resource.Attribute{
					Name:        "nested_group_membership_enabled",
					Description: `(Optional/Computed) Nested group membership enable. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) ActiveDirectory port. Default ` + "`" + `389` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "user_disabled_bit_mask",
					Description: `(Optional) User disabled bit mask. Default ` + "`" + `2` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "user_enabled_attribute",
					Description: `(Optional/Computed) User enable attribute (string)`,
				},
				resource.Attribute{
					Name:        "user_login_attribute",
					Description: `(Optional/Computed) User login attribute. Default ` + "`" + `sAMAccountName` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_name_attribute",
					Description: `(Optional/Computed) User name attribute. Default ` + "`" + `name` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_object_class",
					Description: `(Optional/Computed) User object class. Default ` + "`" + `person` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_search_attribute",
					Description: `(Optional/Computed) User search attribute. Default ` + "`" + `sAMAccountName|sn|givenName` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_search_filter",
					Description: `(Optional/Computed) User search filter (string)`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional/Computed) Enable TLS connection (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_auth_config_azuread",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Auth Config AzureAD resource. This can be used to configure and enable Auth Config AzureAD for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"config",
				"azuread",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required/Sensitive) AzureAD auth application ID (string)`,
				},
				resource.Attribute{
					Name:        "application_secret",
					Description: `(Required/Sensitive) AzureAD auth application secret (string)`,
				},
				resource.Attribute{
					Name:        "auth_endpoint",
					Description: `(Required) AzureAD auth endpoint (string)`,
				},
				resource.Attribute{
					Name:        "graph_endpoint",
					Description: `(Required) AzureAD graph endpoint (string)`,
				},
				resource.Attribute{
					Name:        "rancher_url",
					Description: `(Required) Rancher URL (string). "<rancher_url>/verify-auth-azure"`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) AzureAD tenant ID (string)`,
				},
				resource.Attribute{
					Name:        "token_endpoint",
					Description: `(Required) AzureAD token endpoint (string)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) AzureAD endpoint. Default ` + "`" + `https://login.microsoftonline.com/` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) Access mode for auth. ` + "`" + `required` + "`" + `, ` + "`" + `restricted` + "`" + `, ` + "`" + `unrestricted` + "`" + ` are supported. Default ` + "`" + `unrestricted` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "allowed_principal_ids",
					Description: `(Optional) Allowed principal ids for auth. Required if ` + "`" + `access_mode` + "`" + ` is ` + "`" + `required` + "`" + ` or ` + "`" + `restricted` + "`" + `. Ex: ` + "`" + `azuread_user://<USER_ID>` + "`" + ` ` + "`" + `azuread_group://<GROUP_ID>` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable auth config provider. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) Enable TLS connection. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_auth_config_freeipa",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Auth Config FreeIpa resource. This can be used to configure and enable Auth Config FreeIpa for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"config",
				"freeipa",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "servers",
					Description: `(Required) FreeIpa servers list (list)`,
				},
				resource.Attribute{
					Name:        "service_account_distinguished_name",
					Description: `(Required/Sensitive) Service account DN for access FreeIpa service (string)`,
				},
				resource.Attribute{
					Name:        "service_account_password",
					Description: `(Required/Sensitive) Service account password for access FreeIpa service (string)`,
				},
				resource.Attribute{
					Name:        "user_search_base",
					Description: `(Required) User search base DN (string)`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) Access mode for auth. ` + "`" + `required` + "`" + `, ` + "`" + `restricted` + "`" + `, ` + "`" + `unrestricted` + "`" + ` are supported. Default ` + "`" + `unrestricted` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "allowed_principal_ids",
					Description: `(Optional) Allowed principal ids for auth. Required if ` + "`" + `access_mode` + "`" + ` is ` + "`" + `required` + "`" + ` or ` + "`" + `restricted` + "`" + `. Ex: ` + "`" + `freeipa_user://<DN>` + "`" + ` ` + "`" + `freeipa_group://<DN>` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) Base64 encoded CA certificate for TLS if self-signed. Use filebase64(<FILE>) for encoding file (string)`,
				},
				resource.Attribute{
					Name:        "connection_timeout",
					Description: `(Optional) FreeIpa connection timeout. Default ` + "`" + `5000` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable auth config provider. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "group_dn_attribute",
					Description: `(Optional/Computed) Group DN attribute. Default ` + "`" + `entryDN` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_member_mapping_attribute",
					Description: `(Optional/Computed) Group member mapping attribute. Default ` + "`" + `member` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_member_user_attribute",
					Description: `(Optional/Computed) Group member user attribute. Default ` + "`" + `entryDN` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_name_attribute",
					Description: `(Optional/Computed) Group name attribute. Default ` + "`" + `cn` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_object_class",
					Description: `(Optional/Computed) Group object class. Default ` + "`" + `groupOfNames` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_search_attribute",
					Description: `(Optional/Computed) Group search attribute. Default ` + "`" + `cn` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_search_base",
					Description: `(Optional/Computed) Group search base (string)`,
				},
				resource.Attribute{
					Name:        "nested_group_membership_enabled",
					Description: `(Optional/Computed) Nested group membership enable. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) FreeIpa port. Default ` + "`" + `389` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "user_disabled_bit_mask",
					Description: `(Optional/Computed) User disabled bit mask (int)`,
				},
				resource.Attribute{
					Name:        "user_enabled_attribute",
					Description: `(Optional/Computed) User enable attribute (string)`,
				},
				resource.Attribute{
					Name:        "user_login_attribute",
					Description: `(Optional/Computed) User login attribute. Default ` + "`" + `uid` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_member_attribute",
					Description: `(Optional/Computed) User member attribute. Default ` + "`" + `memberOf` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_name_attribute",
					Description: `(Optional/Computed) User name attribute. Default ` + "`" + `givenName` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_object_class",
					Description: `(Optional/Computed) User object class. Default ` + "`" + `inetorgperson` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_search_attribute",
					Description: `(Optional/Computed) User search attribute. Default ` + "`" + `uid|sn|givenName` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional/Computed) Enable TLS connection (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_auth_config_github",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Auth Config Github resource. This can be used to configure and enable Auth Config Github for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"config",
				"github",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required/Sensitive) Github auth Client ID (string)`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Required/Sensitive) Github auth Client secret (string)`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Github hostname to connect. Default ` + "`" + `github.com` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) Access mode for auth. ` + "`" + `required` + "`" + `, ` + "`" + `restricted` + "`" + `, ` + "`" + `unrestricted` + "`" + ` are supported. Default ` + "`" + `unrestricted` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "allowed_principal_ids",
					Description: `(Optional) Allowed principal ids for auth. Required if ` + "`" + `access_mode` + "`" + ` is ` + "`" + `required` + "`" + ` or ` + "`" + `restricted` + "`" + `. Ex: ` + "`" + `github_user://<USER_ID>` + "`" + ` ` + "`" + `github_team://<GROUP_ID>` + "`" + ` ` + "`" + `github_org://<ORG_ID>` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable auth config provider. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) Enable TLS connection. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_auth_config_keycloak",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Auth Config KeyCloak resource. This can be used to configure and enable Auth Config KeyCloak for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"config",
				"keycloak",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name_field",
					Description: `(Required) KeyCloak display name field (string)`,
				},
				resource.Attribute{
					Name:        "groups_field",
					Description: `(Required) KeyCloak group field (string)`,
				},
				resource.Attribute{
					Name:        "idp_metadata_content",
					Description: `(Required/Sensitive) KeyCloak IDP metadata content (string)`,
				},
				resource.Attribute{
					Name:        "rancher_api_host",
					Description: `(Required) Rancher url. Schema needs to be specified, ` + "`" + `https://<RANCHER_API_HOST>` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "sp_cert",
					Description: `(Required/Sensitive) KeyCloak SP cert (string)`,
				},
				resource.Attribute{
					Name:        "sp_key",
					Description: `(Required/Sensitive) KeyCloak SP key (string)`,
				},
				resource.Attribute{
					Name:        "uid_field",
					Description: `(Required) KeyCloak UID field (string)`,
				},
				resource.Attribute{
					Name:        "user_name_field",
					Description: `(Required) KeyCloak user name field (string)`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) Access mode for auth. ` + "`" + `required` + "`" + `, ` + "`" + `restricted` + "`" + `, ` + "`" + `unrestricted` + "`" + ` are supported. Default ` + "`" + `unrestricted` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "allowed_principal_ids",
					Description: `(Optional) Allowed principal ids for auth. Required if ` + "`" + `access_mode` + "`" + ` is ` + "`" + `required` + "`" + ` or ` + "`" + `restricted` + "`" + `. Ex: ` + "`" + `keycloak_user://<USER_ID>` + "`" + ` ` + "`" + `keycloak_group://<GROUP_ID>` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable auth config provider. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_auth_config_okta",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Auth Config OKTA resource. This can be used to configure and enable Auth Config OKTA for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"config",
				"okta",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name_field",
					Description: `(Required) OKTA display name field (string)`,
				},
				resource.Attribute{
					Name:        "groups_field",
					Description: `(Required) OKTA group field (string)`,
				},
				resource.Attribute{
					Name:        "idp_metadata_content",
					Description: `(Required/Sensitive) OKTA IDP metadata content (string)`,
				},
				resource.Attribute{
					Name:        "rancher_api_host",
					Description: `(Required) Rancher url. Schema needs to be specified, ` + "`" + `https://<RANCHER_API_HOST>` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "sp_cert",
					Description: `(Required/Sensitive) OKTA SP cert (string)`,
				},
				resource.Attribute{
					Name:        "sp_key",
					Description: `(Required/Sensitive) OKTA SP key (string)`,
				},
				resource.Attribute{
					Name:        "uid_field",
					Description: `(Required) OKTA UID field (string)`,
				},
				resource.Attribute{
					Name:        "user_name_field",
					Description: `(Required) OKTA user name field (string)`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) Access mode for auth. ` + "`" + `required` + "`" + `, ` + "`" + `restricted` + "`" + `, ` + "`" + `unrestricted` + "`" + ` are supported. Default ` + "`" + `unrestricted` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "allowed_principal_ids",
					Description: `(Optional) Allowed principal ids for auth. Required if ` + "`" + `access_mode` + "`" + ` is ` + "`" + `required` + "`" + ` or ` + "`" + `restricted` + "`" + `. Ex: ` + "`" + `okta_user://<USER_ID>` + "`" + ` ` + "`" + `okta_group://<GROUP_ID>` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable auth config provider. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_auth_config_openldap",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Auth Config OpenLdap resource. This can be used to configure and enable Auth Config OpenLdap for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"config",
				"openldap",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "servers",
					Description: `(Required) OpenLdap servers list (list)`,
				},
				resource.Attribute{
					Name:        "service_account_distinguished_name",
					Description: `(Required/Sensitive) Service account DN for access OpenLdap service (string)`,
				},
				resource.Attribute{
					Name:        "service_account_password",
					Description: `(Required/Sensitive) Service account password for access OpenLdap service (string)`,
				},
				resource.Attribute{
					Name:        "user_search_base",
					Description: `(Required) User search base DN (string)`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) Access mode for auth. ` + "`" + `required` + "`" + `, ` + "`" + `restricted` + "`" + `, ` + "`" + `unrestricted` + "`" + ` are supported. Default ` + "`" + `unrestricted` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "allowed_principal_ids",
					Description: `(Optional) Allowed principal ids for auth. Required if ` + "`" + `access_mode` + "`" + ` is ` + "`" + `required` + "`" + ` or ` + "`" + `restricted` + "`" + `. Ex: ` + "`" + `openldap_user://<DN>` + "`" + ` ` + "`" + `openldap_group://<DN>` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) Base64 encoded CA certificate for TLS if self-signed. Use filebase64(<FILE>) for encoding file (string)`,
				},
				resource.Attribute{
					Name:        "connection_timeout",
					Description: `(Optional) OpenLdap connection timeout. Default ` + "`" + `5000` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable auth config provider. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "group_dn_attribute",
					Description: `(Optional/Computed) Group DN attribute. Default ` + "`" + `entryDN` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_member_mapping_attribute",
					Description: `(Optional/Computed) Group member mapping attribute. Default ` + "`" + `member` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_member_user_attribute",
					Description: `(Optional/Computed) Group member user attribute. Default ` + "`" + `entryDN` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_name_attribute",
					Description: `(Optional/Computed) Group name attribute. Default ` + "`" + `cn` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_object_class",
					Description: `(Optional/Computed) Group object class. Default ` + "`" + `groupOfNames` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_search_attribute",
					Description: `(Optional/Computed) Group search attribute. Default ` + "`" + `cn` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_search_base",
					Description: `(Optional/Computed) Group search base (string)`,
				},
				resource.Attribute{
					Name:        "nested_group_membership_enabled",
					Description: `(Optional/Computed) Nested group membership enable. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) OpenLdap port. Default ` + "`" + `389` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "user_disabled_bit_mask",
					Description: `(Optional/Computed) User disabled bit mask (int)`,
				},
				resource.Attribute{
					Name:        "user_enabled_attribute",
					Description: `(Optional/Computed) User enable attribute (string)`,
				},
				resource.Attribute{
					Name:        "user_login_attribute",
					Description: `(Optional/Computed) User login attribute. Default ` + "`" + `uid` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_member_attribute",
					Description: `(Optional/Computed) User member attribute. Default ` + "`" + `memberOf` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_name_attribute",
					Description: `(Optional/Computed) User name attribute. Default ` + "`" + `givenName` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_object_class",
					Description: `(Optional/Computed) User object class. Default ` + "`" + `inetorgperson` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "user_search_attribute",
					Description: `(Optional/Computed) User search attribute. Default ` + "`" + `uid|sn|givenName` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional/Computed) Enable TLS connection (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_auth_config_ping",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Auth Config Ping resource. This can be used to configure and enable Auth Config Ping for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"config",
				"ping",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name_field",
					Description: `(Required) Ping display name field (string)`,
				},
				resource.Attribute{
					Name:        "groups_field",
					Description: `(Required) Ping group field (string)`,
				},
				resource.Attribute{
					Name:        "idp_metadata_content",
					Description: `(Required/Sensitive) Ping IDP metadata content (string)`,
				},
				resource.Attribute{
					Name:        "rancher_api_host",
					Description: `(Required) Rancher url. Schema needs to be specified, ` + "`" + `https://<RANCHER_API_HOST>` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "sp_cert",
					Description: `(Required/Sensitive) Ping SP cert (string)`,
				},
				resource.Attribute{
					Name:        "sp_key",
					Description: `(Required/Sensitive) Ping SP key (string)`,
				},
				resource.Attribute{
					Name:        "uid_field",
					Description: `(Required) Ping UID field (string)`,
				},
				resource.Attribute{
					Name:        "user_name_field",
					Description: `(Required) Ping user name field (string)`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) Access mode for auth. ` + "`" + `required` + "`" + `, ` + "`" + `restricted` + "`" + `, ` + "`" + `unrestricted` + "`" + ` are supported. Default ` + "`" + `unrestricted` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "allowed_principal_ids",
					Description: `(Optional) Allowed principal ids for auth. Required if ` + "`" + `access_mode` + "`" + ` is ` + "`" + `required` + "`" + ` or ` + "`" + `restricted` + "`" + `. Ex: ` + "`" + `ping_user://<USER_ID>` + "`" + ` ` + "`" + `ping_group://<GROUP_ID>` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable auth config provider. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the resource (string)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) The type of the resource (string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_bootstrap",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 bootstrap resource. This can be used to bootstrap Rancher v2 environments and output information.`,
			Description:      ``,
			Keywords: []string{
				"bootstrap",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "current_password",
					Description: `(Optional/computed/sensitive) Current password for Admin user. Just needed for recover if admin password has been changed from other resources and token is expired (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/computed/sensitive) Password for Admin user or random generated if empty (string)`,
				},
				resource.Attribute{
					Name:        "telemetry",
					Description: `(Optional) Send telemetry anonymous data. Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) TTL in seconds for generated admin token. Default: ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "token_update",
					Description: `(Optional) Regenerate admin token. Default: ` + "`" + `false` + "`" + ` (bool) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Computed) Generated API token for Admin User (string)`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `(Computed) Generated API token id for Admin User (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) URL set as server-url (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Computed) Admin username (string)`,
				},
				resource.Attribute{
					Name:        "temp_token",
					Description: `(Computed) Generated API temporary token as helper. Should be empty (string)`,
				},
				resource.Attribute{
					Name:        "temp_token_id",
					Description: `(Computed) Generated API temporary token id as helper. Should be empty (string)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Computed) Generated API token for Admin User (string)`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `(Computed) Generated API token id for Admin User (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) URL set as server-url (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Computed) Admin username (string)`,
				},
				resource.Attribute{
					Name:        "temp_token",
					Description: `(Computed) Generated API temporary token as helper. Should be empty (string)`,
				},
				resource.Attribute{
					Name:        "temp_token_id",
					Description: `(Computed) Generated API temporary token id as helper. Should be empty (string)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_catalog",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Catalog resource. This can be used to create cluster, global and/or project catalogs for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"catalog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the catalog (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The url of the catalog repo (string)`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Optional) The branch of the catalog repo to use. Default ` + "`" + `master` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional/ForceNew) The cluster id of the catalog. Mandatory if ` + "`" + `scope = cluster` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A catalog description (string)`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Optional) The kind of the catalog. Just helm by the moment (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) The password to access the catalog if needed (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional/ForceNew) The project id of the catalog. Mandatory if ` + "`" + `scope = project` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the catalog. ` + "`" + `cluster` + "`" + `, ` + "`" + `global` + "`" + `, and ` + "`" + `project` + "`" + ` are supported. Default ` + "`" + `global` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional/Sensitive) The username to access the catalog if needed (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for the catalog (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for the catalog (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_catalog` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating catalogs. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for catalog modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting catalogs. ## Import Catalogs can be imported using the Rancher Catalog ID and its scope. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_catalog.foo <scope>.<catalog_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_catalog` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating catalogs. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for catalog modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting catalogs. ## Import Catalogs can be imported using the Rancher Catalog ID and its scope. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_catalog.foo <scope>.<catalog_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_certificate",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 certificate resource. This can be used to create certificates for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "certs",
					Description: `(Required) Base64 encoded public certs (string)`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required/Sensitive) Base64 encoded private key (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required/ForceNew) The project id where the certificate should be created (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A certificate description (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional/ForceNew) The name of the certificate (string)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Optional/ForceNew) The namespace id where the namespaced certificate should be created (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for certificate object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for certificate object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_certificate` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating registries. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for certificate modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting registries.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_certificate` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating registries. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for certificate modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting registries.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cloud_credential",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2.2.x Cloud Credential resource. This can be used to create Cloud Credential for Rancher v2.2 node templates and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"credential",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Cloud Credential (string)`,
				},
				resource.Attribute{
					Name:        "amazonec2_credential_config",
					Description: `(Optional) AWS config for the Cloud Credential (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "azure_credential_config",
					Description: `(Optional) Azure config for the Cloud Credential (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the Cloud Credential (string)`,
				},
				resource.Attribute{
					Name:        "digitalocean_credential_config",
					Description: `(Optional) DigitalOcean config for the Cloud Credential (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "openstack_credential_config",
					Description: `(Optional) OpenStack config for the Cloud Credential (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "vsphere_credential_config",
					Description: `(Optional) vSphere config for the Cloud Credential (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Annotations for Cloud Credential object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for Cloud Credential object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Computed) The driver of the Cloud Credential (string) ## Nested blocks ### ` + "`" + `amazonec2_credential_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required/Sensitive) AWS access key (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required/Sensitive) AWS secret key (string) ### ` + "`" + `azure_credential_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required/Sensitive) Azure Service Principal Account ID (string)`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Required/Sensitive) Azure Service Principal Account password (string)`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required/Sensitive) Azure Subscription ID (string) ### ` + "`" + `digitalocean_credential_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `(Required/Sensitive) DigitalOcean access token (string) ### ` + "`" + `openstack_credential_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) OpenStack password (string) ### ` + "`" + `vsphere_credential_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) vSphere password (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) vSphere username (string)`,
				},
				resource.Attribute{
					Name:        "vcenter",
					Description: `(Required) vSphere IP/hostname for vCenter (string)`,
				},
				resource.Attribute{
					Name:        "vcenter_port",
					Description: `(Optional) vSphere Port for vCenter. Default ` + "`" + `443` + "`" + ` (string) ## Timeouts ` + "`" + `rancher2_cloud_credential` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cloud credentials. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cloud credential modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cloud credentials.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Computed) The driver of the Cloud Credential (string) ## Nested blocks ### ` + "`" + `amazonec2_credential_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required/Sensitive) AWS access key (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required/Sensitive) AWS secret key (string) ### ` + "`" + `azure_credential_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required/Sensitive) Azure Service Principal Account ID (string)`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Required/Sensitive) Azure Service Principal Account password (string)`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required/Sensitive) Azure Subscription ID (string) ### ` + "`" + `digitalocean_credential_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `(Required/Sensitive) DigitalOcean access token (string) ### ` + "`" + `openstack_credential_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) OpenStack password (string) ### ` + "`" + `vsphere_credential_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) vSphere password (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) vSphere username (string)`,
				},
				resource.Attribute{
					Name:        "vcenter",
					Description: `(Required) vSphere IP/hostname for vCenter (string)`,
				},
				resource.Attribute{
					Name:        "vcenter_port",
					Description: `(Optional) vSphere Port for vCenter. Default ` + "`" + `443` + "`" + ` (string) ## Timeouts ` + "`" + `rancher2_cloud_credential` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cloud credentials. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cloud credential modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cloud credentials.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cluster",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Cluster resource. This can be used to create Clusters for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Cluster (string)`,
				},
				resource.Attribute{
					Name:        "rke_config",
					Description: `(Optional/Computed) The RKE configuration for ` + "`" + `rke` + "`" + ` Clusters. Conflicts with ` + "`" + `aks_config` + "`" + `, ` + "`" + `eks_config` + "`" + ` and ` + "`" + `gke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "aks_config",
					Description: `(Optional) The Azure AKS configuration for ` + "`" + `aks` + "`" + ` Clusters. Conflicts with ` + "`" + `eks_config` + "`" + `, ` + "`" + `gke_config` + "`" + ` and ` + "`" + `rke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "eks_config",
					Description: `(Optional) The Amazon EKS configuration for ` + "`" + `eks` + "`" + ` Clusters. Conflicts with ` + "`" + `aks_config` + "`" + `, ` + "`" + `gke_config` + "`" + ` and ` + "`" + `rke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "gke_config",
					Description: `(Optional) The Google GKE configuration for ` + "`" + `gke` + "`" + ` Clusters. Conflicts with ` + "`" + `aks_config` + "`" + `, ` + "`" + `eks_config` + "`" + ` and ` + "`" + `rke_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for Cluster (string)`,
				},
				resource.Attribute{
					Name:        "cluster_auth_endpoint",
					Description: `(Optional/Computed) Enabling the [local cluster authorized endpoint](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#local-cluster-auth-endpoint) allows direct communication with the cluster, bypassing the Rancher API proxy. (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "cluster_monitoring_input",
					Description: `(Optional/Computed) Cluster monitoring config. Any parameter defined in [rancher-monitoring charts](https://github.com/rancher/system-charts/tree/dev/charts/rancher-monitoring) could be configured (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "cluster_template_answers",
					Description: `(Optional) Cluster template answers. Just for Rancher v2.3.x and above (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "cluster_template_id",
					Description: `(Optional) Cluster template ID. Just for Rancher v2.3.x and above (string)`,
				},
				resource.Attribute{
					Name:        "cluster_template_questions",
					Description: `(Optional) Cluster template questions. Just for Rancher v2.3.x and above (list)`,
				},
				resource.Attribute{
					Name:        "cluster_template_revision_id",
					Description: `(Optional) Cluster template revision ID. Just for Rancher v2.3.x and above (string)`,
				},
				resource.Attribute{
					Name:        "default_pod_security_policy_template_id",
					Description: `(Optional/Computed) [Default pod security policy template id](https://rancher.com/docs/rancher/v2.x/en/cluster-provisioning/rke-clusters/options/#pod-security-policy-support) (string)`,
				},
				resource.Attribute{
					Name:        "desired_agent_image",
					Description: `(Optional/Computed) Desired agent image. Just for Rancher v2.3.x and above (string)`,
				},
				resource.Attribute{
					Name:        "desired_auth_image",
					Description: `(Optional/Computed) Desired auth image. Just for Rancher v2.3.x and above (string)`,
				},
				resource.Attribute{
					Name:        "docker_root_dir",
					Description: `(Optional/Computed) Desired auth image. Just for Rancher v2.3.x and above (string)`,
				},
				resource.Attribute{
					Name:        "enable_cluster_alerting",
					Description: `(Optional) Enable built-in cluster alerting. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_cluster_monitoring",
					Description: `(Optional) Enable built-in cluster monitoring. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_cluster_istio",
					Description: `(Optional) Enable built-in cluster istio. Default ` + "`" + `false` + "`" + `. Just for Rancher v2.3.x and above (bool)`,
				},
				resource.Attribute{
					Name:        "enable_network_policy",
					Description: `(Optional) Enable project network isolation. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for Node Pool object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for Node Pool object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "cluster_registration_token",
					Description: `(Computed) Cluster Registration Token generated for the cluster (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "default_project_id",
					Description: `(Computed) Default project ID for the cluster (string)`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Computed) The driver used for the Cluster. ` + "`" + `imported` + "`" + `, ` + "`" + `azurekubernetesservice` + "`" + `, ` + "`" + `amazonelasticcontainerservice` + "`" + `, ` + "`" + `googlekubernetesengine` + "`" + ` and ` + "`" + `rancherKubernetesEngine` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `(Computed) Kube Config generated for the cluster (string)`,
				},
				resource.Attribute{
					Name:        "system_project_id",
					Description: `(Computed) System project ID for the cluster (string) ## Nested blocks ### ` + "`" + `rke_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "addon_job_timeout",
					Description: `(Optional/Computed) Duration in seconds of addon job (int)`,
				},
				resource.Attribute{
					Name:        "addons",
					Description: `(Optional) Addons descripton to deploy on RKE cluster.`,
				},
				resource.Attribute{
					Name:        "addons_include",
					Description: `(Optional) Addons yaml manifests to deploy on RKE cluster (list)`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional/Computed) Kubernetes cluster authentication (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "authorization",
					Description: `(Optional/Computed) Kubernetes cluster authorization (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "bastion_host",
					Description: `(Optional/Computed) RKE bastion host (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Optional/Computed) RKE cloud provider [rke-cloud-providers](https://rancher.com/docs/rke/v0.1.x/en/config-options/cloud-providers/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional/Computed) RKE dns add-on. Just for Rancher v2.2.x (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "ignore_docker_version",
					Description: `(Optional) Ignore docker version. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional/Computed) Kubernetes ingress configuration (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `(Optional/Computed) Kubernetes version to deploy (string)`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional/Computed) Kubernetes cluster monitoring (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional/Computed) Kubernetes cluster networking (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `(Optional) RKE cluster nodes (list)`,
				},
				resource.Attribute{
					Name:        "prefix_path",
					Description: `(Optional/Computed) Prefix to customize Kubernetes path (string)`,
				},
				resource.Attribute{
					Name:        "private_registries",
					Description: `(Optional) private registries for docker images (list)`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional/Computed) Kubernetes cluster services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "ssh_agent_auth",
					Description: `(Optional) Use ssh agent auth. Default ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional/Computed) Cluster level SSH private key path (string) #### ` + "`" + `authentication` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "sans",
					Description: `(Optional/Computed) RKE sans for authentication ([]string)`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional/Computed) RKE strategy for authentication (string) #### ` + "`" + `authorization` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) RKE mode for authorization. ` + "`" + `rbac` + "`" + ` and ` + "`" + `none` + "`" + ` modes are available. Default ` + "`" + `rbac` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) RKE options for authorization (map) #### ` + "`" + `bastion_host` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Address ip for the bastion host (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) User to connect bastion host (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port for bastion host. Default ` + "`" + `22` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_agent_auth",
					Description: `(Optional) Use ssh agent auth. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `(Optional/Computed/Sensitive) Bastion host SSH private key (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional/Computed) Bastion host SSH private key path (string) #### ` + "`" + `cloud_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "aws_cloud_provider",
					Description: `(Optional/Computed) RKE AWS Cloud Provider config for Cloud Provider [rke-aws-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/aws/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "azure_cloud_provider",
					Description: `(Optional/Computed) RKE Azure Cloud Provider config for Cloud Provider [rke-azure-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/azure/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "custom_cloud_provider",
					Description: `(Optional/Computed) RKE Custom Cloud Provider config for Cloud Provider (string) (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional/Computed) RKE sans for Cloud Provider. ` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + `, ` + "`" + `custom` + "`" + `, ` + "`" + `openstack` + "`" + `, ` + "`" + `vsphere` + "`" + ` are supported. (string)`,
				},
				resource.Attribute{
					Name:        "openstack_cloud_provider",
					Description: `(Optional/Computed) RKE Openstack Cloud Provider config for Cloud Provider [rke-openstack-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/openstack/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "vsphere_cloud_provider",
					Description: `(Optional/Computed) RKE Vsphere Cloud Provider config for Cloud Provider [rke-vsphere-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/vsphere/) Extra argument ` + "`" + `name` + "`" + ` is required on ` + "`" + `virtual_center` + "`" + ` configuration. (list maxitems:1) ##### ` + "`" + `aws_cloud_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "service_override",
					Description: `(Optional) (list) ###### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "disable_security_group_ingress",
					Description: `(Optional) Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "disable_strict_zone_check",
					Description: `(Optional) Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "elb_security_group",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_cluster_id",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_cluster_tag",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional/Computed/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional/Computed) (string) ###### ` + "`" + `service_override` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "signing_method",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "signing_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "signing_region",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional/Computed) (string) ##### ` + "`" + `azure_cloud_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "aad_client_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_secret",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_cert_password",
					Description: `(Optional/Computed/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_cert_path",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "cloud",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_duration",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_exponent",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_jitter",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_retries",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit_bucket",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit_qps",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "maximum_load_balancer_rule_count",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "primary_availability_set_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "primary_scale_set_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "route_table_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "use_instance_metadata",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "use_managed_identity_extension",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "vnet_resource_group",
					Description: `(Optional/Computed) (string) ##### ` + "`" + `openstack_cloud_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Required) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "block_storage",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `(Optional/Computed) (list maxitems:1) ###### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "auth_url",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "ca_file",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional/Computed/Sensitive) Required if ` + "`" + `domain_name` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional/Computed) Required if ` + "`" + `domain_id` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional/Computed/Sensitive) Required if ` + "`" + `tenant_name` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `(Optional/Computed) Required if ` + "`" + `tenant_id` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "trust_id",
					Description: `(Optional/Computed/Sensitive) (string) ###### ` + "`" + `block_storage` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "bs_version",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "ignore_volume_az",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "trust_device_path",
					Description: `(Optional/Computed) (string) ###### ` + "`" + `load_balancer` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "create_monitor",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "floating_network_id",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "lb_provider",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "lb_version",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "manage_security_groups",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "monitor_delay",
					Description: `(Optional/Computed) Default ` + "`" + `60s` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "monitor_max_retries",
					Description: `(Optional/Computed) Default 5 (int)`,
				},
				resource.Attribute{
					Name:        "monitor_timeout",
					Description: `(Optional/Computed) Default ` + "`" + `30s` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "use_octavia",
					Description: `(Optional/Computed) (bool) ###### ` + "`" + `metadata` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "request_timeout",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "search_order",
					Description: `(Optional/Computed) (string) ###### ` + "`" + `route` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Optional/Computed) (string) ##### ` + "`" + `vsphere_cloud_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "virtual_center",
					Description: `(Required) (List)`,
				},
				resource.Attribute{
					Name:        "workspace",
					Description: `(Required) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional/Computed) (list maxitems:1) ###### ` + "`" + `virtual_center` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of virtualcenter config for Vsphere Cloud Provider config (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "soap_roundtrip_count",
					Description: `(Optional/Computed) (int) ###### ` + "`" + `workspace` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "default_datastore",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "resourcepool_path",
					Description: `(Optional/Computed) (string) ###### ` + "`" + `disk` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "scsi_controller_type",
					Description: `(Optional/Computed) (string) ###### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "insecure_flag",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "soap_roundtrip_count",
					Description: `(Optional/Computed) (int) ###### ` + "`" + `network` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "public_network",
					Description: `(Optional/Computed) (string) #### ` + "`" + `dns` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional/Computed) DNS add-on node selector (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional) DNS add-on provider. ` + "`" + `kube-dns` + "`" + `, ` + "`" + `coredns` + "`" + ` (default), and ` + "`" + `none` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "reverse_cidrs",
					Description: `(Optional/Computed) DNS add-on reverse cidr (list)`,
				},
				resource.Attribute{
					Name:        "upstream_nameservers",
					Description: `(Optional/Computed) DNS add-on upstream nameservers (list) #### ` + "`" + `ingress` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for RKE Ingress (map)`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional/Computed) Node selector for RKE Ingress (map)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) RKE options for Ingress (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional/Computed) Provider for RKE Ingress (string) #### ` + "`" + `monitoring` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) RKE options for monitoring (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional/Computed) Provider for RKE monitoring (string) #### ` + "`" + `network` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "calico_network_provider",
					Description: `(Optional/Computed) Calico provider config for RKE network (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "canal_network_provider",
					Description: `(Optional/Computed) Canal provider config for RKE network (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "flannel_network_provider",
					Description: `(Optional/Computed) Flannel provider config for RKE network (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "weave_network_provider",
					Description: `(Optional/Computed) Weave provider config for RKE network (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) RKE options for network (map)`,
				},
				resource.Attribute{
					Name:        "plugin",
					Description: `(Optional/Computed) Plugin for RKE network. ` + "`" + `canal` + "`" + ` (default), ` + "`" + `flannel` + "`" + `, ` + "`" + `calico` + "`" + ` and ` + "`" + `weave` + "`" + ` are supported. (string) ##### ` + "`" + `calico_network_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Optional/Computed) RKE options for Calico network provider (string) ##### ` + "`" + `canal_network_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "iface",
					Description: `(Optional/Computed) Iface config Canal network provider (string) ##### ` + "`" + `flannel_network_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "iface",
					Description: `(Optional/Computed) Iface config Flannel network provider (string) ##### ` + "`" + `weave_network_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Computed) Password config Weave network provider (string) #### ` + "`" + `nodes` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Address ip for node (string)`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Requires) Roles for the node. ` + "`" + `controlplane` + "`" + `, ` + "`" + `etcd` + "`" + ` and ` + "`" + `worker` + "`" + ` are supported. (list)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required/Sensitive) User to connect node (string)`,
				},
				resource.Attribute{
					Name:        "docker_socket",
					Description: `(Optional/Computed) Docker socket for node (string)`,
				},
				resource.Attribute{
					Name:        "hostname_override",
					Description: `(Optional) Hostname override for node (string)`,
				},
				resource.Attribute{
					Name:        "internal_address",
					Description: `(Optional) Internal ip for node (string)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels for the node (map)`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `(Optional) Id for the node (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port for node. Default ` + "`" + `22` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_agent_auth",
					Description: `(Optional) Use ssh agent auth. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `(Optional/Computed/Sensitive) Node SSH private key (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional/Computed) Node SSH private key path (string) #### ` + "`" + `private_registries` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Registry URL (string)`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Set as default registry. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) Registry password (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional/Sensitive) Registry user (string) #### ` + "`" + `services` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "etcd",
					Description: `(Optional/Computed) Etcd options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kube_api",
					Description: `(Optional/Computed) Kube API options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kube_controller",
					Description: `(Optional/Computed) Kube Controller options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kubelet",
					Description: `(Optional/Computed) Kubelet options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kubeproxy",
					Description: `(Optional/Computed) Kubeproxy options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional/Computed) Scheduler options for RKE services (list maxitems:1) ##### ` + "`" + `etcd` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "backup_config",
					Description: `(Optional/Computed) Backup options for etcd service. Just for Rancher v2.2.x (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `(Optional/Computed) TLS CA certificate for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `(Optional/Computed/Sensitive) TLS certificate for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "creation",
					Description: `(Optional/Computed) Creation option for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "external_urls",
					Description: `(Optional) External urls for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for etcd service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `(Optional) Etcd service GID. Default: ` + "`" + `0` + "`" + `. For Rancher v2.3.x or above (int)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional/Computed/Sensitive) TLS key for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional/Computed) Path for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Optional/Computed) Retention option for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional/Computed) Snapshot option for etcd service (bool)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Etcd service UID. Default: ` + "`" + `0` + "`" + `. For Rancher v2.3.x or above (int) ###### ` + "`" + `backup_config` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable etcd backup (bool)`,
				},
				resource.Attribute{
					Name:        "interval_hours",
					Description: `(Optional) Interval hours for etcd backup. Default ` + "`" + `12` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Optional) Retention for etcd backup. Default ` + "`" + `6` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "s3_backup_config",
					Description: `(Optional) S3 config options for etcd backup (list maxitems:1) ###### ` + "`" + `s3_backup_config` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional/Sensitive) Access key for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) Bucket name for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "custom_ca",
					Description: `(Optional) Base64 encoded custom CA for S3 service. Use filebase64(<FILE>) for encoding file. Available from Rancher v2.2.5 (string)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) Folder for S3 service. Available from Rancher v2.2.7 (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional/Sensitive) Secret key for S3 service (string) ##### ` + "`" + `kube_api` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "always_pull_images",
					Description: `(Optional) Enable [AlwaysPullImages](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#alwayspullimages) Admission controller plugin. [Rancher docs](https://rancher.com/docs/rke/latest/en/config-options/services/#kubernetes-api-server-options) Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kube API service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for kube API service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for kube API service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kube API service (string)`,
				},
				resource.Attribute{
					Name:        "pod_security_policy",
					Description: `(Optional) Pod Security Policy option for kube API service. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "service_cluster_ip_range",
					Description: `(Optional/Computed) Service Cluster IP Range option for kube API service (string)`,
				},
				resource.Attribute{
					Name:        "service_node_port_range",
					Description: `(Optional/Computed) Service Node Port Range option for kube API service (string) ##### ` + "`" + `kube_controller` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `(Optional/Computed) Cluster CIDR option for kube controller service (string)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kube controller service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for kube controller service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for kube controller service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kube controller service (string)`,
				},
				resource.Attribute{
					Name:        "service_cluster_ip_range",
					Description: `(Optional/Computed) Service Cluster ip Range option for kube controller service (string) ##### ` + "`" + `kubelet` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_dns_server",
					Description: `(Optional/Computed) Cluster DNS Server option for kubelet service (string)`,
				},
				resource.Attribute{
					Name:        "cluster_domain",
					Description: `(Optional/Computed) Cluster Domain option for kubelet service (string)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kubelet service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for kubelet service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for kubelet service (list)`,
				},
				resource.Attribute{
					Name:        "fail_swap_on",
					Description: `(Optional/Computed) Enable or disable failing when swap on is not supported (bool)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kubelet service (string)`,
				},
				resource.Attribute{
					Name:        "infra_container_image",
					Description: `(Optional/Computed) Infra container image for kubelet service (string) ##### ` + "`" + `kubeproxy` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kubeproxy service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for kubeproxy service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for kubeproxy service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kubeproxy service (string) ##### ` + "`" + `scheduler` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for scheduler service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for scheduler service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for scheduler service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for scheduler service (string) ### ` + "`" + `aks_config` + "`" + ` #### Arguments The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "agent_dns_prefix",
					Description: `(Required) DNS prefix to be used to create the FQDN for the agent pool (string)`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required/Sensitive) Azure client ID to use (string)`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Required/Sensitive) Azure client secret associated with the \"client id\" (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `(Required) Specify the version of Kubernetes. To check available versions exec ` + "`" + `az aks get-versions -l eastus -o table` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "master_dns_prefix",
					Description: `(Required) DNS prefix to use the Kubernetes cluster control pane (string)`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Required) The name of the Cluster resource group (string)`,
				},
				resource.Attribute{
					Name:        "ssh_public_key_contents",
					Description: `(Required) Contents of the SSH public key used to authenticate with Linux hosts (string)`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) The name of an existing Azure Virtual Subnet. Composite of agent virtual network subnet ID (string)`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) Subscription credentials which uniquely identify Microsoft Azure subscription (string)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Azure tenant ID to use (string)`,
				},
				resource.Attribute{
					Name:        "virtual_network",
					Description: `(Required) The name of an existing Azure Virtual Network. Composite of agent virtual network subnet ID (string)`,
				},
				resource.Attribute{
					Name:        "virtual_network_resource_group",
					Description: `(Required) The resource group of an existing Azure Virtual Network. Composite of agent virtual network subnet ID (string)`,
				},
				resource.Attribute{
					Name:        "add_client_app_id",
					Description: `(Optional/Sensitive) The ID of an Azure Active Directory client application of type \"Native\". This application is for user login via kubectl (string)`,
				},
				resource.Attribute{
					Name:        "add_server_app_id",
					Description: `(Optional/Sensitive) The ID of an Azure Active Directory server application of type \"Web app/API\". This application represents the managed cluster's apiserver (Server application) (string)`,
				},
				resource.Attribute{
					Name:        "aad_server_app_secret",
					Description: `(Optional/Sensitive) The secret of an Azure Active Directory server application (string)`,
				},
				resource.Attribute{
					Name:        "aad_tenant_id",
					Description: `(Optional/Sensitive) The ID of an Azure Active Directory tenant (string)`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `(Optional) The administrator username to use for Linux hosts. Default ` + "`" + `azureuser` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "agent_os_disk_size",
					Description: `(Optional) GB size to be used to specify the disk for every machine in the agent pool. If you specify 0, it will apply the default according to the \"agent vm size\" specified. Default ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "agent_pool_name",
					Description: `(Optional) Name for the agent pool, upto 12 alphanumeric characters. Default ` + "`" + `agentpool0` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "agent_storage_profile",
					Description: `(Optional) Storage profile specifies what kind of storage used on machine in the agent pool. Chooses from [ManagedDisks StorageAccount]. Default ` + "`" + `ManagedDisks` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "agent_vm_size",
					Description: `(Optional) Size of machine in the agent pool. Default ` + "`" + `Standard_D1_v2` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "auth_base_url",
					Description: `(Optional) Different authentication API url to use. Default ` + "`" + `https://login.microsoftonline.com/` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `(Optional) Different resource management API url to use. Default ` + "`" + `https://management.azure.com/` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional) Number of machines (VMs) in the agent pool. Allowed values must be in the range of 1 to 100 (inclusive). Default ` + "`" + `1` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "dns_service_ip",
					Description: `(Optional) An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes Service address range specified in \"service cidr\". Default ` + "`" + `10.0.0.10` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "docker_bridge_cidr",
					Description: `(Required) A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes Service address range specified in \"service cidr\". Default ` + "`" + `172.17.0.1/16` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "enable_http_application_routing",
					Description: `(Optional) Enable the Kubernetes ingress with automatic public DNS name creation. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_monitoring",
					Description: `(Optional) Turn on Azure Log Analytics monitoring. Uses the Log Analytics \"Default\" workspace if it exists, else creates one. if using an existing workspace, specifies \"log analytics workspace resource id\". Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Azure Kubernetes cluster location. Default ` + "`" + `eastus` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "log_analytics_workspace",
					Description: `(Optional) The name of an existing Azure Log Analytics Workspace to use for storing monitoring data. If not specified, uses '{resource group}-{subscription id}-{location code}' (string)`,
				},
				resource.Attribute{
					Name:        "log_analytics_workspace_resource_group",
					Description: `(Optional) The resource group of an existing Azure Log Analytics Workspace to use for storing monitoring data. If not specified, uses the 'Cluster' resource group (string)`,
				},
				resource.Attribute{
					Name:        "max_pods",
					Description: `(Optional) Maximum number of pods that can run on a node. Default ` + "`" + `110` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "network_plugin",
					Description: `(Optional) Network plugin used for building Kubernetes network. Chooses from ` + "`" + `azure` + "`" + ` or ` + "`" + `kubenet` + "`" + `. Default ` + "`" + `azure` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "network_policy",
					Description: `(Optional) Network policy used for building Kubernetes network. Chooses from ` + "`" + `calico` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "pod_cidr",
					Description: `(Optional) A CIDR notation IP range from which to assign Kubernetes Pod IPs when \"network plugin\" is specified in \"kubenet\". Default ` + "`" + `172.244.0.0/16` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "service_cidr",
					Description: `(Optional) A CIDR notation IP range from which to assign Kubernetes Service cluster IPs. It must not overlap with any Subnet IP ranges. Default ` + "`" + `10.0.0.0/16` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional/Computed) Tags for Kubernetes cluster. For example, foo=bar (map) ### ` + "`" + `eks_config` + "`" + ` #### Arguments The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required/Sensitive) The AWS Client ID to use (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `(Required) The Kubernetes master version (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required/Sensitive) The AWS Client Secret associated with the Client ID (string)`,
				},
				resource.Attribute{
					Name:        "ami",
					Description: `(Optional) AMI ID to use for the worker nodes instead of the default (string)`,
				},
				resource.Attribute{
					Name:        "associate_worker_node_public_ip",
					Description: `(Optional) Associate public ip EKS worker nodes. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "desired_nodes",
					Description: `(Optional) The desired number of worker nodes. Just for Rancher v2.3.x and above. Default ` + "`" + `3` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The type of machine to use for worker nodes. Default ` + "`" + `t2.medium` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "key_pair_name",
					Description: `(Optional) Allow user to specify key name to use. Just for Rancher v2.2.7 and above (string)`,
				},
				resource.Attribute{
					Name:        "maximum_nodes",
					Description: `(Optional) The maximum number of worker nodes. Default ` + "`" + `3` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "minimum_nodes",
					Description: `(Optional) The minimum number of worker nodes. Default ` + "`" + `1` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "node_volume_size",
					Description: `(Optional) The volume size for each node. Default ` + "`" + `20` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The AWS Region to create the EKS cluster in. Default ` + "`" + `us-west-2` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) List of security groups to use for the cluster. If it's not specified Rancher will create a new security group (list)`,
				},
				resource.Attribute{
					Name:        "service_role",
					Description: `(Optional) The service role to use to perform the cluster operations in AWS. If it's not specified Rancher will create a new service role (string)`,
				},
				resource.Attribute{
					Name:        "session_token",
					Description: `(Optional/Sensitive) A session token to use with the client key and secret if applicable (string)`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) List of subnets in the virtual network to use. If it's not specified Rancher will create 3 news subnets (list)`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional/Computed) Pass user-data to the nodes to perform automated configuration tasks (string)`,
				},
				resource.Attribute{
					Name:        "virtual_network",
					Description: `(Optional) The name of the virtual network to use. If it's not specified Rancher will create a new VPC (string) ### ` + "`" + `gke_config` + "`" + ` #### Arguments The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "cluster_ipv4_cidr",
					Description: `(Required) The IP address range of the container pods (string)`,
				},
				resource.Attribute{
					Name:        "credential",
					Description: `(Required/Sensitive) The contents of the GC credential file (string)`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Required) Type of the disk attached to each node (string)`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `(Required) The image to use for the worker nodes (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_cluster_ipv4_cidr_block",
					Description: `(Required) The IP address range for the cluster pod IPs (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_cluster_secondary_range_name",
					Description: `(Required) The name of the secondary range to be used for the cluster CIDR block (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_node_ipv4_cidr_block",
					Description: `(Required) The IP address range of the instance IPs in this cluster (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_services_ipv4_cidr_block",
					Description: `(Required) The IP address range of the services IPs in this cluster (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_services_secondary_range_name",
					Description: `(Required) The name of the secondary range to be used for the services CIDR block (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_subnetwork_name",
					Description: `(Required) A custom subnetwork name to be used if createSubnetwork is true (string)`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(Required) Locations for GKE cluster (list)`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required) Machine type for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `(Required) Maintenance window for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "master_ipv4_cidr_block",
					Description: `(Required) The IP range in CIDR notation to use for the hosted master network (string)`,
				},
				resource.Attribute{
					Name:        "master_version",
					Description: `(Required) Master version for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Network for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `(Required) The ID of the cluster node pool (string)`,
				},
				resource.Attribute{
					Name:        "node_version",
					Description: `(Required) Node version for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "oauth_scopes",
					Description: `(Required) The set of Google API scopes to be made available on all of the node VMs under the default service account (list)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Required) The Google Cloud Platform Service Account to be used by the node VMs (string)`,
				},
				resource.Attribute{
					Name:        "sub_network",
					Description: `(Required) Subnetwork for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this cluster (string)`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional) Size of the disk attached to each node. Default ` + "`" + `100` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "enable_alpha_feature",
					Description: `(Optional) To enable Kubernetes alpha feature. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_auto_repair",
					Description: `(Optional) Specifies whether the node auto-repair is enabled for the node pool. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_auto_upgrade",
					Description: `(Optional) Specifies whether node auto-upgrade is enabled for the node pool. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_horizontal_pod_autoscaling",
					Description: `(Optional) Enable horizontal pod autoscaling for the cluster. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_http_load_balancing",
					Description: `(Optional) Enable HTTP load balancing on GKE cluster. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_kubernetes_dashboard",
					Description: `(Optional) Whether to enable the Kubernetes dashboard. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_legacy_abac",
					Description: `(Optional) Whether to enable legacy abac on the cluster. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_network_policy_config",
					Description: `(Optional) Enable stackdriver logging. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_nodepool_autoscaling",
					Description: `(Optional) Enable nodepool autoscaling. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_private_endpoint",
					Description: `(Optional) Whether the master's internal IP address is used as the cluster endpoint. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_private_nodes",
					Description: `(Optional) Whether nodes have internal IP address only. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_stackdriver_logging",
					Description: `(Optional) Enable stackdriver monitoring. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_stackdriver_monitoring",
					Description: `(Optional) Enable stackdriver monitoring on GKE cluster (bool)`,
				},
				resource.Attribute{
					Name:        "ip_policy_create_subnetwork",
					Description: `(Optional) Whether a new subnetwork will be created automatically for the cluster. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "issue_client_certificate",
					Description: `(Optional) Issue a client certificate. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "kubernetes_dashboard",
					Description: `(Optional) Enable the Kubernetes dashboard. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) The map of Kubernetes labels to be applied to each node (map)`,
				},
				resource.Attribute{
					Name:        "local_ssd_count",
					Description: `(Optional) The number of local SSD disks to be attached to the node. Default ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "master_authorized_network_cidr_blocks",
					Description: `(Optional) Define up to 10 external networks that could access Kubernetes master through HTTPS (list)`,
				},
				resource.Attribute{
					Name:        "max_node_count",
					Description: `(Optional) Maximum number of nodes in the NodePool. Must be >= minNodeCount. There has to enough quota to scale up the cluster. Default ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "min_node_count",
					Description: `(Optional) Minimmum number of nodes in the NodePool. Must be >= 1 and <= maxNodeCount. Default ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Optional) Node count for GKE cluster. Default ` + "`" + `3` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) Whether the nodes are created as preemptible VM instances. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "resource_labels",
					Description: `(Optional/Computed) The map of Kubernetes labels to be applied to each cluster (map)`,
				},
				resource.Attribute{
					Name:        "use_ip_aliases",
					Description: `(Optional) Whether alias IPs will be used for pod IPs in the cluster. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "taints",
					Description: `(Required) List of Kubernetes taints to be applied to each node (list)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) Zone GKE cluster (string) ### ` + "`" + `cluster_auth_endpoint` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ca_certs",
					Description: `(Optional) CA certs for the authorized cluster endpoint (string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable the authorized cluster endpoint. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Optional) FQDN for the authorized cluster endpoint (string) ### ` + "`" + `cluster_monitoring_input` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "answers",
					Description: `(Optional/Computed) Key/value answers for monitor input (map) ======= ### ` + "`" + `cluster_template_answers` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) Cluster ID to apply answer (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID to apply answer (string)`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Key/values for answer (map) ### ` + "`" + `cluster_template_questions` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) Default variable value (string)`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional) Required variable. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Variable type. ` + "`" + `boolean` + "`" + `, ` + "`" + `int` + "`" + ` and ` + "`" + `string` + "`" + ` are allowed. Default ` + "`" + `string` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "variable",
					Description: `(Optional) Variable name (string) >>>>>>> c6a2cbc... Feat: added rancher2_cluster_template datasource and resource. For rancher V2.3.x. Doc files ### ` + "`" + `cluster_registration_token` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Computed) Cluster ID (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) Name of cluster registration token (string)`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Computed) Command to execute in a imported k8s cluster (string)`,
				},
				resource.Attribute{
					Name:        "insecure_command",
					Description: `(Computed) Insecure command to execute in a imported k8s cluster (string)`,
				},
				resource.Attribute{
					Name:        "manifest_url",
					Description: `(Computed) K8s manifest url to execute with ` + "`" + `kubectl` + "`" + ` to import an existing k8s cluster (string)`,
				},
				resource.Attribute{
					Name:        "node_command",
					Description: `(Computed) Node command to execute in linux nodes for custom k8s cluster (string)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Computed) Token for cluster registration token object (string)`,
				},
				resource.Attribute{
					Name:        "windows_node_command",
					Description: `(Computed) Node command to execute in windows nodes for custom k8s cluster (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for cluster registration token object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for cluster registration token object (map) ## Timeouts ` + "`" + `rancher2_cluster` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for creating clusters. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for cluster modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for deleting clusters. ## Import Clusters can be imported using the Rancher Cluster ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster.foo <cluster> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "cluster_registration_token",
					Description: `(Computed) Cluster Registration Token generated for the cluster (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "default_project_id",
					Description: `(Computed) Default project ID for the cluster (string)`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Computed) The driver used for the Cluster. ` + "`" + `imported` + "`" + `, ` + "`" + `azurekubernetesservice` + "`" + `, ` + "`" + `amazonelasticcontainerservice` + "`" + `, ` + "`" + `googlekubernetesengine` + "`" + ` and ` + "`" + `rancherKubernetesEngine` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `(Computed) Kube Config generated for the cluster (string)`,
				},
				resource.Attribute{
					Name:        "system_project_id",
					Description: `(Computed) System project ID for the cluster (string) ## Nested blocks ### ` + "`" + `rke_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "addon_job_timeout",
					Description: `(Optional/Computed) Duration in seconds of addon job (int)`,
				},
				resource.Attribute{
					Name:        "addons",
					Description: `(Optional) Addons descripton to deploy on RKE cluster.`,
				},
				resource.Attribute{
					Name:        "addons_include",
					Description: `(Optional) Addons yaml manifests to deploy on RKE cluster (list)`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional/Computed) Kubernetes cluster authentication (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "authorization",
					Description: `(Optional/Computed) Kubernetes cluster authorization (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "bastion_host",
					Description: `(Optional/Computed) RKE bastion host (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Optional/Computed) RKE cloud provider [rke-cloud-providers](https://rancher.com/docs/rke/v0.1.x/en/config-options/cloud-providers/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional/Computed) RKE dns add-on. Just for Rancher v2.2.x (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "ignore_docker_version",
					Description: `(Optional) Ignore docker version. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional/Computed) Kubernetes ingress configuration (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `(Optional/Computed) Kubernetes version to deploy (string)`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional/Computed) Kubernetes cluster monitoring (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional/Computed) Kubernetes cluster networking (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `(Optional) RKE cluster nodes (list)`,
				},
				resource.Attribute{
					Name:        "prefix_path",
					Description: `(Optional/Computed) Prefix to customize Kubernetes path (string)`,
				},
				resource.Attribute{
					Name:        "private_registries",
					Description: `(Optional) private registries for docker images (list)`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional/Computed) Kubernetes cluster services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "ssh_agent_auth",
					Description: `(Optional) Use ssh agent auth. Default ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional/Computed) Cluster level SSH private key path (string) #### ` + "`" + `authentication` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "sans",
					Description: `(Optional/Computed) RKE sans for authentication ([]string)`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional/Computed) RKE strategy for authentication (string) #### ` + "`" + `authorization` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) RKE mode for authorization. ` + "`" + `rbac` + "`" + ` and ` + "`" + `none` + "`" + ` modes are available. Default ` + "`" + `rbac` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) RKE options for authorization (map) #### ` + "`" + `bastion_host` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Address ip for the bastion host (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) User to connect bastion host (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port for bastion host. Default ` + "`" + `22` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_agent_auth",
					Description: `(Optional) Use ssh agent auth. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `(Optional/Computed/Sensitive) Bastion host SSH private key (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional/Computed) Bastion host SSH private key path (string) #### ` + "`" + `cloud_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "aws_cloud_provider",
					Description: `(Optional/Computed) RKE AWS Cloud Provider config for Cloud Provider [rke-aws-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/aws/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "azure_cloud_provider",
					Description: `(Optional/Computed) RKE Azure Cloud Provider config for Cloud Provider [rke-azure-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/azure/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "custom_cloud_provider",
					Description: `(Optional/Computed) RKE Custom Cloud Provider config for Cloud Provider (string) (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional/Computed) RKE sans for Cloud Provider. ` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + `, ` + "`" + `custom` + "`" + `, ` + "`" + `openstack` + "`" + `, ` + "`" + `vsphere` + "`" + ` are supported. (string)`,
				},
				resource.Attribute{
					Name:        "openstack_cloud_provider",
					Description: `(Optional/Computed) RKE Openstack Cloud Provider config for Cloud Provider [rke-openstack-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/openstack/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "vsphere_cloud_provider",
					Description: `(Optional/Computed) RKE Vsphere Cloud Provider config for Cloud Provider [rke-vsphere-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/vsphere/) Extra argument ` + "`" + `name` + "`" + ` is required on ` + "`" + `virtual_center` + "`" + ` configuration. (list maxitems:1) ##### ` + "`" + `aws_cloud_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "service_override",
					Description: `(Optional) (list) ###### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "disable_security_group_ingress",
					Description: `(Optional) Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "disable_strict_zone_check",
					Description: `(Optional) Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "elb_security_group",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_cluster_id",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_cluster_tag",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional/Computed/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional/Computed) (string) ###### ` + "`" + `service_override` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "signing_method",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "signing_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "signing_region",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional/Computed) (string) ##### ` + "`" + `azure_cloud_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "aad_client_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_secret",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_cert_password",
					Description: `(Optional/Computed/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_cert_path",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "cloud",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_duration",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_exponent",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_jitter",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_retries",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit_bucket",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit_qps",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "maximum_load_balancer_rule_count",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "primary_availability_set_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "primary_scale_set_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "route_table_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "use_instance_metadata",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "use_managed_identity_extension",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "vnet_resource_group",
					Description: `(Optional/Computed) (string) ##### ` + "`" + `openstack_cloud_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Required) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "block_storage",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `(Optional/Computed) (list maxitems:1) ###### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "auth_url",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "ca_file",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional/Computed/Sensitive) Required if ` + "`" + `domain_name` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional/Computed) Required if ` + "`" + `domain_id` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional/Computed/Sensitive) Required if ` + "`" + `tenant_name` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `(Optional/Computed) Required if ` + "`" + `tenant_id` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "trust_id",
					Description: `(Optional/Computed/Sensitive) (string) ###### ` + "`" + `block_storage` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "bs_version",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "ignore_volume_az",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "trust_device_path",
					Description: `(Optional/Computed) (string) ###### ` + "`" + `load_balancer` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "create_monitor",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "floating_network_id",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "lb_provider",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "lb_version",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "manage_security_groups",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "monitor_delay",
					Description: `(Optional/Computed) Default ` + "`" + `60s` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "monitor_max_retries",
					Description: `(Optional/Computed) Default 5 (int)`,
				},
				resource.Attribute{
					Name:        "monitor_timeout",
					Description: `(Optional/Computed) Default ` + "`" + `30s` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "use_octavia",
					Description: `(Optional/Computed) (bool) ###### ` + "`" + `metadata` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "request_timeout",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "search_order",
					Description: `(Optional/Computed) (string) ###### ` + "`" + `route` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Optional/Computed) (string) ##### ` + "`" + `vsphere_cloud_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "virtual_center",
					Description: `(Required) (List)`,
				},
				resource.Attribute{
					Name:        "workspace",
					Description: `(Required) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional/Computed) (list maxitems:1) ###### ` + "`" + `virtual_center` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of virtualcenter config for Vsphere Cloud Provider config (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "soap_roundtrip_count",
					Description: `(Optional/Computed) (int) ###### ` + "`" + `workspace` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "default_datastore",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "resourcepool_path",
					Description: `(Optional/Computed) (string) ###### ` + "`" + `disk` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "scsi_controller_type",
					Description: `(Optional/Computed) (string) ###### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "insecure_flag",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "soap_roundtrip_count",
					Description: `(Optional/Computed) (int) ###### ` + "`" + `network` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "public_network",
					Description: `(Optional/Computed) (string) #### ` + "`" + `dns` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional/Computed) DNS add-on node selector (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional) DNS add-on provider. ` + "`" + `kube-dns` + "`" + `, ` + "`" + `coredns` + "`" + ` (default), and ` + "`" + `none` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "reverse_cidrs",
					Description: `(Optional/Computed) DNS add-on reverse cidr (list)`,
				},
				resource.Attribute{
					Name:        "upstream_nameservers",
					Description: `(Optional/Computed) DNS add-on upstream nameservers (list) #### ` + "`" + `ingress` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for RKE Ingress (map)`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional/Computed) Node selector for RKE Ingress (map)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) RKE options for Ingress (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional/Computed) Provider for RKE Ingress (string) #### ` + "`" + `monitoring` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) RKE options for monitoring (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional/Computed) Provider for RKE monitoring (string) #### ` + "`" + `network` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "calico_network_provider",
					Description: `(Optional/Computed) Calico provider config for RKE network (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "canal_network_provider",
					Description: `(Optional/Computed) Canal provider config for RKE network (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "flannel_network_provider",
					Description: `(Optional/Computed) Flannel provider config for RKE network (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "weave_network_provider",
					Description: `(Optional/Computed) Weave provider config for RKE network (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) RKE options for network (map)`,
				},
				resource.Attribute{
					Name:        "plugin",
					Description: `(Optional/Computed) Plugin for RKE network. ` + "`" + `canal` + "`" + ` (default), ` + "`" + `flannel` + "`" + `, ` + "`" + `calico` + "`" + ` and ` + "`" + `weave` + "`" + ` are supported. (string) ##### ` + "`" + `calico_network_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Optional/Computed) RKE options for Calico network provider (string) ##### ` + "`" + `canal_network_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "iface",
					Description: `(Optional/Computed) Iface config Canal network provider (string) ##### ` + "`" + `flannel_network_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "iface",
					Description: `(Optional/Computed) Iface config Flannel network provider (string) ##### ` + "`" + `weave_network_provider` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Computed) Password config Weave network provider (string) #### ` + "`" + `nodes` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Address ip for node (string)`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Requires) Roles for the node. ` + "`" + `controlplane` + "`" + `, ` + "`" + `etcd` + "`" + ` and ` + "`" + `worker` + "`" + ` are supported. (list)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required/Sensitive) User to connect node (string)`,
				},
				resource.Attribute{
					Name:        "docker_socket",
					Description: `(Optional/Computed) Docker socket for node (string)`,
				},
				resource.Attribute{
					Name:        "hostname_override",
					Description: `(Optional) Hostname override for node (string)`,
				},
				resource.Attribute{
					Name:        "internal_address",
					Description: `(Optional) Internal ip for node (string)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels for the node (map)`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `(Optional) Id for the node (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port for node. Default ` + "`" + `22` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_agent_auth",
					Description: `(Optional) Use ssh agent auth. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `(Optional/Computed/Sensitive) Node SSH private key (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional/Computed) Node SSH private key path (string) #### ` + "`" + `private_registries` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Registry URL (string)`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Set as default registry. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) Registry password (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional/Sensitive) Registry user (string) #### ` + "`" + `services` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "etcd",
					Description: `(Optional/Computed) Etcd options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kube_api",
					Description: `(Optional/Computed) Kube API options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kube_controller",
					Description: `(Optional/Computed) Kube Controller options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kubelet",
					Description: `(Optional/Computed) Kubelet options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kubeproxy",
					Description: `(Optional/Computed) Kubeproxy options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional/Computed) Scheduler options for RKE services (list maxitems:1) ##### ` + "`" + `etcd` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "backup_config",
					Description: `(Optional/Computed) Backup options for etcd service. Just for Rancher v2.2.x (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `(Optional/Computed) TLS CA certificate for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `(Optional/Computed/Sensitive) TLS certificate for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "creation",
					Description: `(Optional/Computed) Creation option for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "external_urls",
					Description: `(Optional) External urls for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for etcd service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `(Optional) Etcd service GID. Default: ` + "`" + `0` + "`" + `. For Rancher v2.3.x or above (int)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional/Computed/Sensitive) TLS key for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional/Computed) Path for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Optional/Computed) Retention option for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional/Computed) Snapshot option for etcd service (bool)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Etcd service UID. Default: ` + "`" + `0` + "`" + `. For Rancher v2.3.x or above (int) ###### ` + "`" + `backup_config` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable etcd backup (bool)`,
				},
				resource.Attribute{
					Name:        "interval_hours",
					Description: `(Optional) Interval hours for etcd backup. Default ` + "`" + `12` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Optional) Retention for etcd backup. Default ` + "`" + `6` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "s3_backup_config",
					Description: `(Optional) S3 config options for etcd backup (list maxitems:1) ###### ` + "`" + `s3_backup_config` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional/Sensitive) Access key for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) Bucket name for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "custom_ca",
					Description: `(Optional) Base64 encoded custom CA for S3 service. Use filebase64(<FILE>) for encoding file. Available from Rancher v2.2.5 (string)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) Folder for S3 service. Available from Rancher v2.2.7 (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional/Sensitive) Secret key for S3 service (string) ##### ` + "`" + `kube_api` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "always_pull_images",
					Description: `(Optional) Enable [AlwaysPullImages](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#alwayspullimages) Admission controller plugin. [Rancher docs](https://rancher.com/docs/rke/latest/en/config-options/services/#kubernetes-api-server-options) Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kube API service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for kube API service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for kube API service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kube API service (string)`,
				},
				resource.Attribute{
					Name:        "pod_security_policy",
					Description: `(Optional) Pod Security Policy option for kube API service. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "service_cluster_ip_range",
					Description: `(Optional/Computed) Service Cluster IP Range option for kube API service (string)`,
				},
				resource.Attribute{
					Name:        "service_node_port_range",
					Description: `(Optional/Computed) Service Node Port Range option for kube API service (string) ##### ` + "`" + `kube_controller` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `(Optional/Computed) Cluster CIDR option for kube controller service (string)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kube controller service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for kube controller service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for kube controller service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kube controller service (string)`,
				},
				resource.Attribute{
					Name:        "service_cluster_ip_range",
					Description: `(Optional/Computed) Service Cluster ip Range option for kube controller service (string) ##### ` + "`" + `kubelet` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_dns_server",
					Description: `(Optional/Computed) Cluster DNS Server option for kubelet service (string)`,
				},
				resource.Attribute{
					Name:        "cluster_domain",
					Description: `(Optional/Computed) Cluster Domain option for kubelet service (string)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kubelet service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for kubelet service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for kubelet service (list)`,
				},
				resource.Attribute{
					Name:        "fail_swap_on",
					Description: `(Optional/Computed) Enable or disable failing when swap on is not supported (bool)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kubelet service (string)`,
				},
				resource.Attribute{
					Name:        "infra_container_image",
					Description: `(Optional/Computed) Infra container image for kubelet service (string) ##### ` + "`" + `kubeproxy` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kubeproxy service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for kubeproxy service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for kubeproxy service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kubeproxy service (string) ##### ` + "`" + `scheduler` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for scheduler service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional) Extra binds for scheduler service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional) Extra environment for scheduler service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for scheduler service (string) ### ` + "`" + `aks_config` + "`" + ` #### Arguments The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "agent_dns_prefix",
					Description: `(Required) DNS prefix to be used to create the FQDN for the agent pool (string)`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required/Sensitive) Azure client ID to use (string)`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Required/Sensitive) Azure client secret associated with the \"client id\" (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `(Required) Specify the version of Kubernetes. To check available versions exec ` + "`" + `az aks get-versions -l eastus -o table` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "master_dns_prefix",
					Description: `(Required) DNS prefix to use the Kubernetes cluster control pane (string)`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Required) The name of the Cluster resource group (string)`,
				},
				resource.Attribute{
					Name:        "ssh_public_key_contents",
					Description: `(Required) Contents of the SSH public key used to authenticate with Linux hosts (string)`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) The name of an existing Azure Virtual Subnet. Composite of agent virtual network subnet ID (string)`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) Subscription credentials which uniquely identify Microsoft Azure subscription (string)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Azure tenant ID to use (string)`,
				},
				resource.Attribute{
					Name:        "virtual_network",
					Description: `(Required) The name of an existing Azure Virtual Network. Composite of agent virtual network subnet ID (string)`,
				},
				resource.Attribute{
					Name:        "virtual_network_resource_group",
					Description: `(Required) The resource group of an existing Azure Virtual Network. Composite of agent virtual network subnet ID (string)`,
				},
				resource.Attribute{
					Name:        "add_client_app_id",
					Description: `(Optional/Sensitive) The ID of an Azure Active Directory client application of type \"Native\". This application is for user login via kubectl (string)`,
				},
				resource.Attribute{
					Name:        "add_server_app_id",
					Description: `(Optional/Sensitive) The ID of an Azure Active Directory server application of type \"Web app/API\". This application represents the managed cluster's apiserver (Server application) (string)`,
				},
				resource.Attribute{
					Name:        "aad_server_app_secret",
					Description: `(Optional/Sensitive) The secret of an Azure Active Directory server application (string)`,
				},
				resource.Attribute{
					Name:        "aad_tenant_id",
					Description: `(Optional/Sensitive) The ID of an Azure Active Directory tenant (string)`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `(Optional) The administrator username to use for Linux hosts. Default ` + "`" + `azureuser` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "agent_os_disk_size",
					Description: `(Optional) GB size to be used to specify the disk for every machine in the agent pool. If you specify 0, it will apply the default according to the \"agent vm size\" specified. Default ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "agent_pool_name",
					Description: `(Optional) Name for the agent pool, upto 12 alphanumeric characters. Default ` + "`" + `agentpool0` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "agent_storage_profile",
					Description: `(Optional) Storage profile specifies what kind of storage used on machine in the agent pool. Chooses from [ManagedDisks StorageAccount]. Default ` + "`" + `ManagedDisks` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "agent_vm_size",
					Description: `(Optional) Size of machine in the agent pool. Default ` + "`" + `Standard_D1_v2` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "auth_base_url",
					Description: `(Optional) Different authentication API url to use. Default ` + "`" + `https://login.microsoftonline.com/` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `(Optional) Different resource management API url to use. Default ` + "`" + `https://management.azure.com/` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional) Number of machines (VMs) in the agent pool. Allowed values must be in the range of 1 to 100 (inclusive). Default ` + "`" + `1` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "dns_service_ip",
					Description: `(Optional) An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes Service address range specified in \"service cidr\". Default ` + "`" + `10.0.0.10` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "docker_bridge_cidr",
					Description: `(Required) A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes Service address range specified in \"service cidr\". Default ` + "`" + `172.17.0.1/16` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "enable_http_application_routing",
					Description: `(Optional) Enable the Kubernetes ingress with automatic public DNS name creation. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_monitoring",
					Description: `(Optional) Turn on Azure Log Analytics monitoring. Uses the Log Analytics \"Default\" workspace if it exists, else creates one. if using an existing workspace, specifies \"log analytics workspace resource id\". Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Azure Kubernetes cluster location. Default ` + "`" + `eastus` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "log_analytics_workspace",
					Description: `(Optional) The name of an existing Azure Log Analytics Workspace to use for storing monitoring data. If not specified, uses '{resource group}-{subscription id}-{location code}' (string)`,
				},
				resource.Attribute{
					Name:        "log_analytics_workspace_resource_group",
					Description: `(Optional) The resource group of an existing Azure Log Analytics Workspace to use for storing monitoring data. If not specified, uses the 'Cluster' resource group (string)`,
				},
				resource.Attribute{
					Name:        "max_pods",
					Description: `(Optional) Maximum number of pods that can run on a node. Default ` + "`" + `110` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "network_plugin",
					Description: `(Optional) Network plugin used for building Kubernetes network. Chooses from ` + "`" + `azure` + "`" + ` or ` + "`" + `kubenet` + "`" + `. Default ` + "`" + `azure` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "network_policy",
					Description: `(Optional) Network policy used for building Kubernetes network. Chooses from ` + "`" + `calico` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "pod_cidr",
					Description: `(Optional) A CIDR notation IP range from which to assign Kubernetes Pod IPs when \"network plugin\" is specified in \"kubenet\". Default ` + "`" + `172.244.0.0/16` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "service_cidr",
					Description: `(Optional) A CIDR notation IP range from which to assign Kubernetes Service cluster IPs. It must not overlap with any Subnet IP ranges. Default ` + "`" + `10.0.0.0/16` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional/Computed) Tags for Kubernetes cluster. For example, foo=bar (map) ### ` + "`" + `eks_config` + "`" + ` #### Arguments The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required/Sensitive) The AWS Client ID to use (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `(Required) The Kubernetes master version (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required/Sensitive) The AWS Client Secret associated with the Client ID (string)`,
				},
				resource.Attribute{
					Name:        "ami",
					Description: `(Optional) AMI ID to use for the worker nodes instead of the default (string)`,
				},
				resource.Attribute{
					Name:        "associate_worker_node_public_ip",
					Description: `(Optional) Associate public ip EKS worker nodes. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "desired_nodes",
					Description: `(Optional) The desired number of worker nodes. Just for Rancher v2.3.x and above. Default ` + "`" + `3` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The type of machine to use for worker nodes. Default ` + "`" + `t2.medium` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "key_pair_name",
					Description: `(Optional) Allow user to specify key name to use. Just for Rancher v2.2.7 and above (string)`,
				},
				resource.Attribute{
					Name:        "maximum_nodes",
					Description: `(Optional) The maximum number of worker nodes. Default ` + "`" + `3` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "minimum_nodes",
					Description: `(Optional) The minimum number of worker nodes. Default ` + "`" + `1` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "node_volume_size",
					Description: `(Optional) The volume size for each node. Default ` + "`" + `20` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The AWS Region to create the EKS cluster in. Default ` + "`" + `us-west-2` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) List of security groups to use for the cluster. If it's not specified Rancher will create a new security group (list)`,
				},
				resource.Attribute{
					Name:        "service_role",
					Description: `(Optional) The service role to use to perform the cluster operations in AWS. If it's not specified Rancher will create a new service role (string)`,
				},
				resource.Attribute{
					Name:        "session_token",
					Description: `(Optional/Sensitive) A session token to use with the client key and secret if applicable (string)`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) List of subnets in the virtual network to use. If it's not specified Rancher will create 3 news subnets (list)`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional/Computed) Pass user-data to the nodes to perform automated configuration tasks (string)`,
				},
				resource.Attribute{
					Name:        "virtual_network",
					Description: `(Optional) The name of the virtual network to use. If it's not specified Rancher will create a new VPC (string) ### ` + "`" + `gke_config` + "`" + ` #### Arguments The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "cluster_ipv4_cidr",
					Description: `(Required) The IP address range of the container pods (string)`,
				},
				resource.Attribute{
					Name:        "credential",
					Description: `(Required/Sensitive) The contents of the GC credential file (string)`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Required) Type of the disk attached to each node (string)`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `(Required) The image to use for the worker nodes (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_cluster_ipv4_cidr_block",
					Description: `(Required) The IP address range for the cluster pod IPs (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_cluster_secondary_range_name",
					Description: `(Required) The name of the secondary range to be used for the cluster CIDR block (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_node_ipv4_cidr_block",
					Description: `(Required) The IP address range of the instance IPs in this cluster (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_services_ipv4_cidr_block",
					Description: `(Required) The IP address range of the services IPs in this cluster (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_services_secondary_range_name",
					Description: `(Required) The name of the secondary range to be used for the services CIDR block (string)`,
				},
				resource.Attribute{
					Name:        "ip_policy_subnetwork_name",
					Description: `(Required) A custom subnetwork name to be used if createSubnetwork is true (string)`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(Required) Locations for GKE cluster (list)`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required) Machine type for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `(Required) Maintenance window for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "master_ipv4_cidr_block",
					Description: `(Required) The IP range in CIDR notation to use for the hosted master network (string)`,
				},
				resource.Attribute{
					Name:        "master_version",
					Description: `(Required) Master version for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Network for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `(Required) The ID of the cluster node pool (string)`,
				},
				resource.Attribute{
					Name:        "node_version",
					Description: `(Required) Node version for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "oauth_scopes",
					Description: `(Required) The set of Google API scopes to be made available on all of the node VMs under the default service account (list)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Required) The Google Cloud Platform Service Account to be used by the node VMs (string)`,
				},
				resource.Attribute{
					Name:        "sub_network",
					Description: `(Required) Subnetwork for GKE cluster (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this cluster (string)`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional) Size of the disk attached to each node. Default ` + "`" + `100` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "enable_alpha_feature",
					Description: `(Optional) To enable Kubernetes alpha feature. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_auto_repair",
					Description: `(Optional) Specifies whether the node auto-repair is enabled for the node pool. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_auto_upgrade",
					Description: `(Optional) Specifies whether node auto-upgrade is enabled for the node pool. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_horizontal_pod_autoscaling",
					Description: `(Optional) Enable horizontal pod autoscaling for the cluster. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_http_load_balancing",
					Description: `(Optional) Enable HTTP load balancing on GKE cluster. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_kubernetes_dashboard",
					Description: `(Optional) Whether to enable the Kubernetes dashboard. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_legacy_abac",
					Description: `(Optional) Whether to enable legacy abac on the cluster. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_network_policy_config",
					Description: `(Optional) Enable stackdriver logging. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_nodepool_autoscaling",
					Description: `(Optional) Enable nodepool autoscaling. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_private_endpoint",
					Description: `(Optional) Whether the master's internal IP address is used as the cluster endpoint. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_private_nodes",
					Description: `(Optional) Whether nodes have internal IP address only. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_stackdriver_logging",
					Description: `(Optional) Enable stackdriver monitoring. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_stackdriver_monitoring",
					Description: `(Optional) Enable stackdriver monitoring on GKE cluster (bool)`,
				},
				resource.Attribute{
					Name:        "ip_policy_create_subnetwork",
					Description: `(Optional) Whether a new subnetwork will be created automatically for the cluster. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "issue_client_certificate",
					Description: `(Optional) Issue a client certificate. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "kubernetes_dashboard",
					Description: `(Optional) Enable the Kubernetes dashboard. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) The map of Kubernetes labels to be applied to each node (map)`,
				},
				resource.Attribute{
					Name:        "local_ssd_count",
					Description: `(Optional) The number of local SSD disks to be attached to the node. Default ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "master_authorized_network_cidr_blocks",
					Description: `(Optional) Define up to 10 external networks that could access Kubernetes master through HTTPS (list)`,
				},
				resource.Attribute{
					Name:        "max_node_count",
					Description: `(Optional) Maximum number of nodes in the NodePool. Must be >= minNodeCount. There has to enough quota to scale up the cluster. Default ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "min_node_count",
					Description: `(Optional) Minimmum number of nodes in the NodePool. Must be >= 1 and <= maxNodeCount. Default ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Optional) Node count for GKE cluster. Default ` + "`" + `3` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) Whether the nodes are created as preemptible VM instances. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "resource_labels",
					Description: `(Optional/Computed) The map of Kubernetes labels to be applied to each cluster (map)`,
				},
				resource.Attribute{
					Name:        "use_ip_aliases",
					Description: `(Optional) Whether alias IPs will be used for pod IPs in the cluster. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "taints",
					Description: `(Required) List of Kubernetes taints to be applied to each node (list)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) Zone GKE cluster (string) ### ` + "`" + `cluster_auth_endpoint` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ca_certs",
					Description: `(Optional) CA certs for the authorized cluster endpoint (string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable the authorized cluster endpoint. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Optional) FQDN for the authorized cluster endpoint (string) ### ` + "`" + `cluster_monitoring_input` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "answers",
					Description: `(Optional/Computed) Key/value answers for monitor input (map) ======= ### ` + "`" + `cluster_template_answers` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) Cluster ID to apply answer (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID to apply answer (string)`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Key/values for answer (map) ### ` + "`" + `cluster_template_questions` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) Default variable value (string)`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional) Required variable. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Variable type. ` + "`" + `boolean` + "`" + `, ` + "`" + `int` + "`" + ` and ` + "`" + `string` + "`" + ` are allowed. Default ` + "`" + `string` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "variable",
					Description: `(Optional) Variable name (string) >>>>>>> c6a2cbc... Feat: added rancher2_cluster_template datasource and resource. For rancher V2.3.x. Doc files ### ` + "`" + `cluster_registration_token` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Computed) Cluster ID (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) Name of cluster registration token (string)`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Computed) Command to execute in a imported k8s cluster (string)`,
				},
				resource.Attribute{
					Name:        "insecure_command",
					Description: `(Computed) Insecure command to execute in a imported k8s cluster (string)`,
				},
				resource.Attribute{
					Name:        "manifest_url",
					Description: `(Computed) K8s manifest url to execute with ` + "`" + `kubectl` + "`" + ` to import an existing k8s cluster (string)`,
				},
				resource.Attribute{
					Name:        "node_command",
					Description: `(Computed) Node command to execute in linux nodes for custom k8s cluster (string)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Computed) Token for cluster registration token object (string)`,
				},
				resource.Attribute{
					Name:        "windows_node_command",
					Description: `(Computed) Node command to execute in windows nodes for custom k8s cluster (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Computed) Annotations for cluster registration token object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Computed) Labels for cluster registration token object (map) ## Timeouts ` + "`" + `rancher2_cluster` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for creating clusters. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for cluster modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for deleting clusters. ## Import Clusters can be imported using the Rancher Cluster ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster.foo <cluster> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cluster_alert_group",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Cluster Alert Group resource. This can be used to create Cluster Alert Group for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"cluster",
				"alert",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The cluster id where create cluster alert group (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The cluster alert group name (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The cluster alert group description (string)`,
				},
				resource.Attribute{
					Name:        "group_interval_seconds",
					Description: `(Optional) The cluster alert group interval seconds. Default: ` + "`" + `180` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "group_wait_seconds",
					Description: `(Optional) The cluster alert group wait seconds. Default: ` + "`" + `180` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `(Optional) The cluster alert group recipients (list)`,
				},
				resource.Attribute{
					Name:        "repeat_interval_seconds",
					Description: `(Optional) The cluster alert group wait seconds. Default: ` + "`" + `3600` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) The cluster alert group annotations (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) The cluster alert group labels (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `recipients` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "notifier_id",
					Description: `(Required) Recipient notifier ID (string)`,
				},
				resource.Attribute{
					Name:        "recipient",
					Description: `(Optional/Computed) Recipient (string) #### Attributes`,
				},
				resource.Attribute{
					Name:        "notifier_type",
					Description: `(Computed) Recipient notifier ID. Supported values : ` + "`" + `"pagerduty" | "slack" | "email" | "webhook" | "wechat"` + "`" + ` (string) ## Timeouts ` + "`" + `rancher2_cluster_alert_group` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster alert groups. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster alert group modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster alert groups. ## Import Cluster Alert Group can be imported using the Rancher cluster alert group ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_alert_group.foo <rancher2_cluster_alert_group_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `recipients` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "notifier_id",
					Description: `(Required) Recipient notifier ID (string)`,
				},
				resource.Attribute{
					Name:        "recipient",
					Description: `(Optional/Computed) Recipient (string) #### Attributes`,
				},
				resource.Attribute{
					Name:        "notifier_type",
					Description: `(Computed) Recipient notifier ID. Supported values : ` + "`" + `"pagerduty" | "slack" | "email" | "webhook" | "wechat"` + "`" + ` (string) ## Timeouts ` + "`" + `rancher2_cluster_alert_group` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster alert groups. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster alert group modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster alert groups. ## Import Cluster Alert Group can be imported using the Rancher cluster alert group ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_alert_group.foo <rancher2_cluster_alert_group_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cluster_alert_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Cluster Alert Rule resource. This can be used to create Cluster Alert Rule for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"cluster",
				"alert",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The cluster id where create cluster alert rule (string)`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The cluster alert rule alert group ID (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The cluster alert rule name (string)`,
				},
				resource.Attribute{
					Name:        "event_rule",
					Description: `(Optional) The cluster alert rule event rule. ConflictsWith: ` + "`" + `"metric_rule", "node_rule", "system_service_rule"` + "`" + `` + "`" + ` (list Maxitems:1)`,
				},
				resource.Attribute{
					Name:        "group_interval_seconds",
					Description: `(Optional) The cluster alert rule group interval seconds. Default: ` + "`" + `180` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "group_wait_seconds",
					Description: `(Optional) The cluster alert rule group wait seconds. Default: ` + "`" + `180` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "inherited",
					Description: `(Optional) The cluster alert rule inherited. Default: ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "metric_rule",
					Description: `(Optional) The cluster alert rule metric rule. ConflictsWith: ` + "`" + `"event_rule", "node_rule", "system_service_rule"` + "`" + `` + "`" + ` (list Maxitems:1)`,
				},
				resource.Attribute{
					Name:        "node_rule",
					Description: `(Optional) The cluster alert rule node rule. ConflictsWith: ` + "`" + `"event_rule", "metric_rule", "system_service_rule"` + "`" + `` + "`" + ` (list Maxitems:1)`,
				},
				resource.Attribute{
					Name:        "repeat_interval_seconds",
					Description: `(Optional) The cluster alert rule wait seconds. Default: ` + "`" + `3600` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Optional) The cluster alert rule severity. Supported values : ` + "`" + `"critical" | "info" | "warning"` + "`" + `. Default: ` + "`" + `critical` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "system_service_rule",
					Description: `(Optional) The cluster alert rule system service rule. ConflictsWith: ` + "`" + `"event_rule", "metric_rule", "node_rule"` + "`" + `` + "`" + ` (list Maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) The cluster alert rule annotations (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) The cluster alert rule labels (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `event_rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "resource_kind",
					Description: `(Required) Resource kind. Supported values : ` + "`" + `"DaemonSet" | "Deployment" | "Node" | "Pod" | "StatefulSet"` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `(Optional) Event type. Supported values : ` + "`" + `"Warning" | "Normal"` + "`" + `. Default: ` + "`" + `Warning` + "`" + ` (string) ### ` + "`" + `metric_rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) Metric rule duration (string)`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `(Required) Metric rule expression (string)`,
				},
				resource.Attribute{
					Name:        "threshold_value",
					Description: `(Required) Metric rule threshold value (float64)`,
				},
				resource.Attribute{
					Name:        "comparison",
					Description: `(Optional) Metric rule comparison. Supported values : ` + "`" + `"equal" | "greater-or-equal" | "greater-than" | "less-or-equal" | "less-than" | "not-equal"` + "`" + `. Default: ` + "`" + `equal` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Metric rule description (string) ### ` + "`" + `node_rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "cpu_threshold",
					Description: `(Optional) Node rule cpu threshold. Default: ` + "`" + `70` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional) Node rule condition. Supported values : ` + "`" + `"cpu" | "mem" | "notready"` + "`" + `. Default: ` + "`" + `notready` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "mem_threshold",
					Description: `(Optional) Node rule mem threshold. Default: ` + "`" + `70` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `(Optional) Node ID (string)`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) Node rule selector (map) ### ` + "`" + `system_service_rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional) System service rule condition. Supported values : ` + "`" + `"controller-manager" | "etcd" | "scheduler"` + "`" + `. Default: ` + "`" + `scheduler` + "`" + ` (string) ## Timeouts ` + "`" + `rancher2_cluster_alert_rule` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster alert rules. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster alert rule modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster alert rules. ## Import Cluster Alert Rule can be imported using the Rancher cluster alert rule ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_alert_rule.foo <rancher2_cluster_alert_rule_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `event_rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "resource_kind",
					Description: `(Required) Resource kind. Supported values : ` + "`" + `"DaemonSet" | "Deployment" | "Node" | "Pod" | "StatefulSet"` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `(Optional) Event type. Supported values : ` + "`" + `"Warning" | "Normal"` + "`" + `. Default: ` + "`" + `Warning` + "`" + ` (string) ### ` + "`" + `metric_rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) Metric rule duration (string)`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `(Required) Metric rule expression (string)`,
				},
				resource.Attribute{
					Name:        "threshold_value",
					Description: `(Required) Metric rule threshold value (float64)`,
				},
				resource.Attribute{
					Name:        "comparison",
					Description: `(Optional) Metric rule comparison. Supported values : ` + "`" + `"equal" | "greater-or-equal" | "greater-than" | "less-or-equal" | "less-than" | "not-equal"` + "`" + `. Default: ` + "`" + `equal` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Metric rule description (string) ### ` + "`" + `node_rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "cpu_threshold",
					Description: `(Optional) Node rule cpu threshold. Default: ` + "`" + `70` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional) Node rule condition. Supported values : ` + "`" + `"cpu" | "mem" | "notready"` + "`" + `. Default: ` + "`" + `notready` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "mem_threshold",
					Description: `(Optional) Node rule mem threshold. Default: ` + "`" + `70` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `(Optional) Node ID (string)`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) Node rule selector (map) ### ` + "`" + `system_service_rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional) System service rule condition. Supported values : ` + "`" + `"controller-manager" | "etcd" | "scheduler"` + "`" + `. Default: ` + "`" + `scheduler` + "`" + ` (string) ## Timeouts ` + "`" + `rancher2_cluster_alert_rule` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster alert rules. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster alert rule modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster alert rules. ## Import Cluster Alert Rule can be imported using the Rancher cluster alert rule ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_alert_rule.foo <rancher2_cluster_alert_rule_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cluster_driver",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Cluster Driver resource. This can be used to create Cluster Driver for Rancher v2 Kontainer Engine clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"cluster",
				"driver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "active",
					Description: `(Required) Specify the cluster driver state (bool)`,
				},
				resource.Attribute{
					Name:        "builtin",
					Description: `(Required) Specify whether the cluster driver is an internal cluster driver or not (bool)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the cluster driver (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL to download the machine driver binary for 64-bit Linux (string)`,
				},
				resource.Attribute{
					Name:        "actual_url",
					Description: `(Optional) Actual url of the cluster driver (string)`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `(Optional) Verify that the downloaded driver matches the expected checksum (string)`,
				},
				resource.Attribute{
					Name:        "ui_url",
					Description: `(Optional) The URL to load for customized Add Clusters screen for this driver (string)`,
				},
				resource.Attribute{
					Name:        "whitelist_domains",
					Description: `(Optional) Domains to whitelist for the ui (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_cluster_driver` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster drivers. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster driver modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster drivers. ## Import Cluster Driver can be imported using the Rancher Cluster Driver ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_driver.foo <cluster_driver_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_cluster_driver` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster drivers. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster driver modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster drivers. ## Import Cluster Driver can be imported using the Rancher Cluster Driver ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_driver.foo <cluster_driver_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cluster_logging",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Cluster Logging resource. This can be used to configure Cluster Logging for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"cluster",
				"logging",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The cluster id to configure logging (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the cluster logging config (string)`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind of the Cluster Logging. ` + "`" + `elasticsearch` + "`" + `, ` + "`" + `fluentd` + "`" + `, ` + "`" + `kafka` + "`" + `, ` + "`" + `splunk` + "`" + ` and ` + "`" + `syslog` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "elasticsearch_config",
					Description: `(Optional) The elasticsearch config for Cluster Logging. For ` + "`" + `kind = elasticsearch` + "`" + `. Conflicts with ` + "`" + `fluentd_config` + "`" + `, ` + "`" + `kafka_config` + "`" + `, ` + "`" + `splunk_config` + "`" + ` and ` + "`" + `syslog_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "fluentd_config",
					Description: `(Optional) The fluentd config for Cluster Logging. For ` + "`" + `kind = fluentd` + "`" + `. Conflicts with ` + "`" + `elasticsearch_config` + "`" + `, ` + "`" + `kafka_config` + "`" + `, ` + "`" + `splunk_config` + "`" + ` and ` + "`" + `syslog_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kafka_config",
					Description: `(Optional) The kafka config for Cluster Logging. For ` + "`" + `kind = kafka` + "`" + `. Conflicts with ` + "`" + `elasticsearch_config` + "`" + `, ` + "`" + `fluentd_config` + "`" + `, ` + "`" + `splunk_config` + "`" + ` and ` + "`" + `syslog_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Optional) The namespace id from cluster logging (string)`,
				},
				resource.Attribute{
					Name:        "output_flush_interval",
					Description: `(Optional) How often buffered logs would be flushed. Default: ` + "`" + `3` + "`" + ` seconds (int)`,
				},
				resource.Attribute{
					Name:        "output_tags",
					Description: `(Optional/computed) The output tags for Cluster Logging (map)`,
				},
				resource.Attribute{
					Name:        "splunk_config",
					Description: `(Optional) The splunk config for Cluster Logging. For ` + "`" + `kind = splunk` + "`" + `. Conflicts with ` + "`" + `elasticsearch_config` + "`" + `, ` + "`" + `fluentd_config` + "`" + `, ` + "`" + `kafka_config` + "`" + `, and ` + "`" + `syslog_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "syslog_config",
					Description: `(Optional) The syslog config for Cluster Logging. For ` + "`" + `kind = syslog` + "`" + `. Conflicts with ` + "`" + `elasticsearch_config` + "`" + `, ` + "`" + `fluentd_config` + "`" + `, ` + "`" + `kafka_config` + "`" + `, and ` + "`" + `splunk_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for Cluster Logging object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for Cluster Logging object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `elasticsearch_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the elascticsearch service. Must include protocol, ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "auth_password",
					Description: `(Optional/Sensitive) User password for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "auth_username",
					Description: `(Optional/Sensitive) Username for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_key_pass",
					Description: `(Optional/Sensitive) SSL client key password for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "date_format",
					Description: `(Optional) Date format for the elascticsearch logs. Default: ` + "`" + `YYYY-MM-DD` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "index_prefix",
					Description: `(Optional) Index prefix for the elascticsearch logs. Default: ` + "`" + `local` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the elascticsearch service (bool)`,
				},
				resource.Attribute{
					Name:        "ssl_version",
					Description: `(Optional) SSL version for the elascticsearch service (string) ### ` + "`" + `fluentd_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fluent_servers",
					Description: `(Required) Servers for the fluentd service (list)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "compress",
					Description: `(Optional) Compress data for the fluentd service (bool)`,
				},
				resource.Attribute{
					Name:        "enable_tls",
					Description: `(Optional) Enable TLS for the fluentd service (bool) #### ` + "`" + `fluent_servers` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Hostname of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) User password of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "shared_key",
					Description: `(Optional/Sensitive) Shared key of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "standby",
					Description: `(Optional) Standby server of the fluentd service (bool)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional/Sensitive) Username of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight of the fluentd server (int) ### ` + "`" + `kafka_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Required) Topic to publish on the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "broker_endpoints",
					Description: `(Optional) Kafka endpoints for kafka service. Conflicts with ` + "`" + `zookeeper_endpoint` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "zookeeper_endpoint",
					Description: `(Optional) Zookeeper endpoint for kafka service. Conflicts with ` + "`" + `broker_endpoints` + "`" + ` (string) ### ` + "`" + `splunk_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the splunk service. Must include protocol, ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required/Sensitive) Token for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_key_pass",
					Description: `(Optional/Sensitive) SSL client key password for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) Index prefix for the splunk logs (string)`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Date format for the splunk logs (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the splunk service (bool) ### ` + "`" + `syslog_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "program",
					Description: `(Optional) Program for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol for the syslog service. ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + ` are supported. Default: ` + "`" + `udp` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Optional) Date format for the syslog logs. ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notice` + "`" + `, ` + "`" + `info` + "`" + ` and ` + "`" + `debug` + "`" + ` are supported. Default: ` + "`" + `notice` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the syslog service (bool)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional/Sensitive) Token for the syslog service (string) ## Timeouts ` + "`" + `rancher2_cluster_logging` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster logging configurations. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster logging configuration modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster logging configurations. ## Import Cluster Logging can be imported using the Rancher Cluster Logging ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_logging.foo <cluster_logging_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `elasticsearch_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the elascticsearch service. Must include protocol, ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "auth_password",
					Description: `(Optional/Sensitive) User password for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "auth_username",
					Description: `(Optional/Sensitive) Username for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_key_pass",
					Description: `(Optional/Sensitive) SSL client key password for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "date_format",
					Description: `(Optional) Date format for the elascticsearch logs. Default: ` + "`" + `YYYY-MM-DD` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "index_prefix",
					Description: `(Optional) Index prefix for the elascticsearch logs. Default: ` + "`" + `local` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the elascticsearch service (bool)`,
				},
				resource.Attribute{
					Name:        "ssl_version",
					Description: `(Optional) SSL version for the elascticsearch service (string) ### ` + "`" + `fluentd_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fluent_servers",
					Description: `(Required) Servers for the fluentd service (list)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "compress",
					Description: `(Optional) Compress data for the fluentd service (bool)`,
				},
				resource.Attribute{
					Name:        "enable_tls",
					Description: `(Optional) Enable TLS for the fluentd service (bool) #### ` + "`" + `fluent_servers` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Hostname of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) User password of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "shared_key",
					Description: `(Optional/Sensitive) Shared key of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "standby",
					Description: `(Optional) Standby server of the fluentd service (bool)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional/Sensitive) Username of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight of the fluentd server (int) ### ` + "`" + `kafka_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Required) Topic to publish on the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "broker_endpoints",
					Description: `(Optional) Kafka endpoints for kafka service. Conflicts with ` + "`" + `zookeeper_endpoint` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "zookeeper_endpoint",
					Description: `(Optional) Zookeeper endpoint for kafka service. Conflicts with ` + "`" + `broker_endpoints` + "`" + ` (string) ### ` + "`" + `splunk_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the splunk service. Must include protocol, ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required/Sensitive) Token for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_key_pass",
					Description: `(Optional/Sensitive) SSL client key password for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) Index prefix for the splunk logs (string)`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Date format for the splunk logs (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the splunk service (bool) ### ` + "`" + `syslog_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "program",
					Description: `(Optional) Program for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol for the syslog service. ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + ` are supported. Default: ` + "`" + `udp` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Optional) Date format for the syslog logs. ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notice` + "`" + `, ` + "`" + `info` + "`" + ` and ` + "`" + `debug` + "`" + ` are supported. Default: ` + "`" + `notice` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the syslog service (bool)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional/Sensitive) Token for the syslog service (string) ## Timeouts ` + "`" + `rancher2_cluster_logging` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster logging configurations. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster logging configuration modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster logging configurations. ## Import Cluster Logging can be imported using the Rancher Cluster Logging ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_logging.foo <cluster_logging_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cluster_role_template_binding",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Cluster Role Template Binding resource. This can be used to create Cluster Role Template Bindings for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"cluster",
				"role",
				"template",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The cluster id where bind cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "role_template_id",
					Description: `(Required) The role template id from create cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) The group ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "group_principal_id",
					Description: `(Optional) The group_principal ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional) The user ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_principal_id",
					Description: `(Optional) The user_principal ID to assign cluster role template binding (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for cluster role template binding (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for cluster role template binding (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_cluster_role_template_binding` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster role template bindings. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster role template binding modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster role template bindings. ## Import Cluster Role Template Bindings can be imported using the Rancher cluster Role Template Binding ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_role_template_binding.foo <cluster_role_template_binding_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_cluster_role_template_binding` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster role template bindings. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster role template binding modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster role template bindings. ## Import Cluster Role Template Bindings can be imported using the Rancher cluster Role Template Binding ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_role_template_binding.foo <cluster_role_template_binding_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_cluster_template",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Cluster Template resource. This can be used to create Cluster Templates for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"cluster",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The cluster template name (string)`,
				},
				resource.Attribute{
					Name:        "decription",
					Description: `(Optional) The cluster template description (string)`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Cluster template members (list)`,
				},
				resource.Attribute{
					Name:        "template_revisions",
					Description: `(Optional/Computed) Cluster template revisions (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for the cluster template (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for the cluster template (map) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "default_revision_id",
					Description: `(Computed) Default cluster template revision ID (string) ## Nested blocks ### ` + "`" + `members` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) Member access type. Valid values: ` + "`" + `["read-only" | "owner"]` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_principal_id",
					Description: `(Optional) Member group principal id (string)`,
				},
				resource.Attribute{
					Name:        "user_principal_id",
					Description: `(Optional) Member user principal id (string) ### ` + "`" + `template_revisions` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The cluster template revision name (string)`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `(Optional) Cluster configuration (list maxitem: 1)`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) Default cluster template revision. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable cluster template revision. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "questions",
					Description: `(Optional) Cluster template questions (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for the cluster template revision (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for the cluster template revision (map) #### Attributes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The cluster template revision ID (string)`,
				},
				resource.Attribute{
					Name:        "cluster_template_id",
					Description: `(Computed) Cluster template ID (string) #### ` + "`" + `cluster_config` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_auth_endpoint",
					Description: `(Optional/Computed) Local cluster auth endpoint (list maxitems: 1)`,
				},
				resource.Attribute{
					Name:        "default_cluster_role_for_project_members",
					Description: `(Optional/Computed) Default cluster role for project members (string)`,
				},
				resource.Attribute{
					Name:        "default_pod_security_policy_template_id",
					Description: `(Optional/Computed) Default pod security policy template ID (string)`,
				},
				resource.Attribute{
					Name:        "desired_agent_image",
					Description: `(Optional/Computed) Desired agent image (string)`,
				},
				resource.Attribute{
					Name:        "desired_auth_image",
					Description: `(Optional/Computed) Desired auth image (string)`,
				},
				resource.Attribute{
					Name:        "docker_root_dir",
					Description: `(Optional/Computed) Desired auth image (string)`,
				},
				resource.Attribute{
					Name:        "enable_cluster_alerting",
					Description: `(Optional) Enable built-in cluster alerting. Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_cluster_monitoring",
					Description: `(Optional) Enable built-in cluster monitoring. Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_network_policy",
					Description: `(Optional) Enable project network isolation. Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "rke_config",
					Description: `(Optional/Computed) Rancher Kubernetes Engine Config (list maxitems: 1)`,
				},
				resource.Attribute{
					Name:        "windows_prefered_cluster",
					Description: `(Optional) Windows prefered cluster. Default: ` + "`" + `false` + "`" + ` (bool) #### ` + "`" + `questions` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) Default variable value (string)`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional) Required variable. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Variable type. ` + "`" + `boolean` + "`" + `, ` + "`" + `int` + "`" + ` and ` + "`" + `string` + "`" + ` are allowed. Default ` + "`" + `string` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "variable",
					Description: `(Optional) Variable name (string) ## Timeouts ` + "`" + `rancher2_cluster_template` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster templates. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster template modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster templates. ## Import Cluster Template can be imported using the rancher Cluster Template ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_template.foo <cluster_template_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "default_revision_id",
					Description: `(Computed) Default cluster template revision ID (string) ## Nested blocks ### ` + "`" + `members` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) Member access type. Valid values: ` + "`" + `["read-only" | "owner"]` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_principal_id",
					Description: `(Optional) Member group principal id (string)`,
				},
				resource.Attribute{
					Name:        "user_principal_id",
					Description: `(Optional) Member user principal id (string) ### ` + "`" + `template_revisions` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The cluster template revision name (string)`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `(Optional) Cluster configuration (list maxitem: 1)`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) Default cluster template revision. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable cluster template revision. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "questions",
					Description: `(Optional) Cluster template questions (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for the cluster template revision (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for the cluster template revision (map) #### Attributes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The cluster template revision ID (string)`,
				},
				resource.Attribute{
					Name:        "cluster_template_id",
					Description: `(Computed) Cluster template ID (string) #### ` + "`" + `cluster_config` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_auth_endpoint",
					Description: `(Optional/Computed) Local cluster auth endpoint (list maxitems: 1)`,
				},
				resource.Attribute{
					Name:        "default_cluster_role_for_project_members",
					Description: `(Optional/Computed) Default cluster role for project members (string)`,
				},
				resource.Attribute{
					Name:        "default_pod_security_policy_template_id",
					Description: `(Optional/Computed) Default pod security policy template ID (string)`,
				},
				resource.Attribute{
					Name:        "desired_agent_image",
					Description: `(Optional/Computed) Desired agent image (string)`,
				},
				resource.Attribute{
					Name:        "desired_auth_image",
					Description: `(Optional/Computed) Desired auth image (string)`,
				},
				resource.Attribute{
					Name:        "docker_root_dir",
					Description: `(Optional/Computed) Desired auth image (string)`,
				},
				resource.Attribute{
					Name:        "enable_cluster_alerting",
					Description: `(Optional) Enable built-in cluster alerting. Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_cluster_monitoring",
					Description: `(Optional) Enable built-in cluster monitoring. Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "enable_network_policy",
					Description: `(Optional) Enable project network isolation. Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "rke_config",
					Description: `(Optional/Computed) Rancher Kubernetes Engine Config (list maxitems: 1)`,
				},
				resource.Attribute{
					Name:        "windows_prefered_cluster",
					Description: `(Optional) Windows prefered cluster. Default: ` + "`" + `false` + "`" + ` (bool) #### ` + "`" + `questions` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) Default variable value (string)`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional) Required variable. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Variable type. ` + "`" + `boolean` + "`" + `, ` + "`" + `int` + "`" + ` and ` + "`" + `string` + "`" + ` are allowed. Default ` + "`" + `string` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "variable",
					Description: `(Optional) Variable name (string) ## Timeouts ` + "`" + `rancher2_cluster_template` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cluster templates. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cluster template modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cluster templates. ## Import Cluster Template can be imported using the rancher Cluster Template ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_cluster_template.foo <cluster_template_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_etcd_backup",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2.2.x Etcd Backup resource. This can be used to create Etcd Backup for Rancher v2.2 node templates and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"etcd",
				"backup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID to config Etcd Backup (string)`,
				},
				resource.Attribute{
					Name:        "backup_config",
					Description: `(Optional/Computed) Backup config for etcd backup (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Optional/Computed) Filename of the Etcd Backup (string)`,
				},
				resource.Attribute{
					Name:        "manual",
					Description: `(Optional) Manual execution of the Etcd Backup. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Etcd Backup (string)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Optional/Computed) Description for the Etcd Backup (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Annotations for Etcd Backup object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for Etcd Backup object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `backup_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable etcd backup (bool)`,
				},
				resource.Attribute{
					Name:        "interval_hours",
					Description: `(Optional) Interval hours for etcd backup. Default ` + "`" + `12` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Optional) Retention for etcd backup. Default ` + "`" + `6` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "s3_backup_config",
					Description: `(Optional) S3 config options for etcd backup. Valid for ` + "`" + `imported` + "`" + ` and ` + "`" + `rke` + "`" + ` clusters. (list maxitems:1) #### ` + "`" + `s3_backup_config` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional/Sensitive) Access key for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) Bucket name for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "custom_ca",
					Description: `(Optional) Base64 encoded custom CA for S3 service. Use filebase64(<FILE>) for encoding file. Available from Rancher v2.2.5 (string)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) Folder for S3 service. Available from Rancher v2.2.7 (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional/Sensitive) Secret key for S3 service (string) ## Timeouts ` + "`" + `rancher2_etcd_backup` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cloud credentials. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cloud credential modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cloud credentials. ## Import Etcd Backup can be imported using the Rancher etcd backup ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_etcd_backup.foo <etcd_backup_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `backup_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable etcd backup (bool)`,
				},
				resource.Attribute{
					Name:        "interval_hours",
					Description: `(Optional) Interval hours for etcd backup. Default ` + "`" + `12` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Optional) Retention for etcd backup. Default ` + "`" + `6` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "s3_backup_config",
					Description: `(Optional) S3 config options for etcd backup. Valid for ` + "`" + `imported` + "`" + ` and ` + "`" + `rke` + "`" + ` clusters. (list maxitems:1) #### ` + "`" + `s3_backup_config` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional/Sensitive) Access key for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) Bucket name for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "custom_ca",
					Description: `(Optional) Base64 encoded custom CA for S3 service. Use filebase64(<FILE>) for encoding file. Available from Rancher v2.2.5 (string)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) Folder for S3 service. Available from Rancher v2.2.7 (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional/Sensitive) Secret key for S3 service (string) ## Timeouts ` + "`" + `rancher2_etcd_backup` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating cloud credentials. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for cloud credential modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting cloud credentials. ## Import Etcd Backup can be imported using the Rancher etcd backup ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_etcd_backup.foo <etcd_backup_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_global_role_binding",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Global Role Binding resource. This can be used to create Global Role Bindings for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"role",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "global_role_id",
					Description: `(Required/ForceNew) The role id from create global role binding (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required/ForceNew) The user ID to assign global role binding (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional/Computed/ForceNew) The name of the global role binding (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for global role binding (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for global role binding (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_global_role_binding` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for creating global role bindings. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for global role binding modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for deleting global role bindings. ## Import Global Role Bindings can be imported using the Rancher Global Role Binding ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_global_role_binding.foo <global_role_binding_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_global_role_binding` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for creating global role bindings. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for global role binding modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for deleting global role bindings. ## Import Global Role Bindings can be imported using the Rancher Global Role Binding ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_global_role_binding.foo <global_role_binding_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_multi_cluster_app",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 multi_cluster_app resource. This can be used to deploy multi cluster apps on Rancher v2.`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cluster",
				"app",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Required) The multi cluster app catalog name (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required/ForceNew) The multi cluster app name (string)`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) The multi cluster app roles (list)`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) The multi cluster app target projects (list)`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) The multi cluster app template name (string)`,
				},
				resource.Attribute{
					Name:        "answers",
					Description: `(Optional/Computed) The multi cluster app answers (list)`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) The multi cluster app answers (list)`,
				},
				resource.Attribute{
					Name:        "revision_history_limit",
					Description: `(Computed) The multi cluster app revision history limit. Default ` + "`" + `10` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "revision_id",
					Description: `(Optional/Computed) Current revision id for the multi cluster app (string)`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `(Optional/Computed) The multi cluster app template version. Default: ` + "`" + `latest` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "upgrade_strategy",
					Description: `(Optional/Computed) The multi cluster app upgrade strategy (list MaxItems:1)`,
				},
				resource.Attribute{
					Name:        "wait",
					Description: `(Optional) Wait until the multi cluster app is active. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for multi cluster app object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for multi cluster app object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "template_version_id",
					Description: `(Computed) The multi cluster app template version ID (string) ## Nested blocks ### ` + "`" + `targets` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID for target (string)`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Computed) App ID for target (string)`,
				},
				resource.Attribute{
					Name:        "health_state",
					Description: `(Computed) App health state for target (string)`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Computed) App state for target (string) ### ` + "`" + `answers` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) Cluster ID for answer (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID for target (string)`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Key/values for answer (map) ### ` + "`" + `members` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) Member access type. Valid values: ` + "`" + `["member" | "owner" | "read-only"]` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_principal_id",
					Description: `(Optional) Member group principal id (string)`,
				},
				resource.Attribute{
					Name:        "user_principal_id",
					Description: `(Optional) Member user principal id (string) ### ` + "`" + `upgrade_strategy` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "rolling_update",
					Description: `(Optional) Upgrade strategy rolling update (list MaxItems:1) #### ` + "`" + `rolling_update` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "batch_size",
					Description: `(Optional) Rolling update batch size. Default ` + "`" + `1` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Rolling update interval. Default ` + "`" + `1` + "`" + ` (int) ## Timeouts ` + "`" + `rancher2_app` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating apps. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for app modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting apps. ## Import Multi cluster app can be imported using the multi cluster app ID in the format ` + "`" + `<multi_cluster_app_name>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_multi_cluster_app.foo <multi_cluster_app_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "template_version_id",
					Description: `(Computed) The multi cluster app template version ID (string) ## Nested blocks ### ` + "`" + `targets` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID for target (string)`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Computed) App ID for target (string)`,
				},
				resource.Attribute{
					Name:        "health_state",
					Description: `(Computed) App health state for target (string)`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Computed) App state for target (string) ### ` + "`" + `answers` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) Cluster ID for answer (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID for target (string)`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Key/values for answer (map) ### ` + "`" + `members` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) Member access type. Valid values: ` + "`" + `["member" | "owner" | "read-only"]` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "group_principal_id",
					Description: `(Optional) Member group principal id (string)`,
				},
				resource.Attribute{
					Name:        "user_principal_id",
					Description: `(Optional) Member user principal id (string) ### ` + "`" + `upgrade_strategy` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "rolling_update",
					Description: `(Optional) Upgrade strategy rolling update (list MaxItems:1) #### ` + "`" + `rolling_update` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "batch_size",
					Description: `(Optional) Rolling update batch size. Default ` + "`" + `1` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Rolling update interval. Default ` + "`" + `1` + "`" + ` (int) ## Timeouts ` + "`" + `rancher2_app` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating apps. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for app modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting apps. ## Import Multi cluster app can be imported using the multi cluster app ID in the format ` + "`" + `<multi_cluster_app_name>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_multi_cluster_app.foo <multi_cluster_app_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_namespace",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Namespace resource. This can be used to create namespaces for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the namespace (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project id where assign namespace. It's on the form ` + "`" + `project_id=<cluster_id>:<id>` + "`" + `. Updating ` + "`" + `<id>` + "`" + ` part on same ` + "`" + `<cluster_id>` + "`" + ` namespace will be moved between projects (string)`,
				},
				resource.Attribute{
					Name:        "container_resource_limit",
					Description: `(Optional) Default containers resource limits on namespace (List maxitem:1)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A namespace description (string)`,
				},
				resource.Attribute{
					Name:        "resource_quota",
					Description: `(Optional) Resource quota for namespace. Rancher v2.1.x or higher (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "wait_for_cluster",
					Description: `(Optional) Wait for cluster becomes active. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for Node Pool object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for Node Pool object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `container_resource_limit` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "limits_cpu",
					Description: `(Optional) CPU limit for containers (string)`,
				},
				resource.Attribute{
					Name:        "limits_memory",
					Description: `(Optional) Memory limit for containers (string)`,
				},
				resource.Attribute{
					Name:        "requests_cpu",
					Description: `(Optional) CPU reservation for containers (string)`,
				},
				resource.Attribute{
					Name:        "requests_memory",
					Description: `(Optional) Memory reservation for containers (string) ### ` + "`" + `resource_quota` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Required) Resource quota limit for namespace (list maxitems:1) #### ` + "`" + `limit` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "config_maps",
					Description: `(Optional) Limit for config maps in namespace (string)`,
				},
				resource.Attribute{
					Name:        "limits_cpu",
					Description: `(Optional) Limit for limits cpu in namespace (string)`,
				},
				resource.Attribute{
					Name:        "limits_memory",
					Description: `(Optional) Limit for limits memory in namespace (string)`,
				},
				resource.Attribute{
					Name:        "persistent_volume_claims",
					Description: `(Optional) Limit for persistent volume claims in namespace (string)`,
				},
				resource.Attribute{
					Name:        "pods",
					Description: `(Optional) Limit for pods in namespace (string)`,
				},
				resource.Attribute{
					Name:        "replication_controllers",
					Description: `(Optional) Limit for replication controllers in namespace (string)`,
				},
				resource.Attribute{
					Name:        "requests_cpu",
					Description: `(Optional) Limit for requests cpu in namespace (string)`,
				},
				resource.Attribute{
					Name:        "requests_memory",
					Description: `(Optional) Limit for requests memory in namespace (string)`,
				},
				resource.Attribute{
					Name:        "requests_storage",
					Description: `(Optional) Limit for requests storage in namespace (string)`,
				},
				resource.Attribute{
					Name:        "secrets",
					Description: `(Optional) Limit for secrets in namespace (string)`,
				},
				resource.Attribute{
					Name:        "services_load_balancers",
					Description: `(Optional) Limit for services load balancers in namespace (string)`,
				},
				resource.Attribute{
					Name:        "services_node_ports",
					Description: `(Optional) Limit for services node ports in namespace (string) More info at [resource-quotas](https://rancher.com/docs/rancher/v2.x/en/k8s-in-rancher/projects-and-namespaces/resource-quotas/) ## Timeouts ` + "`" + `rancher2_namespace` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating namespaces. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for namespace modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting namespaces. ## Import Namespaces can be imported using the namespace ID in the format ` + "`" + `<project_id>.<namespace_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_namespace.foo <project_id>.<namespace_id> ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `<project_id>` + "`" + ` is in the format ` + "`" + `<cluster_id>:<id>` + "`" + `, but <id> part is optional: - If full project_id is provided, ` + "`" + `<project_id>=<cluster_id>:<id>` + "`" + `, the namespace'll be assigned to corresponding cluster project once it's imported. - If ` + "`" + `<id>` + "`" + ` part is omitted ` + "`" + `<project_id>=<cluster_id>` + "`" + `, the namespace'll not be assigned to any project. To move it into a project, ` + "`" + `<project_id>=<cluster_id>:<id>` + "`" + ` needs to be updated in tf file. Namespace movement is only supported inside same ` + "`" + `cluster_id` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `container_resource_limit` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "limits_cpu",
					Description: `(Optional) CPU limit for containers (string)`,
				},
				resource.Attribute{
					Name:        "limits_memory",
					Description: `(Optional) Memory limit for containers (string)`,
				},
				resource.Attribute{
					Name:        "requests_cpu",
					Description: `(Optional) CPU reservation for containers (string)`,
				},
				resource.Attribute{
					Name:        "requests_memory",
					Description: `(Optional) Memory reservation for containers (string) ### ` + "`" + `resource_quota` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Required) Resource quota limit for namespace (list maxitems:1) #### ` + "`" + `limit` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "config_maps",
					Description: `(Optional) Limit for config maps in namespace (string)`,
				},
				resource.Attribute{
					Name:        "limits_cpu",
					Description: `(Optional) Limit for limits cpu in namespace (string)`,
				},
				resource.Attribute{
					Name:        "limits_memory",
					Description: `(Optional) Limit for limits memory in namespace (string)`,
				},
				resource.Attribute{
					Name:        "persistent_volume_claims",
					Description: `(Optional) Limit for persistent volume claims in namespace (string)`,
				},
				resource.Attribute{
					Name:        "pods",
					Description: `(Optional) Limit for pods in namespace (string)`,
				},
				resource.Attribute{
					Name:        "replication_controllers",
					Description: `(Optional) Limit for replication controllers in namespace (string)`,
				},
				resource.Attribute{
					Name:        "requests_cpu",
					Description: `(Optional) Limit for requests cpu in namespace (string)`,
				},
				resource.Attribute{
					Name:        "requests_memory",
					Description: `(Optional) Limit for requests memory in namespace (string)`,
				},
				resource.Attribute{
					Name:        "requests_storage",
					Description: `(Optional) Limit for requests storage in namespace (string)`,
				},
				resource.Attribute{
					Name:        "secrets",
					Description: `(Optional) Limit for secrets in namespace (string)`,
				},
				resource.Attribute{
					Name:        "services_load_balancers",
					Description: `(Optional) Limit for services load balancers in namespace (string)`,
				},
				resource.Attribute{
					Name:        "services_node_ports",
					Description: `(Optional) Limit for services node ports in namespace (string) More info at [resource-quotas](https://rancher.com/docs/rancher/v2.x/en/k8s-in-rancher/projects-and-namespaces/resource-quotas/) ## Timeouts ` + "`" + `rancher2_namespace` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating namespaces. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for namespace modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting namespaces. ## Import Namespaces can be imported using the namespace ID in the format ` + "`" + `<project_id>.<namespace_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_namespace.foo <project_id>.<namespace_id> ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `<project_id>` + "`" + ` is in the format ` + "`" + `<cluster_id>:<id>` + "`" + `, but <id> part is optional: - If full project_id is provided, ` + "`" + `<project_id>=<cluster_id>:<id>` + "`" + `, the namespace'll be assigned to corresponding cluster project once it's imported. - If ` + "`" + `<id>` + "`" + ` part is omitted ` + "`" + `<project_id>=<cluster_id>` + "`" + `, the namespace'll not be assigned to any project. To move it into a project, ` + "`" + `<project_id>=<cluster_id>:<id>` + "`" + ` needs to be updated in tf file. Namespace movement is only supported inside same ` + "`" + `cluster_id` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_node_driver",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Node Driver resource. This can be used to create Node Driver for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"node",
				"driver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "active",
					Description: `(Required) Specify if the node driver state (bool)`,
				},
				resource.Attribute{
					Name:        "builtin",
					Description: `(Required) Specify wheter the node driver is an internal node driver or not (bool)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the node driver (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL to download the machine driver binary for 64-bit Linux (string)`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `(Optional) Verify that the downloaded driver matches the expected checksum (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the node driver (string)`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Optional) External ID (string)`,
				},
				resource.Attribute{
					Name:        "ui_url",
					Description: `(Optional) The URL to load for customized Add Nodes screen for this driver (string)`,
				},
				resource.Attribute{
					Name:        "whitelist_domains",
					Description: `(Optional) Domains to whitelist for the ui (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_node_driver` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating node drivers. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for node driver modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting node drivers. ## Import Node Driver can be imported using the Rancher Node Driver ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_node_driver.foo <node_driver_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_node_driver` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating node drivers. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for node driver modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting node drivers. ## Import Node Driver can be imported using the Rancher Node Driver ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_node_driver.foo <node_driver_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_node_pool",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Node Pool resource. This can be used to create Node pool, using Node template for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"node",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The RKE cluster id to use Node Pool (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Node Pool (string)`,
				},
				resource.Attribute{
					Name:        "hostname_prefix",
					Description: `(Required) The prefix for created nodes of the Node Pool (string)`,
				},
				resource.Attribute{
					Name:        "node_template_id",
					Description: `(Required) The Node Template ID to use for node creation (string)`,
				},
				resource.Attribute{
					Name:        "control_plane",
					Description: `(Optional) RKE control plane role for created nodes (bool)`,
				},
				resource.Attribute{
					Name:        "etcd",
					Description: `(Optional) RKE etcd role for created nodes (bool)`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Optional) The number of nodes to create on Node Pool. Default ` + "`" + `1` + "`" + `. Only values >= 1 allowed (int)`,
				},
				resource.Attribute{
					Name:        "worker",
					Description: `(Optional) RKE role role for created nodes (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for Node Pool object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for Node Pool object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_node_pool` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating node pools. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for node pool modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting node pools. ## Import Node Pool can be imported using the Rancher Node Pool ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_node_pool.foo <node_pool_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_node_pool` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating node pools. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for node pool modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting node pools. ## Import Node Pool can be imported using the Rancher Node Pool ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_node_pool.foo <node_pool_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_node_template",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Node Template resource. This can be used to create Node template for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"node",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Node Template (string)`,
				},
				resource.Attribute{
					Name:        "amazonec2_config",
					Description: `(Optional) AWS config for the Node Template (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "auth_certificate_authority",
					Description: `(Optional/Sensitive) Auth certificate authority for the Node Template (string)`,
				},
				resource.Attribute{
					Name:        "auth_key",
					Description: `(Optional/Sensitive) Auth key for the Node Template (string)`,
				},
				resource.Attribute{
					Name:        "azure_config",
					Description: `(Optional) Azure config for the Node Template (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "cloud_credential_id",
					Description: `(Optional) Cloud credential ID for the Node Template. Required from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the Node Template (string)`,
				},
				resource.Attribute{
					Name:        "digitalocean_config",
					Description: `(Optional) Digitalocean config for the Node Template (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "engine_env",
					Description: `(Optional) Engine environment for the node template (string)`,
				},
				resource.Attribute{
					Name:        "engine_insecure_registry",
					Description: `(Optional) Insecure registry for the node template (list)`,
				},
				resource.Attribute{
					Name:        "engine_install_url",
					Description: `(Optional) Docker engine install URL for the node template. Default ` + "`" + `https://releases.rancher.com/install-docker/18.09.sh` + "`" + `. Available install docker versions at ` + "`" + `https://github.com/rancher/install-docker` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "engine_label",
					Description: `(Optional) Engine label for the node template (string)`,
				},
				resource.Attribute{
					Name:        "engine_opt",
					Description: `(Optional) Engine options for the node template (map)`,
				},
				resource.Attribute{
					Name:        "engine_registry_mirror",
					Description: `(Optional) Engine registry mirror for the node template (list)`,
				},
				resource.Attribute{
					Name:        "engine_storage_driver",
					Description: `(Optional) Engine storage driver for the node template (string)`,
				},
				resource.Attribute{
					Name:        "openstack_config",
					Description: `(Optional) Openstack config for the Node Template (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "use_internal_ip_address",
					Description: `(Optional) Engine storage driver for the node template (bool)`,
				},
				resource.Attribute{
					Name:        "vsphere_config",
					Description: `(Optional) vSphere config for the Node Template (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Annotations for Node Template object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for Node Template object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Computed) The driver of the node template (string) ## Nested blocks ### ` + "`" + `amazonec2_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ami",
					Description: `(Required) AWS machine image (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) AWS region. (string)`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Required) AWS VPC security group. (list)`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) AWS VPC subnet id (string)`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) AWS VPC id. (string)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) AWS zone for instance (i.e. a,b,c,d,e) (string)`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional/Sensitive) AWS access key. Required on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "block_duration_minutes",
					Description: `(Optional) AWS spot instance duration in minutes (60, 120, 180, 240, 300, or 360). Default ` + "`" + `0` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) AWS root device name. Default ` + "`" + `/dev/sda1` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) Optional endpoint URL (hostname only or fully qualified URI) (string)`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `(Optional) AWS IAM Instance Profile (string)`,
				},
				resource.Attribute{
					Name:        "insecure_transport",
					Description: `(Optional) Disable SSL when sending requests (bool)`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) AWS instance type. Default ` + "`" + `t2.micro` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `(Optional) AWS keypair to use; requires --amazonec2-ssh-keypath (string)`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional) Set this flag to enable CloudWatch monitoring. Deafult ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "open_port",
					Description: `(Optional) Make the specified port number accessible from the Internet. (list)`,
				},
				resource.Attribute{
					Name:        "private_address_only",
					Description: `(Optional) Only use a private IP address. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "request_spot_instance",
					Description: `(Optional) Set this flag to request spot instance. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Set retry count for recoverable failures (use -1 to disable). Default ` + "`" + `5` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "root_size",
					Description: `(Optional) AWS root disk size (in GB). Default ` + "`" + `16` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional/Sensitive) AWS secret key. Required on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "security_group_readonly",
					Description: `(Optional) Skip adding default rules to security groups (bool)`,
				},
				resource.Attribute{
					Name:        "session_token",
					Description: `(Optional/Sensitive) AWS Session Token (string)`,
				},
				resource.Attribute{
					Name:        "spot_price",
					Description: `(Optional) AWS spot instance bid price (in dollar). Default ` + "`" + `0.50` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_keypath",
					Description: `(Optional) SSH Key for Instance (string)`,
				},
				resource.Attribute{
					Name:        "ssh_user",
					Description: `(Optional) Set the name of the ssh user (string)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) AWS Tags (e.g. key1,value1,key2,value2) (string)`,
				},
				resource.Attribute{
					Name:        "use_ebs_optimized_instance",
					Description: `(Optional) Create an EBS optimized instance. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "use_private_address",
					Description: `(Optional) Force the usage of private IP address. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "userdata",
					Description: `(Optional) Path to file with cloud-init user data (string)`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) Amazon EBS volume type. Default ` + "`" + `gp2` + "`" + ` (string) ### ` + "`" + `azure_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional/Sensitive) Azure Service Principal Account ID. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional/Sensitive) Azure Service Principal Account password. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Optional/Sensitive) Azure Subscription ID. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "availability_set",
					Description: `(Optional) Azure Availability Set to place the virtual machine into. Default ` + "`" + `docker-machine` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `(Optional) Path to file with custom-data (string)`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) Disk size if using managed disk. Just for Rancher v2.3.x and above. Default ` + "`" + `30` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional) A unique DNS label for the public IP adddress (string)`,
				},
				resource.Attribute{
					Name:        "docker_port",
					Description: `(Optional) Port number for Docker engine. Default ` + "`" + `2376` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) Azure environment (e.g. AzurePublicCloud, AzureChinaCloud). Default ` + "`" + `AzurePublicCloud` + "`" + ` (string) ` + "`" + `fault_domain_count` + "`" + ` - (Optional) Fault domain count to use for availability set. Default ` + "`" + `3` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Azure virtual machine OS image. Default ` + "`" + `canonical:UbuntuServer:18.04-LTS:latest` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Azure region to create the virtual machine. Default ` + "`" + `westus` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "managed_disks",
					Description: `(Optional) Configures VM and availability set for managed disks. Just for Rancher v2.3.x and above. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "no_public_ip",
					Description: `(Optional) Do not create a public IP address for the machine. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "open_port",
					Description: `(Optional) Make the specified port number accessible from the Internet. (list)`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `(Optional) Specify a static private IP address for the machine. (string)`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Optional) Azure Resource Group name (will be created if missing). Default ` + "`" + `docker-machine` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size for Azure Virtual Machine. Default ` + "`" + `Standard_A2` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_user",
					Description: `(Optional) Username for SSH login (string)`,
				},
				resource.Attribute{
					Name:        "static_public_ip",
					Description: `(Optional) Assign a static public IP address to the machine. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) Type of Storage Account to host the OS Disk for the machine. Default ` + "`" + `Standard_LRS` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) Azure Subnet Name to be used within the Virtual Network. Default ` + "`" + `docker-machine` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "subnet_prefix",
					Description: `(Optional) Private CIDR block to be used for the new subnet, should comply RFC 1918. Default ` + "`" + `192.168.0.0/16` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "update_domain_count",
					Description: `(Optional) Update domain count to use for availability set. Default ` + "`" + `5` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "use_private_ip",
					Description: `(Optional) Use private IP address of the machine to connect. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "vnet",
					Description: `(Optional) Azure Virtual Network name to connect the virtual machine (in [resourcegroup:]name format). Default ` + "`" + `docker-machine-vnet` + "`" + ` (string) ### ` + "`" + `digitalocean_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `(Optional/Sensitive) Digital Ocean access token. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `(Optional) Enable backups for droplet. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Digital Ocean Image. Default ` + "`" + `ubuntu-16-04-x64` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) Enable ipv6 for droplet. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional) Enable monitoring for droplet. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `(Optional) Enable private networking for droplet. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Digital Ocean region. Default ` + "`" + `nyc3` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Digital Ocean size. Default ` + "`" + `s-1vcpu-1gb` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_fingerprint",
					Description: `(Optional/Sensitive) SSH key fingerprint (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional) SSH private key path (string)`,
				},
				resource.Attribute{
					Name:        "ssh_port",
					Description: `(Optional) SSH port. Default ` + "`" + `22` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_user",
					Description: `(Optional) SSH username. Default ` + "`" + `root` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Comma-separated list of tags to apply to the Droplet (string)`,
				},
				resource.Attribute{
					Name:        "userdata",
					Description: `(Optional) Path to file with cloud-init user-data (string) ### ` + "`" + `openstack_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "auth_url",
					Description: `(Required) OpenStack authentication URL (string)`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) OpenStack availability zone (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) OpenStack region name (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) OpenStack username (string)`,
				},
				resource.Attribute{
					Name:        "cacert",
					Description: `(Optional) CA certificate bundle to verify against (string)`,
				},
				resource.Attribute{
					Name:        "config_drive",
					Description: `(Optional) Enables the OpenStack config drive for the instance. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "endpoint_type",
					Description: `(Optional) OpenStack endpoint type. adminURL, internalURL or publicURL (string)`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "floating_ip_pool",
					Description: `(Optional) OpenStack floating IP pool to get an IP from to assign to the instance (string)`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "insecure",
					Description: `(Optional) Disable TLS credential checking. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) OpenStack version of IP address assigned for the machine Default ` + "`" + `4` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `(Optional) OpenStack keypair to use to SSH to the instance (string)`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "net_name",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "nova_network",
					Description: `(Optional) Use the nova networking services instead of neutron (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) OpenStack password. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "private_key_file",
					Description: `(Optional) Private keyfile absolute path to use for SSH (string)`,
				},
				resource.Attribute{
					Name:        "sec_groups",
					Description: `(Optional) OpenStack comma separated security groups for the machine (string)`,
				},
				resource.Attribute{
					Name:        "ssh_port",
					Description: `(Optional) OpenStack SSH port`,
				},
				resource.Attribute{
					Name:        "ssh_user",
					Description: `(Optional) OpenStack SSH user`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "user_data_file",
					Description: `(Optional) File containing an openstack userdata script (string) >`,
				},
				resource.Attribute{
					Name:        "boot2docker_url",
					Description: `(Optional) vSphere URL for boot2docker iso image. Default ` + "`" + `https://releases.rancher.com/os/latest/rancheros-vmware.iso` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "cfgparam",
					Description: `(Optional) vSphere vm configuration parameters (used for guestinfo) (list)`,
				},
				resource.Attribute{
					Name:        "cloudinit",
					Description: `(Optional) vSphere cloud-init file or url to set in the guestinfo (string)`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `(Optional) vSphere CPU number for docker VM. Default ` + "`" + `2` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) vSphere datacenter for docker VM (string)`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Optional) vSphere datastore for docker VM (string)`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) vSphere size of disk for docker VM (in MB). Default ` + "`" + `20480` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) vSphere folder for the docker VM. This folder must already exist in the datacenter (string)`,
				},
				resource.Attribute{
					Name:        "hostsystem",
					Description: `(Optional) vSphere compute resource where the docker VM will be instantiated. This can be omitted if using a cluster with DRS (string)`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Optional) vSphere size of memory for docker VM (in MB). Default ` + "`" + `2048` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) vSphere network where the docker VM will be attached (list)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) vSphere password. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Optional) vSphere resource pool for docker VM (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional/Sensitive) vSphere username. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "vapp_ip_allocation_policy",
					Description: `(Optional) vSphere vApp IP allocation policy. Supported values are: ` + "`" + `dhcp` + "`" + `, ` + "`" + `fixed` + "`" + `, ` + "`" + `transient` + "`" + ` and ` + "`" + `fixedAllocated` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "vapp_ip_protocol",
					Description: `(Optional) vSphere vApp IP protocol for this deployment. Supported values are: ` + "`" + `IPv4` + "`" + ` and ` + "`" + `IPv6` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "vapp_property",
					Description: `(Optional) vSphere vApp properties (list)`,
				},
				resource.Attribute{
					Name:        "vapp_transport",
					Description: `(Optional) vSphere OVF environment transports to use for properties. Supported values are: ` + "`" + `iso` + "`" + ` and ` + "`" + `com.vmware.guestInfo` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "vcenter",
					Description: `(Optional/Sensitive) vSphere IP/hostname for vCenter. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "vcenter_port",
					Description: `(Optional/Sensitive) vSphere Port for vCenter. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x. Default ` + "`" + `443` + "`" + ` (string) ## Timeouts ` + "`" + `rancher2_node_template` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating node templates. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for node template modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting node templates. ## Import Node Template can be imported using the Rancher Node Template ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_node_template.foo <node_template_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Computed) The driver of the node template (string) ## Nested blocks ### ` + "`" + `amazonec2_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ami",
					Description: `(Required) AWS machine image (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) AWS region. (string)`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Required) AWS VPC security group. (list)`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) AWS VPC subnet id (string)`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) AWS VPC id. (string)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) AWS zone for instance (i.e. a,b,c,d,e) (string)`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional/Sensitive) AWS access key. Required on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "block_duration_minutes",
					Description: `(Optional) AWS spot instance duration in minutes (60, 120, 180, 240, 300, or 360). Default ` + "`" + `0` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) AWS root device name. Default ` + "`" + `/dev/sda1` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) Optional endpoint URL (hostname only or fully qualified URI) (string)`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `(Optional) AWS IAM Instance Profile (string)`,
				},
				resource.Attribute{
					Name:        "insecure_transport",
					Description: `(Optional) Disable SSL when sending requests (bool)`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) AWS instance type. Default ` + "`" + `t2.micro` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `(Optional) AWS keypair to use; requires --amazonec2-ssh-keypath (string)`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional) Set this flag to enable CloudWatch monitoring. Deafult ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "open_port",
					Description: `(Optional) Make the specified port number accessible from the Internet. (list)`,
				},
				resource.Attribute{
					Name:        "private_address_only",
					Description: `(Optional) Only use a private IP address. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "request_spot_instance",
					Description: `(Optional) Set this flag to request spot instance. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Set retry count for recoverable failures (use -1 to disable). Default ` + "`" + `5` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "root_size",
					Description: `(Optional) AWS root disk size (in GB). Default ` + "`" + `16` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional/Sensitive) AWS secret key. Required on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "security_group_readonly",
					Description: `(Optional) Skip adding default rules to security groups (bool)`,
				},
				resource.Attribute{
					Name:        "session_token",
					Description: `(Optional/Sensitive) AWS Session Token (string)`,
				},
				resource.Attribute{
					Name:        "spot_price",
					Description: `(Optional) AWS spot instance bid price (in dollar). Default ` + "`" + `0.50` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_keypath",
					Description: `(Optional) SSH Key for Instance (string)`,
				},
				resource.Attribute{
					Name:        "ssh_user",
					Description: `(Optional) Set the name of the ssh user (string)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) AWS Tags (e.g. key1,value1,key2,value2) (string)`,
				},
				resource.Attribute{
					Name:        "use_ebs_optimized_instance",
					Description: `(Optional) Create an EBS optimized instance. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "use_private_address",
					Description: `(Optional) Force the usage of private IP address. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "userdata",
					Description: `(Optional) Path to file with cloud-init user data (string)`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) Amazon EBS volume type. Default ` + "`" + `gp2` + "`" + ` (string) ### ` + "`" + `azure_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional/Sensitive) Azure Service Principal Account ID. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional/Sensitive) Azure Service Principal Account password. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Optional/Sensitive) Azure Subscription ID. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "availability_set",
					Description: `(Optional) Azure Availability Set to place the virtual machine into. Default ` + "`" + `docker-machine` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `(Optional) Path to file with custom-data (string)`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) Disk size if using managed disk. Just for Rancher v2.3.x and above. Default ` + "`" + `30` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional) A unique DNS label for the public IP adddress (string)`,
				},
				resource.Attribute{
					Name:        "docker_port",
					Description: `(Optional) Port number for Docker engine. Default ` + "`" + `2376` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) Azure environment (e.g. AzurePublicCloud, AzureChinaCloud). Default ` + "`" + `AzurePublicCloud` + "`" + ` (string) ` + "`" + `fault_domain_count` + "`" + ` - (Optional) Fault domain count to use for availability set. Default ` + "`" + `3` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Azure virtual machine OS image. Default ` + "`" + `canonical:UbuntuServer:18.04-LTS:latest` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Azure region to create the virtual machine. Default ` + "`" + `westus` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "managed_disks",
					Description: `(Optional) Configures VM and availability set for managed disks. Just for Rancher v2.3.x and above. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "no_public_ip",
					Description: `(Optional) Do not create a public IP address for the machine. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "open_port",
					Description: `(Optional) Make the specified port number accessible from the Internet. (list)`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `(Optional) Specify a static private IP address for the machine. (string)`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Optional) Azure Resource Group name (will be created if missing). Default ` + "`" + `docker-machine` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size for Azure Virtual Machine. Default ` + "`" + `Standard_A2` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_user",
					Description: `(Optional) Username for SSH login (string)`,
				},
				resource.Attribute{
					Name:        "static_public_ip",
					Description: `(Optional) Assign a static public IP address to the machine. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) Type of Storage Account to host the OS Disk for the machine. Default ` + "`" + `Standard_LRS` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) Azure Subnet Name to be used within the Virtual Network. Default ` + "`" + `docker-machine` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "subnet_prefix",
					Description: `(Optional) Private CIDR block to be used for the new subnet, should comply RFC 1918. Default ` + "`" + `192.168.0.0/16` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "update_domain_count",
					Description: `(Optional) Update domain count to use for availability set. Default ` + "`" + `5` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "use_private_ip",
					Description: `(Optional) Use private IP address of the machine to connect. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "vnet",
					Description: `(Optional) Azure Virtual Network name to connect the virtual machine (in [resourcegroup:]name format). Default ` + "`" + `docker-machine-vnet` + "`" + ` (string) ### ` + "`" + `digitalocean_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `(Optional/Sensitive) Digital Ocean access token. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `(Optional) Enable backups for droplet. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Digital Ocean Image. Default ` + "`" + `ubuntu-16-04-x64` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) Enable ipv6 for droplet. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional) Enable monitoring for droplet. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "private_networking",
					Description: `(Optional) Enable private networking for droplet. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Digital Ocean region. Default ` + "`" + `nyc3` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Digital Ocean size. Default ` + "`" + `s-1vcpu-1gb` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_fingerprint",
					Description: `(Optional/Sensitive) SSH key fingerprint (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional) SSH private key path (string)`,
				},
				resource.Attribute{
					Name:        "ssh_port",
					Description: `(Optional) SSH port. Default ` + "`" + `22` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_user",
					Description: `(Optional) SSH username. Default ` + "`" + `root` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Comma-separated list of tags to apply to the Droplet (string)`,
				},
				resource.Attribute{
					Name:        "userdata",
					Description: `(Optional) Path to file with cloud-init user-data (string) ### ` + "`" + `openstack_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "auth_url",
					Description: `(Required) OpenStack authentication URL (string)`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) OpenStack availability zone (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) OpenStack region name (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) OpenStack username (string)`,
				},
				resource.Attribute{
					Name:        "cacert",
					Description: `(Optional) CA certificate bundle to verify against (string)`,
				},
				resource.Attribute{
					Name:        "config_drive",
					Description: `(Optional) Enables the OpenStack config drive for the instance. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "endpoint_type",
					Description: `(Optional) OpenStack endpoint type. adminURL, internalURL or publicURL (string)`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "floating_ip_pool",
					Description: `(Optional) OpenStack floating IP pool to get an IP from to assign to the instance (string)`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "insecure",
					Description: `(Optional) Disable TLS credential checking. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) OpenStack version of IP address assigned for the machine Default ` + "`" + `4` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `(Optional) OpenStack keypair to use to SSH to the instance (string)`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "net_name",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "nova_network",
					Description: `(Optional) Use the nova networking services instead of neutron (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) OpenStack password. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "private_key_file",
					Description: `(Optional) Private keyfile absolute path to use for SSH (string)`,
				},
				resource.Attribute{
					Name:        "sec_groups",
					Description: `(Optional) OpenStack comma separated security groups for the machine (string)`,
				},
				resource.Attribute{
					Name:        "ssh_port",
					Description: `(Optional) OpenStack SSH port`,
				},
				resource.Attribute{
					Name:        "ssh_user",
					Description: `(Optional) OpenStack SSH user`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `(Required`,
				},
				resource.Attribute{
					Name:        "user_data_file",
					Description: `(Optional) File containing an openstack userdata script (string) >`,
				},
				resource.Attribute{
					Name:        "boot2docker_url",
					Description: `(Optional) vSphere URL for boot2docker iso image. Default ` + "`" + `https://releases.rancher.com/os/latest/rancheros-vmware.iso` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "cfgparam",
					Description: `(Optional) vSphere vm configuration parameters (used for guestinfo) (list)`,
				},
				resource.Attribute{
					Name:        "cloudinit",
					Description: `(Optional) vSphere cloud-init file or url to set in the guestinfo (string)`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `(Optional) vSphere CPU number for docker VM. Default ` + "`" + `2` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) vSphere datacenter for docker VM (string)`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Optional) vSphere datastore for docker VM (string)`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) vSphere size of disk for docker VM (in MB). Default ` + "`" + `20480` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) vSphere folder for the docker VM. This folder must already exist in the datacenter (string)`,
				},
				resource.Attribute{
					Name:        "hostsystem",
					Description: `(Optional) vSphere compute resource where the docker VM will be instantiated. This can be omitted if using a cluster with DRS (string)`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Optional) vSphere size of memory for docker VM (in MB). Default ` + "`" + `2048` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) vSphere network where the docker VM will be attached (list)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) vSphere password. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Optional) vSphere resource pool for docker VM (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional/Sensitive) vSphere username. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "vapp_ip_allocation_policy",
					Description: `(Optional) vSphere vApp IP allocation policy. Supported values are: ` + "`" + `dhcp` + "`" + `, ` + "`" + `fixed` + "`" + `, ` + "`" + `transient` + "`" + ` and ` + "`" + `fixedAllocated` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "vapp_ip_protocol",
					Description: `(Optional) vSphere vApp IP protocol for this deployment. Supported values are: ` + "`" + `IPv4` + "`" + ` and ` + "`" + `IPv6` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "vapp_property",
					Description: `(Optional) vSphere vApp properties (list)`,
				},
				resource.Attribute{
					Name:        "vapp_transport",
					Description: `(Optional) vSphere OVF environment transports to use for properties. Supported values are: ` + "`" + `iso` + "`" + ` and ` + "`" + `com.vmware.guestInfo` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "vcenter",
					Description: `(Optional/Sensitive) vSphere IP/hostname for vCenter. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x (string)`,
				},
				resource.Attribute{
					Name:        "vcenter_port",
					Description: `(Optional/Sensitive) vSphere Port for vCenter. Mandatory on Rancher v2.0.x and v2.1.x. Use ` + "`" + `rancher2_cloud_credential` + "`" + ` from Rancher v2.2.x. Default ` + "`" + `443` + "`" + ` (string) ## Timeouts ` + "`" + `rancher2_node_template` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating node templates. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for node template modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting node templates. ## Import Node Template can be imported using the Rancher Node Template ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_node_template.foo <node_template_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_notifier",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Notifier resource. This can be used to create notifiers for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"notifier",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the notifier (string)`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required/ForceNew) The cluster id where create notifier (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The notifier description (string)`,
				},
				resource.Attribute{
					Name:        "pagerduty_config",
					Description: `(Optional) Pagerduty config for notifier (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "slack_config",
					Description: `(Optional) Slack config for notifier (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "smtp_config",
					Description: `(Optional) SMTP config for notifier (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "webhook_config",
					Description: `(Optional) Webhook config for notifier (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "wechat_config",
					Description: `(Optional) Wechat config for notifier (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for notifier object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for notifier object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `pagerduty_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `(Required) Pagerduty service key (string)`,
				},
				resource.Attribute{
					Name:        "proxy_url",
					Description: `(Optional) Pagerduty proxy url (string) ### ` + "`" + `slack_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_recipient",
					Description: `(Required) Slack default recipient (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Slack url (string)`,
				},
				resource.Attribute{
					Name:        "proxy_url",
					Description: `(Optional) Slack proxy url (string) ### ` + "`" + `smtp_config` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_recipient",
					Description: `(Required) SMTP default recipient (string)`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) SMTP host (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) SMTP port (int)`,
				},
				resource.Attribute{
					Name:        "sender",
					Description: `(Required) SMTP sender (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) SMTP password (string)`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional/Sensitive) SMTP tls. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional/Sensitive) SMTP username (string) #### Arguments ### ` + "`" + `webhook_config` + "`" + ``,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Webhook url (string)`,
				},
				resource.Attribute{
					Name:        "proxy_url",
					Description: `(Optional) Webhook proxy url (string) #### Arguments ### ` + "`" + `wechat_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "agent",
					Description: `(Required) Wechat agent ID (string)`,
				},
				resource.Attribute{
					Name:        "corp",
					Description: `(Required) Wechat corporation ID (string)`,
				},
				resource.Attribute{
					Name:        "default_recipient",
					Description: `(Required) Wechat default recipient (string)`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Required/Sensitive) Wechat agent ID (string)`,
				},
				resource.Attribute{
					Name:        "proxy_url",
					Description: `(Optional) Wechat proxy url (string)`,
				},
				resource.Attribute{
					Name:        "recipient_type",
					Description: `(Optional) Wechat recipient type. Allowed values: ` + "`" + `party` + "`" + ` | ` + "`" + `tag` + "`" + ` | ` + "`" + `user` + "`" + ` (string) ## Timeouts ` + "`" + `rancher2_notifier` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating notifiers. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for notifier modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting notifiers. ## Import Notifiers can be imported using the Rancher nNtifier ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_notifier.foo <notifier_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `pagerduty_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `(Required) Pagerduty service key (string)`,
				},
				resource.Attribute{
					Name:        "proxy_url",
					Description: `(Optional) Pagerduty proxy url (string) ### ` + "`" + `slack_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_recipient",
					Description: `(Required) Slack default recipient (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Slack url (string)`,
				},
				resource.Attribute{
					Name:        "proxy_url",
					Description: `(Optional) Slack proxy url (string) ### ` + "`" + `smtp_config` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_recipient",
					Description: `(Required) SMTP default recipient (string)`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) SMTP host (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) SMTP port (int)`,
				},
				resource.Attribute{
					Name:        "sender",
					Description: `(Required) SMTP sender (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) SMTP password (string)`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional/Sensitive) SMTP tls. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional/Sensitive) SMTP username (string) #### Arguments ### ` + "`" + `webhook_config` + "`" + ``,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Webhook url (string)`,
				},
				resource.Attribute{
					Name:        "proxy_url",
					Description: `(Optional) Webhook proxy url (string) #### Arguments ### ` + "`" + `wechat_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "agent",
					Description: `(Required) Wechat agent ID (string)`,
				},
				resource.Attribute{
					Name:        "corp",
					Description: `(Required) Wechat corporation ID (string)`,
				},
				resource.Attribute{
					Name:        "default_recipient",
					Description: `(Required) Wechat default recipient (string)`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Required/Sensitive) Wechat agent ID (string)`,
				},
				resource.Attribute{
					Name:        "proxy_url",
					Description: `(Optional) Wechat proxy url (string)`,
				},
				resource.Attribute{
					Name:        "recipient_type",
					Description: `(Optional) Wechat recipient type. Allowed values: ` + "`" + `party` + "`" + ` | ` + "`" + `tag` + "`" + ` | ` + "`" + `user` + "`" + ` (string) ## Timeouts ` + "`" + `rancher2_notifier` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating notifiers. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for notifier modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting notifiers. ## Import Notifiers can be imported using the Rancher nNtifier ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_notifier.foo <notifier_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_project",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Project resource. This can be used to create projects for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project (string)`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The cluster id where create project (string)`,
				},
				resource.Attribute{
					Name:        "container_resource_limit",
					Description: `(Optional) Default containers resource limits on project (List maxitem:1)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A project description (string)`,
				},
				resource.Attribute{
					Name:        "enable_project_monitoring",
					Description: `(Optional) Enable built-in project monitoring. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "pod_security_policy_template_id",
					Description: `(Optional) Default Pod Security Policy ID for the project (string)`,
				},
				resource.Attribute{
					Name:        "project_monitoring_input",
					Description: `(Optional/Computed) Project monitoring config. Any parameter defined in [rancher-monitoring charts](https://github.com/rancher/system-charts/tree/dev/charts/rancher-monitoring) could be configured (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "resource_quota",
					Description: `(Optional) Resource quota for project. Rancher v2.1.x or higher (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "wait_for_cluster",
					Description: `(Optional) Wait for cluster becomes active. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for Node Pool object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for Node Pool object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `container_resource_limit` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "limits_cpu",
					Description: `(Optional) CPU limit for containers (string)`,
				},
				resource.Attribute{
					Name:        "limits_memory",
					Description: `(Optional) Memory limit for containers (string)`,
				},
				resource.Attribute{
					Name:        "requests_cpu",
					Description: `(Optional) CPU reservation for containers (string)`,
				},
				resource.Attribute{
					Name:        "requests_memory",
					Description: `(Optional) Memory reservation for containers (string) ### ` + "`" + `project_monitoring_input` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "answers",
					Description: `(Optional/Computed) Key/value answers for monitor input (map) ### ` + "`" + `resource_quota` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "project_limit",
					Description: `(Required) Resource quota limit for project (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "namespace_default_limit",
					Description: `(Required) Default resource quota limit for namespaces in project (list maxitems:1) #### ` + "`" + `project_limit` + "`" + ` and ` + "`" + `namespace_default_limit` + "`" + ` ##### Arguments The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "config_maps",
					Description: `(Optional) Limit for config maps in project (string)`,
				},
				resource.Attribute{
					Name:        "limits_cpu",
					Description: `(Optional) Limit for limits cpu in project (string)`,
				},
				resource.Attribute{
					Name:        "limits_memory",
					Description: `(Optional) Limit for limits memory in project (string)`,
				},
				resource.Attribute{
					Name:        "persistent_volume_claims",
					Description: `(Optional) Limit for persistent volume claims in project (string)`,
				},
				resource.Attribute{
					Name:        "pods",
					Description: `(Optional) Limit for pods in project (string)`,
				},
				resource.Attribute{
					Name:        "replication_controllers",
					Description: `(Optional) Limit for replication controllers in project (string)`,
				},
				resource.Attribute{
					Name:        "requests_cpu",
					Description: `(Optional) Limit for requests cpu in project (string)`,
				},
				resource.Attribute{
					Name:        "requests_memory",
					Description: `(Optional) Limit for requests memory in project (string)`,
				},
				resource.Attribute{
					Name:        "requests_storage",
					Description: `(Optional) Limit for requests storage in project (string)`,
				},
				resource.Attribute{
					Name:        "secrets",
					Description: `(Optional) Limit for secrets in project (string)`,
				},
				resource.Attribute{
					Name:        "services_load_balancers",
					Description: `(Optional) Limit for services load balancers in project (string)`,
				},
				resource.Attribute{
					Name:        "services_node_ports",
					Description: `(Optional) Limit for services node ports in project (string) More info at [resource-quotas](https://rancher.com/docs/rancher/v2.x/en/k8s-in-rancher/projects-and-namespaces/resource-quotas/) ## Timeouts ` + "`" + `rancher2_project` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating projects. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for project modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting projects. ## Import Projects can be imported using the Rancher Project ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_project.foo <project_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `container_resource_limit` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "limits_cpu",
					Description: `(Optional) CPU limit for containers (string)`,
				},
				resource.Attribute{
					Name:        "limits_memory",
					Description: `(Optional) Memory limit for containers (string)`,
				},
				resource.Attribute{
					Name:        "requests_cpu",
					Description: `(Optional) CPU reservation for containers (string)`,
				},
				resource.Attribute{
					Name:        "requests_memory",
					Description: `(Optional) Memory reservation for containers (string) ### ` + "`" + `project_monitoring_input` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "answers",
					Description: `(Optional/Computed) Key/value answers for monitor input (map) ### ` + "`" + `resource_quota` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "project_limit",
					Description: `(Required) Resource quota limit for project (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "namespace_default_limit",
					Description: `(Required) Default resource quota limit for namespaces in project (list maxitems:1) #### ` + "`" + `project_limit` + "`" + ` and ` + "`" + `namespace_default_limit` + "`" + ` ##### Arguments The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "config_maps",
					Description: `(Optional) Limit for config maps in project (string)`,
				},
				resource.Attribute{
					Name:        "limits_cpu",
					Description: `(Optional) Limit for limits cpu in project (string)`,
				},
				resource.Attribute{
					Name:        "limits_memory",
					Description: `(Optional) Limit for limits memory in project (string)`,
				},
				resource.Attribute{
					Name:        "persistent_volume_claims",
					Description: `(Optional) Limit for persistent volume claims in project (string)`,
				},
				resource.Attribute{
					Name:        "pods",
					Description: `(Optional) Limit for pods in project (string)`,
				},
				resource.Attribute{
					Name:        "replication_controllers",
					Description: `(Optional) Limit for replication controllers in project (string)`,
				},
				resource.Attribute{
					Name:        "requests_cpu",
					Description: `(Optional) Limit for requests cpu in project (string)`,
				},
				resource.Attribute{
					Name:        "requests_memory",
					Description: `(Optional) Limit for requests memory in project (string)`,
				},
				resource.Attribute{
					Name:        "requests_storage",
					Description: `(Optional) Limit for requests storage in project (string)`,
				},
				resource.Attribute{
					Name:        "secrets",
					Description: `(Optional) Limit for secrets in project (string)`,
				},
				resource.Attribute{
					Name:        "services_load_balancers",
					Description: `(Optional) Limit for services load balancers in project (string)`,
				},
				resource.Attribute{
					Name:        "services_node_ports",
					Description: `(Optional) Limit for services node ports in project (string) More info at [resource-quotas](https://rancher.com/docs/rancher/v2.x/en/k8s-in-rancher/projects-and-namespaces/resource-quotas/) ## Timeouts ` + "`" + `rancher2_project` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating projects. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for project modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting projects. ## Import Projects can be imported using the Rancher Project ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_project.foo <project_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_project_logging",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Project Logging resource. This can be used to create Project Logging for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"project",
				"logging",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project id to configure logging (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Project Logging config (string)`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind of the Project Logging. ` + "`" + `elasticsearch` + "`" + `, ` + "`" + `fluentd` + "`" + `, ` + "`" + `kafka` + "`" + `, ` + "`" + `splunk` + "`" + ` and ` + "`" + `syslog` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "elasticsearch_config",
					Description: `(Optional) The elasticsearch config for Project Logging. For ` + "`" + `kind = elasticsearch` + "`" + `. Conflicts with ` + "`" + `fluentd_config` + "`" + `, ` + "`" + `kafka_config` + "`" + `, ` + "`" + `splunk_config` + "`" + ` and ` + "`" + `syslog_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "fluentd_config",
					Description: `(Optional) The fluentd config for Project Logging. For ` + "`" + `kind = fluentd` + "`" + `. Conflicts with ` + "`" + `elasticsearch_config` + "`" + `, ` + "`" + `kafka_config` + "`" + `, ` + "`" + `splunk_config` + "`" + ` and ` + "`" + `syslog_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kafka_config",
					Description: `(Optional) The kafka config for Project Logging. For ` + "`" + `kind = kafka` + "`" + `. Conflicts with ` + "`" + `elasticsearch_config` + "`" + `, ` + "`" + `fluentd_config` + "`" + `, ` + "`" + `splunk_config` + "`" + ` and ` + "`" + `syslog_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Optional) The namespace id from Project logging (string)`,
				},
				resource.Attribute{
					Name:        "output_flush_interval",
					Description: `(Optional) How often buffered logs would be flushed. Default: ` + "`" + `3` + "`" + ` seconds (int)`,
				},
				resource.Attribute{
					Name:        "output_tags",
					Description: `(Optional/computed) The output tags for Project Logging (map)`,
				},
				resource.Attribute{
					Name:        "splunk_config",
					Description: `(Optional) The splunk config for Project Logging. For ` + "`" + `kind = splunk` + "`" + `. Conflicts with ` + "`" + `elasticsearch_config` + "`" + `, ` + "`" + `fluentd_config` + "`" + `, ` + "`" + `kafka_config` + "`" + `, and ` + "`" + `syslog_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "syslog_config",
					Description: `(Optional) The syslog config for Project Logging. For ` + "`" + `kind = syslog` + "`" + `. Conflicts with ` + "`" + `elasticsearch_config` + "`" + `, ` + "`" + `fluentd_config` + "`" + `, ` + "`" + `kafka_config` + "`" + `, and ` + "`" + `splunk_config` + "`" + ` (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for Project Logging object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for Project Logging object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `elasticsearch_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the elascticsearch service. Must include protocol, ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "auth_password",
					Description: `(Optional/Sensitive) User password for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "auth_username",
					Description: `(Optional/Sensitive) Username for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_key_pass",
					Description: `(Optional/Sensitive) SSL client key password for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "date_format",
					Description: `(Optional) Date format for the elascticsearch logs. Default: ` + "`" + `YYYY-MM-DD` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "index_prefix",
					Description: `(Optional) Index prefix for the elascticsearch logs. Default: ` + "`" + `local` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the elascticsearch service (bool)`,
				},
				resource.Attribute{
					Name:        "ssl_version",
					Description: `(Optional) SSL version for the elascticsearch service (string) ### ` + "`" + `fluentd_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fluent_servers",
					Description: `(Required) Servers for the fluentd service (list)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "compress",
					Description: `(Optional) Compress data for the fluentd service (bool)`,
				},
				resource.Attribute{
					Name:        "enable_tls",
					Description: `(Optional) Enable TLS for the fluentd service (bool) #### ` + "`" + `fluent_servers` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Hostname of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) User password of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "shared_key",
					Description: `(Optional/Sensitive) Shared key of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "standby",
					Description: `(Optional) Standby server of the fluentd service (bool)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional/Sensitive) Username of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight of the fluentd server (int) ### ` + "`" + `kafka_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Required) Topic to publish on the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "broker_endpoints",
					Description: `(Optional) Kafka endpoints for kafka service. Conflicts with ` + "`" + `zookeeper_endpoint` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "zookeeper_endpoint",
					Description: `(Optional) Zookeeper endpoint for kafka service. Conflicts with ` + "`" + `broker_endpoints` + "`" + ` (string) ### ` + "`" + `splunk_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the splunk service. Must include protocol, ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required/Sensitive) Token for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_key_pass",
					Description: `(Optional/Sensitive) SSL client key password for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) Index prefix for the splunk logs (string)`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Date format for the splunk logs (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the splunk service (bool) ### ` + "`" + `syslog_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "program",
					Description: `(Optional) Program for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol for the syslog service. ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + ` are supported. Default: ` + "`" + `udp` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Optional) Date format for the syslog logs. ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notice` + "`" + `, ` + "`" + `info` + "`" + ` and ` + "`" + `debug` + "`" + ` are supported. Default: ` + "`" + `notice` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the syslog service (bool)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional/Sensitive) Token for the syslog service (string) ## Timeouts ` + "`" + `rancher2_project_logging` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating project logging configurations. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for project logging configuration modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting project logging configurations. ## Import Project Logging can be imported using the Rancher Project Logging ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_project_logging.foo <project_logging_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `elasticsearch_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the elascticsearch service. Must include protocol, ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "auth_password",
					Description: `(Optional/Sensitive) User password for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "auth_username",
					Description: `(Optional/Sensitive) Username for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "client_key_pass",
					Description: `(Optional/Sensitive) SSL client key password for the elascticsearch service (string)`,
				},
				resource.Attribute{
					Name:        "date_format",
					Description: `(Optional) Date format for the elascticsearch logs. Default: ` + "`" + `YYYY-MM-DD` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "index_prefix",
					Description: `(Optional) Index prefix for the elascticsearch logs. Default: ` + "`" + `local` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the elascticsearch service (bool)`,
				},
				resource.Attribute{
					Name:        "ssl_version",
					Description: `(Optional) SSL version for the elascticsearch service (string) ### ` + "`" + `fluentd_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fluent_servers",
					Description: `(Required) Servers for the fluentd service (list)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "compress",
					Description: `(Optional) Compress data for the fluentd service (bool)`,
				},
				resource.Attribute{
					Name:        "enable_tls",
					Description: `(Optional) Enable TLS for the fluentd service (bool) #### ` + "`" + `fluent_servers` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Hostname of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) User password of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "shared_key",
					Description: `(Optional/Sensitive) Shared key of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "standby",
					Description: `(Optional) Standby server of the fluentd service (bool)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional/Sensitive) Username of the fluentd service (string)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight of the fluentd server (int) ### ` + "`" + `kafka_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Required) Topic to publish on the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "broker_endpoints",
					Description: `(Optional) Kafka endpoints for kafka service. Conflicts with ` + "`" + `zookeeper_endpoint` + "`" + ` (list)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the kafka service (string)`,
				},
				resource.Attribute{
					Name:        "zookeeper_endpoint",
					Description: `(Optional) Zookeeper endpoint for kafka service. Conflicts with ` + "`" + `broker_endpoints` + "`" + ` (string) ### ` + "`" + `splunk_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the splunk service. Must include protocol, ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required/Sensitive) Token for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "client_key_pass",
					Description: `(Optional/Sensitive) SSL client key password for the splunk service (string)`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) Index prefix for the splunk logs (string)`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Date format for the splunk logs (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the splunk service (bool) ### ` + "`" + `syslog_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Endpoint of the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional/Sensitive) SSL certificate for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional/Sensitive) SSL client certificate for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional/Sensitive) SSL client key for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "program",
					Description: `(Optional) Program for the syslog service (string)`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol for the syslog service. ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + ` are supported. Default: ` + "`" + `udp` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Optional) Date format for the syslog logs. ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notice` + "`" + `, ` + "`" + `info` + "`" + ` and ` + "`" + `debug` + "`" + ` are supported. Default: ` + "`" + `notice` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssl_verify",
					Description: `(Optional) SSL verify for the syslog service (bool)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional/Sensitive) Token for the syslog service (string) ## Timeouts ` + "`" + `rancher2_project_logging` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating project logging configurations. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for project logging configuration modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting project logging configurations. ## Import Project Logging can be imported using the Rancher Project Logging ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_project_logging.foo <project_logging_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_project_role_template_binding",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Project Role Template Binding resource. This can be used to create Project Role Template Bindings for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"project",
				"role",
				"template",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required/ForceNew) The project id where bind project role template (string)`,
				},
				resource.Attribute{
					Name:        "role_template_id",
					Description: `(Required/ForceNew) The role template id from create project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) The group ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "group_principal_id",
					Description: `(Optional/Computed) The group_principal ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional) The user ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "user_principal_id",
					Description: `(Optional/Computed) The user_principal ID to assign project role template binding (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the resource (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the resource (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_project_role_template_binding` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating project role template bindings. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for project role template binding modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting project role template bindings. ## Import Project Role Template Bindings can be imported using the Rancher Project Role Template Binding ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_project_role_template_binding.foo <project_role_template_binding_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_project_role_template_binding` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating project role template bindings. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for project role template binding modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting project role template bindings. ## Import Project Role Template Bindings can be imported using the Rancher Project Role Template Binding ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_project_role_template_binding.foo <project_role_template_binding_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_registry",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 docker registry resource. This can be used to create docker registries for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"registry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required/ForceNew) The name of the registry (string)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required/ForceNew) The project id where to assign the registry (string)`,
				},
				resource.Attribute{
					Name:        "registries",
					Description: `(Required) Registries data for registry (list)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A registry description (string)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Optional) The namespace id where to assign the namespaced registry (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for Registry object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for Registry object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `registries` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Address for registry.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for the registry (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Username for the registry (string) ## Timeouts ` + "`" + `rancher2_registry` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating registries. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for registry modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting registries. ## Import Registries can be imported using the registry ID in the format ` + "`" + `<namespace_id>.<project_id>.<registry_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_registry.foo <namespace_id>.<project_id>.<registry_id> ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `<namespace_id>` + "`" + ` is optional, just needed for namespaced registry.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Nested blocks ### ` + "`" + `registries` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Address for registry.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for the registry (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Username for the registry (string) ## Timeouts ` + "`" + `rancher2_registry` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating registries. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for registry modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting registries. ## Import Registries can be imported using the registry ID in the format ` + "`" + `<namespace_id>.<project_id>.<registry_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_registry.foo <namespace_id>.<project_id>.<registry_id> ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `<namespace_id>` + "`" + ` is optional, just needed for namespaced registry.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_role_template",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Role Template resource. This can be used to create Role template for Rancher v2 RKE clusters and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"role",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Role template name (string)`,
				},
				resource.Attribute{
					Name:        "administrative",
					Description: `(Optional) Administrative role template. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Optional) Role template context. ` + "`" + `cluster` + "`" + ` and ` + "`" + `project` + "`" + ` values are supported. Default: ` + "`" + `cluster` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "default_role",
					Description: `(Optional) Default role template for new created cluster or project. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional/Computed) Role template description (string)`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `(Optional) External role template. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "hidden",
					Description: `(Optional) Hidden role template. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `(Optional) Locked role template. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "role_template_ids",
					Description: `(Optional/Computed) Inherit role template IDs (list)`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional/Computed) Role template policy rules (list)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for role template object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for role template object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "builtin",
					Description: `(Computed) Builtin role template (string) ## Nested blocks ### ` + "`" + `rules` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "api_groups",
					Description: `(Optional) Policy rule api groups (list)`,
				},
				resource.Attribute{
					Name:        "non_resource_urls",
					Description: `(Optional) Policy rule non resource urls (list)`,
				},
				resource.Attribute{
					Name:        "resource_names",
					Description: `(Optional) Policy rule resource names (list)`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) Policy rule resources (list)`,
				},
				resource.Attribute{
					Name:        "verbs",
					Description: `(Optional) Policy rule verbs. ` + "`" + `create` + "`" + `, ` + "`" + `delete` + "`" + `, ` + "`" + `get` + "`" + `, ` + "`" + `list` + "`" + `, ` + "`" + `patch` + "`" + `, ` + "`" + `update` + "`" + `, ` + "`" + `watch` + "`" + ` and ` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "builtin",
					Description: `(Computed) Builtin role template (string) ## Nested blocks ### ` + "`" + `rules` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "api_groups",
					Description: `(Optional) Policy rule api groups (list)`,
				},
				resource.Attribute{
					Name:        "non_resource_urls",
					Description: `(Optional) Policy rule non resource urls (list)`,
				},
				resource.Attribute{
					Name:        "resource_names",
					Description: `(Optional) Policy rule resource names (list)`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) Policy rule resources (list)`,
				},
				resource.Attribute{
					Name:        "verbs",
					Description: `(Optional) Policy rule verbs. ` + "`" + `create` + "`" + `, ` + "`" + `delete` + "`" + `, ` + "`" + `get` + "`" + `, ` + "`" + `list` + "`" + `, ` + "`" + `patch` + "`" + `, ` + "`" + `update` + "`" + `, ` + "`" + `watch` + "`" + ` and ` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_secret",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 secret resource. This can be used to create secrets for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "data",
					Description: `(Required) Secret key/value data. Base64 encoding required for values (map)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required/ForceNew) The project id where to assign the secret (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A secret description (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional/ForceNew) The name of the secret (string)`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Optional/ForceNew) The namespace id where to assign the namespaced secret (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for secret object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for secret object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_secret` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating registries. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for secret modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting registries. ## Import Secrets can be imported using the secret ID in the format ` + "`" + `<namespace_id>.<project_id>.<secret_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_secret.foo <namespace_id>.<project_id>.<secret_id> ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `<namespace_id>` + "`" + ` is optional, just needed for namespaced secret.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Timeouts ` + "`" + `rancher2_secret` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating registries. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for secret modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting registries. ## Import Secrets can be imported using the secret ID in the format ` + "`" + `<namespace_id>.<project_id>.<secret_id>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_secret.foo <namespace_id>.<project_id>.<secret_id> ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `<namespace_id>` + "`" + ` is optional, just needed for namespaced secret.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_setting",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Setting resource. This can be used to create settings for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the setting (string)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the setting (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for setting object (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for setting object (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Import Setting can be imported using the Rancher setting ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_setting.foo <setting_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string) ## Import Setting can be imported using the Rancher setting ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_setting.foo <setting_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_token",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 Token resource. This can be used to create tokens for Rancher v2 provider user and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional/ForceNew) Cluster ID for scoped token (string)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional/ForceNew) Token description (string)`,
				},
				resource.Attribute{
					Name:        "renew",
					Description: `(Optional/ForceNew) Renew token if expired or disabled. If ` + "`" + `true` + "`" + `, a terraform diff would be generated to renew the token if it's disabled or expired. If ` + "`" + `false` + "`" + `, the token will not be renewed. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional/ForceNew) Token time to live in seconds. Default ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations of the token (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels of the token (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Computed) Token access key part (string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Computed) Token is enabled (bool)`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `(Computed) Token is expired (bool)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) Token name (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Computed/Sensitive) Token secret key part (string)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Computed/Sensitive) Token value (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Computed) Token user ID (string) ## Timeouts ` + "`" + `rancher2_token` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for creating tokens. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for token modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for deleting tokens.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Computed) Token access key part (string)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Computed) Token is enabled (bool)`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `(Computed) Token is expired (bool)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) Token name (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Computed/Sensitive) Token secret key part (string)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Computed/Sensitive) Token value (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Computed) Token user ID (string) ## Timeouts ` + "`" + `rancher2_token` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for creating tokens. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for token modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for deleting tokens.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_user",
			Category:         "Resources",
			ShortDescription: `Provides a Rancher v2 User resource. This can be used to create Users for Rancher v2 environments and retrieve their information.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required/ForceNew) The user username (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/ForceNew) The user password (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The user full name (string)`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional/Computed) Annotations for global role binding (map)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional/Computed) Labels for global role binding (map) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "principal_ids",
					Description: `(Computed) The user principal IDs (list) ## Timeouts ` + "`" + `rancher2_user` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for creating users. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for user modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for deleting users. ## Import Users can be imported using the Rancher User ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_user.foo <user_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "principal_ids",
					Description: `(Computed) The user principal IDs (list) ## Timeouts ` + "`" + `rancher2_user` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for creating users. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for user modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for deleting users. ## Import Users can be imported using the Rancher User ID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rancher2_user.foo <user_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"rancher2_app":                           0,
		"rancher2_auth_config_adfs":              1,
		"rancher2_auth_config_activedirectory":   2,
		"rancher2_auth_config_azuread":           3,
		"rancher2_auth_config_freeipa":           4,
		"rancher2_auth_config_github":            5,
		"rancher2_auth_config_keycloak":          6,
		"rancher2_auth_config_okta":              7,
		"rancher2_auth_config_openldap":          8,
		"rancher2_auth_config_ping":              9,
		"rancher2_bootstrap":                     10,
		"rancher2_catalog":                       11,
		"rancher2_certificate":                   12,
		"rancher2_cloud_credential":              13,
		"rancher2_cluster":                       14,
		"rancher2_cluster_alert_group":           15,
		"rancher2_cluster_alert_rule":            16,
		"rancher2_cluster_driver":                17,
		"rancher2_cluster_logging":               18,
		"rancher2_cluster_role_template_binding": 19,
		"rancher2_cluster_template":              20,
		"rancher2_etcd_backup":                   21,
		"rancher2_global_role_binding":           22,
		"rancher2_multi_cluster_app":             23,
		"rancher2_namespace":                     24,
		"rancher2_node_driver":                   25,
		"rancher2_node_pool":                     26,
		"rancher2_node_template":                 27,
		"rancher2_notifier":                      28,
		"rancher2_project":                       29,
		"rancher2_project_logging":               30,
		"rancher2_project_role_template_binding": 31,
		"rancher2_registry":                      32,
		"rancher2_role_template":                 33,
		"rancher2_secret":                        34,
		"rancher2_setting":                       35,
		"rancher2_token":                         36,
		"rancher2_user":                          37,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
